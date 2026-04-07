# ChangeHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **int32** | The file version of the change history. | 
**ContinueVersion** | Pointer to **bool** | Specifies whether to start a new version or continue revision of the change history. | [optional] 

## Methods

### NewChangeHistory

`func NewChangeHistory(version int32, ) *ChangeHistory`

NewChangeHistory instantiates a new ChangeHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeHistoryWithDefaults

`func NewChangeHistoryWithDefaults() *ChangeHistory`

NewChangeHistoryWithDefaults instantiates a new ChangeHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *ChangeHistory) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ChangeHistory) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ChangeHistory) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetContinueVersion

`func (o *ChangeHistory) GetContinueVersion() bool`

GetContinueVersion returns the ContinueVersion field if non-nil, zero value otherwise.

### GetContinueVersionOk

`func (o *ChangeHistory) GetContinueVersionOk() (*bool, bool)`

GetContinueVersionOk returns a tuple with the ContinueVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueVersion

`func (o *ChangeHistory) SetContinueVersion(v bool)`

SetContinueVersion sets ContinueVersion field to given value.

### HasContinueVersion

`func (o *ChangeHistory) HasContinueVersion() bool`

HasContinueVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


