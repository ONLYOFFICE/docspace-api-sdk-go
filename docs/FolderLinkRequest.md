# FolderLinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkId** | Pointer to **string** | The folder link ID. | [optional] 
**Access** | Pointer to [**FileShare**](FileShare.md) |  | [optional] 
**ExpirationDate** | Pointer to [**ApiDateTime**](ApiDateTime.md) |  | [optional] 
**Title** | Pointer to **NullableString** | The link name. | [optional] 
**Password** | Pointer to **NullableString** | The link password. | [optional] 
**DenyDownload** | Pointer to **bool** | Specifies if downloading the file from the link is disabled or not. | [optional] 
**Internal** | Pointer to **bool** | The link scope, whether it is internal or not. | [optional] 
**Primary** | Pointer to **bool** | Specifies whether the folder link is primary or not. | [optional] 

## Methods

### NewFolderLinkRequest

`func NewFolderLinkRequest() *FolderLinkRequest`

NewFolderLinkRequest instantiates a new FolderLinkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFolderLinkRequestWithDefaults

`func NewFolderLinkRequestWithDefaults() *FolderLinkRequest`

NewFolderLinkRequestWithDefaults instantiates a new FolderLinkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkId

`func (o *FolderLinkRequest) GetLinkId() string`

GetLinkId returns the LinkId field if non-nil, zero value otherwise.

### GetLinkIdOk

`func (o *FolderLinkRequest) GetLinkIdOk() (*string, bool)`

GetLinkIdOk returns a tuple with the LinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkId

`func (o *FolderLinkRequest) SetLinkId(v string)`

SetLinkId sets LinkId field to given value.

### HasLinkId

`func (o *FolderLinkRequest) HasLinkId() bool`

HasLinkId returns a boolean if a field has been set.

### GetAccess

`func (o *FolderLinkRequest) GetAccess() FileShare`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *FolderLinkRequest) GetAccessOk() (*FileShare, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *FolderLinkRequest) SetAccess(v FileShare)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *FolderLinkRequest) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetExpirationDate

`func (o *FolderLinkRequest) GetExpirationDate() ApiDateTime`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *FolderLinkRequest) GetExpirationDateOk() (*ApiDateTime, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *FolderLinkRequest) SetExpirationDate(v ApiDateTime)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *FolderLinkRequest) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetTitle

`func (o *FolderLinkRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FolderLinkRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FolderLinkRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FolderLinkRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *FolderLinkRequest) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *FolderLinkRequest) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetPassword

`func (o *FolderLinkRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *FolderLinkRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *FolderLinkRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *FolderLinkRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *FolderLinkRequest) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *FolderLinkRequest) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetDenyDownload

`func (o *FolderLinkRequest) GetDenyDownload() bool`

GetDenyDownload returns the DenyDownload field if non-nil, zero value otherwise.

### GetDenyDownloadOk

`func (o *FolderLinkRequest) GetDenyDownloadOk() (*bool, bool)`

GetDenyDownloadOk returns a tuple with the DenyDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyDownload

`func (o *FolderLinkRequest) SetDenyDownload(v bool)`

SetDenyDownload sets DenyDownload field to given value.

### HasDenyDownload

`func (o *FolderLinkRequest) HasDenyDownload() bool`

HasDenyDownload returns a boolean if a field has been set.

### GetInternal

`func (o *FolderLinkRequest) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *FolderLinkRequest) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *FolderLinkRequest) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *FolderLinkRequest) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetPrimary

`func (o *FolderLinkRequest) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *FolderLinkRequest) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *FolderLinkRequest) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *FolderLinkRequest) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


