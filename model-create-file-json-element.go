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

// checks if the CreateFileJsonElement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateFileJsonElement{}

// CreateFileJsonElement The parameters for creating a file.
type CreateFileJsonElement struct {
	// The file title for creation.
	Title NullableString `json:"title"`
	TemplateId *CreateFileJsonElementTemplateId `json:"templateId,omitempty"`
	// Specifies whether to allow creating a file of an external extension or not.
	EnableExternalExt *bool `json:"enableExternalExt,omitempty"`
	// The form ID for creation.
	FormId *int32 `json:"formId,omitempty"`
}

type _CreateFileJsonElement CreateFileJsonElement

// NewCreateFileJsonElement instantiates a new CreateFileJsonElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFileJsonElement(title NullableString) *CreateFileJsonElement {
	this := CreateFileJsonElement{}
	this.Title = title
	return &this
}

// NewCreateFileJsonElementWithDefaults instantiates a new CreateFileJsonElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFileJsonElementWithDefaults() *CreateFileJsonElement {
	this := CreateFileJsonElement{}
	return &this
}

// GetTitle returns the Title field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreateFileJsonElement) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}

	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateFileJsonElement) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// SetTitle sets field value
func (o *CreateFileJsonElement) SetTitle(v string) {
	o.Title.Set(&v)
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *CreateFileJsonElement) GetTemplateId() CreateFileJsonElementTemplateId {
	if o == nil || IsNil(o.TemplateId) {
		var ret CreateFileJsonElementTemplateId
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFileJsonElement) GetTemplateIdOk() (*CreateFileJsonElementTemplateId, bool) {
	if o == nil || IsNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *CreateFileJsonElement) IsTemplateIdSet() bool {
	if o != nil && !IsNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given CreateFileJsonElementTemplateId and assigns it to the TemplateId field.
func (o *CreateFileJsonElement) SetTemplateId(v CreateFileJsonElementTemplateId) {
	o.TemplateId = &v
}

// GetEnableExternalExt returns the EnableExternalExt field value if set, zero value otherwise.
func (o *CreateFileJsonElement) GetEnableExternalExt() bool {
	if o == nil || IsNil(o.EnableExternalExt) {
		var ret bool
		return ret
	}
	return *o.EnableExternalExt
}

// GetEnableExternalExtOk returns a tuple with the EnableExternalExt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFileJsonElement) GetEnableExternalExtOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableExternalExt) {
		return nil, false
	}
	return o.EnableExternalExt, true
}

// HasEnableExternalExt returns a boolean if a field has been set.
func (o *CreateFileJsonElement) IsEnableExternalExtSet() bool {
	if o != nil && !IsNil(o.EnableExternalExt) {
		return true
	}

	return false
}

// SetEnableExternalExt gets a reference to the given bool and assigns it to the EnableExternalExt field.
func (o *CreateFileJsonElement) SetEnableExternalExt(v bool) {
	o.EnableExternalExt = &v
}

// GetFormId returns the FormId field value if set, zero value otherwise.
func (o *CreateFileJsonElement) GetFormId() int32 {
	if o == nil || IsNil(o.FormId) {
		var ret int32
		return ret
	}
	return *o.FormId
}

// GetFormIdOk returns a tuple with the FormId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFileJsonElement) GetFormIdOk() (*int32, bool) {
	if o == nil || IsNil(o.FormId) {
		return nil, false
	}
	return o.FormId, true
}

// HasFormId returns a boolean if a field has been set.
func (o *CreateFileJsonElement) IsFormIdSet() bool {
	if o != nil && !IsNil(o.FormId) {
		return true
	}

	return false
}

// SetFormId gets a reference to the given int32 and assigns it to the FormId field.
func (o *CreateFileJsonElement) SetFormId(v int32) {
	o.FormId = &v
}

func (o CreateFileJsonElement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateFileJsonElement) ToMap() (map[string]interface{}, error) {
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

func (o *CreateFileJsonElement) UnmarshalJSON(data []byte) (err error) {
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

	varCreateFileJsonElement := _CreateFileJsonElement{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varCreateFileJsonElement)

	if err != nil {
		return err
	}

	*o = CreateFileJsonElement(varCreateFileJsonElement)

	return err
}

type NullableCreateFileJsonElement struct {
	value *CreateFileJsonElement
	isSet bool
}

func (v NullableCreateFileJsonElement) Get() *CreateFileJsonElement {
	return v.value
}

func (v *NullableCreateFileJsonElement) Set(val *CreateFileJsonElement) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFileJsonElement) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFileJsonElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFileJsonElement(val *CreateFileJsonElement) *NullableCreateFileJsonElement {
	return &NullableCreateFileJsonElement{value: val, isSet: true}
}

func (v NullableCreateFileJsonElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFileJsonElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

