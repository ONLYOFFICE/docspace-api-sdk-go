# AuditTrailActionMapperDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageAction** | Pointer to **NullableString** | The action name to send as the `action` filter of `GET api/2.0/security/audit/events/filter`, and the value  that comes back as `actionId` on an event. | [optional] 
**ActionType** | Pointer to **NullableString** | The kind of change the action makes, accepted by the `actionType` filter of the same operation. | [optional] 
**Entity** | Pointer to **NullableString** | The kind of object the action applies to, accepted by the `entryType` filter. It is `None` for an action  that targets no object, such as a settings change, and an action with a second object type reports only the  first one here. | [optional] 

## Methods

### NewAuditTrailActionMapperDto

`func NewAuditTrailActionMapperDto() *AuditTrailActionMapperDto`

NewAuditTrailActionMapperDto instantiates a new AuditTrailActionMapperDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditTrailActionMapperDtoWithDefaults

`func NewAuditTrailActionMapperDtoWithDefaults() *AuditTrailActionMapperDto`

NewAuditTrailActionMapperDtoWithDefaults instantiates a new AuditTrailActionMapperDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageAction

`func (o *AuditTrailActionMapperDto) GetMessageAction() string`

GetMessageAction returns the MessageAction field if non-nil, zero value otherwise.

### GetMessageActionOk

`func (o *AuditTrailActionMapperDto) GetMessageActionOk() (*string, bool)`

GetMessageActionOk returns a tuple with the MessageAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageAction

`func (o *AuditTrailActionMapperDto) SetMessageAction(v string)`

SetMessageAction sets MessageAction field to given value.

### HasMessageAction

`func (o *AuditTrailActionMapperDto) HasMessageAction() bool`

HasMessageAction returns a boolean if a field has been set.

### SetMessageActionNil

`func (o *AuditTrailActionMapperDto) SetMessageActionNil(b bool)`

 SetMessageActionNil sets the value for MessageAction to be an explicit nil

### UnsetMessageAction
`func (o *AuditTrailActionMapperDto) UnsetMessageAction()`

UnsetMessageAction ensures that no value is present for MessageAction, not even an explicit nil
### GetActionType

`func (o *AuditTrailActionMapperDto) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *AuditTrailActionMapperDto) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *AuditTrailActionMapperDto) SetActionType(v string)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *AuditTrailActionMapperDto) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### SetActionTypeNil

`func (o *AuditTrailActionMapperDto) SetActionTypeNil(b bool)`

 SetActionTypeNil sets the value for ActionType to be an explicit nil

### UnsetActionType
`func (o *AuditTrailActionMapperDto) UnsetActionType()`

UnsetActionType ensures that no value is present for ActionType, not even an explicit nil
### GetEntity

`func (o *AuditTrailActionMapperDto) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *AuditTrailActionMapperDto) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *AuditTrailActionMapperDto) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *AuditTrailActionMapperDto) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### SetEntityNil

`func (o *AuditTrailActionMapperDto) SetEntityNil(b bool)`

 SetEntityNil sets the value for Entity to be an explicit nil

### UnsetEntity
`func (o *AuditTrailActionMapperDto) UnsetEntity()`

UnsetEntity ensures that no value is present for Entity, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


