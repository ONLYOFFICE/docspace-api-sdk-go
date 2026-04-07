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
)

// checks if the UserInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserInfo{}

// UserInfo The user information.
type UserInfo struct {
	// The user ID.
	Id *string `json:"id,omitempty"`
	// The user's first name.
	FirstName NullableString `json:"firstName,omitempty"`
	// The user's last name.
	LastName NullableString `json:"lastName,omitempty"`
	// The user username.
	UserName NullableString `json:"userName,omitempty"`
	// The user birthday.
	BirthDate NullableTime `json:"birthDate,omitempty"`
	// The user sex (male or female).
	Sex NullableBool `json:"sex,omitempty"`
	Status *EmployeeStatus `json:"status,omitempty"`
	ActivationStatus *EmployeeActivationStatus `json:"activationStatus,omitempty"`
	// The date and time when the user account was terminated.
	TerminatedDate NullableTime `json:"terminatedDate,omitempty"`
	// The user title.
	Title NullableString `json:"title,omitempty"`
	// The user registration date.
	WorkFromDate NullableTime `json:"workFromDate,omitempty"`
	// The user email address.
	Email NullableString `json:"email,omitempty"`
	// The list of user contacts in the string format.
	Contacts NullableString `json:"contacts,omitempty"`
	// The list of user contacts.
	ContactsList []string `json:"contactsList,omitempty"`
	// The user location.
	Location NullableString `json:"location,omitempty"`
	// The user notes.
	Notes NullableString `json:"notes,omitempty"`
	// Specifies if the user account was removed or not.
	Removed *bool `json:"removed,omitempty"`
	// The date and time when the user account was last modified.
	LastModified *time.Time `json:"lastModified,omitempty"`
	// The tenant ID.
	TenantId *int32 `json:"tenantId,omitempty"`
	// Specifies if the user is active or not.
	IsActive *bool `json:"isActive,omitempty"`
	// The user culture code.
	CultureName NullableString `json:"cultureName,omitempty"`
	// The user mobile phone.
	MobilePhone NullableString `json:"mobilePhone,omitempty"`
	MobilePhoneActivationStatus *MobilePhoneActivationStatus `json:"mobilePhoneActivationStatus,omitempty"`
	// The LDAP user identifier.
	Sid NullableString `json:"sid,omitempty"`
	// The LDAP user quota attribute.
	LdapQouta *int64 `json:"ldapQouta,omitempty"`
	// The SSO SAML user identifier.
	SsoNameId NullableString `json:"ssoNameId,omitempty"`
	// The SSO SAML user session identifier.
	SsoSessionId NullableString `json:"ssoSessionId,omitempty"`
	// The date and time when the user account was created.
	CreateDate *time.Time `json:"createDate,omitempty"`
	// The ID of the user who created the current user account.
	CreatedBy NullableString `json:"createdBy,omitempty"`
	// Specifies if tips, updates and offers are allowed to be sent to the user or not.
	Spam NullableBool `json:"spam,omitempty"`
	// Indicates whether the activation status of the employee or recipient is unchecked or inactive.  Depending on the context, this property evaluates the activation or eligibility status accordingly.
	CheckActivation *bool `json:"checkActivation,omitempty"`
}

// NewUserInfo instantiates a new UserInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserInfo() *UserInfo {
	this := UserInfo{}
	return &this
}

// NewUserInfoWithDefaults instantiates a new UserInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserInfoWithDefaults() *UserInfo {
	this := UserInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserInfo) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInfo) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserInfo) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UserInfo) SetId(v string) {
	o.Id = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserInfo) GetFirstName() string {
	if o == nil || IsNil(o.FirstName.Get()) {
		var ret string
		return ret
	}
	return *o.FirstName.Get()
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserInfo) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstName.Get(), o.FirstName.IsSet()
}

// HasFirstName returns a boolean if a field has been set.
func (o *UserInfo) IsFirstNameSet() bool {
	if o != nil && o.FirstName.IsSet() {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given NullableString and assigns it to the FirstName field.
func (o *UserInfo) SetFirstName(v string) {
	o.FirstName.Set(&v)
}
// SetFirstNameNil sets the value for FirstName to be an explicit nil
func (o *UserInfo) SetFirstNameNil() {
	o.FirstName.Set(nil)
}

// UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
func (o *UserInfo) UnsetFirstName() {
	o.FirstName.Unset()
}

// GetLastName returns the LastName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserInfo) GetLastName() string {
	if o == nil || IsNil(o.LastName.Get()) {
		var ret string
		return ret
	}
	return *o.LastName.Get()
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserInfo) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastName.Get(), o.LastName.IsSet()
}

