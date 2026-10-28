# WebPluginDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** | The plugin's manifest name, which is what every other operation of this group addresses it by and what  makes it unique within the portal - an installation-wide plugin wins the name over a portal one. | 
**Version** | **NullableString** | The plugin's own version from its manifest. The portal does not compare it against anything; it is there  for a person to read. | 
**MinDocSpaceVersion** | Pointer to **NullableString** | The oldest portal version the plugin declares it works with. It is a claim from the manifest and is not  enforced, so a plugin can be loaded on an older portal and simply misbehave; compare it with the `version`  of `GET api/2.0/settings`. | [optional] 
**Description** | **NullableString** | The plugin's description from its manifest, in the language the manifest was written in. The translations  of it are in `descriptionLocale`. | 
**License** | **NullableString** | The licence the plugin is published under, as its manifest states it. Nothing checks it. | 
**Author** | **NullableString** | Who wrote the plugin, as its manifest states it - not the portal member who uploaded it, who is  `createBy`. | 
**HomePage** | **NullableString** | The plugin's own page, for a person to read more about it. It is empty when the manifest names none. | 
**PluginName** | **NullableString** | The global the plugin registers itself under in the browser once its script has run, which is how a  client reaches it. It is distinct from `name`, the identifier the portal uses. | 
**Scopes** | **NullableString** | Which parts of the interface the plugin hooks into, as one comma-separated string rather than a list. | 
**Image** | **NullableString** | The plugin's icon exactly as its manifest declares it, which is normally a file name inside the plugin's  own package rather than an absolute address - resolve it against the directory `url` points into. | 
**CreateBy** | [**EmployeeDto**](EmployeeDto.md) | The portal member who uploaded the plugin. For a plugin that ships with the installation it is an empty  profile, since no member put it there. | 
**CreateOn** | **time.Time** | When the plugin was uploaded. It stays at its zero value for a plugin that ships with the installation. | 
**Enabled** | **bool** | Whether the portal loads the plugin. It is the state this portal stored, so an installation-wide plugin  can be on for one portal and off for another. | 
**System** | **bool** | Whether the plugin ships with the installation rather than having been uploaded here. A system plugin  cannot be deleted through `DELETE api/2.0/settings/webplugins/{name}`, only switched off. | 
**Url** | **NullableString** | The address of the plugin's script, which a client loads to run it. It ends in a `hash` query taken from  `version`, so the address changes whenever the plugin is updated and an old one may be cached. | 
**CssUrl** | **NullableString** | The absolute address of the plugin's stylesheet, empty for a plugin that ships none. | 
**Settings** | **NullableString** | The settings string the portal keeps for the plugin, stored and returned verbatim - only the plugin knows  its shape. It is empty until `PUT api/2.0/settings/webplugins/{name}` saves one. | 
**NameLocale** | Pointer to **map[string]string** | The plugin's name translated, keyed by culture name. A culture that is missing falls back to `name`, and  the whole map is empty for a plugin that ships no translations. | [optional] 
**DescriptionLocale** | Pointer to **map[string]string** | The plugin's description translated, keyed the same way as `nameLocale` and falling back to  `description`. | [optional] 
**Runtime** | Pointer to **NullableString** | How the script at `url` is to be loaded - as an ES module or as a classic script. It is empty for a  plugin whose manifest does not say, which a client treats as a classic script. | [optional] 

## Methods

### NewWebPluginDto

`func NewWebPluginDto(name NullableString, version NullableString, description NullableString, license NullableString, author NullableString, homePage NullableString, pluginName NullableString, scopes NullableString, image NullableString, createBy EmployeeDto, createOn time.Time, enabled bool, system bool, url NullableString, cssUrl NullableString, settings NullableString, ) *WebPluginDto`

NewWebPluginDto instantiates a new WebPluginDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebPluginDtoWithDefaults

`func NewWebPluginDtoWithDefaults() *WebPluginDto`

NewWebPluginDtoWithDefaults instantiates a new WebPluginDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WebPluginDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebPluginDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebPluginDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *WebPluginDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *WebPluginDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetVersion

`func (o *WebPluginDto) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WebPluginDto) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WebPluginDto) SetVersion(v string)`

SetVersion sets Version field to given value.


### SetVersionNil

