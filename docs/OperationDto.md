# OperationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **NullableTime** | The date when the operation took place. | [optional] 
**Service** | Pointer to **NullableString** | The service related to the operation. | [optional] 
**Description** | Pointer to **NullableString** | The brief operation description. | [optional] 
**Details** | Pointer to **NullableString** | The detailed information about the operation. | [optional] 
**ServiceUnit** | Pointer to **NullableString** | The service unit. | [optional] 
**Quantity** | Pointer to **int32** | The quantity of the service used. | [optional] 
**Currency** | Pointer to **NullableString** | The three-character ISO 4217 currency symbol of the operation. | [optional] 
**Credit** | Pointer to **float64** | The credit amount of the operation. | [optional] 
**Debit** | Pointer to **float64** | The debit amount of the operation. | [optional] 
**ParticipantName** | Pointer to **NullableString** | The participant original name. | [optional] 
**ParticipantDisplayName** | Pointer to **NullableString** | The participant display name. | [optional] 
**AgentId** | Pointer to **NullableString** | AI Agent id. | [optional] 
**AgentTitle** | Pointer to **NullableString** | AI Agent name. | [optional] 
**Type** | Pointer to [**OperationType**](OperationType.md) | Type of the operation | [optional] 

## Methods

### NewOperationDto

`func NewOperationDto() *OperationDto`

NewOperationDto instantiates a new OperationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationDtoWithDefaults

`func NewOperationDtoWithDefaults() *OperationDto`

NewOperationDtoWithDefaults instantiates a new OperationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *OperationDto) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *OperationDto) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *OperationDto) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *OperationDto) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *OperationDto) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *OperationDto) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetService

`func (o *OperationDto) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *OperationDto) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *OperationDto) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *OperationDto) HasService() bool`

HasService returns a boolean if a field has been set.

### SetServiceNil

`func (o *OperationDto) SetServiceNil(b bool)`

 SetServiceNil sets the value for Service to be an explicit nil

### UnsetService
`func (o *OperationDto) UnsetService()`

UnsetService ensures that no value is present for Service, not even an explicit nil
### GetDescription

`func (o *OperationDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OperationDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OperationDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OperationDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *OperationDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *OperationDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDetails

`func (o *OperationDto) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *OperationDto) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *OperationDto) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *OperationDto) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *OperationDto) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *OperationDto) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetServiceUnit

`func (o *OperationDto) GetServiceUnit() string`

GetServiceUnit returns the ServiceUnit field if non-nil, zero value otherwise.

### GetServiceUnitOk

`func (o *OperationDto) GetServiceUnitOk() (*string, bool)`

GetServiceUnitOk returns a tuple with the ServiceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUnit

`func (o *OperationDto) SetServiceUnit(v string)`

SetServiceUnit sets ServiceUnit field to given value.

### HasServiceUnit

`func (o *OperationDto) HasServiceUnit() bool`

HasServiceUnit returns a boolean if a field has been set.

### SetServiceUnitNil

`func (o *OperationDto) SetServiceUnitNil(b bool)`

 SetServiceUnitNil sets the value for ServiceUnit to be an explicit nil

### UnsetServiceUnit
`func (o *OperationDto) UnsetServiceUnit()`

UnsetServiceUnit ensures that no value is present for ServiceUnit, not even an explicit nil
### GetQuantity

`func (o *OperationDto) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OperationDto) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OperationDto) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *OperationDto) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetCurrency

`func (o *OperationDto) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *OperationDto) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *OperationDto) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *OperationDto) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *OperationDto) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *OperationDto) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetCredit

