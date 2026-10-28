# AuthenticationTokenDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **NullableString** | The token to put in the `Authorization` header of later calls. It is empty whenever a second factor is  still outstanding, which is what `sms` or `tfa` then says; the same token is also set as a portal cookie by  the call that issued it, so a browser client does not have to carry it itself. | [optional] 
**Expires** | Pointer to **time.Time** | When the token stops being accepted. It stays at its zero value when `session=true` tied the token to the  browser session instead of to a fixed moment. On the two operations that only send an SMS it carries a  different meaning: there is no token, and this is the moment the code that was just sent expires. | [optional] 
**Sms** | Pointer to **bool** | Whether an SMS code is the second factor in play. Next to an empty `token` it means the code has to be sent  to `POST api/2.0/authentication/{code}` before a token is issued; next to a filled `token` it means the  code just accepted was an SMS one. | [optional] 
**PhoneNoise** | Pointer to **NullableString** | The stored phone number with its middle digits masked, filled in only while `sms` is set and a number is  already activated for the user. It is there to be shown to the person signing in, not to be sent back. | [optional] 
**Tfa** | Pointer to **bool** | Whether an authenticator app is the second factor in play, with the same two readings as `sms`. | [optional] 
**TfaKey** | Pointer to **NullableString** | The secret to enrol in an authenticator app, in the manual-entry form. It is filled in only while `tfa` is  set and the app has not been connected yet, which is the one moment the secret is handed out; once the app  is connected it stays empty. `GET api/2.0/settings/tfaapp/setup` returns the same secret with a QR code. | [optional] 
**ConfirmUrl** | Pointer to **NullableString** | The confirmation link the client has to open to get past the second factor. It points at phone activation  while no number is activated, at authenticator-app activation while the app is not connected, and at the  plain code prompt once either is in place. It is empty in an answer that already carries a token. | [optional] 

## Methods

### NewAuthenticationTokenDto

`func NewAuthenticationTokenDto() *AuthenticationTokenDto`

NewAuthenticationTokenDto instantiates a new AuthenticationTokenDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationTokenDtoWithDefaults

`func NewAuthenticationTokenDtoWithDefaults() *AuthenticationTokenDto`

NewAuthenticationTokenDtoWithDefaults instantiates a new AuthenticationTokenDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *AuthenticationTokenDto) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AuthenticationTokenDto) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AuthenticationTokenDto) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AuthenticationTokenDto) HasToken() bool`

HasToken returns a boolean if a field has been set.

### SetTokenNil

`func (o *AuthenticationTokenDto) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *AuthenticationTokenDto) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetExpires

`func (o *AuthenticationTokenDto) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *AuthenticationTokenDto) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *AuthenticationTokenDto) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *AuthenticationTokenDto) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetSms

`func (o *AuthenticationTokenDto) GetSms() bool`

GetSms returns the Sms field if non-nil, zero value otherwise.

### GetSmsOk

`func (o *AuthenticationTokenDto) GetSmsOk() (*bool, bool)`

GetSmsOk returns a tuple with the Sms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSms

`func (o *AuthenticationTokenDto) SetSms(v bool)`

SetSms sets Sms field to given value.

### HasSms

`func (o *AuthenticationTokenDto) HasSms() bool`

HasSms returns a boolean if a field has been set.

### GetPhoneNoise

`func (o *AuthenticationTokenDto) GetPhoneNoise() string`

GetPhoneNoise returns the PhoneNoise field if non-nil, zero value otherwise.

### GetPhoneNoiseOk

`func (o *AuthenticationTokenDto) GetPhoneNoiseOk() (*string, bool)`

GetPhoneNoiseOk returns a tuple with the PhoneNoise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNoise

`func (o *AuthenticationTokenDto) SetPhoneNoise(v string)`

SetPhoneNoise sets PhoneNoise field to given value.

### HasPhoneNoise

`func (o *AuthenticationTokenDto) HasPhoneNoise() bool`

HasPhoneNoise returns a boolean if a field has been set.

### SetPhoneNoiseNil

`func (o *AuthenticationTokenDto) SetPhoneNoiseNil(b bool)`

 SetPhoneNoiseNil sets the value for PhoneNoise to be an explicit nil

### UnsetPhoneNoise
`func (o *AuthenticationTokenDto) UnsetPhoneNoise()`

UnsetPhoneNoise ensures that no value is present for PhoneNoise, not even an explicit nil
### GetTfa

`func (o *AuthenticationTokenDto) GetTfa() bool`

GetTfa returns the Tfa field if non-nil, zero value otherwise.

### GetTfaOk

`func (o *AuthenticationTokenDto) GetTfaOk() (*bool, bool)`

GetTfaOk returns a tuple with the Tfa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTfa

`func (o *AuthenticationTokenDto) SetTfa(v bool)`

SetTfa sets Tfa field to given value.

### HasTfa

`func (o *AuthenticationTokenDto) HasTfa() bool`

HasTfa returns a boolean if a field has been set.

### GetTfaKey

`func (o *AuthenticationTokenDto) GetTfaKey() string`

GetTfaKey returns the TfaKey field if non-nil, zero value otherwise.

### GetTfaKeyOk

`func (o *AuthenticationTokenDto) GetTfaKeyOk() (*string, bool)`

GetTfaKeyOk returns a tuple with the TfaKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTfaKey

`func (o *AuthenticationTokenDto) SetTfaKey(v string)`

SetTfaKey sets TfaKey field to given value.

### HasTfaKey

`func (o *AuthenticationTokenDto) HasTfaKey() bool`

HasTfaKey returns a boolean if a field has been set.

### SetTfaKeyNil

`func (o *AuthenticationTokenDto) SetTfaKeyNil(b bool)`

 SetTfaKeyNil sets the value for TfaKey to be an explicit nil

### UnsetTfaKey
`func (o *AuthenticationTokenDto) UnsetTfaKey()`

UnsetTfaKey ensures that no value is present for TfaKey, not even an explicit nil
### GetConfirmUrl

`func (o *AuthenticationTokenDto) GetConfirmUrl() string`

GetConfirmUrl returns the ConfirmUrl field if non-nil, zero value otherwise.

### GetConfirmUrlOk

`func (o *AuthenticationTokenDto) GetConfirmUrlOk() (*string, bool)`

GetConfirmUrlOk returns a tuple with the ConfirmUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmUrl

`func (o *AuthenticationTokenDto) SetConfirmUrl(v string)`

SetConfirmUrl sets ConfirmUrl field to given value.

### HasConfirmUrl

`func (o *AuthenticationTokenDto) HasConfirmUrl() bool`

HasConfirmUrl returns a boolean if a field has been set.

### SetConfirmUrlNil

`func (o *AuthenticationTokenDto) SetConfirmUrlNil(b bool)`

 SetConfirmUrlNil sets the value for ConfirmUrl to be an explicit nil

### UnsetConfirmUrl
`func (o *AuthenticationTokenDto) UnsetConfirmUrl()`

UnsetConfirmUrl ensures that no value is present for ConfirmUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


