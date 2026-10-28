# CheckDocServiceUrlRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocServiceUrl** | **NullableString** | The public address of the Document Server, the one a browser loads the editor from. An empty value drops the  portal's own setting, so the address configured for the deployment takes over again. A value with no scheme is  stored with `http://` prepended, and an absolute address may not carry a query string. | 
**DocServiceUrlInternal** | Pointer to **NullableString** | The address the portal itself uses for its server-to-server calls to the Document Server, for deployments  where that traffic stays inside the private network. Left empty, those calls go to the public address instead. | [optional] 
**DocServiceUrlPortal** | Pointer to **NullableString** | The address of this portal as the Document Server has to call it back on in order to fetch and save a  document. Set it when the Document Server cannot resolve the portal by its public name; left empty, the  portal's own resolved address is used. | [optional] 
**DocServiceSignatureSecret** | Pointer to **NullableString** | The shared secret that requests between the portal and the Document Server are signed with; it has to be the  same value the Document Server itself is configured with, otherwise the verification of the new settings  fails. It is write-only: the document service location is reported without it. | [optional] 
**DocServiceSignatureHeader** | Pointer to **NullableString** | The name of the HTTP header the signature travels in, which has to match the header the Document Server  expects. A secret without a header is not a usable pair and is rejected. | [optional] 
**DocServiceSslVerification** | Pointer to **NullableBool** | Whether the portal validates the TLS certificate of the Document Server. With verification on, a self-signed  certificate breaks the connection; with it off, any certificate is accepted, which is meant for test  deployments only. Omitting the field turns verification on. | [optional] 

## Methods

### NewCheckDocServiceUrlRequestDto

`func NewCheckDocServiceUrlRequestDto(docServiceUrl NullableString, ) *CheckDocServiceUrlRequestDto`

NewCheckDocServiceUrlRequestDto instantiates a new CheckDocServiceUrlRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckDocServiceUrlRequestDtoWithDefaults

`func NewCheckDocServiceUrlRequestDtoWithDefaults() *CheckDocServiceUrlRequestDto`

NewCheckDocServiceUrlRequestDtoWithDefaults instantiates a new CheckDocServiceUrlRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocServiceUrl

`func (o *CheckDocServiceUrlRequestDto) GetDocServiceUrl() string`

GetDocServiceUrl returns the DocServiceUrl field if non-nil, zero value otherwise.

### GetDocServiceUrlOk

`func (o *CheckDocServiceUrlRequestDto) GetDocServiceUrlOk() (*string, bool)`

GetDocServiceUrlOk returns a tuple with the DocServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocServiceUrl

`func (o *CheckDocServiceUrlRequestDto) SetDocServiceUrl(v string)`

SetDocServiceUrl sets DocServiceUrl field to given value.


### SetDocServiceUrlNil

`func (o *CheckDocServiceUrlRequestDto) SetDocServiceUrlNil(b bool)`

 SetDocServiceUrlNil sets the value for DocServiceUrl to be an explicit nil

### UnsetDocServiceUrl
`func (o *CheckDocServiceUrlRequestDto) UnsetDocServiceUrl()`

UnsetDocServiceUrl ensures that no value is present for DocServiceUrl, not even an explicit nil
### GetDocServiceUrlInternal

`func (o *CheckDocServiceUrlRequestDto) GetDocServiceUrlInternal() string`

GetDocServiceUrlInternal returns the DocServiceUrlInternal field if non-nil, zero value otherwise.

### GetDocServiceUrlInternalOk

`func (o *CheckDocServiceUrlRequestDto) GetDocServiceUrlInternalOk() (*string, bool)`

GetDocServiceUrlInternalOk returns a tuple with the DocServiceUrlInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocServiceUrlInternal

`func (o *CheckDocServiceUrlRequestDto) SetDocServiceUrlInternal(v string)`

SetDocServiceUrlInternal sets DocServiceUrlInternal field to given value.

### HasDocServiceUrlInternal

`func (o *CheckDocServiceUrlRequestDto) HasDocServiceUrlInternal() bool`

HasDocServiceUrlInternal returns a boolean if a field has been set.

### SetDocServiceUrlInternalNil

`func (o *CheckDocServiceUrlRequestDto) SetDocServiceUrlInternalNil(b bool)`

 SetDocServiceUrlInternalNil sets the value for DocServiceUrlInternal to be an explicit nil

### UnsetDocServiceUrlInternal
`func (o *CheckDocServiceUrlRequestDto) UnsetDocServiceUrlInternal()`

UnsetDocServiceUrlInternal ensures that no value is present for DocServiceUrlInternal, not even an explicit nil
### GetDocServiceUrlPortal

`func (o *CheckDocServiceUrlRequestDto) GetDocServiceUrlPortal() string`

GetDocServiceUrlPortal returns the DocServiceUrlPortal field if non-nil, zero value otherwise.

### GetDocServiceUrlPortalOk

`func (o *CheckDocServiceUrlRequestDto) GetDocServiceUrlPortalOk() (*string, bool)`

GetDocServiceUrlPortalOk returns a tuple with the DocServiceUrlPortal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocServiceUrlPortal

