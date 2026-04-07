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

// checks if the EditHistoryDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditHistoryDto{}

// EditHistoryDto The file editing history parameters.
type EditHistoryDto struct {
	// The document ID.
	Id *int32 `json:"id,omitempty"`
	// The document identifier used to unambiguously identify the document file.
	Key NullableString `json:"key,omitempty"`
	// The document version number.
	Version *int32 `json:"version,omitempty"`
	// The document version group.
	VersionGroup *int32 `json:"versionGroup,omitempty"`
	User *EditHistoryAuthor `json:"user,omitempty"`
	Created *ApiDateTime `json:"created,omitempty"`
	// The file history changes in the string format.
	ChangesHistory NullableString `json:"changesHistory,omitempty"`
	// The list of file history changes.
	Changes []EditHistoryChangesWrapper `json:"changes,omitempty"`
	// The current server version number.
	ServerVersion NullableString `json:"serverVersion,omitempty"`
}

// NewEditHistoryDto instantiates a new EditHistoryDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditHistoryDto() *EditHistoryDto {
	this := EditHistoryDto{}
	return &this
}

// NewEditHistoryDtoWithDefaults instantiates a new EditHistoryDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditHistoryDtoWithDefaults() *EditHistoryDto {
	this := EditHistoryDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EditHistoryDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditHistoryDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EditHistoryDto) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *EditHistoryDto) SetId(v int32) {
	o.Id = &v
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditHistoryDto) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditHistoryDto) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *EditHistoryDto) IsKeySet() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *EditHistoryDto) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *EditHistoryDto) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *EditHistoryDto) UnsetKey() {
	o.Key.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *EditHistoryDto) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditHistoryDto) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *EditHistoryDto) IsVersionSet() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *EditHistoryDto) SetVersion(v int32) {
	o.Version = &v
}

// GetVersionGroup returns the VersionGroup field value if set, zero value otherwise.
func (o *EditHistoryDto) GetVersionGroup() int32 {
	if o == nil || IsNil(o.VersionGroup) {
		var ret int32
		return ret
	}
	return *o.VersionGroup
}

// GetVersionGroupOk returns a tuple with the VersionGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditHistoryDto) GetVersionGroupOk() (*int32, bool) {
	if o == nil || IsNil(o.VersionGroup) {
		return nil, false
	}
	return o.VersionGroup, true
}

// HasVersionGroup returns a boolean if a field has been set.
func (o *EditHistoryDto) IsVersionGroupSet() bool {
	if o != nil && !IsNil(o.VersionGroup) {
		return true
	}

	return false
}

// SetVersionGroup gets a reference to the given int32 and assigns it to the VersionGroup field.
func (o *EditHistoryDto) SetVersionGroup(v int32) {
	o.VersionGroup = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *EditHistoryDto) GetUser() EditHistoryAuthor {
	if o == nil || IsNil(o.User) {
		var ret EditHistoryAuthor
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditHistoryDto) GetUserOk() (*EditHistoryAuthor, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *EditHistoryDto) IsUserSet() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given EditHistoryAuthor and assigns it to the User field.
func (o *EditHistoryDto) SetUser(v EditHistoryAuthor) {
	o.User = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *EditHistoryDto) GetCreated() ApiDateTime {
	if o == nil || IsNil(o.Created) {
		var ret ApiDateTime
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditHistoryDto) GetCreatedOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *EditHistoryDto) IsCreatedSet() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given ApiDateTime and assigns it to the Created field.
func (o *EditHistoryDto) SetCreated(v ApiDateTime) {
	o.Created = &v
}

// GetChangesHistory returns the ChangesHistory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditHistoryDto) GetChangesHistory() string {
	if o == nil || IsNil(o.ChangesHistory.Get()) {
		var ret string
		return ret
	}
	return *o.ChangesHistory.Get()
}

// GetChangesHistoryOk returns a tuple with the ChangesHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditHistoryDto) GetChangesHistoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChangesHistory.Get(), o.ChangesHistory.IsSet()
}

