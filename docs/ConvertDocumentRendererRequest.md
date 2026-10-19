# ConvertDocumentRendererRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TextAssociation** | Pointer to **NullableString** | The rendering mode: blockChar reads the text by single characters, blockLine by separate lines, plainLine as  plain text with a paragraph per line, plainParagraph as plain text with lines combined into paragraphs. | [optional] 

## Methods

### NewConvertDocumentRendererRequest

`func NewConvertDocumentRendererRequest() *ConvertDocumentRendererRequest`

NewConvertDocumentRendererRequest instantiates a new ConvertDocumentRendererRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvertDocumentRendererRequestWithDefaults

`func NewConvertDocumentRendererRequestWithDefaults() *ConvertDocumentRendererRequest`

NewConvertDocumentRendererRequestWithDefaults instantiates a new ConvertDocumentRendererRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTextAssociation

`func (o *ConvertDocumentRendererRequest) GetTextAssociation() string`

GetTextAssociation returns the TextAssociation field if non-nil, zero value otherwise.

### GetTextAssociationOk

`func (o *ConvertDocumentRendererRequest) GetTextAssociationOk() (*string, bool)`

GetTextAssociationOk returns a tuple with the TextAssociation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextAssociation

`func (o *ConvertDocumentRendererRequest) SetTextAssociation(v string)`

SetTextAssociation sets TextAssociation field to given value.

### HasTextAssociation

`func (o *ConvertDocumentRendererRequest) HasTextAssociation() bool`

HasTextAssociation returns a boolean if a field has been set.

### SetTextAssociationNil

`func (o *ConvertDocumentRendererRequest) SetTextAssociationNil(b bool)`

 SetTextAssociationNil sets the value for TextAssociation to be an explicit nil

### UnsetTextAssociation
`func (o *ConvertDocumentRendererRequest) UnsetTextAssociation()`

UnsetTextAssociation ensures that no value is present for TextAssociation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


