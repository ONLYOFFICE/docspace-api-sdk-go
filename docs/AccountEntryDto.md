# AccountEntryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The group ID. | 
**DisplayName** | Pointer to **NullableString** | The HTML-encoded user's display name formatted according to the default format for the current culture. | [optional] 
**Avatar** | Pointer to **NullableString** | The user avatar. | [optional] 
**AvatarOriginal** | Pointer to **NullableString** | The user original size avatar. | [optional] 
**AvatarMax** | Pointer to **NullableString** | The user maximum size avatar. | [optional] 
**AvatarMedium** | Pointer to **NullableString** | The user medium size avatar. | [optional] 
**AvatarSmall** | Pointer to **NullableString** | The user small size avatar. | [optional] 
**ProfileUrl** | Pointer to **NullableString** | The user profile URL. | [optional] 
**HasAvatar** | Pointer to **bool** | Specifies if the user has an avatar or not. | [optional] 
**IsAnonim** | Pointer to **bool** | Specifies if the user is anonymous or not. | [optional] 
**FirstName** | Pointer to **NullableString** | The user first name. | [optional] 
**LastName** | Pointer to **NullableString** | The user last name. | [optional] 
**UserName** | Pointer to **NullableString** | The user username. | [optional] 
**Email** | Pointer to **NullableString** | The user email. | [optional] 
**Contacts** | Pointer to [**[]Contact**](Contact.md) | The list of user contacts. | [optional] 
**Status** | Pointer to [**EmployeeStatus**](EmployeeStatus.md) | The user status. | [optional] 
**ActivationStatus** | Pointer to [**EmployeeActivationStatus**](EmployeeActivationStatus.md) | The user activation status. | [optional] 
**Terminated** | Pointer to [**ApiDateTime**](ApiDateTime.md) | The date when the user account was terminated. | [optional] 
**Department** | Pointer to **NullableString** | The user department. | [optional] 
**Groups** | Pointer to [**[]GroupSummaryDto**](GroupSummaryDto.md) | The list of user groups. | [optional] 
**Location** | Pointer to **NullableString** | The user location. | [optional] 
**Notes** | Pointer to **NullableString** | The user notes. | [optional] 
**IsAdmin** | Pointer to **bool** | Specifies if the user is an administrator or not. | [optional] 
**IsRoomAdmin** | Pointer to **bool** | Specifies if the user is a room administrator or not. | [optional] 
**IsLDAP** | **bool** | Specifies if the LDAP settings are enabled for the group or not. | 
**ListAdminModules** | Pointer to **[]string** | The list of the administrator modules. | [optional] 
**IsOwner** | Pointer to **bool** | Specifies if the user is a portal owner or not. | [optional] 
**IsVisitor** | Pointer to **bool** | Specifies if the user is a portal visitor or not. | [optional] 
**IsCollaborator** | Pointer to **bool** | Specifies if the user is a portal collaborator or not. | [optional] 
**CultureName** | Pointer to **NullableString** | The user culture code. | [optional] 
**MobilePhone** | Pointer to **NullableString** | The user mobile phone number. | [optional] 
**MobilePhoneActivationStatus** | Pointer to [**MobilePhoneActivationStatus**](MobilePhoneActivationStatus.md) | The mobile phone activation status. | [optional] 
**IsSSO** | Pointer to **bool** | Specifies if the SSO settings are enabled for the user or not. | [optional] 
**Theme** | Pointer to [**DarkThemeSettingsType**](DarkThemeSettingsType.md) | The user theme settings. | [optional] 
**QuotaLimit** | Pointer to **NullableInt64** | The user quota limit. | [optional] 
**UsedSpace** | Pointer to **NullableFloat64** | The portal used space of the user. | [optional] 
**Shared** | Pointer to **NullableBool** | Specifies whether the group can be shared or not. | [optional] 
**IsCustomQuota** | Pointer to **NullableBool** | Specifies if the user has a custom quota or not. | [optional] 
**AiServicesLimit** | Pointer to [**UserServiceLimitDto**](UserServiceLimitDto.md) | The spending limit on the AI services in force for this user - the personal one, or else the limit for every  user - with what the user spent against it. Null when no such limit applies to the user, when the portal has no  billing service, or when the caller is neither a DocSpace administrator nor the user; the limit of the portal as  a whole is not reflected here. The value, and the spending in particular, can lag by a few minutes. | [optional] 
**LoginEventId** | Pointer to **NullableInt32** | The current login event ID. | [optional] 
**AuthCookieLifetime** | Pointer to **NullableFloat64** | The auth cookie lifetime in seconds. | [optional] 
**CreatedBy** | Pointer to [**EmployeeDto**](EmployeeDto.md) | The user who created the current user. | [optional] 
**RegistrationDate** | Pointer to [**ApiDateTime**](ApiDateTime.md) | The user registration date. | [optional] 
**HasPersonalFolder** | Pointer to **NullableBool** | Specifies if the user has a personal folder or not. | [optional] 
**TfaAppEnabled** | Pointer to **NullableBool** | Indicates whether the user has enabled two-factor authentication (TFA) using an authentication app. | [optional] 
**Name** | **NullableString** | The group name. | 
**Parent** | Pointer to **NullableString** | The parent group ID. | [optional] 
**Category** | **string** | The group category ID. | 
**IsSystem** | Pointer to **NullableBool** | Indicates whether the group is a system group. | [optional] 
**Manager** | Pointer to [**EmployeeFullDto**](EmployeeFullDto.md) | The group manager full information. | [optional] 
**Members** | Pointer to [**[]EmployeeFullDto**](EmployeeFullDto.md) | The list of group members. | [optional] 
**MembersCount** | Pointer to **int32** | The number of group members. | [optional] 

