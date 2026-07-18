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

// checks if the CreateRoomRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRoomRequestDto{}

// CreateRoomRequestDto The request parameters for creating a room.
type CreateRoomRequestDto struct {
	// The room name.
	Title NullableString `json:"title"`
	// The room quota.
	Quota NullableInt64 `json:"quota,omitempty"`
	// Specifies whether to create a room with indexing.
	Indexing NullableBool `json:"indexing,omitempty"`
	// Specifies whether to deny downloads from the room.
	DenyDownload NullableBool `json:"denyDownload,omitempty"`
	Lifetime *RoomDataLifetimeDto `json:"lifetime,omitempty"`
	Watermark *WatermarkRequestDto `json:"watermark,omitempty"`
	Logo *LogoRequest `json:"logo,omitempty"`
	// The list of tags.
	Tags []string `json:"tags,omitempty"`
	// The room color.
	Color NullableString `json:"color,omitempty"`
	// The room cover.
	Cover NullableString `json:"cover,omitempty"`
	RoomType RoomType `json:"roomType"`
	// Specifies whether the room to be created is private or not.
	Private *bool `json:"private,omitempty"`
	// The collection of sharing parameters.
	Share []FileShareParams `json:"share,omitempty"`
	ChatSettings *ChatSettings `json:"chatSettings,omitempty"`
	// Specifies whether to send form data to external database.
	SendFormToExternalDB NullableBool `json:"sendFormToExternalDB,omitempty"`
	// Specifies whether to save form data as XLSX file.
	SaveFormAsXLSX NullableBool `json:"saveFormAsXLSX,omitempty"`
}

type _CreateRoomRequestDto CreateRoomRequestDto

// NewCreateRoomRequestDto instantiates a new CreateRoomRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRoomRequestDto(title NullableString, roomType RoomType) *CreateRoomRequestDto {
	this := CreateRoomRequestDto{}
	this.Title = title
	this.RoomType = roomType
	return &this
}

// NewCreateRoomRequestDtoWithDefaults instantiates a new CreateRoomRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRoomRequestDtoWithDefaults() *CreateRoomRequestDto {
	this := CreateRoomRequestDto{}
	return &this
}

// GetTitle returns the Title field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreateRoomRequestDto) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}

	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateRoomRequestDto) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// SetTitle sets field value
func (o *CreateRoomRequestDto) SetTitle(v string) {
	o.Title.Set(&v)
}

// GetQuota returns the Quota field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateRoomRequestDto) GetQuota() int64 {
	if o == nil || IsNil(o.Quota.Get()) {
		var ret int64
		return ret
	}
	return *o.Quota.Get()
}

// GetQuotaOk returns a tuple with the Quota field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateRoomRequestDto) GetQuotaOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Quota.Get(), o.Quota.IsSet()
}

// HasQuota returns a boolean if a field has been set.
func (o *CreateRoomRequestDto) IsQuotaSet() bool {
	if o != nil && o.Quota.IsSet() {
		return true
	}

	return false
}

// SetQuota gets a reference to the given NullableInt64 and assigns it to the Quota field.
func (o *CreateRoomRequestDto) SetQuota(v int64) {
	o.Quota.Set(&v)
}
// SetQuotaNil sets the value for Quota to be an explicit nil
func (o *CreateRoomRequestDto) SetQuotaNil() {
	o.Quota.Set(nil)
}

// UnsetQuota ensures that no value is present for Quota, not even an explicit nil
func (o *CreateRoomRequestDto) UnsetQuota() {
	o.Quota.Unset()
}

// GetIndexing returns the Indexing field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateRoomRequestDto) GetIndexing() bool {
	if o == nil || IsNil(o.Indexing.Get()) {
		var ret bool
		return ret
	}
	return *o.Indexing.Get()
}

// GetIndexingOk returns a tuple with the Indexing field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateRoomRequestDto) GetIndexingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Indexing.Get(), o.Indexing.IsSet()
}

// HasIndexing returns a boolean if a field has been set.
func (o *CreateRoomRequestDto) IsIndexingSet() bool {
	if o != nil && o.Indexing.IsSet() {
		return true
	}

	return false
}

// SetIndexing gets a reference to the given NullableBool and assigns it to the Indexing field.
func (o *CreateRoomRequestDto) SetIndexing(v bool) {
	o.Indexing.Set(&v)
}
// SetIndexingNil sets the value for Indexing to be an explicit nil
func (o *CreateRoomRequestDto) SetIndexingNil() {
	o.Indexing.Set(nil)
}

// UnsetIndexing ensures that no value is present for Indexing, not even an explicit nil
func (o *CreateRoomRequestDto) UnsetIndexing() {
	o.Indexing.Unset()
}

