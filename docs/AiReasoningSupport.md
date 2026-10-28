# AiReasoningSupport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Thinks** | **bool** | Whether the model can think at all. False hides the whole control. | 
**CanDisable** | **bool** | Whether `off` really turns thinking off. False means the model thinks always and off only drops to its lowest depth (or leaves the default depth, where there is no knob). | 
**Depths** | [**[]AiReasoningDepth**](AiReasoningDepth.md) | Depths the model distinguishes, lowest first. Empty when thinking is an on/off switch with no depth (or the model doesn't think). A level not listed is clamped to the nearest one — see `clampReasoningLevel`. | 
**DefaultDepth** | Pointer to [**AiReasoningDepth**](AiReasoningDepth.md) | The depth the model runs at when nothing asks for one — what a stored `off` means on a model that cannot be switched off. Known only where a catalogue reports it (OpenRouter's `default_effort`); otherwise `DEFAULT_REASONING_LEVEL` clamped to `depths` is assumed. | [optional] 

## Methods

### NewAiReasoningSupport

`func NewAiReasoningSupport(thinks bool, canDisable bool, depths []AiReasoningDepth, ) *AiReasoningSupport`

NewAiReasoningSupport instantiates a new AiReasoningSupport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiReasoningSupportWithDefaults

`func NewAiReasoningSupportWithDefaults() *AiReasoningSupport`

NewAiReasoningSupportWithDefaults instantiates a new AiReasoningSupport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThinks

`func (o *AiReasoningSupport) GetThinks() bool`

GetThinks returns the Thinks field if non-nil, zero value otherwise.

### GetThinksOk

`func (o *AiReasoningSupport) GetThinksOk() (*bool, bool)`

GetThinksOk returns a tuple with the Thinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThinks

`func (o *AiReasoningSupport) SetThinks(v bool)`

SetThinks sets Thinks field to given value.


### GetCanDisable

`func (o *AiReasoningSupport) GetCanDisable() bool`

GetCanDisable returns the CanDisable field if non-nil, zero value otherwise.

### GetCanDisableOk

`func (o *AiReasoningSupport) GetCanDisableOk() (*bool, bool)`

GetCanDisableOk returns a tuple with the CanDisable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDisable

`func (o *AiReasoningSupport) SetCanDisable(v bool)`

SetCanDisable sets CanDisable field to given value.


### GetDepths

`func (o *AiReasoningSupport) GetDepths() []AiReasoningDepth`

GetDepths returns the Depths field if non-nil, zero value otherwise.

### GetDepthsOk

`func (o *AiReasoningSupport) GetDepthsOk() (*[]AiReasoningDepth, bool)`

GetDepthsOk returns a tuple with the Depths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepths

`func (o *AiReasoningSupport) SetDepths(v []AiReasoningDepth)`

SetDepths sets Depths field to given value.


### GetDefaultDepth

`func (o *AiReasoningSupport) GetDefaultDepth() AiReasoningDepth`

GetDefaultDepth returns the DefaultDepth field if non-nil, zero value otherwise.

### GetDefaultDepthOk

`func (o *AiReasoningSupport) GetDefaultDepthOk() (*AiReasoningDepth, bool)`

GetDefaultDepthOk returns a tuple with the DefaultDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDepth

`func (o *AiReasoningSupport) SetDefaultDepth(v AiReasoningDepth)`

SetDefaultDepth sets DefaultDepth field to given value.

### HasDefaultDepth

`func (o *AiReasoningSupport) HasDefaultDepth() bool`

HasDefaultDepth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


