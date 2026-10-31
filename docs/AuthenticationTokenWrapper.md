# AuthenticationTokenWrapper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | Pointer to [**AuthenticationTokenDto**](AuthenticationTokenDto.md) | The AuthenticationTokenDto object returned by the operation. | [optional] 
**Count** | Pointer to **int32** | How many items this response carries: the length of the page for a list, 1 for a single object | [optional] 
**Total** | Pointer to **int64** | How many items the whole list holds across all pages, for working out how many pages to request. Sent only by the operations that page their results, and left out of every other response | [optional] 
**Links** | Pointer to [**[]GetPortalPrices200ResponseLinksInner**](GetPortalPrices200ResponseLinksInner.md) | List of links related to the response | [optional] 
**Status** | Pointer to **int32** | HTTP status code of the response | [optional] 
**StatusCode** | Pointer to **int32** | HTTP status code of the response (duplicate of status) | [optional] 

## Methods

### NewAuthenticationTokenWrapper

`func NewAuthenticationTokenWrapper() *AuthenticationTokenWrapper`

NewAuthenticationTokenWrapper instantiates a new AuthenticationTokenWrapper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationTokenWrapperWithDefaults

`func NewAuthenticationTokenWrapperWithDefaults() *AuthenticationTokenWrapper`

NewAuthenticationTokenWrapperWithDefaults instantiates a new AuthenticationTokenWrapper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponse

`func (o *AuthenticationTokenWrapper) GetResponse() AuthenticationTokenDto`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *AuthenticationTokenWrapper) GetResponseOk() (*AuthenticationTokenDto, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *AuthenticationTokenWrapper) SetResponse(v AuthenticationTokenDto)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *AuthenticationTokenWrapper) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetCount

`func (o *AuthenticationTokenWrapper) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AuthenticationTokenWrapper) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AuthenticationTokenWrapper) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *AuthenticationTokenWrapper) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotal

`func (o *AuthenticationTokenWrapper) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AuthenticationTokenWrapper) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AuthenticationTokenWrapper) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *AuthenticationTokenWrapper) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetLinks

`func (o *AuthenticationTokenWrapper) GetLinks() []GetPortalPrices200ResponseLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AuthenticationTokenWrapper) GetLinksOk() (*[]GetPortalPrices200ResponseLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AuthenticationTokenWrapper) SetLinks(v []GetPortalPrices200ResponseLinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AuthenticationTokenWrapper) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetStatus

`func (o *AuthenticationTokenWrapper) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AuthenticationTokenWrapper) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AuthenticationTokenWrapper) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AuthenticationTokenWrapper) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCode

`func (o *AuthenticationTokenWrapper) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *AuthenticationTokenWrapper) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *AuthenticationTokenWrapper) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *AuthenticationTokenWrapper) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


