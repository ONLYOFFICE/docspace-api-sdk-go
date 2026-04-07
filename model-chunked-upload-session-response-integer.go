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
)

// checks if the ChunkedUploadSessionResponseInteger type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChunkedUploadSessionResponseInteger{}

// ChunkedUploadSessionResponseInteger Represents the response returned from a chunked upload session.
type ChunkedUploadSessionResponseInteger struct {
	// The unique identifier for the entity.
	Id NullableString `json:"id,omitempty"`
	// Represents the hierarchical path of folders associated with a chunked upload session.
	Path []int32 `json:"path,omitempty"`
	// The timestamp indicating when the chunked upload session was created.
	Created *time.Time `json:"created,omitempty"`
	// The date and time when the chunked upload session is set to expire.
	Expired *time.Time `json:"expired,omitempty"`
	// Represents the URI or path of the chunked upload session's current location.
	Location NullableString `json:"location,omitempty"`
	// The total size, in bytes, of the file being uploaded in the chunked upload session.
	BytesTotal *int64 `json:"bytes_total,omitempty"`
}

// NewChunkedUploadSessionResponseInteger instantiates a new ChunkedUploadSessionResponseInteger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChunkedUploadSessionResponseInteger() *ChunkedUploadSessionResponseInteger {
	this := ChunkedUploadSessionResponseInteger{}
	return &this
}

// NewChunkedUploadSessionResponseIntegerWithDefaults instantiates a new ChunkedUploadSessionResponseInteger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChunkedUploadSessionResponseIntegerWithDefaults() *ChunkedUploadSessionResponseInteger {
	this := ChunkedUploadSessionResponseInteger{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChunkedUploadSessionResponseInteger) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChunkedUploadSessionResponseInteger) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ChunkedUploadSessionResponseInteger) IsIdSet() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *ChunkedUploadSessionResponseInteger) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ChunkedUploadSessionResponseInteger) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ChunkedUploadSessionResponseInteger) UnsetId() {
	o.Id.Unset()
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChunkedUploadSessionResponseInteger) GetPath() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChunkedUploadSessionResponseInteger) GetPathOk() ([]int32, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ChunkedUploadSessionResponseInteger) IsPathSet() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given []int32 and assigns it to the Path field.
func (o *ChunkedUploadSessionResponseInteger) SetPath(v []int32) {
	o.Path = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ChunkedUploadSessionResponseInteger) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChunkedUploadSessionResponseInteger) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ChunkedUploadSessionResponseInteger) IsCreatedSet() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *ChunkedUploadSessionResponseInteger) SetCreated(v time.Time) {
	o.Created = &v
}

// GetExpired returns the Expired field value if set, zero value otherwise.
func (o *ChunkedUploadSessionResponseInteger) GetExpired() time.Time {
	if o == nil || IsNil(o.Expired) {
		var ret time.Time
		return ret
	}
	return *o.Expired
}

// GetExpiredOk returns a tuple with the Expired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChunkedUploadSessionResponseInteger) GetExpiredOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Expired) {
		return nil, false
	}
	return o.Expired, true
}

// HasExpired returns a boolean if a field has been set.
func (o *ChunkedUploadSessionResponseInteger) IsExpiredSet() bool {
	if o != nil && !IsNil(o.Expired) {
		return true
	}

	return false
}

// SetExpired gets a reference to the given time.Time and assigns it to the Expired field.
func (o *ChunkedUploadSessionResponseInteger) SetExpired(v time.Time) {
	o.Expired = &v
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChunkedUploadSessionResponseInteger) GetLocation() string {
	if o == nil || IsNil(o.Location.Get()) {
		var ret string
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChunkedUploadSessionResponseInteger) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *ChunkedUploadSessionResponseInteger) IsLocationSet() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableString and assigns it to the Location field.
func (o *ChunkedUploadSessionResponseInteger) SetLocation(v string) {
	o.Location.Set(&v)
}
// SetLocationNil sets the value for Location to be an explicit nil
func (o *ChunkedUploadSessionResponseInteger) SetLocationNil() {
	o.Location.Set(nil)
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *ChunkedUploadSessionResponseInteger) UnsetLocation() {
	o.Location.Unset()
}

// GetBytesTotal returns the BytesTotal field value if set, zero value otherwise.
func (o *ChunkedUploadSessionResponseInteger) GetBytesTotal() int64 {
	if o == nil || IsNil(o.BytesTotal) {
		var ret int64
		return ret
	}
	return *o.BytesTotal
}

// GetBytesTotalOk returns a tuple with the BytesTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChunkedUploadSessionResponseInteger) GetBytesTotalOk() (*int64, bool) {
	if o == nil || IsNil(o.BytesTotal) {
		return nil, false
	}
	return o.BytesTotal, true
}

// HasBytesTotal returns a boolean if a field has been set.
func (o *ChunkedUploadSessionResponseInteger) IsBytesTotalSet() bool {
	if o != nil && !IsNil(o.BytesTotal) {
		return true
	}

	return false
}

// SetBytesTotal gets a reference to the given int64 and assigns it to the BytesTotal field.
func (o *ChunkedUploadSessionResponseInteger) SetBytesTotal(v int64) {
	o.BytesTotal = &v
}

func (o ChunkedUploadSessionResponseInteger) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChunkedUploadSessionResponseInteger) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Expired) {
		toSerialize["expired"] = o.Expired
	}
	if o.Location.IsSet() {
		toSerialize["location"] = o.Location.Get()
	}
	if !IsNil(o.BytesTotal) {
		toSerialize["bytes_total"] = o.BytesTotal
	}
	return toSerialize, nil
}

type NullableChunkedUploadSessionResponseInteger struct {
	value *ChunkedUploadSessionResponseInteger
	isSet bool
}

func (v NullableChunkedUploadSessionResponseInteger) Get() *ChunkedUploadSessionResponseInteger {
	return v.value
}

func (v *NullableChunkedUploadSessionResponseInteger) Set(val *ChunkedUploadSessionResponseInteger) {
	v.value = val
	v.isSet = true
}

func (v NullableChunkedUploadSessionResponseInteger) IsSet() bool {
	return v.isSet
}

func (v *NullableChunkedUploadSessionResponseInteger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChunkedUploadSessionResponseInteger(val *ChunkedUploadSessionResponseInteger) *NullableChunkedUploadSessionResponseInteger {
	return &NullableChunkedUploadSessionResponseInteger{value: val, isSet: true}
}

func (v NullableChunkedUploadSessionResponseInteger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChunkedUploadSessionResponseInteger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

