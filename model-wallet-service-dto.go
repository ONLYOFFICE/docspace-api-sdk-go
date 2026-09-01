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

// checks if the WalletServiceDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WalletServiceDto{}

// WalletServiceDto The wallet service information.
type WalletServiceDto struct {
	// The quota ID.
	Id int32 `json:"id"`
	// The quota title.
	Title *string `json:"title,omitempty"`
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
	DueDate *time.Time `json:"dueDate,omitempty"`
	// The list of inner services.
	InnerServices []WalletServiceDto `json:"innerServices,omitempty"`
	// The service name.
	ServiceName NullableString `json:"serviceName,omitempty"`
}

type _WalletServiceDto WalletServiceDto

// NewWalletServiceDto instantiates a new WalletServiceDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletServiceDto(id int32, price PriceDto, nonProfit bool, free bool, trial bool, features []TenantQuotaFeatureDto) *WalletServiceDto {
	this := WalletServiceDto{}
	this.Id = id
	this.Price = price
	this.NonProfit = nonProfit
	this.Free = free
	this.Trial = trial
	this.Features = features
	return &this
}

// NewWalletServiceDtoWithDefaults instantiates a new WalletServiceDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletServiceDtoWithDefaults() *WalletServiceDto {
	this := WalletServiceDto{}
	return &this
}

// GetId returns the Id field value
func (o *WalletServiceDto) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WalletServiceDto) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WalletServiceDto) SetId(v int32) {
	o.Id = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *WalletServiceDto) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletServiceDto) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *WalletServiceDto) IsTitleSet() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *WalletServiceDto) SetTitle(v string) {
	o.Title = &v
}

// GetPrice returns the Price field value
func (o *WalletServiceDto) GetPrice() PriceDto {
	if o == nil {
		var ret PriceDto
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *WalletServiceDto) GetPriceOk() (*PriceDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *WalletServiceDto) SetPrice(v PriceDto) {
	o.Price = v
}

// GetNonProfit returns the NonProfit field value
func (o *WalletServiceDto) GetNonProfit() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.NonProfit
}

// GetNonProfitOk returns a tuple with the NonProfit field value
// and a boolean to check if the value has been set.
func (o *WalletServiceDto) GetNonProfitOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NonProfit, true
}

// SetNonProfit sets field value
func (o *WalletServiceDto) SetNonProfit(v bool) {
	o.NonProfit = v
}

// GetFree returns the Free field value
func (o *WalletServiceDto) GetFree() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Free
}

// GetFreeOk returns a tuple with the Free field value
// and a boolean to check if the value has been set.
func (o *WalletServiceDto) GetFreeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Free, true
}

// SetFree sets field value
func (o *WalletServiceDto) SetFree(v bool) {
	o.Free = v
}

// GetTrial returns the Trial field value
func (o *WalletServiceDto) GetTrial() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Trial
}

// GetTrialOk returns a tuple with the Trial field value
// and a boolean to check if the value has been set.
func (o *WalletServiceDto) GetTrialOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Trial, true
}

// SetTrial sets field value
func (o *WalletServiceDto) SetTrial(v bool) {
	o.Trial = v
}

// GetFeatures returns the Features field value
func (o *WalletServiceDto) GetFeatures() []TenantQuotaFeatureDto {
	if o == nil {
		var ret []TenantQuotaFeatureDto
		return ret
	}

	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value
// and a boolean to check if the value has been set.
func (o *WalletServiceDto) GetFeaturesOk() ([]TenantQuotaFeatureDto, bool) {
	if o == nil {
		return nil, false
	}
	return o.Features, true
}

// SetFeatures sets field value
func (o *WalletServiceDto) SetFeatures(v []TenantQuotaFeatureDto) {
	o.Features = v
}

// GetUsersQuota returns the UsersQuota field value if set, zero value otherwise.
func (o *WalletServiceDto) GetUsersQuota() TenantEntityQuotaSettings {
	if o == nil || IsNil(o.UsersQuota) {
		var ret TenantEntityQuotaSettings
		return ret
	}
	return *o.UsersQuota
}

// GetUsersQuotaOk returns a tuple with the UsersQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletServiceDto) GetUsersQuotaOk() (*TenantEntityQuotaSettings, bool) {
	if o == nil || IsNil(o.UsersQuota) {
		return nil, false
	}
	return o.UsersQuota, true
}

