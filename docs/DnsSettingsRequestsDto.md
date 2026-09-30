# DnsSettingsRequestsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsName** | Pointer to **NullableString** | The domain the portal is to be reachable under, as a bare hostname without a scheme. It must not collide with  the reserved base domain of the installation, and a name that fails validation is refused without disturbing  the mapping in force. It is read only while `enable` is true. | [optional] 
**Enable** | Pointer to **bool** | Whether the custom domain is put in force. Setting it false clears the mapping and ignores `dnsName`; setting  it true also stops the previous domain from answering and rewrites any Content Security Policy entry that  named it. | [optional] 

## Methods

### NewDnsSettingsRequestsDto

`func NewDnsSettingsRequestsDto() *DnsSettingsRequestsDto`

NewDnsSettingsRequestsDto instantiates a new DnsSettingsRequestsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsSettingsRequestsDtoWithDefaults

`func NewDnsSettingsRequestsDtoWithDefaults() *DnsSettingsRequestsDto`

NewDnsSettingsRequestsDtoWithDefaults instantiates a new DnsSettingsRequestsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsName

`func (o *DnsSettingsRequestsDto) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *DnsSettingsRequestsDto) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *DnsSettingsRequestsDto) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *DnsSettingsRequestsDto) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### SetDnsNameNil

`func (o *DnsSettingsRequestsDto) SetDnsNameNil(b bool)`

 SetDnsNameNil sets the value for DnsName to be an explicit nil

### UnsetDnsName
`func (o *DnsSettingsRequestsDto) UnsetDnsName()`

UnsetDnsName ensures that no value is present for DnsName, not even an explicit nil
### GetEnable

`func (o *DnsSettingsRequestsDto) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *DnsSettingsRequestsDto) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *DnsSettingsRequestsDto) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *DnsSettingsRequestsDto) HasEnable() bool`

HasEnable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


