# StartNewChatBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **NullableString** | The initial user message to send to the AI assistant. | 
**ContextFolderId** | Pointer to **NullableInt32** | The optional collection of file identifiers to attach as context for the AI model. | [optional] 
**Files** | Pointer to [**[]ContinueChatBodyFilesInner**](ContinueChatBodyFilesInner.md) | The list of attached files. | [optional] 

## Methods

### NewStartNewChatBody

`func NewStartNewChatBody(message NullableString, ) *StartNewChatBody`

NewStartNewChatBody instantiates a new StartNewChatBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartNewChatBodyWithDefaults

`func NewStartNewChatBodyWithDefaults() *StartNewChatBody`

NewStartNewChatBodyWithDefaults instantiates a new StartNewChatBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *StartNewChatBody) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *StartNewChatBody) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *StartNewChatBody) SetMessage(v string)`

SetMessage sets Message field to given value.


### SetMessageNil

`func (o *StartNewChatBody) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *StartNewChatBody) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetContextFolderId

`func (o *StartNewChatBody) GetContextFolderId() int32`

GetContextFolderId returns the ContextFolderId field if non-nil, zero value otherwise.

### GetContextFolderIdOk

`func (o *StartNewChatBody) GetContextFolderIdOk() (*int32, bool)`

GetContextFolderIdOk returns a tuple with the ContextFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextFolderId

`func (o *StartNewChatBody) SetContextFolderId(v int32)`

SetContextFolderId sets ContextFolderId field to given value.

### HasContextFolderId

`func (o *StartNewChatBody) HasContextFolderId() bool`

HasContextFolderId returns a boolean if a field has been set.

### SetContextFolderIdNil

`func (o *StartNewChatBody) SetContextFolderIdNil(b bool)`

 SetContextFolderIdNil sets the value for ContextFolderId to be an explicit nil

### UnsetContextFolderId
`func (o *StartNewChatBody) UnsetContextFolderId()`

UnsetContextFolderId ensures that no value is present for ContextFolderId, not even an explicit nil
### GetFiles

`func (o *StartNewChatBody) GetFiles() []ContinueChatBodyFilesInner`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *StartNewChatBody) GetFilesOk() (*[]ContinueChatBodyFilesInner, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *StartNewChatBody) SetFiles(v []ContinueChatBodyFilesInner)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *StartNewChatBody) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### SetFilesNil

`func (o *StartNewChatBody) SetFilesNil(b bool)`

 SetFilesNil sets the value for Files to be an explicit nil

### UnsetFiles
`func (o *StartNewChatBody) UnsetFiles()`

UnsetFiles ensures that no value is present for Files, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


