# SaveFormRoleMappingDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FormId** | **int32** | The PDF form the roles belong to. This is the value the operation reads, rather than the identifier in its  route, and the two are to be sent the same. | 
**Roles** | [**[]FormRole**](FormRole.md) | The roles with the account taking each of them and the sequence number that decides the turn: the same number  means the roles may be filled in parallel, different ones make a queue. The whole set is replaced on every  call, and an empty set resets the filling. | 

## Methods

### NewSaveFormRoleMappingDto

`func NewSaveFormRoleMappingDto(formId int32, roles []FormRole, ) *SaveFormRoleMappingDto`

NewSaveFormRoleMappingDto instantiates a new SaveFormRoleMappingDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveFormRoleMappingDtoWithDefaults

`func NewSaveFormRoleMappingDtoWithDefaults() *SaveFormRoleMappingDto`

NewSaveFormRoleMappingDtoWithDefaults instantiates a new SaveFormRoleMappingDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormId

`func (o *SaveFormRoleMappingDto) GetFormId() int32`

GetFormId returns the FormId field if non-nil, zero value otherwise.

### GetFormIdOk

`func (o *SaveFormRoleMappingDto) GetFormIdOk() (*int32, bool)`

GetFormIdOk returns a tuple with the FormId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormId

`func (o *SaveFormRoleMappingDto) SetFormId(v int32)`

SetFormId sets FormId field to given value.


### GetRoles

`func (o *SaveFormRoleMappingDto) GetRoles() []FormRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *SaveFormRoleMappingDto) GetRolesOk() (*[]FormRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *SaveFormRoleMappingDto) SetRoles(v []FormRole)`

SetRoles sets Roles field to given value.


### SetRolesNil

`func (o *SaveFormRoleMappingDto) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *SaveFormRoleMappingDto) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


