# DocsCloudIpFilterConfigDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]DocsCloudIpFilterRuleDto**](DocsCloudIpFilterRuleDto.md) | The IP filter rules. | [optional] 

## Methods

### NewDocsCloudIpFilterConfigDto

`func NewDocsCloudIpFilterConfigDto() *DocsCloudIpFilterConfigDto`

NewDocsCloudIpFilterConfigDto instantiates a new DocsCloudIpFilterConfigDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocsCloudIpFilterConfigDtoWithDefaults

`func NewDocsCloudIpFilterConfigDtoWithDefaults() *DocsCloudIpFilterConfigDto`

NewDocsCloudIpFilterConfigDtoWithDefaults instantiates a new DocsCloudIpFilterConfigDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *DocsCloudIpFilterConfigDto) GetRules() []DocsCloudIpFilterRuleDto`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *DocsCloudIpFilterConfigDto) GetRulesOk() (*[]DocsCloudIpFilterRuleDto, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *DocsCloudIpFilterConfigDto) SetRules(v []DocsCloudIpFilterRuleDto)`

SetRules sets Rules field to given value.

### HasRules

`func (o *DocsCloudIpFilterConfigDto) HasRules() bool`

HasRules returns a boolean if a field has been set.

### SetRulesNil

`func (o *DocsCloudIpFilterConfigDto) SetRulesNil(b bool)`

 SetRulesNil sets the value for Rules to be an explicit nil

### UnsetRules
`func (o *DocsCloudIpFilterConfigDto) UnsetRules()`

UnsetRules ensures that no value is present for Rules, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


