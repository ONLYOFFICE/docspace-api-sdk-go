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

// checks if the FileShareLink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileShareLink{}

// FileShareLink A sharing link of a file, a folder or a room, with everything set on it.
type FileShareLink struct {
	// The identifier of the link, the one to send back as `linkId` to change or delete it.
	Id *string `json:"id,omitempty"`
	// The name the link is listed under, which its author is free to choose and to leave empty.
	Title NullableString `json:"title,omitempty"`
	// The shortened address to hand out. Opening it is what turns the link into access; the address stays the same  while the link exists.
	ShareLink NullableString `json:"shareLink,omitempty"`
	// The moment the link stops working, written with the offset of the portal time zone. Null when the link was  left without an end.
	ExpirationDate *ApiDateTime `json:"expirationDate,omitempty"`
	// Which of the two jobs the link does: letting somebody into the room as a member, or handing out the entry  itself. The counters of uses are filled in for the first kind only.
	LinkType *LinkType `json:"linkType,omitempty"`
	// The password a visitor has to send before the link resolves, readable only by those who may manage the link.  Empty when the link asks for none.
	Password NullableString `json:"password,omitempty"`
	// Whether visitors coming through this link may only read the entry in the editor and not download or print it.
	DenyDownload NullableBool `json:"denyDownload,omitempty"`
	// Whether the moment in `expirationDate` has already passed, which leaves the link in place but refuses  everybody who opens it.
	IsExpired NullableBool `json:"isExpired,omitempty"`
	// Whether this is the one link the entry always keeps: a public or a form-filling room is given it at creation,  and deleting it there only makes a new one.
	Primary *bool `json:"primary,omitempty"`
	// Whether the visitor has to sign in to the portal before the link resolves, as opposed to it being open to  anybody who has the address.
	Internal NullableBool `json:"internal,omitempty"`
	// The key that stands for this link in the calls that resolve it, such as `GET api/2.0/files/share/{key}`. It is  filled in for links that hand out the entry, and empty for the ones that invite into a room.
	RequestToken NullableString `json:"requestToken,omitempty"`
	// How many accounts may still join the room through this invitation link in total. Null on a link that hands out  the entry, where nothing is counted.
	MaxUseCount NullableInt32 `json:"maxUseCount,omitempty"`
	// How many accounts have already joined through this invitation link. Once it reaches `maxUseCount` the link  stops letting anybody else in. Null on a link that hands out the entry.
	CurrentUseCount NullableInt32 `json:"currentUseCount,omitempty"`
}

// NewFileShareLink instantiates a new FileShareLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileShareLink() *FileShareLink {
	this := FileShareLink{}
	return &this
}

// NewFileShareLinkWithDefaults instantiates a new FileShareLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileShareLinkWithDefaults() *FileShareLink {
	this := FileShareLink{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FileShareLink) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileShareLink) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FileShareLink) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FileShareLink) SetId(v string) {
	o.Id = &v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileShareLink) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileShareLink) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *FileShareLink) IsTitleSet() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *FileShareLink) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *FileShareLink) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *FileShareLink) UnsetTitle() {
	o.Title.Unset()
}

// GetShareLink returns the ShareLink field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileShareLink) GetShareLink() string {
	if o == nil || IsNil(o.ShareLink.Get()) {
		var ret string
		return ret
	}
	return *o.ShareLink.Get()
}

// GetShareLinkOk returns a tuple with the ShareLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileShareLink) GetShareLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShareLink.Get(), o.ShareLink.IsSet()
}

// HasShareLink returns a boolean if a field has been set.
func (o *FileShareLink) IsShareLinkSet() bool {
	if o != nil && o.ShareLink.IsSet() {
		return true
	}

	return false
}

// SetShareLink gets a reference to the given NullableString and assigns it to the ShareLink field.
func (o *FileShareLink) SetShareLink(v string) {
	o.ShareLink.Set(&v)
}
// SetShareLinkNil sets the value for ShareLink to be an explicit nil
func (o *FileShareLink) SetShareLinkNil() {
	o.ShareLink.Set(nil)
}

// UnsetShareLink ensures that no value is present for ShareLink, not even an explicit nil
func (o *FileShareLink) UnsetShareLink() {
	o.ShareLink.Unset()
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *FileShareLink) GetExpirationDate() ApiDateTime {
	if o == nil || IsNil(o.ExpirationDate) {
		var ret ApiDateTime
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileShareLink) GetExpirationDateOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *FileShareLink) IsExpirationDateSet() bool {
	if o != nil && !IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given ApiDateTime and assigns it to the ExpirationDate field.
func (o *FileShareLink) SetExpirationDate(v ApiDateTime) {
	o.ExpirationDate = &v
}

// GetLinkType returns the LinkType field value if set, zero value otherwise.
func (o *FileShareLink) GetLinkType() LinkType {
	if o == nil || IsNil(o.LinkType) {
		var ret LinkType
		return ret
	}
	return *o.LinkType
}

// GetLinkTypeOk returns a tuple with the LinkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileShareLink) GetLinkTypeOk() (*LinkType, bool) {
	if o == nil || IsNil(o.LinkType) {
		return nil, false
	}
	return o.LinkType, true
}

