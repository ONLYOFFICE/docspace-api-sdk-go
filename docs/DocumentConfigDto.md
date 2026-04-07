# DocumentConfigDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileType** | Pointer to **NullableString** | The file type of the document. | [optional] 
**Info** | Pointer to [**InfoConfigDto**](InfoConfigDto.md) |  | [optional] 
**IsLinkedForMe** | Pointer to **bool** | Specifies if the documnet is linked for current user. | [optional] 
**Key** | Pointer to **NullableString** | The document key. | [optional] 
**Permissions** | Pointer to [**PermissionsConfig**](PermissionsConfig.md) |  | [optional] 
**SharedLinkParam** | Pointer to **NullableString** | The shared link parameter of the document. | [optional] 
**SharedLinkKey** | Pointer to **NullableString** | The shared link key of the document. | [optional] 
**ReferenceData** | Pointer to [**FileReferenceData**](FileReferenceData.md) |  | [optional] 
**Title** | Pointer to **NullableString** | The document title. | [optional] 
**Url** | Pointer to **NullableString** | The document url. | [optional] 
**IsForm** | Pointer to **bool** | Indicates whether this is a form. | [optional] 
**Options** | Pointer to [**Options**](Options.md) |  | [optional] 

## Methods

### NewDocumentConfigDto

`func NewDocumentConfigDto() *DocumentConfigDto`

NewDocumentConfigDto instantiates a new DocumentConfigDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentConfigDtoWithDefaults

`func NewDocumentConfigDtoWithDefaults() *DocumentConfigDto`

NewDocumentConfigDtoWithDefaults instantiates a new DocumentConfigDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileType

`func (o *DocumentConfigDto) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *DocumentConfigDto) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *DocumentConfigDto) SetFileType(v string)`

SetFileType sets FileType field to given value.

### HasFileType

`func (o *DocumentConfigDto) HasFileType() bool`

HasFileType returns a boolean if a field has been set.

### SetFileTypeNil

`func (o *DocumentConfigDto) SetFileTypeNil(b bool)`

 SetFileTypeNil sets the value for FileType to be an explicit nil

### UnsetFileType
`func (o *DocumentConfigDto) UnsetFileType()`

UnsetFileType ensures that no value is present for FileType, not even an explicit nil
### GetInfo

`func (o *DocumentConfigDto) GetInfo() InfoConfigDto`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *DocumentConfigDto) GetInfoOk() (*InfoConfigDto, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *DocumentConfigDto) SetInfo(v InfoConfigDto)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *DocumentConfigDto) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetIsLinkedForMe

`func (o *DocumentConfigDto) GetIsLinkedForMe() bool`

GetIsLinkedForMe returns the IsLinkedForMe field if non-nil, zero value otherwise.

### GetIsLinkedForMeOk

`func (o *DocumentConfigDto) GetIsLinkedForMeOk() (*bool, bool)`

GetIsLinkedForMeOk returns a tuple with the IsLinkedForMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLinkedForMe

`func (o *DocumentConfigDto) SetIsLinkedForMe(v bool)`

SetIsLinkedForMe sets IsLinkedForMe field to given value.

### HasIsLinkedForMe

`func (o *DocumentConfigDto) HasIsLinkedForMe() bool`

HasIsLinkedForMe returns a boolean if a field has been set.

### GetKey

