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

// checks if the AuditEventDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuditEventDto{}

// AuditEventDto One entry of the portal audit trail: who changed what, from where, and where it belongs in the product.
type AuditEventDto struct {
	// The ID of the recorded entry. Nothing accepts it as an argument - no operation fetches a single audit event  - so it serves only to tell two otherwise identical entries apart.
	Id *int32 `json:"id,omitempty"`
	// When the action happened, in the portal time zone. The `from` and `to` filters are read as UTC instants, so  the two do not line up on a portal that is not on UTC.
	Date *ApiDateTime `json:"date,omitempty"`
	// The display name of the user who acted, taken from the account as it stands now rather than as it stood  when the entry was written. A localised placeholder stands in when there is no account to read: a portal  background job, an anonymous guest, or a user who has since been deleted.
	User NullableString `json:"user,omitempty"`
	// The ID of the user who acted, which is what the `userId` filter of this operation matches on. It stays  readable after the account is deleted, which is when `user` falls back to a placeholder.
	UserId *string `json:"userId,omitempty"`
	// The whole event as a readable sentence in the portal language, with the names of the objects involved  substituted into it. On the two `audit/.../last` operations each substituted value is cut to 50 characters;  the filtered operations substitute them in full. It is empty when the build has no wording for the action.
	Action NullableString `json:"action,omitempty"`
	// The action itself, as the `action` filter of this operation spells it and as  `GET api/2.0/security/audit/mappers` lists it under `messageAction`. Use this rather than parsing `action`,  which is prose and changes with the portal language.
	ActionId *MessageAction `json:"actionId,omitempty"`
	// The IP address the request came from, with the port stripped off. It is empty for an action a portal  background job performed, which has no request behind it.
	Ip NullableString `json:"ip,omitempty"`
	// The English name of the country the IP address is located in, empty when the address cannot be located -  the normal outcome for private and loopback addresses.
	Country NullableString `json:"country,omitempty"`
	// The city the IP address is located in, empty under the same conditions as `country`.
	City NullableString `json:"city,omitempty"`
	// The browser and its version as parsed from the user agent of the request, empty when the client sent none  that could be parsed or when no request was involved.
	Browser NullableString `json:"browser,omitempty"`
	// The operating system as parsed from the same user agent, empty under the same conditions as `browser`.
	Platform NullableString `json:"platform,omitempty"`
	// Where in the portal the action was made from: the referrer of the request, or that request's own path when  it carried no referrer. Long values are cut off at 512 characters.
	Page NullableString `json:"page,omitempty"`
	// The kind of change the action stands for, as the `actionType` filter of this operation spells it. It is  derived from `actionId`, not stored per entry, so it is the same on every entry of one action.
	ActionType *ActionType `json:"actionType,omitempty"`
	// The product the action belongs to. It cannot be filtered on here; the tree that groups actions by product  is `GET api/2.0/security/audit/mappers`.
	Product *ProductType `json:"product,omitempty"`
	// The location inside that product, as the `moduleType` filter of this operation spells it. It is also  derived from `actionId` rather than stored per entry.
	Location *LocationType `json:"location,omitempty"`
	// The objects the action was applied to, as the trail recorded them - a title, an account, an ID - one string  each. It is empty for an action that targets nothing, such as a settings change, and the `target` filter of  this operation matches one of these values in full.
	Target []string `json:"target,omitempty"`
	// The kinds of object the action applies to, holding at most two entries and none at all for an action that  targets nothing. Only the first of them can be filtered on, through `entryType`.
	Entries []EntryType `json:"entries,omitempty"`
	// Where the action took place, spelled out in the portal language rather than as a code: for a Documents  event the room or the root folder it happened in, and for anything else the name of the module. Nothing  filters on it.
	Context NullableString `json:"context,omitempty"`
}

// NewAuditEventDto instantiates a new AuditEventDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuditEventDto() *AuditEventDto {
	this := AuditEventDto{}
	return &this
}

