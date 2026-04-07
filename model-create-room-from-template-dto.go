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

// checks if the CreateRoomFromTemplateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRoomFromTemplateDto{}

// CreateRoomFromTemplateDto The parameters for creating a room from a template.
type CreateRoomFromTemplateDto struct {
	// The template ID from which the room to be created.
	TemplateId int32 `json:"templateId"`
	// The room title.
	Title NullableString `json:"title"`
	Logo *LogoRequest `json:"logo,omitempty"`
	// Specifies whether to copy a logo or not.
	CopyLogo *bool `json:"copyLogo,omitempty"`
	// The collection of tags.
	Tags []string `json:"tags,omitempty"`
	// The color of the room to be created.
	Color NullableString `json:"color,omitempty"`
	// The cover of the room to be created.
	Cover NullableString `json:"cover,omitempty"`
	// The room quota.
	Quota NullableInt64 `json:"quota,omitempty"`
	// Specifies whether to create a room with indexing.
	Indexing NullableBool `json:"indexing,omitempty"`
	// Specifies whether to deny downloads from the room.
	DenyDownload NullableBool `json:"denyDownload,omitempty"`
	Lifetime *RoomDataLifetimeDto `json:"lifetime,omitempty"`
	Watermark *WatermarkRequestDto `json:"watermark,omitempty"`
	// Specifies whether the room to be created is private or not.
	Private NullableBool `json:"private,omitempty"`
}

type _CreateRoomFromTemplateDto CreateRoomFromTemplateDto

// NewCreateRoomFromTemplateDto instantiates a new CreateRoomFromTemplateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRoomFromTemplateDto(templateId int32, title NullableString) *CreateRoomFromTemplateDto {
	this := CreateRoomFromTemplateDto{}
	this.TemplateId = templateId
	this.Title = title
	return &this
}

// NewCreateRoomFromTemplateDtoWithDefaults instantiates a new CreateRoomFromTemplateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRoomFromTemplateDtoWithDefaults() *CreateRoomFromTemplateDto {
	this := CreateRoomFromTemplateDto{}
	return &this
}

// GetTemplateId returns the TemplateId field value
func (o *CreateRoomFromTemplateDto) GetTemplateId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value
// and a boolean to check if the value has been set.
func (o *CreateRoomFromTemplateDto) GetTemplateIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateId, true
}

// SetTemplateId sets field value
func (o *CreateRoomFromTemplateDto) SetTemplateId(v int32) {
	o.TemplateId = v
}

// GetTitle returns the Title field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CreateRoomFromTemplateDto) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}

	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateRoomFromTemplateDto) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// SetTitle sets field value
func (o *CreateRoomFromTemplateDto) SetTitle(v string) {
	o.Title.Set(&v)
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *CreateRoomFromTemplateDto) GetLogo() LogoRequest {
	if o == nil || IsNil(o.Logo) {
		var ret LogoRequest
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRoomFromTemplateDto) GetLogoOk() (*LogoRequest, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *CreateRoomFromTemplateDto) IsLogoSet() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given LogoRequest and assigns it to the Logo field.
func (o *CreateRoomFromTemplateDto) SetLogo(v LogoRequest) {
	o.Logo = &v
}

// GetCopyLogo returns the CopyLogo field value if set, zero value otherwise.
func (o *CreateRoomFromTemplateDto) GetCopyLogo() bool {
	if o == nil || IsNil(o.CopyLogo) {
		var ret bool
		return ret
	}
	return *o.CopyLogo
}

// GetCopyLogoOk returns a tuple with the CopyLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRoomFromTemplateDto) GetCopyLogoOk() (*bool, bool) {
	if o == nil || IsNil(o.CopyLogo) {
		return nil, false
	}
	return o.CopyLogo, true
}

// HasCopyLogo returns a boolean if a field has been set.
func (o *CreateRoomFromTemplateDto) IsCopyLogoSet() bool {
	if o != nil && !IsNil(o.CopyLogo) {
		return true
	}

	return false
}

// SetCopyLogo gets a reference to the given bool and assigns it to the CopyLogo field.
func (o *CreateRoomFromTemplateDto) SetCopyLogo(v bool) {
	o.CopyLogo = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateRoomFromTemplateDto) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateRoomFromTemplateDto) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateRoomFromTemplateDto) IsTagsSet() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *CreateRoomFromTemplateDto) SetTags(v []string) {
	o.Tags = v
}

// GetColor returns the Color field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateRoomFromTemplateDto) GetColor() string {
	if o == nil || IsNil(o.Color.Get()) {
		var ret string
		return ret
	}
	return *o.Color.Get()
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateRoomFromTemplateDto) GetColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Color.Get(), o.Color.IsSet()
}

// HasColor returns a boolean if a field has been set.
func (o *CreateRoomFromTemplateDto) IsColorSet() bool {
	if o != nil && o.Color.IsSet() {
		return true
	}

	return false
}

