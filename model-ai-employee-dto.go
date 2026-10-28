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

// checks if the AiEmployeeDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiEmployeeDto{}

// AiEmployeeDto The user parameters.
type AiEmployeeDto struct {
	// The user ID.
	Id *string `json:"id,omitempty"`
	// The HTML-encoded user's display name formatted according to the default format for the current culture.
	DisplayName NullableString `json:"displayName,omitempty"`
	// The user avatar.
	Avatar NullableString `json:"avatar,omitempty"`
	// The user original size avatar.
	AvatarOriginal NullableString `json:"avatarOriginal,omitempty"`
	// The user maximum size avatar.
	AvatarMax NullableString `json:"avatarMax,omitempty"`
	// The user medium size avatar.
	AvatarMedium NullableString `json:"avatarMedium,omitempty"`
	// The user small size avatar.
	AvatarSmall NullableString `json:"avatarSmall,omitempty"`
	// The user profile URL.
	ProfileUrl NullableString `json:"profileUrl,omitempty"`
	// Specifies if the user has an avatar or not.
	HasAvatar *bool `json:"hasAvatar,omitempty"`
	// Specifies if the user is anonymous or not.
	IsAnonim *bool `json:"isAnonim,omitempty"`
}

// NewAiEmployeeDto instantiates a new AiEmployeeDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiEmployeeDto() *AiEmployeeDto {
	this := AiEmployeeDto{}
	return &this
}

// NewAiEmployeeDtoWithDefaults instantiates a new AiEmployeeDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiEmployeeDtoWithDefaults() *AiEmployeeDto {
	this := AiEmployeeDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AiEmployeeDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiEmployeeDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AiEmployeeDto) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AiEmployeeDto) SetId(v string) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiEmployeeDto) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName.Get()) {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiEmployeeDto) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AiEmployeeDto) IsDisplayNameSet() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *AiEmployeeDto) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *AiEmployeeDto) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *AiEmployeeDto) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetAvatar returns the Avatar field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiEmployeeDto) GetAvatar() string {
	if o == nil || IsNil(o.Avatar.Get()) {
		var ret string
		return ret
	}
	return *o.Avatar.Get()
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiEmployeeDto) GetAvatarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Avatar.Get(), o.Avatar.IsSet()
}

// HasAvatar returns a boolean if a field has been set.
func (o *AiEmployeeDto) IsAvatarSet() bool {
	if o != nil && o.Avatar.IsSet() {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given NullableString and assigns it to the Avatar field.
func (o *AiEmployeeDto) SetAvatar(v string) {
	o.Avatar.Set(&v)
}
// SetAvatarNil sets the value for Avatar to be an explicit nil
func (o *AiEmployeeDto) SetAvatarNil() {
	o.Avatar.Set(nil)
}

// UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
func (o *AiEmployeeDto) UnsetAvatar() {
	o.Avatar.Unset()
}

// GetAvatarOriginal returns the AvatarOriginal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiEmployeeDto) GetAvatarOriginal() string {
	if o == nil || IsNil(o.AvatarOriginal.Get()) {
		var ret string
		return ret
	}
	return *o.AvatarOriginal.Get()
}

// GetAvatarOriginalOk returns a tuple with the AvatarOriginal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiEmployeeDto) GetAvatarOriginalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvatarOriginal.Get(), o.AvatarOriginal.IsSet()
}

// HasAvatarOriginal returns a boolean if a field has been set.
func (o *AiEmployeeDto) IsAvatarOriginalSet() bool {
	if o != nil && o.AvatarOriginal.IsSet() {
		return true
	}

	return false
}

// SetAvatarOriginal gets a reference to the given NullableString and assigns it to the AvatarOriginal field.
func (o *AiEmployeeDto) SetAvatarOriginal(v string) {
	o.AvatarOriginal.Set(&v)
}
// SetAvatarOriginalNil sets the value for AvatarOriginal to be an explicit nil
func (o *AiEmployeeDto) SetAvatarOriginalNil() {
	o.AvatarOriginal.Set(nil)
}

// UnsetAvatarOriginal ensures that no value is present for AvatarOriginal, not even an explicit nil
func (o *AiEmployeeDto) UnsetAvatarOriginal() {
	o.AvatarOriginal.Unset()
}

