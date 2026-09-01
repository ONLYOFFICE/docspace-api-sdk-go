# IPRestrictionsSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** | Specifies if the IP restrictions are enabled or not. | [optional] 
**LastModified** | Pointer to **time.Time** | The date and time when the settings were last modified. | [optional] 

## Methods

### NewIPRestrictionsSettings

`func NewIPRestrictionsSettings() *IPRestrictionsSettings`

NewIPRestrictionsSettings instantiates a new IPRestrictionsSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPRestrictionsSettingsWithDefaults

`func NewIPRestrictionsSettingsWithDefaults() *IPRestrictionsSettings`

NewIPRestrictionsSettingsWithDefaults instantiates a new IPRestrictionsSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *IPRestrictionsSettings) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *IPRestrictionsSettings) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *IPRestrictionsSettings) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *IPRestrictionsSettings) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetLastModified

`func (o *IPRestrictionsSettings) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *IPRestrictionsSettings) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *IPRestrictionsSettings) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *IPRestrictionsSettings) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


