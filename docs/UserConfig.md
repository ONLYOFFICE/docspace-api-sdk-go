# UserConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The user ID. | [optional] 
**Name** | Pointer to **NullableString** | The full name of the user. | [optional] 
**Image** | Pointer to **NullableString** | The path to the user's avatar. | [optional] 
**Roles** | Pointer to **[]string** | Roles | [optional] 
**CustomerId** | Pointer to **NullableString** | Customer identifier associated with the user. | [optional] 

## Methods

### NewUserConfig

`func NewUserConfig() *UserConfig`

NewUserConfig instantiates a new UserConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserConfigWithDefaults

`func NewUserConfigWithDefaults() *UserConfig`

NewUserConfigWithDefaults instantiates a new UserConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *UserConfig) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *UserConfig) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *UserConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UserConfig) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UserConfig) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetImage

`func (o *UserConfig) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *UserConfig) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *UserConfig) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *UserConfig) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *UserConfig) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *UserConfig) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetRoles

`func (o *UserConfig) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserConfig) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserConfig) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UserConfig) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *UserConfig) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *UserConfig) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetCustomerId

`func (o *UserConfig) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *UserConfig) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *UserConfig) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *UserConfig) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *UserConfig) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *UserConfig) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


