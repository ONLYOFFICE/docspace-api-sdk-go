# MessageDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The unique identifier of the message. | [optional] 
**Role** | Pointer to [**Role**](Role.md) |  | [optional] 
**Contents** | Pointer to [**[]MessageContentDto**](MessageContentDto.md) | The ordered collection of content blocks that make up the message body (text, tool calls, or attachments). | [optional] 
**CreatedOn** | Pointer to [**ApiDateTime**](ApiDateTime.md) |  | [optional] 

## Methods

### NewMessageDto

`func NewMessageDto() *MessageDto`

NewMessageDto instantiates a new MessageDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageDtoWithDefaults

`func NewMessageDtoWithDefaults() *MessageDto`

NewMessageDtoWithDefaults instantiates a new MessageDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MessageDto) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessageDto) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessageDto) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *MessageDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRole

`func (o *MessageDto) GetRole() Role`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *MessageDto) GetRoleOk() (*Role, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *MessageDto) SetRole(v Role)`

SetRole sets Role field to given value.

### HasRole

`func (o *MessageDto) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetContents

`func (o *MessageDto) GetContents() []MessageContentDto`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *MessageDto) GetContentsOk() (*[]MessageContentDto, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *MessageDto) SetContents(v []MessageContentDto)`

SetContents sets Contents field to given value.

### HasContents

`func (o *MessageDto) HasContents() bool`

HasContents returns a boolean if a field has been set.

### SetContentsNil

`func (o *MessageDto) SetContentsNil(b bool)`

 SetContentsNil sets the value for Contents to be an explicit nil

### UnsetContents
`func (o *MessageDto) UnsetContents()`

UnsetContents ensures that no value is present for Contents, not even an explicit nil
### GetCreatedOn

`func (o *MessageDto) GetCreatedOn() ApiDateTime`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *MessageDto) GetCreatedOnOk() (*ApiDateTime, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *MessageDto) SetCreatedOn(v ApiDateTime)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *MessageDto) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


