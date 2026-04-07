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

// checks if the Module type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Module{}

// Module The module information.
type Module struct {
	// The module ID.
	Id *string `json:"id,omitempty"`
	// The module product class name.
	AppName NullableString `json:"appName,omitempty"`
	// The module product class name.
	Title NullableString `json:"title,omitempty"`
	// The URL to the module start page.
	Link NullableString `json:"link,omitempty"`
	// The module icon URL.
	IconUrl NullableString `json:"iconUrl,omitempty"`
	// The module large image URL.
	ImageUrl NullableString `json:"imageUrl,omitempty"`
	// The module help URL.
	HelpUrl NullableString `json:"helpUrl,omitempty"`
	// The module description.
	Description NullableString `json:"description,omitempty"`
	// Specifies if the module is primary or not.
	IsPrimary *bool `json:"isPrimary,omitempty"`
}

// NewModule instantiates a new Module object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModule() *Module {
	this := Module{}
	return &this
}

// NewModuleWithDefaults instantiates a new Module object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModuleWithDefaults() *Module {
	this := Module{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Module) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Module) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Module) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Module) SetId(v string) {
	o.Id = &v
}

// GetAppName returns the AppName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Module) GetAppName() string {
	if o == nil || IsNil(o.AppName.Get()) {
		var ret string
		return ret
	}
	return *o.AppName.Get()
}

// GetAppNameOk returns a tuple with the AppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Module) GetAppNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppName.Get(), o.AppName.IsSet()
}

// HasAppName returns a boolean if a field has been set.
func (o *Module) IsAppNameSet() bool {
	if o != nil && o.AppName.IsSet() {
		return true
	}

	return false
}

// SetAppName gets a reference to the given NullableString and assigns it to the AppName field.
func (o *Module) SetAppName(v string) {
	o.AppName.Set(&v)
}
// SetAppNameNil sets the value for AppName to be an explicit nil
func (o *Module) SetAppNameNil() {
	o.AppName.Set(nil)
}

// UnsetAppName ensures that no value is present for AppName, not even an explicit nil
func (o *Module) UnsetAppName() {
	o.AppName.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Module) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Module) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *Module) IsTitleSet() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *Module) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *Module) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *Module) UnsetTitle() {
	o.Title.Unset()
}

// GetLink returns the Link field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Module) GetLink() string {
	if o == nil || IsNil(o.Link.Get()) {
		var ret string
		return ret
	}
	return *o.Link.Get()
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Module) GetLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Link.Get(), o.Link.IsSet()
}

// HasLink returns a boolean if a field has been set.
func (o *Module) IsLinkSet() bool {
	if o != nil && o.Link.IsSet() {
		return true
	}

	return false
}

// SetLink gets a reference to the given NullableString and assigns it to the Link field.
func (o *Module) SetLink(v string) {
	o.Link.Set(&v)
}
// SetLinkNil sets the value for Link to be an explicit nil
func (o *Module) SetLinkNil() {
	o.Link.Set(nil)
}

// UnsetLink ensures that no value is present for Link, not even an explicit nil
func (o *Module) UnsetLink() {
	o.Link.Unset()
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Module) GetIconUrl() string {
	if o == nil || IsNil(o.IconUrl.Get()) {
		var ret string
		return ret
	}
	return *o.IconUrl.Get()
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Module) GetIconUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IconUrl.Get(), o.IconUrl.IsSet()
}

// HasIconUrl returns a boolean if a field has been set.
func (o *Module) IsIconUrlSet() bool {
	if o != nil && o.IconUrl.IsSet() {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given NullableString and assigns it to the IconUrl field.
func (o *Module) SetIconUrl(v string) {
	o.IconUrl.Set(&v)
}
// SetIconUrlNil sets the value for IconUrl to be an explicit nil
func (o *Module) SetIconUrlNil() {
	o.IconUrl.Set(nil)
}

// UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
func (o *Module) UnsetIconUrl() {
	o.IconUrl.Unset()
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Module) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ImageUrl.Get()
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Module) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageUrl.Get(), o.ImageUrl.IsSet()
}

