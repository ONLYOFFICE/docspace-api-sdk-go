# AuditTrailTypesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to **[]string** | Every action name the build can record, spelled as the `action` filter of  `GET api/2.0/security/audit/events/filter` and `GET api/2.0/security/audit/login/filter` expects it. It is  the whole vocabulary, not the actions this portal has recorded, and only a handful of the names are the  sign-in actions the login filter accepts. | [optional] 
**ActionTypes** | Pointer to **[]string** | The kinds of change an action can stand for, spelled as the `actionType` filter of  `GET api/2.0/security/audit/events/filter` expects it. | [optional] 
**ProductTypes** | Pointer to **[]string** | The products an action can belong to, spelled as the `productType` filter of  `GET api/2.0/security/audit/mappers` expects it. The audit trail itself cannot be filtered by product. | [optional] 
**ModuleTypes** | Pointer to **[]string** | The locations inside those products, spelled as the `moduleType` filter of  `GET api/2.0/security/audit/events/filter` and `GET api/2.0/security/audit/mappers` expects it. | [optional] 
**EntryTypes** | Pointer to **[]string** | The kinds of object an action can be applied to, spelled as the `entryType` filter of  `GET api/2.0/security/audit/events/filter` expects it. | [optional] 

## Methods

### NewAuditTrailTypesDto

`func NewAuditTrailTypesDto() *AuditTrailTypesDto`

NewAuditTrailTypesDto instantiates a new AuditTrailTypesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditTrailTypesDtoWithDefaults

`func NewAuditTrailTypesDtoWithDefaults() *AuditTrailTypesDto`

NewAuditTrailTypesDtoWithDefaults instantiates a new AuditTrailTypesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *AuditTrailTypesDto) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *AuditTrailTypesDto) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *AuditTrailTypesDto) SetActions(v []string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *AuditTrailTypesDto) HasActions() bool`

HasActions returns a boolean if a field has been set.

### SetActionsNil

`func (o *AuditTrailTypesDto) SetActionsNil(b bool)`

 SetActionsNil sets the value for Actions to be an explicit nil

### UnsetActions
`func (o *AuditTrailTypesDto) UnsetActions()`

UnsetActions ensures that no value is present for Actions, not even an explicit nil
### GetActionTypes

`func (o *AuditTrailTypesDto) GetActionTypes() []string`

GetActionTypes returns the ActionTypes field if non-nil, zero value otherwise.

### GetActionTypesOk

`func (o *AuditTrailTypesDto) GetActionTypesOk() (*[]string, bool)`

GetActionTypesOk returns a tuple with the ActionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTypes

`func (o *AuditTrailTypesDto) SetActionTypes(v []string)`

SetActionTypes sets ActionTypes field to given value.

### HasActionTypes

`func (o *AuditTrailTypesDto) HasActionTypes() bool`

HasActionTypes returns a boolean if a field has been set.

### SetActionTypesNil

`func (o *AuditTrailTypesDto) SetActionTypesNil(b bool)`

 SetActionTypesNil sets the value for ActionTypes to be an explicit nil

### UnsetActionTypes
`func (o *AuditTrailTypesDto) UnsetActionTypes()`

UnsetActionTypes ensures that no value is present for ActionTypes, not even an explicit nil
### GetProductTypes

`func (o *AuditTrailTypesDto) GetProductTypes() []string`

GetProductTypes returns the ProductTypes field if non-nil, zero value otherwise.

### GetProductTypesOk

`func (o *AuditTrailTypesDto) GetProductTypesOk() (*[]string, bool)`

GetProductTypesOk returns a tuple with the ProductTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTypes

`func (o *AuditTrailTypesDto) SetProductTypes(v []string)`

SetProductTypes sets ProductTypes field to given value.

### HasProductTypes

`func (o *AuditTrailTypesDto) HasProductTypes() bool`

HasProductTypes returns a boolean if a field has been set.

### SetProductTypesNil

`func (o *AuditTrailTypesDto) SetProductTypesNil(b bool)`

 SetProductTypesNil sets the value for ProductTypes to be an explicit nil

### UnsetProductTypes
`func (o *AuditTrailTypesDto) UnsetProductTypes()`

UnsetProductTypes ensures that no value is present for ProductTypes, not even an explicit nil
### GetModuleTypes

`func (o *AuditTrailTypesDto) GetModuleTypes() []string`

GetModuleTypes returns the ModuleTypes field if non-nil, zero value otherwise.

### GetModuleTypesOk

`func (o *AuditTrailTypesDto) GetModuleTypesOk() (*[]string, bool)`

GetModuleTypesOk returns a tuple with the ModuleTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleTypes

`func (o *AuditTrailTypesDto) SetModuleTypes(v []string)`

SetModuleTypes sets ModuleTypes field to given value.

### HasModuleTypes

`func (o *AuditTrailTypesDto) HasModuleTypes() bool`

HasModuleTypes returns a boolean if a field has been set.

### SetModuleTypesNil

`func (o *AuditTrailTypesDto) SetModuleTypesNil(b bool)`

 SetModuleTypesNil sets the value for ModuleTypes to be an explicit nil

### UnsetModuleTypes
`func (o *AuditTrailTypesDto) UnsetModuleTypes()`

UnsetModuleTypes ensures that no value is present for ModuleTypes, not even an explicit nil
### GetEntryTypes

`func (o *AuditTrailTypesDto) GetEntryTypes() []string`

GetEntryTypes returns the EntryTypes field if non-nil, zero value otherwise.

### GetEntryTypesOk

`func (o *AuditTrailTypesDto) GetEntryTypesOk() (*[]string, bool)`

GetEntryTypesOk returns a tuple with the EntryTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryTypes

`func (o *AuditTrailTypesDto) SetEntryTypes(v []string)`

SetEntryTypes sets EntryTypes field to given value.

### HasEntryTypes

`func (o *AuditTrailTypesDto) HasEntryTypes() bool`

HasEntryTypes returns a boolean if a field has been set.

### SetEntryTypesNil

`func (o *AuditTrailTypesDto) SetEntryTypesNil(b bool)`

 SetEntryTypesNil sets the value for EntryTypes to be an explicit nil

### UnsetEntryTypes
`func (o *AuditTrailTypesDto) UnsetEntryTypes()`

UnsetEntryTypes ensures that no value is present for EntryTypes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


