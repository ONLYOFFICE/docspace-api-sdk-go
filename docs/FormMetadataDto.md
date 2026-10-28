# FormMetadataDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **NullableString** | The form field key. | [optional] 
**Type** | Pointer to **NullableString** | The form field type. | [optional] 
**Format** | Pointer to **NullableString** | The form field format. | [optional] 
**PossibleValues** | Pointer to **[]string** | The list of possible values for the form field. | [optional] 

## Methods

### NewFormMetadataDto

`func NewFormMetadataDto() *FormMetadataDto`

NewFormMetadataDto instantiates a new FormMetadataDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormMetadataDtoWithDefaults

`func NewFormMetadataDtoWithDefaults() *FormMetadataDto`

NewFormMetadataDtoWithDefaults instantiates a new FormMetadataDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *FormMetadataDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FormMetadataDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FormMetadataDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *FormMetadataDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *FormMetadataDto) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *FormMetadataDto) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetType

`func (o *FormMetadataDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormMetadataDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormMetadataDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FormMetadataDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *FormMetadataDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *FormMetadataDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetFormat

`func (o *FormMetadataDto) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *FormMetadataDto) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *FormMetadataDto) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *FormMetadataDto) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *FormMetadataDto) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *FormMetadataDto) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil
### GetPossibleValues

`func (o *FormMetadataDto) GetPossibleValues() []string`

GetPossibleValues returns the PossibleValues field if non-nil, zero value otherwise.

### GetPossibleValuesOk

`func (o *FormMetadataDto) GetPossibleValuesOk() (*[]string, bool)`

GetPossibleValuesOk returns a tuple with the PossibleValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPossibleValues

`func (o *FormMetadataDto) SetPossibleValues(v []string)`

SetPossibleValues sets PossibleValues field to given value.

### HasPossibleValues

`func (o *FormMetadataDto) HasPossibleValues() bool`

HasPossibleValues returns a boolean if a field has been set.

### SetPossibleValuesNil

`func (o *FormMetadataDto) SetPossibleValuesNil(b bool)`

 SetPossibleValuesNil sets the value for PossibleValues to be an explicit nil

### UnsetPossibleValues
`func (o *FormMetadataDto) UnsetPossibleValues()`

UnsetPossibleValues ensures that no value is present for PossibleValues, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


