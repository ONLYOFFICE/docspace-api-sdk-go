# DocServiceUrlDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **NullableString** | The version of the document service. | 
**DocServiceUrlApi** | **NullableString** | The document service URL API. | 
**DocServiceUrl** | **NullableString** | The document service URL. | 
**DocServicePreloadUrl** | **NullableString** | The URL used to preload the document service scripts. | 
**DocServiceUrlInternal** | **NullableString** | The internal document service URL. | 
**DocServicePortalUrl** | **NullableString** | The document service portal URL. | 
**DocServiceSignatureHeader** | **NullableString** | The document service signature header. | 
**DocServiceSslVerification** | **bool** | Specifies if the document service SSL verification is enabled. | 
**IsDefault** | **bool** | Specifies if the document service is default. | 

## Methods

### NewDocServiceUrlDto

`func NewDocServiceUrlDto(version NullableString, docServiceUrlApi NullableString, docServiceUrl NullableString, docServicePreloadUrl NullableString, docServiceUrlInternal NullableString, docServicePortalUrl NullableString, docServiceSignatureHeader NullableString, docServiceSslVerification bool, isDefault bool, ) *DocServiceUrlDto`

NewDocServiceUrlDto instantiates a new DocServiceUrlDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocServiceUrlDtoWithDefaults

`func NewDocServiceUrlDtoWithDefaults() *DocServiceUrlDto`

NewDocServiceUrlDtoWithDefaults instantiates a new DocServiceUrlDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *DocServiceUrlDto) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DocServiceUrlDto) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DocServiceUrlDto) SetVersion(v string)`

SetVersion sets Version field to given value.


### SetVersionNil

`func (o *DocServiceUrlDto) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *DocServiceUrlDto) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetDocServiceUrlApi

`func (o *DocServiceUrlDto) GetDocServiceUrlApi() string`

GetDocServiceUrlApi returns the DocServiceUrlApi field if non-nil, zero value otherwise.

### GetDocServiceUrlApiOk

`func (o *DocServiceUrlDto) GetDocServiceUrlApiOk() (*string, bool)`

GetDocServiceUrlApiOk returns a tuple with the DocServiceUrlApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocServiceUrlApi

`func (o *DocServiceUrlDto) SetDocServiceUrlApi(v string)`

SetDocServiceUrlApi sets DocServiceUrlApi field to given value.


### SetDocServiceUrlApiNil

`func (o *DocServiceUrlDto) SetDocServiceUrlApiNil(b bool)`

 SetDocServiceUrlApiNil sets the value for DocServiceUrlApi to be an explicit nil

### UnsetDocServiceUrlApi
`func (o *DocServiceUrlDto) UnsetDocServiceUrlApi()`

UnsetDocServiceUrlApi ensures that no value is present for DocServiceUrlApi, not even an explicit nil
### GetDocServiceUrl

`func (o *DocServiceUrlDto) GetDocServiceUrl() string`

GetDocServiceUrl returns the DocServiceUrl field if non-nil, zero value otherwise.

### GetDocServiceUrlOk

`func (o *DocServiceUrlDto) GetDocServiceUrlOk() (*string, bool)`

GetDocServiceUrlOk returns a tuple with the DocServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocServiceUrl

`func (o *DocServiceUrlDto) SetDocServiceUrl(v string)`

SetDocServiceUrl sets DocServiceUrl field to given value.


### SetDocServiceUrlNil

`func (o *DocServiceUrlDto) SetDocServiceUrlNil(b bool)`

 SetDocServiceUrlNil sets the value for DocServiceUrl to be an explicit nil

### UnsetDocServiceUrl
`func (o *DocServiceUrlDto) UnsetDocServiceUrl()`

UnsetDocServiceUrl ensures that no value is present for DocServiceUrl, not even an explicit nil
### GetDocServicePreloadUrl

`func (o *DocServiceUrlDto) GetDocServicePreloadUrl() string`

GetDocServicePreloadUrl returns the DocServicePreloadUrl field if non-nil, zero value otherwise.

### GetDocServicePreloadUrlOk

`func (o *DocServiceUrlDto) GetDocServicePreloadUrlOk() (*string, bool)`

GetDocServicePreloadUrlOk returns a tuple with the DocServicePreloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocServicePreloadUrl

`func (o *DocServiceUrlDto) SetDocServicePreloadUrl(v string)`

SetDocServicePreloadUrl sets DocServicePreloadUrl field to given value.


### SetDocServicePreloadUrlNil

