# CustomFilterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | The state to reach: `true` turns the mode on, so that the sorting and filtering each person applies stays  visible to that person alone, and drops the others out of a running editing session; `false` turns it off and  makes filtering shared again. | [optional] 

## Methods

### NewCustomFilterRequest

`func NewCustomFilterRequest() *CustomFilterRequest`

NewCustomFilterRequest instantiates a new CustomFilterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFilterRequestWithDefaults

`func NewCustomFilterRequestWithDefaults() *CustomFilterRequest`

NewCustomFilterRequestWithDefaults instantiates a new CustomFilterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *CustomFilterRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustomFilterRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustomFilterRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustomFilterRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


