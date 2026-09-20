# ConfirmData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **NullableString** | The address the confirmation link was issued for. It has to be the same address the key was signed with, and  a value that is not an email address fails the request with 400. | [optional] 
**First** | Pointer to **NullableBool** | Whether the link is being followed for the first time, taken from the `first` parameter of the confirmation  URL. It is part of what the key was signed over, so passing a different value invalidates the key rather than  changing behaviour. | [optional] 
**Key** | Pointer to **NullableString** | The `key` parameter of the confirmation URL, copied verbatim. It is bound to the address and to the moment it  was issued, so it stops being accepted once the portal email key lifetime has passed. | [optional] 

## Methods

### NewConfirmData

`func NewConfirmData() *ConfirmData`

NewConfirmData instantiates a new ConfirmData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfirmDataWithDefaults

`func NewConfirmDataWithDefaults() *ConfirmData`

NewConfirmDataWithDefaults instantiates a new ConfirmData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *ConfirmData) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ConfirmData) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ConfirmData) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ConfirmData) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *ConfirmData) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ConfirmData) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetFirst

`func (o *ConfirmData) GetFirst() bool`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *ConfirmData) GetFirstOk() (*bool, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *ConfirmData) SetFirst(v bool)`

SetFirst sets First field to given value.

### HasFirst

`func (o *ConfirmData) HasFirst() bool`

HasFirst returns a boolean if a field has been set.

### SetFirstNil

`func (o *ConfirmData) SetFirstNil(b bool)`

 SetFirstNil sets the value for First to be an explicit nil

### UnsetFirst
`func (o *ConfirmData) UnsetFirst()`

UnsetFirst ensures that no value is present for First, not even an explicit nil
### GetKey

`func (o *ConfirmData) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ConfirmData) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ConfirmData) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ConfirmData) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *ConfirmData) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *ConfirmData) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


