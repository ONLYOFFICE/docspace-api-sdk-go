# LinkAiProjectSourcesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoomIds** | **[]int32** | IDs of the Instructions and Knowledge rooms to link. Each must be an active room the caller can read; rooms already linked are skipped. | 

## Methods

### NewLinkAiProjectSourcesRequest

`func NewLinkAiProjectSourcesRequest(roomIds []int32, ) *LinkAiProjectSourcesRequest`

NewLinkAiProjectSourcesRequest instantiates a new LinkAiProjectSourcesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkAiProjectSourcesRequestWithDefaults

`func NewLinkAiProjectSourcesRequestWithDefaults() *LinkAiProjectSourcesRequest`

NewLinkAiProjectSourcesRequestWithDefaults instantiates a new LinkAiProjectSourcesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoomIds

`func (o *LinkAiProjectSourcesRequest) GetRoomIds() []int32`

GetRoomIds returns the RoomIds field if non-nil, zero value otherwise.

### GetRoomIdsOk

`func (o *LinkAiProjectSourcesRequest) GetRoomIdsOk() (*[]int32, bool)`

GetRoomIdsOk returns a tuple with the RoomIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomIds

`func (o *LinkAiProjectSourcesRequest) SetRoomIds(v []int32)`

SetRoomIds sets RoomIds field to given value.


### SetRoomIdsNil

`func (o *LinkAiProjectSourcesRequest) SetRoomIdsNil(b bool)`

 SetRoomIdsNil sets the value for RoomIds to be an explicit nil

### UnsetRoomIds
`func (o *LinkAiProjectSourcesRequest) UnsetRoomIds()`

UnsetRoomIds ensures that no value is present for RoomIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


