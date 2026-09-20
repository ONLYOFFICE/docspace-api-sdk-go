# IAccountEntryDto

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

### NewIAccountEntryDto

`func NewIAccountEntryDto(id string, isLDAP bool, name NullableString, category string, ) *IAccountEntryDto`

NewIAccountEntryDto instantiates a new IAccountEntryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIAccountEntryDtoWithDefaults

`func NewIAccountEntryDtoWithDefaults() *IAccountEntryDto`

NewIAccountEntryDtoWithDefaults instantiates a new IAccountEntryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IAccountEntryDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IAccountEntryDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IAccountEntryDto) SetId(v string)`

SetId sets Id field to given value.


### GetDisplayName

`func (o *IAccountEntryDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IAccountEntryDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IAccountEntryDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IAccountEntryDto) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *IAccountEntryDto) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *IAccountEntryDto) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetAvatar

`func (o *IAccountEntryDto) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *IAccountEntryDto) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *IAccountEntryDto) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *IAccountEntryDto) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### SetAvatarNil

`func (o *IAccountEntryDto) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *IAccountEntryDto) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetAvatarOriginal

`func (o *IAccountEntryDto) GetAvatarOriginal() string`

GetAvatarOriginal returns the AvatarOriginal field if non-nil, zero value otherwise.

### GetAvatarOriginalOk

`func (o *IAccountEntryDto) GetAvatarOriginalOk() (*string, bool)`

GetAvatarOriginalOk returns a tuple with the AvatarOriginal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarOriginal

`func (o *IAccountEntryDto) SetAvatarOriginal(v string)`

SetAvatarOriginal sets AvatarOriginal field to given value.

### HasAvatarOriginal

`func (o *IAccountEntryDto) HasAvatarOriginal() bool`

HasAvatarOriginal returns a boolean if a field has been set.

### SetAvatarOriginalNil

`func (o *IAccountEntryDto) SetAvatarOriginalNil(b bool)`

 SetAvatarOriginalNil sets the value for AvatarOriginal to be an explicit nil

### UnsetAvatarOriginal
`func (o *IAccountEntryDto) UnsetAvatarOriginal()`

UnsetAvatarOriginal ensures that no value is present for AvatarOriginal, not even an explicit nil
### GetAvatarMax

`func (o *IAccountEntryDto) GetAvatarMax() string`

GetAvatarMax returns the AvatarMax field if non-nil, zero value otherwise.

### GetAvatarMaxOk

`func (o *IAccountEntryDto) GetAvatarMaxOk() (*string, bool)`

GetAvatarMaxOk returns a tuple with the AvatarMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarMax

`func (o *IAccountEntryDto) SetAvatarMax(v string)`

SetAvatarMax sets AvatarMax field to given value.

### HasAvatarMax

`func (o *IAccountEntryDto) HasAvatarMax() bool`

HasAvatarMax returns a boolean if a field has been set.

### SetAvatarMaxNil

`func (o *IAccountEntryDto) SetAvatarMaxNil(b bool)`

 SetAvatarMaxNil sets the value for AvatarMax to be an explicit nil

### UnsetAvatarMax
`func (o *IAccountEntryDto) UnsetAvatarMax()`

UnsetAvatarMax ensures that no value is present for AvatarMax, not even an explicit nil
### GetAvatarMedium

`func (o *IAccountEntryDto) GetAvatarMedium() string`

GetAvatarMedium returns the AvatarMedium field if non-nil, zero value otherwise.

### GetAvatarMediumOk

`func (o *IAccountEntryDto) GetAvatarMediumOk() (*string, bool)`

GetAvatarMediumOk returns a tuple with the AvatarMedium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarMedium

`func (o *IAccountEntryDto) SetAvatarMedium(v string)`

SetAvatarMedium sets AvatarMedium field to given value.

### HasAvatarMedium

`func (o *IAccountEntryDto) HasAvatarMedium() bool`

HasAvatarMedium returns a boolean if a field has been set.

### SetAvatarMediumNil

