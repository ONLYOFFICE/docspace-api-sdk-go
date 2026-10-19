# TfaValidateRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **NullableString** | The code to check - either one from the authenticator application or one of the account's unused backup  codes, which is spent by the check. A wrong code is refused with 400 and counts against the portal login  attempt limit. | 
**Session** | Pointer to **bool** | Whether the sign-in that follows is tied to the browser session. When it is, the session ends with the  browser rather than lasting for the portal session lifetime. | [optional] 

## Methods

### NewTfaValidateRequestDto

`func NewTfaValidateRequestDto(code NullableString, ) *TfaValidateRequestDto`

NewTfaValidateRequestDto instantiates a new TfaValidateRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTfaValidateRequestDtoWithDefaults

`func NewTfaValidateRequestDtoWithDefaults() *TfaValidateRequestDto`

NewTfaValidateRequestDtoWithDefaults instantiates a new TfaValidateRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *TfaValidateRequestDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TfaValidateRequestDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TfaValidateRequestDto) SetCode(v string)`

SetCode sets Code field to given value.


### SetCodeNil

`func (o *TfaValidateRequestDto) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *TfaValidateRequestDto) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetSession

`func (o *TfaValidateRequestDto) GetSession() bool`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *TfaValidateRequestDto) GetSessionOk() (*bool, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *TfaValidateRequestDto) SetSession(v bool)`

SetSession sets Session field to given value.

### HasSession

`func (o *TfaValidateRequestDto) HasSession() bool`

HasSession returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


