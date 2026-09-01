# FilesStatisticsFolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** | The folder title. | [optional] 
**UsedSpace** | Pointer to **int64** | The used space in the folder. | [optional] 

## Methods

### NewFilesStatisticsFolder

`func NewFilesStatisticsFolder() *FilesStatisticsFolder`

NewFilesStatisticsFolder instantiates a new FilesStatisticsFolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesStatisticsFolderWithDefaults

`func NewFilesStatisticsFolderWithDefaults() *FilesStatisticsFolder`

NewFilesStatisticsFolderWithDefaults instantiates a new FilesStatisticsFolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *FilesStatisticsFolder) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FilesStatisticsFolder) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FilesStatisticsFolder) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FilesStatisticsFolder) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *FilesStatisticsFolder) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *FilesStatisticsFolder) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetUsedSpace

`func (o *FilesStatisticsFolder) GetUsedSpace() int64`

GetUsedSpace returns the UsedSpace field if non-nil, zero value otherwise.

### GetUsedSpaceOk

`func (o *FilesStatisticsFolder) GetUsedSpaceOk() (*int64, bool)`

GetUsedSpaceOk returns a tuple with the UsedSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedSpace

`func (o *FilesStatisticsFolder) SetUsedSpace(v int64)`

SetUsedSpace sets UsedSpace field to given value.

### HasUsedSpace

`func (o *FilesStatisticsFolder) HasUsedSpace() bool`

HasUsedSpace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


