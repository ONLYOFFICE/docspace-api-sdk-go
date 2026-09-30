# ItemKeyValuePairBooleanString

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **bool** | The left half of the pair. Where the pair configures something, this is the identifier the value belongs to -  a setting name, a module id, a logo slot; where the pair reports the result of a call, this is the result  itself, such as the flag telling whether the call succeeded. Which of the two it is, and which keys are  accepted, is stated by the operation that sends or returns the pair. | [optional] 
**Value** | Pointer to **NullableString** | The right half of the pair: what is assigned to the key next to it, or what is reported for it. Its meaning  and its accepted values follow from the key, so read them from the operation that sends or returns the pair. | [optional] 

## Methods

### NewItemKeyValuePairBooleanString

`func NewItemKeyValuePairBooleanString() *ItemKeyValuePairBooleanString`

NewItemKeyValuePairBooleanString instantiates a new ItemKeyValuePairBooleanString object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemKeyValuePairBooleanStringWithDefaults

`func NewItemKeyValuePairBooleanStringWithDefaults() *ItemKeyValuePairBooleanString`

NewItemKeyValuePairBooleanStringWithDefaults instantiates a new ItemKeyValuePairBooleanString object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ItemKeyValuePairBooleanString) GetKey() bool`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ItemKeyValuePairBooleanString) GetKeyOk() (*bool, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ItemKeyValuePairBooleanString) SetKey(v bool)`

SetKey sets Key field to given value.

### HasKey

`func (o *ItemKeyValuePairBooleanString) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *ItemKeyValuePairBooleanString) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ItemKeyValuePairBooleanString) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ItemKeyValuePairBooleanString) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ItemKeyValuePairBooleanString) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ItemKeyValuePairBooleanString) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ItemKeyValuePairBooleanString) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


