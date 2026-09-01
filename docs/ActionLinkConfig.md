# ActionLinkConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**ActionConfig**](ActionConfig.md) | The information about the action in the document that will be scrolled to. | [optional] 

## Methods

### NewActionLinkConfig

`func NewActionLinkConfig() *ActionLinkConfig`

NewActionLinkConfig instantiates a new ActionLinkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionLinkConfigWithDefaults

`func NewActionLinkConfigWithDefaults() *ActionLinkConfig`

NewActionLinkConfigWithDefaults instantiates a new ActionLinkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ActionLinkConfig) GetAction() ActionConfig`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ActionLinkConfig) GetActionOk() (*ActionConfig, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ActionLinkConfig) SetAction(v ActionConfig)`

SetAction sets Action field to given value.

### HasAction

`func (o *ActionLinkConfig) HasAction() bool`

HasAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


