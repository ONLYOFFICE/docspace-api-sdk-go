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

// checks if the FormGalleryDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormGalleryDto{}

// FormGalleryDto Where the ready-made form templates are served from, for browsing them and for submitting new ones.
type FormGalleryDto struct {
	// The path under `domain` that the gallery's own listing API is reached at. It is joined to `domain` by the  client; the portal only relays the values from its configuration.
	Path NullableString `json:"path"`
	// The address of the gallery service, which is a service of the vendor rather than part of the portal. Every  field of this object is empty on an installation that configures no gallery, and a client should then not  offer the gallery at all.
	Domain NullableString `json:"domain"`
	// The file extension to ask the gallery for, which decides which rendition of a template is downloaded when  several are published.
	Ext NullableString `json:"ext"`
	// The path used for submitting a form of one's own to the gallery, the counterpart of `path` for the upload  side. The four `upload` fields are empty when the installation allows browsing but not submitting.
	UploadPath NullableString `json:"uploadPath"`
	// The address the submission is sent to, which may differ from `domain`.
	UploadDomain NullableString `json:"uploadDomain"`
	// The file extension a submitted form has to carry.
	UploadExt NullableString `json:"uploadExt"`
	// The page a person is sent to in order to follow up on a submission, joined to `uploadDomain` the same way  as `uploadPath`.
	UploadDashboard NullableString `json:"uploadDashboard"`
}

type _FormGalleryDto FormGalleryDto

// NewFormGalleryDto instantiates a new FormGalleryDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormGalleryDto(path NullableString, domain NullableString, ext NullableString, uploadPath NullableString, uploadDomain NullableString, uploadExt NullableString, uploadDashboard NullableString) *FormGalleryDto {
	this := FormGalleryDto{}
	this.Path = path
	this.Domain = domain
	this.Ext = ext
	this.UploadPath = uploadPath
	this.UploadDomain = uploadDomain
	this.UploadExt = uploadExt
	this.UploadDashboard = uploadDashboard
	return &this
}

// NewFormGalleryDtoWithDefaults instantiates a new FormGalleryDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormGalleryDtoWithDefaults() *FormGalleryDto {
	this := FormGalleryDto{}
	return &this
}

// GetPath returns the Path field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FormGalleryDto) GetPath() string {
	if o == nil || o.Path.Get() == nil {
		var ret string
		return ret
	}

	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FormGalleryDto) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// SetPath sets field value
func (o *FormGalleryDto) SetPath(v string) {
	o.Path.Set(&v)
}

// GetDomain returns the Domain field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FormGalleryDto) GetDomain() string {
	if o == nil || o.Domain.Get() == nil {
		var ret string
		return ret
	}

	return *o.Domain.Get()
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FormGalleryDto) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Domain.Get(), o.Domain.IsSet()
}

// SetDomain sets field value
func (o *FormGalleryDto) SetDomain(v string) {
	o.Domain.Set(&v)
}

// GetExt returns the Ext field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FormGalleryDto) GetExt() string {
	if o == nil || o.Ext.Get() == nil {
		var ret string
		return ret
	}

	return *o.Ext.Get()
}

// GetExtOk returns a tuple with the Ext field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FormGalleryDto) GetExtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ext.Get(), o.Ext.IsSet()
}

// SetExt sets field value
func (o *FormGalleryDto) SetExt(v string) {
	o.Ext.Set(&v)
}

// GetUploadPath returns the UploadPath field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FormGalleryDto) GetUploadPath() string {
	if o == nil || o.UploadPath.Get() == nil {
		var ret string
		return ret
	}

	return *o.UploadPath.Get()
}

// GetUploadPathOk returns a tuple with the UploadPath field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FormGalleryDto) GetUploadPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UploadPath.Get(), o.UploadPath.IsSet()
}

// SetUploadPath sets field value
func (o *FormGalleryDto) SetUploadPath(v string) {
	o.UploadPath.Set(&v)
}

// GetUploadDomain returns the UploadDomain field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FormGalleryDto) GetUploadDomain() string {
	if o == nil || o.UploadDomain.Get() == nil {
		var ret string
		return ret
	}

	return *o.UploadDomain.Get()
}

// GetUploadDomainOk returns a tuple with the UploadDomain field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FormGalleryDto) GetUploadDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UploadDomain.Get(), o.UploadDomain.IsSet()
}

// SetUploadDomain sets field value
func (o *FormGalleryDto) SetUploadDomain(v string) {
	o.UploadDomain.Set(&v)
}

// GetUploadExt returns the UploadExt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FormGalleryDto) GetUploadExt() string {
	if o == nil || o.UploadExt.Get() == nil {
		var ret string
		return ret
	}

	return *o.UploadExt.Get()
}

// GetUploadExtOk returns a tuple with the UploadExt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FormGalleryDto) GetUploadExtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UploadExt.Get(), o.UploadExt.IsSet()
}

// SetUploadExt sets field value
func (o *FormGalleryDto) SetUploadExt(v string) {
	o.UploadExt.Set(&v)
}

// GetUploadDashboard returns the UploadDashboard field value
// If the value is explicit nil, the zero value for string will be returned
func (o *FormGalleryDto) GetUploadDashboard() string {
	if o == nil || o.UploadDashboard.Get() == nil {
		var ret string
		return ret
	}

	return *o.UploadDashboard.Get()
}

// GetUploadDashboardOk returns a tuple with the UploadDashboard field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FormGalleryDto) GetUploadDashboardOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UploadDashboard.Get(), o.UploadDashboard.IsSet()
}

// SetUploadDashboard sets field value
func (o *FormGalleryDto) SetUploadDashboard(v string) {
	o.UploadDashboard.Set(&v)
}

func (o FormGalleryDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormGalleryDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["path"] = o.Path.Get()
	toSerialize["domain"] = o.Domain.Get()
	toSerialize["ext"] = o.Ext.Get()
	toSerialize["uploadPath"] = o.UploadPath.Get()
	toSerialize["uploadDomain"] = o.UploadDomain.Get()
	toSerialize["uploadExt"] = o.UploadExt.Get()
	toSerialize["uploadDashboard"] = o.UploadDashboard.Get()
	return toSerialize, nil
}

func (o *FormGalleryDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"path",
		"domain",
		"ext",
		"uploadPath",
		"uploadDomain",
		"uploadExt",
		"uploadDashboard",
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

	varFormGalleryDto := _FormGalleryDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varFormGalleryDto)

	if err != nil {
		return err
	}

	*o = FormGalleryDto(varFormGalleryDto)

	return err
}

type NullableFormGalleryDto struct {
	value *FormGalleryDto
	isSet bool
}

func (v NullableFormGalleryDto) Get() *FormGalleryDto {
	return v.value
}

func (v *NullableFormGalleryDto) Set(val *FormGalleryDto) {
	v.value = val
	v.isSet = true
}

func (v NullableFormGalleryDto) IsSet() bool {
	return v.isSet
}

func (v *NullableFormGalleryDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormGalleryDto(val *FormGalleryDto) *NullableFormGalleryDto {
	return &NullableFormGalleryDto{value: val, isSet: true}
}

func (v NullableFormGalleryDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormGalleryDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

