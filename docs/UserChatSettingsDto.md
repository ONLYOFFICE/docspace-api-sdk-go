# UserChatSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebSearchEnabled** | Pointer to **bool** | Indicates whether the AI assistant is allowed to perform web searches when generating responses in this room. | [optional] 
**ReasoningEffort** | Pointer to [**ChatReasoningEffort**](ChatReasoningEffort.md) |  | [optional] 

## Methods

### NewUserChatSettingsDto

`func NewUserChatSettingsDto() *UserChatSettingsDto`

NewUserChatSettingsDto instantiates a new UserChatSettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserChatSettingsDtoWithDefaults

`func NewUserChatSettingsDtoWithDefaults() *UserChatSettingsDto`

NewUserChatSettingsDtoWithDefaults instantiates a new UserChatSettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebSearchEnabled

`func (o *UserChatSettingsDto) GetWebSearchEnabled() bool`

GetWebSearchEnabled returns the WebSearchEnabled field if non-nil, zero value otherwise.

### GetWebSearchEnabledOk

`func (o *UserChatSettingsDto) GetWebSearchEnabledOk() (*bool, bool)`

GetWebSearchEnabledOk returns a tuple with the WebSearchEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebSearchEnabled

`func (o *UserChatSettingsDto) SetWebSearchEnabled(v bool)`

SetWebSearchEnabled sets WebSearchEnabled field to given value.

### HasWebSearchEnabled

`func (o *UserChatSettingsDto) HasWebSearchEnabled() bool`

HasWebSearchEnabled returns a boolean if a field has been set.

### GetReasoningEffort

`func (o *UserChatSettingsDto) GetReasoningEffort() ChatReasoningEffort`

GetReasoningEffort returns the ReasoningEffort field if non-nil, zero value otherwise.

### GetReasoningEffortOk

`func (o *UserChatSettingsDto) GetReasoningEffortOk() (*ChatReasoningEffort, bool)`

GetReasoningEffortOk returns a tuple with the ReasoningEffort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoningEffort

`func (o *UserChatSettingsDto) SetReasoningEffort(v ChatReasoningEffort)`

SetReasoningEffort sets ReasoningEffort field to given value.

### HasReasoningEffort

`func (o *UserChatSettingsDto) HasReasoningEffort() bool`

HasReasoningEffort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


