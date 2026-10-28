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

// checks if the LoginEventDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoginEventDto{}

// LoginEventDto One entry of the portal login history: a sign-in, a sign-out or a failed attempt, and where it came from.
type LoginEventDto struct {
	// The ID of the recorded sign-in. When the entry is a successful sign-in that is still open, this is also  the value `GET api/2.0/security/activeconnections` reports as the connection's `id`.
	Id *int32 `json:"id,omitempty"`
	// When the attempt was made, in the portal time zone. The `from` and `to` filters are read as UTC instants,  so the two do not line up on a portal that is not on UTC.
	Date *ApiDateTime `json:"date,omitempty"`
	// The display name of the account the attempt was made against, taken from the account as it stands now  rather than as it stood at the time. A localised placeholder stands in when there is no account to read,  which is the usual case for a failed attempt on an address nobody owns.
	User NullableString `json:"user,omitempty"`
	// The ID of that account, which is what the `userId` filter of this operation matches on. It is the empty  GUID when the attempt could not be tied to an account.
	UserId *string `json:"userId,omitempty"`
	// The login string as it was typed - normally the email address. It is the only field that survives a failed  attempt against an unknown account, which makes it the one to read when `user` is a placeholder.
	Login NullableString `json:"login,omitempty"`
	// The event as a readable sentence in the portal language. On `GET api/2.0/security/audit/login/last` each  substituted value is cut to 50 characters; the filtered operation substitutes them in full.
	Action NullableString `json:"action,omitempty"`
	// What happened, as the `action` filter of this operation spells it: a successful sign-in, a failed one, a  sign-out. Use this rather than parsing `action`, which is prose and changes with the portal language.
	ActionId *MessageAction `json:"actionId,omitempty"`
	// The IP address the attempt came from, with the port stripped off.
	Ip NullableString `json:"ip,omitempty"`
	// The English name of the country the IP address is located in, empty when the address cannot be located -  the normal outcome for private and loopback addresses.
	Country NullableString `json:"country,omitempty"`
	// The city the IP address is located in, empty under the same conditions as `country`.
	City NullableString `json:"city,omitempty"`
	// The browser and its version as parsed from the user agent of the attempt, empty when the client sent none  that could be parsed.
	Browser NullableString `json:"browser,omitempty"`
	// The operating system as parsed from the same user agent, empty under the same conditions as `browser`.
	Platform NullableString `json:"platform,omitempty"`
	// Where in the portal the attempt was made from: the referrer of the request, or that request's own path  when it carried no referrer. Long values are cut off at 512 characters.
	Page NullableString `json:"page,omitempty"`
}

// NewLoginEventDto instantiates a new LoginEventDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginEventDto() *LoginEventDto {
	this := LoginEventDto{}
	return &this
}

// NewLoginEventDtoWithDefaults instantiates a new LoginEventDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginEventDtoWithDefaults() *LoginEventDto {
	this := LoginEventDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LoginEventDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginEventDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LoginEventDto) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *LoginEventDto) SetId(v int32) {
	o.Id = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *LoginEventDto) GetDate() ApiDateTime {
	if o == nil || IsNil(o.Date) {
		var ret ApiDateTime
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginEventDto) GetDateOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *LoginEventDto) IsDateSet() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given ApiDateTime and assigns it to the Date field.
func (o *LoginEventDto) SetDate(v ApiDateTime) {
	o.Date = &v
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoginEventDto) GetUser() string {
	if o == nil || IsNil(o.User.Get()) {
		var ret string
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoginEventDto) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *LoginEventDto) IsUserSet() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableString and assigns it to the User field.
func (o *LoginEventDto) SetUser(v string) {
	o.User.Set(&v)
}
// SetUserNil sets the value for User to be an explicit nil
func (o *LoginEventDto) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *LoginEventDto) UnsetUser() {
	o.User.Unset()
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *LoginEventDto) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginEventDto) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *LoginEventDto) IsUserIdSet() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *LoginEventDto) SetUserId(v string) {
	o.UserId = &v
}

// GetLogin returns the Login field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoginEventDto) GetLogin() string {
	if o == nil || IsNil(o.Login.Get()) {
		var ret string
		return ret
	}
	return *o.Login.Get()
}

// GetLoginOk returns a tuple with the Login field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoginEventDto) GetLoginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Login.Get(), o.Login.IsSet()
}

