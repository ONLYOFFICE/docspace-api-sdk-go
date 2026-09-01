# PasswordSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinLength** | **int32** | The minimum number of characters required for valid passwords. | 
**UpperCase** | **bool** | Specifies whether the password should contain the uppercase letters or not. | 
**Digits** | **bool** | Specifies whether the password should contain the digits or not. | 
**SpecSymbols** | **bool** | Specifies whether the password should contain the special symbols or not. | 
**AllowedCharactersRegexStr** | **NullableString** | The allowed password characters in the regex string format. | 
**DigitsRegexStr** | **NullableString** | The password digits in the regex string format. | 
**UpperCaseRegexStr** | **NullableString** | The password uppercase letters in the regex string format. | 
**SpecSymbolsRegexStr** | **NullableString** | The passaword special symbols in the regex string format. | 

## Methods

### NewPasswordSettingsDto

`func NewPasswordSettingsDto(minLength int32, upperCase bool, digits bool, specSymbols bool, allowedCharactersRegexStr NullableString, digitsRegexStr NullableString, upperCaseRegexStr NullableString, specSymbolsRegexStr NullableString, ) *PasswordSettingsDto`

NewPasswordSettingsDto instantiates a new PasswordSettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordSettingsDtoWithDefaults

`func NewPasswordSettingsDtoWithDefaults() *PasswordSettingsDto`

NewPasswordSettingsDtoWithDefaults instantiates a new PasswordSettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinLength

`func (o *PasswordSettingsDto) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *PasswordSettingsDto) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *PasswordSettingsDto) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.


### GetUpperCase

`func (o *PasswordSettingsDto) GetUpperCase() bool`

GetUpperCase returns the UpperCase field if non-nil, zero value otherwise.

### GetUpperCaseOk

`func (o *PasswordSettingsDto) GetUpperCaseOk() (*bool, bool)`

GetUpperCaseOk returns a tuple with the UpperCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperCase

`func (o *PasswordSettingsDto) SetUpperCase(v bool)`

SetUpperCase sets UpperCase field to given value.


### GetDigits

`func (o *PasswordSettingsDto) GetDigits() bool`

GetDigits returns the Digits field if non-nil, zero value otherwise.

### GetDigitsOk

`func (o *PasswordSettingsDto) GetDigitsOk() (*bool, bool)`

GetDigitsOk returns a tuple with the Digits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigits

`func (o *PasswordSettingsDto) SetDigits(v bool)`

SetDigits sets Digits field to given value.


### GetSpecSymbols

`func (o *PasswordSettingsDto) GetSpecSymbols() bool`

GetSpecSymbols returns the SpecSymbols field if non-nil, zero value otherwise.

### GetSpecSymbolsOk

`func (o *PasswordSettingsDto) GetSpecSymbolsOk() (*bool, bool)`

GetSpecSymbolsOk returns a tuple with the SpecSymbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecSymbols

`func (o *PasswordSettingsDto) SetSpecSymbols(v bool)`

SetSpecSymbols sets SpecSymbols field to given value.


### GetAllowedCharactersRegexStr

`func (o *PasswordSettingsDto) GetAllowedCharactersRegexStr() string`

GetAllowedCharactersRegexStr returns the AllowedCharactersRegexStr field if non-nil, zero value otherwise.

### GetAllowedCharactersRegexStrOk

`func (o *PasswordSettingsDto) GetAllowedCharactersRegexStrOk() (*string, bool)`

GetAllowedCharactersRegexStrOk returns a tuple with the AllowedCharactersRegexStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCharactersRegexStr

`func (o *PasswordSettingsDto) SetAllowedCharactersRegexStr(v string)`

SetAllowedCharactersRegexStr sets AllowedCharactersRegexStr field to given value.


### SetAllowedCharactersRegexStrNil

`func (o *PasswordSettingsDto) SetAllowedCharactersRegexStrNil(b bool)`

 SetAllowedCharactersRegexStrNil sets the value for AllowedCharactersRegexStr to be an explicit nil

### UnsetAllowedCharactersRegexStr
`func (o *PasswordSettingsDto) UnsetAllowedCharactersRegexStr()`

UnsetAllowedCharactersRegexStr ensures that no value is present for AllowedCharactersRegexStr, not even an explicit nil
### GetDigitsRegexStr

`func (o *PasswordSettingsDto) GetDigitsRegexStr() string`

GetDigitsRegexStr returns the DigitsRegexStr field if non-nil, zero value otherwise.

### GetDigitsRegexStrOk

`func (o *PasswordSettingsDto) GetDigitsRegexStrOk() (*string, bool)`

GetDigitsRegexStrOk returns a tuple with the DigitsRegexStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitsRegexStr

`func (o *PasswordSettingsDto) SetDigitsRegexStr(v string)`

SetDigitsRegexStr sets DigitsRegexStr field to given value.


### SetDigitsRegexStrNil

`func (o *PasswordSettingsDto) SetDigitsRegexStrNil(b bool)`

 SetDigitsRegexStrNil sets the value for DigitsRegexStr to be an explicit nil

### UnsetDigitsRegexStr
`func (o *PasswordSettingsDto) UnsetDigitsRegexStr()`

UnsetDigitsRegexStr ensures that no value is present for DigitsRegexStr, not even an explicit nil
### GetUpperCaseRegexStr

`func (o *PasswordSettingsDto) GetUpperCaseRegexStr() string`

GetUpperCaseRegexStr returns the UpperCaseRegexStr field if non-nil, zero value otherwise.

### GetUpperCaseRegexStrOk

`func (o *PasswordSettingsDto) GetUpperCaseRegexStrOk() (*string, bool)`

GetUpperCaseRegexStrOk returns a tuple with the UpperCaseRegexStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperCaseRegexStr

`func (o *PasswordSettingsDto) SetUpperCaseRegexStr(v string)`

SetUpperCaseRegexStr sets UpperCaseRegexStr field to given value.


### SetUpperCaseRegexStrNil

`func (o *PasswordSettingsDto) SetUpperCaseRegexStrNil(b bool)`

 SetUpperCaseRegexStrNil sets the value for UpperCaseRegexStr to be an explicit nil

### UnsetUpperCaseRegexStr
`func (o *PasswordSettingsDto) UnsetUpperCaseRegexStr()`

UnsetUpperCaseRegexStr ensures that no value is present for UpperCaseRegexStr, not even an explicit nil
### GetSpecSymbolsRegexStr

`func (o *PasswordSettingsDto) GetSpecSymbolsRegexStr() string`

GetSpecSymbolsRegexStr returns the SpecSymbolsRegexStr field if non-nil, zero value otherwise.

### GetSpecSymbolsRegexStrOk

`func (o *PasswordSettingsDto) GetSpecSymbolsRegexStrOk() (*string, bool)`

GetSpecSymbolsRegexStrOk returns a tuple with the SpecSymbolsRegexStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecSymbolsRegexStr

`func (o *PasswordSettingsDto) SetSpecSymbolsRegexStr(v string)`

SetSpecSymbolsRegexStr sets SpecSymbolsRegexStr field to given value.


### SetSpecSymbolsRegexStrNil

`func (o *PasswordSettingsDto) SetSpecSymbolsRegexStrNil(b bool)`

 SetSpecSymbolsRegexStrNil sets the value for SpecSymbolsRegexStr to be an explicit nil

### UnsetSpecSymbolsRegexStr
`func (o *PasswordSettingsDto) UnsetSpecSymbolsRegexStr()`

UnsetSpecSymbolsRegexStr ensures that no value is present for SpecSymbolsRegexStr, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


