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

// checks if the CreateThirdPartyRoom type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateThirdPartyRoom{}

// CreateThirdPartyRoom The parameters for creating a third-party room.
type CreateThirdPartyRoom struct {
	// Specifies whether to create a third-party room as a new folder or not.
	CreateAsNewFolder *bool `json:"createAsNewFolder,omitempty"`
	// The third-party room name to be created.
	Title NullableString `json:"title"`
	RoomType RoomType `json:"roomType"`
	// Specifies whether to create the private third-party room or not.
	Private *bool `json:"private,omitempty"`
	// Specifies whether to create the third-party room with indexing.
	Indexing *bool `json:"indexing,omitempty"`
	// Specifies whether to deny downloads from the third-party room.
	DenyDownload *bool `json:"denyDownload,omitempty"`
	// The color of the third-party room.
	Color NullableString `json:"color,omitempty"`
	// The cover of the third-party room.
	Cover NullableString `json:"cover,omitempty"`
	// The list of tags of the third-party room.
	Tags []string `json:"tags,omitempty"`
	Logo *LogoRequest `json:"logo,omitempty"`
}

type _CreateThirdPartyRoom CreateThirdPartyRoom

// NewCreateThirdPartyRoom instantiates a new CreateThirdPartyRoom object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateThirdPartyRoom(title NullableString, roomType RoomType) *CreateThirdPartyRoom {
	this := CreateThirdPartyRoom{}
	this.Title = title
	this.RoomType = roomType
	return &this
}

// NewCreateThirdPartyRoomWithDefaults instantiates a new CreateThirdPartyRoom object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateThirdPartyRoomWithDefaults() *CreateThirdPartyRoom {
	this := CreateThirdPartyRoom{}
	return &this
}

// GetCreateAsNewFolder returns the CreateAsNewFolder field value if set, zero value otherwise.
func (o *CreateThirdPartyRoom) GetCreateAsNewFolder() bool {
	if o == nil || IsNil(o.CreateAsNewFolder) {
		var ret bool
		return ret
	}
	return *o.CreateAsNewFolder
}

// GetCreateAsNewFolderOk returns a tuple with the CreateAsNewFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateThirdPartyRoom) GetCreateAsNewFolderOk() (*bool, bool) {
	if o == nil || IsNil(o.CreateAsNewFolder) {
		return nil, false
	}
	return o.CreateAsNewFolder, true
}

// HasCreateAsNewFolder returns a boolean if a field has been set.
func (o *CreateThirdPartyRoom) IsCreateAsNewFolderSet() bool {
	if o != nil && !IsNil(o.CreateAsNewFolder) {
		return true
	}

	return false
}

// SetCreateAsNewFolder gets a reference to the given bool and assigns it to the CreateAsNewFolder field.
func (o *CreateThirdPartyRoom) SetCreateAsNewFolder(v bool) {
	o.CreateAsNewFolder = &v
}

// GetTitle returns the Title field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreateThirdPartyRoom) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}

	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateThirdPartyRoom) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// SetTitle sets field value
func (o *CreateThirdPartyRoom) SetTitle(v string) {
	o.Title.Set(&v)
}

// GetRoomType returns the RoomType field value
func (o *CreateThirdPartyRoom) GetRoomType() RoomType {
	if o == nil {
		var ret RoomType
		return ret
	}

	return o.RoomType
}

// GetRoomTypeOk returns a tuple with the RoomType field value
// and a boolean to check if the value has been set.
func (o *CreateThirdPartyRoom) GetRoomTypeOk() (*RoomType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoomType, true
}

// SetRoomType sets field value
func (o *CreateThirdPartyRoom) SetRoomType(v RoomType) {
	o.RoomType = v
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *CreateThirdPartyRoom) GetPrivate() bool {
	if o == nil || IsNil(o.Private) {
		var ret bool
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateThirdPartyRoom) GetPrivateOk() (*bool, bool) {
	if o == nil || IsNil(o.Private) {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *CreateThirdPartyRoom) IsPrivateSet() bool {
	if o != nil && !IsNil(o.Private) {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given bool and assigns it to the Private field.
func (o *CreateThirdPartyRoom) SetPrivate(v bool) {
	o.Private = &v
}

// GetIndexing returns the Indexing field value if set, zero value otherwise.
func (o *CreateThirdPartyRoom) GetIndexing() bool {
	if o == nil || IsNil(o.Indexing) {
		var ret bool
		return ret
	}
	return *o.Indexing
}

// GetIndexingOk returns a tuple with the Indexing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateThirdPartyRoom) GetIndexingOk() (*bool, bool) {
	if o == nil || IsNil(o.Indexing) {
		return nil, false
	}
	return o.Indexing, true
}

// HasIndexing returns a boolean if a field has been set.
func (o *CreateThirdPartyRoom) IsIndexingSet() bool {
	if o != nil && !IsNil(o.Indexing) {
		return true
	}

	return false
}

// SetIndexing gets a reference to the given bool and assigns it to the Indexing field.
func (o *CreateThirdPartyRoom) SetIndexing(v bool) {
	o.Indexing = &v
}

// GetDenyDownload returns the DenyDownload field value if set, zero value otherwise.
func (o *CreateThirdPartyRoom) GetDenyDownload() bool {
	if o == nil || IsNil(o.DenyDownload) {
		var ret bool
		return ret
	}
	return *o.DenyDownload
}

// GetDenyDownloadOk returns a tuple with the DenyDownload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateThirdPartyRoom) GetDenyDownloadOk() (*bool, bool) {
	if o == nil || IsNil(o.DenyDownload) {
		return nil, false
	}
	return o.DenyDownload, true
}

// HasDenyDownload returns a boolean if a field has been set.
func (o *CreateThirdPartyRoom) IsDenyDownloadSet() bool {
	if o != nil && !IsNil(o.DenyDownload) {
		return true
	}

	return false
}

// SetDenyDownload gets a reference to the given bool and assigns it to the DenyDownload field.
func (o *CreateThirdPartyRoom) SetDenyDownload(v bool) {
	o.DenyDownload = &v
}

// GetColor returns the Color field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateThirdPartyRoom) GetColor() string {
	if o == nil || IsNil(o.Color.Get()) {
		var ret string
		return ret
	}
	return *o.Color.Get()
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateThirdPartyRoom) GetColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Color.Get(), o.Color.IsSet()
}

