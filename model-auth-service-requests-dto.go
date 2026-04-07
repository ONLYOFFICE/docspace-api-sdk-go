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

// checks if the AuthServiceRequestsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthServiceRequestsDto{}

// AuthServiceRequestsDto The request parameters for handling the authorization service.
type AuthServiceRequestsDto struct {
	// The name of the authorization service.
	Name NullableString `json:"name,omitempty"`
	// The user-friendly display title of the authorization service.
	Title NullableString `json:"title,omitempty"`
	// The brief description of the authorization service.
	Description NullableString `json:"description,omitempty"`
	// The detailed instructions for configuring or using the authorization service.
	Instruction NullableString `json:"instruction,omitempty"`
	// Specifies whether the authorization service can be configured by the user.
	CanSet *bool `json:"canSet,omitempty"`
	// Specifies whether the authorization service is paid or not.
	Paid *bool `json:"paid,omitempty"`
	// The collection of authorization keys associated with the authorization service.
	Props []AuthKey `json:"props,omitempty"`
}

// NewAuthServiceRequestsDto instantiates a new AuthServiceRequestsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthServiceRequestsDto() *AuthServiceRequestsDto {
	this := AuthServiceRequestsDto{}
	return &this
}

// NewAuthServiceRequestsDtoWithDefaults instantiates a new AuthServiceRequestsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthServiceRequestsDtoWithDefaults() *AuthServiceRequestsDto {
	this := AuthServiceRequestsDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthServiceRequestsDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthServiceRequestsDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *AuthServiceRequestsDto) IsNameSet() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *AuthServiceRequestsDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *AuthServiceRequestsDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *AuthServiceRequestsDto) UnsetName() {
	o.Name.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthServiceRequestsDto) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthServiceRequestsDto) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *AuthServiceRequestsDto) IsTitleSet() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *AuthServiceRequestsDto) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *AuthServiceRequestsDto) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *AuthServiceRequestsDto) UnsetTitle() {
	o.Title.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthServiceRequestsDto) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthServiceRequestsDto) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *AuthServiceRequestsDto) IsDescriptionSet() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *AuthServiceRequestsDto) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *AuthServiceRequestsDto) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *AuthServiceRequestsDto) UnsetDescription() {
	o.Description.Unset()
}

// GetInstruction returns the Instruction field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthServiceRequestsDto) GetInstruction() string {
	if o == nil || IsNil(o.Instruction.Get()) {
		var ret string
		return ret
	}
	return *o.Instruction.Get()
}

// GetInstructionOk returns a tuple with the Instruction field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthServiceRequestsDto) GetInstructionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Instruction.Get(), o.Instruction.IsSet()
}

// HasInstruction returns a boolean if a field has been set.
func (o *AuthServiceRequestsDto) IsInstructionSet() bool {
	if o != nil && o.Instruction.IsSet() {
		return true
	}

	return false
}

// SetInstruction gets a reference to the given NullableString and assigns it to the Instruction field.
func (o *AuthServiceRequestsDto) SetInstruction(v string) {
	o.Instruction.Set(&v)
}
// SetInstructionNil sets the value for Instruction to be an explicit nil
func (o *AuthServiceRequestsDto) SetInstructionNil() {
	o.Instruction.Set(nil)
}

// UnsetInstruction ensures that no value is present for Instruction, not even an explicit nil
func (o *AuthServiceRequestsDto) UnsetInstruction() {
	o.Instruction.Unset()
}

// GetCanSet returns the CanSet field value if set, zero value otherwise.
func (o *AuthServiceRequestsDto) GetCanSet() bool {
	if o == nil || IsNil(o.CanSet) {
		var ret bool
		return ret
	}
	return *o.CanSet
}

// GetCanSetOk returns a tuple with the CanSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthServiceRequestsDto) GetCanSetOk() (*bool, bool) {
	if o == nil || IsNil(o.CanSet) {
		return nil, false
	}
	return o.CanSet, true
}

// HasCanSet returns a boolean if a field has been set.
func (o *AuthServiceRequestsDto) IsCanSetSet() bool {
	if o != nil && !IsNil(o.CanSet) {
		return true
	}

	return false
}

// SetCanSet gets a reference to the given bool and assigns it to the CanSet field.
func (o *AuthServiceRequestsDto) SetCanSet(v bool) {
	o.CanSet = &v
}

// GetPaid returns the Paid field value if set, zero value otherwise.
func (o *AuthServiceRequestsDto) GetPaid() bool {
	if o == nil || IsNil(o.Paid) {
		var ret bool
		return ret
	}
	return *o.Paid
}

// GetPaidOk returns a tuple with the Paid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthServiceRequestsDto) GetPaidOk() (*bool, bool) {
	if o == nil || IsNil(o.Paid) {
		return nil, false
	}
	return o.Paid, true
}

// HasPaid returns a boolean if a field has been set.
func (o *AuthServiceRequestsDto) IsPaidSet() bool {
	if o != nil && !IsNil(o.Paid) {
		return true
	}

	return false
}

// SetPaid gets a reference to the given bool and assigns it to the Paid field.
func (o *AuthServiceRequestsDto) SetPaid(v bool) {
	o.Paid = &v
}

// GetProps returns the Props field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthServiceRequestsDto) GetProps() []AuthKey {
	if o == nil {
		var ret []AuthKey
		return ret
	}
	return o.Props
}

// GetPropsOk returns a tuple with the Props field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthServiceRequestsDto) GetPropsOk() ([]AuthKey, bool) {
	if o == nil || IsNil(o.Props) {
		return nil, false
	}
	return o.Props, true
}

// HasProps returns a boolean if a field has been set.
func (o *AuthServiceRequestsDto) IsPropsSet() bool {
	if o != nil && !IsNil(o.Props) {
		return true
	}

	return false
}

// SetProps gets a reference to the given []AuthKey and assigns it to the Props field.
func (o *AuthServiceRequestsDto) SetProps(v []AuthKey) {
	o.Props = v
}

func (o AuthServiceRequestsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthServiceRequestsDto) ToMap() (map[string]interface{}, error) {
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

type NullableAuthServiceRequestsDto struct {
	value *AuthServiceRequestsDto
	isSet bool
}

func (v NullableAuthServiceRequestsDto) Get() *AuthServiceRequestsDto {
	return v.value
}

func (v *NullableAuthServiceRequestsDto) Set(val *AuthServiceRequestsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthServiceRequestsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthServiceRequestsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthServiceRequestsDto(val *AuthServiceRequestsDto) *NullableAuthServiceRequestsDto {
	return &NullableAuthServiceRequestsDto{value: val, isSet: true}
}

func (v NullableAuthServiceRequestsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthServiceRequestsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

