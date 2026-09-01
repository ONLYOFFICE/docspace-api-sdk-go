# QuotaDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The quota ID. | 
**Title** | Pointer to **NullableString** | The quota title. | [optional] 
**Price** | [**PriceDto**](PriceDto.md) | The price parameters. | 
**NonProfit** | **bool** | Specifies if the quota is nonprofit or not. | 
**Free** | **bool** | Specifies if the quota is free or not. | 
**Trial** | **bool** | Specifies if the quota is trial or not. | 
**Features** | [**[]TenantQuotaFeatureDto**](TenantQuotaFeatureDto.md) | The list of tenant quota features. | 
**UsersQuota** | Pointer to [**TenantEntityQuotaSettings**](TenantEntityQuotaSettings.md) | The user quota. | [optional] 
**RoomsQuota** | Pointer to [**TenantEntityQuotaSettings**](TenantEntityQuotaSettings.md) | The room quota. | [optional] 
**AiAgentsQuota** | Pointer to [**TenantEntityQuotaSettings**](TenantEntityQuotaSettings.md) | The ai agent quota. | [optional] 
**TenantCustomQuota** | Pointer to [**TenantQuotaSettings**](TenantQuotaSettings.md) | The tenant custom quota. | [optional] 
**DueDate** | Pointer to **NullableTime** | The due date. | [optional] 

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

`func (o *QuotaDto) GetUsersQuota() TenantEntityQuotaSettings`

GetUsersQuota returns the UsersQuota field if non-nil, zero value otherwise.

### GetUsersQuotaOk

`func (o *QuotaDto) GetUsersQuotaOk() (*TenantEntityQuotaSettings, bool)`

GetUsersQuotaOk returns a tuple with the UsersQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersQuota

`func (o *QuotaDto) SetUsersQuota(v TenantEntityQuotaSettings)`

SetUsersQuota sets UsersQuota field to given value.

### HasUsersQuota

`func (o *QuotaDto) HasUsersQuota() bool`

HasUsersQuota returns a boolean if a field has been set.

### GetRoomsQuota

`func (o *QuotaDto) GetRoomsQuota() TenantEntityQuotaSettings`

GetRoomsQuota returns the RoomsQuota field if non-nil, zero value otherwise.

### GetRoomsQuotaOk

`func (o *QuotaDto) GetRoomsQuotaOk() (*TenantEntityQuotaSettings, bool)`

GetRoomsQuotaOk returns a tuple with the RoomsQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomsQuota

`func (o *QuotaDto) SetRoomsQuota(v TenantEntityQuotaSettings)`

SetRoomsQuota sets RoomsQuota field to given value.

### HasRoomsQuota

`func (o *QuotaDto) HasRoomsQuota() bool`

HasRoomsQuota returns a boolean if a field has been set.

### GetAiAgentsQuota

`func (o *QuotaDto) GetAiAgentsQuota() TenantEntityQuotaSettings`

GetAiAgentsQuota returns the AiAgentsQuota field if non-nil, zero value otherwise.

### GetAiAgentsQuotaOk

`func (o *QuotaDto) GetAiAgentsQuotaOk() (*TenantEntityQuotaSettings, bool)`

GetAiAgentsQuotaOk returns a tuple with the AiAgentsQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiAgentsQuota

`func (o *QuotaDto) SetAiAgentsQuota(v TenantEntityQuotaSettings)`

SetAiAgentsQuota sets AiAgentsQuota field to given value.

### HasAiAgentsQuota

`func (o *QuotaDto) HasAiAgentsQuota() bool`

HasAiAgentsQuota returns a boolean if a field has been set.

### GetTenantCustomQuota

`func (o *QuotaDto) GetTenantCustomQuota() TenantQuotaSettings`

GetTenantCustomQuota returns the TenantCustomQuota field if non-nil, zero value otherwise.

### GetTenantCustomQuotaOk

`func (o *QuotaDto) GetTenantCustomQuotaOk() (*TenantQuotaSettings, bool)`

GetTenantCustomQuotaOk returns a tuple with the TenantCustomQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantCustomQuota

`func (o *QuotaDto) SetTenantCustomQuota(v TenantQuotaSettings)`

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


