# TenantDomainValidator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Regex** | Pointer to **NullableString** | The regex string to validate a domain. | [optional] [readonly] 
**MinLength** | Pointer to **int32** | The minimum length of the valid domain. | [optional] [readonly] 
**MaxLength** | Pointer to **int32** | The maximum length of the valid domain. | [optional] [readonly] 

## Methods

### NewTenantDomainValidator

`func NewTenantDomainValidator() *TenantDomainValidator`

NewTenantDomainValidator instantiates a new TenantDomainValidator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantDomainValidatorWithDefaults

`func NewTenantDomainValidatorWithDefaults() *TenantDomainValidator`

NewTenantDomainValidatorWithDefaults instantiates a new TenantDomainValidator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegex

`func (o *TenantDomainValidator) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *TenantDomainValidator) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *TenantDomainValidator) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *TenantDomainValidator) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### SetRegexNil

`func (o *TenantDomainValidator) SetRegexNil(b bool)`

 SetRegexNil sets the value for Regex to be an explicit nil

### UnsetRegex
`func (o *TenantDomainValidator) UnsetRegex()`

UnsetRegex ensures that no value is present for Regex, not even an explicit nil
### GetMinLength

`func (o *TenantDomainValidator) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *TenantDomainValidator) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *TenantDomainValidator) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *TenantDomainValidator) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### GetMaxLength

`func (o *TenantDomainValidator) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *TenantDomainValidator) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *TenantDomainValidator) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *TenantDomainValidator) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


