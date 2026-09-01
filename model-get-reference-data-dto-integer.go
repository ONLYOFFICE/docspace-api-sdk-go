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

// checks if the GetReferenceDataDtoInteger type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetReferenceDataDtoInteger{}

// GetReferenceDataDtoInteger The request parameters for getting reference data.
type GetReferenceDataDtoInteger struct {
	// The unique document identifier used by the service to get a link to the file.
	FileKey NullableString `json:"fileKey"`
	// The unique system identifier.
	InstanceId NullableString `json:"instanceId"`
	// The source file ID.
	SourceFileId *int32 `json:"sourceFileId,omitempty"`
	// The file name or relative path for the formula editor.
	Path NullableString `json:"path,omitempty"`
	// The file link.
	Link NullableString `json:"link,omitempty"`
}

type _GetReferenceDataDtoInteger GetReferenceDataDtoInteger

// NewGetReferenceDataDtoInteger instantiates a new GetReferenceDataDtoInteger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetReferenceDataDtoInteger(fileKey NullableString, instanceId NullableString) *GetReferenceDataDtoInteger {
	this := GetReferenceDataDtoInteger{}
	this.FileKey = fileKey
	this.InstanceId = instanceId
	return &this
}

// NewGetReferenceDataDtoIntegerWithDefaults instantiates a new GetReferenceDataDtoInteger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetReferenceDataDtoIntegerWithDefaults() *GetReferenceDataDtoInteger {
	this := GetReferenceDataDtoInteger{}
	return &this
}

// GetFileKey returns the FileKey field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GetReferenceDataDtoInteger) GetFileKey() string {
	if o == nil || o.FileKey.Get() == nil {
		var ret string
		return ret
	}

	return *o.FileKey.Get()
}

// GetFileKeyOk returns a tuple with the FileKey field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetReferenceDataDtoInteger) GetFileKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileKey.Get(), o.FileKey.IsSet()
}

// SetFileKey sets field value
func (o *GetReferenceDataDtoInteger) SetFileKey(v string) {
	o.FileKey.Set(&v)
}

// GetInstanceId returns the InstanceId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GetReferenceDataDtoInteger) GetInstanceId() string {
	if o == nil || o.InstanceId.Get() == nil {
		var ret string
		return ret
	}

	return *o.InstanceId.Get()
}

// GetInstanceIdOk returns a tuple with the InstanceId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetReferenceDataDtoInteger) GetInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstanceId.Get(), o.InstanceId.IsSet()
}

// SetInstanceId sets field value
func (o *GetReferenceDataDtoInteger) SetInstanceId(v string) {
	o.InstanceId.Set(&v)
}

// GetSourceFileId returns the SourceFileId field value if set, zero value otherwise.
func (o *GetReferenceDataDtoInteger) GetSourceFileId() int32 {
	if o == nil || IsNil(o.SourceFileId) {
		var ret int32
		return ret
	}
	return *o.SourceFileId
}

// GetSourceFileIdOk returns a tuple with the SourceFileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReferenceDataDtoInteger) GetSourceFileIdOk() (*int32, bool) {
	if o == nil || IsNil(o.SourceFileId) {
		return nil, false
	}
	return o.SourceFileId, true
}

// HasSourceFileId returns a boolean if a field has been set.
func (o *GetReferenceDataDtoInteger) IsSourceFileIdSet() bool {
	if o != nil && !IsNil(o.SourceFileId) {
		return true
	}

	return false
}

// SetSourceFileId gets a reference to the given int32 and assigns it to the SourceFileId field.
func (o *GetReferenceDataDtoInteger) SetSourceFileId(v int32) {
	o.SourceFileId = &v
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetReferenceDataDtoInteger) GetPath() string {
	if o == nil || IsNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetReferenceDataDtoInteger) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *GetReferenceDataDtoInteger) IsPathSet() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *GetReferenceDataDtoInteger) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *GetReferenceDataDtoInteger) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *GetReferenceDataDtoInteger) UnsetPath() {
	o.Path.Unset()
}

// GetLink returns the Link field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetReferenceDataDtoInteger) GetLink() string {
	if o == nil || IsNil(o.Link.Get()) {
		var ret string
		return ret
	}
	return *o.Link.Get()
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetReferenceDataDtoInteger) GetLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Link.Get(), o.Link.IsSet()
}

// HasLink returns a boolean if a field has been set.
func (o *GetReferenceDataDtoInteger) IsLinkSet() bool {
	if o != nil && o.Link.IsSet() {
		return true
	}

	return false
}

// SetLink gets a reference to the given NullableString and assigns it to the Link field.
func (o *GetReferenceDataDtoInteger) SetLink(v string) {
	o.Link.Set(&v)
}
// SetLinkNil sets the value for Link to be an explicit nil
func (o *GetReferenceDataDtoInteger) SetLinkNil() {
	o.Link.Set(nil)
}

// UnsetLink ensures that no value is present for Link, not even an explicit nil
func (o *GetReferenceDataDtoInteger) UnsetLink() {
	o.Link.Unset()
}

func (o GetReferenceDataDtoInteger) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetReferenceDataDtoInteger) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fileKey"] = o.FileKey.Get()
	toSerialize["instanceId"] = o.InstanceId.Get()
	if !IsNil(o.SourceFileId) {
		toSerialize["sourceFileId"] = o.SourceFileId
	}
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	if o.Link.IsSet() {
		toSerialize["link"] = o.Link.Get()
	}
	return toSerialize, nil
}

func (o *GetReferenceDataDtoInteger) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fileKey",
		"instanceId",
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

	varGetReferenceDataDtoInteger := _GetReferenceDataDtoInteger{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetReferenceDataDtoInteger)

	if err != nil {
		return err
	}

	*o = GetReferenceDataDtoInteger(varGetReferenceDataDtoInteger)

	return err
}

type NullableGetReferenceDataDtoInteger struct {
	value *GetReferenceDataDtoInteger
	isSet bool
}

func (v NullableGetReferenceDataDtoInteger) Get() *GetReferenceDataDtoInteger {
	return v.value
}

func (v *NullableGetReferenceDataDtoInteger) Set(val *GetReferenceDataDtoInteger) {
	v.value = val
	v.isSet = true
}

func (v NullableGetReferenceDataDtoInteger) IsSet() bool {
	return v.isSet
}

func (v *NullableGetReferenceDataDtoInteger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetReferenceDataDtoInteger(val *GetReferenceDataDtoInteger) *NullableGetReferenceDataDtoInteger {
	return &NullableGetReferenceDataDtoInteger{value: val, isSet: true}
}

func (v NullableGetReferenceDataDtoInteger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetReferenceDataDtoInteger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

