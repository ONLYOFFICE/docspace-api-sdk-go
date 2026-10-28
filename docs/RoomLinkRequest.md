# RoomLinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkId** | Pointer to **string** | Which link to change, taken from `GET api/2.0/files/rooms/{id}/links`. Leaving it out creates a link, and an  identifier the room does not know creates a link carrying that identifier. | [optional] 
**Access** | Pointer to [**FileShare**](FileShare.md) | What whoever opens the link may do in the room. The value 0 revokes the link instead of changing it, and the  levels a room accepts depend on its kind. | [optional] 
**ExpirationDate** | Pointer to [**ApiDateTime**](ApiDateTime.md) | When the link stops working, written with the offset of the portal time zone. A date already past is dropped  silently for an external link and refused for an invitation link, and a date further ahead than the portal  allows is refused as well; leaving it out means the link does not expire. | [optional] 
**Internal** | Pointer to **bool** | Whether the external link works only for people already signed in to the portal. With it off the link opens  the room for anyone who has the address, subject to the password. | [optional] 
**Title** | Pointer to **NullableString** | The name the link is shown under in the room. An empty value is accepted and the portal names the link itself,  so the answer is what tells the caller the name in use. | [optional] 
**LinkType** | Pointer to [**LinkType**](LinkType.md) | Which kind of link to create: an invitation link makes whoever opens it a member of the room, while an  external link opens the room without an account. It is fixed when the link is created and is ignored on later  changes. | [optional] 
**Password** | Pointer to **NullableString** | The password an external link asks for before it opens the room. An empty value leaves the link open to anyone  who has the address, and the password is never returned when links are listed. | [optional] 
**DenyDownload** | Pointer to **bool** | Whether people arriving through the link are stopped from downloading and printing what they open. They can  still read the documents in the editor. | [optional] 
**MaxUseCount** | Pointer to **NullableInt32** | How many people an invitation link may still let in before it stops working. A value below the number of  people who already used it is refused, and leaving it out puts no ceiling on the link. | [optional] 
**CurrentUseCount** | Pointer to **int32** | How many people have already joined through this invitation link. The value is kept by the portal: it is  reported back when links are listed and anything sent here is ignored. | [optional] 

## Methods

### NewRoomLinkRequest

`func NewRoomLinkRequest() *RoomLinkRequest`

NewRoomLinkRequest instantiates a new RoomLinkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoomLinkRequestWithDefaults

`func NewRoomLinkRequestWithDefaults() *RoomLinkRequest`

NewRoomLinkRequestWithDefaults instantiates a new RoomLinkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkId

`func (o *RoomLinkRequest) GetLinkId() string`

GetLinkId returns the LinkId field if non-nil, zero value otherwise.

### GetLinkIdOk

`func (o *RoomLinkRequest) GetLinkIdOk() (*string, bool)`

GetLinkIdOk returns a tuple with the LinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkId

`func (o *RoomLinkRequest) SetLinkId(v string)`

SetLinkId sets LinkId field to given value.

### HasLinkId

`func (o *RoomLinkRequest) HasLinkId() bool`

HasLinkId returns a boolean if a field has been set.

### GetAccess

`func (o *RoomLinkRequest) GetAccess() FileShare`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *RoomLinkRequest) GetAccessOk() (*FileShare, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *RoomLinkRequest) SetAccess(v FileShare)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *RoomLinkRequest) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetExpirationDate

`func (o *RoomLinkRequest) GetExpirationDate() ApiDateTime`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *RoomLinkRequest) GetExpirationDateOk() (*ApiDateTime, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *RoomLinkRequest) SetExpirationDate(v ApiDateTime)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *RoomLinkRequest) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetInternal

`func (o *RoomLinkRequest) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *RoomLinkRequest) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *RoomLinkRequest) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *RoomLinkRequest) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetTitle

`func (o *RoomLinkRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RoomLinkRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RoomLinkRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *RoomLinkRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *RoomLinkRequest) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *RoomLinkRequest) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetLinkType

`func (o *RoomLinkRequest) GetLinkType() LinkType`

GetLinkType returns the LinkType field if non-nil, zero value otherwise.

### GetLinkTypeOk

`func (o *RoomLinkRequest) GetLinkTypeOk() (*LinkType, bool)`

GetLinkTypeOk returns a tuple with the LinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkType

`func (o *RoomLinkRequest) SetLinkType(v LinkType)`

SetLinkType sets LinkType field to given value.

### HasLinkType

`func (o *RoomLinkRequest) HasLinkType() bool`

HasLinkType returns a boolean if a field has been set.

### GetPassword

`func (o *RoomLinkRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RoomLinkRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RoomLinkRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *RoomLinkRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *RoomLinkRequest) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *RoomLinkRequest) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetDenyDownload

`func (o *RoomLinkRequest) GetDenyDownload() bool`

GetDenyDownload returns the DenyDownload field if non-nil, zero value otherwise.

### GetDenyDownloadOk

`func (o *RoomLinkRequest) GetDenyDownloadOk() (*bool, bool)`

GetDenyDownloadOk returns a tuple with the DenyDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyDownload

`func (o *RoomLinkRequest) SetDenyDownload(v bool)`

SetDenyDownload sets DenyDownload field to given value.

### HasDenyDownload

`func (o *RoomLinkRequest) HasDenyDownload() bool`

HasDenyDownload returns a boolean if a field has been set.

### GetMaxUseCount

`func (o *RoomLinkRequest) GetMaxUseCount() int32`

GetMaxUseCount returns the MaxUseCount field if non-nil, zero value otherwise.

### GetMaxUseCountOk

`func (o *RoomLinkRequest) GetMaxUseCountOk() (*int32, bool)`

GetMaxUseCountOk returns a tuple with the MaxUseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUseCount

`func (o *RoomLinkRequest) SetMaxUseCount(v int32)`

SetMaxUseCount sets MaxUseCount field to given value.

### HasMaxUseCount

`func (o *RoomLinkRequest) HasMaxUseCount() bool`

HasMaxUseCount returns a boolean if a field has been set.

### SetMaxUseCountNil

`func (o *RoomLinkRequest) SetMaxUseCountNil(b bool)`

 SetMaxUseCountNil sets the value for MaxUseCount to be an explicit nil

### UnsetMaxUseCount
`func (o *RoomLinkRequest) UnsetMaxUseCount()`

UnsetMaxUseCount ensures that no value is present for MaxUseCount, not even an explicit nil
### GetCurrentUseCount

`func (o *RoomLinkRequest) GetCurrentUseCount() int32`

GetCurrentUseCount returns the CurrentUseCount field if non-nil, zero value otherwise.

### GetCurrentUseCountOk

`func (o *RoomLinkRequest) GetCurrentUseCountOk() (*int32, bool)`

GetCurrentUseCountOk returns a tuple with the CurrentUseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUseCount

`func (o *RoomLinkRequest) SetCurrentUseCount(v int32)`

SetCurrentUseCount sets CurrentUseCount field to given value.

### HasCurrentUseCount

`func (o *RoomLinkRequest) HasCurrentUseCount() bool`

HasCurrentUseCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


