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
	"time"
	"bytes"
	"fmt"
)

// checks if the WebPluginDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebPluginDto{}

// WebPluginDto The web plugin information.
type WebPluginDto struct {
	// The web plugin name.
	Name NullableString `json:"name"`
	// The web plugin version.
	Version NullableString `json:"version"`
	// The minimum version of DocSpace with which the plugin is guaranteed to work.
	MinDocSpaceVersion NullableString `json:"minDocSpaceVersion,omitempty"`
	// The web plugin description.
	Description NullableString `json:"description"`
	// The web plugin license.
	License NullableString `json:"license"`
	// The web plugin author.
	Author NullableString `json:"author"`
	// The web plugin home page URL.
	HomePage NullableString `json:"homePage"`
	// The name by which the web plugin is registered in the window object.
	PluginName NullableString `json:"pluginName"`
	// The web plugin scopes.
	Scopes NullableString `json:"scopes"`
	// The web plugin image.
	Image NullableString `json:"image"`
	CreateBy EmployeeDto `json:"createBy"`
	// The date and time when the web plugin was created.
	CreateOn time.Time `json:"createOn"`
	// Specifies if the web plugin is enabled or not.
	Enabled bool `json:"enabled"`
	// Specifies if the web plugin is system or not.
	System bool `json:"system"`
	// The web plugin URL.
	Url NullableString `json:"url"`
	// The web plugin css URL.
	CssUrl NullableString `json:"cssUrl"`
	// The web plugin settings.
	Settings NullableString `json:"settings"`
	// The web plugin localized name.
	NameLocale map[string]string `json:"nameLocale,omitempty"`
	// The web plugin localized description.
	DescriptionLocale map[string]string `json:"descriptionLocale,omitempty"`
}

type _WebPluginDto WebPluginDto

// NewWebPluginDto instantiates a new WebPluginDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebPluginDto(name NullableString, version NullableString, description NullableString, license NullableString, author NullableString, homePage NullableString, pluginName NullableString, scopes NullableString, image NullableString, createBy EmployeeDto, createOn time.Time, enabled bool, system bool, url NullableString, cssUrl NullableString, settings NullableString) *WebPluginDto {
	this := WebPluginDto{}
	this.Name = name
	this.Version = version
	this.Description = description
	this.License = license
	this.Author = author
	this.HomePage = homePage
	this.PluginName = pluginName
	this.Scopes = scopes
	this.Image = image
	this.CreateBy = createBy
	this.CreateOn = createOn
	this.Enabled = enabled
	this.System = system
	this.Url = url
	this.CssUrl = cssUrl
	this.Settings = settings
	return &this
}

// NewWebPluginDtoWithDefaults instantiates a new WebPluginDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebPluginDtoWithDefaults() *WebPluginDto {
	this := WebPluginDto{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WebPluginDto) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebPluginDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *WebPluginDto) SetName(v string) {
	o.Name.Set(&v)
}

// GetVersion returns the Version field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WebPluginDto) GetVersion() string {
	if o == nil || o.Version.Get() == nil {
		var ret string
		return ret
	}

	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebPluginDto) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// SetVersion sets field value
func (o *WebPluginDto) SetVersion(v string) {
	o.Version.Set(&v)
}

// GetMinDocSpaceVersion returns the MinDocSpaceVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebPluginDto) GetMinDocSpaceVersion() string {
	if o == nil || IsNil(o.MinDocSpaceVersion.Get()) {
		var ret string
		return ret
	}
	return *o.MinDocSpaceVersion.Get()
}

// GetMinDocSpaceVersionOk returns a tuple with the MinDocSpaceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebPluginDto) GetMinDocSpaceVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinDocSpaceVersion.Get(), o.MinDocSpaceVersion.IsSet()
}

// HasMinDocSpaceVersion returns a boolean if a field has been set.
func (o *WebPluginDto) IsMinDocSpaceVersionSet() bool {
	if o != nil && o.MinDocSpaceVersion.IsSet() {
		return true
	}

	return false
}

// SetMinDocSpaceVersion gets a reference to the given NullableString and assigns it to the MinDocSpaceVersion field.
func (o *WebPluginDto) SetMinDocSpaceVersion(v string) {
	o.MinDocSpaceVersion.Set(&v)
}
// SetMinDocSpaceVersionNil sets the value for MinDocSpaceVersion to be an explicit nil
func (o *WebPluginDto) SetMinDocSpaceVersionNil() {
	o.MinDocSpaceVersion.Set(nil)
}