## Methods

### NewAccountEntryDto

`func NewAccountEntryDto(id string, isLDAP bool, name NullableString, category string, ) *AccountEntryDto`

NewAccountEntryDto instantiates a new AccountEntryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountEntryDtoWithDefaults

`func NewAccountEntryDtoWithDefaults() *AccountEntryDto`

NewAccountEntryDtoWithDefaults instantiates a new AccountEntryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountEntryDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountEntryDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountEntryDto) SetId(v string)`

SetId sets Id field to given value.


### GetDisplayName

`func (o *AccountEntryDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AccountEntryDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AccountEntryDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AccountEntryDto) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *AccountEntryDto) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *AccountEntryDto) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetAvatar

`func (o *AccountEntryDto) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *AccountEntryDto) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *AccountEntryDto) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *AccountEntryDto) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### SetAvatarNil

`func (o *AccountEntryDto) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *AccountEntryDto) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetAvatarOriginal

`func (o *AccountEntryDto) GetAvatarOriginal() string`

GetAvatarOriginal returns the AvatarOriginal field if non-nil, zero value otherwise.

### GetAvatarOriginalOk

`func (o *AccountEntryDto) GetAvatarOriginalOk() (*string, bool)`

GetAvatarOriginalOk returns a tuple with the AvatarOriginal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarOriginal

`func (o *AccountEntryDto) SetAvatarOriginal(v string)`

SetAvatarOriginal sets AvatarOriginal field to given value.

### HasAvatarOriginal

`func (o *AccountEntryDto) HasAvatarOriginal() bool`

HasAvatarOriginal returns a boolean if a field has been set.

### SetAvatarOriginalNil

`func (o *AccountEntryDto) SetAvatarOriginalNil(b bool)`

 SetAvatarOriginalNil sets the value for AvatarOriginal to be an explicit nil

### UnsetAvatarOriginal
`func (o *AccountEntryDto) UnsetAvatarOriginal()`

UnsetAvatarOriginal ensures that no value is present for AvatarOriginal, not even an explicit nil
### GetAvatarMax

`func (o *AccountEntryDto) GetAvatarMax() string`

GetAvatarMax returns the AvatarMax field if non-nil, zero value otherwise.

### GetAvatarMaxOk

`func (o *AccountEntryDto) GetAvatarMaxOk() (*string, bool)`

GetAvatarMaxOk returns a tuple with the AvatarMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarMax

`func (o *AccountEntryDto) SetAvatarMax(v string)`

SetAvatarMax sets AvatarMax field to given value.

### HasAvatarMax

`func (o *AccountEntryDto) HasAvatarMax() bool`

HasAvatarMax returns a boolean if a field has been set.

### SetAvatarMaxNil

`func (o *AccountEntryDto) SetAvatarMaxNil(b bool)`

 SetAvatarMaxNil sets the value for AvatarMax to be an explicit nil

### UnsetAvatarMax
`func (o *AccountEntryDto) UnsetAvatarMax()`

UnsetAvatarMax ensures that no value is present for AvatarMax, not even an explicit nil
### GetAvatarMedium

`func (o *AccountEntryDto) GetAvatarMedium() string`

GetAvatarMedium returns the AvatarMedium field if non-nil, zero value otherwise.

### GetAvatarMediumOk

`func (o *AccountEntryDto) GetAvatarMediumOk() (*string, bool)`

GetAvatarMediumOk returns a tuple with the AvatarMedium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarMedium

`func (o *AccountEntryDto) SetAvatarMedium(v string)`

SetAvatarMedium sets AvatarMedium field to given value.

### HasAvatarMedium

`func (o *AccountEntryDto) HasAvatarMedium() bool`

HasAvatarMedium returns a boolean if a field has been set.

### SetAvatarMediumNil

