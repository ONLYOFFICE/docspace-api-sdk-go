# WebSearchSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Indicates whether web search is currently enabled. | [optional] 
**Type** | Pointer to [**EngineType**](EngineType.md) |  | [optional] 
**NeedReset** | Pointer to **bool** | Indicates whether the web search API key needs to be reconfigured. | [optional] 

## Methods

### NewWebSearchSettingsDto

`func NewWebSearchSettingsDto() *WebSearchSettingsDto`

NewWebSearchSettingsDto instantiates a new WebSearchSettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebSearchSettingsDtoWithDefaults

`func NewWebSearchSettingsDtoWithDefaults() *WebSearchSettingsDto`

NewWebSearchSettingsDtoWithDefaults instantiates a new WebSearchSettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *WebSearchSettingsDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WebSearchSettingsDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WebSearchSettingsDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WebSearchSettingsDto) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetType

`func (o *WebSearchSettingsDto) GetType() EngineType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebSearchSettingsDto) GetTypeOk() (*EngineType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebSearchSettingsDto) SetType(v EngineType)`

SetType sets Type field to given value.

### HasType

`func (o *WebSearchSettingsDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNeedReset

`func (o *WebSearchSettingsDto) GetNeedReset() bool`

GetNeedReset returns the NeedReset field if non-nil, zero value otherwise.

### GetNeedResetOk

`func (o *WebSearchSettingsDto) GetNeedResetOk() (*bool, bool)`

GetNeedResetOk returns a tuple with the NeedReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedReset

`func (o *WebSearchSettingsDto) SetNeedReset(v bool)`

SetNeedReset sets NeedReset field to given value.

### HasNeedReset

`func (o *WebSearchSettingsDto) HasNeedReset() bool`

HasNeedReset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


