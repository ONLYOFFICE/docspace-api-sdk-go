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
	"bytes"
	"fmt"
)

// checks if the PaymentSettingsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentSettingsDto{}

// PaymentSettingsDto The payment settings parameters.
type PaymentSettingsDto struct {
	// The email address for sales inquiries and support.
	SalesEmail NullableString `json:"salesEmail"`
	// The URL for accessing the feedback and support resources.
	FeedbackAndSupportUrl NullableString `json:"feedbackAndSupportUrl,omitempty"`
	// The URL for purchasing or upgrading the product.
	BuyUrl NullableString `json:"buyUrl"`
	// Indicates whether the system is running in standalone mode.
	Standalone bool `json:"standalone"`
	CurrentLicense CurrentLicenseInfo `json:"currentLicense"`
	// The maximum quota quantity.
	Max int32 `json:"max"`
}

type _PaymentSettingsDto PaymentSettingsDto

// NewPaymentSettingsDto instantiates a new PaymentSettingsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentSettingsDto(salesEmail NullableString, buyUrl NullableString, standalone bool, currentLicense CurrentLicenseInfo, max int32) *PaymentSettingsDto {
	this := PaymentSettingsDto{}
	this.SalesEmail = salesEmail
	this.BuyUrl = buyUrl
	this.Standalone = standalone
	this.CurrentLicense = currentLicense
	this.Max = max
	return &this
}

// NewPaymentSettingsDtoWithDefaults instantiates a new PaymentSettingsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentSettingsDtoWithDefaults() *PaymentSettingsDto {
	this := PaymentSettingsDto{}
	return &this
}

// GetSalesEmail returns the SalesEmail field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaymentSettingsDto) GetSalesEmail() string {
	if o == nil || o.SalesEmail.Get() == nil {
		var ret string
		return ret
	}

	return *o.SalesEmail.Get()
}

// GetSalesEmailOk returns a tuple with the SalesEmail field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentSettingsDto) GetSalesEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SalesEmail.Get(), o.SalesEmail.IsSet()
}

// SetSalesEmail sets field value
func (o *PaymentSettingsDto) SetSalesEmail(v string) {
	o.SalesEmail.Set(&v)
}

// GetFeedbackAndSupportUrl returns the FeedbackAndSupportUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentSettingsDto) GetFeedbackAndSupportUrl() string {
	if o == nil || IsNil(o.FeedbackAndSupportUrl.Get()) {
		var ret string
		return ret
	}
	return *o.FeedbackAndSupportUrl.Get()
}

// GetFeedbackAndSupportUrlOk returns a tuple with the FeedbackAndSupportUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentSettingsDto) GetFeedbackAndSupportUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FeedbackAndSupportUrl.Get(), o.FeedbackAndSupportUrl.IsSet()
}

// HasFeedbackAndSupportUrl returns a boolean if a field has been set.
func (o *PaymentSettingsDto) IsFeedbackAndSupportUrlSet() bool {
	if o != nil && o.FeedbackAndSupportUrl.IsSet() {
		return true
	}

	return false
}

// SetFeedbackAndSupportUrl gets a reference to the given NullableString and assigns it to the FeedbackAndSupportUrl field.
func (o *PaymentSettingsDto) SetFeedbackAndSupportUrl(v string) {
	o.FeedbackAndSupportUrl.Set(&v)
}
// SetFeedbackAndSupportUrlNil sets the value for FeedbackAndSupportUrl to be an explicit nil
func (o *PaymentSettingsDto) SetFeedbackAndSupportUrlNil() {
	o.FeedbackAndSupportUrl.Set(nil)
}

// UnsetFeedbackAndSupportUrl ensures that no value is present for FeedbackAndSupportUrl, not even an explicit nil
func (o *PaymentSettingsDto) UnsetFeedbackAndSupportUrl() {
	o.FeedbackAndSupportUrl.Unset()
}

// GetBuyUrl returns the BuyUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaymentSettingsDto) GetBuyUrl() string {
	if o == nil || o.BuyUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.BuyUrl.Get()
}

// GetBuyUrlOk returns a tuple with the BuyUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentSettingsDto) GetBuyUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BuyUrl.Get(), o.BuyUrl.IsSet()
}

// SetBuyUrl sets field value
func (o *PaymentSettingsDto) SetBuyUrl(v string) {
	o.BuyUrl.Set(&v)
}

// GetStandalone returns the Standalone field value
func (o *PaymentSettingsDto) GetStandalone() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Standalone
}

// GetStandaloneOk returns a tuple with the Standalone field value
// and a boolean to check if the value has been set.
func (o *PaymentSettingsDto) GetStandaloneOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Standalone, true
}

// SetStandalone sets field value
func (o *PaymentSettingsDto) SetStandalone(v bool) {
	o.Standalone = v
}

// GetCurrentLicense returns the CurrentLicense field value
func (o *PaymentSettingsDto) GetCurrentLicense() CurrentLicenseInfo {
	if o == nil {
		var ret CurrentLicenseInfo
		return ret
	}

	return o.CurrentLicense
}

// GetCurrentLicenseOk returns a tuple with the CurrentLicense field value
// and a boolean to check if the value has been set.
func (o *PaymentSettingsDto) GetCurrentLicenseOk() (*CurrentLicenseInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentLicense, true
}

// SetCurrentLicense sets field value
func (o *PaymentSettingsDto) SetCurrentLicense(v CurrentLicenseInfo) {
	o.CurrentLicense = v
}

// GetMax returns the Max field value
func (o *PaymentSettingsDto) GetMax() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Max
}

// GetMaxOk returns a tuple with the Max field value
// and a boolean to check if the value has been set.
func (o *PaymentSettingsDto) GetMaxOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Max, true
}

// SetMax sets field value
func (o *PaymentSettingsDto) SetMax(v int32) {
	o.Max = v
}

func (o PaymentSettingsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentSettingsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["salesEmail"] = o.SalesEmail.Get()
	if o.FeedbackAndSupportUrl.IsSet() {
		toSerialize["feedbackAndSupportUrl"] = o.FeedbackAndSupportUrl.Get()
	}
	toSerialize["buyUrl"] = o.BuyUrl.Get()
	toSerialize["standalone"] = o.Standalone
	toSerialize["currentLicense"] = o.CurrentLicense
	toSerialize["max"] = o.Max
	return toSerialize, nil
}

func (o *PaymentSettingsDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"salesEmail",
		"buyUrl",
		"standalone",
		"currentLicense",
		"max",
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

	varPaymentSettingsDto := _PaymentSettingsDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varPaymentSettingsDto)

	if err != nil {
		return err
	}

	*o = PaymentSettingsDto(varPaymentSettingsDto)

	return err
}

type NullablePaymentSettingsDto struct {
	value *PaymentSettingsDto
	isSet bool
}

func (v NullablePaymentSettingsDto) Get() *PaymentSettingsDto {
	return v.value
}

func (v *NullablePaymentSettingsDto) Set(val *PaymentSettingsDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentSettingsDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentSettingsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentSettingsDto(val *PaymentSettingsDto) *NullablePaymentSettingsDto {
	return &NullablePaymentSettingsDto{value: val, isSet: true}
}

func (v NullablePaymentSettingsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentSettingsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

