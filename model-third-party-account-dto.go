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

// checks if the ThirdPartyAccountDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThirdPartyAccountDto{}

// ThirdPartyAccountDto A third-party storage account the caller has connected.
type ThirdPartyAccountDto struct {
	// Whether the account is attached to the legacy Common section, which is the case only for accounts inherited  from an older portal.
	Corporate *bool `json:"corporate,omitempty"`
	// Whether the account is attached to the Rooms section, room templates and the archive counted in. This is where  `POST api/2.0/files/thirdparty` puts every account it connects.
	RoomsStorage *bool `json:"roomsStorage,omitempty"`
	// The name the account is shown under in the portal, as it was saved when the account was connected.
	CustomerTitle NullableString `json:"customer_title,omitempty"`
	// The account ID to send to `DELETE api/2.0/files/thirdparty/{providerId}`, or as `providerId` to  re-authenticate the account.
	ProviderId NullableInt32 `json:"provider_id,omitempty"`
	// The storage service behind the account. `WebDav` stands for every WebDAV preset, so it does not tell which of  them was chosen when the account was connected.
	ProviderKey NullableString `json:"provider_key,omitempty"`
}

// NewThirdPartyAccountDto instantiates a new ThirdPartyAccountDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThirdPartyAccountDto() *ThirdPartyAccountDto {
	this := ThirdPartyAccountDto{}
	return &this
}

// NewThirdPartyAccountDtoWithDefaults instantiates a new ThirdPartyAccountDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThirdPartyAccountDtoWithDefaults() *ThirdPartyAccountDto {
	this := ThirdPartyAccountDto{}
	return &this
}

// GetCorporate returns the Corporate field value if set, zero value otherwise.
func (o *ThirdPartyAccountDto) GetCorporate() bool {
	if o == nil || IsNil(o.Corporate) {
		var ret bool
		return ret
	}
	return *o.Corporate
}

// GetCorporateOk returns a tuple with the Corporate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyAccountDto) GetCorporateOk() (*bool, bool) {
	if o == nil || IsNil(o.Corporate) {
		return nil, false
	}
	return o.Corporate, true
}

// HasCorporate returns a boolean if a field has been set.
func (o *ThirdPartyAccountDto) IsCorporateSet() bool {
	if o != nil && !IsNil(o.Corporate) {
		return true
	}

	return false
}

// SetCorporate gets a reference to the given bool and assigns it to the Corporate field.
func (o *ThirdPartyAccountDto) SetCorporate(v bool) {
	o.Corporate = &v
}

// GetRoomsStorage returns the RoomsStorage field value if set, zero value otherwise.
func (o *ThirdPartyAccountDto) GetRoomsStorage() bool {
	if o == nil || IsNil(o.RoomsStorage) {
		var ret bool
		return ret
	}
	return *o.RoomsStorage
}

// GetRoomsStorageOk returns a tuple with the RoomsStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyAccountDto) GetRoomsStorageOk() (*bool, bool) {
	if o == nil || IsNil(o.RoomsStorage) {
		return nil, false
	}
	return o.RoomsStorage, true
}

// HasRoomsStorage returns a boolean if a field has been set.
func (o *ThirdPartyAccountDto) IsRoomsStorageSet() bool {
	if o != nil && !IsNil(o.RoomsStorage) {
		return true
	}

	return false
}

// SetRoomsStorage gets a reference to the given bool and assigns it to the RoomsStorage field.
func (o *ThirdPartyAccountDto) SetRoomsStorage(v bool) {
	o.RoomsStorage = &v
}

// GetCustomerTitle returns the CustomerTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyAccountDto) GetCustomerTitle() string {
	if o == nil || IsNil(o.CustomerTitle.Get()) {
		var ret string
		return ret
	}
	return *o.CustomerTitle.Get()
}

// GetCustomerTitleOk returns a tuple with the CustomerTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyAccountDto) GetCustomerTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerTitle.Get(), o.CustomerTitle.IsSet()
}

// HasCustomerTitle returns a boolean if a field has been set.
func (o *ThirdPartyAccountDto) IsCustomerTitleSet() bool {
	if o != nil && o.CustomerTitle.IsSet() {
		return true
	}

	return false
}

