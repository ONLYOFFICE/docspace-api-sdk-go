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

// checks if the PageableClientInfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PageableClientInfoResponse{}

// PageableClientInfoResponse One page of consent-facing client info together with the next-page cursor.
type PageableClientInfoResponse struct {
	// The items on this page, at most as many as the requested limit. An empty array means there is nothing further to read.
	Data []ClientInfoResponse `json:"data,omitempty"`
	// The page size that was applied to this request, between 1 and 50.
	Limit *int32 `json:"limit,omitempty"`
	// The cursor to send back as last_client_id to ask for the next page, together with last_created_on. It is null when the page is empty.
	LastClientId *string `json:"last_client_id,omitempty"`
	// The cursor to send back as last_created_on to ask for the next page, together with last_client_id. It is null when the page is empty.
	LastCreatedOn *time.Time `json:"last_created_on,omitempty"`
}

// NewPageableClientInfoResponse instantiates a new PageableClientInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPageableClientInfoResponse() *PageableClientInfoResponse {
	this := PageableClientInfoResponse{}
	return &this
}

// NewPageableClientInfoResponseWithDefaults instantiates a new PageableClientInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageableClientInfoResponseWithDefaults() *PageableClientInfoResponse {
	this := PageableClientInfoResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PageableClientInfoResponse) GetData() []ClientInfoResponse {
	if o == nil || IsNil(o.Data) {
		var ret []ClientInfoResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageableClientInfoResponse) GetDataOk() ([]ClientInfoResponse, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PageableClientInfoResponse) IsDataSet() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []ClientInfoResponse and assigns it to the Data field.
func (o *PageableClientInfoResponse) SetData(v []ClientInfoResponse) {
	o.Data = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *PageableClientInfoResponse) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageableClientInfoResponse) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *PageableClientInfoResponse) IsLimitSet() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *PageableClientInfoResponse) SetLimit(v int32) {
	o.Limit = &v
}

// GetLastClientId returns the LastClientId field value if set, zero value otherwise.
func (o *PageableClientInfoResponse) GetLastClientId() string {
	if o == nil || IsNil(o.LastClientId) {
		var ret string
		return ret
	}
	return *o.LastClientId
}

// GetLastClientIdOk returns a tuple with the LastClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageableClientInfoResponse) GetLastClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.LastClientId) {
		return nil, false
	}
	return o.LastClientId, true
}

// HasLastClientId returns a boolean if a field has been set.
func (o *PageableClientInfoResponse) IsLastClientIdSet() bool {
	if o != nil && !IsNil(o.LastClientId) {
		return true
	}

	return false
}

// SetLastClientId gets a reference to the given string and assigns it to the LastClientId field.
func (o *PageableClientInfoResponse) SetLastClientId(v string) {
	o.LastClientId = &v
}

// GetLastCreatedOn returns the LastCreatedOn field value if set, zero value otherwise.
func (o *PageableClientInfoResponse) GetLastCreatedOn() time.Time {
	if o == nil || IsNil(o.LastCreatedOn) {
		var ret time.Time
		return ret
	}
	return *o.LastCreatedOn
}

// GetLastCreatedOnOk returns a tuple with the LastCreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageableClientInfoResponse) GetLastCreatedOnOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastCreatedOn) {
		return nil, false
	}
	return o.LastCreatedOn, true
}

// HasLastCreatedOn returns a boolean if a field has been set.
func (o *PageableClientInfoResponse) IsLastCreatedOnSet() bool {
	if o != nil && !IsNil(o.LastCreatedOn) {
		return true
	}

	return false
}

// SetLastCreatedOn gets a reference to the given time.Time and assigns it to the LastCreatedOn field.
func (o *PageableClientInfoResponse) SetLastCreatedOn(v time.Time) {
	o.LastCreatedOn = &v
}

func (o PageableClientInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PageableClientInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.LastClientId) {
		toSerialize["last_client_id"] = o.LastClientId
	}
	if !IsNil(o.LastCreatedOn) {
		toSerialize["last_created_on"] = o.LastCreatedOn
	}
	return toSerialize, nil
}

type NullablePageableClientInfoResponse struct {
	value *PageableClientInfoResponse
	isSet bool
}

func (v NullablePageableClientInfoResponse) Get() *PageableClientInfoResponse {
	return v.value
}

func (v *NullablePageableClientInfoResponse) Set(val *PageableClientInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePageableClientInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePageableClientInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePageableClientInfoResponse(val *PageableClientInfoResponse) *NullablePageableClientInfoResponse {
	return &NullablePageableClientInfoResponse{value: val, isSet: true}
}

func (v NullablePageableClientInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePageableClientInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

