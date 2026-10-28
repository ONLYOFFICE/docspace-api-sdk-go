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

// checks if the FileDtoAllOfViewAccessibility type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileDtoAllOfViewAccessibility{}

// FileDtoAllOfViewAccessibility Which ways of opening this format the portal supports at all - its own editor, the picture viewer, the media  player and so on. It answers whether the format can be shown, not whether this account may do it; rights are  reported in `security`.
type FileDtoAllOfViewAccessibility struct {
	ImageView *bool `json:"ImageView,omitempty"`
	MediaView *bool `json:"MediaView,omitempty"`
	WebView *bool `json:"WebView,omitempty"`
	WebEdit *bool `json:"WebEdit,omitempty"`
	WebReview *bool `json:"WebReview,omitempty"`
	WebCustomFilterEditing *bool `json:"WebCustomFilterEditing,omitempty"`
	WebRestrictedEditing *bool `json:"WebRestrictedEditing,omitempty"`
	WebComment *bool `json:"WebComment,omitempty"`
	CanConvert *bool `json:"CanConvert,omitempty"`
	MustConvert *bool `json:"MustConvert,omitempty"`
}

// NewFileDtoAllOfViewAccessibility instantiates a new FileDtoAllOfViewAccessibility object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileDtoAllOfViewAccessibility() *FileDtoAllOfViewAccessibility {
	this := FileDtoAllOfViewAccessibility{}
	return &this
}

// NewFileDtoAllOfViewAccessibilityWithDefaults instantiates a new FileDtoAllOfViewAccessibility object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileDtoAllOfViewAccessibilityWithDefaults() *FileDtoAllOfViewAccessibility {
	this := FileDtoAllOfViewAccessibility{}
	return &this
}

// GetImageView returns the ImageView field value if set, zero value otherwise.
func (o *FileDtoAllOfViewAccessibility) GetImageView() bool {
	if o == nil || IsNil(o.ImageView) {
		var ret bool
		return ret
	}
	return *o.ImageView
}

// GetImageViewOk returns a tuple with the ImageView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoAllOfViewAccessibility) GetImageViewOk() (*bool, bool) {
	if o == nil || IsNil(o.ImageView) {
		return nil, false
	}
	return o.ImageView, true
}

// HasImageView returns a boolean if a field has been set.
func (o *FileDtoAllOfViewAccessibility) IsImageViewSet() bool {
	if o != nil && !IsNil(o.ImageView) {
		return true
	}

	return false
}

// SetImageView gets a reference to the given bool and assigns it to the ImageView field.
func (o *FileDtoAllOfViewAccessibility) SetImageView(v bool) {
	o.ImageView = &v
}

// GetMediaView returns the MediaView field value if set, zero value otherwise.
func (o *FileDtoAllOfViewAccessibility) GetMediaView() bool {
	if o == nil || IsNil(o.MediaView) {
		var ret bool
		return ret
	}
	return *o.MediaView
}

// GetMediaViewOk returns a tuple with the MediaView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoAllOfViewAccessibility) GetMediaViewOk() (*bool, bool) {
	if o == nil || IsNil(o.MediaView) {
		return nil, false
	}
	return o.MediaView, true
}

// HasMediaView returns a boolean if a field has been set.
func (o *FileDtoAllOfViewAccessibility) IsMediaViewSet() bool {
	if o != nil && !IsNil(o.MediaView) {
		return true
	}

	return false
}

// SetMediaView gets a reference to the given bool and assigns it to the MediaView field.
func (o *FileDtoAllOfViewAccessibility) SetMediaView(v bool) {
	o.MediaView = &v
}

// GetWebView returns the WebView field value if set, zero value otherwise.
func (o *FileDtoAllOfViewAccessibility) GetWebView() bool {
	if o == nil || IsNil(o.WebView) {
		var ret bool
		return ret
	}
	return *o.WebView
}

// GetWebViewOk returns a tuple with the WebView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoAllOfViewAccessibility) GetWebViewOk() (*bool, bool) {
	if o == nil || IsNil(o.WebView) {
		return nil, false
	}
	return o.WebView, true
}

// HasWebView returns a boolean if a field has been set.
func (o *FileDtoAllOfViewAccessibility) IsWebViewSet() bool {
	if o != nil && !IsNil(o.WebView) {
		return true
	}

	return false
}

// SetWebView gets a reference to the given bool and assigns it to the WebView field.
func (o *FileDtoAllOfViewAccessibility) SetWebView(v bool) {
	o.WebView = &v
}

// GetWebEdit returns the WebEdit field value if set, zero value otherwise.
func (o *FileDtoAllOfViewAccessibility) GetWebEdit() bool {
	if o == nil || IsNil(o.WebEdit) {
		var ret bool
		return ret
	}
	return *o.WebEdit
}

