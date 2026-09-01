# AiAgentsResetQuotaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoomIds** | [**[]AiAgentsUpdateQuotaRequestRoomIdsInner**](AiAgentsUpdateQuotaRequestRoomIdsInner.md) | Agent (room) ids to reset to the tenant default quota. | 

## Methods

### NewAiAgentsResetQuotaRequest

`func NewAiAgentsResetQuotaRequest(roomIds []AiAgentsUpdateQuotaRequestRoomIdsInner, ) *AiAgentsResetQuotaRequest`

NewAiAgentsResetQuotaRequest instantiates a new AiAgentsResetQuotaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiAgentsResetQuotaRequestWithDefaults

`func NewAiAgentsResetQuotaRequestWithDefaults() *AiAgentsResetQuotaRequest`

NewAiAgentsResetQuotaRequestWithDefaults instantiates a new AiAgentsResetQuotaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoomIds

`func (o *AiAgentsResetQuotaRequest) GetRoomIds() []AiAgentsUpdateQuotaRequestRoomIdsInner`

GetRoomIds returns the RoomIds field if non-nil, zero value otherwise.

### GetRoomIdsOk

`func (o *AiAgentsResetQuotaRequest) GetRoomIdsOk() (*[]AiAgentsUpdateQuotaRequestRoomIdsInner, bool)`

GetRoomIdsOk returns a tuple with the RoomIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomIds

`func (o *AiAgentsResetQuotaRequest) SetRoomIds(v []AiAgentsUpdateQuotaRequestRoomIdsInner)`

SetRoomIds sets RoomIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


