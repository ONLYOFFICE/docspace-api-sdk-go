# AiThreadMessageLike

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Storage-assigned message id (absent on inbound drafts). | [optional] 
**Role** | **string** | Message author role. | 
**Content** | [**AiThreadMessageLikeContent**](AiThreadMessageLikeContent.md) |  | 
**CreatedAt** | Pointer to **string** | Creation timestamp, ISO-8601 on the wire. | [optional] 
**Status** | Pointer to [**AiThreadMessageLikeStatus**](AiThreadMessageLikeStatus.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Arbitrary per-message metadata. | [optional] 
**Attachments** | Pointer to **[]map[string]interface{}** | Attachments linked to the message. | [optional] 

## Methods

### NewAiThreadMessageLike

`func NewAiThreadMessageLike(role string, content AiThreadMessageLikeContent, ) *AiThreadMessageLike`

NewAiThreadMessageLike instantiates a new AiThreadMessageLike object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiThreadMessageLikeWithDefaults

`func NewAiThreadMessageLikeWithDefaults() *AiThreadMessageLike`

NewAiThreadMessageLikeWithDefaults instantiates a new AiThreadMessageLike object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AiThreadMessageLike) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiThreadMessageLike) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiThreadMessageLike) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AiThreadMessageLike) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRole

`func (o *AiThreadMessageLike) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AiThreadMessageLike) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AiThreadMessageLike) SetRole(v string)`

SetRole sets Role field to given value.


### GetContent

`func (o *AiThreadMessageLike) GetContent() AiThreadMessageLikeContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *AiThreadMessageLike) GetContentOk() (*AiThreadMessageLikeContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *AiThreadMessageLike) SetContent(v AiThreadMessageLikeContent)`

SetContent sets Content field to given value.


### GetCreatedAt

`func (o *AiThreadMessageLike) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AiThreadMessageLike) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AiThreadMessageLike) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AiThreadMessageLike) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetStatus

`func (o *AiThreadMessageLike) GetStatus() AiThreadMessageLikeStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AiThreadMessageLike) GetStatusOk() (*AiThreadMessageLikeStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AiThreadMessageLike) SetStatus(v AiThreadMessageLikeStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AiThreadMessageLike) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMetadata

`func (o *AiThreadMessageLike) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AiThreadMessageLike) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AiThreadMessageLike) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AiThreadMessageLike) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetAttachments

`func (o *AiThreadMessageLike) GetAttachments() []map[string]interface{}`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *AiThreadMessageLike) GetAttachmentsOk() (*[]map[string]interface{}, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *AiThreadMessageLike) SetAttachments(v []map[string]interface{})`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *AiThreadMessageLike) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


