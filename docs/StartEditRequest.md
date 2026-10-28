# StartEditRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EditingAlone** | Pointer to **bool** | Claims the file for this caller alone: the session is opened without asking the document service to track  co-editing, and the call is refused when anybody else already has the file open. Left off, an ordinary  co-editing session is opened and others may join it. | [optional] 

## Methods

### NewStartEditRequest

`func NewStartEditRequest() *StartEditRequest`

NewStartEditRequest instantiates a new StartEditRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartEditRequestWithDefaults

`func NewStartEditRequestWithDefaults() *StartEditRequest`

NewStartEditRequestWithDefaults instantiates a new StartEditRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEditingAlone

`func (o *StartEditRequest) GetEditingAlone() bool`

GetEditingAlone returns the EditingAlone field if non-nil, zero value otherwise.

### GetEditingAloneOk

`func (o *StartEditRequest) GetEditingAloneOk() (*bool, bool)`

GetEditingAloneOk returns a tuple with the EditingAlone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditingAlone

`func (o *StartEditRequest) SetEditingAlone(v bool)`

SetEditingAlone sets EditingAlone field to given value.

### HasEditingAlone

`func (o *StartEditRequest) HasEditingAlone() bool`

HasEditingAlone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


