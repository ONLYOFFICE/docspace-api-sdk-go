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

// checks if the EmployeeFullDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmployeeFullDto{}

// EmployeeFullDto The full list of user parameters.
type EmployeeFullDto struct {
	// The user ID.
	Id *string `json:"id,omitempty"`
	// The HTML-encoded user's display name formatted according to the default format for the current culture.
	DisplayName NullableString `json:"displayName,omitempty"`
	// The user title.
	Title NullableString `json:"title,omitempty"`
	// The user avatar.
	Avatar NullableString `json:"avatar,omitempty"`
	// The user original size avatar.
	AvatarOriginal NullableString `json:"avatarOriginal,omitempty"`
	// The user maximum size avatar.
	AvatarMax NullableString `json:"avatarMax,omitempty"`
	// The user medium size avatar.
	AvatarMedium NullableString `json:"avatarMedium,omitempty"`
	// The user small size avatar.
	AvatarSmall NullableString `json:"avatarSmall,omitempty"`
	// The user profile URL.
	ProfileUrl NullableString `json:"profileUrl,omitempty"`
	// Specifies if the user has an avatar or not.
	HasAvatar *bool `json:"hasAvatar,omitempty"`
	// Specifies if the user is anonymous or not.
	IsAnonim *bool `json:"isAnonim,omitempty"`
	// The user first name.
	FirstName NullableString `json:"firstName,omitempty"`
	// The user last name.
	LastName NullableString `json:"lastName,omitempty"`
	// The user username.
	UserName NullableString `json:"userName,omitempty"`
	// The user email.
	Email NullableString `json:"email,omitempty"`
	// The list of user contacts.
	Contacts []Contact `json:"contacts,omitempty"`
	Birthday *ApiDateTime `json:"birthday,omitempty"`
	// The user sex.
	Sex NullableString `json:"sex,omitempty"`
	Status *EmployeeStatus `json:"status,omitempty"`
	ActivationStatus *EmployeeActivationStatus `json:"activationStatus,omitempty"`
	Terminated *ApiDateTime `json:"terminated,omitempty"`
	// The user department.
	Department NullableString `json:"department,omitempty"`
	WorkFrom *ApiDateTime `json:"workFrom,omitempty"`
	// The list of user groups.
	Groups []GroupSummaryDto `json:"groups,omitempty"`
	// The user location.
	Location NullableString `json:"location,omitempty"`
	// The user notes.
	Notes NullableString `json:"notes,omitempty"`
	// Specifies if the user is an administrator or not.
	IsAdmin *bool `json:"isAdmin,omitempty"`
	// Specifies if the user is a room administrator or not.
	IsRoomAdmin *bool `json:"isRoomAdmin,omitempty"`
	// Specifies if the LDAP settings are enabled for the user or not.
	IsLDAP *bool `json:"isLDAP,omitempty"`
	// The list of the administrator modules.
	ListAdminModules []string `json:"listAdminModules,omitempty"`
	// Specifies if the user is a portal owner or not.
	IsOwner *bool `json:"isOwner,omitempty"`
	// Specifies if the user is a portal visitor or not.
	IsVisitor *bool `json:"isVisitor,omitempty"`
	// Specifies if the user is a portal collaborator or not.
	IsCollaborator *bool `json:"isCollaborator,omitempty"`
	// The user culture code.
	CultureName NullableString `json:"cultureName,omitempty"`
	// The user mobile phone number.
	MobilePhone NullableString `json:"mobilePhone,omitempty"`
	MobilePhoneActivationStatus *MobilePhoneActivationStatus `json:"mobilePhoneActivationStatus,omitempty"`
	// Specifies if the SSO settings are enabled for the user or not.
	IsSSO *bool `json:"isSSO,omitempty"`
	Theme *DarkThemeSettingsType `json:"theme,omitempty"`
	// The user quota limit.
	QuotaLimit NullableInt64 `json:"quotaLimit,omitempty"`
	// The portal used space of the user.
	UsedSpace NullableFloat64 `json:"usedSpace,omitempty"`
	// Specifies if the user has access rights.
	Shared NullableBool `json:"shared,omitempty"`
	// Specifies if the user has a custom quota or not.
	IsCustomQuota NullableBool `json:"isCustomQuota,omitempty"`
	// The current login event ID.
	LoginEventId NullableInt32 `json:"loginEventId,omitempty"`
	// The auth cookie lifetime in seconds.
	AuthCookieLifetime NullableFloat64 `json:"authCookieLifetime,omitempty"`
	CreatedBy *EmployeeDto `json:"createdBy,omitempty"`
	RegistrationDate *ApiDateTime `json:"registrationDate,omitempty"`
	// Specifies if the user has a personal folder or not.
	HasPersonalFolder NullableBool `json:"hasPersonalFolder,omitempty"`
	// Indicates whether the user has enabled two-factor authentication (TFA) using an authentication app.
	TfaAppEnabled NullableBool `json:"tfaAppEnabled,omitempty"`
}

// NewEmployeeFullDto instantiates a new EmployeeFullDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmployeeFullDto() *EmployeeFullDto {
	this := EmployeeFullDto{}
	return &this
}

// NewEmployeeFullDtoWithDefaults instantiates a new EmployeeFullDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmployeeFullDtoWithDefaults() *EmployeeFullDto {
	this := EmployeeFullDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EmployeeFullDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployeeFullDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EmployeeFullDto) SetId(v string) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployeeFullDto) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName.Get()) {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployeeFullDto) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsDisplayNameSet() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *EmployeeFullDto) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *EmployeeFullDto) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *EmployeeFullDto) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployeeFullDto) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployeeFullDto) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsTitleSet() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *EmployeeFullDto) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *EmployeeFullDto) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *EmployeeFullDto) UnsetTitle() {
	o.Title.Unset()
}

// GetAvatar returns the Avatar field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployeeFullDto) GetAvatar() string {
	if o == nil || IsNil(o.Avatar.Get()) {
		var ret string
		return ret
	}
	return *o.Avatar.Get()
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployeeFullDto) GetAvatarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Avatar.Get(), o.Avatar.IsSet()
}

// HasAvatar returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsAvatarSet() bool {
	if o != nil && o.Avatar.IsSet() {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given NullableString and assigns it to the Avatar field.
func (o *EmployeeFullDto) SetAvatar(v string) {
	o.Avatar.Set(&v)
}
// SetAvatarNil sets the value for Avatar to be an explicit nil
func (o *EmployeeFullDto) SetAvatarNil() {
	o.Avatar.Set(nil)
}

// UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
func (o *EmployeeFullDto) UnsetAvatar() {
	o.Avatar.Unset()
}

// GetAvatarOriginal returns the AvatarOriginal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployeeFullDto) GetAvatarOriginal() string {
	if o == nil || IsNil(o.AvatarOriginal.Get()) {
		var ret string
		return ret
	}
	return *o.AvatarOriginal.Get()
}

