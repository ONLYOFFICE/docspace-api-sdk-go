# TenantAiAccessSettingsRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether AI is available on the portal at all - chat, agents and vectorization together. Switching it off  hides the AI Agents folder and makes every AI endpoint unreachable for all members at once, not only for the  caller, and the change is pushed to connected clients rather than waiting for their next request. | [optional] 

## Methods

### NewTenantAiAccessSettingsRequestDto

`func NewTenantAiAccessSettingsRequestDto() *TenantAiAccessSettingsRequestDto`

NewTenantAiAccessSettingsRequestDto instantiates a new TenantAiAccessSettingsRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantAiAccessSettingsRequestDtoWithDefaults

`func NewTenantAiAccessSettingsRequestDtoWithDefaults() *TenantAiAccessSettingsRequestDto`

NewTenantAiAccessSettingsRequestDtoWithDefaults instantiates a new TenantAiAccessSettingsRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *TenantAiAccessSettingsRequestDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TenantAiAccessSettingsRequestDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *TenantAiAccessSettingsRequestDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *TenantAiAccessSettingsRequestDto) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


