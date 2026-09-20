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

// checks if the AiFolderContentDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiFolderContentDto{}

// AiFolderContentDto One page of the contents of a folder or of a section: its entries split into files and folders, the folder itself,  and the counters needed to page through the rest.
type AiFolderContentDto struct {
	// The file entries of this page. It is empty when the folder holds no files, when the filters matched none of  them, and in the sections that list rooms only.
	Files []AiFileEntryBaseDto `json:"files,omitempty"`
	// The folder entries of this page. In a section of rooms these entries are the rooms themselves, which is where  their type, tags, logo and quota are read from.
	Folders []AiFileEntryBaseDto `json:"folders,omitempty"`
	// The folder or section the page was read from, with its own title, type and access rights. It describes the  container, not the entries, and is filled in even when the page is empty.
	Current *AiFolderDto `json:"current,omitempty"`
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

type _AiFolderContentDto AiFolderContentDto

// NewAiFolderContentDto instantiates a new AiFolderContentDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiFolderContentDto(pathParts interface{}, total int32) *AiFolderContentDto {
	this := AiFolderContentDto{}
	this.PathParts = pathParts
	this.Total = total
	return &this
}

// NewAiFolderContentDtoWithDefaults instantiates a new AiFolderContentDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiFolderContentDtoWithDefaults() *AiFolderContentDto {
	this := AiFolderContentDto{}
	return &this
}

// GetFiles returns the Files field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiFolderContentDto) GetFiles() []AiFileEntryBaseDto {
	if o == nil {
		var ret []AiFileEntryBaseDto
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiFolderContentDto) GetFilesOk() ([]AiFileEntryBaseDto, bool) {
	if o == nil || IsNil(o.Files) {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *AiFolderContentDto) IsFilesSet() bool {
	if o != nil && !IsNil(o.Files) {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []AiFileEntryBaseDto and assigns it to the Files field.
func (o *AiFolderContentDto) SetFiles(v []AiFileEntryBaseDto) {
	o.Files = v
}

// GetFolders returns the Folders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiFolderContentDto) GetFolders() []AiFileEntryBaseDto {
	if o == nil {
		var ret []AiFileEntryBaseDto
		return ret
	}
	return o.Folders
}

// GetFoldersOk returns a tuple with the Folders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiFolderContentDto) GetFoldersOk() ([]AiFileEntryBaseDto, bool) {
	if o == nil || IsNil(o.Folders) {
		return nil, false
	}
	return o.Folders, true
}

// HasFolders returns a boolean if a field has been set.
func (o *AiFolderContentDto) IsFoldersSet() bool {
	if o != nil && !IsNil(o.Folders) {
		return true
	}

	return false
}

// SetFolders gets a reference to the given []AiFileEntryBaseDto and assigns it to the Folders field.
func (o *AiFolderContentDto) SetFolders(v []AiFileEntryBaseDto) {
	o.Folders = v
}

// GetCurrent returns the Current field value if set, zero value otherwise.
func (o *AiFolderContentDto) GetCurrent() AiFolderDto {
	if o == nil || IsNil(o.Current) {
		var ret AiFolderDto
		return ret
	}
	return *o.Current
}

// GetCurrentOk returns a tuple with the Current field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderContentDto) GetCurrentOk() (*AiFolderDto, bool) {
	if o == nil || IsNil(o.Current) {
		return nil, false
	}
	return o.Current, true
}

// HasCurrent returns a boolean if a field has been set.
func (o *AiFolderContentDto) IsCurrentSet() bool {
	if o != nil && !IsNil(o.Current) {
		return true
	}

	return false
}

// SetCurrent gets a reference to the given AiFolderDto and assigns it to the Current field.
func (o *AiFolderContentDto) SetCurrent(v AiFolderDto) {
	o.Current = &v
}

// GetPathParts returns the PathParts field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *AiFolderContentDto) GetPathParts() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.PathParts
}

// GetPathPartsOk returns a tuple with the PathParts field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiFolderContentDto) GetPathPartsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PathParts) {
		return nil, false
	}
	return &o.PathParts, true
}

// SetPathParts sets field value
func (o *AiFolderContentDto) SetPathParts(v interface{}) {
	o.PathParts = v
}

// GetStartIndex returns the StartIndex field value if set, zero value otherwise.
func (o *AiFolderContentDto) GetStartIndex() int32 {
	if o == nil || IsNil(o.StartIndex) {
		var ret int32
		return ret
	}
	return *o.StartIndex
}

// GetStartIndexOk returns a tuple with the StartIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderContentDto) GetStartIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.StartIndex) {
		return nil, false
	}
	return o.StartIndex, true
}

// HasStartIndex returns a boolean if a field has been set.
func (o *AiFolderContentDto) IsStartIndexSet() bool {
	if o != nil && !IsNil(o.StartIndex) {
		return true
	}

	return false
}

// SetStartIndex gets a reference to the given int32 and assigns it to the StartIndex field.
func (o *AiFolderContentDto) SetStartIndex(v int32) {
	o.StartIndex = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *AiFolderContentDto) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderContentDto) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *AiFolderContentDto) IsCountSet() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *AiFolderContentDto) SetCount(v int32) {
	o.Count = &v
}

// GetTotal returns the Total field value
func (o *AiFolderContentDto) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *AiFolderContentDto) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *AiFolderContentDto) SetTotal(v int32) {
	o.Total = v
}

// GetNew returns the New field value if set, zero value otherwise.
func (o *AiFolderContentDto) GetNew() int32 {
	if o == nil || IsNil(o.New) {
		var ret int32
		return ret
	}
	return *o.New
}

// GetNewOk returns a tuple with the New field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFolderContentDto) GetNewOk() (*int32, bool) {
	if o == nil || IsNil(o.New) {
		return nil, false
	}
	return o.New, true
}

// HasNew returns a boolean if a field has been set.
func (o *AiFolderContentDto) IsNewSet() bool {
	if o != nil && !IsNil(o.New) {
		return true
	}

	return false
}

// SetNew gets a reference to the given int32 and assigns it to the New field.
func (o *AiFolderContentDto) SetNew(v int32) {
	o.New = &v
}

func (o AiFolderContentDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiFolderContentDto) ToMap() (map[string]interface{}, error) {
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

func (o *AiFolderContentDto) UnmarshalJSON(data []byte) (err error) {
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

	varAiFolderContentDto := _AiFolderContentDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiFolderContentDto)

	if err != nil {
		return err
	}

	*o = AiFolderContentDto(varAiFolderContentDto)

	return err
}

type NullableAiFolderContentDto struct {
	value *AiFolderContentDto
	isSet bool
}

func (v NullableAiFolderContentDto) Get() *AiFolderContentDto {
	return v.value
}

func (v *NullableAiFolderContentDto) Set(val *AiFolderContentDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAiFolderContentDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAiFolderContentDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiFolderContentDto(val *AiFolderContentDto) *NullableAiFolderContentDto {
	return &NullableAiFolderContentDto{value: val, isSet: true}
}

func (v NullableAiFolderContentDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiFolderContentDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

