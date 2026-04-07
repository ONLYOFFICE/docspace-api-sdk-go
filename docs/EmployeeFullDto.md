# EmployeeFullDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The user ID. | [optional] 
**DisplayName** | Pointer to **NullableString** | The HTML-encoded user's display name formatted according to the default format for the current culture. | [optional] 
**Title** | Pointer to **NullableString** | The user title. | [optional] 
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
**Birthday** | Pointer to [**ApiDateTime**](ApiDateTime.md) |  | [optional] 
**Sex** | Pointer to **NullableString** | The user sex. | [optional] 
**Status** | Pointer to [**EmployeeStatus**](EmployeeStatus.md) |  | [optional] 
**ActivationStatus** | Pointer to [**EmployeeActivationStatus**](EmployeeActivationStatus.md) |  | [optional] 
**Terminated** | Pointer to [**ApiDateTime**](ApiDateTime.md) |  | [optional] 
**Department** | Pointer to **NullableString** | The user department. | [optional] 
**WorkFrom** | Pointer to [**ApiDateTime**](ApiDateTime.md) |  | [optional] 
**Groups** | Pointer to [**[]GroupSummaryDto**](GroupSummaryDto.md) | The list of user groups. | [optional] 
**Location** | Pointer to **NullableString** | The user location. | [optional] 
**Notes** | Pointer to **NullableString** | The user notes. | [optional] 
**IsAdmin** | Pointer to **bool** | Specifies if the user is an administrator or not. | [optional] 
**IsRoomAdmin** | Pointer to **bool** | Specifies if the user is a room administrator or not. | [optional] 
**IsLDAP** | Pointer to **bool** | Specifies if the LDAP settings are enabled for the user or not. | [optional] 
**ListAdminModules** | Pointer to **[]string** | The list of the administrator modules. | [optional] 
**IsOwner** | Pointer to **bool** | Specifies if the user is a portal owner or not. | [optional] 
**IsVisitor** | Pointer to **bool** | Specifies if the user is a portal visitor or not. | [optional] 
**IsCollaborator** | Pointer to **bool** | Specifies if the user is a portal collaborator or not. | [optional] 
**CultureName** | Pointer to **NullableString** | The user culture code. | [optional] 
**MobilePhone** | Pointer to **NullableString** | The user mobile phone number. | [optional] 
**MobilePhoneActivationStatus** | Pointer to [**MobilePhoneActivationStatus**](MobilePhoneActivationStatus.md) |  | [optional] 
**IsSSO** | Pointer to **bool** | Specifies if the SSO settings are enabled for the user or not. | [optional] 
**Theme** | Pointer to [**DarkThemeSettingsType**](DarkThemeSettingsType.md) |  | [optional] 
**QuotaLimit** | Pointer to **NullableInt64** | The user quota limit. | [optional] 
**UsedSpace** | Pointer to **NullableFloat64** | The portal used space of the user. | [optional] 
**Shared** | Pointer to **NullableBool** | Specifies if the user has access rights. | [optional] 
**IsCustomQuota** | Pointer to **NullableBool** | Specifies if the user has a custom quota or not. | [optional] 
**LoginEventId** | Pointer to **NullableInt32** | The current login event ID. | [optional] 
**AuthCookieLifetime** | Pointer to **NullableFloat64** | The auth cookie lifetime in seconds. | [optional] 
**CreatedBy** | Pointer to [**EmployeeDto**](EmployeeDto.md) |  | [optional] 
**RegistrationDate** | Pointer to [**ApiDateTime**](ApiDateTime.md) |  | [optional] 
**HasPersonalFolder** | Pointer to **NullableBool** | Specifies if the user has a personal folder or not. | [optional] 
**TfaAppEnabled** | Pointer to **NullableBool** | Indicates whether the user has enabled two-factor authentication (TFA) using an authentication app. | [optional] 

## Methods

### NewEmployeeFullDto

`func NewEmployeeFullDto() *EmployeeFullDto`

NewEmployeeFullDto instantiates a new EmployeeFullDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployeeFullDtoWithDefaults

`func NewEmployeeFullDtoWithDefaults() *EmployeeFullDto`

NewEmployeeFullDtoWithDefaults instantiates a new EmployeeFullDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmployeeFullDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmployeeFullDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmployeeFullDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EmployeeFullDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *EmployeeFullDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EmployeeFullDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EmployeeFullDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *EmployeeFullDto) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *EmployeeFullDto) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *EmployeeFullDto) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetTitle

