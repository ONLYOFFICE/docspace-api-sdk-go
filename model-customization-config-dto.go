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

// checks if the CustomizationConfigDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomizationConfigDto{}

// CustomizationConfigDto The customization config parameters.
type CustomizationConfigDto struct {
	// Specifies if the customization is about.
	About *bool `json:"about,omitempty"`
	Customer *CustomerConfigDto `json:"customer,omitempty"`
	Anonymous *AnonymousConfigDto `json:"anonymous,omitempty"`
	Feedback *FeedbackConfig `json:"feedback,omitempty"`
	// Specifies if the customization should be force saved.
	Forcesave NullableBool `json:"forcesave,omitempty"`
	Goback *GobackConfig `json:"goback,omitempty"`
	Review *ReviewConfig `json:"review,omitempty"`
	Logo *LogoConfigDto `json:"logo,omitempty"`
	// Specifies if the share should be mentioned.
	MentionShare *bool `json:"mentionShare,omitempty"`
	SubmitForm *SubmitForm `json:"submitForm,omitempty"`
	StartFillingForm *StartFillingForm `json:"startFillingForm,omitempty"`
}

// NewCustomizationConfigDto instantiates a new CustomizationConfigDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomizationConfigDto() *CustomizationConfigDto {
	this := CustomizationConfigDto{}
	return &this
}

// NewCustomizationConfigDtoWithDefaults instantiates a new CustomizationConfigDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomizationConfigDtoWithDefaults() *CustomizationConfigDto {
	this := CustomizationConfigDto{}
	return &this
}

// GetAbout returns the About field value if set, zero value otherwise.
func (o *CustomizationConfigDto) GetAbout() bool {
	if o == nil || IsNil(o.About) {
		var ret bool
		return ret
	}
	return *o.About
}

// GetAboutOk returns a tuple with the About field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomizationConfigDto) GetAboutOk() (*bool, bool) {
	if o == nil || IsNil(o.About) {
		return nil, false
	}
	return o.About, true
}

// HasAbout returns a boolean if a field has been set.
func (o *CustomizationConfigDto) IsAboutSet() bool {
	if o != nil && !IsNil(o.About) {
		return true
	}

	return false
}

// SetAbout gets a reference to the given bool and assigns it to the About field.
func (o *CustomizationConfigDto) SetAbout(v bool) {
	o.About = &v
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *CustomizationConfigDto) GetCustomer() CustomerConfigDto {
	if o == nil || IsNil(o.Customer) {
		var ret CustomerConfigDto
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomizationConfigDto) GetCustomerOk() (*CustomerConfigDto, bool) {
	if o == nil || IsNil(o.Customer) {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *CustomizationConfigDto) IsCustomerSet() bool {
	if o != nil && !IsNil(o.Customer) {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given CustomerConfigDto and assigns it to the Customer field.
func (o *CustomizationConfigDto) SetCustomer(v CustomerConfigDto) {
	o.Customer = &v
}

// GetAnonymous returns the Anonymous field value if set, zero value otherwise.
func (o *CustomizationConfigDto) GetAnonymous() AnonymousConfigDto {
	if o == nil || IsNil(o.Anonymous) {
		var ret AnonymousConfigDto
		return ret
	}
	return *o.Anonymous
}

// GetAnonymousOk returns a tuple with the Anonymous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomizationConfigDto) GetAnonymousOk() (*AnonymousConfigDto, bool) {
	if o == nil || IsNil(o.Anonymous) {
		return nil, false
	}
	return o.Anonymous, true
}

// HasAnonymous returns a boolean if a field has been set.
func (o *CustomizationConfigDto) IsAnonymousSet() bool {
	if o != nil && !IsNil(o.Anonymous) {
		return true
	}

	return false
}

// SetAnonymous gets a reference to the given AnonymousConfigDto and assigns it to the Anonymous field.
func (o *CustomizationConfigDto) SetAnonymous(v AnonymousConfigDto) {
	o.Anonymous = &v
}

// GetFeedback returns the Feedback field value if set, zero value otherwise.
func (o *CustomizationConfigDto) GetFeedback() FeedbackConfig {
	if o == nil || IsNil(o.Feedback) {
		var ret FeedbackConfig
		return ret
	}
	return *o.Feedback
}

// GetFeedbackOk returns a tuple with the Feedback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomizationConfigDto) GetFeedbackOk() (*FeedbackConfig, bool) {
	if o == nil || IsNil(o.Feedback) {
		return nil, false
	}
	return o.Feedback, true
}

// HasFeedback returns a boolean if a field has been set.
func (o *CustomizationConfigDto) IsFeedbackSet() bool {
	if o != nil && !IsNil(o.Feedback) {
		return true
	}

	return false
}

// SetFeedback gets a reference to the given FeedbackConfig and assigns it to the Feedback field.
func (o *CustomizationConfigDto) SetFeedback(v FeedbackConfig) {
	o.Feedback = &v
}

// GetForcesave returns the Forcesave field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomizationConfigDto) GetForcesave() bool {
	if o == nil || IsNil(o.Forcesave.Get()) {
		var ret bool
		return ret
	}
	return *o.Forcesave.Get()
}

// GetForcesaveOk returns a tuple with the Forcesave field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomizationConfigDto) GetForcesaveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Forcesave.Get(), o.Forcesave.IsSet()
}

