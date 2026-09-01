# AiAttachmentsSaveFileRequestInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | Storage path/key of the file. | 
**Content** | **string** | File contents. | 
**Type** | **float32** | File type discriminator. | 
**Title** | Pointer to **string** | Optional display title. | [optional] 

## Methods

### NewAiAttachmentsSaveFileRequestInput

`func NewAiAttachmentsSaveFileRequestInput(path string, content string, type_ float32, ) *AiAttachmentsSaveFileRequestInput`

NewAiAttachmentsSaveFileRequestInput instantiates a new AiAttachmentsSaveFileRequestInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiAttachmentsSaveFileRequestInputWithDefaults

`func NewAiAttachmentsSaveFileRequestInputWithDefaults() *AiAttachmentsSaveFileRequestInput`

NewAiAttachmentsSaveFileRequestInputWithDefaults instantiates a new AiAttachmentsSaveFileRequestInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *AiAttachmentsSaveFileRequestInput) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AiAttachmentsSaveFileRequestInput) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AiAttachmentsSaveFileRequestInput) SetPath(v string)`

SetPath sets Path field to given value.


### GetContent

`func (o *AiAttachmentsSaveFileRequestInput) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *AiAttachmentsSaveFileRequestInput) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *AiAttachmentsSaveFileRequestInput) SetContent(v string)`

SetContent sets Content field to given value.


### GetType

`func (o *AiAttachmentsSaveFileRequestInput) GetType() float32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AiAttachmentsSaveFileRequestInput) GetTypeOk() (*float32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AiAttachmentsSaveFileRequestInput) SetType(v float32)`

SetType sets Type field to given value.


### GetTitle

`func (o *AiAttachmentsSaveFileRequestInput) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AiAttachmentsSaveFileRequestInput) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AiAttachmentsSaveFileRequestInput) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AiAttachmentsSaveFileRequestInput) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


