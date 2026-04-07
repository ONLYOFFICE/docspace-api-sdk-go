# ConfigurationDtoInteger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Document** | [**DocumentConfigDto**](DocumentConfigDto.md) |  | 
**DocumentType** | **NullableString** | The document type. | 
**EditorConfig** | [**EditorConfigurationDto**](EditorConfigurationDto.md) |  | 
**EditorType** | [**EditorType**](EditorType.md) |  | 
**EditorUrl** | **NullableString** | The editor URL. | 
**Token** | Pointer to **NullableString** | The token of the file configuration. | [optional] 
**Type** | Pointer to **NullableString** | The platform type. | [optional] 
**File** | [**FileDtoInteger**](FileDtoInteger.md) |  | 
**ErrorMessage** | Pointer to **NullableString** | The error message. | [optional] 
**StartFilling** | Pointer to **NullableBool** | Specifies if the file filling has started or not. | [optional] 
**FillingStatus** | Pointer to **NullableBool** | The file filling status. | [optional] 
**StartFillingMode** | Pointer to [**StartFillingMode**](StartFillingMode.md) |  | [optional] 
**FillingSessionId** | Pointer to **NullableString** | The file filling session ID. | [optional] 
**QuotaExceededScope** | Pointer to [**QuotaScope**](QuotaScope.md) |  | [optional] 
**GenerationToolCallState** | Pointer to [**EditorToolCallStateDto**](EditorToolCallStateDto.md) |  | [optional] 

## Methods

### NewConfigurationDtoInteger

`func NewConfigurationDtoInteger(document DocumentConfigDto, documentType NullableString, editorConfig EditorConfigurationDto, editorType EditorType, editorUrl NullableString, file FileDtoInteger, ) *ConfigurationDtoInteger`

NewConfigurationDtoInteger instantiates a new ConfigurationDtoInteger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationDtoIntegerWithDefaults

`func NewConfigurationDtoIntegerWithDefaults() *ConfigurationDtoInteger`

NewConfigurationDtoIntegerWithDefaults instantiates a new ConfigurationDtoInteger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocument

`func (o *ConfigurationDtoInteger) GetDocument() DocumentConfigDto`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *ConfigurationDtoInteger) GetDocumentOk() (*DocumentConfigDto, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *ConfigurationDtoInteger) SetDocument(v DocumentConfigDto)`

SetDocument sets Document field to given value.


### GetDocumentType

`func (o *ConfigurationDtoInteger) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *ConfigurationDtoInteger) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *ConfigurationDtoInteger) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.


### SetDocumentTypeNil

`func (o *ConfigurationDtoInteger) SetDocumentTypeNil(b bool)`

 SetDocumentTypeNil sets the value for DocumentType to be an explicit nil

### UnsetDocumentType
`func (o *ConfigurationDtoInteger) UnsetDocumentType()`

UnsetDocumentType ensures that no value is present for DocumentType, not even an explicit nil
### GetEditorConfig

`func (o *ConfigurationDtoInteger) GetEditorConfig() EditorConfigurationDto`

GetEditorConfig returns the EditorConfig field if non-nil, zero value otherwise.

### GetEditorConfigOk

`func (o *ConfigurationDtoInteger) GetEditorConfigOk() (*EditorConfigurationDto, bool)`

GetEditorConfigOk returns a tuple with the EditorConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditorConfig

`func (o *ConfigurationDtoInteger) SetEditorConfig(v EditorConfigurationDto)`

SetEditorConfig sets EditorConfig field to given value.


### GetEditorType

`func (o *ConfigurationDtoInteger) GetEditorType() EditorType`

GetEditorType returns the EditorType field if non-nil, zero value otherwise.

### GetEditorTypeOk

`func (o *ConfigurationDtoInteger) GetEditorTypeOk() (*EditorType, bool)`

GetEditorTypeOk returns a tuple with the EditorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditorType

`func (o *ConfigurationDtoInteger) SetEditorType(v EditorType)`

SetEditorType sets EditorType field to given value.


### GetEditorUrl

`func (o *ConfigurationDtoInteger) GetEditorUrl() string`

GetEditorUrl returns the EditorUrl field if non-nil, zero value otherwise.

### GetEditorUrlOk

`func (o *ConfigurationDtoInteger) GetEditorUrlOk() (*string, bool)`

GetEditorUrlOk returns a tuple with the EditorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditorUrl

`func (o *ConfigurationDtoInteger) SetEditorUrl(v string)`

SetEditorUrl sets EditorUrl field to given value.


### SetEditorUrlNil

`func (o *ConfigurationDtoInteger) SetEditorUrlNil(b bool)`

 SetEditorUrlNil sets the value for EditorUrl to be an explicit nil

### UnsetEditorUrl
`func (o *ConfigurationDtoInteger) UnsetEditorUrl()`

UnsetEditorUrl ensures that no value is present for EditorUrl, not even an explicit nil
### GetToken

`func (o *ConfigurationDtoInteger) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ConfigurationDtoInteger) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ConfigurationDtoInteger) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ConfigurationDtoInteger) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *ConfigurationDtoInteger) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *ConfigurationDtoInteger) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetType

`func (o *ConfigurationDtoInteger) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConfigurationDtoInteger) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConfigurationDtoInteger) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ConfigurationDtoInteger) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ConfigurationDtoInteger) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ConfigurationDtoInteger) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetFile

`func (o *ConfigurationDtoInteger) GetFile() FileDtoInteger`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *ConfigurationDtoInteger) GetFileOk() (*FileDtoInteger, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *ConfigurationDtoInteger) SetFile(v FileDtoInteger)`

SetFile sets File field to given value.


### GetErrorMessage

`func (o *ConfigurationDtoInteger) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ConfigurationDtoInteger) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ConfigurationDtoInteger) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ConfigurationDtoInteger) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *ConfigurationDtoInteger) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *ConfigurationDtoInteger) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetStartFilling

