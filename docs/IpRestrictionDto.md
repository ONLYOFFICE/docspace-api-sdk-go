# IpRestrictionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The IP restiction ID. | [optional] 
**TenantId** | Pointer to **int32** | The tenant ID. | [optional] 
**Ip** | **NullableString** | The IP address. | 
**ForAdmin** | Pointer to **bool** | Specifies if the IP address is for administrator users only or not. | [optional] 

## Methods

### NewIpRestrictionDto

`func NewIpRestrictionDto(ip NullableString, ) *IpRestrictionDto`

NewIpRestrictionDto instantiates a new IpRestrictionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpRestrictionDtoWithDefaults

`func NewIpRestrictionDtoWithDefaults() *IpRestrictionDto`

NewIpRestrictionDtoWithDefaults instantiates a new IpRestrictionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IpRestrictionDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IpRestrictionDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IpRestrictionDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *IpRestrictionDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTenantId

`func (o *IpRestrictionDto) GetTenantId() int32`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *IpRestrictionDto) GetTenantIdOk() (*int32, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *IpRestrictionDto) SetTenantId(v int32)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *IpRestrictionDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetIp

`func (o *IpRestrictionDto) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *IpRestrictionDto) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *IpRestrictionDto) SetIp(v string)`

SetIp sets Ip field to given value.


### SetIpNil

`func (o *IpRestrictionDto) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *IpRestrictionDto) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetForAdmin

`func (o *IpRestrictionDto) GetForAdmin() bool`

GetForAdmin returns the ForAdmin field if non-nil, zero value otherwise.

### GetForAdminOk

`func (o *IpRestrictionDto) GetForAdminOk() (*bool, bool)`

GetForAdminOk returns a tuple with the ForAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForAdmin

`func (o *IpRestrictionDto) SetForAdmin(v bool)`

SetForAdmin sets ForAdmin field to given value.

### HasForAdmin

`func (o *IpRestrictionDto) HasForAdmin() bool`

HasForAdmin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


