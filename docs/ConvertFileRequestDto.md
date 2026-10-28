# ConvertFileRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | **int32** | The id of the portal file to convert, as reported by a folder listing such as `GET api/2.0/files/{folderId}`. | 
**FolderId** | Pointer to **NullableInt32** | The id of the folder to save the converted file in, as reported by a folder listing such as  `GET api/2.0/files/{folderId}`. | [optional] 
**CodePage** | Pointer to **NullableInt32** | The encoding of a csv or txt source, as a code page number: 1251 for Cyrillic, 65001 for UTF-8, and so on.  Leaving it out lets the document service guess, which garbles non-latin text. | [optional] 
**Delimiter** | Pointer to **NullableInt32** | The character separating the values of a csv source: 0 - none, 1 - tab, 2 - semicolon, 3 - colon, 4 - comma,  5 - space. | [optional] 
**DocumentLayout** | Pointer to [**ConvertDocumentLayoutRequest**](ConvertDocumentLayoutRequest.md) | The layout of forms printed as pdf documents or images. | [optional] 
**DocumentRenderer** | Pointer to [**ConvertDocumentRendererRequest**](ConvertDocumentRendererRequest.md) | How the text of a pdf, xps or oxps source is read back. | [optional] 
**Outputtype** | **NullableString** | The extension of the format to convert to, without the dot. The formats a source can become are listed under  `extsConvertible` in `GET api/2.0/files/settings`, keyed by the source extension, both with a leading dot there. | 
**Password** | Pointer to **NullableString** | The password of the source document, for a document that is protected with one. | [optional] 
**Region** | Pointer to **NullableString** | The culture code whose currency, date and time formats a spreadsheet converted to PDF is shown in. | [optional] 
**Thumbnail** | Pointer to [**ConvertThumbnailRequest**](ConvertThumbnailRequest.md) | How a thumbnail is fitted, when the target format is an image. | [optional] 
**SpreadsheetLayout** | Pointer to [**ConvertSpreadsheetLayoutRequest**](ConvertSpreadsheetLayoutRequest.md) | How a spreadsheet is laid out on the page, when it is converted to PDF. | [optional] 
**Pdf** | Pointer to [**ConvertPdfRequest**](ConvertPdfRequest.md) | The PDF settings, used when the target format is PDF. | [optional] 

## Methods

### NewConvertFileRequestDto

`func NewConvertFileRequestDto(fileId int32, outputtype NullableString, ) *ConvertFileRequestDto`

NewConvertFileRequestDto instantiates a new ConvertFileRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvertFileRequestDtoWithDefaults

`func NewConvertFileRequestDtoWithDefaults() *ConvertFileRequestDto`

NewConvertFileRequestDtoWithDefaults instantiates a new ConvertFileRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileId

