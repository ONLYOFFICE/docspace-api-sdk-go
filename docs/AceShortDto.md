# AceShortDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to **NullableString** | Who or what the line stands for, as a display string: the display name of a member, the name of a group, or  the title given to a shared link when `isLink` is true. It is empty when the subject has no name to show - a  shared link that was never given a title, for instance. | [optional] 
**Permissions** | Pointer to **NullableString** | The access level of that subject as a localized label, not a code: inside a room it usually names the role the  subject holds there (Viewer, Editor, Room Manager), while outside a room it names the access itself  (Read Only, Full Access). The wording comes from the portal resources and is translated for the current  language, so show it to a person rather than compare it in code. | [optional] 
**IsLink** | Pointer to **bool** | Whether the line stands for a shared link instead of a member or a group. Clients use it to draw a link badge  where they would otherwise draw an avatar. | [optional] 

## Methods

### NewAceShortDto

`func NewAceShortDto() *AceShortDto`

NewAceShortDto instantiates a new AceShortDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAceShortDtoWithDefaults

`func NewAceShortDtoWithDefaults() *AceShortDto`

NewAceShortDtoWithDefaults instantiates a new AceShortDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *AceShortDto) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AceShortDto) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AceShortDto) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *AceShortDto) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *AceShortDto) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *AceShortDto) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetPermissions

`func (o *AceShortDto) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *AceShortDto) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *AceShortDto) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *AceShortDto) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *AceShortDto) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *AceShortDto) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetIsLink

`func (o *AceShortDto) GetIsLink() bool`

GetIsLink returns the IsLink field if non-nil, zero value otherwise.

### GetIsLinkOk

`func (o *AceShortDto) GetIsLinkOk() (*bool, bool)`

GetIsLinkOk returns a tuple with the IsLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLink

`func (o *AceShortDto) SetIsLink(v bool)`

SetIsLink sets IsLink field to given value.

### HasIsLink

`func (o *AceShortDto) HasIsLink() bool`

HasIsLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


