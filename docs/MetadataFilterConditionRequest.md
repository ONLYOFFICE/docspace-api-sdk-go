# MetadataFilterConditionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldId** | Pointer to **int32** | The ID of the template field the condition is on. A custom field is addressed by ASC.Files.Core.MetadataFilterConditionRequest.Name instead. | [optional] 
**Name** | Pointer to **NullableString** | The name of a custom field, for the conditions on the custom fields, which have no identifier outside. Either the  field ID or the name is given; the name is matched without regard to case. | [optional] 
**Op** | Pointer to **NullableString** | The operator, one of ASC.Files.Core.MetadataFilterOperators: eq or contains for a string field, eq, range,  gt, gte, lt or lte for a number field, range, gt, gte, lt or lte for a date field, in for a  choice field. Optional: an omitted operator is read as eq for a string, as range for a date, as in for a  choice, and for a number as eq when a value is given or as range when a bound is. A present one has to belong  to the field type and to the arguments: eq takes a value, range takes bounds. | [optional] 
**Value** | Pointer to **NullableString** | The value of a single-valued condition: the whole or a part of a string field, the exact value of a number field,  or the bound of a gt, gte, lt or lte condition on a number or a date field. A JSON number is accepted as  well as a string. A date given without a time (2026-06-30) names the whole day (UTC). | [optional] 
**From** | Pointer to **NullableString** | The inclusive lower bound of a range. A date given without a time (2026-06-01) is the start of that day (UTC). | [optional] 
**To** | Pointer to **NullableString** | The inclusive upper bound of a range. A date given without a time (2026-06-30) covers the whole day (UTC);  a value with a time is an instant and is taken as is. | [optional] 
**OptionIds** | Pointer to **[]string** | The options any of which the choice field must hold. | [optional] 

## Methods

### NewMetadataFilterConditionRequest

`func NewMetadataFilterConditionRequest() *MetadataFilterConditionRequest`

NewMetadataFilterConditionRequest instantiates a new MetadataFilterConditionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataFilterConditionRequestWithDefaults

`func NewMetadataFilterConditionRequestWithDefaults() *MetadataFilterConditionRequest`

NewMetadataFilterConditionRequestWithDefaults instantiates a new MetadataFilterConditionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldId

`func (o *MetadataFilterConditionRequest) GetFieldId() int32`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *MetadataFilterConditionRequest) GetFieldIdOk() (*int32, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *MetadataFilterConditionRequest) SetFieldId(v int32)`

SetFieldId sets FieldId field to given value.

### HasFieldId

`func (o *MetadataFilterConditionRequest) HasFieldId() bool`

HasFieldId returns a boolean if a field has been set.

### GetName

`func (o *MetadataFilterConditionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetadataFilterConditionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetadataFilterConditionRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetadataFilterConditionRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MetadataFilterConditionRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MetadataFilterConditionRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOp

`func (o *MetadataFilterConditionRequest) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *MetadataFilterConditionRequest) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *MetadataFilterConditionRequest) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *MetadataFilterConditionRequest) HasOp() bool`

HasOp returns a boolean if a field has been set.

### SetOpNil

`func (o *MetadataFilterConditionRequest) SetOpNil(b bool)`

 SetOpNil sets the value for Op to be an explicit nil

### UnsetOp
`func (o *MetadataFilterConditionRequest) UnsetOp()`

UnsetOp ensures that no value is present for Op, not even an explicit nil
### GetValue

`func (o *MetadataFilterConditionRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MetadataFilterConditionRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MetadataFilterConditionRequest) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *MetadataFilterConditionRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *MetadataFilterConditionRequest) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *MetadataFilterConditionRequest) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetFrom

`func (o *MetadataFilterConditionRequest) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *MetadataFilterConditionRequest) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *MetadataFilterConditionRequest) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *MetadataFilterConditionRequest) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### SetFromNil

`func (o *MetadataFilterConditionRequest) SetFromNil(b bool)`

 SetFromNil sets the value for From to be an explicit nil

### UnsetFrom
`func (o *MetadataFilterConditionRequest) UnsetFrom()`

UnsetFrom ensures that no value is present for From, not even an explicit nil
### GetTo

`func (o *MetadataFilterConditionRequest) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *MetadataFilterConditionRequest) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *MetadataFilterConditionRequest) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *MetadataFilterConditionRequest) HasTo() bool`

HasTo returns a boolean if a field has been set.

### SetToNil

`func (o *MetadataFilterConditionRequest) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *MetadataFilterConditionRequest) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil
### GetOptionIds

`func (o *MetadataFilterConditionRequest) GetOptionIds() []string`

GetOptionIds returns the OptionIds field if non-nil, zero value otherwise.

### GetOptionIdsOk

`func (o *MetadataFilterConditionRequest) GetOptionIdsOk() (*[]string, bool)`

GetOptionIdsOk returns a tuple with the OptionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionIds

`func (o *MetadataFilterConditionRequest) SetOptionIds(v []string)`

SetOptionIds sets OptionIds field to given value.

### HasOptionIds

`func (o *MetadataFilterConditionRequest) HasOptionIds() bool`

HasOptionIds returns a boolean if a field has been set.

### SetOptionIdsNil

`func (o *MetadataFilterConditionRequest) SetOptionIdsNil(b bool)`

 SetOptionIdsNil sets the value for OptionIds to be an explicit nil

### UnsetOptionIds
`func (o *MetadataFilterConditionRequest) UnsetOptionIds()`

UnsetOptionIds ensures that no value is present for OptionIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


