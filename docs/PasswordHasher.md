# PasswordHasher

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | Pointer to **int32** | The password hash size. | [optional] [readonly] 
**Iterations** | Pointer to **int32** | The number of iterations to generate the ppassword hash. | [optional] [readonly] 
**Salt** | Pointer to **NullableString** | The salt to generate the ppassword hash. | [optional] [readonly] 

## Methods

### NewPasswordHasher

`func NewPasswordHasher() *PasswordHasher`

NewPasswordHasher instantiates a new PasswordHasher object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordHasherWithDefaults

`func NewPasswordHasherWithDefaults() *PasswordHasher`

NewPasswordHasherWithDefaults instantiates a new PasswordHasher object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *PasswordHasher) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PasswordHasher) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PasswordHasher) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *PasswordHasher) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetIterations

`func (o *PasswordHasher) GetIterations() int32`

GetIterations returns the Iterations field if non-nil, zero value otherwise.

### GetIterationsOk

`func (o *PasswordHasher) GetIterationsOk() (*int32, bool)`

GetIterationsOk returns a tuple with the Iterations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterations

`func (o *PasswordHasher) SetIterations(v int32)`

SetIterations sets Iterations field to given value.

### HasIterations

`func (o *PasswordHasher) HasIterations() bool`

HasIterations returns a boolean if a field has been set.

### GetSalt

`func (o *PasswordHasher) GetSalt() string`

GetSalt returns the Salt field if non-nil, zero value otherwise.

### GetSaltOk

`func (o *PasswordHasher) GetSaltOk() (*string, bool)`

GetSaltOk returns a tuple with the Salt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalt

`func (o *PasswordHasher) SetSalt(v string)`

SetSalt sets Salt field to given value.

### HasSalt

`func (o *PasswordHasher) HasSalt() bool`

HasSalt returns a boolean if a field has been set.

### SetSaltNil

`func (o *PasswordHasher) SetSaltNil(b bool)`

 SetSaltNil sets the value for Salt to be an explicit nil

### UnsetSalt
`func (o *PasswordHasher) UnsetSalt()`

UnsetSalt ensures that no value is present for Salt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