`func (o *EmployeeFullDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EmployeeFullDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EmployeeFullDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *EmployeeFullDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *EmployeeFullDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *EmployeeFullDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetAvatar

`func (o *EmployeeFullDto) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *EmployeeFullDto) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *EmployeeFullDto) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *EmployeeFullDto) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### SetAvatarNil

`func (o *EmployeeFullDto) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *EmployeeFullDto) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetAvatarOriginal

`func (o *EmployeeFullDto) GetAvatarOriginal() string`

GetAvatarOriginal returns the AvatarOriginal field if non-nil, zero value otherwise.

### GetAvatarOriginalOk

`func (o *EmployeeFullDto) GetAvatarOriginalOk() (*string, bool)`

GetAvatarOriginalOk returns a tuple with the AvatarOriginal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarOriginal

`func (o *EmployeeFullDto) SetAvatarOriginal(v string)`

SetAvatarOriginal sets AvatarOriginal field to given value.

### HasAvatarOriginal

`func (o *EmployeeFullDto) HasAvatarOriginal() bool`

HasAvatarOriginal returns a boolean if a field has been set.

### SetAvatarOriginalNil

`func (o *EmployeeFullDto) SetAvatarOriginalNil(b bool)`

 SetAvatarOriginalNil sets the value for AvatarOriginal to be an explicit nil

### UnsetAvatarOriginal
`func (o *EmployeeFullDto) UnsetAvatarOriginal()`

UnsetAvatarOriginal ensures that no value is present for AvatarOriginal, not even an explicit nil
### GetAvatarMax

`func (o *EmployeeFullDto) GetAvatarMax() string`

GetAvatarMax returns the AvatarMax field if non-nil, zero value otherwise.

### GetAvatarMaxOk

`func (o *EmployeeFullDto) GetAvatarMaxOk() (*string, bool)`

GetAvatarMaxOk returns a tuple with the AvatarMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarMax

`func (o *EmployeeFullDto) SetAvatarMax(v string)`

SetAvatarMax sets AvatarMax field to given value.

### HasAvatarMax

`func (o *EmployeeFullDto) HasAvatarMax() bool`

HasAvatarMax returns a boolean if a field has been set.

### SetAvatarMaxNil

`func (o *EmployeeFullDto) SetAvatarMaxNil(b bool)`

 SetAvatarMaxNil sets the value for AvatarMax to be an explicit nil

### UnsetAvatarMax
`func (o *EmployeeFullDto) UnsetAvatarMax()`

UnsetAvatarMax ensures that no value is present for AvatarMax, not even an explicit nil
### GetAvatarMedium

`func (o *EmployeeFullDto) GetAvatarMedium() string`

GetAvatarMedium returns the AvatarMedium field if non-nil, zero value otherwise.

### GetAvatarMediumOk

`func (o *EmployeeFullDto) GetAvatarMediumOk() (*string, bool)`

GetAvatarMediumOk returns a tuple with the AvatarMedium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarMedium

`func (o *EmployeeFullDto) SetAvatarMedium(v string)`

SetAvatarMedium sets AvatarMedium field to given value.

### HasAvatarMedium

`func (o *EmployeeFullDto) HasAvatarMedium() bool`

HasAvatarMedium returns a boolean if a field has been set.

### SetAvatarMediumNil

`func (o *EmployeeFullDto) SetAvatarMediumNil(b bool)`

 SetAvatarMediumNil sets the value for AvatarMedium to be an explicit nil

### UnsetAvatarMedium
`func (o *EmployeeFullDto) UnsetAvatarMedium()`

UnsetAvatarMedium ensures that no value is present for AvatarMedium, not even an explicit nil
### GetAvatarSmall

`func (o *EmployeeFullDto) GetAvatarSmall() string`

GetAvatarSmall returns the AvatarSmall field if non-nil, zero value otherwise.

### GetAvatarSmallOk

`func (o *EmployeeFullDto) GetAvatarSmallOk() (*string, bool)`

GetAvatarSmallOk returns a tuple with the AvatarSmall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarSmall

`func (o *EmployeeFullDto) SetAvatarSmall(v string)`

SetAvatarSmall sets AvatarSmall field to given value.

### HasAvatarSmall

`func (o *EmployeeFullDto) HasAvatarSmall() bool`

HasAvatarSmall returns a boolean if a field has been set.

### SetAvatarSmallNil

`func (o *EmployeeFullDto) SetAvatarSmallNil(b bool)`

 SetAvatarSmallNil sets the value for AvatarSmall to be an explicit nil

### UnsetAvatarSmall
`func (o *EmployeeFullDto) UnsetAvatarSmall()`

UnsetAvatarSmall ensures that no value is present for AvatarSmall, not even an explicit nil
### GetProfileUrl

`func (o *EmployeeFullDto) GetProfileUrl() string`

GetProfileUrl returns the ProfileUrl field if non-nil, zero value otherwise.

### GetProfileUrlOk

`func (o *EmployeeFullDto) GetProfileUrlOk() (*string, bool)`

GetProfileUrlOk returns a tuple with the ProfileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileUrl

`func (o *EmployeeFullDto) SetProfileUrl(v string)`

SetProfileUrl sets ProfileUrl field to given value.

### HasProfileUrl

`func (o *EmployeeFullDto) HasProfileUrl() bool`

HasProfileUrl returns a boolean if a field has been set.

### SetProfileUrlNil

`func (o *EmployeeFullDto) SetProfileUrlNil(b bool)`

 SetProfileUrlNil sets the value for ProfileUrl to be an explicit nil

### UnsetProfileUrl
`func (o *EmployeeFullDto) UnsetProfileUrl()`

UnsetProfileUrl ensures that no value is present for ProfileUrl, not even an explicit nil
### GetHasAvatar

`func (o *EmployeeFullDto) GetHasAvatar() bool`

GetHasAvatar returns the HasAvatar field if non-nil, zero value otherwise.

### GetHasAvatarOk

`func (o *EmployeeFullDto) GetHasAvatarOk() (*bool, bool)`

GetHasAvatarOk returns a tuple with the HasAvatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAvatar

`func (o *EmployeeFullDto) SetHasAvatar(v bool)`

SetHasAvatar sets HasAvatar field to given value.

### HasHasAvatar

`func (o *EmployeeFullDto) HasHasAvatar() bool`

HasHasAvatar returns a boolean if a field has been set.

### GetIsAnonim

`func (o *EmployeeFullDto) GetIsAnonim() bool`

GetIsAnonim returns the IsAnonim field if non-nil, zero value otherwise.

### GetIsAnonimOk

`func (o *EmployeeFullDto) GetIsAnonimOk() (*bool, bool)`

GetIsAnonimOk returns a tuple with the IsAnonim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnonim

`func (o *EmployeeFullDto) SetIsAnonim(v bool)`

SetIsAnonim sets IsAnonim field to given value.

### HasIsAnonim

`func (o *EmployeeFullDto) HasIsAnonim() bool`

HasIsAnonim returns a boolean if a field has been set.

### GetFirstName

`func (o *EmployeeFullDto) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *EmployeeFullDto) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *EmployeeFullDto) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *EmployeeFullDto) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *EmployeeFullDto) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *EmployeeFullDto) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *EmployeeFullDto) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *EmployeeFullDto) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *EmployeeFullDto) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *EmployeeFullDto) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *EmployeeFullDto) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *EmployeeFullDto) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetUserName

