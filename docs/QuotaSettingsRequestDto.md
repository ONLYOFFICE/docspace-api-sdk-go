# QuotaSettingsRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableQuota** | Pointer to **bool** | Whether the limit is enforced at all. While it is false the size is ignored and nothing created afterwards  carries a limit; objects that already have one keep it either way. | [optional] 
**DefaultQuota** | [**QuotaSettingsRequestDtoDefaultQuota**](QuotaSettingsRequestDtoDefaultQuota.md) |  | 

## Methods

### NewQuotaSettingsRequestDto

`func NewQuotaSettingsRequestDto(defaultQuota QuotaSettingsRequestDtoDefaultQuota, ) *QuotaSettingsRequestDto`

NewQuotaSettingsRequestDto instantiates a new QuotaSettingsRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotaSettingsRequestDtoWithDefaults

`func NewQuotaSettingsRequestDtoWithDefaults() *QuotaSettingsRequestDto`

NewQuotaSettingsRequestDtoWithDefaults instantiates a new QuotaSettingsRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableQuota

`func (o *QuotaSettingsRequestDto) GetEnableQuota() bool`

GetEnableQuota returns the EnableQuota field if non-nil, zero value otherwise.

### GetEnableQuotaOk

`func (o *QuotaSettingsRequestDto) GetEnableQuotaOk() (*bool, bool)`

GetEnableQuotaOk returns a tuple with the EnableQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableQuota

`func (o *QuotaSettingsRequestDto) SetEnableQuota(v bool)`

SetEnableQuota sets EnableQuota field to given value.

### HasEnableQuota

`func (o *QuotaSettingsRequestDto) HasEnableQuota() bool`

HasEnableQuota returns a boolean if a field has been set.

### GetDefaultQuota

`func (o *QuotaSettingsRequestDto) GetDefaultQuota() QuotaSettingsRequestDtoDefaultQuota`

GetDefaultQuota returns the DefaultQuota field if non-nil, zero value otherwise.

### GetDefaultQuotaOk

`func (o *QuotaSettingsRequestDto) GetDefaultQuotaOk() (*QuotaSettingsRequestDtoDefaultQuota, bool)`

GetDefaultQuotaOk returns a tuple with the DefaultQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultQuota

`func (o *QuotaSettingsRequestDto) SetDefaultQuota(v QuotaSettingsRequestDtoDefaultQuota)`

SetDefaultQuota sets DefaultQuota field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


