# SaveFormRoleMappingDtoInteger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FormId** | **int32** | The form ID. | 
**Roles** | [**[]FormRole**](FormRole.md) | The collection of roles. | 

## Methods

### NewSaveFormRoleMappingDtoInteger

`func NewSaveFormRoleMappingDtoInteger(formId int32, roles []FormRole, ) *SaveFormRoleMappingDtoInteger`

NewSaveFormRoleMappingDtoInteger instantiates a new SaveFormRoleMappingDtoInteger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveFormRoleMappingDtoIntegerWithDefaults

`func NewSaveFormRoleMappingDtoIntegerWithDefaults() *SaveFormRoleMappingDtoInteger`

NewSaveFormRoleMappingDtoIntegerWithDefaults instantiates a new SaveFormRoleMappingDtoInteger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormId

`func (o *SaveFormRoleMappingDtoInteger) GetFormId() int32`

GetFormId returns the FormId field if non-nil, zero value otherwise.

### GetFormIdOk

`func (o *SaveFormRoleMappingDtoInteger) GetFormIdOk() (*int32, bool)`

GetFormIdOk returns a tuple with the FormId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormId

`func (o *SaveFormRoleMappingDtoInteger) SetFormId(v int32)`

SetFormId sets FormId field to given value.


### GetRoles

`func (o *SaveFormRoleMappingDtoInteger) GetRoles() []FormRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *SaveFormRoleMappingDtoInteger) GetRolesOk() (*[]FormRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *SaveFormRoleMappingDtoInteger) SetRoles(v []FormRole)`

SetRoles sets Roles field to given value.


### SetRolesNil

`func (o *SaveFormRoleMappingDtoInteger) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *SaveFormRoleMappingDtoInteger) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


