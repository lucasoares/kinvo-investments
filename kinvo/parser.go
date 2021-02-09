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

package kinvo

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"strconv"
	"strings"
	"time"
)

func ParseProducts(file *excelize.File) []*Product {
	rows, err := getRows(file, 0)
	if err != nil {
		fmt.Println("Error loading product sheet.")

		return nil
	}

	result := make([]*Product, 0)

	// Skip header
	rows.Next()
	_, _ = rows.Columns()

	for rows.Next() {
		row, _ := rows.Columns()

		product := Product{}
		for i := range row {
			value := row[i]

			switch i {
			case 0:
				product.Name = value
			case 1:
				product.AssetClass = value
			case 2:
				product.Broker = value
			case 3:
				product.FirstApplication, _ = time.Parse("02/01/2006", value)
			case 4:
				product.Investment = getNumber(value)
			case 5:
				product.Balance = getNumber(value)
			case 6:
				product.ProfitPercentage = getNumber(value)
			case 7:
				product.PortfolioPercentage = getNumber(value)
			}
		}

		result = append(result, &product)
	}

	return result
}

func ParseRecommendation(file *excelize.File) []*Recommendation {
	rows, err := getRows(file, 1)
	if err != nil {
		fmt.Println("Recommendations not found.")

		return nil
	}

	result := make([]*Recommendation, 0)

	for rows.Next() {
		row, _ := rows.Columns()

		if len(row) == 0 {
			break
		}

		if isNumber(row[0]) {
			continue
		}

		recommendation := Recommendation{
			Name: row[0],
		}

		result = append(result, &recommendation)
	}

	return result
}

func getRows(file *excelize.File, index int) (*excelize.Rows, error) {
	rows, err := file.Rows(file.GetSheetName(index))

	if err != nil || rows == nil {
		return nil, err
	}

	return rows, nil
}

func isNumber(value string) bool {
	_, err := formatNumber(value)

	return err == nil
}

func getNumber(value string) float64 {
	result, err := formatNumber(value)

	if err != nil {
		fmt.Println("Error parsing value.", err.Error())
	}

	return result
}

func formatNumber(value string) (float64, error) {
	if strings.Contains(value, ",") {
		value = strings.Replace(value, ".", "", -1)
		value = strings.Replace(value, ",", ".", 1)
	}

	value = strings.Replace(value, "%", "", 1)

	return strconv.ParseFloat(value, 64)
}
