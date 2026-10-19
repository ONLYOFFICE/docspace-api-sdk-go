# UserConfigDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The account the changes are recorded under. Two sessions carrying the same value are taken by the editors for  the same person. | [optional] 
**Name** | Pointer to **NullableString** | The name shown next to the changes and in the list of participants. | [optional] 
**Image** | Pointer to **NullableString** | An absolute address of the avatar shown for this participant. | [optional] 
**Roles** | Pointer to **[]string** | The filling roles this participant holds in the form being filled out. It is set only for a form in a virtual  data room, where the role decides which fields open for them. | [optional] 
**CustomerId** | Pointer to **NullableString** | Identifies the paying customer this participant belongs to, on deployments where the editors are licensed per  customer. | [optional] 

## Methods

### NewUserConfigDto

`func NewUserConfigDto() *UserConfigDto`

NewUserConfigDto instantiates a new UserConfigDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserConfigDtoWithDefaults

`func NewUserConfigDtoWithDefaults() *UserConfigDto`

NewUserConfigDtoWithDefaults instantiates a new UserConfigDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserConfigDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserConfigDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserConfigDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserConfigDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *UserConfigDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *UserConfigDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *UserConfigDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserConfigDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserConfigDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserConfigDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UserConfigDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UserConfigDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetImage

`func (o *UserConfigDto) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *UserConfigDto) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *UserConfigDto) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *UserConfigDto) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *UserConfigDto) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *UserConfigDto) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetRoles

`func (o *UserConfigDto) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserConfigDto) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserConfigDto) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UserConfigDto) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *UserConfigDto) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *UserConfigDto) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetCustomerId

`func (o *UserConfigDto) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *UserConfigDto) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *UserConfigDto) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *UserConfigDto) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *UserConfigDto) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *UserConfigDto) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


