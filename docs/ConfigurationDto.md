# ConfigurationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Document** | [**DocumentConfigDto**](DocumentConfigDto.md) | The document as the editors address it: its revision key, title, type, download address and the permissions of  this caller on it. | 
**DocumentType** | **NullableString** | The editor family the file opens in - `word`, `cell`, `slide`, `pdf` or `diagram`. It comes back empty for a  format no editor handles. | 
**EditorConfig** | [**EditorConfigurationDto**](EditorConfigurationDto.md) | How the editor is set up for this opening: the mode, the language, the interface customization, the callback  the editors save through, and the account they attribute changes to. | 
**EditorType** | [**EditorType**](EditorType.md) | The layout the configuration was actually built for. It echoes the requested one except where the room  overruled it, as the templates folder does by forcing the embedded viewer. | 
**EditorUrl** | **NullableString** | The address of the editor api script the client has to load, with the shard key of this document already  appended. Load it as it is given rather than assembling it by hand. | 
**Token** | Pointer to **NullableString** | Signs this whole configuration so that the editors can trust it; anything a client changes in the  configuration invalidates it. It stays empty on a portal that has no signature secret configured for the  document service. | [optional] 
**Type** | Pointer to **NullableString** | The layout spelled as a lowercase word - `desktop`, `mobile` or `embedded` - the same value the editor type  carries as a number. | [optional] 
**File** | [**FileDto**](FileDto.md) | The file the configuration was built for, in the same shape the file listings report it. | 
**ErrorMessage** | Pointer to **NullableString** | Filled in when the document could not be prepared for opening; the rest of the configuration should then not  be handed to the editors. | [optional] 
**StartFilling** | Pointer to **NullableBool** | Whether this caller may start a filling session on the form from inside the editor. It stays empty when the  file is not a form opened where starting is possible at all. | [optional] 
**FillingStatus** | Pointer to **NullableBool** | True once the caller holds a role in the running filling session of this form. It stays empty outside a  virtual data room, where roles are the only place it is set. | [optional] 
**StartFillingMode** | Pointer to [**StartFillingMode**](StartFillingMode.md) | Which filling button the editor offers: none at all, sharing the form out for others to fill, starting a  filling session, or starting one inside the form-filling room. | [optional] 
**FillingSessionId** | Pointer to **NullableString** | Identifies the filling session this opening belongs to, and is empty when the document is not opened as part  of one. Submissions made in the editor are collected under it. | [optional] 
**QuotaExceededScope** | Pointer to [**QuotaScope**](QuotaScope.md) | Names the quota that ran out - the user, the room or the portal - and is set only when the document had to be  opened read-only because of it. | [optional] 
**GenerationToolCallState** | Pointer to [**EditorToolCallStateDto**](EditorToolCallStateDto.md) | The generation the editor should run as soon as the document opens. It is set only for a document an AI agent  produced and left waiting for its content, and is empty for every other file. | [optional] 

## Methods

### NewConfigurationDto

`func NewConfigurationDto(document DocumentConfigDto, documentType NullableString, editorConfig EditorConfigurationDto, editorType EditorType, editorUrl NullableString, file FileDto, ) *ConfigurationDto`

NewConfigurationDto instantiates a new ConfigurationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationDtoWithDefaults

`func NewConfigurationDtoWithDefaults() *ConfigurationDto`

NewConfigurationDtoWithDefaults instantiates a new ConfigurationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocument

`func (o *ConfigurationDto) GetDocument() DocumentConfigDto`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *ConfigurationDto) GetDocumentOk() (*DocumentConfigDto, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *ConfigurationDto) SetDocument(v DocumentConfigDto)`

SetDocument sets Document field to given value.


### GetDocumentType

`func (o *ConfigurationDto) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *ConfigurationDto) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *ConfigurationDto) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.


### SetDocumentTypeNil

`func (o *ConfigurationDto) SetDocumentTypeNil(b bool)`

 SetDocumentTypeNil sets the value for DocumentType to be an explicit nil

### UnsetDocumentType
`func (o *ConfigurationDto) UnsetDocumentType()`

