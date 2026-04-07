# ContentDisposition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DispositionType** | Pointer to **NullableString** |  | [optional] 
**Parameters** | Pointer to **[]interface{}** |  | [optional] [readonly] 
**FileName** | Pointer to **NullableString** |  | [optional] 
**CreationDate** | Pointer to **time.Time** |  | [optional] 
**ModificationDate** | Pointer to **time.Time** |  | [optional] 
**Inline** | Pointer to **bool** |  | [optional] 
**ReadDate** | Pointer to **time.Time** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 

## Methods

### NewContentDisposition

`func NewContentDisposition() *ContentDisposition`

NewContentDisposition instantiates a new ContentDisposition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentDispositionWithDefaults

`func NewContentDispositionWithDefaults() *ContentDisposition`

NewContentDispositionWithDefaults instantiates a new ContentDisposition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDispositionType

`func (o *ContentDisposition) GetDispositionType() string`

GetDispositionType returns the DispositionType field if non-nil, zero value otherwise.

### GetDispositionTypeOk

`func (o *ContentDisposition) GetDispositionTypeOk() (*string, bool)`

GetDispositionTypeOk returns a tuple with the DispositionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispositionType

`func (o *ContentDisposition) SetDispositionType(v string)`

SetDispositionType sets DispositionType field to given value.

### HasDispositionType

`func (o *ContentDisposition) HasDispositionType() bool`

HasDispositionType returns a boolean if a field has been set.

### SetDispositionTypeNil

`func (o *ContentDisposition) SetDispositionTypeNil(b bool)`

 SetDispositionTypeNil sets the value for DispositionType to be an explicit nil

### UnsetDispositionType
`func (o *ContentDisposition) UnsetDispositionType()`

UnsetDispositionType ensures that no value is present for DispositionType, not even an explicit nil
### GetParameters

`func (o *ContentDisposition) GetParameters() []interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ContentDisposition) GetParametersOk() (*[]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ContentDisposition) SetParameters(v []interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *ContentDisposition) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *ContentDisposition) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *ContentDisposition) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil
### GetFileName

`func (o *ContentDisposition) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *ContentDisposition) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *ContentDisposition) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *ContentDisposition) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *ContentDisposition) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *ContentDisposition) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetCreationDate

`func (o *ContentDisposition) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *ContentDisposition) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *ContentDisposition) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *ContentDisposition) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetModificationDate

`func (o *ContentDisposition) GetModificationDate() time.Time`

GetModificationDate returns the ModificationDate field if non-nil, zero value otherwise.

### GetModificationDateOk

`func (o *ContentDisposition) GetModificationDateOk() (*time.Time, bool)`

GetModificationDateOk returns a tuple with the ModificationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationDate

`func (o *ContentDisposition) SetModificationDate(v time.Time)`

SetModificationDate sets ModificationDate field to given value.

### HasModificationDate

`func (o *ContentDisposition) HasModificationDate() bool`

HasModificationDate returns a boolean if a field has been set.

### GetInline

`func (o *ContentDisposition) GetInline() bool`

GetInline returns the Inline field if non-nil, zero value otherwise.

### GetInlineOk

`func (o *ContentDisposition) GetInlineOk() (*bool, bool)`

GetInlineOk returns a tuple with the Inline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInline

`func (o *ContentDisposition) SetInline(v bool)`

SetInline sets Inline field to given value.

### HasInline

`func (o *ContentDisposition) HasInline() bool`

HasInline returns a boolean if a field has been set.

### GetReadDate

`func (o *ContentDisposition) GetReadDate() time.Time`

GetReadDate returns the ReadDate field if non-nil, zero value otherwise.

### GetReadDateOk

`func (o *ContentDisposition) GetReadDateOk() (*time.Time, bool)`

GetReadDateOk returns a tuple with the ReadDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadDate

`func (o *ContentDisposition) SetReadDate(v time.Time)`

SetReadDate sets ReadDate field to given value.

### HasReadDate

`func (o *ContentDisposition) HasReadDate() bool`

HasReadDate returns a boolean if a field has been set.

### GetSize

`func (o *ContentDisposition) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ContentDisposition) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ContentDisposition) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *ContentDisposition) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


