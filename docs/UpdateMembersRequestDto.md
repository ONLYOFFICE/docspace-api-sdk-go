# UpdateMembersRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserIds** | Pointer to **[]string** | The accounts the operation applies to. System accounts are dropped from the list without an error, and the  remaining ones are processed in the order they are given. | [optional] 
**ResendAll** | Pointer to **bool** | Reaches every pending account of the portal instead of the ones in `userIds`. It is read only by  `PUT api/2.0/people/invite` and is ignored by every other operation that binds this body. | [optional] 

## Methods

### NewUpdateMembersRequestDto

`func NewUpdateMembersRequestDto() *UpdateMembersRequestDto`

NewUpdateMembersRequestDto instantiates a new UpdateMembersRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMembersRequestDtoWithDefaults

`func NewUpdateMembersRequestDtoWithDefaults() *UpdateMembersRequestDto`

NewUpdateMembersRequestDtoWithDefaults instantiates a new UpdateMembersRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserIds

`func (o *UpdateMembersRequestDto) GetUserIds() []string`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *UpdateMembersRequestDto) GetUserIdsOk() (*[]string, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *UpdateMembersRequestDto) SetUserIds(v []string)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *UpdateMembersRequestDto) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.

### SetUserIdsNil

`func (o *UpdateMembersRequestDto) SetUserIdsNil(b bool)`

 SetUserIdsNil sets the value for UserIds to be an explicit nil

### UnsetUserIds
`func (o *UpdateMembersRequestDto) UnsetUserIds()`

UnsetUserIds ensures that no value is present for UserIds, not even an explicit nil
### GetResendAll

`func (o *UpdateMembersRequestDto) GetResendAll() bool`

GetResendAll returns the ResendAll field if non-nil, zero value otherwise.

### GetResendAllOk

`func (o *UpdateMembersRequestDto) GetResendAllOk() (*bool, bool)`

GetResendAllOk returns a tuple with the ResendAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResendAll

`func (o *UpdateMembersRequestDto) SetResendAll(v bool)`

SetResendAll sets ResendAll field to given value.

### HasResendAll

`func (o *UpdateMembersRequestDto) HasResendAll() bool`

HasResendAll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