// SetCustomerTitle gets a reference to the given NullableString and assigns it to the CustomerTitle field.
func (o *ThirdPartyAccountDto) SetCustomerTitle(v string) {
	o.CustomerTitle.Set(&v)
}
// SetCustomerTitleNil sets the value for CustomerTitle to be an explicit nil
func (o *ThirdPartyAccountDto) SetCustomerTitleNil() {
	o.CustomerTitle.Set(nil)
}

// UnsetCustomerTitle ensures that no value is present for CustomerTitle, not even an explicit nil
func (o *ThirdPartyAccountDto) UnsetCustomerTitle() {
	o.CustomerTitle.Unset()
}

// GetProviderId returns the ProviderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyAccountDto) GetProviderId() int32 {
	if o == nil || IsNil(o.ProviderId.Get()) {
		var ret int32
		return ret
	}
	return *o.ProviderId.Get()
}

// GetProviderIdOk returns a tuple with the ProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyAccountDto) GetProviderIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderId.Get(), o.ProviderId.IsSet()
}

// HasProviderId returns a boolean if a field has been set.
func (o *ThirdPartyAccountDto) IsProviderIdSet() bool {
	if o != nil && o.ProviderId.IsSet() {
		return true
	}

	return false
}

// SetProviderId gets a reference to the given NullableInt32 and assigns it to the ProviderId field.
func (o *ThirdPartyAccountDto) SetProviderId(v int32) {
	o.ProviderId.Set(&v)
}
// SetProviderIdNil sets the value for ProviderId to be an explicit nil
func (o *ThirdPartyAccountDto) SetProviderIdNil() {
	o.ProviderId.Set(nil)
}

// UnsetProviderId ensures that no value is present for ProviderId, not even an explicit nil
func (o *ThirdPartyAccountDto) UnsetProviderId() {
	o.ProviderId.Unset()
}

// GetProviderKey returns the ProviderKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyAccountDto) GetProviderKey() string {
	if o == nil || IsNil(o.ProviderKey.Get()) {
		var ret string
		return ret
	}
	return *o.ProviderKey.Get()
}

// GetProviderKeyOk returns a tuple with the ProviderKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyAccountDto) GetProviderKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderKey.Get(), o.ProviderKey.IsSet()
}

// HasProviderKey returns a boolean if a field has been set.
func (o *ThirdPartyAccountDto) IsProviderKeySet() bool {
	if o != nil && o.ProviderKey.IsSet() {
		return true
	}

	return false
}

// SetProviderKey gets a reference to the given NullableString and assigns it to the ProviderKey field.
func (o *ThirdPartyAccountDto) SetProviderKey(v string) {
	o.ProviderKey.Set(&v)
}
// SetProviderKeyNil sets the value for ProviderKey to be an explicit nil
func (o *ThirdPartyAccountDto) SetProviderKeyNil() {
	o.ProviderKey.Set(nil)
}

// UnsetProviderKey ensures that no value is present for ProviderKey, not even an explicit nil
func (o *ThirdPartyAccountDto) UnsetProviderKey() {
	o.ProviderKey.Unset()
}

func (o ThirdPartyAccountDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThirdPartyAccountDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Corporate) {
		toSerialize["corporate"] = o.Corporate
	}
	if !IsNil(o.RoomsStorage) {
		toSerialize["roomsStorage"] = o.RoomsStorage
	}
	if o.CustomerTitle.IsSet() {
		toSerialize["customer_title"] = o.CustomerTitle.Get()
	}
	if o.ProviderId.IsSet() {
		toSerialize["provider_id"] = o.ProviderId.Get()
	}
	if o.ProviderKey.IsSet() {
		toSerialize["provider_key"] = o.ProviderKey.Get()
	}
	return toSerialize, nil
}

type NullableThirdPartyAccountDto struct {
	value *ThirdPartyAccountDto
	isSet bool
}

func (v NullableThirdPartyAccountDto) Get() *ThirdPartyAccountDto {
	return v.value
}

func (v *NullableThirdPartyAccountDto) Set(val *ThirdPartyAccountDto) {
	v.value = val
	v.isSet = true
}

func (v NullableThirdPartyAccountDto) IsSet() bool {
	return v.isSet
}

func (v *NullableThirdPartyAccountDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThirdPartyAccountDto(val *ThirdPartyAccountDto) *NullableThirdPartyAccountDto {
	return &NullableThirdPartyAccountDto{value: val, isSet: true}
}

func (v NullableThirdPartyAccountDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThirdPartyAccountDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

