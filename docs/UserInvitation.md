# UserInvitation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsersIds** | Pointer to **[]string** | The list of user IDs. | [optional] 
**ResendAll** | Pointer to **bool** | Specifies whether to resend all user invitations or not. | [optional] 

## Methods

### NewUserInvitation

`func NewUserInvitation() *UserInvitation`

NewUserInvitation instantiates a new UserInvitation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserInvitationWithDefaults

`func NewUserInvitationWithDefaults() *UserInvitation`

NewUserInvitationWithDefaults instantiates a new UserInvitation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsersIds

`func (o *UserInvitation) GetUsersIds() []string`

GetUsersIds returns the UsersIds field if non-nil, zero value otherwise.

### GetUsersIdsOk

`func (o *UserInvitation) GetUsersIdsOk() (*[]string, bool)`

GetUsersIdsOk returns a tuple with the UsersIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersIds

`func (o *UserInvitation) SetUsersIds(v []string)`

SetUsersIds sets UsersIds field to given value.

### HasUsersIds

`func (o *UserInvitation) HasUsersIds() bool`

HasUsersIds returns a boolean if a field has been set.

### SetUsersIdsNil

`func (o *UserInvitation) SetUsersIdsNil(b bool)`

 SetUsersIdsNil sets the value for UsersIds to be an explicit nil

### UnsetUsersIds
`func (o *UserInvitation) UnsetUsersIds()`

UnsetUsersIds ensures that no value is present for UsersIds, not even an explicit nil
### GetResendAll

`func (o *UserInvitation) GetResendAll() bool`

GetResendAll returns the ResendAll field if non-nil, zero value otherwise.

### GetResendAllOk

`func (o *UserInvitation) GetResendAllOk() (*bool, bool)`

GetResendAllOk returns a tuple with the ResendAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResendAll

`func (o *UserInvitation) SetResendAll(v bool)`

SetResendAll sets ResendAll field to given value.

### HasResendAll

`func (o *UserInvitation) HasResendAll() bool`

HasResendAll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