// HasLastName returns a boolean if a field has been set.
func (o *UserInfo) IsLastNameSet() bool {
	if o != nil && o.LastName.IsSet() {
		return true
	}

	return false
}

// SetLastName gets a reference to the given NullableString and assigns it to the LastName field.
func (o *UserInfo) SetLastName(v string) {
	o.LastName.Set(&v)
}
// SetLastNameNil sets the value for LastName to be an explicit nil
func (o *UserInfo) SetLastNameNil() {
	o.LastName.Set(nil)
}

// UnsetLastName ensures that no value is present for LastName, not even an explicit nil
func (o *UserInfo) UnsetLastName() {
	o.LastName.Unset()
}

// GetUserName returns the UserName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserInfo) GetUserName() string {
	if o == nil || IsNil(o.UserName.Get()) {
		var ret string
		return ret
	}
	return *o.UserName.Get()
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserInfo) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserName.Get(), o.UserName.IsSet()
}

// HasUserName returns a boolean if a field has been set.
func (o *UserInfo) IsUserNameSet() bool {
	if o != nil && o.UserName.IsSet() {
		return true
	}

	return false
}

// SetUserName gets a reference to the given NullableString and assigns it to the UserName field.
func (o *UserInfo) SetUserName(v string) {
	o.UserName.Set(&v)
}
// SetUserNameNil sets the value for UserName to be an explicit nil
func (o *UserInfo) SetUserNameNil() {
	o.UserName.Set(nil)
}

// UnsetUserName ensures that no value is present for UserName, not even an explicit nil
func (o *UserInfo) UnsetUserName() {
	o.UserName.Unset()
}

// GetBirthDate returns the BirthDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserInfo) GetBirthDate() time.Time {
	if o == nil || IsNil(o.BirthDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.BirthDate.Get()
}

// GetBirthDateOk returns a tuple with the BirthDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserInfo) GetBirthDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.BirthDate.Get(), o.BirthDate.IsSet()
}

// HasBirthDate returns a boolean if a field has been set.
func (o *UserInfo) IsBirthDateSet() bool {
	if o != nil && o.BirthDate.IsSet() {
		return true
	}

	return false
}

// SetBirthDate gets a reference to the given NullableTime and assigns it to the BirthDate field.
func (o *UserInfo) SetBirthDate(v time.Time) {
	o.BirthDate.Set(&v)
}
// SetBirthDateNil sets the value for BirthDate to be an explicit nil
func (o *UserInfo) SetBirthDateNil() {
	o.BirthDate.Set(nil)
}

// UnsetBirthDate ensures that no value is present for BirthDate, not even an explicit nil
func (o *UserInfo) UnsetBirthDate() {
	o.BirthDate.Unset()
}

// GetSex returns the Sex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserInfo) GetSex() bool {
	if o == nil || IsNil(o.Sex.Get()) {
		var ret bool
		return ret
	}
	return *o.Sex.Get()
}

// GetSexOk returns a tuple with the Sex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserInfo) GetSexOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sex.Get(), o.Sex.IsSet()
}

// HasSex returns a boolean if a field has been set.
func (o *UserInfo) IsSexSet() bool {
	if o != nil && o.Sex.IsSet() {
		return true
	}

	return false
}

// SetSex gets a reference to the given NullableBool and assigns it to the Sex field.
func (o *UserInfo) SetSex(v bool) {
	o.Sex.Set(&v)
}
// SetSexNil sets the value for Sex to be an explicit nil
func (o *UserInfo) SetSexNil() {
	o.Sex.Set(nil)
}

