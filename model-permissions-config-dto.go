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

// checks if the PermissionsConfigDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionsConfigDto{}

// PermissionsConfigDto The permissions configuration parameters.
type PermissionsConfigDto struct {
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

// NewPermissionsConfigDto instantiates a new PermissionsConfigDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionsConfigDto() *PermissionsConfigDto {
	this := PermissionsConfigDto{}
	return &this
}

// NewPermissionsConfigDtoWithDefaults instantiates a new PermissionsConfigDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionsConfigDtoWithDefaults() *PermissionsConfigDto {
	this := PermissionsConfigDto{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *PermissionsConfigDto) GetComment() bool {
	if o == nil || IsNil(o.Comment) {
		var ret bool
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsConfigDto) GetCommentOk() (*bool, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *PermissionsConfigDto) IsCommentSet() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given bool and assigns it to the Comment field.
func (o *PermissionsConfigDto) SetComment(v bool) {
	o.Comment = &v
}

// GetChat returns the Chat field value if set, zero value otherwise.
func (o *PermissionsConfigDto) GetChat() bool {
	if o == nil || IsNil(o.Chat) {
		var ret bool
		return ret
	}
	return *o.Chat
}

// GetChatOk returns a tuple with the Chat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsConfigDto) GetChatOk() (*bool, bool) {
	if o == nil || IsNil(o.Chat) {
		return nil, false
	}
	return o.Chat, true
}

// HasChat returns a boolean if a field has been set.
func (o *PermissionsConfigDto) IsChatSet() bool {
	if o != nil && !IsNil(o.Chat) {
		return true
	}

	return false
}

// SetChat gets a reference to the given bool and assigns it to the Chat field.
func (o *PermissionsConfigDto) SetChat(v bool) {
	o.Chat = &v
}

// GetDownload returns the Download field value if set, zero value otherwise.
func (o *PermissionsConfigDto) GetDownload() bool {
	if o == nil || IsNil(o.Download) {
		var ret bool
		return ret
	}
	return *o.Download
}

// GetDownloadOk returns a tuple with the Download field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsConfigDto) GetDownloadOk() (*bool, bool) {
	if o == nil || IsNil(o.Download) {
		return nil, false
	}
	return o.Download, true
}

// HasDownload returns a boolean if a field has been set.
func (o *PermissionsConfigDto) IsDownloadSet() bool {
	if o != nil && !IsNil(o.Download) {
		return true
	}

	return false
}

// SetDownload gets a reference to the given bool and assigns it to the Download field.
func (o *PermissionsConfigDto) SetDownload(v bool) {
	o.Download = &v
}

// GetEdit returns the Edit field value if set, zero value otherwise.
func (o *PermissionsConfigDto) GetEdit() bool {
	if o == nil || IsNil(o.Edit) {
		var ret bool
		return ret
	}
	return *o.Edit
}

// GetEditOk returns a tuple with the Edit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsConfigDto) GetEditOk() (*bool, bool) {
	if o == nil || IsNil(o.Edit) {
		return nil, false
	}
	return o.Edit, true
}

// HasEdit returns a boolean if a field has been set.
func (o *PermissionsConfigDto) IsEditSet() bool {
	if o != nil && !IsNil(o.Edit) {
		return true
	}

	return false
}

// SetEdit gets a reference to the given bool and assigns it to the Edit field.
func (o *PermissionsConfigDto) SetEdit(v bool) {
	o.Edit = &v
}

// GetFillForms returns the FillForms field value if set, zero value otherwise.
func (o *PermissionsConfigDto) GetFillForms() bool {
	if o == nil || IsNil(o.FillForms) {
		var ret bool
		return ret
	}
	return *o.FillForms
}

// GetFillFormsOk returns a tuple with the FillForms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsConfigDto) GetFillFormsOk() (*bool, bool) {
	if o == nil || IsNil(o.FillForms) {
		return nil, false
	}
	return o.FillForms, true
}

// HasFillForms returns a boolean if a field has been set.
func (o *PermissionsConfigDto) IsFillFormsSet() bool {
	if o != nil && !IsNil(o.FillForms) {
		return true
	}

	return false
}

// SetFillForms gets a reference to the given bool and assigns it to the FillForms field.
func (o *PermissionsConfigDto) SetFillForms(v bool) {
	o.FillForms = &v
}

// GetModifyFilter returns the ModifyFilter field value if set, zero value otherwise.
func (o *PermissionsConfigDto) GetModifyFilter() bool {
	if o == nil || IsNil(o.ModifyFilter) {
		var ret bool
		return ret
	}
	return *o.ModifyFilter
}

