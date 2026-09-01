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

// checks if the AiFolderContentDtoInteger type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiFolderContentDtoInteger{}

// AiFolderContentDtoInteger The folder content information.
type AiFolderContentDtoInteger struct {
	// The list of files in the folder.
	Files []AiFileEntryBaseDto `json:"files,omitempty"`
	// The list of folders in the folder.
	Folders []AiFileEntryBaseDto `json:"folders,omitempty"`
	// The current folder information.
	Current *AiFolderDtoInteger `json:"current,omitempty"`
	PathParts interface{} `json:"pathParts"`
	// The folder start index.
	StartIndex *int32 `json:"startIndex,omitempty"`
	// The number of folder elements.
	Count *int32 `json:"count,omitempty"`
	// The total number of elements in the folder.
	Total int32 `json:"total"`
	// The new element index in the folder.
	New *int32 `json:"new,omitempty"`
}

type _AiFolderContentDtoInteger AiFolderContentDtoInteger

// NewAiFolderContentDtoInteger instantiates a new AiFolderContentDtoInteger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiFolderContentDtoInteger(pathParts interface{}, total int32) *AiFolderContentDtoInteger {
	this := AiFolderContentDtoInteger{}
	this.PathParts = pathParts
	this.Total = total
	return &this
}

// NewAiFolderContentDtoIntegerWithDefaults instantiates a new AiFolderContentDtoInteger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiFolderContentDtoIntegerWithDefaults() *AiFolderContentDtoInteger {
	this := AiFolderContentDtoInteger{}
	return &this
}

// GetFiles returns the Files field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiFolderContentDtoInteger) GetFiles() []AiFileEntryBaseDto {
	if o == nil {
		var ret []AiFileEntryBaseDto
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiFolderContentDtoInteger) GetFilesOk() ([]AiFileEntryBaseDto, bool) {
	if o == nil || IsNil(o.Files) {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *AiFolderContentDtoInteger) IsFilesSet() bool {
	if o != nil && !IsNil(o.Files) {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []AiFileEntryBaseDto and assigns it to the Files field.
func (o *AiFolderContentDtoInteger) SetFiles(v []AiFileEntryBaseDto) {
	o.Files = v
}

// GetFolders returns the Folders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiFolderContentDtoInteger) GetFolders() []AiFileEntryBaseDto {
	if o == nil {
		var ret []AiFileEntryBaseDto
		return ret
	}
	return o.Folders
}

// GetFoldersOk returns a tuple with the Folders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiFolderContentDtoInteger) GetFoldersOk() ([]AiFileEntryBaseDto, bool) {
	if o == nil || IsNil(o.Folders) {
		return nil, false
	}
	return o.Folders, true
}

// HasFolders returns a boolean if a field has been set.
func (o *AiFolderContentDtoInteger) IsFoldersSet() bool {
	if o != nil && !IsNil(o.Folders) {
		return true
	}

	return false
}

// SetFolders gets a reference to the given []AiFileEntryBaseDto and assigns it to the Folders field.
func (o *AiFolderContentDtoInteger) SetFolders(v []AiFileEntryBaseDto) {
	o.Folders = v
}

// GetCurrent returns the Current field value if set, zero value otherwise.
func (o *AiFolderContentDtoInteger) GetCurrent() AiFolderDtoInteger {
	if o == nil || IsNil(o.Current) {
		var ret AiFolderDtoInteger
		return ret
	}
	return *o.Current
}

// GetCurrentOk returns a tuple with the Current field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderContentDtoInteger) GetCurrentOk() (*AiFolderDtoInteger, bool) {
	if o == nil || IsNil(o.Current) {
		return nil, false
	}
	return o.Current, true
}

// HasCurrent returns a boolean if a field has been set.
func (o *AiFolderContentDtoInteger) IsCurrentSet() bool {
	if o != nil && !IsNil(o.Current) {
		return true
	}

	return false
}

// SetCurrent gets a reference to the given AiFolderDtoInteger and assigns it to the Current field.
func (o *AiFolderContentDtoInteger) SetCurrent(v AiFolderDtoInteger) {
	o.Current = &v
}

// GetPathParts returns the PathParts field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *AiFolderContentDtoInteger) GetPathParts() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.PathParts
}