// HasImageUrl returns a boolean if a field has been set.
func (o *Module) IsImageUrlSet() bool {
	if o != nil && o.ImageUrl.IsSet() {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given NullableString and assigns it to the ImageUrl field.
func (o *Module) SetImageUrl(v string) {
	o.ImageUrl.Set(&v)
}
// SetImageUrlNil sets the value for ImageUrl to be an explicit nil
func (o *Module) SetImageUrlNil() {
	o.ImageUrl.Set(nil)
}

// UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
func (o *Module) UnsetImageUrl() {
	o.ImageUrl.Unset()
}

// GetHelpUrl returns the HelpUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Module) GetHelpUrl() string {
	if o == nil || IsNil(o.HelpUrl.Get()) {
		var ret string
		return ret
	}
	return *o.HelpUrl.Get()
}

// GetHelpUrlOk returns a tuple with the HelpUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Module) GetHelpUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HelpUrl.Get(), o.HelpUrl.IsSet()
}

// HasHelpUrl returns a boolean if a field has been set.
func (o *Module) IsHelpUrlSet() bool {
	if o != nil && o.HelpUrl.IsSet() {
		return true
	}

	return false
}

// SetHelpUrl gets a reference to the given NullableString and assigns it to the HelpUrl field.
func (o *Module) SetHelpUrl(v string) {
	o.HelpUrl.Set(&v)
}
// SetHelpUrlNil sets the value for HelpUrl to be an explicit nil
func (o *Module) SetHelpUrlNil() {
	o.HelpUrl.Set(nil)
}

// UnsetHelpUrl ensures that no value is present for HelpUrl, not even an explicit nil
func (o *Module) UnsetHelpUrl() {
	o.HelpUrl.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Module) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Module) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Module) IsDescriptionSet() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Module) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Module) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Module) UnsetDescription() {
	o.Description.Unset()
}

// GetIsPrimary returns the IsPrimary field value if set, zero value otherwise.
func (o *Module) GetIsPrimary() bool {
	if o == nil || IsNil(o.IsPrimary) {
		var ret bool
		return ret
	}
	return *o.IsPrimary
}

// GetIsPrimaryOk returns a tuple with the IsPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Module) GetIsPrimaryOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPrimary) {
		return nil, false
	}
	return o.IsPrimary, true
}

// HasIsPrimary returns a boolean if a field has been set.
func (o *Module) IsIsPrimarySet() bool {
	if o != nil && !IsNil(o.IsPrimary) {
		return true
	}

	return false
}

// SetIsPrimary gets a reference to the given bool and assigns it to the IsPrimary field.
func (o *Module) SetIsPrimary(v bool) {
	o.IsPrimary = &v
}

func (o Module) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Module) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.AppName.IsSet() {
		toSerialize["appName"] = o.AppName.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Link.IsSet() {
		toSerialize["link"] = o.Link.Get()
	}
	if o.IconUrl.IsSet() {
		toSerialize["iconUrl"] = o.IconUrl.Get()
	}
	if o.ImageUrl.IsSet() {
		toSerialize["imageUrl"] = o.ImageUrl.Get()
	}
	if o.HelpUrl.IsSet() {
		toSerialize["helpUrl"] = o.HelpUrl.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.IsPrimary) {
		toSerialize["isPrimary"] = o.IsPrimary
	}
	return toSerialize, nil
}

type NullableModule struct {
	value *Module
	isSet bool
}

func (v NullableModule) Get() *Module {
	return v.value
}

func (v *NullableModule) Set(val *Module) {
	v.value = val
	v.isSet = true
}

func (v NullableModule) IsSet() bool {
	return v.isSet
}

func (v *NullableModule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModule(val *Module) *NullableModule {
	return &NullableModule{value: val, isSet: true}
}

func (v NullableModule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

