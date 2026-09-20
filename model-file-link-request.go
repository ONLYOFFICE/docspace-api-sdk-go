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

// checks if the FileLinkRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileLinkRequest{}

// FileLinkRequest The settings of an external link to a file.
type FileLinkRequest struct {
	// The link to rewrite, as reported by `GET api/2.0/files/file/{id}/links`. An identifier that is not yet in use,  the empty one included, creates a link instead.
	LinkId *string `json:"linkId,omitempty"`
	// The rights the link grants to whoever follows it. The value that denies everything revokes the link.
	Access *FileShare `json:"access,omitempty"`
	// The moment the link stops working, read in the time zone of the portal. A date more than a few years ahead is  rejected as an invalid request; left out, the link does not expire on its own.
	ExpirationDate *ApiDateTime `json:"expirationDate,omitempty"`
	// The name the link carries in the sharing list of the file, for the people who manage it; it is not shown to  whoever follows the link.
	Title NullableString `json:"title,omitempty"`
	// Who may follow the link: `true` admits only accounts that are signed in to the portal, `false` admits anybody  who has the address.
	Internal *bool `json:"internal,omitempty"`
	// Whether this link becomes the primary link of the file - the one the Copy link action of a client hands out.  A file has one primary link at a time.
	Primary *bool `json:"primary,omitempty"`
	// What a visitor may do with the content: `true` leaves them with viewing in the browser, `false` lets them  download and print it as their rights allow.
	DenyDownload *bool `json:"denyDownload,omitempty"`
	// The secret a visitor has to type before the file opens; left out, the link opens without one.
	Password NullableString `json:"password,omitempty"`
}

// NewFileLinkRequest instantiates a new FileLinkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileLinkRequest() *FileLinkRequest {
	this := FileLinkRequest{}
	return &this
}

// NewFileLinkRequestWithDefaults instantiates a new FileLinkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileLinkRequestWithDefaults() *FileLinkRequest {
	this := FileLinkRequest{}
	return &this
}

// GetLinkId returns the LinkId field value if set, zero value otherwise.
func (o *FileLinkRequest) GetLinkId() string {
	if o == nil || IsNil(o.LinkId) {
		var ret string
		return ret
	}
	return *o.LinkId
}

// GetLinkIdOk returns a tuple with the LinkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileLinkRequest) GetLinkIdOk() (*string, bool) {
	if o == nil || IsNil(o.LinkId) {
		return nil, false
	}
	return o.LinkId, true
}

// HasLinkId returns a boolean if a field has been set.
func (o *FileLinkRequest) IsLinkIdSet() bool {
	if o != nil && !IsNil(o.LinkId) {
		return true
	}

	return false
}

// SetLinkId gets a reference to the given string and assigns it to the LinkId field.
func (o *FileLinkRequest) SetLinkId(v string) {
	o.LinkId = &v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *FileLinkRequest) GetAccess() FileShare {
	if o == nil || IsNil(o.Access) {
		var ret FileShare
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileLinkRequest) GetAccessOk() (*FileShare, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *FileLinkRequest) IsAccessSet() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given FileShare and assigns it to the Access field.
func (o *FileLinkRequest) SetAccess(v FileShare) {
	o.Access = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *FileLinkRequest) GetExpirationDate() ApiDateTime {
	if o == nil || IsNil(o.ExpirationDate) {
		var ret ApiDateTime
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileLinkRequest) GetExpirationDateOk() (*ApiDateTime, bool) {
	if o == nil || IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *FileLinkRequest) IsExpirationDateSet() bool {
	if o != nil && !IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given ApiDateTime and assigns it to the ExpirationDate field.
func (o *FileLinkRequest) SetExpirationDate(v ApiDateTime) {
	o.ExpirationDate = &v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileLinkRequest) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileLinkRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *FileLinkRequest) IsTitleSet() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *FileLinkRequest) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *FileLinkRequest) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *FileLinkRequest) UnsetTitle() {
	o.Title.Unset()
}

