# GeneratePresentationToolCallParametersDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Topic** | Pointer to **NullableString** | What the generated presentation is about. | [optional] 
**SlideCount** | Pointer to **NullableString** | How many slides to generate, as the request spelled it. | [optional] 
**Style** | Pointer to **NullableString** | The visual style the slides should be generated in. | [optional] 

## Methods

### NewGeneratePresentationToolCallParametersDto

`func NewGeneratePresentationToolCallParametersDto() *GeneratePresentationToolCallParametersDto`

NewGeneratePresentationToolCallParametersDto instantiates a new GeneratePresentationToolCallParametersDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeneratePresentationToolCallParametersDtoWithDefaults

`func NewGeneratePresentationToolCallParametersDtoWithDefaults() *GeneratePresentationToolCallParametersDto`

NewGeneratePresentationToolCallParametersDtoWithDefaults instantiates a new GeneratePresentationToolCallParametersDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopic

`func (o *GeneratePresentationToolCallParametersDto) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *GeneratePresentationToolCallParametersDto) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *GeneratePresentationToolCallParametersDto) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *GeneratePresentationToolCallParametersDto) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### SetTopicNil

`func (o *GeneratePresentationToolCallParametersDto) SetTopicNil(b bool)`

 SetTopicNil sets the value for Topic to be an explicit nil

### UnsetTopic
`func (o *GeneratePresentationToolCallParametersDto) UnsetTopic()`

UnsetTopic ensures that no value is present for Topic, not even an explicit nil
### GetSlideCount

`func (o *GeneratePresentationToolCallParametersDto) GetSlideCount() string`

GetSlideCount returns the SlideCount field if non-nil, zero value otherwise.

### GetSlideCountOk

`func (o *GeneratePresentationToolCallParametersDto) GetSlideCountOk() (*string, bool)`

GetSlideCountOk returns a tuple with the SlideCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlideCount

`func (o *GeneratePresentationToolCallParametersDto) SetSlideCount(v string)`

SetSlideCount sets SlideCount field to given value.

### HasSlideCount

`func (o *GeneratePresentationToolCallParametersDto) HasSlideCount() bool`

HasSlideCount returns a boolean if a field has been set.

### SetSlideCountNil

`func (o *GeneratePresentationToolCallParametersDto) SetSlideCountNil(b bool)`

 SetSlideCountNil sets the value for SlideCount to be an explicit nil

### UnsetSlideCount
`func (o *GeneratePresentationToolCallParametersDto) UnsetSlideCount()`

UnsetSlideCount ensures that no value is present for SlideCount, not even an explicit nil
### GetStyle

`func (o *GeneratePresentationToolCallParametersDto) GetStyle() string`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *GeneratePresentationToolCallParametersDto) GetStyleOk() (*string, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *GeneratePresentationToolCallParametersDto) SetStyle(v string)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *GeneratePresentationToolCallParametersDto) HasStyle() bool`

HasStyle returns a boolean if a field has been set.

### SetStyleNil

`func (o *GeneratePresentationToolCallParametersDto) SetStyleNil(b bool)`

 SetStyleNil sets the value for Style to be an explicit nil

### UnsetStyle
`func (o *GeneratePresentationToolCallParametersDto) UnsetStyle()`

UnsetStyle ensures that no value is present for Style, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


