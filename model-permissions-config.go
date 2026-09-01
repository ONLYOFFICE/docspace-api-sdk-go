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

// checks if the PermissionsConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionsConfig{}

// PermissionsConfig The permissions configuration parameters.
type PermissionsConfig struct {
	// Defines if the document can be commented or not.
	Comment *bool `json:"comment,omitempty"`
	// Defines if the chat functionality is enabled in the document or not.
	Chat *bool `json:"chat,omitempty"`
	// Defines if the document can be downloaded or only viewed or edited online.
	Download *bool `json:"download,omitempty"`
	// Defines if the document can be edited or only viewed.
	Edit *bool `json:"edit,omitempty"`
	// Defines if the forms can be filled.
	FillForms *bool `json:"fillForms,omitempty"`
	// Defines if the filter can be applied globally (true) affecting all the other users,  or locally (false), i.e. for the current user only.
	ModifyFilter *bool `json:"modifyFilter,omitempty"`
	// Defines if the Protection tab on the toolbar and the Protect button in the left menu are displayedor hidden.
	Protect *bool `json:"protect,omitempty"`
	// Defines if the document can be printed or not.
	Print *bool `json:"print,omitempty"`
	// Defines if the document can be reviewed or not.
	Review *bool `json:"review,omitempty"`
	// Defines if the content can be copied to the clipboard or not.
	Copy *bool `json:"copy,omitempty"`
}

// NewPermissionsConfig instantiates a new PermissionsConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionsConfig() *PermissionsConfig {
	this := PermissionsConfig{}
	return &this
}

// NewPermissionsConfigWithDefaults instantiates a new PermissionsConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionsConfigWithDefaults() *PermissionsConfig {
	this := PermissionsConfig{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *PermissionsConfig) GetComment() bool {
	if o == nil || IsNil(o.Comment) {
		var ret bool
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsConfig) GetCommentOk() (*bool, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *PermissionsConfig) IsCommentSet() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given bool and assigns it to the Comment field.
func (o *PermissionsConfig) SetComment(v bool) {
	o.Comment = &v
}

// GetChat returns the Chat field value if set, zero value otherwise.
func (o *PermissionsConfig) GetChat() bool {
	if o == nil || IsNil(o.Chat) {
		var ret bool
		return ret
	}
	return *o.Chat
}

// GetChatOk returns a tuple with the Chat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsConfig) GetChatOk() (*bool, bool) {
	if o == nil || IsNil(o.Chat) {
		return nil, false
	}
	return o.Chat, true
}

// HasChat returns a boolean if a field has been set.
func (o *PermissionsConfig) IsChatSet() bool {
	if o != nil && !IsNil(o.Chat) {
		return true
	}

	return false
}

// SetChat gets a reference to the given bool and assigns it to the Chat field.
func (o *PermissionsConfig) SetChat(v bool) {
	o.Chat = &v
}

// GetDownload returns the Download field value if set, zero value otherwise.
func (o *PermissionsConfig) GetDownload() bool {
	if o == nil || IsNil(o.Download) {
		var ret bool
		return ret
	}
	return *o.Download
}

// GetDownloadOk returns a tuple with the Download field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsConfig) GetDownloadOk() (*bool, bool) {
	if o == nil || IsNil(o.Download) {
		return nil, false
	}
	return o.Download, true
}

// HasDownload returns a boolean if a field has been set.
func (o *PermissionsConfig) IsDownloadSet() bool {
	if o != nil && !IsNil(o.Download) {
		return true
	}

	return false
}

// SetDownload gets a reference to the given bool and assigns it to the Download field.
func (o *PermissionsConfig) SetDownload(v bool) {
	o.Download = &v
}

// GetEdit returns the Edit field value if set, zero value otherwise.
func (o *PermissionsConfig) GetEdit() bool {
	if o == nil || IsNil(o.Edit) {
		var ret bool
		return ret
	}
	return *o.Edit
}

// GetEditOk returns a tuple with the Edit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsConfig) GetEditOk() (*bool, bool) {
	if o == nil || IsNil(o.Edit) {
		return nil, false
	}
	return o.Edit, true
}

// HasEdit returns a boolean if a field has been set.
func (o *PermissionsConfig) IsEditSet() bool {
	if o != nil && !IsNil(o.Edit) {
		return true
	}

	return false
}

// SetEdit gets a reference to the given bool and assigns it to the Edit field.
func (o *PermissionsConfig) SetEdit(v bool) {
	o.Edit = &v
}

// GetFillForms returns the FillForms field value if set, zero value otherwise.
func (o *PermissionsConfig) GetFillForms() bool {
	if o == nil || IsNil(o.FillForms) {
		var ret bool
		return ret
	}
	return *o.FillForms
}

// GetFillFormsOk returns a tuple with the FillForms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsConfig) GetFillFormsOk() (*bool, bool) {
	if o == nil || IsNil(o.FillForms) {
		return nil, false
	}
	return o.FillForms, true
}

// HasFillForms returns a boolean if a field has been set.
func (o *PermissionsConfig) IsFillFormsSet() bool {
	if o != nil && !IsNil(o.FillForms) {
		return true
	}

	return false
}

// SetFillForms gets a reference to the given bool and assigns it to the FillForms field.
func (o *PermissionsConfig) SetFillForms(v bool) {
	o.FillForms = &v
}

