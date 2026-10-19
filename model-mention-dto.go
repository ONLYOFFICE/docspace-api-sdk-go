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

// checks if the MentionDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MentionDto{}

// MentionDto A portal member the editor can offer when the author types a mention: who they are, where the notification goes  and how to show them in the suggestion list.
type MentionDto struct {
	// The account itself, as the portal stores it.
	User *PortalUserDto `json:"user,omitempty"`
	// Where a mention notification for this user is delivered.
	Email NullableString `json:"email,omitempty"`
	// The account id as text, the same value the account object carries; it is what identifies the user in a sharing  request built from this list.
	Id NullableString `json:"id,omitempty"`
	// An absolute address of the medium-sized avatar. A generated default avatar is reported when the user never  uploaded one, so the field is never empty.
	Image NullableString `json:"image,omitempty"`
	// Not filled in by the operations that return this list: it always comes back false. Whether a user can already  open the document has to be read from the sharing settings of the file.
	HasAccess *bool `json:"hasAccess,omitempty"`
	// The name to display, assembled the way the portal is configured to show names.
	Name NullableString `json:"name,omitempty"`
}

// NewMentionDto instantiates a new MentionDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMentionDto() *MentionDto {
	this := MentionDto{}
	return &this
}

// NewMentionDtoWithDefaults instantiates a new MentionDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMentionDtoWithDefaults() *MentionDto {
	this := MentionDto{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *MentionDto) GetUser() PortalUserDto {
	if o == nil || IsNil(o.User) {
		var ret PortalUserDto
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MentionDto) GetUserOk() (*PortalUserDto, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *MentionDto) IsUserSet() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given PortalUserDto and assigns it to the User field.
func (o *MentionDto) SetUser(v PortalUserDto) {
	o.User = &v
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MentionDto) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MentionDto) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *MentionDto) IsEmailSet() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *MentionDto) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *MentionDto) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *MentionDto) UnsetEmail() {
	o.Email.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MentionDto) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MentionDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *MentionDto) IsIdSet() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *MentionDto) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *MentionDto) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *MentionDto) UnsetId() {
	o.Id.Unset()
}

// GetImage returns the Image field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MentionDto) GetImage() string {
	if o == nil || IsNil(o.Image.Get()) {
		var ret string
		return ret
	}
	return *o.Image.Get()
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MentionDto) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Image.Get(), o.Image.IsSet()
}

// HasImage returns a boolean if a field has been set.
func (o *MentionDto) IsImageSet() bool {
	if o != nil && o.Image.IsSet() {
		return true
	}

	return false
}

// SetImage gets a reference to the given NullableString and assigns it to the Image field.
func (o *MentionDto) SetImage(v string) {
	o.Image.Set(&v)
}
// SetImageNil sets the value for Image to be an explicit nil
func (o *MentionDto) SetImageNil() {
	o.Image.Set(nil)
}

// UnsetImage ensures that no value is present for Image, not even an explicit nil
func (o *MentionDto) UnsetImage() {
	o.Image.Unset()
}

// GetHasAccess returns the HasAccess field value if set, zero value otherwise.
func (o *MentionDto) GetHasAccess() bool {
	if o == nil || IsNil(o.HasAccess) {
		var ret bool
		return ret
	}
	return *o.HasAccess
}

// GetHasAccessOk returns a tuple with the HasAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MentionDto) GetHasAccessOk() (*bool, bool) {
	if o == nil || IsNil(o.HasAccess) {
		return nil, false
	}
	return o.HasAccess, true
}

// HasHasAccess returns a boolean if a field has been set.
func (o *MentionDto) IsHasAccessSet() bool {
	if o != nil && !IsNil(o.HasAccess) {
		return true
	}

	return false
}

// SetHasAccess gets a reference to the given bool and assigns it to the HasAccess field.
func (o *MentionDto) SetHasAccess(v bool) {
	o.HasAccess = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MentionDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MentionDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *MentionDto) IsNameSet() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *MentionDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *MentionDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *MentionDto) UnsetName() {
	o.Name.Unset()
}

func (o MentionDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MentionDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Image.IsSet() {
		toSerialize["image"] = o.Image.Get()
	}
	if !IsNil(o.HasAccess) {
		toSerialize["hasAccess"] = o.HasAccess
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return toSerialize, nil
}

type NullableMentionDto struct {
	value *MentionDto
	isSet bool
}

func (v NullableMentionDto) Get() *MentionDto {
	return v.value
}

func (v *NullableMentionDto) Set(val *MentionDto) {
	v.value = val
	v.isSet = true
}

func (v NullableMentionDto) IsSet() bool {
	return v.isSet
}

func (v *NullableMentionDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMentionDto(val *MentionDto) *NullableMentionDto {
	return &NullableMentionDto{value: val, isSet: true}
}

func (v NullableMentionDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMentionDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

