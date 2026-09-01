# AccountInfoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **NullableString** | The account provider. | 
**Url** | **NullableString** | The account URL. | 
**Linked** | **bool** | Specifies if an account is linked with other profiles or not. | 

## Methods

### NewAccountInfoDto

`func NewAccountInfoDto(provider NullableString, url NullableString, linked bool, ) *AccountInfoDto`

NewAccountInfoDto instantiates a new AccountInfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountInfoDtoWithDefaults

`func NewAccountInfoDtoWithDefaults() *AccountInfoDto`

NewAccountInfoDtoWithDefaults instantiates a new AccountInfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *AccountInfoDto) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AccountInfoDto) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AccountInfoDto) SetProvider(v string)`

SetProvider sets Provider field to given value.


### SetProviderNil

`func (o *AccountInfoDto) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *AccountInfoDto) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetUrl

`func (o *AccountInfoDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AccountInfoDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AccountInfoDto) SetUrl(v string)`

SetUrl sets Url field to given value.


### SetUrlNil

`func (o *AccountInfoDto) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *AccountInfoDto) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetLinked

`func (o *AccountInfoDto) GetLinked() bool`

GetLinked returns the Linked field if non-nil, zero value otherwise.

### GetLinkedOk

`func (o *AccountInfoDto) GetLinkedOk() (*bool, bool)`

GetLinkedOk returns a tuple with the Linked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinked

`func (o *AccountInfoDto) SetLinked(v bool)`

SetLinked sets Linked field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


