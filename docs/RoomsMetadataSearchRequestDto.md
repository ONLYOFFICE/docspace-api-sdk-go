# RoomsMetadataSearchRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetadataTemplateId** | Pointer to **NullableInt32** | The ID of the metadata template the rooms must be assigned to. On its own it narrows the listing to the rooms  carrying the template; together with the conditions it also pins the template the filtered fields belong to. | [optional] 
**MetadataFilters** | Pointer to [**[]MetadataFilterConditionRequest**](MetadataFilterConditionRequest.md) | The metadata filter conditions, combined with AND. A custom field is addressed by its name instead of the field ID. | [optional] 
**IncludeHiddenMetadata** | Pointer to **NullableBool** | Whether the text search, and the metadataMatches it reports, also look at the values of the hidden metadata  templates (created with visible false for the metadata a service keeps for itself, which the web client never  shows). False, the default, leaves them out, so an entry is not found by such a value; true takes them in.  metadataTemplateId and metadataFilters are not concerned: a hidden template or field named there is searched  either way. | [optional] 
**FilterValue** | Pointer to **NullableString** | The text to search for in the room titles and in the string metadata values of the rooms, the custom fields and  the template fields alike; a room found by its metadata says where in metadataMatches. | [optional] 
**SearchArea** | Pointer to [**SearchArea**](SearchArea.md) | The section to search in: the active rooms (the default), the archive or the templates. | [optional] 
**Type** | Pointer to [**[]RoomType**](RoomType.md) | The room types to search among. | [optional] 
**Count** | Pointer to **int32** | The number of rooms to return, from 1 to 100. | [optional] 
**StartIndex** | Pointer to **int32** | The zero-based index of the first room to return. | [optional] 
**SortBy** | Pointer to **NullableString** | The field to sort by, a name of the SortedByType values. | [optional] 
**SortOrder** | Pointer to [**SortOrder**](SortOrder.md) | The sort order. | [optional] 

## Methods

### NewRoomsMetadataSearchRequestDto

`func NewRoomsMetadataSearchRequestDto() *RoomsMetadataSearchRequestDto`

NewRoomsMetadataSearchRequestDto instantiates a new RoomsMetadataSearchRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoomsMetadataSearchRequestDtoWithDefaults

`func NewRoomsMetadataSearchRequestDtoWithDefaults() *RoomsMetadataSearchRequestDto`

NewRoomsMetadataSearchRequestDtoWithDefaults instantiates a new RoomsMetadataSearchRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadataTemplateId

`func (o *RoomsMetadataSearchRequestDto) GetMetadataTemplateId() int32`

GetMetadataTemplateId returns the MetadataTemplateId field if non-nil, zero value otherwise.

### GetMetadataTemplateIdOk

`func (o *RoomsMetadataSearchRequestDto) GetMetadataTemplateIdOk() (*int32, bool)`

GetMetadataTemplateIdOk returns a tuple with the MetadataTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataTemplateId

`func (o *RoomsMetadataSearchRequestDto) SetMetadataTemplateId(v int32)`

SetMetadataTemplateId sets MetadataTemplateId field to given value.

### HasMetadataTemplateId

`func (o *RoomsMetadataSearchRequestDto) HasMetadataTemplateId() bool`

HasMetadataTemplateId returns a boolean if a field has been set.

### SetMetadataTemplateIdNil

`func (o *RoomsMetadataSearchRequestDto) SetMetadataTemplateIdNil(b bool)`

 SetMetadataTemplateIdNil sets the value for MetadataTemplateId to be an explicit nil

### UnsetMetadataTemplateId
`func (o *RoomsMetadataSearchRequestDto) UnsetMetadataTemplateId()`

UnsetMetadataTemplateId ensures that no value is present for MetadataTemplateId, not even an explicit nil
### GetMetadataFilters

`func (o *RoomsMetadataSearchRequestDto) GetMetadataFilters() []MetadataFilterConditionRequest`

GetMetadataFilters returns the MetadataFilters field if non-nil, zero value otherwise.

### GetMetadataFiltersOk

`func (o *RoomsMetadataSearchRequestDto) GetMetadataFiltersOk() (*[]MetadataFilterConditionRequest, bool)`

GetMetadataFiltersOk returns a tuple with the MetadataFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataFilters

`func (o *RoomsMetadataSearchRequestDto) SetMetadataFilters(v []MetadataFilterConditionRequest)`

SetMetadataFilters sets MetadataFilters field to given value.

### HasMetadataFilters

`func (o *RoomsMetadataSearchRequestDto) HasMetadataFilters() bool`

HasMetadataFilters returns a boolean if a field has been set.

### SetMetadataFiltersNil

`func (o *RoomsMetadataSearchRequestDto) SetMetadataFiltersNil(b bool)`

 SetMetadataFiltersNil sets the value for MetadataFilters to be an explicit nil

### UnsetMetadataFilters
`func (o *RoomsMetadataSearchRequestDto) UnsetMetadataFilters()`

UnsetMetadataFilters ensures that no value is present for MetadataFilters, not even an explicit nil
### GetIncludeHiddenMetadata

`func (o *RoomsMetadataSearchRequestDto) GetIncludeHiddenMetadata() bool`

