# IpRestrictionEntryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | **NullableString** | The IPv4 or IPv6 address. | 
**ForAdmin** | Pointer to **bool** | Whether the address admits administrators only. | [optional] 

## Methods

### NewIpRestrictionEntryDto

`func NewIpRestrictionEntryDto(ip NullableString, ) *IpRestrictionEntryDto`

NewIpRestrictionEntryDto instantiates a new IpRestrictionEntryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpRestrictionEntryDtoWithDefaults

`func NewIpRestrictionEntryDtoWithDefaults() *IpRestrictionEntryDto`

NewIpRestrictionEntryDtoWithDefaults instantiates a new IpRestrictionEntryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *IpRestrictionEntryDto) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *IpRestrictionEntryDto) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *IpRestrictionEntryDto) SetIp(v string)`

SetIp sets Ip field to given value.


### SetIpNil

`func (o *IpRestrictionEntryDto) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *IpRestrictionEntryDto) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetForAdmin

`func (o *IpRestrictionEntryDto) GetForAdmin() bool`

GetForAdmin returns the ForAdmin field if non-nil, zero value otherwise.

### GetForAdminOk

`func (o *IpRestrictionEntryDto) GetForAdminOk() (*bool, bool)`

GetForAdminOk returns a tuple with the ForAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForAdmin

`func (o *IpRestrictionEntryDto) SetForAdmin(v bool)`

SetForAdmin sets ForAdmin field to given value.

### HasForAdmin

`func (o *IpRestrictionEntryDto) HasForAdmin() bool`

HasForAdmin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


