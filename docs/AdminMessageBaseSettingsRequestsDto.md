# AdminMessageBaseSettingsRequestsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **NullableString** | The email address used for sending administrator messages. | 
**Culture** | Pointer to **NullableString** | The locale identifier for message localization. | [optional] 

## Methods

### NewAdminMessageBaseSettingsRequestsDto

`func NewAdminMessageBaseSettingsRequestsDto(email NullableString, ) *AdminMessageBaseSettingsRequestsDto`

NewAdminMessageBaseSettingsRequestsDto instantiates a new AdminMessageBaseSettingsRequestsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminMessageBaseSettingsRequestsDtoWithDefaults

`func NewAdminMessageBaseSettingsRequestsDtoWithDefaults() *AdminMessageBaseSettingsRequestsDto`

NewAdminMessageBaseSettingsRequestsDtoWithDefaults instantiates a new AdminMessageBaseSettingsRequestsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *AdminMessageBaseSettingsRequestsDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AdminMessageBaseSettingsRequestsDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AdminMessageBaseSettingsRequestsDto) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *AdminMessageBaseSettingsRequestsDto) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *AdminMessageBaseSettingsRequestsDto) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetCulture

`func (o *AdminMessageBaseSettingsRequestsDto) GetCulture() string`

GetCulture returns the Culture field if non-nil, zero value otherwise.

### GetCultureOk

`func (o *AdminMessageBaseSettingsRequestsDto) GetCultureOk() (*string, bool)`

GetCultureOk returns a tuple with the Culture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCulture

`func (o *AdminMessageBaseSettingsRequestsDto) SetCulture(v string)`

SetCulture sets Culture field to given value.

### HasCulture

`func (o *AdminMessageBaseSettingsRequestsDto) HasCulture() bool`

HasCulture returns a boolean if a field has been set.

### SetCultureNil

`func (o *AdminMessageBaseSettingsRequestsDto) SetCultureNil(b bool)`

 SetCultureNil sets the value for Culture to be an explicit nil

### UnsetCulture
`func (o *AdminMessageBaseSettingsRequestsDto) UnsetCulture()`

UnsetCulture ensures that no value is present for Culture, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


