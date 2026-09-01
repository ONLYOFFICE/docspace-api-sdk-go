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

// checks if the ConfigurationDtoInteger type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigurationDtoInteger{}

// ConfigurationDtoInteger The configuration parameters.
type ConfigurationDtoInteger struct {
	// The document configuration.
	Document DocumentConfigDto `json:"document"`
	// The document type.
	DocumentType NullableString `json:"documentType"`
	// The editor configuration.
	EditorConfig EditorConfigurationDto `json:"editorConfig"`
	// The editor type.
	EditorType EditorType `json:"editorType"`
	// The editor URL.
	EditorUrl NullableString `json:"editorUrl"`
	// The token of the file configuration.
	Token NullableString `json:"token,omitempty"`
	// The platform type.
	Type NullableString `json:"type,omitempty"`
	// The file parameters.
	File FileDtoInteger `json:"file"`
	// The error message.
	ErrorMessage NullableString `json:"errorMessage,omitempty"`
	// Specifies if the file filling has started or not.
	StartFilling NullableBool `json:"startFilling,omitempty"`
	// The file filling status.
	FillingStatus NullableBool `json:"fillingStatus,omitempty"`
	// The start filling mode.
	StartFillingMode *StartFillingMode `json:"startFillingMode,omitempty"`
	// The file filling session ID.
	FillingSessionId NullableString `json:"fillingSessionId,omitempty"`
	// Indicates which quota scope has been exceeded.
	QuotaExceededScope *QuotaScope `json:"quotaExceededScope,omitempty"`
	// The generation tool call state. Used to run the agent flow in the editor.
	GenerationToolCallState *EditorToolCallStateDto `json:"generationToolCallState,omitempty"`
}

type _ConfigurationDtoInteger ConfigurationDtoInteger

// NewConfigurationDtoInteger instantiates a new ConfigurationDtoInteger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurationDtoInteger(document DocumentConfigDto, documentType NullableString, editorConfig EditorConfigurationDto, editorType EditorType, editorUrl NullableString, file FileDtoInteger) *ConfigurationDtoInteger {
	this := ConfigurationDtoInteger{}
	this.Document = document
	this.DocumentType = documentType
	this.EditorConfig = editorConfig
	this.EditorType = editorType
	this.EditorUrl = editorUrl
	this.File = file
	return &this
}

// NewConfigurationDtoIntegerWithDefaults instantiates a new ConfigurationDtoInteger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationDtoIntegerWithDefaults() *ConfigurationDtoInteger {
	this := ConfigurationDtoInteger{}
	return &this
}

// GetDocument returns the Document field value
func (o *ConfigurationDtoInteger) GetDocument() DocumentConfigDto {
	if o == nil {
		var ret DocumentConfigDto
		return ret
	}

	return o.Document
}

// GetDocumentOk returns a tuple with the Document field value
// and a boolean to check if the value has been set.
func (o *ConfigurationDtoInteger) GetDocumentOk() (*DocumentConfigDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Document, true
}

// SetDocument sets field value
func (o *ConfigurationDtoInteger) SetDocument(v DocumentConfigDto) {
	o.Document = v
}

// GetDocumentType returns the DocumentType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ConfigurationDtoInteger) GetDocumentType() string {
	if o == nil || o.DocumentType.Get() == nil {
		var ret string
		return ret
	}

	return *o.DocumentType.Get()
}

// GetDocumentTypeOk returns a tuple with the DocumentType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigurationDtoInteger) GetDocumentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DocumentType.Get(), o.DocumentType.IsSet()
}

// SetDocumentType sets field value
func (o *ConfigurationDtoInteger) SetDocumentType(v string) {
	o.DocumentType.Set(&v)
}

// GetEditorConfig returns the EditorConfig field value
func (o *ConfigurationDtoInteger) GetEditorConfig() EditorConfigurationDto {
	if o == nil {
		var ret EditorConfigurationDto
		return ret
	}

	return o.EditorConfig
}

// GetEditorConfigOk returns a tuple with the EditorConfig field value
// and a boolean to check if the value has been set.
func (o *ConfigurationDtoInteger) GetEditorConfigOk() (*EditorConfigurationDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EditorConfig, true
}

// SetEditorConfig sets field value
func (o *ConfigurationDtoInteger) SetEditorConfig(v EditorConfigurationDto) {
	o.EditorConfig = v
}

// GetEditorType returns the EditorType field value
func (o *ConfigurationDtoInteger) GetEditorType() EditorType {
	if o == nil {
		var ret EditorType
		return ret
	}

	return o.EditorType
}

