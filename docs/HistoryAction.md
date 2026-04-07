# HistoryAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to [**MessageAction**](MessageAction.md) |  | [optional] 
**Key** | Pointer to **NullableString** | The action performed on the file. | [optional] 

## Methods

### NewHistoryAction

`func NewHistoryAction() *HistoryAction`

NewHistoryAction instantiates a new HistoryAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryActionWithDefaults

`func NewHistoryActionWithDefaults() *HistoryAction`

NewHistoryActionWithDefaults instantiates a new HistoryAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HistoryAction) GetId() MessageAction`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HistoryAction) GetIdOk() (*MessageAction, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HistoryAction) SetId(v MessageAction)`

SetId sets Id field to given value.

### HasId

`func (o *HistoryAction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *HistoryAction) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *HistoryAction) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *HistoryAction) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *HistoryAction) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *HistoryAction) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *HistoryAction) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


