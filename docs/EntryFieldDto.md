# EntryFieldDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The field ID. | [optional] 
**Name** | Pointer to **NullableString** | The field name. | [optional] 
**Type** | Pointer to [**MetadataFieldType**](MetadataFieldType.md) | The field type. | [optional] 
**Options** | Pointer to [**[]MetadataFieldOptionDto**](MetadataFieldOptionDto.md) | The choice options of the field. | [optional] 
**Order** | Pointer to **int32** | The field display order inside the template. | [optional] 
**Value** | Pointer to [**MetadataValueDto**](MetadataValueDto.md) | The value of the field on the entry, or `null` when the entry holds no value for it. | [optional] 

## Methods

### NewEntryFieldDto

`func NewEntryFieldDto() *EntryFieldDto`

NewEntryFieldDto instantiates a new EntryFieldDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntryFieldDtoWithDefaults

`func NewEntryFieldDtoWithDefaults() *EntryFieldDto`

NewEntryFieldDtoWithDefaults instantiates a new EntryFieldDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EntryFieldDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntryFieldDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntryFieldDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EntryFieldDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *EntryFieldDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntryFieldDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntryFieldDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EntryFieldDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *EntryFieldDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *EntryFieldDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *EntryFieldDto) GetType() MetadataFieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntryFieldDto) GetTypeOk() (*MetadataFieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntryFieldDto) SetType(v MetadataFieldType)`

SetType sets Type field to given value.

### HasType

`func (o *EntryFieldDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOptions

`func (o *EntryFieldDto) GetOptions() []MetadataFieldOptionDto`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *EntryFieldDto) GetOptionsOk() (*[]MetadataFieldOptionDto, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *EntryFieldDto) SetOptions(v []MetadataFieldOptionDto)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *EntryFieldDto) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *EntryFieldDto) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *EntryFieldDto) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetOrder

`func (o *EntryFieldDto) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *EntryFieldDto) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *EntryFieldDto) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *EntryFieldDto) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetValue

`func (o *EntryFieldDto) GetValue() MetadataValueDto`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EntryFieldDto) GetValueOk() (*MetadataValueDto, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EntryFieldDto) SetValue(v MetadataValueDto)`

SetValue sets Value field to given value.

### HasValue

`func (o *EntryFieldDto) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


