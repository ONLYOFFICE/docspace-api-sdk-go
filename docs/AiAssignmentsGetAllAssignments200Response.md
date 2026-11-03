# AiAssignmentsGetAllAssignments200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignments** | [**map[string]AiProfile**](AiProfile.md) | The bound profiles keyed by action type, as `GET api/2.0/ai/profiles/list` describes them. A binding whose profile no longer exists is left out. | 
**ProfilesTotal** | **int32** | How many profiles the portal has. | 
**ChatProfilesTotal** | **int32** | How many of those profiles can chat; a profile without declared capabilities counts as one that can. | 

## Methods

### NewAiAssignmentsGetAllAssignments200Response

`func NewAiAssignmentsGetAllAssignments200Response(assignments map[string]AiProfile, profilesTotal int32, chatProfilesTotal int32, ) *AiAssignmentsGetAllAssignments200Response`

NewAiAssignmentsGetAllAssignments200Response instantiates a new AiAssignmentsGetAllAssignments200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiAssignmentsGetAllAssignments200ResponseWithDefaults

`func NewAiAssignmentsGetAllAssignments200ResponseWithDefaults() *AiAssignmentsGetAllAssignments200Response`

NewAiAssignmentsGetAllAssignments200ResponseWithDefaults instantiates a new AiAssignmentsGetAllAssignments200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignments

`func (o *AiAssignmentsGetAllAssignments200Response) GetAssignments() map[string]AiProfile`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *AiAssignmentsGetAllAssignments200Response) GetAssignmentsOk() (*map[string]AiProfile, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *AiAssignmentsGetAllAssignments200Response) SetAssignments(v map[string]AiProfile)`

SetAssignments sets Assignments field to given value.


### GetProfilesTotal

`func (o *AiAssignmentsGetAllAssignments200Response) GetProfilesTotal() int32`

GetProfilesTotal returns the ProfilesTotal field if non-nil, zero value otherwise.

### GetProfilesTotalOk

`func (o *AiAssignmentsGetAllAssignments200Response) GetProfilesTotalOk() (*int32, bool)`

GetProfilesTotalOk returns a tuple with the ProfilesTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilesTotal

`func (o *AiAssignmentsGetAllAssignments200Response) SetProfilesTotal(v int32)`

SetProfilesTotal sets ProfilesTotal field to given value.


### GetChatProfilesTotal

`func (o *AiAssignmentsGetAllAssignments200Response) GetChatProfilesTotal() int32`

GetChatProfilesTotal returns the ChatProfilesTotal field if non-nil, zero value otherwise.

### GetChatProfilesTotalOk

`func (o *AiAssignmentsGetAllAssignments200Response) GetChatProfilesTotalOk() (*int32, bool)`

GetChatProfilesTotalOk returns a tuple with the ChatProfilesTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatProfilesTotal

`func (o *AiAssignmentsGetAllAssignments200Response) SetChatProfilesTotal(v int32)`

SetChatProfilesTotal sets ChatProfilesTotal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


