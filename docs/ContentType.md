# ContentType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Boundary** | Pointer to **NullableString** |  | [optional] 
**CharSet** | Pointer to **NullableString** |  | [optional] 
**MediaType** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Parameters** | Pointer to **[]interface{}** |  | [optional] [readonly] 

## Methods

### NewContentType

`func NewContentType() *ContentType`

NewContentType instantiates a new ContentType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentTypeWithDefaults

`func NewContentTypeWithDefaults() *ContentType`

NewContentTypeWithDefaults instantiates a new ContentType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoundary

`func (o *ContentType) GetBoundary() string`

GetBoundary returns the Boundary field if non-nil, zero value otherwise.

### GetBoundaryOk

`func (o *ContentType) GetBoundaryOk() (*string, bool)`

GetBoundaryOk returns a tuple with the Boundary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundary

`func (o *ContentType) SetBoundary(v string)`

SetBoundary sets Boundary field to given value.

### HasBoundary

`func (o *ContentType) HasBoundary() bool`

HasBoundary returns a boolean if a field has been set.

### SetBoundaryNil

`func (o *ContentType) SetBoundaryNil(b bool)`

 SetBoundaryNil sets the value for Boundary to be an explicit nil

### UnsetBoundary
`func (o *ContentType) UnsetBoundary()`

UnsetBoundary ensures that no value is present for Boundary, not even an explicit nil
### GetCharSet

`func (o *ContentType) GetCharSet() string`

GetCharSet returns the CharSet field if non-nil, zero value otherwise.

### GetCharSetOk

`func (o *ContentType) GetCharSetOk() (*string, bool)`

GetCharSetOk returns a tuple with the CharSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharSet

`func (o *ContentType) SetCharSet(v string)`

SetCharSet sets CharSet field to given value.

### HasCharSet

`func (o *ContentType) HasCharSet() bool`

HasCharSet returns a boolean if a field has been set.

### SetCharSetNil

`func (o *ContentType) SetCharSetNil(b bool)`

 SetCharSetNil sets the value for CharSet to be an explicit nil

### UnsetCharSet
`func (o *ContentType) UnsetCharSet()`

UnsetCharSet ensures that no value is present for CharSet, not even an explicit nil
### GetMediaType

`func (o *ContentType) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *ContentType) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *ContentType) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *ContentType) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### SetMediaTypeNil

`func (o *ContentType) SetMediaTypeNil(b bool)`

 SetMediaTypeNil sets the value for MediaType to be an explicit nil

### UnsetMediaType
`func (o *ContentType) UnsetMediaType()`

UnsetMediaType ensures that no value is present for MediaType, not even an explicit nil
### GetName

`func (o *ContentType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContentType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContentType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContentType) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ContentType) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ContentType) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetParameters

`func (o *ContentType) GetParameters() []interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ContentType) GetParametersOk() (*[]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ContentType) SetParameters(v []interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *ContentType) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *ContentType) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *ContentType) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


