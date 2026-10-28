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

// checks if the DocServiceUrlDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocServiceUrlDto{}

// DocServiceUrlDto The document service location as this portal has it configured, together with the editor entry points a client  needs in order to open a document.
type DocServiceUrlDto struct {
	// The editor version the running Document Server reported. It is filled in only when the version was asked for,  and comes back empty otherwise. When the Document Server does not answer, a fallback version is reported  rather than an error, so a value here is no proof that the server is reachable.
	Version NullableString `json:"version"`
	// The absolute URL of the editor api script that a client has to load before it can open a document. It is  derived from the public Document Server address unless the deployment overrides it separately.
	DocServiceUrlApi NullableString `json:"docServiceUrlApi"`
	// The public Document Server address a browser loads the editor from. Empty means no document server is  configured for this portal, and documents cannot be opened for editing or viewing.
	DocServiceUrl NullableString `json:"docServiceUrl"`
	// The absolute URL of a page a client may load in advance to warm the editor scripts up. Loading it is optional  and changes nothing on the portal.
	DocServicePreloadUrl NullableString `json:"docServicePreloadUrl"`
	// The address the portal uses for its own server-to-server calls to the Document Server. When no private-network  address is configured, it repeats the public one.
	DocServiceUrlInternal NullableString `json:"docServiceUrlInternal"`
	// The address the Document Server is told to call this portal back on. Empty means nothing overrides it and the  portal's own resolved address is used.
	DocServicePortalUrl NullableString `json:"docServicePortalUrl"`
	// The name of the HTTP header that carries the signature on requests between the portal and the Document Server.  The secret itself is not part of the answer, so this only tells a client whether request signing is set up and  under which header.
	DocServiceSignatureHeader NullableString `json:"docServiceSignatureHeader"`
	// Whether the portal validates the TLS certificate of the Document Server. False means any certificate is  accepted, which is expected only in a test deployment.
	DocServiceSslVerification bool `json:"docServiceSslVerification"`
	// Whether every one of these settings is still the one the deployment ships with. False means at least one of  the addresses, the signature settings or SSL verification has been overridden for this portal.
	IsDefault bool `json:"isDefault"`
}

type _DocServiceUrlDto DocServiceUrlDto

// NewDocServiceUrlDto instantiates a new DocServiceUrlDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocServiceUrlDto(version NullableString, docServiceUrlApi NullableString, docServiceUrl NullableString, docServicePreloadUrl NullableString, docServiceUrlInternal NullableString, docServicePortalUrl NullableString, docServiceSignatureHeader NullableString, docServiceSslVerification bool, isDefault bool) *DocServiceUrlDto {
	this := DocServiceUrlDto{}
	this.Version = version
	this.DocServiceUrlApi = docServiceUrlApi
	this.DocServiceUrl = docServiceUrl
	this.DocServicePreloadUrl = docServicePreloadUrl
	this.DocServiceUrlInternal = docServiceUrlInternal
	this.DocServicePortalUrl = docServicePortalUrl
	this.DocServiceSignatureHeader = docServiceSignatureHeader
	this.DocServiceSslVerification = docServiceSslVerification
	this.IsDefault = isDefault
	return &this
}

// NewDocServiceUrlDtoWithDefaults instantiates a new DocServiceUrlDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocServiceUrlDtoWithDefaults() *DocServiceUrlDto {
	this := DocServiceUrlDto{}
	return &this
}

// GetVersion returns the Version field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DocServiceUrlDto) GetVersion() string {
	if o == nil || o.Version.Get() == nil {
		var ret string
		return ret
	}

	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocServiceUrlDto) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// SetVersion sets field value
func (o *DocServiceUrlDto) SetVersion(v string) {
	o.Version.Set(&v)
}

// GetDocServiceUrlApi returns the DocServiceUrlApi field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DocServiceUrlDto) GetDocServiceUrlApi() string {
	if o == nil || o.DocServiceUrlApi.Get() == nil {
		var ret string
		return ret
	}

	return *o.DocServiceUrlApi.Get()
}

// GetDocServiceUrlApiOk returns a tuple with the DocServiceUrlApi field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocServiceUrlDto) GetDocServiceUrlApiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DocServiceUrlApi.Get(), o.DocServiceUrlApi.IsSet()
}

// SetDocServiceUrlApi sets field value
func (o *DocServiceUrlDto) SetDocServiceUrlApi(v string) {
	o.DocServiceUrlApi.Set(&v)
}

// GetDocServiceUrl returns the DocServiceUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DocServiceUrlDto) GetDocServiceUrl() string {
	if o == nil || o.DocServiceUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.DocServiceUrl.Get()
}

// GetDocServiceUrlOk returns a tuple with the DocServiceUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocServiceUrlDto) GetDocServiceUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DocServiceUrl.Get(), o.DocServiceUrl.IsSet()
}

// SetDocServiceUrl sets field value
func (o *DocServiceUrlDto) SetDocServiceUrl(v string) {
	o.DocServiceUrl.Set(&v)
}

// GetDocServicePreloadUrl returns the DocServicePreloadUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DocServiceUrlDto) GetDocServicePreloadUrl() string {
	if o == nil || o.DocServicePreloadUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.DocServicePreloadUrl.Get()
}

