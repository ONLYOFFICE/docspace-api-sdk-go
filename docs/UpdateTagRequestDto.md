# UpdateTagRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OldName** | **NullableString** | The name of the tag to rename, matched against the catalog exactly as it is stored rather than searched for.  Read the stored spelling from `GET api/2.0/files/tags`. | 
**NewName** | **NullableString** | The name to store instead. It has to be free: names are unique across the portal, so a name another tag  already carries is refused, and merging two tags this way is not possible. | 

## Methods

### NewUpdateTagRequestDto

`func NewUpdateTagRequestDto(oldName NullableString, newName NullableString, ) *UpdateTagRequestDto`

NewUpdateTagRequestDto instantiates a new UpdateTagRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTagRequestDtoWithDefaults

`func NewUpdateTagRequestDtoWithDefaults() *UpdateTagRequestDto`

NewUpdateTagRequestDtoWithDefaults instantiates a new UpdateTagRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOldName

`func (o *UpdateTagRequestDto) GetOldName() string`

GetOldName returns the OldName field if non-nil, zero value otherwise.

### GetOldNameOk

`func (o *UpdateTagRequestDto) GetOldNameOk() (*string, bool)`

GetOldNameOk returns a tuple with the OldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldName

`func (o *UpdateTagRequestDto) SetOldName(v string)`

SetOldName sets OldName field to given value.


### SetOldNameNil

`func (o *UpdateTagRequestDto) SetOldNameNil(b bool)`

 SetOldNameNil sets the value for OldName to be an explicit nil

### UnsetOldName
`func (o *UpdateTagRequestDto) UnsetOldName()`

UnsetOldName ensures that no value is present for OldName, not even an explicit nil
### GetNewName

`func (o *UpdateTagRequestDto) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *UpdateTagRequestDto) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *UpdateTagRequestDto) SetNewName(v string)`

SetNewName sets NewName field to given value.


### SetNewNameNil

`func (o *UpdateTagRequestDto) SetNewNameNil(b bool)`

 SetNewNameNil sets the value for NewName to be an explicit nil

### UnsetNewName
`func (o *UpdateTagRequestDto) UnsetNewName()`

UnsetNewName ensures that no value is present for NewName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


