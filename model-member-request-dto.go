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

// checks if the MemberRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MemberRequestDto{}

// MemberRequestDto The user request parameters.
type MemberRequestDto struct {
	// The user password.
	Password NullableString `json:"password,omitempty"`
	// The user password hash.
	PasswordHash NullableString `json:"passwordHash,omitempty"`
	// The user email address.
	Email NullableString `json:"email,omitempty"`
	Type *EmployeeType `json:"type,omitempty"`
	// Specifies if this is a guest or a user.
	IsUser NullableBool `json:"isUser,omitempty"`
	// The user first name.
	FirstName NullableString `json:"firstName,omitempty"`
	// The user last name.
	LastName NullableString `json:"lastName,omitempty"`
	// The list of the user departments IDs.
	Department []string `json:"department,omitempty"`
	// The user title.
	Title NullableString `json:"title,omitempty"`
	// The user location.
	Location NullableString `json:"location,omitempty"`
	Sex *SexEnum `json:"sex,omitempty"`
	Birthday *ApiDateTime `json:"birthday,omitempty"`
	Worksfrom *ApiDateTime `json:"worksfrom,omitempty"`
	// The user comment.
	Comment NullableString `json:"comment,omitempty"`
	// The list of the user contacts.
	Contacts []Contact `json:"contacts,omitempty"`
	// The avatar photo URL.
	Files NullableString `json:"files,omitempty"`
	// Specifies if the user is added via the invitation link or not.
	FromInviteLink *bool `json:"fromInviteLink,omitempty"`
	// The user key.
	Key NullableString `json:"key,omitempty"`
	// The user culture code.
	CultureName NullableString `json:"cultureName,omitempty"`
	// The user target ID.
	Target *string `json:"target,omitempty"`
	// Specifies if tips, updates and offers are allowed to be sent to the user or not.
	Spam NullableBool `json:"spam,omitempty"`
}

// NewMemberRequestDto instantiates a new MemberRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemberRequestDto() *MemberRequestDto {
	this := MemberRequestDto{}
	return &this
}

// NewMemberRequestDtoWithDefaults instantiates a new MemberRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemberRequestDtoWithDefaults() *MemberRequestDto {
	this := MemberRequestDto{}
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemberRequestDto) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemberRequestDto) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *MemberRequestDto) IsPasswordSet() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *MemberRequestDto) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *MemberRequestDto) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *MemberRequestDto) UnsetPassword() {
	o.Password.Unset()
}

// GetPasswordHash returns the PasswordHash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemberRequestDto) GetPasswordHash() string {
	if o == nil || IsNil(o.PasswordHash.Get()) {
		var ret string
		return ret
	}
	return *o.PasswordHash.Get()
}

// GetPasswordHashOk returns a tuple with the PasswordHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemberRequestDto) GetPasswordHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PasswordHash.Get(), o.PasswordHash.IsSet()
}

// HasPasswordHash returns a boolean if a field has been set.
func (o *MemberRequestDto) IsPasswordHashSet() bool {
	if o != nil && o.PasswordHash.IsSet() {
		return true
	}

	return false
}

// SetPasswordHash gets a reference to the given NullableString and assigns it to the PasswordHash field.
func (o *MemberRequestDto) SetPasswordHash(v string) {
	o.PasswordHash.Set(&v)
}
// SetPasswordHashNil sets the value for PasswordHash to be an explicit nil
func (o *MemberRequestDto) SetPasswordHashNil() {
	o.PasswordHash.Set(nil)
}

// UnsetPasswordHash ensures that no value is present for PasswordHash, not even an explicit nil
func (o *MemberRequestDto) UnsetPasswordHash() {
	o.PasswordHash.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemberRequestDto) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemberRequestDto) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *MemberRequestDto) IsEmailSet() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *MemberRequestDto) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *MemberRequestDto) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *MemberRequestDto) UnsetEmail() {
	o.Email.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MemberRequestDto) GetType() EmployeeType {
	if o == nil || IsNil(o.Type) {
		var ret EmployeeType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberRequestDto) GetTypeOk() (*EmployeeType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MemberRequestDto) IsTypeSet() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given EmployeeType and assigns it to the Type field.
func (o *MemberRequestDto) SetType(v EmployeeType) {
	o.Type = &v
}

// GetIsUser returns the IsUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemberRequestDto) GetIsUser() bool {
	if o == nil || IsNil(o.IsUser.Get()) {
		var ret bool
		return ret
	}
	return *o.IsUser.Get()
}

// GetIsUserOk returns a tuple with the IsUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemberRequestDto) GetIsUserOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsUser.Get(), o.IsUser.IsSet()
}

// HasIsUser returns a boolean if a field has been set.
func (o *MemberRequestDto) IsIsUserSet() bool {
	if o != nil && o.IsUser.IsSet() {
		return true
	}

	return false
}

