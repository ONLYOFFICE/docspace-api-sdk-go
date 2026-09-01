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
	"bytes"
	"fmt"
)

// checks if the QuotaDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuotaDto{}

// QuotaDto The quota information.
type QuotaDto struct {
	// The quota ID.
	Id int32 `json:"id"`
	// The quota title.
	Title NullableString `json:"title,omitempty"`
	// The price parameters.
	Price PriceDto `json:"price"`
	// Specifies if the quota is nonprofit or not.
	NonProfit bool `json:"nonProfit"`
	// Specifies if the quota is free or not.
	Free bool `json:"free"`
	// Specifies if the quota is trial or not.
	Trial bool `json:"trial"`
	// The list of tenant quota features.
	Features []TenantQuotaFeatureDto `json:"features"`
	// The user quota.
	UsersQuota *TenantEntityQuotaSettings `json:"usersQuota,omitempty"`
	// The room quota.
	RoomsQuota *TenantEntityQuotaSettings `json:"roomsQuota,omitempty"`
	// The ai agent quota.
	AiAgentsQuota *TenantEntityQuotaSettings `json:"aiAgentsQuota,omitempty"`
	// The tenant custom quota.
	TenantCustomQuota *TenantQuotaSettings `json:"tenantCustomQuota,omitempty"`
	// The due date.
	DueDate NullableTime `json:"dueDate,omitempty"`
}

type _QuotaDto QuotaDto

// NewQuotaDto instantiates a new QuotaDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuotaDto(id int32, price PriceDto, nonProfit bool, free bool, trial bool, features []TenantQuotaFeatureDto) *QuotaDto {
	this := QuotaDto{}
	this.Id = id
	this.Price = price
	this.NonProfit = nonProfit
	this.Free = free
	this.Trial = trial
	this.Features = features
	return &this
}

// NewQuotaDtoWithDefaults instantiates a new QuotaDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuotaDtoWithDefaults() *QuotaDto {
	this := QuotaDto{}
	return &this
}

// GetId returns the Id field value
func (o *QuotaDto) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *QuotaDto) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *QuotaDto) SetId(v int32) {
	o.Id = v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QuotaDto) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QuotaDto) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *QuotaDto) IsTitleSet() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *QuotaDto) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *QuotaDto) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *QuotaDto) UnsetTitle() {
	o.Title.Unset()
}

// GetPrice returns the Price field value
func (o *QuotaDto) GetPrice() PriceDto {
	if o == nil {
		var ret PriceDto
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *QuotaDto) GetPriceOk() (*PriceDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *QuotaDto) SetPrice(v PriceDto) {
	o.Price = v
}

// GetNonProfit returns the NonProfit field value
func (o *QuotaDto) GetNonProfit() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.NonProfit
}

// GetNonProfitOk returns a tuple with the NonProfit field value
// and a boolean to check if the value has been set.
func (o *QuotaDto) GetNonProfitOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NonProfit, true
}

// SetNonProfit sets field value
func (o *QuotaDto) SetNonProfit(v bool) {
	o.NonProfit = v
}

// GetFree returns the Free field value
func (o *QuotaDto) GetFree() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Free
}

// GetFreeOk returns a tuple with the Free field value
// and a boolean to check if the value has been set.
func (o *QuotaDto) GetFreeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Free, true
}

// SetFree sets field value
func (o *QuotaDto) SetFree(v bool) {
	o.Free = v
}

// GetTrial returns the Trial field value
func (o *QuotaDto) GetTrial() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Trial
}

// GetTrialOk returns a tuple with the Trial field value
// and a boolean to check if the value has been set.
func (o *QuotaDto) GetTrialOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Trial, true
}

// SetTrial sets field value
func (o *QuotaDto) SetTrial(v bool) {
	o.Trial = v
}

// GetFeatures returns the Features field value
// If the value is explicit nil, the zero value for []TenantQuotaFeatureDto will be returned
func (o *QuotaDto) GetFeatures() []TenantQuotaFeatureDto {
	if o == nil {
		var ret []TenantQuotaFeatureDto
		return ret
	}

	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QuotaDto) GetFeaturesOk() ([]TenantQuotaFeatureDto, bool) {
	if o == nil || IsNil(o.Features) {
		return nil, false
	}
	return o.Features, true
}

// SetFeatures sets field value
func (o *QuotaDto) SetFeatures(v []TenantQuotaFeatureDto) {
	o.Features = v
}

// GetUsersQuota returns the UsersQuota field value if set, zero value otherwise.
func (o *QuotaDto) GetUsersQuota() TenantEntityQuotaSettings {
	if o == nil || IsNil(o.UsersQuota) {
		var ret TenantEntityQuotaSettings
		return ret
	}
	return *o.UsersQuota
}

// GetUsersQuotaOk returns a tuple with the UsersQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaDto) GetUsersQuotaOk() (*TenantEntityQuotaSettings, bool) {
	if o == nil || IsNil(o.UsersQuota) {
		return nil, false
	}
	return o.UsersQuota, true
}

// HasUsersQuota returns a boolean if a field has been set.
func (o *QuotaDto) IsUsersQuotaSet() bool {
	if o != nil && !IsNil(o.UsersQuota) {
		return true
	}

	return false
}

// SetUsersQuota gets a reference to the given TenantEntityQuotaSettings and assigns it to the UsersQuota field.
func (o *QuotaDto) SetUsersQuota(v TenantEntityQuotaSettings) {
	o.UsersQuota = &v
}

// GetRoomsQuota returns the RoomsQuota field value if set, zero value otherwise.
func (o *QuotaDto) GetRoomsQuota() TenantEntityQuotaSettings {
	if o == nil || IsNil(o.RoomsQuota) {
		var ret TenantEntityQuotaSettings
		return ret
	}
	return *o.RoomsQuota
}

