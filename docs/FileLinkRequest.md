# FileLinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkId** | Pointer to **string** | The link to rewrite, as reported by `GET api/2.0/files/file/{id}/links`. An identifier that is not yet in use,  the empty one included, creates a link instead. | [optional] 
**Access** | Pointer to [**FileShare**](FileShare.md) | The rights the link grants to whoever follows it. The value that denies everything revokes the link. | [optional] 
**ExpirationDate** | Pointer to [**ApiDateTime**](ApiDateTime.md) | The moment the link stops working, read in the time zone of the portal. A date more than a few years ahead is  rejected as an invalid request; left out, the link does not expire on its own. | [optional] 
**Title** | Pointer to **NullableString** | The name the link carries in the sharing list of the file, for the people who manage it; it is not shown to  whoever follows the link. | [optional] 
**Internal** | Pointer to **bool** | Who may follow the link: `true` admits only accounts that are signed in to the portal, `false` admits anybody  who has the address. | [optional] 
**Primary** | Pointer to **bool** | Whether this link becomes the primary link of the file - the one the Copy link action of a client hands out.  A file has one primary link at a time. | [optional] 
**DenyDownload** | Pointer to **bool** | What a visitor may do with the content: `true` leaves them with viewing in the browser, `false` lets them  download and print it as their rights allow. | [optional] 
**Password** | Pointer to **NullableString** | The secret a visitor has to type before the file opens; left out, the link opens without one. | [optional] 

## Methods

### NewFileLinkRequest

`func NewFileLinkRequest() *FileLinkRequest`

NewFileLinkRequest instantiates a new FileLinkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileLinkRequestWithDefaults

`func NewFileLinkRequestWithDefaults() *FileLinkRequest`

NewFileLinkRequestWithDefaults instantiates a new FileLinkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkId

`func (o *FileLinkRequest) GetLinkId() string`

GetLinkId returns the LinkId field if non-nil, zero value otherwise.

### GetLinkIdOk

`func (o *FileLinkRequest) GetLinkIdOk() (*string, bool)`

GetLinkIdOk returns a tuple with the LinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkId

`func (o *FileLinkRequest) SetLinkId(v string)`

SetLinkId sets LinkId field to given value.

### HasLinkId

`func (o *FileLinkRequest) HasLinkId() bool`

HasLinkId returns a boolean if a field has been set.

### GetAccess

`func (o *FileLinkRequest) GetAccess() FileShare`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *FileLinkRequest) GetAccessOk() (*FileShare, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *FileLinkRequest) SetAccess(v FileShare)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *FileLinkRequest) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetExpirationDate

`func (o *FileLinkRequest) GetExpirationDate() ApiDateTime`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *FileLinkRequest) GetExpirationDateOk() (*ApiDateTime, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *FileLinkRequest) SetExpirationDate(v ApiDateTime)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *FileLinkRequest) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetTitle

`func (o *FileLinkRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FileLinkRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FileLinkRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FileLinkRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *FileLinkRequest) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *FileLinkRequest) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetInternal

`func (o *FileLinkRequest) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *FileLinkRequest) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *FileLinkRequest) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *FileLinkRequest) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetPrimary

`func (o *FileLinkRequest) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *FileLinkRequest) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *FileLinkRequest) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *FileLinkRequest) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetDenyDownload

`func (o *FileLinkRequest) GetDenyDownload() bool`

GetDenyDownload returns the DenyDownload field if non-nil, zero value otherwise.

### GetDenyDownloadOk

`func (o *FileLinkRequest) GetDenyDownloadOk() (*bool, bool)`

GetDenyDownloadOk returns a tuple with the DenyDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyDownload

`func (o *FileLinkRequest) SetDenyDownload(v bool)`

SetDenyDownload sets DenyDownload field to given value.

### HasDenyDownload

`func (o *FileLinkRequest) HasDenyDownload() bool`

HasDenyDownload returns a boolean if a field has been set.

### GetPassword

`func (o *FileLinkRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *FileLinkRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *FileLinkRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *FileLinkRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *FileLinkRequest) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *FileLinkRequest) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


