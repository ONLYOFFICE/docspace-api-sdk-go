# UpdateRoomsQuotaRequestDtoInteger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoomIds** | Pointer to [**[]DuplicateRequestDtoAllOfFileIds**](DuplicateRequestDtoAllOfFileIds.md) | The list of room IDs. | [optional] 
**Quota** | Pointer to **int64** | The room quota. | [optional] 

## Methods

### NewUpdateRoomsQuotaRequestDtoInteger

`func NewUpdateRoomsQuotaRequestDtoInteger() *UpdateRoomsQuotaRequestDtoInteger`

NewUpdateRoomsQuotaRequestDtoInteger instantiates a new UpdateRoomsQuotaRequestDtoInteger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRoomsQuotaRequestDtoIntegerWithDefaults

`func NewUpdateRoomsQuotaRequestDtoIntegerWithDefaults() *UpdateRoomsQuotaRequestDtoInteger`

NewUpdateRoomsQuotaRequestDtoIntegerWithDefaults instantiates a new UpdateRoomsQuotaRequestDtoInteger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoomIds

`func (o *UpdateRoomsQuotaRequestDtoInteger) GetRoomIds() []DuplicateRequestDtoAllOfFileIds`

GetRoomIds returns the RoomIds field if non-nil, zero value otherwise.

### GetRoomIdsOk

`func (o *UpdateRoomsQuotaRequestDtoInteger) GetRoomIdsOk() (*[]DuplicateRequestDtoAllOfFileIds, bool)`

GetRoomIdsOk returns a tuple with the RoomIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomIds

`func (o *UpdateRoomsQuotaRequestDtoInteger) SetRoomIds(v []DuplicateRequestDtoAllOfFileIds)`

SetRoomIds sets RoomIds field to given value.

### HasRoomIds

`func (o *UpdateRoomsQuotaRequestDtoInteger) HasRoomIds() bool`

HasRoomIds returns a boolean if a field has been set.

### SetRoomIdsNil

`func (o *UpdateRoomsQuotaRequestDtoInteger) SetRoomIdsNil(b bool)`

 SetRoomIdsNil sets the value for RoomIds to be an explicit nil

### UnsetRoomIds
`func (o *UpdateRoomsQuotaRequestDtoInteger) UnsetRoomIds()`

UnsetRoomIds ensures that no value is present for RoomIds, not even an explicit nil
### GetQuota

`func (o *UpdateRoomsQuotaRequestDtoInteger) GetQuota() int64`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *UpdateRoomsQuotaRequestDtoInteger) GetQuotaOk() (*int64, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *UpdateRoomsQuotaRequestDtoInteger) SetQuota(v int64)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *UpdateRoomsQuotaRequestDtoInteger) HasQuota() bool`

HasQuota returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


