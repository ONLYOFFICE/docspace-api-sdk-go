# IpRestrictionsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpRestrictions** | [**[]IpRestrictionBase**](IpRestrictionBase.md) | The allowed addresses, each entry pairing a single IPv4 or IPv6 address with the flag that limits it to  administrators. This is the whole list that is to hold afterwards: entries not repeated here are deleted.  Ranges written as `from-to` and CIDR blocks are refused with 400, even though the portal matches such forms  when they are already stored. Enforcement spares only the portal owner and the installation own networks, so  a list without the caller address locks the remaining administrators out. | 
**Enable** | Pointer to **NullableBool** | Whether the list is enforced. Leaving it out follows the list - on when addresses are sent, off when the list  is empty - and sending `true` with an empty list is refused with 400, since that would admit nobody. | [optional] 

## Methods

### NewIpRestrictionsDto

`func NewIpRestrictionsDto(ipRestrictions []IpRestrictionBase, ) *IpRestrictionsDto`

NewIpRestrictionsDto instantiates a new IpRestrictionsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpRestrictionsDtoWithDefaults

`func NewIpRestrictionsDtoWithDefaults() *IpRestrictionsDto`

NewIpRestrictionsDtoWithDefaults instantiates a new IpRestrictionsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpRestrictions

`func (o *IpRestrictionsDto) GetIpRestrictions() []IpRestrictionBase`

GetIpRestrictions returns the IpRestrictions field if non-nil, zero value otherwise.

### GetIpRestrictionsOk

`func (o *IpRestrictionsDto) GetIpRestrictionsOk() (*[]IpRestrictionBase, bool)`

GetIpRestrictionsOk returns a tuple with the IpRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRestrictions

`func (o *IpRestrictionsDto) SetIpRestrictions(v []IpRestrictionBase)`

SetIpRestrictions sets IpRestrictions field to given value.


### SetIpRestrictionsNil

`func (o *IpRestrictionsDto) SetIpRestrictionsNil(b bool)`

 SetIpRestrictionsNil sets the value for IpRestrictions to be an explicit nil

### UnsetIpRestrictions
`func (o *IpRestrictionsDto) UnsetIpRestrictions()`

UnsetIpRestrictions ensures that no value is present for IpRestrictions, not even an explicit nil
### GetEnable

`func (o *IpRestrictionsDto) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *IpRestrictionsDto) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *IpRestrictionsDto) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *IpRestrictionsDto) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### SetEnableNil

`func (o *IpRestrictionsDto) SetEnableNil(b bool)`

 SetEnableNil sets the value for Enable to be an explicit nil

### UnsetEnable
`func (o *IpRestrictionsDto) UnsetEnable()`

UnsetEnable ensures that no value is present for Enable, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


