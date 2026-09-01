# WalletServiceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The quota ID. | 
**Title** | Pointer to **string** | The quota title. | [optional] 
**Price** | [**PriceDto**](PriceDto.md) | The price parameters. | 
**NonProfit** | **bool** | Specifies if the quota is nonprofit or not. | 
**Free** | **bool** | Specifies if the quota is free or not. | 
**Trial** | **bool** | Specifies if the quota is trial or not. | 
**Features** | [**[]TenantQuotaFeatureDto**](TenantQuotaFeatureDto.md) | The list of tenant quota features. | 
**UsersQuota** | Pointer to [**TenantEntityQuotaSettings**](TenantEntityQuotaSettings.md) | The user quota. | [optional] 
**RoomsQuota** | Pointer to [**TenantEntityQuotaSettings**](TenantEntityQuotaSettings.md) | The room quota. | [optional] 
**AiAgentsQuota** | Pointer to [**TenantEntityQuotaSettings**](TenantEntityQuotaSettings.md) | The ai agent quota. | [optional] 
**TenantCustomQuota** | Pointer to [**TenantQuotaSettings**](TenantQuotaSettings.md) | The tenant custom quota. | [optional] 
**DueDate** | Pointer to **time.Time** | The due date. | [optional] 
**InnerServices** | Pointer to [**[]WalletServiceDto**](WalletServiceDto.md) | The list of inner services. | [optional] 
**ServiceName** | Pointer to **NullableString** | The service name. | [optional] 

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


