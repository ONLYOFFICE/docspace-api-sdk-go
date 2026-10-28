# AiErrorData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | The rejected field. | 
**Message** | **string** | The human-readable reason the field was rejected. | 
**Code** | Pointer to [**AiErrorCode**](AiErrorCode.md) | Optional: producers that know the reason set it (the prompts engine does). | [optional] 

## Methods

### NewAiErrorData

`func NewAiErrorData(field string, message string, ) *AiErrorData`

NewAiErrorData instantiates a new AiErrorData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiErrorDataWithDefaults

`func NewAiErrorDataWithDefaults() *AiErrorData`

NewAiErrorDataWithDefaults instantiates a new AiErrorData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *AiErrorData) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *AiErrorData) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *AiErrorData) SetField(v string)`

SetField sets Field field to given value.


### GetMessage

`func (o *AiErrorData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AiErrorData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AiErrorData) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetCode

`func (o *AiErrorData) GetCode() AiErrorCode`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AiErrorData) GetCodeOk() (*AiErrorCode, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AiErrorData) SetCode(v AiErrorCode)`

SetCode sets Code field to given value.

### HasCode

`func (o *AiErrorData) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


