# ReviewConfigDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReviewDisplay** | Pointer to **NullableString** | How the editors render tracked changes at first: with the markup, in a simplified markup, as the final text,  or as the original text. A session that may not write opens on the final text. | [optional] [readonly] 

## Methods

### NewReviewConfigDto

`func NewReviewConfigDto() *ReviewConfigDto`

NewReviewConfigDto instantiates a new ReviewConfigDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewConfigDtoWithDefaults

`func NewReviewConfigDtoWithDefaults() *ReviewConfigDto`

NewReviewConfigDtoWithDefaults instantiates a new ReviewConfigDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReviewDisplay

`func (o *ReviewConfigDto) GetReviewDisplay() string`

GetReviewDisplay returns the ReviewDisplay field if non-nil, zero value otherwise.

### GetReviewDisplayOk

`func (o *ReviewConfigDto) GetReviewDisplayOk() (*string, bool)`

GetReviewDisplayOk returns a tuple with the ReviewDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewDisplay

`func (o *ReviewConfigDto) SetReviewDisplay(v string)`

SetReviewDisplay sets ReviewDisplay field to given value.

### HasReviewDisplay

`func (o *ReviewConfigDto) HasReviewDisplay() bool`

HasReviewDisplay returns a boolean if a field has been set.

### SetReviewDisplayNil

`func (o *ReviewConfigDto) SetReviewDisplayNil(b bool)`

 SetReviewDisplayNil sets the value for ReviewDisplay to be an explicit nil

### UnsetReviewDisplay
`func (o *ReviewConfigDto) UnsetReviewDisplay()`

UnsetReviewDisplay ensures that no value is present for ReviewDisplay, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


