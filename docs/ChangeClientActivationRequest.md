# ChangeClientActivationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **bool** | Whether the client may obtain tokens from now on. Sending false leaves the registration and the already issued tokens in place but refuses new authorization requests; sending true allows them again. | 

## Methods

### NewChangeClientActivationRequest

`func NewChangeClientActivationRequest(status bool, ) *ChangeClientActivationRequest`

NewChangeClientActivationRequest instantiates a new ChangeClientActivationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeClientActivationRequestWithDefaults

`func NewChangeClientActivationRequestWithDefaults() *ChangeClientActivationRequest`

NewChangeClientActivationRequestWithDefaults instantiates a new ChangeClientActivationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ChangeClientActivationRequest) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChangeClientActivationRequest) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChangeClientActivationRequest) SetStatus(v bool)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