GetIncludeHiddenMetadata returns the IncludeHiddenMetadata field if non-nil, zero value otherwise.

### GetIncludeHiddenMetadataOk

`func (o *RoomsMetadataSearchRequestDto) GetIncludeHiddenMetadataOk() (*bool, bool)`

GetIncludeHiddenMetadataOk returns a tuple with the IncludeHiddenMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeHiddenMetadata

`func (o *RoomsMetadataSearchRequestDto) SetIncludeHiddenMetadata(v bool)`

SetIncludeHiddenMetadata sets IncludeHiddenMetadata field to given value.

### HasIncludeHiddenMetadata

`func (o *RoomsMetadataSearchRequestDto) HasIncludeHiddenMetadata() bool`

HasIncludeHiddenMetadata returns a boolean if a field has been set.

### SetIncludeHiddenMetadataNil

`func (o *RoomsMetadataSearchRequestDto) SetIncludeHiddenMetadataNil(b bool)`

 SetIncludeHiddenMetadataNil sets the value for IncludeHiddenMetadata to be an explicit nil

### UnsetIncludeHiddenMetadata
`func (o *RoomsMetadataSearchRequestDto) UnsetIncludeHiddenMetadata()`

UnsetIncludeHiddenMetadata ensures that no value is present for IncludeHiddenMetadata, not even an explicit nil
### GetFilterValue

`func (o *RoomsMetadataSearchRequestDto) GetFilterValue() string`

GetFilterValue returns the FilterValue field if non-nil, zero value otherwise.

### GetFilterValueOk

`func (o *RoomsMetadataSearchRequestDto) GetFilterValueOk() (*string, bool)`

GetFilterValueOk returns a tuple with the FilterValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterValue

`func (o *RoomsMetadataSearchRequestDto) SetFilterValue(v string)`

SetFilterValue sets FilterValue field to given value.

### HasFilterValue

`func (o *RoomsMetadataSearchRequestDto) HasFilterValue() bool`

HasFilterValue returns a boolean if a field has been set.

### SetFilterValueNil

`func (o *RoomsMetadataSearchRequestDto) SetFilterValueNil(b bool)`

 SetFilterValueNil sets the value for FilterValue to be an explicit nil

### UnsetFilterValue
`func (o *RoomsMetadataSearchRequestDto) UnsetFilterValue()`

UnsetFilterValue ensures that no value is present for FilterValue, not even an explicit nil
### GetSearchArea

`func (o *RoomsMetadataSearchRequestDto) GetSearchArea() SearchArea`

GetSearchArea returns the SearchArea field if non-nil, zero value otherwise.

### GetSearchAreaOk

`func (o *RoomsMetadataSearchRequestDto) GetSearchAreaOk() (*SearchArea, bool)`

GetSearchAreaOk returns a tuple with the SearchArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchArea

`func (o *RoomsMetadataSearchRequestDto) SetSearchArea(v SearchArea)`

SetSearchArea sets SearchArea field to given value.

### HasSearchArea

`func (o *RoomsMetadataSearchRequestDto) HasSearchArea() bool`

HasSearchArea returns a boolean if a field has been set.

### GetType

`func (o *RoomsMetadataSearchRequestDto) GetType() []RoomType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoomsMetadataSearchRequestDto) GetTypeOk() (*[]RoomType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoomsMetadataSearchRequestDto) SetType(v []RoomType)`

SetType sets Type field to given value.

### HasType

`func (o *RoomsMetadataSearchRequestDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *RoomsMetadataSearchRequestDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *RoomsMetadataSearchRequestDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetCount

`func (o *RoomsMetadataSearchRequestDto) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *RoomsMetadataSearchRequestDto) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *RoomsMetadataSearchRequestDto) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *RoomsMetadataSearchRequestDto) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetStartIndex

`func (o *RoomsMetadataSearchRequestDto) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *RoomsMetadataSearchRequestDto) GetStartIndexOk() (*int32, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *RoomsMetadataSearchRequestDto) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.

### HasStartIndex

`func (o *RoomsMetadataSearchRequestDto) HasStartIndex() bool`

HasStartIndex returns a boolean if a field has been set.

### GetSortBy

`func (o *RoomsMetadataSearchRequestDto) GetSortBy() string`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *RoomsMetadataSearchRequestDto) GetSortByOk() (*string, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *RoomsMetadataSearchRequestDto) SetSortBy(v string)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *RoomsMetadataSearchRequestDto) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.

### SetSortByNil

`func (o *RoomsMetadataSearchRequestDto) SetSortByNil(b bool)`

 SetSortByNil sets the value for SortBy to be an explicit nil

### UnsetSortBy
`func (o *RoomsMetadataSearchRequestDto) UnsetSortBy()`

UnsetSortBy ensures that no value is present for SortBy, not even an explicit nil
### GetSortOrder

`func (o *RoomsMetadataSearchRequestDto) GetSortOrder() SortOrder`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *RoomsMetadataSearchRequestDto) GetSortOrderOk() (*SortOrder, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *RoomsMetadataSearchRequestDto) SetSortOrder(v SortOrder)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *RoomsMetadataSearchRequestDto) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


