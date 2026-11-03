# AiContextRoomRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cloud** | **string** | The account the room belongs to; `docspace` for a room of this portal. | 
**Id** | **string** | The room ID. | 
**Skills** | Pointer to [**[]AiRoomSkill**](AiRoomSkill.md) | The skills of the room the round may use, as `GET api/2.0/ai/context/get-room-skills` lists them. | [optional] 

## Methods

### NewAiContextRoomRef

`func NewAiContextRoomRef(cloud string, id string, ) *AiContextRoomRef`

NewAiContextRoomRef instantiates a new AiContextRoomRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiContextRoomRefWithDefaults

`func NewAiContextRoomRefWithDefaults() *AiContextRoomRef`

NewAiContextRoomRefWithDefaults instantiates a new AiContextRoomRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloud

`func (o *AiContextRoomRef) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *AiContextRoomRef) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *AiContextRoomRef) SetCloud(v string)`

SetCloud sets Cloud field to given value.


### GetId

`func (o *AiContextRoomRef) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiContextRoomRef) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiContextRoomRef) SetId(v string)`

SetId sets Id field to given value.


### GetSkills

`func (o *AiContextRoomRef) GetSkills() []AiRoomSkill`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *AiContextRoomRef) GetSkillsOk() (*[]AiRoomSkill, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkills

`func (o *AiContextRoomRef) SetSkills(v []AiRoomSkill)`

SetSkills sets Skills field to given value.

### HasSkills

`func (o *AiContextRoomRef) HasSkills() bool`

HasSkills returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


