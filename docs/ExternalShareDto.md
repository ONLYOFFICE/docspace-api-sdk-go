# ExternalShareDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**Status**](Status.md) |  | 
**Id** | Pointer to **NullableString** | The external data ID. | [optional] 
**Title** | Pointer to **NullableString** | The external data title. | [optional] 
**Type** | Pointer to [**FileEntryType**](FileEntryType.md) |  | [optional] 
**TenantId** | **int32** | The tenant ID. | 
**EntityId** | Pointer to **NullableString** | The unique identifier of the shared entity. | [optional] 
**EntityTitle** | Pointer to **NullableString** | The title of the shared entity. | [optional] 
**EntityType** | Pointer to [**FileEntryType**](FileEntryType.md) |  | [optional] 
**IsRoom** | Pointer to **NullableBool** | Indicates whether the entity represents a room. | [optional] 
**Shared** | **bool** | Specifies whether to share the external data or not. | 
**LinkId** | **string** | The link ID of the external data. | 
**IsAuthenticated** | **bool** | Specifies whether the user is authenticated or not. | 
**IsRoomMember** | Pointer to **bool** | The room ID of the external data. | [optional] 

## Methods

### NewExternalShareDto

`func NewExternalShareDto(status Status, tenantId int32, shared bool, linkId string, isAuthenticated bool, ) *ExternalShareDto`

NewExternalShareDto instantiates a new ExternalShareDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalShareDtoWithDefaults

`func NewExternalShareDtoWithDefaults() *ExternalShareDto`

NewExternalShareDtoWithDefaults instantiates a new ExternalShareDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ExternalShareDto) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExternalShareDto) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExternalShareDto) SetStatus(v Status)`

SetStatus sets Status field to given value.


### GetId

`func (o *ExternalShareDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalShareDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalShareDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExternalShareDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ExternalShareDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ExternalShareDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTitle

`func (o *ExternalShareDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ExternalShareDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ExternalShareDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ExternalShareDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ExternalShareDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ExternalShareDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetType

`func (o *ExternalShareDto) GetType() FileEntryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExternalShareDto) GetTypeOk() (*FileEntryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExternalShareDto) SetType(v FileEntryType)`

SetType sets Type field to given value.

### HasType

`func (o *ExternalShareDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTenantId

`func (o *ExternalShareDto) GetTenantId() int32`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ExternalShareDto) GetTenantIdOk() (*int32, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ExternalShareDto) SetTenantId(v int32)`

SetTenantId sets TenantId field to given value.


### GetEntityId

`func (o *ExternalShareDto) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *ExternalShareDto) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *ExternalShareDto) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *ExternalShareDto) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### SetEntityIdNil

`func (o *ExternalShareDto) SetEntityIdNil(b bool)`

 SetEntityIdNil sets the value for EntityId to be an explicit nil

### UnsetEntityId
`func (o *ExternalShareDto) UnsetEntityId()`

UnsetEntityId ensures that no value is present for EntityId, not even an explicit nil
### GetEntityTitle

`func (o *ExternalShareDto) GetEntityTitle() string`

GetEntityTitle returns the EntityTitle field if non-nil, zero value otherwise.

### GetEntityTitleOk

`func (o *ExternalShareDto) GetEntityTitleOk() (*string, bool)`

GetEntityTitleOk returns a tuple with the EntityTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityTitle

`func (o *ExternalShareDto) SetEntityTitle(v string)`

SetEntityTitle sets EntityTitle field to given value.

### HasEntityTitle

`func (o *ExternalShareDto) HasEntityTitle() bool`

HasEntityTitle returns a boolean if a field has been set.

### SetEntityTitleNil

`func (o *ExternalShareDto) SetEntityTitleNil(b bool)`

 SetEntityTitleNil sets the value for EntityTitle to be an explicit nil

### UnsetEntityTitle
`func (o *ExternalShareDto) UnsetEntityTitle()`

UnsetEntityTitle ensures that no value is present for EntityTitle, not even an explicit nil
### GetEntityType

`func (o *ExternalShareDto) GetEntityType() FileEntryType`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *ExternalShareDto) GetEntityTypeOk() (*FileEntryType, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *ExternalShareDto) SetEntityType(v FileEntryType)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *ExternalShareDto) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetIsRoom

`func (o *ExternalShareDto) GetIsRoom() bool`

GetIsRoom returns the IsRoom field if non-nil, zero value otherwise.

### GetIsRoomOk

`func (o *ExternalShareDto) GetIsRoomOk() (*bool, bool)`

GetIsRoomOk returns a tuple with the IsRoom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRoom

`func (o *ExternalShareDto) SetIsRoom(v bool)`

SetIsRoom sets IsRoom field to given value.

### HasIsRoom

`func (o *ExternalShareDto) HasIsRoom() bool`

HasIsRoom returns a boolean if a field has been set.

### SetIsRoomNil

`func (o *ExternalShareDto) SetIsRoomNil(b bool)`

 SetIsRoomNil sets the value for IsRoom to be an explicit nil

### UnsetIsRoom
`func (o *ExternalShareDto) UnsetIsRoom()`

UnsetIsRoom ensures that no value is present for IsRoom, not even an explicit nil
### GetShared

`func (o *ExternalShareDto) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *ExternalShareDto) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *ExternalShareDto) SetShared(v bool)`

SetShared sets Shared field to given value.


### GetLinkId

`func (o *ExternalShareDto) GetLinkId() string`

GetLinkId returns the LinkId field if non-nil, zero value otherwise.

### GetLinkIdOk

`func (o *ExternalShareDto) GetLinkIdOk() (*string, bool)`

GetLinkIdOk returns a tuple with the LinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkId

`func (o *ExternalShareDto) SetLinkId(v string)`

SetLinkId sets LinkId field to given value.


### GetIsAuthenticated

`func (o *ExternalShareDto) GetIsAuthenticated() bool`

GetIsAuthenticated returns the IsAuthenticated field if non-nil, zero value otherwise.

### GetIsAuthenticatedOk

`func (o *ExternalShareDto) GetIsAuthenticatedOk() (*bool, bool)`

GetIsAuthenticatedOk returns a tuple with the IsAuthenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAuthenticated

`func (o *ExternalShareDto) SetIsAuthenticated(v bool)`

SetIsAuthenticated sets IsAuthenticated field to given value.


### GetIsRoomMember

`func (o *ExternalShareDto) GetIsRoomMember() bool`

GetIsRoomMember returns the IsRoomMember field if non-nil, zero value otherwise.

### GetIsRoomMemberOk

`func (o *ExternalShareDto) GetIsRoomMemberOk() (*bool, bool)`

GetIsRoomMemberOk returns a tuple with the IsRoomMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRoomMember

`func (o *ExternalShareDto) SetIsRoomMember(v bool)`

SetIsRoomMember sets IsRoomMember field to given value.

### HasIsRoomMember

`func (o *ExternalShareDto) HasIsRoomMember() bool`

HasIsRoomMember returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


