# WebItemSecurityRequestsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** | The module the rule applies to, given as a GUID. A value that is not a GUID fails the request as invalid. | 
**Enabled** | Pointer to **bool** | Whether the module may be opened. It decides the outcome only while `subjects` names somebody: an empty  `subjects` array is stored as access for everyone whatever this flag says. | [optional] 
**Subjects** | Pointer to **[]string** | The users and groups the rule is stored for, given by their IDs. This is the whole allow-list that is to hold  afterwards and not a list of additions - what was stored before is dropped. Leaving it out applies `enabled`  to everyone and skips the audit trail entry, while sending it empty stores access for everyone. | [optional] 

## Methods

### NewWebItemSecurityRequestsDto

`func NewWebItemSecurityRequestsDto(id NullableString, ) *WebItemSecurityRequestsDto`

NewWebItemSecurityRequestsDto instantiates a new WebItemSecurityRequestsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebItemSecurityRequestsDtoWithDefaults

`func NewWebItemSecurityRequestsDtoWithDefaults() *WebItemSecurityRequestsDto`

NewWebItemSecurityRequestsDtoWithDefaults instantiates a new WebItemSecurityRequestsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebItemSecurityRequestsDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebItemSecurityRequestsDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebItemSecurityRequestsDto) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *WebItemSecurityRequestsDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *WebItemSecurityRequestsDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetEnabled

`func (o *WebItemSecurityRequestsDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WebItemSecurityRequestsDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WebItemSecurityRequestsDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WebItemSecurityRequestsDto) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSubjects

`func (o *WebItemSecurityRequestsDto) GetSubjects() []string`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *WebItemSecurityRequestsDto) GetSubjectsOk() (*[]string, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *WebItemSecurityRequestsDto) SetSubjects(v []string)`

SetSubjects sets Subjects field to given value.

### HasSubjects

`func (o *WebItemSecurityRequestsDto) HasSubjects() bool`

HasSubjects returns a boolean if a field has been set.

### SetSubjectsNil

`func (o *WebItemSecurityRequestsDto) SetSubjectsNil(b bool)`

 SetSubjectsNil sets the value for Subjects to be an explicit nil

### UnsetSubjects
`func (o *WebItemSecurityRequestsDto) UnsetSubjects()`

UnsetSubjects ensures that no value is present for Subjects, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


