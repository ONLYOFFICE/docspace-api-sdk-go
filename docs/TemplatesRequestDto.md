# TemplatesRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileIds** | Pointer to **[]int32** | The list of file IDs. | [optional] 

## Methods

### NewTemplatesRequestDto

`func NewTemplatesRequestDto() *TemplatesRequestDto`

NewTemplatesRequestDto instantiates a new TemplatesRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplatesRequestDtoWithDefaults

`func NewTemplatesRequestDtoWithDefaults() *TemplatesRequestDto`

NewTemplatesRequestDtoWithDefaults instantiates a new TemplatesRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileIds

`func (o *TemplatesRequestDto) GetFileIds() []int32`

GetFileIds returns the FileIds field if non-nil, zero value otherwise.

### GetFileIdsOk

`func (o *TemplatesRequestDto) GetFileIdsOk() (*[]int32, bool)`

GetFileIdsOk returns a tuple with the FileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileIds

`func (o *TemplatesRequestDto) SetFileIds(v []int32)`

SetFileIds sets FileIds field to given value.

### HasFileIds

`func (o *TemplatesRequestDto) HasFileIds() bool`

HasFileIds returns a boolean if a field has been set.

### SetFileIdsNil

`func (o *TemplatesRequestDto) SetFileIdsNil(b bool)`

 SetFileIdsNil sets the value for FileIds to be an explicit nil

### UnsetFileIds
`func (o *TemplatesRequestDto) UnsetFileIds()`

UnsetFileIds ensures that no value is present for FileIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