// HasLinkType returns a boolean if a field has been set.
func (o *FileShareLink) IsLinkTypeSet() bool {
	if o != nil && !IsNil(o.LinkType) {
		return true
	}

	return false
}

// SetLinkType gets a reference to the given LinkType and assigns it to the LinkType field.
func (o *FileShareLink) SetLinkType(v LinkType) {
	o.LinkType = &v
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileShareLink) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileShareLink) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *FileShareLink) IsPasswordSet() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *FileShareLink) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *FileShareLink) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *FileShareLink) UnsetPassword() {
	o.Password.Unset()
}

// GetDenyDownload returns the DenyDownload field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileShareLink) GetDenyDownload() bool {
	if o == nil || IsNil(o.DenyDownload.Get()) {
		var ret bool
		return ret
	}
	return *o.DenyDownload.Get()
}

// GetDenyDownloadOk returns a tuple with the DenyDownload field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileShareLink) GetDenyDownloadOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DenyDownload.Get(), o.DenyDownload.IsSet()
}

// HasDenyDownload returns a boolean if a field has been set.
func (o *FileShareLink) IsDenyDownloadSet() bool {
	if o != nil && o.DenyDownload.IsSet() {
		return true
	}

	return false
}

// SetDenyDownload gets a reference to the given NullableBool and assigns it to the DenyDownload field.
func (o *FileShareLink) SetDenyDownload(v bool) {
	o.DenyDownload.Set(&v)
}
// SetDenyDownloadNil sets the value for DenyDownload to be an explicit nil
func (o *FileShareLink) SetDenyDownloadNil() {
	o.DenyDownload.Set(nil)
}

// UnsetDenyDownload ensures that no value is present for DenyDownload, not even an explicit nil
func (o *FileShareLink) UnsetDenyDownload() {
	o.DenyDownload.Unset()
}

// GetIsExpired returns the IsExpired field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileShareLink) GetIsExpired() bool {
	if o == nil || IsNil(o.IsExpired.Get()) {
		var ret bool
		return ret
	}
	return *o.IsExpired.Get()
}

// GetIsExpiredOk returns a tuple with the IsExpired field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileShareLink) GetIsExpiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsExpired.Get(), o.IsExpired.IsSet()
}

// HasIsExpired returns a boolean if a field has been set.
func (o *FileShareLink) IsIsExpiredSet() bool {
	if o != nil && o.IsExpired.IsSet() {
		return true
	}

	return false
}

// SetIsExpired gets a reference to the given NullableBool and assigns it to the IsExpired field.
func (o *FileShareLink) SetIsExpired(v bool) {
	o.IsExpired.Set(&v)
}
// SetIsExpiredNil sets the value for IsExpired to be an explicit nil
func (o *FileShareLink) SetIsExpiredNil() {
	o.IsExpired.Set(nil)
}

// UnsetIsExpired ensures that no value is present for IsExpired, not even an explicit nil
func (o *FileShareLink) UnsetIsExpired() {
	o.IsExpired.Unset()
}

// GetPrimary returns the Primary field value if set, zero value otherwise.
func (o *FileShareLink) GetPrimary() bool {
	if o == nil || IsNil(o.Primary) {
		var ret bool
		return ret
	}
	return *o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileShareLink) GetPrimaryOk() (*bool, bool) {
	if o == nil || IsNil(o.Primary) {
		return nil, false
	}
	return o.Primary, true
}

// HasPrimary returns a boolean if a field has been set.
func (o *FileShareLink) IsPrimarySet() bool {
	if o != nil && !IsNil(o.Primary) {
		return true
	}

	return false
}

// SetPrimary gets a reference to the given bool and assigns it to the Primary field.
func (o *FileShareLink) SetPrimary(v bool) {
	o.Primary = &v
}

// GetInternal returns the Internal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileShareLink) GetInternal() bool {
	if o == nil || IsNil(o.Internal.Get()) {
		var ret bool
		return ret
	}
	return *o.Internal.Get()
}

// GetInternalOk returns a tuple with the Internal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileShareLink) GetInternalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Internal.Get(), o.Internal.IsSet()
}

// HasInternal returns a boolean if a field has been set.
func (o *FileShareLink) IsInternalSet() bool {
	if o != nil && o.Internal.IsSet() {
		return true
	}

	return false
}

// SetInternal gets a reference to the given NullableBool and assigns it to the Internal field.
func (o *FileShareLink) SetInternal(v bool) {
	o.Internal.Set(&v)
}
// SetInternalNil sets the value for Internal to be an explicit nil
func (o *FileShareLink) SetInternalNil() {
	o.Internal.Set(nil)
}

// UnsetInternal ensures that no value is present for Internal, not even an explicit nil
func (o *FileShareLink) UnsetInternal() {
	o.Internal.Unset()
}

// GetRequestToken returns the RequestToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileShareLink) GetRequestToken() string {
	if o == nil || IsNil(o.RequestToken.Get()) {
		var ret string
		return ret
	}
	return *o.RequestToken.Get()
}

