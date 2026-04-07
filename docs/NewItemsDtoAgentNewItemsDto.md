# NewItemsDtoAgentNewItemsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | [**ApiDateTime**](ApiDateTime.md) |  | 
**Items** | [**[]AgentNewItemsDto**](AgentNewItemsDto.md) | The list of items. | 

## Methods

### NewNewItemsDtoAgentNewItemsDto

`func NewNewItemsDtoAgentNewItemsDto(date ApiDateTime, items []AgentNewItemsDto, ) *NewItemsDtoAgentNewItemsDto`

NewNewItemsDtoAgentNewItemsDto instantiates a new NewItemsDtoAgentNewItemsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewItemsDtoAgentNewItemsDtoWithDefaults

`func NewNewItemsDtoAgentNewItemsDtoWithDefaults() *NewItemsDtoAgentNewItemsDto`

NewNewItemsDtoAgentNewItemsDtoWithDefaults instantiates a new NewItemsDtoAgentNewItemsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *NewItemsDtoAgentNewItemsDto) GetDate() ApiDateTime`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *NewItemsDtoAgentNewItemsDto) GetDateOk() (*ApiDateTime, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *NewItemsDtoAgentNewItemsDto) SetDate(v ApiDateTime)`

SetDate sets Date field to given value.


### GetItems

`func (o *NewItemsDtoAgentNewItemsDto) GetItems() []AgentNewItemsDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *NewItemsDtoAgentNewItemsDto) GetItemsOk() (*[]AgentNewItemsDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *NewItemsDtoAgentNewItemsDto) SetItems(v []AgentNewItemsDto)`

SetItems sets Items field to given value.


### SetItemsNil

`func (o *NewItemsDtoAgentNewItemsDto) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *NewItemsDtoAgentNewItemsDto) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


