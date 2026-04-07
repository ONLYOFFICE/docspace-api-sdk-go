# RoomGroupRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Group name | 
**Icon** | **string** | Group icon | 
**Rooms** | [**[]DuplicateRequestDtoAllOfFileIds**](DuplicateRequestDtoAllOfFileIds.md) | The list of room IDs. | 

## Methods

### NewRoomGroupRequestDto

`func NewRoomGroupRequestDto(name string, icon string, rooms []DuplicateRequestDtoAllOfFileIds, ) *RoomGroupRequestDto`

NewRoomGroupRequestDto instantiates a new RoomGroupRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoomGroupRequestDtoWithDefaults

`func NewRoomGroupRequestDtoWithDefaults() *RoomGroupRequestDto`

NewRoomGroupRequestDtoWithDefaults instantiates a new RoomGroupRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RoomGroupRequestDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoomGroupRequestDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoomGroupRequestDto) SetName(v string)`

SetName sets Name field to given value.


### GetIcon

`func (o *RoomGroupRequestDto) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *RoomGroupRequestDto) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *RoomGroupRequestDto) SetIcon(v string)`

SetIcon sets Icon field to given value.


### GetRooms

`func (o *RoomGroupRequestDto) GetRooms() []DuplicateRequestDtoAllOfFileIds`

GetRooms returns the Rooms field if non-nil, zero value otherwise.

### GetRoomsOk

`func (o *RoomGroupRequestDto) GetRoomsOk() (*[]DuplicateRequestDtoAllOfFileIds, bool)`

GetRoomsOk returns a tuple with the Rooms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRooms

`func (o *RoomGroupRequestDto) SetRooms(v []DuplicateRequestDtoAllOfFileIds)`

SetRooms sets Rooms field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


