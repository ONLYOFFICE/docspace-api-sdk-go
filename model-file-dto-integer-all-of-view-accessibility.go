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

// checks if the FileDtoIntegerAllOfViewAccessibility type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileDtoIntegerAllOfViewAccessibility{}

// FileDtoIntegerAllOfViewAccessibility The file accessibility.
type FileDtoIntegerAllOfViewAccessibility struct {
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

// NewFileDtoIntegerAllOfViewAccessibility instantiates a new FileDtoIntegerAllOfViewAccessibility object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileDtoIntegerAllOfViewAccessibility() *FileDtoIntegerAllOfViewAccessibility {
	this := FileDtoIntegerAllOfViewAccessibility{}
	return &this
}

// NewFileDtoIntegerAllOfViewAccessibilityWithDefaults instantiates a new FileDtoIntegerAllOfViewAccessibility object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileDtoIntegerAllOfViewAccessibilityWithDefaults() *FileDtoIntegerAllOfViewAccessibility {
	this := FileDtoIntegerAllOfViewAccessibility{}
	return &this
}

// GetImageView returns the ImageView field value if set, zero value otherwise.
func (o *FileDtoIntegerAllOfViewAccessibility) GetImageView() bool {
	if o == nil || IsNil(o.ImageView) {
		var ret bool
		return ret
	}
	return *o.ImageView
}

// GetImageViewOk returns a tuple with the ImageView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoIntegerAllOfViewAccessibility) GetImageViewOk() (*bool, bool) {
	if o == nil || IsNil(o.ImageView) {
		return nil, false
	}
	return o.ImageView, true
}

// HasImageView returns a boolean if a field has been set.
func (o *FileDtoIntegerAllOfViewAccessibility) IsImageViewSet() bool {
	if o != nil && !IsNil(o.ImageView) {
		return true
	}

	return false
}

// SetImageView gets a reference to the given bool and assigns it to the ImageView field.
func (o *FileDtoIntegerAllOfViewAccessibility) SetImageView(v bool) {
	o.ImageView = &v
}

// GetMediaView returns the MediaView field value if set, zero value otherwise.
func (o *FileDtoIntegerAllOfViewAccessibility) GetMediaView() bool {
	if o == nil || IsNil(o.MediaView) {
		var ret bool
		return ret
	}
	return *o.MediaView
}

// GetMediaViewOk returns a tuple with the MediaView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoIntegerAllOfViewAccessibility) GetMediaViewOk() (*bool, bool) {
	if o == nil || IsNil(o.MediaView) {
		return nil, false
	}
	return o.MediaView, true
}

// HasMediaView returns a boolean if a field has been set.
func (o *FileDtoIntegerAllOfViewAccessibility) IsMediaViewSet() bool {
	if o != nil && !IsNil(o.MediaView) {
		return true
	}

	return false
}

// SetMediaView gets a reference to the given bool and assigns it to the MediaView field.
func (o *FileDtoIntegerAllOfViewAccessibility) SetMediaView(v bool) {
	o.MediaView = &v
}

// GetWebView returns the WebView field value if set, zero value otherwise.
func (o *FileDtoIntegerAllOfViewAccessibility) GetWebView() bool {
	if o == nil || IsNil(o.WebView) {
		var ret bool
		return ret
	}
	return *o.WebView
}

// GetWebViewOk returns a tuple with the WebView field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoIntegerAllOfViewAccessibility) GetWebViewOk() (*bool, bool) {
	if o == nil || IsNil(o.WebView) {
		return nil, false
	}
	return o.WebView, true
}

// HasWebView returns a boolean if a field has been set.
func (o *FileDtoIntegerAllOfViewAccessibility) IsWebViewSet() bool {
	if o != nil && !IsNil(o.WebView) {
		return true
	}

	return false
}

// SetWebView gets a reference to the given bool and assigns it to the WebView field.
func (o *FileDtoIntegerAllOfViewAccessibility) SetWebView(v bool) {
	o.WebView = &v
}

// GetWebEdit returns the WebEdit field value if set, zero value otherwise.
func (o *FileDtoIntegerAllOfViewAccessibility) GetWebEdit() bool {
	if o == nil || IsNil(o.WebEdit) {
		var ret bool
		return ret
	}
	return *o.WebEdit
}

