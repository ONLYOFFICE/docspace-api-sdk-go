# ConnectionTestResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Error** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewConnectionTestResult

`func NewConnectionTestResult() *ConnectionTestResult`

NewConnectionTestResult instantiates a new ConnectionTestResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionTestResultWithDefaults

`func NewConnectionTestResultWithDefaults() *ConnectionTestResult`

NewConnectionTestResultWithDefaults instantiates a new ConnectionTestResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ConnectionTestResult) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ConnectionTestResult) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ConnectionTestResult) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ConnectionTestResult) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetError

`func (o *ConnectionTestResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ConnectionTestResult) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ConnectionTestResult) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ConnectionTestResult) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *ConnectionTestResult) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *ConnectionTestResult) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


