# AuthKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** | The authorization key name. | 
**Value** | **NullableString** | The authorization key value. | 
**Title** | Pointer to **NullableString** | The authorization key title. | [optional] 
**Type** | Pointer to **NullableString** | The field type: text, password, select, toggle. | [optional] 
**Options** | Pointer to **[]string** | The list of options for select type fields. | [optional] 
**DependsOn** | Pointer to **NullableString** | The name of another key this field depends on for visibility. | [optional] 
**DependsOnValue** | Pointer to **NullableString** | The value of ASC.Web.Studio.UserControls.Management.AuthKey.DependsOn key that makes this field visible. | [optional] 

## Methods

### NewAuthKey

`func NewAuthKey(name NullableString, value NullableString, ) *AuthKey`

NewAuthKey instantiates a new AuthKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthKeyWithDefaults

`func NewAuthKeyWithDefaults() *AuthKey`

NewAuthKeyWithDefaults instantiates a new AuthKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AuthKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthKey) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *AuthKey) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AuthKey) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetValue

`func (o *AuthKey) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AuthKey) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AuthKey) SetValue(v string)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *AuthKey) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *AuthKey) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetTitle

`func (o *AuthKey) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AuthKey) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AuthKey) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AuthKey) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *AuthKey) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *AuthKey) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetType

`func (o *AuthKey) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthKey) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthKey) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AuthKey) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *AuthKey) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *AuthKey) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetOptions

`func (o *AuthKey) GetOptions() []string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AuthKey) GetOptionsOk() (*[]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AuthKey) SetOptions(v []string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *AuthKey) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *AuthKey) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *AuthKey) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetDependsOn

`func (o *AuthKey) GetDependsOn() string`

GetDependsOn returns the DependsOn field if non-nil, zero value otherwise.

### GetDependsOnOk

`func (o *AuthKey) GetDependsOnOk() (*string, bool)`

GetDependsOnOk returns a tuple with the DependsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOn

`func (o *AuthKey) SetDependsOn(v string)`

SetDependsOn sets DependsOn field to given value.

### HasDependsOn

`func (o *AuthKey) HasDependsOn() bool`

HasDependsOn returns a boolean if a field has been set.

### SetDependsOnNil

`func (o *AuthKey) SetDependsOnNil(b bool)`

 SetDependsOnNil sets the value for DependsOn to be an explicit nil

### UnsetDependsOn
`func (o *AuthKey) UnsetDependsOn()`

UnsetDependsOn ensures that no value is present for DependsOn, not even an explicit nil
### GetDependsOnValue

`func (o *AuthKey) GetDependsOnValue() string`

GetDependsOnValue returns the DependsOnValue field if non-nil, zero value otherwise.

### GetDependsOnValueOk

`func (o *AuthKey) GetDependsOnValueOk() (*string, bool)`

GetDependsOnValueOk returns a tuple with the DependsOnValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOnValue

`func (o *AuthKey) SetDependsOnValue(v string)`

SetDependsOnValue sets DependsOnValue field to given value.

### HasDependsOnValue

`func (o *AuthKey) HasDependsOnValue() bool`

HasDependsOnValue returns a boolean if a field has been set.

### SetDependsOnValueNil

`func (o *AuthKey) SetDependsOnValueNil(b bool)`

 SetDependsOnValueNil sets the value for DependsOnValue to be an explicit nil

### UnsetDependsOnValue
`func (o *AuthKey) UnsetDependsOnValue()`

UnsetDependsOnValue ensures that no value is present for DependsOnValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


