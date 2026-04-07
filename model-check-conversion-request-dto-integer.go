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

// checks if the CheckConversionRequestDtoInteger type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckConversionRequestDtoInteger{}

// CheckConversionRequestDtoInteger The parameters for checking file conversion.
type CheckConversionRequestDtoInteger struct {
	// The file ID to check conversion proccess.
	FileId *int32 `json:"fileId,omitempty"`
	// Specifies if the conversion process is synchronous or not.
	Sync *bool `json:"sync,omitempty"`
	// Specifies whether to start a conversion process or not.
	StartConvert *bool `json:"startConvert,omitempty"`
	// The file version that is converted.
	Version *int32 `json:"version,omitempty"`
	// The password of the converted file.
	Password NullableString `json:"password,omitempty"`
	// The conversion output type.
	OutputType NullableString `json:"outputType,omitempty"`
	// Specifies whether to create a new file if it exists or not.
	CreateNewIfExist *bool `json:"createNewIfExist,omitempty"`
}

// NewCheckConversionRequestDtoInteger instantiates a new CheckConversionRequestDtoInteger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckConversionRequestDtoInteger() *CheckConversionRequestDtoInteger {
	this := CheckConversionRequestDtoInteger{}
	return &this
}

// NewCheckConversionRequestDtoIntegerWithDefaults instantiates a new CheckConversionRequestDtoInteger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckConversionRequestDtoIntegerWithDefaults() *CheckConversionRequestDtoInteger {
	this := CheckConversionRequestDtoInteger{}
	return &this
}

// GetFileId returns the FileId field value if set, zero value otherwise.
func (o *CheckConversionRequestDtoInteger) GetFileId() int32 {
	if o == nil || IsNil(o.FileId) {
		var ret int32
		return ret
	}
	return *o.FileId
}

// GetFileIdOk returns a tuple with the FileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckConversionRequestDtoInteger) GetFileIdOk() (*int32, bool) {
	if o == nil || IsNil(o.FileId) {
		return nil, false
	}
	return o.FileId, true
}

// HasFileId returns a boolean if a field has been set.
func (o *CheckConversionRequestDtoInteger) IsFileIdSet() bool {
	if o != nil && !IsNil(o.FileId) {
		return true
	}

	return false
}

// SetFileId gets a reference to the given int32 and assigns it to the FileId field.
func (o *CheckConversionRequestDtoInteger) SetFileId(v int32) {
	o.FileId = &v
}

// GetSync returns the Sync field value if set, zero value otherwise.
func (o *CheckConversionRequestDtoInteger) GetSync() bool {
	if o == nil || IsNil(o.Sync) {
		var ret bool
		return ret
	}
	return *o.Sync
}

// GetSyncOk returns a tuple with the Sync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckConversionRequestDtoInteger) GetSyncOk() (*bool, bool) {
	if o == nil || IsNil(o.Sync) {
		return nil, false
	}
	return o.Sync, true
}

// HasSync returns a boolean if a field has been set.
func (o *CheckConversionRequestDtoInteger) IsSyncSet() bool {
	if o != nil && !IsNil(o.Sync) {
		return true
	}

	return false
}

// SetSync gets a reference to the given bool and assigns it to the Sync field.
func (o *CheckConversionRequestDtoInteger) SetSync(v bool) {
	o.Sync = &v
}

// GetStartConvert returns the StartConvert field value if set, zero value otherwise.
func (o *CheckConversionRequestDtoInteger) GetStartConvert() bool {
	if o == nil || IsNil(o.StartConvert) {
		var ret bool
		return ret
	}
	return *o.StartConvert
}

// GetStartConvertOk returns a tuple with the StartConvert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckConversionRequestDtoInteger) GetStartConvertOk() (*bool, bool) {
	if o == nil || IsNil(o.StartConvert) {
		return nil, false
	}
	return o.StartConvert, true
}

// HasStartConvert returns a boolean if a field has been set.
func (o *CheckConversionRequestDtoInteger) IsStartConvertSet() bool {
	if o != nil && !IsNil(o.StartConvert) {
		return true
	}

	return false
}

