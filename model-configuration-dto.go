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

// checks if the ConfigurationDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigurationDto{}

// ConfigurationDto Everything an editor client needs in order to open one document: the document itself, the editor setup for this  caller, and the signature that lets the editors trust both.
type ConfigurationDto struct {
	// The document as the editors address it: its revision key, title, type, download address and the permissions of  this caller on it.
	Document DocumentConfigDto `json:"document"`
	// The editor family the file opens in - `word`, `cell`, `slide`, `pdf` or `diagram`. It comes back empty for a  format no editor handles.
	DocumentType NullableString `json:"documentType"`
	// How the editor is set up for this opening: the mode, the language, the interface customization, the callback  the editors save through, and the account they attribute changes to.
	EditorConfig EditorConfigurationDto `json:"editorConfig"`
	// The layout the configuration was actually built for. It echoes the requested one except where the room  overruled it, as the templates folder does by forcing the embedded viewer.
	EditorType EditorType `json:"editorType"`
	// The address of the editor api script the client has to load, with the shard key of this document already  appended. Load it as it is given rather than assembling it by hand.
	EditorUrl NullableString `json:"editorUrl"`
	// Signs this whole configuration so that the editors can trust it; anything a client changes in the  configuration invalidates it. It stays empty on a portal that has no signature secret configured for the  document service.
	Token NullableString `json:"token,omitempty"`
	// The layout spelled as a lowercase word - `desktop`, `mobile` or `embedded` - the same value the editor type  carries as a number.
	Type NullableString `json:"type,omitempty"`
	// The file the configuration was built for, in the same shape the file listings report it.
	File FileDto `json:"file"`
	// Filled in when the document could not be prepared for opening; the rest of the configuration should then not  be handed to the editors.
	ErrorMessage NullableString `json:"errorMessage,omitempty"`
	// Whether this caller may start a filling session on the form from inside the editor. It stays empty when the  file is not a form opened where starting is possible at all.
	StartFilling NullableBool `json:"startFilling,omitempty"`
	// True once the caller holds a role in the running filling session of this form. It stays empty outside a  virtual data room, where roles are the only place it is set.
	FillingStatus NullableBool `json:"fillingStatus,omitempty"`
	// Which filling button the editor offers: none at all, sharing the form out for others to fill, starting a  filling session, or starting one inside the form-filling room.
	StartFillingMode *StartFillingMode `json:"startFillingMode,omitempty"`
	// Identifies the filling session this opening belongs to, and is empty when the document is not opened as part  of one. Submissions made in the editor are collected under it.
	FillingSessionId NullableString `json:"fillingSessionId,omitempty"`
	// Names the quota that ran out - the user, the room or the portal - and is set only when the document had to be  opened read-only because of it.
	QuotaExceededScope *QuotaScope `json:"quotaExceededScope,omitempty"`
	// The generation the editor should run as soon as the document opens. It is set only for a document an AI agent  produced and left waiting for its content, and is empty for every other file.
	GenerationToolCallState *EditorToolCallStateDto `json:"generationToolCallState,omitempty"`
}

type _ConfigurationDto ConfigurationDto

// NewConfigurationDto instantiates a new ConfigurationDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurationDto(document DocumentConfigDto, documentType NullableString, editorConfig EditorConfigurationDto, editorType EditorType, editorUrl NullableString, file FileDto) *ConfigurationDto {
	this := ConfigurationDto{}
	this.Document = document
	this.DocumentType = documentType
	this.EditorConfig = editorConfig
	this.EditorType = editorType
	this.EditorUrl = editorUrl
	this.File = file
	return &this
}

// NewConfigurationDtoWithDefaults instantiates a new ConfigurationDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationDtoWithDefaults() *ConfigurationDto {
	this := ConfigurationDto{}
	return &this
}

// GetDocument returns the Document field value
func (o *ConfigurationDto) GetDocument() DocumentConfigDto {
	if o == nil {
		var ret DocumentConfigDto
		return ret
	}

	return o.Document
}

// GetDocumentOk returns a tuple with the Document field value
// and a boolean to check if the value has been set.
func (o *ConfigurationDto) GetDocumentOk() (*DocumentConfigDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Document, true
}