// SetColor gets a reference to the given NullableString and assigns it to the Color field.
func (o *CreateRoomFromTemplateDto) SetColor(v string) {
	o.Color.Set(&v)
}
// SetColorNil sets the value for Color to be an explicit nil
func (o *CreateRoomFromTemplateDto) SetColorNil() {
	o.Color.Set(nil)
}

// UnsetColor ensures that no value is present for Color, not even an explicit nil
func (o *CreateRoomFromTemplateDto) UnsetColor() {
	o.Color.Unset()
}

// GetCover returns the Cover field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateRoomFromTemplateDto) GetCover() string {
	if o == nil || IsNil(o.Cover.Get()) {
		var ret string
		return ret
	}
	return *o.Cover.Get()
}

// GetCoverOk returns a tuple with the Cover field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateRoomFromTemplateDto) GetCoverOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cover.Get(), o.Cover.IsSet()
}

// HasCover returns a boolean if a field has been set.
func (o *CreateRoomFromTemplateDto) IsCoverSet() bool {
	if o != nil && o.Cover.IsSet() {
		return true
	}

	return false
}

// SetCover gets a reference to the given NullableString and assigns it to the Cover field.
func (o *CreateRoomFromTemplateDto) SetCover(v string) {
	o.Cover.Set(&v)
}
// SetCoverNil sets the value for Cover to be an explicit nil
func (o *CreateRoomFromTemplateDto) SetCoverNil() {
	o.Cover.Set(nil)
}

// UnsetCover ensures that no value is present for Cover, not even an explicit nil
func (o *CreateRoomFromTemplateDto) UnsetCover() {
	o.Cover.Unset()
}

// GetQuota returns the Quota field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateRoomFromTemplateDto) GetQuota() int64 {
	if o == nil || IsNil(o.Quota.Get()) {
		var ret int64
		return ret
	}
	return *o.Quota.Get()
}

// GetQuotaOk returns a tuple with the Quota field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateRoomFromTemplateDto) GetQuotaOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Quota.Get(), o.Quota.IsSet()
}

// HasQuota returns a boolean if a field has been set.
func (o *CreateRoomFromTemplateDto) IsQuotaSet() bool {
	if o != nil && o.Quota.IsSet() {
		return true
	}

	return false
}

// SetQuota gets a reference to the given NullableInt64 and assigns it to the Quota field.
func (o *CreateRoomFromTemplateDto) SetQuota(v int64) {
	o.Quota.Set(&v)
}
// SetQuotaNil sets the value for Quota to be an explicit nil
func (o *CreateRoomFromTemplateDto) SetQuotaNil() {
	o.Quota.Set(nil)
}

// UnsetQuota ensures that no value is present for Quota, not even an explicit nil
func (o *CreateRoomFromTemplateDto) UnsetQuota() {
	o.Quota.Unset()
}

// GetIndexing returns the Indexing field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateRoomFromTemplateDto) GetIndexing() bool {
	if o == nil || IsNil(o.Indexing.Get()) {
		var ret bool
		return ret
	}
	return *o.Indexing.Get()
}

// GetIndexingOk returns a tuple with the Indexing field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateRoomFromTemplateDto) GetIndexingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Indexing.Get(), o.Indexing.IsSet()
}

// HasIndexing returns a boolean if a field has been set.
func (o *CreateRoomFromTemplateDto) IsIndexingSet() bool {
	if o != nil && o.Indexing.IsSet() {
		return true
	}

	return false
}

// SetIndexing gets a reference to the given NullableBool and assigns it to the Indexing field.
func (o *CreateRoomFromTemplateDto) SetIndexing(v bool) {
	o.Indexing.Set(&v)
}
// SetIndexingNil sets the value for Indexing to be an explicit nil
func (o *CreateRoomFromTemplateDto) SetIndexingNil() {
	o.Indexing.Set(nil)
}

// UnsetIndexing ensures that no value is present for Indexing, not even an explicit nil
func (o *CreateRoomFromTemplateDto) UnsetIndexing() {
	o.Indexing.Unset()
}

// GetDenyDownload returns the DenyDownload field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateRoomFromTemplateDto) GetDenyDownload() bool {
	if o == nil || IsNil(o.DenyDownload.Get()) {
		var ret bool
		return ret
	}
	return *o.DenyDownload.Get()
}

// GetDenyDownloadOk returns a tuple with the DenyDownload field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateRoomFromTemplateDto) GetDenyDownloadOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DenyDownload.Get(), o.DenyDownload.IsSet()
}

// HasDenyDownload returns a boolean if a field has been set.
func (o *CreateRoomFromTemplateDto) IsDenyDownloadSet() bool {
	if o != nil && o.DenyDownload.IsSet() {
		return true
	}

	return false
}

// SetDenyDownload gets a reference to the given NullableBool and assigns it to the DenyDownload field.
func (o *CreateRoomFromTemplateDto) SetDenyDownload(v bool) {
	o.DenyDownload.Set(&v)
}
// SetDenyDownloadNil sets the value for DenyDownload to be an explicit nil
func (o *CreateRoomFromTemplateDto) SetDenyDownloadNil() {
	o.DenyDownload.Set(nil)
}

