# CspDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domains** | **[]string** | The list of CSP domains. | 
**Header** | **NullableString** | The CSP header. | 

## Methods

### NewCspDto

`func NewCspDto(domains []string, header NullableString, ) *CspDto`

NewCspDto instantiates a new CspDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCspDtoWithDefaults

`func NewCspDtoWithDefaults() *CspDto`

NewCspDtoWithDefaults instantiates a new CspDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomains

`func (o *CspDto) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *CspDto) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *CspDto) SetDomains(v []string)`

SetDomains sets Domains field to given value.


### SetDomainsNil

`func (o *CspDto) SetDomainsNil(b bool)`

 SetDomainsNil sets the value for Domains to be an explicit nil

### UnsetDomains
`func (o *CspDto) UnsetDomains()`

UnsetDomains ensures that no value is present for Domains, not even an explicit nil
### GetHeader

`func (o *CspDto) GetHeader() string`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *CspDto) GetHeaderOk() (*string, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *CspDto) SetHeader(v string)`

SetHeader sets Header field to given value.


### SetHeaderNil

`func (o *CspDto) SetHeaderNil(b bool)`

 SetHeaderNil sets the value for Header to be an explicit nil

### UnsetHeader
`func (o *CspDto) UnsetHeader()`

UnsetHeader ensures that no value is present for Header, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


