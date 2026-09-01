# NewItemsDtoRoomNewItemsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **NullableTime** | The date and time when the new item was created. | 
**Items** | [**[]RoomNewItemsDto**](RoomNewItemsDto.md) | The list of items. | 

## Methods

### NewNewItemsDtoRoomNewItemsDto

`func NewNewItemsDtoRoomNewItemsDto(date NullableTime, items []RoomNewItemsDto, ) *NewItemsDtoRoomNewItemsDto`

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

`func (o *NewItemsDtoRoomNewItemsDto) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *NewItemsDtoRoomNewItemsDto) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *NewItemsDtoRoomNewItemsDto) SetDate(v time.Time)`

SetDate sets Date field to given value.


### SetDateNil

`func (o *NewItemsDtoRoomNewItemsDto) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *NewItemsDtoRoomNewItemsDto) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
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