// GetEditorTypeOk returns a tuple with the EditorType field value
// and a boolean to check if the value has been set.
func (o *ConfigurationDtoInteger) GetEditorTypeOk() (*EditorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EditorType, true
}

// SetEditorType sets field value
func (o *ConfigurationDtoInteger) SetEditorType(v EditorType) {
	o.EditorType = v
}

// GetEditorUrl returns the EditorUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ConfigurationDtoInteger) GetEditorUrl() string {
	if o == nil || o.EditorUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.EditorUrl.Get()
}

// GetEditorUrlOk returns a tuple with the EditorUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigurationDtoInteger) GetEditorUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EditorUrl.Get(), o.EditorUrl.IsSet()
}

// SetEditorUrl sets field value
func (o *ConfigurationDtoInteger) SetEditorUrl(v string) {
	o.EditorUrl.Set(&v)
}

// GetToken returns the Token field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigurationDtoInteger) GetToken() string {
	if o == nil || IsNil(o.Token.Get()) {
		var ret string
		return ret
	}
	return *o.Token.Get()
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigurationDtoInteger) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Token.Get(), o.Token.IsSet()
}

// HasToken returns a boolean if a field has been set.
func (o *ConfigurationDtoInteger) IsTokenSet() bool {
	if o != nil && o.Token.IsSet() {
		return true
	}

	return false
}

// SetToken gets a reference to the given NullableString and assigns it to the Token field.
func (o *ConfigurationDtoInteger) SetToken(v string) {
	o.Token.Set(&v)
}
// SetTokenNil sets the value for Token to be an explicit nil
func (o *ConfigurationDtoInteger) SetTokenNil() {
	o.Token.Set(nil)
}

// UnsetToken ensures that no value is present for Token, not even an explicit nil
func (o *ConfigurationDtoInteger) UnsetToken() {
	o.Token.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigurationDtoInteger) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigurationDtoInteger) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *ConfigurationDtoInteger) IsTypeSet() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *ConfigurationDtoInteger) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *ConfigurationDtoInteger) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *ConfigurationDtoInteger) UnsetType() {
	o.Type.Unset()
}

// GetFile returns the File field value
func (o *ConfigurationDtoInteger) GetFile() FileDtoInteger {
	if o == nil {
		var ret FileDtoInteger
		return ret
	}

	return o.File
}

// GetFileOk returns a tuple with the File field value
// and a boolean to check if the value has been set.
func (o *ConfigurationDtoInteger) GetFileOk() (*FileDtoInteger, bool) {
	if o == nil {
		return nil, false
	}
	return &o.File, true
}

// SetFile sets field value
func (o *ConfigurationDtoInteger) SetFile(v FileDtoInteger) {
	o.File = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigurationDtoInteger) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage.Get()) {
		var ret string
		return ret
	}
	return *o.ErrorMessage.Get()
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigurationDtoInteger) GetErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorMessage.Get(), o.ErrorMessage.IsSet()
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *ConfigurationDtoInteger) IsErrorMessageSet() bool {
	if o != nil && o.ErrorMessage.IsSet() {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given NullableString and assigns it to the ErrorMessage field.
func (o *ConfigurationDtoInteger) SetErrorMessage(v string) {
	o.ErrorMessage.Set(&v)
}
// SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil
func (o *ConfigurationDtoInteger) SetErrorMessageNil() {
	o.ErrorMessage.Set(nil)
}

// UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
func (o *ConfigurationDtoInteger) UnsetErrorMessage() {
	o.ErrorMessage.Unset()
}

// GetStartFilling returns the StartFilling field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigurationDtoInteger) GetStartFilling() bool {
	if o == nil || IsNil(o.StartFilling.Get()) {
		var ret bool
		return ret
	}
	return *o.StartFilling.Get()
}

// GetStartFillingOk returns a tuple with the StartFilling field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigurationDtoInteger) GetStartFillingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartFilling.Get(), o.StartFilling.IsSet()
}

// HasStartFilling returns a boolean if a field has been set.
func (o *ConfigurationDtoInteger) IsStartFillingSet() bool {
	if o != nil && o.StartFilling.IsSet() {
		return true
	}

	return false
}

