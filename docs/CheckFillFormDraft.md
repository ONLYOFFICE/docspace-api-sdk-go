# CheckFillFormDraft

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **int32** | The revision of the form to open. Pass 0 for the current revision; a positive number addresses that entry of  the file history and is accepted only from a caller who may read the history, so a member who only has  fill-forms access must send 0. | 
**Action** | Pointer to **NullableString** | What the caller intends to do with the form. `view` asks for a read-only address and `embedded` for an address  to be shown inside a frame; both only resolve the address and leave the file untouched. Leave it out to enter  the filling flow, where the personal draft is created or reused. The value is matched case-insensitively, and  anything else behaves like an empty value. | [optional] 
**RequestView** | Pointer to **bool** | Whether the caller asked for a read-only address. The server derives it from `action` being `view` and ignores  any value sent with the request. | [optional] [readonly] 
**RequestEmbedded** | Pointer to **bool** | Whether the caller asked for an address to be shown inside a frame. The server derives it from `action` being  `embedded` and ignores any value sent with the request. | [optional] [readonly] 

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


