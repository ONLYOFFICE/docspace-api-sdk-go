# RoomNewItemsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Room** | Pointer to [**FileEntryBaseDto**](FileEntryBaseDto.md) | The room file entry. | [optional] 
**Items** | Pointer to [**[]FileEntryBaseDto**](FileEntryBaseDto.md) | The list of file entry items. | [optional] 

## Methods

### NewRoomNewItemsDto

`func NewRoomNewItemsDto() *RoomNewItemsDto`

NewRoomNewItemsDto instantiates a new RoomNewItemsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoomNewItemsDtoWithDefaults

`func NewRoomNewItemsDtoWithDefaults() *RoomNewItemsDto`

NewRoomNewItemsDtoWithDefaults instantiates a new RoomNewItemsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoom

`func (o *RoomNewItemsDto) GetRoom() FileEntryBaseDto`

GetRoom returns the Room field if non-nil, zero value otherwise.

### GetRoomOk

`func (o *RoomNewItemsDto) GetRoomOk() (*FileEntryBaseDto, bool)`

GetRoomOk returns a tuple with the Room field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoom

`func (o *RoomNewItemsDto) SetRoom(v FileEntryBaseDto)`

SetRoom sets Room field to given value.

### HasRoom

`func (o *RoomNewItemsDto) HasRoom() bool`

HasRoom returns a boolean if a field has been set.

### GetItems

`func (o *RoomNewItemsDto) GetItems() []FileEntryBaseDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *RoomNewItemsDto) GetItemsOk() (*[]FileEntryBaseDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *RoomNewItemsDto) SetItems(v []FileEntryBaseDto)`

SetItems sets Items field to given value.

### HasItems

`func (o *RoomNewItemsDto) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *RoomNewItemsDto) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *RoomNewItemsDto) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


