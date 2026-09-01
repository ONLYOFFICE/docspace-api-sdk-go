# FireBaseUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The Firebase user ID. | [optional] 
**UserId** | Pointer to **string** | The user ID. | [optional] 
**TenantId** | Pointer to **int32** | The tenant ID. | [optional] 
**FirebaseDeviceToken** | Pointer to **NullableString** | The Firebase device token. | [optional] 
**Application** | Pointer to **NullableString** | The Firebase application. | [optional] 
**IsSubscribed** | Pointer to **NullableBool** | Specifies if the user is subscribed to the push notifications or not. | [optional] 
**Tenant** | Pointer to [**DbTenant**](DbTenant.md) | The database tenant parameters. | [optional] 

## Methods

### NewFireBaseUser

`func NewFireBaseUser() *FireBaseUser`

NewFireBaseUser instantiates a new FireBaseUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFireBaseUserWithDefaults

`func NewFireBaseUserWithDefaults() *FireBaseUser`

NewFireBaseUserWithDefaults instantiates a new FireBaseUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FireBaseUser) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FireBaseUser) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FireBaseUser) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FireBaseUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *FireBaseUser) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *FireBaseUser) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *FireBaseUser) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *FireBaseUser) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetTenantId

`func (o *FireBaseUser) GetTenantId() int32`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *FireBaseUser) GetTenantIdOk() (*int32, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *FireBaseUser) SetTenantId(v int32)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *FireBaseUser) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetFirebaseDeviceToken

`func (o *FireBaseUser) GetFirebaseDeviceToken() string`

GetFirebaseDeviceToken returns the FirebaseDeviceToken field if non-nil, zero value otherwise.

### GetFirebaseDeviceTokenOk

`func (o *FireBaseUser) GetFirebaseDeviceTokenOk() (*string, bool)`

GetFirebaseDeviceTokenOk returns a tuple with the FirebaseDeviceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirebaseDeviceToken

`func (o *FireBaseUser) SetFirebaseDeviceToken(v string)`

SetFirebaseDeviceToken sets FirebaseDeviceToken field to given value.

### HasFirebaseDeviceToken

`func (o *FireBaseUser) HasFirebaseDeviceToken() bool`

HasFirebaseDeviceToken returns a boolean if a field has been set.

### SetFirebaseDeviceTokenNil

`func (o *FireBaseUser) SetFirebaseDeviceTokenNil(b bool)`

 SetFirebaseDeviceTokenNil sets the value for FirebaseDeviceToken to be an explicit nil

### UnsetFirebaseDeviceToken
`func (o *FireBaseUser) UnsetFirebaseDeviceToken()`

UnsetFirebaseDeviceToken ensures that no value is present for FirebaseDeviceToken, not even an explicit nil
### GetApplication

`func (o *FireBaseUser) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *FireBaseUser) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *FireBaseUser) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *FireBaseUser) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### SetApplicationNil

`func (o *FireBaseUser) SetApplicationNil(b bool)`

 SetApplicationNil sets the value for Application to be an explicit nil

### UnsetApplication
`func (o *FireBaseUser) UnsetApplication()`

UnsetApplication ensures that no value is present for Application, not even an explicit nil
### GetIsSubscribed

`func (o *FireBaseUser) GetIsSubscribed() bool`

GetIsSubscribed returns the IsSubscribed field if non-nil, zero value otherwise.

### GetIsSubscribedOk

`func (o *FireBaseUser) GetIsSubscribedOk() (*bool, bool)`

GetIsSubscribedOk returns a tuple with the IsSubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubscribed

`func (o *FireBaseUser) SetIsSubscribed(v bool)`

SetIsSubscribed sets IsSubscribed field to given value.

### HasIsSubscribed

`func (o *FireBaseUser) HasIsSubscribed() bool`

HasIsSubscribed returns a boolean if a field has been set.

### SetIsSubscribedNil

`func (o *FireBaseUser) SetIsSubscribedNil(b bool)`

 SetIsSubscribedNil sets the value for IsSubscribed to be an explicit nil

### UnsetIsSubscribed
`func (o *FireBaseUser) UnsetIsSubscribed()`

UnsetIsSubscribed ensures that no value is present for IsSubscribed, not even an explicit nil
### GetTenant

`func (o *FireBaseUser) GetTenant() DbTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *FireBaseUser) GetTenantOk() (*DbTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *FireBaseUser) SetTenant(v DbTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *FireBaseUser) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