// GetAvatarOriginalOk returns a tuple with the AvatarOriginal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployeeFullDto) GetAvatarOriginalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvatarOriginal.Get(), o.AvatarOriginal.IsSet()
}

// HasAvatarOriginal returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsAvatarOriginalSet() bool {
	if o != nil && o.AvatarOriginal.IsSet() {
		return true
	}

	return false
}

// SetAvatarOriginal gets a reference to the given NullableString and assigns it to the AvatarOriginal field.
func (o *EmployeeFullDto) SetAvatarOriginal(v string) {
	o.AvatarOriginal.Set(&v)
}
// SetAvatarOriginalNil sets the value for AvatarOriginal to be an explicit nil
func (o *EmployeeFullDto) SetAvatarOriginalNil() {
	o.AvatarOriginal.Set(nil)
}

// UnsetAvatarOriginal ensures that no value is present for AvatarOriginal, not even an explicit nil
func (o *EmployeeFullDto) UnsetAvatarOriginal() {
	o.AvatarOriginal.Unset()
}

// GetAvatarMax returns the AvatarMax field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployeeFullDto) GetAvatarMax() string {
	if o == nil || IsNil(o.AvatarMax.Get()) {
		var ret string
		return ret
	}
	return *o.AvatarMax.Get()
}

// GetAvatarMaxOk returns a tuple with the AvatarMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployeeFullDto) GetAvatarMaxOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvatarMax.Get(), o.AvatarMax.IsSet()
}

// HasAvatarMax returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsAvatarMaxSet() bool {
	if o != nil && o.AvatarMax.IsSet() {
		return true
	}

	return false
}

// SetAvatarMax gets a reference to the given NullableString and assigns it to the AvatarMax field.
func (o *EmployeeFullDto) SetAvatarMax(v string) {
	o.AvatarMax.Set(&v)
}
// SetAvatarMaxNil sets the value for AvatarMax to be an explicit nil
func (o *EmployeeFullDto) SetAvatarMaxNil() {
	o.AvatarMax.Set(nil)
}

// UnsetAvatarMax ensures that no value is present for AvatarMax, not even an explicit nil
func (o *EmployeeFullDto) UnsetAvatarMax() {
	o.AvatarMax.Unset()
}

// GetAvatarMedium returns the AvatarMedium field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployeeFullDto) GetAvatarMedium() string {
	if o == nil || IsNil(o.AvatarMedium.Get()) {
		var ret string
		return ret
	}
	return *o.AvatarMedium.Get()
}

// GetAvatarMediumOk returns a tuple with the AvatarMedium field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployeeFullDto) GetAvatarMediumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvatarMedium.Get(), o.AvatarMedium.IsSet()
}

// HasAvatarMedium returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsAvatarMediumSet() bool {
	if o != nil && o.AvatarMedium.IsSet() {
		return true
	}

	return false
}

// SetAvatarMedium gets a reference to the given NullableString and assigns it to the AvatarMedium field.
func (o *EmployeeFullDto) SetAvatarMedium(v string) {
	o.AvatarMedium.Set(&v)
}
// SetAvatarMediumNil sets the value for AvatarMedium to be an explicit nil
func (o *EmployeeFullDto) SetAvatarMediumNil() {
	o.AvatarMedium.Set(nil)
}

// UnsetAvatarMedium ensures that no value is present for AvatarMedium, not even an explicit nil
func (o *EmployeeFullDto) UnsetAvatarMedium() {
	o.AvatarMedium.Unset()
}

// GetAvatarSmall returns the AvatarSmall field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployeeFullDto) GetAvatarSmall() string {
	if o == nil || IsNil(o.AvatarSmall.Get()) {
		var ret string
		return ret
	}
	return *o.AvatarSmall.Get()
}

// GetAvatarSmallOk returns a tuple with the AvatarSmall field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployeeFullDto) GetAvatarSmallOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvatarSmall.Get(), o.AvatarSmall.IsSet()
}

// HasAvatarSmall returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsAvatarSmallSet() bool {
	if o != nil && o.AvatarSmall.IsSet() {
		return true
	}

	return false
}

// SetAvatarSmall gets a reference to the given NullableString and assigns it to the AvatarSmall field.
func (o *EmployeeFullDto) SetAvatarSmall(v string) {
	o.AvatarSmall.Set(&v)
}
// SetAvatarSmallNil sets the value for AvatarSmall to be an explicit nil
func (o *EmployeeFullDto) SetAvatarSmallNil() {
	o.AvatarSmall.Set(nil)
}

// UnsetAvatarSmall ensures that no value is present for AvatarSmall, not even an explicit nil
func (o *EmployeeFullDto) UnsetAvatarSmall() {
	o.AvatarSmall.Unset()
}

// GetProfileUrl returns the ProfileUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployeeFullDto) GetProfileUrl() string {
	if o == nil || IsNil(o.ProfileUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ProfileUrl.Get()
}

// GetProfileUrlOk returns a tuple with the ProfileUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployeeFullDto) GetProfileUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfileUrl.Get(), o.ProfileUrl.IsSet()
}

// HasProfileUrl returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsProfileUrlSet() bool {
	if o != nil && o.ProfileUrl.IsSet() {
		return true
	}

	return false
}

// SetProfileUrl gets a reference to the given NullableString and assigns it to the ProfileUrl field.
func (o *EmployeeFullDto) SetProfileUrl(v string) {
	o.ProfileUrl.Set(&v)
}
// SetProfileUrlNil sets the value for ProfileUrl to be an explicit nil
func (o *EmployeeFullDto) SetProfileUrlNil() {
	o.ProfileUrl.Set(nil)
}

// UnsetProfileUrl ensures that no value is present for ProfileUrl, not even an explicit nil
func (o *EmployeeFullDto) UnsetProfileUrl() {
	o.ProfileUrl.Unset()
}

// GetHasAvatar returns the HasAvatar field value if set, zero value otherwise.
func (o *EmployeeFullDto) GetHasAvatar() bool {
	if o == nil || IsNil(o.HasAvatar) {
		var ret bool
		return ret
	}
	return *o.HasAvatar
}

// GetHasAvatarOk returns a tuple with the HasAvatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployeeFullDto) GetHasAvatarOk() (*bool, bool) {
	if o == nil || IsNil(o.HasAvatar) {
		return nil, false
	}
	return o.HasAvatar, true
}

// HasHasAvatar returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsHasAvatarSet() bool {
	if o != nil && !IsNil(o.HasAvatar) {
		return true
	}

	return false
}

