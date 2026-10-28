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
	"bytes"
	"fmt"
)

// checks if the RoomTemplateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoomTemplateDto{}

// RoomTemplateDto The parameters of a room template built from an existing room.
type RoomTemplateDto struct {
	// The identifier of the room the template is built from. Take it from the room listing of  `GET api/2.0/files/rooms`; a folder identifier is not accepted.
	RoomId int32 `json:"roomId"`
	// The title the template is saved under in the Templates section. Characters that a folder name cannot contain  are replaced with an underscore on save, and two templates may share a title.
	Title string `json:"title"`
	// A picture of the caller's own for the template, cropped out of an image already placed in the temporary  storage.
	Logo *LogoRequest `json:"logo,omitempty"`
	// Whether the template takes over the picture already set on the source room. When false the template gets no  picture from that room.
	CopyLogo *bool `json:"copyLogo,omitempty"`
	// The email addresses of the portal members who are granted read access to the finished template.
	Share []string `json:"share,omitempty"`
	// The identifiers of the portal groups whose members are granted read access to the finished template.
	Groups []string `json:"groups,omitempty"`
	// Whether the finished template is shared with everyone allowed to create rooms. When false it stays reachable  only for the recipients named for it.
	Public *bool `json:"public,omitempty"`
	// The labels attached to the template and shown next to it in listings.
	Tags []string `json:"tags,omitempty"`
	// The accent colour of the generated cover, written as six hexadecimal digits with no leading hash sign. When it  is left empty a colour is picked at random.
	Color NullableString `json:"color,omitempty"`
	// The identifier of a built-in cover picture, as listed by `GET api/2.0/files/rooms/covers`. When it is left  empty the template gets no cover.
	Cover NullableString `json:"cover,omitempty"`
	// The storage limit assigned to the template, in bytes. When it is not set the template keeps the limit of the  source room.
	Quota NullableInt64 `json:"quota,omitempty"`
}

type _RoomTemplateDto RoomTemplateDto

// NewRoomTemplateDto instantiates a new RoomTemplateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoomTemplateDto(roomId int32, title string) *RoomTemplateDto {
	this := RoomTemplateDto{}
	this.RoomId = roomId
	this.Title = title
	return &this
}

// NewRoomTemplateDtoWithDefaults instantiates a new RoomTemplateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoomTemplateDtoWithDefaults() *RoomTemplateDto {
	this := RoomTemplateDto{}
	return &this
}

// GetRoomId returns the RoomId field value
func (o *RoomTemplateDto) GetRoomId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RoomId
}

// GetRoomIdOk returns a tuple with the RoomId field value
// and a boolean to check if the value has been set.
func (o *RoomTemplateDto) GetRoomIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoomId, true
}

// SetRoomId sets field value
func (o *RoomTemplateDto) SetRoomId(v int32) {
	o.RoomId = v
}

// GetTitle returns the Title field value
func (o *RoomTemplateDto) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *RoomTemplateDto) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *RoomTemplateDto) SetTitle(v string) {
	o.Title = v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *RoomTemplateDto) GetLogo() LogoRequest {
	if o == nil || IsNil(o.Logo) {
		var ret LogoRequest
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomTemplateDto) GetLogoOk() (*LogoRequest, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *RoomTemplateDto) IsLogoSet() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given LogoRequest and assigns it to the Logo field.
func (o *RoomTemplateDto) SetLogo(v LogoRequest) {
	o.Logo = &v
}

// GetCopyLogo returns the CopyLogo field value if set, zero value otherwise.
func (o *RoomTemplateDto) GetCopyLogo() bool {
	if o == nil || IsNil(o.CopyLogo) {
		var ret bool
		return ret
	}
	return *o.CopyLogo
}

// GetCopyLogoOk returns a tuple with the CopyLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomTemplateDto) GetCopyLogoOk() (*bool, bool) {
	if o == nil || IsNil(o.CopyLogo) {
		return nil, false
	}
	return o.CopyLogo, true
}

// HasCopyLogo returns a boolean if a field has been set.
func (o *RoomTemplateDto) IsCopyLogoSet() bool {
	if o != nil && !IsNil(o.CopyLogo) {
		return true
	}

	return false
}

// SetCopyLogo gets a reference to the given bool and assigns it to the CopyLogo field.
func (o *RoomTemplateDto) SetCopyLogo(v bool) {
	o.CopyLogo = &v
}

// GetShare returns the Share field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoomTemplateDto) GetShare() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Share
}

// GetShareOk returns a tuple with the Share field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoomTemplateDto) GetShareOk() ([]string, bool) {
	if o == nil || IsNil(o.Share) {
		return nil, false
	}
	return o.Share, true
}

// HasShare returns a boolean if a field has been set.
func (o *RoomTemplateDto) IsShareSet() bool {
	if o != nil && !IsNil(o.Share) {
		return true
	}

	return false
}

// SetShare gets a reference to the given []string and assigns it to the Share field.
func (o *RoomTemplateDto) SetShare(v []string) {
	o.Share = v
}

// GetGroups returns the Groups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoomTemplateDto) GetGroups() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoomTemplateDto) GetGroupsOk() ([]string, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *RoomTemplateDto) IsGroupsSet() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []string and assigns it to the Groups field.
func (o *RoomTemplateDto) SetGroups(v []string) {
	o.Groups = v
}

