# AiProjectModelProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The profile id. | 
**Name** | **string** | The profile name; empty when the profile is gone. | 
**Provider** | **NullableString** | The provider type of the profile; `null` when the profile is gone. | 
**Available** | **bool** | `false` when the profile is gone or cannot chat: the project's rounds are refused. | 

## Methods

### NewAiProjectModelProfile

`func NewAiProjectModelProfile(id string, name string, provider NullableString, available bool, ) *AiProjectModelProfile`

NewAiProjectModelProfile instantiates a new AiProjectModelProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiProjectModelProfileWithDefaults

`func NewAiProjectModelProfileWithDefaults() *AiProjectModelProfile`

NewAiProjectModelProfileWithDefaults instantiates a new AiProjectModelProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AiProjectModelProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiProjectModelProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiProjectModelProfile) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AiProjectModelProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AiProjectModelProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AiProjectModelProfile) SetName(v string)`

SetName sets Name field to given value.


### GetProvider

`func (o *AiProjectModelProfile) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AiProjectModelProfile) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AiProjectModelProfile) SetProvider(v string)`

SetProvider sets Provider field to given value.


### SetProviderNil

`func (o *AiProjectModelProfile) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *AiProjectModelProfile) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetAvailable

`func (o *AiProjectModelProfile) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *AiProjectModelProfile) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *AiProjectModelProfile) SetAvailable(v bool)`

SetAvailable sets Available field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


