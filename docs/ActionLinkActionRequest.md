# ActionLinkActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **NullableString** | The anchor value produced by the editor, opaque to the portal: it names the comment, the mention or the  place the document is scrolled to. | [optional] 
**Type** | Pointer to **NullableString** | What the anchor points at, as the editor names it - a comment thread, for instance. | [optional] 

## Methods

### NewActionLinkActionRequest

`func NewActionLinkActionRequest() *ActionLinkActionRequest`

NewActionLinkActionRequest instantiates a new ActionLinkActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionLinkActionRequestWithDefaults

`func NewActionLinkActionRequestWithDefaults() *ActionLinkActionRequest`

NewActionLinkActionRequestWithDefaults instantiates a new ActionLinkActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ActionLinkActionRequest) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ActionLinkActionRequest) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ActionLinkActionRequest) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *ActionLinkActionRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ActionLinkActionRequest) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ActionLinkActionRequest) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetType

`func (o *ActionLinkActionRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActionLinkActionRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActionLinkActionRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ActionLinkActionRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ActionLinkActionRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ActionLinkActionRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


