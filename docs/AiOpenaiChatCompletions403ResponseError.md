# AiOpenaiChatCompletions403ResponseError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Human-readable description of the failure. | 
**Type** | **string** | OpenAI error class, for example `invalid_request_error`. | 
**Code** | Pointer to **NullableString** | Machine-readable code, when the provider supplies one. | [optional] 
**Param** | Pointer to **NullableString** | The request parameter at fault, when the failure names one. | [optional] 

## Methods

### NewAiOpenaiChatCompletions403ResponseError

`func NewAiOpenaiChatCompletions403ResponseError(message string, type_ string, ) *AiOpenaiChatCompletions403ResponseError`

NewAiOpenaiChatCompletions403ResponseError instantiates a new AiOpenaiChatCompletions403ResponseError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiOpenaiChatCompletions403ResponseErrorWithDefaults

`func NewAiOpenaiChatCompletions403ResponseErrorWithDefaults() *AiOpenaiChatCompletions403ResponseError`

NewAiOpenaiChatCompletions403ResponseErrorWithDefaults instantiates a new AiOpenaiChatCompletions403ResponseError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *AiOpenaiChatCompletions403ResponseError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AiOpenaiChatCompletions403ResponseError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AiOpenaiChatCompletions403ResponseError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetType

`func (o *AiOpenaiChatCompletions403ResponseError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AiOpenaiChatCompletions403ResponseError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AiOpenaiChatCompletions403ResponseError) SetType(v string)`

SetType sets Type field to given value.


### GetCode

`func (o *AiOpenaiChatCompletions403ResponseError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AiOpenaiChatCompletions403ResponseError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AiOpenaiChatCompletions403ResponseError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AiOpenaiChatCompletions403ResponseError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *AiOpenaiChatCompletions403ResponseError) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *AiOpenaiChatCompletions403ResponseError) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetParam

`func (o *AiOpenaiChatCompletions403ResponseError) GetParam() string`

GetParam returns the Param field if non-nil, zero value otherwise.

### GetParamOk

`func (o *AiOpenaiChatCompletions403ResponseError) GetParamOk() (*string, bool)`

GetParamOk returns a tuple with the Param field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParam

`func (o *AiOpenaiChatCompletions403ResponseError) SetParam(v string)`

SetParam sets Param field to given value.

### HasParam

`func (o *AiOpenaiChatCompletions403ResponseError) HasParam() bool`

HasParam returns a boolean if a field has been set.

### SetParamNil

`func (o *AiOpenaiChatCompletions403ResponseError) SetParamNil(b bool)`

 SetParamNil sets the value for Param to be an explicit nil

### UnsetParam
`func (o *AiOpenaiChatCompletions403ResponseError) UnsetParam()`

UnsetParam ensures that no value is present for Param, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