// SetDocument sets field value
func (o *ConfigurationDto) SetDocument(v DocumentConfigDto) {
	o.Document = v
}

// GetDocumentType returns the DocumentType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ConfigurationDto) GetDocumentType() string {
	if o == nil || o.DocumentType.Get() == nil {
		var ret string
		return ret
	}

	return *o.DocumentType.Get()
}

// GetDocumentTypeOk returns a tuple with the DocumentType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigurationDto) GetDocumentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DocumentType.Get(), o.DocumentType.IsSet()
}

// SetDocumentType sets field value
func (o *ConfigurationDto) SetDocumentType(v string) {
	o.DocumentType.Set(&v)
}

// GetEditorConfig returns the EditorConfig field value
func (o *ConfigurationDto) GetEditorConfig() EditorConfigurationDto {
	if o == nil {
		var ret EditorConfigurationDto
		return ret
	}

	return o.EditorConfig
}

// GetEditorConfigOk returns a tuple with the EditorConfig field value
// and a boolean to check if the value has been set.
func (o *ConfigurationDto) GetEditorConfigOk() (*EditorConfigurationDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EditorConfig, true
}

// SetEditorConfig sets field value
func (o *ConfigurationDto) SetEditorConfig(v EditorConfigurationDto) {
	o.EditorConfig = v
}

// GetEditorType returns the EditorType field value
func (o *ConfigurationDto) GetEditorType() EditorType {
	if o == nil {
		var ret EditorType
		return ret
	}

	return o.EditorType
}

// GetEditorTypeOk returns a tuple with the EditorType field value
// and a boolean to check if the value has been set.
func (o *ConfigurationDto) GetEditorTypeOk() (*EditorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EditorType, true
}

// SetEditorType sets field value
func (o *ConfigurationDto) SetEditorType(v EditorType) {
	o.EditorType = v
}

// GetEditorUrl returns the EditorUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ConfigurationDto) GetEditorUrl() string {
	if o == nil || o.EditorUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.EditorUrl.Get()
}

// GetEditorUrlOk returns a tuple with the EditorUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigurationDto) GetEditorUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EditorUrl.Get(), o.EditorUrl.IsSet()
}

// SetEditorUrl sets field value
func (o *ConfigurationDto) SetEditorUrl(v string) {
	o.EditorUrl.Set(&v)
}

// GetToken returns the Token field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigurationDto) GetToken() string {
	if o == nil || IsNil(o.Token.Get()) {
		var ret string
		return ret
	}
	return *o.Token.Get()
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigurationDto) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Token.Get(), o.Token.IsSet()
}

// HasToken returns a boolean if a field has been set.
func (o *ConfigurationDto) IsTokenSet() bool {
	if o != nil && o.Token.IsSet() {
		return true
	}

	return false
}

// SetToken gets a reference to the given NullableString and assigns it to the Token field.
func (o *ConfigurationDto) SetToken(v string) {
	o.Token.Set(&v)
}
// SetTokenNil sets the value for Token to be an explicit nil
func (o *ConfigurationDto) SetTokenNil() {
	o.Token.Set(nil)
}

// UnsetToken ensures that no value is present for Token, not even an explicit nil
func (o *ConfigurationDto) UnsetToken() {
	o.Token.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigurationDto) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigurationDto) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *ConfigurationDto) IsTypeSet() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *ConfigurationDto) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *ConfigurationDto) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *ConfigurationDto) UnsetType() {
	o.Type.Unset()
}

// GetFile returns the File field value
func (o *ConfigurationDto) GetFile() FileDto {
	if o == nil {
		var ret FileDto
		return ret
	}

	return o.File
}

// GetFileOk returns a tuple with the File field value
// and a boolean to check if the value has been set.
func (o *ConfigurationDto) GetFileOk() (*FileDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.File, true
}

