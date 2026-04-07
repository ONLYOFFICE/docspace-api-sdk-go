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

// checks if the CreateAgentRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAgentRequestDto{}

// CreateAgentRequestDto Request to create a new AI agent room.
type CreateAgentRequestDto struct {
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
	// Specifies whether the room to be created is private or not.
	Private *bool `json:"private,omitempty"`
	// The collection of sharing parameters.
	Share []FileShareParams `json:"share,omitempty"`
	ChatSettings ChatSettings `json:"chatSettings"`
	// Specifies whether to attach default tools to the agent or not.
	AttachDefaultTools *bool `json:"attachDefaultTools,omitempty"`
}

type _CreateAgentRequestDto CreateAgentRequestDto

// NewCreateAgentRequestDto instantiates a new CreateAgentRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAgentRequestDto(title NullableString, chatSettings ChatSettings) *CreateAgentRequestDto {
	this := CreateAgentRequestDto{}
	this.Title = title
	this.ChatSettings = chatSettings
	return &this
}

// NewCreateAgentRequestDtoWithDefaults instantiates a new CreateAgentRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAgentRequestDtoWithDefaults() *CreateAgentRequestDto {
	this := CreateAgentRequestDto{}
	return &this
}

// GetTitle returns the Title field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreateAgentRequestDto) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}

	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAgentRequestDto) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// SetTitle sets field value
func (o *CreateAgentRequestDto) SetTitle(v string) {
	o.Title.Set(&v)
}

// GetQuota returns the Quota field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAgentRequestDto) GetQuota() int64 {
	if o == nil || IsNil(o.Quota.Get()) {
		var ret int64
		return ret
	}
	return *o.Quota.Get()
}

// GetQuotaOk returns a tuple with the Quota field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAgentRequestDto) GetQuotaOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Quota.Get(), o.Quota.IsSet()
}

// HasQuota returns a boolean if a field has been set.
func (o *CreateAgentRequestDto) IsQuotaSet() bool {
	if o != nil && o.Quota.IsSet() {
		return true
	}

	return false
}

// SetQuota gets a reference to the given NullableInt64 and assigns it to the Quota field.
func (o *CreateAgentRequestDto) SetQuota(v int64) {
	o.Quota.Set(&v)
}
// SetQuotaNil sets the value for Quota to be an explicit nil
func (o *CreateAgentRequestDto) SetQuotaNil() {
	o.Quota.Set(nil)
}

// UnsetQuota ensures that no value is present for Quota, not even an explicit nil
func (o *CreateAgentRequestDto) UnsetQuota() {
	o.Quota.Unset()
}

// GetIndexing returns the Indexing field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAgentRequestDto) GetIndexing() bool {
	if o == nil || IsNil(o.Indexing.Get()) {
		var ret bool
		return ret
	}
	return *o.Indexing.Get()
}

// GetIndexingOk returns a tuple with the Indexing field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAgentRequestDto) GetIndexingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Indexing.Get(), o.Indexing.IsSet()
}

// HasIndexing returns a boolean if a field has been set.
func (o *CreateAgentRequestDto) IsIndexingSet() bool {
	if o != nil && o.Indexing.IsSet() {
		return true
	}

	return false
}

// SetIndexing gets a reference to the given NullableBool and assigns it to the Indexing field.
func (o *CreateAgentRequestDto) SetIndexing(v bool) {
	o.Indexing.Set(&v)
}
// SetIndexingNil sets the value for Indexing to be an explicit nil
func (o *CreateAgentRequestDto) SetIndexingNil() {
	o.Indexing.Set(nil)
}

// UnsetIndexing ensures that no value is present for Indexing, not even an explicit nil
func (o *CreateAgentRequestDto) UnsetIndexing() {
	o.Indexing.Unset()
}

// GetDenyDownload returns the DenyDownload field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAgentRequestDto) GetDenyDownload() bool {
	if o == nil || IsNil(o.DenyDownload.Get()) {
		var ret bool
		return ret
	}
	return *o.DenyDownload.Get()
}

// GetDenyDownloadOk returns a tuple with the DenyDownload field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAgentRequestDto) GetDenyDownloadOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DenyDownload.Get(), o.DenyDownload.IsSet()
}

// HasDenyDownload returns a boolean if a field has been set.
func (o *CreateAgentRequestDto) IsDenyDownloadSet() bool {
	if o != nil && o.DenyDownload.IsSet() {
		return true
	}

	return false
}

// SetDenyDownload gets a reference to the given NullableBool and assigns it to the DenyDownload field.
func (o *CreateAgentRequestDto) SetDenyDownload(v bool) {
	o.DenyDownload.Set(&v)
}
// SetDenyDownloadNil sets the value for DenyDownload to be an explicit nil
func (o *CreateAgentRequestDto) SetDenyDownloadNil() {
	o.DenyDownload.Set(nil)
}