// SetIsUser gets a reference to the given NullableBool and assigns it to the IsUser field.
func (o *MemberRequestDto) SetIsUser(v bool) {
	o.IsUser.Set(&v)
}
// SetIsUserNil sets the value for IsUser to be an explicit nil
func (o *MemberRequestDto) SetIsUserNil() {
	o.IsUser.Set(nil)
}

// UnsetIsUser ensures that no value is present for IsUser, not even an explicit nil
func (o *MemberRequestDto) UnsetIsUser() {
	o.IsUser.Unset()
}

// GetFirstName returns the FirstName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemberRequestDto) GetFirstName() string {
	if o == nil || IsNil(o.FirstName.Get()) {
		var ret string
		return ret
	}
	return *o.FirstName.Get()
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemberRequestDto) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstName.Get(), o.FirstName.IsSet()
}

// HasFirstName returns a boolean if a field has been set.
func (o *MemberRequestDto) IsFirstNameSet() bool {
	if o != nil && o.FirstName.IsSet() {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given NullableString and assigns it to the FirstName field.
func (o *MemberRequestDto) SetFirstName(v string) {
	o.FirstName.Set(&v)
}
// SetFirstNameNil sets the value for FirstName to be an explicit nil
func (o *MemberRequestDto) SetFirstNameNil() {
	o.FirstName.Set(nil)
}

// UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
func (o *MemberRequestDto) UnsetFirstName() {
	o.FirstName.Unset()
}

// GetLastName returns the LastName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemberRequestDto) GetLastName() string {
	if o == nil || IsNil(o.LastName.Get()) {
		var ret string
		return ret
	}
	return *o.LastName.Get()
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemberRequestDto) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastName.Get(), o.LastName.IsSet()
}

// HasLastName returns a boolean if a field has been set.
func (o *MemberRequestDto) IsLastNameSet() bool {
	if o != nil && o.LastName.IsSet() {
		return true
	}

	return false
}

// SetLastName gets a reference to the given NullableString and assigns it to the LastName field.
func (o *MemberRequestDto) SetLastName(v string) {
	o.LastName.Set(&v)
}
// SetLastNameNil sets the value for LastName to be an explicit nil
func (o *MemberRequestDto) SetLastNameNil() {
	o.LastName.Set(nil)
}

// UnsetLastName ensures that no value is present for LastName, not even an explicit nil
func (o *MemberRequestDto) UnsetLastName() {
	o.LastName.Unset()
}

// GetDepartment returns the Department field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemberRequestDto) GetDepartment() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Department
}

// GetDepartmentOk returns a tuple with the Department field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemberRequestDto) GetDepartmentOk() ([]string, bool) {
	if o == nil || IsNil(o.Department) {
		return nil, false
	}
	return o.Department, true
}

// HasDepartment returns a boolean if a field has been set.
func (o *MemberRequestDto) IsDepartmentSet() bool {
	if o != nil && !IsNil(o.Department) {
		return true
	}

	return false
}

// SetDepartment gets a reference to the given []string and assigns it to the Department field.
func (o *MemberRequestDto) SetDepartment(v []string) {
	o.Department = v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemberRequestDto) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemberRequestDto) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *MemberRequestDto) IsTitleSet() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *MemberRequestDto) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *MemberRequestDto) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *MemberRequestDto) UnsetTitle() {
	o.Title.Unset()
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemberRequestDto) GetLocation() string {
	if o == nil || IsNil(o.Location.Get()) {
		var ret string
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemberRequestDto) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *MemberRequestDto) IsLocationSet() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableString and assigns it to the Location field.
func (o *MemberRequestDto) SetLocation(v string) {
	o.Location.Set(&v)
}
// SetLocationNil sets the value for Location to be an explicit nil
func (o *MemberRequestDto) SetLocationNil() {
	o.Location.Set(nil)
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *MemberRequestDto) UnsetLocation() {
	o.Location.Unset()
}

// GetSex returns the Sex field value if set, zero value otherwise.
func (o *MemberRequestDto) GetSex() SexEnum {
	if o == nil || IsNil(o.Sex) {
		var ret SexEnum
		return ret
	}
	return *o.Sex
}

// GetSexOk returns a tuple with the Sex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberRequestDto) GetSexOk() (*SexEnum, bool) {
	if o == nil || IsNil(o.Sex) {
		return nil, false
	}
	return o.Sex, true
}

// HasSex returns a boolean if a field has been set.
func (o *MemberRequestDto) IsSexSet() bool {
	if o != nil && !IsNil(o.Sex) {
		return true
	}

	return false
}

// SetSex gets a reference to the given SexEnum and assigns it to the Sex field.
func (o *MemberRequestDto) SetSex(v SexEnum) {
	o.Sex = &v
}

