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

// checks if the ThirdPartyCheckConversionRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThirdPartyCheckConversionRequestDto{}

// ThirdPartyCheckConversionRequestDto The parameters of one file conversion.
type ThirdPartyCheckConversionRequestDto struct {
	// The file to convert. It is taken from the route of the operation, so a value sent in the body is overwritten.
	FileId NullableString `json:"fileId,omitempty"`
	// How to wait for the result: `true` converts inside the request and answers with the finished result, which is  only sensible for small documents, while `false` queues the conversion and answers with an entry to poll.
	Sync *bool `json:"sync,omitempty"`
	// Whether the conversion is to be started. It is set by the operation itself, so a value sent in the body is  overwritten.
	StartConvert *bool `json:"startConvert,omitempty"`
	// The version to convert; 0 or less means the current version.
	Version *int32 `json:"version,omitempty"`
	// The password that opens the source document, for a file that is protected by one; anything else may be left  out.
	Password NullableString `json:"password,omitempty"`
	// The extension of the format to convert into, without the dot, and one the portal can produce from that  source format; left out, the default of the portal for that kind of document is used.
	OutputType NullableString `json:"outputType,omitempty"`
	// Where the result goes when the file has been converted before: `true` creates another file beside the source,  `false` replaces the converted file that already exists.
	CreateNewIfExist *bool `json:"createNewIfExist,omitempty"`
}

// NewThirdPartyCheckConversionRequestDto instantiates a new ThirdPartyCheckConversionRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThirdPartyCheckConversionRequestDto() *ThirdPartyCheckConversionRequestDto {
	this := ThirdPartyCheckConversionRequestDto{}
	return &this
}

// NewThirdPartyCheckConversionRequestDtoWithDefaults instantiates a new ThirdPartyCheckConversionRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThirdPartyCheckConversionRequestDtoWithDefaults() *ThirdPartyCheckConversionRequestDto {
	this := ThirdPartyCheckConversionRequestDto{}
	return &this
}

// GetFileId returns the FileId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyCheckConversionRequestDto) GetFileId() string {
	if o == nil || IsNil(o.FileId.Get()) {
		var ret string
		return ret
	}
	return *o.FileId.Get()
}

// GetFileIdOk returns a tuple with the FileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyCheckConversionRequestDto) GetFileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileId.Get(), o.FileId.IsSet()
}

// HasFileId returns a boolean if a field has been set.
func (o *ThirdPartyCheckConversionRequestDto) IsFileIdSet() bool {
	if o != nil && o.FileId.IsSet() {
		return true
	}

	return false
}

// SetFileId gets a reference to the given NullableString and assigns it to the FileId field.
func (o *ThirdPartyCheckConversionRequestDto) SetFileId(v string) {
	o.FileId.Set(&v)
}
// SetFileIdNil sets the value for FileId to be an explicit nil
func (o *ThirdPartyCheckConversionRequestDto) SetFileIdNil() {
	o.FileId.Set(nil)
}

// UnsetFileId ensures that no value is present for FileId, not even an explicit nil
func (o *ThirdPartyCheckConversionRequestDto) UnsetFileId() {
	o.FileId.Unset()
}

// GetSync returns the Sync field value if set, zero value otherwise.
func (o *ThirdPartyCheckConversionRequestDto) GetSync() bool {
	if o == nil || IsNil(o.Sync) {
		var ret bool
		return ret
	}
	return *o.Sync
}

// GetSyncOk returns a tuple with the Sync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyCheckConversionRequestDto) GetSyncOk() (*bool, bool) {
	if o == nil || IsNil(o.Sync) {
		return nil, false
	}
	return o.Sync, true
}

// HasSync returns a boolean if a field has been set.
func (o *ThirdPartyCheckConversionRequestDto) IsSyncSet() bool {
	if o != nil && !IsNil(o.Sync) {
		return true
	}

	return false
}

// SetSync gets a reference to the given bool and assigns it to the Sync field.
func (o *ThirdPartyCheckConversionRequestDto) SetSync(v bool) {
	o.Sync = &v
}

// GetStartConvert returns the StartConvert field value if set, zero value otherwise.
func (o *ThirdPartyCheckConversionRequestDto) GetStartConvert() bool {
	if o == nil || IsNil(o.StartConvert) {
		var ret bool
		return ret
	}
	return *o.StartConvert
}

// GetStartConvertOk returns a tuple with the StartConvert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyCheckConversionRequestDto) GetStartConvertOk() (*bool, bool) {
	if o == nil || IsNil(o.StartConvert) {
		return nil, false
	}
	return o.StartConvert, true
}

