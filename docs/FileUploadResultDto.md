# FileUploadResultDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** | Specifies if the upload operation is successful or not. | [optional] 
**Data** | Pointer to **interface{}** |  | [optional] 
**Message** | Pointer to **NullableString** | The file upload result message. | [optional] 

## Methods

### NewFileUploadResultDto

`func NewFileUploadResultDto() *FileUploadResultDto`

NewFileUploadResultDto instantiates a new FileUploadResultDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileUploadResultDtoWithDefaults

`func NewFileUploadResultDtoWithDefaults() *FileUploadResultDto`

NewFileUploadResultDtoWithDefaults instantiates a new FileUploadResultDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *FileUploadResultDto) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *FileUploadResultDto) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *FileUploadResultDto) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *FileUploadResultDto) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetData

`func (o *FileUploadResultDto) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FileUploadResultDto) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FileUploadResultDto) SetData(v interface{})`

SetData sets Data field to given value.

### HasData

`func (o *FileUploadResultDto) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *FileUploadResultDto) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *FileUploadResultDto) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetMessage

`func (o *FileUploadResultDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FileUploadResultDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FileUploadResultDto) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *FileUploadResultDto) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *FileUploadResultDto) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *FileUploadResultDto) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


