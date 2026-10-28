# SsoCertificateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SelfSigned** | Pointer to **bool** | Specifies if a certificate is self-signed or not. | [optional] 
**Crt** | Pointer to **NullableString** | The CRT certificate file. | [optional] 
**Key** | Pointer to **NullableString** | The certificate key. | [optional] 
**Action** | Pointer to **NullableString** | The certificate action. | [optional] 
**DomainName** | Pointer to **NullableString** | The certificate domain name. | [optional] 
**StartDate** | Pointer to **time.Time** | The certificate start date. | [optional] 
**ExpiredDate** | Pointer to **time.Time** | The certificate expiration date. | [optional] 

## Methods

### NewSsoCertificateDto

`func NewSsoCertificateDto() *SsoCertificateDto`

NewSsoCertificateDto instantiates a new SsoCertificateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoCertificateDtoWithDefaults

`func NewSsoCertificateDtoWithDefaults() *SsoCertificateDto`

NewSsoCertificateDtoWithDefaults instantiates a new SsoCertificateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelfSigned

`func (o *SsoCertificateDto) GetSelfSigned() bool`

GetSelfSigned returns the SelfSigned field if non-nil, zero value otherwise.

### GetSelfSignedOk

`func (o *SsoCertificateDto) GetSelfSignedOk() (*bool, bool)`

GetSelfSignedOk returns a tuple with the SelfSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfSigned

`func (o *SsoCertificateDto) SetSelfSigned(v bool)`

SetSelfSigned sets SelfSigned field to given value.

### HasSelfSigned

`func (o *SsoCertificateDto) HasSelfSigned() bool`

HasSelfSigned returns a boolean if a field has been set.

### GetCrt

`func (o *SsoCertificateDto) GetCrt() string`

GetCrt returns the Crt field if non-nil, zero value otherwise.

### GetCrtOk

`func (o *SsoCertificateDto) GetCrtOk() (*string, bool)`

GetCrtOk returns a tuple with the Crt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrt

`func (o *SsoCertificateDto) SetCrt(v string)`

SetCrt sets Crt field to given value.

### HasCrt

`func (o *SsoCertificateDto) HasCrt() bool`

HasCrt returns a boolean if a field has been set.

### SetCrtNil

`func (o *SsoCertificateDto) SetCrtNil(b bool)`

 SetCrtNil sets the value for Crt to be an explicit nil

### UnsetCrt
`func (o *SsoCertificateDto) UnsetCrt()`

UnsetCrt ensures that no value is present for Crt, not even an explicit nil
### GetKey

`func (o *SsoCertificateDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SsoCertificateDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SsoCertificateDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SsoCertificateDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *SsoCertificateDto) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *SsoCertificateDto) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetAction

`func (o *SsoCertificateDto) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SsoCertificateDto) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SsoCertificateDto) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *SsoCertificateDto) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *SsoCertificateDto) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *SsoCertificateDto) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetDomainName

`func (o *SsoCertificateDto) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *SsoCertificateDto) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *SsoCertificateDto) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *SsoCertificateDto) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### SetDomainNameNil

`func (o *SsoCertificateDto) SetDomainNameNil(b bool)`

 SetDomainNameNil sets the value for DomainName to be an explicit nil

### UnsetDomainName
`func (o *SsoCertificateDto) UnsetDomainName()`

UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil
### GetStartDate

`func (o *SsoCertificateDto) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *SsoCertificateDto) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *SsoCertificateDto) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *SsoCertificateDto) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetExpiredDate

`func (o *SsoCertificateDto) GetExpiredDate() time.Time`

GetExpiredDate returns the ExpiredDate field if non-nil, zero value otherwise.

### GetExpiredDateOk

`func (o *SsoCertificateDto) GetExpiredDateOk() (*time.Time, bool)`

GetExpiredDateOk returns a tuple with the ExpiredDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredDate

`func (o *SsoCertificateDto) SetExpiredDate(v time.Time)`

SetExpiredDate sets ExpiredDate field to given value.

### HasExpiredDate

`func (o *SsoCertificateDto) HasExpiredDate() bool`

HasExpiredDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


