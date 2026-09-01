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

// checks if the GetPortalPrices200ResponseLinksInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPortalPrices200ResponseLinksInner{}

// GetPortalPrices200ResponseLinksInner struct for GetPortalPrices200ResponseLinksInner
type GetPortalPrices200ResponseLinksInner struct {
	// URL of the link
	Href *string `json:"href,omitempty"`
	// Action associated with the link
	Action *string `json:"action,omitempty"`
}

// NewGetPortalPrices200ResponseLinksInner instantiates a new GetPortalPrices200ResponseLinksInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPortalPrices200ResponseLinksInner() *GetPortalPrices200ResponseLinksInner {
	this := GetPortalPrices200ResponseLinksInner{}
	return &this
}

// NewGetPortalPrices200ResponseLinksInnerWithDefaults instantiates a new GetPortalPrices200ResponseLinksInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPortalPrices200ResponseLinksInnerWithDefaults() *GetPortalPrices200ResponseLinksInner {
	this := GetPortalPrices200ResponseLinksInner{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *GetPortalPrices200ResponseLinksInner) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortalPrices200ResponseLinksInner) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *GetPortalPrices200ResponseLinksInner) IsHrefSet() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *GetPortalPrices200ResponseLinksInner) SetHref(v string) {
	o.Href = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *GetPortalPrices200ResponseLinksInner) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPortalPrices200ResponseLinksInner) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *GetPortalPrices200ResponseLinksInner) IsActionSet() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *GetPortalPrices200ResponseLinksInner) SetAction(v string) {
	o.Action = &v
}

func (o GetPortalPrices200ResponseLinksInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPortalPrices200ResponseLinksInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	return toSerialize, nil
}

type NullableGetPortalPrices200ResponseLinksInner struct {
	value *GetPortalPrices200ResponseLinksInner
	isSet bool
}

func (v NullableGetPortalPrices200ResponseLinksInner) Get() *GetPortalPrices200ResponseLinksInner {
	return v.value
}

func (v *NullableGetPortalPrices200ResponseLinksInner) Set(val *GetPortalPrices200ResponseLinksInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPortalPrices200ResponseLinksInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPortalPrices200ResponseLinksInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPortalPrices200ResponseLinksInner(val *GetPortalPrices200ResponseLinksInner) *NullableGetPortalPrices200ResponseLinksInner {
	return &NullableGetPortalPrices200ResponseLinksInner{value: val, isSet: true}
}

func (v NullableGetPortalPrices200ResponseLinksInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPortalPrices200ResponseLinksInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