// HasColor returns a boolean if a field has been set.
func (o *CreateThirdPartyRoom) IsColorSet() bool {
	if o != nil && o.Color.IsSet() {
		return true
	}

	return false
}

// SetColor gets a reference to the given NullableString and assigns it to the Color field.
func (o *CreateThirdPartyRoom) SetColor(v string) {
	o.Color.Set(&v)
}
// SetColorNil sets the value for Color to be an explicit nil
func (o *CreateThirdPartyRoom) SetColorNil() {
	o.Color.Set(nil)
}

// UnsetColor ensures that no value is present for Color, not even an explicit nil
func (o *CreateThirdPartyRoom) UnsetColor() {
	o.Color.Unset()
}

// GetCover returns the Cover field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateThirdPartyRoom) GetCover() string {
	if o == nil || IsNil(o.Cover.Get()) {
		var ret string
		return ret
	}
	return *o.Cover.Get()
}

// GetCoverOk returns a tuple with the Cover field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateThirdPartyRoom) GetCoverOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cover.Get(), o.Cover.IsSet()
}

// HasCover returns a boolean if a field has been set.
func (o *CreateThirdPartyRoom) IsCoverSet() bool {
	if o != nil && o.Cover.IsSet() {
		return true
	}

	return false
}

// SetCover gets a reference to the given NullableString and assigns it to the Cover field.
func (o *CreateThirdPartyRoom) SetCover(v string) {
	o.Cover.Set(&v)
}
// SetCoverNil sets the value for Cover to be an explicit nil
func (o *CreateThirdPartyRoom) SetCoverNil() {
	o.Cover.Set(nil)
}

// UnsetCover ensures that no value is present for Cover, not even an explicit nil
func (o *CreateThirdPartyRoom) UnsetCover() {
	o.Cover.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateThirdPartyRoom) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateThirdPartyRoom) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateThirdPartyRoom) IsTagsSet() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *CreateThirdPartyRoom) SetTags(v []string) {
	o.Tags = v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *CreateThirdPartyRoom) GetLogo() LogoRequest {
	if o == nil || IsNil(o.Logo) {
		var ret LogoRequest
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateThirdPartyRoom) GetLogoOk() (*LogoRequest, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *CreateThirdPartyRoom) IsLogoSet() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given LogoRequest and assigns it to the Logo field.
func (o *CreateThirdPartyRoom) SetLogo(v LogoRequest) {
	o.Logo = &v
}

func (o CreateThirdPartyRoom) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateThirdPartyRoom) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreateAsNewFolder) {
		toSerialize["createAsNewFolder"] = o.CreateAsNewFolder
	}
	toSerialize["title"] = o.Title.Get()
	toSerialize["roomType"] = o.RoomType
	if !IsNil(o.Private) {
		toSerialize["private"] = o.Private
	}
	if !IsNil(o.Indexing) {
		toSerialize["indexing"] = o.Indexing
	}
	if !IsNil(o.DenyDownload) {
		toSerialize["denyDownload"] = o.DenyDownload
	}
	if o.Color.IsSet() {
		toSerialize["color"] = o.Color.Get()
	}
	if o.Cover.IsSet() {
		toSerialize["cover"] = o.Cover.Get()
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}
	return toSerialize, nil
}

func (o *CreateThirdPartyRoom) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
		"roomType",
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

	varCreateThirdPartyRoom := _CreateThirdPartyRoom{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varCreateThirdPartyRoom)

	if err != nil {
		return err
	}

	*o = CreateThirdPartyRoom(varCreateThirdPartyRoom)

	return err
}

type NullableCreateThirdPartyRoom struct {
	value *CreateThirdPartyRoom
	isSet bool
}

func (v NullableCreateThirdPartyRoom) Get() *CreateThirdPartyRoom {
	return v.value
}

func (v *NullableCreateThirdPartyRoom) Set(val *CreateThirdPartyRoom) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateThirdPartyRoom) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateThirdPartyRoom) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateThirdPartyRoom(val *CreateThirdPartyRoom) *NullableCreateThirdPartyRoom {
	return &NullableCreateThirdPartyRoom{value: val, isSet: true}
}

func (v NullableCreateThirdPartyRoom) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateThirdPartyRoom) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

