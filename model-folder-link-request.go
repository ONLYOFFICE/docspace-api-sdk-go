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
	"time"
)

// checks if the FolderLinkRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FolderLinkRequest{}

// FolderLinkRequest The folder link parameters.
type FolderLinkRequest struct {
	// The folder link ID.
	LinkId *string `json:"linkId,omitempty"`
	// The link sharing rights.
	Access *FileShare `json:"access,omitempty"`
	// The link expiration date.
	ExpirationDate NullableTime `json:"expirationDate,omitempty"`
	// The link name.
	Title NullableString `json:"title,omitempty"`
	// The link password.
	Password NullableString `json:"password,omitempty"`
	// Specifies if downloading the file from the link is disabled or not.
	DenyDownload *bool `json:"denyDownload,omitempty"`
	// The link scope, whether it is internal or not.
	Internal *bool `json:"internal,omitempty"`
	// Specifies whether the folder link is primary or not.
	Primary *bool `json:"primary,omitempty"`
}

// NewFolderLinkRequest instantiates a new FolderLinkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFolderLinkRequest() *FolderLinkRequest {
	this := FolderLinkRequest{}
	return &this
}

// NewFolderLinkRequestWithDefaults instantiates a new FolderLinkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFolderLinkRequestWithDefaults() *FolderLinkRequest {
	this := FolderLinkRequest{}
	return &this
}

// GetLinkId returns the LinkId field value if set, zero value otherwise.
func (o *FolderLinkRequest) GetLinkId() string {
	if o == nil || IsNil(o.LinkId) {
		var ret string
		return ret
	}
	return *o.LinkId
}

// GetLinkIdOk returns a tuple with the LinkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderLinkRequest) GetLinkIdOk() (*string, bool) {
	if o == nil || IsNil(o.LinkId) {
		return nil, false
	}
	return o.LinkId, true
}

// HasLinkId returns a boolean if a field has been set.
func (o *FolderLinkRequest) IsLinkIdSet() bool {
	if o != nil && !IsNil(o.LinkId) {
		return true
	}

	return false
}

// SetLinkId gets a reference to the given string and assigns it to the LinkId field.
func (o *FolderLinkRequest) SetLinkId(v string) {
	o.LinkId = &v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *FolderLinkRequest) GetAccess() FileShare {
	if o == nil || IsNil(o.Access) {
		var ret FileShare
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderLinkRequest) GetAccessOk() (*FileShare, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *FolderLinkRequest) IsAccessSet() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given FileShare and assigns it to the Access field.
func (o *FolderLinkRequest) SetAccess(v FileShare) {
	o.Access = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderLinkRequest) GetExpirationDate() time.Time {
	if o == nil || IsNil(o.ExpirationDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDate.Get()
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderLinkRequest) GetExpirationDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpirationDate.Get(), o.ExpirationDate.IsSet()
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *FolderLinkRequest) IsExpirationDateSet() bool {
	if o != nil && o.ExpirationDate.IsSet() {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given NullableTime and assigns it to the ExpirationDate field.
func (o *FolderLinkRequest) SetExpirationDate(v time.Time) {
	o.ExpirationDate.Set(&v)
}
// SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil
func (o *FolderLinkRequest) SetExpirationDateNil() {
	o.ExpirationDate.Set(nil)
}

// UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
func (o *FolderLinkRequest) UnsetExpirationDate() {
	o.ExpirationDate.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderLinkRequest) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderLinkRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *FolderLinkRequest) IsTitleSet() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *FolderLinkRequest) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *FolderLinkRequest) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *FolderLinkRequest) UnsetTitle() {
	o.Title.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FolderLinkRequest) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FolderLinkRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *FolderLinkRequest) IsPasswordSet() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *FolderLinkRequest) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *FolderLinkRequest) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *FolderLinkRequest) UnsetPassword() {
	o.Password.Unset()
}

// GetDenyDownload returns the DenyDownload field value if set, zero value otherwise.
func (o *FolderLinkRequest) GetDenyDownload() bool {
	if o == nil || IsNil(o.DenyDownload) {
		var ret bool
		return ret
	}
	return *o.DenyDownload
}

// GetDenyDownloadOk returns a tuple with the DenyDownload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderLinkRequest) GetDenyDownloadOk() (*bool, bool) {
	if o == nil || IsNil(o.DenyDownload) {
		return nil, false
	}
	return o.DenyDownload, true
}

// HasDenyDownload returns a boolean if a field has been set.
func (o *FolderLinkRequest) IsDenyDownloadSet() bool {
	if o != nil && !IsNil(o.DenyDownload) {
		return true
	}

	return false
}

// SetDenyDownload gets a reference to the given bool and assigns it to the DenyDownload field.
func (o *FolderLinkRequest) SetDenyDownload(v bool) {
	o.DenyDownload = &v
}

// GetInternal returns the Internal field value if set, zero value otherwise.
func (o *FolderLinkRequest) GetInternal() bool {
	if o == nil || IsNil(o.Internal) {
		var ret bool
		return ret
	}
	return *o.Internal
}

// GetInternalOk returns a tuple with the Internal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderLinkRequest) GetInternalOk() (*bool, bool) {
	if o == nil || IsNil(o.Internal) {
		return nil, false
	}
	return o.Internal, true
}

// HasInternal returns a boolean if a field has been set.
func (o *FolderLinkRequest) IsInternalSet() bool {
	if o != nil && !IsNil(o.Internal) {
		return true
	}

	return false
}

// SetInternal gets a reference to the given bool and assigns it to the Internal field.
func (o *FolderLinkRequest) SetInternal(v bool) {
	o.Internal = &v
}

// GetPrimary returns the Primary field value if set, zero value otherwise.
func (o *FolderLinkRequest) GetPrimary() bool {
	if o == nil || IsNil(o.Primary) {
		var ret bool
		return ret
	}
	return *o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderLinkRequest) GetPrimaryOk() (*bool, bool) {
	if o == nil || IsNil(o.Primary) {
		return nil, false
	}
	return o.Primary, true
}

// HasPrimary returns a boolean if a field has been set.
func (o *FolderLinkRequest) IsPrimarySet() bool {
	if o != nil && !IsNil(o.Primary) {
		return true
	}

	return false
}

// SetPrimary gets a reference to the given bool and assigns it to the Primary field.
func (o *FolderLinkRequest) SetPrimary(v bool) {
	o.Primary = &v
}

func (o FolderLinkRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FolderLinkRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LinkId) {
		toSerialize["linkId"] = o.LinkId
	}
	if !IsNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	if o.ExpirationDate.IsSet() {
		toSerialize["expirationDate"] = o.ExpirationDate.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if !IsNil(o.DenyDownload) {
		toSerialize["denyDownload"] = o.DenyDownload
	}
	if !IsNil(o.Internal) {
		toSerialize["internal"] = o.Internal
	}
	if !IsNil(o.Primary) {
		toSerialize["primary"] = o.Primary
	}
	return toSerialize, nil
}

type NullableFolderLinkRequest struct {
	value *FolderLinkRequest
	isSet bool
}

func (v NullableFolderLinkRequest) Get() *FolderLinkRequest {
	return v.value
}

func (v *NullableFolderLinkRequest) Set(val *FolderLinkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFolderLinkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFolderLinkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFolderLinkRequest(val *FolderLinkRequest) *NullableFolderLinkRequest {
	return &NullableFolderLinkRequest{value: val, isSet: true}
}

func (v NullableFolderLinkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFolderLinkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