`func (o *AccountEntryDto) SetAvatarMediumNil(b bool)`

 SetAvatarMediumNil sets the value for AvatarMedium to be an explicit nil

### UnsetAvatarMedium
`func (o *AccountEntryDto) UnsetAvatarMedium()`

UnsetAvatarMedium ensures that no value is present for AvatarMedium, not even an explicit nil
### GetAvatarSmall

`func (o *AccountEntryDto) GetAvatarSmall() string`

GetAvatarSmall returns the AvatarSmall field if non-nil, zero value otherwise.

### GetAvatarSmallOk

`func (o *AccountEntryDto) GetAvatarSmallOk() (*string, bool)`

GetAvatarSmallOk returns a tuple with the AvatarSmall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarSmall

`func (o *AccountEntryDto) SetAvatarSmall(v string)`

SetAvatarSmall sets AvatarSmall field to given value.

### HasAvatarSmall

`func (o *AccountEntryDto) HasAvatarSmall() bool`

HasAvatarSmall returns a boolean if a field has been set.

### SetAvatarSmallNil

`func (o *AccountEntryDto) SetAvatarSmallNil(b bool)`

 SetAvatarSmallNil sets the value for AvatarSmall to be an explicit nil

### UnsetAvatarSmall
`func (o *AccountEntryDto) UnsetAvatarSmall()`

UnsetAvatarSmall ensures that no value is present for AvatarSmall, not even an explicit nil
### GetProfileUrl

`func (o *AccountEntryDto) GetProfileUrl() string`

GetProfileUrl returns the ProfileUrl field if non-nil, zero value otherwise.

### GetProfileUrlOk

`func (o *AccountEntryDto) GetProfileUrlOk() (*string, bool)`

GetProfileUrlOk returns a tuple with the ProfileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileUrl

`func (o *AccountEntryDto) SetProfileUrl(v string)`

SetProfileUrl sets ProfileUrl field to given value.

### HasProfileUrl

`func (o *AccountEntryDto) HasProfileUrl() bool`

HasProfileUrl returns a boolean if a field has been set.

### SetProfileUrlNil

`func (o *AccountEntryDto) SetProfileUrlNil(b bool)`

 SetProfileUrlNil sets the value for ProfileUrl to be an explicit nil

### UnsetProfileUrl
`func (o *AccountEntryDto) UnsetProfileUrl()`

UnsetProfileUrl ensures that no value is present for ProfileUrl, not even an explicit nil
### GetHasAvatar

`func (o *AccountEntryDto) GetHasAvatar() bool`

GetHasAvatar returns the HasAvatar field if non-nil, zero value otherwise.

### GetHasAvatarOk

`func (o *AccountEntryDto) GetHasAvatarOk() (*bool, bool)`

GetHasAvatarOk returns a tuple with the HasAvatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAvatar

`func (o *AccountEntryDto) SetHasAvatar(v bool)`

SetHasAvatar sets HasAvatar field to given value.

### HasHasAvatar

`func (o *AccountEntryDto) HasHasAvatar() bool`

HasHasAvatar returns a boolean if a field has been set.

### GetIsAnonim

`func (o *AccountEntryDto) GetIsAnonim() bool`

GetIsAnonim returns the IsAnonim field if non-nil, zero value otherwise.

### GetIsAnonimOk

`func (o *AccountEntryDto) GetIsAnonimOk() (*bool, bool)`

GetIsAnonimOk returns a tuple with the IsAnonim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnonim

`func (o *AccountEntryDto) SetIsAnonim(v bool)`

SetIsAnonim sets IsAnonim field to given value.

### HasIsAnonim

`func (o *AccountEntryDto) HasIsAnonim() bool`

HasIsAnonim returns a boolean if a field has been set.

### GetFirstName

`func (o *AccountEntryDto) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *AccountEntryDto) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *AccountEntryDto) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *AccountEntryDto) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *AccountEntryDto) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *AccountEntryDto) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *AccountEntryDto) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *AccountEntryDto) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *AccountEntryDto) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *AccountEntryDto) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *AccountEntryDto) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *AccountEntryDto) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetUserName

`func (o *AccountEntryDto) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *AccountEntryDto) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *AccountEntryDto) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *AccountEntryDto) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *AccountEntryDto) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *AccountEntryDto) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetEmail

