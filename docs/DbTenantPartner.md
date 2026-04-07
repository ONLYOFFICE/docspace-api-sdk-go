# DbTenantPartner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | Pointer to **int32** | The tenant ID. | [optional] 
**PartnerId** | Pointer to **NullableString** | The partner ID. | [optional] 
**AffiliateId** | Pointer to **NullableString** | The affiliate ID. | [optional] 
**Campaign** | Pointer to **NullableString** | The tenant partner campaign. | [optional] 

## Methods

### NewDbTenantPartner

`func NewDbTenantPartner() *DbTenantPartner`

NewDbTenantPartner instantiates a new DbTenantPartner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbTenantPartnerWithDefaults

`func NewDbTenantPartnerWithDefaults() *DbTenantPartner`

NewDbTenantPartnerWithDefaults instantiates a new DbTenantPartner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *DbTenantPartner) GetTenantId() int32`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DbTenantPartner) GetTenantIdOk() (*int32, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DbTenantPartner) SetTenantId(v int32)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DbTenantPartner) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetPartnerId

`func (o *DbTenantPartner) GetPartnerId() string`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *DbTenantPartner) GetPartnerIdOk() (*string, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *DbTenantPartner) SetPartnerId(v string)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *DbTenantPartner) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### SetPartnerIdNil

`func (o *DbTenantPartner) SetPartnerIdNil(b bool)`

 SetPartnerIdNil sets the value for PartnerId to be an explicit nil

### UnsetPartnerId
`func (o *DbTenantPartner) UnsetPartnerId()`

UnsetPartnerId ensures that no value is present for PartnerId, not even an explicit nil
### GetAffiliateId

`func (o *DbTenantPartner) GetAffiliateId() string`

GetAffiliateId returns the AffiliateId field if non-nil, zero value otherwise.

### GetAffiliateIdOk

`func (o *DbTenantPartner) GetAffiliateIdOk() (*string, bool)`

GetAffiliateIdOk returns a tuple with the AffiliateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliateId

`func (o *DbTenantPartner) SetAffiliateId(v string)`

SetAffiliateId sets AffiliateId field to given value.

### HasAffiliateId

`func (o *DbTenantPartner) HasAffiliateId() bool`

HasAffiliateId returns a boolean if a field has been set.

### SetAffiliateIdNil

`func (o *DbTenantPartner) SetAffiliateIdNil(b bool)`

 SetAffiliateIdNil sets the value for AffiliateId to be an explicit nil

### UnsetAffiliateId
`func (o *DbTenantPartner) UnsetAffiliateId()`

UnsetAffiliateId ensures that no value is present for AffiliateId, not even an explicit nil
### GetCampaign

`func (o *DbTenantPartner) GetCampaign() string`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *DbTenantPartner) GetCampaignOk() (*string, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *DbTenantPartner) SetCampaign(v string)`

SetCampaign sets Campaign field to given value.

### HasCampaign

`func (o *DbTenantPartner) HasCampaign() bool`

HasCampaign returns a boolean if a field has been set.

### SetCampaignNil

`func (o *DbTenantPartner) SetCampaignNil(b bool)`

 SetCampaignNil sets the value for Campaign to be an explicit nil

### UnsetCampaign
`func (o *DbTenantPartner) UnsetCampaign()`

UnsetCampaign ensures that no value is present for Campaign, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


