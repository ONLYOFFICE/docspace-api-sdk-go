# StartReassignRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromUserId** | **string** | The user ID whose data will be reassigned to another user. | 
**ToUserId** | **string** | The user ID to whom all the data will be reassigned. | 
**DeleteProfile** | Pointer to **bool** | Specifies whether to delete a profile when the data reassignment will be finished or not. | [optional] 

## Methods

### NewStartReassignRequestDto

`func NewStartReassignRequestDto(fromUserId string, toUserId string, ) *StartReassignRequestDto`

NewStartReassignRequestDto instantiates a new StartReassignRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartReassignRequestDtoWithDefaults

`func NewStartReassignRequestDtoWithDefaults() *StartReassignRequestDto`

NewStartReassignRequestDtoWithDefaults instantiates a new StartReassignRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromUserId

`func (o *StartReassignRequestDto) GetFromUserId() string`

GetFromUserId returns the FromUserId field if non-nil, zero value otherwise.

### GetFromUserIdOk

`func (o *StartReassignRequestDto) GetFromUserIdOk() (*string, bool)`

GetFromUserIdOk returns a tuple with the FromUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromUserId

`func (o *StartReassignRequestDto) SetFromUserId(v string)`

SetFromUserId sets FromUserId field to given value.


### GetToUserId

`func (o *StartReassignRequestDto) GetToUserId() string`

GetToUserId returns the ToUserId field if non-nil, zero value otherwise.

### GetToUserIdOk

`func (o *StartReassignRequestDto) GetToUserIdOk() (*string, bool)`

GetToUserIdOk returns a tuple with the ToUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToUserId

`func (o *StartReassignRequestDto) SetToUserId(v string)`

SetToUserId sets ToUserId field to given value.


### GetDeleteProfile

`func (o *StartReassignRequestDto) GetDeleteProfile() bool`

GetDeleteProfile returns the DeleteProfile field if non-nil, zero value otherwise.

### GetDeleteProfileOk

`func (o *StartReassignRequestDto) GetDeleteProfileOk() (*bool, bool)`

GetDeleteProfileOk returns a tuple with the DeleteProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProfile

`func (o *StartReassignRequestDto) SetDeleteProfile(v bool)`

SetDeleteProfile sets DeleteProfile field to given value.

### HasDeleteProfile

`func (o *StartReassignRequestDto) HasDeleteProfile() bool`

HasDeleteProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


