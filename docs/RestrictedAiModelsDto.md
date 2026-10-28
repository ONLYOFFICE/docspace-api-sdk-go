# RestrictedAiModelsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Models** | **[]string** | The identifiers of the models the portal is not allowed to use. | 

## Methods

### NewRestrictedAiModelsDto

`func NewRestrictedAiModelsDto(models []string, ) *RestrictedAiModelsDto`

NewRestrictedAiModelsDto instantiates a new RestrictedAiModelsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestrictedAiModelsDtoWithDefaults

`func NewRestrictedAiModelsDtoWithDefaults() *RestrictedAiModelsDto`

NewRestrictedAiModelsDtoWithDefaults instantiates a new RestrictedAiModelsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModels

`func (o *RestrictedAiModelsDto) GetModels() []string`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *RestrictedAiModelsDto) GetModelsOk() (*[]string, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *RestrictedAiModelsDto) SetModels(v []string)`

SetModels sets Models field to given value.


### SetModelsNil

`func (o *RestrictedAiModelsDto) SetModelsNil(b bool)`

 SetModelsNil sets the value for Models to be an explicit nil

### UnsetModels
`func (o *RestrictedAiModelsDto) UnsetModels()`

UnsetModels ensures that no value is present for Models, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


