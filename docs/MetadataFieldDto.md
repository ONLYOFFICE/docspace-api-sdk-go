# MetadataFieldDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The field ID. | [optional] 
**TemplateId** | Pointer to **int32** | The ID of the template the field belongs to. | [optional] 
**Name** | Pointer to **NullableString** | The field name. | [optional] 
**Type** | Pointer to [**MetadataFieldType**](MetadataFieldType.md) | The field type. | [optional] 
**Options** | Pointer to [**[]MetadataFieldOptionDto**](MetadataFieldOptionDto.md) | The choice options of the field. | [optional] 
**Order** | Pointer to **int32** | The field display order inside the template. | [optional] 

## Methods

### NewMetadataFieldDto

`func NewMetadataFieldDto() *MetadataFieldDto`

NewMetadataFieldDto instantiates a new MetadataFieldDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataFieldDtoWithDefaults

`func NewMetadataFieldDtoWithDefaults() *MetadataFieldDto`

NewMetadataFieldDtoWithDefaults instantiates a new MetadataFieldDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MetadataFieldDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetadataFieldDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetadataFieldDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MetadataFieldDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTemplateId

`func (o *MetadataFieldDto) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *MetadataFieldDto) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *MetadataFieldDto) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *MetadataFieldDto) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetName

`func (o *MetadataFieldDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetadataFieldDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetadataFieldDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetadataFieldDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MetadataFieldDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MetadataFieldDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *MetadataFieldDto) GetType() MetadataFieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetadataFieldDto) GetTypeOk() (*MetadataFieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetadataFieldDto) SetType(v MetadataFieldType)`

SetType sets Type field to given value.

### HasType

`func (o *MetadataFieldDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOptions

`func (o *MetadataFieldDto) GetOptions() []MetadataFieldOptionDto`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *MetadataFieldDto) GetOptionsOk() (*[]MetadataFieldOptionDto, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *MetadataFieldDto) SetOptions(v []MetadataFieldOptionDto)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *MetadataFieldDto) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *MetadataFieldDto) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *MetadataFieldDto) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetOrder

`func (o *MetadataFieldDto) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *MetadataFieldDto) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *MetadataFieldDto) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *MetadataFieldDto) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


