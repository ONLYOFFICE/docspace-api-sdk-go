# AiAgentNewItemsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agent** | [**AiFileEntryBaseDto**](AiFileEntryBaseDto.md) | The agent file entry. | 
**Items** | [**[]AiFileEntryBaseDto**](AiFileEntryBaseDto.md) | The list of file entry items. | 

## Methods

### NewAiAgentNewItemsDto

`func NewAiAgentNewItemsDto(agent AiFileEntryBaseDto, items []AiFileEntryBaseDto, ) *AiAgentNewItemsDto`

NewAiAgentNewItemsDto instantiates a new AiAgentNewItemsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiAgentNewItemsDtoWithDefaults

`func NewAiAgentNewItemsDtoWithDefaults() *AiAgentNewItemsDto`

NewAiAgentNewItemsDtoWithDefaults instantiates a new AiAgentNewItemsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgent

`func (o *AiAgentNewItemsDto) GetAgent() AiFileEntryBaseDto`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *AiAgentNewItemsDto) GetAgentOk() (*AiFileEntryBaseDto, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *AiAgentNewItemsDto) SetAgent(v AiFileEntryBaseDto)`

SetAgent sets Agent field to given value.


### GetItems

`func (o *AiAgentNewItemsDto) GetItems() []AiFileEntryBaseDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AiAgentNewItemsDto) GetItemsOk() (*[]AiFileEntryBaseDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AiAgentNewItemsDto) SetItems(v []AiFileEntryBaseDto)`

SetItems sets Items field to given value.


### SetItemsNil

`func (o *AiAgentNewItemsDto) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *AiAgentNewItemsDto) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


