# TelegramStatusDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**RegStatus**](RegStatus.md) |  | 
**Username** | Pointer to **NullableString** | The Telegram username. | [optional] 

## Methods

### NewTelegramStatusDto

`func NewTelegramStatusDto(status RegStatus, ) *TelegramStatusDto`

NewTelegramStatusDto instantiates a new TelegramStatusDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelegramStatusDtoWithDefaults

`func NewTelegramStatusDtoWithDefaults() *TelegramStatusDto`

NewTelegramStatusDtoWithDefaults instantiates a new TelegramStatusDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *TelegramStatusDto) GetStatus() RegStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TelegramStatusDto) GetStatusOk() (*RegStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TelegramStatusDto) SetStatus(v RegStatus)`

SetStatus sets Status field to given value.


### GetUsername

`func (o *TelegramStatusDto) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *TelegramStatusDto) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *TelegramStatusDto) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *TelegramStatusDto) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *TelegramStatusDto) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *TelegramStatusDto) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


