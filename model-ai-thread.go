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

// checks if the AiThread type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiThread{}

// AiThread Chat conversation metadata. Represents a single chat session (thread).
type AiThread struct {
	// Unique thread identifier (UUID).
	ThreadId string `json:"threadId"`
	// Optional thread title. Auto-generated from the first message if not set.
	Title *string `json:"title,omitempty"`
	// Timestamp (ms since epoch) of the last message in this thread. Used for sorting.
	LastEditDate *float32 `json:"lastEditDate,omitempty"`
	// Provider configuration at the time of last message. Used for thread-level provider display.
	Provider *AiTProvider `json:"provider,omitempty"`
	// Model info at the time of last message.
	Model *AiModel `json:"model,omitempty"`
	// ID of the profile used for this thread. Links to `Profile.id`.
	ProfileId *string `json:"profileId,omitempty"`
}

type _AiThread AiThread

// NewAiThread instantiates a new AiThread object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiThread(threadId string) *AiThread {
	this := AiThread{}
	this.ThreadId = threadId
	return &this
}

// NewAiThreadWithDefaults instantiates a new AiThread object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiThreadWithDefaults() *AiThread {
	this := AiThread{}
	return &this
}

// GetThreadId returns the ThreadId field value
func (o *AiThread) GetThreadId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ThreadId
}

// GetThreadIdOk returns a tuple with the ThreadId field value
// and a boolean to check if the value has been set.
func (o *AiThread) GetThreadIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThreadId, true
}

// SetThreadId sets field value
func (o *AiThread) SetThreadId(v string) {
	o.ThreadId = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AiThread) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiThread) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AiThread) IsTitleSet() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AiThread) SetTitle(v string) {
	o.Title = &v
}

// GetLastEditDate returns the LastEditDate field value if set, zero value otherwise.
func (o *AiThread) GetLastEditDate() float32 {
	if o == nil || IsNil(o.LastEditDate) {
		var ret float32
		return ret
	}
	return *o.LastEditDate
}

// GetLastEditDateOk returns a tuple with the LastEditDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiThread) GetLastEditDateOk() (*float32, bool) {
	if o == nil || IsNil(o.LastEditDate) {
		return nil, false
	}
	return o.LastEditDate, true
}

// HasLastEditDate returns a boolean if a field has been set.
func (o *AiThread) IsLastEditDateSet() bool {
	if o != nil && !IsNil(o.LastEditDate) {
		return true
	}

	return false
}

// SetLastEditDate gets a reference to the given float32 and assigns it to the LastEditDate field.
func (o *AiThread) SetLastEditDate(v float32) {
	o.LastEditDate = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *AiThread) GetProvider() AiTProvider {
	if o == nil || IsNil(o.Provider) {
		var ret AiTProvider
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiThread) GetProviderOk() (*AiTProvider, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *AiThread) IsProviderSet() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given AiTProvider and assigns it to the Provider field.
func (o *AiThread) SetProvider(v AiTProvider) {
	o.Provider = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *AiThread) GetModel() AiModel {
	if o == nil || IsNil(o.Model) {
		var ret AiModel
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiThread) GetModelOk() (*AiModel, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *AiThread) IsModelSet() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given AiModel and assigns it to the Model field.
func (o *AiThread) SetModel(v AiModel) {
	o.Model = &v
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *AiThread) GetProfileId() string {
	if o == nil || IsNil(o.ProfileId) {
		var ret string
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiThread) GetProfileIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProfileId) {
		return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *AiThread) IsProfileIdSet() bool {
	if o != nil && !IsNil(o.ProfileId) {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given string and assigns it to the ProfileId field.
func (o *AiThread) SetProfileId(v string) {
	o.ProfileId = &v
}

func (o AiThread) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiThread) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["threadId"] = o.ThreadId
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.LastEditDate) {
		toSerialize["lastEditDate"] = o.LastEditDate
	}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !IsNil(o.ProfileId) {
		toSerialize["profileId"] = o.ProfileId
	}
	return toSerialize, nil
}

func (o *AiThread) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"threadId",
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

	varAiThread := _AiThread{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiThread)

	if err != nil {
		return err
	}

	*o = AiThread(varAiThread)

	return err
}

type NullableAiThread struct {
	value *AiThread
	isSet bool
}

func (v NullableAiThread) Get() *AiThread {
	return v.value
}

func (v *NullableAiThread) Set(val *AiThread) {
	v.value = val
	v.isSet = true
}

func (v NullableAiThread) IsSet() bool {
	return v.isSet
}

func (v *NullableAiThread) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiThread(val *AiThread) *NullableAiThread {
	return &NullableAiThread{value: val, isSet: true}
}

func (v NullableAiThread) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiThread) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

