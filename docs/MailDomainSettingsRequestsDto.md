# MailDomainSettingsRequestsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**TenantTrustedDomainsType**](TenantTrustedDomainsType.md) | How trusted domains are decided: no domain is trusted, every domain is, or only the ones listed in `domains`.  Only the custom mode reads `domains`; under the other two the list is ignored rather than refused. | 
**Domains** | **[]string** | The trusted domains, as bare hostnames such as `example.com` without a scheme or an `@`. This is the whole  list that is to hold afterwards and not a list of additions. Each entry is lowercased before it is stored,  and one entry that is not a valid hostname - or an empty list in the custom mode - fails the whole call  without saving anything. | 
**InviteUsersAsVisitors** | **bool** | What a user joining through a trusted domain becomes: `true` admits them as a guest, `false` as a full  member. It applies to joins made from now on and does not change anybody who has already joined. | 

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


