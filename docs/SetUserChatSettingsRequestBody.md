# SetUserChatSettingsRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebSearchEnabled** | Pointer to **NullableBool** | Indicates whether the AI assistant is allowed to perform web searches when generating responses. | [optional] 
**ReasoningEffort** | Pointer to [**ChatReasoningEffort**](ChatReasoningEffort.md) |  | [optional] 

## Methods

### NewSetUserChatSettingsRequestBody

`func NewSetUserChatSettingsRequestBody() *SetUserChatSettingsRequestBody`

NewSetUserChatSettingsRequestBody instantiates a new SetUserChatSettingsRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetUserChatSettingsRequestBodyWithDefaults

`func NewSetUserChatSettingsRequestBodyWithDefaults() *SetUserChatSettingsRequestBody`

NewSetUserChatSettingsRequestBodyWithDefaults instantiates a new SetUserChatSettingsRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebSearchEnabled

`func (o *SetUserChatSettingsRequestBody) GetWebSearchEnabled() bool`

GetWebSearchEnabled returns the WebSearchEnabled field if non-nil, zero value otherwise.

### GetWebSearchEnabledOk

`func (o *SetUserChatSettingsRequestBody) GetWebSearchEnabledOk() (*bool, bool)`

GetWebSearchEnabledOk returns a tuple with the WebSearchEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebSearchEnabled

`func (o *SetUserChatSettingsRequestBody) SetWebSearchEnabled(v bool)`

SetWebSearchEnabled sets WebSearchEnabled field to given value.

### HasWebSearchEnabled

`func (o *SetUserChatSettingsRequestBody) HasWebSearchEnabled() bool`

HasWebSearchEnabled returns a boolean if a field has been set.

### SetWebSearchEnabledNil

`func (o *SetUserChatSettingsRequestBody) SetWebSearchEnabledNil(b bool)`

 SetWebSearchEnabledNil sets the value for WebSearchEnabled to be an explicit nil

### UnsetWebSearchEnabled
`func (o *SetUserChatSettingsRequestBody) UnsetWebSearchEnabled()`

UnsetWebSearchEnabled ensures that no value is present for WebSearchEnabled, not even an explicit nil
### GetReasoningEffort

`func (o *SetUserChatSettingsRequestBody) GetReasoningEffort() ChatReasoningEffort`

GetReasoningEffort returns the ReasoningEffort field if non-nil, zero value otherwise.

### GetReasoningEffortOk

`func (o *SetUserChatSettingsRequestBody) GetReasoningEffortOk() (*ChatReasoningEffort, bool)`

GetReasoningEffortOk returns a tuple with the ReasoningEffort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoningEffort

`func (o *SetUserChatSettingsRequestBody) SetReasoningEffort(v ChatReasoningEffort)`

SetReasoningEffort sets ReasoningEffort field to given value.

### HasReasoningEffort

`func (o *SetUserChatSettingsRequestBody) HasReasoningEffort() bool`

HasReasoningEffort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