// NewAuditEventDtoWithDefaults instantiates a new AuditEventDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuditEventDtoWithDefaults() *AuditEventDto {
	this := AuditEventDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuditEventDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditEventDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuditEventDto) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AuditEventDto) SetId(v int32) {
	o.Id = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *AuditEventDto) GetDate() ApiDateTime {
	if o == nil || IsNil(o.Date) {
		var ret ApiDateTime
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditEventDto) GetDateOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *AuditEventDto) IsDateSet() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given ApiDateTime and assigns it to the Date field.
func (o *AuditEventDto) SetDate(v ApiDateTime) {
	o.Date = &v
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditEventDto) GetUser() string {
	if o == nil || IsNil(o.User.Get()) {
		var ret string
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditEventDto) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *AuditEventDto) IsUserSet() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableString and assigns it to the User field.
func (o *AuditEventDto) SetUser(v string) {
	o.User.Set(&v)
}
// SetUserNil sets the value for User to be an explicit nil
func (o *AuditEventDto) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *AuditEventDto) UnsetUser() {
	o.User.Unset()
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *AuditEventDto) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditEventDto) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *AuditEventDto) IsUserIdSet() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *AuditEventDto) SetUserId(v string) {
	o.UserId = &v
}

// GetAction returns the Action field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditEventDto) GetAction() string {
	if o == nil || IsNil(o.Action.Get()) {
		var ret string
		return ret
	}
	return *o.Action.Get()
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditEventDto) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Action.Get(), o.Action.IsSet()
}

// HasAction returns a boolean if a field has been set.
func (o *AuditEventDto) IsActionSet() bool {
	if o != nil && o.Action.IsSet() {
		return true
	}

	return false
}

// SetAction gets a reference to the given NullableString and assigns it to the Action field.
func (o *AuditEventDto) SetAction(v string) {
	o.Action.Set(&v)
}
// SetActionNil sets the value for Action to be an explicit nil
func (o *AuditEventDto) SetActionNil() {
	o.Action.Set(nil)
}

// UnsetAction ensures that no value is present for Action, not even an explicit nil
func (o *AuditEventDto) UnsetAction() {
	o.Action.Unset()
}

// GetActionId returns the ActionId field value if set, zero value otherwise.
func (o *AuditEventDto) GetActionId() MessageAction {
	if o == nil || IsNil(o.ActionId) {
		var ret MessageAction
		return ret
	}
	return *o.ActionId
}

// GetActionIdOk returns a tuple with the ActionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditEventDto) GetActionIdOk() (*MessageAction, bool) {
	if o == nil || IsNil(o.ActionId) {
		return nil, false
	}
	return o.ActionId, true
}

// HasActionId returns a boolean if a field has been set.
func (o *AuditEventDto) IsActionIdSet() bool {
	if o != nil && !IsNil(o.ActionId) {
		return true
	}

	return false
}

// SetActionId gets a reference to the given MessageAction and assigns it to the ActionId field.
func (o *AuditEventDto) SetActionId(v MessageAction) {
	o.ActionId = &v
}

// GetIp returns the Ip field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditEventDto) GetIp() string {
	if o == nil || IsNil(o.Ip.Get()) {
		var ret string
		return ret
	}
	return *o.Ip.Get()
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditEventDto) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ip.Get(), o.Ip.IsSet()
}

// HasIp returns a boolean if a field has been set.
func (o *AuditEventDto) IsIpSet() bool {
	if o != nil && o.Ip.IsSet() {
		return true
	}

	return false
}

// SetIp gets a reference to the given NullableString and assigns it to the Ip field.
func (o *AuditEventDto) SetIp(v string) {
	o.Ip.Set(&v)
}
// SetIpNil sets the value for Ip to be an explicit nil
func (o *AuditEventDto) SetIpNil() {
	o.Ip.Set(nil)
}

// UnsetIp ensures that no value is present for Ip, not even an explicit nil
func (o *AuditEventDto) UnsetIp() {
	o.Ip.Unset()
}

// GetCountry returns the Country field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditEventDto) GetCountry() string {
	if o == nil || IsNil(o.Country.Get()) {
		var ret string
		return ret
	}
	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditEventDto) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// HasCountry returns a boolean if a field has been set.
func (o *AuditEventDto) IsCountrySet() bool {
	if o != nil && o.Country.IsSet() {
		return true
	}

	return false
}

