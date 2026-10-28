# LicensorDetailsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyName** | Pointer to **NullableString** | The vendor name the About page shows and the letters sign off with. Until details are saved it holds  whatever the installation ships as its built-in vendor, and it is empty on an installation that ships none. | [optional] 
**Site** | Pointer to **NullableString** | The address the vendor name links to, as an absolute URL with its scheme. Empty under the same conditions  as `companyName`. | [optional] 
**Email** | Pointer to **NullableString** | The mailbox the About page offers for reaching the vendor. It is not the portal's own support address, and  it is empty under the same conditions as `companyName`. | [optional] 
**Address** | Pointer to **NullableString** | The postal address of the vendor as one free-form line, in the shape it was saved in - no structure is  imposed on it. | [optional] 
**Phone** | Pointer to **NullableString** | The telephone number of the vendor in the shape it was saved in, with no dialling format enforced. | [optional] 
**IsLicensor** | Pointer to **bool** | Whether these details are those of the licensor of the product itself rather than of a reseller. Saving  through `POST api/2.0/settings/rebranding/company` always clears it, so only details that came with the  installation can report `true`. The name starts with a capital letter, unlike the other fields. | [optional] 
**HideAbout** | Pointer to **bool** | Whether the About page is hidden from the interface. A plan that does not include branding cannot switch it  on: the value is stored as `false` in that case, so it can come back different from what was saved. | [optional] 
**LastModified** | Pointer to **time.Time** | When these details were last stored. Details that were never stored report the moment they were read; the  built-in ONLYOFFICE entry and the answer of the reset operation report `0001-01-01T00:00:00`. | [optional] 

## Methods

### NewLicensorDetailsDto

`func NewLicensorDetailsDto() *LicensorDetailsDto`

NewLicensorDetailsDto instantiates a new LicensorDetailsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicensorDetailsDtoWithDefaults

`func NewLicensorDetailsDtoWithDefaults() *LicensorDetailsDto`

NewLicensorDetailsDtoWithDefaults instantiates a new LicensorDetailsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyName

`func (o *LicensorDetailsDto) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *LicensorDetailsDto) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *LicensorDetailsDto) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *LicensorDetailsDto) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *LicensorDetailsDto) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *LicensorDetailsDto) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetSite

`func (o *LicensorDetailsDto) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *LicensorDetailsDto) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *LicensorDetailsDto) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *LicensorDetailsDto) HasSite() bool`

HasSite returns a boolean if a field has been set.

### SetSiteNil

`func (o *LicensorDetailsDto) SetSiteNil(b bool)`

 SetSiteNil sets the value for Site to be an explicit nil

### UnsetSite
`func (o *LicensorDetailsDto) UnsetSite()`

UnsetSite ensures that no value is present for Site, not even an explicit nil
### GetEmail

`func (o *LicensorDetailsDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *LicensorDetailsDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *LicensorDetailsDto) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *LicensorDetailsDto) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *LicensorDetailsDto) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *LicensorDetailsDto) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetAddress

`func (o *LicensorDetailsDto) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *LicensorDetailsDto) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *LicensorDetailsDto) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *LicensorDetailsDto) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *LicensorDetailsDto) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *LicensorDetailsDto) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetPhone

`func (o *LicensorDetailsDto) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *LicensorDetailsDto) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *LicensorDetailsDto) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *LicensorDetailsDto) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *LicensorDetailsDto) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *LicensorDetailsDto) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetIsLicensor

`func (o *LicensorDetailsDto) GetIsLicensor() bool`

GetIsLicensor returns the IsLicensor field if non-nil, zero value otherwise.

### GetIsLicensorOk

`func (o *LicensorDetailsDto) GetIsLicensorOk() (*bool, bool)`

GetIsLicensorOk returns a tuple with the IsLicensor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLicensor

`func (o *LicensorDetailsDto) SetIsLicensor(v bool)`

SetIsLicensor sets IsLicensor field to given value.

### HasIsLicensor

`func (o *LicensorDetailsDto) HasIsLicensor() bool`

HasIsLicensor returns a boolean if a field has been set.

### GetHideAbout

`func (o *LicensorDetailsDto) GetHideAbout() bool`

GetHideAbout returns the HideAbout field if non-nil, zero value otherwise.

### GetHideAboutOk

`func (o *LicensorDetailsDto) GetHideAboutOk() (*bool, bool)`

GetHideAboutOk returns a tuple with the HideAbout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideAbout

`func (o *LicensorDetailsDto) SetHideAbout(v bool)`

SetHideAbout sets HideAbout field to given value.

### HasHideAbout

`func (o *LicensorDetailsDto) HasHideAbout() bool`

HasHideAbout returns a boolean if a field has been set.

### GetLastModified

`func (o *LicensorDetailsDto) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *LicensorDetailsDto) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *LicensorDetailsDto) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *LicensorDetailsDto) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


