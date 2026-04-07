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

// checks if the SecurityRequestsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityRequestsDto{}

// SecurityRequestsDto The request parameters for managing user security and access permissions.
type SecurityRequestsDto struct {
	// The product ID for which permissions are being set.
	ProductId string `json:"productId"`
	// The ID of the user whose permissions are being configured.
	UserId string `json:"userId"`
	// Specifies whether the user has administrative privileges.
	Administrator *bool `json:"administrator,omitempty"`
}

type _SecurityRequestsDto SecurityRequestsDto

// NewSecurityRequestsDto instantiates a new SecurityRequestsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityRequestsDto(productId string, userId string) *SecurityRequestsDto {
	this := SecurityRequestsDto{}
	this.ProductId = productId
	this.UserId = userId
	return &this
}

// NewSecurityRequestsDtoWithDefaults instantiates a new SecurityRequestsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityRequestsDtoWithDefaults() *SecurityRequestsDto {
	this := SecurityRequestsDto{}
	return &this
}

// GetProductId returns the ProductId field value
func (o *SecurityRequestsDto) GetProductId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value
// and a boolean to check if the value has been set.
func (o *SecurityRequestsDto) GetProductIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductId, true
}

// SetProductId sets field value
func (o *SecurityRequestsDto) SetProductId(v string) {
	o.ProductId = v
}

// GetUserId returns the UserId field value
func (o *SecurityRequestsDto) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *SecurityRequestsDto) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *SecurityRequestsDto) SetUserId(v string) {
	o.UserId = v
}

// GetAdministrator returns the Administrator field value if set, zero value otherwise.
func (o *SecurityRequestsDto) GetAdministrator() bool {
	if o == nil || IsNil(o.Administrator) {
		var ret bool
		return ret
	}
	return *o.Administrator
}

// GetAdministratorOk returns a tuple with the Administrator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityRequestsDto) GetAdministratorOk() (*bool, bool) {
	if o == nil || IsNil(o.Administrator) {
		return nil, false
	}
	return o.Administrator, true
}

// HasAdministrator returns a boolean if a field has been set.
func (o *SecurityRequestsDto) IsAdministratorSet() bool {
	if o != nil && !IsNil(o.Administrator) {
		return true
	}

	return false
}

// SetAdministrator gets a reference to the given bool and assigns it to the Administrator field.
func (o *SecurityRequestsDto) SetAdministrator(v bool) {
	o.Administrator = &v
}

func (o SecurityRequestsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityRequestsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["productId"] = o.ProductId
	toSerialize["userId"] = o.UserId
	if !IsNil(o.Administrator) {
		toSerialize["administrator"] = o.Administrator
	}
	return toSerialize, nil
}

func (o *SecurityRequestsDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"productId",
		"userId",
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

	varSecurityRequestsDto := _SecurityRequestsDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varSecurityRequestsDto)

	if err != nil {
		return err
	}

	*o = SecurityRequestsDto(varSecurityRequestsDto)

	return err
}

type NullableSecurityRequestsDto struct {
	value *SecurityRequestsDto
	isSet bool
}

func (v NullableSecurityRequestsDto) Get() *SecurityRequestsDto {
	return v.value
}

func (v *NullableSecurityRequestsDto) Set(val *SecurityRequestsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityRequestsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityRequestsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityRequestsDto(val *SecurityRequestsDto) *NullableSecurityRequestsDto {
	return &NullableSecurityRequestsDto{value: val, isSet: true}
}

func (v NullableSecurityRequestsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityRequestsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