// SetHasAvatar gets a reference to the given bool and assigns it to the HasAvatar field.
func (o *EmployeeFullDto) SetHasAvatar(v bool) {
	o.HasAvatar = &v
}

// GetIsAnonim returns the IsAnonim field value if set, zero value otherwise.
func (o *EmployeeFullDto) GetIsAnonim() bool {
	if o == nil || IsNil(o.IsAnonim) {
		var ret bool
		return ret
	}
	return *o.IsAnonim
}

// GetIsAnonimOk returns a tuple with the IsAnonim field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployeeFullDto) GetIsAnonimOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAnonim) {
		return nil, false
	}
	return o.IsAnonim, true
}

// HasIsAnonim returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsIsAnonimSet() bool {
	if o != nil && !IsNil(o.IsAnonim) {
		return true
	}

	return false
}

// SetIsAnonim gets a reference to the given bool and assigns it to the IsAnonim field.
func (o *EmployeeFullDto) SetIsAnonim(v bool) {
	o.IsAnonim = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployeeFullDto) GetFirstName() string {
	if o == nil || IsNil(o.FirstName.Get()) {
		var ret string
		return ret
	}
	return *o.FirstName.Get()
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployeeFullDto) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstName.Get(), o.FirstName.IsSet()
}

// HasFirstName returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsFirstNameSet() bool {
	if o != nil && o.FirstName.IsSet() {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given NullableString and assigns it to the FirstName field.
func (o *EmployeeFullDto) SetFirstName(v string) {
	o.FirstName.Set(&v)
}
// SetFirstNameNil sets the value for FirstName to be an explicit nil
func (o *EmployeeFullDto) SetFirstNameNil() {
	o.FirstName.Set(nil)
}

// UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
func (o *EmployeeFullDto) UnsetFirstName() {
	o.FirstName.Unset()
}

// GetLastName returns the LastName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployeeFullDto) GetLastName() string {
	if o == nil || IsNil(o.LastName.Get()) {
		var ret string
		return ret
	}
	return *o.LastName.Get()
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployeeFullDto) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastName.Get(), o.LastName.IsSet()
}

// HasLastName returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsLastNameSet() bool {
	if o != nil && o.LastName.IsSet() {
		return true
	}

	return false
}

// SetLastName gets a reference to the given NullableString and assigns it to the LastName field.
func (o *EmployeeFullDto) SetLastName(v string) {
	o.LastName.Set(&v)
}
// SetLastNameNil sets the value for LastName to be an explicit nil
func (o *EmployeeFullDto) SetLastNameNil() {
	o.LastName.Set(nil)
}

// UnsetLastName ensures that no value is present for LastName, not even an explicit nil
func (o *EmployeeFullDto) UnsetLastName() {
	o.LastName.Unset()
}

// GetUserName returns the UserName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployeeFullDto) GetUserName() string {
	if o == nil || IsNil(o.UserName.Get()) {
		var ret string
		return ret
	}
	return *o.UserName.Get()
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployeeFullDto) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserName.Get(), o.UserName.IsSet()
}

// HasUserName returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsUserNameSet() bool {
	if o != nil && o.UserName.IsSet() {
		return true
	}

	return false
}

// SetUserName gets a reference to the given NullableString and assigns it to the UserName field.
func (o *EmployeeFullDto) SetUserName(v string) {
	o.UserName.Set(&v)
}
// SetUserNameNil sets the value for UserName to be an explicit nil
func (o *EmployeeFullDto) SetUserNameNil() {
	o.UserName.Set(nil)
}

// UnsetUserName ensures that no value is present for UserName, not even an explicit nil
func (o *EmployeeFullDto) UnsetUserName() {
	o.UserName.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployeeFullDto) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployeeFullDto) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsEmailSet() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *EmployeeFullDto) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *EmployeeFullDto) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *EmployeeFullDto) UnsetEmail() {
	o.Email.Unset()
}

// GetContacts returns the Contacts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployeeFullDto) GetContacts() []Contact {
	if o == nil {
		var ret []Contact
		return ret
	}
	return o.Contacts
}

// GetContactsOk returns a tuple with the Contacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployeeFullDto) GetContactsOk() ([]Contact, bool) {
	if o == nil || IsNil(o.Contacts) {
		return nil, false
	}
	return o.Contacts, true
}

// HasContacts returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsContactsSet() bool {
	if o != nil && !IsNil(o.Contacts) {
		return true
	}

	return false
}

// SetContacts gets a reference to the given []Contact and assigns it to the Contacts field.
func (o *EmployeeFullDto) SetContacts(v []Contact) {
	o.Contacts = v
}

// GetBirthday returns the Birthday field value if set, zero value otherwise.
func (o *EmployeeFullDto) GetBirthday() ApiDateTime {
	if o == nil || IsNil(o.Birthday) {
		var ret ApiDateTime
		return ret
	}
	return *o.Birthday
}

// GetBirthdayOk returns a tuple with the Birthday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployeeFullDto) GetBirthdayOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.Birthday) {
		return nil, false
	}
	return o.Birthday, true
}

// HasBirthday returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsBirthdaySet() bool {
	if o != nil && !IsNil(o.Birthday) {
		return true
	}

	return false
}

// SetBirthday gets a reference to the given ApiDateTime and assigns it to the Birthday field.
func (o *EmployeeFullDto) SetBirthday(v ApiDateTime) {
	o.Birthday = &v
}

// GetSex returns the Sex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployeeFullDto) GetSex() string {
	if o == nil || IsNil(o.Sex.Get()) {
		var ret string
		return ret
	}
	return *o.Sex.Get()
}

// GetSexOk returns a tuple with the Sex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployeeFullDto) GetSexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sex.Get(), o.Sex.IsSet()
}

// HasSex returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsSexSet() bool {
	if o != nil && o.Sex.IsSet() {
		return true
	}

	return false
}

// SetSex gets a reference to the given NullableString and assigns it to the Sex field.
func (o *EmployeeFullDto) SetSex(v string) {
	o.Sex.Set(&v)
}
// SetSexNil sets the value for Sex to be an explicit nil
func (o *EmployeeFullDto) SetSexNil() {
	o.Sex.Set(nil)
}

