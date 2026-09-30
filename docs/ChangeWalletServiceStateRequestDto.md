# ChangeWalletServiceStateRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | Pointer to [**TenantWalletService**](TenantWalletService.md) | The service being switched, given by its catalogue name. Switching it on only makes it available to the  portal; its units are still bought with `PUT api/2.0/portal/payment/updatewallet`. | [optional] 
**Enabled** | Pointer to **bool** | Which way the service is switched: `true` makes it available to the portal, `false` withdraws it. Setting the  state the service already has changes nothing. | [optional] 

## Methods

### NewChangeWalletServiceStateRequestDto

`func NewChangeWalletServiceStateRequestDto() *ChangeWalletServiceStateRequestDto`

NewChangeWalletServiceStateRequestDto instantiates a new ChangeWalletServiceStateRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeWalletServiceStateRequestDtoWithDefaults

`func NewChangeWalletServiceStateRequestDtoWithDefaults() *ChangeWalletServiceStateRequestDto`

NewChangeWalletServiceStateRequestDtoWithDefaults instantiates a new ChangeWalletServiceStateRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *ChangeWalletServiceStateRequestDto) GetService() TenantWalletService`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ChangeWalletServiceStateRequestDto) GetServiceOk() (*TenantWalletService, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ChangeWalletServiceStateRequestDto) SetService(v TenantWalletService)`

SetService sets Service field to given value.

### HasService

`func (o *ChangeWalletServiceStateRequestDto) HasService() bool`

HasService returns a boolean if a field has been set.

### GetEnabled

`func (o *ChangeWalletServiceStateRequestDto) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ChangeWalletServiceStateRequestDto) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ChangeWalletServiceStateRequestDto) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ChangeWalletServiceStateRequestDto) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


