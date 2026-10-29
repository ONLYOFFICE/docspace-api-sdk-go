# WebhookRetryRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **[]int32** | The delivery records to send again, by the `id` values `GET api/2.0/settings/webhooks/log` reports. The bound  is per call: a longer batch has to be split into several calls. | [optional] 

## Methods

### NewWebhookRetryRequestDto

`func NewWebhookRetryRequestDto() *WebhookRetryRequestDto`

NewWebhookRetryRequestDto instantiates a new WebhookRetryRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookRetryRequestDtoWithDefaults

`func NewWebhookRetryRequestDtoWithDefaults() *WebhookRetryRequestDto`

NewWebhookRetryRequestDtoWithDefaults instantiates a new WebhookRetryRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *WebhookRetryRequestDto) GetIds() []int32`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *WebhookRetryRequestDto) GetIdsOk() (*[]int32, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *WebhookRetryRequestDto) SetIds(v []int32)`

SetIds sets Ids field to given value.

### HasIds

`func (o *WebhookRetryRequestDto) HasIds() bool`

HasIds returns a boolean if a field has been set.

### SetIdsNil

`func (o *WebhookRetryRequestDto) SetIdsNil(b bool)`

 SetIdsNil sets the value for Ids to be an explicit nil

### UnsetIds
`func (o *WebhookRetryRequestDto) UnsetIds()`

UnsetIds ensures that no value is present for Ids, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


