# TfaSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** | Which method this entry describes: `sms` for a code sent by text message, `app` for a code from an  authenticator application. It is the value `PUT api/2.0/settings/tfaapp` takes as its `type`, and no other  value ever appears here. | 
**Title** | **NullableString** | The label for the method in the portal language, meant for a button or a radio option. It is not stable  enough to branch on - match `id` for that. | 
**Enabled** | **bool** | Whether this method is the portal's current policy. At most one entry can have it set, and none has it  while the portal challenges nobody. It says nothing about the caller's own account, which may be exempt  through `trustedIps` or forced through `mandatoryUsers`. | 
**Available** | **bool** | Whether the method could be switched on at all. For `sms` it is `false` until the installation has a  working SMS provider, so a method can be offered here and still be impossible to enable; for `app` it is  always `true`. | 
**TrustedIps** | Pointer to **[]string** | The addresses that skip the challenge, each either a single address, a `from-to` pair or a CIDR range. It  is empty when no address is exempt, which means every account is challenged. | [optional] 
**MandatoryUsers** | Pointer to **[]string** | The accounts that are challenged even from a trusted address, by user ID. Empty means the exemption in  `trustedIps` holds for everyone. | [optional] 
**MandatoryGroups** | Pointer to **[]string** | The groups whose members are challenged even from a trusted address, by group ID, with the same reading of  an empty list as `mandatoryUsers`. | [optional] 

## Methods

### NewTfaSettingsDto

`func NewTfaSettingsDto(id NullableString, title NullableString, enabled bool, available bool, ) *TfaSettingsDto`

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


### GetAvailable

`func (o *TfaSettingsDto) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *TfaSettingsDto) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *TfaSettingsDto) SetAvailable(v bool)`

SetAvailable sets Available field to given value.


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