// GetAvatarMax returns the AvatarMax field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiEmployeeDto) GetAvatarMax() string {
	if o == nil || IsNil(o.AvatarMax.Get()) {
		var ret string
		return ret
	}
	return *o.AvatarMax.Get()
}

// GetAvatarMaxOk returns a tuple with the AvatarMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiEmployeeDto) GetAvatarMaxOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvatarMax.Get(), o.AvatarMax.IsSet()
}

// HasAvatarMax returns a boolean if a field has been set.
func (o *AiEmployeeDto) IsAvatarMaxSet() bool {
	if o != nil && o.AvatarMax.IsSet() {
		return true
	}

	return false
}

// SetAvatarMax gets a reference to the given NullableString and assigns it to the AvatarMax field.
func (o *AiEmployeeDto) SetAvatarMax(v string) {
	o.AvatarMax.Set(&v)
}
// SetAvatarMaxNil sets the value for AvatarMax to be an explicit nil
func (o *AiEmployeeDto) SetAvatarMaxNil() {
	o.AvatarMax.Set(nil)
}

// UnsetAvatarMax ensures that no value is present for AvatarMax, not even an explicit nil
func (o *AiEmployeeDto) UnsetAvatarMax() {
	o.AvatarMax.Unset()
}

// GetAvatarMedium returns the AvatarMedium field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiEmployeeDto) GetAvatarMedium() string {
	if o == nil || IsNil(o.AvatarMedium.Get()) {
		var ret string
		return ret
	}
	return *o.AvatarMedium.Get()
}

// GetAvatarMediumOk returns a tuple with the AvatarMedium field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiEmployeeDto) GetAvatarMediumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvatarMedium.Get(), o.AvatarMedium.IsSet()
}

// HasAvatarMedium returns a boolean if a field has been set.
func (o *AiEmployeeDto) IsAvatarMediumSet() bool {
	if o != nil && o.AvatarMedium.IsSet() {
		return true
	}

	return false
}

// SetAvatarMedium gets a reference to the given NullableString and assigns it to the AvatarMedium field.
func (o *AiEmployeeDto) SetAvatarMedium(v string) {
	o.AvatarMedium.Set(&v)
}
// SetAvatarMediumNil sets the value for AvatarMedium to be an explicit nil
func (o *AiEmployeeDto) SetAvatarMediumNil() {
	o.AvatarMedium.Set(nil)
}

// UnsetAvatarMedium ensures that no value is present for AvatarMedium, not even an explicit nil
func (o *AiEmployeeDto) UnsetAvatarMedium() {
	o.AvatarMedium.Unset()
}

// GetAvatarSmall returns the AvatarSmall field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiEmployeeDto) GetAvatarSmall() string {
	if o == nil || IsNil(o.AvatarSmall.Get()) {
		var ret string
		return ret
	}
	return *o.AvatarSmall.Get()
}

// GetAvatarSmallOk returns a tuple with the AvatarSmall field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiEmployeeDto) GetAvatarSmallOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvatarSmall.Get(), o.AvatarSmall.IsSet()
}

// HasAvatarSmall returns a boolean if a field has been set.
func (o *AiEmployeeDto) IsAvatarSmallSet() bool {
	if o != nil && o.AvatarSmall.IsSet() {
		return true
	}

	return false
}

// SetAvatarSmall gets a reference to the given NullableString and assigns it to the AvatarSmall field.
func (o *AiEmployeeDto) SetAvatarSmall(v string) {
	o.AvatarSmall.Set(&v)
}
// SetAvatarSmallNil sets the value for AvatarSmall to be an explicit nil
func (o *AiEmployeeDto) SetAvatarSmallNil() {
	o.AvatarSmall.Set(nil)
}

// UnsetAvatarSmall ensures that no value is present for AvatarSmall, not even an explicit nil
func (o *AiEmployeeDto) UnsetAvatarSmall() {
	o.AvatarSmall.Unset()
}

// GetProfileUrl returns the ProfileUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiEmployeeDto) GetProfileUrl() string {
	if o == nil || IsNil(o.ProfileUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ProfileUrl.Get()
}

// GetProfileUrlOk returns a tuple with the ProfileUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiEmployeeDto) GetProfileUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfileUrl.Get(), o.ProfileUrl.IsSet()
}

