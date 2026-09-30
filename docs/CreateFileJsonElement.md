# CreateFileJsonElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **NullableString** | The title of the new file. The extension in it decides the format, and one of a known text, spreadsheet or  presentation format is rewritten to the DOCX, XLSX or PPTX of the portal unless `enableExternalExt` says  otherwise; a title with no extension gets DOCX added. | 
**TemplateId** | Pointer to [**CreateFileJsonElementTemplateId**](CreateFileJsonElementTemplateId.md) |  | [optional] 
**EnableExternalExt** | Pointer to **bool** | Whether the extension of the title is kept as it is: `true` stores the title verbatim, `false` rewrites a  known foreign format to the format the portal edits itself. | [optional] 
**FormId** | Pointer to **int32** | A ready form from the form gallery of the portal to copy instead of a template, named by the identifier the  gallery reports for it. It takes precedence over `templateId`; 0 means no form. | [optional] 

## Methods

### NewCreateFileJsonElement

`func NewCreateFileJsonElement(title NullableString, ) *CreateFileJsonElement`

NewCreateFileJsonElement instantiates a new CreateFileJsonElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFileJsonElementWithDefaults

`func NewCreateFileJsonElementWithDefaults() *CreateFileJsonElement`

NewCreateFileJsonElementWithDefaults instantiates a new CreateFileJsonElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CreateFileJsonElement) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateFileJsonElement) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateFileJsonElement) SetTitle(v string)`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *CreateFileJsonElement) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CreateFileJsonElement) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetTemplateId

`func (o *CreateFileJsonElement) GetTemplateId() CreateFileJsonElementTemplateId`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *CreateFileJsonElement) GetTemplateIdOk() (*CreateFileJsonElementTemplateId, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *CreateFileJsonElement) SetTemplateId(v CreateFileJsonElementTemplateId)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *CreateFileJsonElement) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetEnableExternalExt

`func (o *CreateFileJsonElement) GetEnableExternalExt() bool`

GetEnableExternalExt returns the EnableExternalExt field if non-nil, zero value otherwise.

### GetEnableExternalExtOk

`func (o *CreateFileJsonElement) GetEnableExternalExtOk() (*bool, bool)`

GetEnableExternalExtOk returns a tuple with the EnableExternalExt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableExternalExt

`func (o *CreateFileJsonElement) SetEnableExternalExt(v bool)`

SetEnableExternalExt sets EnableExternalExt field to given value.

### HasEnableExternalExt

`func (o *CreateFileJsonElement) HasEnableExternalExt() bool`

HasEnableExternalExt returns a boolean if a field has been set.

### GetFormId

`func (o *CreateFileJsonElement) GetFormId() int32`

GetFormId returns the FormId field if non-nil, zero value otherwise.

### GetFormIdOk

`func (o *CreateFileJsonElement) GetFormIdOk() (*int32, bool)`

GetFormIdOk returns a tuple with the FormId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormId

`func (o *CreateFileJsonElement) SetFormId(v int32)`

SetFormId sets FormId field to given value.

### HasFormId

`func (o *CreateFileJsonElement) HasFormId() bool`

HasFormId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


