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

// checks if the PageableResponseClientInfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PageableResponseClientInfoResponse{}

// PageableResponseClientInfoResponse The response containing paginated client information.
type PageableResponseClientInfoResponse struct {
	// The paginated client data.
	Data map[string]interface{} `json:"data,omitempty"`
	// The maximum number of results returned per page.
	Limit *int32 `json:"limit,omitempty"`
	// The identifier of the last retrieved client.
	LastClientId *string `json:"last_client_id,omitempty"`
	// The creation date of the last retrieved client.
	LastCreatedOn *time.Time `json:"last_created_on,omitempty"`
}

// NewPageableResponseClientInfoResponse instantiates a new PageableResponseClientInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPageableResponseClientInfoResponse() *PageableResponseClientInfoResponse {
	this := PageableResponseClientInfoResponse{}
	return &this
}

// NewPageableResponseClientInfoResponseWithDefaults instantiates a new PageableResponseClientInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageableResponseClientInfoResponseWithDefaults() *PageableResponseClientInfoResponse {
	this := PageableResponseClientInfoResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PageableResponseClientInfoResponse) GetData() map[string]interface{} {
	if o == nil || IsNil(o.Data) {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageableResponseClientInfoResponse) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PageableResponseClientInfoResponse) IsDataSet() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *PageableResponseClientInfoResponse) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *PageableResponseClientInfoResponse) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageableResponseClientInfoResponse) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *PageableResponseClientInfoResponse) IsLimitSet() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *PageableResponseClientInfoResponse) SetLimit(v int32) {
	o.Limit = &v
}

// GetLastClientId returns the LastClientId field value if set, zero value otherwise.
func (o *PageableResponseClientInfoResponse) GetLastClientId() string {
	if o == nil || IsNil(o.LastClientId) {
		var ret string
		return ret
	}
	return *o.LastClientId
}

// GetLastClientIdOk returns a tuple with the LastClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageableResponseClientInfoResponse) GetLastClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.LastClientId) {
		return nil, false
	}
	return o.LastClientId, true
}

// HasLastClientId returns a boolean if a field has been set.
func (o *PageableResponseClientInfoResponse) IsLastClientIdSet() bool {
	if o != nil && !IsNil(o.LastClientId) {
		return true
	}

	return false
}

// SetLastClientId gets a reference to the given string and assigns it to the LastClientId field.
func (o *PageableResponseClientInfoResponse) SetLastClientId(v string) {
	o.LastClientId = &v
}

// GetLastCreatedOn returns the LastCreatedOn field value if set, zero value otherwise.
func (o *PageableResponseClientInfoResponse) GetLastCreatedOn() time.Time {
	if o == nil || IsNil(o.LastCreatedOn) {
		var ret time.Time
		return ret
	}
	return *o.LastCreatedOn
}

// GetLastCreatedOnOk returns a tuple with the LastCreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageableResponseClientInfoResponse) GetLastCreatedOnOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastCreatedOn) {
		return nil, false
	}
	return o.LastCreatedOn, true
}

// HasLastCreatedOn returns a boolean if a field has been set.
func (o *PageableResponseClientInfoResponse) IsLastCreatedOnSet() bool {
	if o != nil && !IsNil(o.LastCreatedOn) {
		return true
	}

	return false
}

// SetLastCreatedOn gets a reference to the given time.Time and assigns it to the LastCreatedOn field.
func (o *PageableResponseClientInfoResponse) SetLastCreatedOn(v time.Time) {
	o.LastCreatedOn = &v
}

func (o PageableResponseClientInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PageableResponseClientInfoResponse) ToMap() (map[string]interface{}, error) {
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

type NullablePageableResponseClientInfoResponse struct {
	value *PageableResponseClientInfoResponse
	isSet bool
}

func (v NullablePageableResponseClientInfoResponse) Get() *PageableResponseClientInfoResponse {
	return v.value
}

func (v *NullablePageableResponseClientInfoResponse) Set(val *PageableResponseClientInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePageableResponseClientInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePageableResponseClientInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePageableResponseClientInfoResponse(val *PageableResponseClientInfoResponse) *NullablePageableResponseClientInfoResponse {
	return &NullablePageableResponseClientInfoResponse{value: val, isSet: true}
}

func (v NullablePageableResponseClientInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePageableResponseClientInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

