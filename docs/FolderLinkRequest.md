# FolderLinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkId** | Pointer to **string** | Which link the request addresses: the identifier of an existing link rewrites that link, while an identifier  that is not in use, the empty one included, creates a new link. Take an existing identifier from  `GET api/2.0/files/folder/{id}/links`. | [optional] 
**Access** | Pointer to [**FileShare**](FileShare.md) | The rights a visitor following the link is given. The value that grants nothing revokes the link instead of  setting it, and the answer is then empty. | [optional] 
**ExpirationDate** | Pointer to [**ApiDateTime**](ApiDateTime.md) | The moment the link stops working, sent as an ISO-8601 stamp. A moment that lies in the past is ignored,  and leaving the field out gives the link no expiry. | [optional] 
**Title** | Pointer to **NullableString** | The name the link is listed under for the people who manage the folder; a visitor following it never sees the  name. | [optional] 
**Password** | Pointer to **NullableString** | The secret a visitor has to enter before the link opens. Leave it out for a link that opens without one; the  secret itself is never given back, only the fact that one is set. | [optional] 
**DenyDownload** | Pointer to **bool** | Whether visitors are left with viewing alone: with true downloading and copying through the link are blocked,  with false they are allowed. | [optional] 
**Internal** | Pointer to **bool** | Whether the link admits signed-in portal members only: with true a visitor has to sign in before the link  opens, with false anyone holding the address may follow it. | [optional] 
**Primary** | Pointer to **bool** | Whether this link becomes the primary link of the folder, the one the Copy link action of a client hands  out; a folder has one primary link at a time. | [optional] 

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


