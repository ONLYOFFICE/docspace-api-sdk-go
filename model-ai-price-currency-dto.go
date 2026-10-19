// (c) Copyright Ascensio System SIA 2026
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

package docspace_api_sdk

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AiPriceCurrencyDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiPriceCurrencyDto{}

// AiPriceCurrencyDto The currency the AI prices are quoted in.
type AiPriceCurrencyDto struct {
	// The ISO 4217 code of the currency the prices are quoted in.
	Code NullableString `json:"code"`
	// The display symbol of the currency.
	Symbol NullableString `json:"symbol"`
}

type _AiPriceCurrencyDto AiPriceCurrencyDto

// NewAiPriceCurrencyDto instantiates a new AiPriceCurrencyDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiPriceCurrencyDto(code NullableString, symbol NullableString) *AiPriceCurrencyDto {
	this := AiPriceCurrencyDto{}
	this.Code = code
	this.Symbol = symbol
	return &this
}

// NewAiPriceCurrencyDtoWithDefaults instantiates a new AiPriceCurrencyDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiPriceCurrencyDtoWithDefaults() *AiPriceCurrencyDto {
	this := AiPriceCurrencyDto{}
	return &this
}

// GetCode returns the Code field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AiPriceCurrencyDto) GetCode() string {
	if o == nil || o.Code.Get() == nil {
		var ret string
		return ret
	}

	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiPriceCurrencyDto) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// SetCode sets field value
func (o *AiPriceCurrencyDto) SetCode(v string) {
	o.Code.Set(&v)
}

// GetSymbol returns the Symbol field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AiPriceCurrencyDto) GetSymbol() string {
	if o == nil || o.Symbol.Get() == nil {
		var ret string
		return ret
	}

	return *o.Symbol.Get()
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiPriceCurrencyDto) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Symbol.Get(), o.Symbol.IsSet()
}

// SetSymbol sets field value
func (o *AiPriceCurrencyDto) SetSymbol(v string) {
	o.Symbol.Set(&v)
}

func (o AiPriceCurrencyDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiPriceCurrencyDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code.Get()
	toSerialize["symbol"] = o.Symbol.Get()
	return toSerialize, nil
}

func (o *AiPriceCurrencyDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"symbol",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAiPriceCurrencyDto := _AiPriceCurrencyDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiPriceCurrencyDto)

	if err != nil {
		return err
	}

	*o = AiPriceCurrencyDto(varAiPriceCurrencyDto)

	return err
}

type NullableAiPriceCurrencyDto struct {
	value *AiPriceCurrencyDto
	isSet bool
}

func (v NullableAiPriceCurrencyDto) Get() *AiPriceCurrencyDto {
	return v.value
}

func (v *NullableAiPriceCurrencyDto) Set(val *AiPriceCurrencyDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAiPriceCurrencyDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAiPriceCurrencyDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiPriceCurrencyDto(val *AiPriceCurrencyDto) *NullableAiPriceCurrencyDto {
	return &NullableAiPriceCurrencyDto{value: val, isSet: true}
}

func (v NullableAiPriceCurrencyDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiPriceCurrencyDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

