# DocsCloudTenantWrapper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | Pointer to [**DocsCloudTenant**](DocsCloudTenant.md) | The DocsCloudTenant object returned by the operation. | [optional] 
**Count** | Pointer to **int32** | The total number of items in the response | [optional] 
**Links** | Pointer to [**[]GetPortalPrices200ResponseLinksInner**](GetPortalPrices200ResponseLinksInner.md) | List of links related to the response | [optional] 
**Status** | Pointer to **int32** | HTTP status code of the response | [optional] 
**StatusCode** | Pointer to **int32** | HTTP status code of the response (duplicate of status) | [optional] 

## Methods

### NewDocsCloudTenantWrapper

`func NewDocsCloudTenantWrapper() *DocsCloudTenantWrapper`

NewDocsCloudTenantWrapper instantiates a new DocsCloudTenantWrapper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocsCloudTenantWrapperWithDefaults

`func NewDocsCloudTenantWrapperWithDefaults() *DocsCloudTenantWrapper`

NewDocsCloudTenantWrapperWithDefaults instantiates a new DocsCloudTenantWrapper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponse

`func (o *DocsCloudTenantWrapper) GetResponse() DocsCloudTenant`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *DocsCloudTenantWrapper) GetResponseOk() (*DocsCloudTenant, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *DocsCloudTenantWrapper) SetResponse(v DocsCloudTenant)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *DocsCloudTenantWrapper) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetCount

`func (o *DocsCloudTenantWrapper) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DocsCloudTenantWrapper) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DocsCloudTenantWrapper) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *DocsCloudTenantWrapper) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetLinks

`func (o *DocsCloudTenantWrapper) GetLinks() []GetPortalPrices200ResponseLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DocsCloudTenantWrapper) GetLinksOk() (*[]GetPortalPrices200ResponseLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DocsCloudTenantWrapper) SetLinks(v []GetPortalPrices200ResponseLinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DocsCloudTenantWrapper) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetStatus

`func (o *DocsCloudTenantWrapper) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DocsCloudTenantWrapper) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DocsCloudTenantWrapper) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DocsCloudTenantWrapper) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCode

`func (o *DocsCloudTenantWrapper) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *DocsCloudTenantWrapper) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *DocsCloudTenantWrapper) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *DocsCloudTenantWrapper) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


