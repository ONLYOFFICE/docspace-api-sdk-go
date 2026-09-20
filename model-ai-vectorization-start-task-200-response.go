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
	"fmt"
)

// checks if the AiVectorizationStartTask200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AiVectorizationStartTask200Response{}

// AiVectorizationStartTask200Response struct for AiVectorizationStartTask200Response
type AiVectorizationStartTask200Response struct {
	// Envelope field from the internal service; 0 for this operation.
	Count int32 `json:"count"`
	// Envelope status flag from the internal service.
	Status int32 `json:"status"`
	// HTTP status the internal service answered with.
	StatusCode int32 `json:"statusCode"`
	AdditionalProperties map[string]interface{}
}

type _AiVectorizationStartTask200Response AiVectorizationStartTask200Response

// NewAiVectorizationStartTask200Response instantiates a new AiVectorizationStartTask200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiVectorizationStartTask200Response(count int32, status int32, statusCode int32) *AiVectorizationStartTask200Response {
	this := AiVectorizationStartTask200Response{}
	this.Count = count
	this.Status = status
	this.StatusCode = statusCode
	return &this
}

// NewAiVectorizationStartTask200ResponseWithDefaults instantiates a new AiVectorizationStartTask200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiVectorizationStartTask200ResponseWithDefaults() *AiVectorizationStartTask200Response {
	this := AiVectorizationStartTask200Response{}
	return &this
}

// GetCount returns the Count field value
func (o *AiVectorizationStartTask200Response) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *AiVectorizationStartTask200Response) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *AiVectorizationStartTask200Response) SetCount(v int32) {
	o.Count = v
}

// GetStatus returns the Status field value
func (o *AiVectorizationStartTask200Response) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AiVectorizationStartTask200Response) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AiVectorizationStartTask200Response) SetStatus(v int32) {
	o.Status = v
}

// GetStatusCode returns the StatusCode field value
func (o *AiVectorizationStartTask200Response) GetStatusCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value
// and a boolean to check if the value has been set.
func (o *AiVectorizationStartTask200Response) GetStatusCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusCode, true
}

// SetStatusCode sets field value
func (o *AiVectorizationStartTask200Response) SetStatusCode(v int32) {
	o.StatusCode = v
}

func (o AiVectorizationStartTask200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AiVectorizationStartTask200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count
	toSerialize["status"] = o.Status
	toSerialize["statusCode"] = o.StatusCode

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AiVectorizationStartTask200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"count",
		"status",
		"statusCode",
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

	varAiVectorizationStartTask200Response := _AiVectorizationStartTask200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAiVectorizationStartTask200Response)

	if err != nil {
		return err
	}

	*o = AiVectorizationStartTask200Response(varAiVectorizationStartTask200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "count")
		delete(additionalProperties, "status")
		delete(additionalProperties, "statusCode")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAiVectorizationStartTask200Response struct {
	value *AiVectorizationStartTask200Response
	isSet bool
}

func (v NullableAiVectorizationStartTask200Response) Get() *AiVectorizationStartTask200Response {
	return v.value
}

func (v *NullableAiVectorizationStartTask200Response) Set(val *AiVectorizationStartTask200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAiVectorizationStartTask200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAiVectorizationStartTask200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiVectorizationStartTask200Response(val *AiVectorizationStartTask200Response) *NullableAiVectorizationStartTask200Response {
	return &NullableAiVectorizationStartTask200Response{value: val, isSet: true}
}

func (v NullableAiVectorizationStartTask200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAiVectorizationStartTask200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

