# RemoveProviderRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | **[]int32** | The set of AI provider identifiers to delete. | 

## Methods

### NewRemoveProviderRequestDto

`func NewRemoveProviderRequestDto(ids []int32, ) *RemoveProviderRequestDto`

NewRemoveProviderRequestDto instantiates a new RemoveProviderRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveProviderRequestDtoWithDefaults

`func NewRemoveProviderRequestDtoWithDefaults() *RemoveProviderRequestDto`

NewRemoveProviderRequestDtoWithDefaults instantiates a new RemoveProviderRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *RemoveProviderRequestDto) GetIds() []int32`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *RemoveProviderRequestDto) GetIdsOk() (*[]int32, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *RemoveProviderRequestDto) SetIds(v []int32)`

SetIds sets Ids field to given value.


### SetIdsNil

`func (o *RemoveProviderRequestDto) SetIdsNil(b bool)`

 SetIdsNil sets the value for Ids to be an explicit nil

### UnsetIds
`func (o *RemoveProviderRequestDto) UnsetIds()`

UnsetIds ensures that no value is present for Ids, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


