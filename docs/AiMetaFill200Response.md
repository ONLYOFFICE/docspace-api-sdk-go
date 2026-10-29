# AiMetaFill200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | **bool** | Whether the values were only proposed (`true`) or written as well. | 
**Metadata** | [**AiEntryMetadata**](AiEntryMetadata.md) | The entry's metadata after the write (or as it stands, with `dryRun`), as `GET api/2.0/files/metadata/{kind}/{id}` returns it. | 
**Filled** | [**[]AiMetaFill200ResponseFilledInner**](AiMetaFill200ResponseFilledInner.md) | The values written, or proposed with `dryRun`: a typed value for a template field, a string for a custom field. | 
**Unresolved** | [**[]AiMetaFill200ResponseUnresolvedInner**](AiMetaFill200ResponseUnresolvedInner.md) | The fields left as they were, with the reason for each. | 
**Sources** | [**[]AiMetaFill200ResponseSourcesInner**](AiMetaFill200ResponseSourcesInner.md) | The documents the values were read from. | 
**Skipped** | [**[]AiMetaGenerate200ResponseSkippedInner**](AiMetaGenerate200ResponseSkippedInner.md) | The entries that were left out, with the reason for each. | 

## Methods

### NewAiMetaFill200Response

`func NewAiMetaFill200Response(dryRun bool, metadata AiEntryMetadata, filled []AiMetaFill200ResponseFilledInner, unresolved []AiMetaFill200ResponseUnresolvedInner, sources []AiMetaFill200ResponseSourcesInner, skipped []AiMetaGenerate200ResponseSkippedInner, ) *AiMetaFill200Response`

NewAiMetaFill200Response instantiates a new AiMetaFill200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiMetaFill200ResponseWithDefaults

`func NewAiMetaFill200ResponseWithDefaults() *AiMetaFill200Response`

NewAiMetaFill200ResponseWithDefaults instantiates a new AiMetaFill200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *AiMetaFill200Response) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *AiMetaFill200Response) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *AiMetaFill200Response) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.


### GetMetadata

`func (o *AiMetaFill200Response) GetMetadata() AiEntryMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AiMetaFill200Response) GetMetadataOk() (*AiEntryMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AiMetaFill200Response) SetMetadata(v AiEntryMetadata)`

SetMetadata sets Metadata field to given value.


### GetFilled

`func (o *AiMetaFill200Response) GetFilled() []AiMetaFill200ResponseFilledInner`

GetFilled returns the Filled field if non-nil, zero value otherwise.

### GetFilledOk

`func (o *AiMetaFill200Response) GetFilledOk() (*[]AiMetaFill200ResponseFilledInner, bool)`

GetFilledOk returns a tuple with the Filled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilled

`func (o *AiMetaFill200Response) SetFilled(v []AiMetaFill200ResponseFilledInner)`

SetFilled sets Filled field to given value.


### GetUnresolved

`func (o *AiMetaFill200Response) GetUnresolved() []AiMetaFill200ResponseUnresolvedInner`

GetUnresolved returns the Unresolved field if non-nil, zero value otherwise.

### GetUnresolvedOk

`func (o *AiMetaFill200Response) GetUnresolvedOk() (*[]AiMetaFill200ResponseUnresolvedInner, bool)`

GetUnresolvedOk returns a tuple with the Unresolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnresolved

`func (o *AiMetaFill200Response) SetUnresolved(v []AiMetaFill200ResponseUnresolvedInner)`

SetUnresolved sets Unresolved field to given value.


### GetSources

`func (o *AiMetaFill200Response) GetSources() []AiMetaFill200ResponseSourcesInner`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *AiMetaFill200Response) GetSourcesOk() (*[]AiMetaFill200ResponseSourcesInner, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *AiMetaFill200Response) SetSources(v []AiMetaFill200ResponseSourcesInner)`

SetSources sets Sources field to given value.


### GetSkipped

`func (o *AiMetaFill200Response) GetSkipped() []AiMetaGenerate200ResponseSkippedInner`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *AiMetaFill200Response) GetSkippedOk() (*[]AiMetaGenerate200ResponseSkippedInner, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *AiMetaFill200Response) SetSkipped(v []AiMetaGenerate200ResponseSkippedInner)`

SetSkipped sets Skipped field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


