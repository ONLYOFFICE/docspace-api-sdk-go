# FirebaseRequestsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirebaseDeviceToken** | Pointer to **NullableString** | The Firebase device token. | [optional] 
**IsSubscribed** | Pointer to **bool** | Specifies whether the user is subscribed to the push notifications or not. | [optional] 

## Methods

### NewFirebaseRequestsDto

`func NewFirebaseRequestsDto() *FirebaseRequestsDto`

NewFirebaseRequestsDto instantiates a new FirebaseRequestsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirebaseRequestsDtoWithDefaults

`func NewFirebaseRequestsDtoWithDefaults() *FirebaseRequestsDto`

NewFirebaseRequestsDtoWithDefaults instantiates a new FirebaseRequestsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirebaseDeviceToken

`func (o *FirebaseRequestsDto) GetFirebaseDeviceToken() string`

GetFirebaseDeviceToken returns the FirebaseDeviceToken field if non-nil, zero value otherwise.

### GetFirebaseDeviceTokenOk

`func (o *FirebaseRequestsDto) GetFirebaseDeviceTokenOk() (*string, bool)`

GetFirebaseDeviceTokenOk returns a tuple with the FirebaseDeviceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirebaseDeviceToken

`func (o *FirebaseRequestsDto) SetFirebaseDeviceToken(v string)`

SetFirebaseDeviceToken sets FirebaseDeviceToken field to given value.

### HasFirebaseDeviceToken

`func (o *FirebaseRequestsDto) HasFirebaseDeviceToken() bool`

HasFirebaseDeviceToken returns a boolean if a field has been set.

### SetFirebaseDeviceTokenNil

`func (o *FirebaseRequestsDto) SetFirebaseDeviceTokenNil(b bool)`

 SetFirebaseDeviceTokenNil sets the value for FirebaseDeviceToken to be an explicit nil

### UnsetFirebaseDeviceToken
`func (o *FirebaseRequestsDto) UnsetFirebaseDeviceToken()`

UnsetFirebaseDeviceToken ensures that no value is present for FirebaseDeviceToken, not even an explicit nil
### GetIsSubscribed

`func (o *FirebaseRequestsDto) GetIsSubscribed() bool`

GetIsSubscribed returns the IsSubscribed field if non-nil, zero value otherwise.

### GetIsSubscribedOk

`func (o *FirebaseRequestsDto) GetIsSubscribedOk() (*bool, bool)`

GetIsSubscribedOk returns a tuple with the IsSubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubscribed

`func (o *FirebaseRequestsDto) SetIsSubscribed(v bool)`

SetIsSubscribed sets IsSubscribed field to given value.

### HasIsSubscribed

`func (o *FirebaseRequestsDto) HasIsSubscribed() bool`

HasIsSubscribed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