// HasChangesHistory returns a boolean if a field has been set.
func (o *EditHistoryDto) IsChangesHistorySet() bool {
	if o != nil && o.ChangesHistory.IsSet() {
		return true
	}

	return false
}

// SetChangesHistory gets a reference to the given NullableString and assigns it to the ChangesHistory field.
func (o *EditHistoryDto) SetChangesHistory(v string) {
	o.ChangesHistory.Set(&v)
}
// SetChangesHistoryNil sets the value for ChangesHistory to be an explicit nil
func (o *EditHistoryDto) SetChangesHistoryNil() {
	o.ChangesHistory.Set(nil)
}

// UnsetChangesHistory ensures that no value is present for ChangesHistory, not even an explicit nil
func (o *EditHistoryDto) UnsetChangesHistory() {
	o.ChangesHistory.Unset()
}

// GetChanges returns the Changes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditHistoryDto) GetChanges() []EditHistoryChangesWrapper {
	if o == nil {
		var ret []EditHistoryChangesWrapper
		return ret
	}
	return o.Changes
}

// GetChangesOk returns a tuple with the Changes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditHistoryDto) GetChangesOk() ([]EditHistoryChangesWrapper, bool) {
	if o == nil || IsNil(o.Changes) {
		return nil, false
	}
	return o.Changes, true
}

// HasChanges returns a boolean if a field has been set.
func (o *EditHistoryDto) IsChangesSet() bool {
	if o != nil && !IsNil(o.Changes) {
		return true
	}

	return false
}

// SetChanges gets a reference to the given []EditHistoryChangesWrapper and assigns it to the Changes field.
func (o *EditHistoryDto) SetChanges(v []EditHistoryChangesWrapper) {
	o.Changes = v
}

// GetServerVersion returns the ServerVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EditHistoryDto) GetServerVersion() string {
	if o == nil || IsNil(o.ServerVersion.Get()) {
		var ret string
		return ret
	}
	return *o.ServerVersion.Get()
}

// GetServerVersionOk returns a tuple with the ServerVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EditHistoryDto) GetServerVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServerVersion.Get(), o.ServerVersion.IsSet()
}

// HasServerVersion returns a boolean if a field has been set.
func (o *EditHistoryDto) IsServerVersionSet() bool {
	if o != nil && o.ServerVersion.IsSet() {
		return true
	}

	return false
}

// SetServerVersion gets a reference to the given NullableString and assigns it to the ServerVersion field.
func (o *EditHistoryDto) SetServerVersion(v string) {
	o.ServerVersion.Set(&v)
}
// SetServerVersionNil sets the value for ServerVersion to be an explicit nil
func (o *EditHistoryDto) SetServerVersionNil() {
	o.ServerVersion.Set(nil)
}

// UnsetServerVersion ensures that no value is present for ServerVersion, not even an explicit nil
func (o *EditHistoryDto) UnsetServerVersion() {
	o.ServerVersion.Unset()
}

func (o EditHistoryDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditHistoryDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.VersionGroup) {
		toSerialize["versionGroup"] = o.VersionGroup
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if o.ChangesHistory.IsSet() {
		toSerialize["changesHistory"] = o.ChangesHistory.Get()
	}
	if o.Changes != nil {
		toSerialize["changes"] = o.Changes
	}
	if o.ServerVersion.IsSet() {
		toSerialize["serverVersion"] = o.ServerVersion.Get()
	}
	return toSerialize, nil
}

type NullableEditHistoryDto struct {
	value *EditHistoryDto
	isSet bool
}

func (v NullableEditHistoryDto) Get() *EditHistoryDto {
	return v.value
}

func (v *NullableEditHistoryDto) Set(val *EditHistoryDto) {
	v.value = val
	v.isSet = true
}

func (v NullableEditHistoryDto) IsSet() bool {
	return v.isSet
}

func (v *NullableEditHistoryDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditHistoryDto(val *EditHistoryDto) *NullableEditHistoryDto {
	return &NullableEditHistoryDto{value: val, isSet: true}
}

func (v NullableEditHistoryDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditHistoryDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

