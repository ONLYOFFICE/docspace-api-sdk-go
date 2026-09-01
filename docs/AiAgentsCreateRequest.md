# AiAgentsCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | **string** | Profile id bound to the agent. | 
**Prompt** | **string** | Agent system prompt; stored as the room's `chatSettings.prompt`. | 
**Private** | Pointer to **bool** | Whether the agent room is private. | [optional] 
**Share** | Pointer to **[]map[string]interface{}** | Initial share entries (`FileShareParams`). | [optional] 
**AttachDefaultTools** | Pointer to **bool** | Whether to attach the default DocSpace MCP tool server. | [optional] 
**Title** | Pointer to **string** | Agent (room) title. | [optional] 
**Quota** | Pointer to **float32** | Room quota in bytes. | [optional] 
**Indexing** | Pointer to **bool** | Whether room content is indexed for search. | [optional] 
**DenyDownload** | Pointer to **bool** | Whether downloading room content is denied. | [optional] 
**Lifetime** | Pointer to **map[string]interface{}** | Room data lifetime policy (`RoomDataLifetimeDto`). | [optional] 
**Watermark** | Pointer to **map[string]interface{}** | Watermark settings (`WatermarkRequestDto`). | [optional] 
**Logo** | Pointer to **map[string]interface{}** | Room logo (`LogoRequest`). | [optional] 
**Tags** | Pointer to **[]string** | Room tags. | [optional] 
**Color** | Pointer to **string** | Room accent color. | [optional] 
**Cover** | Pointer to **string** | Room cover image id. | [optional] 

## Methods

### NewAiAgentsCreateRequest

`func NewAiAgentsCreateRequest(profileId string, prompt string, ) *AiAgentsCreateRequest`

NewAiAgentsCreateRequest instantiates a new AiAgentsCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiAgentsCreateRequestWithDefaults

`func NewAiAgentsCreateRequestWithDefaults() *AiAgentsCreateRequest`

NewAiAgentsCreateRequestWithDefaults instantiates a new AiAgentsCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *AiAgentsCreateRequest) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *AiAgentsCreateRequest) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *AiAgentsCreateRequest) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetPrompt

`func (o *AiAgentsCreateRequest) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *AiAgentsCreateRequest) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *AiAgentsCreateRequest) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.


### GetPrivate

`func (o *AiAgentsCreateRequest) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *AiAgentsCreateRequest) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *AiAgentsCreateRequest) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *AiAgentsCreateRequest) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetShare

`func (o *AiAgentsCreateRequest) GetShare() []map[string]interface{}`

GetShare returns the Share field if non-nil, zero value otherwise.

### GetShareOk

`func (o *AiAgentsCreateRequest) GetShareOk() (*[]map[string]interface{}, bool)`

GetShareOk returns a tuple with the Share field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShare

`func (o *AiAgentsCreateRequest) SetShare(v []map[string]interface{})`

SetShare sets Share field to given value.

### HasShare

`func (o *AiAgentsCreateRequest) HasShare() bool`

HasShare returns a boolean if a field has been set.

### GetAttachDefaultTools

`func (o *AiAgentsCreateRequest) GetAttachDefaultTools() bool`

GetAttachDefaultTools returns the AttachDefaultTools field if non-nil, zero value otherwise.

### GetAttachDefaultToolsOk

`func (o *AiAgentsCreateRequest) GetAttachDefaultToolsOk() (*bool, bool)`

GetAttachDefaultToolsOk returns a tuple with the AttachDefaultTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachDefaultTools

`func (o *AiAgentsCreateRequest) SetAttachDefaultTools(v bool)`

SetAttachDefaultTools sets AttachDefaultTools field to given value.

### HasAttachDefaultTools

`func (o *AiAgentsCreateRequest) HasAttachDefaultTools() bool`

HasAttachDefaultTools returns a boolean if a field has been set.

### GetTitle

