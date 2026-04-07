# EncryptionKeysConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CryptoEngineId** | Pointer to **NullableString** | The crypto engine ID of the encryption key. | [optional] [readonly] 
**PrivateKeyEnc** | Pointer to **NullableString** | The private key. | [optional] 
**PublicKey** | Pointer to **NullableString** | The public key. | [optional] 

## Methods

### NewEncryptionKeysConfig

`func NewEncryptionKeysConfig() *EncryptionKeysConfig`

NewEncryptionKeysConfig instantiates a new EncryptionKeysConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptionKeysConfigWithDefaults

`func NewEncryptionKeysConfigWithDefaults() *EncryptionKeysConfig`

NewEncryptionKeysConfigWithDefaults instantiates a new EncryptionKeysConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCryptoEngineId

`func (o *EncryptionKeysConfig) GetCryptoEngineId() string`

GetCryptoEngineId returns the CryptoEngineId field if non-nil, zero value otherwise.

### GetCryptoEngineIdOk

`func (o *EncryptionKeysConfig) GetCryptoEngineIdOk() (*string, bool)`

GetCryptoEngineIdOk returns a tuple with the CryptoEngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoEngineId

`func (o *EncryptionKeysConfig) SetCryptoEngineId(v string)`

SetCryptoEngineId sets CryptoEngineId field to given value.

### HasCryptoEngineId

`func (o *EncryptionKeysConfig) HasCryptoEngineId() bool`

HasCryptoEngineId returns a boolean if a field has been set.

### SetCryptoEngineIdNil

`func (o *EncryptionKeysConfig) SetCryptoEngineIdNil(b bool)`

 SetCryptoEngineIdNil sets the value for CryptoEngineId to be an explicit nil

### UnsetCryptoEngineId
`func (o *EncryptionKeysConfig) UnsetCryptoEngineId()`

UnsetCryptoEngineId ensures that no value is present for CryptoEngineId, not even an explicit nil
### GetPrivateKeyEnc

`func (o *EncryptionKeysConfig) GetPrivateKeyEnc() string`

GetPrivateKeyEnc returns the PrivateKeyEnc field if non-nil, zero value otherwise.

### GetPrivateKeyEncOk

`func (o *EncryptionKeysConfig) GetPrivateKeyEncOk() (*string, bool)`

GetPrivateKeyEncOk returns a tuple with the PrivateKeyEnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyEnc

`func (o *EncryptionKeysConfig) SetPrivateKeyEnc(v string)`

SetPrivateKeyEnc sets PrivateKeyEnc field to given value.

### HasPrivateKeyEnc

`func (o *EncryptionKeysConfig) HasPrivateKeyEnc() bool`

HasPrivateKeyEnc returns a boolean if a field has been set.

### SetPrivateKeyEncNil

`func (o *EncryptionKeysConfig) SetPrivateKeyEncNil(b bool)`

 SetPrivateKeyEncNil sets the value for PrivateKeyEnc to be an explicit nil

### UnsetPrivateKeyEnc
`func (o *EncryptionKeysConfig) UnsetPrivateKeyEnc()`

UnsetPrivateKeyEnc ensures that no value is present for PrivateKeyEnc, not even an explicit nil
### GetPublicKey

`func (o *EncryptionKeysConfig) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *EncryptionKeysConfig) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *EncryptionKeysConfig) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *EncryptionKeysConfig) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### SetPublicKeyNil

`func (o *EncryptionKeysConfig) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *EncryptionKeysConfig) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


