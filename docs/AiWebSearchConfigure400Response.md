# AiWebSearchConfigure400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | True when the configuration was persisted. | 
**Config** | Pointer to [**AiWebSearchConfig**](AiWebSearchConfig.md) | The persisted web-search configuration. Present on success. | [optional] 
**Error** | **string** | The error message, ready to be shown to the caller. | 

## Methods

### NewAiWebSearchConfigure400Response

`func NewAiWebSearchConfigure400Response(success bool, error_ string, ) *AiWebSearchConfigure400Response`

NewAiWebSearchConfigure400Response instantiates a new AiWebSearchConfigure400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiWebSearchConfigure400ResponseWithDefaults

`func NewAiWebSearchConfigure400ResponseWithDefaults() *AiWebSearchConfigure400Response`

NewAiWebSearchConfigure400ResponseWithDefaults instantiates a new AiWebSearchConfigure400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *AiWebSearchConfigure400Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AiWebSearchConfigure400Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AiWebSearchConfigure400Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetConfig

`func (o *AiWebSearchConfigure400Response) GetConfig() AiWebSearchConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AiWebSearchConfigure400Response) GetConfigOk() (*AiWebSearchConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AiWebSearchConfigure400Response) SetConfig(v AiWebSearchConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AiWebSearchConfigure400Response) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetError

`func (o *AiWebSearchConfigure400Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AiWebSearchConfigure400Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AiWebSearchConfigure400Response) SetError(v string)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


