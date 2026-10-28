# MetadataValueRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldId** | **int32** | The field ID. | 
**StringValue** | Pointer to **NullableString** | The string value. | [optional] 
**NumberValue** | Pointer to **NullableInt64** | The number value. | [optional] 
**DateValue** | Pointer to **NullableTime** | The date value. A value without a time zone offset is treated as UTC, the same way the metadata filters treat their date bounds. | [optional] 
**OptionIds** | Pointer to **[]string** | The selected choice option IDs. | [optional] 

## Methods

### NewMetadataValueRequest

`func NewMetadataValueRequest(fieldId int32, ) *MetadataValueRequest`

NewMetadataValueRequest instantiates a new MetadataValueRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataValueRequestWithDefaults

`func NewMetadataValueRequestWithDefaults() *MetadataValueRequest`

NewMetadataValueRequestWithDefaults instantiates a new MetadataValueRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldId

`func (o *MetadataValueRequest) GetFieldId() int32`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *MetadataValueRequest) GetFieldIdOk() (*int32, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *MetadataValueRequest) SetFieldId(v int32)`

SetFieldId sets FieldId field to given value.


### GetStringValue

`func (o *MetadataValueRequest) GetStringValue() string`

GetStringValue returns the StringValue field if non-nil, zero value otherwise.

### GetStringValueOk

`func (o *MetadataValueRequest) GetStringValueOk() (*string, bool)`

GetStringValueOk returns a tuple with the StringValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringValue

`func (o *MetadataValueRequest) SetStringValue(v string)`

SetStringValue sets StringValue field to given value.

### HasStringValue

`func (o *MetadataValueRequest) HasStringValue() bool`

HasStringValue returns a boolean if a field has been set.

### SetStringValueNil

`func (o *MetadataValueRequest) SetStringValueNil(b bool)`

 SetStringValueNil sets the value for StringValue to be an explicit nil

### UnsetStringValue
`func (o *MetadataValueRequest) UnsetStringValue()`

UnsetStringValue ensures that no value is present for StringValue, not even an explicit nil
### GetNumberValue

`func (o *MetadataValueRequest) GetNumberValue() int64`

GetNumberValue returns the NumberValue field if non-nil, zero value otherwise.

### GetNumberValueOk

`func (o *MetadataValueRequest) GetNumberValueOk() (*int64, bool)`

GetNumberValueOk returns a tuple with the NumberValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberValue

`func (o *MetadataValueRequest) SetNumberValue(v int64)`

SetNumberValue sets NumberValue field to given value.

### HasNumberValue

`func (o *MetadataValueRequest) HasNumberValue() bool`

HasNumberValue returns a boolean if a field has been set.

### SetNumberValueNil

`func (o *MetadataValueRequest) SetNumberValueNil(b bool)`

 SetNumberValueNil sets the value for NumberValue to be an explicit nil

### UnsetNumberValue
`func (o *MetadataValueRequest) UnsetNumberValue()`

UnsetNumberValue ensures that no value is present for NumberValue, not even an explicit nil
### GetDateValue

`func (o *MetadataValueRequest) GetDateValue() time.Time`

GetDateValue returns the DateValue field if non-nil, zero value otherwise.

### GetDateValueOk

`func (o *MetadataValueRequest) GetDateValueOk() (*time.Time, bool)`

GetDateValueOk returns a tuple with the DateValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateValue

`func (o *MetadataValueRequest) SetDateValue(v time.Time)`

SetDateValue sets DateValue field to given value.

### HasDateValue

`func (o *MetadataValueRequest) HasDateValue() bool`

HasDateValue returns a boolean if a field has been set.

### SetDateValueNil

`func (o *MetadataValueRequest) SetDateValueNil(b bool)`

 SetDateValueNil sets the value for DateValue to be an explicit nil

### UnsetDateValue
`func (o *MetadataValueRequest) UnsetDateValue()`

UnsetDateValue ensures that no value is present for DateValue, not even an explicit nil
### GetOptionIds

`func (o *MetadataValueRequest) GetOptionIds() []string`

GetOptionIds returns the OptionIds field if non-nil, zero value otherwise.

### GetOptionIdsOk

`func (o *MetadataValueRequest) GetOptionIdsOk() (*[]string, bool)`

GetOptionIdsOk returns a tuple with the OptionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionIds

`func (o *MetadataValueRequest) SetOptionIds(v []string)`

SetOptionIds sets OptionIds field to given value.

### HasOptionIds

`func (o *MetadataValueRequest) HasOptionIds() bool`

HasOptionIds returns a boolean if a field has been set.

### SetOptionIdsNil

`func (o *MetadataValueRequest) SetOptionIdsNil(b bool)`

 SetOptionIdsNil sets the value for OptionIds to be an explicit nil

### UnsetOptionIds
`func (o *MetadataValueRequest) UnsetOptionIds()`

UnsetOptionIds ensures that no value is present for OptionIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


