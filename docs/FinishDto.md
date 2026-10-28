# FinishDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSendWelcomeEmail** | **bool** | Whether every imported account that has not been activated yet is mailed its activation link. Setting it  requires the finished job to still be in the queue, so the import must not have been cleared first; the  letters go out again on each call, and already active accounts are skipped either way. Setting it false ends  the import quietly and leaves inviting those people for later. | 

## Methods

### NewFinishDto

`func NewFinishDto(isSendWelcomeEmail bool, ) *FinishDto`

NewFinishDto instantiates a new FinishDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinishDtoWithDefaults

`func NewFinishDtoWithDefaults() *FinishDto`

NewFinishDtoWithDefaults instantiates a new FinishDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSendWelcomeEmail

`func (o *FinishDto) GetIsSendWelcomeEmail() bool`

GetIsSendWelcomeEmail returns the IsSendWelcomeEmail field if non-nil, zero value otherwise.

### GetIsSendWelcomeEmailOk

`func (o *FinishDto) GetIsSendWelcomeEmailOk() (*bool, bool)`

GetIsSendWelcomeEmailOk returns a tuple with the IsSendWelcomeEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSendWelcomeEmail

`func (o *FinishDto) SetIsSendWelcomeEmail(v bool)`

SetIsSendWelcomeEmail sets IsSendWelcomeEmail field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


