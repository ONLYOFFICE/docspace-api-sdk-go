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

// checks if the UpdateMemberRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateMemberRequestDto{}

// UpdateMemberRequestDto The request parameters for updating the user information.
type UpdateMemberRequestDto struct {
	// The user ID.
	UserId NullableString `json:"userId,omitempty"`
	// Specifies whether to disable a user or not.
	Disable NullableBool `json:"disable,omitempty"`
	// The user email address.
	Email NullableString `json:"email,omitempty"`
	// Specifies if this is a guest or a user.
	IsUser NullableBool `json:"isUser,omitempty"`
	// The user first name.
	FirstName NullableString `json:"firstName,omitempty"`
	// The user last name.
	LastName NullableString `json:"lastName,omitempty"`
	// The list of the user departments.
	Department []string `json:"department,omitempty"`
	// The user location.
	Location NullableString `json:"location,omitempty"`
	// The user comment.
	Comment NullableString `json:"comment,omitempty"`
	// The list of the user contacts.
	Contacts []Contact `json:"contacts,omitempty"`
	// The user avatar photo URL.
	Files NullableString `json:"files,omitempty"`
	// Specifies if tips, updates and offers are allowed to be sent to the user or not.
	Spam NullableBool `json:"spam,omitempty"`
}

// NewUpdateMemberRequestDto instantiates a new UpdateMemberRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMemberRequestDto() *UpdateMemberRequestDto {
	this := UpdateMemberRequestDto{}
	return &this
}

// NewUpdateMemberRequestDtoWithDefaults instantiates a new UpdateMemberRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMemberRequestDtoWithDefaults() *UpdateMemberRequestDto {
	this := UpdateMemberRequestDto{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateMemberRequestDto) GetUserId() string {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateMemberRequestDto) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *UpdateMemberRequestDto) IsUserIdSet() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableString and assigns it to the UserId field.
func (o *UpdateMemberRequestDto) SetUserId(v string) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *UpdateMemberRequestDto) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *UpdateMemberRequestDto) UnsetUserId() {
	o.UserId.Unset()
}

// GetDisable returns the Disable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateMemberRequestDto) GetDisable() bool {
	if o == nil || IsNil(o.Disable.Get()) {
		var ret bool
		return ret
	}
	return *o.Disable.Get()
}

// GetDisableOk returns a tuple with the Disable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateMemberRequestDto) GetDisableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Disable.Get(), o.Disable.IsSet()
}

// HasDisable returns a boolean if a field has been set.
func (o *UpdateMemberRequestDto) IsDisableSet() bool {
	if o != nil && o.Disable.IsSet() {
		return true
	}

	return false
}

// SetDisable gets a reference to the given NullableBool and assigns it to the Disable field.
func (o *UpdateMemberRequestDto) SetDisable(v bool) {
	o.Disable.Set(&v)
}
// SetDisableNil sets the value for Disable to be an explicit nil
func (o *UpdateMemberRequestDto) SetDisableNil() {
	o.Disable.Set(nil)
}

// UnsetDisable ensures that no value is present for Disable, not even an explicit nil
func (o *UpdateMemberRequestDto) UnsetDisable() {
	o.Disable.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateMemberRequestDto) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateMemberRequestDto) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *UpdateMemberRequestDto) IsEmailSet() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *UpdateMemberRequestDto) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *UpdateMemberRequestDto) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *UpdateMemberRequestDto) UnsetEmail() {
	o.Email.Unset()
}

// GetIsUser returns the IsUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateMemberRequestDto) GetIsUser() bool {
	if o == nil || IsNil(o.IsUser.Get()) {
		var ret bool
		return ret
	}
	return *o.IsUser.Get()
}

// GetIsUserOk returns a tuple with the IsUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateMemberRequestDto) GetIsUserOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsUser.Get(), o.IsUser.IsSet()
}

// HasIsUser returns a boolean if a field has been set.
func (o *UpdateMemberRequestDto) IsIsUserSet() bool {
	if o != nil && o.IsUser.IsSet() {
		return true
	}

	return false
}

