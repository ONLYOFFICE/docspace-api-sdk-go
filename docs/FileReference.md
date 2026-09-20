# FileReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferenceData** | Pointer to [**FileReferenceData**](FileReferenceData.md) | How this document is named when another spreadsheet refers to it. Send it back as it stands to resolve the  reference again. | [optional] 
**Error** | Pointer to **NullableString** | Filled in when the reference resolved to nothing; the rest of the descriptor is then empty and must not be  handed to the editors. | [optional] 
**Path** | Pointer to **NullableString** | The title of the document the reference resolved to. | [optional] 
**Url** | Pointer to **NullableString** | Where the content is fetched from. It is addressed to the host the document service can reach, which on a  deployment with a private editor network is not the address a browser should follow. | [optional] 
**FileType** | Pointer to **NullableString** | The format the content is in, without the leading dot. | [optional] 
**Key** | Pointer to **NullableString** | Identifies the exact revision to the editors: two clients that receive the same key read the same co-editing  session, and the key changes as soon as the document is saved. | [optional] 
**Link** | Pointer to **NullableString** | The address of the document in the portal web editor - the link to put in front of a person, unlike the  download address above. | [optional] 
**Token** | Pointer to **NullableString** | Signs this descriptor so that the editors can trust it. It stays empty on a portal that has no signature  secret configured for the document service. | [optional] 

## Methods

### NewFileReference

`func NewFileReference() *FileReference`

NewFileReference instantiates a new FileReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileReferenceWithDefaults

`func NewFileReferenceWithDefaults() *FileReference`

NewFileReferenceWithDefaults instantiates a new FileReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferenceData

`func (o *FileReference) GetReferenceData() FileReferenceData`

GetReferenceData returns the ReferenceData field if non-nil, zero value otherwise.

### GetReferenceDataOk

`func (o *FileReference) GetReferenceDataOk() (*FileReferenceData, bool)`

GetReferenceDataOk returns a tuple with the ReferenceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceData

`func (o *FileReference) SetReferenceData(v FileReferenceData)`

SetReferenceData sets ReferenceData field to given value.

### HasReferenceData

`func (o *FileReference) HasReferenceData() bool`

HasReferenceData returns a boolean if a field has been set.

### GetError

`func (o *FileReference) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *FileReference) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *FileReference) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *FileReference) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *FileReference) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *FileReference) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetPath

`func (o *FileReference) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FileReference) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FileReference) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *FileReference) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *FileReference) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *FileReference) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetUrl

`func (o *FileReference) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FileReference) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FileReference) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *FileReference) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *FileReference) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *FileReference) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetFileType

`func (o *FileReference) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *FileReference) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *FileReference) SetFileType(v string)`

SetFileType sets FileType field to given value.

### HasFileType

`func (o *FileReference) HasFileType() bool`

HasFileType returns a boolean if a field has been set.

### SetFileTypeNil

`func (o *FileReference) SetFileTypeNil(b bool)`

 SetFileTypeNil sets the value for FileType to be an explicit nil

### UnsetFileType
`func (o *FileReference) UnsetFileType()`

UnsetFileType ensures that no value is present for FileType, not even an explicit nil
### GetKey

`func (o *FileReference) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FileReference) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FileReference) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *FileReference) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *FileReference) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *FileReference) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetLink

`func (o *FileReference) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *FileReference) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *FileReference) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *FileReference) HasLink() bool`

HasLink returns a boolean if a field has been set.

### SetLinkNil

`func (o *FileReference) SetLinkNil(b bool)`

 SetLinkNil sets the value for Link to be an explicit nil

### UnsetLink
`func (o *FileReference) UnsetLink()`

UnsetLink ensures that no value is present for Link, not even an explicit nil
### GetToken

`func (o *FileReference) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FileReference) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FileReference) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *FileReference) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *FileReference) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *FileReference) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


