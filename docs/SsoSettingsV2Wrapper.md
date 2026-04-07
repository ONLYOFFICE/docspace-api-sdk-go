# SsoSettingsV2Wrapper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | Pointer to [**SsoSettingsV2**](SsoSettingsV2.md) |  | [optional] 
**Count** | Pointer to **int32** | The total number of items in the response | [optional] 
**Links** | Pointer to [**[]GetPortalPrices200ResponseLinksInner**](GetPortalPrices200ResponseLinksInner.md) | List of links related to the response | [optional] 
**Status** | Pointer to **int32** | HTTP status code of the response | [optional] 
**StatusCode** | Pointer to **int32** | HTTP status code of the response (duplicate of status) | [optional] 

## Methods

### NewSsoSettingsV2Wrapper

`func NewSsoSettingsV2Wrapper() *SsoSettingsV2Wrapper`

NewSsoSettingsV2Wrapper instantiates a new SsoSettingsV2Wrapper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoSettingsV2WrapperWithDefaults

`func NewSsoSettingsV2WrapperWithDefaults() *SsoSettingsV2Wrapper`

NewSsoSettingsV2WrapperWithDefaults instantiates a new SsoSettingsV2Wrapper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponse

`func (o *SsoSettingsV2Wrapper) GetResponse() SsoSettingsV2`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *SsoSettingsV2Wrapper) GetResponseOk() (*SsoSettingsV2, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *SsoSettingsV2Wrapper) SetResponse(v SsoSettingsV2)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *SsoSettingsV2Wrapper) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetCount

`func (o *SsoSettingsV2Wrapper) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SsoSettingsV2Wrapper) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SsoSettingsV2Wrapper) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *SsoSettingsV2Wrapper) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetLinks

`func (o *SsoSettingsV2Wrapper) GetLinks() []GetPortalPrices200ResponseLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SsoSettingsV2Wrapper) GetLinksOk() (*[]GetPortalPrices200ResponseLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SsoSettingsV2Wrapper) SetLinks(v []GetPortalPrices200ResponseLinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SsoSettingsV2Wrapper) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetStatus

`func (o *SsoSettingsV2Wrapper) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SsoSettingsV2Wrapper) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SsoSettingsV2Wrapper) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SsoSettingsV2Wrapper) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCode

`func (o *SsoSettingsV2Wrapper) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *SsoSettingsV2Wrapper) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *SsoSettingsV2Wrapper) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *SsoSettingsV2Wrapper) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


