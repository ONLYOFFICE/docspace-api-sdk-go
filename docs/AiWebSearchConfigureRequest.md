# AiWebSearchConfigureRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | [**AiWebSearchConfig**](AiWebSearchConfig.md) |  | 
**EntityId** | Pointer to **string** |  | [optional] 

## Methods

### NewAiWebSearchConfigureRequest

`func NewAiWebSearchConfigureRequest(config AiWebSearchConfig, ) *AiWebSearchConfigureRequest`

NewAiWebSearchConfigureRequest instantiates a new AiWebSearchConfigureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiWebSearchConfigureRequestWithDefaults

`func NewAiWebSearchConfigureRequestWithDefaults() *AiWebSearchConfigureRequest`

NewAiWebSearchConfigureRequestWithDefaults instantiates a new AiWebSearchConfigureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *AiWebSearchConfigureRequest) GetConfig() AiWebSearchConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AiWebSearchConfigureRequest) GetConfigOk() (*AiWebSearchConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AiWebSearchConfigureRequest) SetConfig(v AiWebSearchConfig)`

SetConfig sets Config field to given value.


### GetEntityId

`func (o *AiWebSearchConfigureRequest) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *AiWebSearchConfigureRequest) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *AiWebSearchConfigureRequest) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *AiWebSearchConfigureRequest) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


