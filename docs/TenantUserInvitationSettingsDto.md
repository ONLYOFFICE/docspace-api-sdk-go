# TenantUserInvitationSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowInvitingMembers** | **bool** | Whether new members may be invited through the Contacts section. Switching it off stops new invitations  from being created; links already handed out keep working and members already invited stay. | 
**AllowInvitingGuests** | **bool** | Whether every member, and not only an administrator, may invite an outside guest into a room. It is  independent of `allowInvitingMembers`, and switching it off has the same forward-only effect. | 

## Methods

### NewTenantUserInvitationSettingsDto

`func NewTenantUserInvitationSettingsDto(allowInvitingMembers bool, allowInvitingGuests bool, ) *TenantUserInvitationSettingsDto`

NewTenantUserInvitationSettingsDto instantiates a new TenantUserInvitationSettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantUserInvitationSettingsDtoWithDefaults

`func NewTenantUserInvitationSettingsDtoWithDefaults() *TenantUserInvitationSettingsDto`

NewTenantUserInvitationSettingsDtoWithDefaults instantiates a new TenantUserInvitationSettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowInvitingMembers

`func (o *TenantUserInvitationSettingsDto) GetAllowInvitingMembers() bool`

GetAllowInvitingMembers returns the AllowInvitingMembers field if non-nil, zero value otherwise.

### GetAllowInvitingMembersOk

`func (o *TenantUserInvitationSettingsDto) GetAllowInvitingMembersOk() (*bool, bool)`

GetAllowInvitingMembersOk returns a tuple with the AllowInvitingMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInvitingMembers

`func (o *TenantUserInvitationSettingsDto) SetAllowInvitingMembers(v bool)`

SetAllowInvitingMembers sets AllowInvitingMembers field to given value.


### GetAllowInvitingGuests

`func (o *TenantUserInvitationSettingsDto) GetAllowInvitingGuests() bool`

GetAllowInvitingGuests returns the AllowInvitingGuests field if non-nil, zero value otherwise.

### GetAllowInvitingGuestsOk

`func (o *TenantUserInvitationSettingsDto) GetAllowInvitingGuestsOk() (*bool, bool)`

GetAllowInvitingGuestsOk returns a tuple with the AllowInvitingGuests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInvitingGuests

`func (o *TenantUserInvitationSettingsDto) SetAllowInvitingGuests(v bool)`

SetAllowInvitingGuests sets AllowInvitingGuests field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


