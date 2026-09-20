# EmbeddedConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmbedUrl** | Pointer to **NullableString** | The page to put into the frame. It is empty when the opening carries no external share key, since a framed  viewer cannot authenticate a portal member. | [optional] 
**SaveUrl** | Pointer to **NullableString** | Where the download button of the framed viewer leads. | [optional] [readonly] 
**ShareLinkParam** | Pointer to **NullableString** | The query fragment carrying the external share key, ampersand included, out of which the addresses around it  are built. | [optional] 
**ShareUrl** | Pointer to **NullableString** | The address behind the share button of the framed viewer, the document opened full-screen for reading. It is  empty when the opening carries no external share key. | [optional] 
**ToolbarDocked** | Pointer to **NullableString** | Where the framed viewer puts its toolbar. The portal always asks for the top. | [optional] [readonly] 

## Methods

### NewEmbeddedConfig

`func NewEmbeddedConfig() *EmbeddedConfig`

NewEmbeddedConfig instantiates a new EmbeddedConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbeddedConfigWithDefaults

`func NewEmbeddedConfigWithDefaults() *EmbeddedConfig`

NewEmbeddedConfigWithDefaults instantiates a new EmbeddedConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmbedUrl

`func (o *EmbeddedConfig) GetEmbedUrl() string`

GetEmbedUrl returns the EmbedUrl field if non-nil, zero value otherwise.

### GetEmbedUrlOk

`func (o *EmbeddedConfig) GetEmbedUrlOk() (*string, bool)`

GetEmbedUrlOk returns a tuple with the EmbedUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedUrl

`func (o *EmbeddedConfig) SetEmbedUrl(v string)`

SetEmbedUrl sets EmbedUrl field to given value.

### HasEmbedUrl

`func (o *EmbeddedConfig) HasEmbedUrl() bool`

HasEmbedUrl returns a boolean if a field has been set.

### SetEmbedUrlNil

`func (o *EmbeddedConfig) SetEmbedUrlNil(b bool)`

 SetEmbedUrlNil sets the value for EmbedUrl to be an explicit nil

### UnsetEmbedUrl
`func (o *EmbeddedConfig) UnsetEmbedUrl()`

UnsetEmbedUrl ensures that no value is present for EmbedUrl, not even an explicit nil
### GetSaveUrl

`func (o *EmbeddedConfig) GetSaveUrl() string`

GetSaveUrl returns the SaveUrl field if non-nil, zero value otherwise.

### GetSaveUrlOk

`func (o *EmbeddedConfig) GetSaveUrlOk() (*string, bool)`

GetSaveUrlOk returns a tuple with the SaveUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveUrl

`func (o *EmbeddedConfig) SetSaveUrl(v string)`

SetSaveUrl sets SaveUrl field to given value.

### HasSaveUrl

`func (o *EmbeddedConfig) HasSaveUrl() bool`

HasSaveUrl returns a boolean if a field has been set.

### SetSaveUrlNil

`func (o *EmbeddedConfig) SetSaveUrlNil(b bool)`

 SetSaveUrlNil sets the value for SaveUrl to be an explicit nil

### UnsetSaveUrl
`func (o *EmbeddedConfig) UnsetSaveUrl()`

UnsetSaveUrl ensures that no value is present for SaveUrl, not even an explicit nil
### GetShareLinkParam

`func (o *EmbeddedConfig) GetShareLinkParam() string`

GetShareLinkParam returns the ShareLinkParam field if non-nil, zero value otherwise.

### GetShareLinkParamOk

`func (o *EmbeddedConfig) GetShareLinkParamOk() (*string, bool)`

GetShareLinkParamOk returns a tuple with the ShareLinkParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareLinkParam

`func (o *EmbeddedConfig) SetShareLinkParam(v string)`

SetShareLinkParam sets ShareLinkParam field to given value.

### HasShareLinkParam

`func (o *EmbeddedConfig) HasShareLinkParam() bool`

HasShareLinkParam returns a boolean if a field has been set.

### SetShareLinkParamNil

`func (o *EmbeddedConfig) SetShareLinkParamNil(b bool)`

 SetShareLinkParamNil sets the value for ShareLinkParam to be an explicit nil

### UnsetShareLinkParam
`func (o *EmbeddedConfig) UnsetShareLinkParam()`

UnsetShareLinkParam ensures that no value is present for ShareLinkParam, not even an explicit nil
### GetShareUrl

`func (o *EmbeddedConfig) GetShareUrl() string`

GetShareUrl returns the ShareUrl field if non-nil, zero value otherwise.

### GetShareUrlOk

`func (o *EmbeddedConfig) GetShareUrlOk() (*string, bool)`

GetShareUrlOk returns a tuple with the ShareUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareUrl

`func (o *EmbeddedConfig) SetShareUrl(v string)`

SetShareUrl sets ShareUrl field to given value.

### HasShareUrl

`func (o *EmbeddedConfig) HasShareUrl() bool`

HasShareUrl returns a boolean if a field has been set.

### SetShareUrlNil

`func (o *EmbeddedConfig) SetShareUrlNil(b bool)`

 SetShareUrlNil sets the value for ShareUrl to be an explicit nil

### UnsetShareUrl
`func (o *EmbeddedConfig) UnsetShareUrl()`

UnsetShareUrl ensures that no value is present for ShareUrl, not even an explicit nil
### GetToolbarDocked

`func (o *EmbeddedConfig) GetToolbarDocked() string`

GetToolbarDocked returns the ToolbarDocked field if non-nil, zero value otherwise.

### GetToolbarDockedOk

`func (o *EmbeddedConfig) GetToolbarDockedOk() (*string, bool)`

GetToolbarDockedOk returns a tuple with the ToolbarDocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolbarDocked

`func (o *EmbeddedConfig) SetToolbarDocked(v string)`

SetToolbarDocked sets ToolbarDocked field to given value.

### HasToolbarDocked

`func (o *EmbeddedConfig) HasToolbarDocked() bool`

HasToolbarDocked returns a boolean if a field has been set.

### SetToolbarDockedNil

`func (o *EmbeddedConfig) SetToolbarDockedNil(b bool)`

 SetToolbarDockedNil sets the value for ToolbarDocked to be an explicit nil

### UnsetToolbarDocked
`func (o *EmbeddedConfig) UnsetToolbarDocked()`

UnsetToolbarDocked ensures that no value is present for ToolbarDocked, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