`func (o *EmployeeFullDto) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *EmployeeFullDto) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *EmployeeFullDto) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *EmployeeFullDto) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *EmployeeFullDto) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *EmployeeFullDto) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetEmail

`func (o *EmployeeFullDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EmployeeFullDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EmployeeFullDto) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *EmployeeFullDto) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *EmployeeFullDto) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *EmployeeFullDto) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetContacts

`func (o *EmployeeFullDto) GetContacts() []Contact`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *EmployeeFullDto) GetContactsOk() (*[]Contact, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *EmployeeFullDto) SetContacts(v []Contact)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *EmployeeFullDto) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### SetContactsNil

`func (o *EmployeeFullDto) SetContactsNil(b bool)`

 SetContactsNil sets the value for Contacts to be an explicit nil

### UnsetContacts
`func (o *EmployeeFullDto) UnsetContacts()`

UnsetContacts ensures that no value is present for Contacts, not even an explicit nil
### GetBirthday

`func (o *EmployeeFullDto) GetBirthday() ApiDateTime`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *EmployeeFullDto) GetBirthdayOk() (*ApiDateTime, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *EmployeeFullDto) SetBirthday(v ApiDateTime)`

SetBirthday sets Birthday field to given value.

### HasBirthday

`func (o *EmployeeFullDto) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### GetSex

