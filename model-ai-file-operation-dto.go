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

// checks if the AiFileOperationDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiFileOperationDto{}

// AiFileOperationDto One background file operation of the caller, as it stood when the answer was built.
type AiFileOperationDto struct {
	// The identifier of the operation, the one to pass to `PUT api/2.0/files/fileops/terminate/{id}` to stop it.  Operations belong to the account that started them, so an identifier of somebody else is never listed here.
	Id NullableString `json:"id"`
	// What the operation does with the entries, which also decides what else is reported: only a download fills  `url`, and a deletion leaves `files` and `folders` empty.
	Operation AiFileOperationType `json:"Operation"`
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
	Files []AiFileEntryBaseDto `json:"files,omitempty"`
	// The folders the operation produced or moved, in the order it wrote them down. Empty while nothing has been  written yet and for a deletion.
	Folders []AiFileEntryBaseDto `json:"folders,omitempty"`
	// The state of the background task behind the operation, which tells a task that was cancelled or that crashed  from one that ran to its end.
	Status *AiDistributedTaskStatus `json:"status,omitempty"`
}

type _AiFileOperationDto AiFileOperationDto

// NewAiFileOperationDto instantiates a new AiFileOperationDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiFileOperationDto(id NullableString, operation AiFileOperationType, progress int32, error_ NullableString, processed NullableString, finished bool) *AiFileOperationDto {
	this := AiFileOperationDto{}
	this.Id = id
	this.Operation = operation
	this.Progress = progress
	this.Error = error_
	this.Processed = processed
	this.Finished = finished
	return &this
}

// NewAiFileOperationDtoWithDefaults instantiates a new AiFileOperationDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiFileOperationDtoWithDefaults() *AiFileOperationDto {
	this := AiFileOperationDto{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AiFileOperationDto) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiFileOperationDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *AiFileOperationDto) SetId(v string) {
	o.Id.Set(&v)
}

// GetOperation returns the Operation field value
func (o *AiFileOperationDto) GetOperation() AiFileOperationType {
	if o == nil {
		var ret AiFileOperationType
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *AiFileOperationDto) GetOperationOk() (*AiFileOperationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value
func (o *AiFileOperationDto) SetOperation(v AiFileOperationType) {
	o.Operation = v
}

// GetProgress returns the Progress field value
func (o *AiFileOperationDto) GetProgress() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Progress
}

// GetProgressOk returns a tuple with the Progress field value
// and a boolean to check if the value has been set.
func (o *AiFileOperationDto) GetProgressOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Progress, true
}

// SetProgress sets field value
func (o *AiFileOperationDto) SetProgress(v int32) {
	o.Progress = v
}

// GetError returns the Error field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AiFileOperationDto) GetError() string {
	if o == nil || o.Error.Get() == nil {
		var ret string
		return ret
	}

	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiFileOperationDto) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// SetError sets field value
func (o *AiFileOperationDto) SetError(v string) {
	o.Error.Set(&v)
}

// GetProcessed returns the Processed field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AiFileOperationDto) GetProcessed() string {
	if o == nil || o.Processed.Get() == nil {
		var ret string
		return ret
	}

	return *o.Processed.Get()
}

// GetProcessedOk returns a tuple with the Processed field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiFileOperationDto) GetProcessedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Processed.Get(), o.Processed.IsSet()
}

// SetProcessed sets field value
func (o *AiFileOperationDto) SetProcessed(v string) {
	o.Processed.Set(&v)
}

// GetFinished returns the Finished field value
func (o *AiFileOperationDto) GetFinished() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Finished
}

// GetFinishedOk returns a tuple with the Finished field value
// and a boolean to check if the value has been set.
func (o *AiFileOperationDto) GetFinishedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Finished, true
}

// SetFinished sets field value
func (o *AiFileOperationDto) SetFinished(v bool) {
	o.Finished = v
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiFileOperationDto) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiFileOperationDto) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *AiFileOperationDto) IsUrlSet() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *AiFileOperationDto) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *AiFileOperationDto) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *AiFileOperationDto) UnsetUrl() {
	o.Url.Unset()
}

// GetFiles returns the Files field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiFileOperationDto) GetFiles() []AiFileEntryBaseDto {
	if o == nil {
		var ret []AiFileEntryBaseDto
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiFileOperationDto) GetFilesOk() ([]AiFileEntryBaseDto, bool) {
	if o == nil || IsNil(o.Files) {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *AiFileOperationDto) IsFilesSet() bool {
	if o != nil && !IsNil(o.Files) {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []AiFileEntryBaseDto and assigns it to the Files field.
func (o *AiFileOperationDto) SetFiles(v []AiFileEntryBaseDto) {
	o.Files = v
}

// GetFolders returns the Folders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiFileOperationDto) GetFolders() []AiFileEntryBaseDto {
	if o == nil {
		var ret []AiFileEntryBaseDto
		return ret
	}
	return o.Folders
}

// GetFoldersOk returns a tuple with the Folders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiFileOperationDto) GetFoldersOk() ([]AiFileEntryBaseDto, bool) {
	if o == nil || IsNil(o.Folders) {
		return nil, false
	}
	return o.Folders, true
}

// HasFolders returns a boolean if a field has been set.
func (o *AiFileOperationDto) IsFoldersSet() bool {
	if o != nil && !IsNil(o.Folders) {
		return true
	}

	return false
}

// SetFolders gets a reference to the given []AiFileEntryBaseDto and assigns it to the Folders field.
func (o *AiFileOperationDto) SetFolders(v []AiFileEntryBaseDto) {
	o.Folders = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AiFileOperationDto) GetStatus() AiDistributedTaskStatus {
	if o == nil || IsNil(o.Status) {
		var ret AiDistributedTaskStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiFileOperationDto) GetStatusOk() (*AiDistributedTaskStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AiFileOperationDto) IsStatusSet() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given AiDistributedTaskStatus and assigns it to the Status field.
func (o *AiFileOperationDto) SetStatus(v AiDistributedTaskStatus) {
	o.Status = &v
}

func (o AiFileOperationDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiFileOperationDto) ToMap() (map[string]interface{}, error) {
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

func (o *AiFileOperationDto) UnmarshalJSON(data []byte) (err error) {
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

	varAiFileOperationDto := _AiFileOperationDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiFileOperationDto)

	if err != nil {
		return err
	}

	*o = AiFileOperationDto(varAiFileOperationDto)

	return err
}

type NullableAiFileOperationDto struct {
	value *AiFileOperationDto
	isSet bool
}

func (v NullableAiFileOperationDto) Get() *AiFileOperationDto {
	return v.value
}

func (v *NullableAiFileOperationDto) Set(val *AiFileOperationDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAiFileOperationDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAiFileOperationDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiFileOperationDto(val *AiFileOperationDto) *NullableAiFileOperationDto {
	return &NullableAiFileOperationDto{value: val, isSet: true}
}

func (v NullableAiFileOperationDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiFileOperationDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

