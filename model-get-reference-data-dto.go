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

// checks if the GetReferenceDataDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetReferenceDataDto{}

// GetReferenceDataDto The body of a spreadsheet reference request: the source spreadsheet, and the three ways of naming the document it  refers to, which are tried in the order they are described.
type GetReferenceDataDto struct {
	// The id of the referenced file as the document service recorded it in the formula. It is tried first, and only  when `instanceId` names this portal.
	FileKey NullableString `json:"fileKey"`
	// The portal the reference was made on, as the document service recorded it. Only the id of this portal makes  the file key resolvable; any other value falls through to the path and the link.
	InstanceId NullableString `json:"instanceId"`
	// The spreadsheet the formula sits in. The path is resolved against it - the referenced file is looked for among  the files lying next to it - and it is the file whose read access is checked.
	SourceFileId *int32 `json:"sourceFileId,omitempty"`
	// The title of the referenced file exactly as the formula spells it, matched against the files lying next to the  source file. It is tried after the file key, and only when no link is given.
	Path NullableString `json:"path,omitempty"`
	// The web address the formula points at, an editor link of this portal or one of its short links. It is tried  last, and an address belonging to another site is not resolved at all but handed back for the client to follow  as it is.
	Link NullableString `json:"link,omitempty"`
}

type _GetReferenceDataDto GetReferenceDataDto

// NewGetReferenceDataDto instantiates a new GetReferenceDataDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetReferenceDataDto(fileKey NullableString, instanceId NullableString) *GetReferenceDataDto {
	this := GetReferenceDataDto{}
	this.FileKey = fileKey
	this.InstanceId = instanceId
	return &this
}

// NewGetReferenceDataDtoWithDefaults instantiates a new GetReferenceDataDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetReferenceDataDtoWithDefaults() *GetReferenceDataDto {
	this := GetReferenceDataDto{}
	return &this
}

// GetFileKey returns the FileKey field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GetReferenceDataDto) GetFileKey() string {
	if o == nil || o.FileKey.Get() == nil {
		var ret string
		return ret
	}

	return *o.FileKey.Get()
}

// GetFileKeyOk returns a tuple with the FileKey field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetReferenceDataDto) GetFileKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileKey.Get(), o.FileKey.IsSet()
}

// SetFileKey sets field value
func (o *GetReferenceDataDto) SetFileKey(v string) {
	o.FileKey.Set(&v)
}

// GetInstanceId returns the InstanceId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *GetReferenceDataDto) GetInstanceId() string {
	if o == nil || o.InstanceId.Get() == nil {
		var ret string
		return ret
	}

	return *o.InstanceId.Get()
}

// GetInstanceIdOk returns a tuple with the InstanceId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetReferenceDataDto) GetInstanceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstanceId.Get(), o.InstanceId.IsSet()
}

// SetInstanceId sets field value
func (o *GetReferenceDataDto) SetInstanceId(v string) {
	o.InstanceId.Set(&v)
}

// GetSourceFileId returns the SourceFileId field value if set, zero value otherwise.
func (o *GetReferenceDataDto) GetSourceFileId() int32 {
	if o == nil || IsNil(o.SourceFileId) {
		var ret int32
		return ret
	}
	return *o.SourceFileId
}

// GetSourceFileIdOk returns a tuple with the SourceFileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReferenceDataDto) GetSourceFileIdOk() (*int32, bool) {
	if o == nil || IsNil(o.SourceFileId) {
		return nil, false
	}
	return o.SourceFileId, true
}

// HasSourceFileId returns a boolean if a field has been set.
func (o *GetReferenceDataDto) IsSourceFileIdSet() bool {
	if o != nil && !IsNil(o.SourceFileId) {
		return true
	}

	return false
}

// SetSourceFileId gets a reference to the given int32 and assigns it to the SourceFileId field.
func (o *GetReferenceDataDto) SetSourceFileId(v int32) {
	o.SourceFileId = &v
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetReferenceDataDto) GetPath() string {
	if o == nil || IsNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetReferenceDataDto) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *GetReferenceDataDto) IsPathSet() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *GetReferenceDataDto) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *GetReferenceDataDto) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *GetReferenceDataDto) UnsetPath() {
	o.Path.Unset()
}

// GetLink returns the Link field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetReferenceDataDto) GetLink() string {
	if o == nil || IsNil(o.Link.Get()) {
		var ret string
		return ret
	}
	return *o.Link.Get()
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetReferenceDataDto) GetLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Link.Get(), o.Link.IsSet()
}

// HasLink returns a boolean if a field has been set.
func (o *GetReferenceDataDto) IsLinkSet() bool {
	if o != nil && o.Link.IsSet() {
		return true
	}

	return false
}

// SetLink gets a reference to the given NullableString and assigns it to the Link field.
func (o *GetReferenceDataDto) SetLink(v string) {
	o.Link.Set(&v)
}
// SetLinkNil sets the value for Link to be an explicit nil
func (o *GetReferenceDataDto) SetLinkNil() {
	o.Link.Set(nil)
}

// UnsetLink ensures that no value is present for Link, not even an explicit nil
func (o *GetReferenceDataDto) UnsetLink() {
	o.Link.Unset()
}

func (o GetReferenceDataDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetReferenceDataDto) ToMap() (map[string]interface{}, error) {
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

func (o *GetReferenceDataDto) UnmarshalJSON(data []byte) (err error) {
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

	varGetReferenceDataDto := _GetReferenceDataDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetReferenceDataDto)

	if err != nil {
		return err
	}

	*o = GetReferenceDataDto(varGetReferenceDataDto)

	return err
}

type NullableGetReferenceDataDto struct {
	value *GetReferenceDataDto
	isSet bool
}

func (v NullableGetReferenceDataDto) Get() *GetReferenceDataDto {
	return v.value
}

func (v *NullableGetReferenceDataDto) Set(val *GetReferenceDataDto) {
	v.value = val
	v.isSet = true
}

func (v NullableGetReferenceDataDto) IsSet() bool {
	return v.isSet
}

func (v *NullableGetReferenceDataDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetReferenceDataDto(val *GetReferenceDataDto) *NullableGetReferenceDataDto {
	return &NullableGetReferenceDataDto{value: val, isSet: true}
}

func (v NullableGetReferenceDataDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetReferenceDataDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