`func (o *IAccountEntryDto) SetAvatarMediumNil(b bool)`

 SetAvatarMediumNil sets the value for AvatarMedium to be an explicit nil

### UnsetAvatarMedium
`func (o *IAccountEntryDto) UnsetAvatarMedium()`

UnsetAvatarMedium ensures that no value is present for AvatarMedium, not even an explicit nil
### GetAvatarSmall

`func (o *IAccountEntryDto) GetAvatarSmall() string`

GetAvatarSmall returns the AvatarSmall field if non-nil, zero value otherwise.

### GetAvatarSmallOk

`func (o *IAccountEntryDto) GetAvatarSmallOk() (*string, bool)`

GetAvatarSmallOk returns a tuple with the AvatarSmall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarSmall

`func (o *IAccountEntryDto) SetAvatarSmall(v string)`

SetAvatarSmall sets AvatarSmall field to given value.

### HasAvatarSmall

`func (o *IAccountEntryDto) HasAvatarSmall() bool`

HasAvatarSmall returns a boolean if a field has been set.

### SetAvatarSmallNil

`func (o *IAccountEntryDto) SetAvatarSmallNil(b bool)`

 SetAvatarSmallNil sets the value for AvatarSmall to be an explicit nil

### UnsetAvatarSmall
`func (o *IAccountEntryDto) UnsetAvatarSmall()`

UnsetAvatarSmall ensures that no value is present for AvatarSmall, not even an explicit nil
### GetProfileUrl

`func (o *IAccountEntryDto) GetProfileUrl() string`

GetProfileUrl returns the ProfileUrl field if non-nil, zero value otherwise.

### GetProfileUrlOk

`func (o *IAccountEntryDto) GetProfileUrlOk() (*string, bool)`

GetProfileUrlOk returns a tuple with the ProfileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileUrl

`func (o *IAccountEntryDto) SetProfileUrl(v string)`

SetProfileUrl sets ProfileUrl field to given value.

### HasProfileUrl

`func (o *IAccountEntryDto) HasProfileUrl() bool`

HasProfileUrl returns a boolean if a field has been set.

### SetProfileUrlNil

`func (o *IAccountEntryDto) SetProfileUrlNil(b bool)`

 SetProfileUrlNil sets the value for ProfileUrl to be an explicit nil

### UnsetProfileUrl
`func (o *IAccountEntryDto) UnsetProfileUrl()`

UnsetProfileUrl ensures that no value is present for ProfileUrl, not even an explicit nil
### GetHasAvatar

`func (o *IAccountEntryDto) GetHasAvatar() bool`

GetHasAvatar returns the HasAvatar field if non-nil, zero value otherwise.

### GetHasAvatarOk

`func (o *IAccountEntryDto) GetHasAvatarOk() (*bool, bool)`

GetHasAvatarOk returns a tuple with the HasAvatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAvatar

`func (o *IAccountEntryDto) SetHasAvatar(v bool)`

SetHasAvatar sets HasAvatar field to given value.

### HasHasAvatar

`func (o *IAccountEntryDto) HasHasAvatar() bool`

HasHasAvatar returns a boolean if a field has been set.

### GetIsAnonim

`func (o *IAccountEntryDto) GetIsAnonim() bool`

GetIsAnonim returns the IsAnonim field if non-nil, zero value otherwise.

### GetIsAnonimOk

`func (o *IAccountEntryDto) GetIsAnonimOk() (*bool, bool)`

GetIsAnonimOk returns a tuple with the IsAnonim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnonim

`func (o *IAccountEntryDto) SetIsAnonim(v bool)`

SetIsAnonim sets IsAnonim field to given value.

### HasIsAnonim

`func (o *IAccountEntryDto) HasIsAnonim() bool`

HasIsAnonim returns a boolean if a field has been set.

### GetFirstName

