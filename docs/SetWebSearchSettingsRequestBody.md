# SetWebSearchSettingsRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Indicates whether web search is enabled for AI chat sessions. | [optional] 
**Type** | Pointer to [**EngineType**](EngineType.md) |  | [optional] 
**Key** | Pointer to **NullableString** | The API key for the selected web search engine. Pass null to keep the existing key unchanged. | [optional] 

## Methods

### NewSetWebSearchSettingsRequestBody

`func NewSetWebSearchSettingsRequestBody() *SetWebSearchSettingsRequestBody`

NewSetWebSearchSettingsRequestBody instantiates a new SetWebSearchSettingsRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetWebSearchSettingsRequestBodyWithDefaults

`func NewSetWebSearchSettingsRequestBodyWithDefaults() *SetWebSearchSettingsRequestBody`

NewSetWebSearchSettingsRequestBodyWithDefaults instantiates a new SetWebSearchSettingsRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SetWebSearchSettingsRequestBody) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SetWebSearchSettingsRequestBody) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SetWebSearchSettingsRequestBody) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SetWebSearchSettingsRequestBody) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetType

`func (o *SetWebSearchSettingsRequestBody) GetType() EngineType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SetWebSearchSettingsRequestBody) GetTypeOk() (*EngineType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SetWebSearchSettingsRequestBody) SetType(v EngineType)`

SetType sets Type field to given value.

### HasType

`func (o *SetWebSearchSettingsRequestBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetKey

`func (o *SetWebSearchSettingsRequestBody) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SetWebSearchSettingsRequestBody) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SetWebSearchSettingsRequestBody) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SetWebSearchSettingsRequestBody) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *SetWebSearchSettingsRequestBody) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *SetWebSearchSettingsRequestBody) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


