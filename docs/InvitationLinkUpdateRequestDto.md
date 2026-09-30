# InvitationLinkUpdateRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The link to change, by the `id` that creating or reading it returned. The role behind that id cannot be  changed here. | 
**Expiration** | Pointer to **NullableTime** | The new deadline, read in the portal time zone. The body is applied as a whole, so leaving it out clears the  deadline rather than keeping the current one; a moment in the past is refused. | [optional] 
**MaxUseCount** | Pointer to **NullableInt32** | The new total number of accounts that may join through the link. It may not be lower than the uses already  spent, which the link reports as `currentUseCount`, and leaving it out removes the limit rather than keeping  the current one. | [optional] 

## Methods

### NewInvitationLinkUpdateRequestDto

`func NewInvitationLinkUpdateRequestDto(id string, ) *InvitationLinkUpdateRequestDto`

NewInvitationLinkUpdateRequestDto instantiates a new InvitationLinkUpdateRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvitationLinkUpdateRequestDtoWithDefaults

`func NewInvitationLinkUpdateRequestDtoWithDefaults() *InvitationLinkUpdateRequestDto`

NewInvitationLinkUpdateRequestDtoWithDefaults instantiates a new InvitationLinkUpdateRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvitationLinkUpdateRequestDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvitationLinkUpdateRequestDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvitationLinkUpdateRequestDto) SetId(v string)`

SetId sets Id field to given value.


### GetExpiration

`func (o *InvitationLinkUpdateRequestDto) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *InvitationLinkUpdateRequestDto) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *InvitationLinkUpdateRequestDto) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *InvitationLinkUpdateRequestDto) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### SetExpirationNil

`func (o *InvitationLinkUpdateRequestDto) SetExpirationNil(b bool)`

 SetExpirationNil sets the value for Expiration to be an explicit nil

### UnsetExpiration
`func (o *InvitationLinkUpdateRequestDto) UnsetExpiration()`

UnsetExpiration ensures that no value is present for Expiration, not even an explicit nil
### GetMaxUseCount

`func (o *InvitationLinkUpdateRequestDto) GetMaxUseCount() int32`

GetMaxUseCount returns the MaxUseCount field if non-nil, zero value otherwise.

### GetMaxUseCountOk

`func (o *InvitationLinkUpdateRequestDto) GetMaxUseCountOk() (*int32, bool)`

GetMaxUseCountOk returns a tuple with the MaxUseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUseCount

`func (o *InvitationLinkUpdateRequestDto) SetMaxUseCount(v int32)`

SetMaxUseCount sets MaxUseCount field to given value.

### HasMaxUseCount

`func (o *InvitationLinkUpdateRequestDto) HasMaxUseCount() bool`

HasMaxUseCount returns a boolean if a field has been set.

### SetMaxUseCountNil

`func (o *InvitationLinkUpdateRequestDto) SetMaxUseCountNil(b bool)`

 SetMaxUseCountNil sets the value for MaxUseCount to be an explicit nil

### UnsetMaxUseCount
`func (o *InvitationLinkUpdateRequestDto) UnsetMaxUseCount()`

UnsetMaxUseCount ensures that no value is present for MaxUseCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


