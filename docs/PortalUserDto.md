# PortalUserDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The user ID. | [optional] 
**FirstName** | Pointer to **NullableString** | The user's first name. | [optional] 
**LastName** | Pointer to **NullableString** | The user's last name. | [optional] 
**UserName** | Pointer to **NullableString** | The user username. | [optional] 
**BirthDate** | Pointer to **NullableTime** | The user birthday. | [optional] 
**Sex** | Pointer to **NullableBool** | The user sex (male or female). | [optional] 
**Status** | Pointer to [**EmployeeStatus**](EmployeeStatus.md) | The user status. | [optional] 
**ActivationStatus** | Pointer to [**EmployeeActivationStatus**](EmployeeActivationStatus.md) | The user activation status. | [optional] 
**TerminatedDate** | Pointer to **NullableTime** | The date and time when the user account was terminated. | [optional] 
**Title** | Pointer to **NullableString** | The user title. | [optional] 
**WorkFromDate** | Pointer to **NullableTime** | The user registration date. | [optional] 
**Email** | Pointer to **NullableString** | The user email address. | [optional] 
**Contacts** | Pointer to **NullableString** | The list of user contacts in the string format. | [optional] 
**ContactsList** | Pointer to **[]string** | The list of user contacts. | [optional] 
**Location** | Pointer to **NullableString** | The user location. | [optional] 
**Notes** | Pointer to **NullableString** | The user notes. | [optional] 
**Removed** | Pointer to **bool** | Specifies if the user account was removed or not. | [optional] 
**LastModified** | Pointer to **time.Time** | The date and time when the user account was last modified. | [optional] 
**TenantId** | Pointer to **int32** | The tenant ID. | [optional] 
**IsActive** | Pointer to **bool** | Specifies if the user is active or not. | [optional] 
**CultureName** | Pointer to **NullableString** | The user culture code. | [optional] 
**MobilePhone** | Pointer to **NullableString** | The user mobile phone. | [optional] 
**MobilePhoneActivationStatus** | Pointer to [**MobilePhoneActivationStatus**](MobilePhoneActivationStatus.md) | The user mobile phone activation status. | [optional] 
**Sid** | Pointer to **NullableString** | The LDAP user identifier. | [optional] 
**LdapQouta** | Pointer to **int64** | The LDAP user quota attribute. | [optional] 
**SsoNameId** | Pointer to **NullableString** | The SSO SAML user identifier. | [optional] 
**SsoSessionId** | Pointer to **NullableString** | The SSO SAML user session identifier. | [optional] 
**CreateDate** | Pointer to **time.Time** | The date and time when the user account was created. | [optional] 
**CreatedBy** | Pointer to **NullableString** | The ID of the user who created the current user account. | [optional] 
**Spam** | Pointer to **NullableBool** | Specifies if tips, updates and offers are allowed to be sent to the user or not. | [optional] 
**CheckActivation** | Pointer to **bool** | Specifies whether the account still has to be activated: the opposite of `isActive`. | [optional] 

## Methods

### NewPortalUserDto

`func NewPortalUserDto() *PortalUserDto`

NewPortalUserDto instantiates a new PortalUserDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortalUserDtoWithDefaults

`func NewPortalUserDtoWithDefaults() *PortalUserDto`

NewPortalUserDtoWithDefaults instantiates a new PortalUserDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortalUserDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortalUserDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortalUserDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PortalUserDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFirstName

`func (o *PortalUserDto) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *PortalUserDto) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *PortalUserDto) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *PortalUserDto) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *PortalUserDto) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *PortalUserDto) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *PortalUserDto) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *PortalUserDto) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *PortalUserDto) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *PortalUserDto) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *PortalUserDto) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *PortalUserDto) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetUserName

`func (o *PortalUserDto) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *PortalUserDto) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *PortalUserDto) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *PortalUserDto) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *PortalUserDto) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *PortalUserDto) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetBirthDate

`func (o *PortalUserDto) GetBirthDate() time.Time`

GetBirthDate returns the BirthDate field if non-nil, zero value otherwise.

### GetBirthDateOk

`func (o *PortalUserDto) GetBirthDateOk() (*time.Time, bool)`

GetBirthDateOk returns a tuple with the BirthDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthDate

`func (o *PortalUserDto) SetBirthDate(v time.Time)`

SetBirthDate sets BirthDate field to given value.

### HasBirthDate

`func (o *PortalUserDto) HasBirthDate() bool`

HasBirthDate returns a boolean if a field has been set.

### SetBirthDateNil

`func (o *PortalUserDto) SetBirthDateNil(b bool)`

 SetBirthDateNil sets the value for BirthDate to be an explicit nil

### UnsetBirthDate
`func (o *PortalUserDto) UnsetBirthDate()`

