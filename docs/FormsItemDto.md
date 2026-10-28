# FormsItemDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **NullableString** | The name of the field as it is written in the form; send it back as `formsItemKey` to keep only              the completed copies whose field of that name holds a value.              <example>first_name</example> | [optional] 
**Type** | Pointer to **NullableString** | The kind of value the field holds, a text box or a checkbox for instance; send it back as              `formsItemType` beside the key.              <example>text</example> | [optional] 

## Methods

### NewFormsItemDto

`func NewFormsItemDto() *FormsItemDto`

NewFormsItemDto instantiates a new FormsItemDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormsItemDtoWithDefaults

`func NewFormsItemDtoWithDefaults() *FormsItemDto`

NewFormsItemDtoWithDefaults instantiates a new FormsItemDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *FormsItemDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FormsItemDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FormsItemDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *FormsItemDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *FormsItemDto) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *FormsItemDto) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetType

`func (o *FormsItemDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormsItemDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormsItemDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FormsItemDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *FormsItemDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *FormsItemDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


