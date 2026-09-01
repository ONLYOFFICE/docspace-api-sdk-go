# DocsCloudUserStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **int32** | The number of active users. | [optional] 
**Internal** | Pointer to **int32** | The number of internal users. | [optional] 
**External** | Pointer to **int32** | The number of external users. | [optional] 
**Remaining** | Pointer to **int32** | The number of remaining users before the limit is reached. | [optional] 
**CriticalRemaining** | Pointer to **bool** | Whether the number of remaining users is critically low. | [optional] 

## Methods

### NewDocsCloudUserStats

`func NewDocsCloudUserStats() *DocsCloudUserStats`

NewDocsCloudUserStats instantiates a new DocsCloudUserStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocsCloudUserStatsWithDefaults

`func NewDocsCloudUserStatsWithDefaults() *DocsCloudUserStats`

NewDocsCloudUserStatsWithDefaults instantiates a new DocsCloudUserStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *DocsCloudUserStats) GetActive() int32`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *DocsCloudUserStats) GetActiveOk() (*int32, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *DocsCloudUserStats) SetActive(v int32)`

SetActive sets Active field to given value.

### HasActive

`func (o *DocsCloudUserStats) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetInternal

`func (o *DocsCloudUserStats) GetInternal() int32`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *DocsCloudUserStats) GetInternalOk() (*int32, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *DocsCloudUserStats) SetInternal(v int32)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *DocsCloudUserStats) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetExternal

`func (o *DocsCloudUserStats) GetExternal() int32`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *DocsCloudUserStats) GetExternalOk() (*int32, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *DocsCloudUserStats) SetExternal(v int32)`

SetExternal sets External field to given value.

### HasExternal

`func (o *DocsCloudUserStats) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### GetRemaining

`func (o *DocsCloudUserStats) GetRemaining() int32`

GetRemaining returns the Remaining field if non-nil, zero value otherwise.

### GetRemainingOk

`func (o *DocsCloudUserStats) GetRemainingOk() (*int32, bool)`

GetRemainingOk returns a tuple with the Remaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemaining

`func (o *DocsCloudUserStats) SetRemaining(v int32)`

SetRemaining sets Remaining field to given value.

### HasRemaining

`func (o *DocsCloudUserStats) HasRemaining() bool`

HasRemaining returns a boolean if a field has been set.

### GetCriticalRemaining

`func (o *DocsCloudUserStats) GetCriticalRemaining() bool`

GetCriticalRemaining returns the CriticalRemaining field if non-nil, zero value otherwise.

### GetCriticalRemainingOk

`func (o *DocsCloudUserStats) GetCriticalRemainingOk() (*bool, bool)`

GetCriticalRemainingOk returns a tuple with the CriticalRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalRemaining

`func (o *DocsCloudUserStats) SetCriticalRemaining(v bool)`

SetCriticalRemaining sets CriticalRemaining field to given value.

### HasCriticalRemaining

`func (o *DocsCloudUserStats) HasCriticalRemaining() bool`

HasCriticalRemaining returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


