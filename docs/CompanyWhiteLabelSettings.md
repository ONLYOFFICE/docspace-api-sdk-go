# CompanyWhiteLabelSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyName** | Pointer to **NullableString** | The company name. | [optional] 
**Site** | Pointer to **NullableString** | The company site. | [optional] 
**Email** | Pointer to **NullableString** | The company email address. | [optional] 
**Address** | Pointer to **NullableString** | The company address. | [optional] 
**Phone** | Pointer to **NullableString** | The company phone number. | [optional] 
**IsLicensor** | Pointer to **bool** | Specifies if a company is a licensor or not. | [optional] 
**HideAbout** | Pointer to **bool** | Specifies if the About page is visible or not | [optional] 
**LastModified** | Pointer to **time.Time** | The timestamp indicating when the settings were last modified. | [optional] 

## Methods

### NewCompanyWhiteLabelSettings

`func NewCompanyWhiteLabelSettings() *CompanyWhiteLabelSettings`

NewCompanyWhiteLabelSettings instantiates a new CompanyWhiteLabelSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyWhiteLabelSettingsWithDefaults

`func NewCompanyWhiteLabelSettingsWithDefaults() *CompanyWhiteLabelSettings`

NewCompanyWhiteLabelSettingsWithDefaults instantiates a new CompanyWhiteLabelSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyName

`func (o *CompanyWhiteLabelSettings) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CompanyWhiteLabelSettings) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CompanyWhiteLabelSettings) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *CompanyWhiteLabelSettings) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *CompanyWhiteLabelSettings) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *CompanyWhiteLabelSettings) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetSite

`func (o *CompanyWhiteLabelSettings) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *CompanyWhiteLabelSettings) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *CompanyWhiteLabelSettings) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *CompanyWhiteLabelSettings) HasSite() bool`

HasSite returns a boolean if a field has been set.

### SetSiteNil

`func (o *CompanyWhiteLabelSettings) SetSiteNil(b bool)`

 SetSiteNil sets the value for Site to be an explicit nil

### UnsetSite
`func (o *CompanyWhiteLabelSettings) UnsetSite()`

UnsetSite ensures that no value is present for Site, not even an explicit nil
### GetEmail

`func (o *CompanyWhiteLabelSettings) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CompanyWhiteLabelSettings) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CompanyWhiteLabelSettings) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CompanyWhiteLabelSettings) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *CompanyWhiteLabelSettings) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *CompanyWhiteLabelSettings) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetAddress

`func (o *CompanyWhiteLabelSettings) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *CompanyWhiteLabelSettings) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *CompanyWhiteLabelSettings) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *CompanyWhiteLabelSettings) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *CompanyWhiteLabelSettings) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *CompanyWhiteLabelSettings) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetPhone

`func (o *CompanyWhiteLabelSettings) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CompanyWhiteLabelSettings) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CompanyWhiteLabelSettings) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CompanyWhiteLabelSettings) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *CompanyWhiteLabelSettings) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *CompanyWhiteLabelSettings) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetIsLicensor

`func (o *CompanyWhiteLabelSettings) GetIsLicensor() bool`

GetIsLicensor returns the IsLicensor field if non-nil, zero value otherwise.

### GetIsLicensorOk

`func (o *CompanyWhiteLabelSettings) GetIsLicensorOk() (*bool, bool)`

GetIsLicensorOk returns a tuple with the IsLicensor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLicensor

`func (o *CompanyWhiteLabelSettings) SetIsLicensor(v bool)`

SetIsLicensor sets IsLicensor field to given value.

### HasIsLicensor

`func (o *CompanyWhiteLabelSettings) HasIsLicensor() bool`

HasIsLicensor returns a boolean if a field has been set.

### GetHideAbout

`func (o *CompanyWhiteLabelSettings) GetHideAbout() bool`

GetHideAbout returns the HideAbout field if non-nil, zero value otherwise.

### GetHideAboutOk

`func (o *CompanyWhiteLabelSettings) GetHideAboutOk() (*bool, bool)`

GetHideAboutOk returns a tuple with the HideAbout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideAbout

`func (o *CompanyWhiteLabelSettings) SetHideAbout(v bool)`

SetHideAbout sets HideAbout field to given value.

### HasHideAbout

`func (o *CompanyWhiteLabelSettings) HasHideAbout() bool`

HasHideAbout returns a boolean if a field has been set.

### GetLastModified

`func (o *CompanyWhiteLabelSettings) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *CompanyWhiteLabelSettings) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *CompanyWhiteLabelSettings) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *CompanyWhiteLabelSettings) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


