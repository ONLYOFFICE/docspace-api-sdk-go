# AiSkillsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rooms** | [**[]AiSkillsRoomDto**](AiSkillsRoomDto.md) | The rooms, sorted by title. | 

## Methods

### NewAiSkillsDto

`func NewAiSkillsDto(rooms []AiSkillsRoomDto, ) *AiSkillsDto`

NewAiSkillsDto instantiates a new AiSkillsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiSkillsDtoWithDefaults

`func NewAiSkillsDtoWithDefaults() *AiSkillsDto`

NewAiSkillsDtoWithDefaults instantiates a new AiSkillsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRooms

`func (o *AiSkillsDto) GetRooms() []AiSkillsRoomDto`

GetRooms returns the Rooms field if non-nil, zero value otherwise.

### GetRoomsOk

`func (o *AiSkillsDto) GetRoomsOk() (*[]AiSkillsRoomDto, bool)`

GetRoomsOk returns a tuple with the Rooms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRooms

`func (o *AiSkillsDto) SetRooms(v []AiSkillsRoomDto)`

SetRooms sets Rooms field to given value.


### SetRoomsNil

`func (o *AiSkillsDto) SetRoomsNil(b bool)`

 SetRoomsNil sets the value for Rooms to be an explicit nil

### UnsetRooms
`func (o *AiSkillsDto) UnsetRooms()`

UnsetRooms ensures that no value is present for Rooms, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


