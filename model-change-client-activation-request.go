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

// checks if the ChangeClientActivationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangeClientActivationRequest{}

// ChangeClientActivationRequest Client activation change request
type ChangeClientActivationRequest struct {
	// The activation status of the client
	Status bool `json:"status"`
}

type _ChangeClientActivationRequest ChangeClientActivationRequest

// NewChangeClientActivationRequest instantiates a new ChangeClientActivationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeClientActivationRequest(status bool) *ChangeClientActivationRequest {
	this := ChangeClientActivationRequest{}
	this.Status = status
	return &this
}

// NewChangeClientActivationRequestWithDefaults instantiates a new ChangeClientActivationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeClientActivationRequestWithDefaults() *ChangeClientActivationRequest {
	this := ChangeClientActivationRequest{}
	return &this
}

// GetStatus returns the Status field value
func (o *ChangeClientActivationRequest) GetStatus() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ChangeClientActivationRequest) GetStatusOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ChangeClientActivationRequest) SetStatus(v bool) {
	o.Status = v
}

func (o ChangeClientActivationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangeClientActivationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *ChangeClientActivationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
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

	varChangeClientActivationRequest := _ChangeClientActivationRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varChangeClientActivationRequest)

	if err != nil {
		return err
	}

	*o = ChangeClientActivationRequest(varChangeClientActivationRequest)

	return err
}

type NullableChangeClientActivationRequest struct {
	value *ChangeClientActivationRequest
	isSet bool
}

func (v NullableChangeClientActivationRequest) Get() *ChangeClientActivationRequest {
	return v.value
}

func (v *NullableChangeClientActivationRequest) Set(val *ChangeClientActivationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeClientActivationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeClientActivationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeClientActivationRequest(val *ChangeClientActivationRequest) *NullableChangeClientActivationRequest {
	return &NullableChangeClientActivationRequest{value: val, isSet: true}
}

func (v NullableChangeClientActivationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeClientActivationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

