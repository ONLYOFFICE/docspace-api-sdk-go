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

// checks if the ErrorApiResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorApiResponse{}

// ErrorApiResponse The error body returned with every failed request.
type ErrorApiResponse struct {
	// The response status flag. Always 1 on an error, as opposed to 0 on success.
	Status *int32 `json:"status,omitempty"`
	// The HTTP status code of the response, repeated in the body.
	StatusCode *int32 `json:"statusCode,omitempty"`
	Error *ErrorApiResponseError `json:"error,omitempty"`
}

// NewErrorApiResponse instantiates a new ErrorApiResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorApiResponse() *ErrorApiResponse {
	this := ErrorApiResponse{}
	return &this
}

// NewErrorApiResponseWithDefaults instantiates a new ErrorApiResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorApiResponseWithDefaults() *ErrorApiResponse {
	this := ErrorApiResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ErrorApiResponse) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorApiResponse) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ErrorApiResponse) IsStatusSet() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *ErrorApiResponse) SetStatus(v int32) {
	o.Status = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *ErrorApiResponse) GetStatusCode() int32 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorApiResponse) GetStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *ErrorApiResponse) IsStatusCodeSet() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *ErrorApiResponse) SetStatusCode(v int32) {
	o.StatusCode = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ErrorApiResponse) GetError() ErrorApiResponseError {
	if o == nil || IsNil(o.Error) {
		var ret ErrorApiResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorApiResponse) GetErrorOk() (*ErrorApiResponseError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ErrorApiResponse) IsErrorSet() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given ErrorApiResponseError and assigns it to the Error field.
func (o *ErrorApiResponse) SetError(v ErrorApiResponseError) {
	o.Error = &v
}

func (o ErrorApiResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorApiResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.StatusCode) {
		toSerialize["statusCode"] = o.StatusCode
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableErrorApiResponse struct {
	value *ErrorApiResponse
	isSet bool
}

func (v NullableErrorApiResponse) Get() *ErrorApiResponse {
	return v.value
}

func (v *NullableErrorApiResponse) Set(val *ErrorApiResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorApiResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorApiResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorApiResponse(val *ErrorApiResponse) *NullableErrorApiResponse {
	return &NullableErrorApiResponse{value: val, isSet: true}
}

func (v NullableErrorApiResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorApiResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

