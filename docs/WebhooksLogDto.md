# WebhooksLogDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The identifier of this delivery record, which is what `PUT api/2.0/settings/webhook/{id}/retry` re-sends.  The automatic retries of the delivery update this same record, while a manual retry gets a record and an  identifier of its own. | 
**EventId** | Pointer to **int32** | The identifier of the event the delivery is for, which the receiver sees as `event.id` and in the  `x-docspace-event-id` header. It is the `id` of the event's first delivery record, so it equals `id` there and  stays the same on every manual retry, which makes it the value to group the records of one event by. | [optional] 
**RetriedBy** | Pointer to [**EmployeeDto**](EmployeeDto.md) | The member who asked for this delivery with a manual retry. It is empty on the first delivery of an event,  which the portal sends by itself. | [optional] 
**ConfigName** | Pointer to **NullableString** | The name of the subscription the delivery belongs to. It is the name as it stands now, so it follows a  later rename of the subscription rather than recording what it was called at the time. | [optional] 
**Trigger** | Pointer to [**WebhookTrigger**](WebhookTrigger.md) | The event that caused the delivery, as a single bit rather than a mask - a delivery is always for one  event, even though a subscription covers several. | [optional] 
**CreationTime** | Pointer to **time.Time** | When this delivery record was queued, as a UTC instant ending in `Z`. For the first delivery that is when the  event happened; for a manual retry it is the  moment of the retry. Records come back newest first by this moment. | [optional] 
**Method** | Pointer to **NullableString** | The HTTP method the delivery was sent with, which is `POST` for every webhook the portal sends. | [optional] 
**Route** | Pointer to **NullableString** | The address the delivery is sent to: the subscription's URL as it stands now. The journal does not keep the  address of each attempt, so after the subscription's URL changes, an older record names the new one too. | [optional] 
**RequestHeaders** | Pointer to **NullableString** | The headers the portal sent, serialised as one string, including the signature header a receiver verifies  the payload with. | [optional] 
**RequestPayload** | Pointer to **NullableString** | The body the portal sent, which is the event payload as JSON text. It is stored as it was sent, so it  still describes the entity as it looked at the time of the event. | [optional] 
**ResponseHeaders** | Pointer to **NullableString** | The headers the target answered with on the latest attempt, serialised the same way as `requestHeaders`. It  is empty until an attempt is answered, and when the target was not reached. | [optional] 
**ResponsePayload** | Pointer to **NullableString** | The body the target answered with on the latest attempt, truncated for storage. Empty under the same  conditions as `responseHeaders`, and also for a target that answers with no body at all. A record the portal  stopped without sending holds the portal's own note here instead - a text that starts with `Not sent:` and  gives the reason - which is not an answer of the target: on a record that never went out `status` is then `0`  and `responseHeaders` is empty, while on one stopped after failed attempts they still describe the last of them. | [optional] 
**Status** | Pointer to **int32** | The HTTP status code the target answered on the latest attempt. It is `0` while the first attempt is still  on its way and on one that never reached the target, so `0` is not a failure code - it is the absence of an  answer. While `nextAttemptOn` is set it is the status of the failed attempt that caused the retry. | [optional] 
**Delivery** | Pointer to **NullableTime** | When the successful answer came back, as a UTC instant ending in `Z` like `creationTime`. It stays empty until an  attempt succeeds, and for good once the portal has given up retrying. | [optional] 
**Attempts** | Pointer to **int32** | How many times the portal has tried to deliver this record so far. A failed attempt is retried with  growing pauses - on the default schedule about 1 minute, 5 minutes, 30 minutes, 2, 5 and 10 hours, though  an installation may configure its own - all of them under the same record and the same  `x-docspace-event-id` header, so a receiver can drop a delivery it has already processed. The one exception  is a portal moved to another region while retries are pending: the moved records get new IDs, so the attempts  made there carry a new `x-docspace-event-id`. | [optional] 
**NextAttemptOn** | Pointer to **NullableTime** | When the next attempt is due, as a UTC instant ending in `Z` like `creationTime`. While an attempt is being sent  right now it holds the moment that attempt's lease expires, a few minutes ahead, rather than a scheduled retry;  it changes again once the attempt ends. Empty means no attempt will follow: the record is delivered, its retries  ran out, a manual retry of the same event took over from it, or the portal stopped it - because the subscription  was disabled or removed or its address rejected, the subscription's creator may no longer receive the event,  the stored body cannot be read or is in the payload format of a previous version, the portal is removed, the  record waited longer than the retry schedule allows, the portal was restored from a backup taken before the  record was answered, or the portal was moved to another region while the record was being sent. A stopped  record that was never attempted says why in `responsePayload`; one that was keeps the target's last answer there. | [optional] 

