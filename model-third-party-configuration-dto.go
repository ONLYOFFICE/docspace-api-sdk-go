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

// checks if the ThirdPartyConfigurationDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThirdPartyConfigurationDto{}

// ThirdPartyConfigurationDto Everything an editor client needs in order to open one document: the document itself, the editor setup for this  caller, and the signature that lets the editors trust both.
type ThirdPartyConfigurationDto struct {
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
	File ThirdPartyFileDto `json:"file"`
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

type _ThirdPartyConfigurationDto ThirdPartyConfigurationDto

// NewThirdPartyConfigurationDto instantiates a new ThirdPartyConfigurationDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThirdPartyConfigurationDto(document DocumentConfigDto, documentType NullableString, editorConfig EditorConfigurationDto, editorType EditorType, editorUrl NullableString, file ThirdPartyFileDto) *ThirdPartyConfigurationDto {
	this := ThirdPartyConfigurationDto{}
	this.Document = document
	this.DocumentType = documentType
	this.EditorConfig = editorConfig
	this.EditorType = editorType
	this.EditorUrl = editorUrl
	this.File = file
	return &this
}

// NewThirdPartyConfigurationDtoWithDefaults instantiates a new ThirdPartyConfigurationDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThirdPartyConfigurationDtoWithDefaults() *ThirdPartyConfigurationDto {
	this := ThirdPartyConfigurationDto{}
	return &this
}

// GetDocument returns the Document field value
func (o *ThirdPartyConfigurationDto) GetDocument() DocumentConfigDto {
	if o == nil {
		var ret DocumentConfigDto
		return ret
	}

	return o.Document
}

// GetDocumentOk returns a tuple with the Document field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyConfigurationDto) GetDocumentOk() (*DocumentConfigDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Document, true
}

// SetDocument sets field value
func (o *ThirdPartyConfigurationDto) SetDocument(v DocumentConfigDto) {
	o.Document = v
}

// GetDocumentType returns the DocumentType field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ThirdPartyConfigurationDto) GetDocumentType() string {
	if o == nil || o.DocumentType.Get() == nil {
		var ret string
		return ret
	}

	return *o.DocumentType.Get()
}

// GetDocumentTypeOk returns a tuple with the DocumentType field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyConfigurationDto) GetDocumentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DocumentType.Get(), o.DocumentType.IsSet()
}

// SetDocumentType sets field value
func (o *ThirdPartyConfigurationDto) SetDocumentType(v string) {
	o.DocumentType.Set(&v)
}

// GetEditorConfig returns the EditorConfig field value
func (o *ThirdPartyConfigurationDto) GetEditorConfig() EditorConfigurationDto {
	if o == nil {
		var ret EditorConfigurationDto
		return ret
	}

	return o.EditorConfig
}

// GetEditorConfigOk returns a tuple with the EditorConfig field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyConfigurationDto) GetEditorConfigOk() (*EditorConfigurationDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EditorConfig, true
}

// SetEditorConfig sets field value
func (o *ThirdPartyConfigurationDto) SetEditorConfig(v EditorConfigurationDto) {
	o.EditorConfig = v
}

// GetEditorType returns the EditorType field value
func (o *ThirdPartyConfigurationDto) GetEditorType() EditorType {
	if o == nil {
		var ret EditorType
		return ret
	}

	return o.EditorType
}

// GetEditorTypeOk returns a tuple with the EditorType field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyConfigurationDto) GetEditorTypeOk() (*EditorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EditorType, true
}

// SetEditorType sets field value
func (o *ThirdPartyConfigurationDto) SetEditorType(v EditorType) {
	o.EditorType = v
}

// GetEditorUrl returns the EditorUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ThirdPartyConfigurationDto) GetEditorUrl() string {
	if o == nil || o.EditorUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.EditorUrl.Get()
}

// GetEditorUrlOk returns a tuple with the EditorUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyConfigurationDto) GetEditorUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EditorUrl.Get(), o.EditorUrl.IsSet()
}

