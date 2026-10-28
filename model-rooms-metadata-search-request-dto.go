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

// checks if the RoomsMetadataSearchRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoomsMetadataSearchRequestDto{}

// RoomsMetadataSearchRequestDto The typed form of the metadata search of the rooms: the same filter the rooms listing takes in the metadataTemplateId  and metadataFilters query parameters, with the conditions as objects instead of a JSON string.
type RoomsMetadataSearchRequestDto struct {
	// The ID of the metadata template the rooms must be assigned to. On its own it narrows the listing to the rooms  carrying the template; together with the conditions it also pins the template the filtered fields belong to.
	MetadataTemplateId NullableInt32 `json:"metadataTemplateId,omitempty"`
	// The metadata filter conditions, combined with AND. A custom field is addressed by its name instead of the field ID.
	MetadataFilters []MetadataFilterConditionRequest `json:"metadataFilters,omitempty"`
	// Whether the text search, and the metadataMatches it reports, also look at the values of the hidden metadata  templates (created with visible false for the metadata a service keeps for itself, which the web client never  shows). False, the default, leaves them out, so an entry is not found by such a value; true takes them in.  metadataTemplateId and metadataFilters are not concerned: a hidden template or field named there is searched  either way.
	IncludeHiddenMetadata NullableBool `json:"includeHiddenMetadata,omitempty"`
	// The text to search for in the room titles and in the string metadata values of the rooms, the custom fields and  the template fields alike; a room found by its metadata says where in metadataMatches.
	FilterValue NullableString `json:"filterValue,omitempty"`
	// The section to search in: the active rooms (the default), the archive or the templates.
	SearchArea *SearchArea `json:"searchArea,omitempty"`
	// The room types to search among.
	Type []RoomType `json:"type,omitempty"`
	// The number of rooms to return, from 1 to 100.
	Count *int32 `json:"count,omitempty"`
	// The zero-based index of the first room to return.
	StartIndex *int32 `json:"startIndex,omitempty"`
	// The field to sort by, a name of the SortedByType values.
	SortBy NullableString `json:"sortBy,omitempty"`
	// The sort order.
	SortOrder *SortOrder `json:"sortOrder,omitempty"`
}

// NewRoomsMetadataSearchRequestDto instantiates a new RoomsMetadataSearchRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoomsMetadataSearchRequestDto() *RoomsMetadataSearchRequestDto {
	this := RoomsMetadataSearchRequestDto{}
	return &this
}

// NewRoomsMetadataSearchRequestDtoWithDefaults instantiates a new RoomsMetadataSearchRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoomsMetadataSearchRequestDtoWithDefaults() *RoomsMetadataSearchRequestDto {
	this := RoomsMetadataSearchRequestDto{}
	return &this
}

// GetMetadataTemplateId returns the MetadataTemplateId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoomsMetadataSearchRequestDto) GetMetadataTemplateId() int32 {
	if o == nil || IsNil(o.MetadataTemplateId.Get()) {
		var ret int32
		return ret
	}
	return *o.MetadataTemplateId.Get()
}

// GetMetadataTemplateIdOk returns a tuple with the MetadataTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoomsMetadataSearchRequestDto) GetMetadataTemplateIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetadataTemplateId.Get(), o.MetadataTemplateId.IsSet()
}

// HasMetadataTemplateId returns a boolean if a field has been set.
func (o *RoomsMetadataSearchRequestDto) IsMetadataTemplateIdSet() bool {
	if o != nil && o.MetadataTemplateId.IsSet() {
		return true
	}

	return false
}

// SetMetadataTemplateId gets a reference to the given NullableInt32 and assigns it to the MetadataTemplateId field.
func (o *RoomsMetadataSearchRequestDto) SetMetadataTemplateId(v int32) {
	o.MetadataTemplateId.Set(&v)
}
// SetMetadataTemplateIdNil sets the value for MetadataTemplateId to be an explicit nil
func (o *RoomsMetadataSearchRequestDto) SetMetadataTemplateIdNil() {
	o.MetadataTemplateId.Set(nil)
}

// UnsetMetadataTemplateId ensures that no value is present for MetadataTemplateId, not even an explicit nil
func (o *RoomsMetadataSearchRequestDto) UnsetMetadataTemplateId() {
	o.MetadataTemplateId.Unset()
}

// GetMetadataFilters returns the MetadataFilters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoomsMetadataSearchRequestDto) GetMetadataFilters() []MetadataFilterConditionRequest {
	if o == nil {
		var ret []MetadataFilterConditionRequest
		return ret
	}
	return o.MetadataFilters
}

// GetMetadataFiltersOk returns a tuple with the MetadataFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoomsMetadataSearchRequestDto) GetMetadataFiltersOk() ([]MetadataFilterConditionRequest, bool) {
	if o == nil || IsNil(o.MetadataFilters) {
		return nil, false
	}
	return o.MetadataFilters, true
}