UnsetDocumentType ensures that no value is present for DocumentType, not even an explicit nil
### GetEditorConfig

`func (o *ConfigurationDto) GetEditorConfig() EditorConfigurationDto`

GetEditorConfig returns the EditorConfig field if non-nil, zero value otherwise.

### GetEditorConfigOk

`func (o *ConfigurationDto) GetEditorConfigOk() (*EditorConfigurationDto, bool)`

GetEditorConfigOk returns a tuple with the EditorConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditorConfig

`func (o *ConfigurationDto) SetEditorConfig(v EditorConfigurationDto)`

SetEditorConfig sets EditorConfig field to given value.


### GetEditorType

`func (o *ConfigurationDto) GetEditorType() EditorType`

GetEditorType returns the EditorType field if non-nil, zero value otherwise.

### GetEditorTypeOk

`func (o *ConfigurationDto) GetEditorTypeOk() (*EditorType, bool)`

GetEditorTypeOk returns a tuple with the EditorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditorType

`func (o *ConfigurationDto) SetEditorType(v EditorType)`

SetEditorType sets EditorType field to given value.


### GetEditorUrl

`func (o *ConfigurationDto) GetEditorUrl() string`

GetEditorUrl returns the EditorUrl field if non-nil, zero value otherwise.

### GetEditorUrlOk

`func (o *ConfigurationDto) GetEditorUrlOk() (*string, bool)`

GetEditorUrlOk returns a tuple with the EditorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditorUrl

`func (o *ConfigurationDto) SetEditorUrl(v string)`

SetEditorUrl sets EditorUrl field to given value.


### SetEditorUrlNil

`func (o *ConfigurationDto) SetEditorUrlNil(b bool)`

 SetEditorUrlNil sets the value for EditorUrl to be an explicit nil

### UnsetEditorUrl
`func (o *ConfigurationDto) UnsetEditorUrl()`

UnsetEditorUrl ensures that no value is present for EditorUrl, not even an explicit nil
### GetToken

`func (o *ConfigurationDto) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ConfigurationDto) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ConfigurationDto) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ConfigurationDto) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *ConfigurationDto) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *ConfigurationDto) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetType

`func (o *ConfigurationDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConfigurationDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConfigurationDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ConfigurationDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ConfigurationDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ConfigurationDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetFile

`func (o *ConfigurationDto) GetFile() FileDto`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *ConfigurationDto) GetFileOk() (*FileDto, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *ConfigurationDto) SetFile(v FileDto)`

SetFile sets File field to given value.


### GetErrorMessage

`func (o *ConfigurationDto) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ConfigurationDto) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ConfigurationDto) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ConfigurationDto) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *ConfigurationDto) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *ConfigurationDto) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetStartFilling

`func (o *ConfigurationDto) GetStartFilling() bool`

GetStartFilling returns the StartFilling field if non-nil, zero value otherwise.

### GetStartFillingOk

`func (o *ConfigurationDto) GetStartFillingOk() (*bool, bool)`

GetStartFillingOk returns a tuple with the StartFilling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartFilling

`func (o *ConfigurationDto) SetStartFilling(v bool)`

SetStartFilling sets StartFilling field to given value.

### HasStartFilling

`func (o *ConfigurationDto) HasStartFilling() bool`

HasStartFilling returns a boolean if a field has been set.

### SetStartFillingNil

`func (o *ConfigurationDto) SetStartFillingNil(b bool)`

 SetStartFillingNil sets the value for StartFilling to be an explicit nil

### UnsetStartFilling
`func (o *ConfigurationDto) UnsetStartFilling()`

UnsetStartFilling ensures that no value is present for StartFilling, not even an explicit nil
### GetFillingStatus

`func (o *ConfigurationDto) GetFillingStatus() bool`

GetFillingStatus returns the FillingStatus field if non-nil, zero value otherwise.

### GetFillingStatusOk

`func (o *ConfigurationDto) GetFillingStatusOk() (*bool, bool)`

GetFillingStatusOk returns a tuple with the FillingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillingStatus

