# ReportDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collection** | Pointer to [**[]OperationDto**](OperationDto.md) | The movements on this page - top-ups, charges, refunds and corrections alike, newest first. It is empty  for a page past the end of the report as well as for a period in which nothing happened. | [optional] 
**Offset** | Pointer to **int32** | How many movements were skipped before this page, echoed from the request so a client need not remember  what it asked for. | [optional] 
**Limit** | Pointer to **int32** | How many movements one page may hold, echoed from the request; it is 25 unless another value was asked  for. A full page is not proof that more exist - compare `currentPage` with `totalPage`. | [optional] 
**TotalQuantity** | Pointer to **int64** | How many movements match the filters in total, across every page. | [optional] 
**TotalPage** | Pointer to **int32** | How many pages those movements come to at the current `limit`. | [optional] 
**CurrentPage** | Pointer to **int32** | Which of those pages this one is, as the billing service numbers them. Page through by advancing `offset`  rather than this value, which nothing accepts as an argument. | [optional] 

## Methods

### NewReportDto

`func NewReportDto() *ReportDto`

NewReportDto instantiates a new ReportDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportDtoWithDefaults

`func NewReportDtoWithDefaults() *ReportDto`

NewReportDtoWithDefaults instantiates a new ReportDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollection

`func (o *ReportDto) GetCollection() []OperationDto`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *ReportDto) GetCollectionOk() (*[]OperationDto, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *ReportDto) SetCollection(v []OperationDto)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *ReportDto) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### SetCollectionNil

`func (o *ReportDto) SetCollectionNil(b bool)`

 SetCollectionNil sets the value for Collection to be an explicit nil

### UnsetCollection
`func (o *ReportDto) UnsetCollection()`

UnsetCollection ensures that no value is present for Collection, not even an explicit nil
### GetOffset

`func (o *ReportDto) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ReportDto) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ReportDto) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ReportDto) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *ReportDto) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ReportDto) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ReportDto) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ReportDto) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetTotalQuantity

`func (o *ReportDto) GetTotalQuantity() int64`

GetTotalQuantity returns the TotalQuantity field if non-nil, zero value otherwise.

### GetTotalQuantityOk

`func (o *ReportDto) GetTotalQuantityOk() (*int64, bool)`

GetTotalQuantityOk returns a tuple with the TotalQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalQuantity

`func (o *ReportDto) SetTotalQuantity(v int64)`

SetTotalQuantity sets TotalQuantity field to given value.

### HasTotalQuantity

`func (o *ReportDto) HasTotalQuantity() bool`

HasTotalQuantity returns a boolean if a field has been set.

### GetTotalPage

`func (o *ReportDto) GetTotalPage() int32`

GetTotalPage returns the TotalPage field if non-nil, zero value otherwise.

### GetTotalPageOk

`func (o *ReportDto) GetTotalPageOk() (*int32, bool)`

GetTotalPageOk returns a tuple with the TotalPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPage

`func (o *ReportDto) SetTotalPage(v int32)`

SetTotalPage sets TotalPage field to given value.

### HasTotalPage

`func (o *ReportDto) HasTotalPage() bool`

HasTotalPage returns a boolean if a field has been set.

### GetCurrentPage

`func (o *ReportDto) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *ReportDto) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *ReportDto) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *ReportDto) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


