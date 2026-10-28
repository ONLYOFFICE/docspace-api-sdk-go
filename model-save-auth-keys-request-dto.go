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

// checks if the SaveAuthKeysRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SaveAuthKeysRequestDto{}

// SaveAuthKeysRequestDto The keys to store for one third-party authorization or storage provider.
type SaveAuthKeysRequestDto struct {
	// The internal key of the provider, such as `google` or `box`. Take it from the `name` of  `GET api/2.0/settings/authservice`.
	Name NullableString `json:"name,omitempty"`
	// Accepted for compatibility with earlier clients and not read: the server keeps its own value.
	Title NullableString `json:"title,omitempty"`
	// Accepted for compatibility with earlier clients and not read: the server keeps its own value.
	Description NullableString `json:"description,omitempty"`
	// Accepted for compatibility with earlier clients and not read: the server keeps its own value.
	Instruction NullableString `json:"instruction,omitempty"`
	// Accepted for compatibility with earlier clients and not read: the server keeps its own value.
	CanSet *bool `json:"canSet,omitempty"`
	// Accepted for compatibility with earlier clients and not read: the server keeps its own value.
	Paid *bool `json:"paid,omitempty"`
	// The keys of the provider with their new values, by the key names `GET api/2.0/settings/authservice` lists in  `props`.
	Props []AuthKeyRequest `json:"props,omitempty"`
}

// NewSaveAuthKeysRequestDto instantiates a new SaveAuthKeysRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSaveAuthKeysRequestDto() *SaveAuthKeysRequestDto {
	this := SaveAuthKeysRequestDto{}
	return &this
}

// NewSaveAuthKeysRequestDtoWithDefaults instantiates a new SaveAuthKeysRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSaveAuthKeysRequestDtoWithDefaults() *SaveAuthKeysRequestDto {
	this := SaveAuthKeysRequestDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SaveAuthKeysRequestDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SaveAuthKeysRequestDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *SaveAuthKeysRequestDto) IsNameSet() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *SaveAuthKeysRequestDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *SaveAuthKeysRequestDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *SaveAuthKeysRequestDto) UnsetName() {
	o.Name.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SaveAuthKeysRequestDto) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SaveAuthKeysRequestDto) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *SaveAuthKeysRequestDto) IsTitleSet() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *SaveAuthKeysRequestDto) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *SaveAuthKeysRequestDto) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *SaveAuthKeysRequestDto) UnsetTitle() {
	o.Title.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SaveAuthKeysRequestDto) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SaveAuthKeysRequestDto) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *SaveAuthKeysRequestDto) IsDescriptionSet() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *SaveAuthKeysRequestDto) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *SaveAuthKeysRequestDto) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *SaveAuthKeysRequestDto) UnsetDescription() {
	o.Description.Unset()
}

// GetInstruction returns the Instruction field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SaveAuthKeysRequestDto) GetInstruction() string {
	if o == nil || IsNil(o.Instruction.Get()) {
		var ret string
		return ret
	}
	return *o.Instruction.Get()
}

// GetInstructionOk returns a tuple with the Instruction field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SaveAuthKeysRequestDto) GetInstructionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Instruction.Get(), o.Instruction.IsSet()
}

// HasInstruction returns a boolean if a field has been set.
func (o *SaveAuthKeysRequestDto) IsInstructionSet() bool {
	if o != nil && o.Instruction.IsSet() {
		return true
	}

	return false
}

// SetInstruction gets a reference to the given NullableString and assigns it to the Instruction field.
func (o *SaveAuthKeysRequestDto) SetInstruction(v string) {
	o.Instruction.Set(&v)
}
// SetInstructionNil sets the value for Instruction to be an explicit nil
func (o *SaveAuthKeysRequestDto) SetInstructionNil() {
	o.Instruction.Set(nil)
}

// UnsetInstruction ensures that no value is present for Instruction, not even an explicit nil
func (o *SaveAuthKeysRequestDto) UnsetInstruction() {
	o.Instruction.Unset()
}

// GetCanSet returns the CanSet field value if set, zero value otherwise.
func (o *SaveAuthKeysRequestDto) GetCanSet() bool {
	if o == nil || IsNil(o.CanSet) {
		var ret bool
		return ret
	}
	return *o.CanSet
}

// GetCanSetOk returns a tuple with the CanSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaveAuthKeysRequestDto) GetCanSetOk() (*bool, bool) {
	if o == nil || IsNil(o.CanSet) {
		return nil, false
	}
	return o.CanSet, true
}

// HasCanSet returns a boolean if a field has been set.
func (o *SaveAuthKeysRequestDto) IsCanSetSet() bool {
	if o != nil && !IsNil(o.CanSet) {
		return true
	}

	return false
}

// SetCanSet gets a reference to the given bool and assigns it to the CanSet field.
func (o *SaveAuthKeysRequestDto) SetCanSet(v bool) {
	o.CanSet = &v
}

// GetPaid returns the Paid field value if set, zero value otherwise.
func (o *SaveAuthKeysRequestDto) GetPaid() bool {
	if o == nil || IsNil(o.Paid) {
		var ret bool
		return ret
	}
	return *o.Paid
}

// GetPaidOk returns a tuple with the Paid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaveAuthKeysRequestDto) GetPaidOk() (*bool, bool) {
	if o == nil || IsNil(o.Paid) {
		return nil, false
	}
	return o.Paid, true
}

// HasPaid returns a boolean if a field has been set.
func (o *SaveAuthKeysRequestDto) IsPaidSet() bool {
	if o != nil && !IsNil(o.Paid) {
		return true
	}

	return false
}

// SetPaid gets a reference to the given bool and assigns it to the Paid field.
func (o *SaveAuthKeysRequestDto) SetPaid(v bool) {
	o.Paid = &v
}

// GetProps returns the Props field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SaveAuthKeysRequestDto) GetProps() []AuthKeyRequest {
	if o == nil {
		var ret []AuthKeyRequest
		return ret
	}
	return o.Props
}

// GetPropsOk returns a tuple with the Props field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SaveAuthKeysRequestDto) GetPropsOk() ([]AuthKeyRequest, bool) {
	if o == nil || IsNil(o.Props) {
		return nil, false
	}
	return o.Props, true
}

// HasProps returns a boolean if a field has been set.
func (o *SaveAuthKeysRequestDto) IsPropsSet() bool {
	if o != nil && !IsNil(o.Props) {
		return true
	}

	return false
}

// SetProps gets a reference to the given []AuthKeyRequest and assigns it to the Props field.
func (o *SaveAuthKeysRequestDto) SetProps(v []AuthKeyRequest) {
	o.Props = v
}

func (o SaveAuthKeysRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SaveAuthKeysRequestDto) ToMap() (map[string]interface{}, error) {
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

type NullableSaveAuthKeysRequestDto struct {
	value *SaveAuthKeysRequestDto
	isSet bool
}

func (v NullableSaveAuthKeysRequestDto) Get() *SaveAuthKeysRequestDto {
	return v.value
}

func (v *NullableSaveAuthKeysRequestDto) Set(val *SaveAuthKeysRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSaveAuthKeysRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSaveAuthKeysRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSaveAuthKeysRequestDto(val *SaveAuthKeysRequestDto) *NullableSaveAuthKeysRequestDto {
	return &NullableSaveAuthKeysRequestDto{value: val, isSet: true}
}

func (v NullableSaveAuthKeysRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSaveAuthKeysRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

