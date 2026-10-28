# AiNewItemsDtoAgentNewItemsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | [**AiApiDateTime**](AiApiDateTime.md) | The day the grouped entries were last changed, written with the offset of the portal time zone. The time part  is the moment of the newest entry of the group. | 
**Items** | [**[]AiAgentNewItemsDto**](AiAgentNewItemsDto.md) | What changed on that day, the most recent first. Folders are left out of it, so an entry here is always a file  or a room that holds them. | 

## Methods

### NewAiNewItemsDtoAgentNewItemsDto

`func NewAiNewItemsDtoAgentNewItemsDto(date AiApiDateTime, items []AiAgentNewItemsDto, ) *AiNewItemsDtoAgentNewItemsDto`

NewAiNewItemsDtoAgentNewItemsDto instantiates a new AiNewItemsDtoAgentNewItemsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiNewItemsDtoAgentNewItemsDtoWithDefaults

`func NewAiNewItemsDtoAgentNewItemsDtoWithDefaults() *AiNewItemsDtoAgentNewItemsDto`

NewAiNewItemsDtoAgentNewItemsDtoWithDefaults instantiates a new AiNewItemsDtoAgentNewItemsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *AiNewItemsDtoAgentNewItemsDto) GetDate() AiApiDateTime`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *AiNewItemsDtoAgentNewItemsDto) GetDateOk() (*AiApiDateTime, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *AiNewItemsDtoAgentNewItemsDto) SetDate(v AiApiDateTime)`

SetDate sets Date field to given value.


### GetItems

`func (o *AiNewItemsDtoAgentNewItemsDto) GetItems() []AiAgentNewItemsDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AiNewItemsDtoAgentNewItemsDto) GetItemsOk() (*[]AiAgentNewItemsDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AiNewItemsDtoAgentNewItemsDto) SetItems(v []AiAgentNewItemsDto)`

SetItems sets Items field to given value.


### SetItemsNil

`func (o *AiNewItemsDtoAgentNewItemsDto) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *AiNewItemsDtoAgentNewItemsDto) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


