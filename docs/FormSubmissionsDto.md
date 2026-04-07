# FormSubmissionsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**[]FormMetadata**](FormMetadata.md) | The form field metadata. | [optional] 
**Submissions** | Pointer to [**[]FormResultsDto**](FormResultsDto.md) | All submissions. | [optional] 

## Methods

### NewFormSubmissionsDto

`func NewFormSubmissionsDto() *FormSubmissionsDto`

NewFormSubmissionsDto instantiates a new FormSubmissionsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormSubmissionsDtoWithDefaults

`func NewFormSubmissionsDtoWithDefaults() *FormSubmissionsDto`

NewFormSubmissionsDtoWithDefaults instantiates a new FormSubmissionsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *FormSubmissionsDto) GetMetadata() []FormMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FormSubmissionsDto) GetMetadataOk() (*[]FormMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FormSubmissionsDto) SetMetadata(v []FormMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FormSubmissionsDto) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *FormSubmissionsDto) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *FormSubmissionsDto) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetSubmissions

`func (o *FormSubmissionsDto) GetSubmissions() []FormResultsDto`

GetSubmissions returns the Submissions field if non-nil, zero value otherwise.

### GetSubmissionsOk

`func (o *FormSubmissionsDto) GetSubmissionsOk() (*[]FormResultsDto, bool)`

GetSubmissionsOk returns a tuple with the Submissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissions

`func (o *FormSubmissionsDto) SetSubmissions(v []FormResultsDto)`

SetSubmissions sets Submissions field to given value.

### HasSubmissions

`func (o *FormSubmissionsDto) HasSubmissions() bool`

HasSubmissions returns a boolean if a field has been set.

### SetSubmissionsNil

`func (o *FormSubmissionsDto) SetSubmissionsNil(b bool)`

 SetSubmissionsNil sets the value for Submissions to be an explicit nil

### UnsetSubmissions
`func (o *FormSubmissionsDto) UnsetSubmissions()`

UnsetSubmissions ensures that no value is present for Submissions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


