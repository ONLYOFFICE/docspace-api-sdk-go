# Module

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The module ID. | [optional] 
**AppName** | Pointer to **NullableString** | The module product class name. | [optional] 
**Title** | Pointer to **NullableString** | The module product class name. | [optional] 
**Link** | Pointer to **NullableString** | The URL to the module start page. | [optional] 
**IconUrl** | Pointer to **NullableString** | The module icon URL. | [optional] 
**ImageUrl** | Pointer to **NullableString** | The module large image URL. | [optional] 
**HelpUrl** | Pointer to **NullableString** | The module help URL. | [optional] 
**Description** | Pointer to **NullableString** | The module description. | [optional] 
**IsPrimary** | Pointer to **bool** | Specifies if the module is primary or not. | [optional] 

## Methods

### NewModule

`func NewModule() *Module`

NewModule instantiates a new Module object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModuleWithDefaults

`func NewModuleWithDefaults() *Module`

NewModuleWithDefaults instantiates a new Module object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Module) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Module) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Module) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Module) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAppName

`func (o *Module) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *Module) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *Module) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *Module) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### SetAppNameNil

`func (o *Module) SetAppNameNil(b bool)`

 SetAppNameNil sets the value for AppName to be an explicit nil

### UnsetAppName
`func (o *Module) UnsetAppName()`

UnsetAppName ensures that no value is present for AppName, not even an explicit nil
### GetTitle

`func (o *Module) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Module) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Module) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Module) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *Module) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *Module) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetLink

`func (o *Module) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *Module) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *Module) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *Module) HasLink() bool`

HasLink returns a boolean if a field has been set.

### SetLinkNil

`func (o *Module) SetLinkNil(b bool)`

 SetLinkNil sets the value for Link to be an explicit nil

### UnsetLink
`func (o *Module) UnsetLink()`

UnsetLink ensures that no value is present for Link, not even an explicit nil
### GetIconUrl

`func (o *Module) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *Module) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *Module) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *Module) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### SetIconUrlNil

`func (o *Module) SetIconUrlNil(b bool)`

 SetIconUrlNil sets the value for IconUrl to be an explicit nil

### UnsetIconUrl
`func (o *Module) UnsetIconUrl()`

UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
### GetImageUrl

`func (o *Module) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *Module) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *Module) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *Module) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *Module) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *Module) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetHelpUrl

`func (o *Module) GetHelpUrl() string`

GetHelpUrl returns the HelpUrl field if non-nil, zero value otherwise.

### GetHelpUrlOk

`func (o *Module) GetHelpUrlOk() (*string, bool)`

GetHelpUrlOk returns a tuple with the HelpUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpUrl

`func (o *Module) SetHelpUrl(v string)`

SetHelpUrl sets HelpUrl field to given value.

### HasHelpUrl

`func (o *Module) HasHelpUrl() bool`

HasHelpUrl returns a boolean if a field has been set.

### SetHelpUrlNil

`func (o *Module) SetHelpUrlNil(b bool)`

 SetHelpUrlNil sets the value for HelpUrl to be an explicit nil

### UnsetHelpUrl
`func (o *Module) UnsetHelpUrl()`

UnsetHelpUrl ensures that no value is present for HelpUrl, not even an explicit nil
### GetDescription

`func (o *Module) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Module) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Module) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Module) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Module) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Module) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsPrimary

`func (o *Module) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *Module) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *Module) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *Module) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


