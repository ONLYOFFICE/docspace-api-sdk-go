# QuotaDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The identifier of the quota, which is what the tariff reports as a quota `id` and what a purchase names.  A negative value belongs to a built-in quota rather than one on the price list. | 
**Title** | Pointer to **NullableString** | The quota name in the portal language, for printing rather than matching. It is empty when this build  ships no wording for the quota, which is normal for a quota that is not on the public price list. | [optional] 
**Price** | [**PriceDto**](PriceDto.md) | What the quota costs, in the currency resolved for the request. Its `value` is empty for a quota that is  not sold for money, which is what `free`, `trial` and `nonProfit` describe. | 
**NonProfit** | **bool** | Whether this is the non-profit quota, which is granted rather than bought. A portal on it cannot buy any  other plan, so a catalogue asked for plans returns this one alone. | 
**Free** | **bool** | Whether this is the free quota a portal falls back to when nothing is paid for. It has no end date and  the tightest limits of any quota. | 
**Trial** | **bool** | Whether this is the trial quota, which grants the paid limits for a while and then expires. A trial is not  extended by paying - a plan has to be bought instead. | 
**Features** | [**[]TenantQuotaFeatureDto**](TenantQuotaFeatureDto.md) | The features the quota includes, each with the limit it grants and, on the portal's own quota, how much of it  is used. A flag that is off is listed with `value` false; an absent feature is one this answer does not report. | 
**UsersQuota** | Pointer to [**EntityQuotaDto**](EntityQuotaDto.md) | The per-member storage allowance an administrator has set on top of the quota, and whether it is applied  at all. It describes the live portal rather than this quota, so every entry of a catalogue listing repeats  the same values. It is filled in only on a server installation or when the portal's plan includes statistics. | [optional] 
**RoomsQuota** | Pointer to [**EntityQuotaDto**](EntityQuotaDto.md) | The same kind of per-room storage override, filled in and read the same way as `usersQuota`. | [optional] 
**AiAgentsQuota** | Pointer to [**EntityQuotaDto**](EntityQuotaDto.md) | The same kind of per-agent storage override for AI agents, filled in and read the same way as  `usersQuota`. | [optional] 
**TenantCustomQuota** | Pointer to [**TenantQuotaSettingsDto**](TenantQuotaSettingsDto.md) | The storage allowance an administrator has set for the portal as a whole, which caps it below what the  quota grants. Filled in under the same conditions as `usersQuota`. | [optional] 
**DueDate** | Pointer to **NullableTime** | When the quota runs out, in UTC. It is empty on a quota from the catalogue, which has no date until it is  bought, and on a quota that never expires. | [optional] 

## Methods

### NewQuotaDto

`func NewQuotaDto(id int32, price PriceDto, nonProfit bool, free bool, trial bool, features []TenantQuotaFeatureDto, ) *QuotaDto`

NewQuotaDto instantiates a new QuotaDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotaDtoWithDefaults

`func NewQuotaDtoWithDefaults() *QuotaDto`

NewQuotaDtoWithDefaults instantiates a new QuotaDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuotaDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuotaDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuotaDto) SetId(v int32)`

SetId sets Id field to given value.


### GetTitle

`func (o *QuotaDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *QuotaDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *QuotaDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *QuotaDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *QuotaDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *QuotaDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetPrice

`func (o *QuotaDto) GetPrice() PriceDto`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *QuotaDto) GetPriceOk() (*PriceDto, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *QuotaDto) SetPrice(v PriceDto)`

SetPrice sets Price field to given value.


### GetNonProfit

`func (o *QuotaDto) GetNonProfit() bool`

GetNonProfit returns the NonProfit field if non-nil, zero value otherwise.

### GetNonProfitOk

`func (o *QuotaDto) GetNonProfitOk() (*bool, bool)`

GetNonProfitOk returns a tuple with the NonProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonProfit

`func (o *QuotaDto) SetNonProfit(v bool)`

SetNonProfit sets NonProfit field to given value.


### GetFree

`func (o *QuotaDto) GetFree() bool`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *QuotaDto) GetFreeOk() (*bool, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *QuotaDto) SetFree(v bool)`

SetFree sets Free field to given value.


### GetTrial

`func (o *QuotaDto) GetTrial() bool`

