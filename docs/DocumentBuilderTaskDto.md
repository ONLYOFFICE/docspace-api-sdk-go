# DocumentBuilderTaskDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** | The Document Builder task ID. | 
**Error** | **NullableString** | The error message occurred during the document building process. | 
**Percentage** | **int32** | The progress percentage of the document building process. | 
**IsCompleted** | **bool** | Specifies whether the document building process is completed or not. | 
**Status** | [**DistributedTaskStatus**](DistributedTaskStatus.md) |  | 
**ResultFileId** | **interface{}** | The result file ID. | 
**ResultFileName** | **NullableString** | The result file name. | 
**ResultFileUrl** | **NullableString** | The result file URL. | 

## Methods

### NewDocumentBuilderTaskDto

`func NewDocumentBuilderTaskDto(id NullableString, error_ NullableString, percentage int32, isCompleted bool, status DistributedTaskStatus, resultFileId interface{}, resultFileName NullableString, resultFileUrl NullableString, ) *DocumentBuilderTaskDto`

NewDocumentBuilderTaskDto instantiates a new DocumentBuilderTaskDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentBuilderTaskDtoWithDefaults

`func NewDocumentBuilderTaskDtoWithDefaults() *DocumentBuilderTaskDto`

NewDocumentBuilderTaskDtoWithDefaults instantiates a new DocumentBuilderTaskDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DocumentBuilderTaskDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DocumentBuilderTaskDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DocumentBuilderTaskDto) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *DocumentBuilderTaskDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DocumentBuilderTaskDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetError

`func (o *DocumentBuilderTaskDto) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DocumentBuilderTaskDto) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DocumentBuilderTaskDto) SetError(v string)`

SetError sets Error field to given value.


### SetErrorNil

`func (o *DocumentBuilderTaskDto) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *DocumentBuilderTaskDto) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetPercentage

`func (o *DocumentBuilderTaskDto) GetPercentage() int32`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *DocumentBuilderTaskDto) GetPercentageOk() (*int32, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *DocumentBuilderTaskDto) SetPercentage(v int32)`

SetPercentage sets Percentage field to given value.


### GetIsCompleted

`func (o *DocumentBuilderTaskDto) GetIsCompleted() bool`

GetIsCompleted returns the IsCompleted field if non-nil, zero value otherwise.

### GetIsCompletedOk

`func (o *DocumentBuilderTaskDto) GetIsCompletedOk() (*bool, bool)`

GetIsCompletedOk returns a tuple with the IsCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompleted

`func (o *DocumentBuilderTaskDto) SetIsCompleted(v bool)`

SetIsCompleted sets IsCompleted field to given value.


### GetStatus

`func (o *DocumentBuilderTaskDto) GetStatus() DistributedTaskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DocumentBuilderTaskDto) GetStatusOk() (*DistributedTaskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DocumentBuilderTaskDto) SetStatus(v DistributedTaskStatus)`

SetStatus sets Status field to given value.


### GetResultFileId

`func (o *DocumentBuilderTaskDto) GetResultFileId() interface{}`

GetResultFileId returns the ResultFileId field if non-nil, zero value otherwise.

### GetResultFileIdOk

`func (o *DocumentBuilderTaskDto) GetResultFileIdOk() (*interface{}, bool)`

GetResultFileIdOk returns a tuple with the ResultFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultFileId

`func (o *DocumentBuilderTaskDto) SetResultFileId(v interface{})`

SetResultFileId sets ResultFileId field to given value.


### SetResultFileIdNil

`func (o *DocumentBuilderTaskDto) SetResultFileIdNil(b bool)`

 SetResultFileIdNil sets the value for ResultFileId to be an explicit nil

### UnsetResultFileId
`func (o *DocumentBuilderTaskDto) UnsetResultFileId()`

UnsetResultFileId ensures that no value is present for ResultFileId, not even an explicit nil
### GetResultFileName

`func (o *DocumentBuilderTaskDto) GetResultFileName() string`

GetResultFileName returns the ResultFileName field if non-nil, zero value otherwise.

### GetResultFileNameOk

`func (o *DocumentBuilderTaskDto) GetResultFileNameOk() (*string, bool)`

GetResultFileNameOk returns a tuple with the ResultFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultFileName

`func (o *DocumentBuilderTaskDto) SetResultFileName(v string)`

SetResultFileName sets ResultFileName field to given value.


### SetResultFileNameNil

`func (o *DocumentBuilderTaskDto) SetResultFileNameNil(b bool)`

 SetResultFileNameNil sets the value for ResultFileName to be an explicit nil

### UnsetResultFileName
`func (o *DocumentBuilderTaskDto) UnsetResultFileName()`

UnsetResultFileName ensures that no value is present for ResultFileName, not even an explicit nil
### GetResultFileUrl

`func (o *DocumentBuilderTaskDto) GetResultFileUrl() string`

GetResultFileUrl returns the ResultFileUrl field if non-nil, zero value otherwise.

### GetResultFileUrlOk

`func (o *DocumentBuilderTaskDto) GetResultFileUrlOk() (*string, bool)`

GetResultFileUrlOk returns a tuple with the ResultFileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultFileUrl

`func (o *DocumentBuilderTaskDto) SetResultFileUrl(v string)`

SetResultFileUrl sets ResultFileUrl field to given value.


### SetResultFileUrlNil

`func (o *DocumentBuilderTaskDto) SetResultFileUrlNil(b bool)`

 SetResultFileUrlNil sets the value for ResultFileUrl to be an explicit nil

### UnsetResultFileUrl
`func (o *DocumentBuilderTaskDto) UnsetResultFileUrl()`

UnsetResultFileUrl ensures that no value is present for ResultFileUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


