# EditHistoryDataDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangesUrl** | Pointer to **NullableString** | The URL address of the file with the document changes data. | [optional] 
**Key** | **NullableString** | The document identifier used to unambiguously identify the document file. | 
**Previous** | Pointer to [**EditHistoryUrl**](EditHistoryUrl.md) | The object of the previous version of the document. | [optional] 
**Token** | Pointer to **NullableString** | The encrypted signature added to the parameter in the form of a token. | [optional] 
**Url** | **NullableString** | The URL address of the current document version. | 
**Version** | **int32** | The document version number. | 
**FileType** | **NullableString** | The document extension. | 

## Methods

### NewEditHistoryDataDto

`func NewEditHistoryDataDto(key NullableString, url NullableString, version int32, fileType NullableString, ) *EditHistoryDataDto`

NewEditHistoryDataDto instantiates a new EditHistoryDataDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditHistoryDataDtoWithDefaults

`func NewEditHistoryDataDtoWithDefaults() *EditHistoryDataDto`

NewEditHistoryDataDtoWithDefaults instantiates a new EditHistoryDataDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangesUrl

`func (o *EditHistoryDataDto) GetChangesUrl() string`

GetChangesUrl returns the ChangesUrl field if non-nil, zero value otherwise.

### GetChangesUrlOk

`func (o *EditHistoryDataDto) GetChangesUrlOk() (*string, bool)`

GetChangesUrlOk returns a tuple with the ChangesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangesUrl

`func (o *EditHistoryDataDto) SetChangesUrl(v string)`

SetChangesUrl sets ChangesUrl field to given value.

### HasChangesUrl

`func (o *EditHistoryDataDto) HasChangesUrl() bool`

HasChangesUrl returns a boolean if a field has been set.

### SetChangesUrlNil

`func (o *EditHistoryDataDto) SetChangesUrlNil(b bool)`

 SetChangesUrlNil sets the value for ChangesUrl to be an explicit nil

### UnsetChangesUrl
`func (o *EditHistoryDataDto) UnsetChangesUrl()`

UnsetChangesUrl ensures that no value is present for ChangesUrl, not even an explicit nil
### GetKey

`func (o *EditHistoryDataDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *EditHistoryDataDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *EditHistoryDataDto) SetKey(v string)`

SetKey sets Key field to given value.


### SetKeyNil

`func (o *EditHistoryDataDto) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *EditHistoryDataDto) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetPrevious

`func (o *EditHistoryDataDto) GetPrevious() EditHistoryUrl`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *EditHistoryDataDto) GetPreviousOk() (*EditHistoryUrl, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *EditHistoryDataDto) SetPrevious(v EditHistoryUrl)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *EditHistoryDataDto) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### GetToken

`func (o *EditHistoryDataDto) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *EditHistoryDataDto) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *EditHistoryDataDto) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *EditHistoryDataDto) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *EditHistoryDataDto) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *EditHistoryDataDto) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetUrl

`func (o *EditHistoryDataDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EditHistoryDataDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EditHistoryDataDto) SetUrl(v string)`

SetUrl sets Url field to given value.


### SetUrlNil

`func (o *EditHistoryDataDto) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *EditHistoryDataDto) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetVersion

`func (o *EditHistoryDataDto) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EditHistoryDataDto) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EditHistoryDataDto) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetFileType

`func (o *EditHistoryDataDto) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *EditHistoryDataDto) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *EditHistoryDataDto) SetFileType(v string)`

SetFileType sets FileType field to given value.


### SetFileTypeNil

`func (o *EditHistoryDataDto) SetFileTypeNil(b bool)`

 SetFileTypeNil sets the value for FileType to be an explicit nil

### UnsetFileType
`func (o *EditHistoryDataDto) UnsetFileType()`

UnsetFileType ensures that no value is present for FileType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


