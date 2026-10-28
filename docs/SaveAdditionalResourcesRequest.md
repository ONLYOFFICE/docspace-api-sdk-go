# SaveAdditionalResourcesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDocsEnabled** | Pointer to **bool** | Whether the getting-started documents are offered. | [optional] 
**HelpCenterEnabled** | Pointer to **bool** | Whether the help center is linked. | [optional] 
**FeedbackAndSupportEnabled** | Pointer to **bool** | Whether the feedback and support link is shown. | [optional] 
**UserForumEnabled** | Pointer to **bool** | Whether the user forum is linked. | [optional] 
**VideoGuidesEnabled** | Pointer to **bool** | Whether the video guides are linked. | [optional] 
**LicenseAgreementsEnabled** | Pointer to **bool** | Whether the license agreements are linked. | [optional] 
**LastModified** | Pointer to **time.Time** | Accepted for compatibility with earlier clients and not read: the server keeps its own value. | [optional] 

## Methods

### NewSaveAdditionalResourcesRequest

`func NewSaveAdditionalResourcesRequest() *SaveAdditionalResourcesRequest`

NewSaveAdditionalResourcesRequest instantiates a new SaveAdditionalResourcesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveAdditionalResourcesRequestWithDefaults

`func NewSaveAdditionalResourcesRequestWithDefaults() *SaveAdditionalResourcesRequest`

NewSaveAdditionalResourcesRequestWithDefaults instantiates a new SaveAdditionalResourcesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDocsEnabled

`func (o *SaveAdditionalResourcesRequest) GetStartDocsEnabled() bool`

GetStartDocsEnabled returns the StartDocsEnabled field if non-nil, zero value otherwise.

### GetStartDocsEnabledOk

`func (o *SaveAdditionalResourcesRequest) GetStartDocsEnabledOk() (*bool, bool)`

GetStartDocsEnabledOk returns a tuple with the StartDocsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDocsEnabled

`func (o *SaveAdditionalResourcesRequest) SetStartDocsEnabled(v bool)`

SetStartDocsEnabled sets StartDocsEnabled field to given value.

### HasStartDocsEnabled

`func (o *SaveAdditionalResourcesRequest) HasStartDocsEnabled() bool`

HasStartDocsEnabled returns a boolean if a field has been set.

### GetHelpCenterEnabled

`func (o *SaveAdditionalResourcesRequest) GetHelpCenterEnabled() bool`

GetHelpCenterEnabled returns the HelpCenterEnabled field if non-nil, zero value otherwise.

### GetHelpCenterEnabledOk

`func (o *SaveAdditionalResourcesRequest) GetHelpCenterEnabledOk() (*bool, bool)`

GetHelpCenterEnabledOk returns a tuple with the HelpCenterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpCenterEnabled

`func (o *SaveAdditionalResourcesRequest) SetHelpCenterEnabled(v bool)`

SetHelpCenterEnabled sets HelpCenterEnabled field to given value.

### HasHelpCenterEnabled

`func (o *SaveAdditionalResourcesRequest) HasHelpCenterEnabled() bool`

HasHelpCenterEnabled returns a boolean if a field has been set.

### GetFeedbackAndSupportEnabled

`func (o *SaveAdditionalResourcesRequest) GetFeedbackAndSupportEnabled() bool`

GetFeedbackAndSupportEnabled returns the FeedbackAndSupportEnabled field if non-nil, zero value otherwise.

### GetFeedbackAndSupportEnabledOk

`func (o *SaveAdditionalResourcesRequest) GetFeedbackAndSupportEnabledOk() (*bool, bool)`

GetFeedbackAndSupportEnabledOk returns a tuple with the FeedbackAndSupportEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackAndSupportEnabled

`func (o *SaveAdditionalResourcesRequest) SetFeedbackAndSupportEnabled(v bool)`

SetFeedbackAndSupportEnabled sets FeedbackAndSupportEnabled field to given value.

### HasFeedbackAndSupportEnabled

`func (o *SaveAdditionalResourcesRequest) HasFeedbackAndSupportEnabled() bool`

HasFeedbackAndSupportEnabled returns a boolean if a field has been set.

### GetUserForumEnabled

`func (o *SaveAdditionalResourcesRequest) GetUserForumEnabled() bool`

GetUserForumEnabled returns the UserForumEnabled field if non-nil, zero value otherwise.

### GetUserForumEnabledOk

`func (o *SaveAdditionalResourcesRequest) GetUserForumEnabledOk() (*bool, bool)`

GetUserForumEnabledOk returns a tuple with the UserForumEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserForumEnabled

`func (o *SaveAdditionalResourcesRequest) SetUserForumEnabled(v bool)`

SetUserForumEnabled sets UserForumEnabled field to given value.

### HasUserForumEnabled

`func (o *SaveAdditionalResourcesRequest) HasUserForumEnabled() bool`

HasUserForumEnabled returns a boolean if a field has been set.

### GetVideoGuidesEnabled

`func (o *SaveAdditionalResourcesRequest) GetVideoGuidesEnabled() bool`

GetVideoGuidesEnabled returns the VideoGuidesEnabled field if non-nil, zero value otherwise.

### GetVideoGuidesEnabledOk

`func (o *SaveAdditionalResourcesRequest) GetVideoGuidesEnabledOk() (*bool, bool)`

GetVideoGuidesEnabledOk returns a tuple with the VideoGuidesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoGuidesEnabled

`func (o *SaveAdditionalResourcesRequest) SetVideoGuidesEnabled(v bool)`

SetVideoGuidesEnabled sets VideoGuidesEnabled field to given value.

### HasVideoGuidesEnabled

`func (o *SaveAdditionalResourcesRequest) HasVideoGuidesEnabled() bool`

HasVideoGuidesEnabled returns a boolean if a field has been set.

### GetLicenseAgreementsEnabled

`func (o *SaveAdditionalResourcesRequest) GetLicenseAgreementsEnabled() bool`

GetLicenseAgreementsEnabled returns the LicenseAgreementsEnabled field if non-nil, zero value otherwise.

### GetLicenseAgreementsEnabledOk

`func (o *SaveAdditionalResourcesRequest) GetLicenseAgreementsEnabledOk() (*bool, bool)`

GetLicenseAgreementsEnabledOk returns a tuple with the LicenseAgreementsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseAgreementsEnabled

`func (o *SaveAdditionalResourcesRequest) SetLicenseAgreementsEnabled(v bool)`

SetLicenseAgreementsEnabled sets LicenseAgreementsEnabled field to given value.

### HasLicenseAgreementsEnabled

`func (o *SaveAdditionalResourcesRequest) HasLicenseAgreementsEnabled() bool`

HasLicenseAgreementsEnabled returns a boolean if a field has been set.

### GetLastModified

`func (o *SaveAdditionalResourcesRequest) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *SaveAdditionalResourcesRequest) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *SaveAdditionalResourcesRequest) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *SaveAdditionalResourcesRequest) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


