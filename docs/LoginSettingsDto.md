# LoginSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttemptCount** | **int32** | How many failed attempts inside one window are tolerated before the offender is blocked. Attempts are  counted per user name and client address together, so one member being blocked leaves the rest of the  portal signing in normally. | 
**BlockTime** | **int32** | How long, in seconds, a blocked user name and address pair stays refused. While the block lasts the  sign-in is refused even once the password is correct. | 
**CheckPeriod** | **int32** | The length, in seconds, of the rolling window the failures are counted over. It is not a request timeout: a  wider window makes the same `attemptCount` stricter, because failures further apart still add up. | 
**IsDefault** | **bool** | Whether the three numbers above still match the ones the installation ships with. It turns `false` as soon  as any of them is saved differently, and `true` again after  `DELETE api/2.0/settings/security/loginsettings`. | 

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


