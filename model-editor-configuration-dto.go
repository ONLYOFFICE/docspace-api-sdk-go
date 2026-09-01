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

// checks if the EditorConfigurationDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditorConfigurationDto{}

// EditorConfigurationDto The editor configuration parameters.
type EditorConfigurationDto struct {
	// The callback URL of the editor.
	CallbackUrl NullableString `json:"callbackUrl,omitempty"`
	// The co-editing configuration parameters.
	CoEditing *CoEditingConfig `json:"coEditing,omitempty"`
	// The creation URL of the editor.
	CreateUrl NullableString `json:"createUrl,omitempty"`
	// The customization configuration.
	Customization *CustomizationConfigDto `json:"customization,omitempty"`
	// The embedded configuration parameters for embedded documents.
	Embedded *EmbeddedConfig `json:"embedded,omitempty"`
	// The encryption keys of the editor configuration.
	EncryptionKeys []EncryptionKeyDto `json:"encryptionKeys,omitempty"`
	// The language of the editor configuration.
	Lang NullableString `json:"lang"`
	// The mode of the editor configuration.
	Mode NullableString `json:"mode"`
	// Specifies if the mode is write of the editor configuration.
	ModeWrite *bool `json:"modeWrite,omitempty"`
	// The plugins configuration.
	Plugins *PluginsConfig `json:"plugins,omitempty"`
	// The recent configuration of the editor.
	Recent []RecentConfig `json:"recent,omitempty"`
	// The templates of the editor configuration.
	Templates []TemplatesConfig `json:"templates,omitempty"`
	// The user configuration of the editor.
	User *UserConfig `json:"user,omitempty"`
}

type _EditorConfigurationDto EditorConfigurationDto

// NewEditorConfigurationDto instantiates a new EditorConfigurationDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditorConfigurationDto(lang NullableString, mode NullableString) *EditorConfigurationDto {
	this := EditorConfigurationDto{}
	this.Lang = lang
	this.Mode = mode
	return &this
}

// NewEditorConfigurationDtoWithDefaults instantiates a new EditorConfigurationDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditorConfigurationDtoWithDefaults() *EditorConfigurationDto {
	this := EditorConfigurationDto{}
	return &this
}

// GetCallbackUrl returns the CallbackUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditorConfigurationDto) GetCallbackUrl() string {
	if o == nil || IsNil(o.CallbackUrl.Get()) {
		var ret string
		return ret
	}
	return *o.CallbackUrl.Get()
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditorConfigurationDto) GetCallbackUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CallbackUrl.Get(), o.CallbackUrl.IsSet()
}

// HasCallbackUrl returns a boolean if a field has been set.
func (o *EditorConfigurationDto) IsCallbackUrlSet() bool {
	if o != nil && o.CallbackUrl.IsSet() {
		return true
	}

	return false
}

// SetCallbackUrl gets a reference to the given NullableString and assigns it to the CallbackUrl field.
func (o *EditorConfigurationDto) SetCallbackUrl(v string) {
	o.CallbackUrl.Set(&v)
}
// SetCallbackUrlNil sets the value for CallbackUrl to be an explicit nil
func (o *EditorConfigurationDto) SetCallbackUrlNil() {
	o.CallbackUrl.Set(nil)
}

// UnsetCallbackUrl ensures that no value is present for CallbackUrl, not even an explicit nil
func (o *EditorConfigurationDto) UnsetCallbackUrl() {
	o.CallbackUrl.Unset()
}

// GetCoEditing returns the CoEditing field value if set, zero value otherwise.
func (o *EditorConfigurationDto) GetCoEditing() CoEditingConfig {
	if o == nil || IsNil(o.CoEditing) {
		var ret CoEditingConfig
		return ret
	}
	return *o.CoEditing
}

// GetCoEditingOk returns a tuple with the CoEditing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditorConfigurationDto) GetCoEditingOk() (*CoEditingConfig, bool) {
	if o == nil || IsNil(o.CoEditing) {
		return nil, false
	}
	return o.CoEditing, true
}

