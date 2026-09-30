# TfaSetupCodeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **NullableString** | The label the authenticator application will list the credential under, which is the caller's own email  address. It identifies the entry to a person, and no application checks it. | [optional] [readonly] 
**ManualEntryKey** | Pointer to **NullableString** | The secret in the base32 form that is typed into an application by hand. It describes the very same  credential as `qrCodeSetupImageUrl`, and repeating the call hands back the same value for the account until  the credential is reset. | [optional] [readonly] 
**QrCodeSetupImageUrl** | Pointer to **NullableString** | The same secret as a scannable image, given as a `data:image/png;base64,` URL that can be rendered  directly - it is not a link to fetch. | [optional] [readonly] 

## Methods

### NewTfaSetupCodeDto

`func NewTfaSetupCodeDto() *TfaSetupCodeDto`

NewTfaSetupCodeDto instantiates a new TfaSetupCodeDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTfaSetupCodeDtoWithDefaults

`func NewTfaSetupCodeDtoWithDefaults() *TfaSetupCodeDto`

NewTfaSetupCodeDtoWithDefaults instantiates a new TfaSetupCodeDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *TfaSetupCodeDto) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *TfaSetupCodeDto) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *TfaSetupCodeDto) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *TfaSetupCodeDto) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *TfaSetupCodeDto) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *TfaSetupCodeDto) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetManualEntryKey

`func (o *TfaSetupCodeDto) GetManualEntryKey() string`

GetManualEntryKey returns the ManualEntryKey field if non-nil, zero value otherwise.

### GetManualEntryKeyOk

`func (o *TfaSetupCodeDto) GetManualEntryKeyOk() (*string, bool)`

GetManualEntryKeyOk returns a tuple with the ManualEntryKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualEntryKey

`func (o *TfaSetupCodeDto) SetManualEntryKey(v string)`

SetManualEntryKey sets ManualEntryKey field to given value.

### HasManualEntryKey

`func (o *TfaSetupCodeDto) HasManualEntryKey() bool`

HasManualEntryKey returns a boolean if a field has been set.

### SetManualEntryKeyNil

`func (o *TfaSetupCodeDto) SetManualEntryKeyNil(b bool)`

 SetManualEntryKeyNil sets the value for ManualEntryKey to be an explicit nil

### UnsetManualEntryKey
`func (o *TfaSetupCodeDto) UnsetManualEntryKey()`

UnsetManualEntryKey ensures that no value is present for ManualEntryKey, not even an explicit nil
### GetQrCodeSetupImageUrl

`func (o *TfaSetupCodeDto) GetQrCodeSetupImageUrl() string`

GetQrCodeSetupImageUrl returns the QrCodeSetupImageUrl field if non-nil, zero value otherwise.

### GetQrCodeSetupImageUrlOk

`func (o *TfaSetupCodeDto) GetQrCodeSetupImageUrlOk() (*string, bool)`

GetQrCodeSetupImageUrlOk returns a tuple with the QrCodeSetupImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrCodeSetupImageUrl

`func (o *TfaSetupCodeDto) SetQrCodeSetupImageUrl(v string)`

SetQrCodeSetupImageUrl sets QrCodeSetupImageUrl field to given value.

### HasQrCodeSetupImageUrl

`func (o *TfaSetupCodeDto) HasQrCodeSetupImageUrl() bool`

HasQrCodeSetupImageUrl returns a boolean if a field has been set.

### SetQrCodeSetupImageUrlNil

`func (o *TfaSetupCodeDto) SetQrCodeSetupImageUrlNil(b bool)`

 SetQrCodeSetupImageUrlNil sets the value for QrCodeSetupImageUrl to be an explicit nil

### UnsetQrCodeSetupImageUrl
`func (o *TfaSetupCodeDto) UnsetQrCodeSetupImageUrl()`

UnsetQrCodeSetupImageUrl ensures that no value is present for QrCodeSetupImageUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