// GetDenyDownload returns the DenyDownload field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateRoomRequestDto) GetDenyDownload() bool {
	if o == nil || IsNil(o.DenyDownload.Get()) {
		var ret bool
		return ret
	}
	return *o.DenyDownload.Get()
}

// GetDenyDownloadOk returns a tuple with the DenyDownload field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateRoomRequestDto) GetDenyDownloadOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DenyDownload.Get(), o.DenyDownload.IsSet()
}

// HasDenyDownload returns a boolean if a field has been set.
func (o *CreateRoomRequestDto) IsDenyDownloadSet() bool {
	if o != nil && o.DenyDownload.IsSet() {
		return true
	}

	return false
}

// SetDenyDownload gets a reference to the given NullableBool and assigns it to the DenyDownload field.
func (o *CreateRoomRequestDto) SetDenyDownload(v bool) {
	o.DenyDownload.Set(&v)
}
// SetDenyDownloadNil sets the value for DenyDownload to be an explicit nil
func (o *CreateRoomRequestDto) SetDenyDownloadNil() {
	o.DenyDownload.Set(nil)
}

// UnsetDenyDownload ensures that no value is present for DenyDownload, not even an explicit nil
func (o *CreateRoomRequestDto) UnsetDenyDownload() {
	o.DenyDownload.Unset()
}

// GetLifetime returns the Lifetime field value if set, zero value otherwise.
func (o *CreateRoomRequestDto) GetLifetime() RoomDataLifetimeDto {
	if o == nil || IsNil(o.Lifetime) {
		var ret RoomDataLifetimeDto
		return ret
	}
	return *o.Lifetime
}

// GetLifetimeOk returns a tuple with the Lifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRoomRequestDto) GetLifetimeOk() (*RoomDataLifetimeDto, bool) {
	if o == nil || IsNil(o.Lifetime) {
		return nil, false
	}
	return o.Lifetime, true
}

// HasLifetime returns a boolean if a field has been set.
func (o *CreateRoomRequestDto) IsLifetimeSet() bool {
	if o != nil && !IsNil(o.Lifetime) {
		return true
	}

	return false
}

// SetLifetime gets a reference to the given RoomDataLifetimeDto and assigns it to the Lifetime field.
func (o *CreateRoomRequestDto) SetLifetime(v RoomDataLifetimeDto) {
	o.Lifetime = &v
}

// GetWatermark returns the Watermark field value if set, zero value otherwise.
func (o *CreateRoomRequestDto) GetWatermark() WatermarkRequestDto {
	if o == nil || IsNil(o.Watermark) {
		var ret WatermarkRequestDto
		return ret
	}
	return *o.Watermark
}

// GetWatermarkOk returns a tuple with the Watermark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRoomRequestDto) GetWatermarkOk() (*WatermarkRequestDto, bool) {
	if o == nil || IsNil(o.Watermark) {
		return nil, false
	}
	return o.Watermark, true
}

// HasWatermark returns a boolean if a field has been set.
func (o *CreateRoomRequestDto) IsWatermarkSet() bool {
	if o != nil && !IsNil(o.Watermark) {
		return true
	}

	return false
}

// SetWatermark gets a reference to the given WatermarkRequestDto and assigns it to the Watermark field.
func (o *CreateRoomRequestDto) SetWatermark(v WatermarkRequestDto) {
	o.Watermark = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *CreateRoomRequestDto) GetLogo() LogoRequest {
	if o == nil || IsNil(o.Logo) {
		var ret LogoRequest
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRoomRequestDto) GetLogoOk() (*LogoRequest, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *CreateRoomRequestDto) IsLogoSet() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given LogoRequest and assigns it to the Logo field.
func (o *CreateRoomRequestDto) SetLogo(v LogoRequest) {
	o.Logo = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateRoomRequestDto) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateRoomRequestDto) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateRoomRequestDto) IsTagsSet() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *CreateRoomRequestDto) SetTags(v []string) {
	o.Tags = v
}

// GetColor returns the Color field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateRoomRequestDto) GetColor() string {
	if o == nil || IsNil(o.Color.Get()) {
		var ret string
		return ret
	}
	return *o.Color.Get()
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateRoomRequestDto) GetColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Color.Get(), o.Color.IsSet()
}

// HasColor returns a boolean if a field has been set.
func (o *CreateRoomRequestDto) IsColorSet() bool {
	if o != nil && o.Color.IsSet() {
		return true
	}

	return false
}

// SetColor gets a reference to the given NullableString and assigns it to the Color field.
func (o *CreateRoomRequestDto) SetColor(v string) {
	o.Color.Set(&v)
}
// SetColorNil sets the value for Color to be an explicit nil
func (o *CreateRoomRequestDto) SetColorNil() {
	o.Color.Set(nil)
}

