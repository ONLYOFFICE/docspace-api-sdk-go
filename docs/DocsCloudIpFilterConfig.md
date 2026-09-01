# DocsCloudIpFilterConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]DocsCloudIpFilterRule**](DocsCloudIpFilterRule.md) | The IP filter rules. | [optional] 

## Methods

### NewDocsCloudIpFilterConfig

`func NewDocsCloudIpFilterConfig() *DocsCloudIpFilterConfig`

NewDocsCloudIpFilterConfig instantiates a new DocsCloudIpFilterConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocsCloudIpFilterConfigWithDefaults

`func NewDocsCloudIpFilterConfigWithDefaults() *DocsCloudIpFilterConfig`

NewDocsCloudIpFilterConfigWithDefaults instantiates a new DocsCloudIpFilterConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *DocsCloudIpFilterConfig) GetRules() []DocsCloudIpFilterRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *DocsCloudIpFilterConfig) GetRulesOk() (*[]DocsCloudIpFilterRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *DocsCloudIpFilterConfig) SetRules(v []DocsCloudIpFilterRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *DocsCloudIpFilterConfig) HasRules() bool`

HasRules returns a boolean if a field has been set.

### SetRulesNil

`func (o *DocsCloudIpFilterConfig) SetRulesNil(b bool)`

 SetRulesNil sets the value for Rules to be an explicit nil

### UnsetRules
`func (o *DocsCloudIpFilterConfig) UnsetRules()`

UnsetRules ensures that no value is present for Rules, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


