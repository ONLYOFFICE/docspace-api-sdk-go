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

// checks if the FirebaseRequestsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FirebaseRequestsDto{}

// FirebaseRequestsDto The Firebase-related request parameters.
type FirebaseRequestsDto struct {
	// The Firebase device token.
	FirebaseDeviceToken NullableString `json:"firebaseDeviceToken,omitempty"`
	// Specifies whether the user is subscribed to the push notifications or not.
	IsSubscribed *bool `json:"isSubscribed,omitempty"`
}

// NewFirebaseRequestsDto instantiates a new FirebaseRequestsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirebaseRequestsDto() *FirebaseRequestsDto {
	this := FirebaseRequestsDto{}
	return &this
}

// NewFirebaseRequestsDtoWithDefaults instantiates a new FirebaseRequestsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirebaseRequestsDtoWithDefaults() *FirebaseRequestsDto {
	this := FirebaseRequestsDto{}
	return &this
}

// GetFirebaseDeviceToken returns the FirebaseDeviceToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirebaseRequestsDto) GetFirebaseDeviceToken() string {
	if o == nil || IsNil(o.FirebaseDeviceToken.Get()) {
		var ret string
		return ret
	}
	return *o.FirebaseDeviceToken.Get()
}

// GetFirebaseDeviceTokenOk returns a tuple with the FirebaseDeviceToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirebaseRequestsDto) GetFirebaseDeviceTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirebaseDeviceToken.Get(), o.FirebaseDeviceToken.IsSet()
}

// HasFirebaseDeviceToken returns a boolean if a field has been set.
func (o *FirebaseRequestsDto) IsFirebaseDeviceTokenSet() bool {
	if o != nil && o.FirebaseDeviceToken.IsSet() {
		return true
	}

	return false
}

// SetFirebaseDeviceToken gets a reference to the given NullableString and assigns it to the FirebaseDeviceToken field.
func (o *FirebaseRequestsDto) SetFirebaseDeviceToken(v string) {
	o.FirebaseDeviceToken.Set(&v)
}
// SetFirebaseDeviceTokenNil sets the value for FirebaseDeviceToken to be an explicit nil
func (o *FirebaseRequestsDto) SetFirebaseDeviceTokenNil() {
	o.FirebaseDeviceToken.Set(nil)
}

// UnsetFirebaseDeviceToken ensures that no value is present for FirebaseDeviceToken, not even an explicit nil
func (o *FirebaseRequestsDto) UnsetFirebaseDeviceToken() {
	o.FirebaseDeviceToken.Unset()
}

// GetIsSubscribed returns the IsSubscribed field value if set, zero value otherwise.
func (o *FirebaseRequestsDto) GetIsSubscribed() bool {
	if o == nil || IsNil(o.IsSubscribed) {
		var ret bool
		return ret
	}
	return *o.IsSubscribed
}

// GetIsSubscribedOk returns a tuple with the IsSubscribed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirebaseRequestsDto) GetIsSubscribedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSubscribed) {
		return nil, false
	}
	return o.IsSubscribed, true
}

// HasIsSubscribed returns a boolean if a field has been set.
func (o *FirebaseRequestsDto) IsIsSubscribedSet() bool {
	if o != nil && !IsNil(o.IsSubscribed) {
		return true
	}

	return false
}

// SetIsSubscribed gets a reference to the given bool and assigns it to the IsSubscribed field.
func (o *FirebaseRequestsDto) SetIsSubscribed(v bool) {
	o.IsSubscribed = &v
}

func (o FirebaseRequestsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FirebaseRequestsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FirebaseDeviceToken.IsSet() {
		toSerialize["firebaseDeviceToken"] = o.FirebaseDeviceToken.Get()
	}
	if !IsNil(o.IsSubscribed) {
		toSerialize["isSubscribed"] = o.IsSubscribed
	}
	return toSerialize, nil
}

type NullableFirebaseRequestsDto struct {
	value *FirebaseRequestsDto
	isSet bool
}

func (v NullableFirebaseRequestsDto) Get() *FirebaseRequestsDto {
	return v.value
}

func (v *NullableFirebaseRequestsDto) Set(val *FirebaseRequestsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableFirebaseRequestsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableFirebaseRequestsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirebaseRequestsDto(val *FirebaseRequestsDto) *NullableFirebaseRequestsDto {
	return &NullableFirebaseRequestsDto{value: val, isSet: true}
}

func (v NullableFirebaseRequestsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirebaseRequestsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

