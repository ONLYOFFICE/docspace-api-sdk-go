# ConfirmDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | [**ValidationResult**](ValidationResult.md) | The confirmation result. | 
**RoomId** | Pointer to **NullableString** | The confirmation room ID. | [optional] 
**Title** | Pointer to **NullableString** | The confirmation title. | [optional] 
**Email** | Pointer to **NullableString** | The confirmation email. | [optional] 
**IsAgent** | Pointer to **bool** | The confirmation is agent. | [optional] 

## Methods

### NewConfirmDto

`func NewConfirmDto(result ValidationResult, ) *ConfirmDto`

NewConfirmDto instantiates a new ConfirmDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfirmDtoWithDefaults

`func NewConfirmDtoWithDefaults() *ConfirmDto`

NewConfirmDtoWithDefaults instantiates a new ConfirmDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ConfirmDto) GetResult() ValidationResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ConfirmDto) GetResultOk() (*ValidationResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ConfirmDto) SetResult(v ValidationResult)`

SetResult sets Result field to given value.


### GetRoomId

`func (o *ConfirmDto) GetRoomId() string`

GetRoomId returns the RoomId field if non-nil, zero value otherwise.

### GetRoomIdOk

`func (o *ConfirmDto) GetRoomIdOk() (*string, bool)`

GetRoomIdOk returns a tuple with the RoomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomId

`func (o *ConfirmDto) SetRoomId(v string)`

SetRoomId sets RoomId field to given value.

### HasRoomId

`func (o *ConfirmDto) HasRoomId() bool`

HasRoomId returns a boolean if a field has been set.

### SetRoomIdNil

`func (o *ConfirmDto) SetRoomIdNil(b bool)`

 SetRoomIdNil sets the value for RoomId to be an explicit nil

### UnsetRoomId
`func (o *ConfirmDto) UnsetRoomId()`

UnsetRoomId ensures that no value is present for RoomId, not even an explicit nil
### GetTitle

`func (o *ConfirmDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ConfirmDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ConfirmDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ConfirmDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ConfirmDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ConfirmDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetEmail

`func (o *ConfirmDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ConfirmDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ConfirmDto) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ConfirmDto) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *ConfirmDto) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ConfirmDto) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetIsAgent

`func (o *ConfirmDto) GetIsAgent() bool`

GetIsAgent returns the IsAgent field if non-nil, zero value otherwise.

### GetIsAgentOk

`func (o *ConfirmDto) GetIsAgentOk() (*bool, bool)`

GetIsAgentOk returns a tuple with the IsAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAgent

`func (o *ConfirmDto) SetIsAgent(v bool)`

SetIsAgent sets IsAgent field to given value.

### HasIsAgent

`func (o *ConfirmDto) HasIsAgent() bool`

HasIsAgent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


