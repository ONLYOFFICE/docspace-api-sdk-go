# TokenDiagnosticsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The name of the authenticated identity. | [optional] 
**Claims** | Pointer to **[]string** | The claims of the identity, each formatted as type:value. | [optional] 

## Methods

### NewTokenDiagnosticsDto

`func NewTokenDiagnosticsDto() *TokenDiagnosticsDto`

NewTokenDiagnosticsDto instantiates a new TokenDiagnosticsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenDiagnosticsDtoWithDefaults

`func NewTokenDiagnosticsDtoWithDefaults() *TokenDiagnosticsDto`

NewTokenDiagnosticsDtoWithDefaults instantiates a new TokenDiagnosticsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TokenDiagnosticsDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokenDiagnosticsDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokenDiagnosticsDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TokenDiagnosticsDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TokenDiagnosticsDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TokenDiagnosticsDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetClaims

`func (o *TokenDiagnosticsDto) GetClaims() []string`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *TokenDiagnosticsDto) GetClaimsOk() (*[]string, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *TokenDiagnosticsDto) SetClaims(v []string)`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *TokenDiagnosticsDto) HasClaims() bool`

HasClaims returns a boolean if a field has been set.

### SetClaimsNil

`func (o *TokenDiagnosticsDto) SetClaimsNil(b bool)`

 SetClaimsNil sets the value for Claims to be an explicit nil

### UnsetClaims
`func (o *TokenDiagnosticsDto) UnsetClaims()`

UnsetClaims ensures that no value is present for Claims, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


