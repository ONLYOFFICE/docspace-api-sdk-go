# FileShareLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier of the shared link. | [optional] 
**Title** | Pointer to **NullableString** | The title of the shared content. | [optional] 
**ShareLink** | Pointer to **NullableString** | The URL for accessing the shared content. | [optional] 
**ExpirationDate** | Pointer to **NullableTime** | The date when the shared link expires. | [optional] 
**LinkType** | Pointer to [**LinkType**](LinkType.md) | The sharing link type (e.g., Invitation). | [optional] 
**Password** | Pointer to **NullableString** | The password protection for accessing the shared content. | [optional] 
**DenyDownload** | Pointer to **NullableBool** | Indicates whether downloading of the shared content is prohibited. | [optional] 
**IsExpired** | Pointer to **NullableBool** | Indicates whether the shared link has expired. | [optional] 
**Primary** | Pointer to **bool** | Indicates whether this is the primary shared link. | [optional] 
**Internal** | Pointer to **NullableBool** | Indicates whether the link is for the internal sharing only. | [optional] 
**RequestToken** | Pointer to **NullableString** | The token for validating access requests. | [optional] 
**MaxUseCount** | Pointer to **NullableInt32** | The maximum number of times the invitation link can be used. | [optional] 
**CurrentUseCount** | Pointer to **NullableInt32** | The current number of times the invitation link has been used. | [optional] 

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

`func (o *FileShareLink) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *FileShareLink) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *FileShareLink) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *FileShareLink) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *FileShareLink) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *FileShareLink) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
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