`func (o *WebPluginDto) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *WebPluginDto) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetMinDocSpaceVersion

`func (o *WebPluginDto) GetMinDocSpaceVersion() string`

GetMinDocSpaceVersion returns the MinDocSpaceVersion field if non-nil, zero value otherwise.

### GetMinDocSpaceVersionOk

`func (o *WebPluginDto) GetMinDocSpaceVersionOk() (*string, bool)`

GetMinDocSpaceVersionOk returns a tuple with the MinDocSpaceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDocSpaceVersion

`func (o *WebPluginDto) SetMinDocSpaceVersion(v string)`

SetMinDocSpaceVersion sets MinDocSpaceVersion field to given value.

### HasMinDocSpaceVersion

`func (o *WebPluginDto) HasMinDocSpaceVersion() bool`

HasMinDocSpaceVersion returns a boolean if a field has been set.

### SetMinDocSpaceVersionNil

`func (o *WebPluginDto) SetMinDocSpaceVersionNil(b bool)`

 SetMinDocSpaceVersionNil sets the value for MinDocSpaceVersion to be an explicit nil

### UnsetMinDocSpaceVersion
`func (o *WebPluginDto) UnsetMinDocSpaceVersion()`

UnsetMinDocSpaceVersion ensures that no value is present for MinDocSpaceVersion, not even an explicit nil
### GetDescription

`func (o *WebPluginDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebPluginDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebPluginDto) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *WebPluginDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WebPluginDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLicense

`func (o *WebPluginDto) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *WebPluginDto) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *WebPluginDto) SetLicense(v string)`

SetLicense sets License field to given value.


### SetLicenseNil

`func (o *WebPluginDto) SetLicenseNil(b bool)`

 SetLicenseNil sets the value for License to be an explicit nil

### UnsetLicense
`func (o *WebPluginDto) UnsetLicense()`

UnsetLicense ensures that no value is present for License, not even an explicit nil
### GetAuthor

`func (o *WebPluginDto) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *WebPluginDto) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *WebPluginDto) SetAuthor(v string)`

SetAuthor sets Author field to given value.


### SetAuthorNil

`func (o *WebPluginDto) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *WebPluginDto) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil
### GetHomePage

`func (o *WebPluginDto) GetHomePage() string`

GetHomePage returns the HomePage field if non-nil, zero value otherwise.

### GetHomePageOk

`func (o *WebPluginDto) GetHomePageOk() (*string, bool)`

GetHomePageOk returns a tuple with the HomePage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomePage

`func (o *WebPluginDto) SetHomePage(v string)`

SetHomePage sets HomePage field to given value.


### SetHomePageNil

`func (o *WebPluginDto) SetHomePageNil(b bool)`

 SetHomePageNil sets the value for HomePage to be an explicit nil

### UnsetHomePage
`func (o *WebPluginDto) UnsetHomePage()`

UnsetHomePage ensures that no value is present for HomePage, not even an explicit nil
### GetPluginName

`func (o *WebPluginDto) GetPluginName() string`

GetPluginName returns the PluginName field if non-nil, zero value otherwise.

### GetPluginNameOk

`func (o *WebPluginDto) GetPluginNameOk() (*string, bool)`

GetPluginNameOk returns a tuple with the PluginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginName

`func (o *WebPluginDto) SetPluginName(v string)`

SetPluginName sets PluginName field to given value.


### SetPluginNameNil

`func (o *WebPluginDto) SetPluginNameNil(b bool)`

 SetPluginNameNil sets the value for PluginName to be an explicit nil

### UnsetPluginName
`func (o *WebPluginDto) UnsetPluginName()`

UnsetPluginName ensures that no value is present for PluginName, not even an explicit nil
### GetScopes

`func (o *WebPluginDto) GetScopes() string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *WebPluginDto) GetScopesOk() (*string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *WebPluginDto) SetScopes(v string)`

SetScopes sets Scopes field to given value.


### SetScopesNil

`func (o *WebPluginDto) SetScopesNil(b bool)`

 SetScopesNil sets the value for Scopes to be an explicit nil

### UnsetScopes
`func (o *WebPluginDto) UnsetScopes()`

UnsetScopes ensures that no value is present for Scopes, not even an explicit nil
### GetImage

`func (o *WebPluginDto) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *WebPluginDto) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *WebPluginDto) SetImage(v string)`

SetImage sets Image field to given value.


### SetImageNil

`func (o *WebPluginDto) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *WebPluginDto) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetCreateBy

`func (o *WebPluginDto) GetCreateBy() EmployeeDto`

GetCreateBy returns the CreateBy field if non-nil, zero value otherwise.

### GetCreateByOk