// SetFile sets field value
func (o *ConfigurationDto) SetFile(v FileDto) {
	o.File = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigurationDto) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage.Get()) {
		var ret string
		return ret
	}
	return *o.ErrorMessage.Get()
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigurationDto) GetErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorMessage.Get(), o.ErrorMessage.IsSet()
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *ConfigurationDto) IsErrorMessageSet() bool {
	if o != nil && o.ErrorMessage.IsSet() {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given NullableString and assigns it to the ErrorMessage field.
func (o *ConfigurationDto) SetErrorMessage(v string) {
	o.ErrorMessage.Set(&v)
}
// SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil
func (o *ConfigurationDto) SetErrorMessageNil() {
	o.ErrorMessage.Set(nil)
}

// UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
func (o *ConfigurationDto) UnsetErrorMessage() {
	o.ErrorMessage.Unset()
}

// GetStartFilling returns the StartFilling field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigurationDto) GetStartFilling() bool {
	if o == nil || IsNil(o.StartFilling.Get()) {
		var ret bool
		return ret
	}
	return *o.StartFilling.Get()
}

// GetStartFillingOk returns a tuple with the StartFilling field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigurationDto) GetStartFillingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartFilling.Get(), o.StartFilling.IsSet()
}

// HasStartFilling returns a boolean if a field has been set.
func (o *ConfigurationDto) IsStartFillingSet() bool {
	if o != nil && o.StartFilling.IsSet() {
		return true
	}

	return false
}

// SetStartFilling gets a reference to the given NullableBool and assigns it to the StartFilling field.
func (o *ConfigurationDto) SetStartFilling(v bool) {
	o.StartFilling.Set(&v)
}
// SetStartFillingNil sets the value for StartFilling to be an explicit nil
func (o *ConfigurationDto) SetStartFillingNil() {
	o.StartFilling.Set(nil)
}

// UnsetStartFilling ensures that no value is present for StartFilling, not even an explicit nil
func (o *ConfigurationDto) UnsetStartFilling() {
	o.StartFilling.Unset()
}

// GetFillingStatus returns the FillingStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigurationDto) GetFillingStatus() bool {
	if o == nil || IsNil(o.FillingStatus.Get()) {
		var ret bool
		return ret
	}
	return *o.FillingStatus.Get()
}

// GetFillingStatusOk returns a tuple with the FillingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigurationDto) GetFillingStatusOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.FillingStatus.Get(), o.FillingStatus.IsSet()
}

// HasFillingStatus returns a boolean if a field has been set.
func (o *ConfigurationDto) IsFillingStatusSet() bool {
	if o != nil && o.FillingStatus.IsSet() {
		return true
	}

	return false
}

// SetFillingStatus gets a reference to the given NullableBool and assigns it to the FillingStatus field.
func (o *ConfigurationDto) SetFillingStatus(v bool) {
	o.FillingStatus.Set(&v)
}
// SetFillingStatusNil sets the value for FillingStatus to be an explicit nil
func (o *ConfigurationDto) SetFillingStatusNil() {
	o.FillingStatus.Set(nil)
}

// UnsetFillingStatus ensures that no value is present for FillingStatus, not even an explicit nil
func (o *ConfigurationDto) UnsetFillingStatus() {
	o.FillingStatus.Unset()
}

// GetStartFillingMode returns the StartFillingMode field value if set, zero value otherwise.
func (o *ConfigurationDto) GetStartFillingMode() StartFillingMode {
	if o == nil || IsNil(o.StartFillingMode) {
		var ret StartFillingMode
		return ret
	}
	return *o.StartFillingMode
}

// GetStartFillingModeOk returns a tuple with the StartFillingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationDto) GetStartFillingModeOk() (*StartFillingMode, bool) {
	if o == nil || IsNil(o.StartFillingMode) {
		return nil, false
	}
	return o.StartFillingMode, true
}

// HasStartFillingMode returns a boolean if a field has been set.
func (o *ConfigurationDto) IsStartFillingModeSet() bool {
	if o != nil && !IsNil(o.StartFillingMode) {
		return true
	}

	return false
}

// SetStartFillingMode gets a reference to the given StartFillingMode and assigns it to the StartFillingMode field.
func (o *ConfigurationDto) SetStartFillingMode(v StartFillingMode) {
	o.StartFillingMode = &v
}

// GetFillingSessionId returns the FillingSessionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigurationDto) GetFillingSessionId() string {
	if o == nil || IsNil(o.FillingSessionId.Get()) {
		var ret string
		return ret
	}
	return *o.FillingSessionId.Get()
}

