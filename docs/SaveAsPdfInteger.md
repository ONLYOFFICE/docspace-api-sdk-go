# SaveAsPdfInteger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FolderId** | **int32** | The folder ID to save the file as PDF. | 
**Title** | **NullableString** | The file title to save as PDF. | 

## Methods

### NewSaveAsPdfInteger

`func NewSaveAsPdfInteger(folderId int32, title NullableString, ) *SaveAsPdfInteger`

NewSaveAsPdfInteger instantiates a new SaveAsPdfInteger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveAsPdfIntegerWithDefaults

`func NewSaveAsPdfIntegerWithDefaults() *SaveAsPdfInteger`

NewSaveAsPdfIntegerWithDefaults instantiates a new SaveAsPdfInteger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFolderId

`func (o *SaveAsPdfInteger) GetFolderId() int32`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *SaveAsPdfInteger) GetFolderIdOk() (*int32, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *SaveAsPdfInteger) SetFolderId(v int32)`

SetFolderId sets FolderId field to given value.


### GetTitle

`func (o *SaveAsPdfInteger) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SaveAsPdfInteger) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SaveAsPdfInteger) SetTitle(v string)`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *SaveAsPdfInteger) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SaveAsPdfInteger) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


