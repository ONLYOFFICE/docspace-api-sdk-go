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

// checks if the FilesStatisticsResultDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilesStatisticsResultDto{}

// FilesStatisticsResultDto The file statistics result parameters.
type FilesStatisticsResultDto struct {
	MyDocumentsUsedSpace *FilesStatisticsFolder `json:"myDocumentsUsedSpace,omitempty"`
	TrashUsedSpace *FilesStatisticsFolder `json:"trashUsedSpace,omitempty"`
	ArchiveUsedSpace *FilesStatisticsFolder `json:"archiveUsedSpace,omitempty"`
	RoomsUsedSpace *FilesStatisticsFolder `json:"roomsUsedSpace,omitempty"`
	AiAgentsUsedSpace *FilesStatisticsFolder `json:"aiAgentsUsedSpace,omitempty"`
}

// NewFilesStatisticsResultDto instantiates a new FilesStatisticsResultDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesStatisticsResultDto() *FilesStatisticsResultDto {
	this := FilesStatisticsResultDto{}
	return &this
}

// NewFilesStatisticsResultDtoWithDefaults instantiates a new FilesStatisticsResultDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesStatisticsResultDtoWithDefaults() *FilesStatisticsResultDto {
	this := FilesStatisticsResultDto{}
	return &this
}

// GetMyDocumentsUsedSpace returns the MyDocumentsUsedSpace field value if set, zero value otherwise.
func (o *FilesStatisticsResultDto) GetMyDocumentsUsedSpace() FilesStatisticsFolder {
	if o == nil || IsNil(o.MyDocumentsUsedSpace) {
		var ret FilesStatisticsFolder
		return ret
	}
	return *o.MyDocumentsUsedSpace
}

// GetMyDocumentsUsedSpaceOk returns a tuple with the MyDocumentsUsedSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesStatisticsResultDto) GetMyDocumentsUsedSpaceOk() (*FilesStatisticsFolder, bool) {
	if o == nil || IsNil(o.MyDocumentsUsedSpace) {
		return nil, false
	}
	return o.MyDocumentsUsedSpace, true
}

// HasMyDocumentsUsedSpace returns a boolean if a field has been set.
func (o *FilesStatisticsResultDto) IsMyDocumentsUsedSpaceSet() bool {
	if o != nil && !IsNil(o.MyDocumentsUsedSpace) {
		return true
	}

	return false
}

// SetMyDocumentsUsedSpace gets a reference to the given FilesStatisticsFolder and assigns it to the MyDocumentsUsedSpace field.
func (o *FilesStatisticsResultDto) SetMyDocumentsUsedSpace(v FilesStatisticsFolder) {
	o.MyDocumentsUsedSpace = &v
}

// GetTrashUsedSpace returns the TrashUsedSpace field value if set, zero value otherwise.
func (o *FilesStatisticsResultDto) GetTrashUsedSpace() FilesStatisticsFolder {
	if o == nil || IsNil(o.TrashUsedSpace) {
		var ret FilesStatisticsFolder
		return ret
	}
	return *o.TrashUsedSpace
}

// GetTrashUsedSpaceOk returns a tuple with the TrashUsedSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesStatisticsResultDto) GetTrashUsedSpaceOk() (*FilesStatisticsFolder, bool) {
	if o == nil || IsNil(o.TrashUsedSpace) {
		return nil, false
	}
	return o.TrashUsedSpace, true
}

// HasTrashUsedSpace returns a boolean if a field has been set.
func (o *FilesStatisticsResultDto) IsTrashUsedSpaceSet() bool {
	if o != nil && !IsNil(o.TrashUsedSpace) {
		return true
	}

	return false
}

// SetTrashUsedSpace gets a reference to the given FilesStatisticsFolder and assigns it to the TrashUsedSpace field.
func (o *FilesStatisticsResultDto) SetTrashUsedSpace(v FilesStatisticsFolder) {
	o.TrashUsedSpace = &v
}

// GetArchiveUsedSpace returns the ArchiveUsedSpace field value if set, zero value otherwise.
func (o *FilesStatisticsResultDto) GetArchiveUsedSpace() FilesStatisticsFolder {
	if o == nil || IsNil(o.ArchiveUsedSpace) {
		var ret FilesStatisticsFolder
		return ret
	}
	return *o.ArchiveUsedSpace
}

// GetArchiveUsedSpaceOk returns a tuple with the ArchiveUsedSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesStatisticsResultDto) GetArchiveUsedSpaceOk() (*FilesStatisticsFolder, bool) {
	if o == nil || IsNil(o.ArchiveUsedSpace) {
		return nil, false
	}
	return o.ArchiveUsedSpace, true
}

