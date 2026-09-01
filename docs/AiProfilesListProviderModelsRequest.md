# AiProfilesListProviderModelsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderType** | [**AiProviderType**](AiProviderType.md) | Provider whose catalog to list. | 
**BaseUrl** | **string** | Provider API base URL. | 
**ApiKey** | **string** | Provider API key. | 

## Methods

### NewAiProfilesListProviderModelsRequest

`func NewAiProfilesListProviderModelsRequest(providerType AiProviderType, baseUrl string, apiKey string, ) *AiProfilesListProviderModelsRequest`

NewAiProfilesListProviderModelsRequest instantiates a new AiProfilesListProviderModelsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiProfilesListProviderModelsRequestWithDefaults

`func NewAiProfilesListProviderModelsRequestWithDefaults() *AiProfilesListProviderModelsRequest`

NewAiProfilesListProviderModelsRequestWithDefaults instantiates a new AiProfilesListProviderModelsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderType

`func (o *AiProfilesListProviderModelsRequest) GetProviderType() AiProviderType`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *AiProfilesListProviderModelsRequest) GetProviderTypeOk() (*AiProviderType, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *AiProfilesListProviderModelsRequest) SetProviderType(v AiProviderType)`

SetProviderType sets ProviderType field to given value.


### GetBaseUrl

`func (o *AiProfilesListProviderModelsRequest) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *AiProfilesListProviderModelsRequest) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *AiProfilesListProviderModelsRequest) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.


### GetApiKey

`func (o *AiProfilesListProviderModelsRequest) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *AiProfilesListProviderModelsRequest) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *AiProfilesListProviderModelsRequest) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


