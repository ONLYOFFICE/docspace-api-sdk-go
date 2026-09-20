# HistoryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The identifier of the record, which tells two records of the same action apart and stays stable as long as the  portal keeps the log. | 
**Action** | [**HistoryAction**](HistoryAction.md) | What happened - the kind of event the record stands for, such as a file being uploaded, renamed, moved or  shared - with the key a client can key its own wording off. | 
**Initiator** | [**EmployeeDto**](EmployeeDto.md) | Who caused the event. For an event caused by a visitor following an external link only the name they gave is  filled in, the account fields staying empty. | 
**Date** | [**ApiDateTime**](ApiDateTime.md) | When the event happened, written with the offset of the portal's time zone. | 
**Data** | Pointer to [**HistoryData**](HistoryData.md) | The history data. Absent for actions that carry no payload of their own - changing a room's  logo, icon colour or cover, whose interpreter returns no data (see  `RoomLogoChangedInterpreter`). It used to be declared required, which put it in the  OpenAPI document's required list while the null-dropping serializer left it out of the  response, so a generated client threw on any history page holding one of those entries. | [optional] 
**Related** | Pointer to [**[]HistoryDto**](HistoryDto.md) | The records folded into this one because they belong to the same action, the separate files of one upload for  instance. It is empty when the record stands alone, and the records inside it carry no further nesting. | [optional] 

## Methods

### NewHistoryDto

`func NewHistoryDto(id int32, action HistoryAction, initiator EmployeeDto, date ApiDateTime, ) *HistoryDto`

NewHistoryDto instantiates a new HistoryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryDtoWithDefaults

`func NewHistoryDtoWithDefaults() *HistoryDto`

NewHistoryDtoWithDefaults instantiates a new HistoryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HistoryDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HistoryDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HistoryDto) SetId(v int32)`

SetId sets Id field to given value.


### GetAction

`func (o *HistoryDto) GetAction() HistoryAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *HistoryDto) GetActionOk() (*HistoryAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *HistoryDto) SetAction(v HistoryAction)`

SetAction sets Action field to given value.


### GetInitiator

`func (o *HistoryDto) GetInitiator() EmployeeDto`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *HistoryDto) GetInitiatorOk() (*EmployeeDto, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *HistoryDto) SetInitiator(v EmployeeDto)`

SetInitiator sets Initiator field to given value.


### GetDate

`func (o *HistoryDto) GetDate() ApiDateTime`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *HistoryDto) GetDateOk() (*ApiDateTime, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *HistoryDto) SetDate(v ApiDateTime)`

SetDate sets Date field to given value.


### GetData

`func (o *HistoryDto) GetData() HistoryData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HistoryDto) GetDataOk() (*HistoryData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HistoryDto) SetData(v HistoryData)`

SetData sets Data field to given value.

### HasData

`func (o *HistoryDto) HasData() bool`

HasData returns a boolean if a field has been set.

### GetRelated

`func (o *HistoryDto) GetRelated() []HistoryDto`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *HistoryDto) GetRelatedOk() (*[]HistoryDto, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *HistoryDto) SetRelated(v []HistoryDto)`

SetRelated sets Related field to given value.

### HasRelated

`func (o *HistoryDto) HasRelated() bool`

HasRelated returns a boolean if a field has been set.

### SetRelatedNil

`func (o *HistoryDto) SetRelatedNil(b bool)`

 SetRelatedNil sets the value for Related to be an explicit nil

### UnsetRelated
`func (o *HistoryDto) UnsetRelated()`

UnsetRelated ensures that no value is present for Related, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


