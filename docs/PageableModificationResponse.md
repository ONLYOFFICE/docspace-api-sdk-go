# PageableModificationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **interface{}** |  | [optional] 
**Limit** | Pointer to **int32** | The page size that was applied to this request, between 1 and 50. | [optional] 
**LastModifiedOn** | Pointer to **time.Time** | The cursor to send back as last_modified_on to ask for the next page. It is null when the page is empty. | [optional] 

## Methods

### NewPageableModificationResponse

`func NewPageableModificationResponse() *PageableModificationResponse`

NewPageableModificationResponse instantiates a new PageableModificationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageableModificationResponseWithDefaults

`func NewPageableModificationResponseWithDefaults() *PageableModificationResponse`

NewPageableModificationResponseWithDefaults instantiates a new PageableModificationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PageableModificationResponse) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PageableModificationResponse) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PageableModificationResponse) SetData(v interface{})`

SetData sets Data field to given value.

### HasData

`func (o *PageableModificationResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *PageableModificationResponse) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *PageableModificationResponse) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetLimit

`func (o *PageableModificationResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PageableModificationResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PageableModificationResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PageableModificationResponse) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetLastModifiedOn

`func (o *PageableModificationResponse) GetLastModifiedOn() time.Time`

GetLastModifiedOn returns the LastModifiedOn field if non-nil, zero value otherwise.

### GetLastModifiedOnOk

`func (o *PageableModificationResponse) GetLastModifiedOnOk() (*time.Time, bool)`

GetLastModifiedOnOk returns a tuple with the LastModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedOn

`func (o *PageableModificationResponse) SetLastModifiedOn(v time.Time)`

SetLastModifiedOn sets LastModifiedOn field to given value.

### HasLastModifiedOn

`func (o *PageableModificationResponse) HasLastModifiedOn() bool`

HasLastModifiedOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


