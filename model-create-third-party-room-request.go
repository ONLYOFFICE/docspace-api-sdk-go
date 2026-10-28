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

// checks if the CreateThirdPartyRoomRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateThirdPartyRoomRequest{}

// CreateThirdPartyRoomRequest The room to be created out of a folder of a connected third-party storage account.
type CreateThirdPartyRoomRequest struct {
	// Creates a new folder named after `title` inside the folder named in the path and turns that subfolder into the  room, leaving the named folder itself untouched. When omitted, the named folder becomes the room and keeps  everything it already holds.
	CreateAsNewFolder *bool `json:"createAsNewFolder,omitempty"`
	// The name the room is shown under. It is stored on the connected account, so it does not have to match the name  of the folder in the storage; with `createAsNewFolder` it is also the name given to the created subfolder.
	Title NullableString `json:"title"`
	// The kind of room the folder becomes, which decides the default access rules of its members and cannot be  changed afterwards.
	RoomType RoomType `json:"roomType"`
	// Restricts the room to the members explicitly invited into it. The flag is kept on the connected storage  account rather than on the folder, so every folder read through that account reports the same value.
	Private *bool `json:"private,omitempty"`
	// Keeps the contents of the room in an explicit numbered order, the one reported as `order` on every entry,  instead of leaving the order to the reader.
	Indexing *bool `json:"indexing,omitempty"`
	// Forbids downloading and printing the contents of the room, which leaves the members with viewing and editing  in the editor only.
	DenyDownload *bool `json:"denyDownload,omitempty"`
	// The background colour drawn behind the cover of the room, as six hexadecimal digits without a leading number  sign. An empty value restores the colour the portal picks by default.
	Color NullableString `json:"color,omitempty"`
	// The drawing shown on the room tile, named by one of the built-in cover identifiers returned by  `GET api/2.0/files/rooms/covers`. An empty value leaves the room without a cover, and any other unknown value  is rejected as an invalid request.
	Cover NullableString `json:"cover,omitempty"`
	// The tags to attach to the room, named by their text. A name that is not in the portal tag catalogue yet is  added to it, and `GET api/2.0/files/tags` lists the names already there.
	Tags []string `json:"tags,omitempty"`
	// The picture to use as the room logo, which has to be uploaded with `POST api/2.0/files/logos` first; leaving  it out keeps the room on its cover and colour.
	Logo *LogoRequest `json:"logo,omitempty"`
}

type _CreateThirdPartyRoomRequest CreateThirdPartyRoomRequest

// NewCreateThirdPartyRoomRequest instantiates a new CreateThirdPartyRoomRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateThirdPartyRoomRequest(title NullableString, roomType RoomType) *CreateThirdPartyRoomRequest {
	this := CreateThirdPartyRoomRequest{}
	this.Title = title
	this.RoomType = roomType
	return &this
}

// NewCreateThirdPartyRoomRequestWithDefaults instantiates a new CreateThirdPartyRoomRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateThirdPartyRoomRequestWithDefaults() *CreateThirdPartyRoomRequest {
	this := CreateThirdPartyRoomRequest{}
	return &this
}

// GetCreateAsNewFolder returns the CreateAsNewFolder field value if set, zero value otherwise.
func (o *CreateThirdPartyRoomRequest) GetCreateAsNewFolder() bool {
	if o == nil || IsNil(o.CreateAsNewFolder) {
		var ret bool
		return ret
	}
	return *o.CreateAsNewFolder
}

// GetCreateAsNewFolderOk returns a tuple with the CreateAsNewFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateThirdPartyRoomRequest) GetCreateAsNewFolderOk() (*bool, bool) {
	if o == nil || IsNil(o.CreateAsNewFolder) {
		return nil, false
	}
	return o.CreateAsNewFolder, true
}

// HasCreateAsNewFolder returns a boolean if a field has been set.
func (o *CreateThirdPartyRoomRequest) IsCreateAsNewFolderSet() bool {
	if o != nil && !IsNil(o.CreateAsNewFolder) {
		return true
	}

	return false
}

