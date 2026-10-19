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

// checks if the MetadataFilterConditionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataFilterConditionRequest{}

// MetadataFilterConditionRequest One metadata filter condition as the clients send it: an element of the metadataFilters JSON of the listings and of  the body of the search endpoints. All conditions are combined with AND.
type MetadataFilterConditionRequest struct {
	// The ID of the template field the condition is on. A custom field is addressed by ASC.Files.Core.MetadataFilterConditionRequest.Name instead.
	FieldId *int32 `json:"fieldId,omitempty"`
	// The name of a custom field, for the conditions on the custom fields, which have no identifier outside. Either the  field ID or the name is given; the name is matched without regard to case.
	Name NullableString `json:"name,omitempty"`
	// The operator, one of ASC.Files.Core.MetadataFilterOperators: eq or contains for a string field, eq, range,  gt, gte, lt or lte for a number field, range, gt, gte, lt or lte for a date field, in for a  choice field. Optional: an omitted operator is read as eq for a string, as range for a date, as in for a  choice, and for a number as eq when a value is given or as range when a bound is. A present one has to belong  to the field type and to the arguments: eq takes a value, range takes bounds.
	Op NullableString `json:"op,omitempty"`
	// The value of a single-valued condition: the whole or a part of a string field, the exact value of a number field,  or the bound of a gt, gte, lt or lte condition on a number or a date field. A JSON number is accepted as  well as a string. A date given without a time (2026-06-30) names the whole day (UTC).
	Value NullableString `json:"value,omitempty"`
	// The inclusive lower bound of a range. A date given without a time (2026-06-01) is the start of that day (UTC).
	From NullableString `json:"from,omitempty"`
	// The inclusive upper bound of a range. A date given without a time (2026-06-30) covers the whole day (UTC);  a value with a time is an instant and is taken as is.
	To NullableString `json:"to,omitempty"`
	// The options any of which the choice field must hold.
	OptionIds []string `json:"optionIds,omitempty"`
}

// NewMetadataFilterConditionRequest instantiates a new MetadataFilterConditionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataFilterConditionRequest() *MetadataFilterConditionRequest {
	this := MetadataFilterConditionRequest{}
	return &this
}

// NewMetadataFilterConditionRequestWithDefaults instantiates a new MetadataFilterConditionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataFilterConditionRequestWithDefaults() *MetadataFilterConditionRequest {
	this := MetadataFilterConditionRequest{}
	return &this
}

// GetFieldId returns the FieldId field value if set, zero value otherwise.
func (o *MetadataFilterConditionRequest) GetFieldId() int32 {
	if o == nil || IsNil(o.FieldId) {
		var ret int32
		return ret
	}
	return *o.FieldId
}

// GetFieldIdOk returns a tuple with the FieldId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataFilterConditionRequest) GetFieldIdOk() (*int32, bool) {
	if o == nil || IsNil(o.FieldId) {
		return nil, false
	}
	return o.FieldId, true
}

// HasFieldId returns a boolean if a field has been set.
func (o *MetadataFilterConditionRequest) IsFieldIdSet() bool {
	if o != nil && !IsNil(o.FieldId) {
		return true
	}

	return false
}

// SetFieldId gets a reference to the given int32 and assigns it to the FieldId field.
func (o *MetadataFilterConditionRequest) SetFieldId(v int32) {
	o.FieldId = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataFilterConditionRequest) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataFilterConditionRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *MetadataFilterConditionRequest) IsNameSet() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *MetadataFilterConditionRequest) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *MetadataFilterConditionRequest) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *MetadataFilterConditionRequest) UnsetName() {
	o.Name.Unset()
}

// GetOp returns the Op field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataFilterConditionRequest) GetOp() string {
	if o == nil || IsNil(o.Op.Get()) {
		var ret string
		return ret
	}
	return *o.Op.Get()
}

// GetOpOk returns a tuple with the Op field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataFilterConditionRequest) GetOpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Op.Get(), o.Op.IsSet()
}

// HasOp returns a boolean if a field has been set.
func (o *MetadataFilterConditionRequest) IsOpSet() bool {
	if o != nil && o.Op.IsSet() {
		return true
	}

	return false
}

// SetOp gets a reference to the given NullableString and assigns it to the Op field.
func (o *MetadataFilterConditionRequest) SetOp(v string) {
	o.Op.Set(&v)
}
// SetOpNil sets the value for Op to be an explicit nil
func (o *MetadataFilterConditionRequest) SetOpNil() {
	o.Op.Set(nil)
}

