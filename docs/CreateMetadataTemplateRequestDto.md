# CreateMetadataTemplateRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** | The template name. | 
**Visible** | Pointer to **bool** | Specifies if the template is visible in the UI pickers. | [optional] 
**Fields** | Pointer to [**[]MetadataFieldRequest**](MetadataFieldRequest.md) | The template metadata fields. | [optional] 

## Methods

### NewCreateMetadataTemplateRequestDto

`func NewCreateMetadataTemplateRequestDto(name NullableString, ) *CreateMetadataTemplateRequestDto`

NewCreateMetadataTemplateRequestDto instantiates a new CreateMetadataTemplateRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMetadataTemplateRequestDtoWithDefaults

`func NewCreateMetadataTemplateRequestDtoWithDefaults() *CreateMetadataTemplateRequestDto`

NewCreateMetadataTemplateRequestDtoWithDefaults instantiates a new CreateMetadataTemplateRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateMetadataTemplateRequestDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateMetadataTemplateRequestDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateMetadataTemplateRequestDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CreateMetadataTemplateRequestDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateMetadataTemplateRequestDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetVisible

`func (o *CreateMetadataTemplateRequestDto) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *CreateMetadataTemplateRequestDto) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *CreateMetadataTemplateRequestDto) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *CreateMetadataTemplateRequestDto) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetFields

`func (o *CreateMetadataTemplateRequestDto) GetFields() []MetadataFieldRequest`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *CreateMetadataTemplateRequestDto) GetFieldsOk() (*[]MetadataFieldRequest, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *CreateMetadataTemplateRequestDto) SetFields(v []MetadataFieldRequest)`

SetFields sets Fields field to given value.

### HasFields

`func (o *CreateMetadataTemplateRequestDto) HasFields() bool`

HasFields returns a boolean if a field has been set.

### SetFieldsNil

`func (o *CreateMetadataTemplateRequestDto) SetFieldsNil(b bool)`

 SetFieldsNil sets the value for Fields to be an explicit nil

### UnsetFields
`func (o *CreateMetadataTemplateRequestDto) UnsetFields()`

UnsetFields ensures that no value is present for Fields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


