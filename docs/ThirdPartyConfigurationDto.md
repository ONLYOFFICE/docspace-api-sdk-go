# ThirdPartyConfigurationDto

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
**File** | [**ThirdPartyFileDto**](ThirdPartyFileDto.md) | The file the configuration was built for, in the same shape the file listings report it. | 
**ErrorMessage** | Pointer to **NullableString** | Filled in when the document could not be prepared for opening; the rest of the configuration should then not  be handed to the editors. | [optional] 
**StartFilling** | Pointer to **NullableBool** | Whether this caller may start a filling session on the form from inside the editor. It stays empty when the  file is not a form opened where starting is possible at all. | [optional] 
**FillingStatus** | Pointer to **NullableBool** | True once the caller holds a role in the running filling session of this form. It stays empty outside a  virtual data room, where roles are the only place it is set. | [optional] 
**StartFillingMode** | Pointer to [**StartFillingMode**](StartFillingMode.md) | Which filling button the editor offers: none at all, sharing the form out for others to fill, starting a  filling session, or starting one inside the form-filling room. | [optional] 
**FillingSessionId** | Pointer to **NullableString** | Identifies the filling session this opening belongs to, and is empty when the document is not opened as part  of one. Submissions made in the editor are collected under it. | [optional] 
**QuotaExceededScope** | Pointer to [**QuotaScope**](QuotaScope.md) | Names the quota that ran out - the user, the room or the portal - and is set only when the document had to be  opened read-only because of it. | [optional] 
**GenerationToolCallState** | Pointer to [**EditorToolCallStateDto**](EditorToolCallStateDto.md) | The generation the editor should run as soon as the document opens. It is set only for a document an AI agent  produced and left waiting for its content, and is empty for every other file. | [optional] 

## Methods

### NewThirdPartyConfigurationDto

`func NewThirdPartyConfigurationDto(document DocumentConfigDto, documentType NullableString, editorConfig EditorConfigurationDto, editorType EditorType, editorUrl NullableString, file ThirdPartyFileDto, ) *ThirdPartyConfigurationDto`

NewThirdPartyConfigurationDto instantiates a new ThirdPartyConfigurationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThirdPartyConfigurationDtoWithDefaults

`func NewThirdPartyConfigurationDtoWithDefaults() *ThirdPartyConfigurationDto`

NewThirdPartyConfigurationDtoWithDefaults instantiates a new ThirdPartyConfigurationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocument

`func (o *ThirdPartyConfigurationDto) GetDocument() DocumentConfigDto`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *ThirdPartyConfigurationDto) GetDocumentOk() (*DocumentConfigDto, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *ThirdPartyConfigurationDto) SetDocument(v DocumentConfigDto)`

SetDocument sets Document field to given value.


### GetDocumentType

`func (o *ThirdPartyConfigurationDto) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *ThirdPartyConfigurationDto) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *ThirdPartyConfigurationDto) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.


### SetDocumentTypeNil

`func (o *ThirdPartyConfigurationDto) SetDocumentTypeNil(b bool)`

 SetDocumentTypeNil sets the value for DocumentType to be an explicit nil

### UnsetDocumentType
`func (o *ThirdPartyConfigurationDto) UnsetDocumentType()`

UnsetDocumentType ensures that no value is present for DocumentType, not even an explicit nil
### GetEditorConfig

`func (o *ThirdPartyConfigurationDto) GetEditorConfig() EditorConfigurationDto`

GetEditorConfig returns the EditorConfig field if non-nil, zero value otherwise.

### GetEditorConfigOk

`func (o *ThirdPartyConfigurationDto) GetEditorConfigOk() (*EditorConfigurationDto, bool)`

GetEditorConfigOk returns a tuple with the EditorConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditorConfig

`func (o *ThirdPartyConfigurationDto) SetEditorConfig(v EditorConfigurationDto)`

SetEditorConfig sets EditorConfig field to given value.


### GetEditorType

`func (o *ThirdPartyConfigurationDto) GetEditorType() EditorType`

GetEditorType returns the EditorType field if non-nil, zero value otherwise.

### GetEditorTypeOk

