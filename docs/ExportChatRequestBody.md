# ExportChatRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FolderId** | [**ExportChatRequestBodyFolderId**](ExportChatRequestBodyFolderId.md) |  | 
**Title** | **NullableString** | The file name (without extension) to use for the exported document. | 

## Methods

### NewExportChatRequestBody

`func NewExportChatRequestBody(folderId ExportChatRequestBodyFolderId, title NullableString, ) *ExportChatRequestBody`

NewExportChatRequestBody instantiates a new ExportChatRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportChatRequestBodyWithDefaults

`func NewExportChatRequestBodyWithDefaults() *ExportChatRequestBody`

NewExportChatRequestBodyWithDefaults instantiates a new ExportChatRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFolderId

`func (o *ExportChatRequestBody) GetFolderId() ExportChatRequestBodyFolderId`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *ExportChatRequestBody) GetFolderIdOk() (*ExportChatRequestBodyFolderId, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *ExportChatRequestBody) SetFolderId(v ExportChatRequestBodyFolderId)`

SetFolderId sets FolderId field to given value.


### GetTitle

`func (o *ExportChatRequestBody) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ExportChatRequestBody) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ExportChatRequestBody) SetTitle(v string)`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *ExportChatRequestBody) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ExportChatRequestBody) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


