# SaveAsPdf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FolderId** | **int32** | The folder the PDF is created in; the caller has to be allowed to create files there. | 
**Title** | **NullableString** | The name of the PDF, without an extension - `.pdf` is appended. Left empty, the name of the source file is  reused with its extension replaced. | 

## Methods

### NewSaveAsPdf

`func NewSaveAsPdf(folderId int32, title NullableString, ) *SaveAsPdf`

NewSaveAsPdf instantiates a new SaveAsPdf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveAsPdfWithDefaults

`func NewSaveAsPdfWithDefaults() *SaveAsPdf`

NewSaveAsPdfWithDefaults instantiates a new SaveAsPdf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFolderId

`func (o *SaveAsPdf) GetFolderId() int32`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *SaveAsPdf) GetFolderIdOk() (*int32, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *SaveAsPdf) SetFolderId(v int32)`

SetFolderId sets FolderId field to given value.


### GetTitle

`func (o *SaveAsPdf) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SaveAsPdf) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SaveAsPdf) SetTitle(v string)`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *SaveAsPdf) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SaveAsPdf) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