`func (o *EmployeeFullDto) GetSex() string`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *EmployeeFullDto) GetSexOk() (*string, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *EmployeeFullDto) SetSex(v string)`

SetSex sets Sex field to given value.

### HasSex

`func (o *EmployeeFullDto) HasSex() bool`

HasSex returns a boolean if a field has been set.

### SetSexNil

`func (o *EmployeeFullDto) SetSexNil(b bool)`

 SetSexNil sets the value for Sex to be an explicit nil

### UnsetSex
`func (o *EmployeeFullDto) UnsetSex()`

UnsetSex ensures that no value is present for Sex, not even an explicit nil
### GetStatus

`func (o *EmployeeFullDto) GetStatus() EmployeeStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EmployeeFullDto) GetStatusOk() (*EmployeeStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EmployeeFullDto) SetStatus(v EmployeeStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EmployeeFullDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetActivationStatus

`func (o *EmployeeFullDto) GetActivationStatus() EmployeeActivationStatus`

GetActivationStatus returns the ActivationStatus field if non-nil, zero value otherwise.

### GetActivationStatusOk

`func (o *EmployeeFullDto) GetActivationStatusOk() (*EmployeeActivationStatus, bool)`

GetActivationStatusOk returns a tuple with the ActivationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationStatus

`func (o *EmployeeFullDto) SetActivationStatus(v EmployeeActivationStatus)`

SetActivationStatus sets ActivationStatus field to given value.

### HasActivationStatus

`func (o *EmployeeFullDto) HasActivationStatus() bool`

HasActivationStatus returns a boolean if a field has been set.

### GetTerminated

`func (o *EmployeeFullDto) GetTerminated() ApiDateTime`

GetTerminated returns the Terminated field if non-nil, zero value otherwise.

### GetTerminatedOk

`func (o *EmployeeFullDto) GetTerminatedOk() (*ApiDateTime, bool)`

GetTerminatedOk returns a tuple with the Terminated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminated

`func (o *EmployeeFullDto) SetTerminated(v ApiDateTime)`

SetTerminated sets Terminated field to given value.

### HasTerminated

`func (o *EmployeeFullDto) HasTerminated() bool`

HasTerminated returns a boolean if a field has been set.

### GetDepartment

`func (o *EmployeeFullDto) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *EmployeeFullDto) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *EmployeeFullDto) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *EmployeeFullDto) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartmentNil

`func (o *EmployeeFullDto) SetDepartmentNil(b bool)`

 SetDepartmentNil sets the value for Department to be an explicit nil

### UnsetDepartment
`func (o *EmployeeFullDto) UnsetDepartment()`

UnsetDepartment ensures that no value is present for Department, not even an explicit nil
### GetWorkFrom

`func (o *EmployeeFullDto) GetWorkFrom() ApiDateTime`

GetWorkFrom returns the WorkFrom field if non-nil, zero value otherwise.

### GetWorkFromOk

`func (o *EmployeeFullDto) GetWorkFromOk() (*ApiDateTime, bool)`

GetWorkFromOk returns a tuple with the WorkFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkFrom

`func (o *EmployeeFullDto) SetWorkFrom(v ApiDateTime)`

SetWorkFrom sets WorkFrom field to given value.

### HasWorkFrom

`func (o *EmployeeFullDto) HasWorkFrom() bool`

HasWorkFrom returns a boolean if a field has been set.

### GetGroups

`func (o *EmployeeFullDto) GetGroups() []GroupSummaryDto`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *EmployeeFullDto) GetGroupsOk() (*[]GroupSummaryDto, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *EmployeeFullDto) SetGroups(v []GroupSummaryDto)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *EmployeeFullDto) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *EmployeeFullDto) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *EmployeeFullDto) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetLocation

`func (o *EmployeeFullDto) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *EmployeeFullDto) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *EmployeeFullDto) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *EmployeeFullDto) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *EmployeeFullDto) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *EmployeeFullDto) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetNotes

`func (o *EmployeeFullDto) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *EmployeeFullDto) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *EmployeeFullDto) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *EmployeeFullDto) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *EmployeeFullDto) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *EmployeeFullDto) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetIsAdmin

`func (o *EmployeeFullDto) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *EmployeeFullDto) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *EmployeeFullDto) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *EmployeeFullDto) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### GetIsRoomAdmin

