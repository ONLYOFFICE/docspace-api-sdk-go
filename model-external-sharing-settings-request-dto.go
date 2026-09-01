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

// checks if the ExternalSharingSettingsRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalSharingSettingsRequestDto{}

// ExternalSharingSettingsRequestDto The Access Control external sharing settings request parameters.
type ExternalSharingSettingsRequestDto struct {
	// Specifies whether external (public) link creation is allowed.
	ExternalShare *bool `json:"externalShare,omitempty"`
	// Specifies the default sharing link type: true = DocSpace users only, false = Anyone with the link.  Relevant only when ExternalShare is true.
	DefaultShareLinkInternal *bool `json:"defaultShareLinkInternal,omitempty"`
	// When external sharing is restricted, specifies whether to apply the restriction to the My Documents section.  Relevant only when ExternalShare is false.
	ExternalShareApplyToDocuments *bool `json:"externalShareApplyToDocuments,omitempty"`
	// When external sharing is restricted, specifies whether to apply the restriction to the Rooms section.  Relevant only when ExternalShare is false.
	ExternalShareApplyToRooms *bool `json:"externalShareApplyToRooms,omitempty"`
	// When external sharing is restricted, specifies whether to block existing public links immediately.  Relevant only when ExternalShare is false.
	BlockExistingLinksOnRestrict *bool `json:"blockExistingLinksOnRestrict,omitempty"`
}

// NewExternalSharingSettingsRequestDto instantiates a new ExternalSharingSettingsRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalSharingSettingsRequestDto() *ExternalSharingSettingsRequestDto {
	this := ExternalSharingSettingsRequestDto{}
	return &this
}

// NewExternalSharingSettingsRequestDtoWithDefaults instantiates a new ExternalSharingSettingsRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalSharingSettingsRequestDtoWithDefaults() *ExternalSharingSettingsRequestDto {
	this := ExternalSharingSettingsRequestDto{}
	return &this
}

// GetExternalShare returns the ExternalShare field value if set, zero value otherwise.
func (o *ExternalSharingSettingsRequestDto) GetExternalShare() bool {
	if o == nil || IsNil(o.ExternalShare) {
		var ret bool
		return ret
	}
	return *o.ExternalShare
}

// GetExternalShareOk returns a tuple with the ExternalShare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalSharingSettingsRequestDto) GetExternalShareOk() (*bool, bool) {
	if o == nil || IsNil(o.ExternalShare) {
		return nil, false
	}
	return o.ExternalShare, true
}

// HasExternalShare returns a boolean if a field has been set.
func (o *ExternalSharingSettingsRequestDto) IsExternalShareSet() bool {
	if o != nil && !IsNil(o.ExternalShare) {
		return true
	}

	return false
}

// SetExternalShare gets a reference to the given bool and assigns it to the ExternalShare field.
func (o *ExternalSharingSettingsRequestDto) SetExternalShare(v bool) {
	o.ExternalShare = &v
}

// GetDefaultShareLinkInternal returns the DefaultShareLinkInternal field value if set, zero value otherwise.
func (o *ExternalSharingSettingsRequestDto) GetDefaultShareLinkInternal() bool {
	if o == nil || IsNil(o.DefaultShareLinkInternal) {
		var ret bool
		return ret
	}
	return *o.DefaultShareLinkInternal
}

// GetDefaultShareLinkInternalOk returns a tuple with the DefaultShareLinkInternal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalSharingSettingsRequestDto) GetDefaultShareLinkInternalOk() (*bool, bool) {
	if o == nil || IsNil(o.DefaultShareLinkInternal) {
		return nil, false
	}
	return o.DefaultShareLinkInternal, true
}

// HasDefaultShareLinkInternal returns a boolean if a field has been set.
func (o *ExternalSharingSettingsRequestDto) IsDefaultShareLinkInternalSet() bool {
	if o != nil && !IsNil(o.DefaultShareLinkInternal) {
		return true
	}

	return false
}

// SetDefaultShareLinkInternal gets a reference to the given bool and assigns it to the DefaultShareLinkInternal field.
func (o *ExternalSharingSettingsRequestDto) SetDefaultShareLinkInternal(v bool) {
	o.DefaultShareLinkInternal = &v
}

// GetExternalShareApplyToDocuments returns the ExternalShareApplyToDocuments field value if set, zero value otherwise.
func (o *ExternalSharingSettingsRequestDto) GetExternalShareApplyToDocuments() bool {
	if o == nil || IsNil(o.ExternalShareApplyToDocuments) {
		var ret bool
		return ret
	}
	return *o.ExternalShareApplyToDocuments
}