// HasArchiveUsedSpace returns a boolean if a field has been set.
func (o *FilesStatisticsResultDto) IsArchiveUsedSpaceSet() bool {
	if o != nil && !IsNil(o.ArchiveUsedSpace) {
		return true
	}

	return false
}

// SetArchiveUsedSpace gets a reference to the given FilesStatisticsFolder and assigns it to the ArchiveUsedSpace field.
func (o *FilesStatisticsResultDto) SetArchiveUsedSpace(v FilesStatisticsFolder) {
	o.ArchiveUsedSpace = &v
}

// GetRoomsUsedSpace returns the RoomsUsedSpace field value if set, zero value otherwise.
func (o *FilesStatisticsResultDto) GetRoomsUsedSpace() FilesStatisticsFolder {
	if o == nil || IsNil(o.RoomsUsedSpace) {
		var ret FilesStatisticsFolder
		return ret
	}
	return *o.RoomsUsedSpace
}

// GetRoomsUsedSpaceOk returns a tuple with the RoomsUsedSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesStatisticsResultDto) GetRoomsUsedSpaceOk() (*FilesStatisticsFolder, bool) {
	if o == nil || IsNil(o.RoomsUsedSpace) {
		return nil, false
	}
	return o.RoomsUsedSpace, true
}

// HasRoomsUsedSpace returns a boolean if a field has been set.
func (o *FilesStatisticsResultDto) IsRoomsUsedSpaceSet() bool {
	if o != nil && !IsNil(o.RoomsUsedSpace) {
		return true
	}

	return false
}

// SetRoomsUsedSpace gets a reference to the given FilesStatisticsFolder and assigns it to the RoomsUsedSpace field.
func (o *FilesStatisticsResultDto) SetRoomsUsedSpace(v FilesStatisticsFolder) {
	o.RoomsUsedSpace = &v
}

// GetAiAgentsUsedSpace returns the AiAgentsUsedSpace field value if set, zero value otherwise.
func (o *FilesStatisticsResultDto) GetAiAgentsUsedSpace() FilesStatisticsFolder {
	if o == nil || IsNil(o.AiAgentsUsedSpace) {
		var ret FilesStatisticsFolder
		return ret
	}
	return *o.AiAgentsUsedSpace
}

// GetAiAgentsUsedSpaceOk returns a tuple with the AiAgentsUsedSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesStatisticsResultDto) GetAiAgentsUsedSpaceOk() (*FilesStatisticsFolder, bool) {
	if o == nil || IsNil(o.AiAgentsUsedSpace) {
		return nil, false
	}
	return o.AiAgentsUsedSpace, true
}

// HasAiAgentsUsedSpace returns a boolean if a field has been set.
func (o *FilesStatisticsResultDto) IsAiAgentsUsedSpaceSet() bool {
	if o != nil && !IsNil(o.AiAgentsUsedSpace) {
		return true
	}

	return false
}

// SetAiAgentsUsedSpace gets a reference to the given FilesStatisticsFolder and assigns it to the AiAgentsUsedSpace field.
func (o *FilesStatisticsResultDto) SetAiAgentsUsedSpace(v FilesStatisticsFolder) {
	o.AiAgentsUsedSpace = &v
}

func (o FilesStatisticsResultDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilesStatisticsResultDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MyDocumentsUsedSpace) {
		toSerialize["myDocumentsUsedSpace"] = o.MyDocumentsUsedSpace
	}
	if !IsNil(o.TrashUsedSpace) {
		toSerialize["trashUsedSpace"] = o.TrashUsedSpace
	}
	if !IsNil(o.ArchiveUsedSpace) {
		toSerialize["archiveUsedSpace"] = o.ArchiveUsedSpace
	}
	if !IsNil(o.RoomsUsedSpace) {
		toSerialize["roomsUsedSpace"] = o.RoomsUsedSpace
	}
	if !IsNil(o.AiAgentsUsedSpace) {
		toSerialize["aiAgentsUsedSpace"] = o.AiAgentsUsedSpace
	}
	return toSerialize, nil
}

type NullableFilesStatisticsResultDto struct {
	value *FilesStatisticsResultDto
	isSet bool
}

func (v NullableFilesStatisticsResultDto) Get() *FilesStatisticsResultDto {
	return v.value
}

func (v *NullableFilesStatisticsResultDto) Set(val *FilesStatisticsResultDto) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesStatisticsResultDto) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesStatisticsResultDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesStatisticsResultDto(val *FilesStatisticsResultDto) *NullableFilesStatisticsResultDto {
	return &NullableFilesStatisticsResultDto{value: val, isSet: true}
}

func (v NullableFilesStatisticsResultDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesStatisticsResultDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

