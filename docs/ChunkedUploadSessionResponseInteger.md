# ChunkedUploadSessionResponseInteger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The unique identifier for the entity. | [optional] 
**Path** | Pointer to **[]int32** | Represents the hierarchical path of folders associated with a chunked upload session. | [optional] 
**Created** | Pointer to **time.Time** | The timestamp indicating when the chunked upload session was created. | [optional] 
**Expired** | Pointer to **time.Time** | The date and time when the chunked upload session is set to expire. | [optional] 
**Location** | Pointer to **NullableString** | Represents the URI or path of the chunked upload session's current location. | [optional] 
**BytesTotal** | Pointer to **int64** | The total size, in bytes, of the file being uploaded in the chunked upload session. | [optional] 

## Methods

### NewChunkedUploadSessionResponseInteger

`func NewChunkedUploadSessionResponseInteger() *ChunkedUploadSessionResponseInteger`

NewChunkedUploadSessionResponseInteger instantiates a new ChunkedUploadSessionResponseInteger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChunkedUploadSessionResponseIntegerWithDefaults

`func NewChunkedUploadSessionResponseIntegerWithDefaults() *ChunkedUploadSessionResponseInteger`

NewChunkedUploadSessionResponseIntegerWithDefaults instantiates a new ChunkedUploadSessionResponseInteger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChunkedUploadSessionResponseInteger) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChunkedUploadSessionResponseInteger) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChunkedUploadSessionResponseInteger) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ChunkedUploadSessionResponseInteger) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ChunkedUploadSessionResponseInteger) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ChunkedUploadSessionResponseInteger) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetPath

`func (o *ChunkedUploadSessionResponseInteger) GetPath() []int32`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ChunkedUploadSessionResponseInteger) GetPathOk() (*[]int32, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ChunkedUploadSessionResponseInteger) SetPath(v []int32)`

SetPath sets Path field to given value.

### HasPath

`func (o *ChunkedUploadSessionResponseInteger) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *ChunkedUploadSessionResponseInteger) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *ChunkedUploadSessionResponseInteger) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetCreated

`func (o *ChunkedUploadSessionResponseInteger) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ChunkedUploadSessionResponseInteger) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ChunkedUploadSessionResponseInteger) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ChunkedUploadSessionResponseInteger) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetExpired

`func (o *ChunkedUploadSessionResponseInteger) GetExpired() time.Time`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *ChunkedUploadSessionResponseInteger) GetExpiredOk() (*time.Time, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *ChunkedUploadSessionResponseInteger) SetExpired(v time.Time)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *ChunkedUploadSessionResponseInteger) HasExpired() bool`

HasExpired returns a boolean if a field has been set.

### GetLocation

`func (o *ChunkedUploadSessionResponseInteger) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ChunkedUploadSessionResponseInteger) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ChunkedUploadSessionResponseInteger) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ChunkedUploadSessionResponseInteger) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *ChunkedUploadSessionResponseInteger) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *ChunkedUploadSessionResponseInteger) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetBytesTotal

`func (o *ChunkedUploadSessionResponseInteger) GetBytesTotal() int64`

GetBytesTotal returns the BytesTotal field if non-nil, zero value otherwise.

### GetBytesTotalOk

`func (o *ChunkedUploadSessionResponseInteger) GetBytesTotalOk() (*int64, bool)`

GetBytesTotalOk returns a tuple with the BytesTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesTotal

`func (o *ChunkedUploadSessionResponseInteger) SetBytesTotal(v int64)`

SetBytesTotal sets BytesTotal field to given value.

### HasBytesTotal

`func (o *ChunkedUploadSessionResponseInteger) HasBytesTotal() bool`

HasBytesTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