`func (o *ConvertFileRequestDto) GetFileId() int32`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *ConvertFileRequestDto) GetFileIdOk() (*int32, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *ConvertFileRequestDto) SetFileId(v int32)`

SetFileId sets FileId field to given value.


### GetFolderId

`func (o *ConvertFileRequestDto) GetFolderId() int32`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *ConvertFileRequestDto) GetFolderIdOk() (*int32, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *ConvertFileRequestDto) SetFolderId(v int32)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *ConvertFileRequestDto) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### SetFolderIdNil

`func (o *ConvertFileRequestDto) SetFolderIdNil(b bool)`

 SetFolderIdNil sets the value for FolderId to be an explicit nil

### UnsetFolderId
`func (o *ConvertFileRequestDto) UnsetFolderId()`

UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
### GetCodePage

`func (o *ConvertFileRequestDto) GetCodePage() int32`

GetCodePage returns the CodePage field if non-nil, zero value otherwise.

### GetCodePageOk

`func (o *ConvertFileRequestDto) GetCodePageOk() (*int32, bool)`

GetCodePageOk returns a tuple with the CodePage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodePage

`func (o *ConvertFileRequestDto) SetCodePage(v int32)`

SetCodePage sets CodePage field to given value.

### HasCodePage

`func (o *ConvertFileRequestDto) HasCodePage() bool`

HasCodePage returns a boolean if a field has been set.

### SetCodePageNil

`func (o *ConvertFileRequestDto) SetCodePageNil(b bool)`

 SetCodePageNil sets the value for CodePage to be an explicit nil

### UnsetCodePage
`func (o *ConvertFileRequestDto) UnsetCodePage()`

UnsetCodePage ensures that no value is present for CodePage, not even an explicit nil
### GetDelimiter

`func (o *ConvertFileRequestDto) GetDelimiter() int32`

GetDelimiter returns the Delimiter field if non-nil, zero value otherwise.

### GetDelimiterOk

`func (o *ConvertFileRequestDto) GetDelimiterOk() (*int32, bool)`

GetDelimiterOk returns a tuple with the Delimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelimiter

`func (o *ConvertFileRequestDto) SetDelimiter(v int32)`

SetDelimiter sets Delimiter field to given value.

### HasDelimiter

`func (o *ConvertFileRequestDto) HasDelimiter() bool`

HasDelimiter returns a boolean if a field has been set.

### SetDelimiterNil

`func (o *ConvertFileRequestDto) SetDelimiterNil(b bool)`

 SetDelimiterNil sets the value for Delimiter to be an explicit nil

### UnsetDelimiter
`func (o *ConvertFileRequestDto) UnsetDelimiter()`

UnsetDelimiter ensures that no value is present for Delimiter, not even an explicit nil
### GetDocumentLayout

`func (o *ConvertFileRequestDto) GetDocumentLayout() ConvertDocumentLayoutRequest`

GetDocumentLayout returns the DocumentLayout field if non-nil, zero value otherwise.

### GetDocumentLayoutOk

`func (o *ConvertFileRequestDto) GetDocumentLayoutOk() (*ConvertDocumentLayoutRequest, bool)`

GetDocumentLayoutOk returns a tuple with the DocumentLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentLayout

`func (o *ConvertFileRequestDto) SetDocumentLayout(v ConvertDocumentLayoutRequest)`

SetDocumentLayout sets DocumentLayout field to given value.

### HasDocumentLayout

`func (o *ConvertFileRequestDto) HasDocumentLayout() bool`

HasDocumentLayout returns a boolean if a field has been set.

### GetDocumentRenderer

`func (o *ConvertFileRequestDto) GetDocumentRenderer() ConvertDocumentRendererRequest`

GetDocumentRenderer returns the DocumentRenderer field if non-nil, zero value otherwise.

### GetDocumentRendererOk

`func (o *ConvertFileRequestDto) GetDocumentRendererOk() (*ConvertDocumentRendererRequest, bool)`

GetDocumentRendererOk returns a tuple with the DocumentRenderer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentRenderer

`func (o *ConvertFileRequestDto) SetDocumentRenderer(v ConvertDocumentRendererRequest)`

SetDocumentRenderer sets DocumentRenderer field to given value.

### HasDocumentRenderer

`func (o *ConvertFileRequestDto) HasDocumentRenderer() bool`

HasDocumentRenderer returns a boolean if a field has been set.

### GetOutputtype

`func (o *ConvertFileRequestDto) GetOutputtype() string`

GetOutputtype returns the Outputtype field if non-nil, zero value otherwise.

### GetOutputtypeOk

`func (o *ConvertFileRequestDto) GetOutputtypeOk() (*string, bool)`

GetOutputtypeOk returns a tuple with the Outputtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputtype

`func (o *ConvertFileRequestDto) SetOutputtype(v string)`

SetOutputtype sets Outputtype field to given value.


### SetOutputtypeNil

`func (o *ConvertFileRequestDto) SetOutputtypeNil(b bool)`

 SetOutputtypeNil sets the value for Outputtype to be an explicit nil

### UnsetOutputtype
`func (o *ConvertFileRequestDto) UnsetOutputtype()`

UnsetOutputtype ensures that no value is present for Outputtype, not even an explicit nil
### GetPassword

`func (o *ConvertFileRequestDto) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ConvertFileRequestDto) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ConvertFileRequestDto) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ConvertFileRequestDto) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *ConvertFileRequestDto) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *ConvertFileRequestDto) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetRegion

`func (o *ConvertFileRequestDto) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ConvertFileRequestDto) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ConvertFileRequestDto) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ConvertFileRequestDto) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *ConvertFileRequestDto) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *ConvertFileRequestDto) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetThumbnail

`func (o *ConvertFileRequestDto) GetThumbnail() ConvertThumbnailRequest`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *ConvertFileRequestDto) GetThumbnailOk() (*ConvertThumbnailRequest, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *ConvertFileRequestDto) SetThumbnail(v ConvertThumbnailRequest)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *ConvertFileRequestDto) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### GetSpreadsheetLayout

`func (o *ConvertFileRequestDto) GetSpreadsheetLayout() ConvertSpreadsheetLayoutRequest`

GetSpreadsheetLayout returns the SpreadsheetLayout field if non-nil, zero value otherwise.

### GetSpreadsheetLayoutOk

`func (o *ConvertFileRequestDto) GetSpreadsheetLayoutOk() (*ConvertSpreadsheetLayoutRequest, bool)`

GetSpreadsheetLayoutOk returns a tuple with the SpreadsheetLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpreadsheetLayout

`func (o *ConvertFileRequestDto) SetSpreadsheetLayout(v ConvertSpreadsheetLayoutRequest)`

SetSpreadsheetLayout sets SpreadsheetLayout field to given value.

### HasSpreadsheetLayout

`func (o *ConvertFileRequestDto) HasSpreadsheetLayout() bool`

HasSpreadsheetLayout returns a boolean if a field has been set.

### GetPdf

`func (o *ConvertFileRequestDto) GetPdf() ConvertPdfRequest`

GetPdf returns the Pdf field if non-nil, zero value otherwise.

### GetPdfOk

`func (o *ConvertFileRequestDto) GetPdfOk() (*ConvertPdfRequest, bool)`

GetPdfOk returns a tuple with the Pdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdf

`func (o *ConvertFileRequestDto) SetPdf(v ConvertPdfRequest)`

SetPdf sets Pdf field to given value.

### HasPdf

`func (o *ConvertFileRequestDto) HasPdf() bool`

HasPdf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


