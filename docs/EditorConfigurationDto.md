# EditorConfigurationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackUrl** | Pointer to **NullableString** | Where the editors post the document back to when they save it. A client must not call it itself; it is the  address the document service uses. | [optional] 
**CoEditing** | Pointer to [**CoEditingConfig**](CoEditingConfig.md) | How co-editing starts out for this session and whether the user may switch it in the interface. | [optional] 
**CreateUrl** | Pointer to **NullableString** | Where the editor sends the user when they ask for a new document of the same type. It is empty when creating  one is not offered here. | [optional] 
**Customization** | Pointer to [**CustomizationConfigDto**](CustomizationConfigDto.md) | How the editor interface is dressed for this portal, this document and this layout. | [optional] 
**Embedded** | Pointer to [**EmbeddedConfig**](EmbeddedConfig.md) | The addresses the framed viewer needs. It is filled in only for the embedded layout. | [optional] 
**EncryptionKeys** | Pointer to [**[]EncryptionKeyDto**](EncryptionKeyDto.md) | The caller's end-to-end encryption keys, added only when the document lies in a private room, so that the  editors can decrypt it in the browser. It is empty everywhere else. | [optional] 
**Lang** | **NullableString** | The culture the editor interface is shown in, taken from the profile of the caller. | 
**Mode** | **NullableString** | `edit` when this session may write the document, `view` when it may only read it. | 
**ModeWrite** | Pointer to **bool** | Whether this session may write; it is what the mode above says in one word. | [optional] 
**Plugins** | Pointer to [**PluginsConfig**](PluginsConfig.md) | Which editor plugins are offered. The portal currently offers none, so the list inside comes back empty. | [optional] 
**Recent** | Pointer to [**[]RecentConfig**](RecentConfig.md) | The documents offered in the editor's recent list. It is left out altogether when there is nothing to offer. | [optional] 
**Templates** | Pointer to [**[]TemplatesConfig**](TemplatesConfig.md) | Always empty: the portal no longer passes creation templates through the editor configuration. | [optional] 
**User** | Pointer to [**UserConfig**](UserConfig.md) | The account the editors attribute changes to. It is empty for an anonymous session opened through an external  link, and the editors then ask for a name themselves. | [optional] 

## Methods

### NewEditorConfigurationDto

`func NewEditorConfigurationDto(lang NullableString, mode NullableString, ) *EditorConfigurationDto`

NewEditorConfigurationDto instantiates a new EditorConfigurationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditorConfigurationDtoWithDefaults

`func NewEditorConfigurationDtoWithDefaults() *EditorConfigurationDto`

NewEditorConfigurationDtoWithDefaults instantiates a new EditorConfigurationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackUrl

`func (o *EditorConfigurationDto) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *EditorConfigurationDto) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *EditorConfigurationDto) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *EditorConfigurationDto) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### SetCallbackUrlNil

`func (o *EditorConfigurationDto) SetCallbackUrlNil(b bool)`

 SetCallbackUrlNil sets the value for CallbackUrl to be an explicit nil

### UnsetCallbackUrl
`func (o *EditorConfigurationDto) UnsetCallbackUrl()`

UnsetCallbackUrl ensures that no value is present for CallbackUrl, not even an explicit nil
### GetCoEditing

`func (o *EditorConfigurationDto) GetCoEditing() CoEditingConfig`

GetCoEditing returns the CoEditing field if non-nil, zero value otherwise.

### GetCoEditingOk

`func (o *EditorConfigurationDto) GetCoEditingOk() (*CoEditingConfig, bool)`

GetCoEditingOk returns a tuple with the CoEditing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoEditing

`func (o *EditorConfigurationDto) SetCoEditing(v CoEditingConfig)`

SetCoEditing sets CoEditing field to given value.

### HasCoEditing

`func (o *EditorConfigurationDto) HasCoEditing() bool`

HasCoEditing returns a boolean if a field has been set.

### GetCreateUrl

`func (o *EditorConfigurationDto) GetCreateUrl() string`

GetCreateUrl returns the CreateUrl field if non-nil, zero value otherwise.

### GetCreateUrlOk

`func (o *EditorConfigurationDto) GetCreateUrlOk() (*string, bool)`

GetCreateUrlOk returns a tuple with the CreateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUrl

`func (o *EditorConfigurationDto) SetCreateUrl(v string)`

SetCreateUrl sets CreateUrl field to given value.

### HasCreateUrl

`func (o *EditorConfigurationDto) HasCreateUrl() bool`

HasCreateUrl returns a boolean if a field has been set.

### SetCreateUrlNil

`func (o *EditorConfigurationDto) SetCreateUrlNil(b bool)`

 SetCreateUrlNil sets the value for CreateUrl to be an explicit nil

### UnsetCreateUrl
`func (o *EditorConfigurationDto) UnsetCreateUrl()`

UnsetCreateUrl ensures that no value is present for CreateUrl, not even an explicit nil
### GetCustomization

`func (o *EditorConfigurationDto) GetCustomization() CustomizationConfigDto`

GetCustomization returns the Customization field if non-nil, zero value otherwise.

### GetCustomizationOk

`func (o *EditorConfigurationDto) GetCustomizationOk() (*CustomizationConfigDto, bool)`

GetCustomizationOk returns a tuple with the Customization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomization

`func (o *EditorConfigurationDto) SetCustomization(v CustomizationConfigDto)`

SetCustomization sets Customization field to given value.

### HasCustomization

`func (o *EditorConfigurationDto) HasCustomization() bool`

HasCustomization returns a boolean if a field has been set.

### GetEmbedded

`func (o *EditorConfigurationDto) GetEmbedded() EmbeddedConfig`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *EditorConfigurationDto) GetEmbeddedOk() (*EmbeddedConfig, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *EditorConfigurationDto) SetEmbedded(v EmbeddedConfig)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *EditorConfigurationDto) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetEncryptionKeys

`func (o *EditorConfigurationDto) GetEncryptionKeys() []EncryptionKeyDto`

GetEncryptionKeys returns the EncryptionKeys field if non-nil, zero value otherwise.

### GetEncryptionKeysOk

`func (o *EditorConfigurationDto) GetEncryptionKeysOk() (*[]EncryptionKeyDto, bool)`

GetEncryptionKeysOk returns a tuple with the EncryptionKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKeys

`func (o *EditorConfigurationDto) SetEncryptionKeys(v []EncryptionKeyDto)`

SetEncryptionKeys sets EncryptionKeys field to given value.

### HasEncryptionKeys

`func (o *EditorConfigurationDto) HasEncryptionKeys() bool`

HasEncryptionKeys returns a boolean if a field has been set.

### SetEncryptionKeysNil

`func (o *EditorConfigurationDto) SetEncryptionKeysNil(b bool)`

 SetEncryptionKeysNil sets the value for EncryptionKeys to be an explicit nil

### UnsetEncryptionKeys
`func (o *EditorConfigurationDto) UnsetEncryptionKeys()`

UnsetEncryptionKeys ensures that no value is present for EncryptionKeys, not even an explicit nil
### GetLang

`func (o *EditorConfigurationDto) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *EditorConfigurationDto) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *EditorConfigurationDto) SetLang(v string)`

SetLang sets Lang field to given value.


### SetLangNil

`func (o *EditorConfigurationDto) SetLangNil(b bool)`

 SetLangNil sets the value for Lang to be an explicit nil

### UnsetLang
`func (o *EditorConfigurationDto) UnsetLang()`

UnsetLang ensures that no value is present for Lang, not even an explicit nil
### GetMode

`func (o *EditorConfigurationDto) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *EditorConfigurationDto) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *EditorConfigurationDto) SetMode(v string)`

SetMode sets Mode field to given value.


### SetModeNil

`func (o *EditorConfigurationDto) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *EditorConfigurationDto) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil
### GetModeWrite

