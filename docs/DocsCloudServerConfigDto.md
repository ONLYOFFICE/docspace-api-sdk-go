# DocsCloudServerConfigDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAnonymousSupport** | Pointer to **bool** | Whether anonymous access is supported. | [optional] 
**FileSizeLimit** | Pointer to **int64** | The maximum file size in bytes. | [optional] 

## Methods

### NewDocsCloudServerConfigDto

`func NewDocsCloudServerConfigDto() *DocsCloudServerConfigDto`

NewDocsCloudServerConfigDto instantiates a new DocsCloudServerConfigDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocsCloudServerConfigDtoWithDefaults

`func NewDocsCloudServerConfigDtoWithDefaults() *DocsCloudServerConfigDto`

NewDocsCloudServerConfigDtoWithDefaults instantiates a new DocsCloudServerConfigDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAnonymousSupport

`func (o *DocsCloudServerConfigDto) GetIsAnonymousSupport() bool`

GetIsAnonymousSupport returns the IsAnonymousSupport field if non-nil, zero value otherwise.

### GetIsAnonymousSupportOk

`func (o *DocsCloudServerConfigDto) GetIsAnonymousSupportOk() (*bool, bool)`

GetIsAnonymousSupportOk returns a tuple with the IsAnonymousSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnonymousSupport

`func (o *DocsCloudServerConfigDto) SetIsAnonymousSupport(v bool)`

SetIsAnonymousSupport sets IsAnonymousSupport field to given value.

### HasIsAnonymousSupport

`func (o *DocsCloudServerConfigDto) HasIsAnonymousSupport() bool`

HasIsAnonymousSupport returns a boolean if a field has been set.

### GetFileSizeLimit

`func (o *DocsCloudServerConfigDto) GetFileSizeLimit() int64`

GetFileSizeLimit returns the FileSizeLimit field if non-nil, zero value otherwise.

### GetFileSizeLimitOk

`func (o *DocsCloudServerConfigDto) GetFileSizeLimitOk() (*int64, bool)`

GetFileSizeLimitOk returns a tuple with the FileSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSizeLimit

`func (o *DocsCloudServerConfigDto) SetFileSizeLimit(v int64)`

SetFileSizeLimit sets FileSizeLimit field to given value.

### HasFileSizeLimit

`func (o *DocsCloudServerConfigDto) HasFileSizeLimit() bool`

HasFileSizeLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