// UnsetSex ensures that no value is present for Sex, not even an explicit nil
func (o *UserInfo) UnsetSex() {
	o.Sex.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UserInfo) GetStatus() EmployeeStatus {
	if o == nil || IsNil(o.Status) {
		var ret EmployeeStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInfo) GetStatusOk() (*EmployeeStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UserInfo) IsStatusSet() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given EmployeeStatus and assigns it to the Status field.
func (o *UserInfo) SetStatus(v EmployeeStatus) {
	o.Status = &v
}

// GetActivationStatus returns the ActivationStatus field value if set, zero value otherwise.
func (o *UserInfo) GetActivationStatus() EmployeeActivationStatus {
	if o == nil || IsNil(o.ActivationStatus) {
		var ret EmployeeActivationStatus
		return ret
	}
	return *o.ActivationStatus
}

// GetActivationStatusOk returns a tuple with the ActivationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInfo) GetActivationStatusOk() (*EmployeeActivationStatus, bool) {
	if o == nil || IsNil(o.ActivationStatus) {
		return nil, false
	}
	return o.ActivationStatus, true
}

// HasActivationStatus returns a boolean if a field has been set.
func (o *UserInfo) IsActivationStatusSet() bool {
	if o != nil && !IsNil(o.ActivationStatus) {
		return true
	}

	return false
}

// SetActivationStatus gets a reference to the given EmployeeActivationStatus and assigns it to the ActivationStatus field.
func (o *UserInfo) SetActivationStatus(v EmployeeActivationStatus) {
	o.ActivationStatus = &v
}

// GetTerminatedDate returns the TerminatedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserInfo) GetTerminatedDate() time.Time {
	if o == nil || IsNil(o.TerminatedDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.TerminatedDate.Get()
}

// GetTerminatedDateOk returns a tuple with the TerminatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserInfo) GetTerminatedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.TerminatedDate.Get(), o.TerminatedDate.IsSet()
}

// HasTerminatedDate returns a boolean if a field has been set.
func (o *UserInfo) IsTerminatedDateSet() bool {
	if o != nil && o.TerminatedDate.IsSet() {
		return true
	}

	return false
}

// SetTerminatedDate gets a reference to the given NullableTime and assigns it to the TerminatedDate field.
func (o *UserInfo) SetTerminatedDate(v time.Time) {
	o.TerminatedDate.Set(&v)
}
// SetTerminatedDateNil sets the value for TerminatedDate to be an explicit nil
func (o *UserInfo) SetTerminatedDateNil() {
	o.TerminatedDate.Set(nil)
}

// UnsetTerminatedDate ensures that no value is present for TerminatedDate, not even an explicit nil
func (o *UserInfo) UnsetTerminatedDate() {
	o.TerminatedDate.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserInfo) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserInfo) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *UserInfo) IsTitleSet() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *UserInfo) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *UserInfo) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *UserInfo) UnsetTitle() {
	o.Title.Unset()
}

// GetWorkFromDate returns the WorkFromDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserInfo) GetWorkFromDate() time.Time {
	if o == nil || IsNil(o.WorkFromDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.WorkFromDate.Get()
}

// GetWorkFromDateOk returns a tuple with the WorkFromDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserInfo) GetWorkFromDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkFromDate.Get(), o.WorkFromDate.IsSet()
}

// HasWorkFromDate returns a boolean if a field has been set.
func (o *UserInfo) IsWorkFromDateSet() bool {
	if o != nil && o.WorkFromDate.IsSet() {
		return true
	}

	return false
}

// SetWorkFromDate gets a reference to the given NullableTime and assigns it to the WorkFromDate field.
func (o *UserInfo) SetWorkFromDate(v time.Time) {
	o.WorkFromDate.Set(&v)
}
// SetWorkFromDateNil sets the value for WorkFromDate to be an explicit nil
func (o *UserInfo) SetWorkFromDateNil() {
	o.WorkFromDate.Set(nil)
}

// UnsetWorkFromDate ensures that no value is present for WorkFromDate, not even an explicit nil
func (o *UserInfo) UnsetWorkFromDate() {
	o.WorkFromDate.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserInfo) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserInfo) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *UserInfo) IsEmailSet() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *UserInfo) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *UserInfo) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *UserInfo) UnsetEmail() {
	o.Email.Unset()
}

// GetContacts returns the Contacts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserInfo) GetContacts() string {
	if o == nil || IsNil(o.Contacts.Get()) {
		var ret string
		return ret
	}
	return *o.Contacts.Get()
}

// GetContactsOk returns a tuple with the Contacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserInfo) GetContactsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Contacts.Get(), o.Contacts.IsSet()
}

// HasContacts returns a boolean if a field has been set.
func (o *UserInfo) IsContactsSet() bool {
	if o != nil && o.Contacts.IsSet() {
		return true
	}

	return false
}

