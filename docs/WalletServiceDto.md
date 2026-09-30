# WalletServiceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The identifier of the quota, which is what the tariff reports as a quota `id` and what a purchase names.  A negative value belongs to a built-in quota rather than one on the price list. | 
**Title** | Pointer to **string** | The quota name in the portal language, for printing rather than matching. It is empty when this build  ships no wording for the quota, which is normal for a quota that is not on the public price list. | [optional] 
**Price** | [**PriceDto**](PriceDto.md) | What the quota costs, in the currency resolved for the request. Its `value` is empty for a quota that is  not sold for money, which is what `free`, `trial` and `nonProfit` describe. | 
**NonProfit** | **bool** | Whether this is the non-profit quota, which is granted rather than bought. A portal on it cannot buy any  other plan, so a catalogue asked for plans returns this one alone. | 
**Free** | **bool** | Whether this is the free quota a portal falls back to when nothing is paid for. It has no end date and  the tightest limits of any quota. | 
**Trial** | **bool** | Whether this is the trial quota, which grants the paid limits for a while and then expires. A trial is not  extended by paying - a plan has to be bought instead. | 
**Features** | [**[]TenantQuotaFeatureDto**](TenantQuotaFeatureDto.md) | The features the quota switches on, each with the limit it grants and, on the quota the portal is  actually on, how much of that limit is already used. A feature that is absent is off, so the list is the  whole truth about what the quota includes. | 
**UsersQuota** | Pointer to [**TenantEntityQuotaSettings**](TenantEntityQuotaSettings.md) | The per-member storage allowance an administrator has set on top of the quota, and whether it is applied  at all. It describes the live portal rather than this quota, so every entry of a catalogue listing repeats  the same values, and it is empty unless the portal is a server installation or its plan includes  statistics. | [optional] 
**RoomsQuota** | Pointer to [**TenantEntityQuotaSettings**](TenantEntityQuotaSettings.md) | The same kind of per-room storage override, filled in and read the same way as `usersQuota`. | [optional] 
**AiAgentsQuota** | Pointer to [**TenantEntityQuotaSettings**](TenantEntityQuotaSettings.md) | The same kind of per-agent storage override for AI agents, filled in and read the same way as  `usersQuota`. | [optional] 
**TenantCustomQuota** | Pointer to [**TenantQuotaSettings**](TenantQuotaSettings.md) | The storage allowance an administrator has set for the portal as a whole, which caps it below what the  quota grants. Filled in under the same conditions as `usersQuota`. | [optional] 
**DueDate** | Pointer to **time.Time** | When the quota runs out, in UTC. It is empty on a quota from the catalogue, which has no date until it is  bought, and on a quota that never expires. | [optional] 
**InnerServices** | Pointer to [**[]WalletServiceDto**](WalletServiceDto.md) | The variants of this service that are folded into it, so a client renders one card per group instead of  one per variant. It is empty when the service has no variants, and always empty in the answer of  `GET api/2.0/portal/payment/walletservice`, which looks one service up on its own. | [optional] 
**ServiceName** | Pointer to **NullableString** | The stable key of the service, which is what the wallet operations take as their `service` argument and  what the usage reports key their entries by. | [optional] 

## Methods

### NewWalletServiceDto

`func NewWalletServiceDto(id int32, price PriceDto, nonProfit bool, free bool, trial bool, features []TenantQuotaFeatureDto, ) *WalletServiceDto`

NewWalletServiceDto instantiates a new WalletServiceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletServiceDtoWithDefaults

`func NewWalletServiceDtoWithDefaults() *WalletServiceDto`

NewWalletServiceDtoWithDefaults instantiates a new WalletServiceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WalletServiceDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WalletServiceDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WalletServiceDto) SetId(v int32)`

SetId sets Id field to given value.


### GetTitle

`func (o *WalletServiceDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WalletServiceDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WalletServiceDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WalletServiceDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetPrice

`func (o *WalletServiceDto) GetPrice() PriceDto`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *WalletServiceDto) GetPriceOk() (*PriceDto, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *WalletServiceDto) SetPrice(v PriceDto)`

SetPrice sets Price field to given value.


### GetNonProfit

`func (o *WalletServiceDto) GetNonProfit() bool`

GetNonProfit returns the NonProfit field if non-nil, zero value otherwise.

### GetNonProfitOk

`func (o *WalletServiceDto) GetNonProfitOk() (*bool, bool)`

GetNonProfitOk returns a tuple with the NonProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonProfit

`func (o *WalletServiceDto) SetNonProfit(v bool)`

SetNonProfit sets NonProfit field to given value.


### GetFree

`func (o *WalletServiceDto) GetFree() bool`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *WalletServiceDto) GetFreeOk() (*bool, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *WalletServiceDto) SetFree(v bool)`

SetFree sets Free field to given value.


### GetTrial

`func (o *WalletServiceDto) GetTrial() bool`

GetTrial returns the Trial field if non-nil, zero value otherwise.

### GetTrialOk

`func (o *WalletServiceDto) GetTrialOk() (*bool, bool)`

GetTrialOk returns a tuple with the Trial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrial

`func (o *WalletServiceDto) SetTrial(v bool)`

