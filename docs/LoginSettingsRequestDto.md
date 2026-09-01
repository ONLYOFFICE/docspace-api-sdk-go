# LoginSettingsRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttemptCount** | Pointer to **int32** | The maximum number of consecutive failed login attempts allowed before triggering account suspension. | [optional] 
**BlockTime** | Pointer to **int32** | The duration (in minutes) for which an account remains suspended after exceeding maximum login attempts. | [optional] 
**CheckPeriod** | Pointer to **int32** | The maximum time (in seconds) allowed for server to process and respond to login requests. | [optional] 

## Methods

### NewLoginSettingsRequestDto

`func NewLoginSettingsRequestDto() *LoginSettingsRequestDto`

NewLoginSettingsRequestDto instantiates a new LoginSettingsRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginSettingsRequestDtoWithDefaults

`func NewLoginSettingsRequestDtoWithDefaults() *LoginSettingsRequestDto`

NewLoginSettingsRequestDtoWithDefaults instantiates a new LoginSettingsRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttemptCount

`func (o *LoginSettingsRequestDto) GetAttemptCount() int32`

GetAttemptCount returns the AttemptCount field if non-nil, zero value otherwise.

### GetAttemptCountOk

`func (o *LoginSettingsRequestDto) GetAttemptCountOk() (*int32, bool)`

GetAttemptCountOk returns a tuple with the AttemptCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptCount

`func (o *LoginSettingsRequestDto) SetAttemptCount(v int32)`

SetAttemptCount sets AttemptCount field to given value.

### HasAttemptCount

`func (o *LoginSettingsRequestDto) HasAttemptCount() bool`

HasAttemptCount returns a boolean if a field has been set.

### GetBlockTime

`func (o *LoginSettingsRequestDto) GetBlockTime() int32`

GetBlockTime returns the BlockTime field if non-nil, zero value otherwise.

### GetBlockTimeOk

`func (o *LoginSettingsRequestDto) GetBlockTimeOk() (*int32, bool)`

GetBlockTimeOk returns a tuple with the BlockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTime

`func (o *LoginSettingsRequestDto) SetBlockTime(v int32)`

SetBlockTime sets BlockTime field to given value.

### HasBlockTime

`func (o *LoginSettingsRequestDto) HasBlockTime() bool`

HasBlockTime returns a boolean if a field has been set.

### GetCheckPeriod

`func (o *LoginSettingsRequestDto) GetCheckPeriod() int32`

GetCheckPeriod returns the CheckPeriod field if non-nil, zero value otherwise.

### GetCheckPeriodOk

`func (o *LoginSettingsRequestDto) GetCheckPeriodOk() (*int32, bool)`

GetCheckPeriodOk returns a tuple with the CheckPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPeriod

`func (o *LoginSettingsRequestDto) SetCheckPeriod(v int32)`

SetCheckPeriod sets CheckPeriod field to given value.

### HasCheckPeriod

`func (o *LoginSettingsRequestDto) HasCheckPeriod() bool`

HasCheckPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


