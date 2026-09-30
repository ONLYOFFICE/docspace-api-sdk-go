# TfaValidateRequestsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **NullableString** | The code to check - either one from the authenticator application or one of the account's unused backup  codes, which is spent by the check. A wrong code is refused with 400 and counts against the portal login  attempt limit. | 
**Session** | Pointer to **bool** | Whether the sign-in that follows is tied to the browser session. When it is, the session ends with the  browser rather than lasting for the portal session lifetime. | [optional] 

## Methods

### NewTfaValidateRequestsDto

`func NewTfaValidateRequestsDto(code NullableString, ) *TfaValidateRequestsDto`

NewTfaValidateRequestsDto instantiates a new TfaValidateRequestsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTfaValidateRequestsDtoWithDefaults

`func NewTfaValidateRequestsDtoWithDefaults() *TfaValidateRequestsDto`

NewTfaValidateRequestsDtoWithDefaults instantiates a new TfaValidateRequestsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *TfaValidateRequestsDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TfaValidateRequestsDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TfaValidateRequestsDto) SetCode(v string)`

SetCode sets Code field to given value.


### SetCodeNil

`func (o *TfaValidateRequestsDto) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *TfaValidateRequestsDto) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetSession

`func (o *TfaValidateRequestsDto) GetSession() bool`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *TfaValidateRequestsDto) GetSessionOk() (*bool, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *TfaValidateRequestsDto) SetSession(v bool)`

SetSession sets Session field to given value.

### HasSession

`func (o *TfaValidateRequestsDto) HasSession() bool`

HasSession returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


