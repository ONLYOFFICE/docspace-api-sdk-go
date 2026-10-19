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

// checks if the ActionLinkActionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActionLinkActionRequest{}

// ActionLinkActionRequest An anchor inside a document, as the editor writes it.
type ActionLinkActionRequest struct {
	// The anchor value produced by the editor, opaque to the portal: it names the comment, the mention or the  place the document is scrolled to.
	Data NullableString `json:"data,omitempty"`
	// What the anchor points at, as the editor names it - a comment thread, for instance.
	Type NullableString `json:"type,omitempty"`
}

// NewActionLinkActionRequest instantiates a new ActionLinkActionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionLinkActionRequest() *ActionLinkActionRequest {
	this := ActionLinkActionRequest{}
	return &this
}

// NewActionLinkActionRequestWithDefaults instantiates a new ActionLinkActionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionLinkActionRequestWithDefaults() *ActionLinkActionRequest {
	this := ActionLinkActionRequest{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActionLinkActionRequest) GetData() string {
	if o == nil || IsNil(o.Data.Get()) {
		var ret string
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActionLinkActionRequest) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *ActionLinkActionRequest) IsDataSet() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullableString and assigns it to the Data field.
func (o *ActionLinkActionRequest) SetData(v string) {
	o.Data.Set(&v)
}
// SetDataNil sets the value for Data to be an explicit nil
func (o *ActionLinkActionRequest) SetDataNil() {
	o.Data.Set(nil)
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *ActionLinkActionRequest) UnsetData() {
	o.Data.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActionLinkActionRequest) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActionLinkActionRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *ActionLinkActionRequest) IsTypeSet() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *ActionLinkActionRequest) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *ActionLinkActionRequest) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *ActionLinkActionRequest) UnsetType() {
	o.Type.Unset()
}

func (o ActionLinkActionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActionLinkActionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data.IsSet() {
		toSerialize["data"] = o.Data.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	return toSerialize, nil
}

type NullableActionLinkActionRequest struct {
	value *ActionLinkActionRequest
	isSet bool
}

func (v NullableActionLinkActionRequest) Get() *ActionLinkActionRequest {
	return v.value
}

func (v *NullableActionLinkActionRequest) Set(val *ActionLinkActionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableActionLinkActionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableActionLinkActionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionLinkActionRequest(val *ActionLinkActionRequest) *NullableActionLinkActionRequest {
	return &NullableActionLinkActionRequest{value: val, isSet: true}
}

func (v NullableActionLinkActionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionLinkActionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

