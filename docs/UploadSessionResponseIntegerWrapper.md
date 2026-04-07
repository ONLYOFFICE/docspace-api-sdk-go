# UploadSessionResponseIntegerWrapper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | Pointer to [**UploadSessionResponseDtoInteger**](UploadSessionResponseDtoInteger.md) |  | [optional] 
**Count** | Pointer to **int32** | The total number of items in the response | [optional] 
**Links** | Pointer to [**[]GetPortalPrices200ResponseLinksInner**](GetPortalPrices200ResponseLinksInner.md) | List of links related to the response | [optional] 
**Status** | Pointer to **int32** | HTTP status code of the response | [optional] 
**StatusCode** | Pointer to **int32** | HTTP status code of the response (duplicate of status) | [optional] 

## Methods

### NewUploadSessionResponseIntegerWrapper

`func NewUploadSessionResponseIntegerWrapper() *UploadSessionResponseIntegerWrapper`

NewUploadSessionResponseIntegerWrapper instantiates a new UploadSessionResponseIntegerWrapper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadSessionResponseIntegerWrapperWithDefaults

`func NewUploadSessionResponseIntegerWrapperWithDefaults() *UploadSessionResponseIntegerWrapper`

NewUploadSessionResponseIntegerWrapperWithDefaults instantiates a new UploadSessionResponseIntegerWrapper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponse

`func (o *UploadSessionResponseIntegerWrapper) GetResponse() UploadSessionResponseDtoInteger`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *UploadSessionResponseIntegerWrapper) GetResponseOk() (*UploadSessionResponseDtoInteger, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *UploadSessionResponseIntegerWrapper) SetResponse(v UploadSessionResponseDtoInteger)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *UploadSessionResponseIntegerWrapper) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetCount

`func (o *UploadSessionResponseIntegerWrapper) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UploadSessionResponseIntegerWrapper) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UploadSessionResponseIntegerWrapper) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *UploadSessionResponseIntegerWrapper) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetLinks

`func (o *UploadSessionResponseIntegerWrapper) GetLinks() []GetPortalPrices200ResponseLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UploadSessionResponseIntegerWrapper) GetLinksOk() (*[]GetPortalPrices200ResponseLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UploadSessionResponseIntegerWrapper) SetLinks(v []GetPortalPrices200ResponseLinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UploadSessionResponseIntegerWrapper) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetStatus

`func (o *UploadSessionResponseIntegerWrapper) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UploadSessionResponseIntegerWrapper) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UploadSessionResponseIntegerWrapper) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UploadSessionResponseIntegerWrapper) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCode

`func (o *UploadSessionResponseIntegerWrapper) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *UploadSessionResponseIntegerWrapper) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *UploadSessionResponseIntegerWrapper) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *UploadSessionResponseIntegerWrapper) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


