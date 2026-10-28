# EditHistoryUrlDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **NullableString** | The document key of that revision. When the file has no earlier revision the portal generates a fresh key for  the template it falls back to, so the value is not always one an earlier revision ever had. | [optional] 
**Url** | Pointer to **NullableString** | The address that revision's content is served from. It is meant for the editing service and carries its own  key, which is valid for a limited time. | [optional] 
**FileType** | Pointer to **NullableString** | The format of that revision, as an extension without the leading dot. | [optional] 

## Methods

### NewEditHistoryUrlDto

`func NewEditHistoryUrlDto() *EditHistoryUrlDto`

NewEditHistoryUrlDto instantiates a new EditHistoryUrlDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditHistoryUrlDtoWithDefaults

`func NewEditHistoryUrlDtoWithDefaults() *EditHistoryUrlDto`

NewEditHistoryUrlDtoWithDefaults instantiates a new EditHistoryUrlDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *EditHistoryUrlDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *EditHistoryUrlDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *EditHistoryUrlDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *EditHistoryUrlDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *EditHistoryUrlDto) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *EditHistoryUrlDto) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetUrl

`func (o *EditHistoryUrlDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EditHistoryUrlDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EditHistoryUrlDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *EditHistoryUrlDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *EditHistoryUrlDto) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *EditHistoryUrlDto) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetFileType

`func (o *EditHistoryUrlDto) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *EditHistoryUrlDto) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *EditHistoryUrlDto) SetFileType(v string)`

SetFileType sets FileType field to given value.

### HasFileType

`func (o *EditHistoryUrlDto) HasFileType() bool`

HasFileType returns a boolean if a field has been set.

### SetFileTypeNil

`func (o *EditHistoryUrlDto) SetFileTypeNil(b bool)`

 SetFileTypeNil sets the value for FileType to be an explicit nil

### UnsetFileType
`func (o *EditHistoryUrlDto) UnsetFileType()`

UnsetFileType ensures that no value is present for FileType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


