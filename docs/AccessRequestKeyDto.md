# AccessRequestKeyDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** | User ID | [optional] 
**PublicKeyId** | Pointer to **string** | Public key ID | [optional] 
**PrivateKeyEnc** | Pointer to **NullableString** | Encrypted private key | [optional] 

## Methods

### NewAccessRequestKeyDto

`func NewAccessRequestKeyDto() *AccessRequestKeyDto`

NewAccessRequestKeyDto instantiates a new AccessRequestKeyDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessRequestKeyDtoWithDefaults

`func NewAccessRequestKeyDtoWithDefaults() *AccessRequestKeyDto`

NewAccessRequestKeyDtoWithDefaults instantiates a new AccessRequestKeyDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *AccessRequestKeyDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AccessRequestKeyDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AccessRequestKeyDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AccessRequestKeyDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetPublicKeyId

`func (o *AccessRequestKeyDto) GetPublicKeyId() string`

GetPublicKeyId returns the PublicKeyId field if non-nil, zero value otherwise.

### GetPublicKeyIdOk

`func (o *AccessRequestKeyDto) GetPublicKeyIdOk() (*string, bool)`

GetPublicKeyIdOk returns a tuple with the PublicKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyId

`func (o *AccessRequestKeyDto) SetPublicKeyId(v string)`

SetPublicKeyId sets PublicKeyId field to given value.

### HasPublicKeyId

`func (o *AccessRequestKeyDto) HasPublicKeyId() bool`

HasPublicKeyId returns a boolean if a field has been set.

### GetPrivateKeyEnc

`func (o *AccessRequestKeyDto) GetPrivateKeyEnc() string`

GetPrivateKeyEnc returns the PrivateKeyEnc field if non-nil, zero value otherwise.

### GetPrivateKeyEncOk

`func (o *AccessRequestKeyDto) GetPrivateKeyEncOk() (*string, bool)`

GetPrivateKeyEncOk returns a tuple with the PrivateKeyEnc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyEnc

`func (o *AccessRequestKeyDto) SetPrivateKeyEnc(v string)`

SetPrivateKeyEnc sets PrivateKeyEnc field to given value.

### HasPrivateKeyEnc

`func (o *AccessRequestKeyDto) HasPrivateKeyEnc() bool`

HasPrivateKeyEnc returns a boolean if a field has been set.

### SetPrivateKeyEncNil

`func (o *AccessRequestKeyDto) SetPrivateKeyEncNil(b bool)`

 SetPrivateKeyEncNil sets the value for PrivateKeyEnc to be an explicit nil

### UnsetPrivateKeyEnc
`func (o *AccessRequestKeyDto) UnsetPrivateKeyEnc()`

UnsetPrivateKeyEnc ensures that no value is present for PrivateKeyEnc, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


