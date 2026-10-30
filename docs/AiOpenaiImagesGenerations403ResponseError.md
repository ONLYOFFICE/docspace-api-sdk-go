# AiOpenaiImagesGenerations403ResponseError

Inherits from [map[string]interface{}](map[string]interface{}.md).

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Human-readable description of the failure. | 
**Type** | **string** | OpenAI error class, for example `invalid_request_error`. | 
**Code** | Pointer to **NullableString** | Machine-readable code, when the provider supplies one. | [optional] 
**Param** | Pointer to **NullableString** | The request parameter at fault, when the failure names one. | [optional] 

## Methods

### NewAiOpenaiImagesGenerations403ResponseError

`func NewAiOpenaiImagesGenerations403ResponseError(message string, type_ string, ) *AiOpenaiImagesGenerations403ResponseError`

NewAiOpenaiImagesGenerations403ResponseError instantiates a new AiOpenaiImagesGenerations403ResponseError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiOpenaiImagesGenerations403ResponseErrorWithDefaults

`func NewAiOpenaiImagesGenerations403ResponseErrorWithDefaults() *AiOpenaiImagesGenerations403ResponseError`

NewAiOpenaiImagesGenerations403ResponseErrorWithDefaults instantiates a new AiOpenaiImagesGenerations403ResponseError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *AiOpenaiImagesGenerations403ResponseError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AiOpenaiImagesGenerations403ResponseError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AiOpenaiImagesGenerations403ResponseError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetType

`func (o *AiOpenaiImagesGenerations403ResponseError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AiOpenaiImagesGenerations403ResponseError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AiOpenaiImagesGenerations403ResponseError) SetType(v string)`

SetType sets Type field to given value.


### GetCode

`func (o *AiOpenaiImagesGenerations403ResponseError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AiOpenaiImagesGenerations403ResponseError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AiOpenaiImagesGenerations403ResponseError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AiOpenaiImagesGenerations403ResponseError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *AiOpenaiImagesGenerations403ResponseError) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *AiOpenaiImagesGenerations403ResponseError) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetParam

`func (o *AiOpenaiImagesGenerations403ResponseError) GetParam() string`

GetParam returns the Param field if non-nil, zero value otherwise.

### GetParamOk

`func (o *AiOpenaiImagesGenerations403ResponseError) GetParamOk() (*string, bool)`

GetParamOk returns a tuple with the Param field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParam

`func (o *AiOpenaiImagesGenerations403ResponseError) SetParam(v string)`

SetParam sets Param field to given value.

### HasParam

`func (o *AiOpenaiImagesGenerations403ResponseError) HasParam() bool`

HasParam returns a boolean if a field has been set.

### SetParamNil

`func (o *AiOpenaiImagesGenerations403ResponseError) SetParamNil(b bool)`

 SetParamNil sets the value for Param to be an explicit nil

### UnsetParam
`func (o *AiOpenaiImagesGenerations403ResponseError) UnsetParam()`

UnsetParam ensures that no value is present for Param, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


