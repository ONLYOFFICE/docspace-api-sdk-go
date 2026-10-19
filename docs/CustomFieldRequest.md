# CustomFieldRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** | The field name. The names are matched without regard to case. | 
**Value** | Pointer to **NullableString** | The field value. Null or empty removes the field from the entry. | [optional] 

## Methods

### NewCustomFieldRequest

`func NewCustomFieldRequest(name NullableString, ) *CustomFieldRequest`

NewCustomFieldRequest instantiates a new CustomFieldRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldRequestWithDefaults

`func NewCustomFieldRequestWithDefaults() *CustomFieldRequest`

NewCustomFieldRequestWithDefaults instantiates a new CustomFieldRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CustomFieldRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomFieldRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomFieldRequest) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *CustomFieldRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CustomFieldRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetValue

`func (o *CustomFieldRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CustomFieldRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CustomFieldRequest) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CustomFieldRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *CustomFieldRequest) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *CustomFieldRequest) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