`func (o *IAccountEntryDto) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *IAccountEntryDto) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *IAccountEntryDto) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *IAccountEntryDto) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *IAccountEntryDto) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *IAccountEntryDto) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *IAccountEntryDto) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *IAccountEntryDto) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *IAccountEntryDto) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *IAccountEntryDto) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *IAccountEntryDto) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *IAccountEntryDto) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetUserName

`func (o *IAccountEntryDto) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *IAccountEntryDto) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *IAccountEntryDto) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *IAccountEntryDto) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *IAccountEntryDto) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *IAccountEntryDto) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetEmail

`func (o *IAccountEntryDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *IAccountEntryDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *IAccountEntryDto) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *IAccountEntryDto) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *IAccountEntryDto) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *IAccountEntryDto) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetContacts

`func (o *IAccountEntryDto) GetContacts() []Contact`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *IAccountEntryDto) GetContactsOk() (*[]Contact, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *IAccountEntryDto) SetContacts(v []Contact)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *IAccountEntryDto) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### SetContactsNil

`func (o *IAccountEntryDto) SetContactsNil(b bool)`

 SetContactsNil sets the value for Contacts to be an explicit nil

### UnsetContacts
`func (o *IAccountEntryDto) UnsetContacts()`

UnsetContacts ensures that no value is present for Contacts, not even an explicit nil
### GetStatus

`func (o *IAccountEntryDto) GetStatus() EmployeeStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IAccountEntryDto) GetStatusOk() (*EmployeeStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IAccountEntryDto) SetStatus(v EmployeeStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IAccountEntryDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetActivationStatus

`func (o *IAccountEntryDto) GetActivationStatus() EmployeeActivationStatus`

GetActivationStatus returns the ActivationStatus field if non-nil, zero value otherwise.

### GetActivationStatusOk

`func (o *IAccountEntryDto) GetActivationStatusOk() (*EmployeeActivationStatus, bool)`

GetActivationStatusOk returns a tuple with the ActivationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationStatus

`func (o *IAccountEntryDto) SetActivationStatus(v EmployeeActivationStatus)`

SetActivationStatus sets ActivationStatus field to given value.

### HasActivationStatus

`func (o *IAccountEntryDto) HasActivationStatus() bool`

HasActivationStatus returns a boolean if a field has been set.

### GetTerminated

`func (o *IAccountEntryDto) GetTerminated() ApiDateTime`

GetTerminated returns the Terminated field if non-nil, zero value otherwise.

### GetTerminatedOk

`func (o *IAccountEntryDto) GetTerminatedOk() (*ApiDateTime, bool)`

GetTerminatedOk returns a tuple with the Terminated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminated

`func (o *IAccountEntryDto) SetTerminated(v ApiDateTime)`

SetTerminated sets Terminated field to given value.

### HasTerminated

`func (o *IAccountEntryDto) HasTerminated() bool`

HasTerminated returns a boolean if a field has been set.

### GetDepartment

`func (o *IAccountEntryDto) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *IAccountEntryDto) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *IAccountEntryDto) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *IAccountEntryDto) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartmentNil

`func (o *IAccountEntryDto) SetDepartmentNil(b bool)`

 SetDepartmentNil sets the value for Department to be an explicit nil

### UnsetDepartment
`func (o *IAccountEntryDto) UnsetDepartment()`

UnsetDepartment ensures that no value is present for Department, not even an explicit nil
### GetGroups

`func (o *IAccountEntryDto) GetGroups() []GroupSummaryDto`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *IAccountEntryDto) GetGroupsOk() (*[]GroupSummaryDto, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *IAccountEntryDto) SetGroups(v []GroupSummaryDto)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *IAccountEntryDto) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *IAccountEntryDto) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *IAccountEntryDto) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetLocation

`func (o *IAccountEntryDto) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *IAccountEntryDto) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *IAccountEntryDto) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *IAccountEntryDto) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *IAccountEntryDto) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *IAccountEntryDto) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetNotes

`func (o *IAccountEntryDto) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *IAccountEntryDto) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *IAccountEntryDto) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *IAccountEntryDto) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *IAccountEntryDto) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *IAccountEntryDto) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetIsAdmin

`func (o *IAccountEntryDto) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *IAccountEntryDto) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *IAccountEntryDto) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *IAccountEntryDto) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### GetIsRoomAdmin

