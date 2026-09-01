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

// checks if the CopyAsJsonElement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CopyAsJsonElement{}

// CopyAsJsonElement The parameters for copying a file.
type CopyAsJsonElement struct {
	// The copied file name.
	DestTitle NullableString `json:"destTitle"`
	DestFolderId CopyAsJsonElementDestFolderId `json:"destFolderId"`
	// Specifies whether to allow creating the copied file of an external extension or not.
	EnableExternalExt *bool `json:"enableExternalExt,omitempty"`
	// The copied file password.
	Password NullableString `json:"password,omitempty"`
	// Specifies whether to convert the file to form or not.
	ToForm *bool `json:"toForm,omitempty"`
}

type _CopyAsJsonElement CopyAsJsonElement

// NewCopyAsJsonElement instantiates a new CopyAsJsonElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCopyAsJsonElement(destTitle NullableString, destFolderId CopyAsJsonElementDestFolderId) *CopyAsJsonElement {
	this := CopyAsJsonElement{}
	this.DestTitle = destTitle
	this.DestFolderId = destFolderId
	return &this
}

// NewCopyAsJsonElementWithDefaults instantiates a new CopyAsJsonElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCopyAsJsonElementWithDefaults() *CopyAsJsonElement {
	this := CopyAsJsonElement{}
	return &this
}

// GetDestTitle returns the DestTitle field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CopyAsJsonElement) GetDestTitle() string {
	if o == nil || o.DestTitle.Get() == nil {
		var ret string
		return ret
	}

	return *o.DestTitle.Get()
}

// GetDestTitleOk returns a tuple with the DestTitle field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CopyAsJsonElement) GetDestTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DestTitle.Get(), o.DestTitle.IsSet()
}

// SetDestTitle sets field value
func (o *CopyAsJsonElement) SetDestTitle(v string) {
	o.DestTitle.Set(&v)
}

// GetDestFolderId returns the DestFolderId field value
func (o *CopyAsJsonElement) GetDestFolderId() CopyAsJsonElementDestFolderId {
	if o == nil {
		var ret CopyAsJsonElementDestFolderId
		return ret
	}

	return o.DestFolderId
}

// GetDestFolderIdOk returns a tuple with the DestFolderId field value
// and a boolean to check if the value has been set.
func (o *CopyAsJsonElement) GetDestFolderIdOk() (*CopyAsJsonElementDestFolderId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestFolderId, true
}

// SetDestFolderId sets field value
func (o *CopyAsJsonElement) SetDestFolderId(v CopyAsJsonElementDestFolderId) {
	o.DestFolderId = v
}

// GetEnableExternalExt returns the EnableExternalExt field value if set, zero value otherwise.
func (o *CopyAsJsonElement) GetEnableExternalExt() bool {
	if o == nil || IsNil(o.EnableExternalExt) {
		var ret bool
		return ret
	}
	return *o.EnableExternalExt
}

// GetEnableExternalExtOk returns a tuple with the EnableExternalExt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopyAsJsonElement) GetEnableExternalExtOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableExternalExt) {
		return nil, false
	}
	return o.EnableExternalExt, true
}

// HasEnableExternalExt returns a boolean if a field has been set.
func (o *CopyAsJsonElement) IsEnableExternalExtSet() bool {
	if o != nil && !IsNil(o.EnableExternalExt) {
		return true
	}

	return false
}

// SetEnableExternalExt gets a reference to the given bool and assigns it to the EnableExternalExt field.
func (o *CopyAsJsonElement) SetEnableExternalExt(v bool) {
	o.EnableExternalExt = &v
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CopyAsJsonElement) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CopyAsJsonElement) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *CopyAsJsonElement) IsPasswordSet() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *CopyAsJsonElement) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *CopyAsJsonElement) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *CopyAsJsonElement) UnsetPassword() {
	o.Password.Unset()
}

// GetToForm returns the ToForm field value if set, zero value otherwise.
func (o *CopyAsJsonElement) GetToForm() bool {
	if o == nil || IsNil(o.ToForm) {
		var ret bool
		return ret
	}
	return *o.ToForm
}

// GetToFormOk returns a tuple with the ToForm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopyAsJsonElement) GetToFormOk() (*bool, bool) {
	if o == nil || IsNil(o.ToForm) {
		return nil, false
	}
	return o.ToForm, true
}

// HasToForm returns a boolean if a field has been set.
func (o *CopyAsJsonElement) IsToFormSet() bool {
	if o != nil && !IsNil(o.ToForm) {
		return true
	}

	return false
}

// SetToForm gets a reference to the given bool and assigns it to the ToForm field.
func (o *CopyAsJsonElement) SetToForm(v bool) {
	o.ToForm = &v
}

func (o CopyAsJsonElement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CopyAsJsonElement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["destTitle"] = o.DestTitle.Get()
	toSerialize["destFolderId"] = o.DestFolderId
	if !IsNil(o.EnableExternalExt) {
		toSerialize["enableExternalExt"] = o.EnableExternalExt
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if !IsNil(o.ToForm) {
		toSerialize["toForm"] = o.ToForm
	}
	return toSerialize, nil
}

func (o *CopyAsJsonElement) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"destTitle",
		"destFolderId",
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

	varCopyAsJsonElement := _CopyAsJsonElement{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varCopyAsJsonElement)

	if err != nil {
		return err
	}

	*o = CopyAsJsonElement(varCopyAsJsonElement)

	return err
}

type NullableCopyAsJsonElement struct {
	value *CopyAsJsonElement
	isSet bool
}

func (v NullableCopyAsJsonElement) Get() *CopyAsJsonElement {
	return v.value
}

func (v *NullableCopyAsJsonElement) Set(val *CopyAsJsonElement) {
	v.value = val
	v.isSet = true
}

func (v NullableCopyAsJsonElement) IsSet() bool {
	return v.isSet
}

func (v *NullableCopyAsJsonElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCopyAsJsonElement(val *CopyAsJsonElement) *NullableCopyAsJsonElement {
	return &NullableCopyAsJsonElement{value: val, isSet: true}
}

func (v NullableCopyAsJsonElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCopyAsJsonElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

