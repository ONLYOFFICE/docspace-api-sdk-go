# AiProjectSourcesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstructionsRooms** | [**[]AiProjectSourceDto**](AiProjectSourceDto.md) | Linked Instructions rooms, sorted by title. | 
**KnowledgeRooms** | [**[]AiProjectSourceDto**](AiProjectSourceDto.md) | Linked Knowledge rooms, sorted by title. | 

## Methods

### NewAiProjectSourcesDto

`func NewAiProjectSourcesDto(instructionsRooms []AiProjectSourceDto, knowledgeRooms []AiProjectSourceDto, ) *AiProjectSourcesDto`

NewAiProjectSourcesDto instantiates a new AiProjectSourcesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiProjectSourcesDtoWithDefaults

`func NewAiProjectSourcesDtoWithDefaults() *AiProjectSourcesDto`

NewAiProjectSourcesDtoWithDefaults instantiates a new AiProjectSourcesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstructionsRooms

`func (o *AiProjectSourcesDto) GetInstructionsRooms() []AiProjectSourceDto`

GetInstructionsRooms returns the InstructionsRooms field if non-nil, zero value otherwise.

### GetInstructionsRoomsOk

`func (o *AiProjectSourcesDto) GetInstructionsRoomsOk() (*[]AiProjectSourceDto, bool)`

GetInstructionsRoomsOk returns a tuple with the InstructionsRooms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructionsRooms

`func (o *AiProjectSourcesDto) SetInstructionsRooms(v []AiProjectSourceDto)`

SetInstructionsRooms sets InstructionsRooms field to given value.


### SetInstructionsRoomsNil

`func (o *AiProjectSourcesDto) SetInstructionsRoomsNil(b bool)`

 SetInstructionsRoomsNil sets the value for InstructionsRooms to be an explicit nil

### UnsetInstructionsRooms
`func (o *AiProjectSourcesDto) UnsetInstructionsRooms()`

UnsetInstructionsRooms ensures that no value is present for InstructionsRooms, not even an explicit nil
### GetKnowledgeRooms

`func (o *AiProjectSourcesDto) GetKnowledgeRooms() []AiProjectSourceDto`

GetKnowledgeRooms returns the KnowledgeRooms field if non-nil, zero value otherwise.

### GetKnowledgeRoomsOk

`func (o *AiProjectSourcesDto) GetKnowledgeRoomsOk() (*[]AiProjectSourceDto, bool)`

GetKnowledgeRoomsOk returns a tuple with the KnowledgeRooms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnowledgeRooms

`func (o *AiProjectSourcesDto) SetKnowledgeRooms(v []AiProjectSourceDto)`

SetKnowledgeRooms sets KnowledgeRooms field to given value.


### SetKnowledgeRoomsNil

`func (o *AiProjectSourcesDto) SetKnowledgeRoomsNil(b bool)`

 SetKnowledgeRoomsNil sets the value for KnowledgeRooms to be an explicit nil

### UnsetKnowledgeRooms
`func (o *AiProjectSourcesDto) UnsetKnowledgeRooms()`

UnsetKnowledgeRooms ensures that no value is present for KnowledgeRooms, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


