# CoverRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | Pointer to **NullableString** | The background colour the room is drawn with while it has no logo, as six hexadecimal digits with no leading  number sign. An empty value restores the default colour of the room type. | [optional] 
**Cover** | Pointer to **NullableString** | The picture drawn on the room while it has no logo, named by an identifier from  `GET api/2.0/files/rooms/covers`. Any other value is rejected, and an empty value leaves the room without a  cover. | [optional] 

## Methods

### NewCoverRequestDto

`func NewCoverRequestDto() *CoverRequestDto`

NewCoverRequestDto instantiates a new CoverRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoverRequestDtoWithDefaults

`func NewCoverRequestDtoWithDefaults() *CoverRequestDto`

NewCoverRequestDtoWithDefaults instantiates a new CoverRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColor

`func (o *CoverRequestDto) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *CoverRequestDto) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *CoverRequestDto) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *CoverRequestDto) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *CoverRequestDto) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *CoverRequestDto) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetCover

`func (o *CoverRequestDto) GetCover() string`

GetCover returns the Cover field if non-nil, zero value otherwise.

### GetCoverOk

`func (o *CoverRequestDto) GetCoverOk() (*string, bool)`

GetCoverOk returns a tuple with the Cover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCover

`func (o *CoverRequestDto) SetCover(v string)`

SetCover sets Cover field to given value.

### HasCover

`func (o *CoverRequestDto) HasCover() bool`

HasCover returns a boolean if a field has been set.

### SetCoverNil

`func (o *CoverRequestDto) SetCoverNil(b bool)`

 SetCoverNil sets the value for Cover to be an explicit nil

### UnsetCover
`func (o *CoverRequestDto) UnsetCover()`

UnsetCover ensures that no value is present for Cover, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