// UnsetMinDocSpaceVersion ensures that no value is present for MinDocSpaceVersion, not even an explicit nil
func (o *WebPluginDto) UnsetMinDocSpaceVersion() {
	o.MinDocSpaceVersion.Unset()
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WebPluginDto) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebPluginDto) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *WebPluginDto) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetLicense returns the License field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WebPluginDto) GetLicense() string {
	if o == nil || o.License.Get() == nil {
		var ret string
		return ret
	}

	return *o.License.Get()
}

// GetLicenseOk returns a tuple with the License field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebPluginDto) GetLicenseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.License.Get(), o.License.IsSet()
}

// SetLicense sets field value
func (o *WebPluginDto) SetLicense(v string) {
	o.License.Set(&v)
}

// GetAuthor returns the Author field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WebPluginDto) GetAuthor() string {
	if o == nil || o.Author.Get() == nil {
		var ret string
		return ret
	}

	return *o.Author.Get()
}

// GetAuthorOk returns a tuple with the Author field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebPluginDto) GetAuthorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Author.Get(), o.Author.IsSet()
}

// SetAuthor sets field value
func (o *WebPluginDto) SetAuthor(v string) {
	o.Author.Set(&v)
}

// GetHomePage returns the HomePage field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WebPluginDto) GetHomePage() string {
	if o == nil || o.HomePage.Get() == nil {
		var ret string
		return ret
	}

	return *o.HomePage.Get()
}

// GetHomePageOk returns a tuple with the HomePage field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebPluginDto) GetHomePageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HomePage.Get(), o.HomePage.IsSet()
}

// SetHomePage sets field value
func (o *WebPluginDto) SetHomePage(v string) {
	o.HomePage.Set(&v)
}

// GetPluginName returns the PluginName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WebPluginDto) GetPluginName() string {
	if o == nil || o.PluginName.Get() == nil {
		var ret string
		return ret
	}

	return *o.PluginName.Get()
}

// GetPluginNameOk returns a tuple with the PluginName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebPluginDto) GetPluginNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PluginName.Get(), o.PluginName.IsSet()
}

// SetPluginName sets field value
func (o *WebPluginDto) SetPluginName(v string) {
	o.PluginName.Set(&v)
}

// GetScopes returns the Scopes field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WebPluginDto) GetScopes() string {
	if o == nil || o.Scopes.Get() == nil {
		var ret string
		return ret
	}

	return *o.Scopes.Get()
}

// GetScopesOk returns a tuple with the Scopes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebPluginDto) GetScopesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scopes.Get(), o.Scopes.IsSet()
}

// SetScopes sets field value
func (o *WebPluginDto) SetScopes(v string) {
	o.Scopes.Set(&v)
}

// GetImage returns the Image field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WebPluginDto) GetImage() string {
	if o == nil || o.Image.Get() == nil {
		var ret string
		return ret
	}

	return *o.Image.Get()
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebPluginDto) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Image.Get(), o.Image.IsSet()
}

// SetImage sets field value
func (o *WebPluginDto) SetImage(v string) {
	o.Image.Set(&v)
}

// GetCreateBy returns the CreateBy field value
func (o *WebPluginDto) GetCreateBy() EmployeeDto {
	if o == nil {
		var ret EmployeeDto
		return ret
	}

	return o.CreateBy
}

// GetCreateByOk returns a tuple with the CreateBy field value
// and a boolean to check if the value has been set.
func (o *WebPluginDto) GetCreateByOk() (*EmployeeDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreateBy, true
}

// SetCreateBy sets field value
func (o *WebPluginDto) SetCreateBy(v EmployeeDto) {
	o.CreateBy = v
}

// GetCreateOn returns the CreateOn field value
func (o *WebPluginDto) GetCreateOn() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreateOn
}

// GetCreateOnOk returns a tuple with the CreateOn field value
// and a boolean to check if the value has been set.
func (o *WebPluginDto) GetCreateOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreateOn, true
}

// SetCreateOn sets field value
func (o *WebPluginDto) SetCreateOn(v time.Time) {
	o.CreateOn = v
}

// GetEnabled returns the Enabled field value
func (o *WebPluginDto) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *WebPluginDto) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *WebPluginDto) SetEnabled(v bool) {
	o.Enabled = v
}

// GetSystem returns the System field value
func (o *WebPluginDto) GetSystem() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.System
}

// GetSystemOk returns a tuple with the System field value
// and a boolean to check if the value has been set.
func (o *WebPluginDto) GetSystemOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.System, true
}

// SetSystem sets field value
func (o *WebPluginDto) SetSystem(v bool) {
	o.System = v
}

// GetUrl returns the Url field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WebPluginDto) GetUrl() string {
	if o == nil || o.Url.Get() == nil {
		var ret string
		return ret
	}

	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebPluginDto) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// SetUrl sets field value
func (o *WebPluginDto) SetUrl(v string) {
	o.Url.Set(&v)
}

