# CspRequestsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domains** | Pointer to **[]string** | The collection of allowed domains in the Content Security Policy (CSP). | [optional] 

## Methods

### NewCspRequestsDto

`func NewCspRequestsDto() *CspRequestsDto`

NewCspRequestsDto instantiates a new CspRequestsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCspRequestsDtoWithDefaults

`func NewCspRequestsDtoWithDefaults() *CspRequestsDto`

NewCspRequestsDtoWithDefaults instantiates a new CspRequestsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomains

`func (o *CspRequestsDto) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *CspRequestsDto) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *CspRequestsDto) SetDomains(v []string)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *CspRequestsDto) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### SetDomainsNil

`func (o *CspRequestsDto) SetDomainsNil(b bool)`

 SetDomainsNil sets the value for Domains to be an explicit nil

### UnsetDomains
`func (o *CspRequestsDto) UnsetDomains()`

UnsetDomains ensures that no value is present for Domains, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


