# ExternalResourcesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminPanel** | Pointer to [**ExternalResourceDto**](ExternalResourceDto.md) | The link to the administration panel. It is returned only to the full administrators of a licensed (Enterprise) server (standalone) portal. | [optional] 
**Api** | Pointer to [**ExternalResourceDto**](ExternalResourceDto.md) | The link to the product API. | [optional] 
**Common** | Pointer to [**ExternalResourceDto**](ExternalResourceDto.md) | The link to the common product information. | [optional] 
**Forum** | Pointer to [**ExternalResourceDto**](ExternalResourceDto.md) | The link to the forum. | [optional] 
**Helpcenter** | Pointer to [**ExternalResourceDto**](ExternalResourceDto.md) | The link to the Help Center. | [optional] 
**Integrations** | Pointer to [**ExternalResourceDto**](ExternalResourceDto.md) | The link to the product integrations. | [optional] 
**Site** | Pointer to [**ExternalResourceDto**](ExternalResourceDto.md) | The link to the product website. | [optional] 
**SocialNetworks** | Pointer to [**ExternalResourceDto**](ExternalResourceDto.md) | The link to the product social nerworks. | [optional] 
**Support** | Pointer to [**ExternalResourceDto**](ExternalResourceDto.md) | The link to the product support. | [optional] 
**Videoguides** | Pointer to [**ExternalResourceDto**](ExternalResourceDto.md) | The link to the video guides. | [optional] 

## Methods

### NewExternalResourcesDto

`func NewExternalResourcesDto() *ExternalResourcesDto`

NewExternalResourcesDto instantiates a new ExternalResourcesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalResourcesDtoWithDefaults

`func NewExternalResourcesDtoWithDefaults() *ExternalResourcesDto`

NewExternalResourcesDtoWithDefaults instantiates a new ExternalResourcesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminPanel

`func (o *ExternalResourcesDto) GetAdminPanel() ExternalResourceDto`

GetAdminPanel returns the AdminPanel field if non-nil, zero value otherwise.

### GetAdminPanelOk

`func (o *ExternalResourcesDto) GetAdminPanelOk() (*ExternalResourceDto, bool)`

GetAdminPanelOk returns a tuple with the AdminPanel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPanel

`func (o *ExternalResourcesDto) SetAdminPanel(v ExternalResourceDto)`

SetAdminPanel sets AdminPanel field to given value.

### HasAdminPanel

`func (o *ExternalResourcesDto) HasAdminPanel() bool`

HasAdminPanel returns a boolean if a field has been set.

### GetApi

`func (o *ExternalResourcesDto) GetApi() ExternalResourceDto`

GetApi returns the Api field if non-nil, zero value otherwise.

### GetApiOk

`func (o *ExternalResourcesDto) GetApiOk() (*ExternalResourceDto, bool)`

GetApiOk returns a tuple with the Api field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApi

`func (o *ExternalResourcesDto) SetApi(v ExternalResourceDto)`

SetApi sets Api field to given value.

### HasApi

`func (o *ExternalResourcesDto) HasApi() bool`

HasApi returns a boolean if a field has been set.

### GetCommon

`func (o *ExternalResourcesDto) GetCommon() ExternalResourceDto`

GetCommon returns the Common field if non-nil, zero value otherwise.

### GetCommonOk

`func (o *ExternalResourcesDto) GetCommonOk() (*ExternalResourceDto, bool)`

GetCommonOk returns a tuple with the Common field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommon

`func (o *ExternalResourcesDto) SetCommon(v ExternalResourceDto)`

SetCommon sets Common field to given value.

### HasCommon

`func (o *ExternalResourcesDto) HasCommon() bool`

HasCommon returns a boolean if a field has been set.

### GetForum

`func (o *ExternalResourcesDto) GetForum() ExternalResourceDto`

GetForum returns the Forum field if non-nil, zero value otherwise.

### GetForumOk

