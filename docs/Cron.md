# Cron

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Period** | Pointer to [**BackupPeriod**](BackupPeriod.md) | The backup period type. | [optional] 
**Hour** | Pointer to **int32** | The time of the day to start the backup process. | [optional] 
**Day** | Pointer to **NullableInt32** | The day of the week to start the backup process. | [optional] 

## Methods

### NewCron

`func NewCron() *Cron`

NewCron instantiates a new Cron object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCronWithDefaults

`func NewCronWithDefaults() *Cron`

NewCronWithDefaults instantiates a new Cron object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriod

`func (o *Cron) GetPeriod() BackupPeriod`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *Cron) GetPeriodOk() (*BackupPeriod, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *Cron) SetPeriod(v BackupPeriod)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *Cron) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetHour

`func (o *Cron) GetHour() int32`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *Cron) GetHourOk() (*int32, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *Cron) SetHour(v int32)`

SetHour sets Hour field to given value.

### HasHour

`func (o *Cron) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetDay

`func (o *Cron) GetDay() int32`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *Cron) GetDayOk() (*int32, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *Cron) SetDay(v int32)`

SetDay sets Day field to given value.

### HasDay

`func (o *Cron) HasDay() bool`

HasDay returns a boolean if a field has been set.

### SetDayNil

`func (o *Cron) SetDayNil(b bool)`

 SetDayNil sets the value for Day to be an explicit nil

### UnsetDay
`func (o *Cron) UnsetDay()`

UnsetDay ensures that no value is present for Day, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