// GetWebEditOk returns a tuple with the WebEdit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoAllOfViewAccessibility) GetWebEditOk() (*bool, bool) {
	if o == nil || IsNil(o.WebEdit) {
		return nil, false
	}
	return o.WebEdit, true
}

// HasWebEdit returns a boolean if a field has been set.
func (o *FileDtoAllOfViewAccessibility) IsWebEditSet() bool {
	if o != nil && !IsNil(o.WebEdit) {
		return true
	}

	return false
}

// SetWebEdit gets a reference to the given bool and assigns it to the WebEdit field.
func (o *FileDtoAllOfViewAccessibility) SetWebEdit(v bool) {
	o.WebEdit = &v
}

// GetWebReview returns the WebReview field value if set, zero value otherwise.
func (o *FileDtoAllOfViewAccessibility) GetWebReview() bool {
	if o == nil || IsNil(o.WebReview) {
		var ret bool
		return ret
	}
	return *o.WebReview
}

// GetWebReviewOk returns a tuple with the WebReview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoAllOfViewAccessibility) GetWebReviewOk() (*bool, bool) {
	if o == nil || IsNil(o.WebReview) {
		return nil, false
	}
	return o.WebReview, true
}

// HasWebReview returns a boolean if a field has been set.
func (o *FileDtoAllOfViewAccessibility) IsWebReviewSet() bool {
	if o != nil && !IsNil(o.WebReview) {
		return true
	}

	return false
}

// SetWebReview gets a reference to the given bool and assigns it to the WebReview field.
func (o *FileDtoAllOfViewAccessibility) SetWebReview(v bool) {
	o.WebReview = &v
}

// GetWebCustomFilterEditing returns the WebCustomFilterEditing field value if set, zero value otherwise.
func (o *FileDtoAllOfViewAccessibility) GetWebCustomFilterEditing() bool {
	if o == nil || IsNil(o.WebCustomFilterEditing) {
		var ret bool
		return ret
	}
	return *o.WebCustomFilterEditing
}

// GetWebCustomFilterEditingOk returns a tuple with the WebCustomFilterEditing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoAllOfViewAccessibility) GetWebCustomFilterEditingOk() (*bool, bool) {
	if o == nil || IsNil(o.WebCustomFilterEditing) {
		return nil, false
	}
	return o.WebCustomFilterEditing, true
}

// HasWebCustomFilterEditing returns a boolean if a field has been set.
func (o *FileDtoAllOfViewAccessibility) IsWebCustomFilterEditingSet() bool {
	if o != nil && !IsNil(o.WebCustomFilterEditing) {
		return true
	}

	return false
}

// SetWebCustomFilterEditing gets a reference to the given bool and assigns it to the WebCustomFilterEditing field.
func (o *FileDtoAllOfViewAccessibility) SetWebCustomFilterEditing(v bool) {
	o.WebCustomFilterEditing = &v
}

// GetWebRestrictedEditing returns the WebRestrictedEditing field value if set, zero value otherwise.
func (o *FileDtoAllOfViewAccessibility) GetWebRestrictedEditing() bool {
	if o == nil || IsNil(o.WebRestrictedEditing) {
		var ret bool
		return ret
	}
	return *o.WebRestrictedEditing
}

// GetWebRestrictedEditingOk returns a tuple with the WebRestrictedEditing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoAllOfViewAccessibility) GetWebRestrictedEditingOk() (*bool, bool) {
	if o == nil || IsNil(o.WebRestrictedEditing) {
		return nil, false
	}
	return o.WebRestrictedEditing, true
}

// HasWebRestrictedEditing returns a boolean if a field has been set.
func (o *FileDtoAllOfViewAccessibility) IsWebRestrictedEditingSet() bool {
	if o != nil && !IsNil(o.WebRestrictedEditing) {
		return true
	}

	return false
}

// SetWebRestrictedEditing gets a reference to the given bool and assigns it to the WebRestrictedEditing field.
func (o *FileDtoAllOfViewAccessibility) SetWebRestrictedEditing(v bool) {
	o.WebRestrictedEditing = &v
}

// GetWebComment returns the WebComment field value if set, zero value otherwise.
func (o *FileDtoAllOfViewAccessibility) GetWebComment() bool {
	if o == nil || IsNil(o.WebComment) {
		var ret bool
		return ret
	}
	return *o.WebComment
}

// GetWebCommentOk returns a tuple with the WebComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoAllOfViewAccessibility) GetWebCommentOk() (*bool, bool) {
	if o == nil || IsNil(o.WebComment) {
		return nil, false
	}
	return o.WebComment, true
}

