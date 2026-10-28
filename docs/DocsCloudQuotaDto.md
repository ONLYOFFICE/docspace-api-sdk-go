# DocsCloudQuotaDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | Pointer to [**[]DocsCloudQuotaUserDto**](DocsCloudQuotaUserDto.md) | The editor users. | [optional] 
**UsersView** | Pointer to [**[]DocsCloudQuotaUserDto**](DocsCloudQuotaUserDto.md) | The viewer users. | [optional] 

## Methods

### NewDocsCloudQuotaDto

`func NewDocsCloudQuotaDto() *DocsCloudQuotaDto`

NewDocsCloudQuotaDto instantiates a new DocsCloudQuotaDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocsCloudQuotaDtoWithDefaults

`func NewDocsCloudQuotaDtoWithDefaults() *DocsCloudQuotaDto`

NewDocsCloudQuotaDtoWithDefaults instantiates a new DocsCloudQuotaDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *DocsCloudQuotaDto) GetUsers() []DocsCloudQuotaUserDto`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *DocsCloudQuotaDto) GetUsersOk() (*[]DocsCloudQuotaUserDto, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *DocsCloudQuotaDto) SetUsers(v []DocsCloudQuotaUserDto)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *DocsCloudQuotaDto) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsersNil

`func (o *DocsCloudQuotaDto) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *DocsCloudQuotaDto) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil
### GetUsersView

`func (o *DocsCloudQuotaDto) GetUsersView() []DocsCloudQuotaUserDto`

GetUsersView returns the UsersView field if non-nil, zero value otherwise.

### GetUsersViewOk

`func (o *DocsCloudQuotaDto) GetUsersViewOk() (*[]DocsCloudQuotaUserDto, bool)`

GetUsersViewOk returns a tuple with the UsersView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersView

`func (o *DocsCloudQuotaDto) SetUsersView(v []DocsCloudQuotaUserDto)`

SetUsersView sets UsersView field to given value.

### HasUsersView

`func (o *DocsCloudQuotaDto) HasUsersView() bool`

HasUsersView returns a boolean if a field has been set.

### SetUsersViewNil

`func (o *DocsCloudQuotaDto) SetUsersViewNil(b bool)`

 SetUsersViewNil sets the value for UsersView to be an explicit nil

### UnsetUsersView
`func (o *DocsCloudQuotaDto) UnsetUsersView()`

UnsetUsersView ensures that no value is present for UsersView, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


