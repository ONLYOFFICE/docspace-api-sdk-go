# EditorToolCallParametersDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **NullableString** | What the generated fillable form should ask for, in the words the request was made in. | 
**Topic** | Pointer to **NullableString** | What the generated presentation is about. | [optional] 
**SlideCount** | Pointer to **NullableString** | How many slides to generate, as the request spelled it. | [optional] 
**Style** | Pointer to **NullableString** | The visual style the slides should be generated in. | [optional] 

## Methods

### NewEditorToolCallParametersDto

`func NewEditorToolCallParametersDto(description NullableString, ) *EditorToolCallParametersDto`

NewEditorToolCallParametersDto instantiates a new EditorToolCallParametersDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditorToolCallParametersDtoWithDefaults

`func NewEditorToolCallParametersDtoWithDefaults() *EditorToolCallParametersDto`

NewEditorToolCallParametersDtoWithDefaults instantiates a new EditorToolCallParametersDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *EditorToolCallParametersDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EditorToolCallParametersDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EditorToolCallParametersDto) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *EditorToolCallParametersDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *EditorToolCallParametersDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTopic

`func (o *EditorToolCallParametersDto) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *EditorToolCallParametersDto) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *EditorToolCallParametersDto) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *EditorToolCallParametersDto) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### SetTopicNil

`func (o *EditorToolCallParametersDto) SetTopicNil(b bool)`

 SetTopicNil sets the value for Topic to be an explicit nil

### UnsetTopic
`func (o *EditorToolCallParametersDto) UnsetTopic()`

UnsetTopic ensures that no value is present for Topic, not even an explicit nil
### GetSlideCount

`func (o *EditorToolCallParametersDto) GetSlideCount() string`

GetSlideCount returns the SlideCount field if non-nil, zero value otherwise.

### GetSlideCountOk

`func (o *EditorToolCallParametersDto) GetSlideCountOk() (*string, bool)`

GetSlideCountOk returns a tuple with the SlideCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlideCount

`func (o *EditorToolCallParametersDto) SetSlideCount(v string)`

SetSlideCount sets SlideCount field to given value.

### HasSlideCount

`func (o *EditorToolCallParametersDto) HasSlideCount() bool`

HasSlideCount returns a boolean if a field has been set.

### SetSlideCountNil

`func (o *EditorToolCallParametersDto) SetSlideCountNil(b bool)`

 SetSlideCountNil sets the value for SlideCount to be an explicit nil

### UnsetSlideCount
`func (o *EditorToolCallParametersDto) UnsetSlideCount()`

UnsetSlideCount ensures that no value is present for SlideCount, not even an explicit nil
### GetStyle

`func (o *EditorToolCallParametersDto) GetStyle() string`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *EditorToolCallParametersDto) GetStyleOk() (*string, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *EditorToolCallParametersDto) SetStyle(v string)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *EditorToolCallParametersDto) HasStyle() bool`

HasStyle returns a boolean if a field has been set.

### SetStyleNil

`func (o *EditorToolCallParametersDto) SetStyleNil(b bool)`

 SetStyleNil sets the value for Style to be an explicit nil

### UnsetStyle
`func (o *EditorToolCallParametersDto) UnsetStyle()`

UnsetStyle ensures that no value is present for Style, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


