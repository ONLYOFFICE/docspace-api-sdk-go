# DisplayRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Set** | Pointer to **bool** | The state to store for the setting the operation addresses: true enables it or shows what it governs, false  disables or hides it. What exactly is affected, and whether the value belongs to the calling account or to the  whole portal, are stated by the operation that binds this body. The portal may store a different value than  the one sent when another setting overrides it, so read the answer rather than assuming. | [optional] 

## Methods

### NewDisplayRequestDto

`func NewDisplayRequestDto() *DisplayRequestDto`

NewDisplayRequestDto instantiates a new DisplayRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisplayRequestDtoWithDefaults

`func NewDisplayRequestDtoWithDefaults() *DisplayRequestDto`

NewDisplayRequestDtoWithDefaults instantiates a new DisplayRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSet

`func (o *DisplayRequestDto) GetSet() bool`

GetSet returns the Set field if non-nil, zero value otherwise.

### GetSetOk

`func (o *DisplayRequestDto) GetSetOk() (*bool, bool)`

GetSetOk returns a tuple with the Set field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSet

`func (o *DisplayRequestDto) SetSet(v bool)`

SetSet sets Set field to given value.

### HasSet

`func (o *DisplayRequestDto) HasSet() bool`

HasSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