GetTrial returns the Trial field if non-nil, zero value otherwise.

### GetTrialOk

`func (o *QuotaDto) GetTrialOk() (*bool, bool)`

GetTrialOk returns a tuple with the Trial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrial

`func (o *QuotaDto) SetTrial(v bool)`

SetTrial sets Trial field to given value.


### GetFeatures

`func (o *QuotaDto) GetFeatures() []TenantQuotaFeatureDto`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *QuotaDto) GetFeaturesOk() (*[]TenantQuotaFeatureDto, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *QuotaDto) SetFeatures(v []TenantQuotaFeatureDto)`

SetFeatures sets Features field to given value.


### SetFeaturesNil

`func (o *QuotaDto) SetFeaturesNil(b bool)`

 SetFeaturesNil sets the value for Features to be an explicit nil

### UnsetFeatures
`func (o *QuotaDto) UnsetFeatures()`

UnsetFeatures ensures that no value is present for Features, not even an explicit nil
### GetUsersQuota

`func (o *QuotaDto) GetUsersQuota() EntityQuotaDto`

GetUsersQuota returns the UsersQuota field if non-nil, zero value otherwise.

### GetUsersQuotaOk

`func (o *QuotaDto) GetUsersQuotaOk() (*EntityQuotaDto, bool)`

GetUsersQuotaOk returns a tuple with the UsersQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersQuota

`func (o *QuotaDto) SetUsersQuota(v EntityQuotaDto)`

SetUsersQuota sets UsersQuota field to given value.

### HasUsersQuota

`func (o *QuotaDto) HasUsersQuota() bool`

HasUsersQuota returns a boolean if a field has been set.

### GetRoomsQuota

`func (o *QuotaDto) GetRoomsQuota() EntityQuotaDto`

GetRoomsQuota returns the RoomsQuota field if non-nil, zero value otherwise.

### GetRoomsQuotaOk

`func (o *QuotaDto) GetRoomsQuotaOk() (*EntityQuotaDto, bool)`

GetRoomsQuotaOk returns a tuple with the RoomsQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomsQuota

`func (o *QuotaDto) SetRoomsQuota(v EntityQuotaDto)`

SetRoomsQuota sets RoomsQuota field to given value.

### HasRoomsQuota

`func (o *QuotaDto) HasRoomsQuota() bool`

HasRoomsQuota returns a boolean if a field has been set.

### GetAiAgentsQuota

`func (o *QuotaDto) GetAiAgentsQuota() EntityQuotaDto`

GetAiAgentsQuota returns the AiAgentsQuota field if non-nil, zero value otherwise.

### GetAiAgentsQuotaOk

`func (o *QuotaDto) GetAiAgentsQuotaOk() (*EntityQuotaDto, bool)`

GetAiAgentsQuotaOk returns a tuple with the AiAgentsQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiAgentsQuota

`func (o *QuotaDto) SetAiAgentsQuota(v EntityQuotaDto)`

SetAiAgentsQuota sets AiAgentsQuota field to given value.

### HasAiAgentsQuota

`func (o *QuotaDto) HasAiAgentsQuota() bool`

HasAiAgentsQuota returns a boolean if a field has been set.

### GetTenantCustomQuota

`func (o *QuotaDto) GetTenantCustomQuota() TenantQuotaSettingsDto`

GetTenantCustomQuota returns the TenantCustomQuota field if non-nil, zero value otherwise.

### GetTenantCustomQuotaOk

`func (o *QuotaDto) GetTenantCustomQuotaOk() (*TenantQuotaSettingsDto, bool)`

GetTenantCustomQuotaOk returns a tuple with the TenantCustomQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantCustomQuota

`func (o *QuotaDto) SetTenantCustomQuota(v TenantQuotaSettingsDto)`

SetTenantCustomQuota sets TenantCustomQuota field to given value.

### HasTenantCustomQuota

`func (o *QuotaDto) HasTenantCustomQuota() bool`

HasTenantCustomQuota returns a boolean if a field has been set.

### GetDueDate

`func (o *QuotaDto) GetDueDate() time.Time`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *QuotaDto) GetDueDateOk() (*time.Time, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *QuotaDto) SetDueDate(v time.Time)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *QuotaDto) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### SetDueDateNil

`func (o *QuotaDto) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *QuotaDto) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