`func (o *AccountEntryDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AccountEntryDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AccountEntryDto) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AccountEntryDto) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *AccountEntryDto) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *AccountEntryDto) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetContacts

`func (o *AccountEntryDto) GetContacts() []Contact`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *AccountEntryDto) GetContactsOk() (*[]Contact, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *AccountEntryDto) SetContacts(v []Contact)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *AccountEntryDto) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### SetContactsNil

`func (o *AccountEntryDto) SetContactsNil(b bool)`

 SetContactsNil sets the value for Contacts to be an explicit nil

### UnsetContacts
`func (o *AccountEntryDto) UnsetContacts()`

UnsetContacts ensures that no value is present for Contacts, not even an explicit nil
### GetStatus

`func (o *AccountEntryDto) GetStatus() EmployeeStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountEntryDto) GetStatusOk() (*EmployeeStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountEntryDto) SetStatus(v EmployeeStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccountEntryDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetActivationStatus

`func (o *AccountEntryDto) GetActivationStatus() EmployeeActivationStatus`

GetActivationStatus returns the ActivationStatus field if non-nil, zero value otherwise.

### GetActivationStatusOk

`func (o *AccountEntryDto) GetActivationStatusOk() (*EmployeeActivationStatus, bool)`

GetActivationStatusOk returns a tuple with the ActivationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationStatus

`func (o *AccountEntryDto) SetActivationStatus(v EmployeeActivationStatus)`

SetActivationStatus sets ActivationStatus field to given value.

### HasActivationStatus

`func (o *AccountEntryDto) HasActivationStatus() bool`

HasActivationStatus returns a boolean if a field has been set.

### GetTerminated

`func (o *AccountEntryDto) GetTerminated() ApiDateTime`

GetTerminated returns the Terminated field if non-nil, zero value otherwise.

### GetTerminatedOk

`func (o *AccountEntryDto) GetTerminatedOk() (*ApiDateTime, bool)`

GetTerminatedOk returns a tuple with the Terminated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminated

`func (o *AccountEntryDto) SetTerminated(v ApiDateTime)`

SetTerminated sets Terminated field to given value.

### HasTerminated

`func (o *AccountEntryDto) HasTerminated() bool`

HasTerminated returns a boolean if a field has been set.

### GetDepartment

`func (o *AccountEntryDto) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *AccountEntryDto) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *AccountEntryDto) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *AccountEntryDto) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartmentNil

`func (o *AccountEntryDto) SetDepartmentNil(b bool)`

 SetDepartmentNil sets the value for Department to be an explicit nil

### UnsetDepartment
`func (o *AccountEntryDto) UnsetDepartment()`

UnsetDepartment ensures that no value is present for Department, not even an explicit nil
### GetGroups

`func (o *AccountEntryDto) GetGroups() []GroupSummaryDto`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *AccountEntryDto) GetGroupsOk() (*[]GroupSummaryDto, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *AccountEntryDto) SetGroups(v []GroupSummaryDto)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *AccountEntryDto) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *AccountEntryDto) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *AccountEntryDto) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetLocation

`func (o *AccountEntryDto) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AccountEntryDto) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AccountEntryDto) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *AccountEntryDto) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *AccountEntryDto) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *AccountEntryDto) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetNotes

`func (o *AccountEntryDto) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *AccountEntryDto) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *AccountEntryDto) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *AccountEntryDto) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *AccountEntryDto) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *AccountEntryDto) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetIsAdmin

`func (o *AccountEntryDto) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *AccountEntryDto) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *AccountEntryDto) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *AccountEntryDto) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### GetIsRoomAdmin

`func (o *AccountEntryDto) GetIsRoomAdmin() bool`

GetIsRoomAdmin returns the IsRoomAdmin field if non-nil, zero value otherwise.

### GetIsRoomAdminOk

`func (o *AccountEntryDto) GetIsRoomAdminOk() (*bool, bool)`

GetIsRoomAdminOk returns a tuple with the IsRoomAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRoomAdmin

`func (o *AccountEntryDto) SetIsRoomAdmin(v bool)`

SetIsRoomAdmin sets IsRoomAdmin field to given value.

### HasIsRoomAdmin

`func (o *AccountEntryDto) HasIsRoomAdmin() bool`

HasIsRoomAdmin returns a boolean if a field has been set.

### GetIsLDAP

`func (o *AccountEntryDto) GetIsLDAP() bool`

GetIsLDAP returns the IsLDAP field if non-nil, zero value otherwise.

### GetIsLDAPOk

`func (o *AccountEntryDto) GetIsLDAPOk() (*bool, bool)`

GetIsLDAPOk returns a tuple with the IsLDAP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLDAP

`func (o *AccountEntryDto) SetIsLDAP(v bool)`

SetIsLDAP sets IsLDAP field to given value.


### GetListAdminModules

`func (o *AccountEntryDto) GetListAdminModules() []string`

GetListAdminModules returns the ListAdminModules field if non-nil, zero value otherwise.

### GetListAdminModulesOk

`func (o *AccountEntryDto) GetListAdminModulesOk() (*[]string, bool)`

GetListAdminModulesOk returns a tuple with the ListAdminModules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListAdminModules

