# AiUserSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatRecommendedModelVisible** | Pointer to **bool** | Indicates whether the recommended model banner is visible in the AI chat for the current user. | [optional] 
**ToolPermissionMode** | Pointer to [**AiToolPermissionMode**](AiToolPermissionMode.md) | How tool calls made by the model are approved for the current user. The default applies while the user has stored nothing. | [optional] 

## Methods

### NewAiUserSettingsDto

`func NewAiUserSettingsDto() *AiUserSettingsDto`

NewAiUserSettingsDto instantiates a new AiUserSettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiUserSettingsDtoWithDefaults

`func NewAiUserSettingsDtoWithDefaults() *AiUserSettingsDto`

NewAiUserSettingsDtoWithDefaults instantiates a new AiUserSettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatRecommendedModelVisible

`func (o *AiUserSettingsDto) GetChatRecommendedModelVisible() bool`

GetChatRecommendedModelVisible returns the ChatRecommendedModelVisible field if non-nil, zero value otherwise.

### GetChatRecommendedModelVisibleOk

`func (o *AiUserSettingsDto) GetChatRecommendedModelVisibleOk() (*bool, bool)`

GetChatRecommendedModelVisibleOk returns a tuple with the ChatRecommendedModelVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatRecommendedModelVisible

`func (o *AiUserSettingsDto) SetChatRecommendedModelVisible(v bool)`

SetChatRecommendedModelVisible sets ChatRecommendedModelVisible field to given value.

### HasChatRecommendedModelVisible

`func (o *AiUserSettingsDto) HasChatRecommendedModelVisible() bool`

HasChatRecommendedModelVisible returns a boolean if a field has been set.

### GetToolPermissionMode

`func (o *AiUserSettingsDto) GetToolPermissionMode() AiToolPermissionMode`

GetToolPermissionMode returns the ToolPermissionMode field if non-nil, zero value otherwise.

### GetToolPermissionModeOk

`func (o *AiUserSettingsDto) GetToolPermissionModeOk() (*AiToolPermissionMode, bool)`

GetToolPermissionModeOk returns a tuple with the ToolPermissionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolPermissionMode

`func (o *AiUserSettingsDto) SetToolPermissionMode(v AiToolPermissionMode)`

SetToolPermissionMode sets ToolPermissionMode field to given value.

### HasToolPermissionMode

`func (o *AiUserSettingsDto) HasToolPermissionMode() bool`

HasToolPermissionMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


