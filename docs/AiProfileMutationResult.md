# AiProfileMutationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | True when the profile was persisted. | 
**Profile** | Pointer to [**AiProfile**](AiProfile.md) | The persisted profile. Present on success. | [optional] 
**Error** | Pointer to [**AiTErrorData**](AiTErrorData.md) | Why the profile was rejected - the name check or the provider credential check. Present on failure. | [optional] 

## Methods

### NewAiProfileMutationResult

`func NewAiProfileMutationResult(success bool, ) *AiProfileMutationResult`

NewAiProfileMutationResult instantiates a new AiProfileMutationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiProfileMutationResultWithDefaults

`func NewAiProfileMutationResultWithDefaults() *AiProfileMutationResult`

NewAiProfileMutationResultWithDefaults instantiates a new AiProfileMutationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *AiProfileMutationResult) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AiProfileMutationResult) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AiProfileMutationResult) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetProfile

`func (o *AiProfileMutationResult) GetProfile() AiProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *AiProfileMutationResult) GetProfileOk() (*AiProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *AiProfileMutationResult) SetProfile(v AiProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *AiProfileMutationResult) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetError

`func (o *AiProfileMutationResult) GetError() AiTErrorData`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AiProfileMutationResult) GetErrorOk() (*AiTErrorData, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AiProfileMutationResult) SetError(v AiTErrorData)`

SetError sets Error field to given value.

### HasError

`func (o *AiProfileMutationResult) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


