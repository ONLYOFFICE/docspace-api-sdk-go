# AuditTrailModuleMapperDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModuleType** | Pointer to **NullableString** | The location inside the product, as the `moduleType` filter of `GET api/2.0/security/audit/events/filter`  spells it. | [optional] 
**Actions** | Pointer to [**[]AuditTrailActionMapperDto**](AuditTrailActionMapperDto.md) | Every action this module can record. Each action appears under exactly one module, so this tree is where a  caller learns which module a given action belongs to. | [optional] 

## Methods

### NewAuditTrailModuleMapperDto

`func NewAuditTrailModuleMapperDto() *AuditTrailModuleMapperDto`

NewAuditTrailModuleMapperDto instantiates a new AuditTrailModuleMapperDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditTrailModuleMapperDtoWithDefaults

`func NewAuditTrailModuleMapperDtoWithDefaults() *AuditTrailModuleMapperDto`

NewAuditTrailModuleMapperDtoWithDefaults instantiates a new AuditTrailModuleMapperDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModuleType

`func (o *AuditTrailModuleMapperDto) GetModuleType() string`

GetModuleType returns the ModuleType field if non-nil, zero value otherwise.

### GetModuleTypeOk

`func (o *AuditTrailModuleMapperDto) GetModuleTypeOk() (*string, bool)`

GetModuleTypeOk returns a tuple with the ModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleType

`func (o *AuditTrailModuleMapperDto) SetModuleType(v string)`

SetModuleType sets ModuleType field to given value.

### HasModuleType

`func (o *AuditTrailModuleMapperDto) HasModuleType() bool`

HasModuleType returns a boolean if a field has been set.

### SetModuleTypeNil

`func (o *AuditTrailModuleMapperDto) SetModuleTypeNil(b bool)`

 SetModuleTypeNil sets the value for ModuleType to be an explicit nil

### UnsetModuleType
`func (o *AuditTrailModuleMapperDto) UnsetModuleType()`

UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
### GetActions

`func (o *AuditTrailModuleMapperDto) GetActions() []AuditTrailActionMapperDto`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *AuditTrailModuleMapperDto) GetActionsOk() (*[]AuditTrailActionMapperDto, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *AuditTrailModuleMapperDto) SetActions(v []AuditTrailActionMapperDto)`

SetActions sets Actions field to given value.

### HasActions

`func (o *AuditTrailModuleMapperDto) HasActions() bool`

HasActions returns a boolean if a field has been set.

### SetActionsNil

`func (o *AuditTrailModuleMapperDto) SetActionsNil(b bool)`

 SetActionsNil sets the value for Actions to be an explicit nil

### UnsetActions
`func (o *AuditTrailModuleMapperDto) UnsetActions()`

UnsetActions ensures that no value is present for Actions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