// SetContacts gets a reference to the given NullableString and assigns it to the Contacts field.
func (o *UserInfo) SetContacts(v string) {
	o.Contacts.Set(&v)
}
// SetContactsNil sets the value for Contacts to be an explicit nil
func (o *UserInfo) SetContactsNil() {
	o.Contacts.Set(nil)
}

// UnsetContacts ensures that no value is present for Contacts, not even an explicit nil
func (o *UserInfo) UnsetContacts() {
	o.Contacts.Unset()
}

// GetContactsList returns the ContactsList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserInfo) GetContactsList() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ContactsList
}

// GetContactsListOk returns a tuple with the ContactsList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserInfo) GetContactsListOk() ([]string, bool) {
	if o == nil || IsNil(o.ContactsList) {
		return nil, false
	}
	return o.ContactsList, true
}

// HasContactsList returns a boolean if a field has been set.
func (o *UserInfo) IsContactsListSet() bool {
	if o != nil && !IsNil(o.ContactsList) {
		return true
	}

	return false
}

// SetContactsList gets a reference to the given []string and assigns it to the ContactsList field.
func (o *UserInfo) SetContactsList(v []string) {
	o.ContactsList = v
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserInfo) GetLocation() string {
	if o == nil || IsNil(o.Location.Get()) {
		var ret string
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserInfo) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *UserInfo) IsLocationSet() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableString and assigns it to the Location field.
func (o *UserInfo) SetLocation(v string) {
	o.Location.Set(&v)
}
// SetLocationNil sets the value for Location to be an explicit nil
func (o *UserInfo) SetLocationNil() {
	o.Location.Set(nil)
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *UserInfo) UnsetLocation() {
	o.Location.Unset()
}

// GetNotes returns the Notes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserInfo) GetNotes() string {
	if o == nil || IsNil(o.Notes.Get()) {
		var ret string
		return ret
	}
	return *o.Notes.Get()
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserInfo) GetNotesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Notes.Get(), o.Notes.IsSet()
}

// HasNotes returns a boolean if a field has been set.
func (o *UserInfo) IsNotesSet() bool {
	if o != nil && o.Notes.IsSet() {
		return true
	}

	return false
}

// SetNotes gets a reference to the given NullableString and assigns it to the Notes field.
func (o *UserInfo) SetNotes(v string) {
	o.Notes.Set(&v)
}
// SetNotesNil sets the value for Notes to be an explicit nil
func (o *UserInfo) SetNotesNil() {
	o.Notes.Set(nil)
}

// UnsetNotes ensures that no value is present for Notes, not even an explicit nil
func (o *UserInfo) UnsetNotes() {
	o.Notes.Unset()
}

// GetRemoved returns the Removed field value if set, zero value otherwise.
func (o *UserInfo) GetRemoved() bool {
	if o == nil || IsNil(o.Removed) {
		var ret bool
		return ret
	}
	return *o.Removed
}

// GetRemovedOk returns a tuple with the Removed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInfo) GetRemovedOk() (*bool, bool) {
	if o == nil || IsNil(o.Removed) {
		return nil, false
	}
	return o.Removed, true
}

// HasRemoved returns a boolean if a field has been set.
func (o *UserInfo) IsRemovedSet() bool {
	if o != nil && !IsNil(o.Removed) {
		return true
	}

	return false
}

// SetRemoved gets a reference to the given bool and assigns it to the Removed field.
func (o *UserInfo) SetRemoved(v bool) {
	o.Removed = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *UserInfo) GetLastModified() time.Time {
	if o == nil || IsNil(o.LastModified) {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInfo) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *UserInfo) IsLastModifiedSet() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *UserInfo) SetLastModified(v time.Time) {
	o.LastModified = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *UserInfo) GetTenantId() int32 {
	if o == nil || IsNil(o.TenantId) {
		var ret int32
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInfo) GetTenantIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *UserInfo) IsTenantIdSet() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given int32 and assigns it to the TenantId field.
func (o *UserInfo) SetTenantId(v int32) {
	o.TenantId = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *UserInfo) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInfo) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *UserInfo) IsIsActiveSet() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *UserInfo) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetCultureName returns the CultureName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserInfo) GetCultureName() string {
	if o == nil || IsNil(o.CultureName.Get()) {
		var ret string
		return ret
	}
	return *o.CultureName.Get()
}

// GetCultureNameOk returns a tuple with the CultureName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserInfo) GetCultureNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CultureName.Get(), o.CultureName.IsSet()
}

