# AiTProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**AiProviderType**](AiProviderType.md) | Provider type identifier. | 
**Name** | **string** | User-defined display name for this provider connection. | 
**Key** | Pointer to **string** | API key or token. Optional for local providers (Ollama, LM Studio). | [optional] 
**BaseUrl** | **string** | Base URL of the provider API. | 

## Methods

### NewAiTProvider

`func NewAiTProvider(type_ AiProviderType, name string, baseUrl string, ) *AiTProvider`

NewAiTProvider instantiates a new AiTProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiTProviderWithDefaults

`func NewAiTProviderWithDefaults() *AiTProvider`

NewAiTProviderWithDefaults instantiates a new AiTProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AiTProvider) GetType() AiProviderType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AiTProvider) GetTypeOk() (*AiProviderType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AiTProvider) SetType(v AiProviderType)`

SetType sets Type field to given value.


### GetName

`func (o *AiTProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AiTProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AiTProvider) SetName(v string)`

SetName sets Name field to given value.


### GetKey

`func (o *AiTProvider) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AiTProvider) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AiTProvider) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *AiTProvider) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetBaseUrl

`func (o *AiTProvider) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *AiTProvider) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *AiTProvider) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


