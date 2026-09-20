# TfaRequestsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TfaRequestsDtoType**](TfaRequestsDtoType.md) | The second factor the portal demands. The two methods are mutually exclusive, so switching one on switches  the other off, and any value outside the defined set is read as switching TFA off rather than refused. | [optional] 
**Id** | Pointer to **string** | The account the request concerns, by portal user ID. Naming the portal owner is refused unless it is the  caller's own account. Where an operation detaches an authenticator application, the empty GUID and the  caller's own ID both mean the caller. | [optional] 
**TrustedIps** | Pointer to **[]string** | The list of IP addresses that bypass TFA verification. Each entry is a single address, an inclusive  from-to range or a CIDR block. This is the whole list that is to hold afterwards, so send the addresses  already trusted along with a new one; an entry that cannot be parsed fails the call with 400, and accounts  named as mandatory still have to pass the challenge even from a trusted address. | [optional] 
**MandatoryUsers** | Pointer to **[]string** | The accounts that must pass the challenge whatever their address, by portal user ID. This is the whole list  that is to hold afterwards - leaving it out clears it rather than keeping it - and naming the portal owner is  refused unless the caller is the owner. | [optional] 
**MandatoryGroups** | Pointer to **[]string** | The groups whose members must pass the challenge whatever their address, by group ID. This is the whole list  that is to hold afterwards - leaving it out clears it rather than keeping it. | [optional] 

## Methods

### NewTfaRequestsDto

`func NewTfaRequestsDto() *TfaRequestsDto`

NewTfaRequestsDto instantiates a new TfaRequestsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTfaRequestsDtoWithDefaults

`func NewTfaRequestsDtoWithDefaults() *TfaRequestsDto`

NewTfaRequestsDtoWithDefaults instantiates a new TfaRequestsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TfaRequestsDto) GetType() TfaRequestsDtoType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TfaRequestsDto) GetTypeOk() (*TfaRequestsDtoType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TfaRequestsDto) SetType(v TfaRequestsDtoType)`

SetType sets Type field to given value.

### HasType

`func (o *TfaRequestsDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *TfaRequestsDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TfaRequestsDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TfaRequestsDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TfaRequestsDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTrustedIps

`func (o *TfaRequestsDto) GetTrustedIps() []string`

GetTrustedIps returns the TrustedIps field if non-nil, zero value otherwise.

### GetTrustedIpsOk

`func (o *TfaRequestsDto) GetTrustedIpsOk() (*[]string, bool)`

GetTrustedIpsOk returns a tuple with the TrustedIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedIps

`func (o *TfaRequestsDto) SetTrustedIps(v []string)`

SetTrustedIps sets TrustedIps field to given value.

### HasTrustedIps

`func (o *TfaRequestsDto) HasTrustedIps() bool`

HasTrustedIps returns a boolean if a field has been set.

### SetTrustedIpsNil

`func (o *TfaRequestsDto) SetTrustedIpsNil(b bool)`

 SetTrustedIpsNil sets the value for TrustedIps to be an explicit nil

### UnsetTrustedIps
`func (o *TfaRequestsDto) UnsetTrustedIps()`

UnsetTrustedIps ensures that no value is present for TrustedIps, not even an explicit nil
### GetMandatoryUsers

`func (o *TfaRequestsDto) GetMandatoryUsers() []string`

GetMandatoryUsers returns the MandatoryUsers field if non-nil, zero value otherwise.

### GetMandatoryUsersOk

`func (o *TfaRequestsDto) GetMandatoryUsersOk() (*[]string, bool)`

GetMandatoryUsersOk returns a tuple with the MandatoryUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatoryUsers

`func (o *TfaRequestsDto) SetMandatoryUsers(v []string)`

SetMandatoryUsers sets MandatoryUsers field to given value.

### HasMandatoryUsers

`func (o *TfaRequestsDto) HasMandatoryUsers() bool`

HasMandatoryUsers returns a boolean if a field has been set.

### SetMandatoryUsersNil

`func (o *TfaRequestsDto) SetMandatoryUsersNil(b bool)`

 SetMandatoryUsersNil sets the value for MandatoryUsers to be an explicit nil

### UnsetMandatoryUsers
`func (o *TfaRequestsDto) UnsetMandatoryUsers()`

UnsetMandatoryUsers ensures that no value is present for MandatoryUsers, not even an explicit nil
### GetMandatoryGroups

`func (o *TfaRequestsDto) GetMandatoryGroups() []string`

GetMandatoryGroups returns the MandatoryGroups field if non-nil, zero value otherwise.

### GetMandatoryGroupsOk

`func (o *TfaRequestsDto) GetMandatoryGroupsOk() (*[]string, bool)`

GetMandatoryGroupsOk returns a tuple with the MandatoryGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatoryGroups

`func (o *TfaRequestsDto) SetMandatoryGroups(v []string)`

SetMandatoryGroups sets MandatoryGroups field to given value.

### HasMandatoryGroups

`func (o *TfaRequestsDto) HasMandatoryGroups() bool`

HasMandatoryGroups returns a boolean if a field has been set.

### SetMandatoryGroupsNil

`func (o *TfaRequestsDto) SetMandatoryGroupsNil(b bool)`

 SetMandatoryGroupsNil sets the value for MandatoryGroups to be an explicit nil

### UnsetMandatoryGroups
`func (o *TfaRequestsDto) UnsetMandatoryGroups()`

UnsetMandatoryGroups ensures that no value is present for MandatoryGroups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


