# EncryptionKeyRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The identifier of the key pair. | [optional] 
**PublicKey** | Pointer to **NullableString** | The public key of the pair, used to encrypt the file keys. | [optional] 
**PrivateKeyEnc** | Pointer to **NullableString** | The private key of the pair, encrypted with the user password. | [optional] 

## Methods

### NewEncryptionKeyRequestDto

`func NewEncryptionKeyRequestDto() *EncryptionKeyRequestDto`

NewEncryptionKeyRequestDto instantiates a new EncryptionKeyRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptionKeyRequestDtoWithDefaults

`func NewEncryptionKeyRequestDtoWithDefaults() *EncryptionKeyRequestDto`

NewEncryptionKeyRequestDtoWithDefaults instantiates a new EncryptionKeyRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EncryptionKeyRequestDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EncryptionKeyRequestDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EncryptionKeyRequestDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EncryptionKeyRequestDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPublicKey

`func (o *EncryptionKeyRequestDto) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *EncryptionKeyRequestDto) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *EncryptionKeyRequestDto) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *EncryptionKeyRequestDto) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### SetPublicKeyNil

`func (o *EncryptionKeyRequestDto) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *EncryptionKeyRequestDto) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
### GetPrivateKeyEnc

`func (o *EncryptionKeyRequestDto) GetPrivateKeyEnc() string`

GetPrivateKeyEnc returns the PrivateKeyEnc field if non-nil, zero value otherwise.

### GetPrivateKeyEncOk

`func (o *EncryptionKeyRequestDto) GetPrivateKeyEncOk() (*string, bool)`

GetPrivateKeyEncOk returns a tuple with the PrivateKeyEnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyEnc

`func (o *EncryptionKeyRequestDto) SetPrivateKeyEnc(v string)`

SetPrivateKeyEnc sets PrivateKeyEnc field to given value.

### HasPrivateKeyEnc

`func (o *EncryptionKeyRequestDto) HasPrivateKeyEnc() bool`

HasPrivateKeyEnc returns a boolean if a field has been set.

### SetPrivateKeyEncNil

`func (o *EncryptionKeyRequestDto) SetPrivateKeyEncNil(b bool)`

 SetPrivateKeyEncNil sets the value for PrivateKeyEnc to be an explicit nil

### UnsetPrivateKeyEnc
`func (o *EncryptionKeyRequestDto) UnsetPrivateKeyEnc()`

UnsetPrivateKeyEnc ensures that no value is present for PrivateKeyEnc, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


