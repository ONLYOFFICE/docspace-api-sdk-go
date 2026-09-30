# OperationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to [**ApiDateTime**](ApiDateTime.md) | When the movement was booked, in the portal time zone - the same zone the `startDate` and `endDate`  filters are read in, so the two do line up here. | [optional] 
**Service** | Pointer to **NullableString** | The wallet service the movement belongs to, by its stable key. It is what the `serviceName` filter  matches on, and it is empty for a movement that belongs to no service, such as a top-up. | [optional] 
**Description** | Pointer to **NullableString** | A one-line summary of the movement in the portal language, already composed from the service and the  quantity - meant to be printed as it is rather than parsed. | [optional] 
**Details** | Pointer to **NullableString** | The longer explanation of the same movement, where the service recorded one. It is empty for a movement  that has nothing to add to `description`. | [optional] 
**ServiceUnit** | Pointer to **NullableString** | What `quantity` counts for this service, in the portal language. AI consumption is reported in tokens  here rather than in the AI credits the service is sold in. | [optional] 
**Quantity** | Pointer to **int32** | How many units the movement covers, in the unit named by `serviceUnit`. It is `0` for a movement that  moves money without consuming a service. | [optional] 
**Currency** | Pointer to **NullableString** | The currency `credit` and `debit` are expressed in, as a three-letter ISO 4217 code. It is the accounting  currency of the wallet, which need not be the currency the subscription is priced in. | [optional] 
**Credit** | Pointer to **float64** | The amount that went into the wallet. It is `0` on a movement that only took money out, so the pair of  `credit` and `debit` is what shows which way the money went; the `credit` and `debit` filters of the  operation select the two directions by exactly this. | [optional] 
**Debit** | Pointer to **float64** | The amount that was taken out of the wallet, `0` on a movement that put money in. | [optional] 
**ParticipantName** | Pointer to **NullableString** | Who caused the movement, as the billing service records them - an internal name, which is what the  `participantName` filter matches on. Show `participantDisplayName` instead. | [optional] 
**ParticipantDisplayName** | Pointer to **NullableString** | The same person as their portal display name. It falls back to `participantName` when the name belongs to  no portal account, so it is never empty while `participantName` is filled. | [optional] 
**SourceType** | Pointer to **NullableString** | What kind of thing an AI operation was run on - an agent, a file, a folder, a room or a form. It is empty  on any movement that is not an AI charge. | [optional] 
**SourceTitle** | Pointer to **NullableString** | The title that thing had when the operation ran, kept as recorded, so it does not follow a later rename.  Empty under the same conditions as `sourceType`. | [optional] 
**SourceId** | Pointer to **NullableString** | The identifier of that thing, to look it up in the module it belongs to. Empty under the same conditions  as `sourceType`. | [optional] 
**TokenUsage** | Pointer to [**OperationTokenUsage**](OperationTokenUsage.md) | The tokens an AI operation consumed, broken down by kind - prompt, completion, cache reads and writes,  reasoning, images. It is `null` on any movement that is not an AI charge, and on an AI charge the billing  service recorded without token counts. | [optional] 
**Type** | Pointer to [**OperationType**](OperationType.md) | What kind of movement this is - a payment, a charge, a refund, a correction. It is what the `type` filter  matches on, and `Unknown` covers a movement the billing service reported under a kind this build does not  recognise. | [optional] 

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

`func (o *OperationDto) GetDate() ApiDateTime`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *OperationDto) GetDateOk() (*ApiDateTime, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *OperationDto) SetDate(v ApiDateTime)`

SetDate sets Date field to given value.

### HasDate

`func (o *OperationDto) HasDate() bool`

HasDate returns a boolean if a field has been set.

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
### GetSourceType

`func (o *OperationDto) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *OperationDto) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *OperationDto) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *OperationDto) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### SetSourceTypeNil

`func (o *OperationDto) SetSourceTypeNil(b bool)`

 SetSourceTypeNil sets the value for SourceType to be an explicit nil

### UnsetSourceType
`func (o *OperationDto) UnsetSourceType()`

UnsetSourceType ensures that no value is present for SourceType, not even an explicit nil
### GetSourceTitle

`func (o *OperationDto) GetSourceTitle() string`

GetSourceTitle returns the SourceTitle field if non-nil, zero value otherwise.

### GetSourceTitleOk

`func (o *OperationDto) GetSourceTitleOk() (*string, bool)`

GetSourceTitleOk returns a tuple with the SourceTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTitle

`func (o *OperationDto) SetSourceTitle(v string)`

SetSourceTitle sets SourceTitle field to given value.

### HasSourceTitle

`func (o *OperationDto) HasSourceTitle() bool`

HasSourceTitle returns a boolean if a field has been set.

### SetSourceTitleNil

`func (o *OperationDto) SetSourceTitleNil(b bool)`

 SetSourceTitleNil sets the value for SourceTitle to be an explicit nil

### UnsetSourceTitle
`func (o *OperationDto) UnsetSourceTitle()`

UnsetSourceTitle ensures that no value is present for SourceTitle, not even an explicit nil
### GetSourceId

`func (o *OperationDto) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *OperationDto) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *OperationDto) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *OperationDto) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *OperationDto) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *OperationDto) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetTokenUsage

`func (o *OperationDto) GetTokenUsage() OperationTokenUsage`

GetTokenUsage returns the TokenUsage field if non-nil, zero value otherwise.

### GetTokenUsageOk

`func (o *OperationDto) GetTokenUsageOk() (*OperationTokenUsage, bool)`

GetTokenUsageOk returns a tuple with the TokenUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUsage

`func (o *OperationDto) SetTokenUsage(v OperationTokenUsage)`

SetTokenUsage sets TokenUsage field to given value.

### HasTokenUsage

`func (o *OperationDto) HasTokenUsage() bool`

HasTokenUsage returns a boolean if a field has been set.

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


