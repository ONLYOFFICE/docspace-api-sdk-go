# FirebaseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | **NullableString** | The Firebase API key. | 
**AuthDomain** | **NullableString** | The Firebase authentication domain. | 
**ProjectId** | **NullableString** | The Firebase project ID. | 
**StorageBucket** | **NullableString** | The Firebase storage bucket. | 
**MessagingSenderId** | **NullableString** | The Firebase messaging sender ID. | 
**AppId** | **NullableString** | The Firebase application ID. | 
**MeasurementId** | **NullableString** | The Firebase measurement ID. | 
**DatabaseURL** | **NullableString** | The Firebase database URL. | 

## Methods

### NewFirebaseDto

`func NewFirebaseDto(apiKey NullableString, authDomain NullableString, projectId NullableString, storageBucket NullableString, messagingSenderId NullableString, appId NullableString, measurementId NullableString, databaseURL NullableString, ) *FirebaseDto`

NewFirebaseDto instantiates a new FirebaseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirebaseDtoWithDefaults

`func NewFirebaseDtoWithDefaults() *FirebaseDto`

NewFirebaseDtoWithDefaults instantiates a new FirebaseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *FirebaseDto) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *FirebaseDto) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *FirebaseDto) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### SetApiKeyNil

`func (o *FirebaseDto) SetApiKeyNil(b bool)`

 SetApiKeyNil sets the value for ApiKey to be an explicit nil

### UnsetApiKey
`func (o *FirebaseDto) UnsetApiKey()`

UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
### GetAuthDomain

`func (o *FirebaseDto) GetAuthDomain() string`

GetAuthDomain returns the AuthDomain field if non-nil, zero value otherwise.

### GetAuthDomainOk

`func (o *FirebaseDto) GetAuthDomainOk() (*string, bool)`

GetAuthDomainOk returns a tuple with the AuthDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthDomain

`func (o *FirebaseDto) SetAuthDomain(v string)`

SetAuthDomain sets AuthDomain field to given value.


### SetAuthDomainNil

`func (o *FirebaseDto) SetAuthDomainNil(b bool)`

 SetAuthDomainNil sets the value for AuthDomain to be an explicit nil

### UnsetAuthDomain
`func (o *FirebaseDto) UnsetAuthDomain()`

UnsetAuthDomain ensures that no value is present for AuthDomain, not even an explicit nil
### GetProjectId

`func (o *FirebaseDto) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *FirebaseDto) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *FirebaseDto) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### SetProjectIdNil

`func (o *FirebaseDto) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *FirebaseDto) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetStorageBucket

`func (o *FirebaseDto) GetStorageBucket() string`

GetStorageBucket returns the StorageBucket field if non-nil, zero value otherwise.

### GetStorageBucketOk

`func (o *FirebaseDto) GetStorageBucketOk() (*string, bool)`

GetStorageBucketOk returns a tuple with the StorageBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageBucket

`func (o *FirebaseDto) SetStorageBucket(v string)`

SetStorageBucket sets StorageBucket field to given value.


### SetStorageBucketNil

`func (o *FirebaseDto) SetStorageBucketNil(b bool)`

 SetStorageBucketNil sets the value for StorageBucket to be an explicit nil

### UnsetStorageBucket
`func (o *FirebaseDto) UnsetStorageBucket()`

UnsetStorageBucket ensures that no value is present for StorageBucket, not even an explicit nil
### GetMessagingSenderId

`func (o *FirebaseDto) GetMessagingSenderId() string`

GetMessagingSenderId returns the MessagingSenderId field if non-nil, zero value otherwise.

### GetMessagingSenderIdOk

`func (o *FirebaseDto) GetMessagingSenderIdOk() (*string, bool)`

GetMessagingSenderIdOk returns a tuple with the MessagingSenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingSenderId

`func (o *FirebaseDto) SetMessagingSenderId(v string)`

SetMessagingSenderId sets MessagingSenderId field to given value.


### SetMessagingSenderIdNil

`func (o *FirebaseDto) SetMessagingSenderIdNil(b bool)`

 SetMessagingSenderIdNil sets the value for MessagingSenderId to be an explicit nil

### UnsetMessagingSenderId
`func (o *FirebaseDto) UnsetMessagingSenderId()`

UnsetMessagingSenderId ensures that no value is present for MessagingSenderId, not even an explicit nil
### GetAppId

`func (o *FirebaseDto) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *FirebaseDto) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *FirebaseDto) SetAppId(v string)`

SetAppId sets AppId field to given value.


### SetAppIdNil

`func (o *FirebaseDto) SetAppIdNil(b bool)`

 SetAppIdNil sets the value for AppId to be an explicit nil

### UnsetAppId
`func (o *FirebaseDto) UnsetAppId()`

UnsetAppId ensures that no value is present for AppId, not even an explicit nil
### GetMeasurementId

`func (o *FirebaseDto) GetMeasurementId() string`

GetMeasurementId returns the MeasurementId field if non-nil, zero value otherwise.

### GetMeasurementIdOk

`func (o *FirebaseDto) GetMeasurementIdOk() (*string, bool)`

GetMeasurementIdOk returns a tuple with the MeasurementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurementId

`func (o *FirebaseDto) SetMeasurementId(v string)`

SetMeasurementId sets MeasurementId field to given value.


### SetMeasurementIdNil

`func (o *FirebaseDto) SetMeasurementIdNil(b bool)`

 SetMeasurementIdNil sets the value for MeasurementId to be an explicit nil

### UnsetMeasurementId
`func (o *FirebaseDto) UnsetMeasurementId()`

UnsetMeasurementId ensures that no value is present for MeasurementId, not even an explicit nil
### GetDatabaseURL

`func (o *FirebaseDto) GetDatabaseURL() string`

GetDatabaseURL returns the DatabaseURL field if non-nil, zero value otherwise.

### GetDatabaseURLOk

`func (o *FirebaseDto) GetDatabaseURLOk() (*string, bool)`

GetDatabaseURLOk returns a tuple with the DatabaseURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseURL

`func (o *FirebaseDto) SetDatabaseURL(v string)`

SetDatabaseURL sets DatabaseURL field to given value.


### SetDatabaseURLNil

`func (o *FirebaseDto) SetDatabaseURLNil(b bool)`

 SetDatabaseURLNil sets the value for DatabaseURL to be an explicit nil

### UnsetDatabaseURL
`func (o *FirebaseDto) UnsetDatabaseURL()`

UnsetDatabaseURL ensures that no value is present for DatabaseURL, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


