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

// checks if the ErrorApiResponseError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorApiResponseError{}

// ErrorApiResponseError What went wrong.
type ErrorApiResponseError struct {
	// The human-readable error message.
	Message *string `json:"message,omitempty"`
	// The .NET type of the underlying exception. Only sent when stack traces are enabled.
	Type *string `json:"type,omitempty"`
	// The stack trace of the underlying exception. Only sent when stack traces are enabled.
	Stack *string `json:"stack,omitempty"`
	// The HRESULT of the underlying exception. Only sent when stack traces are enabled.
	Hresult *int32 `json:"hresult,omitempty"`
}

// NewErrorApiResponseError instantiates a new ErrorApiResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorApiResponseError() *ErrorApiResponseError {
	this := ErrorApiResponseError{}
	return &this
}

// NewErrorApiResponseErrorWithDefaults instantiates a new ErrorApiResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorApiResponseErrorWithDefaults() *ErrorApiResponseError {
	this := ErrorApiResponseError{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ErrorApiResponseError) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorApiResponseError) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ErrorApiResponseError) IsMessageSet() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ErrorApiResponseError) SetMessage(v string) {
	o.Message = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ErrorApiResponseError) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorApiResponseError) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ErrorApiResponseError) IsTypeSet() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ErrorApiResponseError) SetType(v string) {
	o.Type = &v
}

// GetStack returns the Stack field value if set, zero value otherwise.
func (o *ErrorApiResponseError) GetStack() string {
	if o == nil || IsNil(o.Stack) {
		var ret string
		return ret
	}
	return *o.Stack
}

// GetStackOk returns a tuple with the Stack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorApiResponseError) GetStackOk() (*string, bool) {
	if o == nil || IsNil(o.Stack) {
		return nil, false
	}
	return o.Stack, true
}

// HasStack returns a boolean if a field has been set.
func (o *ErrorApiResponseError) IsStackSet() bool {
	if o != nil && !IsNil(o.Stack) {
		return true
	}

	return false
}

// SetStack gets a reference to the given string and assigns it to the Stack field.
func (o *ErrorApiResponseError) SetStack(v string) {
	o.Stack = &v
}

// GetHresult returns the Hresult field value if set, zero value otherwise.
func (o *ErrorApiResponseError) GetHresult() int32 {
	if o == nil || IsNil(o.Hresult) {
		var ret int32
		return ret
	}
	return *o.Hresult
}

// GetHresultOk returns a tuple with the Hresult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorApiResponseError) GetHresultOk() (*int32, bool) {
	if o == nil || IsNil(o.Hresult) {
		return nil, false
	}
	return o.Hresult, true
}

// HasHresult returns a boolean if a field has been set.
func (o *ErrorApiResponseError) IsHresultSet() bool {
	if o != nil && !IsNil(o.Hresult) {
		return true
	}

	return false
}

// SetHresult gets a reference to the given int32 and assigns it to the Hresult field.
func (o *ErrorApiResponseError) SetHresult(v int32) {
	o.Hresult = &v
}

func (o ErrorApiResponseError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorApiResponseError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Stack) {
		toSerialize["stack"] = o.Stack
	}
	if !IsNil(o.Hresult) {
		toSerialize["hresult"] = o.Hresult
	}
	return toSerialize, nil
}

type NullableErrorApiResponseError struct {
	value *ErrorApiResponseError
	isSet bool
}

func (v NullableErrorApiResponseError) Get() *ErrorApiResponseError {
	return v.value
}

func (v *NullableErrorApiResponseError) Set(val *ErrorApiResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorApiResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorApiResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorApiResponseError(val *ErrorApiResponseError) *NullableErrorApiResponseError {
	return &NullableErrorApiResponseError{value: val, isSet: true}
}

func (v NullableErrorApiResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorApiResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

