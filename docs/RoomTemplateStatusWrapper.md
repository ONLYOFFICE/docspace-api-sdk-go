# RoomTemplateStatusWrapper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | Pointer to [**RoomTemplateStatusDto**](RoomTemplateStatusDto.md) | The RoomTemplateStatusDto object returned by the operation. | [optional] 
**Count** | Pointer to **int32** | The total number of items in the response | [optional] 
**Links** | Pointer to [**[]GetPortalPrices200ResponseLinksInner**](GetPortalPrices200ResponseLinksInner.md) | List of links related to the response | [optional] 
**Status** | Pointer to **int32** | HTTP status code of the response | [optional] 
**StatusCode** | Pointer to **int32** | HTTP status code of the response (duplicate of status) | [optional] 

## Methods

### NewRoomTemplateStatusWrapper

`func NewRoomTemplateStatusWrapper() *RoomTemplateStatusWrapper`

NewRoomTemplateStatusWrapper instantiates a new RoomTemplateStatusWrapper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoomTemplateStatusWrapperWithDefaults

`func NewRoomTemplateStatusWrapperWithDefaults() *RoomTemplateStatusWrapper`

NewRoomTemplateStatusWrapperWithDefaults instantiates a new RoomTemplateStatusWrapper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponse

`func (o *RoomTemplateStatusWrapper) GetResponse() RoomTemplateStatusDto`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *RoomTemplateStatusWrapper) GetResponseOk() (*RoomTemplateStatusDto, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *RoomTemplateStatusWrapper) SetResponse(v RoomTemplateStatusDto)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *RoomTemplateStatusWrapper) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetCount

`func (o *RoomTemplateStatusWrapper) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *RoomTemplateStatusWrapper) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *RoomTemplateStatusWrapper) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *RoomTemplateStatusWrapper) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetLinks

`func (o *RoomTemplateStatusWrapper) GetLinks() []GetPortalPrices200ResponseLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RoomTemplateStatusWrapper) GetLinksOk() (*[]GetPortalPrices200ResponseLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RoomTemplateStatusWrapper) SetLinks(v []GetPortalPrices200ResponseLinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RoomTemplateStatusWrapper) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetStatus

`func (o *RoomTemplateStatusWrapper) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RoomTemplateStatusWrapper) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RoomTemplateStatusWrapper) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RoomTemplateStatusWrapper) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCode

`func (o *RoomTemplateStatusWrapper) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *RoomTemplateStatusWrapper) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *RoomTemplateStatusWrapper) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *RoomTemplateStatusWrapper) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


