# IpRestrictionBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | **NullableString** |  | 
**ForAdmin** | Pointer to **bool** |  | [optional] 

## Methods

### NewIpRestrictionBase

`func NewIpRestrictionBase(ip NullableString, ) *IpRestrictionBase`

NewIpRestrictionBase instantiates a new IpRestrictionBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpRestrictionBaseWithDefaults

`func NewIpRestrictionBaseWithDefaults() *IpRestrictionBase`

NewIpRestrictionBaseWithDefaults instantiates a new IpRestrictionBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *IpRestrictionBase) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *IpRestrictionBase) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *IpRestrictionBase) SetIp(v string)`

SetIp sets Ip field to given value.


### SetIpNil

`func (o *IpRestrictionBase) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *IpRestrictionBase) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetForAdmin

`func (o *IpRestrictionBase) GetForAdmin() bool`

GetForAdmin returns the ForAdmin field if non-nil, zero value otherwise.

### GetForAdminOk

`func (o *IpRestrictionBase) GetForAdminOk() (*bool, bool)`

GetForAdminOk returns a tuple with the ForAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForAdmin

`func (o *IpRestrictionBase) SetForAdmin(v bool)`

SetForAdmin sets ForAdmin field to given value.

### HasForAdmin

`func (o *IpRestrictionBase) HasForAdmin() bool`

HasForAdmin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


