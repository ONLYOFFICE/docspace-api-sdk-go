# UpdatePhotoMemberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | Pointer to **NullableString** | The address the portal downloads the new avatar from. It has to be absolute or relative to the portal, and it  has to use HTTPS unless the request itself came over HTTP; an address the portal refuses to fetch is rejected.  It is required - an empty value is answered with 400 rather than clearing the avatar. | [optional] 

## Methods

### NewUpdatePhotoMemberRequest

`func NewUpdatePhotoMemberRequest() *UpdatePhotoMemberRequest`

NewUpdatePhotoMemberRequest instantiates a new UpdatePhotoMemberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePhotoMemberRequestWithDefaults

`func NewUpdatePhotoMemberRequestWithDefaults() *UpdatePhotoMemberRequest`

NewUpdatePhotoMemberRequestWithDefaults instantiates a new UpdatePhotoMemberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *UpdatePhotoMemberRequest) GetFiles() string`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *UpdatePhotoMemberRequest) GetFilesOk() (*string, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *UpdatePhotoMemberRequest) SetFiles(v string)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *UpdatePhotoMemberRequest) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### SetFilesNil

`func (o *UpdatePhotoMemberRequest) SetFilesNil(b bool)`

 SetFilesNil sets the value for Files to be an explicit nil

### UnsetFiles
`func (o *UpdatePhotoMemberRequest) UnsetFiles()`

UnsetFiles ensures that no value is present for Files, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


