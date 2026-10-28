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

// checks if the RoomLinkRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoomLinkRequest{}

// RoomLinkRequest The link of a room to create, change or revoke.
type RoomLinkRequest struct {
	// Which link to change, taken from `GET api/2.0/files/rooms/{id}/links`. Leaving it out creates a link, and an  identifier the room does not know creates a link carrying that identifier.
	LinkId *string `json:"linkId,omitempty"`
	// What whoever opens the link may do in the room. The value 0 revokes the link instead of changing it, and the  levels a room accepts depend on its kind.
	Access *FileShare `json:"access,omitempty"`
	// When the link stops working, written with the offset of the portal time zone. A date already past is dropped  silently for an external link and refused for an invitation link, and a date further ahead than the portal  allows is refused as well; leaving it out means the link does not expire.
	ExpirationDate *ApiDateTime `json:"expirationDate,omitempty"`
	// Whether the external link works only for people already signed in to the portal. With it off the link opens  the room for anyone who has the address, subject to the password.
	Internal *bool `json:"internal,omitempty"`
	// The name the link is shown under in the room. An empty value is accepted and the portal names the link itself,  so the answer is what tells the caller the name in use.
	Title NullableString `json:"title,omitempty"`
	// Which kind of link to create: an invitation link makes whoever opens it a member of the room, while an  external link opens the room without an account. It is fixed when the link is created and is ignored on later  changes.
	LinkType *LinkType `json:"linkType,omitempty"`
	// The password an external link asks for before it opens the room. An empty value leaves the link open to anyone  who has the address, and the password is never returned when links are listed.
	Password NullableString `json:"password,omitempty"`
	// Whether people arriving through the link are stopped from downloading and printing what they open. They can  still read the documents in the editor.
	DenyDownload *bool `json:"denyDownload,omitempty"`
	// How many people an invitation link may still let in before it stops working. A value below the number of  people who already used it is refused, and leaving it out puts no ceiling on the link.
	MaxUseCount NullableInt32 `json:"maxUseCount,omitempty"`
	// How many people have already joined through this invitation link. The value is kept by the portal: it is  reported back when links are listed and anything sent here is ignored.
	CurrentUseCount *int32 `json:"currentUseCount,omitempty"`
}

// NewRoomLinkRequest instantiates a new RoomLinkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoomLinkRequest() *RoomLinkRequest {
	this := RoomLinkRequest{}
	return &this
}

// NewRoomLinkRequestWithDefaults instantiates a new RoomLinkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoomLinkRequestWithDefaults() *RoomLinkRequest {
	this := RoomLinkRequest{}
	return &this
}

// GetLinkId returns the LinkId field value if set, zero value otherwise.
func (o *RoomLinkRequest) GetLinkId() string {
	if o == nil || IsNil(o.LinkId) {
		var ret string
		return ret
	}
	return *o.LinkId
}

// GetLinkIdOk returns a tuple with the LinkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomLinkRequest) GetLinkIdOk() (*string, bool) {
	if o == nil || IsNil(o.LinkId) {
		return nil, false
	}
	return o.LinkId, true
}

// HasLinkId returns a boolean if a field has been set.
func (o *RoomLinkRequest) IsLinkIdSet() bool {
	if o != nil && !IsNil(o.LinkId) {
		return true
	}

	return false
}

// SetLinkId gets a reference to the given string and assigns it to the LinkId field.
func (o *RoomLinkRequest) SetLinkId(v string) {
	o.LinkId = &v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *RoomLinkRequest) GetAccess() FileShare {
	if o == nil || IsNil(o.Access) {
		var ret FileShare
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomLinkRequest) GetAccessOk() (*FileShare, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *RoomLinkRequest) IsAccessSet() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given FileShare and assigns it to the Access field.
func (o *RoomLinkRequest) SetAccess(v FileShare) {
	o.Access = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *RoomLinkRequest) GetExpirationDate() ApiDateTime {
	if o == nil || IsNil(o.ExpirationDate) {
		var ret ApiDateTime
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomLinkRequest) GetExpirationDateOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *RoomLinkRequest) IsExpirationDateSet() bool {
	if o != nil && !IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given ApiDateTime and assigns it to the ExpirationDate field.
func (o *RoomLinkRequest) SetExpirationDate(v ApiDateTime) {
	o.ExpirationDate = &v
}

// GetInternal returns the Internal field value if set, zero value otherwise.
func (o *RoomLinkRequest) GetInternal() bool {
	if o == nil || IsNil(o.Internal) {
		var ret bool
		return ret
	}
	return *o.Internal
}

// GetInternalOk returns a tuple with the Internal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomLinkRequest) GetInternalOk() (*bool, bool) {
	if o == nil || IsNil(o.Internal) {
		return nil, false
	}
	return o.Internal, true
}

// HasInternal returns a boolean if a field has been set.
func (o *RoomLinkRequest) IsInternalSet() bool {
	if o != nil && !IsNil(o.Internal) {
		return true
	}

	return false
}

// SetInternal gets a reference to the given bool and assigns it to the Internal field.
func (o *RoomLinkRequest) SetInternal(v bool) {
	o.Internal = &v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoomLinkRequest) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoomLinkRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *RoomLinkRequest) IsTitleSet() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *RoomLinkRequest) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *RoomLinkRequest) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *RoomLinkRequest) UnsetTitle() {
	o.Title.Unset()
}

// GetLinkType returns the LinkType field value if set, zero value otherwise.
func (o *RoomLinkRequest) GetLinkType() LinkType {
	if o == nil || IsNil(o.LinkType) {
		var ret LinkType
		return ret
	}
	return *o.LinkType
}

