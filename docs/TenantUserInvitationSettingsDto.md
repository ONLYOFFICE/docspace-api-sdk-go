# TenantUserInvitationSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowInvitingMembers** | **bool** | Specifies whether to allow inviting new DocSpace members through the Contacts section. | 
**AllowInvitingGuests** | **bool** | Specifies whether to allow all DocSpace members to invite external guests to the rooms. | 

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


