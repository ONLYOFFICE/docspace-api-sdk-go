# TenantDevToolsAccessSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LimitedAccessForUsers** | Pointer to **bool** | Specifies if the Developer Tools access are limited for users or not. | [optional] 
**LastModified** | Pointer to **time.Time** | The timestamp indicating when the settings were last modified. | [optional] 

## Methods

### NewTenantDevToolsAccessSettingsDto

`func NewTenantDevToolsAccessSettingsDto() *TenantDevToolsAccessSettingsDto`

NewTenantDevToolsAccessSettingsDto instantiates a new TenantDevToolsAccessSettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantDevToolsAccessSettingsDtoWithDefaults

`func NewTenantDevToolsAccessSettingsDtoWithDefaults() *TenantDevToolsAccessSettingsDto`

NewTenantDevToolsAccessSettingsDtoWithDefaults instantiates a new TenantDevToolsAccessSettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimitedAccessForUsers

`func (o *TenantDevToolsAccessSettingsDto) GetLimitedAccessForUsers() bool`

GetLimitedAccessForUsers returns the LimitedAccessForUsers field if non-nil, zero value otherwise.

### GetLimitedAccessForUsersOk

`func (o *TenantDevToolsAccessSettingsDto) GetLimitedAccessForUsersOk() (*bool, bool)`

GetLimitedAccessForUsersOk returns a tuple with the LimitedAccessForUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitedAccessForUsers

`func (o *TenantDevToolsAccessSettingsDto) SetLimitedAccessForUsers(v bool)`

SetLimitedAccessForUsers sets LimitedAccessForUsers field to given value.

### HasLimitedAccessForUsers

`func (o *TenantDevToolsAccessSettingsDto) HasLimitedAccessForUsers() bool`

HasLimitedAccessForUsers returns a boolean if a field has been set.

### GetLastModified

`func (o *TenantDevToolsAccessSettingsDto) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *TenantDevToolsAccessSettingsDto) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *TenantDevToolsAccessSettingsDto) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *TenantDevToolsAccessSettingsDto) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