// UnsetColor ensures that no value is present for Color, not even an explicit nil
func (o *CreateRoomRequestDto) UnsetColor() {
	o.Color.Unset()
}

// GetCover returns the Cover field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateRoomRequestDto) GetCover() string {
	if o == nil || IsNil(o.Cover.Get()) {
		var ret string
		return ret
	}
	return *o.Cover.Get()
}

// GetCoverOk returns a tuple with the Cover field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateRoomRequestDto) GetCoverOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cover.Get(), o.Cover.IsSet()
}

// HasCover returns a boolean if a field has been set.
func (o *CreateRoomRequestDto) IsCoverSet() bool {
	if o != nil && o.Cover.IsSet() {
		return true
	}

	return false
}

// SetCover gets a reference to the given NullableString and assigns it to the Cover field.
func (o *CreateRoomRequestDto) SetCover(v string) {
	o.Cover.Set(&v)
}
// SetCoverNil sets the value for Cover to be an explicit nil
func (o *CreateRoomRequestDto) SetCoverNil() {
	o.Cover.Set(nil)
}

// UnsetCover ensures that no value is present for Cover, not even an explicit nil
func (o *CreateRoomRequestDto) UnsetCover() {
	o.Cover.Unset()
}

// GetRoomType returns the RoomType field value
func (o *CreateRoomRequestDto) GetRoomType() RoomType {
	if o == nil {
		var ret RoomType
		return ret
	}

	return o.RoomType
}

// GetRoomTypeOk returns a tuple with the RoomType field value
// and a boolean to check if the value has been set.
func (o *CreateRoomRequestDto) GetRoomTypeOk() (*RoomType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoomType, true
}

// SetRoomType sets field value
func (o *CreateRoomRequestDto) SetRoomType(v RoomType) {
	o.RoomType = v
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *CreateRoomRequestDto) GetPrivate() bool {
	if o == nil || IsNil(o.Private) {
		var ret bool
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRoomRequestDto) GetPrivateOk() (*bool, bool) {
	if o == nil || IsNil(o.Private) {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *CreateRoomRequestDto) IsPrivateSet() bool {
	if o != nil && !IsNil(o.Private) {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given bool and assigns it to the Private field.
func (o *CreateRoomRequestDto) SetPrivate(v bool) {
	o.Private = &v
}

// GetShare returns the Share field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateRoomRequestDto) GetShare() []FileShareParams {
	if o == nil {
		var ret []FileShareParams
		return ret
	}
	return o.Share
}

// GetShareOk returns a tuple with the Share field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateRoomRequestDto) GetShareOk() ([]FileShareParams, bool) {
	if o == nil || IsNil(o.Share) {
		return nil, false
	}
	return o.Share, true
}

// HasShare returns a boolean if a field has been set.
func (o *CreateRoomRequestDto) IsShareSet() bool {
	if o != nil && !IsNil(o.Share) {
		return true
	}

	return false
}

// SetShare gets a reference to the given []FileShareParams and assigns it to the Share field.
func (o *CreateRoomRequestDto) SetShare(v []FileShareParams) {
	o.Share = v
}

// GetChatSettings returns the ChatSettings field value if set, zero value otherwise.
func (o *CreateRoomRequestDto) GetChatSettings() ChatSettings {
	if o == nil || IsNil(o.ChatSettings) {
		var ret ChatSettings
		return ret
	}
	return *o.ChatSettings
}

// GetChatSettingsOk returns a tuple with the ChatSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRoomRequestDto) GetChatSettingsOk() (*ChatSettings, bool) {
	if o == nil || IsNil(o.ChatSettings) {
		return nil, false
	}
	return o.ChatSettings, true
}

// HasChatSettings returns a boolean if a field has been set.
func (o *CreateRoomRequestDto) IsChatSettingsSet() bool {
	if o != nil && !IsNil(o.ChatSettings) {
		return true
	}

	return false
}

// SetChatSettings gets a reference to the given ChatSettings and assigns it to the ChatSettings field.
func (o *CreateRoomRequestDto) SetChatSettings(v ChatSettings) {
	o.ChatSettings = &v
}

// GetSendFormToExternalDB returns the SendFormToExternalDB field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateRoomRequestDto) GetSendFormToExternalDB() bool {
	if o == nil || IsNil(o.SendFormToExternalDB.Get()) {
		var ret bool
		return ret
	}
	return *o.SendFormToExternalDB.Get()
}

// GetSendFormToExternalDBOk returns a tuple with the SendFormToExternalDB field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateRoomRequestDto) GetSendFormToExternalDBOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SendFormToExternalDB.Get(), o.SendFormToExternalDB.IsSet()
}

