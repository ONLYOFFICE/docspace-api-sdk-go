# ConvertDocumentLayoutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DrawPlaceHolders** | Pointer to **NullableBool** | Whether an empty form field shows its placeholder text in the result; false leaves it blank. | [optional] 
**DrawFormHighlight** | Pointer to **NullableBool** | Whether form fields keep their highlight colour in the result; false draws them without it. | [optional] 
**IsPrint** | Pointer to **NullableBool** | Whether the print mode is turned on. It only applies to a docx converted into pdf: with the print mode off  the highlight flag does nothing and the placeholder flag saves the forms in the pdf. | [optional] 

## Methods

### NewConvertDocumentLayoutRequest

`func NewConvertDocumentLayoutRequest() *ConvertDocumentLayoutRequest`

NewConvertDocumentLayoutRequest instantiates a new ConvertDocumentLayoutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvertDocumentLayoutRequestWithDefaults

`func NewConvertDocumentLayoutRequestWithDefaults() *ConvertDocumentLayoutRequest`

NewConvertDocumentLayoutRequestWithDefaults instantiates a new ConvertDocumentLayoutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDrawPlaceHolders

`func (o *ConvertDocumentLayoutRequest) GetDrawPlaceHolders() bool`

GetDrawPlaceHolders returns the DrawPlaceHolders field if non-nil, zero value otherwise.

### GetDrawPlaceHoldersOk

`func (o *ConvertDocumentLayoutRequest) GetDrawPlaceHoldersOk() (*bool, bool)`

GetDrawPlaceHoldersOk returns a tuple with the DrawPlaceHolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrawPlaceHolders

`func (o *ConvertDocumentLayoutRequest) SetDrawPlaceHolders(v bool)`

SetDrawPlaceHolders sets DrawPlaceHolders field to given value.

### HasDrawPlaceHolders

`func (o *ConvertDocumentLayoutRequest) HasDrawPlaceHolders() bool`

HasDrawPlaceHolders returns a boolean if a field has been set.

### SetDrawPlaceHoldersNil

`func (o *ConvertDocumentLayoutRequest) SetDrawPlaceHoldersNil(b bool)`

 SetDrawPlaceHoldersNil sets the value for DrawPlaceHolders to be an explicit nil

### UnsetDrawPlaceHolders
`func (o *ConvertDocumentLayoutRequest) UnsetDrawPlaceHolders()`

UnsetDrawPlaceHolders ensures that no value is present for DrawPlaceHolders, not even an explicit nil
### GetDrawFormHighlight

`func (o *ConvertDocumentLayoutRequest) GetDrawFormHighlight() bool`

GetDrawFormHighlight returns the DrawFormHighlight field if non-nil, zero value otherwise.

### GetDrawFormHighlightOk

`func (o *ConvertDocumentLayoutRequest) GetDrawFormHighlightOk() (*bool, bool)`

GetDrawFormHighlightOk returns a tuple with the DrawFormHighlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrawFormHighlight

`func (o *ConvertDocumentLayoutRequest) SetDrawFormHighlight(v bool)`

SetDrawFormHighlight sets DrawFormHighlight field to given value.

### HasDrawFormHighlight

`func (o *ConvertDocumentLayoutRequest) HasDrawFormHighlight() bool`

HasDrawFormHighlight returns a boolean if a field has been set.

### SetDrawFormHighlightNil

`func (o *ConvertDocumentLayoutRequest) SetDrawFormHighlightNil(b bool)`

 SetDrawFormHighlightNil sets the value for DrawFormHighlight to be an explicit nil

### UnsetDrawFormHighlight
`func (o *ConvertDocumentLayoutRequest) UnsetDrawFormHighlight()`

UnsetDrawFormHighlight ensures that no value is present for DrawFormHighlight, not even an explicit nil
### GetIsPrint

`func (o *ConvertDocumentLayoutRequest) GetIsPrint() bool`

GetIsPrint returns the IsPrint field if non-nil, zero value otherwise.

### GetIsPrintOk

`func (o *ConvertDocumentLayoutRequest) GetIsPrintOk() (*bool, bool)`

GetIsPrintOk returns a tuple with the IsPrint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrint

`func (o *ConvertDocumentLayoutRequest) SetIsPrint(v bool)`

SetIsPrint sets IsPrint field to given value.

### HasIsPrint

`func (o *ConvertDocumentLayoutRequest) HasIsPrint() bool`

HasIsPrint returns a boolean if a field has been set.

### SetIsPrintNil

`func (o *ConvertDocumentLayoutRequest) SetIsPrintNil(b bool)`

 SetIsPrintNil sets the value for IsPrint to be an explicit nil

### UnsetIsPrint
`func (o *ConvertDocumentLayoutRequest) UnsetIsPrint()`

UnsetIsPrint ensures that no value is present for IsPrint, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


