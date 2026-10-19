# ConvertPdfRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Form** | Pointer to **bool** | Whether the result is a fillable PDF form: true keeps the form fields of the source as fields to fill in,  false saves a plain PDF. | [optional] 

## Methods

### NewConvertPdfRequest

`func NewConvertPdfRequest() *ConvertPdfRequest`

NewConvertPdfRequest instantiates a new ConvertPdfRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvertPdfRequestWithDefaults

`func NewConvertPdfRequestWithDefaults() *ConvertPdfRequest`

NewConvertPdfRequestWithDefaults instantiates a new ConvertPdfRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForm

`func (o *ConvertPdfRequest) GetForm() bool`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *ConvertPdfRequest) GetFormOk() (*bool, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *ConvertPdfRequest) SetForm(v bool)`

SetForm sets Form field to given value.

### HasForm

`func (o *ConvertPdfRequest) HasForm() bool`

HasForm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