// UnsetSex ensures that no value is present for Sex, not even an explicit nil
func (o *EmployeeFullDto) UnsetSex() {
	o.Sex.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EmployeeFullDto) GetStatus() EmployeeStatus {
	if o == nil || IsNil(o.Status) {
		var ret EmployeeStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployeeFullDto) GetStatusOk() (*EmployeeStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsStatusSet() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given EmployeeStatus and assigns it to the Status field.
func (o *EmployeeFullDto) SetStatus(v EmployeeStatus) {
	o.Status = &v
}

// GetActivationStatus returns the ActivationStatus field value if set, zero value otherwise.
func (o *EmployeeFullDto) GetActivationStatus() EmployeeActivationStatus {
	if o == nil || IsNil(o.ActivationStatus) {
		var ret EmployeeActivationStatus
		return ret
	}
	return *o.ActivationStatus
}

// GetActivationStatusOk returns a tuple with the ActivationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployeeFullDto) GetActivationStatusOk() (*EmployeeActivationStatus, bool) {
	if o == nil || IsNil(o.ActivationStatus) {
		return nil, false
	}
	return o.ActivationStatus, true
}

// HasActivationStatus returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsActivationStatusSet() bool {
	if o != nil && !IsNil(o.ActivationStatus) {
		return true
	}

	return false
}

// SetActivationStatus gets a reference to the given EmployeeActivationStatus and assigns it to the ActivationStatus field.
func (o *EmployeeFullDto) SetActivationStatus(v EmployeeActivationStatus) {
	o.ActivationStatus = &v
}

// GetTerminated returns the Terminated field value if set, zero value otherwise.
func (o *EmployeeFullDto) GetTerminated() ApiDateTime {
	if o == nil || IsNil(o.Terminated) {
		var ret ApiDateTime
		return ret
	}
	return *o.Terminated
}

// GetTerminatedOk returns a tuple with the Terminated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployeeFullDto) GetTerminatedOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.Terminated) {
		return nil, false
	}
	return o.Terminated, true
}

// HasTerminated returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsTerminatedSet() bool {
	if o != nil && !IsNil(o.Terminated) {
		return true
	}

	return false
}

// SetTerminated gets a reference to the given ApiDateTime and assigns it to the Terminated field.
func (o *EmployeeFullDto) SetTerminated(v ApiDateTime) {
	o.Terminated = &v
}

// GetDepartment returns the Department field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployeeFullDto) GetDepartment() string {
	if o == nil || IsNil(o.Department.Get()) {
		var ret string
		return ret
	}
	return *o.Department.Get()
}

// GetDepartmentOk returns a tuple with the Department field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployeeFullDto) GetDepartmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Department.Get(), o.Department.IsSet()
}

// HasDepartment returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsDepartmentSet() bool {
	if o != nil && o.Department.IsSet() {
		return true
	}

	return false
}

// SetDepartment gets a reference to the given NullableString and assigns it to the Department field.
func (o *EmployeeFullDto) SetDepartment(v string) {
	o.Department.Set(&v)
}
// SetDepartmentNil sets the value for Department to be an explicit nil
func (o *EmployeeFullDto) SetDepartmentNil() {
	o.Department.Set(nil)
}

// UnsetDepartment ensures that no value is present for Department, not even an explicit nil
func (o *EmployeeFullDto) UnsetDepartment() {
	o.Department.Unset()
}

// GetWorkFrom returns the WorkFrom field value if set, zero value otherwise.
func (o *EmployeeFullDto) GetWorkFrom() ApiDateTime {
	if o == nil || IsNil(o.WorkFrom) {
		var ret ApiDateTime
		return ret
	}
	return *o.WorkFrom
}

// GetWorkFromOk returns a tuple with the WorkFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployeeFullDto) GetWorkFromOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.WorkFrom) {
		return nil, false
	}
	return o.WorkFrom, true
}

// HasWorkFrom returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsWorkFromSet() bool {
	if o != nil && !IsNil(o.WorkFrom) {
		return true
	}

	return false
}

// SetWorkFrom gets a reference to the given ApiDateTime and assigns it to the WorkFrom field.
func (o *EmployeeFullDto) SetWorkFrom(v ApiDateTime) {
	o.WorkFrom = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployeeFullDto) GetGroups() []GroupSummaryDto {
	if o == nil {
		var ret []GroupSummaryDto
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployeeFullDto) GetGroupsOk() ([]GroupSummaryDto, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsGroupsSet() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []GroupSummaryDto and assigns it to the Groups field.
func (o *EmployeeFullDto) SetGroups(v []GroupSummaryDto) {
	o.Groups = v
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployeeFullDto) GetLocation() string {
	if o == nil || IsNil(o.Location.Get()) {
		var ret string
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployeeFullDto) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsLocationSet() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableString and assigns it to the Location field.
func (o *EmployeeFullDto) SetLocation(v string) {
	o.Location.Set(&v)
}
// SetLocationNil sets the value for Location to be an explicit nil
func (o *EmployeeFullDto) SetLocationNil() {
	o.Location.Set(nil)
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *EmployeeFullDto) UnsetLocation() {
	o.Location.Unset()
}

// GetNotes returns the Notes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployeeFullDto) GetNotes() string {
	if o == nil || IsNil(o.Notes.Get()) {
		var ret string
		return ret
	}
	return *o.Notes.Get()
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployeeFullDto) GetNotesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Notes.Get(), o.Notes.IsSet()
}

// HasNotes returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsNotesSet() bool {
	if o != nil && o.Notes.IsSet() {
		return true
	}

	return false
}

// SetNotes gets a reference to the given NullableString and assigns it to the Notes field.
func (o *EmployeeFullDto) SetNotes(v string) {
	o.Notes.Set(&v)
}
// SetNotesNil sets the value for Notes to be an explicit nil
func (o *EmployeeFullDto) SetNotesNil() {
	o.Notes.Set(nil)
}

// UnsetNotes ensures that no value is present for Notes, not even an explicit nil
func (o *EmployeeFullDto) UnsetNotes() {
	o.Notes.Unset()
}

// GetIsAdmin returns the IsAdmin field value if set, zero value otherwise.
func (o *EmployeeFullDto) GetIsAdmin() bool {
	if o == nil || IsNil(o.IsAdmin) {
		var ret bool
		return ret
	}
	return *o.IsAdmin
}

// GetIsAdminOk returns a tuple with the IsAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployeeFullDto) GetIsAdminOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAdmin) {
		return nil, false
	}
	return o.IsAdmin, true
}

// HasIsAdmin returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsIsAdminSet() bool {
	if o != nil && !IsNil(o.IsAdmin) {
		return true
	}

	return false
}

// SetIsAdmin gets a reference to the given bool and assigns it to the IsAdmin field.
func (o *EmployeeFullDto) SetIsAdmin(v bool) {
	o.IsAdmin = &v
}

// GetIsRoomAdmin returns the IsRoomAdmin field value if set, zero value otherwise.
func (o *EmployeeFullDto) GetIsRoomAdmin() bool {
	if o == nil || IsNil(o.IsRoomAdmin) {
		var ret bool
		return ret
	}
	return *o.IsRoomAdmin
}