`func (o *OperationDto) GetCredit() float64`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *OperationDto) GetCreditOk() (*float64, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *OperationDto) SetCredit(v float64)`

SetCredit sets Credit field to given value.

### HasCredit

`func (o *OperationDto) HasCredit() bool`

HasCredit returns a boolean if a field has been set.

### GetDebit

`func (o *OperationDto) GetDebit() float64`

GetDebit returns the Debit field if non-nil, zero value otherwise.

### GetDebitOk

`func (o *OperationDto) GetDebitOk() (*float64, bool)`

GetDebitOk returns a tuple with the Debit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebit

`func (o *OperationDto) SetDebit(v float64)`

SetDebit sets Debit field to given value.

### HasDebit

`func (o *OperationDto) HasDebit() bool`

HasDebit returns a boolean if a field has been set.

### GetParticipantName

`func (o *OperationDto) GetParticipantName() string`

GetParticipantName returns the ParticipantName field if non-nil, zero value otherwise.

### GetParticipantNameOk

`func (o *OperationDto) GetParticipantNameOk() (*string, bool)`

GetParticipantNameOk returns a tuple with the ParticipantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantName

`func (o *OperationDto) SetParticipantName(v string)`

SetParticipantName sets ParticipantName field to given value.

### HasParticipantName

`func (o *OperationDto) HasParticipantName() bool`

HasParticipantName returns a boolean if a field has been set.

### SetParticipantNameNil

`func (o *OperationDto) SetParticipantNameNil(b bool)`

 SetParticipantNameNil sets the value for ParticipantName to be an explicit nil

### UnsetParticipantName
`func (o *OperationDto) UnsetParticipantName()`

UnsetParticipantName ensures that no value is present for ParticipantName, not even an explicit nil
### GetParticipantDisplayName

`func (o *OperationDto) GetParticipantDisplayName() string`

GetParticipantDisplayName returns the ParticipantDisplayName field if non-nil, zero value otherwise.

### GetParticipantDisplayNameOk

`func (o *OperationDto) GetParticipantDisplayNameOk() (*string, bool)`

GetParticipantDisplayNameOk returns a tuple with the ParticipantDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantDisplayName

`func (o *OperationDto) SetParticipantDisplayName(v string)`

SetParticipantDisplayName sets ParticipantDisplayName field to given value.

### HasParticipantDisplayName

`func (o *OperationDto) HasParticipantDisplayName() bool`

HasParticipantDisplayName returns a boolean if a field has been set.

### SetParticipantDisplayNameNil

`func (o *OperationDto) SetParticipantDisplayNameNil(b bool)`

 SetParticipantDisplayNameNil sets the value for ParticipantDisplayName to be an explicit nil

### UnsetParticipantDisplayName
`func (o *OperationDto) UnsetParticipantDisplayName()`

UnsetParticipantDisplayName ensures that no value is present for ParticipantDisplayName, not even an explicit nil
### GetAgentId

`func (o *OperationDto) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *OperationDto) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *OperationDto) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *OperationDto) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### SetAgentIdNil

`func (o *OperationDto) SetAgentIdNil(b bool)`

 SetAgentIdNil sets the value for AgentId to be an explicit nil

### UnsetAgentId
`func (o *OperationDto) UnsetAgentId()`

UnsetAgentId ensures that no value is present for AgentId, not even an explicit nil
### GetAgentTitle

`func (o *OperationDto) GetAgentTitle() string`

GetAgentTitle returns the AgentTitle field if non-nil, zero value otherwise.

### GetAgentTitleOk

`func (o *OperationDto) GetAgentTitleOk() (*string, bool)`

GetAgentTitleOk returns a tuple with the AgentTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentTitle

`func (o *OperationDto) SetAgentTitle(v string)`

SetAgentTitle sets AgentTitle field to given value.

### HasAgentTitle

`func (o *OperationDto) HasAgentTitle() bool`

HasAgentTitle returns a boolean if a field has been set.

### SetAgentTitleNil

`func (o *OperationDto) SetAgentTitleNil(b bool)`

 SetAgentTitleNil sets the value for AgentTitle to be an explicit nil

### UnsetAgentTitle
`func (o *OperationDto) UnsetAgentTitle()`

UnsetAgentTitle ensures that no value is present for AgentTitle, not even an explicit nil
### GetType

`func (o *OperationDto) GetType() OperationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OperationDto) GetTypeOk() (*OperationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OperationDto) SetType(v OperationType)`

SetType sets Type field to given value.

### HasType

`func (o *OperationDto) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


