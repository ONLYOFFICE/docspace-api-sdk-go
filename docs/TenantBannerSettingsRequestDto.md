# TenantBannerSettingsRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hidden** | Pointer to **bool** | Whether the promotional banners are hidden from every user of the portal. The flag is only honoured on a  self-hosted installation; a SaaS portal keeps showing the banners whatever is stored here. | [optional] 

## Methods

### NewTenantBannerSettingsRequestDto

`func NewTenantBannerSettingsRequestDto() *TenantBannerSettingsRequestDto`

NewTenantBannerSettingsRequestDto instantiates a new TenantBannerSettingsRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantBannerSettingsRequestDtoWithDefaults

`func NewTenantBannerSettingsRequestDtoWithDefaults() *TenantBannerSettingsRequestDto`

NewTenantBannerSettingsRequestDtoWithDefaults instantiates a new TenantBannerSettingsRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHidden

`func (o *TenantBannerSettingsRequestDto) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *TenantBannerSettingsRequestDto) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *TenantBannerSettingsRequestDto) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *TenantBannerSettingsRequestDto) HasHidden() bool`

HasHidden returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


