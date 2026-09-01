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

// checks if the ConversationResultDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConversationResultDto{}

// ConversationResultDto The result of file convertion operation.
type ConversationResultDto struct {
	// The conversion operation ID.
	Id NullableString `json:"id"`
	// The conversion operation type.
	Operation FileOperationType `json:"Operation"`
	// The conversion operation progress.
	Progress int32 `json:"progress"`
	// The source file for the conversion.
	Source NullableString `json:"source,omitempty"`
	Result interface{} `json:"result,omitempty"`
	// The conversion operation error message.
	Error NullableString `json:"error,omitempty"`
	// Specifies if the conversion operation is processed or not.
	Processed NullableString `json:"processed,omitempty"`
}

type _ConversationResultDto ConversationResultDto

// NewConversationResultDto instantiates a new ConversationResultDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConversationResultDto(id NullableString, operation FileOperationType, progress int32) *ConversationResultDto {
	this := ConversationResultDto{}
	this.Id = id
	this.Operation = operation
	this.Progress = progress
	return &this
}

// NewConversationResultDtoWithDefaults instantiates a new ConversationResultDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConversationResultDtoWithDefaults() *ConversationResultDto {
	this := ConversationResultDto{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ConversationResultDto) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConversationResultDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *ConversationResultDto) SetId(v string) {
	o.Id.Set(&v)
}

// GetOperation returns the Operation field value
func (o *ConversationResultDto) GetOperation() FileOperationType {
	if o == nil {
		var ret FileOperationType
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *ConversationResultDto) GetOperationOk() (*FileOperationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value
func (o *ConversationResultDto) SetOperation(v FileOperationType) {
	o.Operation = v
}

// GetProgress returns the Progress field value
func (o *ConversationResultDto) GetProgress() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Progress
}

// GetProgressOk returns a tuple with the Progress field value
// and a boolean to check if the value has been set.
func (o *ConversationResultDto) GetProgressOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Progress, true
}

// SetProgress sets field value
func (o *ConversationResultDto) SetProgress(v int32) {
	o.Progress = v
}

// GetSource returns the Source field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConversationResultDto) GetSource() string {
	if o == nil || IsNil(o.Source.Get()) {
		var ret string
		return ret
	}
	return *o.Source.Get()
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConversationResultDto) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Source.Get(), o.Source.IsSet()
}

// HasSource returns a boolean if a field has been set.
func (o *ConversationResultDto) IsSourceSet() bool {
	if o != nil && o.Source.IsSet() {
		return true
	}

	return false
}

// SetSource gets a reference to the given NullableString and assigns it to the Source field.
func (o *ConversationResultDto) SetSource(v string) {
	o.Source.Set(&v)
}
// SetSourceNil sets the value for Source to be an explicit nil
func (o *ConversationResultDto) SetSourceNil() {
	o.Source.Set(nil)
}

// UnsetSource ensures that no value is present for Source, not even an explicit nil
func (o *ConversationResultDto) UnsetSource() {
	o.Source.Unset()
}

// GetResult returns the Result field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConversationResultDto) GetResult() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConversationResultDto) GetResultOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return &o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ConversationResultDto) IsResultSet() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given interface{} and assigns it to the Result field.
func (o *ConversationResultDto) SetResult(v interface{}) {
	o.Result = v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConversationResultDto) GetError() string {
	if o == nil || IsNil(o.Error.Get()) {
		var ret string
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConversationResultDto) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *ConversationResultDto) IsErrorSet() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableString and assigns it to the Error field.
func (o *ConversationResultDto) SetError(v string) {
	o.Error.Set(&v)
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *ConversationResultDto) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *ConversationResultDto) UnsetError() {
	o.Error.Unset()
}

// GetProcessed returns the Processed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConversationResultDto) GetProcessed() string {
	if o == nil || IsNil(o.Processed.Get()) {
		var ret string
		return ret
	}
	return *o.Processed.Get()
}

// GetProcessedOk returns a tuple with the Processed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConversationResultDto) GetProcessedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Processed.Get(), o.Processed.IsSet()
}

// HasProcessed returns a boolean if a field has been set.
func (o *ConversationResultDto) IsProcessedSet() bool {
	if o != nil && o.Processed.IsSet() {
		return true
	}

	return false
}

// SetProcessed gets a reference to the given NullableString and assigns it to the Processed field.
func (o *ConversationResultDto) SetProcessed(v string) {
	o.Processed.Set(&v)
}
// SetProcessedNil sets the value for Processed to be an explicit nil
func (o *ConversationResultDto) SetProcessedNil() {
	o.Processed.Set(nil)
}

// UnsetProcessed ensures that no value is present for Processed, not even an explicit nil
func (o *ConversationResultDto) UnsetProcessed() {
	o.Processed.Unset()
}

func (o ConversationResultDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConversationResultDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id.Get()
	toSerialize["Operation"] = o.Operation
	toSerialize["progress"] = o.Progress
	if o.Source.IsSet() {
		toSerialize["source"] = o.Source.Get()
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	if o.Processed.IsSet() {
		toSerialize["processed"] = o.Processed.Get()
	}
	return toSerialize, nil
}

func (o *ConversationResultDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"Operation",
		"progress",
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

	varConversationResultDto := _ConversationResultDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varConversationResultDto)

	if err != nil {
		return err
	}

	*o = ConversationResultDto(varConversationResultDto)

	return err
}

type NullableConversationResultDto struct {
	value *ConversationResultDto
	isSet bool
}

func (v NullableConversationResultDto) Get() *ConversationResultDto {
	return v.value
}

func (v *NullableConversationResultDto) Set(val *ConversationResultDto) {
	v.value = val
	v.isSet = true
}

func (v NullableConversationResultDto) IsSet() bool {
	return v.isSet
}

func (v *NullableConversationResultDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConversationResultDto(val *ConversationResultDto) *NullableConversationResultDto {
	return &NullableConversationResultDto{value: val, isSet: true}
}

func (v NullableConversationResultDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConversationResultDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

