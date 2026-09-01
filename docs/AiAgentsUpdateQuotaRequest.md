# AiAgentsUpdateQuotaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoomIds** | [**[]AiAgentsUpdateQuotaRequestRoomIdsInner**](AiAgentsUpdateQuotaRequestRoomIdsInner.md) | Agent (room) ids to update. | 
**Quota** | **float32** | New quota in bytes; a negative value disables the custom quota. | 

## Methods

### NewAiAgentsUpdateQuotaRequest

`func NewAiAgentsUpdateQuotaRequest(roomIds []AiAgentsUpdateQuotaRequestRoomIdsInner, quota float32, ) *AiAgentsUpdateQuotaRequest`

NewAiAgentsUpdateQuotaRequest instantiates a new AiAgentsUpdateQuotaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiAgentsUpdateQuotaRequestWithDefaults

`func NewAiAgentsUpdateQuotaRequestWithDefaults() *AiAgentsUpdateQuotaRequest`

NewAiAgentsUpdateQuotaRequestWithDefaults instantiates a new AiAgentsUpdateQuotaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoomIds

`func (o *AiAgentsUpdateQuotaRequest) GetRoomIds() []AiAgentsUpdateQuotaRequestRoomIdsInner`

GetRoomIds returns the RoomIds field if non-nil, zero value otherwise.

### GetRoomIdsOk

`func (o *AiAgentsUpdateQuotaRequest) GetRoomIdsOk() (*[]AiAgentsUpdateQuotaRequestRoomIdsInner, bool)`

GetRoomIdsOk returns a tuple with the RoomIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomIds

`func (o *AiAgentsUpdateQuotaRequest) SetRoomIds(v []AiAgentsUpdateQuotaRequestRoomIdsInner)`

SetRoomIds sets RoomIds field to given value.


### GetQuota

`func (o *AiAgentsUpdateQuotaRequest) GetQuota() float32`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *AiAgentsUpdateQuotaRequest) GetQuotaOk() (*float32, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *AiAgentsUpdateQuotaRequest) SetQuota(v float32)`

SetQuota sets Quota field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


