# OrderByDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAsc** | Pointer to **bool** | Specifies if the order is ascending. | [optional] 
**Property** | Pointer to [**SortedByType**](SortedByType.md) | The parameters by which the files will be sorted. | [optional] 

## Methods

### NewOrderByDto

`func NewOrderByDto() *OrderByDto`

NewOrderByDto instantiates a new OrderByDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderByDtoWithDefaults

`func NewOrderByDtoWithDefaults() *OrderByDto`

NewOrderByDtoWithDefaults instantiates a new OrderByDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAsc

`func (o *OrderByDto) GetIsAsc() bool`

GetIsAsc returns the IsAsc field if non-nil, zero value otherwise.

### GetIsAscOk

`func (o *OrderByDto) GetIsAscOk() (*bool, bool)`

GetIsAscOk returns a tuple with the IsAsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAsc

`func (o *OrderByDto) SetIsAsc(v bool)`

SetIsAsc sets IsAsc field to given value.

### HasIsAsc

`func (o *OrderByDto) HasIsAsc() bool`

HasIsAsc returns a boolean if a field has been set.

### GetProperty

`func (o *OrderByDto) GetProperty() SortedByType`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *OrderByDto) GetPropertyOk() (*SortedByType, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *OrderByDto) SetProperty(v SortedByType)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *OrderByDto) HasProperty() bool`

HasProperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