// UnsetDenyDownload ensures that no value is present for DenyDownload, not even an explicit nil
func (o *CreateAgentRequestDto) UnsetDenyDownload() {
	o.DenyDownload.Unset()
}

// GetLifetime returns the Lifetime field value if set, zero value otherwise.
func (o *CreateAgentRequestDto) GetLifetime() RoomDataLifetimeDto {
	if o == nil || IsNil(o.Lifetime) {
		var ret RoomDataLifetimeDto
		return ret
	}
	return *o.Lifetime
}

// GetLifetimeOk returns a tuple with the Lifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAgentRequestDto) GetLifetimeOk() (*RoomDataLifetimeDto, bool) {
	if o == nil || IsNil(o.Lifetime) {
		return nil, false
	}
	return o.Lifetime, true
}

// HasLifetime returns a boolean if a field has been set.
func (o *CreateAgentRequestDto) IsLifetimeSet() bool {
	if o != nil && !IsNil(o.Lifetime) {
		return true
	}

	return false
}

// SetLifetime gets a reference to the given RoomDataLifetimeDto and assigns it to the Lifetime field.
func (o *CreateAgentRequestDto) SetLifetime(v RoomDataLifetimeDto) {
	o.Lifetime = &v
}

// GetWatermark returns the Watermark field value if set, zero value otherwise.
func (o *CreateAgentRequestDto) GetWatermark() WatermarkRequestDto {
	if o == nil || IsNil(o.Watermark) {
		var ret WatermarkRequestDto
		return ret
	}
	return *o.Watermark
}

// GetWatermarkOk returns a tuple with the Watermark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAgentRequestDto) GetWatermarkOk() (*WatermarkRequestDto, bool) {
	if o == nil || IsNil(o.Watermark) {
		return nil, false
	}
	return o.Watermark, true
}

// HasWatermark returns a boolean if a field has been set.
func (o *CreateAgentRequestDto) IsWatermarkSet() bool {
	if o != nil && !IsNil(o.Watermark) {
		return true
	}

	return false
}

// SetWatermark gets a reference to the given WatermarkRequestDto and assigns it to the Watermark field.
func (o *CreateAgentRequestDto) SetWatermark(v WatermarkRequestDto) {
	o.Watermark = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *CreateAgentRequestDto) GetLogo() LogoRequest {
	if o == nil || IsNil(o.Logo) {
		var ret LogoRequest
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAgentRequestDto) GetLogoOk() (*LogoRequest, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *CreateAgentRequestDto) IsLogoSet() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given LogoRequest and assigns it to the Logo field.
func (o *CreateAgentRequestDto) SetLogo(v LogoRequest) {
	o.Logo = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAgentRequestDto) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAgentRequestDto) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateAgentRequestDto) IsTagsSet() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *CreateAgentRequestDto) SetTags(v []string) {
	o.Tags = v
}

// GetColor returns the Color field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAgentRequestDto) GetColor() string {
	if o == nil || IsNil(o.Color.Get()) {
		var ret string
		return ret
	}
	return *o.Color.Get()
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAgentRequestDto) GetColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Color.Get(), o.Color.IsSet()
}

// HasColor returns a boolean if a field has been set.
func (o *CreateAgentRequestDto) IsColorSet() bool {
	if o != nil && o.Color.IsSet() {
		return true
	}

	return false
}

// SetColor gets a reference to the given NullableString and assigns it to the Color field.
func (o *CreateAgentRequestDto) SetColor(v string) {
	o.Color.Set(&v)
}
// SetColorNil sets the value for Color to be an explicit nil
func (o *CreateAgentRequestDto) SetColorNil() {
	o.Color.Set(nil)
}

// UnsetColor ensures that no value is present for Color, not even an explicit nil
func (o *CreateAgentRequestDto) UnsetColor() {
	o.Color.Unset()
}

// GetCover returns the Cover field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAgentRequestDto) GetCover() string {
	if o == nil || IsNil(o.Cover.Get()) {
		var ret string
		return ret
	}
	return *o.Cover.Get()
}

// GetCoverOk returns a tuple with the Cover field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAgentRequestDto) GetCoverOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cover.Get(), o.Cover.IsSet()
}

// HasCover returns a boolean if a field has been set.
func (o *CreateAgentRequestDto) IsCoverSet() bool {
	if o != nil && o.Cover.IsSet() {
		return true
	}

	return false
}

