# FirebaseDeviceWrapper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | Pointer to [**FirebaseDeviceDto**](FirebaseDeviceDto.md) | The FirebaseDeviceDto object returned by the operation. | [optional] 
**Count** | Pointer to **int32** | How many items this response carries: the length of the page for a list, 1 for a single object | [optional] 
**Total** | Pointer to **int64** | How many items the whole list holds across all pages, for working out how many pages to request. Sent only by the operations that page their results, and left out of every other response | [optional] 
**Links** | Pointer to [**[]GetPortalPrices200ResponseLinksInner**](GetPortalPrices200ResponseLinksInner.md) | List of links related to the response | [optional] 
**Status** | Pointer to **int32** | HTTP status code of the response | [optional] 
**StatusCode** | Pointer to **int32** | HTTP status code of the response (duplicate of status) | [optional] 

## Methods

### NewFirebaseDeviceWrapper

`func NewFirebaseDeviceWrapper() *FirebaseDeviceWrapper`

NewFirebaseDeviceWrapper instantiates a new FirebaseDeviceWrapper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirebaseDeviceWrapperWithDefaults

`func NewFirebaseDeviceWrapperWithDefaults() *FirebaseDeviceWrapper`

NewFirebaseDeviceWrapperWithDefaults instantiates a new FirebaseDeviceWrapper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponse

`func (o *FirebaseDeviceWrapper) GetResponse() FirebaseDeviceDto`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *FirebaseDeviceWrapper) GetResponseOk() (*FirebaseDeviceDto, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *FirebaseDeviceWrapper) SetResponse(v FirebaseDeviceDto)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *FirebaseDeviceWrapper) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetCount

`func (o *FirebaseDeviceWrapper) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *FirebaseDeviceWrapper) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *FirebaseDeviceWrapper) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *FirebaseDeviceWrapper) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotal

`func (o *FirebaseDeviceWrapper) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *FirebaseDeviceWrapper) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *FirebaseDeviceWrapper) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *FirebaseDeviceWrapper) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetLinks

`func (o *FirebaseDeviceWrapper) GetLinks() []GetPortalPrices200ResponseLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FirebaseDeviceWrapper) GetLinksOk() (*[]GetPortalPrices200ResponseLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FirebaseDeviceWrapper) SetLinks(v []GetPortalPrices200ResponseLinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *FirebaseDeviceWrapper) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetStatus

`func (o *FirebaseDeviceWrapper) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FirebaseDeviceWrapper) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FirebaseDeviceWrapper) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FirebaseDeviceWrapper) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCode

`func (o *FirebaseDeviceWrapper) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *FirebaseDeviceWrapper) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *FirebaseDeviceWrapper) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *FirebaseDeviceWrapper) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