// GetIsRoomAdminOk returns a tuple with the IsRoomAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployeeFullDto) GetIsRoomAdminOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRoomAdmin) {
		return nil, false
	}
	return o.IsRoomAdmin, true
}

// HasIsRoomAdmin returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsIsRoomAdminSet() bool {
	if o != nil && !IsNil(o.IsRoomAdmin) {
		return true
	}

	return false
}

// SetIsRoomAdmin gets a reference to the given bool and assigns it to the IsRoomAdmin field.
func (o *EmployeeFullDto) SetIsRoomAdmin(v bool) {
	o.IsRoomAdmin = &v
}

// GetIsLDAP returns the IsLDAP field value if set, zero value otherwise.
func (o *EmployeeFullDto) GetIsLDAP() bool {
	if o == nil || IsNil(o.IsLDAP) {
		var ret bool
		return ret
	}
	return *o.IsLDAP
}

// GetIsLDAPOk returns a tuple with the IsLDAP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployeeFullDto) GetIsLDAPOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLDAP) {
		return nil, false
	}
	return o.IsLDAP, true
}

// HasIsLDAP returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsIsLDAPSet() bool {
	if o != nil && !IsNil(o.IsLDAP) {
		return true
	}

	return false
}

// SetIsLDAP gets a reference to the given bool and assigns it to the IsLDAP field.
func (o *EmployeeFullDto) SetIsLDAP(v bool) {
	o.IsLDAP = &v
}

// GetListAdminModules returns the ListAdminModules field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployeeFullDto) GetListAdminModules() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ListAdminModules
}

// GetListAdminModulesOk returns a tuple with the ListAdminModules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployeeFullDto) GetListAdminModulesOk() ([]string, bool) {
	if o == nil || IsNil(o.ListAdminModules) {
		return nil, false
	}
	return o.ListAdminModules, true
}

// HasListAdminModules returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsListAdminModulesSet() bool {
	if o != nil && !IsNil(o.ListAdminModules) {
		return true
	}

	return false
}

// SetListAdminModules gets a reference to the given []string and assigns it to the ListAdminModules field.
func (o *EmployeeFullDto) SetListAdminModules(v []string) {
	o.ListAdminModules = v
}

// GetIsOwner returns the IsOwner field value if set, zero value otherwise.
func (o *EmployeeFullDto) GetIsOwner() bool {
	if o == nil || IsNil(o.IsOwner) {
		var ret bool
		return ret
	}
	return *o.IsOwner
}

// GetIsOwnerOk returns a tuple with the IsOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployeeFullDto) GetIsOwnerOk() (*bool, bool) {
	if o == nil || IsNil(o.IsOwner) {
		return nil, false
	}
	return o.IsOwner, true
}

// HasIsOwner returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsIsOwnerSet() bool {
	if o != nil && !IsNil(o.IsOwner) {
		return true
	}

	return false
}

// SetIsOwner gets a reference to the given bool and assigns it to the IsOwner field.
func (o *EmployeeFullDto) SetIsOwner(v bool) {
	o.IsOwner = &v
}

// GetIsVisitor returns the IsVisitor field value if set, zero value otherwise.
func (o *EmployeeFullDto) GetIsVisitor() bool {
	if o == nil || IsNil(o.IsVisitor) {
		var ret bool
		return ret
	}
	return *o.IsVisitor
}

// GetIsVisitorOk returns a tuple with the IsVisitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployeeFullDto) GetIsVisitorOk() (*bool, bool) {
	if o == nil || IsNil(o.IsVisitor) {
		return nil, false
	}
	return o.IsVisitor, true
}

// HasIsVisitor returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsIsVisitorSet() bool {
	if o != nil && !IsNil(o.IsVisitor) {
		return true
	}

	return false
}

// SetIsVisitor gets a reference to the given bool and assigns it to the IsVisitor field.
func (o *EmployeeFullDto) SetIsVisitor(v bool) {
	o.IsVisitor = &v
}

// GetIsCollaborator returns the IsCollaborator field value if set, zero value otherwise.
func (o *EmployeeFullDto) GetIsCollaborator() bool {
	if o == nil || IsNil(o.IsCollaborator) {
		var ret bool
		return ret
	}
	return *o.IsCollaborator
}

// GetIsCollaboratorOk returns a tuple with the IsCollaborator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployeeFullDto) GetIsCollaboratorOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCollaborator) {
		return nil, false
	}
	return o.IsCollaborator, true
}

// HasIsCollaborator returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsIsCollaboratorSet() bool {
	if o != nil && !IsNil(o.IsCollaborator) {
		return true
	}

	return false
}

// SetIsCollaborator gets a reference to the given bool and assigns it to the IsCollaborator field.
func (o *EmployeeFullDto) SetIsCollaborator(v bool) {
	o.IsCollaborator = &v
}

// GetCultureName returns the CultureName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployeeFullDto) GetCultureName() string {
	if o == nil || IsNil(o.CultureName.Get()) {
		var ret string
		return ret
	}
	return *o.CultureName.Get()
}

// GetCultureNameOk returns a tuple with the CultureName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployeeFullDto) GetCultureNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CultureName.Get(), o.CultureName.IsSet()
}

// HasCultureName returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsCultureNameSet() bool {
	if o != nil && o.CultureName.IsSet() {
		return true
	}

	return false
}

// SetCultureName gets a reference to the given NullableString and assigns it to the CultureName field.
func (o *EmployeeFullDto) SetCultureName(v string) {
	o.CultureName.Set(&v)
}
// SetCultureNameNil sets the value for CultureName to be an explicit nil
func (o *EmployeeFullDto) SetCultureNameNil() {
	o.CultureName.Set(nil)
}

// UnsetCultureName ensures that no value is present for CultureName, not even an explicit nil
func (o *EmployeeFullDto) UnsetCultureName() {
	o.CultureName.Unset()
}

// GetMobilePhone returns the MobilePhone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployeeFullDto) GetMobilePhone() string {
	if o == nil || IsNil(o.MobilePhone.Get()) {
		var ret string
		return ret
	}
	return *o.MobilePhone.Get()
}

// GetMobilePhoneOk returns a tuple with the MobilePhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployeeFullDto) GetMobilePhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MobilePhone.Get(), o.MobilePhone.IsSet()
}

// HasMobilePhone returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsMobilePhoneSet() bool {
	if o != nil && o.MobilePhone.IsSet() {
		return true
	}

	return false
}

// SetMobilePhone gets a reference to the given NullableString and assigns it to the MobilePhone field.
func (o *EmployeeFullDto) SetMobilePhone(v string) {
	o.MobilePhone.Set(&v)
}
// SetMobilePhoneNil sets the value for MobilePhone to be an explicit nil
func (o *EmployeeFullDto) SetMobilePhoneNil() {
	o.MobilePhone.Set(nil)
}

