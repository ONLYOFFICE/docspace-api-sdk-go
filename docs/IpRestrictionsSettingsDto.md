# IpRestrictionsSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** | Specifies if the IP restrictions are enabled or not. | [optional] 
**LastModified** | Pointer to **time.Time** | The date and time when the settings were last modified. | [optional] 

## Methods

### NewIpRestrictionsSettingsDto

`func NewIpRestrictionsSettingsDto() *IpRestrictionsSettingsDto`

NewIpRestrictionsSettingsDto instantiates a new IpRestrictionsSettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpRestrictionsSettingsDtoWithDefaults

`func NewIpRestrictionsSettingsDtoWithDefaults() *IpRestrictionsSettingsDto`

NewIpRestrictionsSettingsDtoWithDefaults instantiates a new IpRestrictionsSettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *IpRestrictionsSettingsDto) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *IpRestrictionsSettingsDto) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *IpRestrictionsSettingsDto) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *IpRestrictionsSettingsDto) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetLastModified

`func (o *IpRestrictionsSettingsDto) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *IpRestrictionsSettingsDto) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *IpRestrictionsSettingsDto) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *IpRestrictionsSettingsDto) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


