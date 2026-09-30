# CurrentLicenseInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Trial** | **bool** | Whether the portal is on a trial rather than a paid subscription. A trial expires at `dueDate` and is not  extended by paying - a plan has to be bought instead. | 
**DueDate** | **time.Time** | The day the subscription runs out, with the time of day cut off. The largest value a date can hold means  it never runs out, which is how a free or unlimited plan is expressed. | 

## Methods

### NewCurrentLicenseInfo

`func NewCurrentLicenseInfo(trial bool, dueDate time.Time, ) *CurrentLicenseInfo`

NewCurrentLicenseInfo instantiates a new CurrentLicenseInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrentLicenseInfoWithDefaults

`func NewCurrentLicenseInfoWithDefaults() *CurrentLicenseInfo`

NewCurrentLicenseInfoWithDefaults instantiates a new CurrentLicenseInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrial

`func (o *CurrentLicenseInfo) GetTrial() bool`

GetTrial returns the Trial field if non-nil, zero value otherwise.

### GetTrialOk

`func (o *CurrentLicenseInfo) GetTrialOk() (*bool, bool)`

GetTrialOk returns a tuple with the Trial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrial

`func (o *CurrentLicenseInfo) SetTrial(v bool)`

SetTrial sets Trial field to given value.


### GetDueDate

`func (o *CurrentLicenseInfo) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *CurrentLicenseInfo) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *CurrentLicenseInfo) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