// GetCssUrl returns the CssUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WebPluginDto) GetCssUrl() string {
	if o == nil || o.CssUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.CssUrl.Get()
}

// GetCssUrlOk returns a tuple with the CssUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebPluginDto) GetCssUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CssUrl.Get(), o.CssUrl.IsSet()
}

// SetCssUrl sets field value
func (o *WebPluginDto) SetCssUrl(v string) {
	o.CssUrl.Set(&v)
}

// GetSettings returns the Settings field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WebPluginDto) GetSettings() string {
	if o == nil || o.Settings.Get() == nil {
		var ret string
		return ret
	}

	return *o.Settings.Get()
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebPluginDto) GetSettingsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Settings.Get(), o.Settings.IsSet()
}

// SetSettings sets field value
func (o *WebPluginDto) SetSettings(v string) {
	o.Settings.Set(&v)
}

// GetNameLocale returns the NameLocale field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebPluginDto) GetNameLocale() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.NameLocale
}

// GetNameLocaleOk returns a tuple with the NameLocale field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebPluginDto) GetNameLocaleOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.NameLocale) {
		return nil, false
	}
	return &o.NameLocale, true
}

// HasNameLocale returns a boolean if a field has been set.
func (o *WebPluginDto) IsNameLocaleSet() bool {
	if o != nil && !IsNil(o.NameLocale) {
		return true
	}

	return false
}

// SetNameLocale gets a reference to the given map[string]string and assigns it to the NameLocale field.
func (o *WebPluginDto) SetNameLocale(v map[string]string) {
	o.NameLocale = v
}

// GetDescriptionLocale returns the DescriptionLocale field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebPluginDto) GetDescriptionLocale() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.DescriptionLocale
}

// GetDescriptionLocaleOk returns a tuple with the DescriptionLocale field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebPluginDto) GetDescriptionLocaleOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.DescriptionLocale) {
		return nil, false
	}
	return &o.DescriptionLocale, true
}

// HasDescriptionLocale returns a boolean if a field has been set.
func (o *WebPluginDto) IsDescriptionLocaleSet() bool {
	if o != nil && !IsNil(o.DescriptionLocale) {
		return true
	}

	return false
}

// SetDescriptionLocale gets a reference to the given map[string]string and assigns it to the DescriptionLocale field.
func (o *WebPluginDto) SetDescriptionLocale(v map[string]string) {
	o.DescriptionLocale = v
}

func (o WebPluginDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebPluginDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name.Get()
	toSerialize["version"] = o.Version.Get()
	if o.MinDocSpaceVersion.IsSet() {
		toSerialize["minDocSpaceVersion"] = o.MinDocSpaceVersion.Get()
	}
	toSerialize["description"] = o.Description.Get()
	toSerialize["license"] = o.License.Get()
	toSerialize["author"] = o.Author.Get()
	toSerialize["homePage"] = o.HomePage.Get()
	toSerialize["pluginName"] = o.PluginName.Get()
	toSerialize["scopes"] = o.Scopes.Get()
	toSerialize["image"] = o.Image.Get()
	toSerialize["createBy"] = o.CreateBy
	toSerialize["createOn"] = o.CreateOn
	toSerialize["enabled"] = o.Enabled
	toSerialize["system"] = o.System
	toSerialize["url"] = o.Url.Get()
	toSerialize["cssUrl"] = o.CssUrl.Get()
	toSerialize["settings"] = o.Settings.Get()
	if o.NameLocale != nil {
		toSerialize["nameLocale"] = o.NameLocale
	}
	if o.DescriptionLocale != nil {
		toSerialize["descriptionLocale"] = o.DescriptionLocale
	}
	return toSerialize, nil
}

func (o *WebPluginDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"version",
		"description",
		"license",
		"author",
		"homePage",
		"pluginName",
		"scopes",
		"image",
		"createBy",
		"createOn",
		"enabled",
		"system",
		"url",
		"cssUrl",
		"settings",
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

	varWebPluginDto := _WebPluginDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varWebPluginDto)

	if err != nil {
		return err
	}

	*o = WebPluginDto(varWebPluginDto)

	return err
}

type NullableWebPluginDto struct {
	value *WebPluginDto
	isSet bool
}

func (v NullableWebPluginDto) Get() *WebPluginDto {
	return v.value
}

func (v *NullableWebPluginDto) Set(val *WebPluginDto) {
	v.value = val
	v.isSet = true
}

func (v NullableWebPluginDto) IsSet() bool {
	return v.isSet
}

func (v *NullableWebPluginDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebPluginDto(val *WebPluginDto) *NullableWebPluginDto {
	return &NullableWebPluginDto{value: val, isSet: true}
}

func (v NullableWebPluginDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebPluginDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

