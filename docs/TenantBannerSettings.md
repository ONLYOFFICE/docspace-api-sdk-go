# TenantBannerSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hidden** | Pointer to **bool** | The banners visibility flag. | [optional] 
**LastModified** | Pointer to **time.Time** | The timestamp indicating when the settings were last modified. | [optional] 

## Methods

### NewTenantBannerSettings

`func NewTenantBannerSettings() *TenantBannerSettings`

NewTenantBannerSettings instantiates a new TenantBannerSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantBannerSettingsWithDefaults

`func NewTenantBannerSettingsWithDefaults() *TenantBannerSettings`

NewTenantBannerSettingsWithDefaults instantiates a new TenantBannerSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHidden

`func (o *TenantBannerSettings) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *TenantBannerSettings) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *TenantBannerSettings) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *TenantBannerSettings) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetLastModified

`func (o *TenantBannerSettings) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *TenantBannerSettings) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *TenantBannerSettings) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *TenantBannerSettings) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


