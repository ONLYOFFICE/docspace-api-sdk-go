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

// checks if the ReportDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportDto{}

// ReportDto One page of the portal wallet's money movements, with the paging figures needed to walk the rest.
type ReportDto struct {
	// The movements on this page - top-ups, charges, refunds and corrections alike, newest first. It is empty  for a page past the end of the report as well as for a period in which nothing happened.
	Collection []OperationDto `json:"collection,omitempty"`
	// How many movements were skipped before this page, echoed from the request so a client need not remember  what it asked for.
	Offset *int32 `json:"offset,omitempty"`
	// How many movements one page may hold, echoed from the request; it is 25 unless another value was asked  for. A full page is not proof that more exist - compare `currentPage` with `totalPage`.
	Limit *int32 `json:"limit,omitempty"`
	// How many movements match the filters in total, across every page.
	TotalQuantity *int64 `json:"totalQuantity,omitempty"`
	// How many pages those movements come to at the current `limit`.
	TotalPage *int32 `json:"totalPage,omitempty"`
	// Which of those pages this one is, as the billing service numbers them. Page through by advancing `offset`  rather than this value, which nothing accepts as an argument.
	CurrentPage *int32 `json:"currentPage,omitempty"`
}

// NewReportDto instantiates a new ReportDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportDto() *ReportDto {
	this := ReportDto{}
	return &this
}

// NewReportDtoWithDefaults instantiates a new ReportDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportDtoWithDefaults() *ReportDto {
	this := ReportDto{}
	return &this
}

// GetCollection returns the Collection field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportDto) GetCollection() []OperationDto {
	if o == nil {
		var ret []OperationDto
		return ret
	}
	return o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportDto) GetCollectionOk() ([]OperationDto, bool) {
	if o == nil || IsNil(o.Collection) {
		return nil, false
	}
	return o.Collection, true
}

// HasCollection returns a boolean if a field has been set.
func (o *ReportDto) IsCollectionSet() bool {
	if o != nil && !IsNil(o.Collection) {
		return true
	}

	return false
}

// SetCollection gets a reference to the given []OperationDto and assigns it to the Collection field.
func (o *ReportDto) SetCollection(v []OperationDto) {
	o.Collection = v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *ReportDto) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportDto) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *ReportDto) IsOffsetSet() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *ReportDto) SetOffset(v int32) {
	o.Offset = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ReportDto) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportDto) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ReportDto) IsLimitSet() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *ReportDto) SetLimit(v int32) {
	o.Limit = &v
}

// GetTotalQuantity returns the TotalQuantity field value if set, zero value otherwise.
func (o *ReportDto) GetTotalQuantity() int64 {
	if o == nil || IsNil(o.TotalQuantity) {
		var ret int64
		return ret
	}
	return *o.TotalQuantity
}

// GetTotalQuantityOk returns a tuple with the TotalQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportDto) GetTotalQuantityOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalQuantity) {
		return nil, false
	}
	return o.TotalQuantity, true
}

// HasTotalQuantity returns a boolean if a field has been set.
func (o *ReportDto) IsTotalQuantitySet() bool {
	if o != nil && !IsNil(o.TotalQuantity) {
		return true
	}

	return false
}

// SetTotalQuantity gets a reference to the given int64 and assigns it to the TotalQuantity field.
func (o *ReportDto) SetTotalQuantity(v int64) {
	o.TotalQuantity = &v
}

// GetTotalPage returns the TotalPage field value if set, zero value otherwise.
func (o *ReportDto) GetTotalPage() int32 {
	if o == nil || IsNil(o.TotalPage) {
		var ret int32
		return ret
	}
	return *o.TotalPage
}

// GetTotalPageOk returns a tuple with the TotalPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportDto) GetTotalPageOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalPage) {
		return nil, false
	}
	return o.TotalPage, true
}

// HasTotalPage returns a boolean if a field has been set.
func (o *ReportDto) IsTotalPageSet() bool {
	if o != nil && !IsNil(o.TotalPage) {
		return true
	}

	return false
}

// SetTotalPage gets a reference to the given int32 and assigns it to the TotalPage field.
func (o *ReportDto) SetTotalPage(v int32) {
	o.TotalPage = &v
}

// GetCurrentPage returns the CurrentPage field value if set, zero value otherwise.
func (o *ReportDto) GetCurrentPage() int32 {
	if o == nil || IsNil(o.CurrentPage) {
		var ret int32
		return ret
	}
	return *o.CurrentPage
}

// GetCurrentPageOk returns a tuple with the CurrentPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportDto) GetCurrentPageOk() (*int32, bool) {
	if o == nil || IsNil(o.CurrentPage) {
		return nil, false
	}
	return o.CurrentPage, true
}

// HasCurrentPage returns a boolean if a field has been set.
func (o *ReportDto) IsCurrentPageSet() bool {
	if o != nil && !IsNil(o.CurrentPage) {
		return true
	}

	return false
}

// SetCurrentPage gets a reference to the given int32 and assigns it to the CurrentPage field.
func (o *ReportDto) SetCurrentPage(v int32) {
	o.CurrentPage = &v
}

func (o ReportDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportDto) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.TotalQuantity) {
		toSerialize["totalQuantity"] = o.TotalQuantity
	}
	if !IsNil(o.TotalPage) {
		toSerialize["totalPage"] = o.TotalPage
	}
	if !IsNil(o.CurrentPage) {
		toSerialize["currentPage"] = o.CurrentPage
	}
	return toSerialize, nil
}

type NullableReportDto struct {
	value *ReportDto
	isSet bool
}

func (v NullableReportDto) Get() *ReportDto {
	return v.value
}

func (v *NullableReportDto) Set(val *ReportDto) {
	v.value = val
	v.isSet = true
}

func (v NullableReportDto) IsSet() bool {
	return v.isSet
}

func (v *NullableReportDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportDto(val *ReportDto) *NullableReportDto {
	return &NullableReportDto{value: val, isSet: true}
}

func (v NullableReportDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

