# ThirdPartyCheckConversionRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | Pointer to **NullableString** | The file to convert. It is taken from the route of the operation, so a value sent in the body is overwritten. | [optional] 
**Sync** | Pointer to **bool** | How to wait for the result: `true` converts inside the request and answers with the finished result, which is  only sensible for small documents, while `false` queues the conversion and answers with an entry to poll. | [optional] 
**StartConvert** | Pointer to **bool** | Whether the conversion is to be started. It is set by the operation itself, so a value sent in the body is  overwritten. | [optional] 
**Version** | Pointer to **int32** | The version to convert; 0 or less means the current version. | [optional] 
**Password** | Pointer to **NullableString** | The password that opens the source document, for a file that is protected by one; anything else may be left  out. | [optional] 
**OutputType** | Pointer to **NullableString** | The extension of the format to convert into, without the dot, and one the portal can produce from that  source format; left out, the default of the portal for that kind of document is used. | [optional] 
**CreateNewIfExist** | Pointer to **bool** | Where the result goes when the file has been converted before: `true` creates another file beside the source,  `false` replaces the converted file that already exists. | [optional] 

## Methods

### NewThirdPartyCheckConversionRequestDto

`func NewThirdPartyCheckConversionRequestDto() *ThirdPartyCheckConversionRequestDto`

NewThirdPartyCheckConversionRequestDto instantiates a new ThirdPartyCheckConversionRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThirdPartyCheckConversionRequestDtoWithDefaults

`func NewThirdPartyCheckConversionRequestDtoWithDefaults() *ThirdPartyCheckConversionRequestDto`

NewThirdPartyCheckConversionRequestDtoWithDefaults instantiates a new ThirdPartyCheckConversionRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileId

`func (o *ThirdPartyCheckConversionRequestDto) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *ThirdPartyCheckConversionRequestDto) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *ThirdPartyCheckConversionRequestDto) SetFileId(v string)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *ThirdPartyCheckConversionRequestDto) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### SetFileIdNil

`func (o *ThirdPartyCheckConversionRequestDto) SetFileIdNil(b bool)`

 SetFileIdNil sets the value for FileId to be an explicit nil

### UnsetFileId
`func (o *ThirdPartyCheckConversionRequestDto) UnsetFileId()`

UnsetFileId ensures that no value is present for FileId, not even an explicit nil
### GetSync

`func (o *ThirdPartyCheckConversionRequestDto) GetSync() bool`

GetSync returns the Sync field if non-nil, zero value otherwise.

### GetSyncOk

`func (o *ThirdPartyCheckConversionRequestDto) GetSyncOk() (*bool, bool)`

GetSyncOk returns a tuple with the Sync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSync

`func (o *ThirdPartyCheckConversionRequestDto) SetSync(v bool)`

SetSync sets Sync field to given value.

### HasSync

`func (o *ThirdPartyCheckConversionRequestDto) HasSync() bool`

HasSync returns a boolean if a field has been set.

### GetStartConvert

`func (o *ThirdPartyCheckConversionRequestDto) GetStartConvert() bool`

GetStartConvert returns the StartConvert field if non-nil, zero value otherwise.

### GetStartConvertOk

`func (o *ThirdPartyCheckConversionRequestDto) GetStartConvertOk() (*bool, bool)`

GetStartConvertOk returns a tuple with the StartConvert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartConvert

`func (o *ThirdPartyCheckConversionRequestDto) SetStartConvert(v bool)`

SetStartConvert sets StartConvert field to given value.

### HasStartConvert

`func (o *ThirdPartyCheckConversionRequestDto) HasStartConvert() bool`

HasStartConvert returns a boolean if a field has been set.

### GetVersion

`func (o *ThirdPartyCheckConversionRequestDto) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ThirdPartyCheckConversionRequestDto) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ThirdPartyCheckConversionRequestDto) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ThirdPartyCheckConversionRequestDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetPassword

`func (o *ThirdPartyCheckConversionRequestDto) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ThirdPartyCheckConversionRequestDto) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ThirdPartyCheckConversionRequestDto) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ThirdPartyCheckConversionRequestDto) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *ThirdPartyCheckConversionRequestDto) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *ThirdPartyCheckConversionRequestDto) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetOutputType

`func (o *ThirdPartyCheckConversionRequestDto) GetOutputType() string`

GetOutputType returns the OutputType field if non-nil, zero value otherwise.

### GetOutputTypeOk

`func (o *ThirdPartyCheckConversionRequestDto) GetOutputTypeOk() (*string, bool)`

GetOutputTypeOk returns a tuple with the OutputType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputType

`func (o *ThirdPartyCheckConversionRequestDto) SetOutputType(v string)`

SetOutputType sets OutputType field to given value.

### HasOutputType

`func (o *ThirdPartyCheckConversionRequestDto) HasOutputType() bool`

HasOutputType returns a boolean if a field has been set.

### SetOutputTypeNil

`func (o *ThirdPartyCheckConversionRequestDto) SetOutputTypeNil(b bool)`

 SetOutputTypeNil sets the value for OutputType to be an explicit nil

### UnsetOutputType
`func (o *ThirdPartyCheckConversionRequestDto) UnsetOutputType()`

UnsetOutputType ensures that no value is present for OutputType, not even an explicit nil
### GetCreateNewIfExist

`func (o *ThirdPartyCheckConversionRequestDto) GetCreateNewIfExist() bool`

GetCreateNewIfExist returns the CreateNewIfExist field if non-nil, zero value otherwise.

### GetCreateNewIfExistOk

`func (o *ThirdPartyCheckConversionRequestDto) GetCreateNewIfExistOk() (*bool, bool)`

GetCreateNewIfExistOk returns a tuple with the CreateNewIfExist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateNewIfExist

`func (o *ThirdPartyCheckConversionRequestDto) SetCreateNewIfExist(v bool)`

SetCreateNewIfExist sets CreateNewIfExist field to given value.

### HasCreateNewIfExist

`func (o *ThirdPartyCheckConversionRequestDto) HasCreateNewIfExist() bool`

HasCreateNewIfExist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


