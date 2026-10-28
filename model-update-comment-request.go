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

// checks if the UpdateCommentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateCommentRequest{}

// UpdateCommentRequest The comment to store on one version of a file.
type UpdateCommentRequest struct {
	// The version the comment belongs to, as reported by `GET api/2.0/files/file/{fileId}/edit/history`. A version  that does not exist is rejected as an invalid request.
	Version int32 `json:"version"`
	// The note that explains what changed in that version, as the version history shows it. An empty text clears the  note, and a longer one is cut rather than refused, so read the stored text from the answer.
	Comment NullableString `json:"comment,omitempty"`
}

type _UpdateCommentRequest UpdateCommentRequest

// NewUpdateCommentRequest instantiates a new UpdateCommentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCommentRequest(version int32) *UpdateCommentRequest {
	this := UpdateCommentRequest{}
	this.Version = version
	return &this
}

// NewUpdateCommentRequestWithDefaults instantiates a new UpdateCommentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCommentRequestWithDefaults() *UpdateCommentRequest {
	this := UpdateCommentRequest{}
	return &this
}

// GetVersion returns the Version field value
func (o *UpdateCommentRequest) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *UpdateCommentRequest) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *UpdateCommentRequest) SetVersion(v int32) {
	o.Version = v
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateCommentRequest) GetComment() string {
	if o == nil || IsNil(o.Comment.Get()) {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateCommentRequest) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *UpdateCommentRequest) IsCommentSet() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *UpdateCommentRequest) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *UpdateCommentRequest) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *UpdateCommentRequest) UnsetComment() {
	o.Comment.Unset()
}

func (o UpdateCommentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateCommentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["version"] = o.Version
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	return toSerialize, nil
}

func (o *UpdateCommentRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"version",
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

	varUpdateCommentRequest := _UpdateCommentRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varUpdateCommentRequest)

	if err != nil {
		return err
	}

	*o = UpdateCommentRequest(varUpdateCommentRequest)

	return err
}

type NullableUpdateCommentRequest struct {
	value *UpdateCommentRequest
	isSet bool
}

func (v NullableUpdateCommentRequest) Get() *UpdateCommentRequest {
	return v.value
}

func (v *NullableUpdateCommentRequest) Set(val *UpdateCommentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCommentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCommentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCommentRequest(val *UpdateCommentRequest) *NullableUpdateCommentRequest {
	return &NullableUpdateCommentRequest{value: val, isSet: true}
}

func (v NullableUpdateCommentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCommentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