`func (o *ConfigurationDto) SetFillingStatus(v bool)`

SetFillingStatus sets FillingStatus field to given value.

### HasFillingStatus

`func (o *ConfigurationDto) HasFillingStatus() bool`

HasFillingStatus returns a boolean if a field has been set.

### SetFillingStatusNil

`func (o *ConfigurationDto) SetFillingStatusNil(b bool)`

 SetFillingStatusNil sets the value for FillingStatus to be an explicit nil

### UnsetFillingStatus
`func (o *ConfigurationDto) UnsetFillingStatus()`

UnsetFillingStatus ensures that no value is present for FillingStatus, not even an explicit nil
### GetStartFillingMode

`func (o *ConfigurationDto) GetStartFillingMode() StartFillingMode`

GetStartFillingMode returns the StartFillingMode field if non-nil, zero value otherwise.

### GetStartFillingModeOk

`func (o *ConfigurationDto) GetStartFillingModeOk() (*StartFillingMode, bool)`

GetStartFillingModeOk returns a tuple with the StartFillingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartFillingMode

`func (o *ConfigurationDto) SetStartFillingMode(v StartFillingMode)`

SetStartFillingMode sets StartFillingMode field to given value.

### HasStartFillingMode

`func (o *ConfigurationDto) HasStartFillingMode() bool`

HasStartFillingMode returns a boolean if a field has been set.

### GetFillingSessionId

`func (o *ConfigurationDto) GetFillingSessionId() string`

GetFillingSessionId returns the FillingSessionId field if non-nil, zero value otherwise.

### GetFillingSessionIdOk

`func (o *ConfigurationDto) GetFillingSessionIdOk() (*string, bool)`

GetFillingSessionIdOk returns a tuple with the FillingSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillingSessionId

`func (o *ConfigurationDto) SetFillingSessionId(v string)`

SetFillingSessionId sets FillingSessionId field to given value.

### HasFillingSessionId

`func (o *ConfigurationDto) HasFillingSessionId() bool`

HasFillingSessionId returns a boolean if a field has been set.

### SetFillingSessionIdNil

`func (o *ConfigurationDto) SetFillingSessionIdNil(b bool)`

 SetFillingSessionIdNil sets the value for FillingSessionId to be an explicit nil

### UnsetFillingSessionId
`func (o *ConfigurationDto) UnsetFillingSessionId()`

UnsetFillingSessionId ensures that no value is present for FillingSessionId, not even an explicit nil
### GetQuotaExceededScope

`func (o *ConfigurationDto) GetQuotaExceededScope() QuotaScope`

GetQuotaExceededScope returns the QuotaExceededScope field if non-nil, zero value otherwise.

### GetQuotaExceededScopeOk

`func (o *ConfigurationDto) GetQuotaExceededScopeOk() (*QuotaScope, bool)`

GetQuotaExceededScopeOk returns a tuple with the QuotaExceededScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaExceededScope

`func (o *ConfigurationDto) SetQuotaExceededScope(v QuotaScope)`

SetQuotaExceededScope sets QuotaExceededScope field to given value.

### HasQuotaExceededScope

`func (o *ConfigurationDto) HasQuotaExceededScope() bool`

HasQuotaExceededScope returns a boolean if a field has been set.

### GetGenerationToolCallState

`func (o *ConfigurationDto) GetGenerationToolCallState() EditorToolCallStateDto`

GetGenerationToolCallState returns the GenerationToolCallState field if non-nil, zero value otherwise.

### GetGenerationToolCallStateOk

`func (o *ConfigurationDto) GetGenerationToolCallStateOk() (*EditorToolCallStateDto, bool)`

GetGenerationToolCallStateOk returns a tuple with the GenerationToolCallState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationToolCallState

`func (o *ConfigurationDto) SetGenerationToolCallState(v EditorToolCallStateDto)`

SetGenerationToolCallState sets GenerationToolCallState field to given value.

### HasGenerationToolCallState

`func (o *ConfigurationDto) HasGenerationToolCallState() bool`

HasGenerationToolCallState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


