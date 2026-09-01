# TenantAiAccessSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Specifies whether AI functionality is enabled for the tenant.  When set to `false`, all AI features (chat, agents, vectorization) are disabled tenant-wide. | [optional] 
**LastModified** | Pointer to **time.Time** | The timestamp indicating when the settings were last modified. | [optional] 

## Methods

### NewTenantAiAccessSettings

`func NewTenantAiAccessSettings() *TenantAiAccessSettings`

NewTenantAiAccessSettings instantiates a new TenantAiAccessSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantAiAccessSettingsWithDefaults

`func NewTenantAiAccessSettingsWithDefaults() *TenantAiAccessSettings`

NewTenantAiAccessSettingsWithDefaults instantiates a new TenantAiAccessSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *TenantAiAccessSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TenantAiAccessSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *TenantAiAccessSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *TenantAiAccessSettings) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetLastModified

`func (o *TenantAiAccessSettings) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *TenantAiAccessSettings) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *TenantAiAccessSettings) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *TenantAiAccessSettings) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