// GetRoomsQuotaOk returns a tuple with the RoomsQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaDto) GetRoomsQuotaOk() (*TenantEntityQuotaSettings, bool) {
	if o == nil || IsNil(o.RoomsQuota) {
		return nil, false
	}
	return o.RoomsQuota, true
}

// HasRoomsQuota returns a boolean if a field has been set.
func (o *QuotaDto) IsRoomsQuotaSet() bool {
	if o != nil && !IsNil(o.RoomsQuota) {
		return true
	}

	return false
}

// SetRoomsQuota gets a reference to the given TenantEntityQuotaSettings and assigns it to the RoomsQuota field.
func (o *QuotaDto) SetRoomsQuota(v TenantEntityQuotaSettings) {
	o.RoomsQuota = &v
}

// GetAiAgentsQuota returns the AiAgentsQuota field value if set, zero value otherwise.
func (o *QuotaDto) GetAiAgentsQuota() TenantEntityQuotaSettings {
	if o == nil || IsNil(o.AiAgentsQuota) {
		var ret TenantEntityQuotaSettings
		return ret
	}
	return *o.AiAgentsQuota
}

// GetAiAgentsQuotaOk returns a tuple with the AiAgentsQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaDto) GetAiAgentsQuotaOk() (*TenantEntityQuotaSettings, bool) {
	if o == nil || IsNil(o.AiAgentsQuota) {
		return nil, false
	}
	return o.AiAgentsQuota, true
}

// HasAiAgentsQuota returns a boolean if a field has been set.
func (o *QuotaDto) IsAiAgentsQuotaSet() bool {
	if o != nil && !IsNil(o.AiAgentsQuota) {
		return true
	}

	return false
}

// SetAiAgentsQuota gets a reference to the given TenantEntityQuotaSettings and assigns it to the AiAgentsQuota field.
func (o *QuotaDto) SetAiAgentsQuota(v TenantEntityQuotaSettings) {
	o.AiAgentsQuota = &v
}

// GetTenantCustomQuota returns the TenantCustomQuota field value if set, zero value otherwise.
func (o *QuotaDto) GetTenantCustomQuota() TenantQuotaSettings {
	if o == nil || IsNil(o.TenantCustomQuota) {
		var ret TenantQuotaSettings
		return ret
	}
	return *o.TenantCustomQuota
}

// GetTenantCustomQuotaOk returns a tuple with the TenantCustomQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotaDto) GetTenantCustomQuotaOk() (*TenantQuotaSettings, bool) {
	if o == nil || IsNil(o.TenantCustomQuota) {
		return nil, false
	}
	return o.TenantCustomQuota, true
}

// HasTenantCustomQuota returns a boolean if a field has been set.
func (o *QuotaDto) IsTenantCustomQuotaSet() bool {
	if o != nil && !IsNil(o.TenantCustomQuota) {
		return true
	}

	return false
}

// SetTenantCustomQuota gets a reference to the given TenantQuotaSettings and assigns it to the TenantCustomQuota field.
func (o *QuotaDto) SetTenantCustomQuota(v TenantQuotaSettings) {
	o.TenantCustomQuota = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QuotaDto) GetDueDate() time.Time {
	if o == nil || IsNil(o.DueDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DueDate.Get()
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QuotaDto) GetDueDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DueDate.Get(), o.DueDate.IsSet()
}

// HasDueDate returns a boolean if a field has been set.
func (o *QuotaDto) IsDueDateSet() bool {
	if o != nil && o.DueDate.IsSet() {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given NullableTime and assigns it to the DueDate field.
func (o *QuotaDto) SetDueDate(v time.Time) {
	o.DueDate.Set(&v)
}
// SetDueDateNil sets the value for DueDate to be an explicit nil
func (o *QuotaDto) SetDueDateNil() {
	o.DueDate.Set(nil)
}

// UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
func (o *QuotaDto) UnsetDueDate() {
	o.DueDate.Unset()
}

func (o QuotaDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuotaDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	toSerialize["price"] = o.Price
	toSerialize["nonProfit"] = o.NonProfit
	toSerialize["free"] = o.Free
	toSerialize["trial"] = o.Trial
	if o.Features != nil {
		toSerialize["features"] = o.Features
	}
	if !IsNil(o.UsersQuota) {
		toSerialize["usersQuota"] = o.UsersQuota
	}
	if !IsNil(o.RoomsQuota) {
		toSerialize["roomsQuota"] = o.RoomsQuota
	}
	if !IsNil(o.AiAgentsQuota) {
		toSerialize["aiAgentsQuota"] = o.AiAgentsQuota
	}
	if !IsNil(o.TenantCustomQuota) {
		toSerialize["tenantCustomQuota"] = o.TenantCustomQuota
	}
	if o.DueDate.IsSet() {
		toSerialize["dueDate"] = o.DueDate.Get()
	}
	return toSerialize, nil
}

func (o *QuotaDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"price",
		"nonProfit",
		"free",
		"trial",
		"features",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varQuotaDto := _QuotaDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varQuotaDto)

	if err != nil {
		return err
	}

	*o = QuotaDto(varQuotaDto)

	return err
}

type NullableQuotaDto struct {
	value *QuotaDto
	isSet bool
}

func (v NullableQuotaDto) Get() *QuotaDto {
	return v.value
}

func (v *NullableQuotaDto) Set(val *QuotaDto) {
	v.value = val
	v.isSet = true
}

func (v NullableQuotaDto) IsSet() bool {
	return v.isSet
}

func (v *NullableQuotaDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuotaDto(val *QuotaDto) *NullableQuotaDto {
	return &NullableQuotaDto{value: val, isSet: true}
}

func (v NullableQuotaDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuotaDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