`func (o *EmployeeFullDto) GetIsRoomAdmin() bool`

GetIsRoomAdmin returns the IsRoomAdmin field if non-nil, zero value otherwise.

### GetIsRoomAdminOk

`func (o *EmployeeFullDto) GetIsRoomAdminOk() (*bool, bool)`

GetIsRoomAdminOk returns a tuple with the IsRoomAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRoomAdmin

`func (o *EmployeeFullDto) SetIsRoomAdmin(v bool)`

SetIsRoomAdmin sets IsRoomAdmin field to given value.

### HasIsRoomAdmin

`func (o *EmployeeFullDto) HasIsRoomAdmin() bool`

HasIsRoomAdmin returns a boolean if a field has been set.

### GetIsLDAP

`func (o *EmployeeFullDto) GetIsLDAP() bool`

GetIsLDAP returns the IsLDAP field if non-nil, zero value otherwise.

### GetIsLDAPOk

`func (o *EmployeeFullDto) GetIsLDAPOk() (*bool, bool)`

GetIsLDAPOk returns a tuple with the IsLDAP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLDAP

`func (o *EmployeeFullDto) SetIsLDAP(v bool)`

SetIsLDAP sets IsLDAP field to given value.

### HasIsLDAP

`func (o *EmployeeFullDto) HasIsLDAP() bool`

HasIsLDAP returns a boolean if a field has been set.

### GetListAdminModules

`func (o *EmployeeFullDto) GetListAdminModules() []string`

GetListAdminModules returns the ListAdminModules field if non-nil, zero value otherwise.

### GetListAdminModulesOk

`func (o *EmployeeFullDto) GetListAdminModulesOk() (*[]string, bool)`

GetListAdminModulesOk returns a tuple with the ListAdminModules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListAdminModules

`func (o *EmployeeFullDto) SetListAdminModules(v []string)`

SetListAdminModules sets ListAdminModules field to given value.

### HasListAdminModules

`func (o *EmployeeFullDto) HasListAdminModules() bool`

HasListAdminModules returns a boolean if a field has been set.

### SetListAdminModulesNil

`func (o *EmployeeFullDto) SetListAdminModulesNil(b bool)`

 SetListAdminModulesNil sets the value for ListAdminModules to be an explicit nil

### UnsetListAdminModules
`func (o *EmployeeFullDto) UnsetListAdminModules()`

UnsetListAdminModules ensures that no value is present for ListAdminModules, not even an explicit nil
### GetIsOwner

`func (o *EmployeeFullDto) GetIsOwner() bool`

GetIsOwner returns the IsOwner field if non-nil, zero value otherwise.

### GetIsOwnerOk

`func (o *EmployeeFullDto) GetIsOwnerOk() (*bool, bool)`

GetIsOwnerOk returns a tuple with the IsOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOwner

`func (o *EmployeeFullDto) SetIsOwner(v bool)`

SetIsOwner sets IsOwner field to given value.

### HasIsOwner

`func (o *EmployeeFullDto) HasIsOwner() bool`

HasIsOwner returns a boolean if a field has been set.

### GetIsVisitor

`func (o *EmployeeFullDto) GetIsVisitor() bool`

GetIsVisitor returns the IsVisitor field if non-nil, zero value otherwise.

### GetIsVisitorOk

`func (o *EmployeeFullDto) GetIsVisitorOk() (*bool, bool)`

GetIsVisitorOk returns a tuple with the IsVisitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVisitor

`func (o *EmployeeFullDto) SetIsVisitor(v bool)`

SetIsVisitor sets IsVisitor field to given value.

### HasIsVisitor

`func (o *EmployeeFullDto) HasIsVisitor() bool`

HasIsVisitor returns a boolean if a field has been set.

### GetIsCollaborator

`func (o *EmployeeFullDto) GetIsCollaborator() bool`

GetIsCollaborator returns the IsCollaborator field if non-nil, zero value otherwise.

### GetIsCollaboratorOk

`func (o *EmployeeFullDto) GetIsCollaboratorOk() (*bool, bool)`

GetIsCollaboratorOk returns a tuple with the IsCollaborator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCollaborator

`func (o *EmployeeFullDto) SetIsCollaborator(v bool)`

SetIsCollaborator sets IsCollaborator field to given value.