// HasSendFormToExternalDB returns a boolean if a field has been set.
func (o *CreateRoomRequestDto) IsSendFormToExternalDBSet() bool {
	if o != nil && o.SendFormToExternalDB.IsSet() {
		return true
	}

	return false
}

// SetSendFormToExternalDB gets a reference to the given NullableBool and assigns it to the SendFormToExternalDB field.
func (o *CreateRoomRequestDto) SetSendFormToExternalDB(v bool) {
	o.SendFormToExternalDB.Set(&v)
}
// SetSendFormToExternalDBNil sets the value for SendFormToExternalDB to be an explicit nil
func (o *CreateRoomRequestDto) SetSendFormToExternalDBNil() {
	o.SendFormToExternalDB.Set(nil)
}

// UnsetSendFormToExternalDB ensures that no value is present for SendFormToExternalDB, not even an explicit nil
func (o *CreateRoomRequestDto) UnsetSendFormToExternalDB() {
	o.SendFormToExternalDB.Unset()
}

// GetSaveFormAsXLSX returns the SaveFormAsXLSX field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateRoomRequestDto) GetSaveFormAsXLSX() bool {
	if o == nil || IsNil(o.SaveFormAsXLSX.Get()) {
		var ret bool
		return ret
	}
	return *o.SaveFormAsXLSX.Get()
}

// GetSaveFormAsXLSXOk returns a tuple with the SaveFormAsXLSX field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateRoomRequestDto) GetSaveFormAsXLSXOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SaveFormAsXLSX.Get(), o.SaveFormAsXLSX.IsSet()
}

// HasSaveFormAsXLSX returns a boolean if a field has been set.
func (o *CreateRoomRequestDto) IsSaveFormAsXLSXSet() bool {
	if o != nil && o.SaveFormAsXLSX.IsSet() {
		return true
	}

	return false
}

// SetSaveFormAsXLSX gets a reference to the given NullableBool and assigns it to the SaveFormAsXLSX field.
func (o *CreateRoomRequestDto) SetSaveFormAsXLSX(v bool) {
	o.SaveFormAsXLSX.Set(&v)
}
// SetSaveFormAsXLSXNil sets the value for SaveFormAsXLSX to be an explicit nil
func (o *CreateRoomRequestDto) SetSaveFormAsXLSXNil() {
	o.SaveFormAsXLSX.Set(nil)
}

// UnsetSaveFormAsXLSX ensures that no value is present for SaveFormAsXLSX, not even an explicit nil
func (o *CreateRoomRequestDto) UnsetSaveFormAsXLSX() {
	o.SaveFormAsXLSX.Unset()
}

func (o CreateRoomRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRoomRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["title"] = o.Title.Get()
	if o.Quota.IsSet() {
		toSerialize["quota"] = o.Quota.Get()
	}
	if o.Indexing.IsSet() {
		toSerialize["indexing"] = o.Indexing.Get()
	}
	if o.DenyDownload.IsSet() {
		toSerialize["denyDownload"] = o.DenyDownload.Get()
	}
	if !IsNil(o.Lifetime) {
		toSerialize["lifetime"] = o.Lifetime
	}
	if !IsNil(o.Watermark) {
		toSerialize["watermark"] = o.Watermark
	}
	if !IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
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
	toSerialize["roomType"] = o.RoomType
	if !IsNil(o.Private) {
		toSerialize["private"] = o.Private
	}
	if o.Share != nil {
		toSerialize["share"] = o.Share
	}
	if !IsNil(o.ChatSettings) {
		toSerialize["chatSettings"] = o.ChatSettings
	}
	if o.SendFormToExternalDB.IsSet() {
		toSerialize["sendFormToExternalDB"] = o.SendFormToExternalDB.Get()
	}
	if o.SaveFormAsXLSX.IsSet() {
		toSerialize["saveFormAsXLSX"] = o.SaveFormAsXLSX.Get()
	}
	return toSerialize, nil
}

func (o *CreateRoomRequestDto) UnmarshalJSON(data []byte) (err error) {
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

	varCreateRoomRequestDto := _CreateRoomRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varCreateRoomRequestDto)

	if err != nil {
		return err
	}

	*o = CreateRoomRequestDto(varCreateRoomRequestDto)

	return err
}

type NullableCreateRoomRequestDto struct {
	value *CreateRoomRequestDto
	isSet bool
}

func (v NullableCreateRoomRequestDto) Get() *CreateRoomRequestDto {
	return v.value
}

func (v *NullableCreateRoomRequestDto) Set(val *CreateRoomRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRoomRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRoomRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRoomRequestDto(val *CreateRoomRequestDto) *NullableCreateRoomRequestDto {
	return &NullableCreateRoomRequestDto{value: val, isSet: true}
}

func (v NullableCreateRoomRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRoomRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

