# Paragraph

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Align** | Pointer to **int32** | The paragraph align. | [optional] 
**Runs** | Pointer to [**[]Run**](Run.md) | The list of text runs from the paragraph. | [optional] 

## Methods

### NewParagraph

`func NewParagraph() *Paragraph`

NewParagraph instantiates a new Paragraph object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParagraphWithDefaults

`func NewParagraphWithDefaults() *Paragraph`

NewParagraphWithDefaults instantiates a new Paragraph object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlign

`func (o *Paragraph) GetAlign() int32`

GetAlign returns the Align field if non-nil, zero value otherwise.

### GetAlignOk

`func (o *Paragraph) GetAlignOk() (*int32, bool)`

GetAlignOk returns a tuple with the Align field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlign

`func (o *Paragraph) SetAlign(v int32)`

SetAlign sets Align field to given value.

### HasAlign

`func (o *Paragraph) HasAlign() bool`

HasAlign returns a boolean if a field has been set.

### GetRuns

`func (o *Paragraph) GetRuns() []Run`

GetRuns returns the Runs field if non-nil, zero value otherwise.

### GetRunsOk

`func (o *Paragraph) GetRunsOk() (*[]Run, bool)`

GetRunsOk returns a tuple with the Runs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuns

`func (o *Paragraph) SetRuns(v []Run)`

SetRuns sets Runs field to given value.

### HasRuns

`func (o *Paragraph) HasRuns() bool`

HasRuns returns a boolean if a field has been set.

### SetRunsNil

`func (o *Paragraph) SetRunsNil(b bool)`

 SetRunsNil sets the value for Runs to be an explicit nil

### UnsetRuns
`func (o *Paragraph) UnsetRuns()`

UnsetRuns ensures that no value is present for Runs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


