# ActionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **NullableString** | The anchor value produced by the editor, opaque to the portal: it names the comment, the mention or the  place the document is scrolled to. | [optional] 
**Type** | Pointer to **NullableString** | What the anchor points at, as the editor names it - a comment thread, for instance. | [optional] 

## Methods

### NewActionConfig

`func NewActionConfig() *ActionConfig`

NewActionConfig instantiates a new ActionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionConfigWithDefaults

`func NewActionConfigWithDefaults() *ActionConfig`

NewActionConfigWithDefaults instantiates a new ActionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ActionConfig) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ActionConfig) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ActionConfig) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *ActionConfig) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ActionConfig) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ActionConfig) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetType

`func (o *ActionConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActionConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActionConfig) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ActionConfig) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ActionConfig) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ActionConfig) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


