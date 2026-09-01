# DocsCloudQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | Pointer to [**[]DocsCloudQuotaUser**](DocsCloudQuotaUser.md) | The editor users. | [optional] 
**UsersView** | Pointer to [**[]DocsCloudQuotaUser**](DocsCloudQuotaUser.md) | The viewer users. | [optional] 

## Methods

### NewDocsCloudQuota

`func NewDocsCloudQuota() *DocsCloudQuota`

NewDocsCloudQuota instantiates a new DocsCloudQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocsCloudQuotaWithDefaults

`func NewDocsCloudQuotaWithDefaults() *DocsCloudQuota`

NewDocsCloudQuotaWithDefaults instantiates a new DocsCloudQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *DocsCloudQuota) GetUsers() []DocsCloudQuotaUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *DocsCloudQuota) GetUsersOk() (*[]DocsCloudQuotaUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *DocsCloudQuota) SetUsers(v []DocsCloudQuotaUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *DocsCloudQuota) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsersNil

`func (o *DocsCloudQuota) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *DocsCloudQuota) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil
### GetUsersView

`func (o *DocsCloudQuota) GetUsersView() []DocsCloudQuotaUser`

GetUsersView returns the UsersView field if non-nil, zero value otherwise.

### GetUsersViewOk

`func (o *DocsCloudQuota) GetUsersViewOk() (*[]DocsCloudQuotaUser, bool)`

GetUsersViewOk returns a tuple with the UsersView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersView

`func (o *DocsCloudQuota) SetUsersView(v []DocsCloudQuotaUser)`

SetUsersView sets UsersView field to given value.

### HasUsersView

`func (o *DocsCloudQuota) HasUsersView() bool`

HasUsersView returns a boolean if a field has been set.

### SetUsersViewNil

`func (o *DocsCloudQuota) SetUsersViewNil(b bool)`

 SetUsersViewNil sets the value for UsersView to be an explicit nil

### UnsetUsersView
`func (o *DocsCloudQuota) UnsetUsersView()`

UnsetUsersView ensures that no value is present for UsersView, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


