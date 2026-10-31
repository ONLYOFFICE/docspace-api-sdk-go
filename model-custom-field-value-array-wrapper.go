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

// checks if the CustomFieldValueArrayWrapper type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomFieldValueArrayWrapper{}

// CustomFieldValueArrayWrapper The successful API response containing the list of CustomFieldValueDto objects.
type CustomFieldValueArrayWrapper struct {
	// The list of CustomFieldValueDto objects returned by the operation.
	Response []CustomFieldValueDto `json:"response,omitempty"`
	// How many items this response carries: the length of the page for a list, 1 for a single object
	Count *int32 `json:"count,omitempty"`
	// How many items the whole list holds across all pages, for working out how many pages to request. Sent only by the operations that page their results, and left out of every other response
	Total *int64 `json:"total,omitempty"`
	// List of links related to the response
	Links []GetPortalPrices200ResponseLinksInner `json:"links,omitempty"`
	// HTTP status code of the response
	Status *int32 `json:"status,omitempty"`
	// HTTP status code of the response (duplicate of status)
	StatusCode *int32 `json:"statusCode,omitempty"`
}

// NewCustomFieldValueArrayWrapper instantiates a new CustomFieldValueArrayWrapper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomFieldValueArrayWrapper() *CustomFieldValueArrayWrapper {
	this := CustomFieldValueArrayWrapper{}
	return &this
}

// NewCustomFieldValueArrayWrapperWithDefaults instantiates a new CustomFieldValueArrayWrapper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomFieldValueArrayWrapperWithDefaults() *CustomFieldValueArrayWrapper {
	this := CustomFieldValueArrayWrapper{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *CustomFieldValueArrayWrapper) GetResponse() []CustomFieldValueDto {
	if o == nil || IsNil(o.Response) {
		var ret []CustomFieldValueDto
		return ret
	}
	return o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFieldValueArrayWrapper) GetResponseOk() ([]CustomFieldValueDto, bool) {
	if o == nil || IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *CustomFieldValueArrayWrapper) IsResponseSet() bool {
	if o != nil && !IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given []CustomFieldValueDto and assigns it to the Response field.
func (o *CustomFieldValueArrayWrapper) SetResponse(v []CustomFieldValueDto) {
	o.Response = v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *CustomFieldValueArrayWrapper) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFieldValueArrayWrapper) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *CustomFieldValueArrayWrapper) IsCountSet() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *CustomFieldValueArrayWrapper) SetCount(v int32) {
	o.Count = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *CustomFieldValueArrayWrapper) GetTotal() int64 {
	if o == nil || IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFieldValueArrayWrapper) GetTotalOk() (*int64, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *CustomFieldValueArrayWrapper) IsTotalSet() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *CustomFieldValueArrayWrapper) SetTotal(v int64) {
	o.Total = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CustomFieldValueArrayWrapper) GetLinks() []GetPortalPrices200ResponseLinksInner {
	if o == nil || IsNil(o.Links) {
		var ret []GetPortalPrices200ResponseLinksInner
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFieldValueArrayWrapper) GetLinksOk() ([]GetPortalPrices200ResponseLinksInner, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CustomFieldValueArrayWrapper) IsLinksSet() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []GetPortalPrices200ResponseLinksInner and assigns it to the Links field.
func (o *CustomFieldValueArrayWrapper) SetLinks(v []GetPortalPrices200ResponseLinksInner) {
	o.Links = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CustomFieldValueArrayWrapper) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFieldValueArrayWrapper) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CustomFieldValueArrayWrapper) IsStatusSet() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *CustomFieldValueArrayWrapper) SetStatus(v int32) {
	o.Status = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *CustomFieldValueArrayWrapper) GetStatusCode() int32 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFieldValueArrayWrapper) GetStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *CustomFieldValueArrayWrapper) IsStatusCodeSet() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *CustomFieldValueArrayWrapper) SetStatusCode(v int32) {
	o.StatusCode = &v
}

func (o CustomFieldValueArrayWrapper) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomFieldValueArrayWrapper) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Response) {
		toSerialize["response"] = o.Response
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.StatusCode) {
		toSerialize["statusCode"] = o.StatusCode
	}
	return toSerialize, nil
}

type NullableCustomFieldValueArrayWrapper struct {
	value *CustomFieldValueArrayWrapper
	isSet bool
}

func (v NullableCustomFieldValueArrayWrapper) Get() *CustomFieldValueArrayWrapper {
	return v.value
}

func (v *NullableCustomFieldValueArrayWrapper) Set(val *CustomFieldValueArrayWrapper) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomFieldValueArrayWrapper) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomFieldValueArrayWrapper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomFieldValueArrayWrapper(val *CustomFieldValueArrayWrapper) *NullableCustomFieldValueArrayWrapper {
	return &NullableCustomFieldValueArrayWrapper{value: val, isSet: true}
}

func (v NullableCustomFieldValueArrayWrapper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomFieldValueArrayWrapper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