// HasCoEditing returns a boolean if a field has been set.
func (o *EditorConfigurationDto) IsCoEditingSet() bool {
	if o != nil && !IsNil(o.CoEditing) {
		return true
	}

	return false
}

// SetCoEditing gets a reference to the given CoEditingConfig and assigns it to the CoEditing field.
func (o *EditorConfigurationDto) SetCoEditing(v CoEditingConfig) {
	o.CoEditing = &v
}

// GetCreateUrl returns the CreateUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditorConfigurationDto) GetCreateUrl() string {
	if o == nil || IsNil(o.CreateUrl.Get()) {
		var ret string
		return ret
	}
	return *o.CreateUrl.Get()
}

// GetCreateUrlOk returns a tuple with the CreateUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditorConfigurationDto) GetCreateUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreateUrl.Get(), o.CreateUrl.IsSet()
}

// HasCreateUrl returns a boolean if a field has been set.
func (o *EditorConfigurationDto) IsCreateUrlSet() bool {
	if o != nil && o.CreateUrl.IsSet() {
		return true
	}

	return false
}

// SetCreateUrl gets a reference to the given NullableString and assigns it to the CreateUrl field.
func (o *EditorConfigurationDto) SetCreateUrl(v string) {
	o.CreateUrl.Set(&v)
}
// SetCreateUrlNil sets the value for CreateUrl to be an explicit nil
func (o *EditorConfigurationDto) SetCreateUrlNil() {
	o.CreateUrl.Set(nil)
}

// UnsetCreateUrl ensures that no value is present for CreateUrl, not even an explicit nil
func (o *EditorConfigurationDto) UnsetCreateUrl() {
	o.CreateUrl.Unset()
}

// GetCustomization returns the Customization field value if set, zero value otherwise.
func (o *EditorConfigurationDto) GetCustomization() CustomizationConfigDto {
	if o == nil || IsNil(o.Customization) {
		var ret CustomizationConfigDto
		return ret
	}
	return *o.Customization
}

// GetCustomizationOk returns a tuple with the Customization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditorConfigurationDto) GetCustomizationOk() (*CustomizationConfigDto, bool) {
	if o == nil || IsNil(o.Customization) {
		return nil, false
	}
	return o.Customization, true
}

// HasCustomization returns a boolean if a field has been set.
func (o *EditorConfigurationDto) IsCustomizationSet() bool {
	if o != nil && !IsNil(o.Customization) {
		return true
	}

	return false
}

// SetCustomization gets a reference to the given CustomizationConfigDto and assigns it to the Customization field.
func (o *EditorConfigurationDto) SetCustomization(v CustomizationConfigDto) {
	o.Customization = &v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *EditorConfigurationDto) GetEmbedded() EmbeddedConfig {
	if o == nil || IsNil(o.Embedded) {
		var ret EmbeddedConfig
		return ret
	}
	return *o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditorConfigurationDto) GetEmbeddedOk() (*EmbeddedConfig, bool) {
	if o == nil || IsNil(o.Embedded) {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *EditorConfigurationDto) IsEmbeddedSet() bool {
	if o != nil && !IsNil(o.Embedded) {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given EmbeddedConfig and assigns it to the Embedded field.
func (o *EditorConfigurationDto) SetEmbedded(v EmbeddedConfig) {
	o.Embedded = &v
}

// GetEncryptionKeys returns the EncryptionKeys field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditorConfigurationDto) GetEncryptionKeys() []EncryptionKeyDto {
	if o == nil {
		var ret []EncryptionKeyDto
		return ret
	}
	return o.EncryptionKeys
}

// GetEncryptionKeysOk returns a tuple with the EncryptionKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditorConfigurationDto) GetEncryptionKeysOk() ([]EncryptionKeyDto, bool) {
	if o == nil || IsNil(o.EncryptionKeys) {
		return nil, false
	}
	return o.EncryptionKeys, true
}

// HasEncryptionKeys returns a boolean if a field has been set.
func (o *EditorConfigurationDto) IsEncryptionKeysSet() bool {
	if o != nil && !IsNil(o.EncryptionKeys) {
		return true
	}

	return false
}

// SetEncryptionKeys gets a reference to the given []EncryptionKeyDto and assigns it to the EncryptionKeys field.
func (o *EditorConfigurationDto) SetEncryptionKeys(v []EncryptionKeyDto) {
	o.EncryptionKeys = v
}

// GetLang returns the Lang field value
// If the value is explicit nil, the zero value for string will be returned
func (o *EditorConfigurationDto) GetLang() string {
	if o == nil || o.Lang.Get() == nil {
		var ret string
		return ret
	}

	return *o.Lang.Get()
}

// GetLangOk returns a tuple with the Lang field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditorConfigurationDto) GetLangOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Lang.Get(), o.Lang.IsSet()
}

