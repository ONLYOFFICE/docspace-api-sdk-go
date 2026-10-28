# AppDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The application's stable key, declared in the installation configuration - `ai-rooms`, `docs-cloud` and  the like. It is what every other operation of this group addresses an application by, and a client maps it  to a title and an icon of its own; the portal ships no display name for it. | [optional] 
**Enabled** | Pointer to **bool** | Whether the application is switched on for this portal. It is the portal's own flag where one has been  saved, and the default the installation configuration gives the application otherwise. | [optional] 
**Settings** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewAppDto

`func NewAppDto() *AppDto`

NewAppDto instantiates a new AppDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppDtoWithDefaults

`func NewAppDtoWithDefaults() *AppDto`

NewAppDtoWithDefaults instantiates a new AppDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AppDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AppDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AppDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetEnabled

`func (o *AppDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AppDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AppDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AppDto) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSettings

`func (o *AppDto) GetSettings() interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *AppDto) GetSettingsOk() (*interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *AppDto) SetSettings(v interface{})`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *AppDto) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### SetSettingsNil

`func (o *AppDto) SetSettingsNil(b bool)`

 SetSettingsNil sets the value for Settings to be an explicit nil

### UnsetSettings
`func (o *AppDto) UnsetSettings()`

UnsetSettings ensures that no value is present for Settings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


