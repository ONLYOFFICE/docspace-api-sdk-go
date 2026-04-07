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

// checks if the IsDefaultWhiteLabelLogosWrapper type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IsDefaultWhiteLabelLogosWrapper{}

// IsDefaultWhiteLabelLogosWrapper struct for IsDefaultWhiteLabelLogosWrapper
type IsDefaultWhiteLabelLogosWrapper struct {
	Response *IsDefaultWhiteLabelLogosDto `json:"response,omitempty"`
	// The total number of items in the response
	Count *int32 `json:"count,omitempty"`
	// List of links related to the response
	Links []GetPortalPrices200ResponseLinksInner `json:"links,omitempty"`
	// HTTP status code of the response
	Status *int32 `json:"status,omitempty"`
	// HTTP status code of the response (duplicate of status)
	StatusCode *int32 `json:"statusCode,omitempty"`
}

// NewIsDefaultWhiteLabelLogosWrapper instantiates a new IsDefaultWhiteLabelLogosWrapper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIsDefaultWhiteLabelLogosWrapper() *IsDefaultWhiteLabelLogosWrapper {
	this := IsDefaultWhiteLabelLogosWrapper{}
	return &this
}

// NewIsDefaultWhiteLabelLogosWrapperWithDefaults instantiates a new IsDefaultWhiteLabelLogosWrapper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIsDefaultWhiteLabelLogosWrapperWithDefaults() *IsDefaultWhiteLabelLogosWrapper {
	this := IsDefaultWhiteLabelLogosWrapper{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *IsDefaultWhiteLabelLogosWrapper) GetResponse() IsDefaultWhiteLabelLogosDto {
	if o == nil || IsNil(o.Response) {
		var ret IsDefaultWhiteLabelLogosDto
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsDefaultWhiteLabelLogosWrapper) GetResponseOk() (*IsDefaultWhiteLabelLogosDto, bool) {
	if o == nil || IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *IsDefaultWhiteLabelLogosWrapper) IsResponseSet() bool {
	if o != nil && !IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given IsDefaultWhiteLabelLogosDto and assigns it to the Response field.
func (o *IsDefaultWhiteLabelLogosWrapper) SetResponse(v IsDefaultWhiteLabelLogosDto) {
	o.Response = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *IsDefaultWhiteLabelLogosWrapper) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsDefaultWhiteLabelLogosWrapper) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *IsDefaultWhiteLabelLogosWrapper) IsCountSet() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *IsDefaultWhiteLabelLogosWrapper) SetCount(v int32) {
	o.Count = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *IsDefaultWhiteLabelLogosWrapper) GetLinks() []GetPortalPrices200ResponseLinksInner {
	if o == nil || IsNil(o.Links) {
		var ret []GetPortalPrices200ResponseLinksInner
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsDefaultWhiteLabelLogosWrapper) GetLinksOk() ([]GetPortalPrices200ResponseLinksInner, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *IsDefaultWhiteLabelLogosWrapper) IsLinksSet() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []GetPortalPrices200ResponseLinksInner and assigns it to the Links field.
func (o *IsDefaultWhiteLabelLogosWrapper) SetLinks(v []GetPortalPrices200ResponseLinksInner) {
	o.Links = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IsDefaultWhiteLabelLogosWrapper) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsDefaultWhiteLabelLogosWrapper) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IsDefaultWhiteLabelLogosWrapper) IsStatusSet() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *IsDefaultWhiteLabelLogosWrapper) SetStatus(v int32) {
	o.Status = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *IsDefaultWhiteLabelLogosWrapper) GetStatusCode() int32 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsDefaultWhiteLabelLogosWrapper) GetStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *IsDefaultWhiteLabelLogosWrapper) IsStatusCodeSet() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *IsDefaultWhiteLabelLogosWrapper) SetStatusCode(v int32) {
	o.StatusCode = &v
}

func (o IsDefaultWhiteLabelLogosWrapper) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IsDefaultWhiteLabelLogosWrapper) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Response) {
		toSerialize["response"] = o.Response
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
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

type NullableIsDefaultWhiteLabelLogosWrapper struct {
	value *IsDefaultWhiteLabelLogosWrapper
	isSet bool
}

func (v NullableIsDefaultWhiteLabelLogosWrapper) Get() *IsDefaultWhiteLabelLogosWrapper {
	return v.value
}

func (v *NullableIsDefaultWhiteLabelLogosWrapper) Set(val *IsDefaultWhiteLabelLogosWrapper) {
	v.value = val
	v.isSet = true
}

func (v NullableIsDefaultWhiteLabelLogosWrapper) IsSet() bool {
	return v.isSet
}

func (v *NullableIsDefaultWhiteLabelLogosWrapper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIsDefaultWhiteLabelLogosWrapper(val *IsDefaultWhiteLabelLogosWrapper) *NullableIsDefaultWhiteLabelLogosWrapper {
	return &NullableIsDefaultWhiteLabelLogosWrapper{value: val, isSet: true}
}

func (v NullableIsDefaultWhiteLabelLogosWrapper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIsDefaultWhiteLabelLogosWrapper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

