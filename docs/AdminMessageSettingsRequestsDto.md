# AdminMessageSettingsRequestsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **NullableString** | The content of the administrator message to be sent. | 
**Email** | **NullableString** | Email | 
**Culture** | Pointer to **NullableString** | Culture | [optional] 
**RecaptchaType** | Pointer to [**RecaptchaType**](RecaptchaType.md) | The type of CAPTCHA validation used. | [optional] 
**RecaptchaResponse** | Pointer to **NullableString** | The user's response to the CAPTCHA challenge. | [optional] 

## Methods

### NewAdminMessageSettingsRequestsDto

`func NewAdminMessageSettingsRequestsDto(message NullableString, email NullableString, ) *AdminMessageSettingsRequestsDto`

NewAdminMessageSettingsRequestsDto instantiates a new AdminMessageSettingsRequestsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminMessageSettingsRequestsDtoWithDefaults

`func NewAdminMessageSettingsRequestsDtoWithDefaults() *AdminMessageSettingsRequestsDto`

NewAdminMessageSettingsRequestsDtoWithDefaults instantiates a new AdminMessageSettingsRequestsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *AdminMessageSettingsRequestsDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AdminMessageSettingsRequestsDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AdminMessageSettingsRequestsDto) SetMessage(v string)`

SetMessage sets Message field to given value.


### SetMessageNil

`func (o *AdminMessageSettingsRequestsDto) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *AdminMessageSettingsRequestsDto) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetEmail

`func (o *AdminMessageSettingsRequestsDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AdminMessageSettingsRequestsDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AdminMessageSettingsRequestsDto) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *AdminMessageSettingsRequestsDto) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *AdminMessageSettingsRequestsDto) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetCulture

`func (o *AdminMessageSettingsRequestsDto) GetCulture() string`

GetCulture returns the Culture field if non-nil, zero value otherwise.

### GetCultureOk

`func (o *AdminMessageSettingsRequestsDto) GetCultureOk() (*string, bool)`

GetCultureOk returns a tuple with the Culture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCulture

`func (o *AdminMessageSettingsRequestsDto) SetCulture(v string)`

SetCulture sets Culture field to given value.

### HasCulture

`func (o *AdminMessageSettingsRequestsDto) HasCulture() bool`

HasCulture returns a boolean if a field has been set.

### SetCultureNil

`func (o *AdminMessageSettingsRequestsDto) SetCultureNil(b bool)`

 SetCultureNil sets the value for Culture to be an explicit nil

### UnsetCulture
`func (o *AdminMessageSettingsRequestsDto) UnsetCulture()`

UnsetCulture ensures that no value is present for Culture, not even an explicit nil
### GetRecaptchaType

`func (o *AdminMessageSettingsRequestsDto) GetRecaptchaType() RecaptchaType`

GetRecaptchaType returns the RecaptchaType field if non-nil, zero value otherwise.

### GetRecaptchaTypeOk

`func (o *AdminMessageSettingsRequestsDto) GetRecaptchaTypeOk() (*RecaptchaType, bool)`

GetRecaptchaTypeOk returns a tuple with the RecaptchaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecaptchaType

`func (o *AdminMessageSettingsRequestsDto) SetRecaptchaType(v RecaptchaType)`

SetRecaptchaType sets RecaptchaType field to given value.

### HasRecaptchaType

`func (o *AdminMessageSettingsRequestsDto) HasRecaptchaType() bool`

HasRecaptchaType returns a boolean if a field has been set.

### GetRecaptchaResponse

`func (o *AdminMessageSettingsRequestsDto) GetRecaptchaResponse() string`

GetRecaptchaResponse returns the RecaptchaResponse field if non-nil, zero value otherwise.

### GetRecaptchaResponseOk

`func (o *AdminMessageSettingsRequestsDto) GetRecaptchaResponseOk() (*string, bool)`

GetRecaptchaResponseOk returns a tuple with the RecaptchaResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecaptchaResponse

`func (o *AdminMessageSettingsRequestsDto) SetRecaptchaResponse(v string)`

SetRecaptchaResponse sets RecaptchaResponse field to given value.

### HasRecaptchaResponse

`func (o *AdminMessageSettingsRequestsDto) HasRecaptchaResponse() bool`

HasRecaptchaResponse returns a boolean if a field has been set.

### SetRecaptchaResponseNil

`func (o *AdminMessageSettingsRequestsDto) SetRecaptchaResponseNil(b bool)`

 SetRecaptchaResponseNil sets the value for RecaptchaResponse to be an explicit nil

### UnsetRecaptchaResponse
`func (o *AdminMessageSettingsRequestsDto) UnsetRecaptchaResponse()`

UnsetRecaptchaResponse ensures that no value is present for RecaptchaResponse, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


