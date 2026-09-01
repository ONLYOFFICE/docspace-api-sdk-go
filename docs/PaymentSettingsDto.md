# PaymentSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SalesEmail** | **NullableString** | The email address for sales inquiries and support. | 
**FeedbackAndSupportUrl** | Pointer to **NullableString** | The URL for accessing the feedback and support resources. | [optional] 
**BuyUrl** | **NullableString** | The URL for purchasing or upgrading the product. | 
**Standalone** | **bool** | Indicates whether the system is running in standalone mode. | 
**CurrentLicense** | [**CurrentLicenseInfo**](CurrentLicenseInfo.md) | The current license information. | 
**Max** | **int32** | The maximum quota quantity. | 

## Methods

### NewPaymentSettingsDto

`func NewPaymentSettingsDto(salesEmail NullableString, buyUrl NullableString, standalone bool, currentLicense CurrentLicenseInfo, max int32, ) *PaymentSettingsDto`

NewPaymentSettingsDto instantiates a new PaymentSettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentSettingsDtoWithDefaults

`func NewPaymentSettingsDtoWithDefaults() *PaymentSettingsDto`

NewPaymentSettingsDtoWithDefaults instantiates a new PaymentSettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSalesEmail

`func (o *PaymentSettingsDto) GetSalesEmail() string`

GetSalesEmail returns the SalesEmail field if non-nil, zero value otherwise.

### GetSalesEmailOk

`func (o *PaymentSettingsDto) GetSalesEmailOk() (*string, bool)`

GetSalesEmailOk returns a tuple with the SalesEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesEmail

`func (o *PaymentSettingsDto) SetSalesEmail(v string)`

SetSalesEmail sets SalesEmail field to given value.


### SetSalesEmailNil

`func (o *PaymentSettingsDto) SetSalesEmailNil(b bool)`

 SetSalesEmailNil sets the value for SalesEmail to be an explicit nil

### UnsetSalesEmail
`func (o *PaymentSettingsDto) UnsetSalesEmail()`

UnsetSalesEmail ensures that no value is present for SalesEmail, not even an explicit nil
### GetFeedbackAndSupportUrl

`func (o *PaymentSettingsDto) GetFeedbackAndSupportUrl() string`

GetFeedbackAndSupportUrl returns the FeedbackAndSupportUrl field if non-nil, zero value otherwise.

### GetFeedbackAndSupportUrlOk

`func (o *PaymentSettingsDto) GetFeedbackAndSupportUrlOk() (*string, bool)`

GetFeedbackAndSupportUrlOk returns a tuple with the FeedbackAndSupportUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackAndSupportUrl

`func (o *PaymentSettingsDto) SetFeedbackAndSupportUrl(v string)`

SetFeedbackAndSupportUrl sets FeedbackAndSupportUrl field to given value.

### HasFeedbackAndSupportUrl

`func (o *PaymentSettingsDto) HasFeedbackAndSupportUrl() bool`

HasFeedbackAndSupportUrl returns a boolean if a field has been set.

### SetFeedbackAndSupportUrlNil

`func (o *PaymentSettingsDto) SetFeedbackAndSupportUrlNil(b bool)`

 SetFeedbackAndSupportUrlNil sets the value for FeedbackAndSupportUrl to be an explicit nil

### UnsetFeedbackAndSupportUrl
`func (o *PaymentSettingsDto) UnsetFeedbackAndSupportUrl()`

UnsetFeedbackAndSupportUrl ensures that no value is present for FeedbackAndSupportUrl, not even an explicit nil
### GetBuyUrl

`func (o *PaymentSettingsDto) GetBuyUrl() string`

GetBuyUrl returns the BuyUrl field if non-nil, zero value otherwise.

### GetBuyUrlOk

`func (o *PaymentSettingsDto) GetBuyUrlOk() (*string, bool)`

GetBuyUrlOk returns a tuple with the BuyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyUrl

`func (o *PaymentSettingsDto) SetBuyUrl(v string)`

SetBuyUrl sets BuyUrl field to given value.


### SetBuyUrlNil

`func (o *PaymentSettingsDto) SetBuyUrlNil(b bool)`

 SetBuyUrlNil sets the value for BuyUrl to be an explicit nil

### UnsetBuyUrl
`func (o *PaymentSettingsDto) UnsetBuyUrl()`

UnsetBuyUrl ensures that no value is present for BuyUrl, not even an explicit nil
### GetStandalone

`func (o *PaymentSettingsDto) GetStandalone() bool`

GetStandalone returns the Standalone field if non-nil, zero value otherwise.

### GetStandaloneOk

`func (o *PaymentSettingsDto) GetStandaloneOk() (*bool, bool)`

GetStandaloneOk returns a tuple with the Standalone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandalone

`func (o *PaymentSettingsDto) SetStandalone(v bool)`

SetStandalone sets Standalone field to given value.


### GetCurrentLicense

`func (o *PaymentSettingsDto) GetCurrentLicense() CurrentLicenseInfo`

GetCurrentLicense returns the CurrentLicense field if non-nil, zero value otherwise.

### GetCurrentLicenseOk

`func (o *PaymentSettingsDto) GetCurrentLicenseOk() (*CurrentLicenseInfo, bool)`

GetCurrentLicenseOk returns a tuple with the CurrentLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentLicense

`func (o *PaymentSettingsDto) SetCurrentLicense(v CurrentLicenseInfo)`

SetCurrentLicense sets CurrentLicense field to given value.


### GetMax

`func (o *PaymentSettingsDto) GetMax() int32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *PaymentSettingsDto) GetMaxOk() (*int32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *PaymentSettingsDto) SetMax(v int32)`

SetMax sets Max field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


