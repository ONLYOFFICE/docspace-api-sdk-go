# InviteUsersRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invitations** | [**[]UserInvitationRequestDto**](UserInvitationRequestDto.md) | The list of user invitations. | 
**Culture** | Pointer to **NullableString** | The culture code of invitations. | [optional] 

## Methods

### NewInviteUsersRequestDto

`func NewInviteUsersRequestDto(invitations []UserInvitationRequestDto, ) *InviteUsersRequestDto`

NewInviteUsersRequestDto instantiates a new InviteUsersRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInviteUsersRequestDtoWithDefaults

`func NewInviteUsersRequestDtoWithDefaults() *InviteUsersRequestDto`

NewInviteUsersRequestDtoWithDefaults instantiates a new InviteUsersRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvitations

`func (o *InviteUsersRequestDto) GetInvitations() []UserInvitationRequestDto`

GetInvitations returns the Invitations field if non-nil, zero value otherwise.

### GetInvitationsOk

`func (o *InviteUsersRequestDto) GetInvitationsOk() (*[]UserInvitationRequestDto, bool)`

GetInvitationsOk returns a tuple with the Invitations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitations

`func (o *InviteUsersRequestDto) SetInvitations(v []UserInvitationRequestDto)`

SetInvitations sets Invitations field to given value.


### GetCulture

`func (o *InviteUsersRequestDto) GetCulture() string`

GetCulture returns the Culture field if non-nil, zero value otherwise.

### GetCultureOk

`func (o *InviteUsersRequestDto) GetCultureOk() (*string, bool)`

GetCultureOk returns a tuple with the Culture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCulture

`func (o *InviteUsersRequestDto) SetCulture(v string)`

SetCulture sets Culture field to given value.

### HasCulture

`func (o *InviteUsersRequestDto) HasCulture() bool`

HasCulture returns a boolean if a field has been set.

### SetCultureNil

`func (o *InviteUsersRequestDto) SetCultureNil(b bool)`

 SetCultureNil sets the value for Culture to be an explicit nil

### UnsetCulture
`func (o *InviteUsersRequestDto) UnsetCulture()`

UnsetCulture ensures that no value is present for Culture, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


