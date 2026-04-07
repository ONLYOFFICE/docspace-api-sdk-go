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

// checks if the CheckDocServiceUrlRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckDocServiceUrlRequestDto{}

// CheckDocServiceUrlRequestDto The request parameters for checking the document service location.
type CheckDocServiceUrlRequestDto struct {
	// The ONLYOFFICE Docs URL address.
	DocServiceUrl NullableString `json:"docServiceUrl"`
	// The ONLYOFFICE Docs URL address in the local private network.
	DocServiceUrlInternal NullableString `json:"docServiceUrlInternal,omitempty"`
	// The ONLYOFFICE Docs URL address.
	DocServiceUrlPortal NullableString `json:"docServiceUrlPortal,omitempty"`
	// The signature secret of the ONLYOFFICE Docs.
	DocServiceSignatureSecret NullableString `json:"docServiceSignatureSecret,omitempty"`
	// The signature header of the ONLYOFFICE Docs.
	DocServiceSignatureHeader NullableString `json:"docServiceSignatureHeader,omitempty"`
	// Specifies if the SSL verification of the ONLYOFFICE Docs is enabled or not.
	DocServiceSslVerification NullableBool `json:"docServiceSslVerification,omitempty"`
}

type _CheckDocServiceUrlRequestDto CheckDocServiceUrlRequestDto

// NewCheckDocServiceUrlRequestDto instantiates a new CheckDocServiceUrlRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckDocServiceUrlRequestDto(docServiceUrl NullableString) *CheckDocServiceUrlRequestDto {
	this := CheckDocServiceUrlRequestDto{}
	this.DocServiceUrl = docServiceUrl
	return &this
}

// NewCheckDocServiceUrlRequestDtoWithDefaults instantiates a new CheckDocServiceUrlRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckDocServiceUrlRequestDtoWithDefaults() *CheckDocServiceUrlRequestDto {
	this := CheckDocServiceUrlRequestDto{}
	return &this
}

// GetDocServiceUrl returns the DocServiceUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CheckDocServiceUrlRequestDto) GetDocServiceUrl() string {
	if o == nil || o.DocServiceUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.DocServiceUrl.Get()
}

// GetDocServiceUrlOk returns a tuple with the DocServiceUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckDocServiceUrlRequestDto) GetDocServiceUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DocServiceUrl.Get(), o.DocServiceUrl.IsSet()
}

// SetDocServiceUrl sets field value
func (o *CheckDocServiceUrlRequestDto) SetDocServiceUrl(v string) {
	o.DocServiceUrl.Set(&v)
}

// GetDocServiceUrlInternal returns the DocServiceUrlInternal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CheckDocServiceUrlRequestDto) GetDocServiceUrlInternal() string {
	if o == nil || IsNil(o.DocServiceUrlInternal.Get()) {
		var ret string
		return ret
	}
	return *o.DocServiceUrlInternal.Get()
}

// GetDocServiceUrlInternalOk returns a tuple with the DocServiceUrlInternal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckDocServiceUrlRequestDto) GetDocServiceUrlInternalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DocServiceUrlInternal.Get(), o.DocServiceUrlInternal.IsSet()
}

// HasDocServiceUrlInternal returns a boolean if a field has been set.
func (o *CheckDocServiceUrlRequestDto) IsDocServiceUrlInternalSet() bool {
	if o != nil && o.DocServiceUrlInternal.IsSet() {
		return true
	}

	return false
}

// SetDocServiceUrlInternal gets a reference to the given NullableString and assigns it to the DocServiceUrlInternal field.
func (o *CheckDocServiceUrlRequestDto) SetDocServiceUrlInternal(v string) {
	o.DocServiceUrlInternal.Set(&v)
}
// SetDocServiceUrlInternalNil sets the value for DocServiceUrlInternal to be an explicit nil
func (o *CheckDocServiceUrlRequestDto) SetDocServiceUrlInternalNil() {
	o.DocServiceUrlInternal.Set(nil)
}

// UnsetDocServiceUrlInternal ensures that no value is present for DocServiceUrlInternal, not even an explicit nil
func (o *CheckDocServiceUrlRequestDto) UnsetDocServiceUrlInternal() {
	o.DocServiceUrlInternal.Unset()
}