`func (o *IAccountEntryDto) GetIsRoomAdmin() bool`

GetIsRoomAdmin returns the IsRoomAdmin field if non-nil, zero value otherwise.

### GetIsRoomAdminOk

`func (o *IAccountEntryDto) GetIsRoomAdminOk() (*bool, bool)`

GetIsRoomAdminOk returns a tuple with the IsRoomAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRoomAdmin

`func (o *IAccountEntryDto) SetIsRoomAdmin(v bool)`

SetIsRoomAdmin sets IsRoomAdmin field to given value.

### HasIsRoomAdmin

`func (o *IAccountEntryDto) HasIsRoomAdmin() bool`

HasIsRoomAdmin returns a boolean if a field has been set.

### GetIsLDAP

`func (o *IAccountEntryDto) GetIsLDAP() bool`

GetIsLDAP returns the IsLDAP field if non-nil, zero value otherwise.

### GetIsLDAPOk

`func (o *IAccountEntryDto) GetIsLDAPOk() (*bool, bool)`

GetIsLDAPOk returns a tuple with the IsLDAP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLDAP

`func (o *IAccountEntryDto) SetIsLDAP(v bool)`

SetIsLDAP sets IsLDAP field to given value.


### GetListAdminModules

`func (o *IAccountEntryDto) GetListAdminModules() []string`

GetListAdminModules returns the ListAdminModules field if non-nil, zero value otherwise.

### GetListAdminModulesOk

`func (o *IAccountEntryDto) GetListAdminModulesOk() (*[]string, bool)`

GetListAdminModulesOk returns a tuple with the ListAdminModules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListAdminModules

`func (o *IAccountEntryDto) SetListAdminModules(v []string)`

SetListAdminModules sets ListAdminModules field to given value.

### HasListAdminModules

`func (o *IAccountEntryDto) HasListAdminModules() bool`

HasListAdminModules returns a boolean if a field has been set.

### SetListAdminModulesNil

`func (o *IAccountEntryDto) SetListAdminModulesNil(b bool)`

 SetListAdminModulesNil sets the value for ListAdminModules to be an explicit nil

### UnsetListAdminModules
`func (o *IAccountEntryDto) UnsetListAdminModules()`

UnsetListAdminModules ensures that no value is present for ListAdminModules, not even an explicit nil
### GetIsOwner

`func (o *IAccountEntryDto) GetIsOwner() bool`

GetIsOwner returns the IsOwner field if non-nil, zero value otherwise.

### GetIsOwnerOk

`func (o *IAccountEntryDto) GetIsOwnerOk() (*bool, bool)`

GetIsOwnerOk returns a tuple with the IsOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOwner

`func (o *IAccountEntryDto) SetIsOwner(v bool)`

SetIsOwner sets IsOwner field to given value.

### HasIsOwner

`func (o *IAccountEntryDto) HasIsOwner() bool`

HasIsOwner returns a boolean if a field has been set.

### GetIsVisitor

`func (o *IAccountEntryDto) GetIsVisitor() bool`

GetIsVisitor returns the IsVisitor field if non-nil, zero value otherwise.

### GetIsVisitorOk

`func (o *IAccountEntryDto) GetIsVisitorOk() (*bool, bool)`

GetIsVisitorOk returns a tuple with the IsVisitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVisitor

`func (o *IAccountEntryDto) SetIsVisitor(v bool)`

SetIsVisitor sets IsVisitor field to given value.

### HasIsVisitor

`func (o *IAccountEntryDto) HasIsVisitor() bool`

HasIsVisitor returns a boolean if a field has been set.

### GetIsCollaborator

`func (o *IAccountEntryDto) GetIsCollaborator() bool`

GetIsCollaborator returns the IsCollaborator field if non-nil, zero value otherwise.

### GetIsCollaboratorOk

`func (o *IAccountEntryDto) GetIsCollaboratorOk() (*bool, bool)`

GetIsCollaboratorOk returns a tuple with the IsCollaborator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCollaborator

`func (o *IAccountEntryDto) SetIsCollaborator(v bool)`