// UnsetDenyDownload ensures that no value is present for DenyDownload, not even an explicit nil
func (o *CreateRoomFromTemplateDto) UnsetDenyDownload() {
	o.DenyDownload.Unset()
}

// GetLifetime returns the Lifetime field value if set, zero value otherwise.
func (o *CreateRoomFromTemplateDto) GetLifetime() RoomDataLifetimeDto {
	if o == nil || IsNil(o.Lifetime) {
		var ret RoomDataLifetimeDto
		return ret
	}
	return *o.Lifetime
}

// GetLifetimeOk returns a tuple with the Lifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRoomFromTemplateDto) GetLifetimeOk() (*RoomDataLifetimeDto, bool) {
	if o == nil || IsNil(o.Lifetime) {
		return nil, false
	}
	return o.Lifetime, true
}

// HasLifetime returns a boolean if a field has been set.
func (o *CreateRoomFromTemplateDto) IsLifetimeSet() bool {
	if o != nil && !IsNil(o.Lifetime) {
		return true
	}

	return false
}

// SetLifetime gets a reference to the given RoomDataLifetimeDto and assigns it to the Lifetime field.
func (o *CreateRoomFromTemplateDto) SetLifetime(v RoomDataLifetimeDto) {
	o.Lifetime = &v
}

// GetWatermark returns the Watermark field value if set, zero value otherwise.
func (o *CreateRoomFromTemplateDto) GetWatermark() WatermarkRequestDto {
	if o == nil || IsNil(o.Watermark) {
		var ret WatermarkRequestDto
		return ret
	}
	return *o.Watermark
}

// GetWatermarkOk returns a tuple with the Watermark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRoomFromTemplateDto) GetWatermarkOk() (*WatermarkRequestDto, bool) {
	if o == nil || IsNil(o.Watermark) {
		return nil, false
	}
	return o.Watermark, true
}

// HasWatermark returns a boolean if a field has been set.
func (o *CreateRoomFromTemplateDto) IsWatermarkSet() bool {
	if o != nil && !IsNil(o.Watermark) {
		return true
	}

	return false
}

// SetWatermark gets a reference to the given WatermarkRequestDto and assigns it to the Watermark field.
func (o *CreateRoomFromTemplateDto) SetWatermark(v WatermarkRequestDto) {
	o.Watermark = &v
}

// GetPrivate returns the Private field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateRoomFromTemplateDto) GetPrivate() bool {
	if o == nil || IsNil(o.Private.Get()) {
		var ret bool
		return ret
	}
	return *o.Private.Get()
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateRoomFromTemplateDto) GetPrivateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Private.Get(), o.Private.IsSet()
}

// HasPrivate returns a boolean if a field has been set.
func (o *CreateRoomFromTemplateDto) IsPrivateSet() bool {
	if o != nil && o.Private.IsSet() {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given NullableBool and assigns it to the Private field.
func (o *CreateRoomFromTemplateDto) SetPrivate(v bool) {
	o.Private.Set(&v)
}
// SetPrivateNil sets the value for Private to be an explicit nil
func (o *CreateRoomFromTemplateDto) SetPrivateNil() {
	o.Private.Set(nil)
}

// UnsetPrivate ensures that no value is present for Private, not even an explicit nil
func (o *CreateRoomFromTemplateDto) UnsetPrivate() {
	o.Private.Unset()
}

func (o CreateRoomFromTemplateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRoomFromTemplateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["templateId"] = o.TemplateId
	toSerialize["title"] = o.Title.Get()
	if !IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}
	if !IsNil(o.CopyLogo) {
		toSerialize["copyLogo"] = o.CopyLogo
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
	if o.Private.IsSet() {
		toSerialize["private"] = o.Private.Get()
	}
	return toSerialize, nil
}

func (o *CreateRoomFromTemplateDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"templateId",
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

	varCreateRoomFromTemplateDto := _CreateRoomFromTemplateDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varCreateRoomFromTemplateDto)

	if err != nil {
		return err
	}

	*o = CreateRoomFromTemplateDto(varCreateRoomFromTemplateDto)

	return err
}

type NullableCreateRoomFromTemplateDto struct {
	value *CreateRoomFromTemplateDto
	isSet bool
}

func (v NullableCreateRoomFromTemplateDto) Get() *CreateRoomFromTemplateDto {
	return v.value
}

func (v *NullableCreateRoomFromTemplateDto) Set(val *CreateRoomFromTemplateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRoomFromTemplateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRoomFromTemplateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRoomFromTemplateDto(val *CreateRoomFromTemplateDto) *NullableCreateRoomFromTemplateDto {
	return &NullableCreateRoomFromTemplateDto{value: val, isSet: true}
}

func (v NullableCreateRoomFromTemplateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRoomFromTemplateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

