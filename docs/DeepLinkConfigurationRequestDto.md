# DeepLinkConfigurationRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeepLinkSettings** | Pointer to [**DeepLinkSettingsRequestDto**](DeepLinkSettingsRequestDto.md) | The deep link configuration to store. Only its `handlingMode` is read - whether a link always opens in the  browser, always in the native application, or asks the user each time - and a mode outside the defined set is  refused with 400 before anything is stored. | [optional] 

## Methods

### NewDeepLinkConfigurationRequestDto

`func NewDeepLinkConfigurationRequestDto() *DeepLinkConfigurationRequestDto`

NewDeepLinkConfigurationRequestDto instantiates a new DeepLinkConfigurationRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeepLinkConfigurationRequestDtoWithDefaults

`func NewDeepLinkConfigurationRequestDtoWithDefaults() *DeepLinkConfigurationRequestDto`

NewDeepLinkConfigurationRequestDtoWithDefaults instantiates a new DeepLinkConfigurationRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeepLinkSettings

`func (o *DeepLinkConfigurationRequestDto) GetDeepLinkSettings() DeepLinkSettingsRequestDto`

GetDeepLinkSettings returns the DeepLinkSettings field if non-nil, zero value otherwise.

### GetDeepLinkSettingsOk

`func (o *DeepLinkConfigurationRequestDto) GetDeepLinkSettingsOk() (*DeepLinkSettingsRequestDto, bool)`

GetDeepLinkSettingsOk returns a tuple with the DeepLinkSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeepLinkSettings

`func (o *DeepLinkConfigurationRequestDto) SetDeepLinkSettings(v DeepLinkSettingsRequestDto)`

SetDeepLinkSettings sets DeepLinkSettings field to given value.

### HasDeepLinkSettings

`func (o *DeepLinkConfigurationRequestDto) HasDeepLinkSettings() bool`

HasDeepLinkSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


