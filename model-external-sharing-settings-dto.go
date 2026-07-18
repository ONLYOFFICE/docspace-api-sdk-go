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

// checks if the ExternalSharingSettingsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalSharingSettingsDto{}

// ExternalSharingSettingsDto The Access Control external sharing settings.
type ExternalSharingSettingsDto struct {
	// Specifies whether external (public) link creation is allowed.
	ExternalShare *bool `json:"externalShare,omitempty"`
	// Specifies the default sharing link type: true = DocSpace users only, false = Anyone with the link.
	DefaultShareLinkInternal *bool `json:"defaultShareLinkInternal,omitempty"`
	// When external sharing is restricted, specifies whether the restriction applies to the My Documents section.
	ExternalShareApplyToDocuments *bool `json:"externalShareApplyToDocuments,omitempty"`
	// When external sharing is restricted, specifies whether the restriction applies to the Rooms section.
	ExternalShareApplyToRooms *bool `json:"externalShareApplyToRooms,omitempty"`
	// When external sharing is restricted, specifies whether existing public links are blocked immediately.
	BlockExistingLinksOnRestrict *bool `json:"blockExistingLinksOnRestrict,omitempty"`
}

// NewExternalSharingSettingsDto instantiates a new ExternalSharingSettingsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalSharingSettingsDto() *ExternalSharingSettingsDto {
	this := ExternalSharingSettingsDto{}
	return &this
}

// NewExternalSharingSettingsDtoWithDefaults instantiates a new ExternalSharingSettingsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalSharingSettingsDtoWithDefaults() *ExternalSharingSettingsDto {
	this := ExternalSharingSettingsDto{}
	return &this
}

// GetExternalShare returns the ExternalShare field value if set, zero value otherwise.
func (o *ExternalSharingSettingsDto) GetExternalShare() bool {
	if o == nil || IsNil(o.ExternalShare) {
		var ret bool
		return ret
	}
	return *o.ExternalShare
}

// GetExternalShareOk returns a tuple with the ExternalShare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalSharingSettingsDto) GetExternalShareOk() (*bool, bool) {
	if o == nil || IsNil(o.ExternalShare) {
		return nil, false
	}
	return o.ExternalShare, true
}

// HasExternalShare returns a boolean if a field has been set.
func (o *ExternalSharingSettingsDto) IsExternalShareSet() bool {
	if o != nil && !IsNil(o.ExternalShare) {
		return true
	}

	return false
}

// SetExternalShare gets a reference to the given bool and assigns it to the ExternalShare field.
func (o *ExternalSharingSettingsDto) SetExternalShare(v bool) {
	o.ExternalShare = &v
}

// GetDefaultShareLinkInternal returns the DefaultShareLinkInternal field value if set, zero value otherwise.
func (o *ExternalSharingSettingsDto) GetDefaultShareLinkInternal() bool {
	if o == nil || IsNil(o.DefaultShareLinkInternal) {
		var ret bool
		return ret
	}
	return *o.DefaultShareLinkInternal
}

// GetDefaultShareLinkInternalOk returns a tuple with the DefaultShareLinkInternal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalSharingSettingsDto) GetDefaultShareLinkInternalOk() (*bool, bool) {
	if o == nil || IsNil(o.DefaultShareLinkInternal) {
		return nil, false
	}
	return o.DefaultShareLinkInternal, true
}

// HasDefaultShareLinkInternal returns a boolean if a field has been set.
func (o *ExternalSharingSettingsDto) IsDefaultShareLinkInternalSet() bool {
	if o != nil && !IsNil(o.DefaultShareLinkInternal) {
		return true
	}

	return false
}

// SetDefaultShareLinkInternal gets a reference to the given bool and assigns it to the DefaultShareLinkInternal field.
func (o *ExternalSharingSettingsDto) SetDefaultShareLinkInternal(v bool) {
	o.DefaultShareLinkInternal = &v
}

// GetExternalShareApplyToDocuments returns the ExternalShareApplyToDocuments field value if set, zero value otherwise.
func (o *ExternalSharingSettingsDto) GetExternalShareApplyToDocuments() bool {
	if o == nil || IsNil(o.ExternalShareApplyToDocuments) {
		var ret bool
		return ret
	}
	return *o.ExternalShareApplyToDocuments
}

// GetExternalShareApplyToDocumentsOk returns a tuple with the ExternalShareApplyToDocuments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalSharingSettingsDto) GetExternalShareApplyToDocumentsOk() (*bool, bool) {
	if o == nil || IsNil(o.ExternalShareApplyToDocuments) {
		return nil, false
	}
	return o.ExternalShareApplyToDocuments, true
}

