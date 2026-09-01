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

// checks if the ActiveConnectionsItemDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActiveConnectionsItemDto{}

// ActiveConnectionsItemDto The active connection item parameters.
type ActiveConnectionsItemDto struct {
	// The active connection ID.
	Id int32 `json:"id"`
	// The tenant ID.
	TenantId int32 `json:"tenantId"`
	// The user ID.
	UserId string `json:"userId"`
	// Specifies if the active connection has a mobile phone or not.
	Mobile *bool `json:"mobile,omitempty"`
	// The IP address of the active connection.
	Ip NullableString `json:"ip,omitempty"`
	// The active connection country.
	Country NullableString `json:"country,omitempty"`
	// The active connection city.
	City NullableString `json:"city,omitempty"`
	// The active connection browser.
	Browser NullableString `json:"browser,omitempty"`
	// The active connection platform.
	Platform NullableString `json:"platform,omitempty"`
	// The active connection date.
	Date NullableTime `json:"date,omitempty"`
	// The active connection page.
	Page NullableString `json:"page,omitempty"`
}

type _ActiveConnectionsItemDto ActiveConnectionsItemDto

// NewActiveConnectionsItemDto instantiates a new ActiveConnectionsItemDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActiveConnectionsItemDto(id int32, tenantId int32, userId string) *ActiveConnectionsItemDto {
	this := ActiveConnectionsItemDto{}
	this.Id = id
	this.TenantId = tenantId
	this.UserId = userId
	return &this
}

// NewActiveConnectionsItemDtoWithDefaults instantiates a new ActiveConnectionsItemDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActiveConnectionsItemDtoWithDefaults() *ActiveConnectionsItemDto {
	this := ActiveConnectionsItemDto{}
	return &this
}

// GetId returns the Id field value
func (o *ActiveConnectionsItemDto) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ActiveConnectionsItemDto) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ActiveConnectionsItemDto) SetId(v int32) {
	o.Id = v
}

// GetTenantId returns the TenantId field value
func (o *ActiveConnectionsItemDto) GetTenantId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *ActiveConnectionsItemDto) GetTenantIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *ActiveConnectionsItemDto) SetTenantId(v int32) {
	o.TenantId = v
}

// GetUserId returns the UserId field value
func (o *ActiveConnectionsItemDto) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *ActiveConnectionsItemDto) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *ActiveConnectionsItemDto) SetUserId(v string) {
	o.UserId = v
}

// GetMobile returns the Mobile field value if set, zero value otherwise.
func (o *ActiveConnectionsItemDto) GetMobile() bool {
	if o == nil || IsNil(o.Mobile) {
		var ret bool
		return ret
	}
	return *o.Mobile
}

// GetMobileOk returns a tuple with the Mobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveConnectionsItemDto) GetMobileOk() (*bool, bool) {
	if o == nil || IsNil(o.Mobile) {
		return nil, false
	}
	return o.Mobile, true
}

// HasMobile returns a boolean if a field has been set.
func (o *ActiveConnectionsItemDto) IsMobileSet() bool {
	if o != nil && !IsNil(o.Mobile) {
		return true
	}

	return false
}

// SetMobile gets a reference to the given bool and assigns it to the Mobile field.
func (o *ActiveConnectionsItemDto) SetMobile(v bool) {
	o.Mobile = &v
}

// GetIp returns the Ip field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActiveConnectionsItemDto) GetIp() string {
	if o == nil || IsNil(o.Ip.Get()) {
		var ret string
		return ret
	}
	return *o.Ip.Get()
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveConnectionsItemDto) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ip.Get(), o.Ip.IsSet()
}

// HasIp returns a boolean if a field has been set.
func (o *ActiveConnectionsItemDto) IsIpSet() bool {
	if o != nil && o.Ip.IsSet() {
		return true
	}

	return false
}

// SetIp gets a reference to the given NullableString and assigns it to the Ip field.
func (o *ActiveConnectionsItemDto) SetIp(v string) {
	o.Ip.Set(&v)
}
// SetIpNil sets the value for Ip to be an explicit nil
func (o *ActiveConnectionsItemDto) SetIpNil() {
	o.Ip.Set(nil)
}

