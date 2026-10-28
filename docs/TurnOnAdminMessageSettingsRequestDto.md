# TurnOnAdminMessageSettingsRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TurnOn** | Pointer to **bool** | Whether the form is offered. Switching it off hides the form for everybody and makes the operation that  submits it refuse new messages; letters already sent are untouched. | [optional] 

## Methods

### NewTurnOnAdminMessageSettingsRequestDto

`func NewTurnOnAdminMessageSettingsRequestDto() *TurnOnAdminMessageSettingsRequestDto`

NewTurnOnAdminMessageSettingsRequestDto instantiates a new TurnOnAdminMessageSettingsRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTurnOnAdminMessageSettingsRequestDtoWithDefaults

`func NewTurnOnAdminMessageSettingsRequestDtoWithDefaults() *TurnOnAdminMessageSettingsRequestDto`

NewTurnOnAdminMessageSettingsRequestDtoWithDefaults instantiates a new TurnOnAdminMessageSettingsRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTurnOn

`func (o *TurnOnAdminMessageSettingsRequestDto) GetTurnOn() bool`

GetTurnOn returns the TurnOn field if non-nil, zero value otherwise.

### GetTurnOnOk

`func (o *TurnOnAdminMessageSettingsRequestDto) GetTurnOnOk() (*bool, bool)`

GetTurnOnOk returns a tuple with the TurnOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTurnOn

`func (o *TurnOnAdminMessageSettingsRequestDto) SetTurnOn(v bool)`

SetTurnOn sets TurnOn field to given value.

### HasTurnOn

`func (o *TurnOnAdminMessageSettingsRequestDto) HasTurnOn() bool`

HasTurnOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


