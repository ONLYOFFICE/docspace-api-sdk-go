# CheckDestFolderDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**CheckDestFolderResult**](CheckDestFolderResult.md) |  | [optional] 
**Files** | Pointer to [**[]FileEntryBaseDto**](FileEntryBaseDto.md) | The list of files in the destination folder. | [optional] 

## Methods

### NewCheckDestFolderDto

`func NewCheckDestFolderDto() *CheckDestFolderDto`

NewCheckDestFolderDto instantiates a new CheckDestFolderDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckDestFolderDtoWithDefaults

`func NewCheckDestFolderDtoWithDefaults() *CheckDestFolderDto`

NewCheckDestFolderDtoWithDefaults instantiates a new CheckDestFolderDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *CheckDestFolderDto) GetResult() CheckDestFolderResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CheckDestFolderDto) GetResultOk() (*CheckDestFolderResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CheckDestFolderDto) SetResult(v CheckDestFolderResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *CheckDestFolderDto) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetFiles

`func (o *CheckDestFolderDto) GetFiles() []FileEntryBaseDto`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *CheckDestFolderDto) GetFilesOk() (*[]FileEntryBaseDto, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *CheckDestFolderDto) SetFiles(v []FileEntryBaseDto)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *CheckDestFolderDto) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### SetFilesNil

`func (o *CheckDestFolderDto) SetFilesNil(b bool)`

 SetFilesNil sets the value for Files to be an explicit nil

### UnsetFiles
`func (o *CheckDestFolderDto) UnsetFiles()`

UnsetFiles ensures that no value is present for Files, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


