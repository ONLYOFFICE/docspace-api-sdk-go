# UploadResultDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** | Specifies if the upload operation is successful or not. | [optional] 
**Data** | Pointer to **interface{}** |  | [optional] 
**Message** | Pointer to **NullableString** | The message sent after the successful upload operation. | [optional] 

## Methods

### NewUploadResultDto

`func NewUploadResultDto() *UploadResultDto`

NewUploadResultDto instantiates a new UploadResultDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadResultDtoWithDefaults

`func NewUploadResultDtoWithDefaults() *UploadResultDto`

NewUploadResultDtoWithDefaults instantiates a new UploadResultDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *UploadResultDto) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UploadResultDto) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UploadResultDto) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UploadResultDto) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetData

`func (o *UploadResultDto) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UploadResultDto) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UploadResultDto) SetData(v interface{})`

SetData sets Data field to given value.

### HasData

`func (o *UploadResultDto) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *UploadResultDto) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *UploadResultDto) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetMessage

`func (o *UploadResultDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UploadResultDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UploadResultDto) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UploadResultDto) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *UploadResultDto) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *UploadResultDto) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