`func (o *DocumentConfigDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DocumentConfigDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DocumentConfigDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *DocumentConfigDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *DocumentConfigDto) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *DocumentConfigDto) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetPermissions

`func (o *DocumentConfigDto) GetPermissions() PermissionsConfig`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *DocumentConfigDto) GetPermissionsOk() (*PermissionsConfig, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *DocumentConfigDto) SetPermissions(v PermissionsConfig)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *DocumentConfigDto) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetSharedLinkParam

`func (o *DocumentConfigDto) GetSharedLinkParam() string`

GetSharedLinkParam returns the SharedLinkParam field if non-nil, zero value otherwise.

### GetSharedLinkParamOk

`func (o *DocumentConfigDto) GetSharedLinkParamOk() (*string, bool)`

GetSharedLinkParamOk returns a tuple with the SharedLinkParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedLinkParam

`func (o *DocumentConfigDto) SetSharedLinkParam(v string)`

SetSharedLinkParam sets SharedLinkParam field to given value.

### HasSharedLinkParam

`func (o *DocumentConfigDto) HasSharedLinkParam() bool`

HasSharedLinkParam returns a boolean if a field has been set.

### SetSharedLinkParamNil

`func (o *DocumentConfigDto) SetSharedLinkParamNil(b bool)`

 SetSharedLinkParamNil sets the value for SharedLinkParam to be an explicit nil

### UnsetSharedLinkParam
`func (o *DocumentConfigDto) UnsetSharedLinkParam()`

UnsetSharedLinkParam ensures that no value is present for SharedLinkParam, not even an explicit nil
### GetSharedLinkKey

`func (o *DocumentConfigDto) GetSharedLinkKey() string`

GetSharedLinkKey returns the SharedLinkKey field if non-nil, zero value otherwise.

### GetSharedLinkKeyOk

`func (o *DocumentConfigDto) GetSharedLinkKeyOk() (*string, bool)`

GetSharedLinkKeyOk returns a tuple with the SharedLinkKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedLinkKey

`func (o *DocumentConfigDto) SetSharedLinkKey(v string)`

SetSharedLinkKey sets SharedLinkKey field to given value.

### HasSharedLinkKey

`func (o *DocumentConfigDto) HasSharedLinkKey() bool`

HasSharedLinkKey returns a boolean if a field has been set.

### SetSharedLinkKeyNil

`func (o *DocumentConfigDto) SetSharedLinkKeyNil(b bool)`

 SetSharedLinkKeyNil sets the value for SharedLinkKey to be an explicit nil

### UnsetSharedLinkKey
`func (o *DocumentConfigDto) UnsetSharedLinkKey()`

UnsetSharedLinkKey ensures that no value is present for SharedLinkKey, not even an explicit nil
### GetReferenceData

`func (o *DocumentConfigDto) GetReferenceData() FileReferenceData`

GetReferenceData returns the ReferenceData field if non-nil, zero value otherwise.

### GetReferenceDataOk

`func (o *DocumentConfigDto) GetReferenceDataOk() (*FileReferenceData, bool)`

GetReferenceDataOk returns a tuple with the ReferenceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceData

`func (o *DocumentConfigDto) SetReferenceData(v FileReferenceData)`

SetReferenceData sets ReferenceData field to given value.

### HasReferenceData

`func (o *DocumentConfigDto) HasReferenceData() bool`

HasReferenceData returns a boolean if a field has been set.

### GetTitle

`func (o *DocumentConfigDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DocumentConfigDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DocumentConfigDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DocumentConfigDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *DocumentConfigDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *DocumentConfigDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetUrl

`func (o *DocumentConfigDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DocumentConfigDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DocumentConfigDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *DocumentConfigDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *DocumentConfigDto) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *DocumentConfigDto) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetIsForm

`func (o *DocumentConfigDto) GetIsForm() bool`

GetIsForm returns the IsForm field if non-nil, zero value otherwise.

### GetIsFormOk

`func (o *DocumentConfigDto) GetIsFormOk() (*bool, bool)`

GetIsFormOk returns a tuple with the IsForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForm

`func (o *DocumentConfigDto) SetIsForm(v bool)`

SetIsForm sets IsForm field to given value.

### HasIsForm

`func (o *DocumentConfigDto) HasIsForm() bool`

HasIsForm returns a boolean if a field has been set.

### GetOptions

`func (o *DocumentConfigDto) GetOptions() Options`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *DocumentConfigDto) GetOptionsOk() (*Options, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *DocumentConfigDto) SetOptions(v Options)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *DocumentConfigDto) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


