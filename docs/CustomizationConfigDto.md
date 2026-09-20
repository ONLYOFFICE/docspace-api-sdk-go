# CustomizationConfigDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**About** | Pointer to **bool** | Whether the About entry of the editor menu is shown. | [optional] 
**Customer** | Pointer to [**CustomerConfigDto**](CustomerConfigDto.md) | The branding of the organization running the portal. It is filled in on a server installation only and is  empty in the cloud. | [optional] 
**Anonymous** | Pointer to [**AnonymousConfigDto**](AnonymousConfigDto.md) | How an anonymous participant is treated in this session. | [optional] 
**Feedback** | Pointer to [**FeedbackConfig**](FeedbackConfig.md) | The support link the editor offers behind its feedback button. | [optional] 
**Forcesave** | Pointer to **NullableBool** | Whether the editors write intermediate revisions while the document stays open. It is empty when the portal  leaves the decision to the editors themselves. | [optional] 
**Goback** | Pointer to [**GobackConfig**](GobackConfig.md) | Where the editor returns the user to when they leave the document. It is empty when there is nowhere to go  back to, as in an embedded opening. | [optional] 
**Review** | Pointer to [**ReviewConfig**](ReviewConfig.md) | How tracked changes are displayed when the document opens; it depends on whether this session may write. | [optional] 
**Logo** | Pointer to [**LogoConfigDto**](LogoConfigDto.md) | The logo the editor shows, in the variants the current layout and file type need. | [optional] 
**MentionShare** | Pointer to **bool** | Whether mentioning a user who cannot yet open the document offers to share it with them, instead of silently  notifying nobody. | [optional] 
**SubmitForm** | Pointer to [**SubmitForm**](SubmitForm.md) | The submit button of a form: whether it is shown and what it says. | [optional] 
**StartFillingForm** | Pointer to [**StartFillingForm**](StartFillingForm.md) | The button that starts filling out the form. It is empty when this opening offers no such button. | [optional] 
**Ai** | Pointer to [**AIConfig**](AIConfig.md) | The AI configuration settings. | [optional] 

## Methods

### NewCustomizationConfigDto

`func NewCustomizationConfigDto() *CustomizationConfigDto`

NewCustomizationConfigDto instantiates a new CustomizationConfigDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomizationConfigDtoWithDefaults

`func NewCustomizationConfigDtoWithDefaults() *CustomizationConfigDto`

NewCustomizationConfigDtoWithDefaults instantiates a new CustomizationConfigDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbout

`func (o *CustomizationConfigDto) GetAbout() bool`

GetAbout returns the About field if non-nil, zero value otherwise.

### GetAboutOk

`func (o *CustomizationConfigDto) GetAboutOk() (*bool, bool)`

GetAboutOk returns a tuple with the About field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbout

`func (o *CustomizationConfigDto) SetAbout(v bool)`

SetAbout sets About field to given value.

### HasAbout

`func (o *CustomizationConfigDto) HasAbout() bool`

HasAbout returns a boolean if a field has been set.

### GetCustomer

`func (o *CustomizationConfigDto) GetCustomer() CustomerConfigDto`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *CustomizationConfigDto) GetCustomerOk() (*CustomerConfigDto, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *CustomizationConfigDto) SetCustomer(v CustomerConfigDto)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *CustomizationConfigDto) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetAnonymous

`func (o *CustomizationConfigDto) GetAnonymous() AnonymousConfigDto`

GetAnonymous returns the Anonymous field if non-nil, zero value otherwise.

### GetAnonymousOk

`func (o *CustomizationConfigDto) GetAnonymousOk() (*AnonymousConfigDto, bool)`

GetAnonymousOk returns a tuple with the Anonymous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymous

`func (o *CustomizationConfigDto) SetAnonymous(v AnonymousConfigDto)`

SetAnonymous sets Anonymous field to given value.

### HasAnonymous

`func (o *CustomizationConfigDto) HasAnonymous() bool`

HasAnonymous returns a boolean if a field has been set.

### GetFeedback

`func (o *CustomizationConfigDto) GetFeedback() FeedbackConfig`

GetFeedback returns the Feedback field if non-nil, zero value otherwise.

### GetFeedbackOk

`func (o *CustomizationConfigDto) GetFeedbackOk() (*FeedbackConfig, bool)`

GetFeedbackOk returns a tuple with the Feedback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedback

`func (o *CustomizationConfigDto) SetFeedback(v FeedbackConfig)`

SetFeedback sets Feedback field to given value.

### HasFeedback

`func (o *CustomizationConfigDto) HasFeedback() bool`

HasFeedback returns a boolean if a field has been set.

### GetForcesave

`func (o *CustomizationConfigDto) GetForcesave() bool`

GetForcesave returns the Forcesave field if non-nil, zero value otherwise.