`func (o *DocServiceUrlDto) SetDocServicePreloadUrlNil(b bool)`

 SetDocServicePreloadUrlNil sets the value for DocServicePreloadUrl to be an explicit nil

### UnsetDocServicePreloadUrl
`func (o *DocServiceUrlDto) UnsetDocServicePreloadUrl()`

UnsetDocServicePreloadUrl ensures that no value is present for DocServicePreloadUrl, not even an explicit nil
### GetDocServiceUrlInternal

`func (o *DocServiceUrlDto) GetDocServiceUrlInternal() string`

GetDocServiceUrlInternal returns the DocServiceUrlInternal field if non-nil, zero value otherwise.

### GetDocServiceUrlInternalOk

`func (o *DocServiceUrlDto) GetDocServiceUrlInternalOk() (*string, bool)`

GetDocServiceUrlInternalOk returns a tuple with the DocServiceUrlInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocServiceUrlInternal

`func (o *DocServiceUrlDto) SetDocServiceUrlInternal(v string)`

SetDocServiceUrlInternal sets DocServiceUrlInternal field to given value.


### SetDocServiceUrlInternalNil

`func (o *DocServiceUrlDto) SetDocServiceUrlInternalNil(b bool)`

 SetDocServiceUrlInternalNil sets the value for DocServiceUrlInternal to be an explicit nil

### UnsetDocServiceUrlInternal
`func (o *DocServiceUrlDto) UnsetDocServiceUrlInternal()`

UnsetDocServiceUrlInternal ensures that no value is present for DocServiceUrlInternal, not even an explicit nil
### GetDocServicePortalUrl

`func (o *DocServiceUrlDto) GetDocServicePortalUrl() string`

GetDocServicePortalUrl returns the DocServicePortalUrl field if non-nil, zero value otherwise.

### GetDocServicePortalUrlOk

`func (o *DocServiceUrlDto) GetDocServicePortalUrlOk() (*string, bool)`

GetDocServicePortalUrlOk returns a tuple with the DocServicePortalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocServicePortalUrl

`func (o *DocServiceUrlDto) SetDocServicePortalUrl(v string)`

SetDocServicePortalUrl sets DocServicePortalUrl field to given value.


### SetDocServicePortalUrlNil

`func (o *DocServiceUrlDto) SetDocServicePortalUrlNil(b bool)`

 SetDocServicePortalUrlNil sets the value for DocServicePortalUrl to be an explicit nil

### UnsetDocServicePortalUrl
`func (o *DocServiceUrlDto) UnsetDocServicePortalUrl()`

UnsetDocServicePortalUrl ensures that no value is present for DocServicePortalUrl, not even an explicit nil
### GetDocServiceSignatureHeader

`func (o *DocServiceUrlDto) GetDocServiceSignatureHeader() string`

GetDocServiceSignatureHeader returns the DocServiceSignatureHeader field if non-nil, zero value otherwise.

### GetDocServiceSignatureHeaderOk

`func (o *DocServiceUrlDto) GetDocServiceSignatureHeaderOk() (*string, bool)`

GetDocServiceSignatureHeaderOk returns a tuple with the DocServiceSignatureHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocServiceSignatureHeader

`func (o *DocServiceUrlDto) SetDocServiceSignatureHeader(v string)`

SetDocServiceSignatureHeader sets DocServiceSignatureHeader field to given value.


### SetDocServiceSignatureHeaderNil

`func (o *DocServiceUrlDto) SetDocServiceSignatureHeaderNil(b bool)`

 SetDocServiceSignatureHeaderNil sets the value for DocServiceSignatureHeader to be an explicit nil

### UnsetDocServiceSignatureHeader
`func (o *DocServiceUrlDto) UnsetDocServiceSignatureHeader()`

UnsetDocServiceSignatureHeader ensures that no value is present for DocServiceSignatureHeader, not even an explicit nil
### GetDocServiceSslVerification

`func (o *DocServiceUrlDto) GetDocServiceSslVerification() bool`

GetDocServiceSslVerification returns the DocServiceSslVerification field if non-nil, zero value otherwise.

### GetDocServiceSslVerificationOk

`func (o *DocServiceUrlDto) GetDocServiceSslVerificationOk() (*bool, bool)`

GetDocServiceSslVerificationOk returns a tuple with the DocServiceSslVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocServiceSslVerification

`func (o *DocServiceUrlDto) SetDocServiceSslVerification(v bool)`

SetDocServiceSslVerification sets DocServiceSslVerification field to given value.


### GetIsDefault

`func (o *DocServiceUrlDto) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *DocServiceUrlDto) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *DocServiceUrlDto) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


