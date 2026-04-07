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

// checks if the UpdateClientRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateClientRequest{}

// UpdateClientRequest Client update request containing modified client details
type UpdateClientRequest struct {
	// The name of the client
	Name *string `json:"name,omitempty"`
	// The description of the client
	Description *string `json:"description,omitempty"`
	// The logo of the client in base64 format
	Logo *string `json:"logo,omitempty" validate:"regexp=^data:image\\/(?:png|jpeg|jpg|svg\\\\+xml);base64,.*.{1,}"`
	Public *bool `json:"public,omitempty"`
	// Indicates whether PKCE is allowed for the client
	AllowPkce *bool `json:"allow_pkce,omitempty"`
	// Indicates whether client is accessible by third-party tenants
	IsPublic *bool `json:"is_public,omitempty"`
	// The allowed origins for the client
	AllowedOrigins []string `json:"allowed_origins,omitempty"`
}

// NewUpdateClientRequest instantiates a new UpdateClientRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateClientRequest() *UpdateClientRequest {
	this := UpdateClientRequest{}
	return &this
}

// NewUpdateClientRequestWithDefaults instantiates a new UpdateClientRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateClientRequestWithDefaults() *UpdateClientRequest {
	this := UpdateClientRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateClientRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateClientRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateClientRequest) IsNameSet() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateClientRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateClientRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateClientRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateClientRequest) IsDescriptionSet() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateClientRequest) SetDescription(v string) {
	o.Description = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *UpdateClientRequest) GetLogo() string {
	if o == nil || IsNil(o.Logo) {
		var ret string
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateClientRequest) GetLogoOk() (*string, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *UpdateClientRequest) IsLogoSet() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given string and assigns it to the Logo field.
func (o *UpdateClientRequest) SetLogo(v string) {
	o.Logo = &v
}

// GetPublic returns the Public field value if set, zero value otherwise.
func (o *UpdateClientRequest) GetPublic() bool {
	if o == nil || IsNil(o.Public) {
		var ret bool
		return ret
	}
	return *o.Public
}

// GetPublicOk returns a tuple with the Public field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateClientRequest) GetPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.Public) {
		return nil, false
	}
	return o.Public, true
}

// HasPublic returns a boolean if a field has been set.
func (o *UpdateClientRequest) IsPublicSet() bool {
	if o != nil && !IsNil(o.Public) {
		return true
	}

	return false
}

// SetPublic gets a reference to the given bool and assigns it to the Public field.
func (o *UpdateClientRequest) SetPublic(v bool) {
	o.Public = &v
}

// GetAllowPkce returns the AllowPkce field value if set, zero value otherwise.
func (o *UpdateClientRequest) GetAllowPkce() bool {
	if o == nil || IsNil(o.AllowPkce) {
		var ret bool
		return ret
	}
	return *o.AllowPkce
}

// GetAllowPkceOk returns a tuple with the AllowPkce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateClientRequest) GetAllowPkceOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowPkce) {
		return nil, false
	}
	return o.AllowPkce, true
}

// HasAllowPkce returns a boolean if a field has been set.
func (o *UpdateClientRequest) IsAllowPkceSet() bool {
	if o != nil && !IsNil(o.AllowPkce) {
		return true
	}

	return false
}

// SetAllowPkce gets a reference to the given bool and assigns it to the AllowPkce field.
func (o *UpdateClientRequest) SetAllowPkce(v bool) {
	o.AllowPkce = &v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *UpdateClientRequest) GetIsPublic() bool {
	if o == nil || IsNil(o.IsPublic) {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateClientRequest) GetIsPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPublic) {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *UpdateClientRequest) IsIsPublicSet() bool {
	if o != nil && !IsNil(o.IsPublic) {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *UpdateClientRequest) SetIsPublic(v bool) {
	o.IsPublic = &v
}

// GetAllowedOrigins returns the AllowedOrigins field value if set, zero value otherwise.
func (o *UpdateClientRequest) GetAllowedOrigins() []string {
	if o == nil || IsNil(o.AllowedOrigins) {
		var ret []string
		return ret
	}
	return o.AllowedOrigins
}

// GetAllowedOriginsOk returns a tuple with the AllowedOrigins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateClientRequest) GetAllowedOriginsOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowedOrigins) {
		return nil, false
	}
	return o.AllowedOrigins, true
}

// HasAllowedOrigins returns a boolean if a field has been set.
func (o *UpdateClientRequest) IsAllowedOriginsSet() bool {
	if o != nil && !IsNil(o.AllowedOrigins) {
		return true
	}

	return false
}

// SetAllowedOrigins gets a reference to the given []string and assigns it to the AllowedOrigins field.
func (o *UpdateClientRequest) SetAllowedOrigins(v []string) {
	o.AllowedOrigins = v
}

func (o UpdateClientRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateClientRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}
	if !IsNil(o.Public) {
		toSerialize["public"] = o.Public
	}
	if !IsNil(o.AllowPkce) {
		toSerialize["allow_pkce"] = o.AllowPkce
	}
	if !IsNil(o.IsPublic) {
		toSerialize["is_public"] = o.IsPublic
	}
	if !IsNil(o.AllowedOrigins) {
		toSerialize["allowed_origins"] = o.AllowedOrigins
	}
	return toSerialize, nil
}

type NullableUpdateClientRequest struct {
	value *UpdateClientRequest
	isSet bool
}

func (v NullableUpdateClientRequest) Get() *UpdateClientRequest {
	return v.value
}

func (v *NullableUpdateClientRequest) Set(val *UpdateClientRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateClientRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateClientRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateClientRequest(val *UpdateClientRequest) *NullableUpdateClientRequest {
	return &NullableUpdateClientRequest{value: val, isSet: true}
}

func (v NullableUpdateClientRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateClientRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