// GetPathPartsOk returns a tuple with the PathParts field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiFolderContentDtoInteger) GetPathPartsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PathParts) {
		return nil, false
	}
	return &o.PathParts, true
}

// SetPathParts sets field value
func (o *AiFolderContentDtoInteger) SetPathParts(v interface{}) {
	o.PathParts = v
}

// GetStartIndex returns the StartIndex field value if set, zero value otherwise.
func (o *AiFolderContentDtoInteger) GetStartIndex() int32 {
	if o == nil || IsNil(o.StartIndex) {
		var ret int32
		return ret
	}
	return *o.StartIndex
}

// GetStartIndexOk returns a tuple with the StartIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderContentDtoInteger) GetStartIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.StartIndex) {
		return nil, false
	}
	return o.StartIndex, true
}

// HasStartIndex returns a boolean if a field has been set.
func (o *AiFolderContentDtoInteger) IsStartIndexSet() bool {
	if o != nil && !IsNil(o.StartIndex) {
		return true
	}

	return false
}

// SetStartIndex gets a reference to the given int32 and assigns it to the StartIndex field.
func (o *AiFolderContentDtoInteger) SetStartIndex(v int32) {
	o.StartIndex = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *AiFolderContentDtoInteger) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderContentDtoInteger) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *AiFolderContentDtoInteger) IsCountSet() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *AiFolderContentDtoInteger) SetCount(v int32) {
	o.Count = &v
}

// GetTotal returns the Total field value
func (o *AiFolderContentDtoInteger) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *AiFolderContentDtoInteger) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *AiFolderContentDtoInteger) SetTotal(v int32) {
	o.Total = v
}

// GetNew returns the New field value if set, zero value otherwise.
func (o *AiFolderContentDtoInteger) GetNew() int32 {
	if o == nil || IsNil(o.New) {
		var ret int32
		return ret
	}
	return *o.New
}

// GetNewOk returns a tuple with the New field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderContentDtoInteger) GetNewOk() (*int32, bool) {
	if o == nil || IsNil(o.New) {
		return nil, false
	}
	return o.New, true
}

// HasNew returns a boolean if a field has been set.
func (o *AiFolderContentDtoInteger) IsNewSet() bool {
	if o != nil && !IsNil(o.New) {
		return true
	}

	return false
}

// SetNew gets a reference to the given int32 and assigns it to the New field.
func (o *AiFolderContentDtoInteger) SetNew(v int32) {
	o.New = &v
}

func (o AiFolderContentDtoInteger) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiFolderContentDtoInteger) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Files != nil {
		toSerialize["files"] = o.Files
	}
	if o.Folders != nil {
		toSerialize["folders"] = o.Folders
	}
	if !IsNil(o.Current) {
		toSerialize["current"] = o.Current
	}
	if o.PathParts != nil {
		toSerialize["pathParts"] = o.PathParts
	}
	if !IsNil(o.StartIndex) {
		toSerialize["startIndex"] = o.StartIndex
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	toSerialize["total"] = o.Total
	if !IsNil(o.New) {
		toSerialize["new"] = o.New
	}
	return toSerialize, nil
}

func (o *AiFolderContentDtoInteger) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pathParts",
		"total",
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

	varAiFolderContentDtoInteger := _AiFolderContentDtoInteger{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiFolderContentDtoInteger)

	if err != nil {
		return err
	}

	*o = AiFolderContentDtoInteger(varAiFolderContentDtoInteger)

	return err
}

type NullableAiFolderContentDtoInteger struct {
	value *AiFolderContentDtoInteger
	isSet bool
}

func (v NullableAiFolderContentDtoInteger) Get() *AiFolderContentDtoInteger {
	return v.value
}

func (v *NullableAiFolderContentDtoInteger) Set(val *AiFolderContentDtoInteger) {
	v.value = val
	v.isSet = true
}

func (v NullableAiFolderContentDtoInteger) IsSet() bool {
	return v.isSet
}

func (v *NullableAiFolderContentDtoInteger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiFolderContentDtoInteger(val *AiFolderContentDtoInteger) *NullableAiFolderContentDtoInteger {
	return &NullableAiFolderContentDtoInteger{value: val, isSet: true}
}

func (v NullableAiFolderContentDtoInteger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiFolderContentDtoInteger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

