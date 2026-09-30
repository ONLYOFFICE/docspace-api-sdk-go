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

// checks if the EmbeddedConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmbeddedConfig{}

// EmbeddedConfig The addresses the framed viewer needs. It is reported for the embedded layout only.
type EmbeddedConfig struct {
	// The page to put into the frame. It is empty when the opening carries no external share key, since a framed  viewer cannot authenticate a portal member.
	EmbedUrl NullableString `json:"embedUrl,omitempty"`
	// Where the download button of the framed viewer leads.
	SaveUrl NullableString `json:"saveUrl,omitempty"`
	// The query fragment carrying the external share key, ampersand included, out of which the addresses around it  are built.
	ShareLinkParam NullableString `json:"shareLinkParam,omitempty"`
	// The address behind the share button of the framed viewer, the document opened full-screen for reading. It is  empty when the opening carries no external share key.
	ShareUrl NullableString `json:"shareUrl,omitempty"`
	// Where the framed viewer puts its toolbar. The portal always asks for the top.
	ToolbarDocked NullableString `json:"toolbarDocked,omitempty"`
}

// NewEmbeddedConfig instantiates a new EmbeddedConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmbeddedConfig() *EmbeddedConfig {
	this := EmbeddedConfig{}
	return &this
}

// NewEmbeddedConfigWithDefaults instantiates a new EmbeddedConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmbeddedConfigWithDefaults() *EmbeddedConfig {
	this := EmbeddedConfig{}
	return &this
}

// GetEmbedUrl returns the EmbedUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmbeddedConfig) GetEmbedUrl() string {
	if o == nil || IsNil(o.EmbedUrl.Get()) {
		var ret string
		return ret
	}
	return *o.EmbedUrl.Get()
}

// GetEmbedUrlOk returns a tuple with the EmbedUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmbeddedConfig) GetEmbedUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmbedUrl.Get(), o.EmbedUrl.IsSet()
}

// HasEmbedUrl returns a boolean if a field has been set.
func (o *EmbeddedConfig) IsEmbedUrlSet() bool {
	if o != nil && o.EmbedUrl.IsSet() {
		return true
	}

	return false
}

// SetEmbedUrl gets a reference to the given NullableString and assigns it to the EmbedUrl field.
func (o *EmbeddedConfig) SetEmbedUrl(v string) {
	o.EmbedUrl.Set(&v)
}
// SetEmbedUrlNil sets the value for EmbedUrl to be an explicit nil
func (o *EmbeddedConfig) SetEmbedUrlNil() {
	o.EmbedUrl.Set(nil)
}

// UnsetEmbedUrl ensures that no value is present for EmbedUrl, not even an explicit nil
func (o *EmbeddedConfig) UnsetEmbedUrl() {
	o.EmbedUrl.Unset()
}

// GetSaveUrl returns the SaveUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmbeddedConfig) GetSaveUrl() string {
	if o == nil || IsNil(o.SaveUrl.Get()) {
		var ret string
		return ret
	}
	return *o.SaveUrl.Get()
}

// GetSaveUrlOk returns a tuple with the SaveUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmbeddedConfig) GetSaveUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SaveUrl.Get(), o.SaveUrl.IsSet()
}

// HasSaveUrl returns a boolean if a field has been set.
func (o *EmbeddedConfig) IsSaveUrlSet() bool {
	if o != nil && o.SaveUrl.IsSet() {
		return true
	}

	return false
}

// SetSaveUrl gets a reference to the given NullableString and assigns it to the SaveUrl field.
func (o *EmbeddedConfig) SetSaveUrl(v string) {
	o.SaveUrl.Set(&v)
}
// SetSaveUrlNil sets the value for SaveUrl to be an explicit nil
func (o *EmbeddedConfig) SetSaveUrlNil() {
	o.SaveUrl.Set(nil)
}

// UnsetSaveUrl ensures that no value is present for SaveUrl, not even an explicit nil
func (o *EmbeddedConfig) UnsetSaveUrl() {
	o.SaveUrl.Unset()
}

// GetShareLinkParam returns the ShareLinkParam field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmbeddedConfig) GetShareLinkParam() string {
	if o == nil || IsNil(o.ShareLinkParam.Get()) {
		var ret string
		return ret
	}
	return *o.ShareLinkParam.Get()
}

// GetShareLinkParamOk returns a tuple with the ShareLinkParam field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmbeddedConfig) GetShareLinkParamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShareLinkParam.Get(), o.ShareLinkParam.IsSet()
}

