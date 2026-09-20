# WebhookRetryRequestsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **[]int32** | The delivery records to send again, by the identifiers `GET api/2.0/settings/webhooks/log` reports. An  identifier that exists nowhere, and one belonging to another member subscription when the caller is not a  DocSpace administrator, is skipped in silence rather than failing the call, so compare the number of records  that come back against the number sent. An empty list is accepted and queues nothing. | [optional] 

## Methods

### NewWebhookRetryRequestsDto

`func NewWebhookRetryRequestsDto() *WebhookRetryRequestsDto`

NewWebhookRetryRequestsDto instantiates a new WebhookRetryRequestsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookRetryRequestsDtoWithDefaults

`func NewWebhookRetryRequestsDtoWithDefaults() *WebhookRetryRequestsDto`

NewWebhookRetryRequestsDtoWithDefaults instantiates a new WebhookRetryRequestsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *WebhookRetryRequestsDto) GetIds() []int32`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *WebhookRetryRequestsDto) GetIdsOk() (*[]int32, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *WebhookRetryRequestsDto) SetIds(v []int32)`

SetIds sets Ids field to given value.

### HasIds

`func (o *WebhookRetryRequestsDto) HasIds() bool`

HasIds returns a boolean if a field has been set.

### SetIdsNil

`func (o *WebhookRetryRequestsDto) SetIdsNil(b bool)`

 SetIdsNil sets the value for Ids to be an explicit nil

### UnsetIds
`func (o *WebhookRetryRequestsDto) UnsetIds()`

UnsetIds ensures that no value is present for Ids, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