SetIsCollaborator sets IsCollaborator field to given value.

### HasIsCollaborator

`func (o *IAccountEntryDto) HasIsCollaborator() bool`

HasIsCollaborator returns a boolean if a field has been set.

### GetCultureName

`func (o *IAccountEntryDto) GetCultureName() string`

GetCultureName returns the CultureName field if non-nil, zero value otherwise.

### GetCultureNameOk

`func (o *IAccountEntryDto) GetCultureNameOk() (*string, bool)`

GetCultureNameOk returns a tuple with the CultureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCultureName

`func (o *IAccountEntryDto) SetCultureName(v string)`

SetCultureName sets CultureName field to given value.

### HasCultureName

`func (o *IAccountEntryDto) HasCultureName() bool`

HasCultureName returns a boolean if a field has been set.

### SetCultureNameNil

`func (o *IAccountEntryDto) SetCultureNameNil(b bool)`

 SetCultureNameNil sets the value for CultureName to be an explicit nil

### UnsetCultureName
`func (o *IAccountEntryDto) UnsetCultureName()`

UnsetCultureName ensures that no value is present for CultureName, not even an explicit nil
### GetMobilePhone

`func (o *IAccountEntryDto) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *IAccountEntryDto) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *IAccountEntryDto) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *IAccountEntryDto) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### SetMobilePhoneNil

`func (o *IAccountEntryDto) SetMobilePhoneNil(b bool)`

 SetMobilePhoneNil sets the value for MobilePhone to be an explicit nil

### UnsetMobilePhone
`func (o *IAccountEntryDto) UnsetMobilePhone()`

UnsetMobilePhone ensures that no value is present for MobilePhone, not even an explicit nil
### GetMobilePhoneActivationStatus

`func (o *IAccountEntryDto) GetMobilePhoneActivationStatus() MobilePhoneActivationStatus`

GetMobilePhoneActivationStatus returns the MobilePhoneActivationStatus field if non-nil, zero value otherwise.

### GetMobilePhoneActivationStatusOk

`func (o *IAccountEntryDto) GetMobilePhoneActivationStatusOk() (*MobilePhoneActivationStatus, bool)`

GetMobilePhoneActivationStatusOk returns a tuple with the MobilePhoneActivationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhoneActivationStatus

`func (o *IAccountEntryDto) SetMobilePhoneActivationStatus(v MobilePhoneActivationStatus)`

SetMobilePhoneActivationStatus sets MobilePhoneActivationStatus field to given value.

### HasMobilePhoneActivationStatus

`func (o *IAccountEntryDto) HasMobilePhoneActivationStatus() bool`

HasMobilePhoneActivationStatus returns a boolean if a field has been set.

### GetIsSSO

`func (o *IAccountEntryDto) GetIsSSO() bool`

GetIsSSO returns the IsSSO field if non-nil, zero value otherwise.

### GetIsSSOOk

`func (o *IAccountEntryDto) GetIsSSOOk() (*bool, bool)`

GetIsSSOOk returns a tuple with the IsSSO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSSO

`func (o *IAccountEntryDto) SetIsSSO(v bool)`

SetIsSSO sets IsSSO field to given value.

### HasIsSSO

`func (o *IAccountEntryDto) HasIsSSO() bool`

HasIsSSO returns a boolean if a field has been set.

### GetTheme

`func (o *IAccountEntryDto) GetTheme() DarkThemeSettingsType`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *IAccountEntryDto) GetThemeOk() (*DarkThemeSettingsType, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *IAccountEntryDto) SetTheme(v DarkThemeSettingsType)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *IAccountEntryDto) HasTheme() bool`

HasTheme returns a boolean if a field has been set.

### GetQuotaLimit

`func (o *IAccountEntryDto) GetQuotaLimit() int64`

GetQuotaLimit returns the QuotaLimit field if non-nil, zero value otherwise.

### GetQuotaLimitOk

`func (o *IAccountEntryDto) GetQuotaLimitOk() (*int64, bool)`

GetQuotaLimitOk returns a tuple with the QuotaLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaLimit

