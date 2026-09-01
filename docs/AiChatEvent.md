# AiChatEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Emitted once per `sendWithStream` call, immediately after the user message has been persisted by storage and before the assistant stream starts. Carries the storage-assigned `id` and `createdAt`. The UI uses it to render the user bubble — no client-side optimistic placeholder is needed, which keeps the runtime tree free of phantom nodes from index-fallback ids. | 
**Message** | Pointer to [**AiThreadMessageLike**](AiThreadMessageLike.md) | The message the event is about, in the state it has reached. | [optional] 
**MessageId** | Pointer to **string** | The storage identifier of that message. | [optional] 
**Idx** | Pointer to **float32** | The zero-based position of the pending tool call within the message. | [optional] 
**ThreadId** | Pointer to **string** | The thread the event belongs to. | [optional] 
**AutoAllow** | Pointer to **bool** | The consumer should execute the tool without prompting the user. True when the tool is in the persisted always-allow list, or the tool itself opts in via `TMCPItem.requireApproval === false` (host tools default to this). For a client-side tool with a server-side engine, this lets the engine return the pending call already flagged auto-allow so the client runs it and streams the result back without a dialog round-trip. | [optional] 
**ServerExecuted** | Pointer to **bool** | Set when the tool is served by a server-side system source: the consumer must NOT execute it locally — only show the approval UI (unless `autoAllow`) and resume via `approveToolCall` (no `result` needed) / `denyToolCall`. The engine runs it in-engine. | [optional] 
**Title** | Pointer to **string** | The generated thread title. | [optional] 
**ProfileId** | Pointer to **string** | The profile that generated the title, when one was used. | [optional] 

## Methods

### NewAiChatEvent

`func NewAiChatEvent(type_ string, ) *AiChatEvent`

NewAiChatEvent instantiates a new AiChatEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiChatEventWithDefaults

`func NewAiChatEventWithDefaults() *AiChatEvent`

NewAiChatEventWithDefaults instantiates a new AiChatEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AiChatEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AiChatEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AiChatEvent) SetType(v string)`

SetType sets Type field to given value.


### GetMessage

`func (o *AiChatEvent) GetMessage() AiThreadMessageLike`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AiChatEvent) GetMessageOk() (*AiThreadMessageLike, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AiChatEvent) SetMessage(v AiThreadMessageLike)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AiChatEvent) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMessageId

`func (o *AiChatEvent) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *AiChatEvent) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *AiChatEvent) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *AiChatEvent) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetIdx

`func (o *AiChatEvent) GetIdx() float32`

GetIdx returns the Idx field if non-nil, zero value otherwise.

### GetIdxOk

`func (o *AiChatEvent) GetIdxOk() (*float32, bool)`

GetIdxOk returns a tuple with the Idx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdx

`func (o *AiChatEvent) SetIdx(v float32)`

SetIdx sets Idx field to given value.

### HasIdx

`func (o *AiChatEvent) HasIdx() bool`

HasIdx returns a boolean if a field has been set.

### GetThreadId

`func (o *AiChatEvent) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *AiChatEvent) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *AiChatEvent) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.

### HasThreadId

`func (o *AiChatEvent) HasThreadId() bool`

HasThreadId returns a boolean if a field has been set.

### GetAutoAllow

`func (o *AiChatEvent) GetAutoAllow() bool`

GetAutoAllow returns the AutoAllow field if non-nil, zero value otherwise.

### GetAutoAllowOk

`func (o *AiChatEvent) GetAutoAllowOk() (*bool, bool)`

GetAutoAllowOk returns a tuple with the AutoAllow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAllow

`func (o *AiChatEvent) SetAutoAllow(v bool)`

SetAutoAllow sets AutoAllow field to given value.

### HasAutoAllow

`func (o *AiChatEvent) HasAutoAllow() bool`

HasAutoAllow returns a boolean if a field has been set.

### GetServerExecuted

`func (o *AiChatEvent) GetServerExecuted() bool`

GetServerExecuted returns the ServerExecuted field if non-nil, zero value otherwise.

### GetServerExecutedOk

`func (o *AiChatEvent) GetServerExecutedOk() (*bool, bool)`

GetServerExecutedOk returns a tuple with the ServerExecuted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerExecuted

`func (o *AiChatEvent) SetServerExecuted(v bool)`

SetServerExecuted sets ServerExecuted field to given value.

### HasServerExecuted

`func (o *AiChatEvent) HasServerExecuted() bool`

HasServerExecuted returns a boolean if a field has been set.

### GetTitle

`func (o *AiChatEvent) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AiChatEvent) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AiChatEvent) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AiChatEvent) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetProfileId

`func (o *AiChatEvent) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *AiChatEvent) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *AiChatEvent) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *AiChatEvent) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


