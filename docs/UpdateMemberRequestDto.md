# UpdateMemberRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **NullableString** | The user ID. | [optional] 
**Disable** | Pointer to **NullableBool** | Specifies whether to disable a user or not. | [optional] 
**Email** | Pointer to **NullableString** | The user email address. | [optional] 
**IsUser** | Pointer to **NullableBool** | Specifies if this is a guest or a user. | [optional] 
**FirstName** | Pointer to **NullableString** | The user first name. | [optional] 
**LastName** | Pointer to **NullableString** | The user last name. | [optional] 
**Department** | Pointer to **[]string** | The list of the user departments. | [optional] 
**Location** | Pointer to **NullableString** | The user location. | [optional] 
**Comment** | Pointer to **NullableString** | The user comment. | [optional] 
**Contacts** | Pointer to [**[]Contact**](Contact.md) | The list of the user contacts. | [optional] 
**Files** | Pointer to **NullableString** | The user avatar photo URL. | [optional] 
**Spam** | Pointer to **NullableBool** | Specifies if tips, updates and offers are allowed to be sent to the user or not. | [optional] 

## Methods

### NewUpdateMemberRequestDto

`func NewUpdateMemberRequestDto() *UpdateMemberRequestDto`

NewUpdateMemberRequestDto instantiates a new UpdateMemberRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMemberRequestDtoWithDefaults

`func NewUpdateMemberRequestDtoWithDefaults() *UpdateMemberRequestDto`

NewUpdateMemberRequestDtoWithDefaults instantiates a new UpdateMemberRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *UpdateMemberRequestDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UpdateMemberRequestDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UpdateMemberRequestDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UpdateMemberRequestDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *UpdateMemberRequestDto) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *UpdateMemberRequestDto) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetDisable

`func (o *UpdateMemberRequestDto) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *UpdateMemberRequestDto) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *UpdateMemberRequestDto) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *UpdateMemberRequestDto) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### SetDisableNil

`func (o *UpdateMemberRequestDto) SetDisableNil(b bool)`

 SetDisableNil sets the value for Disable to be an explicit nil

### UnsetDisable
`func (o *UpdateMemberRequestDto) UnsetDisable()`

UnsetDisable ensures that no value is present for Disable, not even an explicit nil
### GetEmail

`func (o *UpdateMemberRequestDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateMemberRequestDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateMemberRequestDto) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateMemberRequestDto) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *UpdateMemberRequestDto) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *UpdateMemberRequestDto) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetIsUser

`func (o *UpdateMemberRequestDto) GetIsUser() bool`

GetIsUser returns the IsUser field if non-nil, zero value otherwise.

### GetIsUserOk

`func (o *UpdateMemberRequestDto) GetIsUserOk() (*bool, bool)`

GetIsUserOk returns a tuple with the IsUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUser

`func (o *UpdateMemberRequestDto) SetIsUser(v bool)`

SetIsUser sets IsUser field to given value.

### HasIsUser

`func (o *UpdateMemberRequestDto) HasIsUser() bool`

HasIsUser returns a boolean if a field has been set.

### SetIsUserNil

`func (o *UpdateMemberRequestDto) SetIsUserNil(b bool)`

 SetIsUserNil sets the value for IsUser to be an explicit nil

### UnsetIsUser
`func (o *UpdateMemberRequestDto) UnsetIsUser()`

UnsetIsUser ensures that no value is present for IsUser, not even an explicit nil
### GetFirstName

`func (o *UpdateMemberRequestDto) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UpdateMemberRequestDto) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UpdateMemberRequestDto) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UpdateMemberRequestDto) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *UpdateMemberRequestDto) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *UpdateMemberRequestDto) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *UpdateMemberRequestDto) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UpdateMemberRequestDto) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UpdateMemberRequestDto) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UpdateMemberRequestDto) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *UpdateMemberRequestDto) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *UpdateMemberRequestDto) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetDepartment

`func (o *UpdateMemberRequestDto) GetDepartment() []string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *UpdateMemberRequestDto) GetDepartmentOk() (*[]string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *UpdateMemberRequestDto) SetDepartment(v []string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *UpdateMemberRequestDto) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartmentNil

`func (o *UpdateMemberRequestDto) SetDepartmentNil(b bool)`

 SetDepartmentNil sets the value for Department to be an explicit nil

### UnsetDepartment
`func (o *UpdateMemberRequestDto) UnsetDepartment()`

UnsetDepartment ensures that no value is present for Department, not even an explicit nil
### GetLocation

`func (o *UpdateMemberRequestDto) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *UpdateMemberRequestDto) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *UpdateMemberRequestDto) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *UpdateMemberRequestDto) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *UpdateMemberRequestDto) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *UpdateMemberRequestDto) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetComment

`func (o *UpdateMemberRequestDto) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *UpdateMemberRequestDto) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *UpdateMemberRequestDto) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *UpdateMemberRequestDto) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *UpdateMemberRequestDto) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *UpdateMemberRequestDto) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetContacts

`func (o *UpdateMemberRequestDto) GetContacts() []Contact`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *UpdateMemberRequestDto) GetContactsOk() (*[]Contact, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *UpdateMemberRequestDto) SetContacts(v []Contact)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *UpdateMemberRequestDto) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### SetContactsNil

`func (o *UpdateMemberRequestDto) SetContactsNil(b bool)`

 SetContactsNil sets the value for Contacts to be an explicit nil

### UnsetContacts
`func (o *UpdateMemberRequestDto) UnsetContacts()`

UnsetContacts ensures that no value is present for Contacts, not even an explicit nil
### GetFiles

`func (o *UpdateMemberRequestDto) GetFiles() string`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *UpdateMemberRequestDto) GetFilesOk() (*string, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *UpdateMemberRequestDto) SetFiles(v string)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *UpdateMemberRequestDto) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### SetFilesNil

`func (o *UpdateMemberRequestDto) SetFilesNil(b bool)`

 SetFilesNil sets the value for Files to be an explicit nil

### UnsetFiles
`func (o *UpdateMemberRequestDto) UnsetFiles()`

UnsetFiles ensures that no value is present for Files, not even an explicit nil
### GetSpam

`func (o *UpdateMemberRequestDto) GetSpam() bool`

GetSpam returns the Spam field if non-nil, zero value otherwise.

### GetSpamOk

`func (o *UpdateMemberRequestDto) GetSpamOk() (*bool, bool)`

GetSpamOk returns a tuple with the Spam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpam

`func (o *UpdateMemberRequestDto) SetSpam(v bool)`

SetSpam sets Spam field to given value.

### HasSpam

`func (o *UpdateMemberRequestDto) HasSpam() bool`

HasSpam returns a boolean if a field has been set.

### SetSpamNil

`func (o *UpdateMemberRequestDto) SetSpamNil(b bool)`

 SetSpamNil sets the value for Spam to be an explicit nil

### UnsetSpam
`func (o *UpdateMemberRequestDto) UnsetSpam()`

UnsetSpam ensures that no value is present for Spam, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


