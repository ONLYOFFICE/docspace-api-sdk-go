# SsoEncryptAlgorithmTypeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aes128** | Pointer to **NullableString** | The AES-128-CBC encryption algorithm, which the built-in configuration uses. | [optional] [readonly] 
**Aes256** | Pointer to **NullableString** | The AES-256-CBC encryption algorithm, the strongest of the three. | [optional] [readonly] 
**TriDec** | Pointer to **NullableString** | The Triple DES CBC encryption algorithm, kept for identity providers that support nothing newer. | [optional] [readonly] 

## Methods

### NewSsoEncryptAlgorithmTypeDto

`func NewSsoEncryptAlgorithmTypeDto() *SsoEncryptAlgorithmTypeDto`

NewSsoEncryptAlgorithmTypeDto instantiates a new SsoEncryptAlgorithmTypeDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoEncryptAlgorithmTypeDtoWithDefaults

`func NewSsoEncryptAlgorithmTypeDtoWithDefaults() *SsoEncryptAlgorithmTypeDto`

NewSsoEncryptAlgorithmTypeDtoWithDefaults instantiates a new SsoEncryptAlgorithmTypeDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAes128

`func (o *SsoEncryptAlgorithmTypeDto) GetAes128() string`

GetAes128 returns the Aes128 field if non-nil, zero value otherwise.

### GetAes128Ok

`func (o *SsoEncryptAlgorithmTypeDto) GetAes128Ok() (*string, bool)`

GetAes128Ok returns a tuple with the Aes128 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAes128

`func (o *SsoEncryptAlgorithmTypeDto) SetAes128(v string)`

SetAes128 sets Aes128 field to given value.

### HasAes128

`func (o *SsoEncryptAlgorithmTypeDto) HasAes128() bool`

HasAes128 returns a boolean if a field has been set.

### SetAes128Nil

`func (o *SsoEncryptAlgorithmTypeDto) SetAes128Nil(b bool)`

 SetAes128Nil sets the value for Aes128 to be an explicit nil

### UnsetAes128
`func (o *SsoEncryptAlgorithmTypeDto) UnsetAes128()`

UnsetAes128 ensures that no value is present for Aes128, not even an explicit nil
### GetAes256

`func (o *SsoEncryptAlgorithmTypeDto) GetAes256() string`

GetAes256 returns the Aes256 field if non-nil, zero value otherwise.

### GetAes256Ok

`func (o *SsoEncryptAlgorithmTypeDto) GetAes256Ok() (*string, bool)`

GetAes256Ok returns a tuple with the Aes256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAes256

`func (o *SsoEncryptAlgorithmTypeDto) SetAes256(v string)`

SetAes256 sets Aes256 field to given value.

### HasAes256

`func (o *SsoEncryptAlgorithmTypeDto) HasAes256() bool`

HasAes256 returns a boolean if a field has been set.

### SetAes256Nil

`func (o *SsoEncryptAlgorithmTypeDto) SetAes256Nil(b bool)`

 SetAes256Nil sets the value for Aes256 to be an explicit nil

### UnsetAes256
`func (o *SsoEncryptAlgorithmTypeDto) UnsetAes256()`

UnsetAes256 ensures that no value is present for Aes256, not even an explicit nil
### GetTriDec

`func (o *SsoEncryptAlgorithmTypeDto) GetTriDec() string`

GetTriDec returns the TriDec field if non-nil, zero value otherwise.

### GetTriDecOk

`func (o *SsoEncryptAlgorithmTypeDto) GetTriDecOk() (*string, bool)`

GetTriDecOk returns a tuple with the TriDec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriDec

`func (o *SsoEncryptAlgorithmTypeDto) SetTriDec(v string)`

SetTriDec sets TriDec field to given value.

### HasTriDec

`func (o *SsoEncryptAlgorithmTypeDto) HasTriDec() bool`

HasTriDec returns a boolean if a field has been set.

### SetTriDecNil

`func (o *SsoEncryptAlgorithmTypeDto) SetTriDecNil(b bool)`

 SetTriDecNil sets the value for TriDec to be an explicit nil

### UnsetTriDec
`func (o *SsoEncryptAlgorithmTypeDto) UnsetTriDec()`

UnsetTriDec ensures that no value is present for TriDec, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


