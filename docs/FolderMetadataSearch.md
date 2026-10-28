# FolderMetadataSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetadataTemplateId** | Pointer to **NullableInt32** | The ID of the metadata template the entries must be assigned to. On its own it narrows the listing to the entries  carrying the template; together with the conditions it also pins the template the filtered fields belong to. | [optional] 
**MetadataFilters** | Pointer to [**[]MetadataFilterConditionRequest**](MetadataFilterConditionRequest.md) | The metadata filter conditions, combined with AND. A custom field is addressed by its name instead of the field ID. | [optional] 
**IncludeHiddenMetadata** | Pointer to **NullableBool** | Whether the text search, and the metadataMatches it reports, also look at the values of the hidden metadata  templates (created with visible false for the metadata a service keeps for itself, which the web client never  shows). False, the default, leaves them out, so an entry is not found by such a value; true takes them in.  metadataTemplateId and metadataFilters are not concerned: a hidden template or field named there is searched  either way. | [optional] 
**FilterValue** | Pointer to **NullableString** | The text to search for in the titles and in the string metadata values of the entries, the custom fields and the  template fields alike; an entry found by its metadata says where in metadataMatches. | [optional] 
**WithSubFolders** | Pointer to **NullableBool** | Specifies whether to search the whole subtree of the folder (the default) or its direct children only. | [optional] 
**FilterType** | Pointer to [**FilterType**](FilterType.md) | The filter type. | [optional] 
**Count** | Pointer to **int32** | The number of entries to return, from 1 to 100. | [optional] 
**StartIndex** | Pointer to **int32** | The zero-based index of the first entry to return. | [optional] 
**SortBy** | Pointer to **NullableString** | The field to sort by, a name of the SortedByType values. | [optional] 
**SortOrder** | Pointer to [**SortOrder**](SortOrder.md) | The sort order. | [optional] 

## Methods

### NewFolderMetadataSearch

`func NewFolderMetadataSearch() *FolderMetadataSearch`

NewFolderMetadataSearch instantiates a new FolderMetadataSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFolderMetadataSearchWithDefaults

`func NewFolderMetadataSearchWithDefaults() *FolderMetadataSearch`

NewFolderMetadataSearchWithDefaults instantiates a new FolderMetadataSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadataTemplateId

`func (o *FolderMetadataSearch) GetMetadataTemplateId() int32`

GetMetadataTemplateId returns the MetadataTemplateId field if non-nil, zero value otherwise.

### GetMetadataTemplateIdOk

`func (o *FolderMetadataSearch) GetMetadataTemplateIdOk() (*int32, bool)`

GetMetadataTemplateIdOk returns a tuple with the MetadataTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataTemplateId

`func (o *FolderMetadataSearch) SetMetadataTemplateId(v int32)`

SetMetadataTemplateId sets MetadataTemplateId field to given value.

### HasMetadataTemplateId

`func (o *FolderMetadataSearch) HasMetadataTemplateId() bool`

HasMetadataTemplateId returns a boolean if a field has been set.

### SetMetadataTemplateIdNil

`func (o *FolderMetadataSearch) SetMetadataTemplateIdNil(b bool)`

 SetMetadataTemplateIdNil sets the value for MetadataTemplateId to be an explicit nil

### UnsetMetadataTemplateId
`func (o *FolderMetadataSearch) UnsetMetadataTemplateId()`

UnsetMetadataTemplateId ensures that no value is present for MetadataTemplateId, not even an explicit nil
### GetMetadataFilters

`func (o *FolderMetadataSearch) GetMetadataFilters() []MetadataFilterConditionRequest`

GetMetadataFilters returns the MetadataFilters field if non-nil, zero value otherwise.

### GetMetadataFiltersOk

`func (o *FolderMetadataSearch) GetMetadataFiltersOk() (*[]MetadataFilterConditionRequest, bool)`

GetMetadataFiltersOk returns a tuple with the MetadataFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataFilters

`func (o *FolderMetadataSearch) SetMetadataFilters(v []MetadataFilterConditionRequest)`

SetMetadataFilters sets MetadataFilters field to given value.

### HasMetadataFilters

`func (o *FolderMetadataSearch) HasMetadataFilters() bool`

HasMetadataFilters returns a boolean if a field has been set.

### SetMetadataFiltersNil

`func (o *FolderMetadataSearch) SetMetadataFiltersNil(b bool)`

 SetMetadataFiltersNil sets the value for MetadataFilters to be an explicit nil

### UnsetMetadataFilters
`func (o *FolderMetadataSearch) UnsetMetadataFilters()`

UnsetMetadataFilters ensures that no value is present for MetadataFilters, not even an explicit nil
### GetIncludeHiddenMetadata

`func (o *FolderMetadataSearch) GetIncludeHiddenMetadata() bool`

GetIncludeHiddenMetadata returns the IncludeHiddenMetadata field if non-nil, zero value otherwise.

### GetIncludeHiddenMetadataOk

