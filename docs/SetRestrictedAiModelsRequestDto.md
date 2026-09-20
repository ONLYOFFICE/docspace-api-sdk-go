# SetRestrictedAiModelsRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Models** | **[]string** | The identifiers of the models no user of the portal may pick, taken from  `GET api/2.0/portal/payment/ai-prices`. This is the whole set that is to hold afterwards and not a list of  additions: send the models already barred together with the new one to add a restriction, leave one out to  lift it, and send an empty set to lift them all. | 

## Methods

### NewSetRestrictedAiModelsRequestDto

`func NewSetRestrictedAiModelsRequestDto(models []string, ) *SetRestrictedAiModelsRequestDto`

NewSetRestrictedAiModelsRequestDto instantiates a new SetRestrictedAiModelsRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetRestrictedAiModelsRequestDtoWithDefaults

`func NewSetRestrictedAiModelsRequestDtoWithDefaults() *SetRestrictedAiModelsRequestDto`

NewSetRestrictedAiModelsRequestDtoWithDefaults instantiates a new SetRestrictedAiModelsRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModels

`func (o *SetRestrictedAiModelsRequestDto) GetModels() []string`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *SetRestrictedAiModelsRequestDto) GetModelsOk() (*[]string, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *SetRestrictedAiModelsRequestDto) SetModels(v []string)`

SetModels sets Models field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


