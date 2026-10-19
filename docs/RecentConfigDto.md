# RecentConfigDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Folder** | Pointer to **NullableString** | The folder shown next to the entry, as a readable name rather than an id. | [optional] 
**Title** | Pointer to **NullableString** | The name shown for the entry. | [optional] 
**Url** | Pointer to **NullableString** | Where the entry opens. | [optional] 

## Methods

### NewRecentConfigDto

`func NewRecentConfigDto() *RecentConfigDto`

NewRecentConfigDto instantiates a new RecentConfigDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecentConfigDtoWithDefaults

`func NewRecentConfigDtoWithDefaults() *RecentConfigDto`

NewRecentConfigDtoWithDefaults instantiates a new RecentConfigDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFolder

`func (o *RecentConfigDto) GetFolder() string`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *RecentConfigDto) GetFolderOk() (*string, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *RecentConfigDto) SetFolder(v string)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *RecentConfigDto) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### SetFolderNil

`func (o *RecentConfigDto) SetFolderNil(b bool)`

 SetFolderNil sets the value for Folder to be an explicit nil

### UnsetFolder
`func (o *RecentConfigDto) UnsetFolder()`

UnsetFolder ensures that no value is present for Folder, not even an explicit nil
### GetTitle

`func (o *RecentConfigDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RecentConfigDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RecentConfigDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *RecentConfigDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *RecentConfigDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *RecentConfigDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetUrl

`func (o *RecentConfigDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RecentConfigDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RecentConfigDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *RecentConfigDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *RecentConfigDto) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *RecentConfigDto) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


