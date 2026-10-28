# CookieSettingsRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LifeTime** | Pointer to **int32** | How long, in minutes, a session issued from now on remains valid. A value above 9999 is clamped to 9999  rather than refused, and 0 or less clears the number, which together with `enabled` leaves sessions that  never expire on their own. Any positive value invalidates every session issued before this call, the  caller's included, so the client has to keep the fresh cookie the response carries. | [optional] 
**Enabled** | Pointer to **bool** | Whether the stored lifetime is applied at all. While it is false the number is ignored and an issued session  is honoured for a year; while it is true the connections behind expired sessions are dropped as well. | [optional] 

## Methods

### NewCookieSettingsRequestDto

`func NewCookieSettingsRequestDto() *CookieSettingsRequestDto`

NewCookieSettingsRequestDto instantiates a new CookieSettingsRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCookieSettingsRequestDtoWithDefaults

`func NewCookieSettingsRequestDtoWithDefaults() *CookieSettingsRequestDto`

NewCookieSettingsRequestDtoWithDefaults instantiates a new CookieSettingsRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLifeTime

`func (o *CookieSettingsRequestDto) GetLifeTime() int32`

GetLifeTime returns the LifeTime field if non-nil, zero value otherwise.

### GetLifeTimeOk

`func (o *CookieSettingsRequestDto) GetLifeTimeOk() (*int32, bool)`

GetLifeTimeOk returns a tuple with the LifeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifeTime

`func (o *CookieSettingsRequestDto) SetLifeTime(v int32)`

SetLifeTime sets LifeTime field to given value.

### HasLifeTime

`func (o *CookieSettingsRequestDto) HasLifeTime() bool`

HasLifeTime returns a boolean if a field has been set.

### GetEnabled

`func (o *CookieSettingsRequestDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CookieSettingsRequestDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CookieSettingsRequestDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CookieSettingsRequestDto) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