// SetEditorUrl sets field value
func (o *ThirdPartyConfigurationDto) SetEditorUrl(v string) {
	o.EditorUrl.Set(&v)
}

// GetToken returns the Token field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyConfigurationDto) GetToken() string {
	if o == nil || IsNil(o.Token.Get()) {
		var ret string
		return ret
	}
	return *o.Token.Get()
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyConfigurationDto) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Token.Get(), o.Token.IsSet()
}

// HasToken returns a boolean if a field has been set.
func (o *ThirdPartyConfigurationDto) IsTokenSet() bool {
	if o != nil && o.Token.IsSet() {
		return true
	}

	return false
}

// SetToken gets a reference to the given NullableString and assigns it to the Token field.
func (o *ThirdPartyConfigurationDto) SetToken(v string) {
	o.Token.Set(&v)
}
// SetTokenNil sets the value for Token to be an explicit nil
func (o *ThirdPartyConfigurationDto) SetTokenNil() {
	o.Token.Set(nil)
}

// UnsetToken ensures that no value is present for Token, not even an explicit nil
func (o *ThirdPartyConfigurationDto) UnsetToken() {
	o.Token.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyConfigurationDto) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyConfigurationDto) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *ThirdPartyConfigurationDto) IsTypeSet() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *ThirdPartyConfigurationDto) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *ThirdPartyConfigurationDto) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *ThirdPartyConfigurationDto) UnsetType() {
	o.Type.Unset()
}

// GetFile returns the File field value
func (o *ThirdPartyConfigurationDto) GetFile() ThirdPartyFileDto {
	if o == nil {
		var ret ThirdPartyFileDto
		return ret
	}

	return o.File
}

// GetFileOk returns a tuple with the File field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyConfigurationDto) GetFileOk() (*ThirdPartyFileDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.File, true
}

// SetFile sets field value
func (o *ThirdPartyConfigurationDto) SetFile(v ThirdPartyFileDto) {
	o.File = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyConfigurationDto) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage.Get()) {
		var ret string
		return ret
	}
	return *o.ErrorMessage.Get()
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyConfigurationDto) GetErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorMessage.Get(), o.ErrorMessage.IsSet()
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *ThirdPartyConfigurationDto) IsErrorMessageSet() bool {
	if o != nil && o.ErrorMessage.IsSet() {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given NullableString and assigns it to the ErrorMessage field.
func (o *ThirdPartyConfigurationDto) SetErrorMessage(v string) {
	o.ErrorMessage.Set(&v)
}
// SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil
func (o *ThirdPartyConfigurationDto) SetErrorMessageNil() {
	o.ErrorMessage.Set(nil)
}

// UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
func (o *ThirdPartyConfigurationDto) UnsetErrorMessage() {
	o.ErrorMessage.Unset()
}

// GetStartFilling returns the StartFilling field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyConfigurationDto) GetStartFilling() bool {
	if o == nil || IsNil(o.StartFilling.Get()) {
		var ret bool
		return ret
	}
	return *o.StartFilling.Get()
}

// GetStartFillingOk returns a tuple with the StartFilling field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyConfigurationDto) GetStartFillingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartFilling.Get(), o.StartFilling.IsSet()
}

// HasStartFilling returns a boolean if a field has been set.
func (o *ThirdPartyConfigurationDto) IsStartFillingSet() bool {
	if o != nil && o.StartFilling.IsSet() {
		return true
	}

	return false
}

// SetStartFilling gets a reference to the given NullableBool and assigns it to the StartFilling field.
func (o *ThirdPartyConfigurationDto) SetStartFilling(v bool) {
	o.StartFilling.Set(&v)
}
// SetStartFillingNil sets the value for StartFilling to be an explicit nil
func (o *ThirdPartyConfigurationDto) SetStartFillingNil() {
	o.StartFilling.Set(nil)
}

