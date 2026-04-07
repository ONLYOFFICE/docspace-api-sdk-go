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

// checks if the UpdateServerRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateServerRequestBody{}

// UpdateServerRequestBody Parameters for updating an existing MCP server. All fields are optional — only provided fields will be modified.
type UpdateServerRequestBody struct {
	// New display name for the server. Only letters, numbers, underscores, and hyphens are allowed. Maximum 128 characters.
	Name NullableString `json:"name,omitempty"`
	// New human-readable description of the server's purpose. Maximum 255 characters.
	Description NullableString `json:"description,omitempty"`
	// New base URL of the MCP server endpoint. If changed, the system will re-verify connectivity before saving.
	Endpoint NullableString `json:"endpoint,omitempty"`
	// New HTTP headers to include with every request. If changed alongside the endpoint, connectivity is re-verified.
	Headers map[string]string `json:"headers,omitempty"`
	// Set to true to update the server icon. When true, the Icon field value (or null to remove) will be applied.
	UpdateIcon *bool `json:"updateIcon,omitempty"`
	// New Base64-encoded icon image for the server, or null to remove the existing icon. Only applied when UpdateIcon is true.
	Icon NullableString `json:"icon,omitempty"`
}

// NewUpdateServerRequestBody instantiates a new UpdateServerRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateServerRequestBody() *UpdateServerRequestBody {
	this := UpdateServerRequestBody{}
	return &this
}

// NewUpdateServerRequestBodyWithDefaults instantiates a new UpdateServerRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateServerRequestBodyWithDefaults() *UpdateServerRequestBody {
	this := UpdateServerRequestBody{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateServerRequestBody) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateServerRequestBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *UpdateServerRequestBody) IsNameSet() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *UpdateServerRequestBody) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *UpdateServerRequestBody) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *UpdateServerRequestBody) UnsetName() {
	o.Name.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateServerRequestBody) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateServerRequestBody) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateServerRequestBody) IsDescriptionSet() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *UpdateServerRequestBody) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *UpdateServerRequestBody) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *UpdateServerRequestBody) UnsetDescription() {
	o.Description.Unset()
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateServerRequestBody) GetEndpoint() string {
	if o == nil || IsNil(o.Endpoint.Get()) {
		var ret string
		return ret
	}
	return *o.Endpoint.Get()
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateServerRequestBody) GetEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Endpoint.Get(), o.Endpoint.IsSet()
}

// HasEndpoint returns a boolean if a field has been set.
func (o *UpdateServerRequestBody) IsEndpointSet() bool {
	if o != nil && o.Endpoint.IsSet() {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given NullableString and assigns it to the Endpoint field.
func (o *UpdateServerRequestBody) SetEndpoint(v string) {
	o.Endpoint.Set(&v)
}
// SetEndpointNil sets the value for Endpoint to be an explicit nil
func (o *UpdateServerRequestBody) SetEndpointNil() {
	o.Endpoint.Set(nil)
}

// UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil
func (o *UpdateServerRequestBody) UnsetEndpoint() {
	o.Endpoint.Unset()
}

// GetHeaders returns the Headers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateServerRequestBody) GetHeaders() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateServerRequestBody) GetHeadersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return &o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *UpdateServerRequestBody) IsHeadersSet() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string]string and assigns it to the Headers field.
func (o *UpdateServerRequestBody) SetHeaders(v map[string]string) {
	o.Headers = v
}

// GetUpdateIcon returns the UpdateIcon field value if set, zero value otherwise.
func (o *UpdateServerRequestBody) GetUpdateIcon() bool {
	if o == nil || IsNil(o.UpdateIcon) {
		var ret bool
		return ret
	}
	return *o.UpdateIcon
}

// GetUpdateIconOk returns a tuple with the UpdateIcon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateServerRequestBody) GetUpdateIconOk() (*bool, bool) {
	if o == nil || IsNil(o.UpdateIcon) {
		return nil, false
	}
	return o.UpdateIcon, true
}

// HasUpdateIcon returns a boolean if a field has been set.
func (o *UpdateServerRequestBody) IsUpdateIconSet() bool {
	if o != nil && !IsNil(o.UpdateIcon) {
		return true
	}

	return false
}

// SetUpdateIcon gets a reference to the given bool and assigns it to the UpdateIcon field.
func (o *UpdateServerRequestBody) SetUpdateIcon(v bool) {
	o.UpdateIcon = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateServerRequestBody) GetIcon() string {
	if o == nil || IsNil(o.Icon.Get()) {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateServerRequestBody) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *UpdateServerRequestBody) IsIconSet() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given NullableString and assigns it to the Icon field.
func (o *UpdateServerRequestBody) SetIcon(v string) {
	o.Icon.Set(&v)
}
// SetIconNil sets the value for Icon to be an explicit nil
func (o *UpdateServerRequestBody) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *UpdateServerRequestBody) UnsetIcon() {
	o.Icon.Unset()
}

func (o UpdateServerRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateServerRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Endpoint.IsSet() {
		toSerialize["endpoint"] = o.Endpoint.Get()
	}
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	if !IsNil(o.UpdateIcon) {
		toSerialize["updateIcon"] = o.UpdateIcon
	}
	if o.Icon.IsSet() {
		toSerialize["icon"] = o.Icon.Get()
	}
	return toSerialize, nil
}

type NullableUpdateServerRequestBody struct {
	value *UpdateServerRequestBody
	isSet bool
}

func (v NullableUpdateServerRequestBody) Get() *UpdateServerRequestBody {
	return v.value
}

func (v *NullableUpdateServerRequestBody) Set(val *UpdateServerRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateServerRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateServerRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateServerRequestBody(val *UpdateServerRequestBody) *NullableUpdateServerRequestBody {
	return &NullableUpdateServerRequestBody{value: val, isSet: true}
}

func (v NullableUpdateServerRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateServerRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

