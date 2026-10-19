# SaveCompanyInfoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyName** | Pointer to **NullableString** | The company name. | [optional] 
**Site** | Pointer to **NullableString** | The company website, as an absolute URL. | [optional] 
**Email** | Pointer to **NullableString** | The contact email address. | [optional] 
**Address** | Pointer to **NullableString** | The postal address. | [optional] 
**Phone** | Pointer to **NullableString** | The contact phone number. | [optional] 
**IsLicensor** | Pointer to **bool** | Accepted for compatibility with earlier clients and not read: saved details are never those of the licensor, so the server always stores `false`. | [optional] 
**HideAbout** | Pointer to **bool** | Whether the About page is hidden. | [optional] 
**LastModified** | Pointer to **time.Time** | Accepted for compatibility with earlier clients and not read: the server keeps its own value. | [optional] 

## Methods

### NewSaveCompanyInfoRequest

`func NewSaveCompanyInfoRequest() *SaveCompanyInfoRequest`

NewSaveCompanyInfoRequest instantiates a new SaveCompanyInfoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveCompanyInfoRequestWithDefaults

`func NewSaveCompanyInfoRequestWithDefaults() *SaveCompanyInfoRequest`

NewSaveCompanyInfoRequestWithDefaults instantiates a new SaveCompanyInfoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyName

`func (o *SaveCompanyInfoRequest) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *SaveCompanyInfoRequest) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *SaveCompanyInfoRequest) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *SaveCompanyInfoRequest) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *SaveCompanyInfoRequest) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *SaveCompanyInfoRequest) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetSite

`func (o *SaveCompanyInfoRequest) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *SaveCompanyInfoRequest) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *SaveCompanyInfoRequest) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *SaveCompanyInfoRequest) HasSite() bool`

HasSite returns a boolean if a field has been set.

### SetSiteNil

`func (o *SaveCompanyInfoRequest) SetSiteNil(b bool)`

 SetSiteNil sets the value for Site to be an explicit nil

### UnsetSite
`func (o *SaveCompanyInfoRequest) UnsetSite()`

UnsetSite ensures that no value is present for Site, not even an explicit nil
### GetEmail

`func (o *SaveCompanyInfoRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SaveCompanyInfoRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SaveCompanyInfoRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SaveCompanyInfoRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *SaveCompanyInfoRequest) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *SaveCompanyInfoRequest) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetAddress

`func (o *SaveCompanyInfoRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SaveCompanyInfoRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SaveCompanyInfoRequest) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *SaveCompanyInfoRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *SaveCompanyInfoRequest) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *SaveCompanyInfoRequest) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetPhone

`func (o *SaveCompanyInfoRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *SaveCompanyInfoRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *SaveCompanyInfoRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *SaveCompanyInfoRequest) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *SaveCompanyInfoRequest) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *SaveCompanyInfoRequest) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetIsLicensor

`func (o *SaveCompanyInfoRequest) GetIsLicensor() bool`

GetIsLicensor returns the IsLicensor field if non-nil, zero value otherwise.

### GetIsLicensorOk

`func (o *SaveCompanyInfoRequest) GetIsLicensorOk() (*bool, bool)`

GetIsLicensorOk returns a tuple with the IsLicensor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLicensor

`func (o *SaveCompanyInfoRequest) SetIsLicensor(v bool)`

SetIsLicensor sets IsLicensor field to given value.

### HasIsLicensor

`func (o *SaveCompanyInfoRequest) HasIsLicensor() bool`

HasIsLicensor returns a boolean if a field has been set.

### GetHideAbout

`func (o *SaveCompanyInfoRequest) GetHideAbout() bool`

GetHideAbout returns the HideAbout field if non-nil, zero value otherwise.

### GetHideAboutOk

`func (o *SaveCompanyInfoRequest) GetHideAboutOk() (*bool, bool)`

GetHideAboutOk returns a tuple with the HideAbout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideAbout

`func (o *SaveCompanyInfoRequest) SetHideAbout(v bool)`

SetHideAbout sets HideAbout field to given value.

### HasHideAbout

`func (o *SaveCompanyInfoRequest) HasHideAbout() bool`

HasHideAbout returns a boolean if a field has been set.

### GetLastModified

`func (o *SaveCompanyInfoRequest) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *SaveCompanyInfoRequest) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *SaveCompanyInfoRequest) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *SaveCompanyInfoRequest) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