### HasIsCollaborator

`func (o *EmployeeFullDto) HasIsCollaborator() bool`

HasIsCollaborator returns a boolean if a field has been set.

### GetCultureName

`func (o *EmployeeFullDto) GetCultureName() string`

GetCultureName returns the CultureName field if non-nil, zero value otherwise.

### GetCultureNameOk

`func (o *EmployeeFullDto) GetCultureNameOk() (*string, bool)`

GetCultureNameOk returns a tuple with the CultureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCultureName

`func (o *EmployeeFullDto) SetCultureName(v string)`

SetCultureName sets CultureName field to given value.

### HasCultureName

`func (o *EmployeeFullDto) HasCultureName() bool`

HasCultureName returns a boolean if a field has been set.

### SetCultureNameNil

`func (o *EmployeeFullDto) SetCultureNameNil(b bool)`

 SetCultureNameNil sets the value for CultureName to be an explicit nil

### UnsetCultureName
`func (o *EmployeeFullDto) UnsetCultureName()`

UnsetCultureName ensures that no value is present for CultureName, not even an explicit nil
### GetMobilePhone

`func (o *EmployeeFullDto) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *EmployeeFullDto) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *EmployeeFullDto) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *EmployeeFullDto) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### SetMobilePhoneNil

`func (o *EmployeeFullDto) SetMobilePhoneNil(b bool)`

 SetMobilePhoneNil sets the value for MobilePhone to be an explicit nil

### UnsetMobilePhone
`func (o *EmployeeFullDto) UnsetMobilePhone()`

UnsetMobilePhone ensures that no value is present for MobilePhone, not even an explicit nil
### GetMobilePhoneActivationStatus

`func (o *EmployeeFullDto) GetMobilePhoneActivationStatus() MobilePhoneActivationStatus`

GetMobilePhoneActivationStatus returns the MobilePhoneActivationStatus field if non-nil, zero value otherwise.

### GetMobilePhoneActivationStatusOk

`func (o *EmployeeFullDto) GetMobilePhoneActivationStatusOk() (*MobilePhoneActivationStatus, bool)`

GetMobilePhoneActivationStatusOk returns a tuple with the MobilePhoneActivationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhoneActivationStatus

`func (o *EmployeeFullDto) SetMobilePhoneActivationStatus(v MobilePhoneActivationStatus)`

SetMobilePhoneActivationStatus sets MobilePhoneActivationStatus field to given value.

### HasMobilePhoneActivationStatus

`func (o *EmployeeFullDto) HasMobilePhoneActivationStatus() bool`

HasMobilePhoneActivationStatus returns a boolean if a field has been set.

### GetIsSSO

`func (o *EmployeeFullDto) GetIsSSO() bool`

GetIsSSO returns the IsSSO field if non-nil, zero value otherwise.

### GetIsSSOOk

`func (o *EmployeeFullDto) GetIsSSOOk() (*bool, bool)`

GetIsSSOOk returns a tuple with the IsSSO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSSO

`func (o *EmployeeFullDto) SetIsSSO(v bool)`

SetIsSSO sets IsSSO field to given value.

### HasIsSSO

`func (o *EmployeeFullDto) HasIsSSO() bool`

HasIsSSO returns a boolean if a field has been set.

### GetTheme

`func (o *EmployeeFullDto) GetTheme() DarkThemeSettingsType`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *EmployeeFullDto) GetThemeOk() (*DarkThemeSettingsType, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *EmployeeFullDto) SetTheme(v DarkThemeSettingsType)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *EmployeeFullDto) HasTheme() bool`

HasTheme returns a boolean if a field has been set.

### GetQuotaLimit

`func (o *EmployeeFullDto) GetQuotaLimit() int64`

GetQuotaLimit returns the QuotaLimit field if non-nil, zero value otherwise.

### GetQuotaLimitOk

`func (o *EmployeeFullDto) GetQuotaLimitOk() (*int64, bool)`

GetQuotaLimitOk returns a tuple with the QuotaLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaLimit

`func (o *EmployeeFullDto) SetQuotaLimit(v int64)`

SetQuotaLimit sets QuotaLimit field to given value.

### HasQuotaLimit

`func (o *EmployeeFullDto) HasQuotaLimit() bool`

HasQuotaLimit returns a boolean if a field has been set.

### SetQuotaLimitNil

