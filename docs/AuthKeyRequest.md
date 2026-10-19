# AuthKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** | The key name, as `GET api/2.0/settings/authservice` lists it in `props`. | 
**Value** | **NullableString** | The value to store. An empty string clears the key. | 
**Title** | Pointer to **NullableString** | Accepted for compatibility with earlier clients and not read: the server keeps its own value. | [optional] 
**Type** | Pointer to **NullableString** | Accepted for compatibility with earlier clients and not read: the server keeps its own value. | [optional] 
**Options** | Pointer to **[]string** | Accepted for compatibility with earlier clients and not read: the server keeps its own value. | [optional] 
**DependsOn** | Pointer to **NullableString** | Accepted for compatibility with earlier clients and not read: the server keeps its own value. | [optional] 
**DependsOnValue** | Pointer to **NullableString** | Accepted for compatibility with earlier clients and not read: the server keeps its own value. | [optional] 

## Methods

### NewAuthKeyRequest

`func NewAuthKeyRequest(name NullableString, value NullableString, ) *AuthKeyRequest`

NewAuthKeyRequest instantiates a new AuthKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthKeyRequestWithDefaults

`func NewAuthKeyRequestWithDefaults() *AuthKeyRequest`

NewAuthKeyRequestWithDefaults instantiates a new AuthKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AuthKeyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthKeyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthKeyRequest) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *AuthKeyRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AuthKeyRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetValue

`func (o *AuthKeyRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AuthKeyRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AuthKeyRequest) SetValue(v string)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *AuthKeyRequest) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *AuthKeyRequest) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetTitle

`func (o *AuthKeyRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AuthKeyRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AuthKeyRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AuthKeyRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *AuthKeyRequest) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *AuthKeyRequest) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetType

`func (o *AuthKeyRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthKeyRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthKeyRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AuthKeyRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *AuthKeyRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *AuthKeyRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetOptions

`func (o *AuthKeyRequest) GetOptions() []string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AuthKeyRequest) GetOptionsOk() (*[]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AuthKeyRequest) SetOptions(v []string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *AuthKeyRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *AuthKeyRequest) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *AuthKeyRequest) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetDependsOn

`func (o *AuthKeyRequest) GetDependsOn() string`

GetDependsOn returns the DependsOn field if non-nil, zero value otherwise.

### GetDependsOnOk

`func (o *AuthKeyRequest) GetDependsOnOk() (*string, bool)`

GetDependsOnOk returns a tuple with the DependsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOn

`func (o *AuthKeyRequest) SetDependsOn(v string)`

SetDependsOn sets DependsOn field to given value.

### HasDependsOn

`func (o *AuthKeyRequest) HasDependsOn() bool`

HasDependsOn returns a boolean if a field has been set.

### SetDependsOnNil

`func (o *AuthKeyRequest) SetDependsOnNil(b bool)`

 SetDependsOnNil sets the value for DependsOn to be an explicit nil

### UnsetDependsOn
`func (o *AuthKeyRequest) UnsetDependsOn()`

UnsetDependsOn ensures that no value is present for DependsOn, not even an explicit nil
### GetDependsOnValue

`func (o *AuthKeyRequest) GetDependsOnValue() string`

GetDependsOnValue returns the DependsOnValue field if non-nil, zero value otherwise.

### GetDependsOnValueOk

`func (o *AuthKeyRequest) GetDependsOnValueOk() (*string, bool)`

GetDependsOnValueOk returns a tuple with the DependsOnValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOnValue

`func (o *AuthKeyRequest) SetDependsOnValue(v string)`

SetDependsOnValue sets DependsOnValue field to given value.

### HasDependsOnValue

`func (o *AuthKeyRequest) HasDependsOnValue() bool`

HasDependsOnValue returns a boolean if a field has been set.

### SetDependsOnValueNil

`func (o *AuthKeyRequest) SetDependsOnValueNil(b bool)`

 SetDependsOnValueNil sets the value for DependsOnValue to be an explicit nil

### UnsetDependsOnValue
`func (o *AuthKeyRequest) UnsetDependsOnValue()`

UnsetDependsOnValue ensures that no value is present for DependsOnValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


