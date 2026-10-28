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

// checks if the FileOperationDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileOperationDto{}

// FileOperationDto One background file operation of the caller, as it stood when the answer was built.
type FileOperationDto struct {
	// The identifier of the operation, the one to pass to `PUT api/2.0/files/fileops/terminate/{id}` to stop it.  Operations belong to the account that started them, so an identifier of somebody else is never listed here.
	Id NullableString `json:"id"`
	// What the operation does with the entries, which also decides what else is reported: only a download fills  `url`, and a deletion leaves `files` and `folders` empty.
	Operation FileOperationType `json:"Operation"`
	// How far the operation has come, from 0 to 100. Reaching 100 only means it stopped; whether it did what it was  asked for is told by `error`.
	Progress int32 `json:"progress"`
	// The reason the operation could not finish its work, in the language of the request. Empty when nothing went  wrong, which is the only way to tell a successful operation from a failed one.
	Error NullableString `json:"error"`
	// How many entries the operation has handled so far, written as a decimal number in a string. It counts items,  not percent, and stays behind `progress` on operations that walk into subfolders.
	Processed NullableString `json:"processed"`
	// Whether the operation has stopped running. A finished operation is reported once and then dropped, so the next  read of the operation list no longer contains it.
	Finished bool `json:"finished"`
	// The address the packed archive can be downloaded from once a bulk download has finished. Empty for every other  kind of operation.
	Url NullableString `json:"url,omitempty"`
	// The files the operation produced or moved, in the order it wrote them down. Empty while nothing has been  written yet and for a deletion, which reports no entries at all.
	Files []FileEntryBaseDto `json:"files,omitempty"`
	// The folders the operation produced or moved, in the order it wrote them down. Empty while nothing has been  written yet and for a deletion.
	Folders []FileEntryBaseDto `json:"folders,omitempty"`
	// The state of the background task behind the operation, which tells a task that was cancelled or that crashed  from one that ran to its end.
	Status *DistributedTaskStatus `json:"status,omitempty"`
}

type _FileOperationDto FileOperationDto

// NewFileOperationDto instantiates a new FileOperationDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileOperationDto(id NullableString, operation FileOperationType, progress int32, error_ NullableString, processed NullableString, finished bool) *FileOperationDto {
	this := FileOperationDto{}
	this.Id = id
	this.Operation = operation
	this.Progress = progress
	this.Error = error_
	this.Processed = processed
	this.Finished = finished
	return &this
}

// NewFileOperationDtoWithDefaults instantiates a new FileOperationDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileOperationDtoWithDefaults() *FileOperationDto {
	this := FileOperationDto{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FileOperationDto) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileOperationDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *FileOperationDto) SetId(v string) {
	o.Id.Set(&v)
}

// GetOperation returns the Operation field value
func (o *FileOperationDto) GetOperation() FileOperationType {
	if o == nil {
		var ret FileOperationType
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *FileOperationDto) GetOperationOk() (*FileOperationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value
func (o *FileOperationDto) SetOperation(v FileOperationType) {
	o.Operation = v
}

// GetProgress returns the Progress field value
func (o *FileOperationDto) GetProgress() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Progress
}

// GetProgressOk returns a tuple with the Progress field value
// and a boolean to check if the value has been set.
func (o *FileOperationDto) GetProgressOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Progress, true
}

// SetProgress sets field value
func (o *FileOperationDto) SetProgress(v int32) {
	o.Progress = v
}

// GetError returns the Error field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FileOperationDto) GetError() string {
	if o == nil || o.Error.Get() == nil {
		var ret string
		return ret
	}

	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileOperationDto) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// SetError sets field value
func (o *FileOperationDto) SetError(v string) {
	o.Error.Set(&v)
}

// GetProcessed returns the Processed field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FileOperationDto) GetProcessed() string {
	if o == nil || o.Processed.Get() == nil {
		var ret string
		return ret
	}

	return *o.Processed.Get()
}