// HasCultureName returns a boolean if a field has been set.
func (o *UserInfo) IsCultureNameSet() bool {
	if o != nil && o.CultureName.IsSet() {
		return true
	}

	return false
}

// SetCultureName gets a reference to the given NullableString and assigns it to the CultureName field.
func (o *UserInfo) SetCultureName(v string) {
	o.CultureName.Set(&v)
}
// SetCultureNameNil sets the value for CultureName to be an explicit nil
func (o *UserInfo) SetCultureNameNil() {
	o.CultureName.Set(nil)
}

// UnsetCultureName ensures that no value is present for CultureName, not even an explicit nil
func (o *UserInfo) UnsetCultureName() {
	o.CultureName.Unset()
}

// GetMobilePhone returns the MobilePhone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserInfo) GetMobilePhone() string {
	if o == nil || IsNil(o.MobilePhone.Get()) {
		var ret string
		return ret
	}
	return *o.MobilePhone.Get()
}

// GetMobilePhoneOk returns a tuple with the MobilePhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserInfo) GetMobilePhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MobilePhone.Get(), o.MobilePhone.IsSet()
}

// HasMobilePhone returns a boolean if a field has been set.
func (o *UserInfo) IsMobilePhoneSet() bool {
	if o != nil && o.MobilePhone.IsSet() {
		return true
	}

	return false
}

// SetMobilePhone gets a reference to the given NullableString and assigns it to the MobilePhone field.
func (o *UserInfo) SetMobilePhone(v string) {
	o.MobilePhone.Set(&v)
}
// SetMobilePhoneNil sets the value for MobilePhone to be an explicit nil
func (o *UserInfo) SetMobilePhoneNil() {
	o.MobilePhone.Set(nil)
}

// UnsetMobilePhone ensures that no value is present for MobilePhone, not even an explicit nil
func (o *UserInfo) UnsetMobilePhone() {
	o.MobilePhone.Unset()
}

// GetMobilePhoneActivationStatus returns the MobilePhoneActivationStatus field value if set, zero value otherwise.
func (o *UserInfo) GetMobilePhoneActivationStatus() MobilePhoneActivationStatus {
	if o == nil || IsNil(o.MobilePhoneActivationStatus) {
		var ret MobilePhoneActivationStatus
		return ret
	}
	return *o.MobilePhoneActivationStatus
}

// GetMobilePhoneActivationStatusOk returns a tuple with the MobilePhoneActivationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInfo) GetMobilePhoneActivationStatusOk() (*MobilePhoneActivationStatus, bool) {
	if o == nil || IsNil(o.MobilePhoneActivationStatus) {
		return nil, false
	}
	return o.MobilePhoneActivationStatus, true
}

// HasMobilePhoneActivationStatus returns a boolean if a field has been set.
func (o *UserInfo) IsMobilePhoneActivationStatusSet() bool {
	if o != nil && !IsNil(o.MobilePhoneActivationStatus) {
		return true
	}

	return false
}

// SetMobilePhoneActivationStatus gets a reference to the given MobilePhoneActivationStatus and assigns it to the MobilePhoneActivationStatus field.
func (o *UserInfo) SetMobilePhoneActivationStatus(v MobilePhoneActivationStatus) {
	o.MobilePhoneActivationStatus = &v
}

// GetSid returns the Sid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserInfo) GetSid() string {
	if o == nil || IsNil(o.Sid.Get()) {
		var ret string
		return ret
	}
	return *o.Sid.Get()
}

// GetSidOk returns a tuple with the Sid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserInfo) GetSidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sid.Get(), o.Sid.IsSet()
}

// HasSid returns a boolean if a field has been set.
func (o *UserInfo) IsSidSet() bool {
	if o != nil && o.Sid.IsSet() {
		return true
	}

	return false
}

// SetSid gets a reference to the given NullableString and assigns it to the Sid field.
func (o *UserInfo) SetSid(v string) {
	o.Sid.Set(&v)
}
// SetSidNil sets the value for Sid to be an explicit nil
func (o *UserInfo) SetSidNil() {
	o.Sid.Set(nil)
}

// UnsetSid ensures that no value is present for Sid, not even an explicit nil
func (o *UserInfo) UnsetSid() {
	o.Sid.Unset()
}

