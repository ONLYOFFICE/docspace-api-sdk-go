# WebhookTriggerDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The event name exactly as it appears in a delivered payload, so a receiver can match on it. The entry  named `*` is not an event but the catch-all. | [optional] 
**Id** | Pointer to **int64** | The bit that stands for this event in the `triggers` bitmask of a subscription. Add the bits of the wanted  events together; the catch-all entry has the value `0` and is used on its own rather than added to  anything. | [optional] 
**Available** | Pointer to **bool** | Whether the caller's own role may subscribe to this event - a plain member cannot subscribe to user, group  or room creation, where a room administrator can. An unavailable event is listed all the same, and sending  its bit to `POST api/2.0/settings/webhook` is refused as an invalid request. | [optional] 

## Methods

### NewWebhookTriggerDto

`func NewWebhookTriggerDto() *WebhookTriggerDto`

NewWebhookTriggerDto instantiates a new WebhookTriggerDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookTriggerDtoWithDefaults

`func NewWebhookTriggerDtoWithDefaults() *WebhookTriggerDto`

NewWebhookTriggerDtoWithDefaults instantiates a new WebhookTriggerDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WebhookTriggerDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookTriggerDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookTriggerDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WebhookTriggerDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *WebhookTriggerDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *WebhookTriggerDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetId

`func (o *WebhookTriggerDto) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookTriggerDto) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookTriggerDto) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *WebhookTriggerDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAvailable

`func (o *WebhookTriggerDto) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *WebhookTriggerDto) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *WebhookTriggerDto) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *WebhookTriggerDto) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


