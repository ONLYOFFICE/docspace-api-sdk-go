# RoomGroupRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name to show the group under. Surrounding spaces are trimmed before it is stored, a name that is blank  once trimmed is refused, and the name does not have to differ from the names of the caller's other groups. | 
**Icon** | **string** | The icon of the group, given as the identifier of one of the built-in covers listed by  `GET api/2.0/files/rooms/covers`. An uploaded image cannot be used, and any value that is not one of those  identifiers is refused. | 
**Rooms** | [**[]DuplicateRequestDtoAllOfFileIds**](DuplicateRequestDtoAllOfFileIds.md) | The rooms to gather in the group, each given as a number for a room stored in the portal or as a string for a  room on a connected third-party account. Every identifier has to name a room the caller can read; repeats are  collapsed, and an element of any other shape - a decimal number, a number sent as a string, null - is refused. | 

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