// HasLogin returns a boolean if a field has been set.
func (o *LoginEventDto) IsLoginSet() bool {
	if o != nil && o.Login.IsSet() {
		return true
	}

	return false
}

// SetLogin gets a reference to the given NullableString and assigns it to the Login field.
func (o *LoginEventDto) SetLogin(v string) {
	o.Login.Set(&v)
}
// SetLoginNil sets the value for Login to be an explicit nil
func (o *LoginEventDto) SetLoginNil() {
	o.Login.Set(nil)
}

// UnsetLogin ensures that no value is present for Login, not even an explicit nil
func (o *LoginEventDto) UnsetLogin() {
	o.Login.Unset()
}

// GetAction returns the Action field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoginEventDto) GetAction() string {
	if o == nil || IsNil(o.Action.Get()) {
		var ret string
		return ret
	}
	return *o.Action.Get()
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoginEventDto) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Action.Get(), o.Action.IsSet()
}

// HasAction returns a boolean if a field has been set.
func (o *LoginEventDto) IsActionSet() bool {
	if o != nil && o.Action.IsSet() {
		return true
	}

	return false
}

// SetAction gets a reference to the given NullableString and assigns it to the Action field.
func (o *LoginEventDto) SetAction(v string) {
	o.Action.Set(&v)
}
// SetActionNil sets the value for Action to be an explicit nil
func (o *LoginEventDto) SetActionNil() {
	o.Action.Set(nil)
}

// UnsetAction ensures that no value is present for Action, not even an explicit nil
func (o *LoginEventDto) UnsetAction() {
	o.Action.Unset()
}

// GetActionId returns the ActionId field value if set, zero value otherwise.
func (o *LoginEventDto) GetActionId() MessageAction {
	if o == nil || IsNil(o.ActionId) {
		var ret MessageAction
		return ret
	}
	return *o.ActionId
}

// GetActionIdOk returns a tuple with the ActionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginEventDto) GetActionIdOk() (*MessageAction, bool) {
	if o == nil || IsNil(o.ActionId) {
		return nil, false
	}
	return o.ActionId, true
}

// HasActionId returns a boolean if a field has been set.
func (o *LoginEventDto) IsActionIdSet() bool {
	if o != nil && !IsNil(o.ActionId) {
		return true
	}

	return false
}

// SetActionId gets a reference to the given MessageAction and assigns it to the ActionId field.
func (o *LoginEventDto) SetActionId(v MessageAction) {
	o.ActionId = &v
}

// GetIp returns the Ip field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoginEventDto) GetIp() string {
	if o == nil || IsNil(o.Ip.Get()) {
		var ret string
		return ret
	}
	return *o.Ip.Get()
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoginEventDto) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ip.Get(), o.Ip.IsSet()
}

// HasIp returns a boolean if a field has been set.
func (o *LoginEventDto) IsIpSet() bool {
	if o != nil && o.Ip.IsSet() {
		return true
	}

	return false
}

// SetIp gets a reference to the given NullableString and assigns it to the Ip field.
func (o *LoginEventDto) SetIp(v string) {
	o.Ip.Set(&v)
}
// SetIpNil sets the value for Ip to be an explicit nil
func (o *LoginEventDto) SetIpNil() {
	o.Ip.Set(nil)
}

// UnsetIp ensures that no value is present for Ip, not even an explicit nil
func (o *LoginEventDto) UnsetIp() {
	o.Ip.Unset()
}

// GetCountry returns the Country field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoginEventDto) GetCountry() string {
	if o == nil || IsNil(o.Country.Get()) {
		var ret string
		return ret
	}
	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoginEventDto) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// HasCountry returns a boolean if a field has been set.
func (o *LoginEventDto) IsCountrySet() bool {
	if o != nil && o.Country.IsSet() {
		return true
	}

	return false
}

// SetCountry gets a reference to the given NullableString and assigns it to the Country field.
func (o *LoginEventDto) SetCountry(v string) {
	o.Country.Set(&v)
}
// SetCountryNil sets the value for Country to be an explicit nil
func (o *LoginEventDto) SetCountryNil() {
	o.Country.Set(nil)
}

// UnsetCountry ensures that no value is present for Country, not even an explicit nil
func (o *LoginEventDto) UnsetCountry() {
	o.Country.Unset()
}

// GetCity returns the City field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoginEventDto) GetCity() string {
	if o == nil || IsNil(o.City.Get()) {
		var ret string
		return ret
	}
	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoginEventDto) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// HasCity returns a boolean if a field has been set.
