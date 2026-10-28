# PasswordHashSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | Pointer to **int32** | The length of the hash, in bytes. | [optional] 
**Iterations** | Pointer to **int32** | The number of PBKDF2 iterations. | [optional] 
**Salt** | Pointer to **NullableString** | The salt the installation hashes passwords with. | [optional] 

## Methods

### NewPasswordHashSettingsDto

`func NewPasswordHashSettingsDto() *PasswordHashSettingsDto`

NewPasswordHashSettingsDto instantiates a new PasswordHashSettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordHashSettingsDtoWithDefaults

`func NewPasswordHashSettingsDtoWithDefaults() *PasswordHashSettingsDto`

NewPasswordHashSettingsDtoWithDefaults instantiates a new PasswordHashSettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *PasswordHashSettingsDto) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PasswordHashSettingsDto) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PasswordHashSettingsDto) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *PasswordHashSettingsDto) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetIterations

`func (o *PasswordHashSettingsDto) GetIterations() int32`

GetIterations returns the Iterations field if non-nil, zero value otherwise.

### GetIterationsOk

`func (o *PasswordHashSettingsDto) GetIterationsOk() (*int32, bool)`

GetIterationsOk returns a tuple with the Iterations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterations

`func (o *PasswordHashSettingsDto) SetIterations(v int32)`

SetIterations sets Iterations field to given value.

### HasIterations

`func (o *PasswordHashSettingsDto) HasIterations() bool`

HasIterations returns a boolean if a field has been set.

### GetSalt

`func (o *PasswordHashSettingsDto) GetSalt() string`

GetSalt returns the Salt field if non-nil, zero value otherwise.

### GetSaltOk

`func (o *PasswordHashSettingsDto) GetSaltOk() (*string, bool)`

GetSaltOk returns a tuple with the Salt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalt

`func (o *PasswordHashSettingsDto) SetSalt(v string)`

SetSalt sets Salt field to given value.

### HasSalt

`func (o *PasswordHashSettingsDto) HasSalt() bool`

HasSalt returns a boolean if a field has been set.

### SetSaltNil

`func (o *PasswordHashSettingsDto) SetSaltNil(b bool)`

 SetSaltNil sets the value for Salt to be an explicit nil

### UnsetSalt
`func (o *PasswordHashSettingsDto) UnsetSalt()`

UnsetSalt ensures that no value is present for Salt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


