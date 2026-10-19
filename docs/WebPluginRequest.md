# WebPluginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether the plugin runs in this portal. Switching it on adds the domains its manifest declares to the portal  Content Security Policy and switching it off takes them away again; connected clients are told of the new  state without a reload. | [optional] 
**Settings** | **NullableString** | The configuration the plugin reads at run time, as a JSON document serialised into a string. Its shape is  defined by the plugin and not by the portal, which stores it encrypted for this portal alone. It replaces  whatever was stored rather than merging into it, so send `{}` when there is nothing to keep. | 

## Methods

### NewWebPluginRequest

`func NewWebPluginRequest(settings NullableString, ) *WebPluginRequest`

NewWebPluginRequest instantiates a new WebPluginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebPluginRequestWithDefaults

`func NewWebPluginRequestWithDefaults() *WebPluginRequest`

NewWebPluginRequestWithDefaults instantiates a new WebPluginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *WebPluginRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WebPluginRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WebPluginRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WebPluginRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSettings

`func (o *WebPluginRequest) GetSettings() string`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *WebPluginRequest) GetSettingsOk() (*string, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *WebPluginRequest) SetSettings(v string)`

SetSettings sets Settings field to given value.


### SetSettingsNil

`func (o *WebPluginRequest) SetSettingsNil(b bool)`

 SetSettingsNil sets the value for Settings to be an explicit nil

### UnsetSettings
`func (o *WebPluginRequest) UnsetSettings()`

UnsetSettings ensures that no value is present for Settings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


