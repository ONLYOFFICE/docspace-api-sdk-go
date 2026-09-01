# TenantDeepLinkSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HandlingMode** | Pointer to [**DeepLinkHandlingMode**](DeepLinkHandlingMode.md) | The deep link handling mode. | [optional] 
**LastModified** | Pointer to **time.Time** | The timestamp indicating when the settings were last modified. | [optional] 

## Methods

### NewTenantDeepLinkSettings

`func NewTenantDeepLinkSettings() *TenantDeepLinkSettings`

NewTenantDeepLinkSettings instantiates a new TenantDeepLinkSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantDeepLinkSettingsWithDefaults

`func NewTenantDeepLinkSettingsWithDefaults() *TenantDeepLinkSettings`

NewTenantDeepLinkSettingsWithDefaults instantiates a new TenantDeepLinkSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandlingMode

`func (o *TenantDeepLinkSettings) GetHandlingMode() DeepLinkHandlingMode`

GetHandlingMode returns the HandlingMode field if non-nil, zero value otherwise.

### GetHandlingModeOk

`func (o *TenantDeepLinkSettings) GetHandlingModeOk() (*DeepLinkHandlingMode, bool)`

GetHandlingModeOk returns a tuple with the HandlingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlingMode

`func (o *TenantDeepLinkSettings) SetHandlingMode(v DeepLinkHandlingMode)`

SetHandlingMode sets HandlingMode field to given value.

### HasHandlingMode

`func (o *TenantDeepLinkSettings) HasHandlingMode() bool`

HasHandlingMode returns a boolean if a field has been set.

### GetLastModified

`func (o *TenantDeepLinkSettings) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *TenantDeepLinkSettings) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *TenantDeepLinkSettings) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *TenantDeepLinkSettings) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


