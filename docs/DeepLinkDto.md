# DeepLinkDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AndroidPackageName** | **NullableString** | The package name to look for on Android, and to build a store link from when the application is missing.  All three fields are empty strings on an installation that ships no mobile application, which is the  signal to keep opening links in the browser. | 
**Url** | **NullableString** | The address the client redirects a portal link through so that the application can claim it. It is the  installation's own deep-link host, not a link to any particular document. | 
**IosPackageId** | **NullableString** | The bundle identifier to look for on iOS, used the same way as `androidPackageName`. | 

## Methods

### NewDeepLinkDto

`func NewDeepLinkDto(androidPackageName NullableString, url NullableString, iosPackageId NullableString, ) *DeepLinkDto`

NewDeepLinkDto instantiates a new DeepLinkDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeepLinkDtoWithDefaults

`func NewDeepLinkDtoWithDefaults() *DeepLinkDto`

NewDeepLinkDtoWithDefaults instantiates a new DeepLinkDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAndroidPackageName

`func (o *DeepLinkDto) GetAndroidPackageName() string`

GetAndroidPackageName returns the AndroidPackageName field if non-nil, zero value otherwise.

### GetAndroidPackageNameOk

`func (o *DeepLinkDto) GetAndroidPackageNameOk() (*string, bool)`

GetAndroidPackageNameOk returns a tuple with the AndroidPackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidPackageName

`func (o *DeepLinkDto) SetAndroidPackageName(v string)`

SetAndroidPackageName sets AndroidPackageName field to given value.


### SetAndroidPackageNameNil

`func (o *DeepLinkDto) SetAndroidPackageNameNil(b bool)`

 SetAndroidPackageNameNil sets the value for AndroidPackageName to be an explicit nil

### UnsetAndroidPackageName
`func (o *DeepLinkDto) UnsetAndroidPackageName()`

UnsetAndroidPackageName ensures that no value is present for AndroidPackageName, not even an explicit nil
### GetUrl

`func (o *DeepLinkDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DeepLinkDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DeepLinkDto) SetUrl(v string)`

SetUrl sets Url field to given value.


### SetUrlNil

`func (o *DeepLinkDto) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *DeepLinkDto) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetIosPackageId

`func (o *DeepLinkDto) GetIosPackageId() string`

GetIosPackageId returns the IosPackageId field if non-nil, zero value otherwise.

### GetIosPackageIdOk

`func (o *DeepLinkDto) GetIosPackageIdOk() (*string, bool)`

GetIosPackageIdOk returns a tuple with the IosPackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosPackageId

`func (o *DeepLinkDto) SetIosPackageId(v string)`

SetIosPackageId sets IosPackageId field to given value.


### SetIosPackageIdNil

`func (o *DeepLinkDto) SetIosPackageIdNil(b bool)`

 SetIosPackageIdNil sets the value for IosPackageId to be an explicit nil

### UnsetIosPackageId
`func (o *DeepLinkDto) UnsetIosPackageId()`

UnsetIosPackageId ensures that no value is present for IosPackageId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


