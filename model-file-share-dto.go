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

// checks if the FileShareDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileShareDto{}

// FileShareDto The file sharing information and access rights.
type FileShareDto struct {
	Access *FileShare `json:"access,omitempty"`
	// The user who has the access to the specified file.
	// Deprecated
	SharedTo interface{} `json:"sharedTo,omitempty"`
	SharedToUser *EmployeeFullDto `json:"sharedToUser,omitempty"`
	SharedToGroup *GroupSummaryDto `json:"sharedToGroup,omitempty"`
	SharedLink *FileShareLink `json:"sharedLink,omitempty"`
	// Specifies if the access right is locked or not.
	IsLocked bool `json:"isLocked"`
	// Specifies if the user is an owner of the specified file or not.
	IsOwner bool `json:"isOwner"`
	// Specifies if the user can edit the access to the specified file or not.
	CanEditAccess bool `json:"canEditAccess"`
	// Indicates whether internal editing permissions are granted.
	CanEditInternal bool `json:"canEditInternal"`
	// Determines whether the user has permission to modify the deny download setting for the file share.
	CanEditDenyDownload bool `json:"canEditDenyDownload"`
	// Indicates whether the expiration date of access permissions can be edited.
	CanEditExpirationDate bool `json:"canEditExpirationDate"`
	// Specifies whether the file sharing access can be revoked by the current user.
	CanRevoke bool `json:"canRevoke"`
	SubjectType SubjectType `json:"subjectType"`
}

type _FileShareDto FileShareDto

// NewFileShareDto instantiates a new FileShareDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileShareDto(isLocked bool, isOwner bool, canEditAccess bool, canEditInternal bool, canEditDenyDownload bool, canEditExpirationDate bool, canRevoke bool, subjectType SubjectType) *FileShareDto {
	this := FileShareDto{}
	this.IsLocked = isLocked
	this.IsOwner = isOwner
	this.CanEditAccess = canEditAccess
	this.CanEditInternal = canEditInternal
	this.CanEditDenyDownload = canEditDenyDownload
	this.CanEditExpirationDate = canEditExpirationDate
	this.CanRevoke = canRevoke
	this.SubjectType = subjectType
	return &this
}

// NewFileShareDtoWithDefaults instantiates a new FileShareDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileShareDtoWithDefaults() *FileShareDto {
	this := FileShareDto{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *FileShareDto) GetAccess() FileShare {
	if o == nil || IsNil(o.Access) {
		var ret FileShare
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileShareDto) GetAccessOk() (*FileShare, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *FileShareDto) IsAccessSet() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given FileShare and assigns it to the Access field.
func (o *FileShareDto) SetAccess(v FileShare) {
	o.Access = &v
}

// GetSharedTo returns the SharedTo field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *FileShareDto) GetSharedTo() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SharedTo
}

// GetSharedToOk returns a tuple with the SharedTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *FileShareDto) GetSharedToOk() (*interface{}, bool) {
	if o == nil || IsNil(o.SharedTo) {
		return nil, false
	}
	return &o.SharedTo, true
}

// HasSharedTo returns a boolean if a field has been set.
func (o *FileShareDto) IsSharedToSet() bool {
	if o != nil && !IsNil(o.SharedTo) {
		return true
	}

	return false
}

// SetSharedTo gets a reference to the given interface{} and assigns it to the SharedTo field.
// Deprecated
func (o *FileShareDto) SetSharedTo(v interface{}) {
	o.SharedTo = v
}

// GetSharedToUser returns the SharedToUser field value if set, zero value otherwise.
func (o *FileShareDto) GetSharedToUser() EmployeeFullDto {
	if o == nil || IsNil(o.SharedToUser) {
		var ret EmployeeFullDto
		return ret
	}
	return *o.SharedToUser
}

// GetSharedToUserOk returns a tuple with the SharedToUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileShareDto) GetSharedToUserOk() (*EmployeeFullDto, bool) {
	if o == nil || IsNil(o.SharedToUser) {
		return nil, false
	}
	return o.SharedToUser, true
}

