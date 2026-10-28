# PageableClientResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ClientResponse**](ClientResponse.md) | The items on this page, at most as many as the requested limit. An empty array means there is nothing further to read. | [optional] 
**Limit** | Pointer to **int32** | The page size that was applied to this request, between 1 and 50. | [optional] 
**LastClientId** | Pointer to **string** | The cursor to send back as last_client_id to ask for the next page, together with last_created_on. It is null when the page is empty. | [optional] 
**LastCreatedOn** | Pointer to **time.Time** | The cursor to send back as last_created_on to ask for the next page, together with last_client_id. It is null when the page is empty. | [optional] 

## Methods

### NewPageableClientResponse

`func NewPageableClientResponse() *PageableClientResponse`

NewPageableClientResponse instantiates a new PageableClientResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageableClientResponseWithDefaults

`func NewPageableClientResponseWithDefaults() *PageableClientResponse`

NewPageableClientResponseWithDefaults instantiates a new PageableClientResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PageableClientResponse) GetData() []ClientResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PageableClientResponse) GetDataOk() (*[]ClientResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PageableClientResponse) SetData(v []ClientResponse)`

SetData sets Data field to given value.

### HasData

`func (o *PageableClientResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLimit

`func (o *PageableClientResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PageableClientResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PageableClientResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PageableClientResponse) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetLastClientId

`func (o *PageableClientResponse) GetLastClientId() string`

GetLastClientId returns the LastClientId field if non-nil, zero value otherwise.

### GetLastClientIdOk

`func (o *PageableClientResponse) GetLastClientIdOk() (*string, bool)`

GetLastClientIdOk returns a tuple with the LastClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastClientId

`func (o *PageableClientResponse) SetLastClientId(v string)`

SetLastClientId sets LastClientId field to given value.

### HasLastClientId

`func (o *PageableClientResponse) HasLastClientId() bool`

HasLastClientId returns a boolean if a field has been set.

### GetLastCreatedOn

`func (o *PageableClientResponse) GetLastCreatedOn() time.Time`

GetLastCreatedOn returns the LastCreatedOn field if non-nil, zero value otherwise.

### GetLastCreatedOnOk

`func (o *PageableClientResponse) GetLastCreatedOnOk() (*time.Time, bool)`

GetLastCreatedOnOk returns a tuple with the LastCreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCreatedOn

`func (o *PageableClientResponse) SetLastCreatedOn(v time.Time)`

SetLastCreatedOn sets LastCreatedOn field to given value.

### HasLastCreatedOn

`func (o *PageableClientResponse) HasLastCreatedOn() bool`

HasLastCreatedOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


