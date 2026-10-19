# UpdateFileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** | The new title of the file, without an extension - the stored extension is kept whatever the title says, so a  rename cannot change the format. Left empty, the file keeps its name. | [optional] 
**LastVersion** | Pointer to **int32** | The version to restore on top of the history, as reported by `GET api/2.0/files/file/{fileId}/history`; 0 or  less leaves the versions untouched. | [optional] 

## Methods

### NewUpdateFileRequest

`func NewUpdateFileRequest() *UpdateFileRequest`

NewUpdateFileRequest instantiates a new UpdateFileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFileRequestWithDefaults

`func NewUpdateFileRequestWithDefaults() *UpdateFileRequest`

NewUpdateFileRequestWithDefaults instantiates a new UpdateFileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *UpdateFileRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateFileRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateFileRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpdateFileRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *UpdateFileRequest) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *UpdateFileRequest) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetLastVersion

`func (o *UpdateFileRequest) GetLastVersion() int32`

GetLastVersion returns the LastVersion field if non-nil, zero value otherwise.

### GetLastVersionOk

`func (o *UpdateFileRequest) GetLastVersionOk() (*int32, bool)`

GetLastVersionOk returns a tuple with the LastVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastVersion

`func (o *UpdateFileRequest) SetLastVersion(v int32)`

SetLastVersion sets LastVersion field to given value.

### HasLastVersion

`func (o *UpdateFileRequest) HasLastVersion() bool`

HasLastVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