`func (o *AccountEntryDto) SetListAdminModules(v []string)`

SetListAdminModules sets ListAdminModules field to given value.

### HasListAdminModules

`func (o *AccountEntryDto) HasListAdminModules() bool`

HasListAdminModules returns a boolean if a field has been set.

### SetListAdminModulesNil

`func (o *AccountEntryDto) SetListAdminModulesNil(b bool)`

 SetListAdminModulesNil sets the value for ListAdminModules to be an explicit nil

### UnsetListAdminModules
`func (o *AccountEntryDto) UnsetListAdminModules()`

UnsetListAdminModules ensures that no value is present for ListAdminModules, not even an explicit nil
### GetIsOwner

`func (o *AccountEntryDto) GetIsOwner() bool`

GetIsOwner returns the IsOwner field if non-nil, zero value otherwise.

### GetIsOwnerOk

`func (o *AccountEntryDto) GetIsOwnerOk() (*bool, bool)`

GetIsOwnerOk returns a tuple with the IsOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOwner

`func (o *AccountEntryDto) SetIsOwner(v bool)`

SetIsOwner sets IsOwner field to given value.

### HasIsOwner

`func (o *AccountEntryDto) HasIsOwner() bool`

HasIsOwner returns a boolean if a field has been set.

### GetIsVisitor

`func (o *AccountEntryDto) GetIsVisitor() bool`

GetIsVisitor returns the IsVisitor field if non-nil, zero value otherwise.

### GetIsVisitorOk

`func (o *AccountEntryDto) GetIsVisitorOk() (*bool, bool)`

GetIsVisitorOk returns a tuple with the IsVisitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVisitor

`func (o *AccountEntryDto) SetIsVisitor(v bool)`

SetIsVisitor sets IsVisitor field to given value.

### HasIsVisitor

`func (o *AccountEntryDto) HasIsVisitor() bool`

HasIsVisitor returns a boolean if a field has been set.

### GetIsCollaborator

`func (o *AccountEntryDto) GetIsCollaborator() bool`

GetIsCollaborator returns the IsCollaborator field if non-nil, zero value otherwise.

### GetIsCollaboratorOk

`func (o *AccountEntryDto) GetIsCollaboratorOk() (*bool, bool)`

GetIsCollaboratorOk returns a tuple with the IsCollaborator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCollaborator

`func (o *AccountEntryDto) SetIsCollaborator(v bool)`

SetIsCollaborator sets IsCollaborator field to given value.

### HasIsCollaborator

`func (o *AccountEntryDto) HasIsCollaborator() bool`

HasIsCollaborator returns a boolean if a field has been set.

### GetCultureName

`func (o *AccountEntryDto) GetCultureName() string`

GetCultureName returns the CultureName field if non-nil, zero value otherwise.

### GetCultureNameOk

`func (o *AccountEntryDto) GetCultureNameOk() (*string, bool)`

GetCultureNameOk returns a tuple with the CultureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCultureName

`func (o *AccountEntryDto) SetCultureName(v string)`

SetCultureName sets CultureName field to given value.

### HasCultureName

`func (o *AccountEntryDto) HasCultureName() bool`

HasCultureName returns a boolean if a field has been set.

### SetCultureNameNil

`func (o *AccountEntryDto) SetCultureNameNil(b bool)`

 SetCultureNameNil sets the value for CultureName to be an explicit nil

### UnsetCultureName
`func (o *AccountEntryDto) UnsetCultureName()`

UnsetCultureName ensures that no value is present for CultureName, not even an explicit nil
### GetMobilePhone

`func (o *AccountEntryDto) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *AccountEntryDto) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *AccountEntryDto) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *AccountEntryDto) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### SetMobilePhoneNil

`func (o *AccountEntryDto) SetMobilePhoneNil(b bool)`

 SetMobilePhoneNil sets the value for MobilePhone to be an explicit nil

### UnsetMobilePhone
`func (o *AccountEntryDto) UnsetMobilePhone()`

UnsetMobilePhone ensures that no value is present for MobilePhone, not even an explicit nil
### GetMobilePhoneActivationStatus

`func (o *AccountEntryDto) GetMobilePhoneActivationStatus() MobilePhoneActivationStatus`

GetMobilePhoneActivationStatus returns the MobilePhoneActivationStatus field if non-nil, zero value otherwise.

### GetMobilePhoneActivationStatusOk

`func (o *AccountEntryDto) GetMobilePhoneActivationStatusOk() (*MobilePhoneActivationStatus, bool)`

GetMobilePhoneActivationStatusOk returns a tuple with the MobilePhoneActivationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhoneActivationStatus

`func (o *AccountEntryDto) SetMobilePhoneActivationStatus(v MobilePhoneActivationStatus)`

SetMobilePhoneActivationStatus sets MobilePhoneActivationStatus field to given value.

### HasMobilePhoneActivationStatus

`func (o *AccountEntryDto) HasMobilePhoneActivationStatus() bool`

HasMobilePhoneActivationStatus returns a boolean if a field has been set.

### GetIsSSO

`func (o *AccountEntryDto) GetIsSSO() bool`

GetIsSSO returns the IsSSO field if non-nil, zero value otherwise.

### GetIsSSOOk

`func (o *AccountEntryDto) GetIsSSOOk() (*bool, bool)`

GetIsSSOOk returns a tuple with the IsSSO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSSO

`func (o *AccountEntryDto) SetIsSSO(v bool)`

SetIsSSO sets IsSSO field to given value.

### HasIsSSO

`func (o *AccountEntryDto) HasIsSSO() bool`

HasIsSSO returns a boolean if a field has been set.

### GetTheme

`func (o *AccountEntryDto) GetTheme() DarkThemeSettingsType`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *AccountEntryDto) GetThemeOk() (*DarkThemeSettingsType, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *AccountEntryDto) SetTheme(v DarkThemeSettingsType)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *AccountEntryDto) HasTheme() bool`

