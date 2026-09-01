# GetReferenceDataDtoInteger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileKey** | **NullableString** | The unique document identifier used by the service to get a link to the file. | 
**InstanceId** | **NullableString** | The unique system identifier. | 
**SourceFileId** | Pointer to **int32** | The source file ID. | [optional] 
**Path** | Pointer to **NullableString** | The file name or relative path for the formula editor. | [optional] 
**Link** | Pointer to **NullableString** | The file link. | [optional] 

## Methods

### NewGetReferenceDataDtoInteger

`func NewGetReferenceDataDtoInteger(fileKey NullableString, instanceId NullableString, ) *GetReferenceDataDtoInteger`

NewGetReferenceDataDtoInteger instantiates a new GetReferenceDataDtoInteger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetReferenceDataDtoIntegerWithDefaults

`func NewGetReferenceDataDtoIntegerWithDefaults() *GetReferenceDataDtoInteger`

NewGetReferenceDataDtoIntegerWithDefaults instantiates a new GetReferenceDataDtoInteger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileKey

`func (o *GetReferenceDataDtoInteger) GetFileKey() string`

GetFileKey returns the FileKey field if non-nil, zero value otherwise.

### GetFileKeyOk

`func (o *GetReferenceDataDtoInteger) GetFileKeyOk() (*string, bool)`

GetFileKeyOk returns a tuple with the FileKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileKey

`func (o *GetReferenceDataDtoInteger) SetFileKey(v string)`

SetFileKey sets FileKey field to given value.


### SetFileKeyNil

`func (o *GetReferenceDataDtoInteger) SetFileKeyNil(b bool)`

 SetFileKeyNil sets the value for FileKey to be an explicit nil

### UnsetFileKey
`func (o *GetReferenceDataDtoInteger) UnsetFileKey()`

UnsetFileKey ensures that no value is present for FileKey, not even an explicit nil
### GetInstanceId

`func (o *GetReferenceDataDtoInteger) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *GetReferenceDataDtoInteger) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *GetReferenceDataDtoInteger) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### SetInstanceIdNil

`func (o *GetReferenceDataDtoInteger) SetInstanceIdNil(b bool)`

 SetInstanceIdNil sets the value for InstanceId to be an explicit nil

### UnsetInstanceId
`func (o *GetReferenceDataDtoInteger) UnsetInstanceId()`

UnsetInstanceId ensures that no value is present for InstanceId, not even an explicit nil
### GetSourceFileId

`func (o *GetReferenceDataDtoInteger) GetSourceFileId() int32`

GetSourceFileId returns the SourceFileId field if non-nil, zero value otherwise.

### GetSourceFileIdOk

`func (o *GetReferenceDataDtoInteger) GetSourceFileIdOk() (*int32, bool)`

GetSourceFileIdOk returns a tuple with the SourceFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFileId

`func (o *GetReferenceDataDtoInteger) SetSourceFileId(v int32)`

SetSourceFileId sets SourceFileId field to given value.

### HasSourceFileId

`func (o *GetReferenceDataDtoInteger) HasSourceFileId() bool`

HasSourceFileId returns a boolean if a field has been set.

### GetPath

`func (o *GetReferenceDataDtoInteger) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GetReferenceDataDtoInteger) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GetReferenceDataDtoInteger) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *GetReferenceDataDtoInteger) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *GetReferenceDataDtoInteger) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *GetReferenceDataDtoInteger) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetLink

`func (o *GetReferenceDataDtoInteger) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *GetReferenceDataDtoInteger) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *GetReferenceDataDtoInteger) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *GetReferenceDataDtoInteger) HasLink() bool`

HasLink returns a boolean if a field has been set.

### SetLinkNil

`func (o *GetReferenceDataDtoInteger) SetLinkNil(b bool)`

 SetLinkNil sets the value for Link to be an explicit nil

### UnsetLink
`func (o *GetReferenceDataDtoInteger) UnsetLink()`

UnsetLink ensures that no value is present for Link, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


