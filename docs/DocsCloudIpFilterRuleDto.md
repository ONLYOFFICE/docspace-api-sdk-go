# DocsCloudIpFilterRuleDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **NullableString** | The IP address. | [optional] 
**Allowed** | Pointer to **bool** | Whether the IP address is allowed. | [optional] 

## Methods

### NewDocsCloudIpFilterRuleDto

`func NewDocsCloudIpFilterRuleDto() *DocsCloudIpFilterRuleDto`

NewDocsCloudIpFilterRuleDto instantiates a new DocsCloudIpFilterRuleDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocsCloudIpFilterRuleDtoWithDefaults

`func NewDocsCloudIpFilterRuleDtoWithDefaults() *DocsCloudIpFilterRuleDto`

NewDocsCloudIpFilterRuleDtoWithDefaults instantiates a new DocsCloudIpFilterRuleDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *DocsCloudIpFilterRuleDto) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DocsCloudIpFilterRuleDto) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DocsCloudIpFilterRuleDto) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DocsCloudIpFilterRuleDto) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *DocsCloudIpFilterRuleDto) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *DocsCloudIpFilterRuleDto) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetAllowed

`func (o *DocsCloudIpFilterRuleDto) GetAllowed() bool`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *DocsCloudIpFilterRuleDto) GetAllowedOk() (*bool, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *DocsCloudIpFilterRuleDto) SetAllowed(v bool)`

SetAllowed sets Allowed field to given value.

### HasAllowed

`func (o *DocsCloudIpFilterRuleDto) HasAllowed() bool`

HasAllowed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


