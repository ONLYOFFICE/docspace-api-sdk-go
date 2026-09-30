# HideConfirmConvertRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Save** | Pointer to **bool** | Chooses the prompt to hide rather than the state to store: true hides the prompt that offers to keep a copy in  the original format when a document is converted, false hides the prompt that offers to open the conversion  result. Each of the two flags is stored separately for the calling account, and both are one-way - the portal  can hide a prompt but has no way to show it again. | [optional] 

## Methods

### NewHideConfirmConvertRequestDto

`func NewHideConfirmConvertRequestDto() *HideConfirmConvertRequestDto`

NewHideConfirmConvertRequestDto instantiates a new HideConfirmConvertRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHideConfirmConvertRequestDtoWithDefaults

`func NewHideConfirmConvertRequestDtoWithDefaults() *HideConfirmConvertRequestDto`

NewHideConfirmConvertRequestDtoWithDefaults instantiates a new HideConfirmConvertRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSave

`func (o *HideConfirmConvertRequestDto) GetSave() bool`

GetSave returns the Save field if non-nil, zero value otherwise.

### GetSaveOk

`func (o *HideConfirmConvertRequestDto) GetSaveOk() (*bool, bool)`

GetSaveOk returns a tuple with the Save field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSave

`func (o *HideConfirmConvertRequestDto) SetSave(v bool)`

SetSave sets Save field to given value.

### HasSave

`func (o *HideConfirmConvertRequestDto) HasSave() bool`

HasSave returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


