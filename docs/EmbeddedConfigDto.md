# EmbeddedConfigDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmbedUrl** | Pointer to **NullableString** | The page to put into the frame. It is empty when the opening carries no external share key, since a framed  viewer cannot authenticate a portal member. | [optional] 
**SaveUrl** | Pointer to **NullableString** | Where the download button of the framed viewer leads. | [optional] [readonly] 
**ShareLinkParam** | Pointer to **NullableString** | The query fragment carrying the external share key, ampersand included, out of which the addresses around it  are built. | [optional] 
**ShareUrl** | Pointer to **NullableString** | The address behind the share button of the framed viewer, the document opened full-screen for reading. It is  empty when the opening carries no external share key. | [optional] 
**ToolbarDocked** | Pointer to **NullableString** | Where the framed viewer puts its toolbar. The portal always asks for the top. | [optional] [readonly] 

## Methods

### NewEmbeddedConfigDto

`func NewEmbeddedConfigDto() *EmbeddedConfigDto`

NewEmbeddedConfigDto instantiates a new EmbeddedConfigDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbeddedConfigDtoWithDefaults

`func NewEmbeddedConfigDtoWithDefaults() *EmbeddedConfigDto`

NewEmbeddedConfigDtoWithDefaults instantiates a new EmbeddedConfigDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmbedUrl

`func (o *EmbeddedConfigDto) GetEmbedUrl() string`

GetEmbedUrl returns the EmbedUrl field if non-nil, zero value otherwise.

### GetEmbedUrlOk

`func (o *EmbeddedConfigDto) GetEmbedUrlOk() (*string, bool)`

GetEmbedUrlOk returns a tuple with the EmbedUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedUrl

`func (o *EmbeddedConfigDto) SetEmbedUrl(v string)`

SetEmbedUrl sets EmbedUrl field to given value.

### HasEmbedUrl

`func (o *EmbeddedConfigDto) HasEmbedUrl() bool`

HasEmbedUrl returns a boolean if a field has been set.

### SetEmbedUrlNil

`func (o *EmbeddedConfigDto) SetEmbedUrlNil(b bool)`

 SetEmbedUrlNil sets the value for EmbedUrl to be an explicit nil

### UnsetEmbedUrl
`func (o *EmbeddedConfigDto) UnsetEmbedUrl()`

UnsetEmbedUrl ensures that no value is present for EmbedUrl, not even an explicit nil
### GetSaveUrl

`func (o *EmbeddedConfigDto) GetSaveUrl() string`

GetSaveUrl returns the SaveUrl field if non-nil, zero value otherwise.

### GetSaveUrlOk

`func (o *EmbeddedConfigDto) GetSaveUrlOk() (*string, bool)`

GetSaveUrlOk returns a tuple with the SaveUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveUrl

`func (o *EmbeddedConfigDto) SetSaveUrl(v string)`

SetSaveUrl sets SaveUrl field to given value.

### HasSaveUrl

`func (o *EmbeddedConfigDto) HasSaveUrl() bool`

HasSaveUrl returns a boolean if a field has been set.

### SetSaveUrlNil

`func (o *EmbeddedConfigDto) SetSaveUrlNil(b bool)`

 SetSaveUrlNil sets the value for SaveUrl to be an explicit nil

### UnsetSaveUrl
`func (o *EmbeddedConfigDto) UnsetSaveUrl()`

UnsetSaveUrl ensures that no value is present for SaveUrl, not even an explicit nil
### GetShareLinkParam

`func (o *EmbeddedConfigDto) GetShareLinkParam() string`

GetShareLinkParam returns the ShareLinkParam field if non-nil, zero value otherwise.

### GetShareLinkParamOk

`func (o *EmbeddedConfigDto) GetShareLinkParamOk() (*string, bool)`

GetShareLinkParamOk returns a tuple with the ShareLinkParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareLinkParam

`func (o *EmbeddedConfigDto) SetShareLinkParam(v string)`

SetShareLinkParam sets ShareLinkParam field to given value.

### HasShareLinkParam

`func (o *EmbeddedConfigDto) HasShareLinkParam() bool`

HasShareLinkParam returns a boolean if a field has been set.

### SetShareLinkParamNil

`func (o *EmbeddedConfigDto) SetShareLinkParamNil(b bool)`

 SetShareLinkParamNil sets the value for ShareLinkParam to be an explicit nil

### UnsetShareLinkParam
`func (o *EmbeddedConfigDto) UnsetShareLinkParam()`

UnsetShareLinkParam ensures that no value is present for ShareLinkParam, not even an explicit nil
### GetShareUrl

`func (o *EmbeddedConfigDto) GetShareUrl() string`

GetShareUrl returns the ShareUrl field if non-nil, zero value otherwise.

### GetShareUrlOk

`func (o *EmbeddedConfigDto) GetShareUrlOk() (*string, bool)`

GetShareUrlOk returns a tuple with the ShareUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareUrl

`func (o *EmbeddedConfigDto) SetShareUrl(v string)`

SetShareUrl sets ShareUrl field to given value.

### HasShareUrl

`func (o *EmbeddedConfigDto) HasShareUrl() bool`

HasShareUrl returns a boolean if a field has been set.

### SetShareUrlNil

`func (o *EmbeddedConfigDto) SetShareUrlNil(b bool)`

 SetShareUrlNil sets the value for ShareUrl to be an explicit nil

### UnsetShareUrl
`func (o *EmbeddedConfigDto) UnsetShareUrl()`

UnsetShareUrl ensures that no value is present for ShareUrl, not even an explicit nil
### GetToolbarDocked

`func (o *EmbeddedConfigDto) GetToolbarDocked() string`

GetToolbarDocked returns the ToolbarDocked field if non-nil, zero value otherwise.

### GetToolbarDockedOk

`func (o *EmbeddedConfigDto) GetToolbarDockedOk() (*string, bool)`

GetToolbarDockedOk returns a tuple with the ToolbarDocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolbarDocked

`func (o *EmbeddedConfigDto) SetToolbarDocked(v string)`

SetToolbarDocked sets ToolbarDocked field to given value.

### HasToolbarDocked

`func (o *EmbeddedConfigDto) HasToolbarDocked() bool`

HasToolbarDocked returns a boolean if a field has been set.

### SetToolbarDockedNil

`func (o *EmbeddedConfigDto) SetToolbarDockedNil(b bool)`

 SetToolbarDockedNil sets the value for ToolbarDocked to be an explicit nil

### UnsetToolbarDocked
`func (o *EmbeddedConfigDto) UnsetToolbarDocked()`

UnsetToolbarDocked ensures that no value is present for ToolbarDocked, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