// SetLang sets field value
func (o *EditorConfigurationDto) SetLang(v string) {
	o.Lang.Set(&v)
}

// GetMode returns the Mode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *EditorConfigurationDto) GetMode() string {
	if o == nil || o.Mode.Get() == nil {
		var ret string
		return ret
	}

	return *o.Mode.Get()
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditorConfigurationDto) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mode.Get(), o.Mode.IsSet()
}

// SetMode sets field value
func (o *EditorConfigurationDto) SetMode(v string) {
	o.Mode.Set(&v)
}

// GetModeWrite returns the ModeWrite field value if set, zero value otherwise.
func (o *EditorConfigurationDto) GetModeWrite() bool {
	if o == nil || IsNil(o.ModeWrite) {
		var ret bool
		return ret
	}
	return *o.ModeWrite
}

// GetModeWriteOk returns a tuple with the ModeWrite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditorConfigurationDto) GetModeWriteOk() (*bool, bool) {
	if o == nil || IsNil(o.ModeWrite) {
		return nil, false
	}
	return o.ModeWrite, true
}

// HasModeWrite returns a boolean if a field has been set.
func (o *EditorConfigurationDto) IsModeWriteSet() bool {
	if o != nil && !IsNil(o.ModeWrite) {
		return true
	}

	return false
}

// SetModeWrite gets a reference to the given bool and assigns it to the ModeWrite field.
func (o *EditorConfigurationDto) SetModeWrite(v bool) {
	o.ModeWrite = &v
}

// GetPlugins returns the Plugins field value if set, zero value otherwise.
func (o *EditorConfigurationDto) GetPlugins() PluginsConfig {
	if o == nil || IsNil(o.Plugins) {
		var ret PluginsConfig
		return ret
	}
	return *o.Plugins
}

// GetPluginsOk returns a tuple with the Plugins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditorConfigurationDto) GetPluginsOk() (*PluginsConfig, bool) {
	if o == nil || IsNil(o.Plugins) {
		return nil, false
	}
	return o.Plugins, true
}

// HasPlugins returns a boolean if a field has been set.
func (o *EditorConfigurationDto) IsPluginsSet() bool {
	if o != nil && !IsNil(o.Plugins) {
		return true
	}

	return false
}

// SetPlugins gets a reference to the given PluginsConfig and assigns it to the Plugins field.
func (o *EditorConfigurationDto) SetPlugins(v PluginsConfig) {
	o.Plugins = &v
}

// GetRecent returns the Recent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditorConfigurationDto) GetRecent() []RecentConfig {
	if o == nil {
		var ret []RecentConfig
		return ret
	}
	return o.Recent
}

// GetRecentOk returns a tuple with the Recent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditorConfigurationDto) GetRecentOk() ([]RecentConfig, bool) {
	if o == nil || IsNil(o.Recent) {
		return nil, false
	}
	return o.Recent, true
}

