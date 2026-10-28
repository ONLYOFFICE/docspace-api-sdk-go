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

// checks if the AmazonS3RegionDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AmazonS3RegionDto{}

// AmazonS3RegionDto An Amazon S3 region.
type AmazonS3RegionDto struct {
	// The region code to send as the region value when configuring an Amazon S3 storage or backup target. It is  the one field of this object that is an argument elsewhere; a code the server does not list here cannot be  reached, so pick one from this list rather than typing it.
	SystemName NullableString `json:"systemName,omitempty"`
	// The region name as Amazon writes it, in English regardless of the portal language, for showing in a  picker next to `systemName`.
	DisplayName NullableString `json:"displayName,omitempty"`
	// The Amazon partition the region sits in - the ordinary commercial cloud, the Chinese one, or a government  one. Regions of different partitions are not reachable with the same credentials.
	PartitionName NullableString `json:"partitionName,omitempty"`
	// The domain the partition's service host names end in, which differs from partition to partition.
	PartitionDnsSuffix NullableString `json:"partitionDnsSuffix,omitempty"`
	// The pattern every region code of this partition matches, for validating a code before sending it.
	PartitionRegionRegex NullableString `json:"partitionRegionRegex,omitempty"`
	// How a service host name of the partition is assembled, with `{service}`, `{region}` and `{dnsSuffix}` to  be filled in. It is reference material - the portal builds its own endpoints from `systemName`.
	HostnameTemplate NullableString `json:"hostnameTemplate,omitempty"`
}

// NewAmazonS3RegionDto instantiates a new AmazonS3RegionDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmazonS3RegionDto() *AmazonS3RegionDto {
	this := AmazonS3RegionDto{}
	return &this
}

// NewAmazonS3RegionDtoWithDefaults instantiates a new AmazonS3RegionDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmazonS3RegionDtoWithDefaults() *AmazonS3RegionDto {
	this := AmazonS3RegionDto{}
	return &this
}

// GetSystemName returns the SystemName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AmazonS3RegionDto) GetSystemName() string {
	if o == nil || IsNil(o.SystemName.Get()) {
		var ret string
		return ret
	}
	return *o.SystemName.Get()
}

// GetSystemNameOk returns a tuple with the SystemName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AmazonS3RegionDto) GetSystemNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SystemName.Get(), o.SystemName.IsSet()
}

// HasSystemName returns a boolean if a field has been set.
func (o *AmazonS3RegionDto) IsSystemNameSet() bool {
	if o != nil && o.SystemName.IsSet() {
		return true
	}

	return false
}

// SetSystemName gets a reference to the given NullableString and assigns it to the SystemName field.
func (o *AmazonS3RegionDto) SetSystemName(v string) {
	o.SystemName.Set(&v)
}
// SetSystemNameNil sets the value for SystemName to be an explicit nil
func (o *AmazonS3RegionDto) SetSystemNameNil() {
	o.SystemName.Set(nil)
}

// UnsetSystemName ensures that no value is present for SystemName, not even an explicit nil
func (o *AmazonS3RegionDto) UnsetSystemName() {
	o.SystemName.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AmazonS3RegionDto) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName.Get()) {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AmazonS3RegionDto) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AmazonS3RegionDto) IsDisplayNameSet() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *AmazonS3RegionDto) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *AmazonS3RegionDto) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *AmazonS3RegionDto) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetPartitionName returns the PartitionName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AmazonS3RegionDto) GetPartitionName() string {
	if o == nil || IsNil(o.PartitionName.Get()) {
		var ret string
		return ret
	}
	return *o.PartitionName.Get()
}

// GetPartitionNameOk returns a tuple with the PartitionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AmazonS3RegionDto) GetPartitionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PartitionName.Get(), o.PartitionName.IsSet()
}

// HasPartitionName returns a boolean if a field has been set.
func (o *AmazonS3RegionDto) IsPartitionNameSet() bool {
	if o != nil && o.PartitionName.IsSet() {
		return true
	}

	return false
}

// SetPartitionName gets a reference to the given NullableString and assigns it to the PartitionName field.
func (o *AmazonS3RegionDto) SetPartitionName(v string) {
	o.PartitionName.Set(&v)
}
// SetPartitionNameNil sets the value for PartitionName to be an explicit nil
func (o *AmazonS3RegionDto) SetPartitionNameNil() {
	o.PartitionName.Set(nil)
}

// UnsetPartitionName ensures that no value is present for PartitionName, not even an explicit nil
func (o *AmazonS3RegionDto) UnsetPartitionName() {
	o.PartitionName.Unset()
}

// GetPartitionDnsSuffix returns the PartitionDnsSuffix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AmazonS3RegionDto) GetPartitionDnsSuffix() string {
	if o == nil || IsNil(o.PartitionDnsSuffix.Get()) {
		var ret string
		return ret
	}
	return *o.PartitionDnsSuffix.Get()
}

// GetPartitionDnsSuffixOk returns a tuple with the PartitionDnsSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AmazonS3RegionDto) GetPartitionDnsSuffixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PartitionDnsSuffix.Get(), o.PartitionDnsSuffix.IsSet()
}

