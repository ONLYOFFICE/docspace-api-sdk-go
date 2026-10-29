# ServiceLimitReportDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collection** | Pointer to [**[]ServiceLimitDto**](ServiceLimitDto.md) | The limits of this page, one per user; empty when no limit applies to any user or the page is past the end. | [optional] 
**Offset** | Pointer to **int32** | The number of limits skipped before this page. | [optional] 
**Limit** | Pointer to **int32** | The page size the page was read with. | [optional] 
**TotalPage** | Pointer to **int32** | The number of pages the whole list takes at this page size. | [optional] 
**CurrentPage** | Pointer to **int32** | The number of this page, starting from 1. | [optional] 
**TotalQuantity** | Pointer to **int32** | The number of users with a limit in force across all pages. | [optional] 

## Methods

### NewServiceLimitReportDto

`func NewServiceLimitReportDto() *ServiceLimitReportDto`

NewServiceLimitReportDto instantiates a new ServiceLimitReportDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceLimitReportDtoWithDefaults

`func NewServiceLimitReportDtoWithDefaults() *ServiceLimitReportDto`

NewServiceLimitReportDtoWithDefaults instantiates a new ServiceLimitReportDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollection

`func (o *ServiceLimitReportDto) GetCollection() []ServiceLimitDto`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *ServiceLimitReportDto) GetCollectionOk() (*[]ServiceLimitDto, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *ServiceLimitReportDto) SetCollection(v []ServiceLimitDto)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *ServiceLimitReportDto) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### SetCollectionNil

`func (o *ServiceLimitReportDto) SetCollectionNil(b bool)`

 SetCollectionNil sets the value for Collection to be an explicit nil

### UnsetCollection
`func (o *ServiceLimitReportDto) UnsetCollection()`

UnsetCollection ensures that no value is present for Collection, not even an explicit nil
### GetOffset

`func (o *ServiceLimitReportDto) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ServiceLimitReportDto) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ServiceLimitReportDto) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ServiceLimitReportDto) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *ServiceLimitReportDto) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ServiceLimitReportDto) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ServiceLimitReportDto) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ServiceLimitReportDto) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetTotalPage

`func (o *ServiceLimitReportDto) GetTotalPage() int32`

GetTotalPage returns the TotalPage field if non-nil, zero value otherwise.

### GetTotalPageOk

`func (o *ServiceLimitReportDto) GetTotalPageOk() (*int32, bool)`

GetTotalPageOk returns a tuple with the TotalPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPage

`func (o *ServiceLimitReportDto) SetTotalPage(v int32)`

SetTotalPage sets TotalPage field to given value.

### HasTotalPage

`func (o *ServiceLimitReportDto) HasTotalPage() bool`

HasTotalPage returns a boolean if a field has been set.

### GetCurrentPage

`func (o *ServiceLimitReportDto) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *ServiceLimitReportDto) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *ServiceLimitReportDto) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *ServiceLimitReportDto) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetTotalQuantity

`func (o *ServiceLimitReportDto) GetTotalQuantity() int32`

GetTotalQuantity returns the TotalQuantity field if non-nil, zero value otherwise.

### GetTotalQuantityOk

`func (o *ServiceLimitReportDto) GetTotalQuantityOk() (*int32, bool)`

GetTotalQuantityOk returns a tuple with the TotalQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalQuantity

`func (o *ServiceLimitReportDto) SetTotalQuantity(v int32)`

SetTotalQuantity sets TotalQuantity field to given value.

### HasTotalQuantity

`func (o *ServiceLimitReportDto) HasTotalQuantity() bool`

HasTotalQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


