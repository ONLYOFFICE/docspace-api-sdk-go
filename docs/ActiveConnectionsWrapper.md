# ActiveConnectionsWrapper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | Pointer to [**ActiveConnectionsDto**](ActiveConnectionsDto.md) | The ActiveConnectionsDto object returned by the operation. | [optional] 
**Count** | Pointer to **int32** | The total number of items in the response | [optional] 
**Links** | Pointer to [**[]GetPortalPrices200ResponseLinksInner**](GetPortalPrices200ResponseLinksInner.md) | List of links related to the response | [optional] 
**Status** | Pointer to **int32** | HTTP status code of the response | [optional] 
**StatusCode** | Pointer to **int32** | HTTP status code of the response (duplicate of status) | [optional] 

## Methods

### NewActiveConnectionsWrapper

`func NewActiveConnectionsWrapper() *ActiveConnectionsWrapper`

NewActiveConnectionsWrapper instantiates a new ActiveConnectionsWrapper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveConnectionsWrapperWithDefaults

`func NewActiveConnectionsWrapperWithDefaults() *ActiveConnectionsWrapper`

NewActiveConnectionsWrapperWithDefaults instantiates a new ActiveConnectionsWrapper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponse

`func (o *ActiveConnectionsWrapper) GetResponse() ActiveConnectionsDto`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *ActiveConnectionsWrapper) GetResponseOk() (*ActiveConnectionsDto, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *ActiveConnectionsWrapper) SetResponse(v ActiveConnectionsDto)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *ActiveConnectionsWrapper) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetCount

`func (o *ActiveConnectionsWrapper) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ActiveConnectionsWrapper) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ActiveConnectionsWrapper) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ActiveConnectionsWrapper) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetLinks

`func (o *ActiveConnectionsWrapper) GetLinks() []GetPortalPrices200ResponseLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ActiveConnectionsWrapper) GetLinksOk() (*[]GetPortalPrices200ResponseLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ActiveConnectionsWrapper) SetLinks(v []GetPortalPrices200ResponseLinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ActiveConnectionsWrapper) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetStatus

`func (o *ActiveConnectionsWrapper) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ActiveConnectionsWrapper) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ActiveConnectionsWrapper) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ActiveConnectionsWrapper) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCode

`func (o *ActiveConnectionsWrapper) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ActiveConnectionsWrapper) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ActiveConnectionsWrapper) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *ActiveConnectionsWrapper) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


