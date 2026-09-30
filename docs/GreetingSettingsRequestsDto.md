# GreetingSettingsRequestsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **NullableString** | The caption to store, which is kept as the portal name. An empty value clears the greeting and returns the  portal to the built-in default caption. On a cloud portal with a free or trial plan the text is also matched  against the character rule configured for the installation and a text that breaks it is refused, while a paid  cloud plan and a self-hosted installation apply no such check. | 

## Methods

### NewGreetingSettingsRequestsDto

`func NewGreetingSettingsRequestsDto(title NullableString, ) *GreetingSettingsRequestsDto`

NewGreetingSettingsRequestsDto instantiates a new GreetingSettingsRequestsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGreetingSettingsRequestsDtoWithDefaults

`func NewGreetingSettingsRequestsDtoWithDefaults() *GreetingSettingsRequestsDto`

NewGreetingSettingsRequestsDtoWithDefaults instantiates a new GreetingSettingsRequestsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *GreetingSettingsRequestsDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GreetingSettingsRequestsDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GreetingSettingsRequestsDto) SetTitle(v string)`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *GreetingSettingsRequestsDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *GreetingSettingsRequestsDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