// UnsetMobilePhone ensures that no value is present for MobilePhone, not even an explicit nil
func (o *EmployeeFullDto) UnsetMobilePhone() {
	o.MobilePhone.Unset()
}

// GetMobilePhoneActivationStatus returns the MobilePhoneActivationStatus field value if set, zero value otherwise.
func (o *EmployeeFullDto) GetMobilePhoneActivationStatus() MobilePhoneActivationStatus {
	if o == nil || IsNil(o.MobilePhoneActivationStatus) {
		var ret MobilePhoneActivationStatus
		return ret
	}
	return *o.MobilePhoneActivationStatus
}

// GetMobilePhoneActivationStatusOk returns a tuple with the MobilePhoneActivationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployeeFullDto) GetMobilePhoneActivationStatusOk() (*MobilePhoneActivationStatus, bool) {
	if o == nil || IsNil(o.MobilePhoneActivationStatus) {
		return nil, false
	}
	return o.MobilePhoneActivationStatus, true
}

// HasMobilePhoneActivationStatus returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsMobilePhoneActivationStatusSet() bool {
	if o != nil && !IsNil(o.MobilePhoneActivationStatus) {
		return true
	}

	return false
}

// SetMobilePhoneActivationStatus gets a reference to the given MobilePhoneActivationStatus and assigns it to the MobilePhoneActivationStatus field.
func (o *EmployeeFullDto) SetMobilePhoneActivationStatus(v MobilePhoneActivationStatus) {
	o.MobilePhoneActivationStatus = &v
}

// GetIsSSO returns the IsSSO field value if set, zero value otherwise.
func (o *EmployeeFullDto) GetIsSSO() bool {
	if o == nil || IsNil(o.IsSSO) {
		var ret bool
		return ret
	}
	return *o.IsSSO
}

// GetIsSSOOk returns a tuple with the IsSSO field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployeeFullDto) GetIsSSOOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSSO) {
		return nil, false
	}
	return o.IsSSO, true
}

// HasIsSSO returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsIsSSOSet() bool {
	if o != nil && !IsNil(o.IsSSO) {
		return true
	}

	return false
}

// SetIsSSO gets a reference to the given bool and assigns it to the IsSSO field.
func (o *EmployeeFullDto) SetIsSSO(v bool) {
	o.IsSSO = &v
}

// GetTheme returns the Theme field value if set, zero value otherwise.
func (o *EmployeeFullDto) GetTheme() DarkThemeSettingsType {
	if o == nil || IsNil(o.Theme) {
		var ret DarkThemeSettingsType
		return ret
	}
	return *o.Theme
}

// GetThemeOk returns a tuple with the Theme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployeeFullDto) GetThemeOk() (*DarkThemeSettingsType, bool) {
	if o == nil || IsNil(o.Theme) {
		return nil, false
	}
	return o.Theme, true
}

// HasTheme returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsThemeSet() bool {
	if o != nil && !IsNil(o.Theme) {
		return true
	}

	return false
}

// SetTheme gets a reference to the given DarkThemeSettingsType and assigns it to the Theme field.
func (o *EmployeeFullDto) SetTheme(v DarkThemeSettingsType) {
	o.Theme = &v
}

// GetQuotaLimit returns the QuotaLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployeeFullDto) GetQuotaLimit() int64 {
	if o == nil || IsNil(o.QuotaLimit.Get()) {
		var ret int64
		return ret
	}
	return *o.QuotaLimit.Get()
}

// GetQuotaLimitOk returns a tuple with the QuotaLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployeeFullDto) GetQuotaLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.QuotaLimit.Get(), o.QuotaLimit.IsSet()
}

// HasQuotaLimit returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsQuotaLimitSet() bool {
	if o != nil && o.QuotaLimit.IsSet() {
		return true
	}

	return false
}

// SetQuotaLimit gets a reference to the given NullableInt64 and assigns it to the QuotaLimit field.
func (o *EmployeeFullDto) SetQuotaLimit(v int64) {
	o.QuotaLimit.Set(&v)
}
// SetQuotaLimitNil sets the value for QuotaLimit to be an explicit nil
func (o *EmployeeFullDto) SetQuotaLimitNil() {
	o.QuotaLimit.Set(nil)
}

// UnsetQuotaLimit ensures that no value is present for QuotaLimit, not even an explicit nil
func (o *EmployeeFullDto) UnsetQuotaLimit() {
	o.QuotaLimit.Unset()
}

// GetUsedSpace returns the UsedSpace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployeeFullDto) GetUsedSpace() float64 {
	if o == nil || IsNil(o.UsedSpace.Get()) {
		var ret float64
		return ret
	}
	return *o.UsedSpace.Get()
}

// GetUsedSpaceOk returns a tuple with the UsedSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployeeFullDto) GetUsedSpaceOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UsedSpace.Get(), o.UsedSpace.IsSet()
}

// HasUsedSpace returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsUsedSpaceSet() bool {
	if o != nil && o.UsedSpace.IsSet() {
		return true
	}

	return false
}

// SetUsedSpace gets a reference to the given NullableFloat64 and assigns it to the UsedSpace field.
func (o *EmployeeFullDto) SetUsedSpace(v float64) {
	o.UsedSpace.Set(&v)
}
// SetUsedSpaceNil sets the value for UsedSpace to be an explicit nil
func (o *EmployeeFullDto) SetUsedSpaceNil() {
	o.UsedSpace.Set(nil)
}

// UnsetUsedSpace ensures that no value is present for UsedSpace, not even an explicit nil
func (o *EmployeeFullDto) UnsetUsedSpace() {
	o.UsedSpace.Unset()
}

// GetShared returns the Shared field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployeeFullDto) GetShared() bool {
	if o == nil || IsNil(o.Shared.Get()) {
		var ret bool
		return ret
	}
	return *o.Shared.Get()
}

// GetSharedOk returns a tuple with the Shared field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployeeFullDto) GetSharedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Shared.Get(), o.Shared.IsSet()
}

// HasShared returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsSharedSet() bool {
	if o != nil && o.Shared.IsSet() {
		return true
	}

	return false
}

// SetShared gets a reference to the given NullableBool and assigns it to the Shared field.
func (o *EmployeeFullDto) SetShared(v bool) {
	o.Shared.Set(&v)
}
// SetSharedNil sets the value for Shared to be an explicit nil
func (o *EmployeeFullDto) SetSharedNil() {
	o.Shared.Set(nil)
}

// UnsetShared ensures that no value is present for Shared, not even an explicit nil
func (o *EmployeeFullDto) UnsetShared() {
	o.Shared.Unset()
}

