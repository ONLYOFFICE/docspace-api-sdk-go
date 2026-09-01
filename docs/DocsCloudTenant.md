# DocsCloudTenant

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
**Payment** | Pointer to [**DocsCloudPayment**](DocsCloudPayment.md) | The tenant payment information. | [optional] 

## Methods

### NewDocsCloudTenant

`func NewDocsCloudTenant() *DocsCloudTenant`

NewDocsCloudTenant instantiates a new DocsCloudTenant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocsCloudTenantWithDefaults

`func NewDocsCloudTenantWithDefaults() *DocsCloudTenant`

NewDocsCloudTenantWithDefaults instantiates a new DocsCloudTenant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDedicatedResourceExId

`func (o *DocsCloudTenant) GetDedicatedResourceExId() int32`

GetDedicatedResourceExId returns the DedicatedResourceExId field if non-nil, zero value otherwise.

### GetDedicatedResourceExIdOk

`func (o *DocsCloudTenant) GetDedicatedResourceExIdOk() (*int32, bool)`

GetDedicatedResourceExIdOk returns a tuple with the DedicatedResourceExId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedResourceExId

`func (o *DocsCloudTenant) SetDedicatedResourceExId(v int32)`

SetDedicatedResourceExId sets DedicatedResourceExId field to given value.

### HasDedicatedResourceExId

`func (o *DocsCloudTenant) HasDedicatedResourceExId() bool`

HasDedicatedResourceExId returns a boolean if a field has been set.

### GetAlias

`func (o *DocsCloudTenant) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *DocsCloudTenant) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *DocsCloudTenant) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *DocsCloudTenant) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### SetAliasNil

`func (o *DocsCloudTenant) SetAliasNil(b bool)`

 SetAliasNil sets the value for Alias to be an explicit nil

### UnsetAlias
`func (o *DocsCloudTenant) UnsetAlias()`

UnsetAlias ensures that no value is present for Alias, not even an explicit nil
### GetName

`func (o *DocsCloudTenant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DocsCloudTenant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DocsCloudTenant) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DocsCloudTenant) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DocsCloudTenant) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DocsCloudTenant) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetModifiedDate

`func (o *DocsCloudTenant) GetModifiedDate() time.Time`

GetModifiedDate returns the ModifiedDate field if non-nil, zero value otherwise.

### GetModifiedDateOk

`func (o *DocsCloudTenant) GetModifiedDateOk() (*time.Time, bool)`

GetModifiedDateOk returns a tuple with the ModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedDate

`func (o *DocsCloudTenant) SetModifiedDate(v time.Time)`

SetModifiedDate sets ModifiedDate field to given value.

### HasModifiedDate

`func (o *DocsCloudTenant) HasModifiedDate() bool`

HasModifiedDate returns a boolean if a field has been set.

### GetCustomerId

`func (o *DocsCloudTenant) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *DocsCloudTenant) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *DocsCloudTenant) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *DocsCloudTenant) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *DocsCloudTenant) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *DocsCloudTenant) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetCustomerName

`func (o *DocsCloudTenant) GetCustomerName() string`

GetCustomerName returns the CustomerName field if non-nil, zero value otherwise.

### GetCustomerNameOk

`func (o *DocsCloudTenant) GetCustomerNameOk() (*string, bool)`

GetCustomerNameOk returns a tuple with the CustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerName

`func (o *DocsCloudTenant) SetCustomerName(v string)`

SetCustomerName sets CustomerName field to given value.

### HasCustomerName

`func (o *DocsCloudTenant) HasCustomerName() bool`

HasCustomerName returns a boolean if a field has been set.

### SetCustomerNameNil

`func (o *DocsCloudTenant) SetCustomerNameNil(b bool)`

 SetCustomerNameNil sets the value for CustomerName to be an explicit nil

### UnsetCustomerName
`func (o *DocsCloudTenant) UnsetCustomerName()`

UnsetCustomerName ensures that no value is present for CustomerName, not even an explicit nil
### GetEndDate

`func (o *DocsCloudTenant) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *DocsCloudTenant) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *DocsCloudTenant) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *DocsCloudTenant) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetResourceType

`func (o *DocsCloudTenant) GetResourceType() int32`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *DocsCloudTenant) GetResourceTypeOk() (*int32, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *DocsCloudTenant) SetResourceType(v int32)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *DocsCloudTenant) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetIsActive

`func (o *DocsCloudTenant) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *DocsCloudTenant) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *DocsCloudTenant) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *DocsCloudTenant) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetAddress

`func (o *DocsCloudTenant) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DocsCloudTenant) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DocsCloudTenant) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DocsCloudTenant) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *DocsCloudTenant) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *DocsCloudTenant) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetPayment

`func (o *DocsCloudTenant) GetPayment() DocsCloudPayment`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *DocsCloudTenant) GetPaymentOk() (*DocsCloudPayment, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *DocsCloudTenant) SetPayment(v DocsCloudPayment)`

SetPayment sets Payment field to given value.

### HasPayment

`func (o *DocsCloudTenant) HasPayment() bool`

HasPayment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


