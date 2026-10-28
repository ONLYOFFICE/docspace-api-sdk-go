# ItemKeyValuePairStringBoolean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **NullableString** | The left half of the pair. Where the pair configures something, this is the identifier the value belongs to -  a setting name, a module id, a logo slot; where the pair reports the result of a call, this is the result  itself, such as the flag telling whether the call succeeded. Which of the two it is, and which keys are  accepted, is stated by the operation that sends or returns the pair. | [optional] 
**Value** | Pointer to **bool** | The right half of the pair: what is assigned to the key next to it, or what is reported for it. Its meaning  and its accepted values follow from the key, so read them from the operation that sends or returns the pair. | [optional] 

## Methods

### NewItemKeyValuePairStringBoolean

`func NewItemKeyValuePairStringBoolean() *ItemKeyValuePairStringBoolean`

NewItemKeyValuePairStringBoolean instantiates a new ItemKeyValuePairStringBoolean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemKeyValuePairStringBooleanWithDefaults

`func NewItemKeyValuePairStringBooleanWithDefaults() *ItemKeyValuePairStringBoolean`

NewItemKeyValuePairStringBooleanWithDefaults instantiates a new ItemKeyValuePairStringBoolean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ItemKeyValuePairStringBoolean) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ItemKeyValuePairStringBoolean) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ItemKeyValuePairStringBoolean) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ItemKeyValuePairStringBoolean) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *ItemKeyValuePairStringBoolean) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *ItemKeyValuePairStringBoolean) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetValue

`func (o *ItemKeyValuePairStringBoolean) GetValue() bool`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ItemKeyValuePairStringBoolean) GetValueOk() (*bool, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ItemKeyValuePairStringBoolean) SetValue(v bool)`

SetValue sets Value field to given value.

### HasValue

`func (o *ItemKeyValuePairStringBoolean) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