// GetWebEditOk returns a tuple with the WebEdit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoIntegerAllOfViewAccessibility) GetWebEditOk() (*bool, bool) {
	if o == nil || IsNil(o.WebEdit) {
		return nil, false
	}
	return o.WebEdit, true
}

// HasWebEdit returns a boolean if a field has been set.
func (o *FileDtoIntegerAllOfViewAccessibility) IsWebEditSet() bool {
	if o != nil && !IsNil(o.WebEdit) {
		return true
	}

	return false
}

// SetWebEdit gets a reference to the given bool and assigns it to the WebEdit field.
func (o *FileDtoIntegerAllOfViewAccessibility) SetWebEdit(v bool) {
	o.WebEdit = &v
}

// GetWebReview returns the WebReview field value if set, zero value otherwise.
func (o *FileDtoIntegerAllOfViewAccessibility) GetWebReview() bool {
	if o == nil || IsNil(o.WebReview) {
		var ret bool
		return ret
	}
	return *o.WebReview
}

// GetWebReviewOk returns a tuple with the WebReview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoIntegerAllOfViewAccessibility) GetWebReviewOk() (*bool, bool) {
	if o == nil || IsNil(o.WebReview) {
		return nil, false
	}
	return o.WebReview, true
}

// HasWebReview returns a boolean if a field has been set.
func (o *FileDtoIntegerAllOfViewAccessibility) IsWebReviewSet() bool {
	if o != nil && !IsNil(o.WebReview) {
		return true
	}

	return false
}

// SetWebReview gets a reference to the given bool and assigns it to the WebReview field.
func (o *FileDtoIntegerAllOfViewAccessibility) SetWebReview(v bool) {
	o.WebReview = &v
}

// GetWebCustomFilterEditing returns the WebCustomFilterEditing field value if set, zero value otherwise.
func (o *FileDtoIntegerAllOfViewAccessibility) GetWebCustomFilterEditing() bool {
	if o == nil || IsNil(o.WebCustomFilterEditing) {
		var ret bool
		return ret
	}
	return *o.WebCustomFilterEditing
}

// GetWebCustomFilterEditingOk returns a tuple with the WebCustomFilterEditing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoIntegerAllOfViewAccessibility) GetWebCustomFilterEditingOk() (*bool, bool) {
	if o == nil || IsNil(o.WebCustomFilterEditing) {
		return nil, false
	}
	return o.WebCustomFilterEditing, true
}

// HasWebCustomFilterEditing returns a boolean if a field has been set.
func (o *FileDtoIntegerAllOfViewAccessibility) IsWebCustomFilterEditingSet() bool {
	if o != nil && !IsNil(o.WebCustomFilterEditing) {
		return true
	}

	return false
}

// SetWebCustomFilterEditing gets a reference to the given bool and assigns it to the WebCustomFilterEditing field.
func (o *FileDtoIntegerAllOfViewAccessibility) SetWebCustomFilterEditing(v bool) {
	o.WebCustomFilterEditing = &v
}

// GetWebRestrictedEditing returns the WebRestrictedEditing field value if set, zero value otherwise.
func (o *FileDtoIntegerAllOfViewAccessibility) GetWebRestrictedEditing() bool {
	if o == nil || IsNil(o.WebRestrictedEditing) {
		var ret bool
		return ret
	}
	return *o.WebRestrictedEditing
}

// GetWebRestrictedEditingOk returns a tuple with the WebRestrictedEditing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoIntegerAllOfViewAccessibility) GetWebRestrictedEditingOk() (*bool, bool) {
	if o == nil || IsNil(o.WebRestrictedEditing) {
		return nil, false
	}
	return o.WebRestrictedEditing, true
}

// HasWebRestrictedEditing returns a boolean if a field has been set.
func (o *FileDtoIntegerAllOfViewAccessibility) IsWebRestrictedEditingSet() bool {
	if o != nil && !IsNil(o.WebRestrictedEditing) {
		return true
	}

	return false
}

// SetWebRestrictedEditing gets a reference to the given bool and assigns it to the WebRestrictedEditing field.
func (o *FileDtoIntegerAllOfViewAccessibility) SetWebRestrictedEditing(v bool) {
	o.WebRestrictedEditing = &v
}

