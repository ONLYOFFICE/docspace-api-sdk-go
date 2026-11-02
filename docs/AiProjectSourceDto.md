# AiProjectSourceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID of the room. | 
**Title** | **NullableString** | Room title. | 
**Logo** | Pointer to [**LogoDto**](LogoDto.md) | Room logo. | [optional] 

## Methods

### NewAiProjectSourceDto

`func NewAiProjectSourceDto(id int32, title NullableString, ) *AiProjectSourceDto`

NewAiProjectSourceDto instantiates a new AiProjectSourceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiProjectSourceDtoWithDefaults

`func NewAiProjectSourceDtoWithDefaults() *AiProjectSourceDto`

NewAiProjectSourceDtoWithDefaults instantiates a new AiProjectSourceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AiProjectSourceDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiProjectSourceDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiProjectSourceDto) SetId(v int32)`

SetId sets Id field to given value.


### GetTitle

`func (o *AiProjectSourceDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AiProjectSourceDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AiProjectSourceDto) SetTitle(v string)`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *AiProjectSourceDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *AiProjectSourceDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetLogo

`func (o *AiProjectSourceDto) GetLogo() LogoDto`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *AiProjectSourceDto) GetLogoOk() (*LogoDto, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *AiProjectSourceDto) SetLogo(v LogoDto)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *AiProjectSourceDto) HasLogo() bool`

HasLogo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


