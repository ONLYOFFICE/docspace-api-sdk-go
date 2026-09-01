# InvitationLinkCreateRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmployeeType** | [**EmployeeType**](EmployeeType.md) | The type of employee role for the invitation link (DocSpaceAdmin, RoomAdmin or User). | 
**Expiration** | Pointer to **NullableTime** | The expiration date of the invitation link. | [optional] 
**MaxUseCount** | Pointer to **NullableInt32** | The maximum number of times the invitation link can be used. | [optional] 

## Methods

### NewInvitationLinkCreateRequestDto

`func NewInvitationLinkCreateRequestDto(employeeType EmployeeType, ) *InvitationLinkCreateRequestDto`

NewInvitationLinkCreateRequestDto instantiates a new InvitationLinkCreateRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvitationLinkCreateRequestDtoWithDefaults

`func NewInvitationLinkCreateRequestDtoWithDefaults() *InvitationLinkCreateRequestDto`

NewInvitationLinkCreateRequestDtoWithDefaults instantiates a new InvitationLinkCreateRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmployeeType

`func (o *InvitationLinkCreateRequestDto) GetEmployeeType() EmployeeType`

GetEmployeeType returns the EmployeeType field if non-nil, zero value otherwise.

### GetEmployeeTypeOk

`func (o *InvitationLinkCreateRequestDto) GetEmployeeTypeOk() (*EmployeeType, bool)`

GetEmployeeTypeOk returns a tuple with the EmployeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeType

`func (o *InvitationLinkCreateRequestDto) SetEmployeeType(v EmployeeType)`

SetEmployeeType sets EmployeeType field to given value.


### GetExpiration

`func (o *InvitationLinkCreateRequestDto) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *InvitationLinkCreateRequestDto) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *InvitationLinkCreateRequestDto) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *InvitationLinkCreateRequestDto) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### SetExpirationNil

`func (o *InvitationLinkCreateRequestDto) SetExpirationNil(b bool)`

 SetExpirationNil sets the value for Expiration to be an explicit nil

### UnsetExpiration
`func (o *InvitationLinkCreateRequestDto) UnsetExpiration()`

UnsetExpiration ensures that no value is present for Expiration, not even an explicit nil
### GetMaxUseCount

`func (o *InvitationLinkCreateRequestDto) GetMaxUseCount() int32`

GetMaxUseCount returns the MaxUseCount field if non-nil, zero value otherwise.

### GetMaxUseCountOk

`func (o *InvitationLinkCreateRequestDto) GetMaxUseCountOk() (*int32, bool)`

GetMaxUseCountOk returns a tuple with the MaxUseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUseCount

`func (o *InvitationLinkCreateRequestDto) SetMaxUseCount(v int32)`

SetMaxUseCount sets MaxUseCount field to given value.

### HasMaxUseCount

`func (o *InvitationLinkCreateRequestDto) HasMaxUseCount() bool`

HasMaxUseCount returns a boolean if a field has been set.

### SetMaxUseCountNil

`func (o *InvitationLinkCreateRequestDto) SetMaxUseCountNil(b bool)`

 SetMaxUseCountNil sets the value for MaxUseCount to be an explicit nil

### UnsetMaxUseCount
`func (o *InvitationLinkCreateRequestDto) UnsetMaxUseCount()`

UnsetMaxUseCount ensures that no value is present for MaxUseCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