## Methods

### NewWebhooksLogDto

`func NewWebhooksLogDto(id int32, ) *WebhooksLogDto`

NewWebhooksLogDto instantiates a new WebhooksLogDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhooksLogDtoWithDefaults

`func NewWebhooksLogDtoWithDefaults() *WebhooksLogDto`

NewWebhooksLogDtoWithDefaults instantiates a new WebhooksLogDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhooksLogDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhooksLogDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhooksLogDto) SetId(v int32)`

SetId sets Id field to given value.


### GetEventId

`func (o *WebhooksLogDto) GetEventId() int32`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *WebhooksLogDto) GetEventIdOk() (*int32, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *WebhooksLogDto) SetEventId(v int32)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *WebhooksLogDto) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetRetriedBy

`func (o *WebhooksLogDto) GetRetriedBy() EmployeeDto`

GetRetriedBy returns the RetriedBy field if non-nil, zero value otherwise.

### GetRetriedByOk

`func (o *WebhooksLogDto) GetRetriedByOk() (*EmployeeDto, bool)`

GetRetriedByOk returns a tuple with the RetriedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetriedBy

`func (o *WebhooksLogDto) SetRetriedBy(v EmployeeDto)`

SetRetriedBy sets RetriedBy field to given value.

### HasRetriedBy

`func (o *WebhooksLogDto) HasRetriedBy() bool`

HasRetriedBy returns a boolean if a field has been set.

### GetConfigName

`func (o *WebhooksLogDto) GetConfigName() string`

GetConfigName returns the ConfigName field if non-nil, zero value otherwise.

### GetConfigNameOk

`func (o *WebhooksLogDto) GetConfigNameOk() (*string, bool)`

GetConfigNameOk returns a tuple with the ConfigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigName

`func (o *WebhooksLogDto) SetConfigName(v string)`

SetConfigName sets ConfigName field to given value.

### HasConfigName

`func (o *WebhooksLogDto) HasConfigName() bool`

HasConfigName returns a boolean if a field has been set.

### SetConfigNameNil

`func (o *WebhooksLogDto) SetConfigNameNil(b bool)`

 SetConfigNameNil sets the value for ConfigName to be an explicit nil

### UnsetConfigName
`func (o *WebhooksLogDto) UnsetConfigName()`

UnsetConfigName ensures that no value is present for ConfigName, not even an explicit nil
### GetTrigger

`func (o *WebhooksLogDto) GetTrigger() WebhookTrigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *WebhooksLogDto) GetTriggerOk() (*WebhookTrigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *WebhooksLogDto) SetTrigger(v WebhookTrigger)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *WebhooksLogDto) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### GetCreationTime

`func (o *WebhooksLogDto) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *WebhooksLogDto) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *WebhooksLogDto) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *WebhooksLogDto) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetMethod

`func (o *WebhooksLogDto) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *WebhooksLogDto) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *WebhooksLogDto) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *WebhooksLogDto) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### SetMethodNil

`func (o *WebhooksLogDto) SetMethodNil(b bool)`

 SetMethodNil sets the value for Method to be an explicit nil

### UnsetMethod
`func (o *WebhooksLogDto) UnsetMethod()`

