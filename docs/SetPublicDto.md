# SetPublicDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The room template ID. | 
**Public** | Pointer to **bool** | Specifies whether the room template is public or not. | [optional] 

## Methods

### NewSetPublicDto

`func NewSetPublicDto(id int32, ) *SetPublicDto`

NewSetPublicDto instantiates a new SetPublicDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetPublicDtoWithDefaults

`func NewSetPublicDtoWithDefaults() *SetPublicDto`

NewSetPublicDtoWithDefaults instantiates a new SetPublicDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SetPublicDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SetPublicDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SetPublicDto) SetId(v int32)`

SetId sets Id field to given value.


### GetPublic

`func (o *SetPublicDto) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *SetPublicDto) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *SetPublicDto) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *SetPublicDto) HasPublic() bool`

HasPublic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


