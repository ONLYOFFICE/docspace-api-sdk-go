# AiAgentsUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | Pointer to **string** | Profile id to rebind (optional). | [optional] 
**ChatSettings** | Pointer to **map[string]interface{}** | Chat settings (`ChatSettings`); requires a valid provider/model. | [optional] 
**SendFormToExternalDB** | Pointer to **bool** | Whether form results are sent to an external DB. | [optional] 
**SaveFormAsXLSX** | Pointer to **bool** | Whether forms are saved as XLSX. | [optional] 
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

### NewAiAgentsUpdateRequest

`func NewAiAgentsUpdateRequest() *AiAgentsUpdateRequest`

NewAiAgentsUpdateRequest instantiates a new AiAgentsUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiAgentsUpdateRequestWithDefaults

`func NewAiAgentsUpdateRequestWithDefaults() *AiAgentsUpdateRequest`

NewAiAgentsUpdateRequestWithDefaults instantiates a new AiAgentsUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *AiAgentsUpdateRequest) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *AiAgentsUpdateRequest) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *AiAgentsUpdateRequest) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *AiAgentsUpdateRequest) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetChatSettings

`func (o *AiAgentsUpdateRequest) GetChatSettings() map[string]interface{}`

GetChatSettings returns the ChatSettings field if non-nil, zero value otherwise.

### GetChatSettingsOk

`func (o *AiAgentsUpdateRequest) GetChatSettingsOk() (*map[string]interface{}, bool)`

GetChatSettingsOk returns a tuple with the ChatSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatSettings

`func (o *AiAgentsUpdateRequest) SetChatSettings(v map[string]interface{})`

SetChatSettings sets ChatSettings field to given value.

### HasChatSettings

`func (o *AiAgentsUpdateRequest) HasChatSettings() bool`

HasChatSettings returns a boolean if a field has been set.

### GetSendFormToExternalDB

`func (o *AiAgentsUpdateRequest) GetSendFormToExternalDB() bool`

GetSendFormToExternalDB returns the SendFormToExternalDB field if non-nil, zero value otherwise.

### GetSendFormToExternalDBOk

`func (o *AiAgentsUpdateRequest) GetSendFormToExternalDBOk() (*bool, bool)`

GetSendFormToExternalDBOk returns a tuple with the SendFormToExternalDB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendFormToExternalDB

`func (o *AiAgentsUpdateRequest) SetSendFormToExternalDB(v bool)`

SetSendFormToExternalDB sets SendFormToExternalDB field to given value.

### HasSendFormToExternalDB

`func (o *AiAgentsUpdateRequest) HasSendFormToExternalDB() bool`

HasSendFormToExternalDB returns a boolean if a field has been set.

### GetSaveFormAsXLSX

`func (o *AiAgentsUpdateRequest) GetSaveFormAsXLSX() bool`

GetSaveFormAsXLSX returns the SaveFormAsXLSX field if non-nil, zero value otherwise.

### GetSaveFormAsXLSXOk

`func (o *AiAgentsUpdateRequest) GetSaveFormAsXLSXOk() (*bool, bool)`

GetSaveFormAsXLSXOk returns a tuple with the SaveFormAsXLSX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveFormAsXLSX

`func (o *AiAgentsUpdateRequest) SetSaveFormAsXLSX(v bool)`

SetSaveFormAsXLSX sets SaveFormAsXLSX field to given value.

### HasSaveFormAsXLSX

`func (o *AiAgentsUpdateRequest) HasSaveFormAsXLSX() bool`

HasSaveFormAsXLSX returns a boolean if a field has been set.

### GetTitle

`func (o *AiAgentsUpdateRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AiAgentsUpdateRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AiAgentsUpdateRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AiAgentsUpdateRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetQuota

`func (o *AiAgentsUpdateRequest) GetQuota() float32`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *AiAgentsUpdateRequest) GetQuotaOk() (*float32, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *AiAgentsUpdateRequest) SetQuota(v float32)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *AiAgentsUpdateRequest) HasQuota() bool`

