# InvitationLinkDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the invitation link. | [optional] 
**EmployeeType** | [**EmployeeType**](EmployeeType.md) | The type of employee role for the invitation link. | 
**Expiration** | Pointer to **NullableTime** | The expiration date of the invitation link. | [optional] 
**IsExpired** | Pointer to **bool** | Indicates whether the invitation link has expired. | [optional] 
**MaxUseCount** | Pointer to **NullableInt32** | The maximum number of times the invitation link can be used. | [optional] 
**CurrentUseCount** | Pointer to **int32** | The current number of times the invitation link has been used. | [optional] 
**Url** | Pointer to **NullableString** | The URL of the invitation link. | [optional] 

## Methods

### NewInvitationLinkDto

`func NewInvitationLinkDto(employeeType EmployeeType, ) *InvitationLinkDto`

NewInvitationLinkDto instantiates a new InvitationLinkDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvitationLinkDtoWithDefaults

`func NewInvitationLinkDtoWithDefaults() *InvitationLinkDto`

NewInvitationLinkDtoWithDefaults instantiates a new InvitationLinkDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvitationLinkDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvitationLinkDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvitationLinkDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InvitationLinkDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEmployeeType

`func (o *InvitationLinkDto) GetEmployeeType() EmployeeType`

GetEmployeeType returns the EmployeeType field if non-nil, zero value otherwise.

### GetEmployeeTypeOk

`func (o *InvitationLinkDto) GetEmployeeTypeOk() (*EmployeeType, bool)`

GetEmployeeTypeOk returns a tuple with the EmployeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeType

`func (o *InvitationLinkDto) SetEmployeeType(v EmployeeType)`

SetEmployeeType sets EmployeeType field to given value.


### GetExpiration

`func (o *InvitationLinkDto) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *InvitationLinkDto) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *InvitationLinkDto) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *InvitationLinkDto) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### SetExpirationNil

`func (o *InvitationLinkDto) SetExpirationNil(b bool)`

 SetExpirationNil sets the value for Expiration to be an explicit nil

### UnsetExpiration
`func (o *InvitationLinkDto) UnsetExpiration()`

UnsetExpiration ensures that no value is present for Expiration, not even an explicit nil
### GetIsExpired

`func (o *InvitationLinkDto) GetIsExpired() bool`

GetIsExpired returns the IsExpired field if non-nil, zero value otherwise.

### GetIsExpiredOk

`func (o *InvitationLinkDto) GetIsExpiredOk() (*bool, bool)`

GetIsExpiredOk returns a tuple with the IsExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExpired

`func (o *InvitationLinkDto) SetIsExpired(v bool)`

SetIsExpired sets IsExpired field to given value.

### HasIsExpired

`func (o *InvitationLinkDto) HasIsExpired() bool`

HasIsExpired returns a boolean if a field has been set.

### GetMaxUseCount

`func (o *InvitationLinkDto) GetMaxUseCount() int32`

GetMaxUseCount returns the MaxUseCount field if non-nil, zero value otherwise.

### GetMaxUseCountOk

`func (o *InvitationLinkDto) GetMaxUseCountOk() (*int32, bool)`

GetMaxUseCountOk returns a tuple with the MaxUseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUseCount

`func (o *InvitationLinkDto) SetMaxUseCount(v int32)`

SetMaxUseCount sets MaxUseCount field to given value.

### HasMaxUseCount

`func (o *InvitationLinkDto) HasMaxUseCount() bool`

HasMaxUseCount returns a boolean if a field has been set.

### SetMaxUseCountNil

`func (o *InvitationLinkDto) SetMaxUseCountNil(b bool)`

 SetMaxUseCountNil sets the value for MaxUseCount to be an explicit nil

### UnsetMaxUseCount
`func (o *InvitationLinkDto) UnsetMaxUseCount()`

UnsetMaxUseCount ensures that no value is present for MaxUseCount, not even an explicit nil
### GetCurrentUseCount

`func (o *InvitationLinkDto) GetCurrentUseCount() int32`

GetCurrentUseCount returns the CurrentUseCount field if non-nil, zero value otherwise.

### GetCurrentUseCountOk

`func (o *InvitationLinkDto) GetCurrentUseCountOk() (*int32, bool)`

GetCurrentUseCountOk returns a tuple with the CurrentUseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUseCount

`func (o *InvitationLinkDto) SetCurrentUseCount(v int32)`

SetCurrentUseCount sets CurrentUseCount field to given value.

### HasCurrentUseCount

`func (o *InvitationLinkDto) HasCurrentUseCount() bool`

HasCurrentUseCount returns a boolean if a field has been set.

### GetUrl

`func (o *InvitationLinkDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InvitationLinkDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InvitationLinkDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InvitationLinkDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *InvitationLinkDto) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *InvitationLinkDto) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