`func (o *IAccountEntryDto) SetQuotaLimit(v int64)`

SetQuotaLimit sets QuotaLimit field to given value.

### HasQuotaLimit

`func (o *IAccountEntryDto) HasQuotaLimit() bool`

HasQuotaLimit returns a boolean if a field has been set.

### SetQuotaLimitNil

`func (o *IAccountEntryDto) SetQuotaLimitNil(b bool)`

 SetQuotaLimitNil sets the value for QuotaLimit to be an explicit nil

### UnsetQuotaLimit
`func (o *IAccountEntryDto) UnsetQuotaLimit()`

UnsetQuotaLimit ensures that no value is present for QuotaLimit, not even an explicit nil
### GetUsedSpace

`func (o *IAccountEntryDto) GetUsedSpace() float64`

GetUsedSpace returns the UsedSpace field if non-nil, zero value otherwise.

### GetUsedSpaceOk

`func (o *IAccountEntryDto) GetUsedSpaceOk() (*float64, bool)`

GetUsedSpaceOk returns a tuple with the UsedSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedSpace

`func (o *IAccountEntryDto) SetUsedSpace(v float64)`

SetUsedSpace sets UsedSpace field to given value.

### HasUsedSpace

`func (o *IAccountEntryDto) HasUsedSpace() bool`

HasUsedSpace returns a boolean if a field has been set.

### SetUsedSpaceNil

`func (o *IAccountEntryDto) SetUsedSpaceNil(b bool)`

 SetUsedSpaceNil sets the value for UsedSpace to be an explicit nil

### UnsetUsedSpace
`func (o *IAccountEntryDto) UnsetUsedSpace()`

UnsetUsedSpace ensures that no value is present for UsedSpace, not even an explicit nil
### GetShared

`func (o *IAccountEntryDto) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *IAccountEntryDto) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *IAccountEntryDto) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *IAccountEntryDto) HasShared() bool`

HasShared returns a boolean if a field has been set.

### SetSharedNil

`func (o *IAccountEntryDto) SetSharedNil(b bool)`

 SetSharedNil sets the value for Shared to be an explicit nil

### UnsetShared
`func (o *IAccountEntryDto) UnsetShared()`

UnsetShared ensures that no value is present for Shared, not even an explicit nil
### GetIsCustomQuota

`func (o *IAccountEntryDto) GetIsCustomQuota() bool`

GetIsCustomQuota returns the IsCustomQuota field if non-nil, zero value otherwise.

### GetIsCustomQuotaOk

`func (o *IAccountEntryDto) GetIsCustomQuotaOk() (*bool, bool)`

GetIsCustomQuotaOk returns a tuple with the IsCustomQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomQuota

`func (o *IAccountEntryDto) SetIsCustomQuota(v bool)`

SetIsCustomQuota sets IsCustomQuota field to given value.

### HasIsCustomQuota

`func (o *IAccountEntryDto) HasIsCustomQuota() bool`

HasIsCustomQuota returns a boolean if a field has been set.

### SetIsCustomQuotaNil

`func (o *IAccountEntryDto) SetIsCustomQuotaNil(b bool)`

 SetIsCustomQuotaNil sets the value for IsCustomQuota to be an explicit nil

### UnsetIsCustomQuota
`func (o *IAccountEntryDto) UnsetIsCustomQuota()`

UnsetIsCustomQuota ensures that no value is present for IsCustomQuota, not even an explicit nil
### GetLoginEventId

`func (o *IAccountEntryDto) GetLoginEventId() int32`

GetLoginEventId returns the LoginEventId field if non-nil, zero value otherwise.

### GetLoginEventIdOk

`func (o *IAccountEntryDto) GetLoginEventIdOk() (*int32, bool)`

GetLoginEventIdOk returns a tuple with the LoginEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginEventId

`func (o *IAccountEntryDto) SetLoginEventId(v int32)`

SetLoginEventId sets LoginEventId field to given value.

### HasLoginEventId

`func (o *IAccountEntryDto) HasLoginEventId() bool`

