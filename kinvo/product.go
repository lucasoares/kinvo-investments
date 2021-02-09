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

import "time"

type Product struct {
	Name                string
	AssetClass          string
	Broker              string
	FirstApplication    time.Time
	Investment          float64
	Balance             float64
	ProfitPercentage    float64
	PortfolioPercentage float64

	// Current percentage based only on filtered products.
	CurrentPercentage float64
}
