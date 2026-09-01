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

// checks if the TenantQuotaFeatureDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenantQuotaFeatureDto{}

// TenantQuotaFeatureDto The tenant quota feature parameters.
type TenantQuotaFeatureDto struct {
	// The ID of the tenant quota feature.
	Id NullableString `json:"id,omitempty"`
	// The title of the tenant quota feature.
	Title NullableString `json:"title,omitempty"`
	// The image URL of the tenant quota feature.
	Image NullableString `json:"image,omitempty"`
	Value interface{} `json:"value,omitempty"`
	// The type of the tenant quota feature.
	Type NullableString `json:"type,omitempty"`
	// The used space parameters of the tenant quota feature.
	Used *FeatureUsedDto `json:"used,omitempty"`
	// The price title of the tenant quota feature.
	PriceTitle NullableString `json:"priceTitle,omitempty"`
}

// NewTenantQuotaFeatureDto instantiates a new TenantQuotaFeatureDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantQuotaFeatureDto() *TenantQuotaFeatureDto {
	this := TenantQuotaFeatureDto{}
	return &this
}

// NewTenantQuotaFeatureDtoWithDefaults instantiates a new TenantQuotaFeatureDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantQuotaFeatureDtoWithDefaults() *TenantQuotaFeatureDto {
	this := TenantQuotaFeatureDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantQuotaFeatureDto) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantQuotaFeatureDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *TenantQuotaFeatureDto) IsIdSet() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *TenantQuotaFeatureDto) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *TenantQuotaFeatureDto) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *TenantQuotaFeatureDto) UnsetId() {
	o.Id.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantQuotaFeatureDto) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantQuotaFeatureDto) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *TenantQuotaFeatureDto) IsTitleSet() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *TenantQuotaFeatureDto) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *TenantQuotaFeatureDto) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *TenantQuotaFeatureDto) UnsetTitle() {
	o.Title.Unset()
}

// GetImage returns the Image field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantQuotaFeatureDto) GetImage() string {
	if o == nil || IsNil(o.Image.Get()) {
		var ret string
		return ret
	}
	return *o.Image.Get()
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantQuotaFeatureDto) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Image.Get(), o.Image.IsSet()
}

// HasImage returns a boolean if a field has been set.
func (o *TenantQuotaFeatureDto) IsImageSet() bool {
	if o != nil && o.Image.IsSet() {
		return true
	}

	return false
}

// SetImage gets a reference to the given NullableString and assigns it to the Image field.
func (o *TenantQuotaFeatureDto) SetImage(v string) {
	o.Image.Set(&v)
}
// SetImageNil sets the value for Image to be an explicit nil
func (o *TenantQuotaFeatureDto) SetImageNil() {
	o.Image.Set(nil)
}

// UnsetImage ensures that no value is present for Image, not even an explicit nil
func (o *TenantQuotaFeatureDto) UnsetImage() {
	o.Image.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantQuotaFeatureDto) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantQuotaFeatureDto) GetValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TenantQuotaFeatureDto) IsValueSet() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *TenantQuotaFeatureDto) SetValue(v interface{}) {
	o.Value = v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantQuotaFeatureDto) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantQuotaFeatureDto) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *TenantQuotaFeatureDto) IsTypeSet() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *TenantQuotaFeatureDto) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *TenantQuotaFeatureDto) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *TenantQuotaFeatureDto) UnsetType() {
	o.Type.Unset()
}

// GetUsed returns the Used field value if set, zero value otherwise.
func (o *TenantQuotaFeatureDto) GetUsed() FeatureUsedDto {
	if o == nil || IsNil(o.Used) {
		var ret FeatureUsedDto
		return ret
	}
	return *o.Used
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantQuotaFeatureDto) GetUsedOk() (*FeatureUsedDto, bool) {
	if o == nil || IsNil(o.Used) {
		return nil, false
	}
	return o.Used, true
}

// HasUsed returns a boolean if a field has been set.
func (o *TenantQuotaFeatureDto) IsUsedSet() bool {
	if o != nil && !IsNil(o.Used) {
		return true
	}

	return false
}

// SetUsed gets a reference to the given FeatureUsedDto and assigns it to the Used field.
func (o *TenantQuotaFeatureDto) SetUsed(v FeatureUsedDto) {
	o.Used = &v
}

// GetPriceTitle returns the PriceTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantQuotaFeatureDto) GetPriceTitle() string {
	if o == nil || IsNil(o.PriceTitle.Get()) {
		var ret string
		return ret
	}
	return *o.PriceTitle.Get()
}

// GetPriceTitleOk returns a tuple with the PriceTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantQuotaFeatureDto) GetPriceTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PriceTitle.Get(), o.PriceTitle.IsSet()
}

// HasPriceTitle returns a boolean if a field has been set.
func (o *TenantQuotaFeatureDto) IsPriceTitleSet() bool {
	if o != nil && o.PriceTitle.IsSet() {
		return true
	}

	return false
}

// SetPriceTitle gets a reference to the given NullableString and assigns it to the PriceTitle field.
func (o *TenantQuotaFeatureDto) SetPriceTitle(v string) {
	o.PriceTitle.Set(&v)
}
// SetPriceTitleNil sets the value for PriceTitle to be an explicit nil
func (o *TenantQuotaFeatureDto) SetPriceTitleNil() {
	o.PriceTitle.Set(nil)
}

// UnsetPriceTitle ensures that no value is present for PriceTitle, not even an explicit nil
func (o *TenantQuotaFeatureDto) UnsetPriceTitle() {
	o.PriceTitle.Unset()
}

func (o TenantQuotaFeatureDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantQuotaFeatureDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Image.IsSet() {
		toSerialize["image"] = o.Image.Get()
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if !IsNil(o.Used) {
		toSerialize["used"] = o.Used
	}
	if o.PriceTitle.IsSet() {
		toSerialize["priceTitle"] = o.PriceTitle.Get()
	}
	return toSerialize, nil
}

type NullableTenantQuotaFeatureDto struct {
	value *TenantQuotaFeatureDto
	isSet bool
}

func (v NullableTenantQuotaFeatureDto) Get() *TenantQuotaFeatureDto {
	return v.value
}

func (v *NullableTenantQuotaFeatureDto) Set(val *TenantQuotaFeatureDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantQuotaFeatureDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantQuotaFeatureDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantQuotaFeatureDto(val *TenantQuotaFeatureDto) *NullableTenantQuotaFeatureDto {
	return &NullableTenantQuotaFeatureDto{value: val, isSet: true}
}

func (v NullableTenantQuotaFeatureDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantQuotaFeatureDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

