# TfaSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** | The ID of the TFA configuration. | 
**Title** | **NullableString** | The display name or description of the TFA configuration. | 
**Enabled** | **bool** | Indicates whether the TFA configuration is currently active. | 
**Avaliable** | **bool** | Indicates whether the TFA configuration can be used. | 
**TrustedIps** | Pointer to **[]string** | The list of IP addresses that are exempt from TFA requirements. | [optional] 
**MandatoryUsers** | Pointer to **[]string** | The list of user IDs that are required to use TFA. | [optional] 
**MandatoryGroups** | Pointer to **[]string** | The list of group IDs whose members are required to use TFA. | [optional] 

## Methods

### NewTfaSettingsDto

`func NewTfaSettingsDto(id NullableString, title NullableString, enabled bool, avaliable bool, ) *TfaSettingsDto`

NewTfaSettingsDto instantiates a new TfaSettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTfaSettingsDtoWithDefaults

`func NewTfaSettingsDtoWithDefaults() *TfaSettingsDto`

NewTfaSettingsDtoWithDefaults instantiates a new TfaSettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TfaSettingsDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TfaSettingsDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TfaSettingsDto) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *TfaSettingsDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TfaSettingsDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTitle

`func (o *TfaSettingsDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TfaSettingsDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TfaSettingsDto) SetTitle(v string)`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *TfaSettingsDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *TfaSettingsDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetEnabled

`func (o *TfaSettingsDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TfaSettingsDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *TfaSettingsDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAvaliable

`func (o *TfaSettingsDto) GetAvaliable() bool`

GetAvaliable returns the Avaliable field if non-nil, zero value otherwise.

### GetAvaliableOk

`func (o *TfaSettingsDto) GetAvaliableOk() (*bool, bool)`

GetAvaliableOk returns a tuple with the Avaliable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvaliable

`func (o *TfaSettingsDto) SetAvaliable(v bool)`

SetAvaliable sets Avaliable field to given value.


### GetTrustedIps

`func (o *TfaSettingsDto) GetTrustedIps() []string`

GetTrustedIps returns the TrustedIps field if non-nil, zero value otherwise.

### GetTrustedIpsOk

`func (o *TfaSettingsDto) GetTrustedIpsOk() (*[]string, bool)`

GetTrustedIpsOk returns a tuple with the TrustedIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedIps

`func (o *TfaSettingsDto) SetTrustedIps(v []string)`

SetTrustedIps sets TrustedIps field to given value.

### HasTrustedIps

`func (o *TfaSettingsDto) HasTrustedIps() bool`

HasTrustedIps returns a boolean if a field has been set.

### SetTrustedIpsNil

`func (o *TfaSettingsDto) SetTrustedIpsNil(b bool)`

 SetTrustedIpsNil sets the value for TrustedIps to be an explicit nil

### UnsetTrustedIps
`func (o *TfaSettingsDto) UnsetTrustedIps()`

UnsetTrustedIps ensures that no value is present for TrustedIps, not even an explicit nil
### GetMandatoryUsers

`func (o *TfaSettingsDto) GetMandatoryUsers() []string`

GetMandatoryUsers returns the MandatoryUsers field if non-nil, zero value otherwise.

### GetMandatoryUsersOk

`func (o *TfaSettingsDto) GetMandatoryUsersOk() (*[]string, bool)`

GetMandatoryUsersOk returns a tuple with the MandatoryUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatoryUsers

`func (o *TfaSettingsDto) SetMandatoryUsers(v []string)`

SetMandatoryUsers sets MandatoryUsers field to given value.

### HasMandatoryUsers

`func (o *TfaSettingsDto) HasMandatoryUsers() bool`

HasMandatoryUsers returns a boolean if a field has been set.

### SetMandatoryUsersNil

`func (o *TfaSettingsDto) SetMandatoryUsersNil(b bool)`

 SetMandatoryUsersNil sets the value for MandatoryUsers to be an explicit nil

### UnsetMandatoryUsers
`func (o *TfaSettingsDto) UnsetMandatoryUsers()`

UnsetMandatoryUsers ensures that no value is present for MandatoryUsers, not even an explicit nil
### GetMandatoryGroups

`func (o *TfaSettingsDto) GetMandatoryGroups() []string`

GetMandatoryGroups returns the MandatoryGroups field if non-nil, zero value otherwise.

### GetMandatoryGroupsOk

`func (o *TfaSettingsDto) GetMandatoryGroupsOk() (*[]string, bool)`

GetMandatoryGroupsOk returns a tuple with the MandatoryGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatoryGroups

`func (o *TfaSettingsDto) SetMandatoryGroups(v []string)`

SetMandatoryGroups sets MandatoryGroups field to given value.

### HasMandatoryGroups

`func (o *TfaSettingsDto) HasMandatoryGroups() bool`

HasMandatoryGroups returns a boolean if a field has been set.

### SetMandatoryGroupsNil

`func (o *TfaSettingsDto) SetMandatoryGroupsNil(b bool)`

 SetMandatoryGroupsNil sets the value for MandatoryGroups to be an explicit nil

### UnsetMandatoryGroups
`func (o *TfaSettingsDto) UnsetMandatoryGroups()`

UnsetMandatoryGroups ensures that no value is present for MandatoryGroups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


