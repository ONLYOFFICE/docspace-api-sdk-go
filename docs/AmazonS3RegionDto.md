# AmazonS3RegionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SystemName** | Pointer to **NullableString** | The region code to send as the region value when configuring an Amazon S3 storage or backup target. It is  the one field of this object that is an argument elsewhere; a code the server does not list here cannot be  reached, so pick one from this list rather than typing it. | [optional] 
**DisplayName** | Pointer to **NullableString** | The region name as Amazon writes it, in English regardless of the portal language, for showing in a  picker next to `systemName`. | [optional] 
**PartitionName** | Pointer to **NullableString** | The Amazon partition the region sits in - the ordinary commercial cloud, the Chinese one, or a government  one. Regions of different partitions are not reachable with the same credentials. | [optional] 
**PartitionDnsSuffix** | Pointer to **NullableString** | The domain the partition's service host names end in, which differs from partition to partition. | [optional] 
**PartitionRegionRegex** | Pointer to **NullableString** | The pattern every region code of this partition matches, for validating a code before sending it. | [optional] 
**HostnameTemplate** | Pointer to **NullableString** | How a service host name of the partition is assembled, with `{service}`, `{region}` and `{dnsSuffix}` to  be filled in. It is reference material - the portal builds its own endpoints from `systemName`. | [optional] 

## Methods

### NewAmazonS3RegionDto

`func NewAmazonS3RegionDto() *AmazonS3RegionDto`

NewAmazonS3RegionDto instantiates a new AmazonS3RegionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonS3RegionDtoWithDefaults

`func NewAmazonS3RegionDtoWithDefaults() *AmazonS3RegionDto`

NewAmazonS3RegionDtoWithDefaults instantiates a new AmazonS3RegionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSystemName

`func (o *AmazonS3RegionDto) GetSystemName() string`

GetSystemName returns the SystemName field if non-nil, zero value otherwise.

### GetSystemNameOk

`func (o *AmazonS3RegionDto) GetSystemNameOk() (*string, bool)`

GetSystemNameOk returns a tuple with the SystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemName

`func (o *AmazonS3RegionDto) SetSystemName(v string)`

SetSystemName sets SystemName field to given value.

### HasSystemName

`func (o *AmazonS3RegionDto) HasSystemName() bool`

HasSystemName returns a boolean if a field has been set.

### SetSystemNameNil

`func (o *AmazonS3RegionDto) SetSystemNameNil(b bool)`

 SetSystemNameNil sets the value for SystemName to be an explicit nil

### UnsetSystemName
`func (o *AmazonS3RegionDto) UnsetSystemName()`

UnsetSystemName ensures that no value is present for SystemName, not even an explicit nil
### GetDisplayName

`func (o *AmazonS3RegionDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AmazonS3RegionDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AmazonS3RegionDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AmazonS3RegionDto) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *AmazonS3RegionDto) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *AmazonS3RegionDto) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetPartitionName

`func (o *AmazonS3RegionDto) GetPartitionName() string`

GetPartitionName returns the PartitionName field if non-nil, zero value otherwise.

### GetPartitionNameOk

`func (o *AmazonS3RegionDto) GetPartitionNameOk() (*string, bool)`

GetPartitionNameOk returns a tuple with the PartitionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionName

`func (o *AmazonS3RegionDto) SetPartitionName(v string)`

SetPartitionName sets PartitionName field to given value.

### HasPartitionName

`func (o *AmazonS3RegionDto) HasPartitionName() bool`

HasPartitionName returns a boolean if a field has been set.

### SetPartitionNameNil

`func (o *AmazonS3RegionDto) SetPartitionNameNil(b bool)`

 SetPartitionNameNil sets the value for PartitionName to be an explicit nil

### UnsetPartitionName
`func (o *AmazonS3RegionDto) UnsetPartitionName()`

UnsetPartitionName ensures that no value is present for PartitionName, not even an explicit nil
### GetPartitionDnsSuffix

`func (o *AmazonS3RegionDto) GetPartitionDnsSuffix() string`

GetPartitionDnsSuffix returns the PartitionDnsSuffix field if non-nil, zero value otherwise.

### GetPartitionDnsSuffixOk

`func (o *AmazonS3RegionDto) GetPartitionDnsSuffixOk() (*string, bool)`

GetPartitionDnsSuffixOk returns a tuple with the PartitionDnsSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionDnsSuffix

`func (o *AmazonS3RegionDto) SetPartitionDnsSuffix(v string)`

SetPartitionDnsSuffix sets PartitionDnsSuffix field to given value.

### HasPartitionDnsSuffix

`func (o *AmazonS3RegionDto) HasPartitionDnsSuffix() bool`

HasPartitionDnsSuffix returns a boolean if a field has been set.

### SetPartitionDnsSuffixNil

`func (o *AmazonS3RegionDto) SetPartitionDnsSuffixNil(b bool)`

 SetPartitionDnsSuffixNil sets the value for PartitionDnsSuffix to be an explicit nil

### UnsetPartitionDnsSuffix
`func (o *AmazonS3RegionDto) UnsetPartitionDnsSuffix()`

UnsetPartitionDnsSuffix ensures that no value is present for PartitionDnsSuffix, not even an explicit nil
### GetPartitionRegionRegex

`func (o *AmazonS3RegionDto) GetPartitionRegionRegex() string`

GetPartitionRegionRegex returns the PartitionRegionRegex field if non-nil, zero value otherwise.

### GetPartitionRegionRegexOk

`func (o *AmazonS3RegionDto) GetPartitionRegionRegexOk() (*string, bool)`

GetPartitionRegionRegexOk returns a tuple with the PartitionRegionRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionRegionRegex

`func (o *AmazonS3RegionDto) SetPartitionRegionRegex(v string)`

SetPartitionRegionRegex sets PartitionRegionRegex field to given value.

### HasPartitionRegionRegex

`func (o *AmazonS3RegionDto) HasPartitionRegionRegex() bool`

HasPartitionRegionRegex returns a boolean if a field has been set.

### SetPartitionRegionRegexNil

`func (o *AmazonS3RegionDto) SetPartitionRegionRegexNil(b bool)`

 SetPartitionRegionRegexNil sets the value for PartitionRegionRegex to be an explicit nil

### UnsetPartitionRegionRegex
`func (o *AmazonS3RegionDto) UnsetPartitionRegionRegex()`

UnsetPartitionRegionRegex ensures that no value is present for PartitionRegionRegex, not even an explicit nil
### GetHostnameTemplate

`func (o *AmazonS3RegionDto) GetHostnameTemplate() string`

GetHostnameTemplate returns the HostnameTemplate field if non-nil, zero value otherwise.

### GetHostnameTemplateOk

`func (o *AmazonS3RegionDto) GetHostnameTemplateOk() (*string, bool)`

GetHostnameTemplateOk returns a tuple with the HostnameTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameTemplate

`func (o *AmazonS3RegionDto) SetHostnameTemplate(v string)`

SetHostnameTemplate sets HostnameTemplate field to given value.

### HasHostnameTemplate

`func (o *AmazonS3RegionDto) HasHostnameTemplate() bool`

HasHostnameTemplate returns a boolean if a field has been set.

### SetHostnameTemplateNil

`func (o *AmazonS3RegionDto) SetHostnameTemplateNil(b bool)`

 SetHostnameTemplateNil sets the value for HostnameTemplate to be an explicit nil

### UnsetHostnameTemplate
`func (o *AmazonS3RegionDto) UnsetHostnameTemplate()`

UnsetHostnameTemplate ensures that no value is present for HostnameTemplate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


