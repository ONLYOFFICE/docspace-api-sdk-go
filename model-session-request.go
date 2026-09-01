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
	"time"
	"bytes"
	"fmt"
)

// checks if the SessionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionRequest{}

// SessionRequest The session request parameters.
type SessionRequest struct {
	// The file name.
	FileName NullableString `json:"fileName"`
	// The file size.
	FileSize *int64 `json:"fileSize,omitempty"`
	// The relative path to the file.
	RelativePath NullableString `json:"relativePath,omitempty"`
	// The date and time when the file was created.
	CreateOn NullableTime `json:"createOn,omitempty"`
	// Specifies whether the file is encrypted or not.
	Encrypted *bool `json:"encrypted,omitempty"`
	// Specifies whether to create a new file if it already exists.
	CreateNewIfExist *bool `json:"createNewIfExist,omitempty"`
}

type _SessionRequest SessionRequest

// NewSessionRequest instantiates a new SessionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionRequest(fileName NullableString) *SessionRequest {
	this := SessionRequest{}
	this.FileName = fileName
	return &this
}

// NewSessionRequestWithDefaults instantiates a new SessionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionRequestWithDefaults() *SessionRequest {
	this := SessionRequest{}
	return &this
}

// GetFileName returns the FileName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SessionRequest) GetFileName() string {
	if o == nil || o.FileName.Get() == nil {
		var ret string
		return ret
	}

	return *o.FileName.Get()
}

// GetFileNameOk returns a tuple with the FileName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SessionRequest) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileName.Get(), o.FileName.IsSet()
}

// SetFileName sets field value
func (o *SessionRequest) SetFileName(v string) {
	o.FileName.Set(&v)
}

// GetFileSize returns the FileSize field value if set, zero value otherwise.
func (o *SessionRequest) GetFileSize() int64 {
	if o == nil || IsNil(o.FileSize) {
		var ret int64
		return ret
	}
	return *o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionRequest) GetFileSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.FileSize) {
		return nil, false
	}
	return o.FileSize, true
}

// HasFileSize returns a boolean if a field has been set.
func (o *SessionRequest) IsFileSizeSet() bool {
	if o != nil && !IsNil(o.FileSize) {
		return true
	}

	return false
}

// SetFileSize gets a reference to the given int64 and assigns it to the FileSize field.
func (o *SessionRequest) SetFileSize(v int64) {
	o.FileSize = &v
}

// GetRelativePath returns the RelativePath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SessionRequest) GetRelativePath() string {
	if o == nil || IsNil(o.RelativePath.Get()) {
		var ret string
		return ret
	}
	return *o.RelativePath.Get()
}

// GetRelativePathOk returns a tuple with the RelativePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SessionRequest) GetRelativePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RelativePath.Get(), o.RelativePath.IsSet()
}

// HasRelativePath returns a boolean if a field has been set.
func (o *SessionRequest) IsRelativePathSet() bool {
	if o != nil && o.RelativePath.IsSet() {
		return true
	}

	return false
}

// SetRelativePath gets a reference to the given NullableString and assigns it to the RelativePath field.
func (o *SessionRequest) SetRelativePath(v string) {
	o.RelativePath.Set(&v)
}
// SetRelativePathNil sets the value for RelativePath to be an explicit nil
func (o *SessionRequest) SetRelativePathNil() {
	o.RelativePath.Set(nil)
}

// UnsetRelativePath ensures that no value is present for RelativePath, not even an explicit nil
func (o *SessionRequest) UnsetRelativePath() {
	o.RelativePath.Unset()
}

// GetCreateOn returns the CreateOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SessionRequest) GetCreateOn() time.Time {
	if o == nil || IsNil(o.CreateOn.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CreateOn.Get()
}

// GetCreateOnOk returns a tuple with the CreateOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SessionRequest) GetCreateOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreateOn.Get(), o.CreateOn.IsSet()
}