`func (o *ThirdPartyConfigurationDto) GetEditorTypeOk() (*EditorType, bool)`

GetEditorTypeOk returns a tuple with the EditorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditorType

`func (o *ThirdPartyConfigurationDto) SetEditorType(v EditorType)`

SetEditorType sets EditorType field to given value.


### GetEditorUrl

`func (o *ThirdPartyConfigurationDto) GetEditorUrl() string`

GetEditorUrl returns the EditorUrl field if non-nil, zero value otherwise.

### GetEditorUrlOk

`func (o *ThirdPartyConfigurationDto) GetEditorUrlOk() (*string, bool)`

GetEditorUrlOk returns a tuple with the EditorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditorUrl

`func (o *ThirdPartyConfigurationDto) SetEditorUrl(v string)`

SetEditorUrl sets EditorUrl field to given value.


### SetEditorUrlNil

`func (o *ThirdPartyConfigurationDto) SetEditorUrlNil(b bool)`

 SetEditorUrlNil sets the value for EditorUrl to be an explicit nil

### UnsetEditorUrl
`func (o *ThirdPartyConfigurationDto) UnsetEditorUrl()`

UnsetEditorUrl ensures that no value is present for EditorUrl, not even an explicit nil
### GetToken

`func (o *ThirdPartyConfigurationDto) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ThirdPartyConfigurationDto) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ThirdPartyConfigurationDto) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ThirdPartyConfigurationDto) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *ThirdPartyConfigurationDto) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *ThirdPartyConfigurationDto) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetType

`func (o *ThirdPartyConfigurationDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ThirdPartyConfigurationDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ThirdPartyConfigurationDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ThirdPartyConfigurationDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ThirdPartyConfigurationDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ThirdPartyConfigurationDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetFile

`func (o *ThirdPartyConfigurationDto) GetFile() ThirdPartyFileDto`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *ThirdPartyConfigurationDto) GetFileOk() (*ThirdPartyFileDto, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *ThirdPartyConfigurationDto) SetFile(v ThirdPartyFileDto)`

SetFile sets File field to given value.


### GetErrorMessage

`func (o *ThirdPartyConfigurationDto) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ThirdPartyConfigurationDto) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ThirdPartyConfigurationDto) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ThirdPartyConfigurationDto) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *ThirdPartyConfigurationDto) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *ThirdPartyConfigurationDto) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetStartFilling

`func (o *ThirdPartyConfigurationDto) GetStartFilling() bool`

GetStartFilling returns the StartFilling field if non-nil, zero value otherwise.

### GetStartFillingOk

`func (o *ThirdPartyConfigurationDto) GetStartFillingOk() (*bool, bool)`

GetStartFillingOk returns a tuple with the StartFilling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartFilling

`func (o *ThirdPartyConfigurationDto) SetStartFilling(v bool)`

SetStartFilling sets StartFilling field to given value.

### HasStartFilling

`func (o *ThirdPartyConfigurationDto) HasStartFilling() bool`

HasStartFilling returns a boolean if a field has been set.

### SetStartFillingNil

`func (o *ThirdPartyConfigurationDto) SetStartFillingNil(b bool)`

 SetStartFillingNil sets the value for StartFilling to be an explicit nil

### UnsetStartFilling
`func (o *ThirdPartyConfigurationDto) UnsetStartFilling()`

UnsetStartFilling ensures that no value is present for StartFilling, not even an explicit nil
### GetFillingStatus

`func (o *ThirdPartyConfigurationDto) GetFillingStatus() bool`

GetFillingStatus returns the FillingStatus field if non-nil, zero value otherwise.

### GetFillingStatusOk

`func (o *ThirdPartyConfigurationDto) GetFillingStatusOk() (*bool, bool)`

GetFillingStatusOk returns a tuple with the FillingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillingStatus

`func (o *ThirdPartyConfigurationDto) SetFillingStatus(v bool)`

SetFillingStatus sets FillingStatus field to given value.

### HasFillingStatus

`func (o *ThirdPartyConfigurationDto) HasFillingStatus() bool`

HasFillingStatus returns a boolean if a field has been set.

