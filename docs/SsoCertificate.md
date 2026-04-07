# SsoCertificate

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

### NewSsoCertificate

`func NewSsoCertificate() *SsoCertificate`

NewSsoCertificate instantiates a new SsoCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoCertificateWithDefaults

`func NewSsoCertificateWithDefaults() *SsoCertificate`

NewSsoCertificateWithDefaults instantiates a new SsoCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelfSigned

`func (o *SsoCertificate) GetSelfSigned() bool`

GetSelfSigned returns the SelfSigned field if non-nil, zero value otherwise.

### GetSelfSignedOk

`func (o *SsoCertificate) GetSelfSignedOk() (*bool, bool)`

GetSelfSignedOk returns a tuple with the SelfSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfSigned

`func (o *SsoCertificate) SetSelfSigned(v bool)`

SetSelfSigned sets SelfSigned field to given value.

### HasSelfSigned

`func (o *SsoCertificate) HasSelfSigned() bool`

HasSelfSigned returns a boolean if a field has been set.

### GetCrt

`func (o *SsoCertificate) GetCrt() string`

GetCrt returns the Crt field if non-nil, zero value otherwise.

### GetCrtOk

`func (o *SsoCertificate) GetCrtOk() (*string, bool)`

GetCrtOk returns a tuple with the Crt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrt

`func (o *SsoCertificate) SetCrt(v string)`

SetCrt sets Crt field to given value.

### HasCrt

`func (o *SsoCertificate) HasCrt() bool`

HasCrt returns a boolean if a field has been set.

### SetCrtNil

`func (o *SsoCertificate) SetCrtNil(b bool)`

 SetCrtNil sets the value for Crt to be an explicit nil

### UnsetCrt
`func (o *SsoCertificate) UnsetCrt()`

UnsetCrt ensures that no value is present for Crt, not even an explicit nil
### GetKey

`func (o *SsoCertificate) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SsoCertificate) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SsoCertificate) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SsoCertificate) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *SsoCertificate) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *SsoCertificate) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetAction

`func (o *SsoCertificate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SsoCertificate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SsoCertificate) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *SsoCertificate) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *SsoCertificate) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *SsoCertificate) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetDomainName

`func (o *SsoCertificate) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *SsoCertificate) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *SsoCertificate) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *SsoCertificate) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### SetDomainNameNil

`func (o *SsoCertificate) SetDomainNameNil(b bool)`

 SetDomainNameNil sets the value for DomainName to be an explicit nil

### UnsetDomainName
`func (o *SsoCertificate) UnsetDomainName()`

UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil
### GetStartDate

`func (o *SsoCertificate) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *SsoCertificate) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *SsoCertificate) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *SsoCertificate) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetExpiredDate

`func (o *SsoCertificate) GetExpiredDate() time.Time`

GetExpiredDate returns the ExpiredDate field if non-nil, zero value otherwise.

### GetExpiredDateOk

`func (o *SsoCertificate) GetExpiredDateOk() (*time.Time, bool)`

GetExpiredDateOk returns a tuple with the ExpiredDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredDate

`func (o *SsoCertificate) SetExpiredDate(v time.Time)`

SetExpiredDate sets ExpiredDate field to given value.

### HasExpiredDate

`func (o *SsoCertificate) HasExpiredDate() bool`

HasExpiredDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


