# AiNewItemsDtoAgentNewItemsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **NullableTime** | The date and time when the new item was created. | 
**Items** | [**[]AiAgentNewItemsDto**](AiAgentNewItemsDto.md) | The list of items. | 

## Methods

### NewAiNewItemsDtoAgentNewItemsDto

`func NewAiNewItemsDtoAgentNewItemsDto(date NullableTime, items []AiAgentNewItemsDto, ) *AiNewItemsDtoAgentNewItemsDto`

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

`func (o *AiNewItemsDtoAgentNewItemsDto) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *AiNewItemsDtoAgentNewItemsDto) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *AiNewItemsDtoAgentNewItemsDto) SetDate(v time.Time)`

SetDate sets Date field to given value.


### SetDateNil

`func (o *AiNewItemsDtoAgentNewItemsDto) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *AiNewItemsDtoAgentNewItemsDto) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
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


