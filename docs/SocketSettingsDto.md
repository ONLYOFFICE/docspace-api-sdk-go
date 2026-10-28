# SocketSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **NullableString** | The base address of the Socket.IO hub that pushes file changes, presence and quota alerts, always with a  trailing slash. It is empty when the installation runs no hub, and a client must then fall back to  polling rather than guessing an address. The value comes from the installation's configuration and cannot  be changed through this API. | [optional] 

## Methods

### NewSocketSettingsDto

`func NewSocketSettingsDto() *SocketSettingsDto`

NewSocketSettingsDto instantiates a new SocketSettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocketSettingsDtoWithDefaults

`func NewSocketSettingsDtoWithDefaults() *SocketSettingsDto`

NewSocketSettingsDtoWithDefaults instantiates a new SocketSettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *SocketSettingsDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SocketSettingsDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SocketSettingsDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SocketSettingsDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *SocketSettingsDto) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *SocketSettingsDto) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


