# SettingsRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Set** | Pointer to **bool** | The state to store for the setting the operation addresses: true switches it on, false switches it off. The  flag carries no meaning of its own - what is switched, who is allowed to switch it, and whether the value  belongs to the calling account or to the whole portal are stated by the operation that binds this body. The  answer repeats the value the portal read back afterwards, which is not always the one that was sent. | [optional] 

## Methods

### NewSettingsRequestDto

`func NewSettingsRequestDto() *SettingsRequestDto`

NewSettingsRequestDto instantiates a new SettingsRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsRequestDtoWithDefaults

`func NewSettingsRequestDtoWithDefaults() *SettingsRequestDto`

NewSettingsRequestDtoWithDefaults instantiates a new SettingsRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSet

`func (o *SettingsRequestDto) GetSet() bool`

GetSet returns the Set field if non-nil, zero value otherwise.

### GetSetOk

`func (o *SettingsRequestDto) GetSetOk() (*bool, bool)`

GetSetOk returns a tuple with the Set field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSet

`func (o *SettingsRequestDto) SetSet(v bool)`

SetSet sets Set field to given value.

### HasSet

`func (o *SettingsRequestDto) HasSet() bool`

HasSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