// SetStartConvert gets a reference to the given bool and assigns it to the StartConvert field.
func (o *CheckConversionRequestDtoInteger) SetStartConvert(v bool) {
	o.StartConvert = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CheckConversionRequestDtoInteger) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckConversionRequestDtoInteger) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CheckConversionRequestDtoInteger) IsVersionSet() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *CheckConversionRequestDtoInteger) SetVersion(v int32) {
	o.Version = &v
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CheckConversionRequestDtoInteger) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckConversionRequestDtoInteger) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *CheckConversionRequestDtoInteger) IsPasswordSet() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *CheckConversionRequestDtoInteger) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *CheckConversionRequestDtoInteger) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *CheckConversionRequestDtoInteger) UnsetPassword() {
	o.Password.Unset()
}

// GetOutputType returns the OutputType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CheckConversionRequestDtoInteger) GetOutputType() string {
	if o == nil || IsNil(o.OutputType.Get()) {
		var ret string
		return ret
	}
	return *o.OutputType.Get()
}

// GetOutputTypeOk returns a tuple with the OutputType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckConversionRequestDtoInteger) GetOutputTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OutputType.Get(), o.OutputType.IsSet()
}

// HasOutputType returns a boolean if a field has been set.
func (o *CheckConversionRequestDtoInteger) IsOutputTypeSet() bool {
	if o != nil && o.OutputType.IsSet() {
		return true
	}

	return false
}

// SetOutputType gets a reference to the given NullableString and assigns it to the OutputType field.
func (o *CheckConversionRequestDtoInteger) SetOutputType(v string) {
	o.OutputType.Set(&v)
}
// SetOutputTypeNil sets the value for OutputType to be an explicit nil
func (o *CheckConversionRequestDtoInteger) SetOutputTypeNil() {
	o.OutputType.Set(nil)
}

// UnsetOutputType ensures that no value is present for OutputType, not even an explicit nil
func (o *CheckConversionRequestDtoInteger) UnsetOutputType() {
	o.OutputType.Unset()
}

// GetCreateNewIfExist returns the CreateNewIfExist field value if set, zero value otherwise.
func (o *CheckConversionRequestDtoInteger) GetCreateNewIfExist() bool {
	if o == nil || IsNil(o.CreateNewIfExist) {
		var ret bool
		return ret
	}
	return *o.CreateNewIfExist
}

// GetCreateNewIfExistOk returns a tuple with the CreateNewIfExist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckConversionRequestDtoInteger) GetCreateNewIfExistOk() (*bool, bool) {
	if o == nil || IsNil(o.CreateNewIfExist) {
		return nil, false
	}
	return o.CreateNewIfExist, true
}

// HasCreateNewIfExist returns a boolean if a field has been set.
func (o *CheckConversionRequestDtoInteger) IsCreateNewIfExistSet() bool {
	if o != nil && !IsNil(o.CreateNewIfExist) {
		return true
	}

	return false
}

// SetCreateNewIfExist gets a reference to the given bool and assigns it to the CreateNewIfExist field.
func (o *CheckConversionRequestDtoInteger) SetCreateNewIfExist(v bool) {
	o.CreateNewIfExist = &v
}

func (o CheckConversionRequestDtoInteger) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckConversionRequestDtoInteger) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FileId) {
		toSerialize["fileId"] = o.FileId
	}
	if !IsNil(o.Sync) {
		toSerialize["sync"] = o.Sync
	}
	if !IsNil(o.StartConvert) {
		toSerialize["startConvert"] = o.StartConvert
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if o.OutputType.IsSet() {
		toSerialize["outputType"] = o.OutputType.Get()
	}
	if !IsNil(o.CreateNewIfExist) {
		toSerialize["createNewIfExist"] = o.CreateNewIfExist
	}
	return toSerialize, nil
}

type NullableCheckConversionRequestDtoInteger struct {
	value *CheckConversionRequestDtoInteger
	isSet bool
}

func (v NullableCheckConversionRequestDtoInteger) Get() *CheckConversionRequestDtoInteger {
	return v.value
}

func (v *NullableCheckConversionRequestDtoInteger) Set(val *CheckConversionRequestDtoInteger) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckConversionRequestDtoInteger) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckConversionRequestDtoInteger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckConversionRequestDtoInteger(val *CheckConversionRequestDtoInteger) *NullableCheckConversionRequestDtoInteger {
	return &NullableCheckConversionRequestDtoInteger{value: val, isSet: true}
}

func (v NullableCheckConversionRequestDtoInteger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckConversionRequestDtoInteger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