`func (o *AiAgentsCreateRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AiAgentsCreateRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AiAgentsCreateRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AiAgentsCreateRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetQuota

`func (o *AiAgentsCreateRequest) GetQuota() float32`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *AiAgentsCreateRequest) GetQuotaOk() (*float32, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *AiAgentsCreateRequest) SetQuota(v float32)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *AiAgentsCreateRequest) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### GetIndexing

`func (o *AiAgentsCreateRequest) GetIndexing() bool`

GetIndexing returns the Indexing field if non-nil, zero value otherwise.

### GetIndexingOk

`func (o *AiAgentsCreateRequest) GetIndexingOk() (*bool, bool)`

GetIndexingOk returns a tuple with the Indexing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexing

`func (o *AiAgentsCreateRequest) SetIndexing(v bool)`

SetIndexing sets Indexing field to given value.

### HasIndexing

`func (o *AiAgentsCreateRequest) HasIndexing() bool`

HasIndexing returns a boolean if a field has been set.

### GetDenyDownload

`func (o *AiAgentsCreateRequest) GetDenyDownload() bool`

GetDenyDownload returns the DenyDownload field if non-nil, zero value otherwise.

### GetDenyDownloadOk

`func (o *AiAgentsCreateRequest) GetDenyDownloadOk() (*bool, bool)`

GetDenyDownloadOk returns a tuple with the DenyDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyDownload

`func (o *AiAgentsCreateRequest) SetDenyDownload(v bool)`

SetDenyDownload sets DenyDownload field to given value.

### HasDenyDownload

`func (o *AiAgentsCreateRequest) HasDenyDownload() bool`

HasDenyDownload returns a boolean if a field has been set.

### GetLifetime

`func (o *AiAgentsCreateRequest) GetLifetime() map[string]interface{}`

GetLifetime returns the Lifetime field if non-nil, zero value otherwise.

### GetLifetimeOk

`func (o *AiAgentsCreateRequest) GetLifetimeOk() (*map[string]interface{}, bool)`

GetLifetimeOk returns a tuple with the Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetime

`func (o *AiAgentsCreateRequest) SetLifetime(v map[string]interface{})`

SetLifetime sets Lifetime field to given value.

### HasLifetime

`func (o *AiAgentsCreateRequest) HasLifetime() bool`

HasLifetime returns a boolean if a field has been set.

### GetWatermark

`func (o *AiAgentsCreateRequest) GetWatermark() map[string]interface{}`

GetWatermark returns the Watermark field if non-nil, zero value otherwise.

### GetWatermarkOk

`func (o *AiAgentsCreateRequest) GetWatermarkOk() (*map[string]interface{}, bool)`

GetWatermarkOk returns a tuple with the Watermark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatermark

`func (o *AiAgentsCreateRequest) SetWatermark(v map[string]interface{})`

SetWatermark sets Watermark field to given value.

### HasWatermark

`func (o *AiAgentsCreateRequest) HasWatermark() bool`

HasWatermark returns a boolean if a field has been set.

### GetLogo

`func (o *AiAgentsCreateRequest) GetLogo() map[string]interface{}`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *AiAgentsCreateRequest) GetLogoOk() (*map[string]interface{}, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *AiAgentsCreateRequest) SetLogo(v map[string]interface{})`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *AiAgentsCreateRequest) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetTags

`func (o *AiAgentsCreateRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AiAgentsCreateRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AiAgentsCreateRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AiAgentsCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetColor

`func (o *AiAgentsCreateRequest) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *AiAgentsCreateRequest) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *AiAgentsCreateRequest) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *AiAgentsCreateRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetCover

`func (o *AiAgentsCreateRequest) GetCover() string`

GetCover returns the Cover field if non-nil, zero value otherwise.

### GetCoverOk

`func (o *AiAgentsCreateRequest) GetCoverOk() (*string, bool)`

GetCoverOk returns a tuple with the Cover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCover

`func (o *AiAgentsCreateRequest) SetCover(v string)`

SetCover sets Cover field to given value.

### HasCover

`func (o *AiAgentsCreateRequest) HasCover() bool`

HasCover returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