// GetModifyFilterOk returns a tuple with the ModifyFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsConfigDto) GetModifyFilterOk() (*bool, bool) {
	if o == nil || IsNil(o.ModifyFilter) {
		return nil, false
	}
	return o.ModifyFilter, true
}

// HasModifyFilter returns a boolean if a field has been set.
func (o *PermissionsConfigDto) IsModifyFilterSet() bool {
	if o != nil && !IsNil(o.ModifyFilter) {
		return true
	}

	return false
}

// SetModifyFilter gets a reference to the given bool and assigns it to the ModifyFilter field.
func (o *PermissionsConfigDto) SetModifyFilter(v bool) {
	o.ModifyFilter = &v
}

// GetProtect returns the Protect field value if set, zero value otherwise.
func (o *PermissionsConfigDto) GetProtect() bool {
	if o == nil || IsNil(o.Protect) {
		var ret bool
		return ret
	}
	return *o.Protect
}

// GetProtectOk returns a tuple with the Protect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsConfigDto) GetProtectOk() (*bool, bool) {
	if o == nil || IsNil(o.Protect) {
		return nil, false
	}
	return o.Protect, true
}

// HasProtect returns a boolean if a field has been set.
func (o *PermissionsConfigDto) IsProtectSet() bool {
	if o != nil && !IsNil(o.Protect) {
		return true
	}

	return false
}

// SetProtect gets a reference to the given bool and assigns it to the Protect field.
func (o *PermissionsConfigDto) SetProtect(v bool) {
	o.Protect = &v
}

// GetPrint returns the Print field value if set, zero value otherwise.
func (o *PermissionsConfigDto) GetPrint() bool {
	if o == nil || IsNil(o.Print) {
		var ret bool
		return ret
	}
	return *o.Print
}

// GetPrintOk returns a tuple with the Print field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsConfigDto) GetPrintOk() (*bool, bool) {
	if o == nil || IsNil(o.Print) {
		return nil, false
	}
	return o.Print, true
}

// HasPrint returns a boolean if a field has been set.
func (o *PermissionsConfigDto) IsPrintSet() bool {
	if o != nil && !IsNil(o.Print) {
		return true
	}

	return false
}

// SetPrint gets a reference to the given bool and assigns it to the Print field.
func (o *PermissionsConfigDto) SetPrint(v bool) {
	o.Print = &v
}

// GetReview returns the Review field value if set, zero value otherwise.
func (o *PermissionsConfigDto) GetReview() bool {
	if o == nil || IsNil(o.Review) {
		var ret bool
		return ret
	}
	return *o.Review
}

// GetReviewOk returns a tuple with the Review field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsConfigDto) GetReviewOk() (*bool, bool) {
	if o == nil || IsNil(o.Review) {
		return nil, false
	}
	return o.Review, true
}

// HasReview returns a boolean if a field has been set.
func (o *PermissionsConfigDto) IsReviewSet() bool {
	if o != nil && !IsNil(o.Review) {
		return true
	}

	return false
}

// SetReview gets a reference to the given bool and assigns it to the Review field.
func (o *PermissionsConfigDto) SetReview(v bool) {
	o.Review = &v
}

// GetCopy returns the Copy field value if set, zero value otherwise.
func (o *PermissionsConfigDto) GetCopy() bool {
	if o == nil || IsNil(o.Copy) {
		var ret bool
		return ret
	}
	return *o.Copy
}

// GetCopyOk returns a tuple with the Copy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsConfigDto) GetCopyOk() (*bool, bool) {
	if o == nil || IsNil(o.Copy) {
		return nil, false
	}
	return o.Copy, true
}

// HasCopy returns a boolean if a field has been set.
func (o *PermissionsConfigDto) IsCopySet() bool {
	if o != nil && !IsNil(o.Copy) {
		return true
	}

	return false
}

// SetCopy gets a reference to the given bool and assigns it to the Copy field.
func (o *PermissionsConfigDto) SetCopy(v bool) {
	o.Copy = &v
}

func (o PermissionsConfigDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermissionsConfigDto) ToMap() (map[string]interface{}, error) {
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

type NullablePermissionsConfigDto struct {
	value *PermissionsConfigDto
	isSet bool
}

func (v NullablePermissionsConfigDto) Get() *PermissionsConfigDto {
	return v.value
}

func (v *NullablePermissionsConfigDto) Set(val *PermissionsConfigDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionsConfigDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionsConfigDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionsConfigDto(val *PermissionsConfigDto) *NullablePermissionsConfigDto {
	return &NullablePermissionsConfigDto{value: val, isSet: true}
}

func (v NullablePermissionsConfigDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionsConfigDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

