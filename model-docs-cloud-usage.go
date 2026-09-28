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

// checks if the DocsCloudUsage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocsCloudUsage{}

// DocsCloudUsage Represents the usage statistics of a Docs Connect tenant.
type DocsCloudUsage struct {
	// The date and time the usage statistics are counted from.
	Since *time.Time `json:"since,omitempty"`
	// The number of active users.
	ActiveCount *int32 `json:"activeCount,omitempty"`
}

// NewDocsCloudUsage instantiates a new DocsCloudUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocsCloudUsage() *DocsCloudUsage {
	this := DocsCloudUsage{}
	return &this
}

// NewDocsCloudUsageWithDefaults instantiates a new DocsCloudUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocsCloudUsageWithDefaults() *DocsCloudUsage {
	this := DocsCloudUsage{}
	return &this
}

// GetSince returns the Since field value if set, zero value otherwise.
func (o *DocsCloudUsage) GetSince() time.Time {
	if o == nil || IsNil(o.Since) {
		var ret time.Time
		return ret
	}
	return *o.Since
}

// GetSinceOk returns a tuple with the Since field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudUsage) GetSinceOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Since) {
		return nil, false
	}
	return o.Since, true
}

// HasSince returns a boolean if a field has been set.
func (o *DocsCloudUsage) IsSinceSet() bool {
	if o != nil && !IsNil(o.Since) {
		return true
	}

	return false
}

// SetSince gets a reference to the given time.Time and assigns it to the Since field.
func (o *DocsCloudUsage) SetSince(v time.Time) {
	o.Since = &v
}

// GetActiveCount returns the ActiveCount field value if set, zero value otherwise.
func (o *DocsCloudUsage) GetActiveCount() int32 {
	if o == nil || IsNil(o.ActiveCount) {
		var ret int32
		return ret
	}
	return *o.ActiveCount
}

// GetActiveCountOk returns a tuple with the ActiveCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudUsage) GetActiveCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ActiveCount) {
		return nil, false
	}
	return o.ActiveCount, true
}

// HasActiveCount returns a boolean if a field has been set.
func (o *DocsCloudUsage) IsActiveCountSet() bool {
	if o != nil && !IsNil(o.ActiveCount) {
		return true
	}

	return false
}

// SetActiveCount gets a reference to the given int32 and assigns it to the ActiveCount field.
func (o *DocsCloudUsage) SetActiveCount(v int32) {
	o.ActiveCount = &v
}

func (o DocsCloudUsage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocsCloudUsage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Since) {
		toSerialize["since"] = o.Since
	}
	if !IsNil(o.ActiveCount) {
		toSerialize["activeCount"] = o.ActiveCount
	}
	return toSerialize, nil
}

type NullableDocsCloudUsage struct {
	value *DocsCloudUsage
	isSet bool
}

func (v NullableDocsCloudUsage) Get() *DocsCloudUsage {
	return v.value
}

func (v *NullableDocsCloudUsage) Set(val *DocsCloudUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableDocsCloudUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableDocsCloudUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocsCloudUsage(val *DocsCloudUsage) *NullableDocsCloudUsage {
	return &NullableDocsCloudUsage{value: val, isSet: true}
}

func (v NullableDocsCloudUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocsCloudUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

