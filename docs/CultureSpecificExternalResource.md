# CultureSpecificExternalResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **NullableString** | The external resource domain. | [optional] 
**Entries** | Pointer to **map[string]string** | The external resource entries. | [optional] 

## Methods

### NewCultureSpecificExternalResource

`func NewCultureSpecificExternalResource() *CultureSpecificExternalResource`

NewCultureSpecificExternalResource instantiates a new CultureSpecificExternalResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCultureSpecificExternalResourceWithDefaults

`func NewCultureSpecificExternalResourceWithDefaults() *CultureSpecificExternalResource`

NewCultureSpecificExternalResourceWithDefaults instantiates a new CultureSpecificExternalResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *CultureSpecificExternalResource) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CultureSpecificExternalResource) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CultureSpecificExternalResource) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *CultureSpecificExternalResource) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *CultureSpecificExternalResource) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *CultureSpecificExternalResource) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetEntries

`func (o *CultureSpecificExternalResource) GetEntries() map[string]*string`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *CultureSpecificExternalResource) GetEntriesOk() (*map[string]*string, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *CultureSpecificExternalResource) SetEntries(v map[string]*string)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *CultureSpecificExternalResource) HasEntries() bool`

HasEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