// HasForcesave returns a boolean if a field has been set.
func (o *CustomizationConfigDto) IsForcesaveSet() bool {
	if o != nil && o.Forcesave.IsSet() {
		return true
	}

	return false
}

// SetForcesave gets a reference to the given NullableBool and assigns it to the Forcesave field.
func (o *CustomizationConfigDto) SetForcesave(v bool) {
	o.Forcesave.Set(&v)
}
// SetForcesaveNil sets the value for Forcesave to be an explicit nil
func (o *CustomizationConfigDto) SetForcesaveNil() {
	o.Forcesave.Set(nil)
}

// UnsetForcesave ensures that no value is present for Forcesave, not even an explicit nil
func (o *CustomizationConfigDto) UnsetForcesave() {
	o.Forcesave.Unset()
}

// GetGoback returns the Goback field value if set, zero value otherwise.
func (o *CustomizationConfigDto) GetGoback() GobackConfig {
	if o == nil || IsNil(o.Goback) {
		var ret GobackConfig
		return ret
	}
	return *o.Goback
}

// GetGobackOk returns a tuple with the Goback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomizationConfigDto) GetGobackOk() (*GobackConfig, bool) {
	if o == nil || IsNil(o.Goback) {
		return nil, false
	}
	return o.Goback, true
}

// HasGoback returns a boolean if a field has been set.
func (o *CustomizationConfigDto) IsGobackSet() bool {
	if o != nil && !IsNil(o.Goback) {
		return true
	}

	return false
}

// SetGoback gets a reference to the given GobackConfig and assigns it to the Goback field.
func (o *CustomizationConfigDto) SetGoback(v GobackConfig) {
	o.Goback = &v
}

// GetReview returns the Review field value if set, zero value otherwise.
func (o *CustomizationConfigDto) GetReview() ReviewConfig {
	if o == nil || IsNil(o.Review) {
		var ret ReviewConfig
		return ret
	}
	return *o.Review
}

// GetReviewOk returns a tuple with the Review field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomizationConfigDto) GetReviewOk() (*ReviewConfig, bool) {
	if o == nil || IsNil(o.Review) {
		return nil, false
	}
	return o.Review, true
}

// HasReview returns a boolean if a field has been set.
func (o *CustomizationConfigDto) IsReviewSet() bool {
	if o != nil && !IsNil(o.Review) {
		return true
	}

	return false
}

// SetReview gets a reference to the given ReviewConfig and assigns it to the Review field.
func (o *CustomizationConfigDto) SetReview(v ReviewConfig) {
	o.Review = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *CustomizationConfigDto) GetLogo() LogoConfigDto {
	if o == nil || IsNil(o.Logo) {
		var ret LogoConfigDto
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomizationConfigDto) GetLogoOk() (*LogoConfigDto, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *CustomizationConfigDto) IsLogoSet() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given LogoConfigDto and assigns it to the Logo field.
func (o *CustomizationConfigDto) SetLogo(v LogoConfigDto) {
	o.Logo = &v
}

// GetMentionShare returns the MentionShare field value if set, zero value otherwise.
func (o *CustomizationConfigDto) GetMentionShare() bool {
	if o == nil || IsNil(o.MentionShare) {
		var ret bool
		return ret
	}
	return *o.MentionShare
}

// GetMentionShareOk returns a tuple with the MentionShare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomizationConfigDto) GetMentionShareOk() (*bool, bool) {
	if o == nil || IsNil(o.MentionShare) {
		return nil, false
	}
	return o.MentionShare, true
}

