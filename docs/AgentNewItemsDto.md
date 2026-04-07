# AgentNewItemsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agent** | [**FileEntryBaseDto**](FileEntryBaseDto.md) |  | 
**Items** | [**[]FileEntryBaseDto**](FileEntryBaseDto.md) | The list of file entry items. | 

## Methods

### NewAgentNewItemsDto

`func NewAgentNewItemsDto(agent FileEntryBaseDto, items []FileEntryBaseDto, ) *AgentNewItemsDto`

NewAgentNewItemsDto instantiates a new AgentNewItemsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentNewItemsDtoWithDefaults

`func NewAgentNewItemsDtoWithDefaults() *AgentNewItemsDto`

NewAgentNewItemsDtoWithDefaults instantiates a new AgentNewItemsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgent

`func (o *AgentNewItemsDto) GetAgent() FileEntryBaseDto`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *AgentNewItemsDto) GetAgentOk() (*FileEntryBaseDto, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *AgentNewItemsDto) SetAgent(v FileEntryBaseDto)`

SetAgent sets Agent field to given value.


### GetItems

`func (o *AgentNewItemsDto) GetItems() []FileEntryBaseDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AgentNewItemsDto) GetItemsOk() (*[]FileEntryBaseDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AgentNewItemsDto) SetItems(v []FileEntryBaseDto)`

SetItems sets Items field to given value.


### SetItemsNil

`func (o *AgentNewItemsDto) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *AgentNewItemsDto) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


