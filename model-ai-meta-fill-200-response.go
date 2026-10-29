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

// checks if the AiMetaFill200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiMetaFill200Response{}

// AiMetaFill200Response struct for AiMetaFill200Response
type AiMetaFill200Response struct {
	// Whether the values were only proposed (`true`) or written as well.
	DryRun bool `json:"dryRun"`
	// The entry's metadata after the write (or as it stands, with `dryRun`), as `GET api/2.0/files/metadata/{kind}/{id}` returns it.
	Metadata AiEntryMetadata `json:"metadata"`
	// The values written, or proposed with `dryRun`: a typed value for a template field, a string for a custom field.
	Filled []AiMetaFill200ResponseFilledInner `json:"filled"`
	// The fields left as they were, with the reason for each.
	Unresolved []AiMetaFill200ResponseUnresolvedInner `json:"unresolved"`
	// The documents the values were read from.
	Sources []AiMetaFill200ResponseSourcesInner `json:"sources"`
	// The entries that were left out, with the reason for each.
	Skipped []AiMetaGenerate200ResponseSkippedInner `json:"skipped"`
}

type _AiMetaFill200Response AiMetaFill200Response

// NewAiMetaFill200Response instantiates a new AiMetaFill200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiMetaFill200Response(dryRun bool, metadata AiEntryMetadata, filled []AiMetaFill200ResponseFilledInner, unresolved []AiMetaFill200ResponseUnresolvedInner, sources []AiMetaFill200ResponseSourcesInner, skipped []AiMetaGenerate200ResponseSkippedInner) *AiMetaFill200Response {
	this := AiMetaFill200Response{}
	this.DryRun = dryRun
	this.Metadata = metadata
	this.Filled = filled
	this.Unresolved = unresolved
	this.Sources = sources
	this.Skipped = skipped
	return &this
}

// NewAiMetaFill200ResponseWithDefaults instantiates a new AiMetaFill200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiMetaFill200ResponseWithDefaults() *AiMetaFill200Response {
	this := AiMetaFill200Response{}
	return &this
}

// GetDryRun returns the DryRun field value
func (o *AiMetaFill200Response) GetDryRun() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value
// and a boolean to check if the value has been set.
func (o *AiMetaFill200Response) GetDryRunOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DryRun, true
}

// SetDryRun sets field value
func (o *AiMetaFill200Response) SetDryRun(v bool) {
	o.DryRun = v
}

// GetMetadata returns the Metadata field value
func (o *AiMetaFill200Response) GetMetadata() AiEntryMetadata {
	if o == nil {
		var ret AiEntryMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *AiMetaFill200Response) GetMetadataOk() (*AiEntryMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *AiMetaFill200Response) SetMetadata(v AiEntryMetadata) {
	o.Metadata = v
}

// GetFilled returns the Filled field value
func (o *AiMetaFill200Response) GetFilled() []AiMetaFill200ResponseFilledInner {
	if o == nil {
		var ret []AiMetaFill200ResponseFilledInner
		return ret
	}

	return o.Filled
}

// GetFilledOk returns a tuple with the Filled field value
// and a boolean to check if the value has been set.
func (o *AiMetaFill200Response) GetFilledOk() ([]AiMetaFill200ResponseFilledInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Filled, true
}

// SetFilled sets field value
func (o *AiMetaFill200Response) SetFilled(v []AiMetaFill200ResponseFilledInner) {
	o.Filled = v
}

// GetUnresolved returns the Unresolved field value
func (o *AiMetaFill200Response) GetUnresolved() []AiMetaFill200ResponseUnresolvedInner {
	if o == nil {
		var ret []AiMetaFill200ResponseUnresolvedInner
		return ret
	}

	return o.Unresolved
}

// GetUnresolvedOk returns a tuple with the Unresolved field value
// and a boolean to check if the value has been set.
func (o *AiMetaFill200Response) GetUnresolvedOk() ([]AiMetaFill200ResponseUnresolvedInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Unresolved, true
}

// SetUnresolved sets field value
func (o *AiMetaFill200Response) SetUnresolved(v []AiMetaFill200ResponseUnresolvedInner) {
	o.Unresolved = v
}

// GetSources returns the Sources field value
func (o *AiMetaFill200Response) GetSources() []AiMetaFill200ResponseSourcesInner {
	if o == nil {
		var ret []AiMetaFill200ResponseSourcesInner
		return ret
	}

	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value
// and a boolean to check if the value has been set.
func (o *AiMetaFill200Response) GetSourcesOk() ([]AiMetaFill200ResponseSourcesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sources, true
}

// SetSources sets field value
func (o *AiMetaFill200Response) SetSources(v []AiMetaFill200ResponseSourcesInner) {
	o.Sources = v
}

// GetSkipped returns the Skipped field value
func (o *AiMetaFill200Response) GetSkipped() []AiMetaGenerate200ResponseSkippedInner {
	if o == nil {
		var ret []AiMetaGenerate200ResponseSkippedInner
		return ret
	}

	return o.Skipped
}

// GetSkippedOk returns a tuple with the Skipped field value
// and a boolean to check if the value has been set.
func (o *AiMetaFill200Response) GetSkippedOk() ([]AiMetaGenerate200ResponseSkippedInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Skipped, true
}

// SetSkipped sets field value
func (o *AiMetaFill200Response) SetSkipped(v []AiMetaGenerate200ResponseSkippedInner) {
	o.Skipped = v
}

func (o AiMetaFill200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiMetaFill200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dryRun"] = o.DryRun
	toSerialize["metadata"] = o.Metadata
	toSerialize["filled"] = o.Filled
	toSerialize["unresolved"] = o.Unresolved
	toSerialize["sources"] = o.Sources
	toSerialize["skipped"] = o.Skipped
	return toSerialize, nil
}

func (o *AiMetaFill200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dryRun",
		"metadata",
		"filled",
		"unresolved",
		"sources",
		"skipped",
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

	varAiMetaFill200Response := _AiMetaFill200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiMetaFill200Response)

	if err != nil {
		return err
	}

	*o = AiMetaFill200Response(varAiMetaFill200Response)

	return err
}

type NullableAiMetaFill200Response struct {
	value *AiMetaFill200Response
	isSet bool
}

func (v NullableAiMetaFill200Response) Get() *AiMetaFill200Response {
	return v.value
}

func (v *NullableAiMetaFill200Response) Set(val *AiMetaFill200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAiMetaFill200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAiMetaFill200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiMetaFill200Response(val *AiMetaFill200Response) *NullableAiMetaFill200Response {
	return &NullableAiMetaFill200Response{value: val, isSet: true}
}

func (v NullableAiMetaFill200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiMetaFill200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