SetTrial sets Trial field to given value.


### GetFeatures

`func (o *WalletServiceDto) GetFeatures() []TenantQuotaFeatureDto`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *WalletServiceDto) GetFeaturesOk() (*[]TenantQuotaFeatureDto, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *WalletServiceDto) SetFeatures(v []TenantQuotaFeatureDto)`

SetFeatures sets Features field to given value.


### GetUsersQuota

`func (o *WalletServiceDto) GetUsersQuota() TenantEntityQuotaSettings`

GetUsersQuota returns the UsersQuota field if non-nil, zero value otherwise.

### GetUsersQuotaOk

`func (o *WalletServiceDto) GetUsersQuotaOk() (*TenantEntityQuotaSettings, bool)`

GetUsersQuotaOk returns a tuple with the UsersQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersQuota

`func (o *WalletServiceDto) SetUsersQuota(v TenantEntityQuotaSettings)`

SetUsersQuota sets UsersQuota field to given value.

### HasUsersQuota

`func (o *WalletServiceDto) HasUsersQuota() bool`

HasUsersQuota returns a boolean if a field has been set.

### GetRoomsQuota

`func (o *WalletServiceDto) GetRoomsQuota() TenantEntityQuotaSettings`

GetRoomsQuota returns the RoomsQuota field if non-nil, zero value otherwise.

### GetRoomsQuotaOk

`func (o *WalletServiceDto) GetRoomsQuotaOk() (*TenantEntityQuotaSettings, bool)`

GetRoomsQuotaOk returns a tuple with the RoomsQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomsQuota

`func (o *WalletServiceDto) SetRoomsQuota(v TenantEntityQuotaSettings)`

SetRoomsQuota sets RoomsQuota field to given value.

### HasRoomsQuota

`func (o *WalletServiceDto) HasRoomsQuota() bool`

HasRoomsQuota returns a boolean if a field has been set.

### GetAiAgentsQuota

`func (o *WalletServiceDto) GetAiAgentsQuota() TenantEntityQuotaSettings`

GetAiAgentsQuota returns the AiAgentsQuota field if non-nil, zero value otherwise.

### GetAiAgentsQuotaOk

`func (o *WalletServiceDto) GetAiAgentsQuotaOk() (*TenantEntityQuotaSettings, bool)`

GetAiAgentsQuotaOk returns a tuple with the AiAgentsQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiAgentsQuota

`func (o *WalletServiceDto) SetAiAgentsQuota(v TenantEntityQuotaSettings)`

SetAiAgentsQuota sets AiAgentsQuota field to given value.

### HasAiAgentsQuota

`func (o *WalletServiceDto) HasAiAgentsQuota() bool`

HasAiAgentsQuota returns a boolean if a field has been set.

### GetTenantCustomQuota

`func (o *WalletServiceDto) GetTenantCustomQuota() TenantQuotaSettings`

GetTenantCustomQuota returns the TenantCustomQuota field if non-nil, zero value otherwise.

### GetTenantCustomQuotaOk

`func (o *WalletServiceDto) GetTenantCustomQuotaOk() (*TenantQuotaSettings, bool)`

GetTenantCustomQuotaOk returns a tuple with the TenantCustomQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantCustomQuota

`func (o *WalletServiceDto) SetTenantCustomQuota(v TenantQuotaSettings)`

SetTenantCustomQuota sets TenantCustomQuota field to given value.

### HasTenantCustomQuota

`func (o *WalletServiceDto) HasTenantCustomQuota() bool`

HasTenantCustomQuota returns a boolean if a field has been set.

### GetDueDate

`func (o *WalletServiceDto) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *WalletServiceDto) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *WalletServiceDto) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *WalletServiceDto) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetInnerServices

`func (o *WalletServiceDto) GetInnerServices() []WalletServiceDto`

GetInnerServices returns the InnerServices field if non-nil, zero value otherwise.

### GetInnerServicesOk

`func (o *WalletServiceDto) GetInnerServicesOk() (*[]WalletServiceDto, bool)`

GetInnerServicesOk returns a tuple with the InnerServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerServices

`func (o *WalletServiceDto) SetInnerServices(v []WalletServiceDto)`

SetInnerServices sets InnerServices field to given value.

### HasInnerServices

`func (o *WalletServiceDto) HasInnerServices() bool`

HasInnerServices returns a boolean if a field has been set.

### SetInnerServicesNil

`func (o *WalletServiceDto) SetInnerServicesNil(b bool)`

 SetInnerServicesNil sets the value for InnerServices to be an explicit nil

### UnsetInnerServices
`func (o *WalletServiceDto) UnsetInnerServices()`

UnsetInnerServices ensures that no value is present for InnerServices, not even an explicit nil
### GetServiceName

`func (o *WalletServiceDto) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *WalletServiceDto) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *WalletServiceDto) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *WalletServiceDto) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### SetServiceNameNil

`func (o *WalletServiceDto) SetServiceNameNil(b bool)`

 SetServiceNameNil sets the value for ServiceName to be an explicit nil

### UnsetServiceName
`func (o *WalletServiceDto) UnsetServiceName()`

UnsetServiceName ensures that no value is present for ServiceName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


