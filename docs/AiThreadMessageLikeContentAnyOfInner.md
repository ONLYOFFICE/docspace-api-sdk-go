# AiThreadMessageLikeContentAnyOfInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Part kind: `text`, `reasoning`, `source`, `image`, `file`, `audio`, `data` or `tool-call`. | 
**Text** | Pointer to **string** | The text of a `text` or `reasoning` part. | [optional] 
**ParentId** | Pointer to **string** | Id of the part this one belongs to, on `text`, `reasoning`, `source` and `tool-call` parts. On a `reasoning` part the chat library keeps the provider's thinking signature here. | [optional] 
**SourceType** | Pointer to **string** | A `source` part's source kind; always `url`. | [optional] 
**Id** | Pointer to **string** | A `source` part's id. | [optional] 
**Url** | Pointer to **string** | A `source` part's URL. | [optional] 
**Title** | Pointer to **string** | A `source` part's title. | [optional] 
**Image** | Pointer to **string** | An `image` part's picture, as a URL (usually a `data:` URL). For an image the user attached, a JSON string `{ ref, title, kind: image }` naming the stored attachment instead. | [optional] 
**MimeType** | Pointer to **string** | A `file` part's media type. For a file the user attached, a JSON string `{ ref, title, kind: file, path, type }` naming the stored attachment instead - the chat library carries the reference here. | [optional] 
**Data** | Pointer to **string** | A `file` part's content; empty when the part references an attachment. (A `data` part would carry its payload here too, as any JSON value; neither the chat library nor this service writes one.) | [optional] 
**Filename** | Pointer to **string** | The file name of a `file` or `image` part. | [optional] 
**Name** | Pointer to **string** | An `image` part's display name, or a `data` part's name. | [optional] 
**Audio** | Pointer to [**AiThreadMessageLikeContentAnyOfInnerAudio**](AiThreadMessageLikeContentAnyOfInnerAudio.md) |  | [optional] 
**ToolCallId** | Pointer to **string** | A `tool-call` part's call id. | [optional] 
**ToolName** | Pointer to **string** | A `tool-call` part's tool name. | [optional] 
**Args** | Pointer to **map[string]interface{}** | A `tool-call` part's arguments. | [optional] 
**ArgsText** | Pointer to **string** | A `tool-call` part's arguments as the model streamed them, JSON text. | [optional] 
**Result** | Pointer to **interface{}** |  | [optional] 
**IsError** | Pointer to **bool** | Whether a `tool-call` part's result is an error. | [optional] 
**Artifact** | Pointer to **interface{}** |  | [optional] 
**Interrupt** | Pointer to [**AiThreadMessageLikeContentAnyOfInnerInterrupt**](AiThreadMessageLikeContentAnyOfInnerInterrupt.md) |  | [optional] 
**Messages** | Pointer to **[]map[string]interface{}** | Messages of a sub-agent run inside a `tool-call` part. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Provider-specific data of a `tool-call` part - for a Gemini call, `{ thoughtSignature }`, which has to be sent back with the call. | [optional] 

## Methods

### NewAiThreadMessageLikeContentAnyOfInner

`func NewAiThreadMessageLikeContentAnyOfInner(type_ string, ) *AiThreadMessageLikeContentAnyOfInner`

NewAiThreadMessageLikeContentAnyOfInner instantiates a new AiThreadMessageLikeContentAnyOfInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiThreadMessageLikeContentAnyOfInnerWithDefaults

`func NewAiThreadMessageLikeContentAnyOfInnerWithDefaults() *AiThreadMessageLikeContentAnyOfInner`

NewAiThreadMessageLikeContentAnyOfInnerWithDefaults instantiates a new AiThreadMessageLikeContentAnyOfInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AiThreadMessageLikeContentAnyOfInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AiThreadMessageLikeContentAnyOfInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AiThreadMessageLikeContentAnyOfInner) SetType(v string)`

SetType sets Type field to given value.


### GetText

`func (o *AiThreadMessageLikeContentAnyOfInner) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *AiThreadMessageLikeContentAnyOfInner) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *AiThreadMessageLikeContentAnyOfInner) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *AiThreadMessageLikeContentAnyOfInner) HasText() bool`

HasText returns a boolean if a field has been set.

### GetParentId

`func (o *AiThreadMessageLikeContentAnyOfInner) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *AiThreadMessageLikeContentAnyOfInner) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *AiThreadMessageLikeContentAnyOfInner) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *AiThreadMessageLikeContentAnyOfInner) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetSourceType

