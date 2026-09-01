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

// checks if the FireBaseUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FireBaseUser{}

// FireBaseUser The Firebase user parameters.
type FireBaseUser struct {
	// The Firebase user ID.
	Id *int32 `json:"id,omitempty"`
	// The user ID.
	UserId *string `json:"userId,omitempty"`
	// The tenant ID.
	TenantId *int32 `json:"tenantId,omitempty"`
	// The Firebase device token.
	FirebaseDeviceToken NullableString `json:"firebaseDeviceToken,omitempty"`
	// The Firebase application.
	Application NullableString `json:"application,omitempty"`
	// Specifies if the user is subscribed to the push notifications or not.
	IsSubscribed NullableBool `json:"isSubscribed,omitempty"`
	// The database tenant parameters.
	Tenant *DbTenant `json:"tenant,omitempty"`
}

// NewFireBaseUser instantiates a new FireBaseUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFireBaseUser() *FireBaseUser {
	this := FireBaseUser{}
	return &this
}

// NewFireBaseUserWithDefaults instantiates a new FireBaseUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFireBaseUserWithDefaults() *FireBaseUser {
	this := FireBaseUser{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FireBaseUser) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FireBaseUser) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FireBaseUser) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *FireBaseUser) SetId(v int32) {
	o.Id = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *FireBaseUser) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FireBaseUser) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *FireBaseUser) IsUserIdSet() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *FireBaseUser) SetUserId(v string) {
	o.UserId = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *FireBaseUser) GetTenantId() int32 {
	if o == nil || IsNil(o.TenantId) {
		var ret int32
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FireBaseUser) GetTenantIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *FireBaseUser) IsTenantIdSet() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given int32 and assigns it to the TenantId field.
func (o *FireBaseUser) SetTenantId(v int32) {
	o.TenantId = &v
}

// GetFirebaseDeviceToken returns the FirebaseDeviceToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FireBaseUser) GetFirebaseDeviceToken() string {
	if o == nil || IsNil(o.FirebaseDeviceToken.Get()) {
		var ret string
		return ret
	}
	return *o.FirebaseDeviceToken.Get()
}

// GetFirebaseDeviceTokenOk returns a tuple with the FirebaseDeviceToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FireBaseUser) GetFirebaseDeviceTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirebaseDeviceToken.Get(), o.FirebaseDeviceToken.IsSet()
}

// HasFirebaseDeviceToken returns a boolean if a field has been set.
func (o *FireBaseUser) IsFirebaseDeviceTokenSet() bool {
	if o != nil && o.FirebaseDeviceToken.IsSet() {
		return true
	}

	return false
}

// SetFirebaseDeviceToken gets a reference to the given NullableString and assigns it to the FirebaseDeviceToken field.
func (o *FireBaseUser) SetFirebaseDeviceToken(v string) {
	o.FirebaseDeviceToken.Set(&v)
}
// SetFirebaseDeviceTokenNil sets the value for FirebaseDeviceToken to be an explicit nil
func (o *FireBaseUser) SetFirebaseDeviceTokenNil() {
	o.FirebaseDeviceToken.Set(nil)
}

// UnsetFirebaseDeviceToken ensures that no value is present for FirebaseDeviceToken, not even an explicit nil
func (o *FireBaseUser) UnsetFirebaseDeviceToken() {
	o.FirebaseDeviceToken.Unset()
}

// GetApplication returns the Application field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FireBaseUser) GetApplication() string {
	if o == nil || IsNil(o.Application.Get()) {
		var ret string
		return ret
	}
	return *o.Application.Get()
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FireBaseUser) GetApplicationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Application.Get(), o.Application.IsSet()
}

// HasApplication returns a boolean if a field has been set.
func (o *FireBaseUser) IsApplicationSet() bool {
	if o != nil && o.Application.IsSet() {
		return true
	}

	return false
}

// SetApplication gets a reference to the given NullableString and assigns it to the Application field.
func (o *FireBaseUser) SetApplication(v string) {
	o.Application.Set(&v)
}
// SetApplicationNil sets the value for Application to be an explicit nil
func (o *FireBaseUser) SetApplicationNil() {
	o.Application.Set(nil)
}

// UnsetApplication ensures that no value is present for Application, not even an explicit nil
func (o *FireBaseUser) UnsetApplication() {
	o.Application.Unset()
}

// GetIsSubscribed returns the IsSubscribed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FireBaseUser) GetIsSubscribed() bool {
	if o == nil || IsNil(o.IsSubscribed.Get()) {
		var ret bool
		return ret
	}
	return *o.IsSubscribed.Get()
}

// GetIsSubscribedOk returns a tuple with the IsSubscribed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FireBaseUser) GetIsSubscribedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsSubscribed.Get(), o.IsSubscribed.IsSet()
}

// HasIsSubscribed returns a boolean if a field has been set.
func (o *FireBaseUser) IsIsSubscribedSet() bool {
	if o != nil && o.IsSubscribed.IsSet() {
		return true
	}

	return false
}

// SetIsSubscribed gets a reference to the given NullableBool and assigns it to the IsSubscribed field.
func (o *FireBaseUser) SetIsSubscribed(v bool) {
	o.IsSubscribed.Set(&v)
}
// SetIsSubscribedNil sets the value for IsSubscribed to be an explicit nil
func (o *FireBaseUser) SetIsSubscribedNil() {
	o.IsSubscribed.Set(nil)
}

// UnsetIsSubscribed ensures that no value is present for IsSubscribed, not even an explicit nil
func (o *FireBaseUser) UnsetIsSubscribed() {
	o.IsSubscribed.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *FireBaseUser) GetTenant() DbTenant {
	if o == nil || IsNil(o.Tenant) {
		var ret DbTenant
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FireBaseUser) GetTenantOk() (*DbTenant, bool) {
	if o == nil || IsNil(o.Tenant) {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *FireBaseUser) IsTenantSet() bool {
	if o != nil && !IsNil(o.Tenant) {
		return true
	}

	return false
}

// SetTenant gets a reference to the given DbTenant and assigns it to the Tenant field.
func (o *FireBaseUser) SetTenant(v DbTenant) {
	o.Tenant = &v
}

func (o FireBaseUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FireBaseUser) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Tenant) {
		toSerialize["tenant"] = o.Tenant
	}
	return toSerialize, nil
}

type NullableFireBaseUser struct {
	value *FireBaseUser
	isSet bool
}

func (v NullableFireBaseUser) Get() *FireBaseUser {
	return v.value
}

func (v *NullableFireBaseUser) Set(val *FireBaseUser) {
	v.value = val
	v.isSet = true
}

func (v NullableFireBaseUser) IsSet() bool {
	return v.isSet
}

func (v *NullableFireBaseUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFireBaseUser(val *FireBaseUser) *NullableFireBaseUser {
	return &NullableFireBaseUser{value: val, isSet: true}
}

func (v NullableFireBaseUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFireBaseUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

