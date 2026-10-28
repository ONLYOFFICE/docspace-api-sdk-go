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

// checks if the AiAgentsUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiAgentsUpdateRequest{}

// AiAgentsUpdateRequest struct for AiAgentsUpdateRequest
type AiAgentsUpdateRequest struct {
	// Profile id to rebind (optional).
	ProfileId *string `json:"profileId,omitempty"`
	// Chat settings (`ChatSettings`); requires a valid provider/model.
	ChatSettings map[string]interface{} `json:"chatSettings,omitempty"`
	// Whether form results are sent to an external DB.
	SendFormToExternalDB *bool `json:"sendFormToExternalDB,omitempty"`
	// Whether forms are saved as XLSX.
	SaveFormAsXLSX *bool `json:"saveFormAsXLSX,omitempty"`
	// Agent (room) title.
	Title *string `json:"title,omitempty"`
	// Room quota in bytes.
	Quota *float32 `json:"quota,omitempty"`
	// Whether room content is indexed for search.
	Indexing *bool `json:"indexing,omitempty"`
	// Whether downloading room content is denied.
	DenyDownload *bool `json:"denyDownload,omitempty"`
	// Room data lifetime policy (`RoomDataLifetimeDto`).
	Lifetime map[string]interface{} `json:"lifetime,omitempty"`
	// Watermark settings (`WatermarkRequestDto`).
	Watermark map[string]interface{} `json:"watermark,omitempty"`
	// Room logo (`LogoRequest`).
	Logo map[string]interface{} `json:"logo,omitempty"`
	// Room tags.
	Tags []string `json:"tags,omitempty"`
	// Room accent color.
	Color *string `json:"color,omitempty"`
	// Room cover image id.
	Cover *string `json:"cover,omitempty"`
}

// NewAiAgentsUpdateRequest instantiates a new AiAgentsUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiAgentsUpdateRequest() *AiAgentsUpdateRequest {
	this := AiAgentsUpdateRequest{}
	return &this
}

// NewAiAgentsUpdateRequestWithDefaults instantiates a new AiAgentsUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiAgentsUpdateRequestWithDefaults() *AiAgentsUpdateRequest {
	this := AiAgentsUpdateRequest{}
	return &this
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *AiAgentsUpdateRequest) GetProfileId() string {
	if o == nil || IsNil(o.ProfileId) {
		var ret string
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentsUpdateRequest) GetProfileIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProfileId) {
		return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *AiAgentsUpdateRequest) IsProfileIdSet() bool {
	if o != nil && !IsNil(o.ProfileId) {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given string and assigns it to the ProfileId field.
func (o *AiAgentsUpdateRequest) SetProfileId(v string) {
	o.ProfileId = &v
}

// GetChatSettings returns the ChatSettings field value if set, zero value otherwise.
func (o *AiAgentsUpdateRequest) GetChatSettings() map[string]interface{} {
	if o == nil || IsNil(o.ChatSettings) {
		var ret map[string]interface{}
		return ret
	}
	return o.ChatSettings
}

// GetChatSettingsOk returns a tuple with the ChatSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentsUpdateRequest) GetChatSettingsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ChatSettings) {
		return map[string]interface{}{}, false
	}
	return o.ChatSettings, true
}

// HasChatSettings returns a boolean if a field has been set.
func (o *AiAgentsUpdateRequest) IsChatSettingsSet() bool {
	if o != nil && !IsNil(o.ChatSettings) {
		return true
	}

	return false
}

// SetChatSettings gets a reference to the given map[string]interface{} and assigns it to the ChatSettings field.
func (o *AiAgentsUpdateRequest) SetChatSettings(v map[string]interface{}) {
	o.ChatSettings = v
}

// GetSendFormToExternalDB returns the SendFormToExternalDB field value if set, zero value otherwise.
func (o *AiAgentsUpdateRequest) GetSendFormToExternalDB() bool {
	if o == nil || IsNil(o.SendFormToExternalDB) {
		var ret bool
		return ret
	}
	return *o.SendFormToExternalDB
}

// GetSendFormToExternalDBOk returns a tuple with the SendFormToExternalDB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentsUpdateRequest) GetSendFormToExternalDBOk() (*bool, bool) {
	if o == nil || IsNil(o.SendFormToExternalDB) {
		return nil, false
	}
	return o.SendFormToExternalDB, true
}

// HasSendFormToExternalDB returns a boolean if a field has been set.
func (o *AiAgentsUpdateRequest) IsSendFormToExternalDBSet() bool {
	if o != nil && !IsNil(o.SendFormToExternalDB) {
		return true
	}

	return false
}

// SetSendFormToExternalDB gets a reference to the given bool and assigns it to the SendFormToExternalDB field.
func (o *AiAgentsUpdateRequest) SetSendFormToExternalDB(v bool) {
	o.SendFormToExternalDB = &v
}

