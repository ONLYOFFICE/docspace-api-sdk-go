# ConvertSpreadsheetLayoutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IgnorePrintArea** | Pointer to **NullableBool** | Whether the print area chosen in the spreadsheet is ignored; false converts only that area. | [optional] 
**Orientation** | Pointer to **NullableString** | The page orientation of the output PDF file: portrait or landscape. | [optional] 
**FitToHeight** | Pointer to **NullableInt32** | The number of pages the converted area is fitted into in height; 0 does not fit it. | [optional] 
**Scale** | Pointer to **NullableInt32** | The scale of the output PDF file, in percent. | [optional] 
**FitToWidth** | Pointer to **NullableInt32** | The number of pages the converted area is fitted into in width; 0 does not fit it. | [optional] 
**Headings** | Pointer to **NullableBool** | Whether the row and column headings are printed in the output PDF file. | [optional] 
**GridLines** | Pointer to **NullableBool** | Whether the grid lines are printed in the output PDF file. | [optional] 
**Margins** | Pointer to [**ConvertSpreadsheetMarginsRequest**](ConvertSpreadsheetMarginsRequest.md) | The margins of the output PDF file. | [optional] 
**PageSize** | Pointer to [**ConvertSpreadsheetPageSizeRequest**](ConvertSpreadsheetPageSizeRequest.md) | The page size of the output PDF file. | [optional] 

## Methods

### NewConvertSpreadsheetLayoutRequest

`func NewConvertSpreadsheetLayoutRequest() *ConvertSpreadsheetLayoutRequest`

NewConvertSpreadsheetLayoutRequest instantiates a new ConvertSpreadsheetLayoutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvertSpreadsheetLayoutRequestWithDefaults

`func NewConvertSpreadsheetLayoutRequestWithDefaults() *ConvertSpreadsheetLayoutRequest`

NewConvertSpreadsheetLayoutRequestWithDefaults instantiates a new ConvertSpreadsheetLayoutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIgnorePrintArea

`func (o *ConvertSpreadsheetLayoutRequest) GetIgnorePrintArea() bool`

GetIgnorePrintArea returns the IgnorePrintArea field if non-nil, zero value otherwise.

### GetIgnorePrintAreaOk

`func (o *ConvertSpreadsheetLayoutRequest) GetIgnorePrintAreaOk() (*bool, bool)`

GetIgnorePrintAreaOk returns a tuple with the IgnorePrintArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnorePrintArea

`func (o *ConvertSpreadsheetLayoutRequest) SetIgnorePrintArea(v bool)`

SetIgnorePrintArea sets IgnorePrintArea field to given value.

### HasIgnorePrintArea

`func (o *ConvertSpreadsheetLayoutRequest) HasIgnorePrintArea() bool`

HasIgnorePrintArea returns a boolean if a field has been set.

### SetIgnorePrintAreaNil

`func (o *ConvertSpreadsheetLayoutRequest) SetIgnorePrintAreaNil(b bool)`

 SetIgnorePrintAreaNil sets the value for IgnorePrintArea to be an explicit nil

### UnsetIgnorePrintArea
`func (o *ConvertSpreadsheetLayoutRequest) UnsetIgnorePrintArea()`

UnsetIgnorePrintArea ensures that no value is present for IgnorePrintArea, not even an explicit nil
### GetOrientation

`func (o *ConvertSpreadsheetLayoutRequest) GetOrientation() string`

GetOrientation returns the Orientation field if non-nil, zero value otherwise.

### GetOrientationOk

`func (o *ConvertSpreadsheetLayoutRequest) GetOrientationOk() (*string, bool)`

GetOrientationOk returns a tuple with the Orientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientation

`func (o *ConvertSpreadsheetLayoutRequest) SetOrientation(v string)`

SetOrientation sets Orientation field to given value.

### HasOrientation

`func (o *ConvertSpreadsheetLayoutRequest) HasOrientation() bool`

HasOrientation returns a boolean if a field has been set.

### SetOrientationNil

`func (o *ConvertSpreadsheetLayoutRequest) SetOrientationNil(b bool)`

 SetOrientationNil sets the value for Orientation to be an explicit nil

### UnsetOrientation
`func (o *ConvertSpreadsheetLayoutRequest) UnsetOrientation()`

UnsetOrientation ensures that no value is present for Orientation, not even an explicit nil
### GetFitToHeight

`func (o *ConvertSpreadsheetLayoutRequest) GetFitToHeight() int32`

GetFitToHeight returns the FitToHeight field if non-nil, zero value otherwise.

### GetFitToHeightOk

`func (o *ConvertSpreadsheetLayoutRequest) GetFitToHeightOk() (*int32, bool)`

GetFitToHeightOk returns a tuple with the FitToHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFitToHeight

`func (o *ConvertSpreadsheetLayoutRequest) SetFitToHeight(v int32)`

SetFitToHeight sets FitToHeight field to given value.

### HasFitToHeight

`func (o *ConvertSpreadsheetLayoutRequest) HasFitToHeight() bool`

HasFitToHeight returns a boolean if a field has been set.

### SetFitToHeightNil

`func (o *ConvertSpreadsheetLayoutRequest) SetFitToHeightNil(b bool)`

 SetFitToHeightNil sets the value for FitToHeight to be an explicit nil

### UnsetFitToHeight
`func (o *ConvertSpreadsheetLayoutRequest) UnsetFitToHeight()`

UnsetFitToHeight ensures that no value is present for FitToHeight, not even an explicit nil
### GetScale

