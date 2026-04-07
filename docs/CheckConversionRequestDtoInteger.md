# CheckConversionRequestDtoInteger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | Pointer to **int32** | The file ID to check conversion proccess. | [optional] 
**Sync** | Pointer to **bool** | Specifies if the conversion process is synchronous or not. | [optional] 
**StartConvert** | Pointer to **bool** | Specifies whether to start a conversion process or not. | [optional] 
**Version** | Pointer to **int32** | The file version that is converted. | [optional] 
**Password** | Pointer to **NullableString** | The password of the converted file. | [optional] 
**OutputType** | Pointer to **NullableString** | The conversion output type. | [optional] 
**CreateNewIfExist** | Pointer to **bool** | Specifies whether to create a new file if it exists or not. | [optional] 

## Methods

### NewCheckConversionRequestDtoInteger

`func NewCheckConversionRequestDtoInteger() *CheckConversionRequestDtoInteger`

NewCheckConversionRequestDtoInteger instantiates a new CheckConversionRequestDtoInteger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckConversionRequestDtoIntegerWithDefaults

`func NewCheckConversionRequestDtoIntegerWithDefaults() *CheckConversionRequestDtoInteger`

NewCheckConversionRequestDtoIntegerWithDefaults instantiates a new CheckConversionRequestDtoInteger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileId

`func (o *CheckConversionRequestDtoInteger) GetFileId() int32`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *CheckConversionRequestDtoInteger) GetFileIdOk() (*int32, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *CheckConversionRequestDtoInteger) SetFileId(v int32)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *CheckConversionRequestDtoInteger) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### GetSync

`func (o *CheckConversionRequestDtoInteger) GetSync() bool`

GetSync returns the Sync field if non-nil, zero value otherwise.

### GetSyncOk

`func (o *CheckConversionRequestDtoInteger) GetSyncOk() (*bool, bool)`

GetSyncOk returns a tuple with the Sync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSync

`func (o *CheckConversionRequestDtoInteger) SetSync(v bool)`

SetSync sets Sync field to given value.

### HasSync

`func (o *CheckConversionRequestDtoInteger) HasSync() bool`

HasSync returns a boolean if a field has been set.

### GetStartConvert

`func (o *CheckConversionRequestDtoInteger) GetStartConvert() bool`

GetStartConvert returns the StartConvert field if non-nil, zero value otherwise.

### GetStartConvertOk

`func (o *CheckConversionRequestDtoInteger) GetStartConvertOk() (*bool, bool)`

GetStartConvertOk returns a tuple with the StartConvert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartConvert

`func (o *CheckConversionRequestDtoInteger) SetStartConvert(v bool)`

SetStartConvert sets StartConvert field to given value.

### HasStartConvert

`func (o *CheckConversionRequestDtoInteger) HasStartConvert() bool`

HasStartConvert returns a boolean if a field has been set.

### GetVersion

`func (o *CheckConversionRequestDtoInteger) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CheckConversionRequestDtoInteger) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CheckConversionRequestDtoInteger) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CheckConversionRequestDtoInteger) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetPassword

`func (o *CheckConversionRequestDtoInteger) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CheckConversionRequestDtoInteger) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CheckConversionRequestDtoInteger) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CheckConversionRequestDtoInteger) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *CheckConversionRequestDtoInteger) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *CheckConversionRequestDtoInteger) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetOutputType

`func (o *CheckConversionRequestDtoInteger) GetOutputType() string`

GetOutputType returns the OutputType field if non-nil, zero value otherwise.

### GetOutputTypeOk

`func (o *CheckConversionRequestDtoInteger) GetOutputTypeOk() (*string, bool)`

GetOutputTypeOk returns a tuple with the OutputType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputType

`func (o *CheckConversionRequestDtoInteger) SetOutputType(v string)`

SetOutputType sets OutputType field to given value.

### HasOutputType

`func (o *CheckConversionRequestDtoInteger) HasOutputType() bool`

HasOutputType returns a boolean if a field has been set.

### SetOutputTypeNil

`func (o *CheckConversionRequestDtoInteger) SetOutputTypeNil(b bool)`

 SetOutputTypeNil sets the value for OutputType to be an explicit nil

### UnsetOutputType
`func (o *CheckConversionRequestDtoInteger) UnsetOutputType()`

UnsetOutputType ensures that no value is present for OutputType, not even an explicit nil
### GetCreateNewIfExist

`func (o *CheckConversionRequestDtoInteger) GetCreateNewIfExist() bool`

GetCreateNewIfExist returns the CreateNewIfExist field if non-nil, zero value otherwise.

### GetCreateNewIfExistOk

`func (o *CheckConversionRequestDtoInteger) GetCreateNewIfExistOk() (*bool, bool)`

GetCreateNewIfExistOk returns a tuple with the CreateNewIfExist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateNewIfExist

`func (o *CheckConversionRequestDtoInteger) SetCreateNewIfExist(v bool)`

SetCreateNewIfExist sets CreateNewIfExist field to given value.

### HasCreateNewIfExist

`func (o *CheckConversionRequestDtoInteger) HasCreateNewIfExist() bool`

HasCreateNewIfExist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


