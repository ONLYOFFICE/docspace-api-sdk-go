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

// checks if the AiProviderDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiProviderDto{}

// AiProviderDto AI provider details.
type AiProviderDto struct {
	// AI provider identifier.
	Id *int32 `json:"id,omitempty"`
	// AI provider display title.
	Title NullableString `json:"title"`
	Type *ProviderType `json:"type,omitempty"`
	// API endpoint URL for the AI provider.
	Url NullableString `json:"url,omitempty"`
	CreatedOn ApiDateTime `json:"createdOn"`
	ModifiedOn ApiDateTime `json:"modifiedOn"`
	// Indicates whether the provider's API key needs to be reset.
	NeedReset *bool `json:"needReset,omitempty"`
	// Indicates whether this provider is the default provider for the tenant.
	IsDefault *bool `json:"isDefault,omitempty"`
}

type _AiProviderDto AiProviderDto

// NewAiProviderDto instantiates a new AiProviderDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiProviderDto(title NullableString, createdOn ApiDateTime, modifiedOn ApiDateTime) *AiProviderDto {
	this := AiProviderDto{}
	this.Title = title
	this.CreatedOn = createdOn
	this.ModifiedOn = modifiedOn
	return &this
}

// NewAiProviderDtoWithDefaults instantiates a new AiProviderDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiProviderDtoWithDefaults() *AiProviderDto {
	this := AiProviderDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AiProviderDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiProviderDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AiProviderDto) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AiProviderDto) SetId(v int32) {
	o.Id = &v
}

// GetTitle returns the Title field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AiProviderDto) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}

	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiProviderDto) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// SetTitle sets field value
func (o *AiProviderDto) SetTitle(v string) {
	o.Title.Set(&v)
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AiProviderDto) GetType() ProviderType {
	if o == nil || IsNil(o.Type) {
		var ret ProviderType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiProviderDto) GetTypeOk() (*ProviderType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AiProviderDto) IsTypeSet() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ProviderType and assigns it to the Type field.
func (o *AiProviderDto) SetType(v ProviderType) {
	o.Type = &v
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AiProviderDto) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AiProviderDto) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *AiProviderDto) IsUrlSet() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *AiProviderDto) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *AiProviderDto) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *AiProviderDto) UnsetUrl() {
	o.Url.Unset()
}

// GetCreatedOn returns the CreatedOn field value
func (o *AiProviderDto) GetCreatedOn() ApiDateTime {
	if o == nil {
		var ret ApiDateTime
		return ret
	}

	return o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value
// and a boolean to check if the value has been set.
func (o *AiProviderDto) GetCreatedOnOk() (*ApiDateTime, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedOn, true
}

// SetCreatedOn sets field value
func (o *AiProviderDto) SetCreatedOn(v ApiDateTime) {
	o.CreatedOn = v
}

// GetModifiedOn returns the ModifiedOn field value
func (o *AiProviderDto) GetModifiedOn() ApiDateTime {
	if o == nil {
		var ret ApiDateTime
		return ret
	}

	return o.ModifiedOn
}

// GetModifiedOnOk returns a tuple with the ModifiedOn field value
// and a boolean to check if the value has been set.
func (o *AiProviderDto) GetModifiedOnOk() (*ApiDateTime, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedOn, true
}

// SetModifiedOn sets field value
func (o *AiProviderDto) SetModifiedOn(v ApiDateTime) {
	o.ModifiedOn = v
}

// GetNeedReset returns the NeedReset field value if set, zero value otherwise.
func (o *AiProviderDto) GetNeedReset() bool {
	if o == nil || IsNil(o.NeedReset) {
		var ret bool
		return ret
	}
	return *o.NeedReset
}

// GetNeedResetOk returns a tuple with the NeedReset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiProviderDto) GetNeedResetOk() (*bool, bool) {
	if o == nil || IsNil(o.NeedReset) {
		return nil, false
	}
	return o.NeedReset, true
}

// HasNeedReset returns a boolean if a field has been set.
func (o *AiProviderDto) IsNeedResetSet() bool {
	if o != nil && !IsNil(o.NeedReset) {
		return true
	}

	return false
}

// SetNeedReset gets a reference to the given bool and assigns it to the NeedReset field.
func (o *AiProviderDto) SetNeedReset(v bool) {
	o.NeedReset = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *AiProviderDto) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiProviderDto) GetIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *AiProviderDto) IsIsDefaultSet() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *AiProviderDto) SetIsDefault(v bool) {
	o.IsDefault = &v
}

func (o AiProviderDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiProviderDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["title"] = o.Title.Get()
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	toSerialize["createdOn"] = o.CreatedOn
	toSerialize["modifiedOn"] = o.ModifiedOn
	if !IsNil(o.NeedReset) {
		toSerialize["needReset"] = o.NeedReset
	}
	if !IsNil(o.IsDefault) {
		toSerialize["isDefault"] = o.IsDefault
	}
	return toSerialize, nil
}

func (o *AiProviderDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
		"createdOn",
		"modifiedOn",
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

	varAiProviderDto := _AiProviderDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiProviderDto)

	if err != nil {
		return err
	}

	*o = AiProviderDto(varAiProviderDto)

	return err
}

type NullableAiProviderDto struct {
	value *AiProviderDto
	isSet bool
}

func (v NullableAiProviderDto) Get() *AiProviderDto {
	return v.value
}

func (v *NullableAiProviderDto) Set(val *AiProviderDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAiProviderDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAiProviderDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiProviderDto(val *AiProviderDto) *NullableAiProviderDto {
	return &NullableAiProviderDto{value: val, isSet: true}
}

func (v NullableAiProviderDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiProviderDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

