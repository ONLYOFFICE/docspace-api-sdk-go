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

// checks if the AuthServiceDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthServiceDto{}

// AuthServiceDto One third-party authorization or storage provider and the keys the portal connects to it with.
type AuthServiceDto struct {
	// The internal key of the provider, such as `google` or `box`. It is the `name` that  `POST api/2.0/settings/authservice` takes to select the provider.
	Name NullableString `json:"name,omitempty"`
	// The provider name as it is shown in the interface.
	Title NullableString `json:"title,omitempty"`
	// A sentence about what connecting the provider gives the portal, shown next to it in the interface.
	Description NullableString `json:"description,omitempty"`
	// The steps an administrator has to take on the provider side to obtain the keys, shown in the interface.
	Instruction NullableString `json:"instruction,omitempty"`
	// Whether this provider accepts keys through the API at all. A provider whose keys are fixed by the  installation reports `false`, and saving keys for it is refused.
	CanSet *bool `json:"canSet,omitempty"`
	// Whether the provider is a paid option. A paid one can only be connected while the portal plan includes  third-party storage or the installation is licensed as self-hosted.
	Paid *bool `json:"paid,omitempty"`
	// The keys the provider defines, with the values last saved and how the settings form shows each of them.  It is `null` for a provider that forbids changes (`canSet` is `false`): its keys are not read at all.
	Props []AuthKeyDto `json:"props,omitempty"`
}

// NewAuthServiceDto instantiates a new AuthServiceDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthServiceDto() *AuthServiceDto {
	this := AuthServiceDto{}
	return &this
}

// NewAuthServiceDtoWithDefaults instantiates a new AuthServiceDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthServiceDtoWithDefaults() *AuthServiceDto {
	this := AuthServiceDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthServiceDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthServiceDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *AuthServiceDto) IsNameSet() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *AuthServiceDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *AuthServiceDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *AuthServiceDto) UnsetName() {
	o.Name.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthServiceDto) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthServiceDto) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *AuthServiceDto) IsTitleSet() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *AuthServiceDto) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *AuthServiceDto) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *AuthServiceDto) UnsetTitle() {
	o.Title.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthServiceDto) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthServiceDto) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *AuthServiceDto) IsDescriptionSet() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *AuthServiceDto) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *AuthServiceDto) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *AuthServiceDto) UnsetDescription() {
	o.Description.Unset()
}

// GetInstruction returns the Instruction field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthServiceDto) GetInstruction() string {
	if o == nil || IsNil(o.Instruction.Get()) {
		var ret string
		return ret
	}
	return *o.Instruction.Get()
}

// GetInstructionOk returns a tuple with the Instruction field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthServiceDto) GetInstructionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Instruction.Get(), o.Instruction.IsSet()
}

// HasInstruction returns a boolean if a field has been set.
func (o *AuthServiceDto) IsInstructionSet() bool {
	if o != nil && o.Instruction.IsSet() {
		return true
	}

	return false
}

// SetInstruction gets a reference to the given NullableString and assigns it to the Instruction field.
func (o *AuthServiceDto) SetInstruction(v string) {
	o.Instruction.Set(&v)
}
// SetInstructionNil sets the value for Instruction to be an explicit nil
func (o *AuthServiceDto) SetInstructionNil() {
	o.Instruction.Set(nil)
}

// UnsetInstruction ensures that no value is present for Instruction, not even an explicit nil
func (o *AuthServiceDto) UnsetInstruction() {
	o.Instruction.Unset()
}

// GetCanSet returns the CanSet field value if set, zero value otherwise.
func (o *AuthServiceDto) GetCanSet() bool {
	if o == nil || IsNil(o.CanSet) {
		var ret bool
		return ret
	}
	return *o.CanSet
}

// GetCanSetOk returns a tuple with the CanSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthServiceDto) GetCanSetOk() (*bool, bool) {
	if o == nil || IsNil(o.CanSet) {
		return nil, false
	}
	return o.CanSet, true
}

// HasCanSet returns a boolean if a field has been set.
func (o *AuthServiceDto) IsCanSetSet() bool {
	if o != nil && !IsNil(o.CanSet) {
		return true
	}

	return false
}

// SetCanSet gets a reference to the given bool and assigns it to the CanSet field.
func (o *AuthServiceDto) SetCanSet(v bool) {
	o.CanSet = &v
}

// GetPaid returns the Paid field value if set, zero value otherwise.
func (o *AuthServiceDto) GetPaid() bool {
	if o == nil || IsNil(o.Paid) {
		var ret bool
		return ret
	}
	return *o.Paid
}

// GetPaidOk returns a tuple with the Paid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthServiceDto) GetPaidOk() (*bool, bool) {
	if o == nil || IsNil(o.Paid) {
		return nil, false
	}
	return o.Paid, true
}

// HasPaid returns a boolean if a field has been set.
func (o *AuthServiceDto) IsPaidSet() bool {
	if o != nil && !IsNil(o.Paid) {
		return true
	}

	return false
}

// SetPaid gets a reference to the given bool and assigns it to the Paid field.
func (o *AuthServiceDto) SetPaid(v bool) {
	o.Paid = &v
}

// GetProps returns the Props field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthServiceDto) GetProps() []AuthKeyDto {
	if o == nil {
		var ret []AuthKeyDto
		return ret
	}
	return o.Props
}

// GetPropsOk returns a tuple with the Props field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthServiceDto) GetPropsOk() ([]AuthKeyDto, bool) {
	if o == nil || IsNil(o.Props) {
		return nil, false
	}
	return o.Props, true
}

// HasProps returns a boolean if a field has been set.
func (o *AuthServiceDto) IsPropsSet() bool {
	if o != nil && !IsNil(o.Props) {
		return true
	}

	return false
}

// SetProps gets a reference to the given []AuthKeyDto and assigns it to the Props field.
func (o *AuthServiceDto) SetProps(v []AuthKeyDto) {
	o.Props = v
}

func (o AuthServiceDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthServiceDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Instruction.IsSet() {
		toSerialize["instruction"] = o.Instruction.Get()
	}
	if !IsNil(o.CanSet) {
		toSerialize["canSet"] = o.CanSet
	}
	if !IsNil(o.Paid) {
		toSerialize["paid"] = o.Paid
	}
	if o.Props != nil {
		toSerialize["props"] = o.Props
	}
	return toSerialize, nil
}

type NullableAuthServiceDto struct {
	value *AuthServiceDto
	isSet bool
}

func (v NullableAuthServiceDto) Get() *AuthServiceDto {
	return v.value
}

func (v *NullableAuthServiceDto) Set(val *AuthServiceDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthServiceDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthServiceDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthServiceDto(val *AuthServiceDto) *NullableAuthServiceDto {
	return &NullableAuthServiceDto{value: val, isSet: true}
}

func (v NullableAuthServiceDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthServiceDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

