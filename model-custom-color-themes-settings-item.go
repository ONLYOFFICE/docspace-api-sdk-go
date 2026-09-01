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

// checks if the CustomColorThemesSettingsItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomColorThemesSettingsItem{}

// CustomColorThemesSettingsItem The custom color theme settings.
type CustomColorThemesSettingsItem struct {
	// The custom color theme ID.
	Id *int32 `json:"id,omitempty"`
	// The custom color theme name.
	Name NullableString `json:"name,omitempty"`
	// The custom color theme main colors.
	Main *CustomColorThemesSettingsColorItem `json:"main,omitempty"`
	// The custom color theme text colors.
	Text *CustomColorThemesSettingsColorItem `json:"text,omitempty"`
}

// NewCustomColorThemesSettingsItem instantiates a new CustomColorThemesSettingsItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomColorThemesSettingsItem() *CustomColorThemesSettingsItem {
	this := CustomColorThemesSettingsItem{}
	return &this
}

// NewCustomColorThemesSettingsItemWithDefaults instantiates a new CustomColorThemesSettingsItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomColorThemesSettingsItemWithDefaults() *CustomColorThemesSettingsItem {
	this := CustomColorThemesSettingsItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomColorThemesSettingsItem) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomColorThemesSettingsItem) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomColorThemesSettingsItem) IsIdSet() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CustomColorThemesSettingsItem) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomColorThemesSettingsItem) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomColorThemesSettingsItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CustomColorThemesSettingsItem) IsNameSet() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CustomColorThemesSettingsItem) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *CustomColorThemesSettingsItem) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CustomColorThemesSettingsItem) UnsetName() {
	o.Name.Unset()
}

// GetMain returns the Main field value if set, zero value otherwise.
func (o *CustomColorThemesSettingsItem) GetMain() CustomColorThemesSettingsColorItem {
	if o == nil || IsNil(o.Main) {
		var ret CustomColorThemesSettingsColorItem
		return ret
	}
	return *o.Main
}

// GetMainOk returns a tuple with the Main field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomColorThemesSettingsItem) GetMainOk() (*CustomColorThemesSettingsColorItem, bool) {
	if o == nil || IsNil(o.Main) {
		return nil, false
	}
	return o.Main, true
}

// HasMain returns a boolean if a field has been set.
func (o *CustomColorThemesSettingsItem) IsMainSet() bool {
	if o != nil && !IsNil(o.Main) {
		return true
	}

	return false
}

// SetMain gets a reference to the given CustomColorThemesSettingsColorItem and assigns it to the Main field.
func (o *CustomColorThemesSettingsItem) SetMain(v CustomColorThemesSettingsColorItem) {
	o.Main = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *CustomColorThemesSettingsItem) GetText() CustomColorThemesSettingsColorItem {
	if o == nil || IsNil(o.Text) {
		var ret CustomColorThemesSettingsColorItem
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomColorThemesSettingsItem) GetTextOk() (*CustomColorThemesSettingsColorItem, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *CustomColorThemesSettingsItem) IsTextSet() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given CustomColorThemesSettingsColorItem and assigns it to the Text field.
func (o *CustomColorThemesSettingsItem) SetText(v CustomColorThemesSettingsColorItem) {
	o.Text = &v
}

func (o CustomColorThemesSettingsItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomColorThemesSettingsItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.Main) {
		toSerialize["main"] = o.Main
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	return toSerialize, nil
}

type NullableCustomColorThemesSettingsItem struct {
	value *CustomColorThemesSettingsItem
	isSet bool
}

func (v NullableCustomColorThemesSettingsItem) Get() *CustomColorThemesSettingsItem {
	return v.value
}

func (v *NullableCustomColorThemesSettingsItem) Set(val *CustomColorThemesSettingsItem) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomColorThemesSettingsItem) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomColorThemesSettingsItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomColorThemesSettingsItem(val *CustomColorThemesSettingsItem) *NullableCustomColorThemesSettingsItem {
	return &NullableCustomColorThemesSettingsItem{value: val, isSet: true}
}

func (v NullableCustomColorThemesSettingsItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomColorThemesSettingsItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

