# InvitationLinkDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The identifier to address the link by in `PUT api/2.0/portal/users/invitationlink` and  `DELETE api/2.0/portal/users/invitationlink`. It survives a change of deadline or use limit, so it is  worth storing rather than re-reading. | [optional] 
**EmployeeType** | [**EmployeeType**](EmployeeType.md) | The role an account gets by joining through this link. A portal keeps at most one link per role, and the  role of an existing link cannot be changed - the link has to be deleted and created again. | 
**Expiration** | Pointer to [**ApiDateTime**](ApiDateTime.md) | When the link stops working, in the portal time zone. It is empty for a link that never expires, which is  what omitting the deadline on create or update leaves behind. | [optional] 
**IsExpired** | Pointer to **bool** | Whether that deadline has already passed. A link without a deadline always reports `false`, and an expired  link is still returned rather than treated as gone - it can be revived by moving `expiration`. | [optional] 
**MaxUseCount** | Pointer to **NullableInt32** | How many accounts may join through the link in total. It is empty for a link with no use limit, and an  update may not lower it below `currentUseCount`. | [optional] 
**CurrentUseCount** | Pointer to **int32** | How many accounts have already joined through the link. It only ever grows, and reaching `maxUseCount`  retires the link as surely as a passed deadline. | [optional] 
**Url** | Pointer to **NullableString** | The shortened address to hand to the people being invited. It is signed for the account that read it, so  two administrators are given two different URLs for one and the same link and both of them work; the `id`  above, not this string, is what identifies the link. | [optional] 

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

`func (o *InvitationLinkDto) GetExpiration() ApiDateTime`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *InvitationLinkDto) GetExpirationOk() (*ApiDateTime, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *InvitationLinkDto) SetExpiration(v ApiDateTime)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *InvitationLinkDto) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

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


