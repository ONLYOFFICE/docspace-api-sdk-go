# XlsxReportResponseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Form** | Pointer to [**FileDto**](FileDto.md) | The original form the answers are collected from. It is not the produced spreadsheet - that one arrives with  the task, once the task reports completion. | [optional] 
**Task** | Pointer to [**DocumentBuilderTaskDto**](DocumentBuilderTaskDto.md) | The queued generation. Poll it with `GET api/2.0/files/file/{fileId}/xlsx` until it reports completion, and  take the produced file from it then. | [optional] 
**IsNewFile** | Pointer to **bool** | True when this run creates the report file, false when an existing report is rewritten in place, which means  it keeps its id and the links already shared for it. | [optional] 

## Methods

### NewXlsxReportResponseDto

`func NewXlsxReportResponseDto() *XlsxReportResponseDto`

NewXlsxReportResponseDto instantiates a new XlsxReportResponseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXlsxReportResponseDtoWithDefaults

`func NewXlsxReportResponseDtoWithDefaults() *XlsxReportResponseDto`

NewXlsxReportResponseDtoWithDefaults instantiates a new XlsxReportResponseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForm

`func (o *XlsxReportResponseDto) GetForm() FileDto`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *XlsxReportResponseDto) GetFormOk() (*FileDto, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *XlsxReportResponseDto) SetForm(v FileDto)`

SetForm sets Form field to given value.

### HasForm

`func (o *XlsxReportResponseDto) HasForm() bool`

HasForm returns a boolean if a field has been set.

### GetTask

`func (o *XlsxReportResponseDto) GetTask() DocumentBuilderTaskDto`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *XlsxReportResponseDto) GetTaskOk() (*DocumentBuilderTaskDto, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *XlsxReportResponseDto) SetTask(v DocumentBuilderTaskDto)`

SetTask sets Task field to given value.

### HasTask

`func (o *XlsxReportResponseDto) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetIsNewFile

`func (o *XlsxReportResponseDto) GetIsNewFile() bool`

GetIsNewFile returns the IsNewFile field if non-nil, zero value otherwise.

### GetIsNewFileOk

`func (o *XlsxReportResponseDto) GetIsNewFileOk() (*bool, bool)`

GetIsNewFileOk returns a tuple with the IsNewFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNewFile

`func (o *XlsxReportResponseDto) SetIsNewFile(v bool)`

SetIsNewFile sets IsNewFile field to given value.

### HasIsNewFile

`func (o *XlsxReportResponseDto) HasIsNewFile() bool`

HasIsNewFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


