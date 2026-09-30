# SsoSigningAlgorithmTypeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RsaSha1** | Pointer to **NullableString** | The RSA-SHA1 signing algorithm, which the built-in configuration uses. SHA-1 is the weakest of the three  and some identity providers no longer accept it. | [optional] [readonly] 
**RsaSha256** | Pointer to **NullableString** | The RSA-SHA256 signing algorithm. | [optional] [readonly] 
**RsaSha512** | Pointer to **NullableString** | The RSA-SHA512 signing algorithm. | [optional] [readonly] 

## Methods

### NewSsoSigningAlgorithmTypeDto

`func NewSsoSigningAlgorithmTypeDto() *SsoSigningAlgorithmTypeDto`

NewSsoSigningAlgorithmTypeDto instantiates a new SsoSigningAlgorithmTypeDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoSigningAlgorithmTypeDtoWithDefaults

`func NewSsoSigningAlgorithmTypeDtoWithDefaults() *SsoSigningAlgorithmTypeDto`

NewSsoSigningAlgorithmTypeDtoWithDefaults instantiates a new SsoSigningAlgorithmTypeDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRsaSha1

`func (o *SsoSigningAlgorithmTypeDto) GetRsaSha1() string`

GetRsaSha1 returns the RsaSha1 field if non-nil, zero value otherwise.

### GetRsaSha1Ok

`func (o *SsoSigningAlgorithmTypeDto) GetRsaSha1Ok() (*string, bool)`

GetRsaSha1Ok returns a tuple with the RsaSha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsaSha1

`func (o *SsoSigningAlgorithmTypeDto) SetRsaSha1(v string)`

SetRsaSha1 sets RsaSha1 field to given value.

### HasRsaSha1

`func (o *SsoSigningAlgorithmTypeDto) HasRsaSha1() bool`

HasRsaSha1 returns a boolean if a field has been set.

### SetRsaSha1Nil

`func (o *SsoSigningAlgorithmTypeDto) SetRsaSha1Nil(b bool)`

 SetRsaSha1Nil sets the value for RsaSha1 to be an explicit nil

### UnsetRsaSha1
`func (o *SsoSigningAlgorithmTypeDto) UnsetRsaSha1()`

UnsetRsaSha1 ensures that no value is present for RsaSha1, not even an explicit nil
### GetRsaSha256

`func (o *SsoSigningAlgorithmTypeDto) GetRsaSha256() string`

GetRsaSha256 returns the RsaSha256 field if non-nil, zero value otherwise.

### GetRsaSha256Ok

`func (o *SsoSigningAlgorithmTypeDto) GetRsaSha256Ok() (*string, bool)`

GetRsaSha256Ok returns a tuple with the RsaSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsaSha256

`func (o *SsoSigningAlgorithmTypeDto) SetRsaSha256(v string)`

SetRsaSha256 sets RsaSha256 field to given value.

### HasRsaSha256

`func (o *SsoSigningAlgorithmTypeDto) HasRsaSha256() bool`

HasRsaSha256 returns a boolean if a field has been set.

### SetRsaSha256Nil

`func (o *SsoSigningAlgorithmTypeDto) SetRsaSha256Nil(b bool)`

 SetRsaSha256Nil sets the value for RsaSha256 to be an explicit nil

### UnsetRsaSha256
`func (o *SsoSigningAlgorithmTypeDto) UnsetRsaSha256()`

UnsetRsaSha256 ensures that no value is present for RsaSha256, not even an explicit nil
### GetRsaSha512

`func (o *SsoSigningAlgorithmTypeDto) GetRsaSha512() string`

GetRsaSha512 returns the RsaSha512 field if non-nil, zero value otherwise.

### GetRsaSha512Ok

`func (o *SsoSigningAlgorithmTypeDto) GetRsaSha512Ok() (*string, bool)`

GetRsaSha512Ok returns a tuple with the RsaSha512 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsaSha512

`func (o *SsoSigningAlgorithmTypeDto) SetRsaSha512(v string)`

SetRsaSha512 sets RsaSha512 field to given value.

### HasRsaSha512

`func (o *SsoSigningAlgorithmTypeDto) HasRsaSha512() bool`

HasRsaSha512 returns a boolean if a field has been set.

### SetRsaSha512Nil

`func (o *SsoSigningAlgorithmTypeDto) SetRsaSha512Nil(b bool)`

 SetRsaSha512Nil sets the value for RsaSha512 to be an explicit nil

### UnsetRsaSha512
`func (o *SsoSigningAlgorithmTypeDto) UnsetRsaSha512()`

UnsetRsaSha512 ensures that no value is present for RsaSha512, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


