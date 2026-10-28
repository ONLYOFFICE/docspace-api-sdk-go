# AdditionalResourcesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDocsEnabled** | Pointer to **bool** | Whether the sample documents that ONLYOFFICE ships may be placed in a new user's Documents. Unlike the link  flags below it depends on nothing that has to be configured, so its built-in value is always `true`. | [optional] 
**HelpCenterEnabled** | Pointer to **bool** | Whether the interface may offer the Help Center entry. It is `false` both when the entry was switched off  for the installation and when the installation configures no Help Center address at all; the addresses  themselves are not part of this answer and arrive in `externalResources` of `GET api/2.0/settings`. | [optional] 
**FeedbackAndSupportEnabled** | Pointer to **bool** | Whether the interface may offer the Feedback and Support entry, `false` for the same two reasons as  `helpCenterEnabled`. | [optional] 
**UserForumEnabled** | Pointer to **bool** | Whether the interface may offer the user forum entry, `false` for the same two reasons as  `helpCenterEnabled`. | [optional] 
**VideoGuidesEnabled** | Pointer to **bool** | Whether the interface may offer the Video Guides entry, `false` for the same two reasons as  `helpCenterEnabled`. | [optional] 
**LicenseAgreementsEnabled** | Pointer to **bool** | Whether the interface may offer the License Agreements entry, `false` for the same two reasons as  `helpCenterEnabled`. | [optional] 
**LastModified** | Pointer to **time.Time** | When these flags were last stored. Flags that were never stored report the moment they were read, and the  answer of the reset operation reports `0001-01-01T00:00:00`. | [optional] 

## Methods

### NewAdditionalResourcesDto

`func NewAdditionalResourcesDto() *AdditionalResourcesDto`

NewAdditionalResourcesDto instantiates a new AdditionalResourcesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalResourcesDtoWithDefaults

`func NewAdditionalResourcesDtoWithDefaults() *AdditionalResourcesDto`

NewAdditionalResourcesDtoWithDefaults instantiates a new AdditionalResourcesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDocsEnabled

`func (o *AdditionalResourcesDto) GetStartDocsEnabled() bool`

GetStartDocsEnabled returns the StartDocsEnabled field if non-nil, zero value otherwise.

### GetStartDocsEnabledOk

`func (o *AdditionalResourcesDto) GetStartDocsEnabledOk() (*bool, bool)`

GetStartDocsEnabledOk returns a tuple with the StartDocsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDocsEnabled

`func (o *AdditionalResourcesDto) SetStartDocsEnabled(v bool)`

SetStartDocsEnabled sets StartDocsEnabled field to given value.

### HasStartDocsEnabled

`func (o *AdditionalResourcesDto) HasStartDocsEnabled() bool`

HasStartDocsEnabled returns a boolean if a field has been set.

### GetHelpCenterEnabled

`func (o *AdditionalResourcesDto) GetHelpCenterEnabled() bool`

GetHelpCenterEnabled returns the HelpCenterEnabled field if non-nil, zero value otherwise.

### GetHelpCenterEnabledOk

`func (o *AdditionalResourcesDto) GetHelpCenterEnabledOk() (*bool, bool)`

GetHelpCenterEnabledOk returns a tuple with the HelpCenterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpCenterEnabled

`func (o *AdditionalResourcesDto) SetHelpCenterEnabled(v bool)`

SetHelpCenterEnabled sets HelpCenterEnabled field to given value.

### HasHelpCenterEnabled

`func (o *AdditionalResourcesDto) HasHelpCenterEnabled() bool`

HasHelpCenterEnabled returns a boolean if a field has been set.

### GetFeedbackAndSupportEnabled

`func (o *AdditionalResourcesDto) GetFeedbackAndSupportEnabled() bool`

GetFeedbackAndSupportEnabled returns the FeedbackAndSupportEnabled field if non-nil, zero value otherwise.

### GetFeedbackAndSupportEnabledOk

`func (o *AdditionalResourcesDto) GetFeedbackAndSupportEnabledOk() (*bool, bool)`

GetFeedbackAndSupportEnabledOk returns a tuple with the FeedbackAndSupportEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackAndSupportEnabled

`func (o *AdditionalResourcesDto) SetFeedbackAndSupportEnabled(v bool)`

SetFeedbackAndSupportEnabled sets FeedbackAndSupportEnabled field to given value.

### HasFeedbackAndSupportEnabled

`func (o *AdditionalResourcesDto) HasFeedbackAndSupportEnabled() bool`

HasFeedbackAndSupportEnabled returns a boolean if a field has been set.

### GetUserForumEnabled

`func (o *AdditionalResourcesDto) GetUserForumEnabled() bool`

GetUserForumEnabled returns the UserForumEnabled field if non-nil, zero value otherwise.

### GetUserForumEnabledOk

`func (o *AdditionalResourcesDto) GetUserForumEnabledOk() (*bool, bool)`

GetUserForumEnabledOk returns a tuple with the UserForumEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserForumEnabled

`func (o *AdditionalResourcesDto) SetUserForumEnabled(v bool)`

SetUserForumEnabled sets UserForumEnabled field to given value.

### HasUserForumEnabled

`func (o *AdditionalResourcesDto) HasUserForumEnabled() bool`

HasUserForumEnabled returns a boolean if a field has been set.

### GetVideoGuidesEnabled

`func (o *AdditionalResourcesDto) GetVideoGuidesEnabled() bool`

GetVideoGuidesEnabled returns the VideoGuidesEnabled field if non-nil, zero value otherwise.

### GetVideoGuidesEnabledOk

`func (o *AdditionalResourcesDto) GetVideoGuidesEnabledOk() (*bool, bool)`

GetVideoGuidesEnabledOk returns a tuple with the VideoGuidesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoGuidesEnabled

`func (o *AdditionalResourcesDto) SetVideoGuidesEnabled(v bool)`

SetVideoGuidesEnabled sets VideoGuidesEnabled field to given value.

### HasVideoGuidesEnabled

`func (o *AdditionalResourcesDto) HasVideoGuidesEnabled() bool`

HasVideoGuidesEnabled returns a boolean if a field has been set.

### GetLicenseAgreementsEnabled

`func (o *AdditionalResourcesDto) GetLicenseAgreementsEnabled() bool`

GetLicenseAgreementsEnabled returns the LicenseAgreementsEnabled field if non-nil, zero value otherwise.

### GetLicenseAgreementsEnabledOk

`func (o *AdditionalResourcesDto) GetLicenseAgreementsEnabledOk() (*bool, bool)`

GetLicenseAgreementsEnabledOk returns a tuple with the LicenseAgreementsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseAgreementsEnabled

`func (o *AdditionalResourcesDto) SetLicenseAgreementsEnabled(v bool)`

SetLicenseAgreementsEnabled sets LicenseAgreementsEnabled field to given value.

### HasLicenseAgreementsEnabled

`func (o *AdditionalResourcesDto) HasLicenseAgreementsEnabled() bool`

HasLicenseAgreementsEnabled returns a boolean if a field has been set.

### GetLastModified

`func (o *AdditionalResourcesDto) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *AdditionalResourcesDto) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *AdditionalResourcesDto) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *AdditionalResourcesDto) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