// HasRecent returns a boolean if a field has been set.
func (o *EditorConfigurationDto) IsRecentSet() bool {
	if o != nil && !IsNil(o.Recent) {
		return true
	}

	return false
}

// SetRecent gets a reference to the given []RecentConfig and assigns it to the Recent field.
func (o *EditorConfigurationDto) SetRecent(v []RecentConfig) {
	o.Recent = v
}

// GetTemplates returns the Templates field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditorConfigurationDto) GetTemplates() []TemplatesConfig {
	if o == nil {
		var ret []TemplatesConfig
		return ret
	}
	return o.Templates
}

// GetTemplatesOk returns a tuple with the Templates field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditorConfigurationDto) GetTemplatesOk() ([]TemplatesConfig, bool) {
	if o == nil || IsNil(o.Templates) {
		return nil, false
	}
	return o.Templates, true
}

// HasTemplates returns a boolean if a field has been set.
func (o *EditorConfigurationDto) IsTemplatesSet() bool {
	if o != nil && !IsNil(o.Templates) {
		return true
	}

	return false
}

// SetTemplates gets a reference to the given []TemplatesConfig and assigns it to the Templates field.
func (o *EditorConfigurationDto) SetTemplates(v []TemplatesConfig) {
	o.Templates = v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *EditorConfigurationDto) GetUser() UserConfig {
	if o == nil || IsNil(o.User) {
		var ret UserConfig
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditorConfigurationDto) GetUserOk() (*UserConfig, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *EditorConfigurationDto) IsUserSet() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given UserConfig and assigns it to the User field.
func (o *EditorConfigurationDto) SetUser(v UserConfig) {
	o.User = &v
}

func (o EditorConfigurationDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditorConfigurationDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CallbackUrl.IsSet() {
		toSerialize["callbackUrl"] = o.CallbackUrl.Get()
	}
	if !IsNil(o.CoEditing) {
		toSerialize["coEditing"] = o.CoEditing
	}
	if o.CreateUrl.IsSet() {
		toSerialize["createUrl"] = o.CreateUrl.Get()
	}
	if !IsNil(o.Customization) {
		toSerialize["customization"] = o.Customization
	}
	if !IsNil(o.Embedded) {
		toSerialize["embedded"] = o.Embedded
	}
	if o.EncryptionKeys != nil {
		toSerialize["encryptionKeys"] = o.EncryptionKeys
	}
	toSerialize["lang"] = o.Lang.Get()
	toSerialize["mode"] = o.Mode.Get()
	if !IsNil(o.ModeWrite) {
		toSerialize["modeWrite"] = o.ModeWrite
	}
	if !IsNil(o.Plugins) {
		toSerialize["plugins"] = o.Plugins
	}
	if o.Recent != nil {
		toSerialize["recent"] = o.Recent
	}
	if o.Templates != nil {
		toSerialize["templates"] = o.Templates
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

func (o *EditorConfigurationDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"lang",
		"mode",
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

	varEditorConfigurationDto := _EditorConfigurationDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varEditorConfigurationDto)

	if err != nil {
		return err
	}

	*o = EditorConfigurationDto(varEditorConfigurationDto)

	return err
}

type NullableEditorConfigurationDto struct {
	value *EditorConfigurationDto
	isSet bool
}

func (v NullableEditorConfigurationDto) Get() *EditorConfigurationDto {
	return v.value
}

func (v *NullableEditorConfigurationDto) Set(val *EditorConfigurationDto) {
	v.value = val
	v.isSet = true
}

func (v NullableEditorConfigurationDto) IsSet() bool {
	return v.isSet
}

func (v *NullableEditorConfigurationDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditorConfigurationDto(val *EditorConfigurationDto) *NullableEditorConfigurationDto {
	return &NullableEditorConfigurationDto{value: val, isSet: true}
}

func (v NullableEditorConfigurationDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditorConfigurationDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

