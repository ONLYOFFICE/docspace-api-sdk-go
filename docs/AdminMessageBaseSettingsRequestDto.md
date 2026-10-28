# AdminMessageBaseSettingsRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **NullableString** | The address the join link is sent to. It has to be a well-formed ASCII address rather than an  internationalized one, must not already belong to a member of the portal, and, where the portal trusts named  domains only, has to end with one of them; any of these faults is refused with 400. | 
**Culture** | Pointer to **NullableString** | The language the letter is written in, as a culture name such as `en-US`. A culture the installation does not  have falls back to the portal language rather than failing the call. | [optional] 

## Methods

### NewAdminMessageBaseSettingsRequestDto

`func NewAdminMessageBaseSettingsRequestDto(email NullableString, ) *AdminMessageBaseSettingsRequestDto`

NewAdminMessageBaseSettingsRequestDto instantiates a new AdminMessageBaseSettingsRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminMessageBaseSettingsRequestDtoWithDefaults

`func NewAdminMessageBaseSettingsRequestDtoWithDefaults() *AdminMessageBaseSettingsRequestDto`

NewAdminMessageBaseSettingsRequestDtoWithDefaults instantiates a new AdminMessageBaseSettingsRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *AdminMessageBaseSettingsRequestDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AdminMessageBaseSettingsRequestDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AdminMessageBaseSettingsRequestDto) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *AdminMessageBaseSettingsRequestDto) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *AdminMessageBaseSettingsRequestDto) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetCulture

`func (o *AdminMessageBaseSettingsRequestDto) GetCulture() string`

GetCulture returns the Culture field if non-nil, zero value otherwise.

### GetCultureOk

`func (o *AdminMessageBaseSettingsRequestDto) GetCultureOk() (*string, bool)`

GetCultureOk returns a tuple with the Culture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCulture

`func (o *AdminMessageBaseSettingsRequestDto) SetCulture(v string)`

SetCulture sets Culture field to given value.

### HasCulture

`func (o *AdminMessageBaseSettingsRequestDto) HasCulture() bool`

HasCulture returns a boolean if a field has been set.

### SetCultureNil

`func (o *AdminMessageBaseSettingsRequestDto) SetCultureNil(b bool)`

 SetCultureNil sets the value for Culture to be an explicit nil

### UnsetCulture
`func (o *AdminMessageBaseSettingsRequestDto) UnsetCulture()`

UnsetCulture ensures that no value is present for Culture, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


