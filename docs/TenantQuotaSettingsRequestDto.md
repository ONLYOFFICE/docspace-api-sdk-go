# TenantQuotaSettingsRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **int32** | The tenant the limit applies to, by tenant ID. Only a self-hosted installation has more than one, which is  why the operation is refused on SaaS. | 
**Quota** | Pointer to **int64** | The limit in bytes. A negative value is not a smaller limit but the absence of one: it removes whatever limit  the tenant had. The value is a ceiling on stored data and says nothing about how much of it is already used. | [optional] 

## Methods

### NewTenantQuotaSettingsRequestDto

`func NewTenantQuotaSettingsRequestDto(tenantId int32, ) *TenantQuotaSettingsRequestDto`

NewTenantQuotaSettingsRequestDto instantiates a new TenantQuotaSettingsRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantQuotaSettingsRequestDtoWithDefaults

`func NewTenantQuotaSettingsRequestDtoWithDefaults() *TenantQuotaSettingsRequestDto`

NewTenantQuotaSettingsRequestDtoWithDefaults instantiates a new TenantQuotaSettingsRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *TenantQuotaSettingsRequestDto) GetTenantId() int32`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TenantQuotaSettingsRequestDto) GetTenantIdOk() (*int32, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TenantQuotaSettingsRequestDto) SetTenantId(v int32)`

SetTenantId sets TenantId field to given value.


### GetQuota

`func (o *TenantQuotaSettingsRequestDto) GetQuota() int64`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *TenantQuotaSettingsRequestDto) GetQuotaOk() (*int64, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *TenantQuotaSettingsRequestDto) SetQuota(v int64)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *TenantQuotaSettingsRequestDto) HasQuota() bool`

HasQuota returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