`func (o *EmployeeFullDto) SetQuotaLimitNil(b bool)`

 SetQuotaLimitNil sets the value for QuotaLimit to be an explicit nil

### UnsetQuotaLimit
`func (o *EmployeeFullDto) UnsetQuotaLimit()`

UnsetQuotaLimit ensures that no value is present for QuotaLimit, not even an explicit nil
### GetUsedSpace

`func (o *EmployeeFullDto) GetUsedSpace() float64`

GetUsedSpace returns the UsedSpace field if non-nil, zero value otherwise.

### GetUsedSpaceOk

`func (o *EmployeeFullDto) GetUsedSpaceOk() (*float64, bool)`

GetUsedSpaceOk returns a tuple with the UsedSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedSpace

`func (o *EmployeeFullDto) SetUsedSpace(v float64)`

SetUsedSpace sets UsedSpace field to given value.

### HasUsedSpace

`func (o *EmployeeFullDto) HasUsedSpace() bool`

HasUsedSpace returns a boolean if a field has been set.

### SetUsedSpaceNil

`func (o *EmployeeFullDto) SetUsedSpaceNil(b bool)`

 SetUsedSpaceNil sets the value for UsedSpace to be an explicit nil

### UnsetUsedSpace
`func (o *EmployeeFullDto) UnsetUsedSpace()`

UnsetUsedSpace ensures that no value is present for UsedSpace, not even an explicit nil
### GetShared

`func (o *EmployeeFullDto) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *EmployeeFullDto) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *EmployeeFullDto) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *EmployeeFullDto) HasShared() bool`

HasShared returns a boolean if a field has been set.

### SetSharedNil

`func (o *EmployeeFullDto) SetSharedNil(b bool)`

 SetSharedNil sets the value for Shared to be an explicit nil

### UnsetShared
`func (o *EmployeeFullDto) UnsetShared()`

UnsetShared ensures that no value is present for Shared, not even an explicit nil
### GetIsCustomQuota

`func (o *EmployeeFullDto) GetIsCustomQuota() bool`

GetIsCustomQuota returns the IsCustomQuota field if non-nil, zero value otherwise.

### GetIsCustomQuotaOk

`func (o *EmployeeFullDto) GetIsCustomQuotaOk() (*bool, bool)`

GetIsCustomQuotaOk returns a tuple with the IsCustomQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomQuota

`func (o *EmployeeFullDto) SetIsCustomQuota(v bool)`

SetIsCustomQuota sets IsCustomQuota field to given value.

### HasIsCustomQuota

`func (o *EmployeeFullDto) HasIsCustomQuota() bool`

HasIsCustomQuota returns a boolean if a field has been set.

### SetIsCustomQuotaNil

`func (o *EmployeeFullDto) SetIsCustomQuotaNil(b bool)`

 SetIsCustomQuotaNil sets the value for IsCustomQuota to be an explicit nil

### UnsetIsCustomQuota
`func (o *EmployeeFullDto) UnsetIsCustomQuota()`

UnsetIsCustomQuota ensures that no value is present for IsCustomQuota, not even an explicit nil
### GetLoginEventId

`func (o *EmployeeFullDto) GetLoginEventId() int32`

GetLoginEventId returns the LoginEventId field if non-nil, zero value otherwise.

### GetLoginEventIdOk

`func (o *EmployeeFullDto) GetLoginEventIdOk() (*int32, bool)`

GetLoginEventIdOk returns a tuple with the LoginEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginEventId

`func (o *EmployeeFullDto) SetLoginEventId(v int32)`

SetLoginEventId sets LoginEventId field to given value.

### HasLoginEventId

`func (o *EmployeeFullDto) HasLoginEventId() bool`

HasLoginEventId returns a boolean if a field has been set.

### SetLoginEventIdNil

`func (o *EmployeeFullDto) SetLoginEventIdNil(b bool)`

 SetLoginEventIdNil sets the value for LoginEventId to be an explicit nil

### UnsetLoginEventId
`func (o *EmployeeFullDto) UnsetLoginEventId()`

UnsetLoginEventId ensures that no value is present for LoginEventId, not even an explicit nil
### GetAuthCookieLifetime

`func (o *EmployeeFullDto) GetAuthCookieLifetime() float64`

GetAuthCookieLifetime returns the AuthCookieLifetime field if non-nil, zero value otherwise.

### GetAuthCookieLifetimeOk