// SetStartFilling gets a reference to the given NullableBool and assigns it to the StartFilling field.
func (o *ConfigurationDtoInteger) SetStartFilling(v bool) {
	o.StartFilling.Set(&v)
}
// SetStartFillingNil sets the value for StartFilling to be an explicit nil
func (o *ConfigurationDtoInteger) SetStartFillingNil() {
	o.StartFilling.Set(nil)
}

// UnsetStartFilling ensures that no value is present for StartFilling, not even an explicit nil
func (o *ConfigurationDtoInteger) UnsetStartFilling() {
	o.StartFilling.Unset()
}

// GetFillingStatus returns the FillingStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigurationDtoInteger) GetFillingStatus() bool {
	if o == nil || IsNil(o.FillingStatus.Get()) {
		var ret bool
		return ret
	}
	return *o.FillingStatus.Get()
}

// GetFillingStatusOk returns a tuple with the FillingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigurationDtoInteger) GetFillingStatusOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.FillingStatus.Get(), o.FillingStatus.IsSet()
}

// HasFillingStatus returns a boolean if a field has been set.
func (o *ConfigurationDtoInteger) IsFillingStatusSet() bool {
	if o != nil && o.FillingStatus.IsSet() {
		return true
	}

	return false
}

// SetFillingStatus gets a reference to the given NullableBool and assigns it to the FillingStatus field.
func (o *ConfigurationDtoInteger) SetFillingStatus(v bool) {
	o.FillingStatus.Set(&v)
}
// SetFillingStatusNil sets the value for FillingStatus to be an explicit nil
func (o *ConfigurationDtoInteger) SetFillingStatusNil() {
	o.FillingStatus.Set(nil)
}

// UnsetFillingStatus ensures that no value is present for FillingStatus, not even an explicit nil
func (o *ConfigurationDtoInteger) UnsetFillingStatus() {
	o.FillingStatus.Unset()
}

// GetStartFillingMode returns the StartFillingMode field value if set, zero value otherwise.
func (o *ConfigurationDtoInteger) GetStartFillingMode() StartFillingMode {
	if o == nil || IsNil(o.StartFillingMode) {
		var ret StartFillingMode
		return ret
	}
	return *o.StartFillingMode
}

// GetStartFillingModeOk returns a tuple with the StartFillingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationDtoInteger) GetStartFillingModeOk() (*StartFillingMode, bool) {
	if o == nil || IsNil(o.StartFillingMode) {
		return nil, false
	}
	return o.StartFillingMode, true
}

// HasStartFillingMode returns a boolean if a field has been set.
func (o *ConfigurationDtoInteger) IsStartFillingModeSet() bool {
	if o != nil && !IsNil(o.StartFillingMode) {
		return true
	}

	return false
}

// SetStartFillingMode gets a reference to the given StartFillingMode and assigns it to the StartFillingMode field.
func (o *ConfigurationDtoInteger) SetStartFillingMode(v StartFillingMode) {
	o.StartFillingMode = &v
}

// GetFillingSessionId returns the FillingSessionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigurationDtoInteger) GetFillingSessionId() string {
	if o == nil || IsNil(o.FillingSessionId.Get()) {
		var ret string
		return ret
	}
	return *o.FillingSessionId.Get()
}

// GetFillingSessionIdOk returns a tuple with the FillingSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigurationDtoInteger) GetFillingSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FillingSessionId.Get(), o.FillingSessionId.IsSet()
}

// HasFillingSessionId returns a boolean if a field has been set.
func (o *ConfigurationDtoInteger) IsFillingSessionIdSet() bool {
	if o != nil && o.FillingSessionId.IsSet() {
		return true
	}

	return false
}

// SetFillingSessionId gets a reference to the given NullableString and assigns it to the FillingSessionId field.
func (o *ConfigurationDtoInteger) SetFillingSessionId(v string) {
	o.FillingSessionId.Set(&v)
}
// SetFillingSessionIdNil sets the value for FillingSessionId to be an explicit nil
func (o *ConfigurationDtoInteger) SetFillingSessionIdNil() {
	o.FillingSessionId.Set(nil)
}

// UnsetFillingSessionId ensures that no value is present for FillingSessionId, not even an explicit nil
func (o *ConfigurationDtoInteger) UnsetFillingSessionId() {
	o.FillingSessionId.Unset()
}

// GetQuotaExceededScope returns the QuotaExceededScope field value if set, zero value otherwise.
func (o *ConfigurationDtoInteger) GetQuotaExceededScope() QuotaScope {
	if o == nil || IsNil(o.QuotaExceededScope) {
		var ret QuotaScope
		return ret
	}
	return *o.QuotaExceededScope
}