HasTheme returns a boolean if a field has been set.

### GetQuotaLimit

`func (o *AccountEntryDto) GetQuotaLimit() int64`

GetQuotaLimit returns the QuotaLimit field if non-nil, zero value otherwise.

### GetQuotaLimitOk

`func (o *AccountEntryDto) GetQuotaLimitOk() (*int64, bool)`

GetQuotaLimitOk returns a tuple with the QuotaLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaLimit

`func (o *AccountEntryDto) SetQuotaLimit(v int64)`

SetQuotaLimit sets QuotaLimit field to given value.

### HasQuotaLimit

`func (o *AccountEntryDto) HasQuotaLimit() bool`

HasQuotaLimit returns a boolean if a field has been set.

### SetQuotaLimitNil

`func (o *AccountEntryDto) SetQuotaLimitNil(b bool)`

 SetQuotaLimitNil sets the value for QuotaLimit to be an explicit nil

### UnsetQuotaLimit
`func (o *AccountEntryDto) UnsetQuotaLimit()`

UnsetQuotaLimit ensures that no value is present for QuotaLimit, not even an explicit nil
### GetUsedSpace

`func (o *AccountEntryDto) GetUsedSpace() float64`

GetUsedSpace returns the UsedSpace field if non-nil, zero value otherwise.

### GetUsedSpaceOk

`func (o *AccountEntryDto) GetUsedSpaceOk() (*float64, bool)`

GetUsedSpaceOk returns a tuple with the UsedSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedSpace

`func (o *AccountEntryDto) SetUsedSpace(v float64)`

SetUsedSpace sets UsedSpace field to given value.

### HasUsedSpace

`func (o *AccountEntryDto) HasUsedSpace() bool`

HasUsedSpace returns a boolean if a field has been set.

### SetUsedSpaceNil

`func (o *AccountEntryDto) SetUsedSpaceNil(b bool)`

 SetUsedSpaceNil sets the value for UsedSpace to be an explicit nil

### UnsetUsedSpace
`func (o *AccountEntryDto) UnsetUsedSpace()`

UnsetUsedSpace ensures that no value is present for UsedSpace, not even an explicit nil
### GetShared

`func (o *AccountEntryDto) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *AccountEntryDto) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *AccountEntryDto) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *AccountEntryDto) HasShared() bool`

HasShared returns a boolean if a field has been set.

### SetSharedNil

`func (o *AccountEntryDto) SetSharedNil(b bool)`

 SetSharedNil sets the value for Shared to be an explicit nil

### UnsetShared
`func (o *AccountEntryDto) UnsetShared()`

UnsetShared ensures that no value is present for Shared, not even an explicit nil
### GetIsCustomQuota

`func (o *AccountEntryDto) GetIsCustomQuota() bool`

GetIsCustomQuota returns the IsCustomQuota field if non-nil, zero value otherwise.

### GetIsCustomQuotaOk

`func (o *AccountEntryDto) GetIsCustomQuotaOk() (*bool, bool)`

GetIsCustomQuotaOk returns a tuple with the IsCustomQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomQuota

`func (o *AccountEntryDto) SetIsCustomQuota(v bool)`

SetIsCustomQuota sets IsCustomQuota field to given value.

### HasIsCustomQuota

`func (o *AccountEntryDto) HasIsCustomQuota() bool`

HasIsCustomQuota returns a boolean if a field has been set.

### SetIsCustomQuotaNil

`func (o *AccountEntryDto) SetIsCustomQuotaNil(b bool)`

 SetIsCustomQuotaNil sets the value for IsCustomQuota to be an explicit nil

