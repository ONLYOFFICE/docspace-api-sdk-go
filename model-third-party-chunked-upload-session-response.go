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

// checks if the ThirdPartyChunkedUploadSessionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThirdPartyChunkedUploadSessionResponse{}

// ThirdPartyChunkedUploadSessionResponse The reserved chunked upload: where the parts are sent, how much was declared and when the reservation lapses. No  content of the file is described here.
type ThirdPartyChunkedUploadSessionResponse struct {
	// The identifier of the reserved upload, repeated in the path of every call that follows it - the chunk uploads,  the finalize and the abort. It is thirty-two hexadecimal characters without separators, and it is the only  thing the server checks, so anyone holding it can write into this upload.
	Id NullableString `json:"id,omitempty"`
	// The chain of folders leading to the destination, outermost first and the destination itself last, with folders  the caller cannot read left out. An answer that reports a stored part carries the destination folder alone  instead of the whole chain.
	Path []string `json:"path,omitempty"`
	// The moment the upload was reserved, in UTC.
	Created *time.Time `json:"created,omitempty"`
	// The moment the reservation lapses and the parts buffered for it are dropped, in UTC. It is a gap rather than a  deadline for the whole transfer: every accepted part pushes it twelve hours past that part, so only a long  silence loses the upload.
	Expired *time.Time `json:"expired,omitempty"`
	// The absolute address of the separate chunk handler that also accepts the parts of this upload, kept for  clients written against it. A caller working through this API does not need it and sends the parts to the  session operations instead.
	Location NullableString `json:"location,omitempty"`
	// The size in bytes that was declared when the upload was reserved, echoed back. It is what the arriving parts  are counted against to decide the file is complete, not the amount received so far.
	BytesTotal *int64 `json:"bytes_total,omitempty"`
}

// NewThirdPartyChunkedUploadSessionResponse instantiates a new ThirdPartyChunkedUploadSessionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThirdPartyChunkedUploadSessionResponse() *ThirdPartyChunkedUploadSessionResponse {
	this := ThirdPartyChunkedUploadSessionResponse{}
	return &this
}

// NewThirdPartyChunkedUploadSessionResponseWithDefaults instantiates a new ThirdPartyChunkedUploadSessionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThirdPartyChunkedUploadSessionResponseWithDefaults() *ThirdPartyChunkedUploadSessionResponse {
	this := ThirdPartyChunkedUploadSessionResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyChunkedUploadSessionResponse) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyChunkedUploadSessionResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ThirdPartyChunkedUploadSessionResponse) IsIdSet() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *ThirdPartyChunkedUploadSessionResponse) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ThirdPartyChunkedUploadSessionResponse) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ThirdPartyChunkedUploadSessionResponse) UnsetId() {
	o.Id.Unset()
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyChunkedUploadSessionResponse) GetPath() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyChunkedUploadSessionResponse) GetPathOk() ([]string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ThirdPartyChunkedUploadSessionResponse) IsPathSet() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given []string and assigns it to the Path field.
func (o *ThirdPartyChunkedUploadSessionResponse) SetPath(v []string) {
	o.Path = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ThirdPartyChunkedUploadSessionResponse) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyChunkedUploadSessionResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ThirdPartyChunkedUploadSessionResponse) IsCreatedSet() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *ThirdPartyChunkedUploadSessionResponse) SetCreated(v time.Time) {
	o.Created = &v
}

// GetExpired returns the Expired field value if set, zero value otherwise.
func (o *ThirdPartyChunkedUploadSessionResponse) GetExpired() time.Time {
	if o == nil || IsNil(o.Expired) {
		var ret time.Time
		return ret
	}
	return *o.Expired
}

// GetExpiredOk returns a tuple with the Expired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyChunkedUploadSessionResponse) GetExpiredOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Expired) {
		return nil, false
	}
	return o.Expired, true
}

// HasExpired returns a boolean if a field has been set.
func (o *ThirdPartyChunkedUploadSessionResponse) IsExpiredSet() bool {
	if o != nil && !IsNil(o.Expired) {
		return true
	}

	return false
}

// SetExpired gets a reference to the given time.Time and assigns it to the Expired field.
func (o *ThirdPartyChunkedUploadSessionResponse) SetExpired(v time.Time) {
	o.Expired = &v
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThirdPartyChunkedUploadSessionResponse) GetLocation() string {
	if o == nil || IsNil(o.Location.Get()) {
		var ret string
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThirdPartyChunkedUploadSessionResponse) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *ThirdPartyChunkedUploadSessionResponse) IsLocationSet() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableString and assigns it to the Location field.
func (o *ThirdPartyChunkedUploadSessionResponse) SetLocation(v string) {
	o.Location.Set(&v)
}
// SetLocationNil sets the value for Location to be an explicit nil
func (o *ThirdPartyChunkedUploadSessionResponse) SetLocationNil() {
	o.Location.Set(nil)
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *ThirdPartyChunkedUploadSessionResponse) UnsetLocation() {
	o.Location.Unset()
}

// GetBytesTotal returns the BytesTotal field value if set, zero value otherwise.
func (o *ThirdPartyChunkedUploadSessionResponse) GetBytesTotal() int64 {
	if o == nil || IsNil(o.BytesTotal) {
		var ret int64
		return ret
	}
	return *o.BytesTotal
}

// GetBytesTotalOk returns a tuple with the BytesTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyChunkedUploadSessionResponse) GetBytesTotalOk() (*int64, bool) {
	if o == nil || IsNil(o.BytesTotal) {
		return nil, false
	}
	return o.BytesTotal, true
}

// HasBytesTotal returns a boolean if a field has been set.
func (o *ThirdPartyChunkedUploadSessionResponse) IsBytesTotalSet() bool {
	if o != nil && !IsNil(o.BytesTotal) {
		return true
	}

	return false
}

// SetBytesTotal gets a reference to the given int64 and assigns it to the BytesTotal field.
func (o *ThirdPartyChunkedUploadSessionResponse) SetBytesTotal(v int64) {
	o.BytesTotal = &v
}

func (o ThirdPartyChunkedUploadSessionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThirdPartyChunkedUploadSessionResponse) ToMap() (map[string]interface{}, error) {
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

type NullableThirdPartyChunkedUploadSessionResponse struct {
	value *ThirdPartyChunkedUploadSessionResponse
	isSet bool
}

func (v NullableThirdPartyChunkedUploadSessionResponse) Get() *ThirdPartyChunkedUploadSessionResponse {
	return v.value
}

func (v *NullableThirdPartyChunkedUploadSessionResponse) Set(val *ThirdPartyChunkedUploadSessionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableThirdPartyChunkedUploadSessionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableThirdPartyChunkedUploadSessionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThirdPartyChunkedUploadSessionResponse(val *ThirdPartyChunkedUploadSessionResponse) *NullableThirdPartyChunkedUploadSessionResponse {
	return &NullableThirdPartyChunkedUploadSessionResponse{value: val, isSet: true}
}

func (v NullableThirdPartyChunkedUploadSessionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThirdPartyChunkedUploadSessionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