// GetRequestTokenOk returns a tuple with the RequestToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileShareLink) GetRequestTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestToken.Get(), o.RequestToken.IsSet()
}

// HasRequestToken returns a boolean if a field has been set.
func (o *FileShareLink) IsRequestTokenSet() bool {
	if o != nil && o.RequestToken.IsSet() {
		return true
	}

	return false
}

// SetRequestToken gets a reference to the given NullableString and assigns it to the RequestToken field.
func (o *FileShareLink) SetRequestToken(v string) {
	o.RequestToken.Set(&v)
}
// SetRequestTokenNil sets the value for RequestToken to be an explicit nil
func (o *FileShareLink) SetRequestTokenNil() {
	o.RequestToken.Set(nil)
}

// UnsetRequestToken ensures that no value is present for RequestToken, not even an explicit nil
func (o *FileShareLink) UnsetRequestToken() {
	o.RequestToken.Unset()
}

// GetMaxUseCount returns the MaxUseCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileShareLink) GetMaxUseCount() int32 {
	if o == nil || IsNil(o.MaxUseCount.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxUseCount.Get()
}

// GetMaxUseCountOk returns a tuple with the MaxUseCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileShareLink) GetMaxUseCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxUseCount.Get(), o.MaxUseCount.IsSet()
}

// HasMaxUseCount returns a boolean if a field has been set.
func (o *FileShareLink) IsMaxUseCountSet() bool {
	if o != nil && o.MaxUseCount.IsSet() {
		return true
	}

	return false
}

// SetMaxUseCount gets a reference to the given NullableInt32 and assigns it to the MaxUseCount field.
func (o *FileShareLink) SetMaxUseCount(v int32) {
	o.MaxUseCount.Set(&v)
}
// SetMaxUseCountNil sets the value for MaxUseCount to be an explicit nil
func (o *FileShareLink) SetMaxUseCountNil() {
	o.MaxUseCount.Set(nil)
}

// UnsetMaxUseCount ensures that no value is present for MaxUseCount, not even an explicit nil
func (o *FileShareLink) UnsetMaxUseCount() {
	o.MaxUseCount.Unset()
}

// GetCurrentUseCount returns the CurrentUseCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileShareLink) GetCurrentUseCount() int32 {
	if o == nil || IsNil(o.CurrentUseCount.Get()) {
		var ret int32
		return ret
	}
	return *o.CurrentUseCount.Get()
}

// GetCurrentUseCountOk returns a tuple with the CurrentUseCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileShareLink) GetCurrentUseCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrentUseCount.Get(), o.CurrentUseCount.IsSet()
}

// HasCurrentUseCount returns a boolean if a field has been set.
func (o *FileShareLink) IsCurrentUseCountSet() bool {
	if o != nil && o.CurrentUseCount.IsSet() {
		return true
	}

	return false
}

// SetCurrentUseCount gets a reference to the given NullableInt32 and assigns it to the CurrentUseCount field.
func (o *FileShareLink) SetCurrentUseCount(v int32) {
	o.CurrentUseCount.Set(&v)
}
// SetCurrentUseCountNil sets the value for CurrentUseCount to be an explicit nil
func (o *FileShareLink) SetCurrentUseCountNil() {
	o.CurrentUseCount.Set(nil)
}

// UnsetCurrentUseCount ensures that no value is present for CurrentUseCount, not even an explicit nil
func (o *FileShareLink) UnsetCurrentUseCount() {
	o.CurrentUseCount.Unset()
}

func (o FileShareLink) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileShareLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.ShareLink.IsSet() {
		toSerialize["shareLink"] = o.ShareLink.Get()
	}
	if !IsNil(o.ExpirationDate) {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
	if !IsNil(o.LinkType) {
		toSerialize["linkType"] = o.LinkType
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if o.DenyDownload.IsSet() {
		toSerialize["denyDownload"] = o.DenyDownload.Get()
	}
	if o.IsExpired.IsSet() {
		toSerialize["isExpired"] = o.IsExpired.Get()
	}
	if !IsNil(o.Primary) {
		toSerialize["primary"] = o.Primary
	}
	if o.Internal.IsSet() {
		toSerialize["internal"] = o.Internal.Get()
	}
	if o.RequestToken.IsSet() {
		toSerialize["requestToken"] = o.RequestToken.Get()
	}
	if o.MaxUseCount.IsSet() {
		toSerialize["maxUseCount"] = o.MaxUseCount.Get()
	}
	if o.CurrentUseCount.IsSet() {
		toSerialize["currentUseCount"] = o.CurrentUseCount.Get()
	}
	return toSerialize, nil
}

type NullableFileShareLink struct {
	value *FileShareLink
	isSet bool
}

func (v NullableFileShareLink) Get() *FileShareLink {
	return v.value
}

func (v *NullableFileShareLink) Set(val *FileShareLink) {
	v.value = val
	v.isSet = true
}

func (v NullableFileShareLink) IsSet() bool {
	return v.isSet
}

func (v *NullableFileShareLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileShareLink(val *FileShareLink) *NullableFileShareLink {
	return &NullableFileShareLink{value: val, isSet: true}
}

func (v NullableFileShareLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileShareLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

