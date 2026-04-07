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
	"time"
)

// checks if the ContentDisposition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentDisposition{}

// ContentDisposition struct for ContentDisposition
type ContentDisposition struct {
	DispositionType NullableString `json:"dispositionType,omitempty"`
	Parameters []interface{} `json:"parameters,omitempty"`
	FileName NullableString `json:"fileName,omitempty"`
	CreationDate *time.Time `json:"creationDate,omitempty"`
	ModificationDate *time.Time `json:"modificationDate,omitempty"`
	Inline *bool `json:"inline,omitempty"`
	ReadDate *time.Time `json:"readDate,omitempty"`
	Size *int64 `json:"size,omitempty"`
}

// NewContentDisposition instantiates a new ContentDisposition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentDisposition() *ContentDisposition {
	this := ContentDisposition{}
	return &this
}

// NewContentDispositionWithDefaults instantiates a new ContentDisposition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentDispositionWithDefaults() *ContentDisposition {
	this := ContentDisposition{}
	return &this
}

// GetDispositionType returns the DispositionType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContentDisposition) GetDispositionType() string {
	if o == nil || IsNil(o.DispositionType.Get()) {
		var ret string
		return ret
	}
	return *o.DispositionType.Get()
}

// GetDispositionTypeOk returns a tuple with the DispositionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContentDisposition) GetDispositionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DispositionType.Get(), o.DispositionType.IsSet()
}

// HasDispositionType returns a boolean if a field has been set.
func (o *ContentDisposition) IsDispositionTypeSet() bool {
	if o != nil && o.DispositionType.IsSet() {
		return true
	}

	return false
}

// SetDispositionType gets a reference to the given NullableString and assigns it to the DispositionType field.
func (o *ContentDisposition) SetDispositionType(v string) {
	o.DispositionType.Set(&v)
}
// SetDispositionTypeNil sets the value for DispositionType to be an explicit nil
func (o *ContentDisposition) SetDispositionTypeNil() {
	o.DispositionType.Set(nil)
}

// UnsetDispositionType ensures that no value is present for DispositionType, not even an explicit nil
func (o *ContentDisposition) UnsetDispositionType() {
	o.DispositionType.Unset()
}

// GetParameters returns the Parameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContentDisposition) GetParameters() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContentDisposition) GetParametersOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *ContentDisposition) IsParametersSet() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []interface{} and assigns it to the Parameters field.
func (o *ContentDisposition) SetParameters(v []interface{}) {
	o.Parameters = v
}

// GetFileName returns the FileName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContentDisposition) GetFileName() string {
	if o == nil || IsNil(o.FileName.Get()) {
		var ret string
		return ret
	}
	return *o.FileName.Get()
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContentDisposition) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileName.Get(), o.FileName.IsSet()
}

// HasFileName returns a boolean if a field has been set.
func (o *ContentDisposition) IsFileNameSet() bool {
	if o != nil && o.FileName.IsSet() {
		return true
	}

	return false
}

// SetFileName gets a reference to the given NullableString and assigns it to the FileName field.
func (o *ContentDisposition) SetFileName(v string) {
	o.FileName.Set(&v)
}
// SetFileNameNil sets the value for FileName to be an explicit nil
func (o *ContentDisposition) SetFileNameNil() {
	o.FileName.Set(nil)
}

// UnsetFileName ensures that no value is present for FileName, not even an explicit nil
func (o *ContentDisposition) UnsetFileName() {
	o.FileName.Unset()
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *ContentDisposition) GetCreationDate() time.Time {
	if o == nil || IsNil(o.CreationDate) {
		var ret time.Time
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentDisposition) GetCreationDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationDate) {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *ContentDisposition) IsCreationDateSet() bool {
	if o != nil && !IsNil(o.CreationDate) {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given time.Time and assigns it to the CreationDate field.
func (o *ContentDisposition) SetCreationDate(v time.Time) {
	o.CreationDate = &v
}

// GetModificationDate returns the ModificationDate field value if set, zero value otherwise.
func (o *ContentDisposition) GetModificationDate() time.Time {
	if o == nil || IsNil(o.ModificationDate) {
		var ret time.Time
		return ret
	}
	return *o.ModificationDate
}

// GetModificationDateOk returns a tuple with the ModificationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentDisposition) GetModificationDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ModificationDate) {
		return nil, false
	}
	return o.ModificationDate, true
}

