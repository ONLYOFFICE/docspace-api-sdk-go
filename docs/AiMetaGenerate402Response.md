# AiMetaGenerate402Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** | Short reason, ready to be shown to the caller. | 
**Code** | **string** | Stable classification of the failure: `rate_limit`, `insufficient_funds`, `auth`, `model_not_found`, `timeout`, `network`, `bad_request`, `server`, `unknown`, or a provider-specific code relayed as it came. | 
**ProviderMessage** | Pointer to **string** | The provider's own error message, cut to 500 characters. Absent when the failure never reached the provider, or was a timeout. | [optional] 

## Methods

### NewAiMetaGenerate402Response

`func NewAiMetaGenerate402Response(error_ string, code string, ) *AiMetaGenerate402Response`

NewAiMetaGenerate402Response instantiates a new AiMetaGenerate402Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiMetaGenerate402ResponseWithDefaults

`func NewAiMetaGenerate402ResponseWithDefaults() *AiMetaGenerate402Response`

NewAiMetaGenerate402ResponseWithDefaults instantiates a new AiMetaGenerate402Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *AiMetaGenerate402Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AiMetaGenerate402Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AiMetaGenerate402Response) SetError(v string)`

SetError sets Error field to given value.


### GetCode

`func (o *AiMetaGenerate402Response) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AiMetaGenerate402Response) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AiMetaGenerate402Response) SetCode(v string)`

SetCode sets Code field to given value.


### GetProviderMessage

`func (o *AiMetaGenerate402Response) GetProviderMessage() string`

GetProviderMessage returns the ProviderMessage field if non-nil, zero value otherwise.

### GetProviderMessageOk

`func (o *AiMetaGenerate402Response) GetProviderMessageOk() (*string, bool)`

GetProviderMessageOk returns a tuple with the ProviderMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderMessage

`func (o *AiMetaGenerate402Response) SetProviderMessage(v string)`

SetProviderMessage sets ProviderMessage field to given value.

### HasProviderMessage

`func (o *AiMetaGenerate402Response) HasProviderMessage() bool`

HasProviderMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


