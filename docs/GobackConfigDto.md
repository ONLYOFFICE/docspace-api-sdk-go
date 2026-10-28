# GobackConfigDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **NullableString** | Where the user is taken when they leave the document, normally the folder or the room it lies in. It is empty  when there is nowhere to return to, as in a framed opening. | [optional] 

## Methods

### NewGobackConfigDto

`func NewGobackConfigDto() *GobackConfigDto`

NewGobackConfigDto instantiates a new GobackConfigDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGobackConfigDtoWithDefaults

`func NewGobackConfigDtoWithDefaults() *GobackConfigDto`

NewGobackConfigDtoWithDefaults instantiates a new GobackConfigDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *GobackConfigDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GobackConfigDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GobackConfigDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GobackConfigDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *GobackConfigDto) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *GobackConfigDto) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


