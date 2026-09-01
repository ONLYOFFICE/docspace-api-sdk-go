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

// checks if the DocsCloudServerInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocsCloudServerInfo{}

// DocsCloudServerInfo Represents the DocsCloud server information.
type DocsCloudServerInfo struct {
	// The server version.
	Version NullableString `json:"version,omitempty"`
	// The server package type (Open Source, Enterprise Edition or Developer Edition).
	PackageType NullableString `json:"packageType,omitempty"`
	// The server build date.
	Date *time.Time `json:"date,omitempty"`
}

// NewDocsCloudServerInfo instantiates a new DocsCloudServerInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocsCloudServerInfo() *DocsCloudServerInfo {
	this := DocsCloudServerInfo{}
	return &this
}

// NewDocsCloudServerInfoWithDefaults instantiates a new DocsCloudServerInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocsCloudServerInfoWithDefaults() *DocsCloudServerInfo {
	this := DocsCloudServerInfo{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocsCloudServerInfo) GetVersion() string {
	if o == nil || IsNil(o.Version.Get()) {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocsCloudServerInfo) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *DocsCloudServerInfo) IsVersionSet() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *DocsCloudServerInfo) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *DocsCloudServerInfo) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *DocsCloudServerInfo) UnsetVersion() {
	o.Version.Unset()
}

// GetPackageType returns the PackageType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocsCloudServerInfo) GetPackageType() string {
	if o == nil || IsNil(o.PackageType.Get()) {
		var ret string
		return ret
	}
	return *o.PackageType.Get()
}

// GetPackageTypeOk returns a tuple with the PackageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocsCloudServerInfo) GetPackageTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PackageType.Get(), o.PackageType.IsSet()
}

// HasPackageType returns a boolean if a field has been set.
func (o *DocsCloudServerInfo) IsPackageTypeSet() bool {
	if o != nil && o.PackageType.IsSet() {
		return true
	}

	return false
}

// SetPackageType gets a reference to the given NullableString and assigns it to the PackageType field.
func (o *DocsCloudServerInfo) SetPackageType(v string) {
	o.PackageType.Set(&v)
}
// SetPackageTypeNil sets the value for PackageType to be an explicit nil
func (o *DocsCloudServerInfo) SetPackageTypeNil() {
	o.PackageType.Set(nil)
}

// UnsetPackageType ensures that no value is present for PackageType, not even an explicit nil
func (o *DocsCloudServerInfo) UnsetPackageType() {
	o.PackageType.Unset()
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *DocsCloudServerInfo) GetDate() time.Time {
	if o == nil || IsNil(o.Date) {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocsCloudServerInfo) GetDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *DocsCloudServerInfo) IsDateSet() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *DocsCloudServerInfo) SetDate(v time.Time) {
	o.Date = &v
}

func (o DocsCloudServerInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocsCloudServerInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	if o.PackageType.IsSet() {
		toSerialize["packageType"] = o.PackageType.Get()
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	return toSerialize, nil
}

type NullableDocsCloudServerInfo struct {
	value *DocsCloudServerInfo
	isSet bool
}

func (v NullableDocsCloudServerInfo) Get() *DocsCloudServerInfo {
	return v.value
}

func (v *NullableDocsCloudServerInfo) Set(val *DocsCloudServerInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDocsCloudServerInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDocsCloudServerInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocsCloudServerInfo(val *DocsCloudServerInfo) *NullableDocsCloudServerInfo {
	return &NullableDocsCloudServerInfo{value: val, isSet: true}
}

func (v NullableDocsCloudServerInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocsCloudServerInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

