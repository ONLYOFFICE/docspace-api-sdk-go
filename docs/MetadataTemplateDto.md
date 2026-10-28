# MetadataTemplateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The template ID. | [optional] 
**Name** | Pointer to **NullableString** | The template name. | [optional] 
**Visible** | Pointer to **bool** | Specifies if the template is visible in the UI pickers. | [optional] 
**CreateBy** | Pointer to [**EmployeeDto**](EmployeeDto.md) | The user who created the template. A user who has since left the portal comes back as the lost user placeholder. | [optional] 
**CreateOn** | Pointer to [**ApiDateTime**](ApiDateTime.md) | The template creation date. | [optional] 
**ModifiedBy** | Pointer to [**EmployeeDto**](EmployeeDto.md) | The user who changed the template itself last: its name or its visibility. The creator until somebody changes  it; a change of a field is recorded on the field, not here. | [optional] 
**ModifiedOn** | Pointer to [**ApiDateTime**](ApiDateTime.md) | The date when the template was modified last. | [optional] 
**Fields** | Pointer to [**[]MetadataFieldDto**](MetadataFieldDto.md) | The template metadata fields. | [optional] 

## Methods

### NewMetadataTemplateDto

`func NewMetadataTemplateDto() *MetadataTemplateDto`

NewMetadataTemplateDto instantiates a new MetadataTemplateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataTemplateDtoWithDefaults

`func NewMetadataTemplateDtoWithDefaults() *MetadataTemplateDto`

NewMetadataTemplateDtoWithDefaults instantiates a new MetadataTemplateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MetadataTemplateDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetadataTemplateDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetadataTemplateDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MetadataTemplateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MetadataTemplateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetadataTemplateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetadataTemplateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetadataTemplateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MetadataTemplateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MetadataTemplateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetVisible

`func (o *MetadataTemplateDto) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *MetadataTemplateDto) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *MetadataTemplateDto) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *MetadataTemplateDto) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetCreateBy

`func (o *MetadataTemplateDto) GetCreateBy() EmployeeDto`

GetCreateBy returns the CreateBy field if non-nil, zero value otherwise.

### GetCreateByOk

`func (o *MetadataTemplateDto) GetCreateByOk() (*EmployeeDto, bool)`

GetCreateByOk returns a tuple with the CreateBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBy

`func (o *MetadataTemplateDto) SetCreateBy(v EmployeeDto)`

SetCreateBy sets CreateBy field to given value.

### HasCreateBy

`func (o *MetadataTemplateDto) HasCreateBy() bool`

HasCreateBy returns a boolean if a field has been set.

### GetCreateOn

`func (o *MetadataTemplateDto) GetCreateOn() ApiDateTime`

GetCreateOn returns the CreateOn field if non-nil, zero value otherwise.

### GetCreateOnOk

`func (o *MetadataTemplateDto) GetCreateOnOk() (*ApiDateTime, bool)`

GetCreateOnOk returns a tuple with the CreateOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateOn

`func (o *MetadataTemplateDto) SetCreateOn(v ApiDateTime)`

SetCreateOn sets CreateOn field to given value.

### HasCreateOn

`func (o *MetadataTemplateDto) HasCreateOn() bool`

HasCreateOn returns a boolean if a field has been set.

### GetModifiedBy

`func (o *MetadataTemplateDto) GetModifiedBy() EmployeeDto`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *MetadataTemplateDto) GetModifiedByOk() (*EmployeeDto, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *MetadataTemplateDto) SetModifiedBy(v EmployeeDto)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *MetadataTemplateDto) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetModifiedOn

`func (o *MetadataTemplateDto) GetModifiedOn() ApiDateTime`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *MetadataTemplateDto) GetModifiedOnOk() (*ApiDateTime, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *MetadataTemplateDto) SetModifiedOn(v ApiDateTime)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *MetadataTemplateDto) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### GetFields

`func (o *MetadataTemplateDto) GetFields() []MetadataFieldDto`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *MetadataTemplateDto) GetFieldsOk() (*[]MetadataFieldDto, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *MetadataTemplateDto) SetFields(v []MetadataFieldDto)`

SetFields sets Fields field to given value.

### HasFields

`func (o *MetadataTemplateDto) HasFields() bool`

HasFields returns a boolean if a field has been set.

### SetFieldsNil

`func (o *MetadataTemplateDto) SetFieldsNil(b bool)`

 SetFieldsNil sets the value for Fields to be an explicit nil

### UnsetFields
`func (o *MetadataTemplateDto) UnsetFields()`

UnsetFields ensures that no value is present for Fields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


