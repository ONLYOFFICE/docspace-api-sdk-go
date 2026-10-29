# AiMetaGenerate200ResponseFieldsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Field name, at most 255 characters. | 
**Value** | **NullableString** | The value found in the documents, or null when the text does not state one. | 

## Methods

### NewAiMetaGenerate200ResponseFieldsInner

`func NewAiMetaGenerate200ResponseFieldsInner(name string, value NullableString, ) *AiMetaGenerate200ResponseFieldsInner`

NewAiMetaGenerate200ResponseFieldsInner instantiates a new AiMetaGenerate200ResponseFieldsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiMetaGenerate200ResponseFieldsInnerWithDefaults

`func NewAiMetaGenerate200ResponseFieldsInnerWithDefaults() *AiMetaGenerate200ResponseFieldsInner`

NewAiMetaGenerate200ResponseFieldsInnerWithDefaults instantiates a new AiMetaGenerate200ResponseFieldsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AiMetaGenerate200ResponseFieldsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AiMetaGenerate200ResponseFieldsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AiMetaGenerate200ResponseFieldsInner) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *AiMetaGenerate200ResponseFieldsInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AiMetaGenerate200ResponseFieldsInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AiMetaGenerate200ResponseFieldsInner) SetValue(v string)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *AiMetaGenerate200ResponseFieldsInner) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *AiMetaGenerate200ResponseFieldsInner) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


