# AiSkillDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** | ID of the skill file. | 
**Name** | **NullableString** | Skill name from the frontmatter. | 
**Description** | **NullableString** | Skill description from the frontmatter. | 

## Methods

### NewAiSkillDto

`func NewAiSkillDto(id NullableString, name NullableString, description NullableString, ) *AiSkillDto`

NewAiSkillDto instantiates a new AiSkillDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiSkillDtoWithDefaults

`func NewAiSkillDtoWithDefaults() *AiSkillDto`

NewAiSkillDtoWithDefaults instantiates a new AiSkillDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AiSkillDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiSkillDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiSkillDto) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *AiSkillDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AiSkillDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *AiSkillDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AiSkillDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AiSkillDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *AiSkillDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AiSkillDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *AiSkillDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AiSkillDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AiSkillDto) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *AiSkillDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AiSkillDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


