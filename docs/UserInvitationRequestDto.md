# UserInvitationRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | The address of somebody who has no portal account yet. An invitation is sent to it and an account is created  once it is accepted, so this is the field to use instead of an account identifier when the person is new to  the portal. | [optional] 
**Type** | Pointer to [**EmployeeType**](EmployeeType.md) | The user type. | [optional] 

## Methods

### NewUserInvitationRequestDto

`func NewUserInvitationRequestDto() *UserInvitationRequestDto`

NewUserInvitationRequestDto instantiates a new UserInvitationRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserInvitationRequestDtoWithDefaults

`func NewUserInvitationRequestDtoWithDefaults() *UserInvitationRequestDto`

NewUserInvitationRequestDtoWithDefaults instantiates a new UserInvitationRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UserInvitationRequestDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserInvitationRequestDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserInvitationRequestDto) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserInvitationRequestDto) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetType

`func (o *UserInvitationRequestDto) GetType() EmployeeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserInvitationRequestDto) GetTypeOk() (*EmployeeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserInvitationRequestDto) SetType(v EmployeeType)`

SetType sets Type field to given value.

### HasType

`func (o *UserInvitationRequestDto) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