// HasStartConvert returns a boolean if a field has been set.
func (o *ThirdPartyCheckConversionRequestDto) IsStartConvertSet() bool {
	if o != nil && !IsNil(o.StartConvert) {
		return true
	}

	return false
}

// SetStartConvert gets a reference to the given bool and assigns it to the StartConvert field.
func (o *ThirdPartyCheckConversionRequestDto) SetStartConvert(v bool) {
	o.StartConvert = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ThirdPartyCheckConversionRequestDto) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyCheckConversionRequestDto) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ThirdPartyCheckConversionRequestDto) IsVersionSet() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *ThirdPartyCheckConversionRequestDto) SetVersion(v int32) {
	o.Version = &v
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyCheckConversionRequestDto) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyCheckConversionRequestDto) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *ThirdPartyCheckConversionRequestDto) IsPasswordSet() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *ThirdPartyCheckConversionRequestDto) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *ThirdPartyCheckConversionRequestDto) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *ThirdPartyCheckConversionRequestDto) UnsetPassword() {
	o.Password.Unset()
}

// GetOutputType returns the OutputType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyCheckConversionRequestDto) GetOutputType() string {
	if o == nil || IsNil(o.OutputType.Get()) {
		var ret string
		return ret
	}
	return *o.OutputType.Get()
}

// GetOutputTypeOk returns a tuple with the OutputType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyCheckConversionRequestDto) GetOutputTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OutputType.Get(), o.OutputType.IsSet()
}

// HasOutputType returns a boolean if a field has been set.
func (o *ThirdPartyCheckConversionRequestDto) IsOutputTypeSet() bool {
	if o != nil && o.OutputType.IsSet() {
		return true
	}

	return false
}

// SetOutputType gets a reference to the given NullableString and assigns it to the OutputType field.
func (o *ThirdPartyCheckConversionRequestDto) SetOutputType(v string) {
	o.OutputType.Set(&v)
}
// SetOutputTypeNil sets the value for OutputType to be an explicit nil
func (o *ThirdPartyCheckConversionRequestDto) SetOutputTypeNil() {
	o.OutputType.Set(nil)
}

// UnsetOutputType ensures that no value is present for OutputType, not even an explicit nil
func (o *ThirdPartyCheckConversionRequestDto) UnsetOutputType() {
	o.OutputType.Unset()
}

// GetCreateNewIfExist returns the CreateNewIfExist field value if set, zero value otherwise.
func (o *ThirdPartyCheckConversionRequestDto) GetCreateNewIfExist() bool {
	if o == nil || IsNil(o.CreateNewIfExist) {
		var ret bool
		return ret
	}
	return *o.CreateNewIfExist
}

// GetCreateNewIfExistOk returns a tuple with the CreateNewIfExist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyCheckConversionRequestDto) GetCreateNewIfExistOk() (*bool, bool) {
	if o == nil || IsNil(o.CreateNewIfExist) {
		return nil, false
	}
	return o.CreateNewIfExist, true
}

// HasCreateNewIfExist returns a boolean if a field has been set.
func (o *ThirdPartyCheckConversionRequestDto) IsCreateNewIfExistSet() bool {
	if o != nil && !IsNil(o.CreateNewIfExist) {
		return true
	}

	return false
}

// SetCreateNewIfExist gets a reference to the given bool and assigns it to the CreateNewIfExist field.
func (o *ThirdPartyCheckConversionRequestDto) SetCreateNewIfExist(v bool) {
	o.CreateNewIfExist = &v
}

func (o ThirdPartyCheckConversionRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThirdPartyCheckConversionRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FileId.IsSet() {
		toSerialize["fileId"] = o.FileId.Get()
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

type NullableThirdPartyCheckConversionRequestDto struct {
	value *ThirdPartyCheckConversionRequestDto
	isSet bool
}

func (v NullableThirdPartyCheckConversionRequestDto) Get() *ThirdPartyCheckConversionRequestDto {
	return v.value
}

func (v *NullableThirdPartyCheckConversionRequestDto) Set(val *ThirdPartyCheckConversionRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableThirdPartyCheckConversionRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableThirdPartyCheckConversionRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThirdPartyCheckConversionRequestDto(val *ThirdPartyCheckConversionRequestDto) *NullableThirdPartyCheckConversionRequestDto {
	return &NullableThirdPartyCheckConversionRequestDto{value: val, isSet: true}
}

func (v NullableThirdPartyCheckConversionRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThirdPartyCheckConversionRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

