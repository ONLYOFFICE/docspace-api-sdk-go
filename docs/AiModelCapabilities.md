# AiModelCapabilities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vision** | Pointer to **bool** | Indicates whether the model supports image and vision input. | [optional] 
**ToolCalling** | Pointer to **bool** | Indicates whether the model supports tool (function) calling. | [optional] 
**Thinking** | Pointer to **bool** | Indicates whether the model supports extended thinking and reasoning. | [optional] 

## Methods

### NewAiModelCapabilities

`func NewAiModelCapabilities() *AiModelCapabilities`

NewAiModelCapabilities instantiates a new AiModelCapabilities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiModelCapabilitiesWithDefaults

`func NewAiModelCapabilitiesWithDefaults() *AiModelCapabilities`

NewAiModelCapabilitiesWithDefaults instantiates a new AiModelCapabilities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVision

`func (o *AiModelCapabilities) GetVision() bool`

GetVision returns the Vision field if non-nil, zero value otherwise.

### GetVisionOk

`func (o *AiModelCapabilities) GetVisionOk() (*bool, bool)`

GetVisionOk returns a tuple with the Vision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVision

`func (o *AiModelCapabilities) SetVision(v bool)`

SetVision sets Vision field to given value.

### HasVision

`func (o *AiModelCapabilities) HasVision() bool`

HasVision returns a boolean if a field has been set.

### GetToolCalling

`func (o *AiModelCapabilities) GetToolCalling() bool`

GetToolCalling returns the ToolCalling field if non-nil, zero value otherwise.

### GetToolCallingOk

`func (o *AiModelCapabilities) GetToolCallingOk() (*bool, bool)`

GetToolCallingOk returns a tuple with the ToolCalling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCalling

`func (o *AiModelCapabilities) SetToolCalling(v bool)`

SetToolCalling sets ToolCalling field to given value.

### HasToolCalling

`func (o *AiModelCapabilities) HasToolCalling() bool`

HasToolCalling returns a boolean if a field has been set.

### GetThinking

`func (o *AiModelCapabilities) GetThinking() bool`

GetThinking returns the Thinking field if non-nil, zero value otherwise.

### GetThinkingOk

`func (o *AiModelCapabilities) GetThinkingOk() (*bool, bool)`

GetThinkingOk returns a tuple with the Thinking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThinking

`func (o *AiModelCapabilities) SetThinking(v bool)`

SetThinking sets Thinking field to given value.

### HasThinking

`func (o *AiModelCapabilities) HasThinking() bool`

HasThinking returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


