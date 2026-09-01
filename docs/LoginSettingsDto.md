# LoginSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttemptCount** | **int32** | The maximum number of consecutive failed login attempts allowed before triggering account suspension. | 
**BlockTime** | **int32** | The duration (in minutes) for which an account remains suspended after exceeding maximum login attempts. | 
**CheckPeriod** | **int32** | The maximum time (in seconds) allowed for server to process and respond to login requests. | 
**IsDefault** | **bool** | Specifies whether the login settings are default or not. | 

## Methods

### NewLoginSettingsDto

`func NewLoginSettingsDto(attemptCount int32, blockTime int32, checkPeriod int32, isDefault bool, ) *LoginSettingsDto`

NewLoginSettingsDto instantiates a new LoginSettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginSettingsDtoWithDefaults

`func NewLoginSettingsDtoWithDefaults() *LoginSettingsDto`

NewLoginSettingsDtoWithDefaults instantiates a new LoginSettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttemptCount

`func (o *LoginSettingsDto) GetAttemptCount() int32`

GetAttemptCount returns the AttemptCount field if non-nil, zero value otherwise.

### GetAttemptCountOk

`func (o *LoginSettingsDto) GetAttemptCountOk() (*int32, bool)`

GetAttemptCountOk returns a tuple with the AttemptCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptCount

`func (o *LoginSettingsDto) SetAttemptCount(v int32)`

SetAttemptCount sets AttemptCount field to given value.


### GetBlockTime

`func (o *LoginSettingsDto) GetBlockTime() int32`

GetBlockTime returns the BlockTime field if non-nil, zero value otherwise.

### GetBlockTimeOk

`func (o *LoginSettingsDto) GetBlockTimeOk() (*int32, bool)`

GetBlockTimeOk returns a tuple with the BlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTime

`func (o *LoginSettingsDto) SetBlockTime(v int32)`

SetBlockTime sets BlockTime field to given value.


### GetCheckPeriod

`func (o *LoginSettingsDto) GetCheckPeriod() int32`

GetCheckPeriod returns the CheckPeriod field if non-nil, zero value otherwise.

### GetCheckPeriodOk

`func (o *LoginSettingsDto) GetCheckPeriodOk() (*int32, bool)`

GetCheckPeriodOk returns a tuple with the CheckPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPeriod

`func (o *LoginSettingsDto) SetCheckPeriod(v int32)`

SetCheckPeriod sets CheckPeriod field to given value.


### GetIsDefault

`func (o *LoginSettingsDto) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *LoginSettingsDto) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *LoginSettingsDto) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


