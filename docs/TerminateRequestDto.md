# TerminateRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | The user ID whose data is reassigned/removed. | 

## Methods

### NewTerminateRequestDto

`func NewTerminateRequestDto(userId string, ) *TerminateRequestDto`

NewTerminateRequestDto instantiates a new TerminateRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerminateRequestDtoWithDefaults

`func NewTerminateRequestDtoWithDefaults() *TerminateRequestDto`

NewTerminateRequestDtoWithDefaults instantiates a new TerminateRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *TerminateRequestDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *TerminateRequestDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *TerminateRequestDto) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


