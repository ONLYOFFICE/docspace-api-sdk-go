# UpdateAiProjectModelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | **string** | ID of the AI model profile that answers in the project's chat; take it from the `id` of a profile that `GET api/2.0/ai/profiles/list` returns. | 

## Methods

### NewUpdateAiProjectModelRequest

`func NewUpdateAiProjectModelRequest(profileId string, ) *UpdateAiProjectModelRequest`

NewUpdateAiProjectModelRequest instantiates a new UpdateAiProjectModelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAiProjectModelRequestWithDefaults

`func NewUpdateAiProjectModelRequestWithDefaults() *UpdateAiProjectModelRequest`

NewUpdateAiProjectModelRequestWithDefaults instantiates a new UpdateAiProjectModelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *UpdateAiProjectModelRequest) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *UpdateAiProjectModelRequest) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *UpdateAiProjectModelRequest) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