UnsetBirthDate ensures that no value is present for BirthDate, not even an explicit nil
### GetSex

`func (o *PortalUserDto) GetSex() bool`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *PortalUserDto) GetSexOk() (*bool, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *PortalUserDto) SetSex(v bool)`

SetSex sets Sex field to given value.

### HasSex

`func (o *PortalUserDto) HasSex() bool`

HasSex returns a boolean if a field has been set.

### SetSexNil

`func (o *PortalUserDto) SetSexNil(b bool)`

 SetSexNil sets the value for Sex to be an explicit nil

### UnsetSex
`func (o *PortalUserDto) UnsetSex()`

UnsetSex ensures that no value is present for Sex, not even an explicit nil
### GetStatus

`func (o *PortalUserDto) GetStatus() EmployeeStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PortalUserDto) GetStatusOk() (*EmployeeStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PortalUserDto) SetStatus(v EmployeeStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PortalUserDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetActivationStatus

`func (o *PortalUserDto) GetActivationStatus() EmployeeActivationStatus`

GetActivationStatus returns the ActivationStatus field if non-nil, zero value otherwise.

### GetActivationStatusOk

`func (o *PortalUserDto) GetActivationStatusOk() (*EmployeeActivationStatus, bool)`

GetActivationStatusOk returns a tuple with the ActivationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationStatus

`func (o *PortalUserDto) SetActivationStatus(v EmployeeActivationStatus)`

SetActivationStatus sets ActivationStatus field to given value.

### HasActivationStatus

`func (o *PortalUserDto) HasActivationStatus() bool`

HasActivationStatus returns a boolean if a field has been set.

### GetTerminatedDate

`func (o *PortalUserDto) GetTerminatedDate() time.Time`

GetTerminatedDate returns the TerminatedDate field if non-nil, zero value otherwise.

### GetTerminatedDateOk

`func (o *PortalUserDto) GetTerminatedDateOk() (*time.Time, bool)`

GetTerminatedDateOk returns a tuple with the TerminatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminatedDate

`func (o *PortalUserDto) SetTerminatedDate(v time.Time)`

SetTerminatedDate sets TerminatedDate field to given value.

### HasTerminatedDate

`func (o *PortalUserDto) HasTerminatedDate() bool`

HasTerminatedDate returns a boolean if a field has been set.

### SetTerminatedDateNil

`func (o *PortalUserDto) SetTerminatedDateNil(b bool)`

 SetTerminatedDateNil sets the value for TerminatedDate to be an explicit nil

### UnsetTerminatedDate
`func (o *PortalUserDto) UnsetTerminatedDate()`

UnsetTerminatedDate ensures that no value is present for TerminatedDate, not even an explicit nil
### GetTitle

`func (o *PortalUserDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PortalUserDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PortalUserDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PortalUserDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *PortalUserDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *PortalUserDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetWorkFromDate

`func (o *PortalUserDto) GetWorkFromDate() time.Time`

GetWorkFromDate returns the WorkFromDate field if non-nil, zero value otherwise.

### GetWorkFromDateOk

`func (o *PortalUserDto) GetWorkFromDateOk() (*time.Time, bool)`

GetWorkFromDateOk returns a tuple with the WorkFromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkFromDate

`func (o *PortalUserDto) SetWorkFromDate(v time.Time)`

SetWorkFromDate sets WorkFromDate field to given value.

### HasWorkFromDate

`func (o *PortalUserDto) HasWorkFromDate() bool`

HasWorkFromDate returns a boolean if a field has been set.

### SetWorkFromDateNil

`func (o *PortalUserDto) SetWorkFromDateNil(b bool)`

 SetWorkFromDateNil sets the value for WorkFromDate to be an explicit nil

### UnsetWorkFromDate
`func (o *PortalUserDto) UnsetWorkFromDate()`

UnsetWorkFromDate ensures that no value is present for WorkFromDate, not even an explicit nil
### GetEmail

`func (o *PortalUserDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PortalUserDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PortalUserDto) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PortalUserDto) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *PortalUserDto) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *PortalUserDto) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetContacts

`func (o *PortalUserDto) GetContacts() string`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *PortalUserDto) GetContactsOk() (*string, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *PortalUserDto) SetContacts(v string)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *PortalUserDto) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### SetContactsNil

`func (o *PortalUserDto) SetContactsNil(b bool)`

 SetContactsNil sets the value for Contacts to be an explicit nil

### UnsetContacts
`func (o *PortalUserDto) UnsetContacts()`

UnsetContacts ensures that no value is present for Contacts, not even an explicit nil
### GetContactsList

`func (o *PortalUserDto) GetContactsList() []string`

GetContactsList returns the ContactsList field if non-nil, zero value otherwise.