// SetIsUser gets a reference to the given NullableBool and assigns it to the IsUser field.
func (o *UpdateMemberRequestDto) SetIsUser(v bool) {
	o.IsUser.Set(&v)
}
// SetIsUserNil sets the value for IsUser to be an explicit nil
func (o *UpdateMemberRequestDto) SetIsUserNil() {
	o.IsUser.Set(nil)
}

// UnsetIsUser ensures that no value is present for IsUser, not even an explicit nil
func (o *UpdateMemberRequestDto) UnsetIsUser() {
	o.IsUser.Unset()
}

// GetFirstName returns the FirstName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateMemberRequestDto) GetFirstName() string {
	if o == nil || IsNil(o.FirstName.Get()) {
		var ret string
		return ret
	}
	return *o.FirstName.Get()
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateMemberRequestDto) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstName.Get(), o.FirstName.IsSet()
}

// HasFirstName returns a boolean if a field has been set.
func (o *UpdateMemberRequestDto) IsFirstNameSet() bool {
	if o != nil && o.FirstName.IsSet() {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given NullableString and assigns it to the FirstName field.
func (o *UpdateMemberRequestDto) SetFirstName(v string) {
	o.FirstName.Set(&v)
}
// SetFirstNameNil sets the value for FirstName to be an explicit nil
func (o *UpdateMemberRequestDto) SetFirstNameNil() {
	o.FirstName.Set(nil)
}

// UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
func (o *UpdateMemberRequestDto) UnsetFirstName() {
	o.FirstName.Unset()
}

// GetLastName returns the LastName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateMemberRequestDto) GetLastName() string {
	if o == nil || IsNil(o.LastName.Get()) {
		var ret string
		return ret
	}
	return *o.LastName.Get()
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateMemberRequestDto) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastName.Get(), o.LastName.IsSet()
}

// HasLastName returns a boolean if a field has been set.
func (o *UpdateMemberRequestDto) IsLastNameSet() bool {
	if o != nil && o.LastName.IsSet() {
		return true
	}

	return false
}

// SetLastName gets a reference to the given NullableString and assigns it to the LastName field.
func (o *UpdateMemberRequestDto) SetLastName(v string) {
	o.LastName.Set(&v)
}
// SetLastNameNil sets the value for LastName to be an explicit nil
func (o *UpdateMemberRequestDto) SetLastNameNil() {
	o.LastName.Set(nil)
}

// UnsetLastName ensures that no value is present for LastName, not even an explicit nil
func (o *UpdateMemberRequestDto) UnsetLastName() {
	o.LastName.Unset()
}

// GetDepartment returns the Department field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateMemberRequestDto) GetDepartment() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Department
}

// GetDepartmentOk returns a tuple with the Department field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateMemberRequestDto) GetDepartmentOk() ([]string, bool) {
	if o == nil || IsNil(o.Department) {
		return nil, false
	}
	return o.Department, true
}

// HasDepartment returns a boolean if a field has been set.
func (o *UpdateMemberRequestDto) IsDepartmentSet() bool {
	if o != nil && !IsNil(o.Department) {
		return true
	}

	return false
}

// SetDepartment gets a reference to the given []string and assigns it to the Department field.
func (o *UpdateMemberRequestDto) SetDepartment(v []string) {
	o.Department = v
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateMemberRequestDto) GetLocation() string {
	if o == nil || IsNil(o.Location.Get()) {
		var ret string
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateMemberRequestDto) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *UpdateMemberRequestDto) IsLocationSet() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableString and assigns it to the Location field.
func (o *UpdateMemberRequestDto) SetLocation(v string) {
	o.Location.Set(&v)
}
// SetLocationNil sets the value for Location to be an explicit nil
func (o *UpdateMemberRequestDto) SetLocationNil() {
	o.Location.Set(nil)
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *UpdateMemberRequestDto) UnsetLocation() {
	o.Location.Unset()
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateMemberRequestDto) GetComment() string {
	if o == nil || IsNil(o.Comment.Get()) {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateMemberRequestDto) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *UpdateMemberRequestDto) IsCommentSet() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *UpdateMemberRequestDto) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *UpdateMemberRequestDto) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *UpdateMemberRequestDto) UnsetComment() {
	o.Comment.Unset()
}