// HasMetadataFilters returns a boolean if a field has been set.
func (o *RoomsMetadataSearchRequestDto) IsMetadataFiltersSet() bool {
	if o != nil && !IsNil(o.MetadataFilters) {
		return true
	}

	return false
}

// SetMetadataFilters gets a reference to the given []MetadataFilterConditionRequest and assigns it to the MetadataFilters field.
func (o *RoomsMetadataSearchRequestDto) SetMetadataFilters(v []MetadataFilterConditionRequest) {
	o.MetadataFilters = v
}

// GetIncludeHiddenMetadata returns the IncludeHiddenMetadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoomsMetadataSearchRequestDto) GetIncludeHiddenMetadata() bool {
	if o == nil || IsNil(o.IncludeHiddenMetadata.Get()) {
		var ret bool
		return ret
	}
	return *o.IncludeHiddenMetadata.Get()
}

// GetIncludeHiddenMetadataOk returns a tuple with the IncludeHiddenMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoomsMetadataSearchRequestDto) GetIncludeHiddenMetadataOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IncludeHiddenMetadata.Get(), o.IncludeHiddenMetadata.IsSet()
}

// HasIncludeHiddenMetadata returns a boolean if a field has been set.
func (o *RoomsMetadataSearchRequestDto) IsIncludeHiddenMetadataSet() bool {
	if o != nil && o.IncludeHiddenMetadata.IsSet() {
		return true
	}

	return false
}

// SetIncludeHiddenMetadata gets a reference to the given NullableBool and assigns it to the IncludeHiddenMetadata field.
func (o *RoomsMetadataSearchRequestDto) SetIncludeHiddenMetadata(v bool) {
	o.IncludeHiddenMetadata.Set(&v)
}
// SetIncludeHiddenMetadataNil sets the value for IncludeHiddenMetadata to be an explicit nil
func (o *RoomsMetadataSearchRequestDto) SetIncludeHiddenMetadataNil() {
	o.IncludeHiddenMetadata.Set(nil)
}

// UnsetIncludeHiddenMetadata ensures that no value is present for IncludeHiddenMetadata, not even an explicit nil
func (o *RoomsMetadataSearchRequestDto) UnsetIncludeHiddenMetadata() {
	o.IncludeHiddenMetadata.Unset()
}

// GetFilterValue returns the FilterValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoomsMetadataSearchRequestDto) GetFilterValue() string {
	if o == nil || IsNil(o.FilterValue.Get()) {
		var ret string
		return ret
	}
	return *o.FilterValue.Get()
}

// GetFilterValueOk returns a tuple with the FilterValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoomsMetadataSearchRequestDto) GetFilterValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FilterValue.Get(), o.FilterValue.IsSet()
}

// HasFilterValue returns a boolean if a field has been set.
func (o *RoomsMetadataSearchRequestDto) IsFilterValueSet() bool {
	if o != nil && o.FilterValue.IsSet() {
		return true
	}

	return false
}

// SetFilterValue gets a reference to the given NullableString and assigns it to the FilterValue field.
func (o *RoomsMetadataSearchRequestDto) SetFilterValue(v string) {
	o.FilterValue.Set(&v)
}
// SetFilterValueNil sets the value for FilterValue to be an explicit nil
func (o *RoomsMetadataSearchRequestDto) SetFilterValueNil() {
	o.FilterValue.Set(nil)
}

// UnsetFilterValue ensures that no value is present for FilterValue, not even an explicit nil
func (o *RoomsMetadataSearchRequestDto) UnsetFilterValue() {
	o.FilterValue.Unset()
}

// GetSearchArea returns the SearchArea field value if set, zero value otherwise.
func (o *RoomsMetadataSearchRequestDto) GetSearchArea() SearchArea {
	if o == nil || IsNil(o.SearchArea) {
		var ret SearchArea
		return ret
	}
	return *o.SearchArea
}

// GetSearchAreaOk returns a tuple with the SearchArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomsMetadataSearchRequestDto) GetSearchAreaOk() (*SearchArea, bool) {
	if o == nil || IsNil(o.SearchArea) {
		return nil, false
	}
	return o.SearchArea, true
}

// HasSearchArea returns a boolean if a field has been set.
func (o *RoomsMetadataSearchRequestDto) IsSearchAreaSet() bool {
	if o != nil && !IsNil(o.SearchArea) {
		return true
	}

	return false
}

// SetSearchArea gets a reference to the given SearchArea and assigns it to the SearchArea field.
func (o *RoomsMetadataSearchRequestDto) SetSearchArea(v SearchArea) {
	o.SearchArea = &v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoomsMetadataSearchRequestDto) GetType() []RoomType {
	if o == nil {
		var ret []RoomType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoomsMetadataSearchRequestDto) GetTypeOk() ([]RoomType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RoomsMetadataSearchRequestDto) IsTypeSet() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given []RoomType and assigns it to the Type field.
func (o *RoomsMetadataSearchRequestDto) SetType(v []RoomType) {
	o.Type = v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *RoomsMetadataSearchRequestDto) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomsMetadataSearchRequestDto) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *RoomsMetadataSearchRequestDto) IsCountSet() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *RoomsMetadataSearchRequestDto) SetCount(v int32) {
	o.Count = &v
}