// HasSharedToUser returns a boolean if a field has been set.
func (o *FileShareDto) IsSharedToUserSet() bool {
	if o != nil && !IsNil(o.SharedToUser) {
		return true
	}

	return false
}

// SetSharedToUser gets a reference to the given EmployeeFullDto and assigns it to the SharedToUser field.
func (o *FileShareDto) SetSharedToUser(v EmployeeFullDto) {
	o.SharedToUser = &v
}

// GetSharedToGroup returns the SharedToGroup field value if set, zero value otherwise.
func (o *FileShareDto) GetSharedToGroup() GroupSummaryDto {
	if o == nil || IsNil(o.SharedToGroup) {
		var ret GroupSummaryDto
		return ret
	}
	return *o.SharedToGroup
}

// GetSharedToGroupOk returns a tuple with the SharedToGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileShareDto) GetSharedToGroupOk() (*GroupSummaryDto, bool) {
	if o == nil || IsNil(o.SharedToGroup) {
		return nil, false
	}
	return o.SharedToGroup, true
}

// HasSharedToGroup returns a boolean if a field has been set.
func (o *FileShareDto) IsSharedToGroupSet() bool {
	if o != nil && !IsNil(o.SharedToGroup) {
		return true
	}

	return false
}

// SetSharedToGroup gets a reference to the given GroupSummaryDto and assigns it to the SharedToGroup field.
func (o *FileShareDto) SetSharedToGroup(v GroupSummaryDto) {
	o.SharedToGroup = &v
}

// GetSharedLink returns the SharedLink field value if set, zero value otherwise.
func (o *FileShareDto) GetSharedLink() FileShareLink {
	if o == nil || IsNil(o.SharedLink) {
		var ret FileShareLink
		return ret
	}
	return *o.SharedLink
}

// GetSharedLinkOk returns a tuple with the SharedLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileShareDto) GetSharedLinkOk() (*FileShareLink, bool) {
	if o == nil || IsNil(o.SharedLink) {
		return nil, false
	}
	return o.SharedLink, true
}

// HasSharedLink returns a boolean if a field has been set.
func (o *FileShareDto) IsSharedLinkSet() bool {
	if o != nil && !IsNil(o.SharedLink) {
		return true
	}

	return false
}

// SetSharedLink gets a reference to the given FileShareLink and assigns it to the SharedLink field.
func (o *FileShareDto) SetSharedLink(v FileShareLink) {
	o.SharedLink = &v
}

// GetIsLocked returns the IsLocked field value
func (o *FileShareDto) GetIsLocked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value
// and a boolean to check if the value has been set.
func (o *FileShareDto) GetIsLockedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsLocked, true
}

// SetIsLocked sets field value
func (o *FileShareDto) SetIsLocked(v bool) {
	o.IsLocked = v
}

// GetIsOwner returns the IsOwner field value
func (o *FileShareDto) GetIsOwner() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsOwner
}

// GetIsOwnerOk returns a tuple with the IsOwner field value
// and a boolean to check if the value has been set.
func (o *FileShareDto) GetIsOwnerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsOwner, true
}

// SetIsOwner sets field value
func (o *FileShareDto) SetIsOwner(v bool) {
	o.IsOwner = v
}

// GetCanEditAccess returns the CanEditAccess field value
func (o *FileShareDto) GetCanEditAccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanEditAccess
}

// GetCanEditAccessOk returns a tuple with the CanEditAccess field value
// and a boolean to check if the value has been set.
func (o *FileShareDto) GetCanEditAccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanEditAccess, true
}

// SetCanEditAccess sets field value
func (o *FileShareDto) SetCanEditAccess(v bool) {
	o.CanEditAccess = v
}

// GetCanEditInternal returns the CanEditInternal field value
func (o *FileShareDto) GetCanEditInternal() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanEditInternal
}

// GetCanEditInternalOk returns a tuple with the CanEditInternal field value
// and a boolean to check if the value has been set.
func (o *FileShareDto) GetCanEditInternalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanEditInternal, true
}

// SetCanEditInternal sets field value
func (o *FileShareDto) SetCanEditInternal(v bool) {
	o.CanEditInternal = v
}

