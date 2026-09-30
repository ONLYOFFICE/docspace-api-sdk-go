# FileShareLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The identifier of the link, the one to send back as `linkId` to change or delete it. | [optional] 
**Title** | Pointer to **NullableString** | The name the link is listed under, which its author is free to choose and to leave empty. | [optional] 
**ShareLink** | Pointer to **NullableString** | The shortened address to hand out. Opening it is what turns the link into access; the address stays the same  while the link exists. | [optional] 
**ExpirationDate** | Pointer to [**ApiDateTime**](ApiDateTime.md) | The moment the link stops working, written with the offset of the portal time zone. Null when the link was  left without an end. | [optional] 
**LinkType** | Pointer to [**LinkType**](LinkType.md) | Which of the two jobs the link does: letting somebody into the room as a member, or handing out the entry  itself. The counters of uses are filled in for the first kind only. | [optional] 
**Password** | Pointer to **NullableString** | The password a visitor has to send before the link resolves, readable only by those who may manage the link.  Empty when the link asks for none. | [optional] 
**DenyDownload** | Pointer to **NullableBool** | Whether visitors coming through this link may only read the entry in the editor and not download or print it. | [optional] 
**IsExpired** | Pointer to **NullableBool** | Whether the moment in `expirationDate` has already passed, which leaves the link in place but refuses  everybody who opens it. | [optional] 
**Primary** | Pointer to **bool** | Whether this is the one link the entry always keeps: a public or a form-filling room is given it at creation,  and deleting it there only makes a new one. | [optional] 
**Internal** | Pointer to **NullableBool** | Whether the visitor has to sign in to the portal before the link resolves, as opposed to it being open to  anybody who has the address. | [optional] 
**RequestToken** | Pointer to **NullableString** | The key that stands for this link in the calls that resolve it, such as `GET api/2.0/files/share/{key}`. It is  filled in for links that hand out the entry, and empty for the ones that invite into a room. | [optional] 
**MaxUseCount** | Pointer to **NullableInt32** | How many accounts may still join the room through this invitation link in total. Null on a link that hands out  the entry, where nothing is counted. | [optional] 
**CurrentUseCount** | Pointer to **NullableInt32** | How many accounts have already joined through this invitation link. Once it reaches `maxUseCount` the link  stops letting anybody else in. Null on a link that hands out the entry. | [optional] 

## Methods

### NewFileShareLink

`func NewFileShareLink() *FileShareLink`

NewFileShareLink instantiates a new FileShareLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileShareLinkWithDefaults

`func NewFileShareLinkWithDefaults() *FileShareLink`

NewFileShareLinkWithDefaults instantiates a new FileShareLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FileShareLink) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FileShareLink) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FileShareLink) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FileShareLink) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *FileShareLink) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FileShareLink) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FileShareLink) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FileShareLink) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *FileShareLink) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *FileShareLink) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetShareLink

`func (o *FileShareLink) GetShareLink() string`

GetShareLink returns the ShareLink field if non-nil, zero value otherwise.

### GetShareLinkOk

`func (o *FileShareLink) GetShareLinkOk() (*string, bool)`

GetShareLinkOk returns a tuple with the ShareLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareLink

`func (o *FileShareLink) SetShareLink(v string)`

SetShareLink sets ShareLink field to given value.

### HasShareLink

`func (o *FileShareLink) HasShareLink() bool`

HasShareLink returns a boolean if a field has been set.

### SetShareLinkNil

`func (o *FileShareLink) SetShareLinkNil(b bool)`

 SetShareLinkNil sets the value for ShareLink to be an explicit nil

### UnsetShareLink
`func (o *FileShareLink) UnsetShareLink()`

UnsetShareLink ensures that no value is present for ShareLink, not even an explicit nil
### GetExpirationDate

`func (o *FileShareLink) GetExpirationDate() ApiDateTime`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *FileShareLink) GetExpirationDateOk() (*ApiDateTime, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *FileShareLink) SetExpirationDate(v ApiDateTime)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *FileShareLink) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetLinkType

`func (o *FileShareLink) GetLinkType() LinkType`

GetLinkType returns the LinkType field if non-nil, zero value otherwise.

### GetLinkTypeOk

`func (o *FileShareLink) GetLinkTypeOk() (*LinkType, bool)`

GetLinkTypeOk returns a tuple with the LinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkType

`func (o *FileShareLink) SetLinkType(v LinkType)`

SetLinkType sets LinkType field to given value.

### HasLinkType

`func (o *FileShareLink) HasLinkType() bool`

HasLinkType returns a boolean if a field has been set.

### GetPassword