### GetForcesaveOk

`func (o *CustomizationConfigDto) GetForcesaveOk() (*bool, bool)`

GetForcesaveOk returns a tuple with the Forcesave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcesave

`func (o *CustomizationConfigDto) SetForcesave(v bool)`

SetForcesave sets Forcesave field to given value.

### HasForcesave

`func (o *CustomizationConfigDto) HasForcesave() bool`

HasForcesave returns a boolean if a field has been set.

### SetForcesaveNil

`func (o *CustomizationConfigDto) SetForcesaveNil(b bool)`

 SetForcesaveNil sets the value for Forcesave to be an explicit nil

### UnsetForcesave
`func (o *CustomizationConfigDto) UnsetForcesave()`

UnsetForcesave ensures that no value is present for Forcesave, not even an explicit nil
### GetGoback

`func (o *CustomizationConfigDto) GetGoback() GobackConfig`

GetGoback returns the Goback field if non-nil, zero value otherwise.

### GetGobackOk

`func (o *CustomizationConfigDto) GetGobackOk() (*GobackConfig, bool)`

GetGobackOk returns a tuple with the Goback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoback

`func (o *CustomizationConfigDto) SetGoback(v GobackConfig)`

SetGoback sets Goback field to given value.

### HasGoback

`func (o *CustomizationConfigDto) HasGoback() bool`

HasGoback returns a boolean if a field has been set.

### GetReview

`func (o *CustomizationConfigDto) GetReview() ReviewConfig`

GetReview returns the Review field if non-nil, zero value otherwise.

### GetReviewOk

`func (o *CustomizationConfigDto) GetReviewOk() (*ReviewConfig, bool)`

GetReviewOk returns a tuple with the Review field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReview

`func (o *CustomizationConfigDto) SetReview(v ReviewConfig)`

SetReview sets Review field to given value.

### HasReview

`func (o *CustomizationConfigDto) HasReview() bool`

HasReview returns a boolean if a field has been set.

### GetLogo

`func (o *CustomizationConfigDto) GetLogo() LogoConfigDto`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *CustomizationConfigDto) GetLogoOk() (*LogoConfigDto, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *CustomizationConfigDto) SetLogo(v LogoConfigDto)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *CustomizationConfigDto) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetMentionShare

`func (o *CustomizationConfigDto) GetMentionShare() bool`

GetMentionShare returns the MentionShare field if non-nil, zero value otherwise.

### GetMentionShareOk

`func (o *CustomizationConfigDto) GetMentionShareOk() (*bool, bool)`

GetMentionShareOk returns a tuple with the MentionShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionShare

`func (o *CustomizationConfigDto) SetMentionShare(v bool)`

SetMentionShare sets MentionShare field to given value.

### HasMentionShare

`func (o *CustomizationConfigDto) HasMentionShare() bool`

HasMentionShare returns a boolean if a field has been set.

### GetSubmitForm

`func (o *CustomizationConfigDto) GetSubmitForm() SubmitForm`

GetSubmitForm returns the SubmitForm field if non-nil, zero value otherwise.

### GetSubmitFormOk

`func (o *CustomizationConfigDto) GetSubmitFormOk() (*SubmitForm, bool)`

GetSubmitFormOk returns a tuple with the SubmitForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitForm

`func (o *CustomizationConfigDto) SetSubmitForm(v SubmitForm)`

SetSubmitForm sets SubmitForm field to given value.

### HasSubmitForm

`func (o *CustomizationConfigDto) HasSubmitForm() bool`

HasSubmitForm returns a boolean if a field has been set.

### GetStartFillingForm

`func (o *CustomizationConfigDto) GetStartFillingForm() StartFillingForm`

GetStartFillingForm returns the StartFillingForm field if non-nil, zero value otherwise.

### GetStartFillingFormOk

`func (o *CustomizationConfigDto) GetStartFillingFormOk() (*StartFillingForm, bool)`

GetStartFillingFormOk returns a tuple with the StartFillingForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartFillingForm

`func (o *CustomizationConfigDto) SetStartFillingForm(v StartFillingForm)`

SetStartFillingForm sets StartFillingForm field to given value.

### HasStartFillingForm

`func (o *CustomizationConfigDto) HasStartFillingForm() bool`

HasStartFillingForm returns a boolean if a field has been set.

### GetAi

`func (o *CustomizationConfigDto) GetAi() AIConfig`

GetAi returns the Ai field if non-nil, zero value otherwise.

### GetAiOk

`func (o *CustomizationConfigDto) GetAiOk() (*AIConfig, bool)`

GetAiOk returns a tuple with the Ai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAi

`func (o *CustomizationConfigDto) SetAi(v AIConfig)`

SetAi sets Ai field to given value.

### HasAi

`func (o *CustomizationConfigDto) HasAi() bool`

HasAi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


