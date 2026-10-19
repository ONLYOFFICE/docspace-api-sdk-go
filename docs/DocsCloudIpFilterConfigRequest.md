# DocsCloudIpFilterConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]DocsCloudIpFilterRuleRequest**](DocsCloudIpFilterRuleRequest.md) | The IP filter rules. | [optional] 

## Methods

### NewDocsCloudIpFilterConfigRequest

`func NewDocsCloudIpFilterConfigRequest() *DocsCloudIpFilterConfigRequest`

NewDocsCloudIpFilterConfigRequest instantiates a new DocsCloudIpFilterConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocsCloudIpFilterConfigRequestWithDefaults

`func NewDocsCloudIpFilterConfigRequestWithDefaults() *DocsCloudIpFilterConfigRequest`

NewDocsCloudIpFilterConfigRequestWithDefaults instantiates a new DocsCloudIpFilterConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *DocsCloudIpFilterConfigRequest) GetRules() []DocsCloudIpFilterRuleRequest`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *DocsCloudIpFilterConfigRequest) GetRulesOk() (*[]DocsCloudIpFilterRuleRequest, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *DocsCloudIpFilterConfigRequest) SetRules(v []DocsCloudIpFilterRuleRequest)`

SetRules sets Rules field to given value.

### HasRules

`func (o *DocsCloudIpFilterConfigRequest) HasRules() bool`

HasRules returns a boolean if a field has been set.

### SetRulesNil

`func (o *DocsCloudIpFilterConfigRequest) SetRulesNil(b bool)`

 SetRulesNil sets the value for Rules to be an explicit nil

### UnsetRules
`func (o *DocsCloudIpFilterConfigRequest) UnsetRules()`

UnsetRules ensures that no value is present for Rules, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


