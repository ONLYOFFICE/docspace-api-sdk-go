# MetadataOperationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The operation ID. | [optional] 
**Progress** | Pointer to **float64** | The operation progress percentage. | [optional] 
**IsCompleted** | Pointer to **bool** | Specifies if the operation is completed. | [optional] 
**Error** | Pointer to **NullableString** | The operation error message. | [optional] 
**TemplateIds** | Pointer to **[]int32** | The IDs of the metadata templates the operation propagates to the subtree. Empty when there is no operation to  report, so the answer describes nothing in progress. | [optional] 

## Methods

### NewMetadataOperationDto

`func NewMetadataOperationDto() *MetadataOperationDto`

NewMetadataOperationDto instantiates a new MetadataOperationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataOperationDtoWithDefaults

`func NewMetadataOperationDtoWithDefaults() *MetadataOperationDto`

NewMetadataOperationDtoWithDefaults instantiates a new MetadataOperationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MetadataOperationDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetadataOperationDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetadataOperationDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MetadataOperationDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *MetadataOperationDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MetadataOperationDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetProgress

`func (o *MetadataOperationDto) GetProgress() float64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *MetadataOperationDto) GetProgressOk() (*float64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *MetadataOperationDto) SetProgress(v float64)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *MetadataOperationDto) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetIsCompleted

`func (o *MetadataOperationDto) GetIsCompleted() bool`

GetIsCompleted returns the IsCompleted field if non-nil, zero value otherwise.

### GetIsCompletedOk

`func (o *MetadataOperationDto) GetIsCompletedOk() (*bool, bool)`

GetIsCompletedOk returns a tuple with the IsCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompleted

`func (o *MetadataOperationDto) SetIsCompleted(v bool)`

SetIsCompleted sets IsCompleted field to given value.

### HasIsCompleted

`func (o *MetadataOperationDto) HasIsCompleted() bool`

HasIsCompleted returns a boolean if a field has been set.

### GetError

`func (o *MetadataOperationDto) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MetadataOperationDto) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MetadataOperationDto) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *MetadataOperationDto) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *MetadataOperationDto) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *MetadataOperationDto) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetTemplateIds

`func (o *MetadataOperationDto) GetTemplateIds() []int32`

GetTemplateIds returns the TemplateIds field if non-nil, zero value otherwise.

### GetTemplateIdsOk

`func (o *MetadataOperationDto) GetTemplateIdsOk() (*[]int32, bool)`

GetTemplateIdsOk returns a tuple with the TemplateIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateIds

`func (o *MetadataOperationDto) SetTemplateIds(v []int32)`

SetTemplateIds sets TemplateIds field to given value.

### HasTemplateIds

`func (o *MetadataOperationDto) HasTemplateIds() bool`

HasTemplateIds returns a boolean if a field has been set.

### SetTemplateIdsNil

`func (o *MetadataOperationDto) SetTemplateIdsNil(b bool)`

 SetTemplateIdsNil sets the value for TemplateIds to be an explicit nil

### UnsetTemplateIds
`func (o *MetadataOperationDto) UnsetTemplateIds()`

UnsetTemplateIds ensures that no value is present for TemplateIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


