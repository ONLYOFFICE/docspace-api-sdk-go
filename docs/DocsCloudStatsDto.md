# DocsCloudStatsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PeriodDay** | Pointer to **int32** | The length of the statistics period in days. | [optional] 
**Editor** | Pointer to [**DocsCloudUserStatsDto**](DocsCloudUserStatsDto.md) | The statistics for editor users. | [optional] 
**Viewer** | Pointer to [**DocsCloudUserStatsDto**](DocsCloudUserStatsDto.md) | The statistics for viewer users. | [optional] 

## Methods

### NewDocsCloudStatsDto

`func NewDocsCloudStatsDto() *DocsCloudStatsDto`

NewDocsCloudStatsDto instantiates a new DocsCloudStatsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocsCloudStatsDtoWithDefaults

`func NewDocsCloudStatsDtoWithDefaults() *DocsCloudStatsDto`

NewDocsCloudStatsDtoWithDefaults instantiates a new DocsCloudStatsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriodDay

`func (o *DocsCloudStatsDto) GetPeriodDay() int32`

GetPeriodDay returns the PeriodDay field if non-nil, zero value otherwise.

### GetPeriodDayOk

`func (o *DocsCloudStatsDto) GetPeriodDayOk() (*int32, bool)`

GetPeriodDayOk returns a tuple with the PeriodDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodDay

`func (o *DocsCloudStatsDto) SetPeriodDay(v int32)`

SetPeriodDay sets PeriodDay field to given value.

### HasPeriodDay

`func (o *DocsCloudStatsDto) HasPeriodDay() bool`

HasPeriodDay returns a boolean if a field has been set.

### GetEditor

`func (o *DocsCloudStatsDto) GetEditor() DocsCloudUserStatsDto`

GetEditor returns the Editor field if non-nil, zero value otherwise.

### GetEditorOk

`func (o *DocsCloudStatsDto) GetEditorOk() (*DocsCloudUserStatsDto, bool)`

GetEditorOk returns a tuple with the Editor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditor

`func (o *DocsCloudStatsDto) SetEditor(v DocsCloudUserStatsDto)`

SetEditor sets Editor field to given value.

### HasEditor

`func (o *DocsCloudStatsDto) HasEditor() bool`

HasEditor returns a boolean if a field has been set.

### GetViewer

`func (o *DocsCloudStatsDto) GetViewer() DocsCloudUserStatsDto`

GetViewer returns the Viewer field if non-nil, zero value otherwise.

### GetViewerOk

`func (o *DocsCloudStatsDto) GetViewerOk() (*DocsCloudUserStatsDto, bool)`

GetViewerOk returns a tuple with the Viewer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewer

`func (o *DocsCloudStatsDto) SetViewer(v DocsCloudUserStatsDto)`

SetViewer sets Viewer field to given value.

### HasViewer

`func (o *DocsCloudStatsDto) HasViewer() bool`

HasViewer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


