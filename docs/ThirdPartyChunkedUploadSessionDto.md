# ThirdPartyChunkedUploadSessionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The identifier of the reserved upload, repeated in the path of every call that follows it - the chunk uploads,  the finalize and the abort. It is thirty-two hexadecimal characters without separators, and it is the only  thing the server checks, so anyone holding it can write into this upload. | [optional] 
**Path** | Pointer to **[]string** | The chain of folders leading to the destination, outermost first and the destination itself last, with folders  the caller cannot read left out. An answer that reports a stored part carries the destination folder alone  instead of the whole chain. | [optional] 
**Created** | Pointer to **time.Time** | The moment the upload was reserved, in UTC. | [optional] 
**Expired** | Pointer to **time.Time** | The moment the reservation lapses and the parts buffered for it are dropped, in UTC. It is a gap rather than a  deadline for the whole transfer: every accepted part pushes it twelve hours past that part, so only a long  silence loses the upload. | [optional] 
**Location** | Pointer to **NullableString** | The absolute address of the separate chunk handler that also accepts the parts of this upload, kept for  clients written against it. A caller working through this API does not need it and sends the parts to the  session operations instead. | [optional] 
**BytesTotal** | Pointer to **int64** | The size in bytes that was declared when the upload was reserved, echoed back. It is what the arriving parts  are counted against to decide the file is complete, not the amount received so far. | [optional] 

## Methods

### NewThirdPartyChunkedUploadSessionDto

`func NewThirdPartyChunkedUploadSessionDto() *ThirdPartyChunkedUploadSessionDto`

NewThirdPartyChunkedUploadSessionDto instantiates a new ThirdPartyChunkedUploadSessionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThirdPartyChunkedUploadSessionDtoWithDefaults

`func NewThirdPartyChunkedUploadSessionDtoWithDefaults() *ThirdPartyChunkedUploadSessionDto`

NewThirdPartyChunkedUploadSessionDtoWithDefaults instantiates a new ThirdPartyChunkedUploadSessionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ThirdPartyChunkedUploadSessionDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ThirdPartyChunkedUploadSessionDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ThirdPartyChunkedUploadSessionDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ThirdPartyChunkedUploadSessionDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ThirdPartyChunkedUploadSessionDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ThirdPartyChunkedUploadSessionDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetPath

`func (o *ThirdPartyChunkedUploadSessionDto) GetPath() []string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ThirdPartyChunkedUploadSessionDto) GetPathOk() (*[]string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ThirdPartyChunkedUploadSessionDto) SetPath(v []string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ThirdPartyChunkedUploadSessionDto) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *ThirdPartyChunkedUploadSessionDto) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *ThirdPartyChunkedUploadSessionDto) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetCreated

`func (o *ThirdPartyChunkedUploadSessionDto) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ThirdPartyChunkedUploadSessionDto) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ThirdPartyChunkedUploadSessionDto) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ThirdPartyChunkedUploadSessionDto) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetExpired

`func (o *ThirdPartyChunkedUploadSessionDto) GetExpired() time.Time`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *ThirdPartyChunkedUploadSessionDto) GetExpiredOk() (*time.Time, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *ThirdPartyChunkedUploadSessionDto) SetExpired(v time.Time)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *ThirdPartyChunkedUploadSessionDto) HasExpired() bool`

HasExpired returns a boolean if a field has been set.

### GetLocation

`func (o *ThirdPartyChunkedUploadSessionDto) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ThirdPartyChunkedUploadSessionDto) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ThirdPartyChunkedUploadSessionDto) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ThirdPartyChunkedUploadSessionDto) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *ThirdPartyChunkedUploadSessionDto) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *ThirdPartyChunkedUploadSessionDto) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetBytesTotal

`func (o *ThirdPartyChunkedUploadSessionDto) GetBytesTotal() int64`

GetBytesTotal returns the BytesTotal field if non-nil, zero value otherwise.

### GetBytesTotalOk

`func (o *ThirdPartyChunkedUploadSessionDto) GetBytesTotalOk() (*int64, bool)`

GetBytesTotalOk returns a tuple with the BytesTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesTotal

`func (o *ThirdPartyChunkedUploadSessionDto) SetBytesTotal(v int64)`

SetBytesTotal sets BytesTotal field to given value.

### HasBytesTotal

`func (o *ThirdPartyChunkedUploadSessionDto) HasBytesTotal() bool`

HasBytesTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


