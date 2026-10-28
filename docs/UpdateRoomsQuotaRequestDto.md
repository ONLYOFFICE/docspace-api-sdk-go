# UpdateRoomsQuotaRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoomIds** | Pointer to [**[]DuplicateRequestDtoAllOfFileIds**](DuplicateRequestDtoAllOfFileIds.md) | The rooms to change, named by the identifiers that `GET api/2.0/files/rooms` reports. Only whole numbers are  processed, so identifiers of rooms kept in a connected third-party account are skipped without an error. | [optional] 
**Quota** | Pointer to **int64** | The storage each of the listed rooms may take, in bytes. It has to stay inside the portal own limit, and the  per-room quota feature has to be on, otherwise nothing is changed. | [optional] 

## Methods

### NewUpdateRoomsQuotaRequestDto

`func NewUpdateRoomsQuotaRequestDto() *UpdateRoomsQuotaRequestDto`

NewUpdateRoomsQuotaRequestDto instantiates a new UpdateRoomsQuotaRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRoomsQuotaRequestDtoWithDefaults

`func NewUpdateRoomsQuotaRequestDtoWithDefaults() *UpdateRoomsQuotaRequestDto`

NewUpdateRoomsQuotaRequestDtoWithDefaults instantiates a new UpdateRoomsQuotaRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoomIds

`func (o *UpdateRoomsQuotaRequestDto) GetRoomIds() []DuplicateRequestDtoAllOfFileIds`

GetRoomIds returns the RoomIds field if non-nil, zero value otherwise.

### GetRoomIdsOk

`func (o *UpdateRoomsQuotaRequestDto) GetRoomIdsOk() (*[]DuplicateRequestDtoAllOfFileIds, bool)`

GetRoomIdsOk returns a tuple with the RoomIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomIds

`func (o *UpdateRoomsQuotaRequestDto) SetRoomIds(v []DuplicateRequestDtoAllOfFileIds)`

SetRoomIds sets RoomIds field to given value.

### HasRoomIds

`func (o *UpdateRoomsQuotaRequestDto) HasRoomIds() bool`

HasRoomIds returns a boolean if a field has been set.

### SetRoomIdsNil

`func (o *UpdateRoomsQuotaRequestDto) SetRoomIdsNil(b bool)`

 SetRoomIdsNil sets the value for RoomIds to be an explicit nil

### UnsetRoomIds
`func (o *UpdateRoomsQuotaRequestDto) UnsetRoomIds()`

UnsetRoomIds ensures that no value is present for RoomIds, not even an explicit nil
### GetQuota

`func (o *UpdateRoomsQuotaRequestDto) GetQuota() int64`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *UpdateRoomsQuotaRequestDto) GetQuotaOk() (*int64, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *UpdateRoomsQuotaRequestDto) SetQuota(v int64)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *UpdateRoomsQuotaRequestDto) HasQuota() bool`

HasQuota returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


