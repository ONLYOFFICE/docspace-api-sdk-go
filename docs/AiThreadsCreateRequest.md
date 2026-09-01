# AiThreadsCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Thread title. | 
**ProfileId** | Pointer to **string** | Optional profile to bind. | [optional] 
**EntityId** | Pointer to **string** | Optional entity (room) scope. | [optional] 

## Methods

### NewAiThreadsCreateRequest

`func NewAiThreadsCreateRequest(title string, ) *AiThreadsCreateRequest`

NewAiThreadsCreateRequest instantiates a new AiThreadsCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiThreadsCreateRequestWithDefaults

`func NewAiThreadsCreateRequestWithDefaults() *AiThreadsCreateRequest`

NewAiThreadsCreateRequestWithDefaults instantiates a new AiThreadsCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *AiThreadsCreateRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AiThreadsCreateRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AiThreadsCreateRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetProfileId

`func (o *AiThreadsCreateRequest) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *AiThreadsCreateRequest) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *AiThreadsCreateRequest) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *AiThreadsCreateRequest) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetEntityId

`func (o *AiThreadsCreateRequest) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *AiThreadsCreateRequest) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *AiThreadsCreateRequest) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *AiThreadsCreateRequest) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