// SetCreateAsNewFolder gets a reference to the given bool and assigns it to the CreateAsNewFolder field.
func (o *CreateThirdPartyRoomRequest) SetCreateAsNewFolder(v bool) {
	o.CreateAsNewFolder = &v
}

// GetTitle returns the Title field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreateThirdPartyRoomRequest) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}

	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateThirdPartyRoomRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// SetTitle sets field value
func (o *CreateThirdPartyRoomRequest) SetTitle(v string) {
	o.Title.Set(&v)
}

// GetRoomType returns the RoomType field value
func (o *CreateThirdPartyRoomRequest) GetRoomType() RoomType {
	if o == nil {
		var ret RoomType
		return ret
	}

	return o.RoomType
}

// GetRoomTypeOk returns a tuple with the RoomType field value
// and a boolean to check if the value has been set.
func (o *CreateThirdPartyRoomRequest) GetRoomTypeOk() (*RoomType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoomType, true
}

// SetRoomType sets field value
func (o *CreateThirdPartyRoomRequest) SetRoomType(v RoomType) {
	o.RoomType = v
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *CreateThirdPartyRoomRequest) GetPrivate() bool {
	if o == nil || IsNil(o.Private) {
		var ret bool
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateThirdPartyRoomRequest) GetPrivateOk() (*bool, bool) {
	if o == nil || IsNil(o.Private) {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *CreateThirdPartyRoomRequest) IsPrivateSet() bool {
	if o != nil && !IsNil(o.Private) {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given bool and assigns it to the Private field.
func (o *CreateThirdPartyRoomRequest) SetPrivate(v bool) {
	o.Private = &v
}

// GetIndexing returns the Indexing field value if set, zero value otherwise.
func (o *CreateThirdPartyRoomRequest) GetIndexing() bool {
	if o == nil || IsNil(o.Indexing) {
		var ret bool
		return ret
	}
	return *o.Indexing
}

// GetIndexingOk returns a tuple with the Indexing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateThirdPartyRoomRequest) GetIndexingOk() (*bool, bool) {
	if o == nil || IsNil(o.Indexing) {
		return nil, false
	}
	return o.Indexing, true
}

// HasIndexing returns a boolean if a field has been set.
func (o *CreateThirdPartyRoomRequest) IsIndexingSet() bool {
	if o != nil && !IsNil(o.Indexing) {
		return true
	}

	return false
}

// SetIndexing gets a reference to the given bool and assigns it to the Indexing field.
func (o *CreateThirdPartyRoomRequest) SetIndexing(v bool) {
	o.Indexing = &v
}

// GetDenyDownload returns the DenyDownload field value if set, zero value otherwise.
func (o *CreateThirdPartyRoomRequest) GetDenyDownload() bool {
	if o == nil || IsNil(o.DenyDownload) {
		var ret bool
		return ret
	}
	return *o.DenyDownload
}

// GetDenyDownloadOk returns a tuple with the DenyDownload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateThirdPartyRoomRequest) GetDenyDownloadOk() (*bool, bool) {
	if o == nil || IsNil(o.DenyDownload) {
		return nil, false
	}
	return o.DenyDownload, true
}

// HasDenyDownload returns a boolean if a field has been set.
func (o *CreateThirdPartyRoomRequest) IsDenyDownloadSet() bool {
	if o != nil && !IsNil(o.DenyDownload) {
		return true
	}

	return false
}

// SetDenyDownload gets a reference to the given bool and assigns it to the DenyDownload field.
func (o *CreateThirdPartyRoomRequest) SetDenyDownload(v bool) {
	o.DenyDownload = &v
}

// GetColor returns the Color field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateThirdPartyRoomRequest) GetColor() string {
	if o == nil || IsNil(o.Color.Get()) {
		var ret string
		return ret
	}
	return *o.Color.Get()
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateThirdPartyRoomRequest) GetColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Color.Get(), o.Color.IsSet()
}

