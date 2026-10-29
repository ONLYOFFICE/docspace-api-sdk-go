# AiMetadataValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StringValue** | Pointer to **string** |  | [optional] 
**NumberValue** | Pointer to **float32** |  | [optional] 
**DateValue** | Pointer to **string** | ISO 8601 date-time, UTC midnight of the date. | [optional] 
**OptionIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAiMetadataValue

`func NewAiMetadataValue() *AiMetadataValue`

NewAiMetadataValue instantiates a new AiMetadataValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiMetadataValueWithDefaults

`func NewAiMetadataValueWithDefaults() *AiMetadataValue`

NewAiMetadataValueWithDefaults instantiates a new AiMetadataValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStringValue

`func (o *AiMetadataValue) GetStringValue() string`

GetStringValue returns the StringValue field if non-nil, zero value otherwise.

### GetStringValueOk

`func (o *AiMetadataValue) GetStringValueOk() (*string, bool)`

GetStringValueOk returns a tuple with the StringValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringValue

`func (o *AiMetadataValue) SetStringValue(v string)`

SetStringValue sets StringValue field to given value.

### HasStringValue

`func (o *AiMetadataValue) HasStringValue() bool`

HasStringValue returns a boolean if a field has been set.

### GetNumberValue

`func (o *AiMetadataValue) GetNumberValue() float32`

GetNumberValue returns the NumberValue field if non-nil, zero value otherwise.

### GetNumberValueOk

`func (o *AiMetadataValue) GetNumberValueOk() (*float32, bool)`

GetNumberValueOk returns a tuple with the NumberValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberValue

`func (o *AiMetadataValue) SetNumberValue(v float32)`

SetNumberValue sets NumberValue field to given value.

### HasNumberValue

`func (o *AiMetadataValue) HasNumberValue() bool`

HasNumberValue returns a boolean if a field has been set.

### GetDateValue

`func (o *AiMetadataValue) GetDateValue() string`

GetDateValue returns the DateValue field if non-nil, zero value otherwise.

### GetDateValueOk

`func (o *AiMetadataValue) GetDateValueOk() (*string, bool)`

GetDateValueOk returns a tuple with the DateValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateValue

`func (o *AiMetadataValue) SetDateValue(v string)`

SetDateValue sets DateValue field to given value.

### HasDateValue

`func (o *AiMetadataValue) HasDateValue() bool`

HasDateValue returns a boolean if a field has been set.

### GetOptionIds

`func (o *AiMetadataValue) GetOptionIds() []string`

GetOptionIds returns the OptionIds field if non-nil, zero value otherwise.

### GetOptionIdsOk

`func (o *AiMetadataValue) GetOptionIdsOk() (*[]string, bool)`

GetOptionIdsOk returns a tuple with the OptionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionIds

`func (o *AiMetadataValue) SetOptionIds(v []string)`

SetOptionIds sets OptionIds field to given value.

### HasOptionIds

`func (o *AiMetadataValue) HasOptionIds() bool`

HasOptionIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


