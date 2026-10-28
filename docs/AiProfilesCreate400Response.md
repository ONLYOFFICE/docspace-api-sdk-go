# AiProfilesCreate400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | True when the profile was persisted. | 
**Profile** | Pointer to [**AiProfile**](AiProfile.md) | The persisted profile. Present on success. | [optional] 
**Error** | **string** | The error message, ready to be shown to the caller. | 

## Methods

### NewAiProfilesCreate400Response

`func NewAiProfilesCreate400Response(success bool, error_ string, ) *AiProfilesCreate400Response`

NewAiProfilesCreate400Response instantiates a new AiProfilesCreate400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiProfilesCreate400ResponseWithDefaults

`func NewAiProfilesCreate400ResponseWithDefaults() *AiProfilesCreate400Response`

NewAiProfilesCreate400ResponseWithDefaults instantiates a new AiProfilesCreate400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *AiProfilesCreate400Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AiProfilesCreate400Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AiProfilesCreate400Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetProfile

`func (o *AiProfilesCreate400Response) GetProfile() AiProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *AiProfilesCreate400Response) GetProfileOk() (*AiProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *AiProfilesCreate400Response) SetProfile(v AiProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *AiProfilesCreate400Response) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetError

`func (o *AiProfilesCreate400Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AiProfilesCreate400Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AiProfilesCreate400Response) SetError(v string)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