// GetProcessedOk returns a tuple with the Processed field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileOperationDto) GetProcessedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Processed.Get(), o.Processed.IsSet()
}

// SetProcessed sets field value
func (o *FileOperationDto) SetProcessed(v string) {
	o.Processed.Set(&v)
}

// GetFinished returns the Finished field value
func (o *FileOperationDto) GetFinished() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Finished
}

// GetFinishedOk returns a tuple with the Finished field value
// and a boolean to check if the value has been set.
func (o *FileOperationDto) GetFinishedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Finished, true
}

// SetFinished sets field value
func (o *FileOperationDto) SetFinished(v bool) {
	o.Finished = v
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileOperationDto) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileOperationDto) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *FileOperationDto) IsUrlSet() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *FileOperationDto) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *FileOperationDto) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *FileOperationDto) UnsetUrl() {
	o.Url.Unset()
}

// GetFiles returns the Files field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileOperationDto) GetFiles() []FileEntryBaseDto {
	if o == nil {
		var ret []FileEntryBaseDto
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileOperationDto) GetFilesOk() ([]FileEntryBaseDto, bool) {
	if o == nil || IsNil(o.Files) {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *FileOperationDto) IsFilesSet() bool {
	if o != nil && !IsNil(o.Files) {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []FileEntryBaseDto and assigns it to the Files field.
func (o *FileOperationDto) SetFiles(v []FileEntryBaseDto) {
	o.Files = v
}

// GetFolders returns the Folders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileOperationDto) GetFolders() []FileEntryBaseDto {
	if o == nil {
		var ret []FileEntryBaseDto
		return ret
	}
	return o.Folders
}

// GetFoldersOk returns a tuple with the Folders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileOperationDto) GetFoldersOk() ([]FileEntryBaseDto, bool) {
	if o == nil || IsNil(o.Folders) {
		return nil, false
	}
	return o.Folders, true
}

// HasFolders returns a boolean if a field has been set.
func (o *FileOperationDto) IsFoldersSet() bool {
	if o != nil && !IsNil(o.Folders) {
		return true
	}

	return false
}

// SetFolders gets a reference to the given []FileEntryBaseDto and assigns it to the Folders field.
func (o *FileOperationDto) SetFolders(v []FileEntryBaseDto) {
	o.Folders = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FileOperationDto) GetStatus() DistributedTaskStatus {
	if o == nil || IsNil(o.Status) {
		var ret DistributedTaskStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileOperationDto) GetStatusOk() (*DistributedTaskStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FileOperationDto) IsStatusSet() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given DistributedTaskStatus and assigns it to the Status field.
func (o *FileOperationDto) SetStatus(v DistributedTaskStatus) {
	o.Status = &v
}

func (o FileOperationDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileOperationDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id.Get()
	toSerialize["Operation"] = o.Operation
	toSerialize["progress"] = o.Progress
	toSerialize["error"] = o.Error.Get()
	toSerialize["processed"] = o.Processed.Get()
	toSerialize["finished"] = o.Finished
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.Files != nil {
		toSerialize["files"] = o.Files
	}
	if o.Folders != nil {
		toSerialize["folders"] = o.Folders
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

func (o *FileOperationDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"Operation",
		"progress",
		"error",
		"processed",
		"finished",
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

	varFileOperationDto := _FileOperationDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varFileOperationDto)

	if err != nil {
		return err
	}

	*o = FileOperationDto(varFileOperationDto)

	return err
}

type NullableFileOperationDto struct {
	value *FileOperationDto
	isSet bool
}

func (v NullableFileOperationDto) Get() *FileOperationDto {
	return v.value
}

func (v *NullableFileOperationDto) Set(val *FileOperationDto) {
	v.value = val
	v.isSet = true
}

func (v NullableFileOperationDto) IsSet() bool {
	return v.isSet
}

func (v *NullableFileOperationDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileOperationDto(val *FileOperationDto) *NullableFileOperationDto {
	return &NullableFileOperationDto{value: val, isSet: true}
}

func (v NullableFileOperationDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileOperationDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

