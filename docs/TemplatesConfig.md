# TemplatesConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | Pointer to **NullableString** | The absolute URL to the image for template. | [optional] 
**Title** | Pointer to **NullableString** | The template title that will be displayed in the Create New... menu option. | [optional] 
**Url** | Pointer to **NullableString** | The absolute URL to the document where it will be created and available after creation. | [optional] 

## Methods

### NewTemplatesConfig

`func NewTemplatesConfig() *TemplatesConfig`

NewTemplatesConfig instantiates a new TemplatesConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplatesConfigWithDefaults

`func NewTemplatesConfigWithDefaults() *TemplatesConfig`

NewTemplatesConfigWithDefaults instantiates a new TemplatesConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *TemplatesConfig) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *TemplatesConfig) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *TemplatesConfig) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *TemplatesConfig) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *TemplatesConfig) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *TemplatesConfig) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetTitle

`func (o *TemplatesConfig) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TemplatesConfig) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TemplatesConfig) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TemplatesConfig) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *TemplatesConfig) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *TemplatesConfig) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetUrl

`func (o *TemplatesConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TemplatesConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TemplatesConfig) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TemplatesConfig) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *TemplatesConfig) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *TemplatesConfig) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