// GetBirthday returns the Birthday field value if set, zero value otherwise.
func (o *MemberRequestDto) GetBirthday() ApiDateTime {
	if o == nil || IsNil(o.Birthday) {
		var ret ApiDateTime
		return ret
	}
	return *o.Birthday
}

// GetBirthdayOk returns a tuple with the Birthday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberRequestDto) GetBirthdayOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.Birthday) {
		return nil, false
	}
	return o.Birthday, true
}

// HasBirthday returns a boolean if a field has been set.
func (o *MemberRequestDto) IsBirthdaySet() bool {
	if o != nil && !IsNil(o.Birthday) {
		return true
	}

	return false
}

// SetBirthday gets a reference to the given ApiDateTime and assigns it to the Birthday field.
func (o *MemberRequestDto) SetBirthday(v ApiDateTime) {
	o.Birthday = &v
}

// GetWorksfrom returns the Worksfrom field value if set, zero value otherwise.
func (o *MemberRequestDto) GetWorksfrom() ApiDateTime {
	if o == nil || IsNil(o.Worksfrom) {
		var ret ApiDateTime
		return ret
	}
	return *o.Worksfrom
}

// GetWorksfromOk returns a tuple with the Worksfrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberRequestDto) GetWorksfromOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.Worksfrom) {
		return nil, false
	}
	return o.Worksfrom, true
}

// HasWorksfrom returns a boolean if a field has been set.
func (o *MemberRequestDto) IsWorksfromSet() bool {
	if o != nil && !IsNil(o.Worksfrom) {
		return true
	}

	return false
}

// SetWorksfrom gets a reference to the given ApiDateTime and assigns it to the Worksfrom field.
func (o *MemberRequestDto) SetWorksfrom(v ApiDateTime) {
	o.Worksfrom = &v
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemberRequestDto) GetComment() string {
	if o == nil || IsNil(o.Comment.Get()) {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemberRequestDto) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *MemberRequestDto) IsCommentSet() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *MemberRequestDto) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *MemberRequestDto) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *MemberRequestDto) UnsetComment() {
	o.Comment.Unset()
}

// GetContacts returns the Contacts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemberRequestDto) GetContacts() []Contact {
	if o == nil {
		var ret []Contact
		return ret
	}
	return o.Contacts
}

// GetContactsOk returns a tuple with the Contacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemberRequestDto) GetContactsOk() ([]Contact, bool) {
	if o == nil || IsNil(o.Contacts) {
		return nil, false
	}
	return o.Contacts, true
}

// HasContacts returns a boolean if a field has been set.
func (o *MemberRequestDto) IsContactsSet() bool {
	if o != nil && !IsNil(o.Contacts) {
		return true
	}

	return false
}

// SetContacts gets a reference to the given []Contact and assigns it to the Contacts field.
func (o *MemberRequestDto) SetContacts(v []Contact) {
	o.Contacts = v
}

// GetFiles returns the Files field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemberRequestDto) GetFiles() string {
	if o == nil || IsNil(o.Files.Get()) {
		var ret string
		return ret
	}
	return *o.Files.Get()
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemberRequestDto) GetFilesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Files.Get(), o.Files.IsSet()
}

// HasFiles returns a boolean if a field has been set.
func (o *MemberRequestDto) IsFilesSet() bool {
	if o != nil && o.Files.IsSet() {
		return true
	}

	return false
}

// SetFiles gets a reference to the given NullableString and assigns it to the Files field.
func (o *MemberRequestDto) SetFiles(v string) {
	o.Files.Set(&v)
}
// SetFilesNil sets the value for Files to be an explicit nil
func (o *MemberRequestDto) SetFilesNil() {
	o.Files.Set(nil)
}

// UnsetFiles ensures that no value is present for Files, not even an explicit nil
func (o *MemberRequestDto) UnsetFiles() {
	o.Files.Unset()
}

// GetFromInviteLink returns the FromInviteLink field value if set, zero value otherwise.
func (o *MemberRequestDto) GetFromInviteLink() bool {
	if o == nil || IsNil(o.FromInviteLink) {
		var ret bool
		return ret
	}
	return *o.FromInviteLink
}

// GetFromInviteLinkOk returns a tuple with the FromInviteLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberRequestDto) GetFromInviteLinkOk() (*bool, bool) {
	if o == nil || IsNil(o.FromInviteLink) {
		return nil, false
	}
	return o.FromInviteLink, true
}

// HasFromInviteLink returns a boolean if a field has been set.
func (o *MemberRequestDto) IsFromInviteLinkSet() bool {
	if o != nil && !IsNil(o.FromInviteLink) {
		return true
	}

	return false
}

// SetFromInviteLink gets a reference to the given bool and assigns it to the FromInviteLink field.
func (o *MemberRequestDto) SetFromInviteLink(v bool) {
	o.FromInviteLink = &v
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemberRequestDto) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemberRequestDto) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *MemberRequestDto) IsKeySet() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *MemberRequestDto) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *MemberRequestDto) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *MemberRequestDto) UnsetKey() {
	o.Key.Unset()
}