// GetFillingSessionIdOk returns a tuple with the FillingSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigurationDto) GetFillingSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FillingSessionId.Get(), o.FillingSessionId.IsSet()
}

// HasFillingSessionId returns a boolean if a field has been set.
func (o *ConfigurationDto) IsFillingSessionIdSet() bool {
	if o != nil && o.FillingSessionId.IsSet() {
		return true
	}

	return false
}

// SetFillingSessionId gets a reference to the given NullableString and assigns it to the FillingSessionId field.
func (o *ConfigurationDto) SetFillingSessionId(v string) {
	o.FillingSessionId.Set(&v)
}
// SetFillingSessionIdNil sets the value for FillingSessionId to be an explicit nil
func (o *ConfigurationDto) SetFillingSessionIdNil() {
	o.FillingSessionId.Set(nil)
}

// UnsetFillingSessionId ensures that no value is present for FillingSessionId, not even an explicit nil
func (o *ConfigurationDto) UnsetFillingSessionId() {
	o.FillingSessionId.Unset()
}

// GetQuotaExceededScope returns the QuotaExceededScope field value if set, zero value otherwise.
func (o *ConfigurationDto) GetQuotaExceededScope() QuotaScope {
	if o == nil || IsNil(o.QuotaExceededScope) {
		var ret QuotaScope
		return ret
	}
	return *o.QuotaExceededScope
}

// GetQuotaExceededScopeOk returns a tuple with the QuotaExceededScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationDto) GetQuotaExceededScopeOk() (*QuotaScope, bool) {
	if o == nil || IsNil(o.QuotaExceededScope) {
		return nil, false
	}
	return o.QuotaExceededScope, true
}

// HasQuotaExceededScope returns a boolean if a field has been set.
func (o *ConfigurationDto) IsQuotaExceededScopeSet() bool {
	if o != nil && !IsNil(o.QuotaExceededScope) {
		return true
	}

	return false
}

// SetQuotaExceededScope gets a reference to the given QuotaScope and assigns it to the QuotaExceededScope field.
func (o *ConfigurationDto) SetQuotaExceededScope(v QuotaScope) {
	o.QuotaExceededScope = &v
}

// GetGenerationToolCallState returns the GenerationToolCallState field value if set, zero value otherwise.
func (o *ConfigurationDto) GetGenerationToolCallState() EditorToolCallStateDto {
	if o == nil || IsNil(o.GenerationToolCallState) {
		var ret EditorToolCallStateDto
		return ret
	}
	return *o.GenerationToolCallState
}

// GetGenerationToolCallStateOk returns a tuple with the GenerationToolCallState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationDto) GetGenerationToolCallStateOk() (*EditorToolCallStateDto, bool) {
	if o == nil || IsNil(o.GenerationToolCallState) {
		return nil, false
	}
	return o.GenerationToolCallState, true
}

// HasGenerationToolCallState returns a boolean if a field has been set.
func (o *ConfigurationDto) IsGenerationToolCallStateSet() bool {
	if o != nil && !IsNil(o.GenerationToolCallState) {
		return true
	}

	return false
}

// SetGenerationToolCallState gets a reference to the given EditorToolCallStateDto and assigns it to the GenerationToolCallState field.
func (o *ConfigurationDto) SetGenerationToolCallState(v EditorToolCallStateDto) {
	o.GenerationToolCallState = &v
}

func (o ConfigurationDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigurationDto) ToMap() (map[string]interface{}, error) {
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

func (o *ConfigurationDto) UnmarshalJSON(data []byte) (err error) {
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

	varConfigurationDto := _ConfigurationDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varConfigurationDto)

	if err != nil {
		return err
	}

	*o = ConfigurationDto(varConfigurationDto)

	return err
}

type NullableConfigurationDto struct {
	value *ConfigurationDto
	isSet bool
}

func (v NullableConfigurationDto) Get() *ConfigurationDto {
	return v.value
}

func (v *NullableConfigurationDto) Set(val *ConfigurationDto) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurationDto) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurationDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurationDto(val *ConfigurationDto) *NullableConfigurationDto {
	return &NullableConfigurationDto{value: val, isSet: true}
}

func (v NullableConfigurationDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurationDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

