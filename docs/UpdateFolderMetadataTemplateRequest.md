# UpdateFolderMetadataTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cascade** | Pointer to **bool** | Whether the template cascades from the folder to the folders and files below it: true turns the cascade on,  false turns it off while the template stays on the folder. | [optional] 
**ConflictResolveType** | Pointer to [**MetadataConflictResolveType**](MetadataConflictResolveType.md) | How the cascade treats a value an entry below already holds. | [optional] 

## Methods

### NewUpdateFolderMetadataTemplateRequest

`func NewUpdateFolderMetadataTemplateRequest() *UpdateFolderMetadataTemplateRequest`

NewUpdateFolderMetadataTemplateRequest instantiates a new UpdateFolderMetadataTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFolderMetadataTemplateRequestWithDefaults

`func NewUpdateFolderMetadataTemplateRequestWithDefaults() *UpdateFolderMetadataTemplateRequest`

NewUpdateFolderMetadataTemplateRequestWithDefaults instantiates a new UpdateFolderMetadataTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCascade

`func (o *UpdateFolderMetadataTemplateRequest) GetCascade() bool`

GetCascade returns the Cascade field if non-nil, zero value otherwise.

### GetCascadeOk

`func (o *UpdateFolderMetadataTemplateRequest) GetCascadeOk() (*bool, bool)`

GetCascadeOk returns a tuple with the Cascade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCascade

`func (o *UpdateFolderMetadataTemplateRequest) SetCascade(v bool)`

SetCascade sets Cascade field to given value.

### HasCascade

`func (o *UpdateFolderMetadataTemplateRequest) HasCascade() bool`

HasCascade returns a boolean if a field has been set.

### GetConflictResolveType

`func (o *UpdateFolderMetadataTemplateRequest) GetConflictResolveType() MetadataConflictResolveType`

GetConflictResolveType returns the ConflictResolveType field if non-nil, zero value otherwise.

### GetConflictResolveTypeOk

`func (o *UpdateFolderMetadataTemplateRequest) GetConflictResolveTypeOk() (*MetadataConflictResolveType, bool)`

GetConflictResolveTypeOk returns a tuple with the ConflictResolveType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictResolveType

`func (o *UpdateFolderMetadataTemplateRequest) SetConflictResolveType(v MetadataConflictResolveType)`

SetConflictResolveType sets ConflictResolveType field to given value.

### HasConflictResolveType

`func (o *UpdateFolderMetadataTemplateRequest) HasConflictResolveType() bool`

HasConflictResolveType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


