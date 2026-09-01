# AiResolvedAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | **string** | The identifier of the resolved profile. | 
**Profile** | [**AiProfile**](AiProfile.md) | The resolved profile itself. | 

## Methods

### NewAiResolvedAssignment

`func NewAiResolvedAssignment(profileId string, profile AiProfile, ) *AiResolvedAssignment`

NewAiResolvedAssignment instantiates a new AiResolvedAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiResolvedAssignmentWithDefaults

`func NewAiResolvedAssignmentWithDefaults() *AiResolvedAssignment`

NewAiResolvedAssignmentWithDefaults instantiates a new AiResolvedAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *AiResolvedAssignment) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *AiResolvedAssignment) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *AiResolvedAssignment) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetProfile

`func (o *AiResolvedAssignment) GetProfile() AiProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *AiResolvedAssignment) GetProfileOk() (*AiProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *AiResolvedAssignment) SetProfile(v AiProfile)`

SetProfile sets Profile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