HasLoginEventId returns a boolean if a field has been set.

### SetLoginEventIdNil

`func (o *IAccountEntryDto) SetLoginEventIdNil(b bool)`

 SetLoginEventIdNil sets the value for LoginEventId to be an explicit nil

### UnsetLoginEventId
`func (o *IAccountEntryDto) UnsetLoginEventId()`

UnsetLoginEventId ensures that no value is present for LoginEventId, not even an explicit nil
### GetAuthCookieLifetime

`func (o *IAccountEntryDto) GetAuthCookieLifetime() float64`

GetAuthCookieLifetime returns the AuthCookieLifetime field if non-nil, zero value otherwise.

### GetAuthCookieLifetimeOk

`func (o *IAccountEntryDto) GetAuthCookieLifetimeOk() (*float64, bool)`

GetAuthCookieLifetimeOk returns a tuple with the AuthCookieLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCookieLifetime

`func (o *IAccountEntryDto) SetAuthCookieLifetime(v float64)`

SetAuthCookieLifetime sets AuthCookieLifetime field to given value.

### HasAuthCookieLifetime

`func (o *IAccountEntryDto) HasAuthCookieLifetime() bool`

HasAuthCookieLifetime returns a boolean if a field has been set.

### SetAuthCookieLifetimeNil

`func (o *IAccountEntryDto) SetAuthCookieLifetimeNil(b bool)`

 SetAuthCookieLifetimeNil sets the value for AuthCookieLifetime to be an explicit nil

### UnsetAuthCookieLifetime
`func (o *IAccountEntryDto) UnsetAuthCookieLifetime()`

UnsetAuthCookieLifetime ensures that no value is present for AuthCookieLifetime, not even an explicit nil
### GetCreatedBy

`func (o *IAccountEntryDto) GetCreatedBy() EmployeeDto`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *IAccountEntryDto) GetCreatedByOk() (*EmployeeDto, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *IAccountEntryDto) SetCreatedBy(v EmployeeDto)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *IAccountEntryDto) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetRegistrationDate

`func (o *IAccountEntryDto) GetRegistrationDate() ApiDateTime`

GetRegistrationDate returns the RegistrationDate field if non-nil, zero value otherwise.

### GetRegistrationDateOk

`func (o *IAccountEntryDto) GetRegistrationDateOk() (*ApiDateTime, bool)`

GetRegistrationDateOk returns a tuple with the RegistrationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationDate

`func (o *IAccountEntryDto) SetRegistrationDate(v ApiDateTime)`

SetRegistrationDate sets RegistrationDate field to given value.

### HasRegistrationDate

`func (o *IAccountEntryDto) HasRegistrationDate() bool`

HasRegistrationDate returns a boolean if a field has been set.

### GetHasPersonalFolder

`func (o *IAccountEntryDto) GetHasPersonalFolder() bool`

GetHasPersonalFolder returns the HasPersonalFolder field if non-nil, zero value otherwise.

### GetHasPersonalFolderOk

`func (o *IAccountEntryDto) GetHasPersonalFolderOk() (*bool, bool)`

GetHasPersonalFolderOk returns a tuple with the HasPersonalFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPersonalFolder

`func (o *IAccountEntryDto) SetHasPersonalFolder(v bool)`

SetHasPersonalFolder sets HasPersonalFolder field to given value.

### HasHasPersonalFolder

`func (o *IAccountEntryDto) HasHasPersonalFolder() bool`

HasHasPersonalFolder returns a boolean if a field has been set.

### SetHasPersonalFolderNil

`func (o *IAccountEntryDto) SetHasPersonalFolderNil(b bool)`

 SetHasPersonalFolderNil sets the value for HasPersonalFolder to be an explicit nil

### UnsetHasPersonalFolder
`func (o *IAccountEntryDto) UnsetHasPersonalFolder()`

UnsetHasPersonalFolder ensures that no value is present for HasPersonalFolder, not even an explicit nil
### GetTfaAppEnabled

`func (o *IAccountEntryDto) GetTfaAppEnabled() bool`

