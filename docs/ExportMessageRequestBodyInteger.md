# ExportMessageRequestBodyInteger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FolderId** | **int32** | The identifier of the destination folder where the exported document will be saved. | 
**Title** | **NullableString** | The file name (without extension) to use for the exported document. | 

## Methods

### NewExportMessageRequestBodyInteger

`func NewExportMessageRequestBodyInteger(folderId int32, title NullableString, ) *ExportMessageRequestBodyInteger`

NewExportMessageRequestBodyInteger instantiates a new ExportMessageRequestBodyInteger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportMessageRequestBodyIntegerWithDefaults

`func NewExportMessageRequestBodyIntegerWithDefaults() *ExportMessageRequestBodyInteger`

NewExportMessageRequestBodyIntegerWithDefaults instantiates a new ExportMessageRequestBodyInteger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFolderId

`func (o *ExportMessageRequestBodyInteger) GetFolderId() int32`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *ExportMessageRequestBodyInteger) GetFolderIdOk() (*int32, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *ExportMessageRequestBodyInteger) SetFolderId(v int32)`

SetFolderId sets FolderId field to given value.


### GetTitle

`func (o *ExportMessageRequestBodyInteger) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ExportMessageRequestBodyInteger) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ExportMessageRequestBodyInteger) SetTitle(v string)`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *ExportMessageRequestBodyInteger) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ExportMessageRequestBodyInteger) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


