# FirebaseDeviceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The id of the registration. | [optional] 
**UserId** | Pointer to **string** | The account the device belongs to; always the caller. | [optional] 
**TenantId** | Pointer to **int32** | The portal the registration belongs to; always the current one. | [optional] 
**FirebaseDeviceToken** | Pointer to **NullableString** | The Firebase token the device was issued, as it was sent at registration. | [optional] 
**Application** | Pointer to **NullableString** | The application the registration is for; `doc` for the Documents application. | [optional] 
**IsSubscribed** | Pointer to **NullableBool** | Whether the device is currently sent push notifications. | [optional] 

## Methods

### NewFirebaseDeviceDto

`func NewFirebaseDeviceDto() *FirebaseDeviceDto`

NewFirebaseDeviceDto instantiates a new FirebaseDeviceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirebaseDeviceDtoWithDefaults

`func NewFirebaseDeviceDtoWithDefaults() *FirebaseDeviceDto`

NewFirebaseDeviceDtoWithDefaults instantiates a new FirebaseDeviceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FirebaseDeviceDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FirebaseDeviceDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FirebaseDeviceDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FirebaseDeviceDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *FirebaseDeviceDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *FirebaseDeviceDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *FirebaseDeviceDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *FirebaseDeviceDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetTenantId

`func (o *FirebaseDeviceDto) GetTenantId() int32`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *FirebaseDeviceDto) GetTenantIdOk() (*int32, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *FirebaseDeviceDto) SetTenantId(v int32)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *FirebaseDeviceDto) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetFirebaseDeviceToken

`func (o *FirebaseDeviceDto) GetFirebaseDeviceToken() string`

GetFirebaseDeviceToken returns the FirebaseDeviceToken field if non-nil, zero value otherwise.

### GetFirebaseDeviceTokenOk

`func (o *FirebaseDeviceDto) GetFirebaseDeviceTokenOk() (*string, bool)`

GetFirebaseDeviceTokenOk returns a tuple with the FirebaseDeviceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirebaseDeviceToken

`func (o *FirebaseDeviceDto) SetFirebaseDeviceToken(v string)`

SetFirebaseDeviceToken sets FirebaseDeviceToken field to given value.

### HasFirebaseDeviceToken

`func (o *FirebaseDeviceDto) HasFirebaseDeviceToken() bool`

HasFirebaseDeviceToken returns a boolean if a field has been set.

### SetFirebaseDeviceTokenNil

`func (o *FirebaseDeviceDto) SetFirebaseDeviceTokenNil(b bool)`

 SetFirebaseDeviceTokenNil sets the value for FirebaseDeviceToken to be an explicit nil

### UnsetFirebaseDeviceToken
`func (o *FirebaseDeviceDto) UnsetFirebaseDeviceToken()`

UnsetFirebaseDeviceToken ensures that no value is present for FirebaseDeviceToken, not even an explicit nil
### GetApplication

`func (o *FirebaseDeviceDto) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *FirebaseDeviceDto) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *FirebaseDeviceDto) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *FirebaseDeviceDto) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### SetApplicationNil

`func (o *FirebaseDeviceDto) SetApplicationNil(b bool)`

 SetApplicationNil sets the value for Application to be an explicit nil

### UnsetApplication
`func (o *FirebaseDeviceDto) UnsetApplication()`

UnsetApplication ensures that no value is present for Application, not even an explicit nil
### GetIsSubscribed

`func (o *FirebaseDeviceDto) GetIsSubscribed() bool`

GetIsSubscribed returns the IsSubscribed field if non-nil, zero value otherwise.

### GetIsSubscribedOk

`func (o *FirebaseDeviceDto) GetIsSubscribedOk() (*bool, bool)`

GetIsSubscribedOk returns a tuple with the IsSubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubscribed

`func (o *FirebaseDeviceDto) SetIsSubscribed(v bool)`

SetIsSubscribed sets IsSubscribed field to given value.

### HasIsSubscribed

`func (o *FirebaseDeviceDto) HasIsSubscribed() bool`

HasIsSubscribed returns a boolean if a field has been set.

### SetIsSubscribedNil

`func (o *FirebaseDeviceDto) SetIsSubscribedNil(b bool)`

 SetIsSubscribedNil sets the value for IsSubscribed to be an explicit nil

### UnsetIsSubscribed
`func (o *FirebaseDeviceDto) UnsetIsSubscribed()`

UnsetIsSubscribed ensures that no value is present for IsSubscribed, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


