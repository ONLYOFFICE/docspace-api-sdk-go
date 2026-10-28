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

// checks if the AiAgentsCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiAgentsCreateRequest{}

// AiAgentsCreateRequest struct for AiAgentsCreateRequest
type AiAgentsCreateRequest struct {
	// Profile id bound to the agent.
	ProfileId string `json:"profileId"`
	// Agent system prompt; stored as the room's `chatSettings.prompt`.
	Prompt string `json:"prompt"`
	// Whether the agent room is private.
	Private *bool `json:"private,omitempty"`
	// Initial share entries (`FileShareParams`).
	Share []map[string]interface{} `json:"share,omitempty"`
	// Whether to attach the default DocSpace MCP tool server.
	AttachDefaultTools *bool `json:"attachDefaultTools,omitempty"`
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

type _AiAgentsCreateRequest AiAgentsCreateRequest

// NewAiAgentsCreateRequest instantiates a new AiAgentsCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiAgentsCreateRequest(profileId string, prompt string) *AiAgentsCreateRequest {
	this := AiAgentsCreateRequest{}
	this.ProfileId = profileId
	this.Prompt = prompt
	return &this
}

// NewAiAgentsCreateRequestWithDefaults instantiates a new AiAgentsCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiAgentsCreateRequestWithDefaults() *AiAgentsCreateRequest {
	this := AiAgentsCreateRequest{}
	return &this
}

// GetProfileId returns the ProfileId field value
func (o *AiAgentsCreateRequest) GetProfileId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value
// and a boolean to check if the value has been set.
func (o *AiAgentsCreateRequest) GetProfileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProfileId, true
}

// SetProfileId sets field value
func (o *AiAgentsCreateRequest) SetProfileId(v string) {
	o.ProfileId = v
}

// GetPrompt returns the Prompt field value
func (o *AiAgentsCreateRequest) GetPrompt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prompt
}

// GetPromptOk returns a tuple with the Prompt field value
// and a boolean to check if the value has been set.
func (o *AiAgentsCreateRequest) GetPromptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prompt, true
}

// SetPrompt sets field value
func (o *AiAgentsCreateRequest) SetPrompt(v string) {
	o.Prompt = v
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *AiAgentsCreateRequest) GetPrivate() bool {
	if o == nil || IsNil(o.Private) {
		var ret bool
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentsCreateRequest) GetPrivateOk() (*bool, bool) {
	if o == nil || IsNil(o.Private) {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *AiAgentsCreateRequest) IsPrivateSet() bool {
	if o != nil && !IsNil(o.Private) {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given bool and assigns it to the Private field.
func (o *AiAgentsCreateRequest) SetPrivate(v bool) {
	o.Private = &v
}

// GetShare returns the Share field value if set, zero value otherwise.
func (o *AiAgentsCreateRequest) GetShare() []map[string]interface{} {
	if o == nil || IsNil(o.Share) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Share
}

// GetShareOk returns a tuple with the Share field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentsCreateRequest) GetShareOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Share) {
		return nil, false
	}
	return o.Share, true
}

// HasShare returns a boolean if a field has been set.
func (o *AiAgentsCreateRequest) IsShareSet() bool {
	if o != nil && !IsNil(o.Share) {
		return true
	}

	return false
}

// SetShare gets a reference to the given []map[string]interface{} and assigns it to the Share field.
func (o *AiAgentsCreateRequest) SetShare(v []map[string]interface{}) {
	o.Share = v
}

// GetAttachDefaultTools returns the AttachDefaultTools field value if set, zero value otherwise.
func (o *AiAgentsCreateRequest) GetAttachDefaultTools() bool {
	if o == nil || IsNil(o.AttachDefaultTools) {
		var ret bool
		return ret
	}
	return *o.AttachDefaultTools
}

// GetAttachDefaultToolsOk returns a tuple with the AttachDefaultTools field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentsCreateRequest) GetAttachDefaultToolsOk() (*bool, bool) {
	if o == nil || IsNil(o.AttachDefaultTools) {
		return nil, false
	}
	return o.AttachDefaultTools, true
}

// HasAttachDefaultTools returns a boolean if a field has been set.
func (o *AiAgentsCreateRequest) IsAttachDefaultToolsSet() bool {
	if o != nil && !IsNil(o.AttachDefaultTools) {
		return true
	}

	return false
}

// SetAttachDefaultTools gets a reference to the given bool and assigns it to the AttachDefaultTools field.
func (o *AiAgentsCreateRequest) SetAttachDefaultTools(v bool) {
	o.AttachDefaultTools = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AiAgentsCreateRequest) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentsCreateRequest) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AiAgentsCreateRequest) IsTitleSet() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AiAgentsCreateRequest) SetTitle(v string) {
	o.Title = &v
}

