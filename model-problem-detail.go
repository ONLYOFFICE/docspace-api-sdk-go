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

// checks if the ProblemDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProblemDetail{}

// ProblemDetail struct for ProblemDetail
type ProblemDetail struct {
	Type *string `json:"type,omitempty"`
	Title *string `json:"title,omitempty"`
	Status *int32 `json:"status,omitempty"`
	Detail *string `json:"detail,omitempty"`
	Instance *string `json:"instance,omitempty"`
	Properties map[string]map[string]interface{} `json:"properties,omitempty"`
}

// NewProblemDetail instantiates a new ProblemDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProblemDetail() *ProblemDetail {
	this := ProblemDetail{}
	return &this
}

// NewProblemDetailWithDefaults instantiates a new ProblemDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProblemDetailWithDefaults() *ProblemDetail {
	this := ProblemDetail{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ProblemDetail) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetail) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ProblemDetail) IsTypeSet() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ProblemDetail) SetType(v string) {
	o.Type = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ProblemDetail) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetail) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ProblemDetail) IsTitleSet() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ProblemDetail) SetTitle(v string) {
	o.Title = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ProblemDetail) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetail) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ProblemDetail) IsStatusSet() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *ProblemDetail) SetStatus(v int32) {
	o.Status = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *ProblemDetail) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetail) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *ProblemDetail) IsDetailSet() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *ProblemDetail) SetDetail(v string) {
	o.Detail = &v
}

// GetInstance returns the Instance field value if set, zero value otherwise.
func (o *ProblemDetail) GetInstance() string {
	if o == nil || IsNil(o.Instance) {
		var ret string
		return ret
	}
	return *o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetail) GetInstanceOk() (*string, bool) {
	if o == nil || IsNil(o.Instance) {
		return nil, false
	}
	return o.Instance, true
}

// HasInstance returns a boolean if a field has been set.
func (o *ProblemDetail) IsInstanceSet() bool {
	if o != nil && !IsNil(o.Instance) {
		return true
	}

	return false
}

// SetInstance gets a reference to the given string and assigns it to the Instance field.
func (o *ProblemDetail) SetInstance(v string) {
	o.Instance = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *ProblemDetail) GetProperties() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Properties) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetail) GetPropertiesOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Properties) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *ProblemDetail) IsPropertiesSet() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]map[string]interface{} and assigns it to the Properties field.
func (o *ProblemDetail) SetProperties(v map[string]map[string]interface{}) {
	o.Properties = v
}

func (o ProblemDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProblemDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	if !IsNil(o.Instance) {
		toSerialize["instance"] = o.Instance
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	return toSerialize, nil
}

type NullableProblemDetail struct {
	value *ProblemDetail
	isSet bool
}

func (v NullableProblemDetail) Get() *ProblemDetail {
	return v.value
}

func (v *NullableProblemDetail) Set(val *ProblemDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableProblemDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableProblemDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProblemDetail(val *ProblemDetail) *NullableProblemDetail {
	return &NullableProblemDetail{value: val, isSet: true}
}

func (v NullableProblemDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProblemDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

