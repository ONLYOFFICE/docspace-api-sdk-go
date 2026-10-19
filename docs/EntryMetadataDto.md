# EntryMetadataDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Templates** | Pointer to [**[]EntryTemplateDto**](EntryTemplateDto.md) | The assigned metadata templates, each field carrying its value on the entry. | [optional] 
**CustomFields** | Pointer to [**[]CustomFieldValueDto**](CustomFieldValueDto.md) | The custom fields with their values. | [optional] 

## Methods

### NewEntryMetadataDto

`func NewEntryMetadataDto() *EntryMetadataDto`

NewEntryMetadataDto instantiates a new EntryMetadataDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntryMetadataDtoWithDefaults

`func NewEntryMetadataDtoWithDefaults() *EntryMetadataDto`

NewEntryMetadataDtoWithDefaults instantiates a new EntryMetadataDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplates

`func (o *EntryMetadataDto) GetTemplates() []EntryTemplateDto`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *EntryMetadataDto) GetTemplatesOk() (*[]EntryTemplateDto, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *EntryMetadataDto) SetTemplates(v []EntryTemplateDto)`

SetTemplates sets Templates field to given value.

### HasTemplates

`func (o *EntryMetadataDto) HasTemplates() bool`

HasTemplates returns a boolean if a field has been set.

### SetTemplatesNil

`func (o *EntryMetadataDto) SetTemplatesNil(b bool)`

 SetTemplatesNil sets the value for Templates to be an explicit nil

### UnsetTemplates
`func (o *EntryMetadataDto) UnsetTemplates()`

UnsetTemplates ensures that no value is present for Templates, not even an explicit nil
### GetCustomFields

`func (o *EntryMetadataDto) GetCustomFields() []CustomFieldValueDto`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *EntryMetadataDto) GetCustomFieldsOk() (*[]CustomFieldValueDto, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *EntryMetadataDto) SetCustomFields(v []CustomFieldValueDto)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *EntryMetadataDto) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *EntryMetadataDto) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *EntryMetadataDto) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


