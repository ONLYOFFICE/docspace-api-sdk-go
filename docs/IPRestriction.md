# IPRestriction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | **string** | The IP address. | 
**ForAdmin** | Pointer to **bool** | Specifies if the IP address is for administrator users only or not. | [optional] 
**Id** | Pointer to **int32** | The IP restiction ID. | [optional] 
**TenantId** | Pointer to **int32** | The tenant ID. | [optional] 

## Methods

### NewIPRestriction

`func NewIPRestriction(ip string, ) *IPRestriction`

NewIPRestriction instantiates a new IPRestriction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPRestrictionWithDefaults

`func NewIPRestrictionWithDefaults() *IPRestriction`

NewIPRestrictionWithDefaults instantiates a new IPRestriction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *IPRestriction) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *IPRestriction) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *IPRestriction) SetIp(v string)`

SetIp sets Ip field to given value.


### GetForAdmin

`func (o *IPRestriction) GetForAdmin() bool`

GetForAdmin returns the ForAdmin field if non-nil, zero value otherwise.

### GetForAdminOk

`func (o *IPRestriction) GetForAdminOk() (*bool, bool)`

GetForAdminOk returns a tuple with the ForAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForAdmin

`func (o *IPRestriction) SetForAdmin(v bool)`

SetForAdmin sets ForAdmin field to given value.

### HasForAdmin

`func (o *IPRestriction) HasForAdmin() bool`

HasForAdmin returns a boolean if a field has been set.

### GetId

`func (o *IPRestriction) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IPRestriction) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IPRestriction) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *IPRestriction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTenantId

`func (o *IPRestriction) GetTenantId() int32`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *IPRestriction) GetTenantIdOk() (*int32, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *IPRestriction) SetTenantId(v int32)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *IPRestriction) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