UnsetMethod ensures that no value is present for Method, not even an explicit nil
### GetRoute

`func (o *WebhooksLogDto) GetRoute() string`

GetRoute returns the Route field if non-nil, zero value otherwise.

### GetRouteOk

`func (o *WebhooksLogDto) GetRouteOk() (*string, bool)`

GetRouteOk returns a tuple with the Route field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoute

`func (o *WebhooksLogDto) SetRoute(v string)`

SetRoute sets Route field to given value.

### HasRoute

`func (o *WebhooksLogDto) HasRoute() bool`

HasRoute returns a boolean if a field has been set.

### SetRouteNil

`func (o *WebhooksLogDto) SetRouteNil(b bool)`

 SetRouteNil sets the value for Route to be an explicit nil

### UnsetRoute
`func (o *WebhooksLogDto) UnsetRoute()`

UnsetRoute ensures that no value is present for Route, not even an explicit nil
### GetRequestHeaders

`func (o *WebhooksLogDto) GetRequestHeaders() string`

GetRequestHeaders returns the RequestHeaders field if non-nil, zero value otherwise.

### GetRequestHeadersOk

`func (o *WebhooksLogDto) GetRequestHeadersOk() (*string, bool)`

GetRequestHeadersOk returns a tuple with the RequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestHeaders

`func (o *WebhooksLogDto) SetRequestHeaders(v string)`

SetRequestHeaders sets RequestHeaders field to given value.

### HasRequestHeaders

`func (o *WebhooksLogDto) HasRequestHeaders() bool`

HasRequestHeaders returns a boolean if a field has been set.

### SetRequestHeadersNil

`func (o *WebhooksLogDto) SetRequestHeadersNil(b bool)`

 SetRequestHeadersNil sets the value for RequestHeaders to be an explicit nil

### UnsetRequestHeaders
`func (o *WebhooksLogDto) UnsetRequestHeaders()`

UnsetRequestHeaders ensures that no value is present for RequestHeaders, not even an explicit nil
### GetRequestPayload

`func (o *WebhooksLogDto) GetRequestPayload() string`

GetRequestPayload returns the RequestPayload field if non-nil, zero value otherwise.

### GetRequestPayloadOk

`func (o *WebhooksLogDto) GetRequestPayloadOk() (*string, bool)`

GetRequestPayloadOk returns a tuple with the RequestPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestPayload

`func (o *WebhooksLogDto) SetRequestPayload(v string)`

SetRequestPayload sets RequestPayload field to given value.

### HasRequestPayload

`func (o *WebhooksLogDto) HasRequestPayload() bool`

HasRequestPayload returns a boolean if a field has been set.

### SetRequestPayloadNil

`func (o *WebhooksLogDto) SetRequestPayloadNil(b bool)`

 SetRequestPayloadNil sets the value for RequestPayload to be an explicit nil

### UnsetRequestPayload
`func (o *WebhooksLogDto) UnsetRequestPayload()`

UnsetRequestPayload ensures that no value is present for RequestPayload, not even an explicit nil
### GetResponseHeaders

`func (o *WebhooksLogDto) GetResponseHeaders() string`

GetResponseHeaders returns the ResponseHeaders field if non-nil, zero value otherwise.

### GetResponseHeadersOk

`func (o *WebhooksLogDto) GetResponseHeadersOk() (*string, bool)`

GetResponseHeadersOk returns a tuple with the ResponseHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeaders

`func (o *WebhooksLogDto) SetResponseHeaders(v string)`

SetResponseHeaders sets ResponseHeaders field to given value.

### HasResponseHeaders

`func (o *WebhooksLogDto) HasResponseHeaders() bool`

HasResponseHeaders returns a boolean if a field has been set.

### SetResponseHeadersNil

`func (o *WebhooksLogDto) SetResponseHeadersNil(b bool)`

 SetResponseHeadersNil sets the value for ResponseHeaders to be an explicit nil

### UnsetResponseHeaders
`func (o *WebhooksLogDto) UnsetResponseHeaders()`

