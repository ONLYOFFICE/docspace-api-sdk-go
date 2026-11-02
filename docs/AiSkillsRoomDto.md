# AiSkillsRoomDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID of the room. | 
**Title** | **NullableString** | Room title. | 
**Logo** | [**LogoDto**](LogoDto.md) | Room logo. | 
**Skills** | [**[]AiSkillDto**](AiSkillDto.md) | Skills of the room, sorted by file title. | 

## Methods

### NewAiSkillsRoomDto

`func NewAiSkillsRoomDto(id int32, title NullableString, logo LogoDto, skills []AiSkillDto, ) *AiSkillsRoomDto`

NewAiSkillsRoomDto instantiates a new AiSkillsRoomDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiSkillsRoomDtoWithDefaults

`func NewAiSkillsRoomDtoWithDefaults() *AiSkillsRoomDto`

NewAiSkillsRoomDtoWithDefaults instantiates a new AiSkillsRoomDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AiSkillsRoomDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiSkillsRoomDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiSkillsRoomDto) SetId(v int32)`

SetId sets Id field to given value.


### GetTitle

`func (o *AiSkillsRoomDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AiSkillsRoomDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AiSkillsRoomDto) SetTitle(v string)`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *AiSkillsRoomDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *AiSkillsRoomDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetLogo

`func (o *AiSkillsRoomDto) GetLogo() LogoDto`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *AiSkillsRoomDto) GetLogoOk() (*LogoDto, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *AiSkillsRoomDto) SetLogo(v LogoDto)`

SetLogo sets Logo field to given value.


### GetSkills

`func (o *AiSkillsRoomDto) GetSkills() []AiSkillDto`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *AiSkillsRoomDto) GetSkillsOk() (*[]AiSkillDto, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkills

`func (o *AiSkillsRoomDto) SetSkills(v []AiSkillDto)`

SetSkills sets Skills field to given value.


### SetSkillsNil

`func (o *AiSkillsRoomDto) SetSkillsNil(b bool)`

 SetSkillsNil sets the value for Skills to be an explicit nil

### UnsetSkills
`func (o *AiSkillsRoomDto) UnsetSkills()`

UnsetSkills ensures that no value is present for Skills, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