`func (o *EmployeeFullDto) GetAuthCookieLifetimeOk() (*float64, bool)`

GetAuthCookieLifetimeOk returns a tuple with the AuthCookieLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCookieLifetime

`func (o *EmployeeFullDto) SetAuthCookieLifetime(v float64)`

SetAuthCookieLifetime sets AuthCookieLifetime field to given value.

### HasAuthCookieLifetime

`func (o *EmployeeFullDto) HasAuthCookieLifetime() bool`

HasAuthCookieLifetime returns a boolean if a field has been set.

### SetAuthCookieLifetimeNil

`func (o *EmployeeFullDto) SetAuthCookieLifetimeNil(b bool)`

 SetAuthCookieLifetimeNil sets the value for AuthCookieLifetime to be an explicit nil

### UnsetAuthCookieLifetime
`func (o *EmployeeFullDto) UnsetAuthCookieLifetime()`

UnsetAuthCookieLifetime ensures that no value is present for AuthCookieLifetime, not even an explicit nil
### GetCreatedBy

`func (o *EmployeeFullDto) GetCreatedBy() EmployeeDto`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *EmployeeFullDto) GetCreatedByOk() (*EmployeeDto, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *EmployeeFullDto) SetCreatedBy(v EmployeeDto)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *EmployeeFullDto) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetRegistrationDate

`func (o *EmployeeFullDto) GetRegistrationDate() ApiDateTime`

GetRegistrationDate returns the RegistrationDate field if non-nil, zero value otherwise.

### GetRegistrationDateOk

`func (o *EmployeeFullDto) GetRegistrationDateOk() (*ApiDateTime, bool)`

GetRegistrationDateOk returns a tuple with the RegistrationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationDate

`func (o *EmployeeFullDto) SetRegistrationDate(v ApiDateTime)`

SetRegistrationDate sets RegistrationDate field to given value.

### HasRegistrationDate

`func (o *EmployeeFullDto) HasRegistrationDate() bool`

HasRegistrationDate returns a boolean if a field has been set.

### GetHasPersonalFolder

`func (o *EmployeeFullDto) GetHasPersonalFolder() bool`

GetHasPersonalFolder returns the HasPersonalFolder field if non-nil, zero value otherwise.

### GetHasPersonalFolderOk

`func (o *EmployeeFullDto) GetHasPersonalFolderOk() (*bool, bool)`

GetHasPersonalFolderOk returns a tuple with the HasPersonalFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPersonalFolder

`func (o *EmployeeFullDto) SetHasPersonalFolder(v bool)`

SetHasPersonalFolder sets HasPersonalFolder field to given value.

### HasHasPersonalFolder

`func (o *EmployeeFullDto) HasHasPersonalFolder() bool`

HasHasPersonalFolder returns a boolean if a field has been set.

### SetHasPersonalFolderNil

`func (o *EmployeeFullDto) SetHasPersonalFolderNil(b bool)`

 SetHasPersonalFolderNil sets the value for HasPersonalFolder to be an explicit nil

### UnsetHasPersonalFolder
`func (o *EmployeeFullDto) UnsetHasPersonalFolder()`

UnsetHasPersonalFolder ensures that no value is present for HasPersonalFolder, not even an explicit nil
### GetTfaAppEnabled

`func (o *EmployeeFullDto) GetTfaAppEnabled() bool`

GetTfaAppEnabled returns the TfaAppEnabled field if non-nil, zero value otherwise.

### GetTfaAppEnabledOk

`func (o *EmployeeFullDto) GetTfaAppEnabledOk() (*bool, bool)`

GetTfaAppEnabledOk returns a tuple with the TfaAppEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTfaAppEnabled

`func (o *EmployeeFullDto) SetTfaAppEnabled(v bool)`

SetTfaAppEnabled sets TfaAppEnabled field to given value.

### HasTfaAppEnabled

`func (o *EmployeeFullDto) HasTfaAppEnabled() bool`

HasTfaAppEnabled returns a boolean if a field has been set.

### SetTfaAppEnabledNil

`func (o *EmployeeFullDto) SetTfaAppEnabledNil(b bool)`

 SetTfaAppEnabledNil sets the value for TfaAppEnabled to be an explicit nil

### UnsetTfaAppEnabled
`func (o *EmployeeFullDto) UnsetTfaAppEnabled()`

UnsetTfaAppEnabled ensures that no value is present for TfaAppEnabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


