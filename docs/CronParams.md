# CronParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Period** | Pointer to [**BackupPeriod**](BackupPeriod.md) | The backup period type. | [optional] 
**Hour** | Pointer to **int32** | The time of the day to start the backup process. | [optional] 
**Day** | Pointer to **int32** | The day of the week to start the backup process. | [optional] 

## Methods

### NewCronParams

`func NewCronParams() *CronParams`

NewCronParams instantiates a new CronParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCronParamsWithDefaults

`func NewCronParamsWithDefaults() *CronParams`

NewCronParamsWithDefaults instantiates a new CronParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriod

`func (o *CronParams) GetPeriod() BackupPeriod`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *CronParams) GetPeriodOk() (*BackupPeriod, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *CronParams) SetPeriod(v BackupPeriod)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *CronParams) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetHour

`func (o *CronParams) GetHour() int32`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *CronParams) GetHourOk() (*int32, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *CronParams) SetHour(v int32)`

SetHour sets Hour field to given value.

### HasHour

`func (o *CronParams) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetDay

`func (o *CronParams) GetDay() int32`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *CronParams) GetDayOk() (*int32, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *CronParams) SetDay(v int32)`

SetDay sets Day field to given value.

### HasDay

`func (o *CronParams) HasDay() bool`

HasDay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


