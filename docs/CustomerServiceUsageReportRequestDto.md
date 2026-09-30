# CustomerServiceUsageReportRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceName** | Pointer to **[]string** | The wallet services whose consumption is reported, named the way the billing catalogue names them -  `backup`, `ai-tools`, `ai-search`, `disk-storage`, `docscloud`. Take the values from the `serviceName` field  of `GET api/2.0/portal/payment/walletservices`; the match ignores case, a name this installation does not  sell fails the call with 404, and an omitted list reports every service. A bare string is accepted in place  of an array for backward compatibility. | [optional] 
**StartDate** | Pointer to **NullableTime** | The beginning of the reported period, inclusive. Read in the portal time zone rather than in UTC, and  defaults to the portal creation date. | [optional] 
**EndDate** | Pointer to **NullableTime** | The end of the reported period, inclusive. Read in the portal time zone rather than in UTC, and defaults to  the moment the call is made. | [optional] 
**ParticipantName** | Pointer to **NullableString** | The participant whose consumption is reported - the account the accounting service records as the consumer.  Consumption caused by a portal user carries that user ID here; surrounding whitespace is trimmed, and an  omitted value reports every participant. | [optional] 
**Status** | Pointer to [**OperationStatus**](OperationStatus.md) | The outcome to keep. Consumption that is still being settled is reported as pending and may change later,  while the other outcomes are final; every outcome is reported when this is omitted. | [optional] 
**Metadata** | Pointer to **map[string]string** | The usage annotations a wallet service records alongside its consumption, as the key and value pairs that  must all match for a record to be reported. The keys are chosen by the service that writes them, so read  them off the `metadata` of the records returned by `GET api/2.0/portal/payment/customer/usage` rather than  guessing; an omitted map reports every record. | [optional] 
**OrderBy** | Pointer to **NullableString** | The name of the field the per-service totals are sorted by, spelled as the accounting service names it, such  as `ServiceName` or `StartDate`. Surrounding whitespace is trimmed, and the accounting service applies its  own ordering when this is omitted. | [optional] 
**OrderType** | Pointer to [**OperationOrderType**](OperationOrderType.md) | The direction the field named in `orderBy` is sorted in. Newest or largest first is what the accounting  service does by default, so leaving this out sorts the same way as asking for descending explicitly. | [optional] 

## Methods

### NewCustomerServiceUsageReportRequestDto

`func NewCustomerServiceUsageReportRequestDto() *CustomerServiceUsageReportRequestDto`

NewCustomerServiceUsageReportRequestDto instantiates a new CustomerServiceUsageReportRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerServiceUsageReportRequestDtoWithDefaults

`func NewCustomerServiceUsageReportRequestDtoWithDefaults() *CustomerServiceUsageReportRequestDto`

NewCustomerServiceUsageReportRequestDtoWithDefaults instantiates a new CustomerServiceUsageReportRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceName

`func (o *CustomerServiceUsageReportRequestDto) GetServiceName() []string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *CustomerServiceUsageReportRequestDto) GetServiceNameOk() (*[]string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *CustomerServiceUsageReportRequestDto) SetServiceName(v []string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *CustomerServiceUsageReportRequestDto) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### SetServiceNameNil

`func (o *CustomerServiceUsageReportRequestDto) SetServiceNameNil(b bool)`

 SetServiceNameNil sets the value for ServiceName to be an explicit nil

### UnsetServiceName
`func (o *CustomerServiceUsageReportRequestDto) UnsetServiceName()`

UnsetServiceName ensures that no value is present for ServiceName, not even an explicit nil
### GetStartDate

`func (o *CustomerServiceUsageReportRequestDto) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CustomerServiceUsageReportRequestDto) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CustomerServiceUsageReportRequestDto) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CustomerServiceUsageReportRequestDto) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *CustomerServiceUsageReportRequestDto) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *CustomerServiceUsageReportRequestDto) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *CustomerServiceUsageReportRequestDto) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CustomerServiceUsageReportRequestDto) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CustomerServiceUsageReportRequestDto) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CustomerServiceUsageReportRequestDto) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *CustomerServiceUsageReportRequestDto) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *CustomerServiceUsageReportRequestDto) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetParticipantName

`func (o *CustomerServiceUsageReportRequestDto) GetParticipantName() string`

GetParticipantName returns the ParticipantName field if non-nil, zero value otherwise.

### GetParticipantNameOk

`func (o *CustomerServiceUsageReportRequestDto) GetParticipantNameOk() (*string, bool)`

GetParticipantNameOk returns a tuple with the ParticipantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantName

`func (o *CustomerServiceUsageReportRequestDto) SetParticipantName(v string)`

SetParticipantName sets ParticipantName field to given value.

### HasParticipantName

`func (o *CustomerServiceUsageReportRequestDto) HasParticipantName() bool`

HasParticipantName returns a boolean if a field has been set.

### SetParticipantNameNil

`func (o *CustomerServiceUsageReportRequestDto) SetParticipantNameNil(b bool)`

 SetParticipantNameNil sets the value for ParticipantName to be an explicit nil

### UnsetParticipantName
`func (o *CustomerServiceUsageReportRequestDto) UnsetParticipantName()`

UnsetParticipantName ensures that no value is present for ParticipantName, not even an explicit nil
### GetStatus

`func (o *CustomerServiceUsageReportRequestDto) GetStatus() OperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CustomerServiceUsageReportRequestDto) GetStatusOk() (*OperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CustomerServiceUsageReportRequestDto) SetStatus(v OperationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CustomerServiceUsageReportRequestDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMetadata

`func (o *CustomerServiceUsageReportRequestDto) GetMetadata() map[string]*string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CustomerServiceUsageReportRequestDto) GetMetadataOk() (*map[string]*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CustomerServiceUsageReportRequestDto) SetMetadata(v map[string]*string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CustomerServiceUsageReportRequestDto) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOrderBy

`func (o *CustomerServiceUsageReportRequestDto) GetOrderBy() string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *CustomerServiceUsageReportRequestDto) GetOrderByOk() (*string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *CustomerServiceUsageReportRequestDto) SetOrderBy(v string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *CustomerServiceUsageReportRequestDto) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### SetOrderByNil

`func (o *CustomerServiceUsageReportRequestDto) SetOrderByNil(b bool)`

 SetOrderByNil sets the value for OrderBy to be an explicit nil

### UnsetOrderBy
`func (o *CustomerServiceUsageReportRequestDto) UnsetOrderBy()`

UnsetOrderBy ensures that no value is present for OrderBy, not even an explicit nil
### GetOrderType

`func (o *CustomerServiceUsageReportRequestDto) GetOrderType() OperationOrderType`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *CustomerServiceUsageReportRequestDto) GetOrderTypeOk() (*OperationOrderType, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *CustomerServiceUsageReportRequestDto) SetOrderType(v OperationOrderType)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *CustomerServiceUsageReportRequestDto) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


