# UnlinkAiProjectSourcesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoomIds** | **[]int32** | IDs of the rooms to unlink. IDs that are not linked to the project are ignored. | 

## Methods

### NewUnlinkAiProjectSourcesRequest

`func NewUnlinkAiProjectSourcesRequest(roomIds []int32, ) *UnlinkAiProjectSourcesRequest`

NewUnlinkAiProjectSourcesRequest instantiates a new UnlinkAiProjectSourcesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnlinkAiProjectSourcesRequestWithDefaults

`func NewUnlinkAiProjectSourcesRequestWithDefaults() *UnlinkAiProjectSourcesRequest`

NewUnlinkAiProjectSourcesRequestWithDefaults instantiates a new UnlinkAiProjectSourcesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoomIds

`func (o *UnlinkAiProjectSourcesRequest) GetRoomIds() []int32`

GetRoomIds returns the RoomIds field if non-nil, zero value otherwise.

### GetRoomIdsOk

`func (o *UnlinkAiProjectSourcesRequest) GetRoomIdsOk() (*[]int32, bool)`

GetRoomIdsOk returns a tuple with the RoomIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomIds

`func (o *UnlinkAiProjectSourcesRequest) SetRoomIds(v []int32)`

SetRoomIds sets RoomIds field to given value.


### SetRoomIdsNil

`func (o *UnlinkAiProjectSourcesRequest) SetRoomIdsNil(b bool)`

 SetRoomIdsNil sets the value for RoomIds to be an explicit nil

### UnsetRoomIds
`func (o *UnlinkAiProjectSourcesRequest) UnsetRoomIds()`

UnsetRoomIds ensures that no value is present for RoomIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