// HasColor returns a boolean if a field has been set.
func (o *CreateThirdPartyRoomRequest) IsColorSet() bool {
	if o != nil && o.Color.IsSet() {
		return true
	}

	return false
}

// SetColor gets a reference to the given NullableString and assigns it to the Color field.
func (o *CreateThirdPartyRoomRequest) SetColor(v string) {
	o.Color.Set(&v)
}
// SetColorNil sets the value for Color to be an explicit nil
func (o *CreateThirdPartyRoomRequest) SetColorNil() {
	o.Color.Set(nil)
}

// UnsetColor ensures that no value is present for Color, not even an explicit nil
func (o *CreateThirdPartyRoomRequest) UnsetColor() {
	o.Color.Unset()
}

// GetCover returns the Cover field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateThirdPartyRoomRequest) GetCover() string {
	if o == nil || IsNil(o.Cover.Get()) {
		var ret string
		return ret
	}
	return *o.Cover.Get()
}

// GetCoverOk returns a tuple with the Cover field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateThirdPartyRoomRequest) GetCoverOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cover.Get(), o.Cover.IsSet()
}

// HasCover returns a boolean if a field has been set.
func (o *CreateThirdPartyRoomRequest) IsCoverSet() bool {
	if o != nil && o.Cover.IsSet() {
		return true
	}

	return false
}

// SetCover gets a reference to the given NullableString and assigns it to the Cover field.
func (o *CreateThirdPartyRoomRequest) SetCover(v string) {
	o.Cover.Set(&v)
}
// SetCoverNil sets the value for Cover to be an explicit nil
func (o *CreateThirdPartyRoomRequest) SetCoverNil() {
	o.Cover.Set(nil)
}

// UnsetCover ensures that no value is present for Cover, not even an explicit nil
func (o *CreateThirdPartyRoomRequest) UnsetCover() {
	o.Cover.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateThirdPartyRoomRequest) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateThirdPartyRoomRequest) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateThirdPartyRoomRequest) IsTagsSet() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *CreateThirdPartyRoomRequest) SetTags(v []string) {
	o.Tags = v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *CreateThirdPartyRoomRequest) GetLogo() LogoRequest {
	if o == nil || IsNil(o.Logo) {
		var ret LogoRequest
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateThirdPartyRoomRequest) GetLogoOk() (*LogoRequest, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *CreateThirdPartyRoomRequest) IsLogoSet() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given LogoRequest and assigns it to the Logo field.
func (o *CreateThirdPartyRoomRequest) SetLogo(v LogoRequest) {
	o.Logo = &v
}

func (o CreateThirdPartyRoomRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateThirdPartyRoomRequest) ToMap() (map[string]interface{}, error) {
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

func (o *CreateThirdPartyRoomRequest) UnmarshalJSON(data []byte) (err error) {
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

	varCreateThirdPartyRoomRequest := _CreateThirdPartyRoomRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varCreateThirdPartyRoomRequest)

	if err != nil {
		return err
	}

	*o = CreateThirdPartyRoomRequest(varCreateThirdPartyRoomRequest)

	return err
}

type NullableCreateThirdPartyRoomRequest struct {
	value *CreateThirdPartyRoomRequest
	isSet bool
}

func (v NullableCreateThirdPartyRoomRequest) Get() *CreateThirdPartyRoomRequest {
	return v.value
}

func (v *NullableCreateThirdPartyRoomRequest) Set(val *CreateThirdPartyRoomRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateThirdPartyRoomRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateThirdPartyRoomRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateThirdPartyRoomRequest(val *CreateThirdPartyRoomRequest) *NullableCreateThirdPartyRoomRequest {
	return &NullableCreateThirdPartyRoomRequest{value: val, isSet: true}
}

func (v NullableCreateThirdPartyRoomRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateThirdPartyRoomRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