`func (o *CheckDocServiceUrlRequestDto) SetDocServiceUrlPortal(v string)`

SetDocServiceUrlPortal sets DocServiceUrlPortal field to given value.

### HasDocServiceUrlPortal

`func (o *CheckDocServiceUrlRequestDto) HasDocServiceUrlPortal() bool`

HasDocServiceUrlPortal returns a boolean if a field has been set.

### SetDocServiceUrlPortalNil

`func (o *CheckDocServiceUrlRequestDto) SetDocServiceUrlPortalNil(b bool)`

 SetDocServiceUrlPortalNil sets the value for DocServiceUrlPortal to be an explicit nil

### UnsetDocServiceUrlPortal
`func (o *CheckDocServiceUrlRequestDto) UnsetDocServiceUrlPortal()`

UnsetDocServiceUrlPortal ensures that no value is present for DocServiceUrlPortal, not even an explicit nil
### GetDocServiceSignatureSecret

`func (o *CheckDocServiceUrlRequestDto) GetDocServiceSignatureSecret() string`

GetDocServiceSignatureSecret returns the DocServiceSignatureSecret field if non-nil, zero value otherwise.

### GetDocServiceSignatureSecretOk

`func (o *CheckDocServiceUrlRequestDto) GetDocServiceSignatureSecretOk() (*string, bool)`

GetDocServiceSignatureSecretOk returns a tuple with the DocServiceSignatureSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocServiceSignatureSecret

`func (o *CheckDocServiceUrlRequestDto) SetDocServiceSignatureSecret(v string)`

SetDocServiceSignatureSecret sets DocServiceSignatureSecret field to given value.

### HasDocServiceSignatureSecret

`func (o *CheckDocServiceUrlRequestDto) HasDocServiceSignatureSecret() bool`

HasDocServiceSignatureSecret returns a boolean if a field has been set.

### SetDocServiceSignatureSecretNil

`func (o *CheckDocServiceUrlRequestDto) SetDocServiceSignatureSecretNil(b bool)`

 SetDocServiceSignatureSecretNil sets the value for DocServiceSignatureSecret to be an explicit nil

### UnsetDocServiceSignatureSecret
`func (o *CheckDocServiceUrlRequestDto) UnsetDocServiceSignatureSecret()`

UnsetDocServiceSignatureSecret ensures that no value is present for DocServiceSignatureSecret, not even an explicit nil
### GetDocServiceSignatureHeader

`func (o *CheckDocServiceUrlRequestDto) GetDocServiceSignatureHeader() string`

GetDocServiceSignatureHeader returns the DocServiceSignatureHeader field if non-nil, zero value otherwise.

### GetDocServiceSignatureHeaderOk

`func (o *CheckDocServiceUrlRequestDto) GetDocServiceSignatureHeaderOk() (*string, bool)`

GetDocServiceSignatureHeaderOk returns a tuple with the DocServiceSignatureHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocServiceSignatureHeader

`func (o *CheckDocServiceUrlRequestDto) SetDocServiceSignatureHeader(v string)`

SetDocServiceSignatureHeader sets DocServiceSignatureHeader field to given value.

### HasDocServiceSignatureHeader

`func (o *CheckDocServiceUrlRequestDto) HasDocServiceSignatureHeader() bool`

HasDocServiceSignatureHeader returns a boolean if a field has been set.

### SetDocServiceSignatureHeaderNil

`func (o *CheckDocServiceUrlRequestDto) SetDocServiceSignatureHeaderNil(b bool)`

 SetDocServiceSignatureHeaderNil sets the value for DocServiceSignatureHeader to be an explicit nil

### UnsetDocServiceSignatureHeader
`func (o *CheckDocServiceUrlRequestDto) UnsetDocServiceSignatureHeader()`

UnsetDocServiceSignatureHeader ensures that no value is present for DocServiceSignatureHeader, not even an explicit nil
### GetDocServiceSslVerification

`func (o *CheckDocServiceUrlRequestDto) GetDocServiceSslVerification() bool`

GetDocServiceSslVerification returns the DocServiceSslVerification field if non-nil, zero value otherwise.

### GetDocServiceSslVerificationOk

`func (o *CheckDocServiceUrlRequestDto) GetDocServiceSslVerificationOk() (*bool, bool)`

GetDocServiceSslVerificationOk returns a tuple with the DocServiceSslVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocServiceSslVerification

`func (o *CheckDocServiceUrlRequestDto) SetDocServiceSslVerification(v bool)`

SetDocServiceSslVerification sets DocServiceSslVerification field to given value.

### HasDocServiceSslVerification

`func (o *CheckDocServiceUrlRequestDto) HasDocServiceSslVerification() bool`

HasDocServiceSslVerification returns a boolean if a field has been set.

### SetDocServiceSslVerificationNil

`func (o *CheckDocServiceUrlRequestDto) SetDocServiceSslVerificationNil(b bool)`

 SetDocServiceSslVerificationNil sets the value for DocServiceSslVerification to be an explicit nil

### UnsetDocServiceSslVerification
`func (o *CheckDocServiceUrlRequestDto) UnsetDocServiceSslVerification()`

UnsetDocServiceSslVerification ensures that no value is present for DocServiceSslVerification, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