// GetSaveFormAsXLSX returns the SaveFormAsXLSX field value if set, zero value otherwise.
func (o *AiAgentsUpdateRequest) GetSaveFormAsXLSX() bool {
	if o == nil || IsNil(o.SaveFormAsXLSX) {
		var ret bool
		return ret
	}
	return *o.SaveFormAsXLSX
}

// GetSaveFormAsXLSXOk returns a tuple with the SaveFormAsXLSX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentsUpdateRequest) GetSaveFormAsXLSXOk() (*bool, bool) {
	if o == nil || IsNil(o.SaveFormAsXLSX) {
		return nil, false
	}
	return o.SaveFormAsXLSX, true
}

// HasSaveFormAsXLSX returns a boolean if a field has been set.
func (o *AiAgentsUpdateRequest) IsSaveFormAsXLSXSet() bool {
	if o != nil && !IsNil(o.SaveFormAsXLSX) {
		return true
	}

	return false
}

// SetSaveFormAsXLSX gets a reference to the given bool and assigns it to the SaveFormAsXLSX field.
func (o *AiAgentsUpdateRequest) SetSaveFormAsXLSX(v bool) {
	o.SaveFormAsXLSX = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AiAgentsUpdateRequest) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentsUpdateRequest) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AiAgentsUpdateRequest) IsTitleSet() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AiAgentsUpdateRequest) SetTitle(v string) {
	o.Title = &v
}

// GetQuota returns the Quota field value if set, zero value otherwise.
func (o *AiAgentsUpdateRequest) GetQuota() float32 {
	if o == nil || IsNil(o.Quota) {
		var ret float32
		return ret
	}
	return *o.Quota
}

// GetQuotaOk returns a tuple with the Quota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentsUpdateRequest) GetQuotaOk() (*float32, bool) {
	if o == nil || IsNil(o.Quota) {
		return nil, false
	}
	return o.Quota, true
}

// HasQuota returns a boolean if a field has been set.
func (o *AiAgentsUpdateRequest) IsQuotaSet() bool {
	if o != nil && !IsNil(o.Quota) {
		return true
	}

	return false
}

// SetQuota gets a reference to the given float32 and assigns it to the Quota field.
func (o *AiAgentsUpdateRequest) SetQuota(v float32) {
	o.Quota = &v
}

// GetIndexing returns the Indexing field value if set, zero value otherwise.
func (o *AiAgentsUpdateRequest) GetIndexing() bool {
	if o == nil || IsNil(o.Indexing) {
		var ret bool
		return ret
	}
	return *o.Indexing
}

// GetIndexingOk returns a tuple with the Indexing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentsUpdateRequest) GetIndexingOk() (*bool, bool) {
	if o == nil || IsNil(o.Indexing) {
		return nil, false
	}
	return o.Indexing, true
}

// HasIndexing returns a boolean if a field has been set.
func (o *AiAgentsUpdateRequest) IsIndexingSet() bool {
	if o != nil && !IsNil(o.Indexing) {
		return true
	}

	return false
}

// SetIndexing gets a reference to the given bool and assigns it to the Indexing field.
func (o *AiAgentsUpdateRequest) SetIndexing(v bool) {
	o.Indexing = &v
}

// GetDenyDownload returns the DenyDownload field value if set, zero value otherwise.
func (o *AiAgentsUpdateRequest) GetDenyDownload() bool {
	if o == nil || IsNil(o.DenyDownload) {
		var ret bool
		return ret
	}
	return *o.DenyDownload
}

// GetDenyDownloadOk returns a tuple with the DenyDownload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentsUpdateRequest) GetDenyDownloadOk() (*bool, bool) {
	if o == nil || IsNil(o.DenyDownload) {
		return nil, false
	}
	return o.DenyDownload, true
}

// HasDenyDownload returns a boolean if a field has been set.
func (o *AiAgentsUpdateRequest) IsDenyDownloadSet() bool {
	if o != nil && !IsNil(o.DenyDownload) {
		return true
	}

	return false
}

// SetDenyDownload gets a reference to the given bool and assigns it to the DenyDownload field.
func (o *AiAgentsUpdateRequest) SetDenyDownload(v bool) {
	o.DenyDownload = &v
}

// GetLifetime returns the Lifetime field value if set, zero value otherwise.
func (o *AiAgentsUpdateRequest) GetLifetime() map[string]interface{} {
	if o == nil || IsNil(o.Lifetime) {
		var ret map[string]interface{}
		return ret
	}
	return o.Lifetime
}

// GetLifetimeOk returns a tuple with the Lifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentsUpdateRequest) GetLifetimeOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Lifetime) {
		return map[string]interface{}{}, false
	}
	return o.Lifetime, true
}

