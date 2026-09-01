# WebItemSecurityRequestsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** | The module ID. | 
**Enabled** | Pointer to **bool** | Controls whether the security restrictions are enforced for this module. | [optional] 
**Subjects** | Pointer to **[]string** | The collection of user and group identifiers granted access to the module. | [optional] 

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


