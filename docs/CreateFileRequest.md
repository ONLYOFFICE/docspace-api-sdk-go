# CreateFileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **NullableString** | The title of the new file. The extension in it decides the format, and one of a known text, spreadsheet or  presentation format is rewritten to the DOCX, XLSX or PPTX of the portal unless `enableExternalExt` says  otherwise; a title with no extension gets DOCX added. | 
**TemplateId** | Pointer to [**CreateFileRequestTemplateId**](CreateFileRequestTemplateId.md) |  | [optional] 
**EnableExternalExt** | Pointer to **bool** | Whether the extension of the title is kept as it is: `true` stores the title verbatim, `false` rewrites a  known foreign format to the format the portal edits itself. | [optional] 
**FormId** | Pointer to **int32** | A ready form from the form gallery of the portal to copy instead of a template, named by the identifier the  gallery reports for it. It takes precedence over `templateId`; 0 means no form. | [optional] 

## Methods

### NewCreateFileRequest

`func NewCreateFileRequest(title NullableString, ) *CreateFileRequest`

NewCreateFileRequest instantiates a new CreateFileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFileRequestWithDefaults

`func NewCreateFileRequestWithDefaults() *CreateFileRequest`

NewCreateFileRequestWithDefaults instantiates a new CreateFileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CreateFileRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateFileRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateFileRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *CreateFileRequest) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CreateFileRequest) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetTemplateId

`func (o *CreateFileRequest) GetTemplateId() CreateFileRequestTemplateId`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *CreateFileRequest) GetTemplateIdOk() (*CreateFileRequestTemplateId, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *CreateFileRequest) SetTemplateId(v CreateFileRequestTemplateId)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *CreateFileRequest) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetEnableExternalExt

`func (o *CreateFileRequest) GetEnableExternalExt() bool`

GetEnableExternalExt returns the EnableExternalExt field if non-nil, zero value otherwise.

### GetEnableExternalExtOk

`func (o *CreateFileRequest) GetEnableExternalExtOk() (*bool, bool)`

GetEnableExternalExtOk returns a tuple with the EnableExternalExt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableExternalExt

`func (o *CreateFileRequest) SetEnableExternalExt(v bool)`

SetEnableExternalExt sets EnableExternalExt field to given value.

### HasEnableExternalExt

`func (o *CreateFileRequest) HasEnableExternalExt() bool`

HasEnableExternalExt returns a boolean if a field has been set.

### GetFormId

`func (o *CreateFileRequest) GetFormId() int32`

GetFormId returns the FormId field if non-nil, zero value otherwise.

### GetFormIdOk

`func (o *CreateFileRequest) GetFormIdOk() (*int32, bool)`

GetFormIdOk returns a tuple with the FormId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormId

`func (o *CreateFileRequest) SetFormId(v int32)`

SetFormId sets FormId field to given value.

### HasFormId

`func (o *CreateFileRequest) HasFormId() bool`

HasFormId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


