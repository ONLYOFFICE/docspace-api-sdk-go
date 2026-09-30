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

// checks if the DocumentBuilderTaskDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentBuilderTaskDto{}

// DocumentBuilderTaskDto The state of a background document building task: how far it has got, how it ended, and the file it produced.
type DocumentBuilderTaskDto struct {
	// The identifier of the task. It is derived from the portal, the account and the kind of report, so starting the  same report again while it runs returns this same value, which is how a resumed poll is told from a newly  queued build.
	Id NullableString `json:"id"`
	// The message of the failure that stopped the build. It is filled in only for a task that ended in the failed  state, and stays empty while the task runs and after it succeeds.
	Error NullableString `json:"error"`
	// How far the build has got, from 0 to 100. It advances in a few coarse steps rather than smoothly, so it is a  progress hint and not a measure of the time left; wait on the completion flag instead.
	Percentage int32 `json:"percentage"`
	// True once the task has stopped for any reason, a failure and a cancellation included. It is the field to poll  on, and the status tells those outcomes apart.
	IsCompleted bool `json:"isCompleted"`
	// How the task ended, or that it has not started yet. Read it together with the completion flag: a stopped task  can be a finished build, a cancelled one or a failure, and only this field separates them.
	Status DistributedTaskStatus `json:"status"`
	ResultFileId interface{} `json:"resultFileId"`
	// The name the produced file was saved with, extension included. The name is built from the subject of the  report and is not unique: a second build adds another file instead of replacing the first.
	ResultFileName NullableString `json:"resultFileName"`
	// The address of the produced file in the document editor, relative to the portal root, so prefix it with the  portal address to open it. It stays empty until the build succeeds.
	ResultFileUrl NullableString `json:"resultFileUrl"`
}

type _DocumentBuilderTaskDto DocumentBuilderTaskDto

// NewDocumentBuilderTaskDto instantiates a new DocumentBuilderTaskDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentBuilderTaskDto(id NullableString, error_ NullableString, percentage int32, isCompleted bool, status DistributedTaskStatus, resultFileId interface{}, resultFileName NullableString, resultFileUrl NullableString) *DocumentBuilderTaskDto {
	this := DocumentBuilderTaskDto{}
	this.Id = id
	this.Error = error_
	this.Percentage = percentage
	this.IsCompleted = isCompleted
	this.Status = status
	this.ResultFileId = resultFileId
	this.ResultFileName = resultFileName
	this.ResultFileUrl = resultFileUrl
	return &this
}

// NewDocumentBuilderTaskDtoWithDefaults instantiates a new DocumentBuilderTaskDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentBuilderTaskDtoWithDefaults() *DocumentBuilderTaskDto {
	this := DocumentBuilderTaskDto{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DocumentBuilderTaskDto) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentBuilderTaskDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *DocumentBuilderTaskDto) SetId(v string) {
	o.Id.Set(&v)
}

// GetError returns the Error field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DocumentBuilderTaskDto) GetError() string {
	if o == nil || o.Error.Get() == nil {
		var ret string
		return ret
	}

	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentBuilderTaskDto) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// SetError sets field value
func (o *DocumentBuilderTaskDto) SetError(v string) {
	o.Error.Set(&v)
}

// GetPercentage returns the Percentage field value
func (o *DocumentBuilderTaskDto) GetPercentage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value
// and a boolean to check if the value has been set.
func (o *DocumentBuilderTaskDto) GetPercentageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Percentage, true
}

// SetPercentage sets field value
func (o *DocumentBuilderTaskDto) SetPercentage(v int32) {
	o.Percentage = v
}

// GetIsCompleted returns the IsCompleted field value
func (o *DocumentBuilderTaskDto) GetIsCompleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsCompleted
}

// GetIsCompletedOk returns a tuple with the IsCompleted field value
// and a boolean to check if the value has been set.
func (o *DocumentBuilderTaskDto) GetIsCompletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsCompleted, true
}

