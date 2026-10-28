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

// checks if the FirebaseDeviceDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FirebaseDeviceDto{}

// FirebaseDeviceDto One mobile device of the calling user registered for push notifications.
type FirebaseDeviceDto struct {
	// The id of the registration.
	Id *int32 `json:"id,omitempty"`
	// The account the device belongs to; always the caller.
	UserId *string `json:"userId,omitempty"`
	// The portal the registration belongs to; always the current one.
	TenantId *int32 `json:"tenantId,omitempty"`
	// The Firebase token the device was issued, as it was sent at registration.
	FirebaseDeviceToken NullableString `json:"firebaseDeviceToken,omitempty"`
	// The application the registration is for; `doc` for the Documents application.
	Application NullableString `json:"application,omitempty"`
	// Whether the device is currently sent push notifications.
	IsSubscribed NullableBool `json:"isSubscribed,omitempty"`
}

// NewFirebaseDeviceDto instantiates a new FirebaseDeviceDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirebaseDeviceDto() *FirebaseDeviceDto {
	this := FirebaseDeviceDto{}
	return &this
}

// NewFirebaseDeviceDtoWithDefaults instantiates a new FirebaseDeviceDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirebaseDeviceDtoWithDefaults() *FirebaseDeviceDto {
	this := FirebaseDeviceDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FirebaseDeviceDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirebaseDeviceDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FirebaseDeviceDto) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *FirebaseDeviceDto) SetId(v int32) {
	o.Id = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *FirebaseDeviceDto) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirebaseDeviceDto) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *FirebaseDeviceDto) IsUserIdSet() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *FirebaseDeviceDto) SetUserId(v string) {
	o.UserId = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *FirebaseDeviceDto) GetTenantId() int32 {
	if o == nil || IsNil(o.TenantId) {
		var ret int32
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirebaseDeviceDto) GetTenantIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *FirebaseDeviceDto) IsTenantIdSet() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given int32 and assigns it to the TenantId field.
func (o *FirebaseDeviceDto) SetTenantId(v int32) {
	o.TenantId = &v
}

// GetFirebaseDeviceToken returns the FirebaseDeviceToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirebaseDeviceDto) GetFirebaseDeviceToken() string {
	if o == nil || IsNil(o.FirebaseDeviceToken.Get()) {
		var ret string
		return ret
	}
	return *o.FirebaseDeviceToken.Get()
}

// GetFirebaseDeviceTokenOk returns a tuple with the FirebaseDeviceToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirebaseDeviceDto) GetFirebaseDeviceTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirebaseDeviceToken.Get(), o.FirebaseDeviceToken.IsSet()
}

// HasFirebaseDeviceToken returns a boolean if a field has been set.
func (o *FirebaseDeviceDto) IsFirebaseDeviceTokenSet() bool {
	if o != nil && o.FirebaseDeviceToken.IsSet() {
		return true
	}

	return false
}

// SetFirebaseDeviceToken gets a reference to the given NullableString and assigns it to the FirebaseDeviceToken field.
func (o *FirebaseDeviceDto) SetFirebaseDeviceToken(v string) {
	o.FirebaseDeviceToken.Set(&v)
}
// SetFirebaseDeviceTokenNil sets the value for FirebaseDeviceToken to be an explicit nil
func (o *FirebaseDeviceDto) SetFirebaseDeviceTokenNil() {
	o.FirebaseDeviceToken.Set(nil)
}

// UnsetFirebaseDeviceToken ensures that no value is present for FirebaseDeviceToken, not even an explicit nil
func (o *FirebaseDeviceDto) UnsetFirebaseDeviceToken() {
	o.FirebaseDeviceToken.Unset()
}

// GetApplication returns the Application field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirebaseDeviceDto) GetApplication() string {
	if o == nil || IsNil(o.Application.Get()) {
		var ret string
		return ret
	}
	return *o.Application.Get()
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirebaseDeviceDto) GetApplicationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Application.Get(), o.Application.IsSet()
}

// HasApplication returns a boolean if a field has been set.
func (o *FirebaseDeviceDto) IsApplicationSet() bool {
	if o != nil && o.Application.IsSet() {
		return true
	}

	return false
}

// SetApplication gets a reference to the given NullableString and assigns it to the Application field.
func (o *FirebaseDeviceDto) SetApplication(v string) {
	o.Application.Set(&v)
}
// SetApplicationNil sets the value for Application to be an explicit nil
func (o *FirebaseDeviceDto) SetApplicationNil() {
	o.Application.Set(nil)
}

// UnsetApplication ensures that no value is present for Application, not even an explicit nil
func (o *FirebaseDeviceDto) UnsetApplication() {
	o.Application.Unset()
}

// GetIsSubscribed returns the IsSubscribed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirebaseDeviceDto) GetIsSubscribed() bool {
	if o == nil || IsNil(o.IsSubscribed.Get()) {
		var ret bool
		return ret
	}
	return *o.IsSubscribed.Get()
}

// GetIsSubscribedOk returns a tuple with the IsSubscribed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirebaseDeviceDto) GetIsSubscribedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsSubscribed.Get(), o.IsSubscribed.IsSet()
}

// HasIsSubscribed returns a boolean if a field has been set.
func (o *FirebaseDeviceDto) IsIsSubscribedSet() bool {
	if o != nil && o.IsSubscribed.IsSet() {
		return true
	}

	return false
}

// SetIsSubscribed gets a reference to the given NullableBool and assigns it to the IsSubscribed field.
func (o *FirebaseDeviceDto) SetIsSubscribed(v bool) {
	o.IsSubscribed.Set(&v)
}
// SetIsSubscribedNil sets the value for IsSubscribed to be an explicit nil
func (o *FirebaseDeviceDto) SetIsSubscribedNil() {
	o.IsSubscribed.Set(nil)
}

// UnsetIsSubscribed ensures that no value is present for IsSubscribed, not even an explicit nil
func (o *FirebaseDeviceDto) UnsetIsSubscribed() {
	o.IsSubscribed.Unset()
}

func (o FirebaseDeviceDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FirebaseDeviceDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if o.FirebaseDeviceToken.IsSet() {
		toSerialize["firebaseDeviceToken"] = o.FirebaseDeviceToken.Get()
	}
	if o.Application.IsSet() {
		toSerialize["application"] = o.Application.Get()
	}
	if o.IsSubscribed.IsSet() {
		toSerialize["isSubscribed"] = o.IsSubscribed.Get()
	}
	return toSerialize, nil
}

type NullableFirebaseDeviceDto struct {
	value *FirebaseDeviceDto
	isSet bool
}

func (v NullableFirebaseDeviceDto) Get() *FirebaseDeviceDto {
	return v.value
}

func (v *NullableFirebaseDeviceDto) Set(val *FirebaseDeviceDto) {
	v.value = val
	v.isSet = true
}

func (v NullableFirebaseDeviceDto) IsSet() bool {
	return v.isSet
}

func (v *NullableFirebaseDeviceDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirebaseDeviceDto(val *FirebaseDeviceDto) *NullableFirebaseDeviceDto {
	return &NullableFirebaseDeviceDto{value: val, isSet: true}
}

func (v NullableFirebaseDeviceDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirebaseDeviceDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