`func (o *WebPluginDto) GetCreateByOk() (*EmployeeDto, bool)`

GetCreateByOk returns a tuple with the CreateBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBy

`func (o *WebPluginDto) SetCreateBy(v EmployeeDto)`

SetCreateBy sets CreateBy field to given value.


### GetCreateOn

`func (o *WebPluginDto) GetCreateOn() time.Time`

GetCreateOn returns the CreateOn field if non-nil, zero value otherwise.

### GetCreateOnOk

`func (o *WebPluginDto) GetCreateOnOk() (*time.Time, bool)`

GetCreateOnOk returns a tuple with the CreateOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateOn

`func (o *WebPluginDto) SetCreateOn(v time.Time)`

SetCreateOn sets CreateOn field to given value.


### GetEnabled

`func (o *WebPluginDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WebPluginDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WebPluginDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSystem

`func (o *WebPluginDto) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *WebPluginDto) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *WebPluginDto) SetSystem(v bool)`

SetSystem sets System field to given value.


### GetUrl

`func (o *WebPluginDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebPluginDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebPluginDto) SetUrl(v string)`

SetUrl sets Url field to given value.


### SetUrlNil

`func (o *WebPluginDto) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *WebPluginDto) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetCssUrl

`func (o *WebPluginDto) GetCssUrl() string`

GetCssUrl returns the CssUrl field if non-nil, zero value otherwise.

### GetCssUrlOk

`func (o *WebPluginDto) GetCssUrlOk() (*string, bool)`

GetCssUrlOk returns a tuple with the CssUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssUrl

`func (o *WebPluginDto) SetCssUrl(v string)`

SetCssUrl sets CssUrl field to given value.


### SetCssUrlNil

`func (o *WebPluginDto) SetCssUrlNil(b bool)`

 SetCssUrlNil sets the value for CssUrl to be an explicit nil

### UnsetCssUrl
`func (o *WebPluginDto) UnsetCssUrl()`

UnsetCssUrl ensures that no value is present for CssUrl, not even an explicit nil
### GetSettings

`func (o *WebPluginDto) GetSettings() string`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *WebPluginDto) GetSettingsOk() (*string, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *WebPluginDto) SetSettings(v string)`

SetSettings sets Settings field to given value.


### SetSettingsNil

`func (o *WebPluginDto) SetSettingsNil(b bool)`

 SetSettingsNil sets the value for Settings to be an explicit nil

### UnsetSettings
`func (o *WebPluginDto) UnsetSettings()`

UnsetSettings ensures that no value is present for Settings, not even an explicit nil
### GetNameLocale

`func (o *WebPluginDto) GetNameLocale() map[string]*string`

GetNameLocale returns the NameLocale field if non-nil, zero value otherwise.

### GetNameLocaleOk

`func (o *WebPluginDto) GetNameLocaleOk() (*map[string]*string, bool)`

GetNameLocaleOk returns a tuple with the NameLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLocale

`func (o *WebPluginDto) SetNameLocale(v map[string]*string)`

SetNameLocale sets NameLocale field to given value.

### HasNameLocale

`func (o *WebPluginDto) HasNameLocale() bool`

HasNameLocale returns a boolean if a field has been set.

### GetDescriptionLocale

`func (o *WebPluginDto) GetDescriptionLocale() map[string]*string`

GetDescriptionLocale returns the DescriptionLocale field if non-nil, zero value otherwise.

### GetDescriptionLocaleOk

`func (o *WebPluginDto) GetDescriptionLocaleOk() (*map[string]*string, bool)`

GetDescriptionLocaleOk returns a tuple with the DescriptionLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionLocale

`func (o *WebPluginDto) SetDescriptionLocale(v map[string]*string)`

SetDescriptionLocale sets DescriptionLocale field to given value.

### HasDescriptionLocale

`func (o *WebPluginDto) HasDescriptionLocale() bool`

HasDescriptionLocale returns a boolean if a field has been set.

### GetRuntime

`func (o *WebPluginDto) GetRuntime() string`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *WebPluginDto) GetRuntimeOk() (*string, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *WebPluginDto) SetRuntime(v string)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *WebPluginDto) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### SetRuntimeNil

`func (o *WebPluginDto) SetRuntimeNil(b bool)`

 SetRuntimeNil sets the value for Runtime to be an explicit nil

### UnsetRuntime
`func (o *WebPluginDto) UnsetRuntime()`

UnsetRuntime ensures that no value is present for Runtime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


