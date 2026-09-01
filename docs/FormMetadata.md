# FormMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **NullableString** | The form field key. | [optional] 
**Type** | Pointer to **NullableString** | The form field type. | [optional] 
**Format** | Pointer to **NullableString** | The form field format. | [optional] 
**PossibleValues** | Pointer to **[]string** | The list of possible values for the form field. | [optional] 

## Methods

### NewFormMetadata

`func NewFormMetadata() *FormMetadata`

NewFormMetadata instantiates a new FormMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormMetadataWithDefaults

`func NewFormMetadataWithDefaults() *FormMetadata`

NewFormMetadataWithDefaults instantiates a new FormMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *FormMetadata) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FormMetadata) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FormMetadata) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *FormMetadata) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *FormMetadata) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *FormMetadata) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetType

`func (o *FormMetadata) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormMetadata) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormMetadata) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FormMetadata) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *FormMetadata) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *FormMetadata) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetFormat

`func (o *FormMetadata) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *FormMetadata) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *FormMetadata) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *FormMetadata) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *FormMetadata) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *FormMetadata) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil
### GetPossibleValues

`func (o *FormMetadata) GetPossibleValues() []string`

GetPossibleValues returns the PossibleValues field if non-nil, zero value otherwise.

### GetPossibleValuesOk

`func (o *FormMetadata) GetPossibleValuesOk() (*[]string, bool)`

GetPossibleValuesOk returns a tuple with the PossibleValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPossibleValues

`func (o *FormMetadata) SetPossibleValues(v []string)`

SetPossibleValues sets PossibleValues field to given value.

### HasPossibleValues

`func (o *FormMetadata) HasPossibleValues() bool`

HasPossibleValues returns a boolean if a field has been set.

### SetPossibleValuesNil

`func (o *FormMetadata) SetPossibleValuesNil(b bool)`

 SetPossibleValuesNil sets the value for PossibleValues to be an explicit nil

### UnsetPossibleValues
`func (o *FormMetadata) UnsetPossibleValues()`

UnsetPossibleValues ensures that no value is present for PossibleValues, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