// UnsetStartFilling ensures that no value is present for StartFilling, not even an explicit nil
func (o *ThirdPartyConfigurationDto) UnsetStartFilling() {
	o.StartFilling.Unset()
}

// GetFillingStatus returns the FillingStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyConfigurationDto) GetFillingStatus() bool {
	if o == nil || IsNil(o.FillingStatus.Get()) {
		var ret bool
		return ret
	}
	return *o.FillingStatus.Get()
}

// GetFillingStatusOk returns a tuple with the FillingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyConfigurationDto) GetFillingStatusOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.FillingStatus.Get(), o.FillingStatus.IsSet()
}

// HasFillingStatus returns a boolean if a field has been set.
func (o *ThirdPartyConfigurationDto) IsFillingStatusSet() bool {
	if o != nil && o.FillingStatus.IsSet() {
		return true
	}

	return false
}

// SetFillingStatus gets a reference to the given NullableBool and assigns it to the FillingStatus field.
func (o *ThirdPartyConfigurationDto) SetFillingStatus(v bool) {
	o.FillingStatus.Set(&v)
}
// SetFillingStatusNil sets the value for FillingStatus to be an explicit nil
func (o *ThirdPartyConfigurationDto) SetFillingStatusNil() {
	o.FillingStatus.Set(nil)
}

// UnsetFillingStatus ensures that no value is present for FillingStatus, not even an explicit nil
func (o *ThirdPartyConfigurationDto) UnsetFillingStatus() {
	o.FillingStatus.Unset()
}

// GetStartFillingMode returns the StartFillingMode field value if set, zero value otherwise.
func (o *ThirdPartyConfigurationDto) GetStartFillingMode() StartFillingMode {
	if o == nil || IsNil(o.StartFillingMode) {
		var ret StartFillingMode
		return ret
	}
	return *o.StartFillingMode
}

// GetStartFillingModeOk returns a tuple with the StartFillingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyConfigurationDto) GetStartFillingModeOk() (*StartFillingMode, bool) {
	if o == nil || IsNil(o.StartFillingMode) {
		return nil, false
	}
	return o.StartFillingMode, true
}

// HasStartFillingMode returns a boolean if a field has been set.
func (o *ThirdPartyConfigurationDto) IsStartFillingModeSet() bool {
	if o != nil && !IsNil(o.StartFillingMode) {
		return true
	}

	return false
}

// SetStartFillingMode gets a reference to the given StartFillingMode and assigns it to the StartFillingMode field.
func (o *ThirdPartyConfigurationDto) SetStartFillingMode(v StartFillingMode) {
	o.StartFillingMode = &v
}

// GetFillingSessionId returns the FillingSessionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyConfigurationDto) GetFillingSessionId() string {
	if o == nil || IsNil(o.FillingSessionId.Get()) {
		var ret string
		return ret
	}
	return *o.FillingSessionId.Get()
}

// GetFillingSessionIdOk returns a tuple with the FillingSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyConfigurationDto) GetFillingSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FillingSessionId.Get(), o.FillingSessionId.IsSet()
}

// HasFillingSessionId returns a boolean if a field has been set.
func (o *ThirdPartyConfigurationDto) IsFillingSessionIdSet() bool {
	if o != nil && o.FillingSessionId.IsSet() {
		return true
	}

	return false
}

// SetFillingSessionId gets a reference to the given NullableString and assigns it to the FillingSessionId field.
func (o *ThirdPartyConfigurationDto) SetFillingSessionId(v string) {
	o.FillingSessionId.Set(&v)
}
// SetFillingSessionIdNil sets the value for FillingSessionId to be an explicit nil
func (o *ThirdPartyConfigurationDto) SetFillingSessionIdNil() {
	o.FillingSessionId.Set(nil)
}

