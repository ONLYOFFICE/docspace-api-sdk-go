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

// checks if the DocsCloudStatsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocsCloudStatsDto{}

// DocsCloudStatsDto Represents the usage statistics of a Docs Connect tenant for the current period.
type DocsCloudStatsDto struct {
	// The length of the statistics period in days.
	PeriodDay *int32 `json:"periodDay,omitempty"`
	// The statistics for editor users.
	Editor *DocsCloudUserStatsDto `json:"editor,omitempty"`
	// The statistics for viewer users.
	Viewer *DocsCloudUserStatsDto `json:"viewer,omitempty"`
}

// NewDocsCloudStatsDto instantiates a new DocsCloudStatsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocsCloudStatsDto() *DocsCloudStatsDto {
	this := DocsCloudStatsDto{}
	return &this
}

// NewDocsCloudStatsDtoWithDefaults instantiates a new DocsCloudStatsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocsCloudStatsDtoWithDefaults() *DocsCloudStatsDto {
	this := DocsCloudStatsDto{}
	return &this
}

// GetPeriodDay returns the PeriodDay field value if set, zero value otherwise.
func (o *DocsCloudStatsDto) GetPeriodDay() int32 {
	if o == nil || IsNil(o.PeriodDay) {
		var ret int32
		return ret
	}
	return *o.PeriodDay
}

// GetPeriodDayOk returns a tuple with the PeriodDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudStatsDto) GetPeriodDayOk() (*int32, bool) {
	if o == nil || IsNil(o.PeriodDay) {
		return nil, false
	}
	return o.PeriodDay, true
}

// HasPeriodDay returns a boolean if a field has been set.
func (o *DocsCloudStatsDto) IsPeriodDaySet() bool {
	if o != nil && !IsNil(o.PeriodDay) {
		return true
	}

	return false
}

// SetPeriodDay gets a reference to the given int32 and assigns it to the PeriodDay field.
func (o *DocsCloudStatsDto) SetPeriodDay(v int32) {
	o.PeriodDay = &v
}

// GetEditor returns the Editor field value if set, zero value otherwise.
func (o *DocsCloudStatsDto) GetEditor() DocsCloudUserStatsDto {
	if o == nil || IsNil(o.Editor) {
		var ret DocsCloudUserStatsDto
		return ret
	}
	return *o.Editor
}

// GetEditorOk returns a tuple with the Editor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudStatsDto) GetEditorOk() (*DocsCloudUserStatsDto, bool) {
	if o == nil || IsNil(o.Editor) {
		return nil, false
	}
	return o.Editor, true
}

// HasEditor returns a boolean if a field has been set.
func (o *DocsCloudStatsDto) IsEditorSet() bool {
	if o != nil && !IsNil(o.Editor) {
		return true
	}

	return false
}

// SetEditor gets a reference to the given DocsCloudUserStatsDto and assigns it to the Editor field.
func (o *DocsCloudStatsDto) SetEditor(v DocsCloudUserStatsDto) {
	o.Editor = &v
}

// GetViewer returns the Viewer field value if set, zero value otherwise.
func (o *DocsCloudStatsDto) GetViewer() DocsCloudUserStatsDto {
	if o == nil || IsNil(o.Viewer) {
		var ret DocsCloudUserStatsDto
		return ret
	}
	return *o.Viewer
}

// GetViewerOk returns a tuple with the Viewer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudStatsDto) GetViewerOk() (*DocsCloudUserStatsDto, bool) {
	if o == nil || IsNil(o.Viewer) {
		return nil, false
	}
	return o.Viewer, true
}

// HasViewer returns a boolean if a field has been set.
func (o *DocsCloudStatsDto) IsViewerSet() bool {
	if o != nil && !IsNil(o.Viewer) {
		return true
	}

	return false
}

// SetViewer gets a reference to the given DocsCloudUserStatsDto and assigns it to the Viewer field.
func (o *DocsCloudStatsDto) SetViewer(v DocsCloudUserStatsDto) {
	o.Viewer = &v
}

func (o DocsCloudStatsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocsCloudStatsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PeriodDay) {
		toSerialize["periodDay"] = o.PeriodDay
	}
	if !IsNil(o.Editor) {
		toSerialize["editor"] = o.Editor
	}
	if !IsNil(o.Viewer) {
		toSerialize["viewer"] = o.Viewer
	}
	return toSerialize, nil
}

type NullableDocsCloudStatsDto struct {
	value *DocsCloudStatsDto
	isSet bool
}

func (v NullableDocsCloudStatsDto) Get() *DocsCloudStatsDto {
	return v.value
}

func (v *NullableDocsCloudStatsDto) Set(val *DocsCloudStatsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDocsCloudStatsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDocsCloudStatsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocsCloudStatsDto(val *DocsCloudStatsDto) *NullableDocsCloudStatsDto {
	return &NullableDocsCloudStatsDto{value: val, isSet: true}
}

func (v NullableDocsCloudStatsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocsCloudStatsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

