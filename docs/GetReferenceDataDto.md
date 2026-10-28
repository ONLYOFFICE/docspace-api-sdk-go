# GetReferenceDataDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileKey** | **NullableString** | The id of the referenced file as the document service recorded it in the formula. It is tried first, and only  when `instanceId` names this portal. | 
**InstanceId** | **NullableString** | The portal the reference was made on, as the document service recorded it. Only the id of this portal makes  the file key resolvable; any other value falls through to the path and the link. | 
**SourceFileId** | Pointer to **int32** | The spreadsheet the formula sits in. The path is resolved against it - the referenced file is looked for among  the files lying next to it - and it is the file whose read access is checked. | [optional] 
**Path** | Pointer to **NullableString** | The title of the referenced file exactly as the formula spells it, matched against the files lying next to the  source file. It is tried after the file key, and only when no link is given. | [optional] 
**Link** | Pointer to **NullableString** | The web address the formula points at, an editor link of this portal or one of its short links. It is tried  last, and an address belonging to another site is not resolved at all but handed back for the client to follow  as it is. | [optional] 

## Methods

### NewGetReferenceDataDto

`func NewGetReferenceDataDto(fileKey NullableString, instanceId NullableString, ) *GetReferenceDataDto`

NewGetReferenceDataDto instantiates a new GetReferenceDataDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetReferenceDataDtoWithDefaults

`func NewGetReferenceDataDtoWithDefaults() *GetReferenceDataDto`

NewGetReferenceDataDtoWithDefaults instantiates a new GetReferenceDataDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileKey

`func (o *GetReferenceDataDto) GetFileKey() string`

GetFileKey returns the FileKey field if non-nil, zero value otherwise.

### GetFileKeyOk

`func (o *GetReferenceDataDto) GetFileKeyOk() (*string, bool)`

GetFileKeyOk returns a tuple with the FileKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileKey

`func (o *GetReferenceDataDto) SetFileKey(v string)`

SetFileKey sets FileKey field to given value.


### SetFileKeyNil

`func (o *GetReferenceDataDto) SetFileKeyNil(b bool)`

 SetFileKeyNil sets the value for FileKey to be an explicit nil

### UnsetFileKey
`func (o *GetReferenceDataDto) UnsetFileKey()`

UnsetFileKey ensures that no value is present for FileKey, not even an explicit nil
### GetInstanceId

`func (o *GetReferenceDataDto) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *GetReferenceDataDto) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *GetReferenceDataDto) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### SetInstanceIdNil

`func (o *GetReferenceDataDto) SetInstanceIdNil(b bool)`

 SetInstanceIdNil sets the value for InstanceId to be an explicit nil

### UnsetInstanceId
`func (o *GetReferenceDataDto) UnsetInstanceId()`

UnsetInstanceId ensures that no value is present for InstanceId, not even an explicit nil
### GetSourceFileId

`func (o *GetReferenceDataDto) GetSourceFileId() int32`

GetSourceFileId returns the SourceFileId field if non-nil, zero value otherwise.

### GetSourceFileIdOk

`func (o *GetReferenceDataDto) GetSourceFileIdOk() (*int32, bool)`

GetSourceFileIdOk returns a tuple with the SourceFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFileId

`func (o *GetReferenceDataDto) SetSourceFileId(v int32)`

SetSourceFileId sets SourceFileId field to given value.

### HasSourceFileId

`func (o *GetReferenceDataDto) HasSourceFileId() bool`

HasSourceFileId returns a boolean if a field has been set.

### GetPath

`func (o *GetReferenceDataDto) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GetReferenceDataDto) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GetReferenceDataDto) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *GetReferenceDataDto) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *GetReferenceDataDto) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *GetReferenceDataDto) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetLink

`func (o *GetReferenceDataDto) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *GetReferenceDataDto) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *GetReferenceDataDto) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *GetReferenceDataDto) HasLink() bool`

HasLink returns a boolean if a field has been set.

### SetLinkNil

`func (o *GetReferenceDataDto) SetLinkNil(b bool)`

 SetLinkNil sets the value for Link to be an explicit nil

### UnsetLink
`func (o *GetReferenceDataDto) UnsetLink()`

UnsetLink ensures that no value is present for Link, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


