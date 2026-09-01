# CookieSettingsRequestsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LifeTime** | Pointer to **int32** | The cookie lifetime in minutes. | [optional] 
**Enabled** | Pointer to **bool** | Specifies whether the cookie settings are enabled or disabled. | [optional] 

## Methods

### NewCookieSettingsRequestsDto

`func NewCookieSettingsRequestsDto() *CookieSettingsRequestsDto`

NewCookieSettingsRequestsDto instantiates a new CookieSettingsRequestsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCookieSettingsRequestsDtoWithDefaults

`func NewCookieSettingsRequestsDtoWithDefaults() *CookieSettingsRequestsDto`

NewCookieSettingsRequestsDtoWithDefaults instantiates a new CookieSettingsRequestsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLifeTime

`func (o *CookieSettingsRequestsDto) GetLifeTime() int32`

GetLifeTime returns the LifeTime field if non-nil, zero value otherwise.

### GetLifeTimeOk

`func (o *CookieSettingsRequestsDto) GetLifeTimeOk() (*int32, bool)`

GetLifeTimeOk returns a tuple with the LifeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifeTime

`func (o *CookieSettingsRequestsDto) SetLifeTime(v int32)`

SetLifeTime sets LifeTime field to given value.

### HasLifeTime

`func (o *CookieSettingsRequestsDto) HasLifeTime() bool`

HasLifeTime returns a boolean if a field has been set.

### GetEnabled

`func (o *CookieSettingsRequestsDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CookieSettingsRequestsDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CookieSettingsRequestsDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CookieSettingsRequestsDto) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


