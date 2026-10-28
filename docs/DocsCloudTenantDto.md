# DocsCloudTenantDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DedicatedResourceExId** | Pointer to **int32** | The external ID of the dedicated resource the tenant is hosted on. | [optional] 
**Alias** | Pointer to **NullableString** | The tenant alias. | [optional] 
**Name** | Pointer to **NullableString** | The tenant name. | [optional] 
**ModifiedDate** | Pointer to **time.Time** | The date and time when the tenant was last modified. | [optional] 
**CustomerId** | Pointer to **NullableString** | The customer ID. | [optional] 
**CustomerName** | Pointer to **NullableString** | The customer name. | [optional] 
**EndDate** | Pointer to **time.Time** | The date and time when the tenant subscription ends. | [optional] 
**ResourceType** | Pointer to **int32** | The resource type. | [optional] 
**IsActive** | Pointer to **bool** | Whether the tenant is active (the end date is in the future). | [optional] 
**Address** | Pointer to **NullableString** | The tenant address. | [optional] 
**Payment** | Pointer to [**DocsCloudPaymentDto**](DocsCloudPaymentDto.md) | The tenant payment information. | [optional] 

## Methods

### NewDocsCloudTenantDto

`func NewDocsCloudTenantDto() *DocsCloudTenantDto`

NewDocsCloudTenantDto instantiates a new DocsCloudTenantDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocsCloudTenantDtoWithDefaults

`func NewDocsCloudTenantDtoWithDefaults() *DocsCloudTenantDto`

NewDocsCloudTenantDtoWithDefaults instantiates a new DocsCloudTenantDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDedicatedResourceExId

`func (o *DocsCloudTenantDto) GetDedicatedResourceExId() int32`

GetDedicatedResourceExId returns the DedicatedResourceExId field if non-nil, zero value otherwise.

### GetDedicatedResourceExIdOk

`func (o *DocsCloudTenantDto) GetDedicatedResourceExIdOk() (*int32, bool)`

GetDedicatedResourceExIdOk returns a tuple with the DedicatedResourceExId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedResourceExId

`func (o *DocsCloudTenantDto) SetDedicatedResourceExId(v int32)`

SetDedicatedResourceExId sets DedicatedResourceExId field to given value.

### HasDedicatedResourceExId

`func (o *DocsCloudTenantDto) HasDedicatedResourceExId() bool`

HasDedicatedResourceExId returns a boolean if a field has been set.

### GetAlias

`func (o *DocsCloudTenantDto) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *DocsCloudTenantDto) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *DocsCloudTenantDto) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *DocsCloudTenantDto) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### SetAliasNil

`func (o *DocsCloudTenantDto) SetAliasNil(b bool)`

 SetAliasNil sets the value for Alias to be an explicit nil

### UnsetAlias
`func (o *DocsCloudTenantDto) UnsetAlias()`

UnsetAlias ensures that no value is present for Alias, not even an explicit nil
### GetName

`func (o *DocsCloudTenantDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DocsCloudTenantDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DocsCloudTenantDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DocsCloudTenantDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DocsCloudTenantDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DocsCloudTenantDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetModifiedDate

`func (o *DocsCloudTenantDto) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *DocsCloudTenantDto) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *DocsCloudTenantDto) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *DocsCloudTenantDto) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### GetCustomerId

`func (o *DocsCloudTenantDto) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DocsCloudTenantDto) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DocsCloudTenantDto) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *DocsCloudTenantDto) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *DocsCloudTenantDto) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *DocsCloudTenantDto) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetCustomerName

`func (o *DocsCloudTenantDto) GetCustomerName() string`

GetCustomerName returns the CustomerName field if non-nil, zero value otherwise.

### GetCustomerNameOk

`func (o *DocsCloudTenantDto) GetCustomerNameOk() (*string, bool)`

GetCustomerNameOk returns a tuple with the CustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerName

`func (o *DocsCloudTenantDto) SetCustomerName(v string)`

SetCustomerName sets CustomerName field to given value.

### HasCustomerName

`func (o *DocsCloudTenantDto) HasCustomerName() bool`

HasCustomerName returns a boolean if a field has been set.

### SetCustomerNameNil

`func (o *DocsCloudTenantDto) SetCustomerNameNil(b bool)`

 SetCustomerNameNil sets the value for CustomerName to be an explicit nil

### UnsetCustomerName
`func (o *DocsCloudTenantDto) UnsetCustomerName()`

UnsetCustomerName ensures that no value is present for CustomerName, not even an explicit nil
### GetEndDate

`func (o *DocsCloudTenantDto) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *DocsCloudTenantDto) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *DocsCloudTenantDto) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *DocsCloudTenantDto) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetResourceType

`func (o *DocsCloudTenantDto) GetResourceType() int32`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *DocsCloudTenantDto) GetResourceTypeOk() (*int32, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *DocsCloudTenantDto) SetResourceType(v int32)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *DocsCloudTenantDto) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetIsActive

`func (o *DocsCloudTenantDto) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *DocsCloudTenantDto) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *DocsCloudTenantDto) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *DocsCloudTenantDto) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetAddress

`func (o *DocsCloudTenantDto) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DocsCloudTenantDto) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DocsCloudTenantDto) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DocsCloudTenantDto) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *DocsCloudTenantDto) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *DocsCloudTenantDto) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetPayment

`func (o *DocsCloudTenantDto) GetPayment() DocsCloudPaymentDto`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *DocsCloudTenantDto) GetPaymentOk() (*DocsCloudPaymentDto, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *DocsCloudTenantDto) SetPayment(v DocsCloudPaymentDto)`

SetPayment sets Payment field to given value.

### HasPayment

`func (o *DocsCloudTenantDto) HasPayment() bool`

HasPayment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


