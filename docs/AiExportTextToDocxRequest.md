# AiExportTextToDocxRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Document title (also the file name). | 
**Content** | **string** | Markdown content to convert. | 
**FolderId** | [**AiExportTextToDocxRequestFolderId**](AiExportTextToDocxRequestFolderId.md) |  | 

## Methods

### NewAiExportTextToDocxRequest

`func NewAiExportTextToDocxRequest(title string, content string, folderId AiExportTextToDocxRequestFolderId, ) *AiExportTextToDocxRequest`

NewAiExportTextToDocxRequest instantiates a new AiExportTextToDocxRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiExportTextToDocxRequestWithDefaults

`func NewAiExportTextToDocxRequestWithDefaults() *AiExportTextToDocxRequest`

NewAiExportTextToDocxRequestWithDefaults instantiates a new AiExportTextToDocxRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *AiExportTextToDocxRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AiExportTextToDocxRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AiExportTextToDocxRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetContent

`func (o *AiExportTextToDocxRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *AiExportTextToDocxRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *AiExportTextToDocxRequest) SetContent(v string)`

SetContent sets Content field to given value.


### GetFolderId

`func (o *AiExportTextToDocxRequest) GetFolderId() AiExportTextToDocxRequestFolderId`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *AiExportTextToDocxRequest) GetFolderIdOk() (*AiExportTextToDocxRequestFolderId, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *AiExportTextToDocxRequest) SetFolderId(v AiExportTextToDocxRequestFolderId)`

SetFolderId sets FolderId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


