# AiPromptsUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Prompt id to update. | 
**Updates** | [**AiPromptsUpdateRequestUpdates**](AiPromptsUpdateRequestUpdates.md) |  | 

## Methods

### NewAiPromptsUpdateRequest

`func NewAiPromptsUpdateRequest(id string, updates AiPromptsUpdateRequestUpdates, ) *AiPromptsUpdateRequest`

NewAiPromptsUpdateRequest instantiates a new AiPromptsUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiPromptsUpdateRequestWithDefaults

`func NewAiPromptsUpdateRequestWithDefaults() *AiPromptsUpdateRequest`

NewAiPromptsUpdateRequestWithDefaults instantiates a new AiPromptsUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AiPromptsUpdateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiPromptsUpdateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiPromptsUpdateRequest) SetId(v string)`

SetId sets Id field to given value.


### GetUpdates

`func (o *AiPromptsUpdateRequest) GetUpdates() AiPromptsUpdateRequestUpdates`

GetUpdates returns the Updates field if non-nil, zero value otherwise.

### GetUpdatesOk

`func (o *AiPromptsUpdateRequest) GetUpdatesOk() (*AiPromptsUpdateRequestUpdates, bool)`

GetUpdatesOk returns a tuple with the Updates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdates

`func (o *AiPromptsUpdateRequest) SetUpdates(v AiPromptsUpdateRequestUpdates)`

SetUpdates sets Updates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


