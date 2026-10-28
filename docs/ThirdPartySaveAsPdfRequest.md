# ThirdPartySaveAsPdfRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FolderId** | **NullableString** | The folder the PDF is created in; the caller has to be allowed to create files there. | 
**Title** | **NullableString** | The name of the PDF, without an extension - `.pdf` is appended. Left empty, the name of the source file is  reused with its extension replaced. | 

## Methods

### NewThirdPartySaveAsPdfRequest

`func NewThirdPartySaveAsPdfRequest(folderId NullableString, title NullableString, ) *ThirdPartySaveAsPdfRequest`

NewThirdPartySaveAsPdfRequest instantiates a new ThirdPartySaveAsPdfRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThirdPartySaveAsPdfRequestWithDefaults

`func NewThirdPartySaveAsPdfRequestWithDefaults() *ThirdPartySaveAsPdfRequest`

NewThirdPartySaveAsPdfRequestWithDefaults instantiates a new ThirdPartySaveAsPdfRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFolderId

`func (o *ThirdPartySaveAsPdfRequest) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *ThirdPartySaveAsPdfRequest) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *ThirdPartySaveAsPdfRequest) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.


### SetFolderIdNil

`func (o *ThirdPartySaveAsPdfRequest) SetFolderIdNil(b bool)`

 SetFolderIdNil sets the value for FolderId to be an explicit nil

### UnsetFolderId
`func (o *ThirdPartySaveAsPdfRequest) UnsetFolderId()`

UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
### GetTitle

`func (o *ThirdPartySaveAsPdfRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ThirdPartySaveAsPdfRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ThirdPartySaveAsPdfRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *ThirdPartySaveAsPdfRequest) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ThirdPartySaveAsPdfRequest) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


