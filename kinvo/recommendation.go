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
	"strings"
)

type Recommendation struct {
	Name string
}

func IsSameProduct(recommendation string, product string) bool {
	recommendationSplit := strings.Split(recommendation, "-")

	productSplit := strings.Split(product, "-")

	return matchLists(recommendationSplit, productSplit)
}

func matchLists(first []string, second []string) bool {
	for _, firstName := range first {
		firstName = strings.TrimSpace(firstName)

		for _, secondName := range second {
			secondName = strings.TrimSpace(secondName)

			if strings.EqualFold(firstName, secondName) {
				return true
			}
		}
	}

	return false
}
