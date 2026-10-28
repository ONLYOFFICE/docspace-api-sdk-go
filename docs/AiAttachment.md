# AiAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Storage-assigned UUID. | 
**Kind** | **string** | file | image. | 
**Source** | Pointer to **string** | Origin of the attachment. `user` — uploaded by the user in the composer (the default when unset, for backward compatibility). `tool` — produced by a tool call (e.g. `generate_image`). Lets the integrator's adapter route or apply policies (separate bucket, quotas, TTL, CDN) per source. | [optional] 
**Title** | **string** | Display label (filename or user-visible title). | 
**Content** | Pointer to **string** | Extracted text for files. | [optional] 
**Base64** | Pointer to **string** | Base64 data URL for images. | [optional] 
**Path** | Pointer to **string** | Original host file path (for files). | [optional] 
**Type** | Pointer to **float32** | ONLYOFFICE file type code (for files). | [optional] 
**MessageId** | Pointer to **string** | Owning message id once linked. Unset while the attachment is a draft. | [optional] 
**ThreadId** | Pointer to **string** | Owning thread id once linked. Unset while the attachment is a draft. | [optional] 
**EntityId** | Pointer to **string** | Opaque scope token (entity / room) the attachment was created in. Drafts carry it so an entity switch keeps in-flight composer state isolated; once linked to a message the field is redundant with the thread's own entity binding. | [optional] 
**CreatedAt** | **float32** | Storage-assigned creation timestamp. | 
**CanAnalyze** | Pointer to **bool** | Whether the attached form can be analyzed. | [optional] 
**FormKeys** | Pointer to [**[]AiAttachmentFormKeysInner**](AiAttachmentFormKeysInner.md) | Keys of the fields inside the form. `key` is the field identifier, `text` its human-readable label. | [optional] 

## Methods

### NewAiAttachment

`func NewAiAttachment(id string, kind string, title string, createdAt float32, ) *AiAttachment`

NewAiAttachment instantiates a new AiAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiAttachmentWithDefaults

`func NewAiAttachmentWithDefaults() *AiAttachment`

NewAiAttachmentWithDefaults instantiates a new AiAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AiAttachment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiAttachment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiAttachment) SetId(v string)`

SetId sets Id field to given value.


### GetKind

`func (o *AiAttachment) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AiAttachment) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AiAttachment) SetKind(v string)`

SetKind sets Kind field to given value.


### GetSource

`func (o *AiAttachment) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AiAttachment) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AiAttachment) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *AiAttachment) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTitle

`func (o *AiAttachment) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AiAttachment) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AiAttachment) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetContent

`func (o *AiAttachment) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *AiAttachment) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *AiAttachment) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *AiAttachment) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetBase64

`func (o *AiAttachment) GetBase64() string`

GetBase64 returns the Base64 field if non-nil, zero value otherwise.

### GetBase64Ok

`func (o *AiAttachment) GetBase64Ok() (*string, bool)`

GetBase64Ok returns a tuple with the Base64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase64

`func (o *AiAttachment) SetBase64(v string)`

SetBase64 sets Base64 field to given value.

### HasBase64

`func (o *AiAttachment) HasBase64() bool`

HasBase64 returns a boolean if a field has been set.

### GetPath

`func (o *AiAttachment) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AiAttachment) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AiAttachment) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *AiAttachment) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetType

`func (o *AiAttachment) GetType() float32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AiAttachment) GetTypeOk() (*float32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AiAttachment) SetType(v float32)`

SetType sets Type field to given value.

### HasType

`func (o *AiAttachment) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMessageId

`func (o *AiAttachment) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *AiAttachment) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *AiAttachment) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *AiAttachment) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetThreadId

`func (o *AiAttachment) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *AiAttachment) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *AiAttachment) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.

### HasThreadId

`func (o *AiAttachment) HasThreadId() bool`

HasThreadId returns a boolean if a field has been set.

### GetEntityId

`func (o *AiAttachment) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *AiAttachment) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *AiAttachment) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *AiAttachment) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AiAttachment) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AiAttachment) GetCreatedAtOk() (*float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AiAttachment) SetCreatedAt(v float32)`

SetCreatedAt sets CreatedAt field to given value.


### GetCanAnalyze

`func (o *AiAttachment) GetCanAnalyze() bool`

GetCanAnalyze returns the CanAnalyze field if non-nil, zero value otherwise.

### GetCanAnalyzeOk

`func (o *AiAttachment) GetCanAnalyzeOk() (*bool, bool)`

GetCanAnalyzeOk returns a tuple with the CanAnalyze field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAnalyze

`func (o *AiAttachment) SetCanAnalyze(v bool)`

SetCanAnalyze sets CanAnalyze field to given value.

### HasCanAnalyze

`func (o *AiAttachment) HasCanAnalyze() bool`

HasCanAnalyze returns a boolean if a field has been set.

### GetFormKeys

`func (o *AiAttachment) GetFormKeys() []AiAttachmentFormKeysInner`

GetFormKeys returns the FormKeys field if non-nil, zero value otherwise.

### GetFormKeysOk

`func (o *AiAttachment) GetFormKeysOk() (*[]AiAttachmentFormKeysInner, bool)`

GetFormKeysOk returns a tuple with the FormKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormKeys

`func (o *AiAttachment) SetFormKeys(v []AiAttachmentFormKeysInner)`

SetFormKeys sets FormKeys field to given value.

### HasFormKeys

`func (o *AiAttachment) HasFormKeys() bool`

HasFormKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


