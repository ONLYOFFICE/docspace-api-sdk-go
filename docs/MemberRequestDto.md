# MemberRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **NullableString** | The user password. | [optional] 
**PasswordHash** | Pointer to **NullableString** | The user password hash. | [optional] 
**Email** | Pointer to **NullableString** | The user email address. | [optional] 
**Type** | Pointer to [**EmployeeType**](EmployeeType.md) |  | [optional] 
**IsUser** | Pointer to **NullableBool** | Specifies if this is a guest or a user. | [optional] 
**FirstName** | Pointer to **NullableString** | The user first name. | [optional] 
**LastName** | Pointer to **NullableString** | The user last name. | [optional] 
**Department** | Pointer to **[]string** | The list of the user departments IDs. | [optional] 
**Title** | Pointer to **NullableString** | The user title. | [optional] 
**Location** | Pointer to **NullableString** | The user location. | [optional] 
**Sex** | Pointer to [**SexEnum**](SexEnum.md) |  | [optional] 
**Birthday** | Pointer to [**ApiDateTime**](ApiDateTime.md) |  | [optional] 
**Worksfrom** | Pointer to [**ApiDateTime**](ApiDateTime.md) |  | [optional] 
**Comment** | Pointer to **NullableString** | The user comment. | [optional] 
**Contacts** | Pointer to [**[]Contact**](Contact.md) | The list of the user contacts. | [optional] 
**Files** | Pointer to **NullableString** | The avatar photo URL. | [optional] 
**FromInviteLink** | Pointer to **bool** | Specifies if the user is added via the invitation link or not. | [optional] 
**Key** | Pointer to **NullableString** | The user key. | [optional] 
**CultureName** | Pointer to **NullableString** | The user culture code. | [optional] 
**Target** | Pointer to **string** | The user target ID. | [optional] 
**Spam** | Pointer to **NullableBool** | Specifies if tips, updates and offers are allowed to be sent to the user or not. | [optional] 

## Methods

### NewMemberRequestDto

`func NewMemberRequestDto() *MemberRequestDto`

NewMemberRequestDto instantiates a new MemberRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberRequestDtoWithDefaults

`func NewMemberRequestDtoWithDefaults() *MemberRequestDto`

NewMemberRequestDtoWithDefaults instantiates a new MemberRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *MemberRequestDto) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *MemberRequestDto) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *MemberRequestDto) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *MemberRequestDto) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *MemberRequestDto) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *MemberRequestDto) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPasswordHash

`func (o *MemberRequestDto) GetPasswordHash() string`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *MemberRequestDto) GetPasswordHashOk() (*string, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *MemberRequestDto) SetPasswordHash(v string)`

SetPasswordHash sets PasswordHash field to given value.

### HasPasswordHash

`func (o *MemberRequestDto) HasPasswordHash() bool`

HasPasswordHash returns a boolean if a field has been set.

### SetPasswordHashNil

`func (o *MemberRequestDto) SetPasswordHashNil(b bool)`

 SetPasswordHashNil sets the value for PasswordHash to be an explicit nil

### UnsetPasswordHash
`func (o *MemberRequestDto) UnsetPasswordHash()`

UnsetPasswordHash ensures that no value is present for PasswordHash, not even an explicit nil
### GetEmail

`func (o *MemberRequestDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MemberRequestDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MemberRequestDto) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MemberRequestDto) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *MemberRequestDto) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *MemberRequestDto) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetType

`func (o *MemberRequestDto) GetType() EmployeeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MemberRequestDto) GetTypeOk() (*EmployeeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MemberRequestDto) SetType(v EmployeeType)`

SetType sets Type field to given value.

### HasType

`func (o *MemberRequestDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIsUser

`func (o *MemberRequestDto) GetIsUser() bool`

GetIsUser returns the IsUser field if non-nil, zero value otherwise.

### GetIsUserOk

`func (o *MemberRequestDto) GetIsUserOk() (*bool, bool)`

GetIsUserOk returns a tuple with the IsUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUser

`func (o *MemberRequestDto) SetIsUser(v bool)`

SetIsUser sets IsUser field to given value.

### HasIsUser

`func (o *MemberRequestDto) HasIsUser() bool`

HasIsUser returns a boolean if a field has been set.

### SetIsUserNil

`func (o *MemberRequestDto) SetIsUserNil(b bool)`

 SetIsUserNil sets the value for IsUser to be an explicit nil

### UnsetIsUser
`func (o *MemberRequestDto) UnsetIsUser()`

UnsetIsUser ensures that no value is present for IsUser, not even an explicit nil
### GetFirstName

`func (o *MemberRequestDto) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *MemberRequestDto) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *MemberRequestDto) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *MemberRequestDto) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *MemberRequestDto) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *MemberRequestDto) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *MemberRequestDto) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *MemberRequestDto) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *MemberRequestDto) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *MemberRequestDto) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *MemberRequestDto) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *MemberRequestDto) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetDepartment

`func (o *MemberRequestDto) GetDepartment() []string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *MemberRequestDto) GetDepartmentOk() (*[]string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *MemberRequestDto) SetDepartment(v []string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *MemberRequestDto) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartmentNil

`func (o *MemberRequestDto) SetDepartmentNil(b bool)`

 SetDepartmentNil sets the value for Department to be an explicit nil

### UnsetDepartment
`func (o *MemberRequestDto) UnsetDepartment()`

UnsetDepartment ensures that no value is present for Department, not even an explicit nil
### GetTitle

`func (o *MemberRequestDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MemberRequestDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MemberRequestDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MemberRequestDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *MemberRequestDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *MemberRequestDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetLocation

`func (o *MemberRequestDto) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *MemberRequestDto) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *MemberRequestDto) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *MemberRequestDto) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *MemberRequestDto) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *MemberRequestDto) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetSex

`func (o *MemberRequestDto) GetSex() SexEnum`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *MemberRequestDto) GetSexOk() (*SexEnum, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *MemberRequestDto) SetSex(v SexEnum)`

SetSex sets Sex field to given value.

### HasSex

`func (o *MemberRequestDto) HasSex() bool`

HasSex returns a boolean if a field has been set.

### GetBirthday

`func (o *MemberRequestDto) GetBirthday() ApiDateTime`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *MemberRequestDto) GetBirthdayOk() (*ApiDateTime, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *MemberRequestDto) SetBirthday(v ApiDateTime)`

SetBirthday sets Birthday field to given value.

### HasBirthday

`func (o *MemberRequestDto) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### GetWorksfrom

`func (o *MemberRequestDto) GetWorksfrom() ApiDateTime`

GetWorksfrom returns the Worksfrom field if non-nil, zero value otherwise.

### GetWorksfromOk

`func (o *MemberRequestDto) GetWorksfromOk() (*ApiDateTime, bool)`

GetWorksfromOk returns a tuple with the Worksfrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorksfrom

`func (o *MemberRequestDto) SetWorksfrom(v ApiDateTime)`

SetWorksfrom sets Worksfrom field to given value.

### HasWorksfrom

`func (o *MemberRequestDto) HasWorksfrom() bool`

HasWorksfrom returns a boolean if a field has been set.

### GetComment

`func (o *MemberRequestDto) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *MemberRequestDto) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *MemberRequestDto) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *MemberRequestDto) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *MemberRequestDto) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *MemberRequestDto) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetContacts

`func (o *MemberRequestDto) GetContacts() []Contact`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *MemberRequestDto) GetContactsOk() (*[]Contact, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *MemberRequestDto) SetContacts(v []Contact)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *MemberRequestDto) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### SetContactsNil

`func (o *MemberRequestDto) SetContactsNil(b bool)`

 SetContactsNil sets the value for Contacts to be an explicit nil

### UnsetContacts
`func (o *MemberRequestDto) UnsetContacts()`

UnsetContacts ensures that no value is present for Contacts, not even an explicit nil
### GetFiles

`func (o *MemberRequestDto) GetFiles() string`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *MemberRequestDto) GetFilesOk() (*string, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *MemberRequestDto) SetFiles(v string)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *MemberRequestDto) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### SetFilesNil

`func (o *MemberRequestDto) SetFilesNil(b bool)`

 SetFilesNil sets the value for Files to be an explicit nil

### UnsetFiles
`func (o *MemberRequestDto) UnsetFiles()`

UnsetFiles ensures that no value is present for Files, not even an explicit nil
### GetFromInviteLink

`func (o *MemberRequestDto) GetFromInviteLink() bool`

GetFromInviteLink returns the FromInviteLink field if non-nil, zero value otherwise.

### GetFromInviteLinkOk

`func (o *MemberRequestDto) GetFromInviteLinkOk() (*bool, bool)`

GetFromInviteLinkOk returns a tuple with the FromInviteLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromInviteLink

`func (o *MemberRequestDto) SetFromInviteLink(v bool)`

SetFromInviteLink sets FromInviteLink field to given value.

### HasFromInviteLink

`func (o *MemberRequestDto) HasFromInviteLink() bool`

HasFromInviteLink returns a boolean if a field has been set.

### GetKey

`func (o *MemberRequestDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MemberRequestDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MemberRequestDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *MemberRequestDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *MemberRequestDto) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *MemberRequestDto) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetCultureName

`func (o *MemberRequestDto) GetCultureName() string`

GetCultureName returns the CultureName field if non-nil, zero value otherwise.

### GetCultureNameOk

`func (o *MemberRequestDto) GetCultureNameOk() (*string, bool)`

GetCultureNameOk returns a tuple with the CultureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCultureName

`func (o *MemberRequestDto) SetCultureName(v string)`

SetCultureName sets CultureName field to given value.

### HasCultureName

`func (o *MemberRequestDto) HasCultureName() bool`

HasCultureName returns a boolean if a field has been set.

### SetCultureNameNil

`func (o *MemberRequestDto) SetCultureNameNil(b bool)`

 SetCultureNameNil sets the value for CultureName to be an explicit nil

### UnsetCultureName
`func (o *MemberRequestDto) UnsetCultureName()`

UnsetCultureName ensures that no value is present for CultureName, not even an explicit nil
### GetTarget

`func (o *MemberRequestDto) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *MemberRequestDto) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *MemberRequestDto) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *MemberRequestDto) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetSpam

`func (o *MemberRequestDto) GetSpam() bool`

GetSpam returns the Spam field if non-nil, zero value otherwise.

### GetSpamOk

`func (o *MemberRequestDto) GetSpamOk() (*bool, bool)`

GetSpamOk returns a tuple with the Spam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpam

`func (o *MemberRequestDto) SetSpam(v bool)`

SetSpam sets Spam field to given value.

### HasSpam

`func (o *MemberRequestDto) HasSpam() bool`

HasSpam returns a boolean if a field has been set.

### SetSpamNil

`func (o *MemberRequestDto) SetSpamNil(b bool)`

 SetSpamNil sets the value for Spam to be an explicit nil

### UnsetSpam
`func (o *MemberRequestDto) UnsetSpam()`

UnsetSpam ensures that no value is present for Spam, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


