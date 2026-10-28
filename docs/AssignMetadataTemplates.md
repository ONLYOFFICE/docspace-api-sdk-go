# AssignMetadataTemplates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateIds** | **[]int32** | The metadata template IDs. | 
**Cascade** | Pointer to **bool** | Specifies if the templates are propagated to the folder sub-entries. | [optional] 
**ConflictResolveType** | Pointer to [**MetadataConflictResolveType**](MetadataConflictResolveType.md) | The conflict resolve type for the cascade assignment. | [optional] 

## Methods

### NewAssignMetadataTemplates

`func NewAssignMetadataTemplates(templateIds []int32, ) *AssignMetadataTemplates`

NewAssignMetadataTemplates instantiates a new AssignMetadataTemplates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignMetadataTemplatesWithDefaults

`func NewAssignMetadataTemplatesWithDefaults() *AssignMetadataTemplates`

NewAssignMetadataTemplatesWithDefaults instantiates a new AssignMetadataTemplates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateIds

`func (o *AssignMetadataTemplates) GetTemplateIds() []int32`

GetTemplateIds returns the TemplateIds field if non-nil, zero value otherwise.

### GetTemplateIdsOk

`func (o *AssignMetadataTemplates) GetTemplateIdsOk() (*[]int32, bool)`

GetTemplateIdsOk returns a tuple with the TemplateIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateIds

`func (o *AssignMetadataTemplates) SetTemplateIds(v []int32)`

SetTemplateIds sets TemplateIds field to given value.


### SetTemplateIdsNil

`func (o *AssignMetadataTemplates) SetTemplateIdsNil(b bool)`

 SetTemplateIdsNil sets the value for TemplateIds to be an explicit nil

### UnsetTemplateIds
`func (o *AssignMetadataTemplates) UnsetTemplateIds()`

UnsetTemplateIds ensures that no value is present for TemplateIds, not even an explicit nil
### GetCascade

`func (o *AssignMetadataTemplates) GetCascade() bool`

GetCascade returns the Cascade field if non-nil, zero value otherwise.

### GetCascadeOk

`func (o *AssignMetadataTemplates) GetCascadeOk() (*bool, bool)`

GetCascadeOk returns a tuple with the Cascade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCascade

`func (o *AssignMetadataTemplates) SetCascade(v bool)`

SetCascade sets Cascade field to given value.

### HasCascade

`func (o *AssignMetadataTemplates) HasCascade() bool`

HasCascade returns a boolean if a field has been set.

### GetConflictResolveType

`func (o *AssignMetadataTemplates) GetConflictResolveType() MetadataConflictResolveType`

GetConflictResolveType returns the ConflictResolveType field if non-nil, zero value otherwise.

### GetConflictResolveTypeOk

`func (o *AssignMetadataTemplates) GetConflictResolveTypeOk() (*MetadataConflictResolveType, bool)`

GetConflictResolveTypeOk returns a tuple with the ConflictResolveType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictResolveType

`func (o *AssignMetadataTemplates) SetConflictResolveType(v MetadataConflictResolveType)`

SetConflictResolveType sets ConflictResolveType field to given value.

### HasConflictResolveType

`func (o *AssignMetadataTemplates) HasConflictResolveType() bool`

HasConflictResolveType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


