# TenantWalletServiceSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnabledServices** | Pointer to **[]int32** | The list of the enabled wallet services. | [optional] 
**LastModified** | Pointer to **time.Time** | The date and time when the wallet services settings were last modified. | [optional] 

## Methods

### NewTenantWalletServiceSettingsDto

`func NewTenantWalletServiceSettingsDto() *TenantWalletServiceSettingsDto`

NewTenantWalletServiceSettingsDto instantiates a new TenantWalletServiceSettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantWalletServiceSettingsDtoWithDefaults

`func NewTenantWalletServiceSettingsDtoWithDefaults() *TenantWalletServiceSettingsDto`

NewTenantWalletServiceSettingsDtoWithDefaults instantiates a new TenantWalletServiceSettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabledServices

`func (o *TenantWalletServiceSettingsDto) GetEnabledServices() []int32`

GetEnabledServices returns the EnabledServices field if non-nil, zero value otherwise.

### GetEnabledServicesOk

`func (o *TenantWalletServiceSettingsDto) GetEnabledServicesOk() (*[]int32, bool)`

GetEnabledServicesOk returns a tuple with the EnabledServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledServices

`func (o *TenantWalletServiceSettingsDto) SetEnabledServices(v []int32)`

SetEnabledServices sets EnabledServices field to given value.

### HasEnabledServices

`func (o *TenantWalletServiceSettingsDto) HasEnabledServices() bool`

HasEnabledServices returns a boolean if a field has been set.

### SetEnabledServicesNil

`func (o *TenantWalletServiceSettingsDto) SetEnabledServicesNil(b bool)`

 SetEnabledServicesNil sets the value for EnabledServices to be an explicit nil

### UnsetEnabledServices
`func (o *TenantWalletServiceSettingsDto) UnsetEnabledServices()`

UnsetEnabledServices ensures that no value is present for EnabledServices, not even an explicit nil
### GetLastModified

`func (o *TenantWalletServiceSettingsDto) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *TenantWalletServiceSettingsDto) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *TenantWalletServiceSettingsDto) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *TenantWalletServiceSettingsDto) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