// GetModifyFilter returns the ModifyFilter field value if set, zero value otherwise.
func (o *PermissionsConfig) GetModifyFilter() bool {
	if o == nil || IsNil(o.ModifyFilter) {
		var ret bool
		return ret
	}
	return *o.ModifyFilter
}

// GetModifyFilterOk returns a tuple with the ModifyFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsConfig) GetModifyFilterOk() (*bool, bool) {
	if o == nil || IsNil(o.ModifyFilter) {
		return nil, false
	}
	return o.ModifyFilter, true
}

// HasModifyFilter returns a boolean if a field has been set.
func (o *PermissionsConfig) IsModifyFilterSet() bool {
	if o != nil && !IsNil(o.ModifyFilter) {
		return true
	}

	return false
}

// SetModifyFilter gets a reference to the given bool and assigns it to the ModifyFilter field.
func (o *PermissionsConfig) SetModifyFilter(v bool) {
	o.ModifyFilter = &v
}

// GetProtect returns the Protect field value if set, zero value otherwise.
func (o *PermissionsConfig) GetProtect() bool {
	if o == nil || IsNil(o.Protect) {
		var ret bool
		return ret
	}
	return *o.Protect
}

// GetProtectOk returns a tuple with the Protect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsConfig) GetProtectOk() (*bool, bool) {
	if o == nil || IsNil(o.Protect) {
		return nil, false
	}
	return o.Protect, true
}

// HasProtect returns a boolean if a field has been set.
func (o *PermissionsConfig) IsProtectSet() bool {
	if o != nil && !IsNil(o.Protect) {
		return true
	}

	return false
}

// SetProtect gets a reference to the given bool and assigns it to the Protect field.
func (o *PermissionsConfig) SetProtect(v bool) {
	o.Protect = &v
}

// GetPrint returns the Print field value if set, zero value otherwise.
func (o *PermissionsConfig) GetPrint() bool {
	if o == nil || IsNil(o.Print) {
		var ret bool
		return ret
	}
	return *o.Print
}

// GetPrintOk returns a tuple with the Print field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsConfig) GetPrintOk() (*bool, bool) {
	if o == nil || IsNil(o.Print) {
		return nil, false
	}
	return o.Print, true
}

// HasPrint returns a boolean if a field has been set.
func (o *PermissionsConfig) IsPrintSet() bool {
	if o != nil && !IsNil(o.Print) {
		return true
	}

	return false
}

// SetPrint gets a reference to the given bool and assigns it to the Print field.
func (o *PermissionsConfig) SetPrint(v bool) {
	o.Print = &v
}

// GetReview returns the Review field value if set, zero value otherwise.
func (o *PermissionsConfig) GetReview() bool {
	if o == nil || IsNil(o.Review) {
		var ret bool
		return ret
	}
	return *o.Review
}

// GetReviewOk returns a tuple with the Review field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsConfig) GetReviewOk() (*bool, bool) {
	if o == nil || IsNil(o.Review) {
		return nil, false
	}
	return o.Review, true
}

// HasReview returns a boolean if a field has been set.
func (o *PermissionsConfig) IsReviewSet() bool {
	if o != nil && !IsNil(o.Review) {
		return true
	}

	return false
}

// SetReview gets a reference to the given bool and assigns it to the Review field.
func (o *PermissionsConfig) SetReview(v bool) {
	o.Review = &v
}

// GetCopy returns the Copy field value if set, zero value otherwise.
func (o *PermissionsConfig) GetCopy() bool {
	if o == nil || IsNil(o.Copy) {
		var ret bool
		return ret
	}
	return *o.Copy
}

// GetCopyOk returns a tuple with the Copy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsConfig) GetCopyOk() (*bool, bool) {
	if o == nil || IsNil(o.Copy) {
		return nil, false
	}
	return o.Copy, true
}

// HasCopy returns a boolean if a field has been set.
func (o *PermissionsConfig) IsCopySet() bool {
	if o != nil && !IsNil(o.Copy) {
		return true
	}

	return false
}

// SetCopy gets a reference to the given bool and assigns it to the Copy field.
func (o *PermissionsConfig) SetCopy(v bool) {
	o.Copy = &v
}

func (o PermissionsConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermissionsConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.Chat) {
		toSerialize["chat"] = o.Chat
	}
	if !IsNil(o.Download) {
		toSerialize["download"] = o.Download
	}
	if !IsNil(o.Edit) {
		toSerialize["edit"] = o.Edit
	}
	if !IsNil(o.FillForms) {
		toSerialize["fillForms"] = o.FillForms
	}
	if !IsNil(o.ModifyFilter) {
		toSerialize["modifyFilter"] = o.ModifyFilter
	}
	if !IsNil(o.Protect) {
		toSerialize["protect"] = o.Protect
	}
	if !IsNil(o.Print) {
		toSerialize["print"] = o.Print
	}
	if !IsNil(o.Review) {
		toSerialize["review"] = o.Review
	}
	if !IsNil(o.Copy) {
		toSerialize["copy"] = o.Copy
	}
	return toSerialize, nil
}

type NullablePermissionsConfig struct {
	value *PermissionsConfig
	isSet bool
}

func (v NullablePermissionsConfig) Get() *PermissionsConfig {
	return v.value
}

func (v *NullablePermissionsConfig) Set(val *PermissionsConfig) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionsConfig) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionsConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionsConfig(val *PermissionsConfig) *NullablePermissionsConfig {
	return &NullablePermissionsConfig{value: val, isSet: true}
}

func (v NullablePermissionsConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionsConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

