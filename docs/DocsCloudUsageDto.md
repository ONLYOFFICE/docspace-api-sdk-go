# DocsCloudUsageDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Since** | Pointer to **time.Time** | The date and time the usage statistics are counted from. | [optional] 
**ActiveCount** | Pointer to **int32** | The number of active users. | [optional] 

## Methods

### NewDocsCloudUsageDto

`func NewDocsCloudUsageDto() *DocsCloudUsageDto`

NewDocsCloudUsageDto instantiates a new DocsCloudUsageDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocsCloudUsageDtoWithDefaults

`func NewDocsCloudUsageDtoWithDefaults() *DocsCloudUsageDto`

NewDocsCloudUsageDtoWithDefaults instantiates a new DocsCloudUsageDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSince

`func (o *DocsCloudUsageDto) GetSince() time.Time`

GetSince returns the Since field if non-nil, zero value otherwise.

### GetSinceOk

`func (o *DocsCloudUsageDto) GetSinceOk() (*time.Time, bool)`

GetSinceOk returns a tuple with the Since field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSince

`func (o *DocsCloudUsageDto) SetSince(v time.Time)`

SetSince sets Since field to given value.

### HasSince

`func (o *DocsCloudUsageDto) HasSince() bool`

HasSince returns a boolean if a field has been set.

### GetActiveCount

`func (o *DocsCloudUsageDto) GetActiveCount() int32`

GetActiveCount returns the ActiveCount field if non-nil, zero value otherwise.

### GetActiveCountOk

`func (o *DocsCloudUsageDto) GetActiveCountOk() (*int32, bool)`

GetActiveCountOk returns a tuple with the ActiveCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveCount

`func (o *DocsCloudUsageDto) SetActiveCount(v int32)`

SetActiveCount sets ActiveCount field to given value.

### HasActiveCount

`func (o *DocsCloudUsageDto) HasActiveCount() bool`

HasActiveCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