// HasPartitionDnsSuffix returns a boolean if a field has been set.
func (o *AmazonS3RegionDto) IsPartitionDnsSuffixSet() bool {
	if o != nil && o.PartitionDnsSuffix.IsSet() {
		return true
	}

	return false
}

// SetPartitionDnsSuffix gets a reference to the given NullableString and assigns it to the PartitionDnsSuffix field.
func (o *AmazonS3RegionDto) SetPartitionDnsSuffix(v string) {
	o.PartitionDnsSuffix.Set(&v)
}
// SetPartitionDnsSuffixNil sets the value for PartitionDnsSuffix to be an explicit nil
func (o *AmazonS3RegionDto) SetPartitionDnsSuffixNil() {
	o.PartitionDnsSuffix.Set(nil)
}

// UnsetPartitionDnsSuffix ensures that no value is present for PartitionDnsSuffix, not even an explicit nil
func (o *AmazonS3RegionDto) UnsetPartitionDnsSuffix() {
	o.PartitionDnsSuffix.Unset()
}

// GetPartitionRegionRegex returns the PartitionRegionRegex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AmazonS3RegionDto) GetPartitionRegionRegex() string {
	if o == nil || IsNil(o.PartitionRegionRegex.Get()) {
		var ret string
		return ret
	}
	return *o.PartitionRegionRegex.Get()
}

// GetPartitionRegionRegexOk returns a tuple with the PartitionRegionRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AmazonS3RegionDto) GetPartitionRegionRegexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PartitionRegionRegex.Get(), o.PartitionRegionRegex.IsSet()
}

// HasPartitionRegionRegex returns a boolean if a field has been set.
func (o *AmazonS3RegionDto) IsPartitionRegionRegexSet() bool {
	if o != nil && o.PartitionRegionRegex.IsSet() {
		return true
	}

	return false
}

// SetPartitionRegionRegex gets a reference to the given NullableString and assigns it to the PartitionRegionRegex field.
func (o *AmazonS3RegionDto) SetPartitionRegionRegex(v string) {
	o.PartitionRegionRegex.Set(&v)
}
// SetPartitionRegionRegexNil sets the value for PartitionRegionRegex to be an explicit nil
func (o *AmazonS3RegionDto) SetPartitionRegionRegexNil() {
	o.PartitionRegionRegex.Set(nil)
}

// UnsetPartitionRegionRegex ensures that no value is present for PartitionRegionRegex, not even an explicit nil
func (o *AmazonS3RegionDto) UnsetPartitionRegionRegex() {
	o.PartitionRegionRegex.Unset()
}

// GetHostnameTemplate returns the HostnameTemplate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AmazonS3RegionDto) GetHostnameTemplate() string {
	if o == nil || IsNil(o.HostnameTemplate.Get()) {
		var ret string
		return ret
	}
	return *o.HostnameTemplate.Get()
}

// GetHostnameTemplateOk returns a tuple with the HostnameTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AmazonS3RegionDto) GetHostnameTemplateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HostnameTemplate.Get(), o.HostnameTemplate.IsSet()
}

// HasHostnameTemplate returns a boolean if a field has been set.
func (o *AmazonS3RegionDto) IsHostnameTemplateSet() bool {
	if o != nil && o.HostnameTemplate.IsSet() {
		return true
	}

	return false
}

// SetHostnameTemplate gets a reference to the given NullableString and assigns it to the HostnameTemplate field.
func (o *AmazonS3RegionDto) SetHostnameTemplate(v string) {
	o.HostnameTemplate.Set(&v)
}
// SetHostnameTemplateNil sets the value for HostnameTemplate to be an explicit nil
func (o *AmazonS3RegionDto) SetHostnameTemplateNil() {
	o.HostnameTemplate.Set(nil)
}

// UnsetHostnameTemplate ensures that no value is present for HostnameTemplate, not even an explicit nil
func (o *AmazonS3RegionDto) UnsetHostnameTemplate() {
	o.HostnameTemplate.Unset()
}

func (o AmazonS3RegionDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AmazonS3RegionDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.SystemName.IsSet() {
		toSerialize["systemName"] = o.SystemName.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.PartitionName.IsSet() {
		toSerialize["partitionName"] = o.PartitionName.Get()
	}
	if o.PartitionDnsSuffix.IsSet() {
		toSerialize["partitionDnsSuffix"] = o.PartitionDnsSuffix.Get()
	}
	if o.PartitionRegionRegex.IsSet() {
		toSerialize["partitionRegionRegex"] = o.PartitionRegionRegex.Get()
	}
	if o.HostnameTemplate.IsSet() {
		toSerialize["hostnameTemplate"] = o.HostnameTemplate.Get()
	}
	return toSerialize, nil
}

type NullableAmazonS3RegionDto struct {
	value *AmazonS3RegionDto
	isSet bool
}

func (v NullableAmazonS3RegionDto) Get() *AmazonS3RegionDto {
	return v.value
}

func (v *NullableAmazonS3RegionDto) Set(val *AmazonS3RegionDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAmazonS3RegionDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAmazonS3RegionDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmazonS3RegionDto(val *AmazonS3RegionDto) *NullableAmazonS3RegionDto {
	return &NullableAmazonS3RegionDto{value: val, isSet: true}
}

func (v NullableAmazonS3RegionDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmazonS3RegionDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

