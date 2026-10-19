# PasswordSettingsRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinLength** | **int32** | The shortest password the portal will accept. It has to sit between the floor the installation is configured  with, 8 characters unless it was changed, and the ceiling of 30; a value outside that is refused with 400. | 
**UpperCase** | Pointer to **bool** | Whether a password must contain at least one uppercase letter. There is no partial update on this body, so  leaving the flag out stores it as `false` and drops the requirement. | [optional] 
**Digits** | Pointer to **bool** | Whether a password must contain at least one digit. Leaving the flag out stores it as `false` and drops the  requirement. | [optional] 
**SpecSymbols** | Pointer to **bool** | Whether a password must contain at least one special symbol. Leaving the flag out stores it as `false` and  drops the requirement. | [optional] 

## Methods

### NewPasswordSettingsRequestDto

`func NewPasswordSettingsRequestDto(minLength int32, ) *PasswordSettingsRequestDto`

NewPasswordSettingsRequestDto instantiates a new PasswordSettingsRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordSettingsRequestDtoWithDefaults

`func NewPasswordSettingsRequestDtoWithDefaults() *PasswordSettingsRequestDto`

NewPasswordSettingsRequestDtoWithDefaults instantiates a new PasswordSettingsRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinLength

`func (o *PasswordSettingsRequestDto) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *PasswordSettingsRequestDto) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *PasswordSettingsRequestDto) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.


### GetUpperCase

`func (o *PasswordSettingsRequestDto) GetUpperCase() bool`

GetUpperCase returns the UpperCase field if non-nil, zero value otherwise.

### GetUpperCaseOk

`func (o *PasswordSettingsRequestDto) GetUpperCaseOk() (*bool, bool)`

GetUpperCaseOk returns a tuple with the UpperCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperCase

`func (o *PasswordSettingsRequestDto) SetUpperCase(v bool)`

SetUpperCase sets UpperCase field to given value.

### HasUpperCase

`func (o *PasswordSettingsRequestDto) HasUpperCase() bool`

HasUpperCase returns a boolean if a field has been set.

### GetDigits

`func (o *PasswordSettingsRequestDto) GetDigits() bool`

GetDigits returns the Digits field if non-nil, zero value otherwise.

### GetDigitsOk

`func (o *PasswordSettingsRequestDto) GetDigitsOk() (*bool, bool)`

GetDigitsOk returns a tuple with the Digits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigits

`func (o *PasswordSettingsRequestDto) SetDigits(v bool)`

SetDigits sets Digits field to given value.

### HasDigits

`func (o *PasswordSettingsRequestDto) HasDigits() bool`

HasDigits returns a boolean if a field has been set.

### GetSpecSymbols

`func (o *PasswordSettingsRequestDto) GetSpecSymbols() bool`

GetSpecSymbols returns the SpecSymbols field if non-nil, zero value otherwise.

### GetSpecSymbolsOk

`func (o *PasswordSettingsRequestDto) GetSpecSymbolsOk() (*bool, bool)`

GetSpecSymbolsOk returns a tuple with the SpecSymbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecSymbols

`func (o *PasswordSettingsRequestDto) SetSpecSymbols(v bool)`

SetSpecSymbols sets SpecSymbols field to given value.

### HasSpecSymbols

`func (o *PasswordSettingsRequestDto) HasSpecSymbols() bool`

HasSpecSymbols returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


