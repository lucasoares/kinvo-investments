// Copyright 2021 Lucas Soares
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/jessevdk/go-flags"
	"github.com/lucasoares/kinvo/kinvo"
	"os"
	"sort"
	"strings"
)

type Opts struct {
	File       string `short:"f" long:"file" description:"File to process" required:"true" value-name:"FILE"`
	Broker     string `short:"b" long:"broker" description:"Broker to filter products" required:"true"`
	AssetClass string `short:"c" long:"class" description:"Assert classes to use" required:"true"`
}

func main() {
	var opts Opts

	_, err := flags.ParseArgs(&opts, os.Args[1:])

	if err != nil {
		panic(err)
	}

	file, err := excelize.OpenFile(opts.File)
	if err != nil {
		fmt.Println(err.Error())

		return
	}

	products := kinvo.ParseProducts(file)
	recommendations := kinvo.ParseRecommendation(file)

	filteredProducts, max, total := generateFilteredProducts(opts, products, recommendations)

	if total == 0 {
		fmt.Println("You have 0 invested. You should follow recommendations.")
		return
	}

	sort.SliceStable(filteredProducts, func(i, j int) bool {
		return filteredProducts[i].Balance < filteredProducts[j].Balance
	})

	for i := range filteredProducts {
		product := filteredProducts[i]

		product.CurrentPercentage = (product.Balance * 100) / total

		fmt.Println(fmt.Sprintf("You have %.2f (%.3f%%) of %s. You should invest %.2f.", product.Balance, product.CurrentPercentage, product.Name, max-product.Balance))
	}
}

func generateFilteredProducts(
	opts Opts,
	products []*kinvo.Product,
	recommendations []*kinvo.Recommendation,
) (filteredProducts []*kinvo.Product, maxInvestment float64, totalInvestment float64) {
	result := make([]*kinvo.Product, 0)

	max := float64(0)
	total := float64(0)

	for i := range products {
		product := products[i]

		if !strings.Contains(strings.ToLower(product.Broker), strings.ToLower(opts.Broker)) {
			continue
		}

		if !strings.Contains(strings.ToLower(product.AssetClass), strings.ToLower(opts.AssetClass)) {
			continue
		}

		if product.Balance == 0 {
			continue
		}

		if len(recommendations) > 0 {
			isOnRecommendation := false
			for r := range recommendations {
				recommendation := recommendations[r]

				if kinvo.IsSameProduct(recommendation.Name, product.Name) {
					isOnRecommendation = true
				}
			}

			if !isOnRecommendation {
				fmt.Println(fmt.Sprintf("%s is not on recommendations. Current balance for this product: %.2f.", product.Name, product.Balance))

				continue
			}
		}

		if max < product.Balance {
			max = product.Balance
		}

		total += product.Balance

		result = append(result, product)
	}

	fmt.Println()

	fmt.Println(fmt.Sprintf(
		"You have %d products from %d recommended. Biggest investment: %.2f. Total from filtered products: %.2f.",
		len(result),
		len(recommendations),
		max,
		total,
	))

	fmt.Println()

	for r := range recommendations {
		recommendation := recommendations[r]

		containsProduct := false
		for i := range result {
			if kinvo.IsSameProduct(recommendation.Name, result[i].Name) {
				containsProduct = true
			}
		}

		if !containsProduct {
			result = append(result, &kinvo.Product{
				Name: recommendation.Name,
			})
		}
	}

	return result, max, total
}
