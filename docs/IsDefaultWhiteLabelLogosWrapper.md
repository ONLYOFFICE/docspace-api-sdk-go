# IsDefaultWhiteLabelLogosWrapper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | Pointer to [**IsDefaultWhiteLabelLogosDto**](IsDefaultWhiteLabelLogosDto.md) | The IsDefaultWhiteLabelLogosDto object returned by the operation. | [optional] 
**Count** | Pointer to **int32** | How many items this response carries: the length of the page for a list, 1 for a single object | [optional] 
**Total** | Pointer to **int64** | How many items the whole list holds across all pages, for working out how many pages to request. Sent only by the operations that page their results, and left out of every other response | [optional] 
**Links** | Pointer to [**[]GetPortalPrices200ResponseLinksInner**](GetPortalPrices200ResponseLinksInner.md) | List of links related to the response | [optional] 
**Status** | Pointer to **int32** | HTTP status code of the response | [optional] 
**StatusCode** | Pointer to **int32** | HTTP status code of the response (duplicate of status) | [optional] 

## Methods

### NewIsDefaultWhiteLabelLogosWrapper

`func NewIsDefaultWhiteLabelLogosWrapper() *IsDefaultWhiteLabelLogosWrapper`

NewIsDefaultWhiteLabelLogosWrapper instantiates a new IsDefaultWhiteLabelLogosWrapper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIsDefaultWhiteLabelLogosWrapperWithDefaults

`func NewIsDefaultWhiteLabelLogosWrapperWithDefaults() *IsDefaultWhiteLabelLogosWrapper`

NewIsDefaultWhiteLabelLogosWrapperWithDefaults instantiates a new IsDefaultWhiteLabelLogosWrapper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponse

`func (o *IsDefaultWhiteLabelLogosWrapper) GetResponse() IsDefaultWhiteLabelLogosDto`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *IsDefaultWhiteLabelLogosWrapper) GetResponseOk() (*IsDefaultWhiteLabelLogosDto, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *IsDefaultWhiteLabelLogosWrapper) SetResponse(v IsDefaultWhiteLabelLogosDto)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *IsDefaultWhiteLabelLogosWrapper) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetCount

`func (o *IsDefaultWhiteLabelLogosWrapper) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *IsDefaultWhiteLabelLogosWrapper) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *IsDefaultWhiteLabelLogosWrapper) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *IsDefaultWhiteLabelLogosWrapper) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotal

`func (o *IsDefaultWhiteLabelLogosWrapper) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *IsDefaultWhiteLabelLogosWrapper) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *IsDefaultWhiteLabelLogosWrapper) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *IsDefaultWhiteLabelLogosWrapper) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetLinks

`func (o *IsDefaultWhiteLabelLogosWrapper) GetLinks() []GetPortalPrices200ResponseLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *IsDefaultWhiteLabelLogosWrapper) GetLinksOk() (*[]GetPortalPrices200ResponseLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *IsDefaultWhiteLabelLogosWrapper) SetLinks(v []GetPortalPrices200ResponseLinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *IsDefaultWhiteLabelLogosWrapper) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetStatus

`func (o *IsDefaultWhiteLabelLogosWrapper) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IsDefaultWhiteLabelLogosWrapper) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IsDefaultWhiteLabelLogosWrapper) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IsDefaultWhiteLabelLogosWrapper) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCode

`func (o *IsDefaultWhiteLabelLogosWrapper) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *IsDefaultWhiteLabelLogosWrapper) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *IsDefaultWhiteLabelLogosWrapper) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *IsDefaultWhiteLabelLogosWrapper) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