### GetContactsListOk

`func (o *PortalUserDto) GetContactsListOk() (*[]string, bool)`

GetContactsListOk returns a tuple with the ContactsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactsList

`func (o *PortalUserDto) SetContactsList(v []string)`

SetContactsList sets ContactsList field to given value.

### HasContactsList

`func (o *PortalUserDto) HasContactsList() bool`

HasContactsList returns a boolean if a field has been set.

### SetContactsListNil

`func (o *PortalUserDto) SetContactsListNil(b bool)`

 SetContactsListNil sets the value for ContactsList to be an explicit nil

### UnsetContactsList
`func (o *PortalUserDto) UnsetContactsList()`

UnsetContactsList ensures that no value is present for ContactsList, not even an explicit nil
### GetLocation

`func (o *PortalUserDto) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PortalUserDto) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PortalUserDto) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PortalUserDto) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *PortalUserDto) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *PortalUserDto) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetNotes

`func (o *PortalUserDto) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *PortalUserDto) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *PortalUserDto) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *PortalUserDto) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *PortalUserDto) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *PortalUserDto) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetRemoved

`func (o *PortalUserDto) GetRemoved() bool`

GetRemoved returns the Removed field if non-nil, zero value otherwise.

### GetRemovedOk

`func (o *PortalUserDto) GetRemovedOk() (*bool, bool)`

GetRemovedOk returns a tuple with the Removed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoved

`func (o *PortalUserDto) SetRemoved(v bool)`

SetRemoved sets Removed field to given value.

### HasRemoved

`func (o *PortalUserDto) HasRemoved() bool`

HasRemoved returns a boolean if a field has been set.

### GetLastModified

`func (o *PortalUserDto) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *PortalUserDto) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *PortalUserDto) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *PortalUserDto) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetTenantId

`func (o *PortalUserDto) GetTenantId() int32`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *PortalUserDto) GetTenantIdOk() (*int32, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *PortalUserDto) SetTenantId(v int32)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *PortalUserDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetIsActive

`func (o *PortalUserDto) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *PortalUserDto) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *PortalUserDto) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *PortalUserDto) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetCultureName

`func (o *PortalUserDto) GetCultureName() string`

GetCultureName returns the CultureName field if non-nil, zero value otherwise.

### GetCultureNameOk

`func (o *PortalUserDto) GetCultureNameOk() (*string, bool)`

GetCultureNameOk returns a tuple with the CultureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCultureName

`func (o *PortalUserDto) SetCultureName(v string)`

SetCultureName sets CultureName field to given value.

### HasCultureName

`func (o *PortalUserDto) HasCultureName() bool`

HasCultureName returns a boolean if a field has been set.

### SetCultureNameNil

`func (o *PortalUserDto) SetCultureNameNil(b bool)`

 SetCultureNameNil sets the value for CultureName to be an explicit nil

### UnsetCultureName
`func (o *PortalUserDto) UnsetCultureName()`

UnsetCultureName ensures that no value is present for CultureName, not even an explicit nil
### GetMobilePhone