// HasWebComment returns a boolean if a field has been set.
func (o *FileDtoAllOfViewAccessibility) IsWebCommentSet() bool {
	if o != nil && !IsNil(o.WebComment) {
		return true
	}

	return false
}

// SetWebComment gets a reference to the given bool and assigns it to the WebComment field.
func (o *FileDtoAllOfViewAccessibility) SetWebComment(v bool) {
	o.WebComment = &v
}

// GetCanConvert returns the CanConvert field value if set, zero value otherwise.
func (o *FileDtoAllOfViewAccessibility) GetCanConvert() bool {
	if o == nil || IsNil(o.CanConvert) {
		var ret bool
		return ret
	}
	return *o.CanConvert
}

// GetCanConvertOk returns a tuple with the CanConvert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoAllOfViewAccessibility) GetCanConvertOk() (*bool, bool) {
	if o == nil || IsNil(o.CanConvert) {
		return nil, false
	}
	return o.CanConvert, true
}

// HasCanConvert returns a boolean if a field has been set.
func (o *FileDtoAllOfViewAccessibility) IsCanConvertSet() bool {
	if o != nil && !IsNil(o.CanConvert) {
		return true
	}

	return false
}

// SetCanConvert gets a reference to the given bool and assigns it to the CanConvert field.
func (o *FileDtoAllOfViewAccessibility) SetCanConvert(v bool) {
	o.CanConvert = &v
}

// GetMustConvert returns the MustConvert field value if set, zero value otherwise.
func (o *FileDtoAllOfViewAccessibility) GetMustConvert() bool {
	if o == nil || IsNil(o.MustConvert) {
		var ret bool
		return ret
	}
	return *o.MustConvert
}

// GetMustConvertOk returns a tuple with the MustConvert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoAllOfViewAccessibility) GetMustConvertOk() (*bool, bool) {
	if o == nil || IsNil(o.MustConvert) {
		return nil, false
	}
	return o.MustConvert, true
}

// HasMustConvert returns a boolean if a field has been set.
func (o *FileDtoAllOfViewAccessibility) IsMustConvertSet() bool {
	if o != nil && !IsNil(o.MustConvert) {
		return true
	}

	return false
}

// SetMustConvert gets a reference to the given bool and assigns it to the MustConvert field.
func (o *FileDtoAllOfViewAccessibility) SetMustConvert(v bool) {
	o.MustConvert = &v
}

func (o FileDtoAllOfViewAccessibility) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileDtoAllOfViewAccessibility) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ImageView) {
		toSerialize["ImageView"] = o.ImageView
	}
	if !IsNil(o.MediaView) {
		toSerialize["MediaView"] = o.MediaView
	}
	if !IsNil(o.WebView) {
		toSerialize["WebView"] = o.WebView
	}
	if !IsNil(o.WebEdit) {
		toSerialize["WebEdit"] = o.WebEdit
	}
	if !IsNil(o.WebReview) {
		toSerialize["WebReview"] = o.WebReview
	}
	if !IsNil(o.WebCustomFilterEditing) {
		toSerialize["WebCustomFilterEditing"] = o.WebCustomFilterEditing
	}
	if !IsNil(o.WebRestrictedEditing) {
		toSerialize["WebRestrictedEditing"] = o.WebRestrictedEditing
	}
	if !IsNil(o.WebComment) {
		toSerialize["WebComment"] = o.WebComment
	}
	if !IsNil(o.CanConvert) {
		toSerialize["CanConvert"] = o.CanConvert
	}
	if !IsNil(o.MustConvert) {
		toSerialize["MustConvert"] = o.MustConvert
	}
	return toSerialize, nil
}

type NullableFileDtoAllOfViewAccessibility struct {
	value *FileDtoAllOfViewAccessibility
	isSet bool
}

func (v NullableFileDtoAllOfViewAccessibility) Get() *FileDtoAllOfViewAccessibility {
	return v.value
}

func (v *NullableFileDtoAllOfViewAccessibility) Set(val *FileDtoAllOfViewAccessibility) {
	v.value = val
	v.isSet = true
}

func (v NullableFileDtoAllOfViewAccessibility) IsSet() bool {
	return v.isSet
}

func (v *NullableFileDtoAllOfViewAccessibility) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileDtoAllOfViewAccessibility(val *FileDtoAllOfViewAccessibility) *NullableFileDtoAllOfViewAccessibility {
	return &NullableFileDtoAllOfViewAccessibility{value: val, isSet: true}
}

func (v NullableFileDtoAllOfViewAccessibility) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileDtoAllOfViewAccessibility) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

