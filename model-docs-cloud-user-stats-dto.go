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

// checks if the DocsCloudUserStatsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocsCloudUserStatsDto{}

// DocsCloudUserStatsDto Represents the usage statistics of a single Docs Connect user category (editor or viewer).
type DocsCloudUserStatsDto struct {
	// The number of active users.
	Active *int32 `json:"active,omitempty"`
	// The number of internal users.
	Internal *int32 `json:"internal,omitempty"`
	// The number of external users.
	External *int32 `json:"external,omitempty"`
	// The number of remaining users before the limit is reached.
	Remaining *int32 `json:"remaining,omitempty"`
	// Whether the number of remaining users is critically low.
	CriticalRemaining *bool `json:"criticalRemaining,omitempty"`
}

// NewDocsCloudUserStatsDto instantiates a new DocsCloudUserStatsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocsCloudUserStatsDto() *DocsCloudUserStatsDto {
	this := DocsCloudUserStatsDto{}
	return &this
}

// NewDocsCloudUserStatsDtoWithDefaults instantiates a new DocsCloudUserStatsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocsCloudUserStatsDtoWithDefaults() *DocsCloudUserStatsDto {
	this := DocsCloudUserStatsDto{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *DocsCloudUserStatsDto) GetActive() int32 {
	if o == nil || IsNil(o.Active) {
		var ret int32
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudUserStatsDto) GetActiveOk() (*int32, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *DocsCloudUserStatsDto) IsActiveSet() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given int32 and assigns it to the Active field.
func (o *DocsCloudUserStatsDto) SetActive(v int32) {
	o.Active = &v
}

// GetInternal returns the Internal field value if set, zero value otherwise.
func (o *DocsCloudUserStatsDto) GetInternal() int32 {
	if o == nil || IsNil(o.Internal) {
		var ret int32
		return ret
	}
	return *o.Internal
}

// GetInternalOk returns a tuple with the Internal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudUserStatsDto) GetInternalOk() (*int32, bool) {
	if o == nil || IsNil(o.Internal) {
		return nil, false
	}
	return o.Internal, true
}

// HasInternal returns a boolean if a field has been set.
func (o *DocsCloudUserStatsDto) IsInternalSet() bool {
	if o != nil && !IsNil(o.Internal) {
		return true
	}

	return false
}

// SetInternal gets a reference to the given int32 and assigns it to the Internal field.
func (o *DocsCloudUserStatsDto) SetInternal(v int32) {
	o.Internal = &v
}

// GetExternal returns the External field value if set, zero value otherwise.
func (o *DocsCloudUserStatsDto) GetExternal() int32 {
	if o == nil || IsNil(o.External) {
		var ret int32
		return ret
	}
	return *o.External
}

// GetExternalOk returns a tuple with the External field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudUserStatsDto) GetExternalOk() (*int32, bool) {
	if o == nil || IsNil(o.External) {
		return nil, false
	}
	return o.External, true
}

// HasExternal returns a boolean if a field has been set.
func (o *DocsCloudUserStatsDto) IsExternalSet() bool {
	if o != nil && !IsNil(o.External) {
		return true
	}

	return false
}

// SetExternal gets a reference to the given int32 and assigns it to the External field.
func (o *DocsCloudUserStatsDto) SetExternal(v int32) {
	o.External = &v
}

// GetRemaining returns the Remaining field value if set, zero value otherwise.
func (o *DocsCloudUserStatsDto) GetRemaining() int32 {
	if o == nil || IsNil(o.Remaining) {
		var ret int32
		return ret
	}
	return *o.Remaining
}

// GetRemainingOk returns a tuple with the Remaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudUserStatsDto) GetRemainingOk() (*int32, bool) {
	if o == nil || IsNil(o.Remaining) {
		return nil, false
	}
	return o.Remaining, true
}

// HasRemaining returns a boolean if a field has been set.
func (o *DocsCloudUserStatsDto) IsRemainingSet() bool {
	if o != nil && !IsNil(o.Remaining) {
		return true
	}

	return false
}

// SetRemaining gets a reference to the given int32 and assigns it to the Remaining field.
func (o *DocsCloudUserStatsDto) SetRemaining(v int32) {
	o.Remaining = &v
}

// GetCriticalRemaining returns the CriticalRemaining field value if set, zero value otherwise.
func (o *DocsCloudUserStatsDto) GetCriticalRemaining() bool {
	if o == nil || IsNil(o.CriticalRemaining) {
		var ret bool
		return ret
	}
	return *o.CriticalRemaining
}

// GetCriticalRemainingOk returns a tuple with the CriticalRemaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudUserStatsDto) GetCriticalRemainingOk() (*bool, bool) {
	if o == nil || IsNil(o.CriticalRemaining) {
		return nil, false
	}
	return o.CriticalRemaining, true
}

// HasCriticalRemaining returns a boolean if a field has been set.
func (o *DocsCloudUserStatsDto) IsCriticalRemainingSet() bool {
	if o != nil && !IsNil(o.CriticalRemaining) {
		return true
	}

	return false
}

// SetCriticalRemaining gets a reference to the given bool and assigns it to the CriticalRemaining field.
func (o *DocsCloudUserStatsDto) SetCriticalRemaining(v bool) {
	o.CriticalRemaining = &v
}

func (o DocsCloudUserStatsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocsCloudUserStatsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Internal) {
		toSerialize["internal"] = o.Internal
	}
	if !IsNil(o.External) {
		toSerialize["external"] = o.External
	}
	if !IsNil(o.Remaining) {
		toSerialize["remaining"] = o.Remaining
	}
	if !IsNil(o.CriticalRemaining) {
		toSerialize["criticalRemaining"] = o.CriticalRemaining
	}
	return toSerialize, nil
}

type NullableDocsCloudUserStatsDto struct {
	value *DocsCloudUserStatsDto
	isSet bool
}

func (v NullableDocsCloudUserStatsDto) Get() *DocsCloudUserStatsDto {
	return v.value
}

func (v *NullableDocsCloudUserStatsDto) Set(val *DocsCloudUserStatsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDocsCloudUserStatsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDocsCloudUserStatsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocsCloudUserStatsDto(val *DocsCloudUserStatsDto) *NullableDocsCloudUserStatsDto {
	return &NullableDocsCloudUserStatsDto{value: val, isSet: true}
}

func (v NullableDocsCloudUserStatsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocsCloudUserStatsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

