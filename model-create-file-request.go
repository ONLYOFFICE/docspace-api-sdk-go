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

// checks if the CreateFileRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateFileRequest{}

// CreateFileRequest The parameters of a file that the portal creates from a template or a blank document.
type CreateFileRequest struct {
	// The title of the new file. The extension in it decides the format, and one of a known text, spreadsheet or  presentation format is rewritten to the DOCX, XLSX or PPTX of the portal unless `enableExternalExt` says  otherwise; a title with no extension gets DOCX added.
	Title NullableString `json:"title"`
	TemplateId *CreateFileRequestTemplateId `json:"templateId,omitempty"`
	// Whether the extension of the title is kept as it is: `true` stores the title verbatim, `false` rewrites a  known foreign format to the format the portal edits itself.
	EnableExternalExt *bool `json:"enableExternalExt,omitempty"`
	// A ready form from the form gallery of the portal to copy instead of a template, named by the identifier the  gallery reports for it. It takes precedence over `templateId`; 0 means no form.
	FormId *int32 `json:"formId,omitempty"`
}

type _CreateFileRequest CreateFileRequest

// NewCreateFileRequest instantiates a new CreateFileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFileRequest(title NullableString) *CreateFileRequest {
	this := CreateFileRequest{}
	this.Title = title
	return &this
}

// NewCreateFileRequestWithDefaults instantiates a new CreateFileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFileRequestWithDefaults() *CreateFileRequest {
	this := CreateFileRequest{}
	return &this
}

// GetTitle returns the Title field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreateFileRequest) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}

	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateFileRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// SetTitle sets field value
func (o *CreateFileRequest) SetTitle(v string) {
	o.Title.Set(&v)
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *CreateFileRequest) GetTemplateId() CreateFileRequestTemplateId {
	if o == nil || IsNil(o.TemplateId) {
		var ret CreateFileRequestTemplateId
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFileRequest) GetTemplateIdOk() (*CreateFileRequestTemplateId, bool) {
	if o == nil || IsNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *CreateFileRequest) IsTemplateIdSet() bool {
	if o != nil && !IsNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given CreateFileRequestTemplateId and assigns it to the TemplateId field.
func (o *CreateFileRequest) SetTemplateId(v CreateFileRequestTemplateId) {
	o.TemplateId = &v
}

// GetEnableExternalExt returns the EnableExternalExt field value if set, zero value otherwise.
func (o *CreateFileRequest) GetEnableExternalExt() bool {
	if o == nil || IsNil(o.EnableExternalExt) {
		var ret bool
		return ret
	}
	return *o.EnableExternalExt
}

// GetEnableExternalExtOk returns a tuple with the EnableExternalExt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFileRequest) GetEnableExternalExtOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableExternalExt) {
		return nil, false
	}
	return o.EnableExternalExt, true
}

// HasEnableExternalExt returns a boolean if a field has been set.
func (o *CreateFileRequest) IsEnableExternalExtSet() bool {
	if o != nil && !IsNil(o.EnableExternalExt) {
		return true
	}

	return false
}

// SetEnableExternalExt gets a reference to the given bool and assigns it to the EnableExternalExt field.
func (o *CreateFileRequest) SetEnableExternalExt(v bool) {
	o.EnableExternalExt = &v
}

// GetFormId returns the FormId field value if set, zero value otherwise.
func (o *CreateFileRequest) GetFormId() int32 {
	if o == nil || IsNil(o.FormId) {
		var ret int32
		return ret
	}
	return *o.FormId
}

// GetFormIdOk returns a tuple with the FormId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFileRequest) GetFormIdOk() (*int32, bool) {
	if o == nil || IsNil(o.FormId) {
		return nil, false
	}
	return o.FormId, true
}

// HasFormId returns a boolean if a field has been set.
func (o *CreateFileRequest) IsFormIdSet() bool {
	if o != nil && !IsNil(o.FormId) {
		return true
	}

	return false
}

// SetFormId gets a reference to the given int32 and assigns it to the FormId field.
func (o *CreateFileRequest) SetFormId(v int32) {
	o.FormId = &v
}

func (o CreateFileRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateFileRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["title"] = o.Title.Get()
	if !IsNil(o.TemplateId) {
		toSerialize["templateId"] = o.TemplateId
	}
	if !IsNil(o.EnableExternalExt) {
		toSerialize["enableExternalExt"] = o.EnableExternalExt
	}
	if !IsNil(o.FormId) {
		toSerialize["formId"] = o.FormId
	}
	return toSerialize, nil
}

func (o *CreateFileRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
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

	varCreateFileRequest := _CreateFileRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varCreateFileRequest)

	if err != nil {
		return err
	}

	*o = CreateFileRequest(varCreateFileRequest)

	return err
}

type NullableCreateFileRequest struct {
	value *CreateFileRequest
	isSet bool
}

func (v NullableCreateFileRequest) Get() *CreateFileRequest {
	return v.value
}

func (v *NullableCreateFileRequest) Set(val *CreateFileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFileRequest(val *CreateFileRequest) *NullableCreateFileRequest {
	return &NullableCreateFileRequest{value: val, isSet: true}
}

func (v NullableCreateFileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