// GetCanEditDenyDownload returns the CanEditDenyDownload field value
func (o *FileShareDto) GetCanEditDenyDownload() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanEditDenyDownload
}

// GetCanEditDenyDownloadOk returns a tuple with the CanEditDenyDownload field value
// and a boolean to check if the value has been set.
func (o *FileShareDto) GetCanEditDenyDownloadOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanEditDenyDownload, true
}

// SetCanEditDenyDownload sets field value
func (o *FileShareDto) SetCanEditDenyDownload(v bool) {
	o.CanEditDenyDownload = v
}

// GetCanEditExpirationDate returns the CanEditExpirationDate field value
func (o *FileShareDto) GetCanEditExpirationDate() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanEditExpirationDate
}

// GetCanEditExpirationDateOk returns a tuple with the CanEditExpirationDate field value
// and a boolean to check if the value has been set.
func (o *FileShareDto) GetCanEditExpirationDateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanEditExpirationDate, true
}

// SetCanEditExpirationDate sets field value
func (o *FileShareDto) SetCanEditExpirationDate(v bool) {
	o.CanEditExpirationDate = v
}

// GetCanRevoke returns the CanRevoke field value
func (o *FileShareDto) GetCanRevoke() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanRevoke
}

// GetCanRevokeOk returns a tuple with the CanRevoke field value
// and a boolean to check if the value has been set.
func (o *FileShareDto) GetCanRevokeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanRevoke, true
}

// SetCanRevoke sets field value
func (o *FileShareDto) SetCanRevoke(v bool) {
	o.CanRevoke = v
}

// GetSubjectType returns the SubjectType field value
func (o *FileShareDto) GetSubjectType() SubjectType {
	if o == nil {
		var ret SubjectType
		return ret
	}

	return o.SubjectType
}

// GetSubjectTypeOk returns a tuple with the SubjectType field value
// and a boolean to check if the value has been set.
func (o *FileShareDto) GetSubjectTypeOk() (*SubjectType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubjectType, true
}

// SetSubjectType sets field value
func (o *FileShareDto) SetSubjectType(v SubjectType) {
	o.SubjectType = v
}

func (o FileShareDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileShareDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	if o.SharedTo != nil {
		toSerialize["sharedTo"] = o.SharedTo
	}
	if !IsNil(o.SharedToUser) {
		toSerialize["sharedToUser"] = o.SharedToUser
	}
	if !IsNil(o.SharedToGroup) {
		toSerialize["sharedToGroup"] = o.SharedToGroup
	}
	if !IsNil(o.SharedLink) {
		toSerialize["sharedLink"] = o.SharedLink
	}
	toSerialize["isLocked"] = o.IsLocked
	toSerialize["isOwner"] = o.IsOwner
	toSerialize["canEditAccess"] = o.CanEditAccess
	toSerialize["canEditInternal"] = o.CanEditInternal
	toSerialize["canEditDenyDownload"] = o.CanEditDenyDownload
	toSerialize["canEditExpirationDate"] = o.CanEditExpirationDate
	toSerialize["canRevoke"] = o.CanRevoke
	toSerialize["subjectType"] = o.SubjectType
	return toSerialize, nil
}

func (o *FileShareDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"isLocked",
		"isOwner",
		"canEditAccess",
		"canEditInternal",
		"canEditDenyDownload",
		"canEditExpirationDate",
		"canRevoke",
		"subjectType",
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

	varFileShareDto := _FileShareDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varFileShareDto)

	if err != nil {
		return err
	}

	*o = FileShareDto(varFileShareDto)

	return err
}

type NullableFileShareDto struct {
	value *FileShareDto
	isSet bool
}

func (v NullableFileShareDto) Get() *FileShareDto {
	return v.value
}

func (v *NullableFileShareDto) Set(val *FileShareDto) {
	v.value = val
	v.isSet = true
}

func (v NullableFileShareDto) IsSet() bool {
	return v.isSet
}

func (v *NullableFileShareDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileShareDto(val *FileShareDto) *NullableFileShareDto {
	return &NullableFileShareDto{value: val, isSet: true}
}

func (v NullableFileShareDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileShareDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

