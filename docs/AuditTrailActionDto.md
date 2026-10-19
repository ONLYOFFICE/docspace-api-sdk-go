# AuditTrailActionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageAction** | Pointer to **NullableString** | The action name to send as the `action` filter of `GET api/2.0/security/audit/events/filter`, and the value  that comes back as `actionId` on an event. | [optional] 
**ActionType** | Pointer to **NullableString** | The kind of change the action makes, accepted by the `actionType` filter of the same operation. | [optional] 
**Entity** | Pointer to **NullableString** | The kind of object the action applies to, accepted by the `entryType` filter. It is `None` for an action  that targets no object, such as a settings change, and an action with a second object type reports only the  first one here. | [optional] 

## Methods

### NewAuditTrailActionDto

`func NewAuditTrailActionDto() *AuditTrailActionDto`

NewAuditTrailActionDto instantiates a new AuditTrailActionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditTrailActionDtoWithDefaults

`func NewAuditTrailActionDtoWithDefaults() *AuditTrailActionDto`

NewAuditTrailActionDtoWithDefaults instantiates a new AuditTrailActionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageAction

`func (o *AuditTrailActionDto) GetMessageAction() string`

GetMessageAction returns the MessageAction field if non-nil, zero value otherwise.

### GetMessageActionOk

`func (o *AuditTrailActionDto) GetMessageActionOk() (*string, bool)`

GetMessageActionOk returns a tuple with the MessageAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageAction

`func (o *AuditTrailActionDto) SetMessageAction(v string)`

SetMessageAction sets MessageAction field to given value.

### HasMessageAction

`func (o *AuditTrailActionDto) HasMessageAction() bool`

HasMessageAction returns a boolean if a field has been set.

### SetMessageActionNil

`func (o *AuditTrailActionDto) SetMessageActionNil(b bool)`

 SetMessageActionNil sets the value for MessageAction to be an explicit nil

### UnsetMessageAction
`func (o *AuditTrailActionDto) UnsetMessageAction()`

UnsetMessageAction ensures that no value is present for MessageAction, not even an explicit nil
### GetActionType

`func (o *AuditTrailActionDto) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *AuditTrailActionDto) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *AuditTrailActionDto) SetActionType(v string)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *AuditTrailActionDto) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### SetActionTypeNil

`func (o *AuditTrailActionDto) SetActionTypeNil(b bool)`

 SetActionTypeNil sets the value for ActionType to be an explicit nil

### UnsetActionType
`func (o *AuditTrailActionDto) UnsetActionType()`

UnsetActionType ensures that no value is present for ActionType, not even an explicit nil
### GetEntity

`func (o *AuditTrailActionDto) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *AuditTrailActionDto) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *AuditTrailActionDto) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *AuditTrailActionDto) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### SetEntityNil

`func (o *AuditTrailActionDto) SetEntityNil(b bool)`

 SetEntityNil sets the value for Entity to be an explicit nil

### UnsetEntity
`func (o *AuditTrailActionDto) UnsetEntity()`

UnsetEntity ensures that no value is present for Entity, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