// HasExternalShareApplyToDocuments returns a boolean if a field has been set.
func (o *ExternalSharingSettingsDto) IsExternalShareApplyToDocumentsSet() bool {
	if o != nil && !IsNil(o.ExternalShareApplyToDocuments) {
		return true
	}

	return false
}

// SetExternalShareApplyToDocuments gets a reference to the given bool and assigns it to the ExternalShareApplyToDocuments field.
func (o *ExternalSharingSettingsDto) SetExternalShareApplyToDocuments(v bool) {
	o.ExternalShareApplyToDocuments = &v
}

// GetExternalShareApplyToRooms returns the ExternalShareApplyToRooms field value if set, zero value otherwise.
func (o *ExternalSharingSettingsDto) GetExternalShareApplyToRooms() bool {
	if o == nil || IsNil(o.ExternalShareApplyToRooms) {
		var ret bool
		return ret
	}
	return *o.ExternalShareApplyToRooms
}

// GetExternalShareApplyToRoomsOk returns a tuple with the ExternalShareApplyToRooms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalSharingSettingsDto) GetExternalShareApplyToRoomsOk() (*bool, bool) {
	if o == nil || IsNil(o.ExternalShareApplyToRooms) {
		return nil, false
	}
	return o.ExternalShareApplyToRooms, true
}

// HasExternalShareApplyToRooms returns a boolean if a field has been set.
func (o *ExternalSharingSettingsDto) IsExternalShareApplyToRoomsSet() bool {
	if o != nil && !IsNil(o.ExternalShareApplyToRooms) {
		return true
	}

	return false
}

// SetExternalShareApplyToRooms gets a reference to the given bool and assigns it to the ExternalShareApplyToRooms field.
func (o *ExternalSharingSettingsDto) SetExternalShareApplyToRooms(v bool) {
	o.ExternalShareApplyToRooms = &v
}

// GetBlockExistingLinksOnRestrict returns the BlockExistingLinksOnRestrict field value if set, zero value otherwise.
func (o *ExternalSharingSettingsDto) GetBlockExistingLinksOnRestrict() bool {
	if o == nil || IsNil(o.BlockExistingLinksOnRestrict) {
		var ret bool
		return ret
	}
	return *o.BlockExistingLinksOnRestrict
}

// GetBlockExistingLinksOnRestrictOk returns a tuple with the BlockExistingLinksOnRestrict field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalSharingSettingsDto) GetBlockExistingLinksOnRestrictOk() (*bool, bool) {
	if o == nil || IsNil(o.BlockExistingLinksOnRestrict) {
		return nil, false
	}
	return o.BlockExistingLinksOnRestrict, true
}

// HasBlockExistingLinksOnRestrict returns a boolean if a field has been set.
func (o *ExternalSharingSettingsDto) IsBlockExistingLinksOnRestrictSet() bool {
	if o != nil && !IsNil(o.BlockExistingLinksOnRestrict) {
		return true
	}

	return false
}

// SetBlockExistingLinksOnRestrict gets a reference to the given bool and assigns it to the BlockExistingLinksOnRestrict field.
func (o *ExternalSharingSettingsDto) SetBlockExistingLinksOnRestrict(v bool) {
	o.BlockExistingLinksOnRestrict = &v
}

func (o ExternalSharingSettingsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalSharingSettingsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExternalShare) {
		toSerialize["externalShare"] = o.ExternalShare
	}
	if !IsNil(o.DefaultShareLinkInternal) {
		toSerialize["defaultShareLinkInternal"] = o.DefaultShareLinkInternal
	}
	if !IsNil(o.ExternalShareApplyToDocuments) {
		toSerialize["externalShareApplyToDocuments"] = o.ExternalShareApplyToDocuments
	}
	if !IsNil(o.ExternalShareApplyToRooms) {
		toSerialize["externalShareApplyToRooms"] = o.ExternalShareApplyToRooms
	}
	if !IsNil(o.BlockExistingLinksOnRestrict) {
		toSerialize["blockExistingLinksOnRestrict"] = o.BlockExistingLinksOnRestrict
	}
	return toSerialize, nil
}

type NullableExternalSharingSettingsDto struct {
	value *ExternalSharingSettingsDto
	isSet bool
}

func (v NullableExternalSharingSettingsDto) Get() *ExternalSharingSettingsDto {
	return v.value
}

func (v *NullableExternalSharingSettingsDto) Set(val *ExternalSharingSettingsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalSharingSettingsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalSharingSettingsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalSharingSettingsDto(val *ExternalSharingSettingsDto) *NullableExternalSharingSettingsDto {
	return &NullableExternalSharingSettingsDto{value: val, isSet: true}
}

func (v NullableExternalSharingSettingsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalSharingSettingsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