// HasLifetime returns a boolean if a field has been set.
func (o *AiAgentsUpdateRequest) IsLifetimeSet() bool {
	if o != nil && !IsNil(o.Lifetime) {
		return true
	}

	return false
}

// SetLifetime gets a reference to the given map[string]interface{} and assigns it to the Lifetime field.
func (o *AiAgentsUpdateRequest) SetLifetime(v map[string]interface{}) {
	o.Lifetime = v
}

// GetWatermark returns the Watermark field value if set, zero value otherwise.
func (o *AiAgentsUpdateRequest) GetWatermark() map[string]interface{} {
	if o == nil || IsNil(o.Watermark) {
		var ret map[string]interface{}
		return ret
	}
	return o.Watermark
}

// GetWatermarkOk returns a tuple with the Watermark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentsUpdateRequest) GetWatermarkOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Watermark) {
		return map[string]interface{}{}, false
	}
	return o.Watermark, true
}

// HasWatermark returns a boolean if a field has been set.
func (o *AiAgentsUpdateRequest) IsWatermarkSet() bool {
	if o != nil && !IsNil(o.Watermark) {
		return true
	}

	return false
}

// SetWatermark gets a reference to the given map[string]interface{} and assigns it to the Watermark field.
func (o *AiAgentsUpdateRequest) SetWatermark(v map[string]interface{}) {
	o.Watermark = v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *AiAgentsUpdateRequest) GetLogo() map[string]interface{} {
	if o == nil || IsNil(o.Logo) {
		var ret map[string]interface{}
		return ret
	}
	return o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentsUpdateRequest) GetLogoOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Logo) {
		return map[string]interface{}{}, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *AiAgentsUpdateRequest) IsLogoSet() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given map[string]interface{} and assigns it to the Logo field.
func (o *AiAgentsUpdateRequest) SetLogo(v map[string]interface{}) {
	o.Logo = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *AiAgentsUpdateRequest) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentsUpdateRequest) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AiAgentsUpdateRequest) IsTagsSet() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *AiAgentsUpdateRequest) SetTags(v []string) {
	o.Tags = v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *AiAgentsUpdateRequest) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentsUpdateRequest) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *AiAgentsUpdateRequest) IsColorSet() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *AiAgentsUpdateRequest) SetColor(v string) {
	o.Color = &v
}

// GetCover returns the Cover field value if set, zero value otherwise.
func (o *AiAgentsUpdateRequest) GetCover() string {
	if o == nil || IsNil(o.Cover) {
		var ret string
		return ret
	}
	return *o.Cover
}

// GetCoverOk returns a tuple with the Cover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentsUpdateRequest) GetCoverOk() (*string, bool) {
	if o == nil || IsNil(o.Cover) {
		return nil, false
	}
	return o.Cover, true
}

// HasCover returns a boolean if a field has been set.
func (o *AiAgentsUpdateRequest) IsCoverSet() bool {
	if o != nil && !IsNil(o.Cover) {
		return true
	}

	return false
}

// SetCover gets a reference to the given string and assigns it to the Cover field.
func (o *AiAgentsUpdateRequest) SetCover(v string) {
	o.Cover = &v
}

func (o AiAgentsUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiAgentsUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProfileId) {
		toSerialize["profileId"] = o.ProfileId
	}
	if !IsNil(o.ChatSettings) {
		toSerialize["chatSettings"] = o.ChatSettings
	}
	if !IsNil(o.SendFormToExternalDB) {
		toSerialize["sendFormToExternalDB"] = o.SendFormToExternalDB
	}
	if !IsNil(o.SaveFormAsXLSX) {
		toSerialize["saveFormAsXLSX"] = o.SaveFormAsXLSX
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Quota) {
		toSerialize["quota"] = o.Quota
	}
	if !IsNil(o.Indexing) {
		toSerialize["indexing"] = o.Indexing
	}
	if !IsNil(o.DenyDownload) {
		toSerialize["denyDownload"] = o.DenyDownload
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
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if !IsNil(o.Cover) {
		toSerialize["cover"] = o.Cover
	}
	return toSerialize, nil
}

type NullableAiAgentsUpdateRequest struct {
	value *AiAgentsUpdateRequest
	isSet bool
}

func (v NullableAiAgentsUpdateRequest) Get() *AiAgentsUpdateRequest {
	return v.value
}

func (v *NullableAiAgentsUpdateRequest) Set(val *AiAgentsUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAiAgentsUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAiAgentsUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiAgentsUpdateRequest(val *AiAgentsUpdateRequest) *NullableAiAgentsUpdateRequest {
	return &NullableAiAgentsUpdateRequest{value: val, isSet: true}
}

func (v NullableAiAgentsUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiAgentsUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

