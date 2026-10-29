# AiMetaGenerate200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | [**[]AiMetaGenerate200ResponseFieldsInner**](AiMetaGenerate200ResponseFieldsInner.md) | The proposed custom metadata fields with the values the text supports. | 
**Sources** | [**[]AiMetaGenerate200ResponseSourcesInner**](AiMetaGenerate200ResponseSourcesInner.md) | The documents the proposal was read from. | 
**Skipped** | [**[]AiMetaGenerate200ResponseSkippedInner**](AiMetaGenerate200ResponseSkippedInner.md) | The entries that were left out, with the reason for each. | 

## Methods

### NewAiMetaGenerate200Response

`func NewAiMetaGenerate200Response(fields []AiMetaGenerate200ResponseFieldsInner, sources []AiMetaGenerate200ResponseSourcesInner, skipped []AiMetaGenerate200ResponseSkippedInner, ) *AiMetaGenerate200Response`

NewAiMetaGenerate200Response instantiates a new AiMetaGenerate200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiMetaGenerate200ResponseWithDefaults

`func NewAiMetaGenerate200ResponseWithDefaults() *AiMetaGenerate200Response`

NewAiMetaGenerate200ResponseWithDefaults instantiates a new AiMetaGenerate200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *AiMetaGenerate200Response) GetFields() []AiMetaGenerate200ResponseFieldsInner`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *AiMetaGenerate200Response) GetFieldsOk() (*[]AiMetaGenerate200ResponseFieldsInner, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *AiMetaGenerate200Response) SetFields(v []AiMetaGenerate200ResponseFieldsInner)`

SetFields sets Fields field to given value.


### GetSources

`func (o *AiMetaGenerate200Response) GetSources() []AiMetaGenerate200ResponseSourcesInner`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *AiMetaGenerate200Response) GetSourcesOk() (*[]AiMetaGenerate200ResponseSourcesInner, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *AiMetaGenerate200Response) SetSources(v []AiMetaGenerate200ResponseSourcesInner)`

SetSources sets Sources field to given value.


### GetSkipped

`func (o *AiMetaGenerate200Response) GetSkipped() []AiMetaGenerate200ResponseSkippedInner`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *AiMetaGenerate200Response) GetSkippedOk() (*[]AiMetaGenerate200ResponseSkippedInner, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *AiMetaGenerate200Response) SetSkipped(v []AiMetaGenerate200ResponseSkippedInner)`

SetSkipped sets Skipped field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


