# EmailMemberRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The address to send the password recovery link to. It is required and validated even by  `POST api/2.0/people/guests/share/approve`, which then ignores its value and takes the account from the  confirmation token instead. | 
**RecaptchaType** | Pointer to [**RecaptchaType**](RecaptchaType.md) | Which CAPTCHA the `recaptchaResponse` comes from: `Default` for the web reCAPTCHA, `AndroidV2` or `iOSV2` for  the mobile ones, and `hCaptcha` when the portal is configured with hCaptcha. It matters only for an  unauthenticated request on a portal that has a CAPTCHA. | [optional] 
**RecaptchaResponse** | Pointer to **NullableString** | The user's response to the CAPTCHA challenge. | [optional] 

## Methods

### NewEmailMemberRequestDto

`func NewEmailMemberRequestDto(email string, ) *EmailMemberRequestDto`

NewEmailMemberRequestDto instantiates a new EmailMemberRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailMemberRequestDtoWithDefaults

`func NewEmailMemberRequestDtoWithDefaults() *EmailMemberRequestDto`

NewEmailMemberRequestDtoWithDefaults instantiates a new EmailMemberRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *EmailMemberRequestDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EmailMemberRequestDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EmailMemberRequestDto) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRecaptchaType

`func (o *EmailMemberRequestDto) GetRecaptchaType() RecaptchaType`

GetRecaptchaType returns the RecaptchaType field if non-nil, zero value otherwise.

### GetRecaptchaTypeOk

`func (o *EmailMemberRequestDto) GetRecaptchaTypeOk() (*RecaptchaType, bool)`

GetRecaptchaTypeOk returns a tuple with the RecaptchaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecaptchaType

`func (o *EmailMemberRequestDto) SetRecaptchaType(v RecaptchaType)`

SetRecaptchaType sets RecaptchaType field to given value.

### HasRecaptchaType

`func (o *EmailMemberRequestDto) HasRecaptchaType() bool`

HasRecaptchaType returns a boolean if a field has been set.

### GetRecaptchaResponse

`func (o *EmailMemberRequestDto) GetRecaptchaResponse() string`

GetRecaptchaResponse returns the RecaptchaResponse field if non-nil, zero value otherwise.

### GetRecaptchaResponseOk

`func (o *EmailMemberRequestDto) GetRecaptchaResponseOk() (*string, bool)`

GetRecaptchaResponseOk returns a tuple with the RecaptchaResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecaptchaResponse

`func (o *EmailMemberRequestDto) SetRecaptchaResponse(v string)`

SetRecaptchaResponse sets RecaptchaResponse field to given value.

### HasRecaptchaResponse

`func (o *EmailMemberRequestDto) HasRecaptchaResponse() bool`

HasRecaptchaResponse returns a boolean if a field has been set.

### SetRecaptchaResponseNil

`func (o *EmailMemberRequestDto) SetRecaptchaResponseNil(b bool)`

 SetRecaptchaResponseNil sets the value for RecaptchaResponse to be an explicit nil

### UnsetRecaptchaResponse
`func (o *EmailMemberRequestDto) UnsetRecaptchaResponse()`

UnsetRecaptchaResponse ensures that no value is present for RecaptchaResponse, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


