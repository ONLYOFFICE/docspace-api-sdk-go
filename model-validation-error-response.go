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

// checks if the ValidationErrorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidationErrorResponse{}

// ValidationErrorResponse Response containing validation errors
type ValidationErrorResponse struct {
	// Error type identifier
	Error *string `json:"error,omitempty"`
	// General error message
	Message *string `json:"message,omitempty"`
	// List of field specific validation errors
	Errors []FieldError `json:"errors,omitempty"`
}

// NewValidationErrorResponse instantiates a new ValidationErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidationErrorResponse() *ValidationErrorResponse {
	this := ValidationErrorResponse{}
	return &this
}

// NewValidationErrorResponseWithDefaults instantiates a new ValidationErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidationErrorResponseWithDefaults() *ValidationErrorResponse {
	this := ValidationErrorResponse{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ValidationErrorResponse) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationErrorResponse) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ValidationErrorResponse) IsErrorSet() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *ValidationErrorResponse) SetError(v string) {
	o.Error = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ValidationErrorResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationErrorResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ValidationErrorResponse) IsMessageSet() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ValidationErrorResponse) SetMessage(v string) {
	o.Message = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *ValidationErrorResponse) GetErrors() []FieldError {
	if o == nil || IsNil(o.Errors) {
		var ret []FieldError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationErrorResponse) GetErrorsOk() ([]FieldError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ValidationErrorResponse) IsErrorsSet() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []FieldError and assigns it to the Errors field.
func (o *ValidationErrorResponse) SetErrors(v []FieldError) {
	o.Errors = v
}

func (o ValidationErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidationErrorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableValidationErrorResponse struct {
	value *ValidationErrorResponse
	isSet bool
}

func (v NullableValidationErrorResponse) Get() *ValidationErrorResponse {
	return v.value
}

func (v *NullableValidationErrorResponse) Set(val *ValidationErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableValidationErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableValidationErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidationErrorResponse(val *ValidationErrorResponse) *NullableValidationErrorResponse {
	return &NullableValidationErrorResponse{value: val, isSet: true}
}

func (v NullableValidationErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidationErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

