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

// checks if the ServiceLimitReportDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceLimitReportDto{}

// ServiceLimitReportDto A page of the spending limits in force for the portal users.
type ServiceLimitReportDto struct {
	// The limits of this page, one per user; empty when no limit applies to any user or the page is past the end.
	Collection []ServiceLimitDto `json:"collection,omitempty"`
	// The number of limits skipped before this page.
	Offset *int32 `json:"offset,omitempty"`
	// The page size the page was read with.
	Limit *int32 `json:"limit,omitempty"`
	// The number of pages the whole list takes at this page size.
	TotalPage *int32 `json:"totalPage,omitempty"`
	// The number of this page, starting from 1.
	CurrentPage *int32 `json:"currentPage,omitempty"`
	// The number of users with a limit in force across all pages.
	TotalQuantity *int32 `json:"totalQuantity,omitempty"`
}

// NewServiceLimitReportDto instantiates a new ServiceLimitReportDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceLimitReportDto() *ServiceLimitReportDto {
	this := ServiceLimitReportDto{}
	return &this
}

// NewServiceLimitReportDtoWithDefaults instantiates a new ServiceLimitReportDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceLimitReportDtoWithDefaults() *ServiceLimitReportDto {
	this := ServiceLimitReportDto{}
	return &this
}

// GetCollection returns the Collection field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ServiceLimitReportDto) GetCollection() []ServiceLimitDto {
	if o == nil {
		var ret []ServiceLimitDto
		return ret
	}
	return o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ServiceLimitReportDto) GetCollectionOk() ([]ServiceLimitDto, bool) {
	if o == nil || IsNil(o.Collection) {
		return nil, false
	}
	return o.Collection, true
}

// HasCollection returns a boolean if a field has been set.
func (o *ServiceLimitReportDto) IsCollectionSet() bool {
	if o != nil && !IsNil(o.Collection) {
		return true
	}

	return false
}

// SetCollection gets a reference to the given []ServiceLimitDto and assigns it to the Collection field.
func (o *ServiceLimitReportDto) SetCollection(v []ServiceLimitDto) {
	o.Collection = v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *ServiceLimitReportDto) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLimitReportDto) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *ServiceLimitReportDto) IsOffsetSet() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *ServiceLimitReportDto) SetOffset(v int32) {
	o.Offset = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ServiceLimitReportDto) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLimitReportDto) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ServiceLimitReportDto) IsLimitSet() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *ServiceLimitReportDto) SetLimit(v int32) {
	o.Limit = &v
}

// GetTotalPage returns the TotalPage field value if set, zero value otherwise.
func (o *ServiceLimitReportDto) GetTotalPage() int32 {
	if o == nil || IsNil(o.TotalPage) {
		var ret int32
		return ret
	}
	return *o.TotalPage
}

// GetTotalPageOk returns a tuple with the TotalPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLimitReportDto) GetTotalPageOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalPage) {
		return nil, false
	}
	return o.TotalPage, true
}

// HasTotalPage returns a boolean if a field has been set.
func (o *ServiceLimitReportDto) IsTotalPageSet() bool {
	if o != nil && !IsNil(o.TotalPage) {
		return true
	}

	return false
}

// SetTotalPage gets a reference to the given int32 and assigns it to the TotalPage field.
func (o *ServiceLimitReportDto) SetTotalPage(v int32) {
	o.TotalPage = &v
}

// GetCurrentPage returns the CurrentPage field value if set, zero value otherwise.
func (o *ServiceLimitReportDto) GetCurrentPage() int32 {
	if o == nil || IsNil(o.CurrentPage) {
		var ret int32
		return ret
	}
	return *o.CurrentPage
}

// GetCurrentPageOk returns a tuple with the CurrentPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLimitReportDto) GetCurrentPageOk() (*int32, bool) {
	if o == nil || IsNil(o.CurrentPage) {
		return nil, false
	}
	return o.CurrentPage, true
}

// HasCurrentPage returns a boolean if a field has been set.
func (o *ServiceLimitReportDto) IsCurrentPageSet() bool {
	if o != nil && !IsNil(o.CurrentPage) {
		return true
	}

	return false
}

// SetCurrentPage gets a reference to the given int32 and assigns it to the CurrentPage field.
func (o *ServiceLimitReportDto) SetCurrentPage(v int32) {
	o.CurrentPage = &v
}

// GetTotalQuantity returns the TotalQuantity field value if set, zero value otherwise.
func (o *ServiceLimitReportDto) GetTotalQuantity() int32 {
	if o == nil || IsNil(o.TotalQuantity) {
		var ret int32
		return ret
	}
	return *o.TotalQuantity
}

// GetTotalQuantityOk returns a tuple with the TotalQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceLimitReportDto) GetTotalQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalQuantity) {
		return nil, false
	}
	return o.TotalQuantity, true
}

// HasTotalQuantity returns a boolean if a field has been set.
func (o *ServiceLimitReportDto) IsTotalQuantitySet() bool {
	if o != nil && !IsNil(o.TotalQuantity) {
		return true
	}

	return false
}

// SetTotalQuantity gets a reference to the given int32 and assigns it to the TotalQuantity field.
func (o *ServiceLimitReportDto) SetTotalQuantity(v int32) {
	o.TotalQuantity = &v
}

func (o ServiceLimitReportDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceLimitReportDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Collection != nil {
		toSerialize["collection"] = o.Collection
	}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.TotalPage) {
		toSerialize["totalPage"] = o.TotalPage
	}
	if !IsNil(o.CurrentPage) {
		toSerialize["currentPage"] = o.CurrentPage
	}
	if !IsNil(o.TotalQuantity) {
		toSerialize["totalQuantity"] = o.TotalQuantity
	}
	return toSerialize, nil
}

type NullableServiceLimitReportDto struct {
	value *ServiceLimitReportDto
	isSet bool
}

func (v NullableServiceLimitReportDto) Get() *ServiceLimitReportDto {
	return v.value
}

func (v *NullableServiceLimitReportDto) Set(val *ServiceLimitReportDto) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceLimitReportDto) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceLimitReportDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceLimitReportDto(val *ServiceLimitReportDto) *NullableServiceLimitReportDto {
	return &NullableServiceLimitReportDto{value: val, isSet: true}
}

func (v NullableServiceLimitReportDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceLimitReportDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