UnsetResponseHeaders ensures that no value is present for ResponseHeaders, not even an explicit nil
### GetResponsePayload

`func (o *WebhooksLogDto) GetResponsePayload() string`

GetResponsePayload returns the ResponsePayload field if non-nil, zero value otherwise.

### GetResponsePayloadOk

`func (o *WebhooksLogDto) GetResponsePayloadOk() (*string, bool)`

GetResponsePayloadOk returns a tuple with the ResponsePayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsePayload

`func (o *WebhooksLogDto) SetResponsePayload(v string)`

SetResponsePayload sets ResponsePayload field to given value.

### HasResponsePayload

`func (o *WebhooksLogDto) HasResponsePayload() bool`

HasResponsePayload returns a boolean if a field has been set.

### SetResponsePayloadNil

`func (o *WebhooksLogDto) SetResponsePayloadNil(b bool)`

 SetResponsePayloadNil sets the value for ResponsePayload to be an explicit nil

### UnsetResponsePayload
`func (o *WebhooksLogDto) UnsetResponsePayload()`

UnsetResponsePayload ensures that no value is present for ResponsePayload, not even an explicit nil
### GetStatus

`func (o *WebhooksLogDto) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebhooksLogDto) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebhooksLogDto) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WebhooksLogDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDelivery

`func (o *WebhooksLogDto) GetDelivery() time.Time`

GetDelivery returns the Delivery field if non-nil, zero value otherwise.

### GetDeliveryOk

`func (o *WebhooksLogDto) GetDeliveryOk() (*time.Time, bool)`

GetDeliveryOk returns a tuple with the Delivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivery

`func (o *WebhooksLogDto) SetDelivery(v time.Time)`

SetDelivery sets Delivery field to given value.

### HasDelivery

`func (o *WebhooksLogDto) HasDelivery() bool`

HasDelivery returns a boolean if a field has been set.

### SetDeliveryNil

`func (o *WebhooksLogDto) SetDeliveryNil(b bool)`

 SetDeliveryNil sets the value for Delivery to be an explicit nil

### UnsetDelivery
`func (o *WebhooksLogDto) UnsetDelivery()`

UnsetDelivery ensures that no value is present for Delivery, not even an explicit nil
### GetAttempts

`func (o *WebhooksLogDto) GetAttempts() int32`

GetAttempts returns the Attempts field if non-nil, zero value otherwise.

### GetAttemptsOk

`func (o *WebhooksLogDto) GetAttemptsOk() (*int32, bool)`

GetAttemptsOk returns a tuple with the Attempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempts

`func (o *WebhooksLogDto) SetAttempts(v int32)`

SetAttempts sets Attempts field to given value.

### HasAttempts

`func (o *WebhooksLogDto) HasAttempts() bool`

HasAttempts returns a boolean if a field has been set.

### GetNextAttemptOn

`func (o *WebhooksLogDto) GetNextAttemptOn() time.Time`

GetNextAttemptOn returns the NextAttemptOn field if non-nil, zero value otherwise.

### GetNextAttemptOnOk

`func (o *WebhooksLogDto) GetNextAttemptOnOk() (*time.Time, bool)`

GetNextAttemptOnOk returns a tuple with the NextAttemptOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAttemptOn

`func (o *WebhooksLogDto) SetNextAttemptOn(v time.Time)`

SetNextAttemptOn sets NextAttemptOn field to given value.

### HasNextAttemptOn

`func (o *WebhooksLogDto) HasNextAttemptOn() bool`

HasNextAttemptOn returns a boolean if a field has been set.

### SetNextAttemptOnNil

`func (o *WebhooksLogDto) SetNextAttemptOnNil(b bool)`

 SetNextAttemptOnNil sets the value for NextAttemptOn to be an explicit nil

### UnsetNextAttemptOn
`func (o *WebhooksLogDto) UnsetNextAttemptOn()`

UnsetNextAttemptOn ensures that no value is present for NextAttemptOn, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