// HasCreateOn returns a boolean if a field has been set.
func (o *SessionRequest) IsCreateOnSet() bool {
	if o != nil && o.CreateOn.IsSet() {
		return true
	}

	return false
}

// SetCreateOn gets a reference to the given NullableTime and assigns it to the CreateOn field.
func (o *SessionRequest) SetCreateOn(v time.Time) {
	o.CreateOn.Set(&v)
}
// SetCreateOnNil sets the value for CreateOn to be an explicit nil
func (o *SessionRequest) SetCreateOnNil() {
	o.CreateOn.Set(nil)
}

// UnsetCreateOn ensures that no value is present for CreateOn, not even an explicit nil
func (o *SessionRequest) UnsetCreateOn() {
	o.CreateOn.Unset()
}

// GetEncrypted returns the Encrypted field value if set, zero value otherwise.
func (o *SessionRequest) GetEncrypted() bool {
	if o == nil || IsNil(o.Encrypted) {
		var ret bool
		return ret
	}
	return *o.Encrypted
}

// GetEncryptedOk returns a tuple with the Encrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionRequest) GetEncryptedOk() (*bool, bool) {
	if o == nil || IsNil(o.Encrypted) {
		return nil, false
	}
	return o.Encrypted, true
}

// HasEncrypted returns a boolean if a field has been set.
func (o *SessionRequest) IsEncryptedSet() bool {
	if o != nil && !IsNil(o.Encrypted) {
		return true
	}

	return false
}

// SetEncrypted gets a reference to the given bool and assigns it to the Encrypted field.
func (o *SessionRequest) SetEncrypted(v bool) {
	o.Encrypted = &v
}

// GetCreateNewIfExist returns the CreateNewIfExist field value if set, zero value otherwise.
func (o *SessionRequest) GetCreateNewIfExist() bool {
	if o == nil || IsNil(o.CreateNewIfExist) {
		var ret bool
		return ret
	}
	return *o.CreateNewIfExist
}

// GetCreateNewIfExistOk returns a tuple with the CreateNewIfExist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionRequest) GetCreateNewIfExistOk() (*bool, bool) {
	if o == nil || IsNil(o.CreateNewIfExist) {
		return nil, false
	}
	return o.CreateNewIfExist, true
}

// HasCreateNewIfExist returns a boolean if a field has been set.
func (o *SessionRequest) IsCreateNewIfExistSet() bool {
	if o != nil && !IsNil(o.CreateNewIfExist) {
		return true
	}

	return false
}

// SetCreateNewIfExist gets a reference to the given bool and assigns it to the CreateNewIfExist field.
func (o *SessionRequest) SetCreateNewIfExist(v bool) {
	o.CreateNewIfExist = &v
}

func (o SessionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fileName"] = o.FileName.Get()
	if !IsNil(o.FileSize) {
		toSerialize["fileSize"] = o.FileSize
	}
	if o.RelativePath.IsSet() {
		toSerialize["relativePath"] = o.RelativePath.Get()
	}
	if o.CreateOn.IsSet() {
		toSerialize["createOn"] = o.CreateOn.Get()
	}
	if !IsNil(o.Encrypted) {
		toSerialize["encrypted"] = o.Encrypted
	}
	if !IsNil(o.CreateNewIfExist) {
		toSerialize["createNewIfExist"] = o.CreateNewIfExist
	}
	return toSerialize, nil
}

func (o *SessionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fileName",
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

	varSessionRequest := _SessionRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varSessionRequest)

	if err != nil {
		return err
	}

	*o = SessionRequest(varSessionRequest)

	return err
}

type NullableSessionRequest struct {
	value *SessionRequest
	isSet bool
}

func (v NullableSessionRequest) Get() *SessionRequest {
	return v.value
}

func (v *NullableSessionRequest) Set(val *SessionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionRequest(val *SessionRequest) *NullableSessionRequest {
	return &NullableSessionRequest{value: val, isSet: true}
}

func (v NullableSessionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

