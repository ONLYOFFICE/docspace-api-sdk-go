# ConnectServerRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **NullableString** | OAuth authorization code received from the provider's redirect. Used to exchange for access and refresh tokens. | 

## Methods

### NewConnectServerRequestBody

`func NewConnectServerRequestBody(code NullableString, ) *ConnectServerRequestBody`

NewConnectServerRequestBody instantiates a new ConnectServerRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectServerRequestBodyWithDefaults

`func NewConnectServerRequestBodyWithDefaults() *ConnectServerRequestBody`

NewConnectServerRequestBodyWithDefaults instantiates a new ConnectServerRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ConnectServerRequestBody) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ConnectServerRequestBody) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ConnectServerRequestBody) SetCode(v string)`

SetCode sets Code field to given value.


### SetCodeNil

`func (o *ConnectServerRequestBody) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *ConnectServerRequestBody) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