// UnsetIp ensures that no value is present for Ip, not even an explicit nil
func (o *ActiveConnectionsItemDto) UnsetIp() {
	o.Ip.Unset()
}

// GetCountry returns the Country field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActiveConnectionsItemDto) GetCountry() string {
	if o == nil || IsNil(o.Country.Get()) {
		var ret string
		return ret
	}
	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveConnectionsItemDto) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// HasCountry returns a boolean if a field has been set.
func (o *ActiveConnectionsItemDto) IsCountrySet() bool {
	if o != nil && o.Country.IsSet() {
		return true
	}

	return false
}

// SetCountry gets a reference to the given NullableString and assigns it to the Country field.
func (o *ActiveConnectionsItemDto) SetCountry(v string) {
	o.Country.Set(&v)
}
// SetCountryNil sets the value for Country to be an explicit nil
func (o *ActiveConnectionsItemDto) SetCountryNil() {
	o.Country.Set(nil)
}

// UnsetCountry ensures that no value is present for Country, not even an explicit nil
func (o *ActiveConnectionsItemDto) UnsetCountry() {
	o.Country.Unset()
}

// GetCity returns the City field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActiveConnectionsItemDto) GetCity() string {
	if o == nil || IsNil(o.City.Get()) {
		var ret string
		return ret
	}
	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveConnectionsItemDto) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// HasCity returns a boolean if a field has been set.
func (o *ActiveConnectionsItemDto) IsCitySet() bool {
	if o != nil && o.City.IsSet() {
		return true
	}

	return false
}

// SetCity gets a reference to the given NullableString and assigns it to the City field.
func (o *ActiveConnectionsItemDto) SetCity(v string) {
	o.City.Set(&v)
}
// SetCityNil sets the value for City to be an explicit nil
func (o *ActiveConnectionsItemDto) SetCityNil() {
	o.City.Set(nil)
}

// UnsetCity ensures that no value is present for City, not even an explicit nil
func (o *ActiveConnectionsItemDto) UnsetCity() {
	o.City.Unset()
}

// GetBrowser returns the Browser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActiveConnectionsItemDto) GetBrowser() string {
	if o == nil || IsNil(o.Browser.Get()) {
		var ret string
		return ret
	}
	return *o.Browser.Get()
}

// GetBrowserOk returns a tuple with the Browser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveConnectionsItemDto) GetBrowserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Browser.Get(), o.Browser.IsSet()
}

// HasBrowser returns a boolean if a field has been set.
func (o *ActiveConnectionsItemDto) IsBrowserSet() bool {
	if o != nil && o.Browser.IsSet() {
		return true
	}

	return false
}

// SetBrowser gets a reference to the given NullableString and assigns it to the Browser field.
func (o *ActiveConnectionsItemDto) SetBrowser(v string) {
	o.Browser.Set(&v)
}
// SetBrowserNil sets the value for Browser to be an explicit nil
func (o *ActiveConnectionsItemDto) SetBrowserNil() {
	o.Browser.Set(nil)
}

// UnsetBrowser ensures that no value is present for Browser, not even an explicit nil
func (o *ActiveConnectionsItemDto) UnsetBrowser() {
	o.Browser.Unset()
}

// GetPlatform returns the Platform field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActiveConnectionsItemDto) GetPlatform() string {
	if o == nil || IsNil(o.Platform.Get()) {
		var ret string
		return ret
	}
	return *o.Platform.Get()
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveConnectionsItemDto) GetPlatformOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Platform.Get(), o.Platform.IsSet()
}

