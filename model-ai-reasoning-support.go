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

// checks if the AiReasoningSupport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiReasoningSupport{}

// AiReasoningSupport What one model can do with extended thinking. Providers describe each model through this shape so the UI offers only the choices that change the request, and the request builders clamp to the same table.
type AiReasoningSupport struct {
	// Whether the model can think at all. False hides the whole control.
	Thinks bool `json:"thinks"`
	// Whether `off` really turns thinking off. False means the model thinks always and off only drops to its lowest depth (or leaves the default depth, where there is no knob).
	CanDisable bool `json:"canDisable"`
	// Depths the model distinguishes, lowest first. Empty when thinking is an on/off switch with no depth (or the model doesn't think). A level not listed is clamped to the nearest one — see `clampReasoningLevel`.
	Depths []AiReasoningDepth `json:"depths"`
	// The depth the model runs at when nothing asks for one — what a stored `off` means on a model that cannot be switched off. Known only where a catalogue reports it (OpenRouter's `default_effort`); otherwise `DEFAULT_REASONING_LEVEL` clamped to `depths` is assumed.
	DefaultDepth *AiReasoningDepth `json:"defaultDepth,omitempty"`
}

type _AiReasoningSupport AiReasoningSupport

// NewAiReasoningSupport instantiates a new AiReasoningSupport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiReasoningSupport(thinks bool, canDisable bool, depths []AiReasoningDepth) *AiReasoningSupport {
	this := AiReasoningSupport{}
	this.Thinks = thinks
	this.CanDisable = canDisable
	this.Depths = depths
	return &this
}

// NewAiReasoningSupportWithDefaults instantiates a new AiReasoningSupport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiReasoningSupportWithDefaults() *AiReasoningSupport {
	this := AiReasoningSupport{}
	return &this
}

// GetThinks returns the Thinks field value
func (o *AiReasoningSupport) GetThinks() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Thinks
}

// GetThinksOk returns a tuple with the Thinks field value
// and a boolean to check if the value has been set.
func (o *AiReasoningSupport) GetThinksOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Thinks, true
}

// SetThinks sets field value
func (o *AiReasoningSupport) SetThinks(v bool) {
	o.Thinks = v
}

// GetCanDisable returns the CanDisable field value
func (o *AiReasoningSupport) GetCanDisable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanDisable
}

// GetCanDisableOk returns a tuple with the CanDisable field value
// and a boolean to check if the value has been set.
func (o *AiReasoningSupport) GetCanDisableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanDisable, true
}

// SetCanDisable sets field value
func (o *AiReasoningSupport) SetCanDisable(v bool) {
	o.CanDisable = v
}

// GetDepths returns the Depths field value
func (o *AiReasoningSupport) GetDepths() []AiReasoningDepth {
	if o == nil {
		var ret []AiReasoningDepth
		return ret
	}

	return o.Depths
}

// GetDepthsOk returns a tuple with the Depths field value
// and a boolean to check if the value has been set.
func (o *AiReasoningSupport) GetDepthsOk() ([]AiReasoningDepth, bool) {
	if o == nil {
		return nil, false
	}
	return o.Depths, true
}

// SetDepths sets field value
func (o *AiReasoningSupport) SetDepths(v []AiReasoningDepth) {
	o.Depths = v
}

// GetDefaultDepth returns the DefaultDepth field value if set, zero value otherwise.
func (o *AiReasoningSupport) GetDefaultDepth() AiReasoningDepth {
	if o == nil || IsNil(o.DefaultDepth) {
		var ret AiReasoningDepth
		return ret
	}
	return *o.DefaultDepth
}

// GetDefaultDepthOk returns a tuple with the DefaultDepth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiReasoningSupport) GetDefaultDepthOk() (*AiReasoningDepth, bool) {
	if o == nil || IsNil(o.DefaultDepth) {
		return nil, false
	}
	return o.DefaultDepth, true
}

// HasDefaultDepth returns a boolean if a field has been set.
func (o *AiReasoningSupport) IsDefaultDepthSet() bool {
	if o != nil && !IsNil(o.DefaultDepth) {
		return true
	}

	return false
}

// SetDefaultDepth gets a reference to the given AiReasoningDepth and assigns it to the DefaultDepth field.
func (o *AiReasoningSupport) SetDefaultDepth(v AiReasoningDepth) {
	o.DefaultDepth = &v
}

func (o AiReasoningSupport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiReasoningSupport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["thinks"] = o.Thinks
	toSerialize["canDisable"] = o.CanDisable
	toSerialize["depths"] = o.Depths
	if !IsNil(o.DefaultDepth) {
		toSerialize["defaultDepth"] = o.DefaultDepth
	}
	return toSerialize, nil
}

func (o *AiReasoningSupport) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"thinks",
		"canDisable",
		"depths",
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

	varAiReasoningSupport := _AiReasoningSupport{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiReasoningSupport)

	if err != nil {
		return err
	}

	*o = AiReasoningSupport(varAiReasoningSupport)

	return err
}

type NullableAiReasoningSupport struct {
	value *AiReasoningSupport
	isSet bool
}

func (v NullableAiReasoningSupport) Get() *AiReasoningSupport {
	return v.value
}

func (v *NullableAiReasoningSupport) Set(val *AiReasoningSupport) {
	v.value = val
	v.isSet = true
}

func (v NullableAiReasoningSupport) IsSet() bool {
	return v.isSet
}

func (v *NullableAiReasoningSupport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiReasoningSupport(val *AiReasoningSupport) *NullableAiReasoningSupport {
	return &NullableAiReasoningSupport{value: val, isSet: true}
}

func (v NullableAiReasoningSupport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiReasoningSupport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

