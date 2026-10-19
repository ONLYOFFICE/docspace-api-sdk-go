# AuthKeyDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** | The authorization key name. | 
**Value** | **NullableString** | The authorization key value. | 
**Title** | Pointer to **NullableString** | The authorization key title. | [optional] 
**Type** | Pointer to **NullableString** | The field type: text, password, select, toggle. | [optional] 
**Options** | Pointer to **[]string** | The list of options for select type fields. | [optional] 
**DependsOn** | Pointer to **NullableString** | The name of another key this field depends on for visibility. | [optional] 
**DependsOnValue** | Pointer to **NullableString** | The value of the `dependsOn` key that makes this field visible. | [optional] 

## Methods

### NewAuthKeyDto

`func NewAuthKeyDto(name NullableString, value NullableString, ) *AuthKeyDto`

NewAuthKeyDto instantiates a new AuthKeyDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthKeyDtoWithDefaults

`func NewAuthKeyDtoWithDefaults() *AuthKeyDto`

NewAuthKeyDtoWithDefaults instantiates a new AuthKeyDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AuthKeyDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthKeyDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthKeyDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *AuthKeyDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AuthKeyDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetValue

`func (o *AuthKeyDto) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AuthKeyDto) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AuthKeyDto) SetValue(v string)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *AuthKeyDto) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *AuthKeyDto) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetTitle

`func (o *AuthKeyDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AuthKeyDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AuthKeyDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AuthKeyDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *AuthKeyDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *AuthKeyDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetType

`func (o *AuthKeyDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthKeyDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthKeyDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AuthKeyDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *AuthKeyDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *AuthKeyDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetOptions

`func (o *AuthKeyDto) GetOptions() []string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AuthKeyDto) GetOptionsOk() (*[]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AuthKeyDto) SetOptions(v []string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *AuthKeyDto) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *AuthKeyDto) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *AuthKeyDto) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetDependsOn

`func (o *AuthKeyDto) GetDependsOn() string`

GetDependsOn returns the DependsOn field if non-nil, zero value otherwise.

### GetDependsOnOk

`func (o *AuthKeyDto) GetDependsOnOk() (*string, bool)`

GetDependsOnOk returns a tuple with the DependsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOn

`func (o *AuthKeyDto) SetDependsOn(v string)`

SetDependsOn sets DependsOn field to given value.

### HasDependsOn

`func (o *AuthKeyDto) HasDependsOn() bool`

HasDependsOn returns a boolean if a field has been set.

### SetDependsOnNil

`func (o *AuthKeyDto) SetDependsOnNil(b bool)`

 SetDependsOnNil sets the value for DependsOn to be an explicit nil

### UnsetDependsOn
`func (o *AuthKeyDto) UnsetDependsOn()`

UnsetDependsOn ensures that no value is present for DependsOn, not even an explicit nil
### GetDependsOnValue

`func (o *AuthKeyDto) GetDependsOnValue() string`

GetDependsOnValue returns the DependsOnValue field if non-nil, zero value otherwise.

### GetDependsOnValueOk

`func (o *AuthKeyDto) GetDependsOnValueOk() (*string, bool)`

GetDependsOnValueOk returns a tuple with the DependsOnValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOnValue

`func (o *AuthKeyDto) SetDependsOnValue(v string)`

SetDependsOnValue sets DependsOnValue field to given value.

### HasDependsOnValue

`func (o *AuthKeyDto) HasDependsOnValue() bool`

HasDependsOnValue returns a boolean if a field has been set.

### SetDependsOnValueNil

`func (o *AuthKeyDto) SetDependsOnValueNil(b bool)`

 SetDependsOnValueNil sets the value for DependsOnValue to be an explicit nil

### UnsetDependsOnValue
`func (o *AuthKeyDto) UnsetDependsOnValue()`

UnsetDependsOnValue ensures that no value is present for DependsOnValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


