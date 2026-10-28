# ConnectionTestResultDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** | Specifies whether the connection to the database succeeded. | [optional] 
**Error** | Pointer to **NullableString** | The reason the connection failed, or null when it succeeded. | [optional] 

## Methods

### NewConnectionTestResultDto

`func NewConnectionTestResultDto() *ConnectionTestResultDto`

NewConnectionTestResultDto instantiates a new ConnectionTestResultDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionTestResultDtoWithDefaults

`func NewConnectionTestResultDtoWithDefaults() *ConnectionTestResultDto`

NewConnectionTestResultDtoWithDefaults instantiates a new ConnectionTestResultDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ConnectionTestResultDto) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ConnectionTestResultDto) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ConnectionTestResultDto) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ConnectionTestResultDto) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetError

`func (o *ConnectionTestResultDto) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ConnectionTestResultDto) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ConnectionTestResultDto) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ConnectionTestResultDto) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *ConnectionTestResultDto) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *ConnectionTestResultDto) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


