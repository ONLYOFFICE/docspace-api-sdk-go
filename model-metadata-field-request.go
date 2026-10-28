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

// checks if the MetadataFieldRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataFieldRequest{}

// MetadataFieldRequest The parameters of a metadata field.
type MetadataFieldRequest struct {
	// The field name.
	Name NullableString `json:"name,omitempty"`
	// The field type.
	Type *MetadataFieldType `json:"type,omitempty"`
	// The choice options of the field.
	Options []MetadataFieldOptionRequest `json:"options,omitempty"`
	// The display position of the field inside the template: the fields are shown by it ascending, and equal positions  keep the order of creation.
	Order NullableInt32 `json:"order,omitempty"`
}

// NewMetadataFieldRequest instantiates a new MetadataFieldRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataFieldRequest() *MetadataFieldRequest {
	this := MetadataFieldRequest{}
	return &this
}

// NewMetadataFieldRequestWithDefaults instantiates a new MetadataFieldRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataFieldRequestWithDefaults() *MetadataFieldRequest {
	this := MetadataFieldRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataFieldRequest) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataFieldRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *MetadataFieldRequest) IsNameSet() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *MetadataFieldRequest) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *MetadataFieldRequest) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *MetadataFieldRequest) UnsetName() {
	o.Name.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MetadataFieldRequest) GetType() MetadataFieldType {
	if o == nil || IsNil(o.Type) {
		var ret MetadataFieldType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataFieldRequest) GetTypeOk() (*MetadataFieldType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MetadataFieldRequest) IsTypeSet() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given MetadataFieldType and assigns it to the Type field.
func (o *MetadataFieldRequest) SetType(v MetadataFieldType) {
	o.Type = &v
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataFieldRequest) GetOptions() []MetadataFieldOptionRequest {
	if o == nil {
		var ret []MetadataFieldOptionRequest
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataFieldRequest) GetOptionsOk() ([]MetadataFieldOptionRequest, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *MetadataFieldRequest) IsOptionsSet() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []MetadataFieldOptionRequest and assigns it to the Options field.
func (o *MetadataFieldRequest) SetOptions(v []MetadataFieldOptionRequest) {
	o.Options = v
}

// GetOrder returns the Order field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataFieldRequest) GetOrder() int32 {
	if o == nil || IsNil(o.Order.Get()) {
		var ret int32
		return ret
	}
	return *o.Order.Get()
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataFieldRequest) GetOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Order.Get(), o.Order.IsSet()
}

// HasOrder returns a boolean if a field has been set.
func (o *MetadataFieldRequest) IsOrderSet() bool {
	if o != nil && o.Order.IsSet() {
		return true
	}

	return false
}

// SetOrder gets a reference to the given NullableInt32 and assigns it to the Order field.
func (o *MetadataFieldRequest) SetOrder(v int32) {
	o.Order.Set(&v)
}
// SetOrderNil sets the value for Order to be an explicit nil
func (o *MetadataFieldRequest) SetOrderNil() {
	o.Order.Set(nil)
}

// UnsetOrder ensures that no value is present for Order, not even an explicit nil
func (o *MetadataFieldRequest) UnsetOrder() {
	o.Order.Unset()
}

func (o MetadataFieldRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetadataFieldRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.Order.IsSet() {
		toSerialize["order"] = o.Order.Get()
	}
	return toSerialize, nil
}

type NullableMetadataFieldRequest struct {
	value *MetadataFieldRequest
	isSet bool
}

func (v NullableMetadataFieldRequest) Get() *MetadataFieldRequest {
	return v.value
}

func (v *NullableMetadataFieldRequest) Set(val *MetadataFieldRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataFieldRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataFieldRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataFieldRequest(val *MetadataFieldRequest) *NullableMetadataFieldRequest {
	return &NullableMetadataFieldRequest{value: val, isSet: true}
}

func (v NullableMetadataFieldRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataFieldRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

