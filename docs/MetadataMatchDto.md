# MetadataMatchDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateId** | Pointer to **NullableInt32** | The ID of the metadata template the matching field belongs to. Null for a custom field, which belongs to no  template the API shows. | [optional] 
**FieldId** | Pointer to **NullableInt32** | The ID of the matching template field. Null for a custom field, which is addressed by its name alone. | [optional] 
**FieldName** | Pointer to **NullableString** | The name of the matching field: the template field or the custom field as it is shown in the entry metadata. | [optional] 
**Value** | Pointer to **NullableString** | The stored value that holds the searched text, as it was written, not only the matching part of it. | [optional] 

## Methods

### NewMetadataMatchDto

`func NewMetadataMatchDto() *MetadataMatchDto`

NewMetadataMatchDto instantiates a new MetadataMatchDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataMatchDtoWithDefaults

`func NewMetadataMatchDtoWithDefaults() *MetadataMatchDto`

NewMetadataMatchDtoWithDefaults instantiates a new MetadataMatchDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateId

`func (o *MetadataMatchDto) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *MetadataMatchDto) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *MetadataMatchDto) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *MetadataMatchDto) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### SetTemplateIdNil

`func (o *MetadataMatchDto) SetTemplateIdNil(b bool)`

 SetTemplateIdNil sets the value for TemplateId to be an explicit nil

### UnsetTemplateId
`func (o *MetadataMatchDto) UnsetTemplateId()`

UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
### GetFieldId

`func (o *MetadataMatchDto) GetFieldId() int32`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *MetadataMatchDto) GetFieldIdOk() (*int32, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *MetadataMatchDto) SetFieldId(v int32)`

SetFieldId sets FieldId field to given value.

### HasFieldId

`func (o *MetadataMatchDto) HasFieldId() bool`

HasFieldId returns a boolean if a field has been set.

### SetFieldIdNil

`func (o *MetadataMatchDto) SetFieldIdNil(b bool)`

 SetFieldIdNil sets the value for FieldId to be an explicit nil

### UnsetFieldId
`func (o *MetadataMatchDto) UnsetFieldId()`

UnsetFieldId ensures that no value is present for FieldId, not even an explicit nil
### GetFieldName

`func (o *MetadataMatchDto) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *MetadataMatchDto) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *MetadataMatchDto) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *MetadataMatchDto) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### SetFieldNameNil

`func (o *MetadataMatchDto) SetFieldNameNil(b bool)`

 SetFieldNameNil sets the value for FieldName to be an explicit nil

### UnsetFieldName
`func (o *MetadataMatchDto) UnsetFieldName()`

UnsetFieldName ensures that no value is present for FieldName, not even an explicit nil
### GetValue

`func (o *MetadataMatchDto) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MetadataMatchDto) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MetadataMatchDto) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *MetadataMatchDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *MetadataMatchDto) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *MetadataMatchDto) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


