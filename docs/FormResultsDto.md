# FormResultsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateOn** | Pointer to **time.Time** | When the portal recorded this copy, in UTC: the moment the filled copy was completed and its data indexed, not  the moment the form itself was made. | [optional] 
**FormsData** | Pointer to [**[]FormsItemData**](FormsItemData.md) | The values that were entered into this copy, one entry per field, preceded by an entry keyed `FormNumber` that  carries the number of the copy and is what the submissions are ordered by. Fields holding a picture or a  signature are left out of the record, so a field missing here was not necessarily left blank. | [optional] 

## Methods

### NewFormResultsDto

`func NewFormResultsDto() *FormResultsDto`

NewFormResultsDto instantiates a new FormResultsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormResultsDtoWithDefaults

`func NewFormResultsDtoWithDefaults() *FormResultsDto`

NewFormResultsDtoWithDefaults instantiates a new FormResultsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateOn

`func (o *FormResultsDto) GetCreateOn() time.Time`

GetCreateOn returns the CreateOn field if non-nil, zero value otherwise.

### GetCreateOnOk

`func (o *FormResultsDto) GetCreateOnOk() (*time.Time, bool)`

GetCreateOnOk returns a tuple with the CreateOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateOn

`func (o *FormResultsDto) SetCreateOn(v time.Time)`

SetCreateOn sets CreateOn field to given value.

### HasCreateOn

`func (o *FormResultsDto) HasCreateOn() bool`

HasCreateOn returns a boolean if a field has been set.

### GetFormsData

`func (o *FormResultsDto) GetFormsData() []FormsItemData`

GetFormsData returns the FormsData field if non-nil, zero value otherwise.

### GetFormsDataOk

`func (o *FormResultsDto) GetFormsDataOk() (*[]FormsItemData, bool)`

GetFormsDataOk returns a tuple with the FormsData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormsData

`func (o *FormResultsDto) SetFormsData(v []FormsItemData)`

SetFormsData sets FormsData field to given value.

### HasFormsData

`func (o *FormResultsDto) HasFormsData() bool`

HasFormsData returns a boolean if a field has been set.

### SetFormsDataNil

`func (o *FormResultsDto) SetFormsDataNil(b bool)`

 SetFormsDataNil sets the value for FormsData to be an explicit nil

### UnsetFormsData
`func (o *FormResultsDto) UnsetFormsData()`

UnsetFormsData ensures that no value is present for FormsData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


