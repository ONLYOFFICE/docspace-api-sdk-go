# EncryptionKeyDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The identifier of the key pair. | [optional] 
**UserId** | Pointer to **string** | The identifier of the user the key pair belongs to. | [optional] 
**Date** | Pointer to **time.Time** | The date and time when the key pair was created. | [optional] 
**PublicKey** | Pointer to **NullableString** | The public key of the pair, used to encrypt the file keys. | [optional] 
**PrivateKeyEnc** | Pointer to **NullableString** | The private key of the pair, encrypted with the user password. | [optional] 
**CryptoEngineId** | Pointer to **NullableString** | The identifier of the crypto engine the key pair was issued for. | [optional] 

## Methods

### NewEncryptionKeyDto

`func NewEncryptionKeyDto() *EncryptionKeyDto`

NewEncryptionKeyDto instantiates a new EncryptionKeyDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptionKeyDtoWithDefaults

`func NewEncryptionKeyDtoWithDefaults() *EncryptionKeyDto`

NewEncryptionKeyDtoWithDefaults instantiates a new EncryptionKeyDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EncryptionKeyDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EncryptionKeyDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EncryptionKeyDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EncryptionKeyDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *EncryptionKeyDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *EncryptionKeyDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *EncryptionKeyDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *EncryptionKeyDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetDate

`func (o *EncryptionKeyDto) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *EncryptionKeyDto) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *EncryptionKeyDto) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *EncryptionKeyDto) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetPublicKey

`func (o *EncryptionKeyDto) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *EncryptionKeyDto) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *EncryptionKeyDto) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *EncryptionKeyDto) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### SetPublicKeyNil

`func (o *EncryptionKeyDto) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *EncryptionKeyDto) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
### GetPrivateKeyEnc

`func (o *EncryptionKeyDto) GetPrivateKeyEnc() string`

GetPrivateKeyEnc returns the PrivateKeyEnc field if non-nil, zero value otherwise.

### GetPrivateKeyEncOk

`func (o *EncryptionKeyDto) GetPrivateKeyEncOk() (*string, bool)`

GetPrivateKeyEncOk returns a tuple with the PrivateKeyEnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyEnc

`func (o *EncryptionKeyDto) SetPrivateKeyEnc(v string)`

SetPrivateKeyEnc sets PrivateKeyEnc field to given value.

### HasPrivateKeyEnc

`func (o *EncryptionKeyDto) HasPrivateKeyEnc() bool`

HasPrivateKeyEnc returns a boolean if a field has been set.

### SetPrivateKeyEncNil

`func (o *EncryptionKeyDto) SetPrivateKeyEncNil(b bool)`

 SetPrivateKeyEncNil sets the value for PrivateKeyEnc to be an explicit nil

### UnsetPrivateKeyEnc
`func (o *EncryptionKeyDto) UnsetPrivateKeyEnc()`

UnsetPrivateKeyEnc ensures that no value is present for PrivateKeyEnc, not even an explicit nil
### GetCryptoEngineId

`func (o *EncryptionKeyDto) GetCryptoEngineId() string`

GetCryptoEngineId returns the CryptoEngineId field if non-nil, zero value otherwise.

### GetCryptoEngineIdOk

`func (o *EncryptionKeyDto) GetCryptoEngineIdOk() (*string, bool)`

GetCryptoEngineIdOk returns a tuple with the CryptoEngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCryptoEngineId

`func (o *EncryptionKeyDto) SetCryptoEngineId(v string)`

SetCryptoEngineId sets CryptoEngineId field to given value.

### HasCryptoEngineId

`func (o *EncryptionKeyDto) HasCryptoEngineId() bool`

HasCryptoEngineId returns a boolean if a field has been set.

### SetCryptoEngineIdNil

`func (o *EncryptionKeyDto) SetCryptoEngineIdNil(b bool)`

 SetCryptoEngineIdNil sets the value for CryptoEngineId to be an explicit nil

### UnsetCryptoEngineId
`func (o *EncryptionKeyDto) UnsetCryptoEngineId()`

UnsetCryptoEngineId ensures that no value is present for CryptoEngineId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


