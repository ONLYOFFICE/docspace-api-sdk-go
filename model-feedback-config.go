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

// checks if the FeedbackConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeedbackConfig{}

// FeedbackConfig The settings for the Feedback & Support menu button.
type FeedbackConfig struct {
	// The absolute URL to the website address which will be opened when clicking the Feedback & Support menu button.
	Url NullableString `json:"url,omitempty"`
	// Whether the support button is shown. The portal always asks for it to be shown.
	Visible *bool `json:"visible,omitempty"`
}

// NewFeedbackConfig instantiates a new FeedbackConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeedbackConfig() *FeedbackConfig {
	this := FeedbackConfig{}
	return &this
}

// NewFeedbackConfigWithDefaults instantiates a new FeedbackConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeedbackConfigWithDefaults() *FeedbackConfig {
	this := FeedbackConfig{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeedbackConfig) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeedbackConfig) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *FeedbackConfig) IsUrlSet() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *FeedbackConfig) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *FeedbackConfig) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *FeedbackConfig) UnsetUrl() {
	o.Url.Unset()
}

// GetVisible returns the Visible field value if set, zero value otherwise.
func (o *FeedbackConfig) GetVisible() bool {
	if o == nil || IsNil(o.Visible) {
		var ret bool
		return ret
	}
	return *o.Visible
}

// GetVisibleOk returns a tuple with the Visible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeedbackConfig) GetVisibleOk() (*bool, bool) {
	if o == nil || IsNil(o.Visible) {
		return nil, false
	}
	return o.Visible, true
}

// HasVisible returns a boolean if a field has been set.
func (o *FeedbackConfig) IsVisibleSet() bool {
	if o != nil && !IsNil(o.Visible) {
		return true
	}

	return false
}

// SetVisible gets a reference to the given bool and assigns it to the Visible field.
func (o *FeedbackConfig) SetVisible(v bool) {
	o.Visible = &v
}

func (o FeedbackConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeedbackConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if !IsNil(o.Visible) {
		toSerialize["visible"] = o.Visible
	}
	return toSerialize, nil
}

type NullableFeedbackConfig struct {
	value *FeedbackConfig
	isSet bool
}

func (v NullableFeedbackConfig) Get() *FeedbackConfig {
	return v.value
}

func (v *NullableFeedbackConfig) Set(val *FeedbackConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableFeedbackConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableFeedbackConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeedbackConfig(val *FeedbackConfig) *NullableFeedbackConfig {
	return &NullableFeedbackConfig{value: val, isSet: true}
}

func (v NullableFeedbackConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeedbackConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