`func (o *AiThreadMessageLikeContentAnyOfInner) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *AiThreadMessageLikeContentAnyOfInner) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *AiThreadMessageLikeContentAnyOfInner) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *AiThreadMessageLikeContentAnyOfInner) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetId

`func (o *AiThreadMessageLikeContentAnyOfInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiThreadMessageLikeContentAnyOfInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiThreadMessageLikeContentAnyOfInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AiThreadMessageLikeContentAnyOfInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *AiThreadMessageLikeContentAnyOfInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AiThreadMessageLikeContentAnyOfInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AiThreadMessageLikeContentAnyOfInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AiThreadMessageLikeContentAnyOfInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetTitle

`func (o *AiThreadMessageLikeContentAnyOfInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AiThreadMessageLikeContentAnyOfInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AiThreadMessageLikeContentAnyOfInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AiThreadMessageLikeContentAnyOfInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetImage

`func (o *AiThreadMessageLikeContentAnyOfInner) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *AiThreadMessageLikeContentAnyOfInner) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *AiThreadMessageLikeContentAnyOfInner) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *AiThreadMessageLikeContentAnyOfInner) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetMimeType

`func (o *AiThreadMessageLikeContentAnyOfInner) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *AiThreadMessageLikeContentAnyOfInner) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *AiThreadMessageLikeContentAnyOfInner) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *AiThreadMessageLikeContentAnyOfInner) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetData

`func (o *AiThreadMessageLikeContentAnyOfInner) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AiThreadMessageLikeContentAnyOfInner) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AiThreadMessageLikeContentAnyOfInner) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *AiThreadMessageLikeContentAnyOfInner) HasData() bool`

HasData returns a boolean if a field has been set.

### GetFilename

`func (o *AiThreadMessageLikeContentAnyOfInner) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *AiThreadMessageLikeContentAnyOfInner) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *AiThreadMessageLikeContentAnyOfInner) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *AiThreadMessageLikeContentAnyOfInner) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetName

`func (o *AiThreadMessageLikeContentAnyOfInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AiThreadMessageLikeContentAnyOfInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AiThreadMessageLikeContentAnyOfInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AiThreadMessageLikeContentAnyOfInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAudio

`func (o *AiThreadMessageLikeContentAnyOfInner) GetAudio() AiThreadMessageLikeContentAnyOfInnerAudio`

GetAudio returns the Audio field if non-nil, zero value otherwise.

### GetAudioOk

`func (o *AiThreadMessageLikeContentAnyOfInner) GetAudioOk() (*AiThreadMessageLikeContentAnyOfInnerAudio, bool)`

GetAudioOk returns a tuple with the Audio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudio

`func (o *AiThreadMessageLikeContentAnyOfInner) SetAudio(v AiThreadMessageLikeContentAnyOfInnerAudio)`

SetAudio sets Audio field to given value.

### HasAudio

`func (o *AiThreadMessageLikeContentAnyOfInner) HasAudio() bool`

HasAudio returns a boolean if a field has been set.

### GetToolCallId

`func (o *AiThreadMessageLikeContentAnyOfInner) GetToolCallId() string`

GetToolCallId returns the ToolCallId field if non-nil, zero value otherwise.

### GetToolCallIdOk

`func (o *AiThreadMessageLikeContentAnyOfInner) GetToolCallIdOk() (*string, bool)`

GetToolCallIdOk returns a tuple with the ToolCallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCallId

`func (o *AiThreadMessageLikeContentAnyOfInner) SetToolCallId(v string)`

SetToolCallId sets ToolCallId field to given value.

### HasToolCallId

`func (o *AiThreadMessageLikeContentAnyOfInner) HasToolCallId() bool`

HasToolCallId returns a boolean if a field has been set.

### GetToolName

`func (o *AiThreadMessageLikeContentAnyOfInner) GetToolName() string`

GetToolName returns the ToolName field if non-nil, zero value otherwise.

### GetToolNameOk

`func (o *AiThreadMessageLikeContentAnyOfInner) GetToolNameOk() (*string, bool)`

GetToolNameOk returns a tuple with the ToolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolName

`func (o *AiThreadMessageLikeContentAnyOfInner) SetToolName(v string)`

SetToolName sets ToolName field to given value.

### HasToolName

`func (o *AiThreadMessageLikeContentAnyOfInner) HasToolName() bool`

HasToolName returns a boolean if a field has been set.

### GetArgs

`func (o *AiThreadMessageLikeContentAnyOfInner) GetArgs() map[string]interface{}`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *AiThreadMessageLikeContentAnyOfInner) GetArgsOk() (*map[string]interface{}, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *AiThreadMessageLikeContentAnyOfInner) SetArgs(v map[string]interface{})`

