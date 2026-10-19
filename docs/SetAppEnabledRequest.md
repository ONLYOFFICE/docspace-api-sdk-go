# SetAppEnabledRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Whether the application is available in this portal. Switching it off leaves its settings document stored, so  switching it back on restores the configuration it had; connected clients are told of the new state without a  reload. The field is mandatory: a body without it is rejected with 400 and leaves the application as it was. | 

## Methods

### NewSetAppEnabledRequest

`func NewSetAppEnabledRequest(enabled bool, ) *SetAppEnabledRequest`

NewSetAppEnabledRequest instantiates a new SetAppEnabledRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetAppEnabledRequestWithDefaults

`func NewSetAppEnabledRequestWithDefaults() *SetAppEnabledRequest`

NewSetAppEnabledRequestWithDefaults instantiates a new SetAppEnabledRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SetAppEnabledRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SetAppEnabledRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SetAppEnabledRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


