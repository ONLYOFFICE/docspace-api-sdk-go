# AiRoomSkill

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The skill's ID, as `GET api/2.0/ai/context/get-room-skill` takes it. | 
**Name** | **string** | The skill's name, from its frontmatter. | 
**Description** | **string** | What the skill is for, from its frontmatter; empty when it declares none. | 

## Methods

### NewAiRoomSkill

`func NewAiRoomSkill(id string, name string, description string, ) *AiRoomSkill`

NewAiRoomSkill instantiates a new AiRoomSkill object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiRoomSkillWithDefaults

`func NewAiRoomSkillWithDefaults() *AiRoomSkill`

NewAiRoomSkillWithDefaults instantiates a new AiRoomSkill object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AiRoomSkill) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiRoomSkill) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiRoomSkill) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AiRoomSkill) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AiRoomSkill) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AiRoomSkill) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AiRoomSkill) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AiRoomSkill) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AiRoomSkill) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


