# ErrorApiResponseError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | The human-readable error message. | [optional] 
**Type** | Pointer to **string** | The .NET type of the underlying exception. Only sent when stack traces are enabled. | [optional] 
**Stack** | Pointer to **string** | The stack trace of the underlying exception. Only sent when stack traces are enabled. | [optional] 
**Hresult** | Pointer to **int32** | The HRESULT of the underlying exception. Only sent when stack traces are enabled. | [optional] 

## Methods

### NewErrorApiResponseError

`func NewErrorApiResponseError() *ErrorApiResponseError`

NewErrorApiResponseError instantiates a new ErrorApiResponseError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorApiResponseErrorWithDefaults

`func NewErrorApiResponseErrorWithDefaults() *ErrorApiResponseError`

NewErrorApiResponseErrorWithDefaults instantiates a new ErrorApiResponseError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ErrorApiResponseError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorApiResponseError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorApiResponseError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ErrorApiResponseError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetType

`func (o *ErrorApiResponseError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrorApiResponseError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrorApiResponseError) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ErrorApiResponseError) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStack

`func (o *ErrorApiResponseError) GetStack() string`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *ErrorApiResponseError) GetStackOk() (*string, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *ErrorApiResponseError) SetStack(v string)`

SetStack sets Stack field to given value.

### HasStack

`func (o *ErrorApiResponseError) HasStack() bool`

HasStack returns a boolean if a field has been set.

### GetHresult

`func (o *ErrorApiResponseError) GetHresult() int32`

GetHresult returns the Hresult field if non-nil, zero value otherwise.

### GetHresultOk

`func (o *ErrorApiResponseError) GetHresultOk() (*int32, bool)`

GetHresultOk returns a tuple with the Hresult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHresult

`func (o *ErrorApiResponseError) SetHresult(v int32)`

SetHresult sets Hresult field to given value.

### HasHresult

`func (o *ErrorApiResponseError) HasHresult() bool`

HasHresult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