`func (o *ConvertSpreadsheetLayoutRequest) GetScale() int32`

GetScale returns the Scale field if non-nil, zero value otherwise.

### GetScaleOk

`func (o *ConvertSpreadsheetLayoutRequest) GetScaleOk() (*int32, bool)`

GetScaleOk returns a tuple with the Scale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScale

`func (o *ConvertSpreadsheetLayoutRequest) SetScale(v int32)`

SetScale sets Scale field to given value.

### HasScale

`func (o *ConvertSpreadsheetLayoutRequest) HasScale() bool`

HasScale returns a boolean if a field has been set.

### SetScaleNil

`func (o *ConvertSpreadsheetLayoutRequest) SetScaleNil(b bool)`

 SetScaleNil sets the value for Scale to be an explicit nil

### UnsetScale
`func (o *ConvertSpreadsheetLayoutRequest) UnsetScale()`

UnsetScale ensures that no value is present for Scale, not even an explicit nil
### GetFitToWidth

`func (o *ConvertSpreadsheetLayoutRequest) GetFitToWidth() int32`

GetFitToWidth returns the FitToWidth field if non-nil, zero value otherwise.

### GetFitToWidthOk

`func (o *ConvertSpreadsheetLayoutRequest) GetFitToWidthOk() (*int32, bool)`

GetFitToWidthOk returns a tuple with the FitToWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFitToWidth

`func (o *ConvertSpreadsheetLayoutRequest) SetFitToWidth(v int32)`

SetFitToWidth sets FitToWidth field to given value.

### HasFitToWidth

`func (o *ConvertSpreadsheetLayoutRequest) HasFitToWidth() bool`

HasFitToWidth returns a boolean if a field has been set.

### SetFitToWidthNil

`func (o *ConvertSpreadsheetLayoutRequest) SetFitToWidthNil(b bool)`

 SetFitToWidthNil sets the value for FitToWidth to be an explicit nil

### UnsetFitToWidth
`func (o *ConvertSpreadsheetLayoutRequest) UnsetFitToWidth()`

UnsetFitToWidth ensures that no value is present for FitToWidth, not even an explicit nil
### GetHeadings

`func (o *ConvertSpreadsheetLayoutRequest) GetHeadings() bool`

GetHeadings returns the Headings field if non-nil, zero value otherwise.

### GetHeadingsOk

`func (o *ConvertSpreadsheetLayoutRequest) GetHeadingsOk() (*bool, bool)`

GetHeadingsOk returns a tuple with the Headings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadings

`func (o *ConvertSpreadsheetLayoutRequest) SetHeadings(v bool)`

SetHeadings sets Headings field to given value.

### HasHeadings

`func (o *ConvertSpreadsheetLayoutRequest) HasHeadings() bool`

HasHeadings returns a boolean if a field has been set.

### SetHeadingsNil

`func (o *ConvertSpreadsheetLayoutRequest) SetHeadingsNil(b bool)`

 SetHeadingsNil sets the value for Headings to be an explicit nil

### UnsetHeadings
`func (o *ConvertSpreadsheetLayoutRequest) UnsetHeadings()`

UnsetHeadings ensures that no value is present for Headings, not even an explicit nil
### GetGridLines

`func (o *ConvertSpreadsheetLayoutRequest) GetGridLines() bool`

GetGridLines returns the GridLines field if non-nil, zero value otherwise.

### GetGridLinesOk

`func (o *ConvertSpreadsheetLayoutRequest) GetGridLinesOk() (*bool, bool)`

GetGridLinesOk returns a tuple with the GridLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridLines

`func (o *ConvertSpreadsheetLayoutRequest) SetGridLines(v bool)`

SetGridLines sets GridLines field to given value.

### HasGridLines

`func (o *ConvertSpreadsheetLayoutRequest) HasGridLines() bool`

HasGridLines returns a boolean if a field has been set.

### SetGridLinesNil

`func (o *ConvertSpreadsheetLayoutRequest) SetGridLinesNil(b bool)`

 SetGridLinesNil sets the value for GridLines to be an explicit nil

### UnsetGridLines
`func (o *ConvertSpreadsheetLayoutRequest) UnsetGridLines()`

UnsetGridLines ensures that no value is present for GridLines, not even an explicit nil
### GetMargins

`func (o *ConvertSpreadsheetLayoutRequest) GetMargins() ConvertSpreadsheetMarginsRequest`

GetMargins returns the Margins field if non-nil, zero value otherwise.

### GetMarginsOk

`func (o *ConvertSpreadsheetLayoutRequest) GetMarginsOk() (*ConvertSpreadsheetMarginsRequest, bool)`

GetMarginsOk returns a tuple with the Margins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMargins

`func (o *ConvertSpreadsheetLayoutRequest) SetMargins(v ConvertSpreadsheetMarginsRequest)`

SetMargins sets Margins field to given value.

### HasMargins

`func (o *ConvertSpreadsheetLayoutRequest) HasMargins() bool`

HasMargins returns a boolean if a field has been set.

### GetPageSize

`func (o *ConvertSpreadsheetLayoutRequest) GetPageSize() ConvertSpreadsheetPageSizeRequest`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ConvertSpreadsheetLayoutRequest) GetPageSizeOk() (*ConvertSpreadsheetPageSizeRequest, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ConvertSpreadsheetLayoutRequest) SetPageSize(v ConvertSpreadsheetPageSizeRequest)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ConvertSpreadsheetLayoutRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


