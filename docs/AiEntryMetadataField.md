# AiEntryMetadataField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | 
**Name** | **string** |  | 
**Type** | **float32** | `MetadataFieldType`: 0 String, 1 Date, 2 Number, 3 SingleChoice, 4 MultiChoice. | 
**Options** | [**[]AiEntryMetadataFieldOptionsInner**](AiEntryMetadataFieldOptionsInner.md) |  | 
**Value** | Pointer to [**AiMetadataValue**](AiMetadataValue.md) |  | [optional] 

## Methods

### NewAiEntryMetadataField

`func NewAiEntryMetadataField(id float32, name string, type_ float32, options []AiEntryMetadataFieldOptionsInner, ) *AiEntryMetadataField`

NewAiEntryMetadataField instantiates a new AiEntryMetadataField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiEntryMetadataFieldWithDefaults

`func NewAiEntryMetadataFieldWithDefaults() *AiEntryMetadataField`

NewAiEntryMetadataFieldWithDefaults instantiates a new AiEntryMetadataField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AiEntryMetadataField) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiEntryMetadataField) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiEntryMetadataField) SetId(v float32)`

SetId sets Id field to given value.


### GetName

`func (o *AiEntryMetadataField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AiEntryMetadataField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AiEntryMetadataField) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *AiEntryMetadataField) GetType() float32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AiEntryMetadataField) GetTypeOk() (*float32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AiEntryMetadataField) SetType(v float32)`

SetType sets Type field to given value.


### GetOptions

`func (o *AiEntryMetadataField) GetOptions() []AiEntryMetadataFieldOptionsInner`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AiEntryMetadataField) GetOptionsOk() (*[]AiEntryMetadataFieldOptionsInner, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AiEntryMetadataField) SetOptions(v []AiEntryMetadataFieldOptionsInner)`

SetOptions sets Options field to given value.


### GetValue

`func (o *AiEntryMetadataField) GetValue() AiMetadataValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AiEntryMetadataField) GetValueOk() (*AiMetadataValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AiEntryMetadataField) SetValue(v AiMetadataValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *AiEntryMetadataField) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


