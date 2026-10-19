# DeepLinkSettingsRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HandlingMode** | Pointer to [**DeepLinkHandlingMode**](DeepLinkHandlingMode.md) | Whether a link always opens in the browser, always in the native application, or asks the user each time. | [optional] 
**LastModified** | Pointer to **time.Time** | Accepted for compatibility with earlier clients and not read: the server keeps its own value. | [optional] 

## Methods

### NewDeepLinkSettingsRequestDto

`func NewDeepLinkSettingsRequestDto() *DeepLinkSettingsRequestDto`

NewDeepLinkSettingsRequestDto instantiates a new DeepLinkSettingsRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeepLinkSettingsRequestDtoWithDefaults

`func NewDeepLinkSettingsRequestDtoWithDefaults() *DeepLinkSettingsRequestDto`

NewDeepLinkSettingsRequestDtoWithDefaults instantiates a new DeepLinkSettingsRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandlingMode

`func (o *DeepLinkSettingsRequestDto) GetHandlingMode() DeepLinkHandlingMode`

GetHandlingMode returns the HandlingMode field if non-nil, zero value otherwise.

### GetHandlingModeOk

`func (o *DeepLinkSettingsRequestDto) GetHandlingModeOk() (*DeepLinkHandlingMode, bool)`

GetHandlingModeOk returns a tuple with the HandlingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlingMode

`func (o *DeepLinkSettingsRequestDto) SetHandlingMode(v DeepLinkHandlingMode)`

SetHandlingMode sets HandlingMode field to given value.

### HasHandlingMode

`func (o *DeepLinkSettingsRequestDto) HasHandlingMode() bool`

HasHandlingMode returns a boolean if a field has been set.

### GetLastModified

`func (o *DeepLinkSettingsRequestDto) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *DeepLinkSettingsRequestDto) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *DeepLinkSettingsRequestDto) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *DeepLinkSettingsRequestDto) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


