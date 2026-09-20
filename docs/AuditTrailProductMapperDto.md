# AuditTrailProductMapperDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductType** | Pointer to **NullableString** | The product this branch of the tree belongs to, as the `productType` filter of this operation spells it and  as `GET api/2.0/security/audit/types` lists it under `productTypes`. | [optional] 
**Modules** | Pointer to [**[]AuditTrailModuleMapperDto**](AuditTrailModuleMapperDto.md) | The locations inside the product. It is empty when `moduleType` was passed and this product has no module  of that name, which is why a product can come back with nothing under it. | [optional] 

## Methods

### NewAuditTrailProductMapperDto

`func NewAuditTrailProductMapperDto() *AuditTrailProductMapperDto`

NewAuditTrailProductMapperDto instantiates a new AuditTrailProductMapperDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditTrailProductMapperDtoWithDefaults

`func NewAuditTrailProductMapperDtoWithDefaults() *AuditTrailProductMapperDto`

NewAuditTrailProductMapperDtoWithDefaults instantiates a new AuditTrailProductMapperDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductType

`func (o *AuditTrailProductMapperDto) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *AuditTrailProductMapperDto) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *AuditTrailProductMapperDto) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *AuditTrailProductMapperDto) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### SetProductTypeNil

`func (o *AuditTrailProductMapperDto) SetProductTypeNil(b bool)`

 SetProductTypeNil sets the value for ProductType to be an explicit nil

### UnsetProductType
`func (o *AuditTrailProductMapperDto) UnsetProductType()`

UnsetProductType ensures that no value is present for ProductType, not even an explicit nil
### GetModules

`func (o *AuditTrailProductMapperDto) GetModules() []AuditTrailModuleMapperDto`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *AuditTrailProductMapperDto) GetModulesOk() (*[]AuditTrailModuleMapperDto, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *AuditTrailProductMapperDto) SetModules(v []AuditTrailModuleMapperDto)`

SetModules sets Modules field to given value.

### HasModules

`func (o *AuditTrailProductMapperDto) HasModules() bool`

HasModules returns a boolean if a field has been set.

### SetModulesNil

`func (o *AuditTrailProductMapperDto) SetModulesNil(b bool)`

 SetModulesNil sets the value for Modules to be an explicit nil

### UnsetModules
`func (o *AuditTrailProductMapperDto) UnsetModules()`

UnsetModules ensures that no value is present for Modules, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