// GetIsCustomQuota returns the IsCustomQuota field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployeeFullDto) GetIsCustomQuota() bool {
	if o == nil || IsNil(o.IsCustomQuota.Get()) {
		var ret bool
		return ret
	}
	return *o.IsCustomQuota.Get()
}

// GetIsCustomQuotaOk returns a tuple with the IsCustomQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployeeFullDto) GetIsCustomQuotaOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsCustomQuota.Get(), o.IsCustomQuota.IsSet()
}

// HasIsCustomQuota returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsIsCustomQuotaSet() bool {
	if o != nil && o.IsCustomQuota.IsSet() {
		return true
	}

	return false
}

// SetIsCustomQuota gets a reference to the given NullableBool and assigns it to the IsCustomQuota field.
func (o *EmployeeFullDto) SetIsCustomQuota(v bool) {
	o.IsCustomQuota.Set(&v)
}
// SetIsCustomQuotaNil sets the value for IsCustomQuota to be an explicit nil
func (o *EmployeeFullDto) SetIsCustomQuotaNil() {
	o.IsCustomQuota.Set(nil)
}

// UnsetIsCustomQuota ensures that no value is present for IsCustomQuota, not even an explicit nil
func (o *EmployeeFullDto) UnsetIsCustomQuota() {
	o.IsCustomQuota.Unset()
}

// GetLoginEventId returns the LoginEventId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployeeFullDto) GetLoginEventId() int32 {
	if o == nil || IsNil(o.LoginEventId.Get()) {
		var ret int32
		return ret
	}
	return *o.LoginEventId.Get()
}

// GetLoginEventIdOk returns a tuple with the LoginEventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployeeFullDto) GetLoginEventIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.LoginEventId.Get(), o.LoginEventId.IsSet()
}

// HasLoginEventId returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsLoginEventIdSet() bool {
	if o != nil && o.LoginEventId.IsSet() {
		return true
	}

	return false
}

// SetLoginEventId gets a reference to the given NullableInt32 and assigns it to the LoginEventId field.
func (o *EmployeeFullDto) SetLoginEventId(v int32) {
	o.LoginEventId.Set(&v)
}
// SetLoginEventIdNil sets the value for LoginEventId to be an explicit nil
func (o *EmployeeFullDto) SetLoginEventIdNil() {
	o.LoginEventId.Set(nil)
}

// UnsetLoginEventId ensures that no value is present for LoginEventId, not even an explicit nil
func (o *EmployeeFullDto) UnsetLoginEventId() {
	o.LoginEventId.Unset()
}

// GetAuthCookieLifetime returns the AuthCookieLifetime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployeeFullDto) GetAuthCookieLifetime() float64 {
	if o == nil || IsNil(o.AuthCookieLifetime.Get()) {
		var ret float64
		return ret
	}
	return *o.AuthCookieLifetime.Get()
}

// GetAuthCookieLifetimeOk returns a tuple with the AuthCookieLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployeeFullDto) GetAuthCookieLifetimeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthCookieLifetime.Get(), o.AuthCookieLifetime.IsSet()
}

// HasAuthCookieLifetime returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsAuthCookieLifetimeSet() bool {
	if o != nil && o.AuthCookieLifetime.IsSet() {
		return true
	}

	return false
}

// SetAuthCookieLifetime gets a reference to the given NullableFloat64 and assigns it to the AuthCookieLifetime field.
func (o *EmployeeFullDto) SetAuthCookieLifetime(v float64) {
	o.AuthCookieLifetime.Set(&v)
}
// SetAuthCookieLifetimeNil sets the value for AuthCookieLifetime to be an explicit nil
func (o *EmployeeFullDto) SetAuthCookieLifetimeNil() {
	o.AuthCookieLifetime.Set(nil)
}

// UnsetAuthCookieLifetime ensures that no value is present for AuthCookieLifetime, not even an explicit nil
func (o *EmployeeFullDto) UnsetAuthCookieLifetime() {
	o.AuthCookieLifetime.Unset()
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *EmployeeFullDto) GetCreatedBy() EmployeeDto {
	if o == nil || IsNil(o.CreatedBy) {
		var ret EmployeeDto
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployeeFullDto) GetCreatedByOk() (*EmployeeDto, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsCreatedBySet() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given EmployeeDto and assigns it to the CreatedBy field.
func (o *EmployeeFullDto) SetCreatedBy(v EmployeeDto) {
	o.CreatedBy = &v
}

// GetRegistrationDate returns the RegistrationDate field value if set, zero value otherwise.
func (o *EmployeeFullDto) GetRegistrationDate() ApiDateTime {
	if o == nil || IsNil(o.RegistrationDate) {
		var ret ApiDateTime
		return ret
	}
	return *o.RegistrationDate
}

// GetRegistrationDateOk returns a tuple with the RegistrationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployeeFullDto) GetRegistrationDateOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.RegistrationDate) {
		return nil, false
	}
	return o.RegistrationDate, true
}

// HasRegistrationDate returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsRegistrationDateSet() bool {
	if o != nil && !IsNil(o.RegistrationDate) {
		return true
	}

	return false
}

// SetRegistrationDate gets a reference to the given ApiDateTime and assigns it to the RegistrationDate field.
func (o *EmployeeFullDto) SetRegistrationDate(v ApiDateTime) {
	o.RegistrationDate = &v
}

// GetHasPersonalFolder returns the HasPersonalFolder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployeeFullDto) GetHasPersonalFolder() bool {
	if o == nil || IsNil(o.HasPersonalFolder.Get()) {
		var ret bool
		return ret
	}
	return *o.HasPersonalFolder.Get()
}

// GetHasPersonalFolderOk returns a tuple with the HasPersonalFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployeeFullDto) GetHasPersonalFolderOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.HasPersonalFolder.Get(), o.HasPersonalFolder.IsSet()
}

// HasHasPersonalFolder returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsHasPersonalFolderSet() bool {
	if o != nil && o.HasPersonalFolder.IsSet() {
		return true
	}

	return false
}

// SetHasPersonalFolder gets a reference to the given NullableBool and assigns it to the HasPersonalFolder field.
func (o *EmployeeFullDto) SetHasPersonalFolder(v bool) {
	o.HasPersonalFolder.Set(&v)
}
// SetHasPersonalFolderNil sets the value for HasPersonalFolder to be an explicit nil
func (o *EmployeeFullDto) SetHasPersonalFolderNil() {
	o.HasPersonalFolder.Set(nil)
}

// UnsetHasPersonalFolder ensures that no value is present for HasPersonalFolder, not even an explicit nil
func (o *EmployeeFullDto) UnsetHasPersonalFolder() {
	o.HasPersonalFolder.Unset()
}