`func (o *ExternalResourcesDto) GetForumOk() (*ExternalResourceDto, bool)`

GetForumOk returns a tuple with the Forum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForum

`func (o *ExternalResourcesDto) SetForum(v ExternalResourceDto)`

SetForum sets Forum field to given value.

### HasForum

`func (o *ExternalResourcesDto) HasForum() bool`

HasForum returns a boolean if a field has been set.

### GetHelpcenter

`func (o *ExternalResourcesDto) GetHelpcenter() ExternalResourceDto`

GetHelpcenter returns the Helpcenter field if non-nil, zero value otherwise.

### GetHelpcenterOk

`func (o *ExternalResourcesDto) GetHelpcenterOk() (*ExternalResourceDto, bool)`

GetHelpcenterOk returns a tuple with the Helpcenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpcenter

`func (o *ExternalResourcesDto) SetHelpcenter(v ExternalResourceDto)`

SetHelpcenter sets Helpcenter field to given value.

### HasHelpcenter

`func (o *ExternalResourcesDto) HasHelpcenter() bool`

HasHelpcenter returns a boolean if a field has been set.

### GetIntegrations

`func (o *ExternalResourcesDto) GetIntegrations() ExternalResourceDto`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *ExternalResourcesDto) GetIntegrationsOk() (*ExternalResourceDto, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *ExternalResourcesDto) SetIntegrations(v ExternalResourceDto)`

SetIntegrations sets Integrations field to given value.

### HasIntegrations

`func (o *ExternalResourcesDto) HasIntegrations() bool`

HasIntegrations returns a boolean if a field has been set.

### GetSite

`func (o *ExternalResourcesDto) GetSite() ExternalResourceDto`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *ExternalResourcesDto) GetSiteOk() (*ExternalResourceDto, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *ExternalResourcesDto) SetSite(v ExternalResourceDto)`

SetSite sets Site field to given value.

### HasSite

`func (o *ExternalResourcesDto) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetSocialNetworks

`func (o *ExternalResourcesDto) GetSocialNetworks() ExternalResourceDto`

GetSocialNetworks returns the SocialNetworks field if non-nil, zero value otherwise.

### GetSocialNetworksOk

`func (o *ExternalResourcesDto) GetSocialNetworksOk() (*ExternalResourceDto, bool)`

GetSocialNetworksOk returns a tuple with the SocialNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialNetworks

`func (o *ExternalResourcesDto) SetSocialNetworks(v ExternalResourceDto)`

SetSocialNetworks sets SocialNetworks field to given value.

### HasSocialNetworks

`func (o *ExternalResourcesDto) HasSocialNetworks() bool`

HasSocialNetworks returns a boolean if a field has been set.

### GetSupport

`func (o *ExternalResourcesDto) GetSupport() ExternalResourceDto`

GetSupport returns the Support field if non-nil, zero value otherwise.

### GetSupportOk

`func (o *ExternalResourcesDto) GetSupportOk() (*ExternalResourceDto, bool)`

GetSupportOk returns a tuple with the Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport

`func (o *ExternalResourcesDto) SetSupport(v ExternalResourceDto)`

SetSupport sets Support field to given value.

### HasSupport

`func (o *ExternalResourcesDto) HasSupport() bool`

HasSupport returns a boolean if a field has been set.

### GetVideoguides

`func (o *ExternalResourcesDto) GetVideoguides() ExternalResourceDto`

GetVideoguides returns the Videoguides field if non-nil, zero value otherwise.

### GetVideoguidesOk

`func (o *ExternalResourcesDto) GetVideoguidesOk() (*ExternalResourceDto, bool)`

GetVideoguidesOk returns a tuple with the Videoguides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoguides

`func (o *ExternalResourcesDto) SetVideoguides(v ExternalResourceDto)`

SetVideoguides sets Videoguides field to given value.

### HasVideoguides

`func (o *ExternalResourcesDto) HasVideoguides() bool`

HasVideoguides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


