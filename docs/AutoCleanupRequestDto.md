# AutoCleanupRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Set** | Pointer to **bool** | Whether the caller's trash is cleared automatically: with true an item is removed for good once it has been in  the trash longer than the interval below, with false the portal removes nothing and waits for the trash to be  emptied by hand. | [optional] 
**Gap** | Pointer to [**DateToAutoCleanUp**](DateToAutoCleanUp.md) | How long an item may stay in the trash before it is removed for good. It is written from every request,  including one that switches clearing off, so send it together with the flag instead of expecting the stored  interval to be kept. | [optional] 

## Methods

### NewAutoCleanupRequestDto

`func NewAutoCleanupRequestDto() *AutoCleanupRequestDto`

NewAutoCleanupRequestDto instantiates a new AutoCleanupRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoCleanupRequestDtoWithDefaults

`func NewAutoCleanupRequestDtoWithDefaults() *AutoCleanupRequestDto`

NewAutoCleanupRequestDtoWithDefaults instantiates a new AutoCleanupRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSet

`func (o *AutoCleanupRequestDto) GetSet() bool`

GetSet returns the Set field if non-nil, zero value otherwise.

### GetSetOk

`func (o *AutoCleanupRequestDto) GetSetOk() (*bool, bool)`

GetSetOk returns a tuple with the Set field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSet

`func (o *AutoCleanupRequestDto) SetSet(v bool)`

SetSet sets Set field to given value.

### HasSet

`func (o *AutoCleanupRequestDto) HasSet() bool`

HasSet returns a boolean if a field has been set.

### GetGap

`func (o *AutoCleanupRequestDto) GetGap() DateToAutoCleanUp`

GetGap returns the Gap field if non-nil, zero value otherwise.

### GetGapOk

`func (o *AutoCleanupRequestDto) GetGapOk() (*DateToAutoCleanUp, bool)`

GetGapOk returns a tuple with the Gap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGap

`func (o *AutoCleanupRequestDto) SetGap(v DateToAutoCleanUp)`

SetGap sets Gap field to given value.

### HasGap

`func (o *AutoCleanupRequestDto) HasGap() bool`

HasGap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


