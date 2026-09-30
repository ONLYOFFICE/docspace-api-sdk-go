# AiPreferencesSetReasoningLevelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | [**AiAiReasoningLevel**](AiAiReasoningLevel.md) | New extended-thinking depth; `off` turns deep mode off. | 
**EntityId** | Pointer to **string** |  | [optional] 

## Methods

### NewAiPreferencesSetReasoningLevelRequest

`func NewAiPreferencesSetReasoningLevelRequest(value AiAiReasoningLevel, ) *AiPreferencesSetReasoningLevelRequest`

NewAiPreferencesSetReasoningLevelRequest instantiates a new AiPreferencesSetReasoningLevelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiPreferencesSetReasoningLevelRequestWithDefaults

`func NewAiPreferencesSetReasoningLevelRequestWithDefaults() *AiPreferencesSetReasoningLevelRequest`

NewAiPreferencesSetReasoningLevelRequestWithDefaults instantiates a new AiPreferencesSetReasoningLevelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *AiPreferencesSetReasoningLevelRequest) GetValue() AiAiReasoningLevel`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AiPreferencesSetReasoningLevelRequest) GetValueOk() (*AiAiReasoningLevel, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AiPreferencesSetReasoningLevelRequest) SetValue(v AiAiReasoningLevel)`

SetValue sets Value field to given value.


### GetEntityId

`func (o *AiPreferencesSetReasoningLevelRequest) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *AiPreferencesSetReasoningLevelRequest) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *AiPreferencesSetReasoningLevelRequest) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *AiPreferencesSetReasoningLevelRequest) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