// GetWebComment returns the WebComment field value if set, zero value otherwise.
func (o *FileDtoIntegerAllOfViewAccessibility) GetWebComment() bool {
	if o == nil || IsNil(o.WebComment) {
		var ret bool
		return ret
	}
	return *o.WebComment
}

// GetWebCommentOk returns a tuple with the WebComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoIntegerAllOfViewAccessibility) GetWebCommentOk() (*bool, bool) {
	if o == nil || IsNil(o.WebComment) {
		return nil, false
	}
	return o.WebComment, true
}

// HasWebComment returns a boolean if a field has been set.
func (o *FileDtoIntegerAllOfViewAccessibility) IsWebCommentSet() bool {
	if o != nil && !IsNil(o.WebComment) {
		return true
	}

	return false
}

// SetWebComment gets a reference to the given bool and assigns it to the WebComment field.
func (o *FileDtoIntegerAllOfViewAccessibility) SetWebComment(v bool) {
	o.WebComment = &v
}

// GetCanConvert returns the CanConvert field value if set, zero value otherwise.
func (o *FileDtoIntegerAllOfViewAccessibility) GetCanConvert() bool {
	if o == nil || IsNil(o.CanConvert) {
		var ret bool
		return ret
	}
	return *o.CanConvert
}

// GetCanConvertOk returns a tuple with the CanConvert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoIntegerAllOfViewAccessibility) GetCanConvertOk() (*bool, bool) {
	if o == nil || IsNil(o.CanConvert) {
		return nil, false
	}
	return o.CanConvert, true
}

// HasCanConvert returns a boolean if a field has been set.
func (o *FileDtoIntegerAllOfViewAccessibility) IsCanConvertSet() bool {
	if o != nil && !IsNil(o.CanConvert) {
		return true
	}

	return false
}

// SetCanConvert gets a reference to the given bool and assigns it to the CanConvert field.
func (o *FileDtoIntegerAllOfViewAccessibility) SetCanConvert(v bool) {
	o.CanConvert = &v
}

// GetMustConvert returns the MustConvert field value if set, zero value otherwise.
func (o *FileDtoIntegerAllOfViewAccessibility) GetMustConvert() bool {
	if o == nil || IsNil(o.MustConvert) {
		var ret bool
		return ret
	}
	return *o.MustConvert
}

// GetMustConvertOk returns a tuple with the MustConvert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDtoIntegerAllOfViewAccessibility) GetMustConvertOk() (*bool, bool) {
	if o == nil || IsNil(o.MustConvert) {
		return nil, false
	}
	return o.MustConvert, true
}

// HasMustConvert returns a boolean if a field has been set.
func (o *FileDtoIntegerAllOfViewAccessibility) IsMustConvertSet() bool {
	if o != nil && !IsNil(o.MustConvert) {
		return true
	}

	return false
}

// SetMustConvert gets a reference to the given bool and assigns it to the MustConvert field.
func (o *FileDtoIntegerAllOfViewAccessibility) SetMustConvert(v bool) {
	o.MustConvert = &v
}

func (o FileDtoIntegerAllOfViewAccessibility) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileDtoIntegerAllOfViewAccessibility) ToMap() (map[string]interface{}, error) {
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

type NullableFileDtoIntegerAllOfViewAccessibility struct {
	value *FileDtoIntegerAllOfViewAccessibility
	isSet bool
}

func (v NullableFileDtoIntegerAllOfViewAccessibility) Get() *FileDtoIntegerAllOfViewAccessibility {
	return v.value
}

func (v *NullableFileDtoIntegerAllOfViewAccessibility) Set(val *FileDtoIntegerAllOfViewAccessibility) {
	v.value = val
	v.isSet = true
}

func (v NullableFileDtoIntegerAllOfViewAccessibility) IsSet() bool {
	return v.isSet
}

func (v *NullableFileDtoIntegerAllOfViewAccessibility) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileDtoIntegerAllOfViewAccessibility(val *FileDtoIntegerAllOfViewAccessibility) *NullableFileDtoIntegerAllOfViewAccessibility {
	return &NullableFileDtoIntegerAllOfViewAccessibility{value: val, isSet: true}
}

func (v NullableFileDtoIntegerAllOfViewAccessibility) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileDtoIntegerAllOfViewAccessibility) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

