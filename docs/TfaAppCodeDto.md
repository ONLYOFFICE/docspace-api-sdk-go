# TfaAppCodeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsUsed** | Pointer to **bool** | Whether the code has already been spent. A spent code is kept in the list but is no longer accepted, so  count the entries where this is `false` to know how many fallbacks remain. | [optional] 
**Code** | Pointer to **NullableString** | The code itself, in the form it is typed at sign-in - six characters with the default configuration. It is  stored encrypted and decrypted for this answer, so this is the one place a caller can read it. | [optional] 

## Methods

### NewTfaAppCodeDto

`func NewTfaAppCodeDto() *TfaAppCodeDto`

NewTfaAppCodeDto instantiates a new TfaAppCodeDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTfaAppCodeDtoWithDefaults

`func NewTfaAppCodeDtoWithDefaults() *TfaAppCodeDto`

NewTfaAppCodeDtoWithDefaults instantiates a new TfaAppCodeDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsUsed

`func (o *TfaAppCodeDto) GetIsUsed() bool`

GetIsUsed returns the IsUsed field if non-nil, zero value otherwise.

### GetIsUsedOk

`func (o *TfaAppCodeDto) GetIsUsedOk() (*bool, bool)`

GetIsUsedOk returns a tuple with the IsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUsed

`func (o *TfaAppCodeDto) SetIsUsed(v bool)`

SetIsUsed sets IsUsed field to given value.

### HasIsUsed

`func (o *TfaAppCodeDto) HasIsUsed() bool`

HasIsUsed returns a boolean if a field has been set.

### GetCode

`func (o *TfaAppCodeDto) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TfaAppCodeDto) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TfaAppCodeDto) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *TfaAppCodeDto) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *TfaAppCodeDto) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *TfaAppCodeDto) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


