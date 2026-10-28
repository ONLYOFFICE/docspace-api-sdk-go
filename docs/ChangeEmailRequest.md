# ChangeEmailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **NullableString** | The new address in plain text, up to 255 characters. It is stored in lowercase, and one of this field and  `encEmail` is required. | [optional] 
**EncEmail** | Pointer to **NullableString** | The new address in the encrypted form the confirmation link carries. Pass the value from the link unchanged;  it is used only when `email` is empty. | [optional] 

## Methods

### NewChangeEmailRequest

`func NewChangeEmailRequest() *ChangeEmailRequest`

NewChangeEmailRequest instantiates a new ChangeEmailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeEmailRequestWithDefaults

`func NewChangeEmailRequestWithDefaults() *ChangeEmailRequest`

NewChangeEmailRequestWithDefaults instantiates a new ChangeEmailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *ChangeEmailRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ChangeEmailRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ChangeEmailRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ChangeEmailRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *ChangeEmailRequest) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ChangeEmailRequest) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetEncEmail

`func (o *ChangeEmailRequest) GetEncEmail() string`

GetEncEmail returns the EncEmail field if non-nil, zero value otherwise.

### GetEncEmailOk

`func (o *ChangeEmailRequest) GetEncEmailOk() (*string, bool)`

GetEncEmailOk returns a tuple with the EncEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncEmail

`func (o *ChangeEmailRequest) SetEncEmail(v string)`

SetEncEmail sets EncEmail field to given value.

### HasEncEmail

`func (o *ChangeEmailRequest) HasEncEmail() bool`

HasEncEmail returns a boolean if a field has been set.

### SetEncEmailNil

`func (o *ChangeEmailRequest) SetEncEmailNil(b bool)`

 SetEncEmailNil sets the value for EncEmail to be an explicit nil

### UnsetEncEmail
`func (o *ChangeEmailRequest) UnsetEncEmail()`

UnsetEncEmail ensures that no value is present for EncEmail, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


