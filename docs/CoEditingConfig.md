# CoEditingConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Change** | Pointer to **bool** | Whether the user may switch between the two co-editing modes from the editor interface, or is held to the one  the portal preset. | [optional] 
**Fast** | Pointer to **bool** | Whether other participants see each change as it is typed. Left off, changes are exchanged only when a  participant saves, and the paragraph being edited is locked for the others meanwhile. | [optional] 
**Mode** | Pointer to [**CoEditingConfigMode**](CoEditingConfigMode.md) | The mode the two settings above amount to, as the editors name it. | [optional] 

## Methods

### NewCoEditingConfig

`func NewCoEditingConfig() *CoEditingConfig`

NewCoEditingConfig instantiates a new CoEditingConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoEditingConfigWithDefaults

`func NewCoEditingConfigWithDefaults() *CoEditingConfig`

NewCoEditingConfigWithDefaults instantiates a new CoEditingConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChange

`func (o *CoEditingConfig) GetChange() bool`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *CoEditingConfig) GetChangeOk() (*bool, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *CoEditingConfig) SetChange(v bool)`

SetChange sets Change field to given value.

### HasChange

`func (o *CoEditingConfig) HasChange() bool`

HasChange returns a boolean if a field has been set.

### GetFast

`func (o *CoEditingConfig) GetFast() bool`

GetFast returns the Fast field if non-nil, zero value otherwise.

### GetFastOk

`func (o *CoEditingConfig) GetFastOk() (*bool, bool)`

GetFastOk returns a tuple with the Fast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFast

`func (o *CoEditingConfig) SetFast(v bool)`

SetFast sets Fast field to given value.

### HasFast

`func (o *CoEditingConfig) HasFast() bool`

HasFast returns a boolean if a field has been set.

### GetMode

`func (o *CoEditingConfig) GetMode() CoEditingConfigMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *CoEditingConfig) GetModeOk() (*CoEditingConfigMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *CoEditingConfig) SetMode(v CoEditingConfigMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *CoEditingConfig) HasMode() bool`

HasMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


