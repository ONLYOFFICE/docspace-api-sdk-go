# MetadataValueDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StringValue** | Pointer to **NullableString** | The string value. | [optional] 
**NumberValue** | Pointer to **NullableInt64** | The number value. | [optional] 
**DateValue** | Pointer to [**ApiDateTime**](ApiDateTime.md) | The date value. | [optional] 
**OptionIds** | Pointer to **[]string** | The selected choice option IDs. | [optional] 

## Methods

### NewMetadataValueDto

`func NewMetadataValueDto() *MetadataValueDto`

NewMetadataValueDto instantiates a new MetadataValueDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataValueDtoWithDefaults

`func NewMetadataValueDtoWithDefaults() *MetadataValueDto`

NewMetadataValueDtoWithDefaults instantiates a new MetadataValueDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStringValue

`func (o *MetadataValueDto) GetStringValue() string`

GetStringValue returns the StringValue field if non-nil, zero value otherwise.

### GetStringValueOk

`func (o *MetadataValueDto) GetStringValueOk() (*string, bool)`

GetStringValueOk returns a tuple with the StringValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringValue

`func (o *MetadataValueDto) SetStringValue(v string)`

SetStringValue sets StringValue field to given value.

### HasStringValue

`func (o *MetadataValueDto) HasStringValue() bool`

HasStringValue returns a boolean if a field has been set.

### SetStringValueNil

`func (o *MetadataValueDto) SetStringValueNil(b bool)`

 SetStringValueNil sets the value for StringValue to be an explicit nil

### UnsetStringValue
`func (o *MetadataValueDto) UnsetStringValue()`

UnsetStringValue ensures that no value is present for StringValue, not even an explicit nil
### GetNumberValue

`func (o *MetadataValueDto) GetNumberValue() int64`

GetNumberValue returns the NumberValue field if non-nil, zero value otherwise.

### GetNumberValueOk

`func (o *MetadataValueDto) GetNumberValueOk() (*int64, bool)`

GetNumberValueOk returns a tuple with the NumberValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberValue

`func (o *MetadataValueDto) SetNumberValue(v int64)`

SetNumberValue sets NumberValue field to given value.

### HasNumberValue

`func (o *MetadataValueDto) HasNumberValue() bool`

HasNumberValue returns a boolean if a field has been set.

### SetNumberValueNil

`func (o *MetadataValueDto) SetNumberValueNil(b bool)`

 SetNumberValueNil sets the value for NumberValue to be an explicit nil

### UnsetNumberValue
`func (o *MetadataValueDto) UnsetNumberValue()`

UnsetNumberValue ensures that no value is present for NumberValue, not even an explicit nil
### GetDateValue

`func (o *MetadataValueDto) GetDateValue() ApiDateTime`

GetDateValue returns the DateValue field if non-nil, zero value otherwise.

### GetDateValueOk

`func (o *MetadataValueDto) GetDateValueOk() (*ApiDateTime, bool)`

GetDateValueOk returns a tuple with the DateValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateValue

`func (o *MetadataValueDto) SetDateValue(v ApiDateTime)`

SetDateValue sets DateValue field to given value.

### HasDateValue

`func (o *MetadataValueDto) HasDateValue() bool`

HasDateValue returns a boolean if a field has been set.

### GetOptionIds

`func (o *MetadataValueDto) GetOptionIds() []string`

GetOptionIds returns the OptionIds field if non-nil, zero value otherwise.

### GetOptionIdsOk

`func (o *MetadataValueDto) GetOptionIdsOk() (*[]string, bool)`

GetOptionIdsOk returns a tuple with the OptionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionIds

`func (o *MetadataValueDto) SetOptionIds(v []string)`

SetOptionIds sets OptionIds field to given value.

### HasOptionIds

`func (o *MetadataValueDto) HasOptionIds() bool`

HasOptionIds returns a boolean if a field has been set.

### SetOptionIdsNil

`func (o *MetadataValueDto) SetOptionIdsNil(b bool)`

 SetOptionIdsNil sets the value for OptionIds to be an explicit nil

### UnsetOptionIds
`func (o *MetadataValueDto) UnsetOptionIds()`

UnsetOptionIds ensures that no value is present for OptionIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