// HasModificationDate returns a boolean if a field has been set.
func (o *ContentDisposition) IsModificationDateSet() bool {
	if o != nil && !IsNil(o.ModificationDate) {
		return true
	}

	return false
}

// SetModificationDate gets a reference to the given time.Time and assigns it to the ModificationDate field.
func (o *ContentDisposition) SetModificationDate(v time.Time) {
	o.ModificationDate = &v
}

// GetInline returns the Inline field value if set, zero value otherwise.
func (o *ContentDisposition) GetInline() bool {
	if o == nil || IsNil(o.Inline) {
		var ret bool
		return ret
	}
	return *o.Inline
}

// GetInlineOk returns a tuple with the Inline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentDisposition) GetInlineOk() (*bool, bool) {
	if o == nil || IsNil(o.Inline) {
		return nil, false
	}
	return o.Inline, true
}

// HasInline returns a boolean if a field has been set.
func (o *ContentDisposition) IsInlineSet() bool {
	if o != nil && !IsNil(o.Inline) {
		return true
	}

	return false
}

// SetInline gets a reference to the given bool and assigns it to the Inline field.
func (o *ContentDisposition) SetInline(v bool) {
	o.Inline = &v
}

// GetReadDate returns the ReadDate field value if set, zero value otherwise.
func (o *ContentDisposition) GetReadDate() time.Time {
	if o == nil || IsNil(o.ReadDate) {
		var ret time.Time
		return ret
	}
	return *o.ReadDate
}

// GetReadDateOk returns a tuple with the ReadDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentDisposition) GetReadDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ReadDate) {
		return nil, false
	}
	return o.ReadDate, true
}

// HasReadDate returns a boolean if a field has been set.
func (o *ContentDisposition) IsReadDateSet() bool {
	if o != nil && !IsNil(o.ReadDate) {
		return true
	}

	return false
}

// SetReadDate gets a reference to the given time.Time and assigns it to the ReadDate field.
func (o *ContentDisposition) SetReadDate(v time.Time) {
	o.ReadDate = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *ContentDisposition) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentDisposition) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *ContentDisposition) IsSizeSet() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *ContentDisposition) SetSize(v int64) {
	o.Size = &v
}

func (o ContentDisposition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentDisposition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DispositionType.IsSet() {
		toSerialize["dispositionType"] = o.DispositionType.Get()
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.FileName.IsSet() {
		toSerialize["fileName"] = o.FileName.Get()
	}
	if !IsNil(o.CreationDate) {
		toSerialize["creationDate"] = o.CreationDate
	}
	if !IsNil(o.ModificationDate) {
		toSerialize["modificationDate"] = o.ModificationDate
	}
	if !IsNil(o.Inline) {
		toSerialize["inline"] = o.Inline
	}
	if !IsNil(o.ReadDate) {
		toSerialize["readDate"] = o.ReadDate
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	return toSerialize, nil
}

type NullableContentDisposition struct {
	value *ContentDisposition
	isSet bool
}

func (v NullableContentDisposition) Get() *ContentDisposition {
	return v.value
}

func (v *NullableContentDisposition) Set(val *ContentDisposition) {
	v.value = val
	v.isSet = true
}

func (v NullableContentDisposition) IsSet() bool {
	return v.isSet
}

func (v *NullableContentDisposition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentDisposition(val *ContentDisposition) *NullableContentDisposition {
	return &NullableContentDisposition{value: val, isSet: true}
}

func (v NullableContentDisposition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentDisposition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