SetArgs sets Args field to given value.

### HasArgs

`func (o *AiThreadMessageLikeContentAnyOfInner) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetArgsText

`func (o *AiThreadMessageLikeContentAnyOfInner) GetArgsText() string`

GetArgsText returns the ArgsText field if non-nil, zero value otherwise.

### GetArgsTextOk

`func (o *AiThreadMessageLikeContentAnyOfInner) GetArgsTextOk() (*string, bool)`

GetArgsTextOk returns a tuple with the ArgsText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgsText

`func (o *AiThreadMessageLikeContentAnyOfInner) SetArgsText(v string)`

SetArgsText sets ArgsText field to given value.

### HasArgsText

`func (o *AiThreadMessageLikeContentAnyOfInner) HasArgsText() bool`

HasArgsText returns a boolean if a field has been set.

### GetResult

`func (o *AiThreadMessageLikeContentAnyOfInner) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *AiThreadMessageLikeContentAnyOfInner) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *AiThreadMessageLikeContentAnyOfInner) SetResult(v interface{})`

SetResult sets Result field to given value.

### HasResult

`func (o *AiThreadMessageLikeContentAnyOfInner) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *AiThreadMessageLikeContentAnyOfInner) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *AiThreadMessageLikeContentAnyOfInner) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetIsError

`func (o *AiThreadMessageLikeContentAnyOfInner) GetIsError() bool`

GetIsError returns the IsError field if non-nil, zero value otherwise.

### GetIsErrorOk

`func (o *AiThreadMessageLikeContentAnyOfInner) GetIsErrorOk() (*bool, bool)`

GetIsErrorOk returns a tuple with the IsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsError

`func (o *AiThreadMessageLikeContentAnyOfInner) SetIsError(v bool)`

SetIsError sets IsError field to given value.

### HasIsError

`func (o *AiThreadMessageLikeContentAnyOfInner) HasIsError() bool`

HasIsError returns a boolean if a field has been set.

### GetArtifact

`func (o *AiThreadMessageLikeContentAnyOfInner) GetArtifact() interface{}`

GetArtifact returns the Artifact field if non-nil, zero value otherwise.

### GetArtifactOk

`func (o *AiThreadMessageLikeContentAnyOfInner) GetArtifactOk() (*interface{}, bool)`

GetArtifactOk returns a tuple with the Artifact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifact

`func (o *AiThreadMessageLikeContentAnyOfInner) SetArtifact(v interface{})`

SetArtifact sets Artifact field to given value.

### HasArtifact

`func (o *AiThreadMessageLikeContentAnyOfInner) HasArtifact() bool`

HasArtifact returns a boolean if a field has been set.

### SetArtifactNil

`func (o *AiThreadMessageLikeContentAnyOfInner) SetArtifactNil(b bool)`

 SetArtifactNil sets the value for Artifact to be an explicit nil

### UnsetArtifact
`func (o *AiThreadMessageLikeContentAnyOfInner) UnsetArtifact()`

UnsetArtifact ensures that no value is present for Artifact, not even an explicit nil
### GetInterrupt

`func (o *AiThreadMessageLikeContentAnyOfInner) GetInterrupt() AiThreadMessageLikeContentAnyOfInnerInterrupt`

GetInterrupt returns the Interrupt field if non-nil, zero value otherwise.

### GetInterruptOk

`func (o *AiThreadMessageLikeContentAnyOfInner) GetInterruptOk() (*AiThreadMessageLikeContentAnyOfInnerInterrupt, bool)`

GetInterruptOk returns a tuple with the Interrupt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterrupt

`func (o *AiThreadMessageLikeContentAnyOfInner) SetInterrupt(v AiThreadMessageLikeContentAnyOfInnerInterrupt)`

SetInterrupt sets Interrupt field to given value.

### HasInterrupt

`func (o *AiThreadMessageLikeContentAnyOfInner) HasInterrupt() bool`

HasInterrupt returns a boolean if a field has been set.

### GetMessages

`func (o *AiThreadMessageLikeContentAnyOfInner) GetMessages() []map[string]interface{}`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *AiThreadMessageLikeContentAnyOfInner) GetMessagesOk() (*[]map[string]interface{}, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *AiThreadMessageLikeContentAnyOfInner) SetMessages(v []map[string]interface{})`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *AiThreadMessageLikeContentAnyOfInner) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetMetadata

`func (o *AiThreadMessageLikeContentAnyOfInner) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AiThreadMessageLikeContentAnyOfInner) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AiThreadMessageLikeContentAnyOfInner) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AiThreadMessageLikeContentAnyOfInner) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