GetTfaAppEnabled returns the TfaAppEnabled field if non-nil, zero value otherwise.

### GetTfaAppEnabledOk

`func (o *IAccountEntryDto) GetTfaAppEnabledOk() (*bool, bool)`

GetTfaAppEnabledOk returns a tuple with the TfaAppEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTfaAppEnabled

`func (o *IAccountEntryDto) SetTfaAppEnabled(v bool)`

SetTfaAppEnabled sets TfaAppEnabled field to given value.

### HasTfaAppEnabled

`func (o *IAccountEntryDto) HasTfaAppEnabled() bool`

HasTfaAppEnabled returns a boolean if a field has been set.

### SetTfaAppEnabledNil

`func (o *IAccountEntryDto) SetTfaAppEnabledNil(b bool)`

 SetTfaAppEnabledNil sets the value for TfaAppEnabled to be an explicit nil

### UnsetTfaAppEnabled
`func (o *IAccountEntryDto) UnsetTfaAppEnabled()`

UnsetTfaAppEnabled ensures that no value is present for TfaAppEnabled, not even an explicit nil
### GetName

`func (o *IAccountEntryDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IAccountEntryDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IAccountEntryDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *IAccountEntryDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *IAccountEntryDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetParent

`func (o *IAccountEntryDto) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *IAccountEntryDto) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *IAccountEntryDto) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *IAccountEntryDto) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *IAccountEntryDto) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *IAccountEntryDto) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetCategory

`func (o *IAccountEntryDto) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *IAccountEntryDto) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *IAccountEntryDto) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetIsSystem

`func (o *IAccountEntryDto) GetIsSystem() bool`

GetIsSystem returns the IsSystem field if non-nil, zero value otherwise.

### GetIsSystemOk

`func (o *IAccountEntryDto) GetIsSystemOk() (*bool, bool)`

GetIsSystemOk returns a tuple with the IsSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystem

`func (o *IAccountEntryDto) SetIsSystem(v bool)`

SetIsSystem sets IsSystem field to given value.

### HasIsSystem

`func (o *IAccountEntryDto) HasIsSystem() bool`

HasIsSystem returns a boolean if a field has been set.

### SetIsSystemNil

`func (o *IAccountEntryDto) SetIsSystemNil(b bool)`

 SetIsSystemNil sets the value for IsSystem to be an explicit nil

### UnsetIsSystem
`func (o *IAccountEntryDto) UnsetIsSystem()`

UnsetIsSystem ensures that no value is present for IsSystem, not even an explicit nil
### GetManager

`func (o *IAccountEntryDto) GetManager() EmployeeFullDto`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *IAccountEntryDto) GetManagerOk() (*EmployeeFullDto, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *IAccountEntryDto) SetManager(v EmployeeFullDto)`

SetManager sets Manager field to given value.

### HasManager

`func (o *IAccountEntryDto) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetMembers

`func (o *IAccountEntryDto) GetMembers() []EmployeeFullDto`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *IAccountEntryDto) GetMembersOk() (*[]EmployeeFullDto, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *IAccountEntryDto) SetMembers(v []EmployeeFullDto)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *IAccountEntryDto) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### SetMembersNil

`func (o *IAccountEntryDto) SetMembersNil(b bool)`

 SetMembersNil sets the value for Members to be an explicit nil

### UnsetMembers
`func (o *IAccountEntryDto) UnsetMembers()`

UnsetMembers ensures that no value is present for Members, not even an explicit nil
### GetMembersCount

`func (o *IAccountEntryDto) GetMembersCount() int32`

GetMembersCount returns the MembersCount field if non-nil, zero value otherwise.

### GetMembersCountOk

`func (o *IAccountEntryDto) GetMembersCountOk() (*int32, bool)`

GetMembersCountOk returns a tuple with the MembersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersCount

`func (o *IAccountEntryDto) SetMembersCount(v int32)`

SetMembersCount sets MembersCount field to given value.

### HasMembersCount

`func (o *IAccountEntryDto) HasMembersCount() bool`

HasMembersCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


