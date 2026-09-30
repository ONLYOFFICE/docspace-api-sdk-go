# OwnerIdSettingsRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnerId** | **string** | The member who is to become the portal owner, by user ID. They have to be an active member of this portal and  not a guest; a member who is not a DocSpace administrator yet is promoted to one as part of the transfer, so  the portal needs a paid seat for them. | 

## Methods

### NewOwnerIdSettingsRequestDto

`func NewOwnerIdSettingsRequestDto(ownerId string, ) *OwnerIdSettingsRequestDto`

NewOwnerIdSettingsRequestDto instantiates a new OwnerIdSettingsRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOwnerIdSettingsRequestDtoWithDefaults

`func NewOwnerIdSettingsRequestDtoWithDefaults() *OwnerIdSettingsRequestDto`

NewOwnerIdSettingsRequestDtoWithDefaults instantiates a new OwnerIdSettingsRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnerId

`func (o *OwnerIdSettingsRequestDto) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *OwnerIdSettingsRequestDto) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *OwnerIdSettingsRequestDto) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


