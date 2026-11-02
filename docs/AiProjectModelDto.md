# AiProjectModelDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | Pointer to **NullableString** | ID of the AI model profile selected for the project, as listed by `GET api/2.0/ai/profiles/list`. Null when no model has been selected for the project, or the selected profile no longer exists. | [optional] 

## Methods

### NewAiProjectModelDto

`func NewAiProjectModelDto() *AiProjectModelDto`

NewAiProjectModelDto instantiates a new AiProjectModelDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiProjectModelDtoWithDefaults

`func NewAiProjectModelDtoWithDefaults() *AiProjectModelDto`

NewAiProjectModelDtoWithDefaults instantiates a new AiProjectModelDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *AiProjectModelDto) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *AiProjectModelDto) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *AiProjectModelDto) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *AiProjectModelDto) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### SetProfileIdNil

`func (o *AiProjectModelDto) SetProfileIdNil(b bool)`

 SetProfileIdNil sets the value for ProfileId to be an explicit nil

### UnsetProfileId
`func (o *AiProjectModelDto) UnsetProfileId()`

UnsetProfileId ensures that no value is present for ProfileId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


