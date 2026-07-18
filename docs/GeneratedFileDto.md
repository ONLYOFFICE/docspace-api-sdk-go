# GeneratedFileDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The unique identifier of the created file. | [optional] 
**Title** | **NullableString** | The file title, including extension. | 
**Extension** | **NullableString** | The file extension. | 

## Methods

### NewGeneratedFileDto

`func NewGeneratedFileDto(title NullableString, extension NullableString, ) *GeneratedFileDto`

NewGeneratedFileDto instantiates a new GeneratedFileDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeneratedFileDtoWithDefaults

`func NewGeneratedFileDtoWithDefaults() *GeneratedFileDto`

NewGeneratedFileDtoWithDefaults instantiates a new GeneratedFileDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GeneratedFileDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GeneratedFileDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GeneratedFileDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GeneratedFileDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *GeneratedFileDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GeneratedFileDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GeneratedFileDto) SetTitle(v string)`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *GeneratedFileDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *GeneratedFileDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetExtension

`func (o *GeneratedFileDto) GetExtension() string`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *GeneratedFileDto) GetExtensionOk() (*string, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *GeneratedFileDto) SetExtension(v string)`

SetExtension sets Extension field to given value.


### SetExtensionNil

`func (o *GeneratedFileDto) SetExtensionNil(b bool)`

 SetExtensionNil sets the value for Extension to be an explicit nil

### UnsetExtension
`func (o *GeneratedFileDto) UnsetExtension()`

UnsetExtension ensures that no value is present for Extension, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


