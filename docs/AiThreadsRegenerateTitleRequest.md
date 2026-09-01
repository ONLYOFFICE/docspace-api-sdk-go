# AiThreadsRegenerateTitleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThreadId** | **string** |  | 
**Profile** | [**AiProfile**](AiProfile.md) | Profile used to regenerate the title. | 
**EntityMeta** | Pointer to [**AiThreadsOpenOrCreateRequestEntityMeta**](AiThreadsOpenOrCreateRequestEntityMeta.md) |  | [optional] 

## Methods

### NewAiThreadsRegenerateTitleRequest

`func NewAiThreadsRegenerateTitleRequest(threadId string, profile AiProfile, ) *AiThreadsRegenerateTitleRequest`

NewAiThreadsRegenerateTitleRequest instantiates a new AiThreadsRegenerateTitleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiThreadsRegenerateTitleRequestWithDefaults

`func NewAiThreadsRegenerateTitleRequestWithDefaults() *AiThreadsRegenerateTitleRequest`

NewAiThreadsRegenerateTitleRequestWithDefaults instantiates a new AiThreadsRegenerateTitleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreadId

`func (o *AiThreadsRegenerateTitleRequest) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *AiThreadsRegenerateTitleRequest) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *AiThreadsRegenerateTitleRequest) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.


### GetProfile

`func (o *AiThreadsRegenerateTitleRequest) GetProfile() AiProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *AiThreadsRegenerateTitleRequest) GetProfileOk() (*AiProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *AiThreadsRegenerateTitleRequest) SetProfile(v AiProfile)`

SetProfile sets Profile field to given value.


### GetEntityMeta

`func (o *AiThreadsRegenerateTitleRequest) GetEntityMeta() AiThreadsOpenOrCreateRequestEntityMeta`

GetEntityMeta returns the EntityMeta field if non-nil, zero value otherwise.

### GetEntityMetaOk

`func (o *AiThreadsRegenerateTitleRequest) GetEntityMetaOk() (*AiThreadsOpenOrCreateRequestEntityMeta, bool)`

GetEntityMetaOk returns a tuple with the EntityMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityMeta

`func (o *AiThreadsRegenerateTitleRequest) SetEntityMeta(v AiThreadsOpenOrCreateRequestEntityMeta)`

SetEntityMeta sets EntityMeta field to given value.

### HasEntityMeta

`func (o *AiThreadsRegenerateTitleRequest) HasEntityMeta() bool`

HasEntityMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


