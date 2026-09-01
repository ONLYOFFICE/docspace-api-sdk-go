# AceShortWrapper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to **NullableString** | The name of the user the document will be shared with. | [optional] 
**Permissions** | Pointer to **NullableString** | The access rights for the user with the name above.  Can be Full Access, Read Only, or Deny Access. | [optional] 
**IsLink** | Pointer to **bool** | Specifies whether to change the user icon to the link icon. | [optional] 

## Methods

### NewAceShortWrapper

`func NewAceShortWrapper() *AceShortWrapper`

NewAceShortWrapper instantiates a new AceShortWrapper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAceShortWrapperWithDefaults

`func NewAceShortWrapperWithDefaults() *AceShortWrapper`

NewAceShortWrapperWithDefaults instantiates a new AceShortWrapper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *AceShortWrapper) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AceShortWrapper) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AceShortWrapper) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *AceShortWrapper) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *AceShortWrapper) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *AceShortWrapper) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetPermissions

`func (o *AceShortWrapper) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *AceShortWrapper) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *AceShortWrapper) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *AceShortWrapper) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *AceShortWrapper) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *AceShortWrapper) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetIsLink

`func (o *AceShortWrapper) GetIsLink() bool`

GetIsLink returns the IsLink field if non-nil, zero value otherwise.

### GetIsLinkOk

`func (o *AceShortWrapper) GetIsLinkOk() (*bool, bool)`

GetIsLinkOk returns a tuple with the IsLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLink

`func (o *AceShortWrapper) SetIsLink(v bool)`

SetIsLink sets IsLink field to given value.

### HasIsLink

`func (o *AceShortWrapper) HasIsLink() bool`

HasIsLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


