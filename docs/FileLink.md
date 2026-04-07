# FileLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filetype** | **NullableString** | The type of the file for the source viewed or edited document. | 
**Token** | Pointer to **NullableString** | The encrypted signature added to the config in the form of a token. | [optional] 
**Url** | **NullableString** | The absolute URL where the source viewed or edited document is stored. | 

## Methods

### NewFileLink

`func NewFileLink(filetype NullableString, url NullableString, ) *FileLink`

NewFileLink instantiates a new FileLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileLinkWithDefaults

`func NewFileLinkWithDefaults() *FileLink`

NewFileLinkWithDefaults instantiates a new FileLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiletype

`func (o *FileLink) GetFiletype() string`

GetFiletype returns the Filetype field if non-nil, zero value otherwise.

### GetFiletypeOk

`func (o *FileLink) GetFiletypeOk() (*string, bool)`

GetFiletypeOk returns a tuple with the Filetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiletype

`func (o *FileLink) SetFiletype(v string)`

SetFiletype sets Filetype field to given value.


### SetFiletypeNil

`func (o *FileLink) SetFiletypeNil(b bool)`

 SetFiletypeNil sets the value for Filetype to be an explicit nil

### UnsetFiletype
`func (o *FileLink) UnsetFiletype()`

UnsetFiletype ensures that no value is present for Filetype, not even an explicit nil
### GetToken

`func (o *FileLink) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FileLink) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FileLink) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *FileLink) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *FileLink) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *FileLink) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetUrl

`func (o *FileLink) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FileLink) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FileLink) SetUrl(v string)`

SetUrl sets Url field to given value.


### SetUrlNil

`func (o *FileLink) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *FileLink) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


