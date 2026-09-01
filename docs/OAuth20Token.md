# OAuth20Token

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **NullableString** | Access token | [optional] 
**RefreshToken** | Pointer to **NullableString** | Refresh token | [optional] 
**ExpiresIn** | Pointer to **int64** | Expires in | [optional] 
**ClientId** | Pointer to **NullableString** | Client id | [optional] 
**ClientSecret** | Pointer to **NullableString** | Client secret | [optional] 
**RedirectUri** | Pointer to **NullableString** | Redirect uri | [optional] 
**Timestamp** | Pointer to **time.Time** | Timestamp | [optional] 
**IsExpired** | Pointer to **bool** | Is expired | [optional] [readonly] 

## Methods

### NewOAuth20Token

`func NewOAuth20Token() *OAuth20Token`

NewOAuth20Token instantiates a new OAuth20Token object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth20TokenWithDefaults

`func NewOAuth20TokenWithDefaults() *OAuth20Token`

NewOAuth20TokenWithDefaults instantiates a new OAuth20Token object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *OAuth20Token) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *OAuth20Token) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *OAuth20Token) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *OAuth20Token) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### SetAccessTokenNil

`func (o *OAuth20Token) SetAccessTokenNil(b bool)`

 SetAccessTokenNil sets the value for AccessToken to be an explicit nil

### UnsetAccessToken
`func (o *OAuth20Token) UnsetAccessToken()`

UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
### GetRefreshToken

`func (o *OAuth20Token) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *OAuth20Token) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *OAuth20Token) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *OAuth20Token) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### SetRefreshTokenNil

`func (o *OAuth20Token) SetRefreshTokenNil(b bool)`

 SetRefreshTokenNil sets the value for RefreshToken to be an explicit nil

### UnsetRefreshToken
`func (o *OAuth20Token) UnsetRefreshToken()`

UnsetRefreshToken ensures that no value is present for RefreshToken, not even an explicit nil
### GetExpiresIn

`func (o *OAuth20Token) GetExpiresIn() int64`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *OAuth20Token) GetExpiresInOk() (*int64, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *OAuth20Token) SetExpiresIn(v int64)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *OAuth20Token) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetClientId

`func (o *OAuth20Token) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OAuth20Token) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OAuth20Token) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *OAuth20Token) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### SetClientIdNil

`func (o *OAuth20Token) SetClientIdNil(b bool)`

 SetClientIdNil sets the value for ClientId to be an explicit nil

### UnsetClientId
`func (o *OAuth20Token) UnsetClientId()`

UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
### GetClientSecret

`func (o *OAuth20Token) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *OAuth20Token) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *OAuth20Token) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *OAuth20Token) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### SetClientSecretNil

`func (o *OAuth20Token) SetClientSecretNil(b bool)`

 SetClientSecretNil sets the value for ClientSecret to be an explicit nil

### UnsetClientSecret
`func (o *OAuth20Token) UnsetClientSecret()`

UnsetClientSecret ensures that no value is present for ClientSecret, not even an explicit nil
### GetRedirectUri

`func (o *OAuth20Token) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *OAuth20Token) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *OAuth20Token) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.

### HasRedirectUri

`func (o *OAuth20Token) HasRedirectUri() bool`

HasRedirectUri returns a boolean if a field has been set.

### SetRedirectUriNil

`func (o *OAuth20Token) SetRedirectUriNil(b bool)`

 SetRedirectUriNil sets the value for RedirectUri to be an explicit nil

### UnsetRedirectUri
`func (o *OAuth20Token) UnsetRedirectUri()`

UnsetRedirectUri ensures that no value is present for RedirectUri, not even an explicit nil
### GetTimestamp

`func (o *OAuth20Token) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *OAuth20Token) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *OAuth20Token) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *OAuth20Token) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetIsExpired

`func (o *OAuth20Token) GetIsExpired() bool`

GetIsExpired returns the IsExpired field if non-nil, zero value otherwise.

### GetIsExpiredOk

`func (o *OAuth20Token) GetIsExpiredOk() (*bool, bool)`

GetIsExpiredOk returns a tuple with the IsExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExpired

`func (o *OAuth20Token) SetIsExpired(v bool)`

SetIsExpired sets IsExpired field to given value.

### HasIsExpired

`func (o *OAuth20Token) HasIsExpired() bool`

HasIsExpired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