// HasProfileUrl returns a boolean if a field has been set.
func (o *AiEmployeeDto) IsProfileUrlSet() bool {
	if o != nil && o.ProfileUrl.IsSet() {
		return true
	}

	return false
}

// SetProfileUrl gets a reference to the given NullableString and assigns it to the ProfileUrl field.
func (o *AiEmployeeDto) SetProfileUrl(v string) {
	o.ProfileUrl.Set(&v)
}
// SetProfileUrlNil sets the value for ProfileUrl to be an explicit nil
func (o *AiEmployeeDto) SetProfileUrlNil() {
	o.ProfileUrl.Set(nil)
}

// UnsetProfileUrl ensures that no value is present for ProfileUrl, not even an explicit nil
func (o *AiEmployeeDto) UnsetProfileUrl() {
	o.ProfileUrl.Unset()
}

// GetHasAvatar returns the HasAvatar field value if set, zero value otherwise.
func (o *AiEmployeeDto) GetHasAvatar() bool {
	if o == nil || IsNil(o.HasAvatar) {
		var ret bool
		return ret
	}
	return *o.HasAvatar
}

// GetHasAvatarOk returns a tuple with the HasAvatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiEmployeeDto) GetHasAvatarOk() (*bool, bool) {
	if o == nil || IsNil(o.HasAvatar) {
		return nil, false
	}
	return o.HasAvatar, true
}

// HasHasAvatar returns a boolean if a field has been set.
func (o *AiEmployeeDto) IsHasAvatarSet() bool {
	if o != nil && !IsNil(o.HasAvatar) {
		return true
	}

	return false
}

// SetHasAvatar gets a reference to the given bool and assigns it to the HasAvatar field.
func (o *AiEmployeeDto) SetHasAvatar(v bool) {
	o.HasAvatar = &v
}

// GetIsAnonim returns the IsAnonim field value if set, zero value otherwise.
func (o *AiEmployeeDto) GetIsAnonim() bool {
	if o == nil || IsNil(o.IsAnonim) {
		var ret bool
		return ret
	}
	return *o.IsAnonim
}

// GetIsAnonimOk returns a tuple with the IsAnonim field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiEmployeeDto) GetIsAnonimOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAnonim) {
		return nil, false
	}
	return o.IsAnonim, true
}

// HasIsAnonim returns a boolean if a field has been set.
func (o *AiEmployeeDto) IsIsAnonimSet() bool {
	if o != nil && !IsNil(o.IsAnonim) {
		return true
	}

	return false
}

// SetIsAnonim gets a reference to the given bool and assigns it to the IsAnonim field.
func (o *AiEmployeeDto) SetIsAnonim(v bool) {
	o.IsAnonim = &v
}

func (o AiEmployeeDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiEmployeeDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.Avatar.IsSet() {
		toSerialize["avatar"] = o.Avatar.Get()
	}
	if o.AvatarOriginal.IsSet() {
		toSerialize["avatarOriginal"] = o.AvatarOriginal.Get()
	}
	if o.AvatarMax.IsSet() {
		toSerialize["avatarMax"] = o.AvatarMax.Get()
	}
	if o.AvatarMedium.IsSet() {
		toSerialize["avatarMedium"] = o.AvatarMedium.Get()
	}
	if o.AvatarSmall.IsSet() {
		toSerialize["avatarSmall"] = o.AvatarSmall.Get()
	}
	if o.ProfileUrl.IsSet() {
		toSerialize["profileUrl"] = o.ProfileUrl.Get()
	}
	if !IsNil(o.HasAvatar) {
		toSerialize["hasAvatar"] = o.HasAvatar
	}
	if !IsNil(o.IsAnonim) {
		toSerialize["isAnonim"] = o.IsAnonim
	}
	return toSerialize, nil
}

type NullableAiEmployeeDto struct {
	value *AiEmployeeDto
	isSet bool
}

func (v NullableAiEmployeeDto) Get() *AiEmployeeDto {
	return v.value
}

func (v *NullableAiEmployeeDto) Set(val *AiEmployeeDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAiEmployeeDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAiEmployeeDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiEmployeeDto(val *AiEmployeeDto) *NullableAiEmployeeDto {
	return &NullableAiEmployeeDto{value: val, isSet: true}
}

func (v NullableAiEmployeeDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiEmployeeDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

