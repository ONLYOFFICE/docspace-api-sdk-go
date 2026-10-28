# AdditionalWhiteLabelSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDocsEnabled** | **bool** | Whether the sample documents that ONLYOFFICE ships may be placed in a new user's Documents. Unlike the link  flags below it depends on nothing that has to be configured, so its built-in value is always `true`. | 
**HelpCenterEnabled** | **bool** | Whether the interface may offer the Help Center entry. It is `false` both when the entry was switched off  for the installation and when the installation configures no Help Center address at all; the addresses  themselves are not part of this answer and arrive in `externalResources` of `GET api/2.0/settings`. | 
**FeedbackAndSupportEnabled** | **bool** | Whether the interface may offer the Feedback and Support entry, `false` for the same two reasons as  `helpCenterEnabled`. | 
**UserForumEnabled** | **bool** | Whether the interface may offer the user forum entry, `false` for the same two reasons as  `helpCenterEnabled`. | 
**VideoGuidesEnabled** | **bool** | Whether the interface may offer the Video Guides entry, `false` for the same two reasons as  `helpCenterEnabled`. | 
**LicenseAgreementsEnabled** | **bool** | Whether the interface may offer the License Agreements entry, `false` for the same two reasons as  `helpCenterEnabled`. | 
**IsDefault** | **bool** | Whether all six flags still hold the values the installation starts out with. It turns `false` as soon as  one of them is saved differently and `true` again after `DELETE api/2.0/settings/rebranding/additional`.  Because a link flag starts out off when no address is configured for it, `true` does not mean every entry  is on. | 

## Methods

### NewAdditionalWhiteLabelSettingsDto

`func NewAdditionalWhiteLabelSettingsDto(startDocsEnabled bool, helpCenterEnabled bool, feedbackAndSupportEnabled bool, userForumEnabled bool, videoGuidesEnabled bool, licenseAgreementsEnabled bool, isDefault bool, ) *AdditionalWhiteLabelSettingsDto`

NewAdditionalWhiteLabelSettingsDto instantiates a new AdditionalWhiteLabelSettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalWhiteLabelSettingsDtoWithDefaults

`func NewAdditionalWhiteLabelSettingsDtoWithDefaults() *AdditionalWhiteLabelSettingsDto`

NewAdditionalWhiteLabelSettingsDtoWithDefaults instantiates a new AdditionalWhiteLabelSettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDocsEnabled

`func (o *AdditionalWhiteLabelSettingsDto) GetStartDocsEnabled() bool`

GetStartDocsEnabled returns the StartDocsEnabled field if non-nil, zero value otherwise.

### GetStartDocsEnabledOk

`func (o *AdditionalWhiteLabelSettingsDto) GetStartDocsEnabledOk() (*bool, bool)`

GetStartDocsEnabledOk returns a tuple with the StartDocsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDocsEnabled

`func (o *AdditionalWhiteLabelSettingsDto) SetStartDocsEnabled(v bool)`

SetStartDocsEnabled sets StartDocsEnabled field to given value.


### GetHelpCenterEnabled

`func (o *AdditionalWhiteLabelSettingsDto) GetHelpCenterEnabled() bool`

GetHelpCenterEnabled returns the HelpCenterEnabled field if non-nil, zero value otherwise.

### GetHelpCenterEnabledOk

`func (o *AdditionalWhiteLabelSettingsDto) GetHelpCenterEnabledOk() (*bool, bool)`

GetHelpCenterEnabledOk returns a tuple with the HelpCenterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpCenterEnabled

`func (o *AdditionalWhiteLabelSettingsDto) SetHelpCenterEnabled(v bool)`

SetHelpCenterEnabled sets HelpCenterEnabled field to given value.


### GetFeedbackAndSupportEnabled

`func (o *AdditionalWhiteLabelSettingsDto) GetFeedbackAndSupportEnabled() bool`

GetFeedbackAndSupportEnabled returns the FeedbackAndSupportEnabled field if non-nil, zero value otherwise.

### GetFeedbackAndSupportEnabledOk

`func (o *AdditionalWhiteLabelSettingsDto) GetFeedbackAndSupportEnabledOk() (*bool, bool)`

GetFeedbackAndSupportEnabledOk returns a tuple with the FeedbackAndSupportEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackAndSupportEnabled

`func (o *AdditionalWhiteLabelSettingsDto) SetFeedbackAndSupportEnabled(v bool)`

SetFeedbackAndSupportEnabled sets FeedbackAndSupportEnabled field to given value.


### GetUserForumEnabled

`func (o *AdditionalWhiteLabelSettingsDto) GetUserForumEnabled() bool`

GetUserForumEnabled returns the UserForumEnabled field if non-nil, zero value otherwise.

### GetUserForumEnabledOk

`func (o *AdditionalWhiteLabelSettingsDto) GetUserForumEnabledOk() (*bool, bool)`

GetUserForumEnabledOk returns a tuple with the UserForumEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserForumEnabled

`func (o *AdditionalWhiteLabelSettingsDto) SetUserForumEnabled(v bool)`

SetUserForumEnabled sets UserForumEnabled field to given value.


### GetVideoGuidesEnabled

`func (o *AdditionalWhiteLabelSettingsDto) GetVideoGuidesEnabled() bool`

GetVideoGuidesEnabled returns the VideoGuidesEnabled field if non-nil, zero value otherwise.

### GetVideoGuidesEnabledOk

`func (o *AdditionalWhiteLabelSettingsDto) GetVideoGuidesEnabledOk() (*bool, bool)`

GetVideoGuidesEnabledOk returns a tuple with the VideoGuidesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoGuidesEnabled

`func (o *AdditionalWhiteLabelSettingsDto) SetVideoGuidesEnabled(v bool)`

SetVideoGuidesEnabled sets VideoGuidesEnabled field to given value.


### GetLicenseAgreementsEnabled

`func (o *AdditionalWhiteLabelSettingsDto) GetLicenseAgreementsEnabled() bool`

GetLicenseAgreementsEnabled returns the LicenseAgreementsEnabled field if non-nil, zero value otherwise.

### GetLicenseAgreementsEnabledOk

`func (o *AdditionalWhiteLabelSettingsDto) GetLicenseAgreementsEnabledOk() (*bool, bool)`

GetLicenseAgreementsEnabledOk returns a tuple with the LicenseAgreementsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseAgreementsEnabled

`func (o *AdditionalWhiteLabelSettingsDto) SetLicenseAgreementsEnabled(v bool)`

SetLicenseAgreementsEnabled sets LicenseAgreementsEnabled field to given value.


### GetIsDefault

`func (o *AdditionalWhiteLabelSettingsDto) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *AdditionalWhiteLabelSettingsDto) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *AdditionalWhiteLabelSettingsDto) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