// HasUsersQuota returns a boolean if a field has been set.
func (o *WalletServiceDto) IsUsersQuotaSet() bool {
	if o != nil && !IsNil(o.UsersQuota) {
		return true
	}

	return false
}

// SetUsersQuota gets a reference to the given TenantEntityQuotaSettings and assigns it to the UsersQuota field.
func (o *WalletServiceDto) SetUsersQuota(v TenantEntityQuotaSettings) {
	o.UsersQuota = &v
}

// GetRoomsQuota returns the RoomsQuota field value if set, zero value otherwise.
func (o *WalletServiceDto) GetRoomsQuota() TenantEntityQuotaSettings {
	if o == nil || IsNil(o.RoomsQuota) {
		var ret TenantEntityQuotaSettings
		return ret
	}
	return *o.RoomsQuota
}

// GetRoomsQuotaOk returns a tuple with the RoomsQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletServiceDto) GetRoomsQuotaOk() (*TenantEntityQuotaSettings, bool) {
	if o == nil || IsNil(o.RoomsQuota) {
		return nil, false
	}
	return o.RoomsQuota, true
}

// HasRoomsQuota returns a boolean if a field has been set.
func (o *WalletServiceDto) IsRoomsQuotaSet() bool {
	if o != nil && !IsNil(o.RoomsQuota) {
		return true
	}

	return false
}

// SetRoomsQuota gets a reference to the given TenantEntityQuotaSettings and assigns it to the RoomsQuota field.
func (o *WalletServiceDto) SetRoomsQuota(v TenantEntityQuotaSettings) {
	o.RoomsQuota = &v
}

// GetAiAgentsQuota returns the AiAgentsQuota field value if set, zero value otherwise.
func (o *WalletServiceDto) GetAiAgentsQuota() TenantEntityQuotaSettings {
	if o == nil || IsNil(o.AiAgentsQuota) {
		var ret TenantEntityQuotaSettings
		return ret
	}
	return *o.AiAgentsQuota
}

// GetAiAgentsQuotaOk returns a tuple with the AiAgentsQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletServiceDto) GetAiAgentsQuotaOk() (*TenantEntityQuotaSettings, bool) {
	if o == nil || IsNil(o.AiAgentsQuota) {
		return nil, false
	}
	return o.AiAgentsQuota, true
}

// HasAiAgentsQuota returns a boolean if a field has been set.
func (o *WalletServiceDto) IsAiAgentsQuotaSet() bool {
	if o != nil && !IsNil(o.AiAgentsQuota) {
		return true
	}

	return false
}

// SetAiAgentsQuota gets a reference to the given TenantEntityQuotaSettings and assigns it to the AiAgentsQuota field.
func (o *WalletServiceDto) SetAiAgentsQuota(v TenantEntityQuotaSettings) {
	o.AiAgentsQuota = &v
}

// GetTenantCustomQuota returns the TenantCustomQuota field value if set, zero value otherwise.
func (o *WalletServiceDto) GetTenantCustomQuota() TenantQuotaSettings {
	if o == nil || IsNil(o.TenantCustomQuota) {
		var ret TenantQuotaSettings
		return ret
	}
	return *o.TenantCustomQuota
}

// GetTenantCustomQuotaOk returns a tuple with the TenantCustomQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletServiceDto) GetTenantCustomQuotaOk() (*TenantQuotaSettings, bool) {
	if o == nil || IsNil(o.TenantCustomQuota) {
		return nil, false
	}
	return o.TenantCustomQuota, true
}

// HasTenantCustomQuota returns a boolean if a field has been set.
func (o *WalletServiceDto) IsTenantCustomQuotaSet() bool {
	if o != nil && !IsNil(o.TenantCustomQuota) {
		return true
	}

	return false
}

