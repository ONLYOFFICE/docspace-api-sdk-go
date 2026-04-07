# UpdateProviderBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** | The new display title for the AI provider. If null, the title is not changed. | [optional] 
**Url** | Pointer to **NullableString** | The new API endpoint URL for the AI provider. If null, the URL is not changed. | [optional] 
**Key** | Pointer to **NullableString** | The new authentication API key for the AI provider. If null, the key is not changed. | [optional] 

## Methods

### NewUpdateProviderBody

`func NewUpdateProviderBody() *UpdateProviderBody`

NewUpdateProviderBody instantiates a new UpdateProviderBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProviderBodyWithDefaults

`func NewUpdateProviderBodyWithDefaults() *UpdateProviderBody`

NewUpdateProviderBodyWithDefaults instantiates a new UpdateProviderBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *UpdateProviderBody) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateProviderBody) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateProviderBody) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpdateProviderBody) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *UpdateProviderBody) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *UpdateProviderBody) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetUrl

`func (o *UpdateProviderBody) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UpdateProviderBody) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UpdateProviderBody) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UpdateProviderBody) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *UpdateProviderBody) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *UpdateProviderBody) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetKey

`func (o *UpdateProviderBody) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdateProviderBody) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdateProviderBody) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UpdateProviderBody) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *UpdateProviderBody) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *UpdateProviderBody) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


