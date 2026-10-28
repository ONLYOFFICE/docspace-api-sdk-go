# TenantUserInvitationSettingsRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowInvitingMembers** | Pointer to **bool** | Whether new DocSpace members may be invited through the Contacts section. Switching it off only stops new  invitations being created; links already issued keep working and members already invited stay. | [optional] 
**AllowInvitingGuests** | Pointer to **bool** | Whether every DocSpace member, and not only an administrator, may invite external guests into rooms.  Switching it off leaves the guests already invited in place. | [optional] 

## Methods

### NewTenantUserInvitationSettingsRequestDto

`func NewTenantUserInvitationSettingsRequestDto() *TenantUserInvitationSettingsRequestDto`

NewTenantUserInvitationSettingsRequestDto instantiates a new TenantUserInvitationSettingsRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantUserInvitationSettingsRequestDtoWithDefaults

`func NewTenantUserInvitationSettingsRequestDtoWithDefaults() *TenantUserInvitationSettingsRequestDto`

NewTenantUserInvitationSettingsRequestDtoWithDefaults instantiates a new TenantUserInvitationSettingsRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowInvitingMembers

`func (o *TenantUserInvitationSettingsRequestDto) GetAllowInvitingMembers() bool`

GetAllowInvitingMembers returns the AllowInvitingMembers field if non-nil, zero value otherwise.

### GetAllowInvitingMembersOk

`func (o *TenantUserInvitationSettingsRequestDto) GetAllowInvitingMembersOk() (*bool, bool)`

GetAllowInvitingMembersOk returns a tuple with the AllowInvitingMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInvitingMembers

`func (o *TenantUserInvitationSettingsRequestDto) SetAllowInvitingMembers(v bool)`

SetAllowInvitingMembers sets AllowInvitingMembers field to given value.

### HasAllowInvitingMembers

`func (o *TenantUserInvitationSettingsRequestDto) HasAllowInvitingMembers() bool`

HasAllowInvitingMembers returns a boolean if a field has been set.

### GetAllowInvitingGuests

`func (o *TenantUserInvitationSettingsRequestDto) GetAllowInvitingGuests() bool`

GetAllowInvitingGuests returns the AllowInvitingGuests field if non-nil, zero value otherwise.

### GetAllowInvitingGuestsOk

`func (o *TenantUserInvitationSettingsRequestDto) GetAllowInvitingGuestsOk() (*bool, bool)`

GetAllowInvitingGuestsOk returns a tuple with the AllowInvitingGuests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInvitingGuests

`func (o *TenantUserInvitationSettingsRequestDto) SetAllowInvitingGuests(v bool)`

SetAllowInvitingGuests sets AllowInvitingGuests field to given value.

### HasAllowInvitingGuests

`func (o *TenantUserInvitationSettingsRequestDto) HasAllowInvitingGuests() bool`

HasAllowInvitingGuests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


