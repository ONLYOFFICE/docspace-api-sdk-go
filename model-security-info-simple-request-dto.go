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
)

// checks if the SecurityInfoSimpleRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityInfoSimpleRequestDto{}

// SecurityInfoSimpleRequestDto The parameters of the security information request.
type SecurityInfoSimpleRequestDto struct {
	// The collection of sharing parameters.
	Share []FileShareParams `json:"share,omitempty"`
	// Specifies whether to notify users about the shared file or not.
	Notify *bool `json:"notify,omitempty"`
	// The message to send when notifying about the shared file.
	SharingMessage NullableString `json:"sharingMessage,omitempty"`
}

// NewSecurityInfoSimpleRequestDto instantiates a new SecurityInfoSimpleRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityInfoSimpleRequestDto() *SecurityInfoSimpleRequestDto {
	this := SecurityInfoSimpleRequestDto{}
	return &this
}

// NewSecurityInfoSimpleRequestDtoWithDefaults instantiates a new SecurityInfoSimpleRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityInfoSimpleRequestDtoWithDefaults() *SecurityInfoSimpleRequestDto {
	this := SecurityInfoSimpleRequestDto{}
	return &this
}

// GetShare returns the Share field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecurityInfoSimpleRequestDto) GetShare() []FileShareParams {
	if o == nil {
		var ret []FileShareParams
		return ret
	}
	return o.Share
}

// GetShareOk returns a tuple with the Share field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecurityInfoSimpleRequestDto) GetShareOk() ([]FileShareParams, bool) {
	if o == nil || IsNil(o.Share) {
		return nil, false
	}
	return o.Share, true
}

// HasShare returns a boolean if a field has been set.
func (o *SecurityInfoSimpleRequestDto) IsShareSet() bool {
	if o != nil && !IsNil(o.Share) {
		return true
	}

	return false
}

// SetShare gets a reference to the given []FileShareParams and assigns it to the Share field.
func (o *SecurityInfoSimpleRequestDto) SetShare(v []FileShareParams) {
	o.Share = v
}

// GetNotify returns the Notify field value if set, zero value otherwise.
func (o *SecurityInfoSimpleRequestDto) GetNotify() bool {
	if o == nil || IsNil(o.Notify) {
		var ret bool
		return ret
	}
	return *o.Notify
}

// GetNotifyOk returns a tuple with the Notify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityInfoSimpleRequestDto) GetNotifyOk() (*bool, bool) {
	if o == nil || IsNil(o.Notify) {
		return nil, false
	}
	return o.Notify, true
}

// HasNotify returns a boolean if a field has been set.
func (o *SecurityInfoSimpleRequestDto) IsNotifySet() bool {
	if o != nil && !IsNil(o.Notify) {
		return true
	}

	return false
}

// SetNotify gets a reference to the given bool and assigns it to the Notify field.
func (o *SecurityInfoSimpleRequestDto) SetNotify(v bool) {
	o.Notify = &v
}

// GetSharingMessage returns the SharingMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecurityInfoSimpleRequestDto) GetSharingMessage() string {
	if o == nil || IsNil(o.SharingMessage.Get()) {
		var ret string
		return ret
	}
	return *o.SharingMessage.Get()
}

// GetSharingMessageOk returns a tuple with the SharingMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecurityInfoSimpleRequestDto) GetSharingMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SharingMessage.Get(), o.SharingMessage.IsSet()
}

// HasSharingMessage returns a boolean if a field has been set.
func (o *SecurityInfoSimpleRequestDto) IsSharingMessageSet() bool {
	if o != nil && o.SharingMessage.IsSet() {
		return true
	}

	return false
}

// SetSharingMessage gets a reference to the given NullableString and assigns it to the SharingMessage field.
func (o *SecurityInfoSimpleRequestDto) SetSharingMessage(v string) {
	o.SharingMessage.Set(&v)
}
// SetSharingMessageNil sets the value for SharingMessage to be an explicit nil
func (o *SecurityInfoSimpleRequestDto) SetSharingMessageNil() {
	o.SharingMessage.Set(nil)
}

// UnsetSharingMessage ensures that no value is present for SharingMessage, not even an explicit nil
func (o *SecurityInfoSimpleRequestDto) UnsetSharingMessage() {
	o.SharingMessage.Unset()
}

func (o SecurityInfoSimpleRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityInfoSimpleRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Share != nil {
		toSerialize["share"] = o.Share
	}
	if !IsNil(o.Notify) {
		toSerialize["notify"] = o.Notify
	}
	if o.SharingMessage.IsSet() {
		toSerialize["sharingMessage"] = o.SharingMessage.Get()
	}
	return toSerialize, nil
}

type NullableSecurityInfoSimpleRequestDto struct {
	value *SecurityInfoSimpleRequestDto
	isSet bool
}

func (v NullableSecurityInfoSimpleRequestDto) Get() *SecurityInfoSimpleRequestDto {
	return v.value
}

func (v *NullableSecurityInfoSimpleRequestDto) Set(val *SecurityInfoSimpleRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityInfoSimpleRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityInfoSimpleRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityInfoSimpleRequestDto(val *SecurityInfoSimpleRequestDto) *NullableSecurityInfoSimpleRequestDto {
	return &NullableSecurityInfoSimpleRequestDto{value: val, isSet: true}
}

func (v NullableSecurityInfoSimpleRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityInfoSimpleRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

