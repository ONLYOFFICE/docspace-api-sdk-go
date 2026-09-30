# CookieSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LifeTime** | **int32** | How long, in minutes, a session issued from now on remains valid. It is `1440` on a portal that has never  stored a limit, and that stored number is reported whether or not `enabled` puts it to use. | 
**Enabled** | **bool** | Whether the stored lifetime is applied at all. While it is `false` the number above is ignored and an  issued session is honoured for a year. | 

## Methods

### NewCookieSettingsDto

`func NewCookieSettingsDto(lifeTime int32, enabled bool, ) *CookieSettingsDto`

NewCookieSettingsDto instantiates a new CookieSettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCookieSettingsDtoWithDefaults

`func NewCookieSettingsDtoWithDefaults() *CookieSettingsDto`

NewCookieSettingsDtoWithDefaults instantiates a new CookieSettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLifeTime

`func (o *CookieSettingsDto) GetLifeTime() int32`

GetLifeTime returns the LifeTime field if non-nil, zero value otherwise.

### GetLifeTimeOk

`func (o *CookieSettingsDto) GetLifeTimeOk() (*int32, bool)`

GetLifeTimeOk returns a tuple with the LifeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifeTime

`func (o *CookieSettingsDto) SetLifeTime(v int32)`

SetLifeTime sets LifeTime field to given value.


### GetEnabled

`func (o *CookieSettingsDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CookieSettingsDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CookieSettingsDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


