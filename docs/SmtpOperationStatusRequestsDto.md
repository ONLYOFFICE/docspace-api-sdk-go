# SmtpOperationStatusRequestsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Completed** | Pointer to **bool** | Whether the job has finished. This is the field to poll; the first answer that reports it true also discards  the job, so read `error` out of that same answer rather than calling again. | [optional] 
**Id** | Pointer to **NullableString** | The identifier of the queued job. A portal only ever has one test job at a time, so it names the run rather  than selecting among several. | [optional] 
**Error** | Pointer to **NullableString** | Why the test failed. It stays empty while the job runs and also once the relay has accepted the message, so  an empty value on a finished job is what success looks like; an unreachable relay is reported here after a  30-second connection timeout rather than as a failed request. | [optional] 
**Status** | Pointer to **NullableString** | The step the job has reached, in words - `Connect to host` or `Send test message`, for instance. It is meant  to be shown to a person and is not a fixed set of values to branch on. | [optional] 
**Percents** | Pointer to **int32** | How far the job has got, as a percentage climbing to 100. Reaching 100 says the job ran to the end, not that  the message was accepted - that is what an empty `error` says. | [optional] 

## Methods

### NewSmtpOperationStatusRequestsDto

`func NewSmtpOperationStatusRequestsDto() *SmtpOperationStatusRequestsDto`

NewSmtpOperationStatusRequestsDto instantiates a new SmtpOperationStatusRequestsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmtpOperationStatusRequestsDtoWithDefaults

`func NewSmtpOperationStatusRequestsDtoWithDefaults() *SmtpOperationStatusRequestsDto`

NewSmtpOperationStatusRequestsDtoWithDefaults instantiates a new SmtpOperationStatusRequestsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompleted

`func (o *SmtpOperationStatusRequestsDto) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *SmtpOperationStatusRequestsDto) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *SmtpOperationStatusRequestsDto) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *SmtpOperationStatusRequestsDto) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetId

`func (o *SmtpOperationStatusRequestsDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SmtpOperationStatusRequestsDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SmtpOperationStatusRequestsDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SmtpOperationStatusRequestsDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SmtpOperationStatusRequestsDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SmtpOperationStatusRequestsDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetError

`func (o *SmtpOperationStatusRequestsDto) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SmtpOperationStatusRequestsDto) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SmtpOperationStatusRequestsDto) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *SmtpOperationStatusRequestsDto) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *SmtpOperationStatusRequestsDto) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *SmtpOperationStatusRequestsDto) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetStatus

`func (o *SmtpOperationStatusRequestsDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SmtpOperationStatusRequestsDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SmtpOperationStatusRequestsDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SmtpOperationStatusRequestsDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *SmtpOperationStatusRequestsDto) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *SmtpOperationStatusRequestsDto) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetPercents

`func (o *SmtpOperationStatusRequestsDto) GetPercents() int32`

GetPercents returns the Percents field if non-nil, zero value otherwise.

### GetPercentsOk

`func (o *SmtpOperationStatusRequestsDto) GetPercentsOk() (*int32, bool)`

GetPercentsOk returns a tuple with the Percents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercents

`func (o *SmtpOperationStatusRequestsDto) SetPercents(v int32)`

SetPercents sets Percents field to given value.

### HasPercents

`func (o *SmtpOperationStatusRequestsDto) HasPercents() bool`

HasPercents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


