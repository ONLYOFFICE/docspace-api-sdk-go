# OperationTokenUsageDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalTokens** | Pointer to **int64** | All tokens of the request: prompt plus completion. | [optional] 
**PromptTokens** | Pointer to **int64** | Tokens sent to the model, cached ones included. | [optional] 
**CompletionTokens** | Pointer to **int64** | Tokens the model generated, reasoning ones included. | [optional] 
**CachedTokens** | Pointer to **int64** | Part of the prompt tokens read from the provider cache. | [optional] 
**CacheWriteTokens** | Pointer to **int64** | Part of the prompt tokens written to the provider cache. | [optional] 
**ReasoningTokens** | Pointer to **int64** | Part of the completion tokens the model spent on reasoning. | [optional] 
**ImageTokens** | Pointer to **int64** | Tokens spent on images. | [optional] 

## Methods

### NewOperationTokenUsageDto

`func NewOperationTokenUsageDto() *OperationTokenUsageDto`

NewOperationTokenUsageDto instantiates a new OperationTokenUsageDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationTokenUsageDtoWithDefaults

`func NewOperationTokenUsageDtoWithDefaults() *OperationTokenUsageDto`

NewOperationTokenUsageDtoWithDefaults instantiates a new OperationTokenUsageDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalTokens

`func (o *OperationTokenUsageDto) GetTotalTokens() int64`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *OperationTokenUsageDto) GetTotalTokensOk() (*int64, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *OperationTokenUsageDto) SetTotalTokens(v int64)`

SetTotalTokens sets TotalTokens field to given value.

### HasTotalTokens

`func (o *OperationTokenUsageDto) HasTotalTokens() bool`

HasTotalTokens returns a boolean if a field has been set.

### GetPromptTokens

`func (o *OperationTokenUsageDto) GetPromptTokens() int64`

GetPromptTokens returns the PromptTokens field if non-nil, zero value otherwise.

### GetPromptTokensOk

`func (o *OperationTokenUsageDto) GetPromptTokensOk() (*int64, bool)`

GetPromptTokensOk returns a tuple with the PromptTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokens

`func (o *OperationTokenUsageDto) SetPromptTokens(v int64)`

SetPromptTokens sets PromptTokens field to given value.

### HasPromptTokens

`func (o *OperationTokenUsageDto) HasPromptTokens() bool`

HasPromptTokens returns a boolean if a field has been set.

### GetCompletionTokens

`func (o *OperationTokenUsageDto) GetCompletionTokens() int64`

GetCompletionTokens returns the CompletionTokens field if non-nil, zero value otherwise.

### GetCompletionTokensOk

`func (o *OperationTokenUsageDto) GetCompletionTokensOk() (*int64, bool)`

GetCompletionTokensOk returns a tuple with the CompletionTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTokens

`func (o *OperationTokenUsageDto) SetCompletionTokens(v int64)`

SetCompletionTokens sets CompletionTokens field to given value.

### HasCompletionTokens

`func (o *OperationTokenUsageDto) HasCompletionTokens() bool`

HasCompletionTokens returns a boolean if a field has been set.

### GetCachedTokens

`func (o *OperationTokenUsageDto) GetCachedTokens() int64`

GetCachedTokens returns the CachedTokens field if non-nil, zero value otherwise.

### GetCachedTokensOk

`func (o *OperationTokenUsageDto) GetCachedTokensOk() (*int64, bool)`

GetCachedTokensOk returns a tuple with the CachedTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachedTokens

`func (o *OperationTokenUsageDto) SetCachedTokens(v int64)`

SetCachedTokens sets CachedTokens field to given value.

### HasCachedTokens

`func (o *OperationTokenUsageDto) HasCachedTokens() bool`

HasCachedTokens returns a boolean if a field has been set.

### GetCacheWriteTokens

`func (o *OperationTokenUsageDto) GetCacheWriteTokens() int64`

GetCacheWriteTokens returns the CacheWriteTokens field if non-nil, zero value otherwise.

### GetCacheWriteTokensOk

`func (o *OperationTokenUsageDto) GetCacheWriteTokensOk() (*int64, bool)`

GetCacheWriteTokensOk returns a tuple with the CacheWriteTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheWriteTokens

`func (o *OperationTokenUsageDto) SetCacheWriteTokens(v int64)`

SetCacheWriteTokens sets CacheWriteTokens field to given value.

### HasCacheWriteTokens

`func (o *OperationTokenUsageDto) HasCacheWriteTokens() bool`

HasCacheWriteTokens returns a boolean if a field has been set.

### GetReasoningTokens

`func (o *OperationTokenUsageDto) GetReasoningTokens() int64`

GetReasoningTokens returns the ReasoningTokens field if non-nil, zero value otherwise.

### GetReasoningTokensOk

`func (o *OperationTokenUsageDto) GetReasoningTokensOk() (*int64, bool)`

GetReasoningTokensOk returns a tuple with the ReasoningTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoningTokens

`func (o *OperationTokenUsageDto) SetReasoningTokens(v int64)`

SetReasoningTokens sets ReasoningTokens field to given value.

### HasReasoningTokens

`func (o *OperationTokenUsageDto) HasReasoningTokens() bool`

HasReasoningTokens returns a boolean if a field has been set.

### GetImageTokens

`func (o *OperationTokenUsageDto) GetImageTokens() int64`

GetImageTokens returns the ImageTokens field if non-nil, zero value otherwise.

### GetImageTokensOk

`func (o *OperationTokenUsageDto) GetImageTokensOk() (*int64, bool)`

GetImageTokensOk returns a tuple with the ImageTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTokens

`func (o *OperationTokenUsageDto) SetImageTokens(v int64)`

SetImageTokens sets ImageTokens field to given value.

### HasImageTokens

`func (o *OperationTokenUsageDto) HasImageTokens() bool`

HasImageTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