// GetDocServiceUrlPortal returns the DocServiceUrlPortal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CheckDocServiceUrlRequestDto) GetDocServiceUrlPortal() string {
	if o == nil || IsNil(o.DocServiceUrlPortal.Get()) {
		var ret string
		return ret
	}
	return *o.DocServiceUrlPortal.Get()
}

// GetDocServiceUrlPortalOk returns a tuple with the DocServiceUrlPortal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckDocServiceUrlRequestDto) GetDocServiceUrlPortalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DocServiceUrlPortal.Get(), o.DocServiceUrlPortal.IsSet()
}

// HasDocServiceUrlPortal returns a boolean if a field has been set.
func (o *CheckDocServiceUrlRequestDto) IsDocServiceUrlPortalSet() bool {
	if o != nil && o.DocServiceUrlPortal.IsSet() {
		return true
	}

	return false
}

// SetDocServiceUrlPortal gets a reference to the given NullableString and assigns it to the DocServiceUrlPortal field.
func (o *CheckDocServiceUrlRequestDto) SetDocServiceUrlPortal(v string) {
	o.DocServiceUrlPortal.Set(&v)
}
// SetDocServiceUrlPortalNil sets the value for DocServiceUrlPortal to be an explicit nil
func (o *CheckDocServiceUrlRequestDto) SetDocServiceUrlPortalNil() {
	o.DocServiceUrlPortal.Set(nil)
}

// UnsetDocServiceUrlPortal ensures that no value is present for DocServiceUrlPortal, not even an explicit nil
func (o *CheckDocServiceUrlRequestDto) UnsetDocServiceUrlPortal() {
	o.DocServiceUrlPortal.Unset()
}

// GetDocServiceSignatureSecret returns the DocServiceSignatureSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CheckDocServiceUrlRequestDto) GetDocServiceSignatureSecret() string {
	if o == nil || IsNil(o.DocServiceSignatureSecret.Get()) {
		var ret string
		return ret
	}
	return *o.DocServiceSignatureSecret.Get()
}

// GetDocServiceSignatureSecretOk returns a tuple with the DocServiceSignatureSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckDocServiceUrlRequestDto) GetDocServiceSignatureSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DocServiceSignatureSecret.Get(), o.DocServiceSignatureSecret.IsSet()
}

// HasDocServiceSignatureSecret returns a boolean if a field has been set.
func (o *CheckDocServiceUrlRequestDto) IsDocServiceSignatureSecretSet() bool {
	if o != nil && o.DocServiceSignatureSecret.IsSet() {
		return true
	}

	return false
}

// SetDocServiceSignatureSecret gets a reference to the given NullableString and assigns it to the DocServiceSignatureSecret field.
func (o *CheckDocServiceUrlRequestDto) SetDocServiceSignatureSecret(v string) {
	o.DocServiceSignatureSecret.Set(&v)
}
// SetDocServiceSignatureSecretNil sets the value for DocServiceSignatureSecret to be an explicit nil
func (o *CheckDocServiceUrlRequestDto) SetDocServiceSignatureSecretNil() {
	o.DocServiceSignatureSecret.Set(nil)
}

// UnsetDocServiceSignatureSecret ensures that no value is present for DocServiceSignatureSecret, not even an explicit nil
func (o *CheckDocServiceUrlRequestDto) UnsetDocServiceSignatureSecret() {
	o.DocServiceSignatureSecret.Unset()
}

// GetDocServiceSignatureHeader returns the DocServiceSignatureHeader field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CheckDocServiceUrlRequestDto) GetDocServiceSignatureHeader() string {
	if o == nil || IsNil(o.DocServiceSignatureHeader.Get()) {
		var ret string
		return ret
	}
	return *o.DocServiceSignatureHeader.Get()
}

// GetDocServiceSignatureHeaderOk returns a tuple with the DocServiceSignatureHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckDocServiceUrlRequestDto) GetDocServiceSignatureHeaderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DocServiceSignatureHeader.Get(), o.DocServiceSignatureHeader.IsSet()
}

// HasDocServiceSignatureHeader returns a boolean if a field has been set.
func (o *CheckDocServiceUrlRequestDto) IsDocServiceSignatureHeaderSet() bool {
	if o != nil && o.DocServiceSignatureHeader.IsSet() {
		return true
	}

	return false
}