`func (o *EditorConfigurationDto) GetModeWrite() bool`

GetModeWrite returns the ModeWrite field if non-nil, zero value otherwise.

### GetModeWriteOk

`func (o *EditorConfigurationDto) GetModeWriteOk() (*bool, bool)`

GetModeWriteOk returns a tuple with the ModeWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModeWrite

`func (o *EditorConfigurationDto) SetModeWrite(v bool)`

SetModeWrite sets ModeWrite field to given value.

### HasModeWrite

`func (o *EditorConfigurationDto) HasModeWrite() bool`

HasModeWrite returns a boolean if a field has been set.

### GetPlugins

`func (o *EditorConfigurationDto) GetPlugins() PluginsConfig`

GetPlugins returns the Plugins field if non-nil, zero value otherwise.

### GetPluginsOk

`func (o *EditorConfigurationDto) GetPluginsOk() (*PluginsConfig, bool)`

GetPluginsOk returns a tuple with the Plugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugins

`func (o *EditorConfigurationDto) SetPlugins(v PluginsConfig)`

SetPlugins sets Plugins field to given value.

### HasPlugins

`func (o *EditorConfigurationDto) HasPlugins() bool`

HasPlugins returns a boolean if a field has been set.

### GetRecent

`func (o *EditorConfigurationDto) GetRecent() []RecentConfig`

GetRecent returns the Recent field if non-nil, zero value otherwise.

### GetRecentOk

`func (o *EditorConfigurationDto) GetRecentOk() (*[]RecentConfig, bool)`

GetRecentOk returns a tuple with the Recent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecent

`func (o *EditorConfigurationDto) SetRecent(v []RecentConfig)`

SetRecent sets Recent field to given value.

### HasRecent

`func (o *EditorConfigurationDto) HasRecent() bool`

HasRecent returns a boolean if a field has been set.

### SetRecentNil

`func (o *EditorConfigurationDto) SetRecentNil(b bool)`

 SetRecentNil sets the value for Recent to be an explicit nil

### UnsetRecent
`func (o *EditorConfigurationDto) UnsetRecent()`

UnsetRecent ensures that no value is present for Recent, not even an explicit nil
### GetTemplates

`func (o *EditorConfigurationDto) GetTemplates() []TemplatesConfig`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *EditorConfigurationDto) GetTemplatesOk() (*[]TemplatesConfig, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *EditorConfigurationDto) SetTemplates(v []TemplatesConfig)`

SetTemplates sets Templates field to given value.

### HasTemplates

`func (o *EditorConfigurationDto) HasTemplates() bool`

HasTemplates returns a boolean if a field has been set.

### SetTemplatesNil

`func (o *EditorConfigurationDto) SetTemplatesNil(b bool)`

 SetTemplatesNil sets the value for Templates to be an explicit nil

### UnsetTemplates
`func (o *EditorConfigurationDto) UnsetTemplates()`

UnsetTemplates ensures that no value is present for Templates, not even an explicit nil
### GetUser

`func (o *EditorConfigurationDto) GetUser() UserConfig`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *EditorConfigurationDto) GetUserOk() (*UserConfig, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *EditorConfigurationDto) SetUser(v UserConfig)`

SetUser sets User field to given value.

### HasUser

`func (o *EditorConfigurationDto) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


