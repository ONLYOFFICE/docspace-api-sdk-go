# AutoCleanUpData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAutoCleanUp** | Pointer to **bool** | Specifies whether to permanently delete files in the Trash folder. | [optional] 
**Gap** | Pointer to [**DateToAutoCleanUp**](DateToAutoCleanUp.md) |  | [optional] 

## Methods

### NewAutoCleanUpData

`func NewAutoCleanUpData() *AutoCleanUpData`

NewAutoCleanUpData instantiates a new AutoCleanUpData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoCleanUpDataWithDefaults

`func NewAutoCleanUpDataWithDefaults() *AutoCleanUpData`

NewAutoCleanUpDataWithDefaults instantiates a new AutoCleanUpData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAutoCleanUp

`func (o *AutoCleanUpData) GetIsAutoCleanUp() bool`

GetIsAutoCleanUp returns the IsAutoCleanUp field if non-nil, zero value otherwise.

### GetIsAutoCleanUpOk

`func (o *AutoCleanUpData) GetIsAutoCleanUpOk() (*bool, bool)`

GetIsAutoCleanUpOk returns a tuple with the IsAutoCleanUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoCleanUp

`func (o *AutoCleanUpData) SetIsAutoCleanUp(v bool)`

SetIsAutoCleanUp sets IsAutoCleanUp field to given value.

### HasIsAutoCleanUp

`func (o *AutoCleanUpData) HasIsAutoCleanUp() bool`

HasIsAutoCleanUp returns a boolean if a field has been set.

### GetGap

`func (o *AutoCleanUpData) GetGap() DateToAutoCleanUp`

GetGap returns the Gap field if non-nil, zero value otherwise.

### GetGapOk

`func (o *AutoCleanUpData) GetGapOk() (*DateToAutoCleanUp, bool)`

GetGapOk returns a tuple with the Gap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGap

`func (o *AutoCleanUpData) SetGap(v DateToAutoCleanUp)`

SetGap sets Gap field to given value.

### HasGap

`func (o *AutoCleanUpData) HasGap() bool`

HasGap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


