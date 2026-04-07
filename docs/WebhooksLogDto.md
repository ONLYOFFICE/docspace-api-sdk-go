# WebhooksLogDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The webhook log ID. | 
**ConfigName** | Pointer to **NullableString** | The webhook configuration name. | [optional] 
**Trigger** | Pointer to [**WebhookTrigger**](WebhookTrigger.md) |  | [optional] 
**CreationTime** | Pointer to **time.Time** | The webhook creation time. | [optional] 
**Method** | Pointer to **NullableString** | The webhook method. | [optional] 
**Route** | Pointer to **NullableString** | The webhook route. | [optional] 
**RequestHeaders** | Pointer to **NullableString** | The webhook request headers. | [optional] 
**RequestPayload** | Pointer to **NullableString** | The webhook request payload. | [optional] 
**ResponseHeaders** | Pointer to **NullableString** | The webhook response headers. | [optional] 
**ResponsePayload** | Pointer to **NullableString** | The webhook response payload. | [optional] 
**Status** | Pointer to **int32** | The webhook status. | [optional] 
**Delivery** | Pointer to **NullableTime** | The webhook delivery time. | [optional] 

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


