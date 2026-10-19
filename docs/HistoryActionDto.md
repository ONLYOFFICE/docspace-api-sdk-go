# HistoryActionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to [**MessageAction**](MessageAction.md) | The action performed on the file. | [optional] 
**Key** | Pointer to **NullableString** | The same action as a camel-case key, for a client to look up its own wording by. | [optional] 

## Methods

### NewHistoryActionDto

`func NewHistoryActionDto() *HistoryActionDto`

NewHistoryActionDto instantiates a new HistoryActionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryActionDtoWithDefaults

`func NewHistoryActionDtoWithDefaults() *HistoryActionDto`

NewHistoryActionDtoWithDefaults instantiates a new HistoryActionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HistoryActionDto) GetId() MessageAction`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HistoryActionDto) GetIdOk() (*MessageAction, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HistoryActionDto) SetId(v MessageAction)`

SetId sets Id field to given value.

### HasId

`func (o *HistoryActionDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *HistoryActionDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *HistoryActionDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *HistoryActionDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *HistoryActionDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *HistoryActionDto) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *HistoryActionDto) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