// GetContacts returns the Contacts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateMemberRequestDto) GetContacts() []Contact {
	if o == nil {
		var ret []Contact
		return ret
	}
	return o.Contacts
}

// GetContactsOk returns a tuple with the Contacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateMemberRequestDto) GetContactsOk() ([]Contact, bool) {
	if o == nil || IsNil(o.Contacts) {
		return nil, false
	}
	return o.Contacts, true
}

// HasContacts returns a boolean if a field has been set.
func (o *UpdateMemberRequestDto) IsContactsSet() bool {
	if o != nil && !IsNil(o.Contacts) {
		return true
	}

	return false
}

// SetContacts gets a reference to the given []Contact and assigns it to the Contacts field.
func (o *UpdateMemberRequestDto) SetContacts(v []Contact) {
	o.Contacts = v
}

// GetFiles returns the Files field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateMemberRequestDto) GetFiles() string {
	if o == nil || IsNil(o.Files.Get()) {
		var ret string
		return ret
	}
	return *o.Files.Get()
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateMemberRequestDto) GetFilesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Files.Get(), o.Files.IsSet()
}

// HasFiles returns a boolean if a field has been set.
func (o *UpdateMemberRequestDto) IsFilesSet() bool {
	if o != nil && o.Files.IsSet() {
		return true
	}

	return false
}

// SetFiles gets a reference to the given NullableString and assigns it to the Files field.
func (o *UpdateMemberRequestDto) SetFiles(v string) {
	o.Files.Set(&v)
}
// SetFilesNil sets the value for Files to be an explicit nil
func (o *UpdateMemberRequestDto) SetFilesNil() {
	o.Files.Set(nil)
}

// UnsetFiles ensures that no value is present for Files, not even an explicit nil
func (o *UpdateMemberRequestDto) UnsetFiles() {
	o.Files.Unset()
}

// GetSpam returns the Spam field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateMemberRequestDto) GetSpam() bool {
	if o == nil || IsNil(o.Spam.Get()) {
		var ret bool
		return ret
	}
	return *o.Spam.Get()
}

// GetSpamOk returns a tuple with the Spam field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateMemberRequestDto) GetSpamOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Spam.Get(), o.Spam.IsSet()
}

// HasSpam returns a boolean if a field has been set.
func (o *UpdateMemberRequestDto) IsSpamSet() bool {
	if o != nil && o.Spam.IsSet() {
		return true
	}

	return false
}

// SetSpam gets a reference to the given NullableBool and assigns it to the Spam field.
func (o *UpdateMemberRequestDto) SetSpam(v bool) {
	o.Spam.Set(&v)
}
// SetSpamNil sets the value for Spam to be an explicit nil
func (o *UpdateMemberRequestDto) SetSpamNil() {
	o.Spam.Set(nil)
}

// UnsetSpam ensures that no value is present for Spam, not even an explicit nil
func (o *UpdateMemberRequestDto) UnsetSpam() {
	o.Spam.Unset()
}

func (o UpdateMemberRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateMemberRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.UserId.IsSet() {
		toSerialize["userId"] = o.UserId.Get()
	}
	if o.Disable.IsSet() {
		toSerialize["disable"] = o.Disable.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
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
	if o.Location.IsSet() {
		toSerialize["location"] = o.Location.Get()
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
	if o.Spam.IsSet() {
		toSerialize["spam"] = o.Spam.Get()
	}
	return toSerialize, nil
}

type NullableUpdateMemberRequestDto struct {
	value *UpdateMemberRequestDto
	isSet bool
}

func (v NullableUpdateMemberRequestDto) Get() *UpdateMemberRequestDto {
	return v.value
}

func (v *NullableUpdateMemberRequestDto) Set(val *UpdateMemberRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMemberRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMemberRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMemberRequestDto(val *UpdateMemberRequestDto) *NullableUpdateMemberRequestDto {
	return &NullableUpdateMemberRequestDto{value: val, isSet: true}
}

func (v NullableUpdateMemberRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMemberRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

