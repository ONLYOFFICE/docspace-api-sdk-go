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
)

// checks if the MetadataFieldOptionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataFieldOptionRequest{}

// MetadataFieldOptionRequest The parameters of a metadata field choice option.
type MetadataFieldOptionRequest struct {
	// The option ID. Omit it for a new option.
	Id NullableString `json:"id,omitempty"`
	// The option value.
	Value NullableString `json:"value,omitempty"`
}

// NewMetadataFieldOptionRequest instantiates a new MetadataFieldOptionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataFieldOptionRequest() *MetadataFieldOptionRequest {
	this := MetadataFieldOptionRequest{}
	return &this
}

// NewMetadataFieldOptionRequestWithDefaults instantiates a new MetadataFieldOptionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataFieldOptionRequestWithDefaults() *MetadataFieldOptionRequest {
	this := MetadataFieldOptionRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataFieldOptionRequest) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataFieldOptionRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *MetadataFieldOptionRequest) IsIdSet() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *MetadataFieldOptionRequest) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *MetadataFieldOptionRequest) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *MetadataFieldOptionRequest) UnsetId() {
	o.Id.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataFieldOptionRequest) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataFieldOptionRequest) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *MetadataFieldOptionRequest) IsValueSet() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *MetadataFieldOptionRequest) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *MetadataFieldOptionRequest) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *MetadataFieldOptionRequest) UnsetValue() {
	o.Value.Unset()
}

func (o MetadataFieldOptionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetadataFieldOptionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	return toSerialize, nil
}

type NullableMetadataFieldOptionRequest struct {
	value *MetadataFieldOptionRequest
	isSet bool
}

func (v NullableMetadataFieldOptionRequest) Get() *MetadataFieldOptionRequest {
	return v.value
}

func (v *NullableMetadataFieldOptionRequest) Set(val *MetadataFieldOptionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataFieldOptionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataFieldOptionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataFieldOptionRequest(val *MetadataFieldOptionRequest) *NullableMetadataFieldOptionRequest {
	return &NullableMetadataFieldOptionRequest{value: val, isSet: true}
}

func (v NullableMetadataFieldOptionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataFieldOptionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

