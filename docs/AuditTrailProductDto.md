# AuditTrailProductDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductType** | Pointer to **NullableString** | The product this branch of the tree belongs to, as the `productType` filter of this operation spells it and  as `GET api/2.0/security/audit/types` lists it under `productTypes`. | [optional] 
**Modules** | Pointer to [**[]AuditTrailModuleDto**](AuditTrailModuleDto.md) | The locations inside the product. It is empty when `moduleType` was passed and this product has no module  of that name, which is why a product can come back with nothing under it. | [optional] 

## Methods

### NewAuditTrailProductDto

`func NewAuditTrailProductDto() *AuditTrailProductDto`

NewAuditTrailProductDto instantiates a new AuditTrailProductDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditTrailProductDtoWithDefaults

`func NewAuditTrailProductDtoWithDefaults() *AuditTrailProductDto`

NewAuditTrailProductDtoWithDefaults instantiates a new AuditTrailProductDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductType

`func (o *AuditTrailProductDto) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *AuditTrailProductDto) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *AuditTrailProductDto) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *AuditTrailProductDto) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### SetProductTypeNil

`func (o *AuditTrailProductDto) SetProductTypeNil(b bool)`

 SetProductTypeNil sets the value for ProductType to be an explicit nil

### UnsetProductType
`func (o *AuditTrailProductDto) UnsetProductType()`

UnsetProductType ensures that no value is present for ProductType, not even an explicit nil
### GetModules

`func (o *AuditTrailProductDto) GetModules() []AuditTrailModuleDto`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *AuditTrailProductDto) GetModulesOk() (*[]AuditTrailModuleDto, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *AuditTrailProductDto) SetModules(v []AuditTrailModuleDto)`

SetModules sets Modules field to given value.

### HasModules

`func (o *AuditTrailProductDto) HasModules() bool`

HasModules returns a boolean if a field has been set.

### SetModulesNil

`func (o *AuditTrailProductDto) SetModulesNil(b bool)`

 SetModulesNil sets the value for Modules to be an explicit nil

### UnsetModules
`func (o *AuditTrailProductDto) UnsetModules()`

UnsetModules ensures that no value is present for Modules, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


