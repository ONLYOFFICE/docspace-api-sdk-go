# SubmitFormDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Visible** | Pointer to **bool** | Specifies whether the Complete  & Submit button will be displayed or hidden on the top toolbar. | [optional] 
**ResultMessage** | Pointer to **NullableString** | A message displayed after forms are submitted. | [optional] 

## Methods

### NewSubmitFormDto

`func NewSubmitFormDto() *SubmitFormDto`

NewSubmitFormDto instantiates a new SubmitFormDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitFormDtoWithDefaults

`func NewSubmitFormDtoWithDefaults() *SubmitFormDto`

NewSubmitFormDtoWithDefaults instantiates a new SubmitFormDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVisible

`func (o *SubmitFormDto) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *SubmitFormDto) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *SubmitFormDto) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *SubmitFormDto) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetResultMessage

`func (o *SubmitFormDto) GetResultMessage() string`

GetResultMessage returns the ResultMessage field if non-nil, zero value otherwise.

### GetResultMessageOk

`func (o *SubmitFormDto) GetResultMessageOk() (*string, bool)`

GetResultMessageOk returns a tuple with the ResultMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultMessage

`func (o *SubmitFormDto) SetResultMessage(v string)`

SetResultMessage sets ResultMessage field to given value.

### HasResultMessage

`func (o *SubmitFormDto) HasResultMessage() bool`

HasResultMessage returns a boolean if a field has been set.

### SetResultMessageNil

`func (o *SubmitFormDto) SetResultMessageNil(b bool)`

 SetResultMessageNil sets the value for ResultMessage to be an explicit nil

### UnsetResultMessage
`func (o *SubmitFormDto) UnsetResultMessage()`

UnsetResultMessage ensures that no value is present for ResultMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