// GetStartIndex returns the StartIndex field value if set, zero value otherwise.
func (o *RoomsMetadataSearchRequestDto) GetStartIndex() int32 {
	if o == nil || IsNil(o.StartIndex) {
		var ret int32
		return ret
	}
	return *o.StartIndex
}

// GetStartIndexOk returns a tuple with the StartIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomsMetadataSearchRequestDto) GetStartIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.StartIndex) {
		return nil, false
	}
	return o.StartIndex, true
}

// HasStartIndex returns a boolean if a field has been set.
func (o *RoomsMetadataSearchRequestDto) IsStartIndexSet() bool {
	if o != nil && !IsNil(o.StartIndex) {
		return true
	}

	return false
}

// SetStartIndex gets a reference to the given int32 and assigns it to the StartIndex field.
func (o *RoomsMetadataSearchRequestDto) SetStartIndex(v int32) {
	o.StartIndex = &v
}

// GetSortBy returns the SortBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoomsMetadataSearchRequestDto) GetSortBy() string {
	if o == nil || IsNil(o.SortBy.Get()) {
		var ret string
		return ret
	}
	return *o.SortBy.Get()
}

// GetSortByOk returns a tuple with the SortBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoomsMetadataSearchRequestDto) GetSortByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SortBy.Get(), o.SortBy.IsSet()
}

// HasSortBy returns a boolean if a field has been set.
func (o *RoomsMetadataSearchRequestDto) IsSortBySet() bool {
	if o != nil && o.SortBy.IsSet() {
		return true
	}

	return false
}

// SetSortBy gets a reference to the given NullableString and assigns it to the SortBy field.
func (o *RoomsMetadataSearchRequestDto) SetSortBy(v string) {
	o.SortBy.Set(&v)
}
// SetSortByNil sets the value for SortBy to be an explicit nil
func (o *RoomsMetadataSearchRequestDto) SetSortByNil() {
	o.SortBy.Set(nil)
}

// UnsetSortBy ensures that no value is present for SortBy, not even an explicit nil
func (o *RoomsMetadataSearchRequestDto) UnsetSortBy() {
	o.SortBy.Unset()
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *RoomsMetadataSearchRequestDto) GetSortOrder() SortOrder {
	if o == nil || IsNil(o.SortOrder) {
		var ret SortOrder
		return ret
	}
	return *o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomsMetadataSearchRequestDto) GetSortOrderOk() (*SortOrder, bool) {
	if o == nil || IsNil(o.SortOrder) {
		return nil, false
	}
	return o.SortOrder, true
}

// HasSortOrder returns a boolean if a field has been set.
func (o *RoomsMetadataSearchRequestDto) IsSortOrderSet() bool {
	if o != nil && !IsNil(o.SortOrder) {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given SortOrder and assigns it to the SortOrder field.
func (o *RoomsMetadataSearchRequestDto) SetSortOrder(v SortOrder) {
	o.SortOrder = &v
}

func (o RoomsMetadataSearchRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoomsMetadataSearchRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MetadataTemplateId.IsSet() {
		toSerialize["metadataTemplateId"] = o.MetadataTemplateId.Get()
	}
	if o.MetadataFilters != nil {
		toSerialize["metadataFilters"] = o.MetadataFilters
	}
	if o.IncludeHiddenMetadata.IsSet() {
		toSerialize["includeHiddenMetadata"] = o.IncludeHiddenMetadata.Get()
	}
	if o.FilterValue.IsSet() {
		toSerialize["filterValue"] = o.FilterValue.Get()
	}
	if !IsNil(o.SearchArea) {
		toSerialize["searchArea"] = o.SearchArea
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.StartIndex) {
		toSerialize["startIndex"] = o.StartIndex
	}
	if o.SortBy.IsSet() {
		toSerialize["sortBy"] = o.SortBy.Get()
	}
	if !IsNil(o.SortOrder) {
		toSerialize["sortOrder"] = o.SortOrder
	}
	return toSerialize, nil
}

type NullableRoomsMetadataSearchRequestDto struct {
	value *RoomsMetadataSearchRequestDto
	isSet bool
}

func (v NullableRoomsMetadataSearchRequestDto) Get() *RoomsMetadataSearchRequestDto {
	return v.value
}

func (v *NullableRoomsMetadataSearchRequestDto) Set(val *RoomsMetadataSearchRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableRoomsMetadataSearchRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableRoomsMetadataSearchRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoomsMetadataSearchRequestDto(val *RoomsMetadataSearchRequestDto) *NullableRoomsMetadataSearchRequestDto {
	return &NullableRoomsMetadataSearchRequestDto{value: val, isSet: true}
}

func (v NullableRoomsMetadataSearchRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoomsMetadataSearchRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

