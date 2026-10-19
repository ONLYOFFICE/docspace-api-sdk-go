# FirebaseRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirebaseDeviceToken** | Pointer to **NullableString** | The registration token Firebase issued to the mobile client for this device, obtained on the device itself.  It is kept as an opaque string of up to 255 characters and is never verified here; it identifies the device  and is matched but never changed, and a token belonging to another member or another portal matches nothing. | [optional] 
**IsSubscribed** | Pointer to **bool** | Whether the device is to receive the room activity messages - an invitation, a role change, an archived room,  a new document. On a first registration it is stored as given; on a registration that already exists it is  ignored, because registering does not update, and the subscription is changed with  `PUT api/2.0/settings/push/docsubscribe` instead. | [optional] 

## Methods

### NewFirebaseRequestDto

`func NewFirebaseRequestDto() *FirebaseRequestDto`

NewFirebaseRequestDto instantiates a new FirebaseRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirebaseRequestDtoWithDefaults

`func NewFirebaseRequestDtoWithDefaults() *FirebaseRequestDto`

NewFirebaseRequestDtoWithDefaults instantiates a new FirebaseRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirebaseDeviceToken

`func (o *FirebaseRequestDto) GetFirebaseDeviceToken() string`

GetFirebaseDeviceToken returns the FirebaseDeviceToken field if non-nil, zero value otherwise.

### GetFirebaseDeviceTokenOk

`func (o *FirebaseRequestDto) GetFirebaseDeviceTokenOk() (*string, bool)`

GetFirebaseDeviceTokenOk returns a tuple with the FirebaseDeviceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirebaseDeviceToken

`func (o *FirebaseRequestDto) SetFirebaseDeviceToken(v string)`

SetFirebaseDeviceToken sets FirebaseDeviceToken field to given value.

### HasFirebaseDeviceToken

`func (o *FirebaseRequestDto) HasFirebaseDeviceToken() bool`

HasFirebaseDeviceToken returns a boolean if a field has been set.

### SetFirebaseDeviceTokenNil

`func (o *FirebaseRequestDto) SetFirebaseDeviceTokenNil(b bool)`

 SetFirebaseDeviceTokenNil sets the value for FirebaseDeviceToken to be an explicit nil

### UnsetFirebaseDeviceToken
`func (o *FirebaseRequestDto) UnsetFirebaseDeviceToken()`

UnsetFirebaseDeviceToken ensures that no value is present for FirebaseDeviceToken, not even an explicit nil
### GetIsSubscribed

`func (o *FirebaseRequestDto) GetIsSubscribed() bool`

GetIsSubscribed returns the IsSubscribed field if non-nil, zero value otherwise.

### GetIsSubscribedOk

`func (o *FirebaseRequestDto) GetIsSubscribedOk() (*bool, bool)`

GetIsSubscribedOk returns a tuple with the IsSubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubscribed

`func (o *FirebaseRequestDto) SetIsSubscribed(v bool)`

SetIsSubscribed sets IsSubscribed field to given value.

### HasIsSubscribed

`func (o *FirebaseRequestDto) HasIsSubscribed() bool`

HasIsSubscribed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