// GetCultureName returns the CultureName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemberRequestDto) GetCultureName() string {
	if o == nil || IsNil(o.CultureName.Get()) {
		var ret string
		return ret
	}
	return *o.CultureName.Get()
}

// GetCultureNameOk returns a tuple with the CultureName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemberRequestDto) GetCultureNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CultureName.Get(), o.CultureName.IsSet()
}

// HasCultureName returns a boolean if a field has been set.
func (o *MemberRequestDto) IsCultureNameSet() bool {
	if o != nil && o.CultureName.IsSet() {
		return true
	}

	return false
}

// SetCultureName gets a reference to the given NullableString and assigns it to the CultureName field.
func (o *MemberRequestDto) SetCultureName(v string) {
	o.CultureName.Set(&v)
}
// SetCultureNameNil sets the value for CultureName to be an explicit nil
func (o *MemberRequestDto) SetCultureNameNil() {
	o.CultureName.Set(nil)
}

// UnsetCultureName ensures that no value is present for CultureName, not even an explicit nil
func (o *MemberRequestDto) UnsetCultureName() {
	o.CultureName.Unset()
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *MemberRequestDto) GetTarget() string {
	if o == nil || IsNil(o.Target) {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberRequestDto) GetTargetOk() (*string, bool) {
	if o == nil || IsNil(o.Target) {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *MemberRequestDto) IsTargetSet() bool {
	if o != nil && !IsNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *MemberRequestDto) SetTarget(v string) {
	o.Target = &v
}

// GetSpam returns the Spam field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MemberRequestDto) GetSpam() bool {
	if o == nil || IsNil(o.Spam.Get()) {
		var ret bool
		return ret
	}
	return *o.Spam.Get()
}

// GetSpamOk returns a tuple with the Spam field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MemberRequestDto) GetSpamOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Spam.Get(), o.Spam.IsSet()
}

// HasSpam returns a boolean if a field has been set.
func (o *MemberRequestDto) IsSpamSet() bool {
	if o != nil && o.Spam.IsSet() {
		return true
	}

	return false
}

// SetSpam gets a reference to the given NullableBool and assigns it to the Spam field.
func (o *MemberRequestDto) SetSpam(v bool) {
	o.Spam.Set(&v)
}
// SetSpamNil sets the value for Spam to be an explicit nil
func (o *MemberRequestDto) SetSpamNil() {
	o.Spam.Set(nil)
}

// UnsetSpam ensures that no value is present for Spam, not even an explicit nil
func (o *MemberRequestDto) UnsetSpam() {
	o.Spam.Unset()
}

func (o MemberRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MemberRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if o.PasswordHash.IsSet() {
		toSerialize["passwordHash"] = o.PasswordHash.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.IsUser.IsSet() {
		toSerialize["isUser"] = o.IsUser.Get()
	}
	if o.FirstName.IsSet() {
		toSerialize["firstName"] = o.FirstName.Get()
	}
	if o.LastName.IsSet() {
		toSerialize["lastName"] = o.LastName.Get()
	}
	if o.Department != nil {
		toSerialize["department"] = o.Department
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Location.IsSet() {
		toSerialize["location"] = o.Location.Get()
	}
	if !IsNil(o.Sex) {
		toSerialize["sex"] = o.Sex
	}
	if !IsNil(o.Birthday) {
		toSerialize["birthday"] = o.Birthday
	}
	if !IsNil(o.Worksfrom) {
		toSerialize["worksfrom"] = o.Worksfrom
	}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	if o.Contacts != nil {
		toSerialize["contacts"] = o.Contacts
	}
	if o.Files.IsSet() {
		toSerialize["files"] = o.Files.Get()
	}
	if !IsNil(o.FromInviteLink) {
		toSerialize["fromInviteLink"] = o.FromInviteLink
	}
	if o.Key.IsSet() {
		toSerialize["key"] = o.Key.Get()
	}
	if o.CultureName.IsSet() {
		toSerialize["cultureName"] = o.CultureName.Get()
	}
	if !IsNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	if o.Spam.IsSet() {
		toSerialize["spam"] = o.Spam.Get()
	}
	return toSerialize, nil
}

type NullableMemberRequestDto struct {
	value *MemberRequestDto
	isSet bool
}

func (v NullableMemberRequestDto) Get() *MemberRequestDto {
	return v.value
}

func (v *NullableMemberRequestDto) Set(val *MemberRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableMemberRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableMemberRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemberRequestDto(val *MemberRequestDto) *NullableMemberRequestDto {
	return &NullableMemberRequestDto{value: val, isSet: true}
}

func (v NullableMemberRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemberRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