// GetLdapQouta returns the LdapQouta field value if set, zero value otherwise.
func (o *UserInfo) GetLdapQouta() int64 {
	if o == nil || IsNil(o.LdapQouta) {
		var ret int64
		return ret
	}
	return *o.LdapQouta
}

// GetLdapQoutaOk returns a tuple with the LdapQouta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInfo) GetLdapQoutaOk() (*int64, bool) {
	if o == nil || IsNil(o.LdapQouta) {
		return nil, false
	}
	return o.LdapQouta, true
}

// HasLdapQouta returns a boolean if a field has been set.
func (o *UserInfo) IsLdapQoutaSet() bool {
	if o != nil && !IsNil(o.LdapQouta) {
		return true
	}

	return false
}

// SetLdapQouta gets a reference to the given int64 and assigns it to the LdapQouta field.
func (o *UserInfo) SetLdapQouta(v int64) {
	o.LdapQouta = &v
}

// GetSsoNameId returns the SsoNameId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserInfo) GetSsoNameId() string {
	if o == nil || IsNil(o.SsoNameId.Get()) {
		var ret string
		return ret
	}
	return *o.SsoNameId.Get()
}

// GetSsoNameIdOk returns a tuple with the SsoNameId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserInfo) GetSsoNameIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SsoNameId.Get(), o.SsoNameId.IsSet()
}

// HasSsoNameId returns a boolean if a field has been set.
func (o *UserInfo) IsSsoNameIdSet() bool {
	if o != nil && o.SsoNameId.IsSet() {
		return true
	}

	return false
}

// SetSsoNameId gets a reference to the given NullableString and assigns it to the SsoNameId field.
func (o *UserInfo) SetSsoNameId(v string) {
	o.SsoNameId.Set(&v)
}
// SetSsoNameIdNil sets the value for SsoNameId to be an explicit nil
func (o *UserInfo) SetSsoNameIdNil() {
	o.SsoNameId.Set(nil)
}

// UnsetSsoNameId ensures that no value is present for SsoNameId, not even an explicit nil
func (o *UserInfo) UnsetSsoNameId() {
	o.SsoNameId.Unset()
}

// GetSsoSessionId returns the SsoSessionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserInfo) GetSsoSessionId() string {
	if o == nil || IsNil(o.SsoSessionId.Get()) {
		var ret string
		return ret
	}
	return *o.SsoSessionId.Get()
}

// GetSsoSessionIdOk returns a tuple with the SsoSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserInfo) GetSsoSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SsoSessionId.Get(), o.SsoSessionId.IsSet()
}

// HasSsoSessionId returns a boolean if a field has been set.
func (o *UserInfo) IsSsoSessionIdSet() bool {
	if o != nil && o.SsoSessionId.IsSet() {
		return true
	}

	return false
}

// SetSsoSessionId gets a reference to the given NullableString and assigns it to the SsoSessionId field.
func (o *UserInfo) SetSsoSessionId(v string) {
	o.SsoSessionId.Set(&v)
}
// SetSsoSessionIdNil sets the value for SsoSessionId to be an explicit nil
func (o *UserInfo) SetSsoSessionIdNil() {
	o.SsoSessionId.Set(nil)
}

// UnsetSsoSessionId ensures that no value is present for SsoSessionId, not even an explicit nil
func (o *UserInfo) UnsetSsoSessionId() {
	o.SsoSessionId.Unset()
}

// GetCreateDate returns the CreateDate field value if set, zero value otherwise.
func (o *UserInfo) GetCreateDate() time.Time {
	if o == nil || IsNil(o.CreateDate) {
		var ret time.Time
		return ret
	}
	return *o.CreateDate
}

// GetCreateDateOk returns a tuple with the CreateDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInfo) GetCreateDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreateDate) {
		return nil, false
	}
	return o.CreateDate, true
}

// HasCreateDate returns a boolean if a field has been set.
func (o *UserInfo) IsCreateDateSet() bool {
	if o != nil && !IsNil(o.CreateDate) {
		return true
	}

	return false
}

// SetCreateDate gets a reference to the given time.Time and assigns it to the CreateDate field.
func (o *UserInfo) SetCreateDate(v time.Time) {
	o.CreateDate = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserInfo) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserInfo) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *UserInfo) IsCreatedBySet() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableString and assigns it to the CreatedBy field.