// GetQuotaExceededScopeOk returns a tuple with the QuotaExceededScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationDtoInteger) GetQuotaExceededScopeOk() (*QuotaScope, bool) {
	if o == nil || IsNil(o.QuotaExceededScope) {
		return nil, false
	}
	return o.QuotaExceededScope, true
}

// HasQuotaExceededScope returns a boolean if a field has been set.
func (o *ConfigurationDtoInteger) IsQuotaExceededScopeSet() bool {
	if o != nil && !IsNil(o.QuotaExceededScope) {
		return true
	}

	return false
}

// SetQuotaExceededScope gets a reference to the given QuotaScope and assigns it to the QuotaExceededScope field.
func (o *ConfigurationDtoInteger) SetQuotaExceededScope(v QuotaScope) {
	o.QuotaExceededScope = &v
}

// GetGenerationToolCallState returns the GenerationToolCallState field value if set, zero value otherwise.
func (o *ConfigurationDtoInteger) GetGenerationToolCallState() EditorToolCallStateDto {
	if o == nil || IsNil(o.GenerationToolCallState) {
		var ret EditorToolCallStateDto
		return ret
	}
	return *o.GenerationToolCallState
}

// GetGenerationToolCallStateOk returns a tuple with the GenerationToolCallState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationDtoInteger) GetGenerationToolCallStateOk() (*EditorToolCallStateDto, bool) {
	if o == nil || IsNil(o.GenerationToolCallState) {
		return nil, false
	}
	return o.GenerationToolCallState, true
}

// HasGenerationToolCallState returns a boolean if a field has been set.
func (o *ConfigurationDtoInteger) IsGenerationToolCallStateSet() bool {
	if o != nil && !IsNil(o.GenerationToolCallState) {
		return true
	}

	return false
}

// SetGenerationToolCallState gets a reference to the given EditorToolCallStateDto and assigns it to the GenerationToolCallState field.
func (o *ConfigurationDtoInteger) SetGenerationToolCallState(v EditorToolCallStateDto) {
	o.GenerationToolCallState = &v
}

func (o ConfigurationDtoInteger) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigurationDtoInteger) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["document"] = o.Document
	toSerialize["documentType"] = o.DocumentType.Get()
	toSerialize["editorConfig"] = o.EditorConfig
	toSerialize["editorType"] = o.EditorType
	toSerialize["editorUrl"] = o.EditorUrl.Get()
	if o.Token.IsSet() {
		toSerialize["token"] = o.Token.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	toSerialize["file"] = o.File
	if o.ErrorMessage.IsSet() {
		toSerialize["errorMessage"] = o.ErrorMessage.Get()
	}
	if o.StartFilling.IsSet() {
		toSerialize["startFilling"] = o.StartFilling.Get()
	}
	if o.FillingStatus.IsSet() {
		toSerialize["fillingStatus"] = o.FillingStatus.Get()
	}
	if !IsNil(o.StartFillingMode) {
		toSerialize["startFillingMode"] = o.StartFillingMode
	}
	if o.FillingSessionId.IsSet() {
		toSerialize["fillingSessionId"] = o.FillingSessionId.Get()
	}
	if !IsNil(o.QuotaExceededScope) {
		toSerialize["quotaExceededScope"] = o.QuotaExceededScope
	}
	if !IsNil(o.GenerationToolCallState) {
		toSerialize["generationToolCallState"] = o.GenerationToolCallState
	}
	return toSerialize, nil
}

func (o *ConfigurationDtoInteger) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"document",
		"documentType",
		"editorConfig",
		"editorType",
		"editorUrl",
		"file",
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

	varConfigurationDtoInteger := _ConfigurationDtoInteger{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varConfigurationDtoInteger)

	if err != nil {
		return err
	}

	*o = ConfigurationDtoInteger(varConfigurationDtoInteger)

	return err
}

type NullableConfigurationDtoInteger struct {
	value *ConfigurationDtoInteger
	isSet bool
}

func (v NullableConfigurationDtoInteger) Get() *ConfigurationDtoInteger {
	return v.value
}

func (v *NullableConfigurationDtoInteger) Set(val *ConfigurationDtoInteger) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurationDtoInteger) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurationDtoInteger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurationDtoInteger(val *ConfigurationDtoInteger) *NullableConfigurationDtoInteger {
	return &NullableConfigurationDtoInteger{value: val, isSet: true}
}

func (v NullableConfigurationDtoInteger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurationDtoInteger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

