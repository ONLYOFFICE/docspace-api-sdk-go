# AiPromptsImportBundleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bundle** | [**AiPromptBundle**](AiPromptBundle.md) | Bundle to restore. | 
**Options** | Pointer to [**AiPromptsImportBundleRequestOptions**](AiPromptsImportBundleRequestOptions.md) |  | [optional] 

## Methods

### NewAiPromptsImportBundleRequest

`func NewAiPromptsImportBundleRequest(bundle AiPromptBundle, ) *AiPromptsImportBundleRequest`

NewAiPromptsImportBundleRequest instantiates a new AiPromptsImportBundleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiPromptsImportBundleRequestWithDefaults

`func NewAiPromptsImportBundleRequestWithDefaults() *AiPromptsImportBundleRequest`

NewAiPromptsImportBundleRequestWithDefaults instantiates a new AiPromptsImportBundleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBundle

`func (o *AiPromptsImportBundleRequest) GetBundle() AiPromptBundle`

GetBundle returns the Bundle field if non-nil, zero value otherwise.

### GetBundleOk

`func (o *AiPromptsImportBundleRequest) GetBundleOk() (*AiPromptBundle, bool)`

GetBundleOk returns a tuple with the Bundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundle

`func (o *AiPromptsImportBundleRequest) SetBundle(v AiPromptBundle)`

SetBundle sets Bundle field to given value.


### GetOptions

`func (o *AiPromptsImportBundleRequest) GetOptions() AiPromptsImportBundleRequestOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AiPromptsImportBundleRequest) GetOptionsOk() (*AiPromptsImportBundleRequestOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AiPromptsImportBundleRequest) SetOptions(v AiPromptsImportBundleRequestOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *AiPromptsImportBundleRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


