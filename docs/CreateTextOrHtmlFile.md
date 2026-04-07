# CreateTextOrHtmlFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **NullableString** | The file title for text or HTML file. | 
**Content** | Pointer to **NullableString** | The text or HTML file contents. | [optional] 
**CreateNewIfExist** | Pointer to **bool** | Specifies whether to create a new text or HTML file if it exists or not. | [optional] 

## Methods

### NewCreateTextOrHtmlFile

`func NewCreateTextOrHtmlFile(title NullableString, ) *CreateTextOrHtmlFile`

NewCreateTextOrHtmlFile instantiates a new CreateTextOrHtmlFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTextOrHtmlFileWithDefaults

`func NewCreateTextOrHtmlFileWithDefaults() *CreateTextOrHtmlFile`

NewCreateTextOrHtmlFileWithDefaults instantiates a new CreateTextOrHtmlFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CreateTextOrHtmlFile) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateTextOrHtmlFile) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateTextOrHtmlFile) SetTitle(v string)`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *CreateTextOrHtmlFile) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CreateTextOrHtmlFile) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetContent

`func (o *CreateTextOrHtmlFile) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CreateTextOrHtmlFile) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CreateTextOrHtmlFile) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *CreateTextOrHtmlFile) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *CreateTextOrHtmlFile) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *CreateTextOrHtmlFile) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetCreateNewIfExist

`func (o *CreateTextOrHtmlFile) GetCreateNewIfExist() bool`

GetCreateNewIfExist returns the CreateNewIfExist field if non-nil, zero value otherwise.

### GetCreateNewIfExistOk

`func (o *CreateTextOrHtmlFile) GetCreateNewIfExistOk() (*bool, bool)`

GetCreateNewIfExistOk returns a tuple with the CreateNewIfExist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateNewIfExist

`func (o *CreateTextOrHtmlFile) SetCreateNewIfExist(v bool)`

SetCreateNewIfExist sets CreateNewIfExist field to given value.

### HasCreateNewIfExist

`func (o *CreateTextOrHtmlFile) HasCreateNewIfExist() bool`

HasCreateNewIfExist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


