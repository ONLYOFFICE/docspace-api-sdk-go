# ActionLinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**ActionLinkActionRequest**](ActionLinkActionRequest.md) | The anchor itself. It is passed on to the editor unchanged, so it has to be the value the editor produced for  the comment or the mention it points at. | [optional] 

## Methods

### NewActionLinkRequest

`func NewActionLinkRequest() *ActionLinkRequest`

NewActionLinkRequest instantiates a new ActionLinkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionLinkRequestWithDefaults

`func NewActionLinkRequestWithDefaults() *ActionLinkRequest`

NewActionLinkRequestWithDefaults instantiates a new ActionLinkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ActionLinkRequest) GetAction() ActionLinkActionRequest`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ActionLinkRequest) GetActionOk() (*ActionLinkActionRequest, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ActionLinkRequest) SetAction(v ActionLinkActionRequest)`

SetAction sets Action field to given value.

### HasAction

`func (o *ActionLinkRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


