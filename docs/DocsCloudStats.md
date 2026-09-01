# DocsCloudStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PeriodDay** | Pointer to **int32** | The length of the statistics period in days. | [optional] 
**Editor** | Pointer to [**DocsCloudUserStats**](DocsCloudUserStats.md) | The statistics for editor users. | [optional] 
**Viewer** | Pointer to [**DocsCloudUserStats**](DocsCloudUserStats.md) | The statistics for viewer users. | [optional] 

## Methods

### NewDocsCloudStats

`func NewDocsCloudStats() *DocsCloudStats`

NewDocsCloudStats instantiates a new DocsCloudStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocsCloudStatsWithDefaults

`func NewDocsCloudStatsWithDefaults() *DocsCloudStats`

NewDocsCloudStatsWithDefaults instantiates a new DocsCloudStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriodDay

`func (o *DocsCloudStats) GetPeriodDay() int32`

GetPeriodDay returns the PeriodDay field if non-nil, zero value otherwise.

### GetPeriodDayOk

`func (o *DocsCloudStats) GetPeriodDayOk() (*int32, bool)`

GetPeriodDayOk returns a tuple with the PeriodDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodDay

`func (o *DocsCloudStats) SetPeriodDay(v int32)`

SetPeriodDay sets PeriodDay field to given value.

### HasPeriodDay

`func (o *DocsCloudStats) HasPeriodDay() bool`

HasPeriodDay returns a boolean if a field has been set.

### GetEditor

`func (o *DocsCloudStats) GetEditor() DocsCloudUserStats`

GetEditor returns the Editor field if non-nil, zero value otherwise.

### GetEditorOk

`func (o *DocsCloudStats) GetEditorOk() (*DocsCloudUserStats, bool)`

GetEditorOk returns a tuple with the Editor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditor

`func (o *DocsCloudStats) SetEditor(v DocsCloudUserStats)`

SetEditor sets Editor field to given value.

### HasEditor

`func (o *DocsCloudStats) HasEditor() bool`

HasEditor returns a boolean if a field has been set.

### GetViewer

`func (o *DocsCloudStats) GetViewer() DocsCloudUserStats`

GetViewer returns the Viewer field if non-nil, zero value otherwise.

### GetViewerOk

`func (o *DocsCloudStats) GetViewerOk() (*DocsCloudUserStats, bool)`

GetViewerOk returns a tuple with the Viewer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewer

`func (o *DocsCloudStats) SetViewer(v DocsCloudUserStats)`

SetViewer sets Viewer field to given value.

### HasViewer

`func (o *DocsCloudStats) HasViewer() bool`

HasViewer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