HasQuota returns a boolean if a field has been set.

### GetIndexing

`func (o *AiAgentsUpdateRequest) GetIndexing() bool`

GetIndexing returns the Indexing field if non-nil, zero value otherwise.

### GetIndexingOk

`func (o *AiAgentsUpdateRequest) GetIndexingOk() (*bool, bool)`

GetIndexingOk returns a tuple with the Indexing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexing

`func (o *AiAgentsUpdateRequest) SetIndexing(v bool)`

SetIndexing sets Indexing field to given value.

### HasIndexing

`func (o *AiAgentsUpdateRequest) HasIndexing() bool`

HasIndexing returns a boolean if a field has been set.

### GetDenyDownload

`func (o *AiAgentsUpdateRequest) GetDenyDownload() bool`

GetDenyDownload returns the DenyDownload field if non-nil, zero value otherwise.

### GetDenyDownloadOk

`func (o *AiAgentsUpdateRequest) GetDenyDownloadOk() (*bool, bool)`

GetDenyDownloadOk returns a tuple with the DenyDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyDownload

`func (o *AiAgentsUpdateRequest) SetDenyDownload(v bool)`

SetDenyDownload sets DenyDownload field to given value.

### HasDenyDownload

`func (o *AiAgentsUpdateRequest) HasDenyDownload() bool`

HasDenyDownload returns a boolean if a field has been set.

### GetLifetime

`func (o *AiAgentsUpdateRequest) GetLifetime() map[string]interface{}`

GetLifetime returns the Lifetime field if non-nil, zero value otherwise.

### GetLifetimeOk

`func (o *AiAgentsUpdateRequest) GetLifetimeOk() (*map[string]interface{}, bool)`

GetLifetimeOk returns a tuple with the Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetime

`func (o *AiAgentsUpdateRequest) SetLifetime(v map[string]interface{})`

SetLifetime sets Lifetime field to given value.

### HasLifetime

`func (o *AiAgentsUpdateRequest) HasLifetime() bool`

HasLifetime returns a boolean if a field has been set.

### GetWatermark

`func (o *AiAgentsUpdateRequest) GetWatermark() map[string]interface{}`

GetWatermark returns the Watermark field if non-nil, zero value otherwise.

### GetWatermarkOk

`func (o *AiAgentsUpdateRequest) GetWatermarkOk() (*map[string]interface{}, bool)`

GetWatermarkOk returns a tuple with the Watermark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatermark

`func (o *AiAgentsUpdateRequest) SetWatermark(v map[string]interface{})`

SetWatermark sets Watermark field to given value.

### HasWatermark

`func (o *AiAgentsUpdateRequest) HasWatermark() bool`

HasWatermark returns a boolean if a field has been set.

### GetLogo

`func (o *AiAgentsUpdateRequest) GetLogo() map[string]interface{}`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *AiAgentsUpdateRequest) GetLogoOk() (*map[string]interface{}, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *AiAgentsUpdateRequest) SetLogo(v map[string]interface{})`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *AiAgentsUpdateRequest) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetTags

`func (o *AiAgentsUpdateRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AiAgentsUpdateRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AiAgentsUpdateRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AiAgentsUpdateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetColor

`func (o *AiAgentsUpdateRequest) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *AiAgentsUpdateRequest) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *AiAgentsUpdateRequest) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *AiAgentsUpdateRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetCover

`func (o *AiAgentsUpdateRequest) GetCover() string`

GetCover returns the Cover field if non-nil, zero value otherwise.

### GetCoverOk

`func (o *AiAgentsUpdateRequest) GetCoverOk() (*string, bool)`

GetCoverOk returns a tuple with the Cover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCover

`func (o *AiAgentsUpdateRequest) SetCover(v string)`

SetCover sets Cover field to given value.

### HasCover

`func (o *AiAgentsUpdateRequest) HasCover() bool`

HasCover returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


