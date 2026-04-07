# QuotaSettingsRequestsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableQuota** | Pointer to **bool** | Specifies whether the storage quota restrictions are enabled. | [optional] 
**DefaultQuota** | [**QuotaSettingsRequestsDtoDefaultQuota**](QuotaSettingsRequestsDtoDefaultQuota.md) |  | 

## Methods

### NewQuotaSettingsRequestsDto

`func NewQuotaSettingsRequestsDto(defaultQuota QuotaSettingsRequestsDtoDefaultQuota, ) *QuotaSettingsRequestsDto`

NewQuotaSettingsRequestsDto instantiates a new QuotaSettingsRequestsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotaSettingsRequestsDtoWithDefaults

`func NewQuotaSettingsRequestsDtoWithDefaults() *QuotaSettingsRequestsDto`

NewQuotaSettingsRequestsDtoWithDefaults instantiates a new QuotaSettingsRequestsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableQuota

`func (o *QuotaSettingsRequestsDto) GetEnableQuota() bool`

GetEnableQuota returns the EnableQuota field if non-nil, zero value otherwise.

### GetEnableQuotaOk

`func (o *QuotaSettingsRequestsDto) GetEnableQuotaOk() (*bool, bool)`

GetEnableQuotaOk returns a tuple with the EnableQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableQuota

`func (o *QuotaSettingsRequestsDto) SetEnableQuota(v bool)`

SetEnableQuota sets EnableQuota field to given value.

### HasEnableQuota

`func (o *QuotaSettingsRequestsDto) HasEnableQuota() bool`

HasEnableQuota returns a boolean if a field has been set.

### GetDefaultQuota

`func (o *QuotaSettingsRequestsDto) GetDefaultQuota() QuotaSettingsRequestsDtoDefaultQuota`

GetDefaultQuota returns the DefaultQuota field if non-nil, zero value otherwise.

### GetDefaultQuotaOk

`func (o *QuotaSettingsRequestsDto) GetDefaultQuotaOk() (*QuotaSettingsRequestsDtoDefaultQuota, bool)`

GetDefaultQuotaOk returns a tuple with the DefaultQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultQuota

`func (o *QuotaSettingsRequestsDto) SetDefaultQuota(v QuotaSettingsRequestsDtoDefaultQuota)`

SetDefaultQuota sets DefaultQuota field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


