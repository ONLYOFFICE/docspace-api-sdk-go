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

// checks if the CopyAsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CopyAsRequest{}

// CopyAsRequest The parameters of a file copy that may change the format on the way.
type CopyAsRequest struct {
	// The title of the copy, extension included. That extension decides the format: the same one as the source  copies the content as it is, a different one has it converted first.
	DestTitle NullableString `json:"destTitle"`
	DestFolderId CopyAsRequestDestFolderId `json:"destFolderId"`
	// Whether the extension of the new title may be one the portal does not edit itself.
	EnableExternalExt *bool `json:"enableExternalExt,omitempty"`
	// The password that opens the source document, for a file that is protected by one.
	Password NullableString `json:"password,omitempty"`
	// Whether the copy is to become a PDF form rather than a plain document, which the conversion supports for the  text formats it can read.
	ToForm *bool `json:"toForm,omitempty"`
}

type _CopyAsRequest CopyAsRequest

// NewCopyAsRequest instantiates a new CopyAsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCopyAsRequest(destTitle NullableString, destFolderId CopyAsRequestDestFolderId) *CopyAsRequest {
	this := CopyAsRequest{}
	this.DestTitle = destTitle
	this.DestFolderId = destFolderId
	return &this
}

// NewCopyAsRequestWithDefaults instantiates a new CopyAsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCopyAsRequestWithDefaults() *CopyAsRequest {
	this := CopyAsRequest{}
	return &this
}

// GetDestTitle returns the DestTitle field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CopyAsRequest) GetDestTitle() string {
	if o == nil || o.DestTitle.Get() == nil {
		var ret string
		return ret
	}

	return *o.DestTitle.Get()
}

// GetDestTitleOk returns a tuple with the DestTitle field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CopyAsRequest) GetDestTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DestTitle.Get(), o.DestTitle.IsSet()
}

// SetDestTitle sets field value
func (o *CopyAsRequest) SetDestTitle(v string) {
	o.DestTitle.Set(&v)
}

// GetDestFolderId returns the DestFolderId field value
func (o *CopyAsRequest) GetDestFolderId() CopyAsRequestDestFolderId {
	if o == nil {
		var ret CopyAsRequestDestFolderId
		return ret
	}

	return o.DestFolderId
}

// GetDestFolderIdOk returns a tuple with the DestFolderId field value
// and a boolean to check if the value has been set.
func (o *CopyAsRequest) GetDestFolderIdOk() (*CopyAsRequestDestFolderId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestFolderId, true
}

// SetDestFolderId sets field value
func (o *CopyAsRequest) SetDestFolderId(v CopyAsRequestDestFolderId) {
	o.DestFolderId = v
}

// GetEnableExternalExt returns the EnableExternalExt field value if set, zero value otherwise.
func (o *CopyAsRequest) GetEnableExternalExt() bool {
	if o == nil || IsNil(o.EnableExternalExt) {
		var ret bool
		return ret
	}
	return *o.EnableExternalExt
}

// GetEnableExternalExtOk returns a tuple with the EnableExternalExt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopyAsRequest) GetEnableExternalExtOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableExternalExt) {
		return nil, false
	}
	return o.EnableExternalExt, true
}

// HasEnableExternalExt returns a boolean if a field has been set.
func (o *CopyAsRequest) IsEnableExternalExtSet() bool {
	if o != nil && !IsNil(o.EnableExternalExt) {
		return true
	}

	return false
}

// SetEnableExternalExt gets a reference to the given bool and assigns it to the EnableExternalExt field.
func (o *CopyAsRequest) SetEnableExternalExt(v bool) {
	o.EnableExternalExt = &v
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CopyAsRequest) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CopyAsRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *CopyAsRequest) IsPasswordSet() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *CopyAsRequest) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *CopyAsRequest) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *CopyAsRequest) UnsetPassword() {
	o.Password.Unset()
}

// GetToForm returns the ToForm field value if set, zero value otherwise.
func (o *CopyAsRequest) GetToForm() bool {
	if o == nil || IsNil(o.ToForm) {
		var ret bool
		return ret
	}
	return *o.ToForm
}

// GetToFormOk returns a tuple with the ToForm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CopyAsRequest) GetToFormOk() (*bool, bool) {
	if o == nil || IsNil(o.ToForm) {
		return nil, false
	}
	return o.ToForm, true
}

// HasToForm returns a boolean if a field has been set.
func (o *CopyAsRequest) IsToFormSet() bool {
	if o != nil && !IsNil(o.ToForm) {
		return true
	}

	return false
}

// SetToForm gets a reference to the given bool and assigns it to the ToForm field.
func (o *CopyAsRequest) SetToForm(v bool) {
	o.ToForm = &v
}

func (o CopyAsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CopyAsRequest) ToMap() (map[string]interface{}, error) {
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

func (o *CopyAsRequest) UnmarshalJSON(data []byte) (err error) {
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

	varCopyAsRequest := _CopyAsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varCopyAsRequest)

	if err != nil {
		return err
	}

	*o = CopyAsRequest(varCopyAsRequest)

	return err
}

type NullableCopyAsRequest struct {
	value *CopyAsRequest
	isSet bool
}

func (v NullableCopyAsRequest) Get() *CopyAsRequest {
	return v.value
}

func (v *NullableCopyAsRequest) Set(val *CopyAsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCopyAsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCopyAsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCopyAsRequest(val *CopyAsRequest) *NullableCopyAsRequest {
	return &NullableCopyAsRequest{value: val, isSet: true}
}

func (v NullableCopyAsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCopyAsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

