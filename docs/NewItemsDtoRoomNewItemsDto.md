# NewItemsDtoRoomNewItemsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | [**ApiDateTime**](ApiDateTime.md) | The day the grouped entries were last changed, written with the offset of the portal time zone. The time part  is the moment of the newest entry of the group. | 
**Items** | [**[]RoomNewItemsDto**](RoomNewItemsDto.md) | What changed on that day, the most recent first. Folders are left out of it, so an entry here is always a file  or a room that holds them. | 

## Methods

### NewNewItemsDtoRoomNewItemsDto

`func NewNewItemsDtoRoomNewItemsDto(date ApiDateTime, items []RoomNewItemsDto, ) *NewItemsDtoRoomNewItemsDto`

NewNewItemsDtoRoomNewItemsDto instantiates a new NewItemsDtoRoomNewItemsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewItemsDtoRoomNewItemsDtoWithDefaults

`func NewNewItemsDtoRoomNewItemsDtoWithDefaults() *NewItemsDtoRoomNewItemsDto`

NewNewItemsDtoRoomNewItemsDtoWithDefaults instantiates a new NewItemsDtoRoomNewItemsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *NewItemsDtoRoomNewItemsDto) GetDate() ApiDateTime`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *NewItemsDtoRoomNewItemsDto) GetDateOk() (*ApiDateTime, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *NewItemsDtoRoomNewItemsDto) SetDate(v ApiDateTime)`

SetDate sets Date field to given value.


### GetItems

`func (o *NewItemsDtoRoomNewItemsDto) GetItems() []RoomNewItemsDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *NewItemsDtoRoomNewItemsDto) GetItemsOk() (*[]RoomNewItemsDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *NewItemsDtoRoomNewItemsDto) SetItems(v []RoomNewItemsDto)`

SetItems sets Items field to given value.


### SetItemsNil

`func (o *NewItemsDtoRoomNewItemsDto) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *NewItemsDtoRoomNewItemsDto) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


