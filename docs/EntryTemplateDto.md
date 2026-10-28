# EntryTemplateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The template ID. | [optional] 
**Name** | Pointer to **NullableString** | The template name. | [optional] 
**Visible** | Pointer to **bool** | Specifies if the template is visible in the UI pickers. | [optional] 
**Cascade** | Pointer to **bool** | Whether the template cascades from this entry to the folders and files below it. Only a folder or a room can  cascade; on a file the value is always false. | [optional] 
**ConflictResolveType** | Pointer to [**MetadataConflictResolveType**](MetadataConflictResolveType.md) | How the cascade of this entry treats a value an entry below already holds. Null while the template does not  cascade from this entry. | [optional] 
**Fields** | Pointer to [**[]EntryFieldDto**](EntryFieldDto.md) | The template fields with their values on the entry. | [optional] 

## Methods

### NewEntryTemplateDto

`func NewEntryTemplateDto() *EntryTemplateDto`

NewEntryTemplateDto instantiates a new EntryTemplateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntryTemplateDtoWithDefaults

`func NewEntryTemplateDtoWithDefaults() *EntryTemplateDto`

NewEntryTemplateDtoWithDefaults instantiates a new EntryTemplateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EntryTemplateDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntryTemplateDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntryTemplateDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EntryTemplateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *EntryTemplateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntryTemplateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntryTemplateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EntryTemplateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *EntryTemplateDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *EntryTemplateDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetVisible

`func (o *EntryTemplateDto) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *EntryTemplateDto) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *EntryTemplateDto) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *EntryTemplateDto) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetCascade

`func (o *EntryTemplateDto) GetCascade() bool`

GetCascade returns the Cascade field if non-nil, zero value otherwise.

### GetCascadeOk

`func (o *EntryTemplateDto) GetCascadeOk() (*bool, bool)`

GetCascadeOk returns a tuple with the Cascade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCascade

`func (o *EntryTemplateDto) SetCascade(v bool)`

SetCascade sets Cascade field to given value.

### HasCascade

`func (o *EntryTemplateDto) HasCascade() bool`

HasCascade returns a boolean if a field has been set.

### GetConflictResolveType

`func (o *EntryTemplateDto) GetConflictResolveType() MetadataConflictResolveType`

GetConflictResolveType returns the ConflictResolveType field if non-nil, zero value otherwise.

### GetConflictResolveTypeOk

`func (o *EntryTemplateDto) GetConflictResolveTypeOk() (*MetadataConflictResolveType, bool)`

GetConflictResolveTypeOk returns a tuple with the ConflictResolveType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictResolveType

`func (o *EntryTemplateDto) SetConflictResolveType(v MetadataConflictResolveType)`

SetConflictResolveType sets ConflictResolveType field to given value.

### HasConflictResolveType

`func (o *EntryTemplateDto) HasConflictResolveType() bool`

HasConflictResolveType returns a boolean if a field has been set.

### GetFields

`func (o *EntryTemplateDto) GetFields() []EntryFieldDto`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *EntryTemplateDto) GetFieldsOk() (*[]EntryFieldDto, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *EntryTemplateDto) SetFields(v []EntryFieldDto)`

SetFields sets Fields field to given value.

### HasFields

`func (o *EntryTemplateDto) HasFields() bool`

HasFields returns a boolean if a field has been set.

### SetFieldsNil

`func (o *EntryTemplateDto) SetFieldsNil(b bool)`

 SetFieldsNil sets the value for Fields to be an explicit nil

### UnsetFields
`func (o *EntryTemplateDto) UnsetFields()`

UnsetFields ensures that no value is present for Fields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