### SetFillingStatusNil

`func (o *ThirdPartyConfigurationDto) SetFillingStatusNil(b bool)`

 SetFillingStatusNil sets the value for FillingStatus to be an explicit nil

### UnsetFillingStatus
`func (o *ThirdPartyConfigurationDto) UnsetFillingStatus()`

UnsetFillingStatus ensures that no value is present for FillingStatus, not even an explicit nil
### GetStartFillingMode

`func (o *ThirdPartyConfigurationDto) GetStartFillingMode() StartFillingMode`

GetStartFillingMode returns the StartFillingMode field if non-nil, zero value otherwise.

### GetStartFillingModeOk

`func (o *ThirdPartyConfigurationDto) GetStartFillingModeOk() (*StartFillingMode, bool)`

GetStartFillingModeOk returns a tuple with the StartFillingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartFillingMode

`func (o *ThirdPartyConfigurationDto) SetStartFillingMode(v StartFillingMode)`

SetStartFillingMode sets StartFillingMode field to given value.

### HasStartFillingMode

`func (o *ThirdPartyConfigurationDto) HasStartFillingMode() bool`

HasStartFillingMode returns a boolean if a field has been set.

### GetFillingSessionId

`func (o *ThirdPartyConfigurationDto) GetFillingSessionId() string`

GetFillingSessionId returns the FillingSessionId field if non-nil, zero value otherwise.

### GetFillingSessionIdOk

`func (o *ThirdPartyConfigurationDto) GetFillingSessionIdOk() (*string, bool)`

GetFillingSessionIdOk returns a tuple with the FillingSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillingSessionId

`func (o *ThirdPartyConfigurationDto) SetFillingSessionId(v string)`

SetFillingSessionId sets FillingSessionId field to given value.

### HasFillingSessionId

`func (o *ThirdPartyConfigurationDto) HasFillingSessionId() bool`

HasFillingSessionId returns a boolean if a field has been set.

### SetFillingSessionIdNil

`func (o *ThirdPartyConfigurationDto) SetFillingSessionIdNil(b bool)`

 SetFillingSessionIdNil sets the value for FillingSessionId to be an explicit nil

### UnsetFillingSessionId
`func (o *ThirdPartyConfigurationDto) UnsetFillingSessionId()`

UnsetFillingSessionId ensures that no value is present for FillingSessionId, not even an explicit nil
### GetQuotaExceededScope

`func (o *ThirdPartyConfigurationDto) GetQuotaExceededScope() QuotaScope`

GetQuotaExceededScope returns the QuotaExceededScope field if non-nil, zero value otherwise.

### GetQuotaExceededScopeOk

`func (o *ThirdPartyConfigurationDto) GetQuotaExceededScopeOk() (*QuotaScope, bool)`

GetQuotaExceededScopeOk returns a tuple with the QuotaExceededScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaExceededScope

`func (o *ThirdPartyConfigurationDto) SetQuotaExceededScope(v QuotaScope)`

SetQuotaExceededScope sets QuotaExceededScope field to given value.

### HasQuotaExceededScope

`func (o *ThirdPartyConfigurationDto) HasQuotaExceededScope() bool`

HasQuotaExceededScope returns a boolean if a field has been set.

### GetGenerationToolCallState

`func (o *ThirdPartyConfigurationDto) GetGenerationToolCallState() EditorToolCallStateDto`

GetGenerationToolCallState returns the GenerationToolCallState field if non-nil, zero value otherwise.

### GetGenerationToolCallStateOk

`func (o *ThirdPartyConfigurationDto) GetGenerationToolCallStateOk() (*EditorToolCallStateDto, bool)`

GetGenerationToolCallStateOk returns a tuple with the GenerationToolCallState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationToolCallState

`func (o *ThirdPartyConfigurationDto) SetGenerationToolCallState(v EditorToolCallStateDto)`

SetGenerationToolCallState sets GenerationToolCallState field to given value.

### HasGenerationToolCallState

`func (o *ThirdPartyConfigurationDto) HasGenerationToolCallState() bool`

HasGenerationToolCallState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