func (o *LoginEventDto) IsCitySet() bool {
	if o != nil && o.City.IsSet() {
		return true
	}

	return false
}

// SetCity gets a reference to the given NullableString and assigns it to the City field.
func (o *LoginEventDto) SetCity(v string) {
	o.City.Set(&v)
}
// SetCityNil sets the value for City to be an explicit nil
func (o *LoginEventDto) SetCityNil() {
	o.City.Set(nil)
}

// UnsetCity ensures that no value is present for City, not even an explicit nil
func (o *LoginEventDto) UnsetCity() {
	o.City.Unset()
}

// GetBrowser returns the Browser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoginEventDto) GetBrowser() string {
	if o == nil || IsNil(o.Browser.Get()) {
		var ret string
		return ret
	}
	return *o.Browser.Get()
}

// GetBrowserOk returns a tuple with the Browser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoginEventDto) GetBrowserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Browser.Get(), o.Browser.IsSet()
}

// HasBrowser returns a boolean if a field has been set.
func (o *LoginEventDto) IsBrowserSet() bool {
	if o != nil && o.Browser.IsSet() {
		return true
	}

	return false
}

// SetBrowser gets a reference to the given NullableString and assigns it to the Browser field.
func (o *LoginEventDto) SetBrowser(v string) {
	o.Browser.Set(&v)
}
// SetBrowserNil sets the value for Browser to be an explicit nil
func (o *LoginEventDto) SetBrowserNil() {
	o.Browser.Set(nil)
}

// UnsetBrowser ensures that no value is present for Browser, not even an explicit nil
func (o *LoginEventDto) UnsetBrowser() {
	o.Browser.Unset()
}

// GetPlatform returns the Platform field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoginEventDto) GetPlatform() string {
	if o == nil || IsNil(o.Platform.Get()) {
		var ret string
		return ret
	}
	return *o.Platform.Get()
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoginEventDto) GetPlatformOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Platform.Get(), o.Platform.IsSet()
}

// HasPlatform returns a boolean if a field has been set.
func (o *LoginEventDto) IsPlatformSet() bool {
	if o != nil && o.Platform.IsSet() {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given NullableString and assigns it to the Platform field.
func (o *LoginEventDto) SetPlatform(v string) {
	o.Platform.Set(&v)
}
// SetPlatformNil sets the value for Platform to be an explicit nil
func (o *LoginEventDto) SetPlatformNil() {
	o.Platform.Set(nil)
}

// UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
func (o *LoginEventDto) UnsetPlatform() {
	o.Platform.Unset()
}

// GetPage returns the Page field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoginEventDto) GetPage() string {
	if o == nil || IsNil(o.Page.Get()) {
		var ret string
		return ret
	}
	return *o.Page.Get()
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoginEventDto) GetPageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Page.Get(), o.Page.IsSet()
}

// HasPage returns a boolean if a field has been set.
func (o *LoginEventDto) IsPageSet() bool {
	if o != nil && o.Page.IsSet() {
		return true
	}

	return false
}

// SetPage gets a reference to the given NullableString and assigns it to the Page field.
func (o *LoginEventDto) SetPage(v string) {
	o.Page.Set(&v)
}
// SetPageNil sets the value for Page to be an explicit nil
func (o *LoginEventDto) SetPageNil() {
	o.Page.Set(nil)
}

// UnsetPage ensures that no value is present for Page, not even an explicit nil
func (o *LoginEventDto) UnsetPage() {
	o.Page.Unset()
}

func (o LoginEventDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoginEventDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if o.Login.IsSet() {
		toSerialize["login"] = o.Login.Get()
	}
	if o.Action.IsSet() {
		toSerialize["action"] = o.Action.Get()
	}
	if !IsNil(o.ActionId) {
		toSerialize["actionId"] = o.ActionId
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
	if o.Page.IsSet() {
		toSerialize["page"] = o.Page.Get()
	}
	return toSerialize, nil
}

type NullableLoginEventDto struct {
	value *LoginEventDto
	isSet bool
}

func (v NullableLoginEventDto) Get() *LoginEventDto {
	return v.value
}

func (v *NullableLoginEventDto) Set(val *LoginEventDto) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginEventDto) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginEventDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginEventDto(val *LoginEventDto) *NullableLoginEventDto {
	return &NullableLoginEventDto{value: val, isSet: true}
}

func (v NullableLoginEventDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginEventDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

