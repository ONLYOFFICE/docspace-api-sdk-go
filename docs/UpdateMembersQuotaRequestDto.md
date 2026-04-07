# UpdateMembersQuotaRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserIds** | Pointer to **[]string** | The list of user IDs. | [optional] 
**Quota** | Pointer to [**UpdateMembersQuotaRequestDtoQuota**](UpdateMembersQuotaRequestDtoQuota.md) |  | [optional] 

## Methods

### NewUpdateMembersQuotaRequestDto

`func NewUpdateMembersQuotaRequestDto() *UpdateMembersQuotaRequestDto`

NewUpdateMembersQuotaRequestDto instantiates a new UpdateMembersQuotaRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMembersQuotaRequestDtoWithDefaults

`func NewUpdateMembersQuotaRequestDtoWithDefaults() *UpdateMembersQuotaRequestDto`

NewUpdateMembersQuotaRequestDtoWithDefaults instantiates a new UpdateMembersQuotaRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserIds

`func (o *UpdateMembersQuotaRequestDto) GetUserIds() []string`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *UpdateMembersQuotaRequestDto) GetUserIdsOk() (*[]string, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *UpdateMembersQuotaRequestDto) SetUserIds(v []string)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *UpdateMembersQuotaRequestDto) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.

### SetUserIdsNil

`func (o *UpdateMembersQuotaRequestDto) SetUserIdsNil(b bool)`

 SetUserIdsNil sets the value for UserIds to be an explicit nil

### UnsetUserIds
`func (o *UpdateMembersQuotaRequestDto) UnsetUserIds()`

UnsetUserIds ensures that no value is present for UserIds, not even an explicit nil
### GetQuota

`func (o *UpdateMembersQuotaRequestDto) GetQuota() UpdateMembersQuotaRequestDtoQuota`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *UpdateMembersQuotaRequestDto) GetQuotaOk() (*UpdateMembersQuotaRequestDtoQuota, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *UpdateMembersQuotaRequestDto) SetQuota(v UpdateMembersQuotaRequestDtoQuota)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *UpdateMembersQuotaRequestDto) HasQuota() bool`

HasQuota returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


