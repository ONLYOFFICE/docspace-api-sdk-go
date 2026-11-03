# AiToolPermissionModeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to [**AiToolPermissionMode**](AiToolPermissionMode.md) | How tool calls made by the model are approved for the current user. The default applies while the user has stored nothing. | [optional] 

## Methods

### NewAiToolPermissionModeDto

`func NewAiToolPermissionModeDto() *AiToolPermissionModeDto`

NewAiToolPermissionModeDto instantiates a new AiToolPermissionModeDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiToolPermissionModeDtoWithDefaults

`func NewAiToolPermissionModeDtoWithDefaults() *AiToolPermissionModeDto`

NewAiToolPermissionModeDtoWithDefaults instantiates a new AiToolPermissionModeDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *AiToolPermissionModeDto) GetMode() AiToolPermissionMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *AiToolPermissionModeDto) GetModeOk() (*AiToolPermissionMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *AiToolPermissionModeDto) SetMode(v AiToolPermissionMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *AiToolPermissionModeDto) HasMode() bool`

HasMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


