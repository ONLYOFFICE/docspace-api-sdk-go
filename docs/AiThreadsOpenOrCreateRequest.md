# AiThreadsOpenOrCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThreadId** | Pointer to **string** |  | [optional] 
**Profile** | [**AiProfile**](AiProfile.md) | Profile the title generation runs on. | 
**ProfileId** | **string** |  | 
**FirstMessage** | [**AiThreadMessageLike**](AiThreadMessageLike.md) | First user message a fresh thread derives its title from. | 
**EntityId** | Pointer to **string** | Opaque scope token persisted on a freshly created thread. | [optional] 
**EntityMeta** | Pointer to [**AiThreadsOpenOrCreateRequestEntityMeta**](AiThreadsOpenOrCreateRequestEntityMeta.md) |  | [optional] 

## Methods

### NewAiThreadsOpenOrCreateRequest

`func NewAiThreadsOpenOrCreateRequest(profile AiProfile, profileId string, firstMessage AiThreadMessageLike, ) *AiThreadsOpenOrCreateRequest`

NewAiThreadsOpenOrCreateRequest instantiates a new AiThreadsOpenOrCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiThreadsOpenOrCreateRequestWithDefaults

`func NewAiThreadsOpenOrCreateRequestWithDefaults() *AiThreadsOpenOrCreateRequest`

NewAiThreadsOpenOrCreateRequestWithDefaults instantiates a new AiThreadsOpenOrCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreadId

`func (o *AiThreadsOpenOrCreateRequest) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *AiThreadsOpenOrCreateRequest) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *AiThreadsOpenOrCreateRequest) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.

### HasThreadId

`func (o *AiThreadsOpenOrCreateRequest) HasThreadId() bool`

HasThreadId returns a boolean if a field has been set.

### GetProfile

`func (o *AiThreadsOpenOrCreateRequest) GetProfile() AiProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *AiThreadsOpenOrCreateRequest) GetProfileOk() (*AiProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *AiThreadsOpenOrCreateRequest) SetProfile(v AiProfile)`

SetProfile sets Profile field to given value.


### GetProfileId

`func (o *AiThreadsOpenOrCreateRequest) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *AiThreadsOpenOrCreateRequest) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *AiThreadsOpenOrCreateRequest) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetFirstMessage

`func (o *AiThreadsOpenOrCreateRequest) GetFirstMessage() AiThreadMessageLike`

GetFirstMessage returns the FirstMessage field if non-nil, zero value otherwise.

### GetFirstMessageOk

`func (o *AiThreadsOpenOrCreateRequest) GetFirstMessageOk() (*AiThreadMessageLike, bool)`

GetFirstMessageOk returns a tuple with the FirstMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstMessage

`func (o *AiThreadsOpenOrCreateRequest) SetFirstMessage(v AiThreadMessageLike)`

SetFirstMessage sets FirstMessage field to given value.


### GetEntityId

`func (o *AiThreadsOpenOrCreateRequest) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *AiThreadsOpenOrCreateRequest) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *AiThreadsOpenOrCreateRequest) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *AiThreadsOpenOrCreateRequest) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetEntityMeta

`func (o *AiThreadsOpenOrCreateRequest) GetEntityMeta() AiThreadsOpenOrCreateRequestEntityMeta`

GetEntityMeta returns the EntityMeta field if non-nil, zero value otherwise.

### GetEntityMetaOk

`func (o *AiThreadsOpenOrCreateRequest) GetEntityMetaOk() (*AiThreadsOpenOrCreateRequestEntityMeta, bool)`

GetEntityMetaOk returns a tuple with the EntityMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityMeta

`func (o *AiThreadsOpenOrCreateRequest) SetEntityMeta(v AiThreadsOpenOrCreateRequestEntityMeta)`

SetEntityMeta sets EntityMeta field to given value.

### HasEntityMeta

`func (o *AiThreadsOpenOrCreateRequest) HasEntityMeta() bool`

HasEntityMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