`func (o *FolderMetadataSearch) GetIncludeHiddenMetadataOk() (*bool, bool)`

GetIncludeHiddenMetadataOk returns a tuple with the IncludeHiddenMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeHiddenMetadata

`func (o *FolderMetadataSearch) SetIncludeHiddenMetadata(v bool)`

SetIncludeHiddenMetadata sets IncludeHiddenMetadata field to given value.

### HasIncludeHiddenMetadata

`func (o *FolderMetadataSearch) HasIncludeHiddenMetadata() bool`

HasIncludeHiddenMetadata returns a boolean if a field has been set.

### SetIncludeHiddenMetadataNil

`func (o *FolderMetadataSearch) SetIncludeHiddenMetadataNil(b bool)`

 SetIncludeHiddenMetadataNil sets the value for IncludeHiddenMetadata to be an explicit nil

### UnsetIncludeHiddenMetadata
`func (o *FolderMetadataSearch) UnsetIncludeHiddenMetadata()`

UnsetIncludeHiddenMetadata ensures that no value is present for IncludeHiddenMetadata, not even an explicit nil
### GetFilterValue

`func (o *FolderMetadataSearch) GetFilterValue() string`

GetFilterValue returns the FilterValue field if non-nil, zero value otherwise.

### GetFilterValueOk

`func (o *FolderMetadataSearch) GetFilterValueOk() (*string, bool)`

GetFilterValueOk returns a tuple with the FilterValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterValue

`func (o *FolderMetadataSearch) SetFilterValue(v string)`

SetFilterValue sets FilterValue field to given value.

### HasFilterValue

`func (o *FolderMetadataSearch) HasFilterValue() bool`

HasFilterValue returns a boolean if a field has been set.

### SetFilterValueNil

`func (o *FolderMetadataSearch) SetFilterValueNil(b bool)`

 SetFilterValueNil sets the value for FilterValue to be an explicit nil

### UnsetFilterValue
`func (o *FolderMetadataSearch) UnsetFilterValue()`

UnsetFilterValue ensures that no value is present for FilterValue, not even an explicit nil
### GetWithSubFolders

`func (o *FolderMetadataSearch) GetWithSubFolders() bool`

GetWithSubFolders returns the WithSubFolders field if non-nil, zero value otherwise.

### GetWithSubFoldersOk

`func (o *FolderMetadataSearch) GetWithSubFoldersOk() (*bool, bool)`

GetWithSubFoldersOk returns a tuple with the WithSubFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithSubFolders

`func (o *FolderMetadataSearch) SetWithSubFolders(v bool)`

SetWithSubFolders sets WithSubFolders field to given value.

### HasWithSubFolders

`func (o *FolderMetadataSearch) HasWithSubFolders() bool`

HasWithSubFolders returns a boolean if a field has been set.

### SetWithSubFoldersNil

`func (o *FolderMetadataSearch) SetWithSubFoldersNil(b bool)`

 SetWithSubFoldersNil sets the value for WithSubFolders to be an explicit nil

### UnsetWithSubFolders
`func (o *FolderMetadataSearch) UnsetWithSubFolders()`

UnsetWithSubFolders ensures that no value is present for WithSubFolders, not even an explicit nil
### GetFilterType

`func (o *FolderMetadataSearch) GetFilterType() FilterType`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *FolderMetadataSearch) GetFilterTypeOk() (*FilterType, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *FolderMetadataSearch) SetFilterType(v FilterType)`

SetFilterType sets FilterType field to given value.

### HasFilterType

`func (o *FolderMetadataSearch) HasFilterType() bool`

HasFilterType returns a boolean if a field has been set.

### GetCount

`func (o *FolderMetadataSearch) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *FolderMetadataSearch) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *FolderMetadataSearch) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *FolderMetadataSearch) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetStartIndex

`func (o *FolderMetadataSearch) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *FolderMetadataSearch) GetStartIndexOk() (*int32, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *FolderMetadataSearch) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.

### HasStartIndex

`func (o *FolderMetadataSearch) HasStartIndex() bool`

HasStartIndex returns a boolean if a field has been set.

### GetSortBy

`func (o *FolderMetadataSearch) GetSortBy() string`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *FolderMetadataSearch) GetSortByOk() (*string, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *FolderMetadataSearch) SetSortBy(v string)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *FolderMetadataSearch) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.

### SetSortByNil

`func (o *FolderMetadataSearch) SetSortByNil(b bool)`

 SetSortByNil sets the value for SortBy to be an explicit nil

### UnsetSortBy
`func (o *FolderMetadataSearch) UnsetSortBy()`

UnsetSortBy ensures that no value is present for SortBy, not even an explicit nil
### GetSortOrder

`func (o *FolderMetadataSearch) GetSortOrder() SortOrder`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *FolderMetadataSearch) GetSortOrderOk() (*SortOrder, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *FolderMetadataSearch) SetSortOrder(v SortOrder)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *FolderMetadataSearch) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


