# PageableClientInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ClientInfoResponse**](ClientInfoResponse.md) | The items on this page, at most as many as the requested limit. An empty array means there is nothing further to read. | [optional] 
**Limit** | Pointer to **int32** | The page size that was applied to this request, between 1 and 50. | [optional] 
**LastClientId** | Pointer to **string** | The cursor to send back as last_client_id to ask for the next page, together with last_created_on. It is null when the page is empty. | [optional] 
**LastCreatedOn** | Pointer to **time.Time** | The cursor to send back as last_created_on to ask for the next page, together with last_client_id. It is null when the page is empty. | [optional] 

## Methods

### NewPageableClientInfoResponse

`func NewPageableClientInfoResponse() *PageableClientInfoResponse`

NewPageableClientInfoResponse instantiates a new PageableClientInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageableClientInfoResponseWithDefaults

`func NewPageableClientInfoResponseWithDefaults() *PageableClientInfoResponse`

NewPageableClientInfoResponseWithDefaults instantiates a new PageableClientInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PageableClientInfoResponse) GetData() []ClientInfoResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PageableClientInfoResponse) GetDataOk() (*[]ClientInfoResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PageableClientInfoResponse) SetData(v []ClientInfoResponse)`

SetData sets Data field to given value.

### HasData

`func (o *PageableClientInfoResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLimit

`func (o *PageableClientInfoResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PageableClientInfoResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PageableClientInfoResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PageableClientInfoResponse) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetLastClientId

`func (o *PageableClientInfoResponse) GetLastClientId() string`

GetLastClientId returns the LastClientId field if non-nil, zero value otherwise.

### GetLastClientIdOk

`func (o *PageableClientInfoResponse) GetLastClientIdOk() (*string, bool)`

GetLastClientIdOk returns a tuple with the LastClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastClientId

`func (o *PageableClientInfoResponse) SetLastClientId(v string)`

SetLastClientId sets LastClientId field to given value.

### HasLastClientId

`func (o *PageableClientInfoResponse) HasLastClientId() bool`

HasLastClientId returns a boolean if a field has been set.

### GetLastCreatedOn

`func (o *PageableClientInfoResponse) GetLastCreatedOn() time.Time`

GetLastCreatedOn returns the LastCreatedOn field if non-nil, zero value otherwise.

### GetLastCreatedOnOk

`func (o *PageableClientInfoResponse) GetLastCreatedOnOk() (*time.Time, bool)`

GetLastCreatedOnOk returns a tuple with the LastCreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCreatedOn

`func (o *PageableClientInfoResponse) SetLastCreatedOn(v time.Time)`

SetLastCreatedOn sets LastCreatedOn field to given value.

### HasLastCreatedOn

`func (o *PageableClientInfoResponse) HasLastCreatedOn() bool`

HasLastCreatedOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