// SetTenantCustomQuota gets a reference to the given TenantQuotaSettings and assigns it to the TenantCustomQuota field.
func (o *WalletServiceDto) SetTenantCustomQuota(v TenantQuotaSettings) {
	o.TenantCustomQuota = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *WalletServiceDto) GetDueDate() time.Time {
	if o == nil || IsNil(o.DueDate) {
		var ret time.Time
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletServiceDto) GetDueDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DueDate) {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *WalletServiceDto) IsDueDateSet() bool {
	if o != nil && !IsNil(o.DueDate) {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given time.Time and assigns it to the DueDate field.
func (o *WalletServiceDto) SetDueDate(v time.Time) {
	o.DueDate = &v
}

// GetInnerServices returns the InnerServices field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WalletServiceDto) GetInnerServices() []WalletServiceDto {
	if o == nil {
		var ret []WalletServiceDto
		return ret
	}
	return o.InnerServices
}

// GetInnerServicesOk returns a tuple with the InnerServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WalletServiceDto) GetInnerServicesOk() ([]WalletServiceDto, bool) {
	if o == nil || IsNil(o.InnerServices) {
		return nil, false
	}
	return o.InnerServices, true
}

// HasInnerServices returns a boolean if a field has been set.
func (o *WalletServiceDto) IsInnerServicesSet() bool {
	if o != nil && !IsNil(o.InnerServices) {
		return true
	}

	return false
}

// SetInnerServices gets a reference to the given []WalletServiceDto and assigns it to the InnerServices field.
func (o *WalletServiceDto) SetInnerServices(v []WalletServiceDto) {
	o.InnerServices = v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WalletServiceDto) GetServiceName() string {
	if o == nil || IsNil(o.ServiceName.Get()) {
		var ret string
		return ret
	}
	return *o.ServiceName.Get()
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WalletServiceDto) GetServiceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceName.Get(), o.ServiceName.IsSet()
}

// HasServiceName returns a boolean if a field has been set.
func (o *WalletServiceDto) IsServiceNameSet() bool {
	if o != nil && o.ServiceName.IsSet() {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given NullableString and assigns it to the ServiceName field.
func (o *WalletServiceDto) SetServiceName(v string) {
	o.ServiceName.Set(&v)
}
// SetServiceNameNil sets the value for ServiceName to be an explicit nil
func (o *WalletServiceDto) SetServiceNameNil() {
	o.ServiceName.Set(nil)
}

// UnsetServiceName ensures that no value is present for ServiceName, not even an explicit nil
func (o *WalletServiceDto) UnsetServiceName() {
	o.ServiceName.Unset()
}

func (o WalletServiceDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WalletServiceDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	toSerialize["price"] = o.Price
	toSerialize["nonProfit"] = o.NonProfit
	toSerialize["free"] = o.Free
	toSerialize["trial"] = o.Trial
	toSerialize["features"] = o.Features
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
	if !IsNil(o.DueDate) {
		toSerialize["dueDate"] = o.DueDate
	}
	if o.InnerServices != nil {
		toSerialize["innerServices"] = o.InnerServices
	}
	if o.ServiceName.IsSet() {
		toSerialize["serviceName"] = o.ServiceName.Get()
	}
	return toSerialize, nil
}

func (o *WalletServiceDto) UnmarshalJSON(data []byte) (err error) {
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

	varWalletServiceDto := _WalletServiceDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varWalletServiceDto)

	if err != nil {
		return err
	}

	*o = WalletServiceDto(varWalletServiceDto)

	return err
}

type NullableWalletServiceDto struct {
	value *WalletServiceDto
	isSet bool
}

func (v NullableWalletServiceDto) Get() *WalletServiceDto {
	return v.value
}

func (v *NullableWalletServiceDto) Set(val *WalletServiceDto) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletServiceDto) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletServiceDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletServiceDto(val *WalletServiceDto) *NullableWalletServiceDto {
	return &NullableWalletServiceDto{value: val, isSet: true}
}

func (v NullableWalletServiceDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletServiceDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

