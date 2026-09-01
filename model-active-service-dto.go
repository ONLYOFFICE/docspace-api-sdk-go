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

// checks if the ActiveServiceDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActiveServiceDto{}

// ActiveServiceDto Represents an active wallet service (quota) of the current portal.
type ActiveServiceDto struct {
	// The name of the service.
	Service NullableString `json:"service,omitempty"`
	// The unit of measurement for the service.
	ServiceUnit NullableString `json:"serviceUnit,omitempty"`
	// Indicates whether the service is subscription-based.
	Subscription *bool `json:"subscription,omitempty"`
	// The title of the service.
	Title NullableString `json:"title,omitempty"`
	// The service limit. Populated only for the subscription-based services.
	Limit NullableInt32 `json:"limit,omitempty"`
	// The current service usage. Populated only for the subscription-based services.
	Used NullableInt32 `json:"used,omitempty"`
}

// NewActiveServiceDto instantiates a new ActiveServiceDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActiveServiceDto() *ActiveServiceDto {
	this := ActiveServiceDto{}
	return &this
}

// NewActiveServiceDtoWithDefaults instantiates a new ActiveServiceDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActiveServiceDtoWithDefaults() *ActiveServiceDto {
	this := ActiveServiceDto{}
	return &this
}

// GetService returns the Service field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActiveServiceDto) GetService() string {
	if o == nil || IsNil(o.Service.Get()) {
		var ret string
		return ret
	}
	return *o.Service.Get()
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveServiceDto) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Service.Get(), o.Service.IsSet()
}

// HasService returns a boolean if a field has been set.
func (o *ActiveServiceDto) IsServiceSet() bool {
	if o != nil && o.Service.IsSet() {
		return true
	}

	return false
}

// SetService gets a reference to the given NullableString and assigns it to the Service field.
func (o *ActiveServiceDto) SetService(v string) {
	o.Service.Set(&v)
}
// SetServiceNil sets the value for Service to be an explicit nil
func (o *ActiveServiceDto) SetServiceNil() {
	o.Service.Set(nil)
}

// UnsetService ensures that no value is present for Service, not even an explicit nil
func (o *ActiveServiceDto) UnsetService() {
	o.Service.Unset()
}

// GetServiceUnit returns the ServiceUnit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActiveServiceDto) GetServiceUnit() string {
	if o == nil || IsNil(o.ServiceUnit.Get()) {
		var ret string
		return ret
	}
	return *o.ServiceUnit.Get()
}

// GetServiceUnitOk returns a tuple with the ServiceUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveServiceDto) GetServiceUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceUnit.Get(), o.ServiceUnit.IsSet()
}

// HasServiceUnit returns a boolean if a field has been set.
func (o *ActiveServiceDto) IsServiceUnitSet() bool {
	if o != nil && o.ServiceUnit.IsSet() {
		return true
	}

	return false
}

// SetServiceUnit gets a reference to the given NullableString and assigns it to the ServiceUnit field.
func (o *ActiveServiceDto) SetServiceUnit(v string) {
	o.ServiceUnit.Set(&v)
}
// SetServiceUnitNil sets the value for ServiceUnit to be an explicit nil
func (o *ActiveServiceDto) SetServiceUnitNil() {
	o.ServiceUnit.Set(nil)
}

// UnsetServiceUnit ensures that no value is present for ServiceUnit, not even an explicit nil
func (o *ActiveServiceDto) UnsetServiceUnit() {
	o.ServiceUnit.Unset()
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *ActiveServiceDto) GetSubscription() bool {
	if o == nil || IsNil(o.Subscription) {
		var ret bool
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveServiceDto) GetSubscriptionOk() (*bool, bool) {
	if o == nil || IsNil(o.Subscription) {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *ActiveServiceDto) IsSubscriptionSet() bool {
	if o != nil && !IsNil(o.Subscription) {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given bool and assigns it to the Subscription field.
func (o *ActiveServiceDto) SetSubscription(v bool) {
	o.Subscription = &v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActiveServiceDto) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveServiceDto) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *ActiveServiceDto) IsTitleSet() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *ActiveServiceDto) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *ActiveServiceDto) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *ActiveServiceDto) UnsetTitle() {
	o.Title.Unset()
}

// GetLimit returns the Limit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActiveServiceDto) GetLimit() int32 {
	if o == nil || IsNil(o.Limit.Get()) {
		var ret int32
		return ret
	}
	return *o.Limit.Get()
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveServiceDto) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Limit.Get(), o.Limit.IsSet()
}

// HasLimit returns a boolean if a field has been set.
func (o *ActiveServiceDto) IsLimitSet() bool {
	if o != nil && o.Limit.IsSet() {
		return true
	}

	return false
}

// SetLimit gets a reference to the given NullableInt32 and assigns it to the Limit field.
func (o *ActiveServiceDto) SetLimit(v int32) {
	o.Limit.Set(&v)
}
// SetLimitNil sets the value for Limit to be an explicit nil
func (o *ActiveServiceDto) SetLimitNil() {
	o.Limit.Set(nil)
}

// UnsetLimit ensures that no value is present for Limit, not even an explicit nil
func (o *ActiveServiceDto) UnsetLimit() {
	o.Limit.Unset()
}

// GetUsed returns the Used field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActiveServiceDto) GetUsed() int32 {
	if o == nil || IsNil(o.Used.Get()) {
		var ret int32
		return ret
	}
	return *o.Used.Get()
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveServiceDto) GetUsedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Used.Get(), o.Used.IsSet()
}

// HasUsed returns a boolean if a field has been set.
func (o *ActiveServiceDto) IsUsedSet() bool {
	if o != nil && o.Used.IsSet() {
		return true
	}

	return false
}

// SetUsed gets a reference to the given NullableInt32 and assigns it to the Used field.
func (o *ActiveServiceDto) SetUsed(v int32) {
	o.Used.Set(&v)
}
// SetUsedNil sets the value for Used to be an explicit nil
func (o *ActiveServiceDto) SetUsedNil() {
	o.Used.Set(nil)
}

// UnsetUsed ensures that no value is present for Used, not even an explicit nil
func (o *ActiveServiceDto) UnsetUsed() {
	o.Used.Unset()
}

func (o ActiveServiceDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActiveServiceDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Service.IsSet() {
		toSerialize["service"] = o.Service.Get()
	}
	if o.ServiceUnit.IsSet() {
		toSerialize["serviceUnit"] = o.ServiceUnit.Get()
	}
	if !IsNil(o.Subscription) {
		toSerialize["subscription"] = o.Subscription
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Limit.IsSet() {
		toSerialize["limit"] = o.Limit.Get()
	}
	if o.Used.IsSet() {
		toSerialize["used"] = o.Used.Get()
	}
	return toSerialize, nil
}

type NullableActiveServiceDto struct {
	value *ActiveServiceDto
	isSet bool
}

func (v NullableActiveServiceDto) Get() *ActiveServiceDto {
	return v.value
}

func (v *NullableActiveServiceDto) Set(val *ActiveServiceDto) {
	v.value = val
	v.isSet = true
}

func (v NullableActiveServiceDto) IsSet() bool {
	return v.isSet
}

func (v *NullableActiveServiceDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActiveServiceDto(val *ActiveServiceDto) *NullableActiveServiceDto {
	return &NullableActiveServiceDto{value: val, isSet: true}
}

func (v NullableActiveServiceDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActiveServiceDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

