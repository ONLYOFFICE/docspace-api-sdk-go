# CompanyWhiteLabelSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyName** | **NullableString** | The company name. | 
**Site** | **NullableString** | The company site. | 
**Email** | **NullableString** | The company email address. | 
**Address** | **NullableString** | The company address. | 
**Phone** | **NullableString** | The company phone number. | 
**IsLicensor** | **bool** | Specifies if a company is a licensor or not. | 
**HideAbout** | **bool** | Specifies if the About page is visible or not. | 
**IsDefault** | **bool** | Specifies if these settings are default or not. | 

## Methods

### NewCompanyWhiteLabelSettingsDto

`func NewCompanyWhiteLabelSettingsDto(companyName NullableString, site NullableString, email NullableString, address NullableString, phone NullableString, isLicensor bool, hideAbout bool, isDefault bool, ) *CompanyWhiteLabelSettingsDto`

NewCompanyWhiteLabelSettingsDto instantiates a new CompanyWhiteLabelSettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyWhiteLabelSettingsDtoWithDefaults

`func NewCompanyWhiteLabelSettingsDtoWithDefaults() *CompanyWhiteLabelSettingsDto`

NewCompanyWhiteLabelSettingsDtoWithDefaults instantiates a new CompanyWhiteLabelSettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyName

`func (o *CompanyWhiteLabelSettingsDto) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CompanyWhiteLabelSettingsDto) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CompanyWhiteLabelSettingsDto) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.


### SetCompanyNameNil

`func (o *CompanyWhiteLabelSettingsDto) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *CompanyWhiteLabelSettingsDto) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetSite

`func (o *CompanyWhiteLabelSettingsDto) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *CompanyWhiteLabelSettingsDto) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *CompanyWhiteLabelSettingsDto) SetSite(v string)`

SetSite sets Site field to given value.


### SetSiteNil

`func (o *CompanyWhiteLabelSettingsDto) SetSiteNil(b bool)`

 SetSiteNil sets the value for Site to be an explicit nil

### UnsetSite
`func (o *CompanyWhiteLabelSettingsDto) UnsetSite()`

UnsetSite ensures that no value is present for Site, not even an explicit nil
### GetEmail

`func (o *CompanyWhiteLabelSettingsDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CompanyWhiteLabelSettingsDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CompanyWhiteLabelSettingsDto) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *CompanyWhiteLabelSettingsDto) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *CompanyWhiteLabelSettingsDto) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetAddress

`func (o *CompanyWhiteLabelSettingsDto) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CompanyWhiteLabelSettingsDto) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CompanyWhiteLabelSettingsDto) SetAddress(v string)`

SetAddress sets Address field to given value.


### SetAddressNil

`func (o *CompanyWhiteLabelSettingsDto) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *CompanyWhiteLabelSettingsDto) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetPhone

`func (o *CompanyWhiteLabelSettingsDto) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CompanyWhiteLabelSettingsDto) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CompanyWhiteLabelSettingsDto) SetPhone(v string)`

SetPhone sets Phone field to given value.


### SetPhoneNil

`func (o *CompanyWhiteLabelSettingsDto) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *CompanyWhiteLabelSettingsDto) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetIsLicensor

`func (o *CompanyWhiteLabelSettingsDto) GetIsLicensor() bool`

GetIsLicensor returns the IsLicensor field if non-nil, zero value otherwise.

### GetIsLicensorOk

`func (o *CompanyWhiteLabelSettingsDto) GetIsLicensorOk() (*bool, bool)`

GetIsLicensorOk returns a tuple with the IsLicensor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLicensor

`func (o *CompanyWhiteLabelSettingsDto) SetIsLicensor(v bool)`

SetIsLicensor sets IsLicensor field to given value.


### GetHideAbout

`func (o *CompanyWhiteLabelSettingsDto) GetHideAbout() bool`

GetHideAbout returns the HideAbout field if non-nil, zero value otherwise.

### GetHideAboutOk

`func (o *CompanyWhiteLabelSettingsDto) GetHideAboutOk() (*bool, bool)`

GetHideAboutOk returns a tuple with the HideAbout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideAbout

`func (o *CompanyWhiteLabelSettingsDto) SetHideAbout(v bool)`

SetHideAbout sets HideAbout field to given value.


### GetIsDefault

`func (o *CompanyWhiteLabelSettingsDto) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *CompanyWhiteLabelSettingsDto) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *CompanyWhiteLabelSettingsDto) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