`func (o *PortalUserDto) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *PortalUserDto) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *PortalUserDto) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *PortalUserDto) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### SetMobilePhoneNil

`func (o *PortalUserDto) SetMobilePhoneNil(b bool)`

 SetMobilePhoneNil sets the value for MobilePhone to be an explicit nil

### UnsetMobilePhone
`func (o *PortalUserDto) UnsetMobilePhone()`

UnsetMobilePhone ensures that no value is present for MobilePhone, not even an explicit nil
### GetMobilePhoneActivationStatus

`func (o *PortalUserDto) GetMobilePhoneActivationStatus() MobilePhoneActivationStatus`

GetMobilePhoneActivationStatus returns the MobilePhoneActivationStatus field if non-nil, zero value otherwise.

### GetMobilePhoneActivationStatusOk

`func (o *PortalUserDto) GetMobilePhoneActivationStatusOk() (*MobilePhoneActivationStatus, bool)`

GetMobilePhoneActivationStatusOk returns a tuple with the MobilePhoneActivationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhoneActivationStatus

`func (o *PortalUserDto) SetMobilePhoneActivationStatus(v MobilePhoneActivationStatus)`

SetMobilePhoneActivationStatus sets MobilePhoneActivationStatus field to given value.

### HasMobilePhoneActivationStatus

`func (o *PortalUserDto) HasMobilePhoneActivationStatus() bool`

HasMobilePhoneActivationStatus returns a boolean if a field has been set.

### GetSid

`func (o *PortalUserDto) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *PortalUserDto) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *PortalUserDto) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *PortalUserDto) HasSid() bool`

HasSid returns a boolean if a field has been set.

### SetSidNil

`func (o *PortalUserDto) SetSidNil(b bool)`

 SetSidNil sets the value for Sid to be an explicit nil

### UnsetSid
`func (o *PortalUserDto) UnsetSid()`

UnsetSid ensures that no value is present for Sid, not even an explicit nil
### GetLdapQouta

`func (o *PortalUserDto) GetLdapQouta() int64`

GetLdapQouta returns the LdapQouta field if non-nil, zero value otherwise.

### GetLdapQoutaOk

`func (o *PortalUserDto) GetLdapQoutaOk() (*int64, bool)`

GetLdapQoutaOk returns a tuple with the LdapQouta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapQouta

`func (o *PortalUserDto) SetLdapQouta(v int64)`

SetLdapQouta sets LdapQouta field to given value.

### HasLdapQouta

`func (o *PortalUserDto) HasLdapQouta() bool`

HasLdapQouta returns a boolean if a field has been set.

### GetSsoNameId

`func (o *PortalUserDto) GetSsoNameId() string`

GetSsoNameId returns the SsoNameId field if non-nil, zero value otherwise.

### GetSsoNameIdOk

`func (o *PortalUserDto) GetSsoNameIdOk() (*string, bool)`

GetSsoNameIdOk returns a tuple with the SsoNameId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoNameId

`func (o *PortalUserDto) SetSsoNameId(v string)`

SetSsoNameId sets SsoNameId field to given value.

### HasSsoNameId

`func (o *PortalUserDto) HasSsoNameId() bool`

HasSsoNameId returns a boolean if a field has been set.

### SetSsoNameIdNil

`func (o *PortalUserDto) SetSsoNameIdNil(b bool)`

 SetSsoNameIdNil sets the value for SsoNameId to be an explicit nil

### UnsetSsoNameId
`func (o *PortalUserDto) UnsetSsoNameId()`

UnsetSsoNameId ensures that no value is present for SsoNameId, not even an explicit nil
### GetSsoSessionId

`func (o *PortalUserDto) GetSsoSessionId() string`

GetSsoSessionId returns the SsoSessionId field if non-nil, zero value otherwise.

### GetSsoSessionIdOk

`func (o *PortalUserDto) GetSsoSessionIdOk() (*string, bool)`

GetSsoSessionIdOk returns a tuple with the SsoSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoSessionId

`func (o *PortalUserDto) SetSsoSessionId(v string)`

SetSsoSessionId sets SsoSessionId field to given value.

### HasSsoSessionId

`func (o *PortalUserDto) HasSsoSessionId() bool`

HasSsoSessionId returns a boolean if a field has been set.

### SetSsoSessionIdNil

`func (o *PortalUserDto) SetSsoSessionIdNil(b bool)`

 SetSsoSessionIdNil sets the value for SsoSessionId to be an explicit nil

### UnsetSsoSessionId
`func (o *PortalUserDto) UnsetSsoSessionId()`

UnsetSsoSessionId ensures that no value is present for SsoSessionId, not even an explicit nil
### GetCreateDate

`func (o *PortalUserDto) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *PortalUserDto) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *PortalUserDto) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *PortalUserDto) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetCreatedBy

`func (o *PortalUserDto) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PortalUserDto) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PortalUserDto) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *PortalUserDto) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *PortalUserDto) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *PortalUserDto) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetSpam

`func (o *PortalUserDto) GetSpam() bool`

GetSpam returns the Spam field if non-nil, zero value otherwise.

### GetSpamOk

`func (o *PortalUserDto) GetSpamOk() (*bool, bool)`

GetSpamOk returns a tuple with the Spam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpam

`func (o *PortalUserDto) SetSpam(v bool)`

SetSpam sets Spam field to given value.

### HasSpam

`func (o *PortalUserDto) HasSpam() bool`

HasSpam returns a boolean if a field has been set.

### SetSpamNil

`func (o *PortalUserDto) SetSpamNil(b bool)`

 SetSpamNil sets the value for Spam to be an explicit nil

### UnsetSpam
`func (o *PortalUserDto) UnsetSpam()`

UnsetSpam ensures that no value is present for Spam, not even an explicit nil
### GetCheckActivation

`func (o *PortalUserDto) GetCheckActivation() bool`

GetCheckActivation returns the CheckActivation field if non-nil, zero value otherwise.

### GetCheckActivationOk

`func (o *PortalUserDto) GetCheckActivationOk() (*bool, bool)`

GetCheckActivationOk returns a tuple with the CheckActivation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckActivation

`func (o *PortalUserDto) SetCheckActivation(v bool)`

SetCheckActivation sets CheckActivation field to given value.

### HasCheckActivation

`func (o *PortalUserDto) HasCheckActivation() bool`

HasCheckActivation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


