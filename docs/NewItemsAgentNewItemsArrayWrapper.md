# NewItemsAgentNewItemsArrayWrapper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | Pointer to [**[]NewItemsDtoAgentNewItemsDto**](NewItemsDtoAgentNewItemsDto.md) |  | [optional] 
**Count** | Pointer to **int32** | The total number of items in the response | [optional] 
**Links** | Pointer to [**[]GetPortalPrices200ResponseLinksInner**](GetPortalPrices200ResponseLinksInner.md) | List of links related to the response | [optional] 
**Status** | Pointer to **int32** | HTTP status code of the response | [optional] 
**StatusCode** | Pointer to **int32** | HTTP status code of the response (duplicate of status) | [optional] 

## Methods

### NewNewItemsAgentNewItemsArrayWrapper

`func NewNewItemsAgentNewItemsArrayWrapper() *NewItemsAgentNewItemsArrayWrapper`

NewNewItemsAgentNewItemsArrayWrapper instantiates a new NewItemsAgentNewItemsArrayWrapper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewItemsAgentNewItemsArrayWrapperWithDefaults

`func NewNewItemsAgentNewItemsArrayWrapperWithDefaults() *NewItemsAgentNewItemsArrayWrapper`

NewNewItemsAgentNewItemsArrayWrapperWithDefaults instantiates a new NewItemsAgentNewItemsArrayWrapper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponse

`func (o *NewItemsAgentNewItemsArrayWrapper) GetResponse() []NewItemsDtoAgentNewItemsDto`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *NewItemsAgentNewItemsArrayWrapper) GetResponseOk() (*[]NewItemsDtoAgentNewItemsDto, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *NewItemsAgentNewItemsArrayWrapper) SetResponse(v []NewItemsDtoAgentNewItemsDto)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *NewItemsAgentNewItemsArrayWrapper) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetCount

`func (o *NewItemsAgentNewItemsArrayWrapper) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *NewItemsAgentNewItemsArrayWrapper) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *NewItemsAgentNewItemsArrayWrapper) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *NewItemsAgentNewItemsArrayWrapper) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetLinks

`func (o *NewItemsAgentNewItemsArrayWrapper) GetLinks() []GetPortalPrices200ResponseLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NewItemsAgentNewItemsArrayWrapper) GetLinksOk() (*[]GetPortalPrices200ResponseLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NewItemsAgentNewItemsArrayWrapper) SetLinks(v []GetPortalPrices200ResponseLinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NewItemsAgentNewItemsArrayWrapper) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetStatus

`func (o *NewItemsAgentNewItemsArrayWrapper) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NewItemsAgentNewItemsArrayWrapper) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NewItemsAgentNewItemsArrayWrapper) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NewItemsAgentNewItemsArrayWrapper) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCode

`func (o *NewItemsAgentNewItemsArrayWrapper) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *NewItemsAgentNewItemsArrayWrapper) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *NewItemsAgentNewItemsArrayWrapper) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *NewItemsAgentNewItemsArrayWrapper) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


