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

// checks if the PageableModificationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PageableModificationResponse{}

// PageableModificationResponse One page of results ordered by modification time, together with the cursor that asks for the next page.
type PageableModificationResponse struct {
	Data interface{} `json:"data,omitempty"`
	// The page size that was applied to this request, between 1 and 50.
	Limit *int32 `json:"limit,omitempty"`
	// The cursor to send back as last_modified_on to ask for the next page. It is null when the page is empty.
	LastModifiedOn *time.Time `json:"last_modified_on,omitempty"`
}

// NewPageableModificationResponse instantiates a new PageableModificationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPageableModificationResponse() *PageableModificationResponse {
	this := PageableModificationResponse{}
	return &this
}

// NewPageableModificationResponseWithDefaults instantiates a new PageableModificationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageableModificationResponseWithDefaults() *PageableModificationResponse {
	this := PageableModificationResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PageableModificationResponse) GetData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PageableModificationResponse) GetDataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PageableModificationResponse) IsDataSet() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given interface{} and assigns it to the Data field.
func (o *PageableModificationResponse) SetData(v interface{}) {
	o.Data = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *PageableModificationResponse) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageableModificationResponse) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *PageableModificationResponse) IsLimitSet() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *PageableModificationResponse) SetLimit(v int32) {
	o.Limit = &v
}

// GetLastModifiedOn returns the LastModifiedOn field value if set, zero value otherwise.
func (o *PageableModificationResponse) GetLastModifiedOn() time.Time {
	if o == nil || IsNil(o.LastModifiedOn) {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedOn
}

// GetLastModifiedOnOk returns a tuple with the LastModifiedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageableModificationResponse) GetLastModifiedOnOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModifiedOn) {
		return nil, false
	}
	return o.LastModifiedOn, true
}

// HasLastModifiedOn returns a boolean if a field has been set.
func (o *PageableModificationResponse) IsLastModifiedOnSet() bool {
	if o != nil && !IsNil(o.LastModifiedOn) {
		return true
	}

	return false
}

// SetLastModifiedOn gets a reference to the given time.Time and assigns it to the LastModifiedOn field.
func (o *PageableModificationResponse) SetLastModifiedOn(v time.Time) {
	o.LastModifiedOn = &v
}

func (o PageableModificationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PageableModificationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.LastModifiedOn) {
		toSerialize["last_modified_on"] = o.LastModifiedOn
	}
	return toSerialize, nil
}

type NullablePageableModificationResponse struct {
	value *PageableModificationResponse
	isSet bool
}

func (v NullablePageableModificationResponse) Get() *PageableModificationResponse {
	return v.value
}

func (v *NullablePageableModificationResponse) Set(val *PageableModificationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePageableModificationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePageableModificationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePageableModificationResponse(val *PageableModificationResponse) *NullablePageableModificationResponse {
	return &NullablePageableModificationResponse{value: val, isSet: true}
}

func (v NullablePageableModificationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePageableModificationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