// GetExternalShareApplyToDocumentsOk returns a tuple with the ExternalShareApplyToDocuments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalSharingSettingsRequestDto) GetExternalShareApplyToDocumentsOk() (*bool, bool) {
	if o == nil || IsNil(o.ExternalShareApplyToDocuments) {
		return nil, false
	}
	return o.ExternalShareApplyToDocuments, true
}

// HasExternalShareApplyToDocuments returns a boolean if a field has been set.
func (o *ExternalSharingSettingsRequestDto) IsExternalShareApplyToDocumentsSet() bool {
	if o != nil && !IsNil(o.ExternalShareApplyToDocuments) {
		return true
	}

	return false
}

// SetExternalShareApplyToDocuments gets a reference to the given bool and assigns it to the ExternalShareApplyToDocuments field.
func (o *ExternalSharingSettingsRequestDto) SetExternalShareApplyToDocuments(v bool) {
	o.ExternalShareApplyToDocuments = &v
}

// GetExternalShareApplyToRooms returns the ExternalShareApplyToRooms field value if set, zero value otherwise.
func (o *ExternalSharingSettingsRequestDto) GetExternalShareApplyToRooms() bool {
	if o == nil || IsNil(o.ExternalShareApplyToRooms) {
		var ret bool
		return ret
	}
	return *o.ExternalShareApplyToRooms
}

// GetExternalShareApplyToRoomsOk returns a tuple with the ExternalShareApplyToRooms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalSharingSettingsRequestDto) GetExternalShareApplyToRoomsOk() (*bool, bool) {
	if o == nil || IsNil(o.ExternalShareApplyToRooms) {
		return nil, false
	}
	return o.ExternalShareApplyToRooms, true
}

// HasExternalShareApplyToRooms returns a boolean if a field has been set.
func (o *ExternalSharingSettingsRequestDto) IsExternalShareApplyToRoomsSet() bool {
	if o != nil && !IsNil(o.ExternalShareApplyToRooms) {
		return true
	}

	return false
}

// SetExternalShareApplyToRooms gets a reference to the given bool and assigns it to the ExternalShareApplyToRooms field.
func (o *ExternalSharingSettingsRequestDto) SetExternalShareApplyToRooms(v bool) {
	o.ExternalShareApplyToRooms = &v
}

// GetBlockExistingLinksOnRestrict returns the BlockExistingLinksOnRestrict field value if set, zero value otherwise.
func (o *ExternalSharingSettingsRequestDto) GetBlockExistingLinksOnRestrict() bool {
	if o == nil || IsNil(o.BlockExistingLinksOnRestrict) {
		var ret bool
		return ret
	}
	return *o.BlockExistingLinksOnRestrict
}

// GetBlockExistingLinksOnRestrictOk returns a tuple with the BlockExistingLinksOnRestrict field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalSharingSettingsRequestDto) GetBlockExistingLinksOnRestrictOk() (*bool, bool) {
	if o == nil || IsNil(o.BlockExistingLinksOnRestrict) {
		return nil, false
	}
	return o.BlockExistingLinksOnRestrict, true
}

// HasBlockExistingLinksOnRestrict returns a boolean if a field has been set.
func (o *ExternalSharingSettingsRequestDto) IsBlockExistingLinksOnRestrictSet() bool {
	if o != nil && !IsNil(o.BlockExistingLinksOnRestrict) {
		return true
	}

	return false
}

// SetBlockExistingLinksOnRestrict gets a reference to the given bool and assigns it to the BlockExistingLinksOnRestrict field.
func (o *ExternalSharingSettingsRequestDto) SetBlockExistingLinksOnRestrict(v bool) {
	o.BlockExistingLinksOnRestrict = &v
}

func (o ExternalSharingSettingsRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalSharingSettingsRequestDto) ToMap() (map[string]interface{}, error) {
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

type NullableExternalSharingSettingsRequestDto struct {
	value *ExternalSharingSettingsRequestDto
	isSet bool
}

func (v NullableExternalSharingSettingsRequestDto) Get() *ExternalSharingSettingsRequestDto {
	return v.value
}

func (v *NullableExternalSharingSettingsRequestDto) Set(val *ExternalSharingSettingsRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalSharingSettingsRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalSharingSettingsRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalSharingSettingsRequestDto(val *ExternalSharingSettingsRequestDto) *NullableExternalSharingSettingsRequestDto {
	return &NullableExternalSharingSettingsRequestDto{value: val, isSet: true}
}

func (v NullableExternalSharingSettingsRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalSharingSettingsRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