// GetQuota returns the Quota field value if set, zero value otherwise.
func (o *AiAgentsCreateRequest) GetQuota() float32 {
	if o == nil || IsNil(o.Quota) {
		var ret float32
		return ret
	}
	return *o.Quota
}

// GetQuotaOk returns a tuple with the Quota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentsCreateRequest) GetQuotaOk() (*float32, bool) {
	if o == nil || IsNil(o.Quota) {
		return nil, false
	}
	return o.Quota, true
}

// HasQuota returns a boolean if a field has been set.
func (o *AiAgentsCreateRequest) IsQuotaSet() bool {
	if o != nil && !IsNil(o.Quota) {
		return true
	}

	return false
}

// SetQuota gets a reference to the given float32 and assigns it to the Quota field.
func (o *AiAgentsCreateRequest) SetQuota(v float32) {
	o.Quota = &v
}

// GetIndexing returns the Indexing field value if set, zero value otherwise.
func (o *AiAgentsCreateRequest) GetIndexing() bool {
	if o == nil || IsNil(o.Indexing) {
		var ret bool
		return ret
	}
	return *o.Indexing
}

// GetIndexingOk returns a tuple with the Indexing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentsCreateRequest) GetIndexingOk() (*bool, bool) {
	if o == nil || IsNil(o.Indexing) {
		return nil, false
	}
	return o.Indexing, true
}

// HasIndexing returns a boolean if a field has been set.
func (o *AiAgentsCreateRequest) IsIndexingSet() bool {
	if o != nil && !IsNil(o.Indexing) {
		return true
	}

	return false
}

// SetIndexing gets a reference to the given bool and assigns it to the Indexing field.
func (o *AiAgentsCreateRequest) SetIndexing(v bool) {
	o.Indexing = &v
}

// GetDenyDownload returns the DenyDownload field value if set, zero value otherwise.
func (o *AiAgentsCreateRequest) GetDenyDownload() bool {
	if o == nil || IsNil(o.DenyDownload) {
		var ret bool
		return ret
	}
	return *o.DenyDownload
}

// GetDenyDownloadOk returns a tuple with the DenyDownload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentsCreateRequest) GetDenyDownloadOk() (*bool, bool) {
	if o == nil || IsNil(o.DenyDownload) {
		return nil, false
	}
	return o.DenyDownload, true
}

// HasDenyDownload returns a boolean if a field has been set.
func (o *AiAgentsCreateRequest) IsDenyDownloadSet() bool {
	if o != nil && !IsNil(o.DenyDownload) {
		return true
	}

	return false
}

// SetDenyDownload gets a reference to the given bool and assigns it to the DenyDownload field.
func (o *AiAgentsCreateRequest) SetDenyDownload(v bool) {
	o.DenyDownload = &v
}

// GetLifetime returns the Lifetime field value if set, zero value otherwise.
func (o *AiAgentsCreateRequest) GetLifetime() map[string]interface{} {
	if o == nil || IsNil(o.Lifetime) {
		var ret map[string]interface{}
		return ret
	}
	return o.Lifetime
}

// GetLifetimeOk returns a tuple with the Lifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentsCreateRequest) GetLifetimeOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Lifetime) {
		return map[string]interface{}{}, false
	}
	return o.Lifetime, true
}

// HasLifetime returns a boolean if a field has been set.
func (o *AiAgentsCreateRequest) IsLifetimeSet() bool {
	if o != nil && !IsNil(o.Lifetime) {
		return true
	}

	return false
}