// SetCover gets a reference to the given NullableString and assigns it to the Cover field.
func (o *CreateAgentRequestDto) SetCover(v string) {
	o.Cover.Set(&v)
}
// SetCoverNil sets the value for Cover to be an explicit nil
func (o *CreateAgentRequestDto) SetCoverNil() {
	o.Cover.Set(nil)
}

// UnsetCover ensures that no value is present for Cover, not even an explicit nil
func (o *CreateAgentRequestDto) UnsetCover() {
	o.Cover.Unset()
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *CreateAgentRequestDto) GetPrivate() bool {
	if o == nil || IsNil(o.Private) {
		var ret bool
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAgentRequestDto) GetPrivateOk() (*bool, bool) {
	if o == nil || IsNil(o.Private) {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *CreateAgentRequestDto) IsPrivateSet() bool {
	if o != nil && !IsNil(o.Private) {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given bool and assigns it to the Private field.
func (o *CreateAgentRequestDto) SetPrivate(v bool) {
	o.Private = &v
}

// GetShare returns the Share field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAgentRequestDto) GetShare() []FileShareParams {
	if o == nil {
		var ret []FileShareParams
		return ret
	}
	return o.Share
}

// GetShareOk returns a tuple with the Share field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAgentRequestDto) GetShareOk() ([]FileShareParams, bool) {
	if o == nil || IsNil(o.Share) {
		return nil, false
	}
	return o.Share, true
}

// HasShare returns a boolean if a field has been set.
func (o *CreateAgentRequestDto) IsShareSet() bool {
	if o != nil && !IsNil(o.Share) {
		return true
	}

	return false
}

// SetShare gets a reference to the given []FileShareParams and assigns it to the Share field.
func (o *CreateAgentRequestDto) SetShare(v []FileShareParams) {
	o.Share = v
}

// GetChatSettings returns the ChatSettings field value
func (o *CreateAgentRequestDto) GetChatSettings() ChatSettings {
	if o == nil {
		var ret ChatSettings
		return ret
	}

	return o.ChatSettings
}

// GetChatSettingsOk returns a tuple with the ChatSettings field value
// and a boolean to check if the value has been set.
func (o *CreateAgentRequestDto) GetChatSettingsOk() (*ChatSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChatSettings, true
}

// SetChatSettings sets field value
func (o *CreateAgentRequestDto) SetChatSettings(v ChatSettings) {
	o.ChatSettings = v
}

// GetAttachDefaultTools returns the AttachDefaultTools field value if set, zero value otherwise.
func (o *CreateAgentRequestDto) GetAttachDefaultTools() bool {
	if o == nil || IsNil(o.AttachDefaultTools) {
		var ret bool
		return ret
	}
	return *o.AttachDefaultTools
}

// GetAttachDefaultToolsOk returns a tuple with the AttachDefaultTools field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAgentRequestDto) GetAttachDefaultToolsOk() (*bool, bool) {
	if o == nil || IsNil(o.AttachDefaultTools) {
		return nil, false
	}
	return o.AttachDefaultTools, true
}

// HasAttachDefaultTools returns a boolean if a field has been set.
func (o *CreateAgentRequestDto) IsAttachDefaultToolsSet() bool {
	if o != nil && !IsNil(o.AttachDefaultTools) {
		return true
	}

	return false
}

// SetAttachDefaultTools gets a reference to the given bool and assigns it to the AttachDefaultTools field.
func (o *CreateAgentRequestDto) SetAttachDefaultTools(v bool) {
	o.AttachDefaultTools = &v
}

func (o CreateAgentRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAgentRequestDto) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Private) {
		toSerialize["private"] = o.Private
	}
	if o.Share != nil {
		toSerialize["share"] = o.Share
	}
	toSerialize["chatSettings"] = o.ChatSettings
	if !IsNil(o.AttachDefaultTools) {
		toSerialize["attachDefaultTools"] = o.AttachDefaultTools
	}
	return toSerialize, nil
}

func (o *CreateAgentRequestDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
		"chatSettings",
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

	varCreateAgentRequestDto := _CreateAgentRequestDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varCreateAgentRequestDto)

	if err != nil {
		return err
	}

	*o = CreateAgentRequestDto(varCreateAgentRequestDto)

	return err
}

type NullableCreateAgentRequestDto struct {
	value *CreateAgentRequestDto
	isSet bool
}

func (v NullableCreateAgentRequestDto) Get() *CreateAgentRequestDto {
	return v.value
}

func (v *NullableCreateAgentRequestDto) Set(val *CreateAgentRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAgentRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAgentRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAgentRequestDto(val *CreateAgentRequestDto) *NullableCreateAgentRequestDto {
	return &NullableCreateAgentRequestDto{value: val, isSet: true}
}

func (v NullableCreateAgentRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAgentRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