// GetLinkTypeOk returns a tuple with the LinkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomLinkRequest) GetLinkTypeOk() (*LinkType, bool) {
	if o == nil || IsNil(o.LinkType) {
		return nil, false
	}
	return o.LinkType, true
}

// HasLinkType returns a boolean if a field has been set.
func (o *RoomLinkRequest) IsLinkTypeSet() bool {
	if o != nil && !IsNil(o.LinkType) {
		return true
	}

	return false
}

// SetLinkType gets a reference to the given LinkType and assigns it to the LinkType field.
func (o *RoomLinkRequest) SetLinkType(v LinkType) {
	o.LinkType = &v
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoomLinkRequest) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoomLinkRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *RoomLinkRequest) IsPasswordSet() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *RoomLinkRequest) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *RoomLinkRequest) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *RoomLinkRequest) UnsetPassword() {
	o.Password.Unset()
}

// GetDenyDownload returns the DenyDownload field value if set, zero value otherwise.
func (o *RoomLinkRequest) GetDenyDownload() bool {
	if o == nil || IsNil(o.DenyDownload) {
		var ret bool
		return ret
	}
	return *o.DenyDownload
}

// GetDenyDownloadOk returns a tuple with the DenyDownload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomLinkRequest) GetDenyDownloadOk() (*bool, bool) {
	if o == nil || IsNil(o.DenyDownload) {
		return nil, false
	}
	return o.DenyDownload, true
}

// HasDenyDownload returns a boolean if a field has been set.
func (o *RoomLinkRequest) IsDenyDownloadSet() bool {
	if o != nil && !IsNil(o.DenyDownload) {
		return true
	}

	return false
}

// SetDenyDownload gets a reference to the given bool and assigns it to the DenyDownload field.
func (o *RoomLinkRequest) SetDenyDownload(v bool) {
	o.DenyDownload = &v
}

// GetMaxUseCount returns the MaxUseCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoomLinkRequest) GetMaxUseCount() int32 {
	if o == nil || IsNil(o.MaxUseCount.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxUseCount.Get()
}

// GetMaxUseCountOk returns a tuple with the MaxUseCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoomLinkRequest) GetMaxUseCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxUseCount.Get(), o.MaxUseCount.IsSet()
}

// HasMaxUseCount returns a boolean if a field has been set.
func (o *RoomLinkRequest) IsMaxUseCountSet() bool {
	if o != nil && o.MaxUseCount.IsSet() {
		return true
	}

	return false
}

// SetMaxUseCount gets a reference to the given NullableInt32 and assigns it to the MaxUseCount field.
func (o *RoomLinkRequest) SetMaxUseCount(v int32) {
	o.MaxUseCount.Set(&v)
}
// SetMaxUseCountNil sets the value for MaxUseCount to be an explicit nil
func (o *RoomLinkRequest) SetMaxUseCountNil() {
	o.MaxUseCount.Set(nil)
}

// UnsetMaxUseCount ensures that no value is present for MaxUseCount, not even an explicit nil
func (o *RoomLinkRequest) UnsetMaxUseCount() {
	o.MaxUseCount.Unset()
}

// GetCurrentUseCount returns the CurrentUseCount field value if set, zero value otherwise.
func (o *RoomLinkRequest) GetCurrentUseCount() int32 {
	if o == nil || IsNil(o.CurrentUseCount) {
		var ret int32
		return ret
	}
	return *o.CurrentUseCount
}

// GetCurrentUseCountOk returns a tuple with the CurrentUseCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomLinkRequest) GetCurrentUseCountOk() (*int32, bool) {
	if o == nil || IsNil(o.CurrentUseCount) {
		return nil, false
	}
	return o.CurrentUseCount, true
}

// HasCurrentUseCount returns a boolean if a field has been set.
func (o *RoomLinkRequest) IsCurrentUseCountSet() bool {
	if o != nil && !IsNil(o.CurrentUseCount) {
		return true
	}

	return false
}

// SetCurrentUseCount gets a reference to the given int32 and assigns it to the CurrentUseCount field.
func (o *RoomLinkRequest) SetCurrentUseCount(v int32) {
	o.CurrentUseCount = &v
}

func (o RoomLinkRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoomLinkRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LinkId) {
		toSerialize["linkId"] = o.LinkId
	}
	if !IsNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	if !IsNil(o.ExpirationDate) {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
	if !IsNil(o.Internal) {
		toSerialize["internal"] = o.Internal
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if !IsNil(o.LinkType) {
		toSerialize["linkType"] = o.LinkType
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if !IsNil(o.DenyDownload) {
		toSerialize["denyDownload"] = o.DenyDownload
	}
	if o.MaxUseCount.IsSet() {
		toSerialize["maxUseCount"] = o.MaxUseCount.Get()
	}
	if !IsNil(o.CurrentUseCount) {
		toSerialize["currentUseCount"] = o.CurrentUseCount
	}
	return toSerialize, nil
}

type NullableRoomLinkRequest struct {
	value *RoomLinkRequest
	isSet bool
}

func (v NullableRoomLinkRequest) Get() *RoomLinkRequest {
	return v.value
}

func (v *NullableRoomLinkRequest) Set(val *RoomLinkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRoomLinkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRoomLinkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoomLinkRequest(val *RoomLinkRequest) *NullableRoomLinkRequest {
	return &NullableRoomLinkRequest{value: val, isSet: true}
}

func (v NullableRoomLinkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoomLinkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

