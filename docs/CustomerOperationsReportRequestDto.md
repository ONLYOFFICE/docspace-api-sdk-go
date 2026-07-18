# CustomerOperationsReportRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceName** | Pointer to **NullableString** | The service name. | [optional] 
**StartDate** | Pointer to **NullableTime** | The report start date. | [optional] 
**EndDate** | Pointer to **NullableTime** | The report end date. | [optional] 
**ParticipantName** | Pointer to **NullableString** | The participant name. | [optional] 
**Credit** | Pointer to **NullableBool** | Specifies whether to include credit operations in the report. | [optional] 
**Debit** | Pointer to **NullableBool** | Specifies whether to include debit operations in the report. | [optional] 
**Type** | Pointer to [**OperationType**](OperationType.md) |  | [optional] 
**Status** | Pointer to [**OperationStatus**](OperationStatus.md) |  | [optional] 
**OrderBy** | Pointer to **NullableString** | The field to order by. | [optional] 
**OrderType** | Pointer to [**OperationOrderType**](OperationOrderType.md) |  | [optional] 

## Methods

### NewCustomerOperationsReportRequestDto

`func NewCustomerOperationsReportRequestDto() *CustomerOperationsReportRequestDto`

NewCustomerOperationsReportRequestDto instantiates a new CustomerOperationsReportRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerOperationsReportRequestDtoWithDefaults

`func NewCustomerOperationsReportRequestDtoWithDefaults() *CustomerOperationsReportRequestDto`

NewCustomerOperationsReportRequestDtoWithDefaults instantiates a new CustomerOperationsReportRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceName

`func (o *CustomerOperationsReportRequestDto) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *CustomerOperationsReportRequestDto) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *CustomerOperationsReportRequestDto) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *CustomerOperationsReportRequestDto) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### SetServiceNameNil

`func (o *CustomerOperationsReportRequestDto) SetServiceNameNil(b bool)`

 SetServiceNameNil sets the value for ServiceName to be an explicit nil

### UnsetServiceName
`func (o *CustomerOperationsReportRequestDto) UnsetServiceName()`

UnsetServiceName ensures that no value is present for ServiceName, not even an explicit nil
### GetStartDate

`func (o *CustomerOperationsReportRequestDto) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CustomerOperationsReportRequestDto) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CustomerOperationsReportRequestDto) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CustomerOperationsReportRequestDto) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *CustomerOperationsReportRequestDto) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *CustomerOperationsReportRequestDto) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *CustomerOperationsReportRequestDto) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CustomerOperationsReportRequestDto) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CustomerOperationsReportRequestDto) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CustomerOperationsReportRequestDto) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *CustomerOperationsReportRequestDto) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *CustomerOperationsReportRequestDto) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetParticipantName

`func (o *CustomerOperationsReportRequestDto) GetParticipantName() string`

GetParticipantName returns the ParticipantName field if non-nil, zero value otherwise.

### GetParticipantNameOk

`func (o *CustomerOperationsReportRequestDto) GetParticipantNameOk() (*string, bool)`

GetParticipantNameOk returns a tuple with the ParticipantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantName

`func (o *CustomerOperationsReportRequestDto) SetParticipantName(v string)`

SetParticipantName sets ParticipantName field to given value.

### HasParticipantName

`func (o *CustomerOperationsReportRequestDto) HasParticipantName() bool`

HasParticipantName returns a boolean if a field has been set.

### SetParticipantNameNil

`func (o *CustomerOperationsReportRequestDto) SetParticipantNameNil(b bool)`

 SetParticipantNameNil sets the value for ParticipantName to be an explicit nil

### UnsetParticipantName
`func (o *CustomerOperationsReportRequestDto) UnsetParticipantName()`

UnsetParticipantName ensures that no value is present for ParticipantName, not even an explicit nil
### GetCredit

`func (o *CustomerOperationsReportRequestDto) GetCredit() bool`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *CustomerOperationsReportRequestDto) GetCreditOk() (*bool, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *CustomerOperationsReportRequestDto) SetCredit(v bool)`

SetCredit sets Credit field to given value.

### HasCredit

`func (o *CustomerOperationsReportRequestDto) HasCredit() bool`

HasCredit returns a boolean if a field has been set.

### SetCreditNil

`func (o *CustomerOperationsReportRequestDto) SetCreditNil(b bool)`

 SetCreditNil sets the value for Credit to be an explicit nil

### UnsetCredit
`func (o *CustomerOperationsReportRequestDto) UnsetCredit()`

UnsetCredit ensures that no value is present for Credit, not even an explicit nil
### GetDebit

`func (o *CustomerOperationsReportRequestDto) GetDebit() bool`

GetDebit returns the Debit field if non-nil, zero value otherwise.

### GetDebitOk

`func (o *CustomerOperationsReportRequestDto) GetDebitOk() (*bool, bool)`

GetDebitOk returns a tuple with the Debit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebit

`func (o *CustomerOperationsReportRequestDto) SetDebit(v bool)`

SetDebit sets Debit field to given value.

### HasDebit

`func (o *CustomerOperationsReportRequestDto) HasDebit() bool`

HasDebit returns a boolean if a field has been set.

### SetDebitNil

`func (o *CustomerOperationsReportRequestDto) SetDebitNil(b bool)`

 SetDebitNil sets the value for Debit to be an explicit nil

### UnsetDebit
`func (o *CustomerOperationsReportRequestDto) UnsetDebit()`

UnsetDebit ensures that no value is present for Debit, not even an explicit nil
### GetType

`func (o *CustomerOperationsReportRequestDto) GetType() OperationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomerOperationsReportRequestDto) GetTypeOk() (*OperationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomerOperationsReportRequestDto) SetType(v OperationType)`

SetType sets Type field to given value.

### HasType

`func (o *CustomerOperationsReportRequestDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *CustomerOperationsReportRequestDto) GetStatus() OperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CustomerOperationsReportRequestDto) GetStatusOk() (*OperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CustomerOperationsReportRequestDto) SetStatus(v OperationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CustomerOperationsReportRequestDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOrderBy

`func (o *CustomerOperationsReportRequestDto) GetOrderBy() string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *CustomerOperationsReportRequestDto) GetOrderByOk() (*string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *CustomerOperationsReportRequestDto) SetOrderBy(v string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *CustomerOperationsReportRequestDto) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### SetOrderByNil

`func (o *CustomerOperationsReportRequestDto) SetOrderByNil(b bool)`

 SetOrderByNil sets the value for OrderBy to be an explicit nil

### UnsetOrderBy
`func (o *CustomerOperationsReportRequestDto) UnsetOrderBy()`

UnsetOrderBy ensures that no value is present for OrderBy, not even an explicit nil
### GetOrderType

`func (o *CustomerOperationsReportRequestDto) GetOrderType() OperationOrderType`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *CustomerOperationsReportRequestDto) GetOrderTypeOk() (*OperationOrderType, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *CustomerOperationsReportRequestDto) SetOrderType(v OperationOrderType)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *CustomerOperationsReportRequestDto) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