`func (o *ConfigurationDtoInteger) GetStartFilling() bool`

GetStartFilling returns the StartFilling field if non-nil, zero value otherwise.

### GetStartFillingOk

`func (o *ConfigurationDtoInteger) GetStartFillingOk() (*bool, bool)`

GetStartFillingOk returns a tuple with the StartFilling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartFilling

`func (o *ConfigurationDtoInteger) SetStartFilling(v bool)`

SetStartFilling sets StartFilling field to given value.

### HasStartFilling

`func (o *ConfigurationDtoInteger) HasStartFilling() bool`

HasStartFilling returns a boolean if a field has been set.

### SetStartFillingNil

`func (o *ConfigurationDtoInteger) SetStartFillingNil(b bool)`

 SetStartFillingNil sets the value for StartFilling to be an explicit nil

### UnsetStartFilling
`func (o *ConfigurationDtoInteger) UnsetStartFilling()`

UnsetStartFilling ensures that no value is present for StartFilling, not even an explicit nil
### GetFillingStatus

`func (o *ConfigurationDtoInteger) GetFillingStatus() bool`

GetFillingStatus returns the FillingStatus field if non-nil, zero value otherwise.

### GetFillingStatusOk

`func (o *ConfigurationDtoInteger) GetFillingStatusOk() (*bool, bool)`

GetFillingStatusOk returns a tuple with the FillingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillingStatus

`func (o *ConfigurationDtoInteger) SetFillingStatus(v bool)`

SetFillingStatus sets FillingStatus field to given value.

### HasFillingStatus

`func (o *ConfigurationDtoInteger) HasFillingStatus() bool`

HasFillingStatus returns a boolean if a field has been set.

### SetFillingStatusNil

`func (o *ConfigurationDtoInteger) SetFillingStatusNil(b bool)`

 SetFillingStatusNil sets the value for FillingStatus to be an explicit nil

### UnsetFillingStatus
`func (o *ConfigurationDtoInteger) UnsetFillingStatus()`

UnsetFillingStatus ensures that no value is present for FillingStatus, not even an explicit nil
### GetStartFillingMode

`func (o *ConfigurationDtoInteger) GetStartFillingMode() StartFillingMode`

GetStartFillingMode returns the StartFillingMode field if non-nil, zero value otherwise.

### GetStartFillingModeOk

`func (o *ConfigurationDtoInteger) GetStartFillingModeOk() (*StartFillingMode, bool)`

GetStartFillingModeOk returns a tuple with the StartFillingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartFillingMode

`func (o *ConfigurationDtoInteger) SetStartFillingMode(v StartFillingMode)`

SetStartFillingMode sets StartFillingMode field to given value.

### HasStartFillingMode

`func (o *ConfigurationDtoInteger) HasStartFillingMode() bool`

HasStartFillingMode returns a boolean if a field has been set.

### GetFillingSessionId

`func (o *ConfigurationDtoInteger) GetFillingSessionId() string`

GetFillingSessionId returns the FillingSessionId field if non-nil, zero value otherwise.

### GetFillingSessionIdOk

`func (o *ConfigurationDtoInteger) GetFillingSessionIdOk() (*string, bool)`

GetFillingSessionIdOk returns a tuple with the FillingSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillingSessionId

`func (o *ConfigurationDtoInteger) SetFillingSessionId(v string)`

SetFillingSessionId sets FillingSessionId field to given value.

### HasFillingSessionId

`func (o *ConfigurationDtoInteger) HasFillingSessionId() bool`

HasFillingSessionId returns a boolean if a field has been set.

### SetFillingSessionIdNil

`func (o *ConfigurationDtoInteger) SetFillingSessionIdNil(b bool)`

 SetFillingSessionIdNil sets the value for FillingSessionId to be an explicit nil

### UnsetFillingSessionId
`func (o *ConfigurationDtoInteger) UnsetFillingSessionId()`

UnsetFillingSessionId ensures that no value is present for FillingSessionId, not even an explicit nil
### GetQuotaExceededScope

`func (o *ConfigurationDtoInteger) GetQuotaExceededScope() QuotaScope`

GetQuotaExceededScope returns the QuotaExceededScope field if non-nil, zero value otherwise.

### GetQuotaExceededScopeOk

`func (o *ConfigurationDtoInteger) GetQuotaExceededScopeOk() (*QuotaScope, bool)`

GetQuotaExceededScopeOk returns a tuple with the QuotaExceededScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaExceededScope

`func (o *ConfigurationDtoInteger) SetQuotaExceededScope(v QuotaScope)`

SetQuotaExceededScope sets QuotaExceededScope field to given value.

### HasQuotaExceededScope

`func (o *ConfigurationDtoInteger) HasQuotaExceededScope() bool`

HasQuotaExceededScope returns a boolean if a field has been set.

### GetGenerationToolCallState

`func (o *ConfigurationDtoInteger) GetGenerationToolCallState() EditorToolCallStateDto`

GetGenerationToolCallState returns the GenerationToolCallState field if non-nil, zero value otherwise.

### GetGenerationToolCallStateOk

`func (o *ConfigurationDtoInteger) GetGenerationToolCallStateOk() (*EditorToolCallStateDto, bool)`

GetGenerationToolCallStateOk returns a tuple with the GenerationToolCallState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationToolCallState

`func (o *ConfigurationDtoInteger) SetGenerationToolCallState(v EditorToolCallStateDto)`

SetGenerationToolCallState sets GenerationToolCallState field to given value.

### HasGenerationToolCallState

`func (o *ConfigurationDtoInteger) HasGenerationToolCallState() bool`

HasGenerationToolCallState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


