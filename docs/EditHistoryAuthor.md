# EditHistoryAuthor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** | The author ID. | 
**Name** | Pointer to **NullableString** | The author name. | [optional] 

## Methods

### NewEditHistoryAuthor

`func NewEditHistoryAuthor(id NullableString, ) *EditHistoryAuthor`

NewEditHistoryAuthor instantiates a new EditHistoryAuthor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditHistoryAuthorWithDefaults

`func NewEditHistoryAuthorWithDefaults() *EditHistoryAuthor`

NewEditHistoryAuthorWithDefaults instantiates a new EditHistoryAuthor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EditHistoryAuthor) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EditHistoryAuthor) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EditHistoryAuthor) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *EditHistoryAuthor) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *EditHistoryAuthor) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *EditHistoryAuthor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditHistoryAuthor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditHistoryAuthor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EditHistoryAuthor) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *EditHistoryAuthor) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *EditHistoryAuthor) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