// UnsetOp ensures that no value is present for Op, not even an explicit nil
func (o *MetadataFilterConditionRequest) UnsetOp() {
	o.Op.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataFilterConditionRequest) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataFilterConditionRequest) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *MetadataFilterConditionRequest) IsValueSet() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *MetadataFilterConditionRequest) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *MetadataFilterConditionRequest) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *MetadataFilterConditionRequest) UnsetValue() {
	o.Value.Unset()
}

// GetFrom returns the From field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataFilterConditionRequest) GetFrom() string {
	if o == nil || IsNil(o.From.Get()) {
		var ret string
		return ret
	}
	return *o.From.Get()
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataFilterConditionRequest) GetFromOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.From.Get(), o.From.IsSet()
}

// HasFrom returns a boolean if a field has been set.
func (o *MetadataFilterConditionRequest) IsFromSet() bool {
	if o != nil && o.From.IsSet() {
		return true
	}

	return false
}

// SetFrom gets a reference to the given NullableString and assigns it to the From field.
func (o *MetadataFilterConditionRequest) SetFrom(v string) {
	o.From.Set(&v)
}
// SetFromNil sets the value for From to be an explicit nil
func (o *MetadataFilterConditionRequest) SetFromNil() {
	o.From.Set(nil)
}

// UnsetFrom ensures that no value is present for From, not even an explicit nil
func (o *MetadataFilterConditionRequest) UnsetFrom() {
	o.From.Unset()
}

// GetTo returns the To field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataFilterConditionRequest) GetTo() string {
	if o == nil || IsNil(o.To.Get()) {
		var ret string
		return ret
	}
	return *o.To.Get()
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataFilterConditionRequest) GetToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.To.Get(), o.To.IsSet()
}

// HasTo returns a boolean if a field has been set.
func (o *MetadataFilterConditionRequest) IsToSet() bool {
	if o != nil && o.To.IsSet() {
		return true
	}

	return false
}

// SetTo gets a reference to the given NullableString and assigns it to the To field.
func (o *MetadataFilterConditionRequest) SetTo(v string) {
	o.To.Set(&v)
}
// SetToNil sets the value for To to be an explicit nil
func (o *MetadataFilterConditionRequest) SetToNil() {
	o.To.Set(nil)
}

// UnsetTo ensures that no value is present for To, not even an explicit nil
func (o *MetadataFilterConditionRequest) UnsetTo() {
	o.To.Unset()
}

// GetOptionIds returns the OptionIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataFilterConditionRequest) GetOptionIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OptionIds
}

// GetOptionIdsOk returns a tuple with the OptionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataFilterConditionRequest) GetOptionIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.OptionIds) {
		return nil, false
	}
	return o.OptionIds, true
}

// HasOptionIds returns a boolean if a field has been set.
func (o *MetadataFilterConditionRequest) IsOptionIdsSet() bool {
	if o != nil && !IsNil(o.OptionIds) {
		return true
	}

	return false
}

// SetOptionIds gets a reference to the given []string and assigns it to the OptionIds field.
func (o *MetadataFilterConditionRequest) SetOptionIds(v []string) {
	o.OptionIds = v
}

func (o MetadataFilterConditionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetadataFilterConditionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FieldId) {
		toSerialize["fieldId"] = o.FieldId
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Op.IsSet() {
		toSerialize["op"] = o.Op.Get()
	}
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	if o.From.IsSet() {
		toSerialize["from"] = o.From.Get()
	}
	if o.To.IsSet() {
		toSerialize["to"] = o.To.Get()
	}
	if o.OptionIds != nil {
		toSerialize["optionIds"] = o.OptionIds
	}
	return toSerialize, nil
}

type NullableMetadataFilterConditionRequest struct {
	value *MetadataFilterConditionRequest
	isSet bool
}

func (v NullableMetadataFilterConditionRequest) Get() *MetadataFilterConditionRequest {
	return v.value
}

func (v *NullableMetadataFilterConditionRequest) Set(val *MetadataFilterConditionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataFilterConditionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataFilterConditionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataFilterConditionRequest(val *MetadataFilterConditionRequest) *NullableMetadataFilterConditionRequest {
	return &NullableMetadataFilterConditionRequest{value: val, isSet: true}
}

func (v NullableMetadataFilterConditionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataFilterConditionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