// SetIsCompleted sets field value
func (o *DocumentBuilderTaskDto) SetIsCompleted(v bool) {
	o.IsCompleted = v
}

// GetStatus returns the Status field value
func (o *DocumentBuilderTaskDto) GetStatus() DistributedTaskStatus {
	if o == nil {
		var ret DistributedTaskStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DocumentBuilderTaskDto) GetStatusOk() (*DistributedTaskStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DocumentBuilderTaskDto) SetStatus(v DistributedTaskStatus) {
	o.Status = v
}

// GetResultFileId returns the ResultFileId field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *DocumentBuilderTaskDto) GetResultFileId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.ResultFileId
}

// GetResultFileIdOk returns a tuple with the ResultFileId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentBuilderTaskDto) GetResultFileIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ResultFileId) {
		return nil, false
	}
	return &o.ResultFileId, true
}

// SetResultFileId sets field value
func (o *DocumentBuilderTaskDto) SetResultFileId(v interface{}) {
	o.ResultFileId = v
}

// GetResultFileName returns the ResultFileName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DocumentBuilderTaskDto) GetResultFileName() string {
	if o == nil || o.ResultFileName.Get() == nil {
		var ret string
		return ret
	}

	return *o.ResultFileName.Get()
}

// GetResultFileNameOk returns a tuple with the ResultFileName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentBuilderTaskDto) GetResultFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResultFileName.Get(), o.ResultFileName.IsSet()
}

// SetResultFileName sets field value
func (o *DocumentBuilderTaskDto) SetResultFileName(v string) {
	o.ResultFileName.Set(&v)
}

// GetResultFileUrl returns the ResultFileUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DocumentBuilderTaskDto) GetResultFileUrl() string {
	if o == nil || o.ResultFileUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.ResultFileUrl.Get()
}

// GetResultFileUrlOk returns a tuple with the ResultFileUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentBuilderTaskDto) GetResultFileUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResultFileUrl.Get(), o.ResultFileUrl.IsSet()
}

// SetResultFileUrl sets field value
func (o *DocumentBuilderTaskDto) SetResultFileUrl(v string) {
	o.ResultFileUrl.Set(&v)
}

func (o DocumentBuilderTaskDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentBuilderTaskDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id.Get()
	toSerialize["error"] = o.Error.Get()
	toSerialize["percentage"] = o.Percentage
	toSerialize["isCompleted"] = o.IsCompleted
	toSerialize["status"] = o.Status
	if o.ResultFileId != nil {
		toSerialize["resultFileId"] = o.ResultFileId
	}
	toSerialize["resultFileName"] = o.ResultFileName.Get()
	toSerialize["resultFileUrl"] = o.ResultFileUrl.Get()
	return toSerialize, nil
}

func (o *DocumentBuilderTaskDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"error",
		"percentage",
		"isCompleted",
		"status",
		"resultFileId",
		"resultFileName",
		"resultFileUrl",
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

	varDocumentBuilderTaskDto := _DocumentBuilderTaskDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varDocumentBuilderTaskDto)

	if err != nil {
		return err
	}

	*o = DocumentBuilderTaskDto(varDocumentBuilderTaskDto)

	return err
}

type NullableDocumentBuilderTaskDto struct {
	value *DocumentBuilderTaskDto
	isSet bool
}

func (v NullableDocumentBuilderTaskDto) Get() *DocumentBuilderTaskDto {
	return v.value
}

func (v *NullableDocumentBuilderTaskDto) Set(val *DocumentBuilderTaskDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentBuilderTaskDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentBuilderTaskDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentBuilderTaskDto(val *DocumentBuilderTaskDto) *NullableDocumentBuilderTaskDto {
	return &NullableDocumentBuilderTaskDto{value: val, isSet: true}
}

func (v NullableDocumentBuilderTaskDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentBuilderTaskDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