// SetCountry gets a reference to the given NullableString and assigns it to the Country field.
func (o *AuditEventDto) SetCountry(v string) {
	o.Country.Set(&v)
}
// SetCountryNil sets the value for Country to be an explicit nil
func (o *AuditEventDto) SetCountryNil() {
	o.Country.Set(nil)
}

// UnsetCountry ensures that no value is present for Country, not even an explicit nil
func (o *AuditEventDto) UnsetCountry() {
	o.Country.Unset()
}

// GetCity returns the City field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditEventDto) GetCity() string {
	if o == nil || IsNil(o.City.Get()) {
		var ret string
		return ret
	}
	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditEventDto) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// HasCity returns a boolean if a field has been set.
func (o *AuditEventDto) IsCitySet() bool {
	if o != nil && o.City.IsSet() {
		return true
	}

	return false
}

// SetCity gets a reference to the given NullableString and assigns it to the City field.
func (o *AuditEventDto) SetCity(v string) {
	o.City.Set(&v)
}
// SetCityNil sets the value for City to be an explicit nil
func (o *AuditEventDto) SetCityNil() {
	o.City.Set(nil)
}

// UnsetCity ensures that no value is present for City, not even an explicit nil
func (o *AuditEventDto) UnsetCity() {
	o.City.Unset()
}

// GetBrowser returns the Browser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditEventDto) GetBrowser() string {
	if o == nil || IsNil(o.Browser.Get()) {
		var ret string
		return ret
	}
	return *o.Browser.Get()
}

// GetBrowserOk returns a tuple with the Browser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditEventDto) GetBrowserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Browser.Get(), o.Browser.IsSet()
}

// HasBrowser returns a boolean if a field has been set.
func (o *AuditEventDto) IsBrowserSet() bool {
	if o != nil && o.Browser.IsSet() {
		return true
	}

	return false
}

// SetBrowser gets a reference to the given NullableString and assigns it to the Browser field.
func (o *AuditEventDto) SetBrowser(v string) {
	o.Browser.Set(&v)
}
// SetBrowserNil sets the value for Browser to be an explicit nil
func (o *AuditEventDto) SetBrowserNil() {
	o.Browser.Set(nil)
}

// UnsetBrowser ensures that no value is present for Browser, not even an explicit nil
func (o *AuditEventDto) UnsetBrowser() {
	o.Browser.Unset()
}

// GetPlatform returns the Platform field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditEventDto) GetPlatform() string {
	if o == nil || IsNil(o.Platform.Get()) {
		var ret string
		return ret
	}
	return *o.Platform.Get()
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditEventDto) GetPlatformOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Platform.Get(), o.Platform.IsSet()
}

// HasPlatform returns a boolean if a field has been set.
func (o *AuditEventDto) IsPlatformSet() bool {
	if o != nil && o.Platform.IsSet() {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given NullableString and assigns it to the Platform field.
func (o *AuditEventDto) SetPlatform(v string) {
	o.Platform.Set(&v)
}
// SetPlatformNil sets the value for Platform to be an explicit nil
func (o *AuditEventDto) SetPlatformNil() {
	o.Platform.Set(nil)
}

// UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
func (o *AuditEventDto) UnsetPlatform() {
	o.Platform.Unset()
}

// GetPage returns the Page field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditEventDto) GetPage() string {
	if o == nil || IsNil(o.Page.Get()) {
		var ret string
		return ret
	}
	return *o.Page.Get()
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditEventDto) GetPageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Page.Get(), o.Page.IsSet()
}

// HasPage returns a boolean if a field has been set.
func (o *AuditEventDto) IsPageSet() bool {
	if o != nil && o.Page.IsSet() {
		return true
	}

	return false
}

// SetPage gets a reference to the given NullableString and assigns it to the Page field.
func (o *AuditEventDto) SetPage(v string) {
	o.Page.Set(&v)
}
// SetPageNil sets the value for Page to be an explicit nil
func (o *AuditEventDto) SetPageNil() {
	o.Page.Set(nil)
}

// UnsetPage ensures that no value is present for Page, not even an explicit nil
func (o *AuditEventDto) UnsetPage() {
	o.Page.Unset()
}

// GetActionType returns the ActionType field value if set, zero value otherwise.
func (o *AuditEventDto) GetActionType() ActionType {
	if o == nil || IsNil(o.ActionType) {
		var ret ActionType
		return ret
	}
	return *o.ActionType
}

