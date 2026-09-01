# MailDomainSettingsRequestsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**TenantTrustedDomainsType**](TenantTrustedDomainsType.md) | Defines how trusted domains are handled and validated. | 
**Domains** | **[]string** | The list of authorized email domains that are considered trusted. | 
**InviteUsersAsVisitors** | **bool** | Specifies the default permission level for the invited users (visitors or not). | 

## Methods

### NewMailDomainSettingsRequestsDto

`func NewMailDomainSettingsRequestsDto(type_ TenantTrustedDomainsType, domains []string, inviteUsersAsVisitors bool, ) *MailDomainSettingsRequestsDto`

NewMailDomainSettingsRequestsDto instantiates a new MailDomainSettingsRequestsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailDomainSettingsRequestsDtoWithDefaults

`func NewMailDomainSettingsRequestsDtoWithDefaults() *MailDomainSettingsRequestsDto`

NewMailDomainSettingsRequestsDtoWithDefaults instantiates a new MailDomainSettingsRequestsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MailDomainSettingsRequestsDto) GetType() TenantTrustedDomainsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MailDomainSettingsRequestsDto) GetTypeOk() (*TenantTrustedDomainsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MailDomainSettingsRequestsDto) SetType(v TenantTrustedDomainsType)`

SetType sets Type field to given value.


### GetDomains

`func (o *MailDomainSettingsRequestsDto) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *MailDomainSettingsRequestsDto) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *MailDomainSettingsRequestsDto) SetDomains(v []string)`

SetDomains sets Domains field to given value.


### SetDomainsNil

`func (o *MailDomainSettingsRequestsDto) SetDomainsNil(b bool)`

 SetDomainsNil sets the value for Domains to be an explicit nil

### UnsetDomains
`func (o *MailDomainSettingsRequestsDto) UnsetDomains()`

UnsetDomains ensures that no value is present for Domains, not even an explicit nil
### GetInviteUsersAsVisitors

`func (o *MailDomainSettingsRequestsDto) GetInviteUsersAsVisitors() bool`

GetInviteUsersAsVisitors returns the InviteUsersAsVisitors field if non-nil, zero value otherwise.

### GetInviteUsersAsVisitorsOk

`func (o *MailDomainSettingsRequestsDto) GetInviteUsersAsVisitorsOk() (*bool, bool)`

GetInviteUsersAsVisitorsOk returns a tuple with the InviteUsersAsVisitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteUsersAsVisitors

`func (o *MailDomainSettingsRequestsDto) SetInviteUsersAsVisitors(v bool)`

SetInviteUsersAsVisitors sets InviteUsersAsVisitors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


