# AiProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**AiProviderType**](AiProviderType.md) | Provider type identifier. | 
**Name** | **string** | User-defined display name for this provider connection. | 
**Key** | Pointer to **string** | API key or token. Optional for local providers (Ollama, LM Studio). | [optional] 
**BaseUrl** | **string** | Base URL of the provider API. | 

## Methods

### NewAiProvider

`func NewAiProvider(type_ AiProviderType, name string, baseUrl string, ) *AiProvider`

NewAiProvider instantiates a new AiProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiProviderWithDefaults

`func NewAiProviderWithDefaults() *AiProvider`

NewAiProviderWithDefaults instantiates a new AiProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AiProvider) GetType() AiProviderType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AiProvider) GetTypeOk() (*AiProviderType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AiProvider) SetType(v AiProviderType)`

SetType sets Type field to given value.


### GetName

`func (o *AiProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AiProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AiProvider) SetName(v string)`

SetName sets Name field to given value.


### GetKey

`func (o *AiProvider) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AiProvider) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AiProvider) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *AiProvider) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetBaseUrl

`func (o *AiProvider) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *AiProvider) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *AiProvider) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


