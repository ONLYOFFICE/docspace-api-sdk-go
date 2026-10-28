# ThirdPartyFileWrapper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | Pointer to [**ThirdPartyFileDto**](ThirdPartyFileDto.md) | The ThirdPartyFileDto object returned by the operation. | [optional] 
**Count** | Pointer to **int32** | The total number of items in the response | [optional] 
**Links** | Pointer to [**[]GetPortalPrices200ResponseLinksInner**](GetPortalPrices200ResponseLinksInner.md) | List of links related to the response | [optional] 
**Status** | Pointer to **int32** | HTTP status code of the response | [optional] 
**StatusCode** | Pointer to **int32** | HTTP status code of the response (duplicate of status) | [optional] 

## Methods

### NewThirdPartyFileWrapper

`func NewThirdPartyFileWrapper() *ThirdPartyFileWrapper`

NewThirdPartyFileWrapper instantiates a new ThirdPartyFileWrapper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThirdPartyFileWrapperWithDefaults

`func NewThirdPartyFileWrapperWithDefaults() *ThirdPartyFileWrapper`

NewThirdPartyFileWrapperWithDefaults instantiates a new ThirdPartyFileWrapper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponse

`func (o *ThirdPartyFileWrapper) GetResponse() ThirdPartyFileDto`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *ThirdPartyFileWrapper) GetResponseOk() (*ThirdPartyFileDto, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *ThirdPartyFileWrapper) SetResponse(v ThirdPartyFileDto)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *ThirdPartyFileWrapper) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetCount

`func (o *ThirdPartyFileWrapper) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ThirdPartyFileWrapper) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ThirdPartyFileWrapper) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ThirdPartyFileWrapper) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetLinks

`func (o *ThirdPartyFileWrapper) GetLinks() []GetPortalPrices200ResponseLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ThirdPartyFileWrapper) GetLinksOk() (*[]GetPortalPrices200ResponseLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ThirdPartyFileWrapper) SetLinks(v []GetPortalPrices200ResponseLinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ThirdPartyFileWrapper) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetStatus

`func (o *ThirdPartyFileWrapper) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ThirdPartyFileWrapper) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ThirdPartyFileWrapper) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ThirdPartyFileWrapper) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCode

`func (o *ThirdPartyFileWrapper) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ThirdPartyFileWrapper) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ThirdPartyFileWrapper) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *ThirdPartyFileWrapper) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


