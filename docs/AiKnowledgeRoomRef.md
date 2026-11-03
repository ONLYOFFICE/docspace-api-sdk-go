# AiKnowledgeRoomRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cloud** | **string** | The account the room belongs to; `docspace` for a room of this portal. | 
**Id** | **string** | The room ID; anything that is not a positive integer is ignored. | 

## Methods

### NewAiKnowledgeRoomRef

`func NewAiKnowledgeRoomRef(cloud string, id string, ) *AiKnowledgeRoomRef`

NewAiKnowledgeRoomRef instantiates a new AiKnowledgeRoomRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiKnowledgeRoomRefWithDefaults

`func NewAiKnowledgeRoomRefWithDefaults() *AiKnowledgeRoomRef`

NewAiKnowledgeRoomRefWithDefaults instantiates a new AiKnowledgeRoomRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloud

`func (o *AiKnowledgeRoomRef) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *AiKnowledgeRoomRef) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *AiKnowledgeRoomRef) SetCloud(v string)`

SetCloud sets Cloud field to given value.


### GetId

`func (o *AiKnowledgeRoomRef) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiKnowledgeRoomRef) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiKnowledgeRoomRef) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


