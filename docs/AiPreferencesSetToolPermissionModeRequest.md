# AiPreferencesSetToolPermissionModeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | [**AiChatToolPermissionMode**](AiChatToolPermissionMode.md) | New tool permission mode. | 
**EntityId** | Pointer to **string** |  | [optional] 

## Methods

### NewAiPreferencesSetToolPermissionModeRequest

`func NewAiPreferencesSetToolPermissionModeRequest(value AiChatToolPermissionMode, ) *AiPreferencesSetToolPermissionModeRequest`

NewAiPreferencesSetToolPermissionModeRequest instantiates a new AiPreferencesSetToolPermissionModeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiPreferencesSetToolPermissionModeRequestWithDefaults

`func NewAiPreferencesSetToolPermissionModeRequestWithDefaults() *AiPreferencesSetToolPermissionModeRequest`

NewAiPreferencesSetToolPermissionModeRequestWithDefaults instantiates a new AiPreferencesSetToolPermissionModeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *AiPreferencesSetToolPermissionModeRequest) GetValue() AiChatToolPermissionMode`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AiPreferencesSetToolPermissionModeRequest) GetValueOk() (*AiChatToolPermissionMode, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AiPreferencesSetToolPermissionModeRequest) SetValue(v AiChatToolPermissionMode)`

SetValue sets Value field to given value.


### GetEntityId

`func (o *AiPreferencesSetToolPermissionModeRequest) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *AiPreferencesSetToolPermissionModeRequest) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *AiPreferencesSetToolPermissionModeRequest) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *AiPreferencesSetToolPermissionModeRequest) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