// GetTfaAppEnabled returns the TfaAppEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployeeFullDto) GetTfaAppEnabled() bool {
	if o == nil || IsNil(o.TfaAppEnabled.Get()) {
		var ret bool
		return ret
	}
	return *o.TfaAppEnabled.Get()
}

// GetTfaAppEnabledOk returns a tuple with the TfaAppEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployeeFullDto) GetTfaAppEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.TfaAppEnabled.Get(), o.TfaAppEnabled.IsSet()
}

// HasTfaAppEnabled returns a boolean if a field has been set.
func (o *EmployeeFullDto) IsTfaAppEnabledSet() bool {
	if o != nil && o.TfaAppEnabled.IsSet() {
		return true
	}

	return false
}

// SetTfaAppEnabled gets a reference to the given NullableBool and assigns it to the TfaAppEnabled field.
func (o *EmployeeFullDto) SetTfaAppEnabled(v bool) {
	o.TfaAppEnabled.Set(&v)
}
// SetTfaAppEnabledNil sets the value for TfaAppEnabled to be an explicit nil
func (o *EmployeeFullDto) SetTfaAppEnabledNil() {
	o.TfaAppEnabled.Set(nil)
}

// UnsetTfaAppEnabled ensures that no value is present for TfaAppEnabled, not even an explicit nil
func (o *EmployeeFullDto) UnsetTfaAppEnabled() {
	o.TfaAppEnabled.Unset()
}

func (o EmployeeFullDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmployeeFullDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Avatar.IsSet() {
		toSerialize["avatar"] = o.Avatar.Get()
	}
	if o.AvatarOriginal.IsSet() {
		toSerialize["avatarOriginal"] = o.AvatarOriginal.Get()
	}
	if o.AvatarMax.IsSet() {
		toSerialize["avatarMax"] = o.AvatarMax.Get()
	}
	if o.AvatarMedium.IsSet() {
		toSerialize["avatarMedium"] = o.AvatarMedium.Get()
	}
	if o.AvatarSmall.IsSet() {
		toSerialize["avatarSmall"] = o.AvatarSmall.Get()
	}
	if o.ProfileUrl.IsSet() {
		toSerialize["profileUrl"] = o.ProfileUrl.Get()
	}
	if !IsNil(o.HasAvatar) {
		toSerialize["hasAvatar"] = o.HasAvatar
	}
	if !IsNil(o.IsAnonim) {
		toSerialize["isAnonim"] = o.IsAnonim
	}
	if o.FirstName.IsSet() {
		toSerialize["firstName"] = o.FirstName.Get()
	}
	if o.LastName.IsSet() {
		toSerialize["lastName"] = o.LastName.Get()
	}
	if o.UserName.IsSet() {
		toSerialize["userName"] = o.UserName.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.Contacts != nil {
		toSerialize["contacts"] = o.Contacts
	}
	if !IsNil(o.Birthday) {
		toSerialize["birthday"] = o.Birthday
	}
	if o.Sex.IsSet() {
		toSerialize["sex"] = o.Sex.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.ActivationStatus) {
		toSerialize["activationStatus"] = o.ActivationStatus
	}
	if !IsNil(o.Terminated) {
		toSerialize["terminated"] = o.Terminated
	}
	if o.Department.IsSet() {
		toSerialize["department"] = o.Department.Get()
	}
	if !IsNil(o.WorkFrom) {
		toSerialize["workFrom"] = o.WorkFrom
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if o.Location.IsSet() {
		toSerialize["location"] = o.Location.Get()
	}
	if o.Notes.IsSet() {
		toSerialize["notes"] = o.Notes.Get()
	}
	if !IsNil(o.IsAdmin) {
		toSerialize["isAdmin"] = o.IsAdmin
	}
	if !IsNil(o.IsRoomAdmin) {
		toSerialize["isRoomAdmin"] = o.IsRoomAdmin
	}
	if !IsNil(o.IsLDAP) {
		toSerialize["isLDAP"] = o.IsLDAP
	}
	if o.ListAdminModules != nil {
		toSerialize["listAdminModules"] = o.ListAdminModules
	}
	if !IsNil(o.IsOwner) {
		toSerialize["isOwner"] = o.IsOwner
	}
	if !IsNil(o.IsVisitor) {
		toSerialize["isVisitor"] = o.IsVisitor
	}
	if !IsNil(o.IsCollaborator) {
		toSerialize["isCollaborator"] = o.IsCollaborator
	}
	if o.CultureName.IsSet() {
		toSerialize["cultureName"] = o.CultureName.Get()
	}
	if o.MobilePhone.IsSet() {
		toSerialize["mobilePhone"] = o.MobilePhone.Get()
	}
	if !IsNil(o.MobilePhoneActivationStatus) {
		toSerialize["mobilePhoneActivationStatus"] = o.MobilePhoneActivationStatus
	}
	if !IsNil(o.IsSSO) {
		toSerialize["isSSO"] = o.IsSSO
	}
	if !IsNil(o.Theme) {
		toSerialize["theme"] = o.Theme
	}
	if o.QuotaLimit.IsSet() {
		toSerialize["quotaLimit"] = o.QuotaLimit.Get()
	}
	if o.UsedSpace.IsSet() {
		toSerialize["usedSpace"] = o.UsedSpace.Get()
	}
	if o.Shared.IsSet() {
		toSerialize["shared"] = o.Shared.Get()
	}
	if o.IsCustomQuota.IsSet() {
		toSerialize["isCustomQuota"] = o.IsCustomQuota.Get()
	}
	if o.LoginEventId.IsSet() {
		toSerialize["loginEventId"] = o.LoginEventId.Get()
	}
	if o.AuthCookieLifetime.IsSet() {
		toSerialize["authCookieLifetime"] = o.AuthCookieLifetime.Get()
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !IsNil(o.RegistrationDate) {
		toSerialize["registrationDate"] = o.RegistrationDate
	}
	if o.HasPersonalFolder.IsSet() {
		toSerialize["hasPersonalFolder"] = o.HasPersonalFolder.Get()
	}
	if o.TfaAppEnabled.IsSet() {
		toSerialize["tfaAppEnabled"] = o.TfaAppEnabled.Get()
	}
	return toSerialize, nil
}

type NullableEmployeeFullDto struct {
	value *EmployeeFullDto
	isSet bool
}

func (v NullableEmployeeFullDto) Get() *EmployeeFullDto {
	return v.value
}

func (v *NullableEmployeeFullDto) Set(val *EmployeeFullDto) {
	v.value = val
	v.isSet = true
}

func (v NullableEmployeeFullDto) IsSet() bool {
	return v.isSet
}

func (v *NullableEmployeeFullDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmployeeFullDto(val *EmployeeFullDto) *NullableEmployeeFullDto {
	return &NullableEmployeeFullDto{value: val, isSet: true}
}

func (v NullableEmployeeFullDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmployeeFullDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