// UnsetFillingSessionId ensures that no value is present for FillingSessionId, not even an explicit nil
func (o *ThirdPartyConfigurationDto) UnsetFillingSessionId() {
	o.FillingSessionId.Unset()
}

// GetQuotaExceededScope returns the QuotaExceededScope field value if set, zero value otherwise.
func (o *ThirdPartyConfigurationDto) GetQuotaExceededScope() QuotaScope {
	if o == nil || IsNil(o.QuotaExceededScope) {
		var ret QuotaScope
		return ret
	}
	return *o.QuotaExceededScope
}

// GetQuotaExceededScopeOk returns a tuple with the QuotaExceededScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyConfigurationDto) GetQuotaExceededScopeOk() (*QuotaScope, bool) {
	if o == nil || IsNil(o.QuotaExceededScope) {
		return nil, false
	}
	return o.QuotaExceededScope, true
}

// HasQuotaExceededScope returns a boolean if a field has been set.
func (o *ThirdPartyConfigurationDto) IsQuotaExceededScopeSet() bool {
	if o != nil && !IsNil(o.QuotaExceededScope) {
		return true
	}

	return false
}

// SetQuotaExceededScope gets a reference to the given QuotaScope and assigns it to the QuotaExceededScope field.
func (o *ThirdPartyConfigurationDto) SetQuotaExceededScope(v QuotaScope) {
	o.QuotaExceededScope = &v
}

// GetGenerationToolCallState returns the GenerationToolCallState field value if set, zero value otherwise.
func (o *ThirdPartyConfigurationDto) GetGenerationToolCallState() EditorToolCallStateDto {
	if o == nil || IsNil(o.GenerationToolCallState) {
		var ret EditorToolCallStateDto
		return ret
	}
	return *o.GenerationToolCallState
}

// GetGenerationToolCallStateOk returns a tuple with the GenerationToolCallState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyConfigurationDto) GetGenerationToolCallStateOk() (*EditorToolCallStateDto, bool) {
	if o == nil || IsNil(o.GenerationToolCallState) {
		return nil, false
	}
	return o.GenerationToolCallState, true
}

// HasGenerationToolCallState returns a boolean if a field has been set.
func (o *ThirdPartyConfigurationDto) IsGenerationToolCallStateSet() bool {
	if o != nil && !IsNil(o.GenerationToolCallState) {
		return true
	}

	return false
}

// SetGenerationToolCallState gets a reference to the given EditorToolCallStateDto and assigns it to the GenerationToolCallState field.
func (o *ThirdPartyConfigurationDto) SetGenerationToolCallState(v EditorToolCallStateDto) {
	o.GenerationToolCallState = &v
}

func (o ThirdPartyConfigurationDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThirdPartyConfigurationDto) ToMap() (map[string]interface{}, error) {
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

func (o *ThirdPartyConfigurationDto) UnmarshalJSON(data []byte) (err error) {
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

	varThirdPartyConfigurationDto := _ThirdPartyConfigurationDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varThirdPartyConfigurationDto)

	if err != nil {
		return err
	}

	*o = ThirdPartyConfigurationDto(varThirdPartyConfigurationDto)

	return err
}

type NullableThirdPartyConfigurationDto struct {
	value *ThirdPartyConfigurationDto
	isSet bool
}

func (v NullableThirdPartyConfigurationDto) Get() *ThirdPartyConfigurationDto {
	return v.value
}

func (v *NullableThirdPartyConfigurationDto) Set(val *ThirdPartyConfigurationDto) {
	v.value = val
	v.isSet = true
}

func (v NullableThirdPartyConfigurationDto) IsSet() bool {
	return v.isSet
}

func (v *NullableThirdPartyConfigurationDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThirdPartyConfigurationDto(val *ThirdPartyConfigurationDto) *NullableThirdPartyConfigurationDto {
	return &NullableThirdPartyConfigurationDto{value: val, isSet: true}
}

func (v NullableThirdPartyConfigurationDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThirdPartyConfigurationDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