// HasShareLinkParam returns a boolean if a field has been set.
func (o *EmbeddedConfig) IsShareLinkParamSet() bool {
	if o != nil && o.ShareLinkParam.IsSet() {
		return true
	}

	return false
}

// SetShareLinkParam gets a reference to the given NullableString and assigns it to the ShareLinkParam field.
func (o *EmbeddedConfig) SetShareLinkParam(v string) {
	o.ShareLinkParam.Set(&v)
}
// SetShareLinkParamNil sets the value for ShareLinkParam to be an explicit nil
func (o *EmbeddedConfig) SetShareLinkParamNil() {
	o.ShareLinkParam.Set(nil)
}

// UnsetShareLinkParam ensures that no value is present for ShareLinkParam, not even an explicit nil
func (o *EmbeddedConfig) UnsetShareLinkParam() {
	o.ShareLinkParam.Unset()
}

// GetShareUrl returns the ShareUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmbeddedConfig) GetShareUrl() string {
	if o == nil || IsNil(o.ShareUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ShareUrl.Get()
}

// GetShareUrlOk returns a tuple with the ShareUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmbeddedConfig) GetShareUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShareUrl.Get(), o.ShareUrl.IsSet()
}

// HasShareUrl returns a boolean if a field has been set.
func (o *EmbeddedConfig) IsShareUrlSet() bool {
	if o != nil && o.ShareUrl.IsSet() {
		return true
	}

	return false
}

// SetShareUrl gets a reference to the given NullableString and assigns it to the ShareUrl field.
func (o *EmbeddedConfig) SetShareUrl(v string) {
	o.ShareUrl.Set(&v)
}
// SetShareUrlNil sets the value for ShareUrl to be an explicit nil
func (o *EmbeddedConfig) SetShareUrlNil() {
	o.ShareUrl.Set(nil)
}

// UnsetShareUrl ensures that no value is present for ShareUrl, not even an explicit nil
func (o *EmbeddedConfig) UnsetShareUrl() {
	o.ShareUrl.Unset()
}

// GetToolbarDocked returns the ToolbarDocked field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmbeddedConfig) GetToolbarDocked() string {
	if o == nil || IsNil(o.ToolbarDocked.Get()) {
		var ret string
		return ret
	}
	return *o.ToolbarDocked.Get()
}

// GetToolbarDockedOk returns a tuple with the ToolbarDocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmbeddedConfig) GetToolbarDockedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ToolbarDocked.Get(), o.ToolbarDocked.IsSet()
}

// HasToolbarDocked returns a boolean if a field has been set.
func (o *EmbeddedConfig) IsToolbarDockedSet() bool {
	if o != nil && o.ToolbarDocked.IsSet() {
		return true
	}

	return false
}

// SetToolbarDocked gets a reference to the given NullableString and assigns it to the ToolbarDocked field.
func (o *EmbeddedConfig) SetToolbarDocked(v string) {
	o.ToolbarDocked.Set(&v)
}
// SetToolbarDockedNil sets the value for ToolbarDocked to be an explicit nil
func (o *EmbeddedConfig) SetToolbarDockedNil() {
	o.ToolbarDocked.Set(nil)
}

// UnsetToolbarDocked ensures that no value is present for ToolbarDocked, not even an explicit nil
func (o *EmbeddedConfig) UnsetToolbarDocked() {
	o.ToolbarDocked.Unset()
}

func (o EmbeddedConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmbeddedConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.EmbedUrl.IsSet() {
		toSerialize["embedUrl"] = o.EmbedUrl.Get()
	}
	if o.SaveUrl.IsSet() {
		toSerialize["saveUrl"] = o.SaveUrl.Get()
	}
	if o.ShareLinkParam.IsSet() {
		toSerialize["shareLinkParam"] = o.ShareLinkParam.Get()
	}
	if o.ShareUrl.IsSet() {
		toSerialize["shareUrl"] = o.ShareUrl.Get()
	}
	if o.ToolbarDocked.IsSet() {
		toSerialize["toolbarDocked"] = o.ToolbarDocked.Get()
	}
	return toSerialize, nil
}

type NullableEmbeddedConfig struct {
	value *EmbeddedConfig
	isSet bool
}

func (v NullableEmbeddedConfig) Get() *EmbeddedConfig {
	return v.value
}

func (v *NullableEmbeddedConfig) Set(val *EmbeddedConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableEmbeddedConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableEmbeddedConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmbeddedConfig(val *EmbeddedConfig) *NullableEmbeddedConfig {
	return &NullableEmbeddedConfig{value: val, isSet: true}
}

func (v NullableEmbeddedConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmbeddedConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