### UnsetIsCustomQuota
`func (o *AccountEntryDto) UnsetIsCustomQuota()`

UnsetIsCustomQuota ensures that no value is present for IsCustomQuota, not even an explicit nil
### GetAiServicesLimit

`func (o *AccountEntryDto) GetAiServicesLimit() UserServiceLimitDto`

GetAiServicesLimit returns the AiServicesLimit field if non-nil, zero value otherwise.

### GetAiServicesLimitOk

`func (o *AccountEntryDto) GetAiServicesLimitOk() (*UserServiceLimitDto, bool)`

GetAiServicesLimitOk returns a tuple with the AiServicesLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiServicesLimit

`func (o *AccountEntryDto) SetAiServicesLimit(v UserServiceLimitDto)`

SetAiServicesLimit sets AiServicesLimit field to given value.

### HasAiServicesLimit

`func (o *AccountEntryDto) HasAiServicesLimit() bool`

HasAiServicesLimit returns a boolean if a field has been set.

### GetLoginEventId

`func (o *AccountEntryDto) GetLoginEventId() int32`

GetLoginEventId returns the LoginEventId field if non-nil, zero value otherwise.

### GetLoginEventIdOk

`func (o *AccountEntryDto) GetLoginEventIdOk() (*int32, bool)`

GetLoginEventIdOk returns a tuple with the LoginEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginEventId

`func (o *AccountEntryDto) SetLoginEventId(v int32)`

SetLoginEventId sets LoginEventId field to given value.

### HasLoginEventId

`func (o *AccountEntryDto) HasLoginEventId() bool`

HasLoginEventId returns a boolean if a field has been set.

### SetLoginEventIdNil

`func (o *AccountEntryDto) SetLoginEventIdNil(b bool)`

 SetLoginEventIdNil sets the value for LoginEventId to be an explicit nil

### UnsetLoginEventId
`func (o *AccountEntryDto) UnsetLoginEventId()`

UnsetLoginEventId ensures that no value is present for LoginEventId, not even an explicit nil
### GetAuthCookieLifetime

`func (o *AccountEntryDto) GetAuthCookieLifetime() float64`

GetAuthCookieLifetime returns the AuthCookieLifetime field if non-nil, zero value otherwise.

### GetAuthCookieLifetimeOk

`func (o *AccountEntryDto) GetAuthCookieLifetimeOk() (*float64, bool)`

GetAuthCookieLifetimeOk returns a tuple with the AuthCookieLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCookieLifetime

`func (o *AccountEntryDto) SetAuthCookieLifetime(v float64)`

SetAuthCookieLifetime sets AuthCookieLifetime field to given value.

### HasAuthCookieLifetime

`func (o *AccountEntryDto) HasAuthCookieLifetime() bool`

HasAuthCookieLifetime returns a boolean if a field has been set.

### SetAuthCookieLifetimeNil

`func (o *AccountEntryDto) SetAuthCookieLifetimeNil(b bool)`

 SetAuthCookieLifetimeNil sets the value for AuthCookieLifetime to be an explicit nil

### UnsetAuthCookieLifetime
`func (o *AccountEntryDto) UnsetAuthCookieLifetime()`

UnsetAuthCookieLifetime ensures that no value is present for AuthCookieLifetime, not even an explicit nil
### GetCreatedBy

`func (o *AccountEntryDto) GetCreatedBy() EmployeeDto`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AccountEntryDto) GetCreatedByOk() (*EmployeeDto, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AccountEntryDto) SetCreatedBy(v EmployeeDto)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AccountEntryDto) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetRegistrationDate

`func (o *AccountEntryDto) GetRegistrationDate() ApiDateTime`

GetRegistrationDate returns the RegistrationDate field if non-nil, zero value otherwise.

### GetRegistrationDateOk

`func (o *AccountEntryDto) GetRegistrationDateOk() (*ApiDateTime, bool)`

GetRegistrationDateOk returns a tuple with the RegistrationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationDate

`func (o *AccountEntryDto) SetRegistrationDate(v ApiDateTime)`

SetRegistrationDate sets RegistrationDate field to given value.

### HasRegistrationDate

`func (o *AccountEntryDto) HasRegistrationDate() bool`

HasRegistrationDate returns a boolean if a field has been set.

### GetHasPersonalFolder

`func (o *AccountEntryDto) GetHasPersonalFolder() bool`

GetHasPersonalFolder returns the HasPersonalFolder field if non-nil, zero value otherwise.

### GetHasPersonalFolderOk

`func (o *AccountEntryDto) GetHasPersonalFolderOk() (*bool, bool)`

GetHasPersonalFolderOk returns a tuple with the HasPersonalFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPersonalFolder

`func (o *AccountEntryDto) SetHasPersonalFolder(v bool)`

SetHasPersonalFolder sets HasPersonalFolder field to given value.

### HasHasPersonalFolder

`func (o *AccountEntryDto) HasHasPersonalFolder() bool`

HasHasPersonalFolder returns a boolean if a field has been set.

### SetHasPersonalFolderNil

`func (o *AccountEntryDto) SetHasPersonalFolderNil(b bool)`

 SetHasPersonalFolderNil sets the value for HasPersonalFolder to be an explicit nil

### UnsetHasPersonalFolder
`func (o *AccountEntryDto) UnsetHasPersonalFolder()`

UnsetHasPersonalFolder ensures that no value is present for HasPersonalFolder, not even an explicit nil
### GetTfaAppEnabled

`func (o *AccountEntryDto) GetTfaAppEnabled() bool`

GetTfaAppEnabled returns the TfaAppEnabled field if non-nil, zero value otherwise.

### GetTfaAppEnabledOk

`func (o *AccountEntryDto) GetTfaAppEnabledOk() (*bool, bool)`

GetTfaAppEnabledOk returns a tuple with the TfaAppEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTfaAppEnabled

`func (o *AccountEntryDto) SetTfaAppEnabled(v bool)`

SetTfaAppEnabled sets TfaAppEnabled field to given value.

### HasTfaAppEnabled

`func (o *AccountEntryDto) HasTfaAppEnabled() bool`

HasTfaAppEnabled returns a boolean if a field has been set.

### SetTfaAppEnabledNil

`func (o *AccountEntryDto) SetTfaAppEnabledNil(b bool)`

 SetTfaAppEnabledNil sets the value for TfaAppEnabled to be an explicit nil

### UnsetTfaAppEnabled
`func (o *AccountEntryDto) UnsetTfaAppEnabled()`

UnsetTfaAppEnabled ensures that no value is present for TfaAppEnabled, not even an explicit nil
### GetName

`func (o *AccountEntryDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountEntryDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountEntryDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *AccountEntryDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AccountEntryDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetParent

`func (o *AccountEntryDto) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *AccountEntryDto) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *AccountEntryDto) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *AccountEntryDto) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *AccountEntryDto) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *AccountEntryDto) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetCategory

`func (o *AccountEntryDto) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AccountEntryDto) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AccountEntryDto) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetIsSystem

`func (o *AccountEntryDto) GetIsSystem() bool`

GetIsSystem returns the IsSystem field if non-nil, zero value otherwise.

### GetIsSystemOk

`func (o *AccountEntryDto) GetIsSystemOk() (*bool, bool)`

GetIsSystemOk returns a tuple with the IsSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystem

`func (o *AccountEntryDto) SetIsSystem(v bool)`

SetIsSystem sets IsSystem field to given value.

### HasIsSystem

`func (o *AccountEntryDto) HasIsSystem() bool`

HasIsSystem returns a boolean if a field has been set.

### SetIsSystemNil

`func (o *AccountEntryDto) SetIsSystemNil(b bool)`

 SetIsSystemNil sets the value for IsSystem to be an explicit nil

### UnsetIsSystem
`func (o *AccountEntryDto) UnsetIsSystem()`

UnsetIsSystem ensures that no value is present for IsSystem, not even an explicit nil
### GetManager

`func (o *AccountEntryDto) GetManager() EmployeeFullDto`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *AccountEntryDto) GetManagerOk() (*EmployeeFullDto, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *AccountEntryDto) SetManager(v EmployeeFullDto)`

SetManager sets Manager field to given value.

### HasManager

`func (o *AccountEntryDto) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetMembers

`func (o *AccountEntryDto) GetMembers() []EmployeeFullDto`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *AccountEntryDto) GetMembersOk() (*[]EmployeeFullDto, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *AccountEntryDto) SetMembers(v []EmployeeFullDto)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *AccountEntryDto) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### SetMembersNil

`func (o *AccountEntryDto) SetMembersNil(b bool)`

 SetMembersNil sets the value for Members to be an explicit nil

### UnsetMembers
`func (o *AccountEntryDto) UnsetMembers()`

UnsetMembers ensures that no value is present for Members, not even an explicit nil
### GetMembersCount

`func (o *AccountEntryDto) GetMembersCount() int32`

GetMembersCount returns the MembersCount field if non-nil, zero value otherwise.

### GetMembersCountOk

`func (o *AccountEntryDto) GetMembersCountOk() (*int32, bool)`

GetMembersCountOk returns a tuple with the MembersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersCount

`func (o *AccountEntryDto) SetMembersCount(v int32)`

SetMembersCount sets MembersCount field to given value.

### HasMembersCount

`func (o *AccountEntryDto) HasMembersCount() bool`

HasMembersCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


