# TfaConfirmDataDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **NullableString** | The link to open. Its `type` shows which step it is: phone activation or phone authorization for the SMS  method, and authenticator activation or re-verification for the application method. The whole body is empty  when the portal requires no second factor of the caller. | [optional] 
**CookieName** | Pointer to **NullableString** | The name of the confirmation cookie the link is validated against. It is filled in only for the  authenticator-application method; the SMS method returns `url` alone. | [optional] 
**CookieValue** | Pointer to **NullableString** | The value of that cookie. The call already set it on the response, so it is repeated here only for a client  that does not keep cookies of its own; it is filled in under the same condition as `cookieName`, and a  later call to this operation replaces it. | [optional] 

## Methods

### NewTfaConfirmDataDto

`func NewTfaConfirmDataDto() *TfaConfirmDataDto`

NewTfaConfirmDataDto instantiates a new TfaConfirmDataDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTfaConfirmDataDtoWithDefaults

`func NewTfaConfirmDataDtoWithDefaults() *TfaConfirmDataDto`

NewTfaConfirmDataDtoWithDefaults instantiates a new TfaConfirmDataDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *TfaConfirmDataDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TfaConfirmDataDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TfaConfirmDataDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TfaConfirmDataDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *TfaConfirmDataDto) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *TfaConfirmDataDto) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetCookieName

`func (o *TfaConfirmDataDto) GetCookieName() string`

GetCookieName returns the CookieName field if non-nil, zero value otherwise.

### GetCookieNameOk

`func (o *TfaConfirmDataDto) GetCookieNameOk() (*string, bool)`

GetCookieNameOk returns a tuple with the CookieName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieName

`func (o *TfaConfirmDataDto) SetCookieName(v string)`

SetCookieName sets CookieName field to given value.

### HasCookieName

`func (o *TfaConfirmDataDto) HasCookieName() bool`

HasCookieName returns a boolean if a field has been set.

### SetCookieNameNil

`func (o *TfaConfirmDataDto) SetCookieNameNil(b bool)`

 SetCookieNameNil sets the value for CookieName to be an explicit nil

### UnsetCookieName
`func (o *TfaConfirmDataDto) UnsetCookieName()`

UnsetCookieName ensures that no value is present for CookieName, not even an explicit nil
### GetCookieValue

`func (o *TfaConfirmDataDto) GetCookieValue() string`

GetCookieValue returns the CookieValue field if non-nil, zero value otherwise.

### GetCookieValueOk

`func (o *TfaConfirmDataDto) GetCookieValueOk() (*string, bool)`

GetCookieValueOk returns a tuple with the CookieValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieValue

`func (o *TfaConfirmDataDto) SetCookieValue(v string)`

SetCookieValue sets CookieValue field to given value.

### HasCookieValue

`func (o *TfaConfirmDataDto) HasCookieValue() bool`

HasCookieValue returns a boolean if a field has been set.

### SetCookieValueNil

`func (o *TfaConfirmDataDto) SetCookieValueNil(b bool)`

 SetCookieValueNil sets the value for CookieValue to be an explicit nil

### UnsetCookieValue
`func (o *TfaConfirmDataDto) UnsetCookieValue()`

UnsetCookieValue ensures that no value is present for CookieValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


