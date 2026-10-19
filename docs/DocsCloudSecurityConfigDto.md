# DocsCloudSecurityConfigDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secret** | Pointer to **NullableString** | The security secret. | [optional] 
**Header** | Pointer to **NullableString** | The security header name. | [optional] 

## Methods

### NewDocsCloudSecurityConfigDto

`func NewDocsCloudSecurityConfigDto() *DocsCloudSecurityConfigDto`

NewDocsCloudSecurityConfigDto instantiates a new DocsCloudSecurityConfigDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocsCloudSecurityConfigDtoWithDefaults

`func NewDocsCloudSecurityConfigDtoWithDefaults() *DocsCloudSecurityConfigDto`

NewDocsCloudSecurityConfigDtoWithDefaults instantiates a new DocsCloudSecurityConfigDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecret

`func (o *DocsCloudSecurityConfigDto) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *DocsCloudSecurityConfigDto) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *DocsCloudSecurityConfigDto) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *DocsCloudSecurityConfigDto) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### SetSecretNil

`func (o *DocsCloudSecurityConfigDto) SetSecretNil(b bool)`

 SetSecretNil sets the value for Secret to be an explicit nil

### UnsetSecret
`func (o *DocsCloudSecurityConfigDto) UnsetSecret()`

UnsetSecret ensures that no value is present for Secret, not even an explicit nil
### GetHeader

`func (o *DocsCloudSecurityConfigDto) GetHeader() string`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *DocsCloudSecurityConfigDto) GetHeaderOk() (*string, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *DocsCloudSecurityConfigDto) SetHeader(v string)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *DocsCloudSecurityConfigDto) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### SetHeaderNil

`func (o *DocsCloudSecurityConfigDto) SetHeaderNil(b bool)`

 SetHeaderNil sets the value for Header to be an explicit nil

### UnsetHeader
`func (o *DocsCloudSecurityConfigDto) UnsetHeader()`

UnsetHeader ensures that no value is present for Header, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


