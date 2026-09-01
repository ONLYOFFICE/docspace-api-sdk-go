# CheckFillFormDraft

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **int32** | The file version of the form draft. | 
**Action** | Pointer to **NullableString** | The action with the form draft. | [optional] 
**RequestView** | Pointer to **bool** | Specifies whether to request the form for viewing or not. | [optional] [readonly] 
**RequestEmbedded** | Pointer to **bool** | Specifies whether to request an embedded form or not. | [optional] [readonly] 

## Methods

### NewCheckFillFormDraft

`func NewCheckFillFormDraft(version int32, ) *CheckFillFormDraft`

NewCheckFillFormDraft instantiates a new CheckFillFormDraft object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckFillFormDraftWithDefaults

`func NewCheckFillFormDraftWithDefaults() *CheckFillFormDraft`

NewCheckFillFormDraftWithDefaults instantiates a new CheckFillFormDraft object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *CheckFillFormDraft) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CheckFillFormDraft) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CheckFillFormDraft) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetAction

`func (o *CheckFillFormDraft) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CheckFillFormDraft) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CheckFillFormDraft) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *CheckFillFormDraft) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *CheckFillFormDraft) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *CheckFillFormDraft) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetRequestView

`func (o *CheckFillFormDraft) GetRequestView() bool`

GetRequestView returns the RequestView field if non-nil, zero value otherwise.

### GetRequestViewOk

`func (o *CheckFillFormDraft) GetRequestViewOk() (*bool, bool)`

GetRequestViewOk returns a tuple with the RequestView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestView

`func (o *CheckFillFormDraft) SetRequestView(v bool)`

SetRequestView sets RequestView field to given value.

### HasRequestView

`func (o *CheckFillFormDraft) HasRequestView() bool`

HasRequestView returns a boolean if a field has been set.

### GetRequestEmbedded

`func (o *CheckFillFormDraft) GetRequestEmbedded() bool`

GetRequestEmbedded returns the RequestEmbedded field if non-nil, zero value otherwise.

### GetRequestEmbeddedOk

`func (o *CheckFillFormDraft) GetRequestEmbeddedOk() (*bool, bool)`

GetRequestEmbeddedOk returns a tuple with the RequestEmbedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestEmbedded

`func (o *CheckFillFormDraft) SetRequestEmbedded(v bool)`

SetRequestEmbedded sets RequestEmbedded field to given value.

### HasRequestEmbedded

`func (o *CheckFillFormDraft) HasRequestEmbedded() bool`

HasRequestEmbedded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


