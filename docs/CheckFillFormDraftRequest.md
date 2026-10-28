# CheckFillFormDraftRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **int32** | The revision of the form to open. Pass 0 for the current revision; a positive number addresses that entry of  the file history and is accepted only from a caller who may read the history, so a member who only has  fill-forms access must send 0. | 
**Action** | Pointer to **NullableString** | What the caller intends to do with the form. `view` asks for a read-only address and `embedded` for an address  to be shown inside a frame; both only resolve the address and leave the file untouched. Leave it out to enter  the filling flow, where the personal draft is created or reused. The value is matched case-insensitively, and  anything else behaves like an empty value. | [optional] 
**RequestView** | Pointer to **bool** | Whether the caller asked for a read-only address. The server derives it from `action` being `view` and ignores  any value sent with the request. | [optional] [readonly] 
**RequestEmbedded** | Pointer to **bool** | Whether the caller asked for an address to be shown inside a frame. The server derives it from `action` being  `embedded` and ignores any value sent with the request. | [optional] [readonly] 

## Methods

### NewCheckFillFormDraftRequest

`func NewCheckFillFormDraftRequest(version int32, ) *CheckFillFormDraftRequest`

NewCheckFillFormDraftRequest instantiates a new CheckFillFormDraftRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckFillFormDraftRequestWithDefaults

`func NewCheckFillFormDraftRequestWithDefaults() *CheckFillFormDraftRequest`

NewCheckFillFormDraftRequestWithDefaults instantiates a new CheckFillFormDraftRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *CheckFillFormDraftRequest) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CheckFillFormDraftRequest) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CheckFillFormDraftRequest) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetAction

`func (o *CheckFillFormDraftRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CheckFillFormDraftRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CheckFillFormDraftRequest) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *CheckFillFormDraftRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *CheckFillFormDraftRequest) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *CheckFillFormDraftRequest) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetRequestView

`func (o *CheckFillFormDraftRequest) GetRequestView() bool`

GetRequestView returns the RequestView field if non-nil, zero value otherwise.

### GetRequestViewOk

`func (o *CheckFillFormDraftRequest) GetRequestViewOk() (*bool, bool)`

GetRequestViewOk returns a tuple with the RequestView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestView

`func (o *CheckFillFormDraftRequest) SetRequestView(v bool)`

SetRequestView sets RequestView field to given value.

### HasRequestView

`func (o *CheckFillFormDraftRequest) HasRequestView() bool`

HasRequestView returns a boolean if a field has been set.

### GetRequestEmbedded

`func (o *CheckFillFormDraftRequest) GetRequestEmbedded() bool`

GetRequestEmbedded returns the RequestEmbedded field if non-nil, zero value otherwise.

### GetRequestEmbeddedOk

`func (o *CheckFillFormDraftRequest) GetRequestEmbeddedOk() (*bool, bool)`

GetRequestEmbeddedOk returns a tuple with the RequestEmbedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestEmbedded

`func (o *CheckFillFormDraftRequest) SetRequestEmbedded(v bool)`

SetRequestEmbedded sets RequestEmbedded field to given value.

### HasRequestEmbedded

`func (o *CheckFillFormDraftRequest) HasRequestEmbedded() bool`

HasRequestEmbedded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