// GetDocServicePreloadUrlOk returns a tuple with the DocServicePreloadUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocServiceUrlDto) GetDocServicePreloadUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DocServicePreloadUrl.Get(), o.DocServicePreloadUrl.IsSet()
}

// SetDocServicePreloadUrl sets field value
func (o *DocServiceUrlDto) SetDocServicePreloadUrl(v string) {
	o.DocServicePreloadUrl.Set(&v)
}

// GetDocServiceUrlInternal returns the DocServiceUrlInternal field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DocServiceUrlDto) GetDocServiceUrlInternal() string {
	if o == nil || o.DocServiceUrlInternal.Get() == nil {
		var ret string
		return ret
	}

	return *o.DocServiceUrlInternal.Get()
}

// GetDocServiceUrlInternalOk returns a tuple with the DocServiceUrlInternal field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocServiceUrlDto) GetDocServiceUrlInternalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DocServiceUrlInternal.Get(), o.DocServiceUrlInternal.IsSet()
}

// SetDocServiceUrlInternal sets field value
func (o *DocServiceUrlDto) SetDocServiceUrlInternal(v string) {
	o.DocServiceUrlInternal.Set(&v)
}

// GetDocServicePortalUrl returns the DocServicePortalUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DocServiceUrlDto) GetDocServicePortalUrl() string {
	if o == nil || o.DocServicePortalUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.DocServicePortalUrl.Get()
}

// GetDocServicePortalUrlOk returns a tuple with the DocServicePortalUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocServiceUrlDto) GetDocServicePortalUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DocServicePortalUrl.Get(), o.DocServicePortalUrl.IsSet()
}

// SetDocServicePortalUrl sets field value
func (o *DocServiceUrlDto) SetDocServicePortalUrl(v string) {
	o.DocServicePortalUrl.Set(&v)
}

// GetDocServiceSignatureHeader returns the DocServiceSignatureHeader field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DocServiceUrlDto) GetDocServiceSignatureHeader() string {
	if o == nil || o.DocServiceSignatureHeader.Get() == nil {
		var ret string
		return ret
	}

	return *o.DocServiceSignatureHeader.Get()
}

// GetDocServiceSignatureHeaderOk returns a tuple with the DocServiceSignatureHeader field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocServiceUrlDto) GetDocServiceSignatureHeaderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DocServiceSignatureHeader.Get(), o.DocServiceSignatureHeader.IsSet()
}

// SetDocServiceSignatureHeader sets field value
func (o *DocServiceUrlDto) SetDocServiceSignatureHeader(v string) {
	o.DocServiceSignatureHeader.Set(&v)
}

// GetDocServiceSslVerification returns the DocServiceSslVerification field value
func (o *DocServiceUrlDto) GetDocServiceSslVerification() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DocServiceSslVerification
}

// GetDocServiceSslVerificationOk returns a tuple with the DocServiceSslVerification field value
// and a boolean to check if the value has been set.
func (o *DocServiceUrlDto) GetDocServiceSslVerificationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DocServiceSslVerification, true
}

// SetDocServiceSslVerification sets field value
func (o *DocServiceUrlDto) SetDocServiceSslVerification(v bool) {
	o.DocServiceSslVerification = v
}

// GetIsDefault returns the IsDefault field value
func (o *DocServiceUrlDto) GetIsDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value
// and a boolean to check if the value has been set.
func (o *DocServiceUrlDto) GetIsDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDefault, true
}

// SetIsDefault sets field value
func (o *DocServiceUrlDto) SetIsDefault(v bool) {
	o.IsDefault = v
}

func (o DocServiceUrlDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocServiceUrlDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["version"] = o.Version.Get()
	toSerialize["docServiceUrlApi"] = o.DocServiceUrlApi.Get()
	toSerialize["docServiceUrl"] = o.DocServiceUrl.Get()
	toSerialize["docServicePreloadUrl"] = o.DocServicePreloadUrl.Get()
	toSerialize["docServiceUrlInternal"] = o.DocServiceUrlInternal.Get()
	toSerialize["docServicePortalUrl"] = o.DocServicePortalUrl.Get()
	toSerialize["docServiceSignatureHeader"] = o.DocServiceSignatureHeader.Get()
	toSerialize["docServiceSslVerification"] = o.DocServiceSslVerification
	toSerialize["isDefault"] = o.IsDefault
	return toSerialize, nil
}

func (o *DocServiceUrlDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"version",
		"docServiceUrlApi",
		"docServiceUrl",
		"docServicePreloadUrl",
		"docServiceUrlInternal",
		"docServicePortalUrl",
		"docServiceSignatureHeader",
		"docServiceSslVerification",
		"isDefault",
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

	varDocServiceUrlDto := _DocServiceUrlDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varDocServiceUrlDto)

	if err != nil {
		return err
	}

	*o = DocServiceUrlDto(varDocServiceUrlDto)

	return err
}

type NullableDocServiceUrlDto struct {
	value *DocServiceUrlDto
	isSet bool
}

func (v NullableDocServiceUrlDto) Get() *DocServiceUrlDto {
	return v.value
}

func (v *NullableDocServiceUrlDto) Set(val *DocServiceUrlDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDocServiceUrlDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDocServiceUrlDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocServiceUrlDto(val *DocServiceUrlDto) *NullableDocServiceUrlDto {
	return &NullableDocServiceUrlDto{value: val, isSet: true}
}

func (v NullableDocServiceUrlDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocServiceUrlDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