// HasMentionShare returns a boolean if a field has been set.
func (o *CustomizationConfigDto) IsMentionShareSet() bool {
	if o != nil && !IsNil(o.MentionShare) {
		return true
	}

	return false
}

// SetMentionShare gets a reference to the given bool and assigns it to the MentionShare field.
func (o *CustomizationConfigDto) SetMentionShare(v bool) {
	o.MentionShare = &v
}

// GetSubmitForm returns the SubmitForm field value if set, zero value otherwise.
func (o *CustomizationConfigDto) GetSubmitForm() SubmitForm {
	if o == nil || IsNil(o.SubmitForm) {
		var ret SubmitForm
		return ret
	}
	return *o.SubmitForm
}

// GetSubmitFormOk returns a tuple with the SubmitForm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomizationConfigDto) GetSubmitFormOk() (*SubmitForm, bool) {
	if o == nil || IsNil(o.SubmitForm) {
		return nil, false
	}
	return o.SubmitForm, true
}

// HasSubmitForm returns a boolean if a field has been set.
func (o *CustomizationConfigDto) IsSubmitFormSet() bool {
	if o != nil && !IsNil(o.SubmitForm) {
		return true
	}

	return false
}

// SetSubmitForm gets a reference to the given SubmitForm and assigns it to the SubmitForm field.
func (o *CustomizationConfigDto) SetSubmitForm(v SubmitForm) {
	o.SubmitForm = &v
}

// GetStartFillingForm returns the StartFillingForm field value if set, zero value otherwise.
func (o *CustomizationConfigDto) GetStartFillingForm() StartFillingForm {
	if o == nil || IsNil(o.StartFillingForm) {
		var ret StartFillingForm
		return ret
	}
	return *o.StartFillingForm
}

// GetStartFillingFormOk returns a tuple with the StartFillingForm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomizationConfigDto) GetStartFillingFormOk() (*StartFillingForm, bool) {
	if o == nil || IsNil(o.StartFillingForm) {
		return nil, false
	}
	return o.StartFillingForm, true
}

// HasStartFillingForm returns a boolean if a field has been set.
func (o *CustomizationConfigDto) IsStartFillingFormSet() bool {
	if o != nil && !IsNil(o.StartFillingForm) {
		return true
	}

	return false
}

// SetStartFillingForm gets a reference to the given StartFillingForm and assigns it to the StartFillingForm field.
func (o *CustomizationConfigDto) SetStartFillingForm(v StartFillingForm) {
	o.StartFillingForm = &v
}

func (o CustomizationConfigDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomizationConfigDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.About) {
		toSerialize["about"] = o.About
	}
	if !IsNil(o.Customer) {
		toSerialize["customer"] = o.Customer
	}
	if !IsNil(o.Anonymous) {
		toSerialize["anonymous"] = o.Anonymous
	}
	if !IsNil(o.Feedback) {
		toSerialize["feedback"] = o.Feedback
	}
	if o.Forcesave.IsSet() {
		toSerialize["forcesave"] = o.Forcesave.Get()
	}
	if !IsNil(o.Goback) {
		toSerialize["goback"] = o.Goback
	}
	if !IsNil(o.Review) {
		toSerialize["review"] = o.Review
	}
	if !IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}
	if !IsNil(o.MentionShare) {
		toSerialize["mentionShare"] = o.MentionShare
	}
	if !IsNil(o.SubmitForm) {
		toSerialize["submitForm"] = o.SubmitForm
	}
	if !IsNil(o.StartFillingForm) {
		toSerialize["startFillingForm"] = o.StartFillingForm
	}
	return toSerialize, nil
}

type NullableCustomizationConfigDto struct {
	value *CustomizationConfigDto
	isSet bool
}

func (v NullableCustomizationConfigDto) Get() *CustomizationConfigDto {
	return v.value
}

func (v *NullableCustomizationConfigDto) Set(val *CustomizationConfigDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomizationConfigDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomizationConfigDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomizationConfigDto(val *CustomizationConfigDto) *NullableCustomizationConfigDto {
	return &NullableCustomizationConfigDto{value: val, isSet: true}
}

func (v NullableCustomizationConfigDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomizationConfigDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

