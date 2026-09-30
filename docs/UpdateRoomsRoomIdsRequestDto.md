# UpdateRoomsRoomIdsRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoomIds** | Pointer to [**[]DuplicateRequestDtoAllOfFileIds**](DuplicateRequestDtoAllOfFileIds.md) | The rooms to reset, named by the identifiers that `GET api/2.0/files/rooms` reports. Only whole numbers are  processed, so identifiers of rooms kept in a connected third-party account are skipped without an error. | [optional] 

## Methods

### NewUpdateRoomsRoomIdsRequestDto

`func NewUpdateRoomsRoomIdsRequestDto() *UpdateRoomsRoomIdsRequestDto`

NewUpdateRoomsRoomIdsRequestDto instantiates a new UpdateRoomsRoomIdsRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRoomsRoomIdsRequestDtoWithDefaults

`func NewUpdateRoomsRoomIdsRequestDtoWithDefaults() *UpdateRoomsRoomIdsRequestDto`

NewUpdateRoomsRoomIdsRequestDtoWithDefaults instantiates a new UpdateRoomsRoomIdsRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoomIds

`func (o *UpdateRoomsRoomIdsRequestDto) GetRoomIds() []DuplicateRequestDtoAllOfFileIds`

GetRoomIds returns the RoomIds field if non-nil, zero value otherwise.

### GetRoomIdsOk

`func (o *UpdateRoomsRoomIdsRequestDto) GetRoomIdsOk() (*[]DuplicateRequestDtoAllOfFileIds, bool)`

GetRoomIdsOk returns a tuple with the RoomIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomIds

`func (o *UpdateRoomsRoomIdsRequestDto) SetRoomIds(v []DuplicateRequestDtoAllOfFileIds)`

SetRoomIds sets RoomIds field to given value.

### HasRoomIds

`func (o *UpdateRoomsRoomIdsRequestDto) HasRoomIds() bool`

HasRoomIds returns a boolean if a field has been set.

### SetRoomIdsNil

`func (o *UpdateRoomsRoomIdsRequestDto) SetRoomIdsNil(b bool)`

 SetRoomIdsNil sets the value for RoomIds to be an explicit nil

### UnsetRoomIds
`func (o *UpdateRoomsRoomIdsRequestDto) UnsetRoomIds()`

UnsetRoomIds ensures that no value is present for RoomIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


