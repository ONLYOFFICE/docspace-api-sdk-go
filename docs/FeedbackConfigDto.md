# FeedbackConfigDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **NullableString** | The absolute URL to the website address which will be opened when clicking the Feedback & Support menu button. | [optional] 
**Visible** | Pointer to **bool** | Whether the support button is shown. The portal always asks for it to be shown. | [optional] [readonly] 

## Methods

### NewFeedbackConfigDto

`func NewFeedbackConfigDto() *FeedbackConfigDto`

NewFeedbackConfigDto instantiates a new FeedbackConfigDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedbackConfigDtoWithDefaults

`func NewFeedbackConfigDtoWithDefaults() *FeedbackConfigDto`

NewFeedbackConfigDtoWithDefaults instantiates a new FeedbackConfigDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *FeedbackConfigDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FeedbackConfigDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FeedbackConfigDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *FeedbackConfigDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *FeedbackConfigDto) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *FeedbackConfigDto) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetVisible

`func (o *FeedbackConfigDto) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *FeedbackConfigDto) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *FeedbackConfigDto) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *FeedbackConfigDto) HasVisible() bool`

HasVisible returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