// SetDocServiceSignatureHeader gets a reference to the given NullableString and assigns it to the DocServiceSignatureHeader field.
func (o *CheckDocServiceUrlRequestDto) SetDocServiceSignatureHeader(v string) {
	o.DocServiceSignatureHeader.Set(&v)
}
// SetDocServiceSignatureHeaderNil sets the value for DocServiceSignatureHeader to be an explicit nil
func (o *CheckDocServiceUrlRequestDto) SetDocServiceSignatureHeaderNil() {
	o.DocServiceSignatureHeader.Set(nil)
}

// UnsetDocServiceSignatureHeader ensures that no value is present for DocServiceSignatureHeader, not even an explicit nil
func (o *CheckDocServiceUrlRequestDto) UnsetDocServiceSignatureHeader() {
	o.DocServiceSignatureHeader.Unset()
}

// GetDocServiceSslVerification returns the DocServiceSslVerification field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CheckDocServiceUrlRequestDto) GetDocServiceSslVerification() bool {
	if o == nil || IsNil(o.DocServiceSslVerification.Get()) {
		var ret bool
		return ret
	}
	return *o.DocServiceSslVerification.Get()
}

// GetDocServiceSslVerificationOk returns a tuple with the DocServiceSslVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckDocServiceUrlRequestDto) GetDocServiceSslVerificationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DocServiceSslVerification.Get(), o.DocServiceSslVerification.IsSet()
}

// HasDocServiceSslVerification returns a boolean if a field has been set.
func (o *CheckDocServiceUrlRequestDto) IsDocServiceSslVerificationSet() bool {
	if o != nil && o.DocServiceSslVerification.IsSet() {
		return true
	}

	return false
}

// SetDocServiceSslVerification gets a reference to the given NullableBool and assigns it to the DocServiceSslVerification field.
func (o *CheckDocServiceUrlRequestDto) SetDocServiceSslVerification(v bool) {
	o.DocServiceSslVerification.Set(&v)
}
// SetDocServiceSslVerificationNil sets the value for DocServiceSslVerification to be an explicit nil
func (o *CheckDocServiceUrlRequestDto) SetDocServiceSslVerificationNil() {
	o.DocServiceSslVerification.Set(nil)
}

// UnsetDocServiceSslVerification ensures that no value is present for DocServiceSslVerification, not even an explicit nil
func (o *CheckDocServiceUrlRequestDto) UnsetDocServiceSslVerification() {
	o.DocServiceSslVerification.Unset()
}

func (o CheckDocServiceUrlRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckDocServiceUrlRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["docServiceUrl"] = o.DocServiceUrl.Get()
	if o.DocServiceUrlInternal.IsSet() {
		toSerialize["docServiceUrlInternal"] = o.DocServiceUrlInternal.Get()
	}
	if o.DocServiceUrlPortal.IsSet() {
		toSerialize["docServiceUrlPortal"] = o.DocServiceUrlPortal.Get()
	}
	if o.DocServiceSignatureSecret.IsSet() {
		toSerialize["docServiceSignatureSecret"] = o.DocServiceSignatureSecret.Get()
	}
	if o.DocServiceSignatureHeader.IsSet() {
		toSerialize["docServiceSignatureHeader"] = o.DocServiceSignatureHeader.Get()
	}
	if o.DocServiceSslVerification.IsSet() {
		toSerialize["docServiceSslVerification"] = o.DocServiceSslVerification.Get()
	}
	return toSerialize, nil
}

func (o *CheckDocServiceUrlRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"docServiceUrl",
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

	varCheckDocServiceUrlRequestDto := _CheckDocServiceUrlRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varCheckDocServiceUrlRequestDto)

	if err != nil {
		return err
	}

	*o = CheckDocServiceUrlRequestDto(varCheckDocServiceUrlRequestDto)

	return err
}

type NullableCheckDocServiceUrlRequestDto struct {
	value *CheckDocServiceUrlRequestDto
	isSet bool
}

func (v NullableCheckDocServiceUrlRequestDto) Get() *CheckDocServiceUrlRequestDto {
	return v.value
}

func (v *NullableCheckDocServiceUrlRequestDto) Set(val *CheckDocServiceUrlRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckDocServiceUrlRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckDocServiceUrlRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckDocServiceUrlRequestDto(val *CheckDocServiceUrlRequestDto) *NullableCheckDocServiceUrlRequestDto {
	return &NullableCheckDocServiceUrlRequestDto{value: val, isSet: true}
}

func (v NullableCheckDocServiceUrlRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckDocServiceUrlRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

