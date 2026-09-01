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

// checks if the DocsCloudUserStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocsCloudUserStats{}

// DocsCloudUserStats Represents the usage statistics of a single DocsCloud user category (editor or viewer).
type DocsCloudUserStats struct {
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

// NewDocsCloudUserStats instantiates a new DocsCloudUserStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocsCloudUserStats() *DocsCloudUserStats {
	this := DocsCloudUserStats{}
	return &this
}

// NewDocsCloudUserStatsWithDefaults instantiates a new DocsCloudUserStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocsCloudUserStatsWithDefaults() *DocsCloudUserStats {
	this := DocsCloudUserStats{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *DocsCloudUserStats) GetActive() int32 {
	if o == nil || IsNil(o.Active) {
		var ret int32
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudUserStats) GetActiveOk() (*int32, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *DocsCloudUserStats) IsActiveSet() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given int32 and assigns it to the Active field.
func (o *DocsCloudUserStats) SetActive(v int32) {
	o.Active = &v
}

// GetInternal returns the Internal field value if set, zero value otherwise.
func (o *DocsCloudUserStats) GetInternal() int32 {
	if o == nil || IsNil(o.Internal) {
		var ret int32
		return ret
	}
	return *o.Internal
}

// GetInternalOk returns a tuple with the Internal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudUserStats) GetInternalOk() (*int32, bool) {
	if o == nil || IsNil(o.Internal) {
		return nil, false
	}
	return o.Internal, true
}

// HasInternal returns a boolean if a field has been set.
func (o *DocsCloudUserStats) IsInternalSet() bool {
	if o != nil && !IsNil(o.Internal) {
		return true
	}

	return false
}

// SetInternal gets a reference to the given int32 and assigns it to the Internal field.
func (o *DocsCloudUserStats) SetInternal(v int32) {
	o.Internal = &v
}

// GetExternal returns the External field value if set, zero value otherwise.
func (o *DocsCloudUserStats) GetExternal() int32 {
	if o == nil || IsNil(o.External) {
		var ret int32
		return ret
	}
	return *o.External
}

// GetExternalOk returns a tuple with the External field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudUserStats) GetExternalOk() (*int32, bool) {
	if o == nil || IsNil(o.External) {
		return nil, false
	}
	return o.External, true
}

// HasExternal returns a boolean if a field has been set.
func (o *DocsCloudUserStats) IsExternalSet() bool {
	if o != nil && !IsNil(o.External) {
		return true
	}

	return false
}

// SetExternal gets a reference to the given int32 and assigns it to the External field.
func (o *DocsCloudUserStats) SetExternal(v int32) {
	o.External = &v
}

// GetRemaining returns the Remaining field value if set, zero value otherwise.
func (o *DocsCloudUserStats) GetRemaining() int32 {
	if o == nil || IsNil(o.Remaining) {
		var ret int32
		return ret
	}
	return *o.Remaining
}

// GetRemainingOk returns a tuple with the Remaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudUserStats) GetRemainingOk() (*int32, bool) {
	if o == nil || IsNil(o.Remaining) {
		return nil, false
	}
	return o.Remaining, true
}

// HasRemaining returns a boolean if a field has been set.
func (o *DocsCloudUserStats) IsRemainingSet() bool {
	if o != nil && !IsNil(o.Remaining) {
		return true
	}

	return false
}

// SetRemaining gets a reference to the given int32 and assigns it to the Remaining field.
func (o *DocsCloudUserStats) SetRemaining(v int32) {
	o.Remaining = &v
}

// GetCriticalRemaining returns the CriticalRemaining field value if set, zero value otherwise.
func (o *DocsCloudUserStats) GetCriticalRemaining() bool {
	if o == nil || IsNil(o.CriticalRemaining) {
		var ret bool
		return ret
	}
	return *o.CriticalRemaining
}

// GetCriticalRemainingOk returns a tuple with the CriticalRemaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudUserStats) GetCriticalRemainingOk() (*bool, bool) {
	if o == nil || IsNil(o.CriticalRemaining) {
		return nil, false
	}
	return o.CriticalRemaining, true
}

// HasCriticalRemaining returns a boolean if a field has been set.
func (o *DocsCloudUserStats) IsCriticalRemainingSet() bool {
	if o != nil && !IsNil(o.CriticalRemaining) {
		return true
	}

	return false
}

// SetCriticalRemaining gets a reference to the given bool and assigns it to the CriticalRemaining field.
func (o *DocsCloudUserStats) SetCriticalRemaining(v bool) {
	o.CriticalRemaining = &v
}

func (o DocsCloudUserStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocsCloudUserStats) ToMap() (map[string]interface{}, error) {
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

type NullableDocsCloudUserStats struct {
	value *DocsCloudUserStats
	isSet bool
}

func (v NullableDocsCloudUserStats) Get() *DocsCloudUserStats {
	return v.value
}

func (v *NullableDocsCloudUserStats) Set(val *DocsCloudUserStats) {
	v.value = val
	v.isSet = true
}

func (v NullableDocsCloudUserStats) IsSet() bool {
	return v.isSet
}

func (v *NullableDocsCloudUserStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocsCloudUserStats(val *DocsCloudUserStats) *NullableDocsCloudUserStats {
	return &NullableDocsCloudUserStats{value: val, isSet: true}
}

func (v NullableDocsCloudUserStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocsCloudUserStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