func (o *UserInfo) SetCreatedBy(v string) {
	o.CreatedBy.Set(&v)
}
// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *UserInfo) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *UserInfo) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetSpam returns the Spam field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserInfo) GetSpam() bool {
	if o == nil || IsNil(o.Spam.Get()) {
		var ret bool
		return ret
	}
	return *o.Spam.Get()
}

// GetSpamOk returns a tuple with the Spam field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserInfo) GetSpamOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Spam.Get(), o.Spam.IsSet()
}

// HasSpam returns a boolean if a field has been set.
func (o *UserInfo) IsSpamSet() bool {
	if o != nil && o.Spam.IsSet() {
		return true
	}

	return false
}

// SetSpam gets a reference to the given NullableBool and assigns it to the Spam field.
func (o *UserInfo) SetSpam(v bool) {
	o.Spam.Set(&v)
}
// SetSpamNil sets the value for Spam to be an explicit nil
func (o *UserInfo) SetSpamNil() {
	o.Spam.Set(nil)
}

// UnsetSpam ensures that no value is present for Spam, not even an explicit nil
func (o *UserInfo) UnsetSpam() {
	o.Spam.Unset()
}

// GetCheckActivation returns the CheckActivation field value if set, zero value otherwise.
func (o *UserInfo) GetCheckActivation() bool {
	if o == nil || IsNil(o.CheckActivation) {
		var ret bool
		return ret
	}
	return *o.CheckActivation
}

// GetCheckActivationOk returns a tuple with the CheckActivation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInfo) GetCheckActivationOk() (*bool, bool) {
	if o == nil || IsNil(o.CheckActivation) {
		return nil, false
	}
	return o.CheckActivation, true
}

// HasCheckActivation returns a boolean if a field has been set.
func (o *UserInfo) IsCheckActivationSet() bool {
	if o != nil && !IsNil(o.CheckActivation) {
		return true
	}

	return false
}

// SetCheckActivation gets a reference to the given bool and assigns it to the CheckActivation field.
func (o *UserInfo) SetCheckActivation(v bool) {
	o.CheckActivation = &v
}

func (o UserInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
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
	if o.BirthDate.IsSet() {
		toSerialize["birthDate"] = o.BirthDate.Get()
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
	if o.TerminatedDate.IsSet() {
		toSerialize["terminatedDate"] = o.TerminatedDate.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.WorkFromDate.IsSet() {
		toSerialize["workFromDate"] = o.WorkFromDate.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.Contacts.IsSet() {
		toSerialize["contacts"] = o.Contacts.Get()
	}
	if o.ContactsList != nil {
		toSerialize["contactsList"] = o.ContactsList
	}
	if o.Location.IsSet() {
		toSerialize["location"] = o.Location.Get()
	}
	if o.Notes.IsSet() {
		toSerialize["notes"] = o.Notes.Get()
	}
	if !IsNil(o.Removed) {
		toSerialize["removed"] = o.Removed
	}
	if !IsNil(o.LastModified) {
		toSerialize["lastModified"] = o.LastModified
	}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !IsNil(o.IsActive) {
		toSerialize["isActive"] = o.IsActive
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
	if o.Sid.IsSet() {
		toSerialize["sid"] = o.Sid.Get()
	}
	if !IsNil(o.LdapQouta) {
		toSerialize["ldapQouta"] = o.LdapQouta
	}
	if o.SsoNameId.IsSet() {
		toSerialize["ssoNameId"] = o.SsoNameId.Get()
	}
	if o.SsoSessionId.IsSet() {
		toSerialize["ssoSessionId"] = o.SsoSessionId.Get()
	}
	if !IsNil(o.CreateDate) {
		toSerialize["createDate"] = o.CreateDate
	}
	if o.CreatedBy.IsSet() {
		toSerialize["createdBy"] = o.CreatedBy.Get()
	}
	if o.Spam.IsSet() {
		toSerialize["spam"] = o.Spam.Get()
	}
	if !IsNil(o.CheckActivation) {
		toSerialize["checkActivation"] = o.CheckActivation
	}
	return toSerialize, nil
}

type NullableUserInfo struct {
	value *UserInfo
	isSet bool
}

func (v NullableUserInfo) Get() *UserInfo {
	return v.value
}

func (v *NullableUserInfo) Set(val *UserInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUserInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUserInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserInfo(val *UserInfo) *NullableUserInfo {
	return &NullableUserInfo{value: val, isSet: true}
}

func (v NullableUserInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