// GetInternal returns the Internal field value if set, zero value otherwise.
func (o *FileLinkRequest) GetInternal() bool {
	if o == nil || IsNil(o.Internal) {
		var ret bool
		return ret
	}
	return *o.Internal
}

// GetInternalOk returns a tuple with the Internal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileLinkRequest) GetInternalOk() (*bool, bool) {
	if o == nil || IsNil(o.Internal) {
		return nil, false
	}
	return o.Internal, true
}

// HasInternal returns a boolean if a field has been set.
func (o *FileLinkRequest) IsInternalSet() bool {
	if o != nil && !IsNil(o.Internal) {
		return true
	}

	return false
}

// SetInternal gets a reference to the given bool and assigns it to the Internal field.
func (o *FileLinkRequest) SetInternal(v bool) {
	o.Internal = &v
}

// GetPrimary returns the Primary field value if set, zero value otherwise.
func (o *FileLinkRequest) GetPrimary() bool {
	if o == nil || IsNil(o.Primary) {
		var ret bool
		return ret
	}
	return *o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileLinkRequest) GetPrimaryOk() (*bool, bool) {
	if o == nil || IsNil(o.Primary) {
		return nil, false
	}
	return o.Primary, true
}

// HasPrimary returns a boolean if a field has been set.
func (o *FileLinkRequest) IsPrimarySet() bool {
	if o != nil && !IsNil(o.Primary) {
		return true
	}

	return false
}

// SetPrimary gets a reference to the given bool and assigns it to the Primary field.
func (o *FileLinkRequest) SetPrimary(v bool) {
	o.Primary = &v
}

// GetDenyDownload returns the DenyDownload field value if set, zero value otherwise.
func (o *FileLinkRequest) GetDenyDownload() bool {
	if o == nil || IsNil(o.DenyDownload) {
		var ret bool
		return ret
	}
	return *o.DenyDownload
}

// GetDenyDownloadOk returns a tuple with the DenyDownload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileLinkRequest) GetDenyDownloadOk() (*bool, bool) {
	if o == nil || IsNil(o.DenyDownload) {
		return nil, false
	}
	return o.DenyDownload, true
}

// HasDenyDownload returns a boolean if a field has been set.
func (o *FileLinkRequest) IsDenyDownloadSet() bool {
	if o != nil && !IsNil(o.DenyDownload) {
		return true
	}

	return false
}

// SetDenyDownload gets a reference to the given bool and assigns it to the DenyDownload field.
func (o *FileLinkRequest) SetDenyDownload(v bool) {
	o.DenyDownload = &v
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileLinkRequest) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileLinkRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *FileLinkRequest) IsPasswordSet() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *FileLinkRequest) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *FileLinkRequest) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *FileLinkRequest) UnsetPassword() {
	o.Password.Unset()
}

func (o FileLinkRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileLinkRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LinkId) {
		toSerialize["linkId"] = o.LinkId
	}
	if !IsNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	if !IsNil(o.ExpirationDate) {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if !IsNil(o.Internal) {
		toSerialize["internal"] = o.Internal
	}
	if !IsNil(o.Primary) {
		toSerialize["primary"] = o.Primary
	}
	if !IsNil(o.DenyDownload) {
		toSerialize["denyDownload"] = o.DenyDownload
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	return toSerialize, nil
}

type NullableFileLinkRequest struct {
	value *FileLinkRequest
	isSet bool
}

func (v NullableFileLinkRequest) Get() *FileLinkRequest {
	return v.value
}

func (v *NullableFileLinkRequest) Set(val *FileLinkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFileLinkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFileLinkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileLinkRequest(val *FileLinkRequest) *NullableFileLinkRequest {
	return &NullableFileLinkRequest{value: val, isSet: true}
}

func (v NullableFileLinkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileLinkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