// HasPlatform returns a boolean if a field has been set.
func (o *ActiveConnectionsItemDto) IsPlatformSet() bool {
	if o != nil && o.Platform.IsSet() {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given NullableString and assigns it to the Platform field.
func (o *ActiveConnectionsItemDto) SetPlatform(v string) {
	o.Platform.Set(&v)
}
// SetPlatformNil sets the value for Platform to be an explicit nil
func (o *ActiveConnectionsItemDto) SetPlatformNil() {
	o.Platform.Set(nil)
}

// UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
func (o *ActiveConnectionsItemDto) UnsetPlatform() {
	o.Platform.Unset()
}

// GetDate returns the Date field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActiveConnectionsItemDto) GetDate() time.Time {
	if o == nil || IsNil(o.Date.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Date.Get()
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveConnectionsItemDto) GetDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Date.Get(), o.Date.IsSet()
}

// HasDate returns a boolean if a field has been set.
func (o *ActiveConnectionsItemDto) IsDateSet() bool {
	if o != nil && o.Date.IsSet() {
		return true
	}

	return false
}

// SetDate gets a reference to the given NullableTime and assigns it to the Date field.
func (o *ActiveConnectionsItemDto) SetDate(v time.Time) {
	o.Date.Set(&v)
}
// SetDateNil sets the value for Date to be an explicit nil
func (o *ActiveConnectionsItemDto) SetDateNil() {
	o.Date.Set(nil)
}

// UnsetDate ensures that no value is present for Date, not even an explicit nil
func (o *ActiveConnectionsItemDto) UnsetDate() {
	o.Date.Unset()
}

// GetPage returns the Page field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActiveConnectionsItemDto) GetPage() string {
	if o == nil || IsNil(o.Page.Get()) {
		var ret string
		return ret
	}
	return *o.Page.Get()
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActiveConnectionsItemDto) GetPageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Page.Get(), o.Page.IsSet()
}

// HasPage returns a boolean if a field has been set.
func (o *ActiveConnectionsItemDto) IsPageSet() bool {
	if o != nil && o.Page.IsSet() {
		return true
	}

	return false
}

// SetPage gets a reference to the given NullableString and assigns it to the Page field.
func (o *ActiveConnectionsItemDto) SetPage(v string) {
	o.Page.Set(&v)
}
// SetPageNil sets the value for Page to be an explicit nil
func (o *ActiveConnectionsItemDto) SetPageNil() {
	o.Page.Set(nil)
}

// UnsetPage ensures that no value is present for Page, not even an explicit nil
func (o *ActiveConnectionsItemDto) UnsetPage() {
	o.Page.Unset()
}

func (o ActiveConnectionsItemDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActiveConnectionsItemDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["tenantId"] = o.TenantId
	toSerialize["userId"] = o.UserId
	if !IsNil(o.Mobile) {
		toSerialize["mobile"] = o.Mobile
	}
	if o.Ip.IsSet() {
		toSerialize["ip"] = o.Ip.Get()
	}
	if o.Country.IsSet() {
		toSerialize["country"] = o.Country.Get()
	}
	if o.City.IsSet() {
		toSerialize["city"] = o.City.Get()
	}
	if o.Browser.IsSet() {
		toSerialize["browser"] = o.Browser.Get()
	}
	if o.Platform.IsSet() {
		toSerialize["platform"] = o.Platform.Get()
	}
	if o.Date.IsSet() {
		toSerialize["date"] = o.Date.Get()
	}
	if o.Page.IsSet() {
		toSerialize["page"] = o.Page.Get()
	}
	return toSerialize, nil
}

func (o *ActiveConnectionsItemDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"tenantId",
		"userId",
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

	varActiveConnectionsItemDto := _ActiveConnectionsItemDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varActiveConnectionsItemDto)

	if err != nil {
		return err
	}

	*o = ActiveConnectionsItemDto(varActiveConnectionsItemDto)

	return err
}

type NullableActiveConnectionsItemDto struct {
	value *ActiveConnectionsItemDto
	isSet bool
}

func (v NullableActiveConnectionsItemDto) Get() *ActiveConnectionsItemDto {
	return v.value
}

func (v *NullableActiveConnectionsItemDto) Set(val *ActiveConnectionsItemDto) {
	v.value = val
	v.isSet = true
}

func (v NullableActiveConnectionsItemDto) IsSet() bool {
	return v.isSet
}

func (v *NullableActiveConnectionsItemDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActiveConnectionsItemDto(val *ActiveConnectionsItemDto) *NullableActiveConnectionsItemDto {
	return &NullableActiveConnectionsItemDto{value: val, isSet: true}
}

func (v NullableActiveConnectionsItemDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActiveConnectionsItemDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

