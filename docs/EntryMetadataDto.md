# EntryMetadataDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Templates** | Pointer to [**[]EntryTemplateDto**](EntryTemplateDto.md) | The assigned metadata templates, each field carrying its value on the entry. | [optional] 
**CustomFields** | Pointer to [**[]CustomFieldValueDto**](CustomFieldValueDto.md) | The custom fields with their values. | [optional] 
**HasMetadataFields** | Pointer to **bool** | Whether the entry has a metadata field a user can see and fill, by hand or with the AI filling: a field of a  template in `templates` whose `visible` is true, or a custom field in `customFields`. False when both lists are  empty, when every template is hidden, or when none of the visible ones has a field. It says nothing about the  values: it stays true once every field is filled. The same flag is carried by the entry in a folder listing and  in the file and folder info responses as `hasMetadataFields`, so a client that has the listing needs no call here  to decide whether to offer the filling. | [optional] 

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
### GetHasMetadataFields

`func (o *EntryMetadataDto) GetHasMetadataFields() bool`

GetHasMetadataFields returns the HasMetadataFields field if non-nil, zero value otherwise.

### GetHasMetadataFieldsOk

`func (o *EntryMetadataDto) GetHasMetadataFieldsOk() (*bool, bool)`

GetHasMetadataFieldsOk returns a tuple with the HasMetadataFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMetadataFields

`func (o *EntryMetadataDto) SetHasMetadataFields(v bool)`

SetHasMetadataFields sets HasMetadataFields field to given value.

### HasHasMetadataFields

`func (o *EntryMetadataDto) HasHasMetadataFields() bool`

HasHasMetadataFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