// GetActionTypeOk returns a tuple with the ActionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditEventDto) GetActionTypeOk() (*ActionType, bool) {
	if o == nil || IsNil(o.ActionType) {
		return nil, false
	}
	return o.ActionType, true
}

// HasActionType returns a boolean if a field has been set.
func (o *AuditEventDto) IsActionTypeSet() bool {
	if o != nil && !IsNil(o.ActionType) {
		return true
	}

	return false
}

// SetActionType gets a reference to the given ActionType and assigns it to the ActionType field.
func (o *AuditEventDto) SetActionType(v ActionType) {
	o.ActionType = &v
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *AuditEventDto) GetProduct() ProductType {
	if o == nil || IsNil(o.Product) {
		var ret ProductType
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditEventDto) GetProductOk() (*ProductType, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *AuditEventDto) IsProductSet() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given ProductType and assigns it to the Product field.
func (o *AuditEventDto) SetProduct(v ProductType) {
	o.Product = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *AuditEventDto) GetLocation() LocationType {
	if o == nil || IsNil(o.Location) {
		var ret LocationType
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuditEventDto) GetLocationOk() (*LocationType, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *AuditEventDto) IsLocationSet() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given LocationType and assigns it to the Location field.
func (o *AuditEventDto) SetLocation(v LocationType) {
	o.Location = &v
}

// GetTarget returns the Target field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditEventDto) GetTarget() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditEventDto) GetTargetOk() ([]string, bool) {
	if o == nil || IsNil(o.Target) {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *AuditEventDto) IsTargetSet() bool {
	if o != nil && !IsNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given []string and assigns it to the Target field.
func (o *AuditEventDto) SetTarget(v []string) {
	o.Target = v
}

// GetEntries returns the Entries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditEventDto) GetEntries() []EntryType {
	if o == nil {
		var ret []EntryType
		return ret
	}
	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditEventDto) GetEntriesOk() ([]EntryType, bool) {
	if o == nil || IsNil(o.Entries) {
		return nil, false
	}
	return o.Entries, true
}

// HasEntries returns a boolean if a field has been set.
func (o *AuditEventDto) IsEntriesSet() bool {
	if o != nil && !IsNil(o.Entries) {
		return true
	}

	return false
}

// SetEntries gets a reference to the given []EntryType and assigns it to the Entries field.
func (o *AuditEventDto) SetEntries(v []EntryType) {
	o.Entries = v
}

// GetContext returns the Context field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuditEventDto) GetContext() string {
	if o == nil || IsNil(o.Context.Get()) {
		var ret string
		return ret
	}
	return *o.Context.Get()
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuditEventDto) GetContextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Context.Get(), o.Context.IsSet()
}

// HasContext returns a boolean if a field has been set.
func (o *AuditEventDto) IsContextSet() bool {
	if o != nil && o.Context.IsSet() {
		return true
	}

	return false
}

// SetContext gets a reference to the given NullableString and assigns it to the Context field.
func (o *AuditEventDto) SetContext(v string) {
	o.Context.Set(&v)
}
// SetContextNil sets the value for Context to be an explicit nil
func (o *AuditEventDto) SetContextNil() {
	o.Context.Set(nil)
}

// UnsetContext ensures that no value is present for Context, not even an explicit nil
func (o *AuditEventDto) UnsetContext() {
	o.Context.Unset()
}

func (o AuditEventDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuditEventDto) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ActionType) {
		toSerialize["actionType"] = o.ActionType
	}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	if o.Entries != nil {
		toSerialize["entries"] = o.Entries
	}
	if o.Context.IsSet() {
		toSerialize["context"] = o.Context.Get()
	}
	return toSerialize, nil
}

type NullableAuditEventDto struct {
	value *AuditEventDto
	isSet bool
}

func (v NullableAuditEventDto) Get() *AuditEventDto {
	return v.value
}

func (v *NullableAuditEventDto) Set(val *AuditEventDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAuditEventDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAuditEventDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuditEventDto(val *AuditEventDto) *NullableAuditEventDto {
	return &NullableAuditEventDto{value: val, isSet: true}
}

func (v NullableAuditEventDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuditEventDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

