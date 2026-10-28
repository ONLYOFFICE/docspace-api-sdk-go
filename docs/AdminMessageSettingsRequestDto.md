# AdminMessageSettingsRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **NullableString** | What the sender wants to tell the portal administrators. Markup is stripped before the letter is written, so  a body that carries nothing but markup counts as empty and is refused with 400. | 
**Email** | **NullableString** | The address the sender can be answered at, which the letter is signed with. It has to be a well-formed email  address. | 
**Culture** | Pointer to **NullableString** | The language the letter is written in, as a culture name such as `en-US`. A culture the installation does not  have falls back to the portal language rather than failing the call. | [optional] 
**RecaptchaType** | Pointer to [**RecaptchaType**](RecaptchaType.md) | Which CAPTCHA service the proof in `recaptchaResponse` came from. It has to match the service the  installation is configured with, which `GET api/2.0/capabilities` reports; the default value means the  installation is left to decide. | [optional] 
**RecaptchaResponse** | Pointer to **NullableString** | The token the CAPTCHA widget produced in the browser, passed on unchanged for the portal to verify with the  CAPTCHA service. It is single-use and short-lived, so it cannot be reused for a second message. | [optional] 

## Methods

### NewAdminMessageSettingsRequestDto

`func NewAdminMessageSettingsRequestDto(message NullableString, email NullableString, ) *AdminMessageSettingsRequestDto`

NewAdminMessageSettingsRequestDto instantiates a new AdminMessageSettingsRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminMessageSettingsRequestDtoWithDefaults

`func NewAdminMessageSettingsRequestDtoWithDefaults() *AdminMessageSettingsRequestDto`

NewAdminMessageSettingsRequestDtoWithDefaults instantiates a new AdminMessageSettingsRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *AdminMessageSettingsRequestDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AdminMessageSettingsRequestDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AdminMessageSettingsRequestDto) SetMessage(v string)`

SetMessage sets Message field to given value.


### SetMessageNil

`func (o *AdminMessageSettingsRequestDto) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *AdminMessageSettingsRequestDto) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetEmail

`func (o *AdminMessageSettingsRequestDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AdminMessageSettingsRequestDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AdminMessageSettingsRequestDto) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *AdminMessageSettingsRequestDto) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *AdminMessageSettingsRequestDto) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetCulture

`func (o *AdminMessageSettingsRequestDto) GetCulture() string`

GetCulture returns the Culture field if non-nil, zero value otherwise.

### GetCultureOk

`func (o *AdminMessageSettingsRequestDto) GetCultureOk() (*string, bool)`

GetCultureOk returns a tuple with the Culture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCulture

`func (o *AdminMessageSettingsRequestDto) SetCulture(v string)`

SetCulture sets Culture field to given value.

### HasCulture

`func (o *AdminMessageSettingsRequestDto) HasCulture() bool`

HasCulture returns a boolean if a field has been set.

### SetCultureNil

`func (o *AdminMessageSettingsRequestDto) SetCultureNil(b bool)`

 SetCultureNil sets the value for Culture to be an explicit nil

### UnsetCulture
`func (o *AdminMessageSettingsRequestDto) UnsetCulture()`

UnsetCulture ensures that no value is present for Culture, not even an explicit nil
### GetRecaptchaType

`func (o *AdminMessageSettingsRequestDto) GetRecaptchaType() RecaptchaType`

GetRecaptchaType returns the RecaptchaType field if non-nil, zero value otherwise.

### GetRecaptchaTypeOk

`func (o *AdminMessageSettingsRequestDto) GetRecaptchaTypeOk() (*RecaptchaType, bool)`

GetRecaptchaTypeOk returns a tuple with the RecaptchaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecaptchaType

`func (o *AdminMessageSettingsRequestDto) SetRecaptchaType(v RecaptchaType)`

SetRecaptchaType sets RecaptchaType field to given value.

### HasRecaptchaType

`func (o *AdminMessageSettingsRequestDto) HasRecaptchaType() bool`

HasRecaptchaType returns a boolean if a field has been set.

### GetRecaptchaResponse

`func (o *AdminMessageSettingsRequestDto) GetRecaptchaResponse() string`

GetRecaptchaResponse returns the RecaptchaResponse field if non-nil, zero value otherwise.

### GetRecaptchaResponseOk

`func (o *AdminMessageSettingsRequestDto) GetRecaptchaResponseOk() (*string, bool)`

GetRecaptchaResponseOk returns a tuple with the RecaptchaResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecaptchaResponse

`func (o *AdminMessageSettingsRequestDto) SetRecaptchaResponse(v string)`

SetRecaptchaResponse sets RecaptchaResponse field to given value.

### HasRecaptchaResponse

`func (o *AdminMessageSettingsRequestDto) HasRecaptchaResponse() bool`

HasRecaptchaResponse returns a boolean if a field has been set.

### SetRecaptchaResponseNil

`func (o *AdminMessageSettingsRequestDto) SetRecaptchaResponseNil(b bool)`

 SetRecaptchaResponseNil sets the value for RecaptchaResponse to be an explicit nil

### UnsetRecaptchaResponse
`func (o *AdminMessageSettingsRequestDto) UnsetRecaptchaResponse()`

UnsetRecaptchaResponse ensures that no value is present for RecaptchaResponse, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


