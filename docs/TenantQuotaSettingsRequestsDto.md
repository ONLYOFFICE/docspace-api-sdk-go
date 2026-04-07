# TenantQuotaSettingsRequestsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **int32** | The ID of the tenant whose quota is being configured. | 
**Quota** | Pointer to **int64** | The storage quota limit in bytes allocated to the tenant. | [optional] 

## Methods

### NewTenantQuotaSettingsRequestsDto

`func NewTenantQuotaSettingsRequestsDto(tenantId int32, ) *TenantQuotaSettingsRequestsDto`

NewTenantQuotaSettingsRequestsDto instantiates a new TenantQuotaSettingsRequestsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantQuotaSettingsRequestsDtoWithDefaults

`func NewTenantQuotaSettingsRequestsDtoWithDefaults() *TenantQuotaSettingsRequestsDto`

NewTenantQuotaSettingsRequestsDtoWithDefaults instantiates a new TenantQuotaSettingsRequestsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *TenantQuotaSettingsRequestsDto) GetTenantId() int32`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TenantQuotaSettingsRequestsDto) GetTenantIdOk() (*int32, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TenantQuotaSettingsRequestsDto) SetTenantId(v int32)`

SetTenantId sets TenantId field to given value.


### GetQuota

`func (o *TenantQuotaSettingsRequestsDto) GetQuota() int64`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *TenantQuotaSettingsRequestsDto) GetQuotaOk() (*int64, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *TenantQuotaSettingsRequestsDto) SetQuota(v int64)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *TenantQuotaSettingsRequestsDto) HasQuota() bool`

HasQuota returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


