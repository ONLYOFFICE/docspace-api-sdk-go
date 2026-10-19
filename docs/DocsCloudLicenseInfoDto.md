# DocsCloudLicenseInfoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Valid** | Pointer to **time.Time** | The date and time until which the license is valid. | [optional] 
**Trial** | Pointer to **bool** | Whether the license is a trial. | [optional] 
**BuildDate** | Pointer to **time.Time** | The license build date. | [optional] 

## Methods

### NewDocsCloudLicenseInfoDto

`func NewDocsCloudLicenseInfoDto() *DocsCloudLicenseInfoDto`

NewDocsCloudLicenseInfoDto instantiates a new DocsCloudLicenseInfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocsCloudLicenseInfoDtoWithDefaults

`func NewDocsCloudLicenseInfoDtoWithDefaults() *DocsCloudLicenseInfoDto`

NewDocsCloudLicenseInfoDtoWithDefaults instantiates a new DocsCloudLicenseInfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValid

`func (o *DocsCloudLicenseInfoDto) GetValid() time.Time`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *DocsCloudLicenseInfoDto) GetValidOk() (*time.Time, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *DocsCloudLicenseInfoDto) SetValid(v time.Time)`

SetValid sets Valid field to given value.

### HasValid

`func (o *DocsCloudLicenseInfoDto) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetTrial

`func (o *DocsCloudLicenseInfoDto) GetTrial() bool`

GetTrial returns the Trial field if non-nil, zero value otherwise.

### GetTrialOk

`func (o *DocsCloudLicenseInfoDto) GetTrialOk() (*bool, bool)`

GetTrialOk returns a tuple with the Trial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrial

`func (o *DocsCloudLicenseInfoDto) SetTrial(v bool)`

SetTrial sets Trial field to given value.

### HasTrial

`func (o *DocsCloudLicenseInfoDto) HasTrial() bool`

HasTrial returns a boolean if a field has been set.

### GetBuildDate

`func (o *DocsCloudLicenseInfoDto) GetBuildDate() time.Time`

GetBuildDate returns the BuildDate field if non-nil, zero value otherwise.

### GetBuildDateOk

`func (o *DocsCloudLicenseInfoDto) GetBuildDateOk() (*time.Time, bool)`

GetBuildDateOk returns a tuple with the BuildDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildDate

`func (o *DocsCloudLicenseInfoDto) SetBuildDate(v time.Time)`

SetBuildDate sets BuildDate field to given value.

### HasBuildDate

`func (o *DocsCloudLicenseInfoDto) HasBuildDate() bool`

HasBuildDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


