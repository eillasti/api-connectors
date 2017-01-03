/* 
 * BitMEX API
 *
 * REST API for the BitMEX.com trading platform.<br><br><a href=\"/app/restAPI\">REST Documentation</a><br><a href=\"/app/wsAPI\">Websocket Documentation</a>
 *
 * OpenAPI spec version: 1.2.0
 * Contact: support@bitmex.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package swagger

import (
	"time"
)

type TradeBin struct {

	Timestamp time.Time `json:"timestamp,omitempty"`

	Symbol string `json:"symbol,omitempty"`

	Open float64 `json:"open,omitempty"`

	High float64 `json:"high,omitempty"`

	Low float64 `json:"low,omitempty"`

	Close float64 `json:"close,omitempty"`

	Trades float32 `json:"trades,omitempty"`

	Volume float32 `json:"volume,omitempty"`

	Vwap float64 `json:"vwap,omitempty"`

	LastSize float32 `json:"lastSize,omitempty"`

	Turnover float32 `json:"turnover,omitempty"`

	HomeNotional float64 `json:"homeNotional,omitempty"`

	ForeignNotional float64 `json:"foreignNotional,omitempty"`

	Id float64 `json:"id,omitempty"`
}