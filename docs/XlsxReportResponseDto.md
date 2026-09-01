# XlsxReportResponseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Form** | Pointer to [**FileDtoInteger**](FileDtoInteger.md) | The original form file information. | [optional] 
**Task** | Pointer to [**DocumentBuilderTaskDto**](DocumentBuilderTaskDto.md) | The Document Builder task information. | [optional] 
**IsNewFile** | Pointer to **bool** | Specifies whether the XLSX report file is newly created or an existing file will be updated. | [optional] 

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

`func (o *XlsxReportResponseDto) GetForm() FileDtoInteger`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *XlsxReportResponseDto) GetFormOk() (*FileDtoInteger, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *XlsxReportResponseDto) SetForm(v FileDtoInteger)`

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