// GetPublic returns the Public field value if set, zero value otherwise.
func (o *RoomTemplateDto) GetPublic() bool {
	if o == nil || IsNil(o.Public) {
		var ret bool
		return ret
	}
	return *o.Public
}

// GetPublicOk returns a tuple with the Public field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoomTemplateDto) GetPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.Public) {
		return nil, false
	}
	return o.Public, true
}

// HasPublic returns a boolean if a field has been set.
func (o *RoomTemplateDto) IsPublicSet() bool {
	if o != nil && !IsNil(o.Public) {
		return true
	}

	return false
}

// SetPublic gets a reference to the given bool and assigns it to the Public field.
func (o *RoomTemplateDto) SetPublic(v bool) {
	o.Public = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoomTemplateDto) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoomTemplateDto) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *RoomTemplateDto) IsTagsSet() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *RoomTemplateDto) SetTags(v []string) {
	o.Tags = v
}

// GetColor returns the Color field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoomTemplateDto) GetColor() string {
	if o == nil || IsNil(o.Color.Get()) {
		var ret string
		return ret
	}
	return *o.Color.Get()
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoomTemplateDto) GetColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Color.Get(), o.Color.IsSet()
}

// HasColor returns a boolean if a field has been set.
func (o *RoomTemplateDto) IsColorSet() bool {
	if o != nil && o.Color.IsSet() {
		return true
	}

	return false
}

// SetColor gets a reference to the given NullableString and assigns it to the Color field.
func (o *RoomTemplateDto) SetColor(v string) {
	o.Color.Set(&v)
}
// SetColorNil sets the value for Color to be an explicit nil
func (o *RoomTemplateDto) SetColorNil() {
	o.Color.Set(nil)
}

// UnsetColor ensures that no value is present for Color, not even an explicit nil
func (o *RoomTemplateDto) UnsetColor() {
	o.Color.Unset()
}

// GetCover returns the Cover field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoomTemplateDto) GetCover() string {
	if o == nil || IsNil(o.Cover.Get()) {
		var ret string
		return ret
	}
	return *o.Cover.Get()
}

// GetCoverOk returns a tuple with the Cover field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoomTemplateDto) GetCoverOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cover.Get(), o.Cover.IsSet()
}

// HasCover returns a boolean if a field has been set.
func (o *RoomTemplateDto) IsCoverSet() bool {
	if o != nil && o.Cover.IsSet() {
		return true
	}

	return false
}

// SetCover gets a reference to the given NullableString and assigns it to the Cover field.
func (o *RoomTemplateDto) SetCover(v string) {
	o.Cover.Set(&v)
}
// SetCoverNil sets the value for Cover to be an explicit nil
func (o *RoomTemplateDto) SetCoverNil() {
	o.Cover.Set(nil)
}

// UnsetCover ensures that no value is present for Cover, not even an explicit nil
func (o *RoomTemplateDto) UnsetCover() {
	o.Cover.Unset()
}

// GetQuota returns the Quota field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoomTemplateDto) GetQuota() int64 {
	if o == nil || IsNil(o.Quota.Get()) {
		var ret int64
		return ret
	}
	return *o.Quota.Get()
}

// GetQuotaOk returns a tuple with the Quota field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoomTemplateDto) GetQuotaOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Quota.Get(), o.Quota.IsSet()
}

// HasQuota returns a boolean if a field has been set.
func (o *RoomTemplateDto) IsQuotaSet() bool {
	if o != nil && o.Quota.IsSet() {
		return true
	}

	return false
}

// SetQuota gets a reference to the given NullableInt64 and assigns it to the Quota field.
func (o *RoomTemplateDto) SetQuota(v int64) {
	o.Quota.Set(&v)
}
// SetQuotaNil sets the value for Quota to be an explicit nil
func (o *RoomTemplateDto) SetQuotaNil() {
	o.Quota.Set(nil)
}

// UnsetQuota ensures that no value is present for Quota, not even an explicit nil
func (o *RoomTemplateDto) UnsetQuota() {
	o.Quota.Unset()
}

func (o RoomTemplateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoomTemplateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["roomId"] = o.RoomId
	toSerialize["title"] = o.Title
	if !IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}
	if !IsNil(o.CopyLogo) {
		toSerialize["copyLogo"] = o.CopyLogo
	}
	if o.Share != nil {
		toSerialize["share"] = o.Share
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.Public) {
		toSerialize["public"] = o.Public
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Color.IsSet() {
		toSerialize["color"] = o.Color.Get()
	}
	if o.Cover.IsSet() {
		toSerialize["cover"] = o.Cover.Get()
	}
	if o.Quota.IsSet() {
		toSerialize["quota"] = o.Quota.Get()
	}
	return toSerialize, nil
}

func (o *RoomTemplateDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"roomId",
		"title",
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

	varRoomTemplateDto := _RoomTemplateDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varRoomTemplateDto)

	if err != nil {
		return err
	}

	*o = RoomTemplateDto(varRoomTemplateDto)

	return err
}

type NullableRoomTemplateDto struct {
	value *RoomTemplateDto
	isSet bool
}

func (v NullableRoomTemplateDto) Get() *RoomTemplateDto {
	return v.value
}

func (v *NullableRoomTemplateDto) Set(val *RoomTemplateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableRoomTemplateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableRoomTemplateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoomTemplateDto(val *RoomTemplateDto) *NullableRoomTemplateDto {
	return &NullableRoomTemplateDto{value: val, isSet: true}
}

func (v NullableRoomTemplateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoomTemplateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

