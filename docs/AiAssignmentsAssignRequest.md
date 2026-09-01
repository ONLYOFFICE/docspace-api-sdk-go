# AiAssignmentsAssignRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionType** | [**AiActionType**](AiActionType.md) | Action the assignment applies to. | 
**ProfileId** | **string** | Profile id to bind. | 

## Methods

### NewAiAssignmentsAssignRequest

`func NewAiAssignmentsAssignRequest(actionType AiActionType, profileId string, ) *AiAssignmentsAssignRequest`

NewAiAssignmentsAssignRequest instantiates a new AiAssignmentsAssignRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiAssignmentsAssignRequestWithDefaults

`func NewAiAssignmentsAssignRequestWithDefaults() *AiAssignmentsAssignRequest`

NewAiAssignmentsAssignRequestWithDefaults instantiates a new AiAssignmentsAssignRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionType

`func (o *AiAssignmentsAssignRequest) GetActionType() AiActionType`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *AiAssignmentsAssignRequest) GetActionTypeOk() (*AiActionType, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *AiAssignmentsAssignRequest) SetActionType(v AiActionType)`

SetActionType sets ActionType field to given value.


### GetProfileId

`func (o *AiAssignmentsAssignRequest) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *AiAssignmentsAssignRequest) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *AiAssignmentsAssignRequest) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


