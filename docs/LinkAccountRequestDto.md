# LinkAccountRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SerializedProfile** | Pointer to **NullableString** | The profile a completed provider authorization produced, in the serialized form the login flow hands back.  Pass that value unchanged; it carries the provider, the third-party account ID and the authorization result,  and a hand-written object is not accepted. | [optional] 

## Methods

### NewLinkAccountRequestDto

`func NewLinkAccountRequestDto() *LinkAccountRequestDto`

NewLinkAccountRequestDto instantiates a new LinkAccountRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkAccountRequestDtoWithDefaults

`func NewLinkAccountRequestDtoWithDefaults() *LinkAccountRequestDto`

NewLinkAccountRequestDtoWithDefaults instantiates a new LinkAccountRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerializedProfile

`func (o *LinkAccountRequestDto) GetSerializedProfile() string`

GetSerializedProfile returns the SerializedProfile field if non-nil, zero value otherwise.

### GetSerializedProfileOk

`func (o *LinkAccountRequestDto) GetSerializedProfileOk() (*string, bool)`

GetSerializedProfileOk returns a tuple with the SerializedProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializedProfile

`func (o *LinkAccountRequestDto) SetSerializedProfile(v string)`

SetSerializedProfile sets SerializedProfile field to given value.

### HasSerializedProfile

`func (o *LinkAccountRequestDto) HasSerializedProfile() bool`

HasSerializedProfile returns a boolean if a field has been set.

### SetSerializedProfileNil

`func (o *LinkAccountRequestDto) SetSerializedProfileNil(b bool)`

 SetSerializedProfileNil sets the value for SerializedProfile to be an explicit nil

### UnsetSerializedProfile
`func (o *LinkAccountRequestDto) UnsetSerializedProfile()`

UnsetSerializedProfile ensures that no value is present for SerializedProfile, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


