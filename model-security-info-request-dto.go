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

// checks if the SecurityInfoRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityInfoRequestDto{}

// SecurityInfoRequestDto The entries whose sharing rights are being changed, and the rights to apply to them.
type SecurityInfoRequestDto struct {
	// The folders and rooms whose rights are being changed, identified as a listing operation returns them - a  number on the portal, a string on a connected third-party account.
	FolderIds []DuplicateRequestDtoAllOfFileIds `json:"folderIds,omitempty"`
	// The files whose rights are being changed, identified as a listing operation returns them - a number on the  portal, a string on a connected third-party account.
	FileIds []DuplicateRequestDtoAllOfFileIds `json:"fileIds,omitempty"`
	// One record per account or group whose rights are being set, each naming the subject and the level it gets on  all of the listed entries; a level of `None` takes the access away. An empty collection makes the call change  nothing.
	Share []FileShareParams `json:"share,omitempty"`
	// Set to true to have every account named in `share` emailed about the access it just received; false changes  the rights without telling anyone.
	Notify *bool `json:"notify,omitempty"`
	// The text put into that email, ignored while `notify` is false. Markup is stripped before sending, so only the  plain text of the value survives.
	SharingMessage NullableString `json:"sharingMessage,omitempty"`
}

// NewSecurityInfoRequestDto instantiates a new SecurityInfoRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityInfoRequestDto() *SecurityInfoRequestDto {
	this := SecurityInfoRequestDto{}
	return &this
}

// NewSecurityInfoRequestDtoWithDefaults instantiates a new SecurityInfoRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityInfoRequestDtoWithDefaults() *SecurityInfoRequestDto {
	this := SecurityInfoRequestDto{}
	return &this
}

// GetFolderIds returns the FolderIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecurityInfoRequestDto) GetFolderIds() []DuplicateRequestDtoAllOfFileIds {
	if o == nil {
		var ret []DuplicateRequestDtoAllOfFileIds
		return ret
	}
	return o.FolderIds
}

// GetFolderIdsOk returns a tuple with the FolderIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecurityInfoRequestDto) GetFolderIdsOk() ([]DuplicateRequestDtoAllOfFileIds, bool) {
	if o == nil || IsNil(o.FolderIds) {
		return nil, false
	}
	return o.FolderIds, true
}

// HasFolderIds returns a boolean if a field has been set.
func (o *SecurityInfoRequestDto) IsFolderIdsSet() bool {
	if o != nil && !IsNil(o.FolderIds) {
		return true
	}

	return false
}

// SetFolderIds gets a reference to the given []DuplicateRequestDtoAllOfFileIds and assigns it to the FolderIds field.
func (o *SecurityInfoRequestDto) SetFolderIds(v []DuplicateRequestDtoAllOfFileIds) {
	o.FolderIds = v
}

// GetFileIds returns the FileIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecurityInfoRequestDto) GetFileIds() []DuplicateRequestDtoAllOfFileIds {
	if o == nil {
		var ret []DuplicateRequestDtoAllOfFileIds
		return ret
	}
	return o.FileIds
}

// GetFileIdsOk returns a tuple with the FileIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecurityInfoRequestDto) GetFileIdsOk() ([]DuplicateRequestDtoAllOfFileIds, bool) {
	if o == nil || IsNil(o.FileIds) {
		return nil, false
	}
	return o.FileIds, true
}

// HasFileIds returns a boolean if a field has been set.
func (o *SecurityInfoRequestDto) IsFileIdsSet() bool {
	if o != nil && !IsNil(o.FileIds) {
		return true
	}

	return false
}

// SetFileIds gets a reference to the given []DuplicateRequestDtoAllOfFileIds and assigns it to the FileIds field.
func (o *SecurityInfoRequestDto) SetFileIds(v []DuplicateRequestDtoAllOfFileIds) {
	o.FileIds = v
}

// GetShare returns the Share field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecurityInfoRequestDto) GetShare() []FileShareParams {
	if o == nil {
		var ret []FileShareParams
		return ret
	}
	return o.Share
}

// GetShareOk returns a tuple with the Share field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecurityInfoRequestDto) GetShareOk() ([]FileShareParams, bool) {
	if o == nil || IsNil(o.Share) {
		return nil, false
	}
	return o.Share, true
}

// HasShare returns a boolean if a field has been set.
func (o *SecurityInfoRequestDto) IsShareSet() bool {
	if o != nil && !IsNil(o.Share) {
		return true
	}

	return false
}

// SetShare gets a reference to the given []FileShareParams and assigns it to the Share field.
func (o *SecurityInfoRequestDto) SetShare(v []FileShareParams) {
	o.Share = v
}

// GetNotify returns the Notify field value if set, zero value otherwise.
func (o *SecurityInfoRequestDto) GetNotify() bool {
	if o == nil || IsNil(o.Notify) {
		var ret bool
		return ret
	}
	return *o.Notify
}

// GetNotifyOk returns a tuple with the Notify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityInfoRequestDto) GetNotifyOk() (*bool, bool) {
	if o == nil || IsNil(o.Notify) {
		return nil, false
	}
	return o.Notify, true
}

// HasNotify returns a boolean if a field has been set.
func (o *SecurityInfoRequestDto) IsNotifySet() bool {
	if o != nil && !IsNil(o.Notify) {
		return true
	}

	return false
}

// SetNotify gets a reference to the given bool and assigns it to the Notify field.
func (o *SecurityInfoRequestDto) SetNotify(v bool) {
	o.Notify = &v
}

// GetSharingMessage returns the SharingMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecurityInfoRequestDto) GetSharingMessage() string {
	if o == nil || IsNil(o.SharingMessage.Get()) {
		var ret string
		return ret
	}
	return *o.SharingMessage.Get()
}

// GetSharingMessageOk returns a tuple with the SharingMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SecurityInfoRequestDto) GetSharingMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SharingMessage.Get(), o.SharingMessage.IsSet()
}

// HasSharingMessage returns a boolean if a field has been set.
func (o *SecurityInfoRequestDto) IsSharingMessageSet() bool {
	if o != nil && o.SharingMessage.IsSet() {
		return true
	}

	return false
}

// SetSharingMessage gets a reference to the given NullableString and assigns it to the SharingMessage field.
func (o *SecurityInfoRequestDto) SetSharingMessage(v string) {
	o.SharingMessage.Set(&v)
}
// SetSharingMessageNil sets the value for SharingMessage to be an explicit nil
func (o *SecurityInfoRequestDto) SetSharingMessageNil() {
	o.SharingMessage.Set(nil)
}

// UnsetSharingMessage ensures that no value is present for SharingMessage, not even an explicit nil
func (o *SecurityInfoRequestDto) UnsetSharingMessage() {
	o.SharingMessage.Unset()
}

func (o SecurityInfoRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityInfoRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FolderIds != nil {
		toSerialize["folderIds"] = o.FolderIds
	}
	if o.FileIds != nil {
		toSerialize["fileIds"] = o.FileIds
	}
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

type NullableSecurityInfoRequestDto struct {
	value *SecurityInfoRequestDto
	isSet bool
}

func (v NullableSecurityInfoRequestDto) Get() *SecurityInfoRequestDto {
	return v.value
}

func (v *NullableSecurityInfoRequestDto) Set(val *SecurityInfoRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityInfoRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityInfoRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityInfoRequestDto(val *SecurityInfoRequestDto) *NullableSecurityInfoRequestDto {
	return &NullableSecurityInfoRequestDto{value: val, isSet: true}
}

func (v NullableSecurityInfoRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityInfoRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