`func (o *FileShareLink) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *FileShareLink) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *FileShareLink) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *FileShareLink) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *FileShareLink) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *FileShareLink) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetDenyDownload

`func (o *FileShareLink) GetDenyDownload() bool`

GetDenyDownload returns the DenyDownload field if non-nil, zero value otherwise.

### GetDenyDownloadOk

`func (o *FileShareLink) GetDenyDownloadOk() (*bool, bool)`

GetDenyDownloadOk returns a tuple with the DenyDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyDownload

`func (o *FileShareLink) SetDenyDownload(v bool)`

SetDenyDownload sets DenyDownload field to given value.

### HasDenyDownload

`func (o *FileShareLink) HasDenyDownload() bool`

HasDenyDownload returns a boolean if a field has been set.

### SetDenyDownloadNil

`func (o *FileShareLink) SetDenyDownloadNil(b bool)`

 SetDenyDownloadNil sets the value for DenyDownload to be an explicit nil

### UnsetDenyDownload
`func (o *FileShareLink) UnsetDenyDownload()`

UnsetDenyDownload ensures that no value is present for DenyDownload, not even an explicit nil
### GetIsExpired

`func (o *FileShareLink) GetIsExpired() bool`

GetIsExpired returns the IsExpired field if non-nil, zero value otherwise.

### GetIsExpiredOk

`func (o *FileShareLink) GetIsExpiredOk() (*bool, bool)`

GetIsExpiredOk returns a tuple with the IsExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExpired

`func (o *FileShareLink) SetIsExpired(v bool)`

SetIsExpired sets IsExpired field to given value.

### HasIsExpired

`func (o *FileShareLink) HasIsExpired() bool`

HasIsExpired returns a boolean if a field has been set.

### SetIsExpiredNil

`func (o *FileShareLink) SetIsExpiredNil(b bool)`

 SetIsExpiredNil sets the value for IsExpired to be an explicit nil

### UnsetIsExpired
`func (o *FileShareLink) UnsetIsExpired()`

UnsetIsExpired ensures that no value is present for IsExpired, not even an explicit nil
### GetPrimary

`func (o *FileShareLink) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *FileShareLink) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *FileShareLink) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *FileShareLink) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetInternal

`func (o *FileShareLink) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *FileShareLink) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *FileShareLink) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *FileShareLink) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### SetInternalNil

`func (o *FileShareLink) SetInternalNil(b bool)`

 SetInternalNil sets the value for Internal to be an explicit nil

### UnsetInternal
`func (o *FileShareLink) UnsetInternal()`

UnsetInternal ensures that no value is present for Internal, not even an explicit nil
### GetRequestToken

`func (o *FileShareLink) GetRequestToken() string`

GetRequestToken returns the RequestToken field if non-nil, zero value otherwise.

### GetRequestTokenOk

`func (o *FileShareLink) GetRequestTokenOk() (*string, bool)`

GetRequestTokenOk returns a tuple with the RequestToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestToken

`func (o *FileShareLink) SetRequestToken(v string)`

SetRequestToken sets RequestToken field to given value.

### HasRequestToken

`func (o *FileShareLink) HasRequestToken() bool`

HasRequestToken returns a boolean if a field has been set.

### SetRequestTokenNil

`func (o *FileShareLink) SetRequestTokenNil(b bool)`

 SetRequestTokenNil sets the value for RequestToken to be an explicit nil

### UnsetRequestToken
`func (o *FileShareLink) UnsetRequestToken()`

UnsetRequestToken ensures that no value is present for RequestToken, not even an explicit nil
### GetMaxUseCount

`func (o *FileShareLink) GetMaxUseCount() int32`

GetMaxUseCount returns the MaxUseCount field if non-nil, zero value otherwise.

### GetMaxUseCountOk

`func (o *FileShareLink) GetMaxUseCountOk() (*int32, bool)`

GetMaxUseCountOk returns a tuple with the MaxUseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUseCount

`func (o *FileShareLink) SetMaxUseCount(v int32)`

SetMaxUseCount sets MaxUseCount field to given value.

### HasMaxUseCount

`func (o *FileShareLink) HasMaxUseCount() bool`

HasMaxUseCount returns a boolean if a field has been set.

### SetMaxUseCountNil

`func (o *FileShareLink) SetMaxUseCountNil(b bool)`

 SetMaxUseCountNil sets the value for MaxUseCount to be an explicit nil

### UnsetMaxUseCount
`func (o *FileShareLink) UnsetMaxUseCount()`

UnsetMaxUseCount ensures that no value is present for MaxUseCount, not even an explicit nil
### GetCurrentUseCount

`func (o *FileShareLink) GetCurrentUseCount() int32`

GetCurrentUseCount returns the CurrentUseCount field if non-nil, zero value otherwise.

### GetCurrentUseCountOk

`func (o *FileShareLink) GetCurrentUseCountOk() (*int32, bool)`

GetCurrentUseCountOk returns a tuple with the CurrentUseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUseCount

`func (o *FileShareLink) SetCurrentUseCount(v int32)`

SetCurrentUseCount sets CurrentUseCount field to given value.

### HasCurrentUseCount

`func (o *FileShareLink) HasCurrentUseCount() bool`

HasCurrentUseCount returns a boolean if a field has been set.

### SetCurrentUseCountNil

`func (o *FileShareLink) SetCurrentUseCountNil(b bool)`

 SetCurrentUseCountNil sets the value for CurrentUseCount to be an explicit nil

### UnsetCurrentUseCount
`func (o *FileShareLink) UnsetCurrentUseCount()`

UnsetCurrentUseCount ensures that no value is present for CurrentUseCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


