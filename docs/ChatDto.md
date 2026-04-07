# ChatDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier of the AI chat session. | [optional] 
**Title** | Pointer to **NullableString** | The display title of the chat session. | [optional] 
**CreatedOn** | Pointer to [**ApiDateTime**](ApiDateTime.md) |  | [optional] 
**ModifiedOn** | Pointer to [**ApiDateTime**](ApiDateTime.md) |  | [optional] 
**CreatedBy** | Pointer to [**EmployeeDto**](EmployeeDto.md) |  | [optional] 

## Methods

### NewChatDto

`func NewChatDto() *ChatDto`

NewChatDto instantiates a new ChatDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatDtoWithDefaults

`func NewChatDtoWithDefaults() *ChatDto`

NewChatDtoWithDefaults instantiates a new ChatDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChatDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChatDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChatDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ChatDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *ChatDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ChatDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ChatDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ChatDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ChatDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ChatDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetCreatedOn

`func (o *ChatDto) GetCreatedOn() ApiDateTime`

GetCreatedOn returns the CreatedOn field if non-nil, zero value otherwise.

### GetCreatedOnOk

`func (o *ChatDto) GetCreatedOnOk() (*ApiDateTime, bool)`

GetCreatedOnOk returns a tuple with the CreatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOn

`func (o *ChatDto) SetCreatedOn(v ApiDateTime)`

SetCreatedOn sets CreatedOn field to given value.

### HasCreatedOn

`func (o *ChatDto) HasCreatedOn() bool`

HasCreatedOn returns a boolean if a field has been set.

### GetModifiedOn

`func (o *ChatDto) GetModifiedOn() ApiDateTime`

GetModifiedOn returns the ModifiedOn field if non-nil, zero value otherwise.

### GetModifiedOnOk

`func (o *ChatDto) GetModifiedOnOk() (*ApiDateTime, bool)`

GetModifiedOnOk returns a tuple with the ModifiedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedOn

`func (o *ChatDto) SetModifiedOn(v ApiDateTime)`

SetModifiedOn sets ModifiedOn field to given value.

### HasModifiedOn

`func (o *ChatDto) HasModifiedOn() bool`

HasModifiedOn returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ChatDto) GetCreatedBy() EmployeeDto`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ChatDto) GetCreatedByOk() (*EmployeeDto, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ChatDto) SetCreatedBy(v EmployeeDto)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ChatDto) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