// SetLifetime gets a reference to the given map[string]interface{} and assigns it to the Lifetime field.
func (o *AiAgentsCreateRequest) SetLifetime(v map[string]interface{}) {
	o.Lifetime = v
}

// GetWatermark returns the Watermark field value if set, zero value otherwise.
func (o *AiAgentsCreateRequest) GetWatermark() map[string]interface{} {
	if o == nil || IsNil(o.Watermark) {
		var ret map[string]interface{}
		return ret
	}
	return o.Watermark
}

// GetWatermarkOk returns a tuple with the Watermark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentsCreateRequest) GetWatermarkOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Watermark) {
		return map[string]interface{}{}, false
	}
	return o.Watermark, true
}

// HasWatermark returns a boolean if a field has been set.
func (o *AiAgentsCreateRequest) IsWatermarkSet() bool {
	if o != nil && !IsNil(o.Watermark) {
		return true
	}

	return false
}

// SetWatermark gets a reference to the given map[string]interface{} and assigns it to the Watermark field.
func (o *AiAgentsCreateRequest) SetWatermark(v map[string]interface{}) {
	o.Watermark = v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *AiAgentsCreateRequest) GetLogo() map[string]interface{} {
	if o == nil || IsNil(o.Logo) {
		var ret map[string]interface{}
		return ret
	}
	return o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentsCreateRequest) GetLogoOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Logo) {
		return map[string]interface{}{}, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *AiAgentsCreateRequest) IsLogoSet() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given map[string]interface{} and assigns it to the Logo field.
func (o *AiAgentsCreateRequest) SetLogo(v map[string]interface{}) {
	o.Logo = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *AiAgentsCreateRequest) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentsCreateRequest) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AiAgentsCreateRequest) IsTagsSet() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *AiAgentsCreateRequest) SetTags(v []string) {
	o.Tags = v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *AiAgentsCreateRequest) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentsCreateRequest) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *AiAgentsCreateRequest) IsColorSet() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *AiAgentsCreateRequest) SetColor(v string) {
	o.Color = &v
}

// GetCover returns the Cover field value if set, zero value otherwise.
func (o *AiAgentsCreateRequest) GetCover() string {
	if o == nil || IsNil(o.Cover) {
		var ret string
		return ret
	}
	return *o.Cover
}

// GetCoverOk returns a tuple with the Cover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentsCreateRequest) GetCoverOk() (*string, bool) {
	if o == nil || IsNil(o.Cover) {
		return nil, false
	}
	return o.Cover, true
}

// HasCover returns a boolean if a field has been set.
func (o *AiAgentsCreateRequest) IsCoverSet() bool {
	if o != nil && !IsNil(o.Cover) {
		return true
	}

	return false
}

// SetCover gets a reference to the given string and assigns it to the Cover field.
func (o *AiAgentsCreateRequest) SetCover(v string) {
	o.Cover = &v
}

func (o AiAgentsCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiAgentsCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["profileId"] = o.ProfileId
	toSerialize["prompt"] = o.Prompt
	if !IsNil(o.Private) {
		toSerialize["private"] = o.Private
	}
	if !IsNil(o.Share) {
		toSerialize["share"] = o.Share
	}
	if !IsNil(o.AttachDefaultTools) {
		toSerialize["attachDefaultTools"] = o.AttachDefaultTools
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

func (o *AiAgentsCreateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"profileId",
		"prompt",
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

	varAiAgentsCreateRequest := _AiAgentsCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiAgentsCreateRequest)

	if err != nil {
		return err
	}

	*o = AiAgentsCreateRequest(varAiAgentsCreateRequest)

	return err
}

type NullableAiAgentsCreateRequest struct {
	value *AiAgentsCreateRequest
	isSet bool
}

func (v NullableAiAgentsCreateRequest) Get() *AiAgentsCreateRequest {
	return v.value
}

func (v *NullableAiAgentsCreateRequest) Set(val *AiAgentsCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAiAgentsCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAiAgentsCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiAgentsCreateRequest(val *AiAgentsCreateRequest) *NullableAiAgentsCreateRequest {
	return &NullableAiAgentsCreateRequest{value: val, isSet: true}
}

func (v NullableAiAgentsCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiAgentsCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

