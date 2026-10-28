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

// checks if the FolderContentDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FolderContentDto{}

// FolderContentDto One page of the contents of a folder or of a section: its entries split into files and folders, the folder itself,  and the counters needed to page through the rest.
type FolderContentDto struct {
	// The file entries of this page. It is empty when the folder holds no files, when the filters matched none of  them, and in the sections that list rooms only.
	Files []FileEntryBaseDto `json:"files,omitempty"`
	// The folder entries of this page. In a section of rooms these entries are the rooms themselves, which is where  their type, tags, logo and quota are read from.
	Folders []FileEntryBaseDto `json:"folders,omitempty"`
	// The folder or section the page was read from, with its own title, type and access rights. It describes the  container, not the entries, and is filled in even when the page is empty.
	Current *FolderDto `json:"current,omitempty"`
	PathParts interface{} `json:"pathParts"`
	// The position of the first entry of this page in the whole result, echoing the requested start index. Add the  number of entries received to it to ask for the next page.
	StartIndex *int32 `json:"startIndex,omitempty"`
	// How many entries this page carries, files and folders together. A page shorter than the requested size means  the result is exhausted.
	Count *int32 `json:"count,omitempty"`
	// How many entries matched before paging was applied, across the whole folder. Page until the start index plus  the entries received reaches it.
	Total int32 `json:"total"`
	// How many entries of this folder are marked as new for the caller. It is 0 for every listing when the account  has switched the new-item badges off, so a zero here does not prove that nothing has changed.
	New *int32 `json:"new,omitempty"`
}

type _FolderContentDto FolderContentDto

// NewFolderContentDto instantiates a new FolderContentDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFolderContentDto(pathParts interface{}, total int32) *FolderContentDto {
	this := FolderContentDto{}
	this.PathParts = pathParts
	this.Total = total
	return &this
}

// NewFolderContentDtoWithDefaults instantiates a new FolderContentDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFolderContentDtoWithDefaults() *FolderContentDto {
	this := FolderContentDto{}
	return &this
}

// GetFiles returns the Files field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderContentDto) GetFiles() []FileEntryBaseDto {
	if o == nil {
		var ret []FileEntryBaseDto
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderContentDto) GetFilesOk() ([]FileEntryBaseDto, bool) {
	if o == nil || IsNil(o.Files) {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *FolderContentDto) IsFilesSet() bool {
	if o != nil && !IsNil(o.Files) {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []FileEntryBaseDto and assigns it to the Files field.
func (o *FolderContentDto) SetFiles(v []FileEntryBaseDto) {
	o.Files = v
}

// GetFolders returns the Folders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderContentDto) GetFolders() []FileEntryBaseDto {
	if o == nil {
		var ret []FileEntryBaseDto
		return ret
	}
	return o.Folders
}

// GetFoldersOk returns a tuple with the Folders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderContentDto) GetFoldersOk() ([]FileEntryBaseDto, bool) {
	if o == nil || IsNil(o.Folders) {
		return nil, false
	}
	return o.Folders, true
}

// HasFolders returns a boolean if a field has been set.
func (o *FolderContentDto) IsFoldersSet() bool {
	if o != nil && !IsNil(o.Folders) {
		return true
	}

	return false
}

// SetFolders gets a reference to the given []FileEntryBaseDto and assigns it to the Folders field.
func (o *FolderContentDto) SetFolders(v []FileEntryBaseDto) {
	o.Folders = v
}

// GetCurrent returns the Current field value if set, zero value otherwise.
func (o *FolderContentDto) GetCurrent() FolderDto {
	if o == nil || IsNil(o.Current) {
		var ret FolderDto
		return ret
	}
	return *o.Current
}

// GetCurrentOk returns a tuple with the Current field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderContentDto) GetCurrentOk() (*FolderDto, bool) {
	if o == nil || IsNil(o.Current) {
		return nil, false
	}
	return o.Current, true
}

// HasCurrent returns a boolean if a field has been set.
func (o *FolderContentDto) IsCurrentSet() bool {
	if o != nil && !IsNil(o.Current) {
		return true
	}

	return false
}

// SetCurrent gets a reference to the given FolderDto and assigns it to the Current field.
func (o *FolderContentDto) SetCurrent(v FolderDto) {
	o.Current = &v
}

// GetPathParts returns the PathParts field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *FolderContentDto) GetPathParts() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.PathParts
}

// GetPathPartsOk returns a tuple with the PathParts field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderContentDto) GetPathPartsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PathParts) {
		return nil, false
	}
	return &o.PathParts, true
}

// SetPathParts sets field value
func (o *FolderContentDto) SetPathParts(v interface{}) {
	o.PathParts = v
}

// GetStartIndex returns the StartIndex field value if set, zero value otherwise.
func (o *FolderContentDto) GetStartIndex() int32 {
	if o == nil || IsNil(o.StartIndex) {
		var ret int32
		return ret
	}
	return *o.StartIndex
}

// GetStartIndexOk returns a tuple with the StartIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderContentDto) GetStartIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.StartIndex) {
		return nil, false
	}
	return o.StartIndex, true
}

// HasStartIndex returns a boolean if a field has been set.
func (o *FolderContentDto) IsStartIndexSet() bool {
	if o != nil && !IsNil(o.StartIndex) {
		return true
	}

	return false
}

// SetStartIndex gets a reference to the given int32 and assigns it to the StartIndex field.
func (o *FolderContentDto) SetStartIndex(v int32) {
	o.StartIndex = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *FolderContentDto) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderContentDto) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *FolderContentDto) IsCountSet() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *FolderContentDto) SetCount(v int32) {
	o.Count = &v
}

// GetTotal returns the Total field value
func (o *FolderContentDto) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *FolderContentDto) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *FolderContentDto) SetTotal(v int32) {
	o.Total = v
}

// GetNew returns the New field value if set, zero value otherwise.
func (o *FolderContentDto) GetNew() int32 {
	if o == nil || IsNil(o.New) {
		var ret int32
		return ret
	}
	return *o.New
}

// GetNewOk returns a tuple with the New field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderContentDto) GetNewOk() (*int32, bool) {
	if o == nil || IsNil(o.New) {
		return nil, false
	}
	return o.New, true
}

// HasNew returns a boolean if a field has been set.
func (o *FolderContentDto) IsNewSet() bool {
	if o != nil && !IsNil(o.New) {
		return true
	}

	return false
}

// SetNew gets a reference to the given int32 and assigns it to the New field.
func (o *FolderContentDto) SetNew(v int32) {
	o.New = &v
}

func (o FolderContentDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FolderContentDto) ToMap() (map[string]interface{}, error) {
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

func (o *FolderContentDto) UnmarshalJSON(data []byte) (err error) {
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

	varFolderContentDto := _FolderContentDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varFolderContentDto)

	if err != nil {
		return err
	}

	*o = FolderContentDto(varFolderContentDto)

	return err
}

type NullableFolderContentDto struct {
	value *FolderContentDto
	isSet bool
}

func (v NullableFolderContentDto) Get() *FolderContentDto {
	return v.value
}

func (v *NullableFolderContentDto) Set(val *FolderContentDto) {
	v.value = val
	v.isSet = true
}

func (v NullableFolderContentDto) IsSet() bool {
	return v.isSet
}

func (v *NullableFolderContentDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFolderContentDto(val *FolderContentDto) *NullableFolderContentDto {
	return &NullableFolderContentDto{value: val, isSet: true}
}

func (v NullableFolderContentDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFolderContentDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

