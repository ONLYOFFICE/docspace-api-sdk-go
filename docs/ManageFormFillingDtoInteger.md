# ManageFormFillingDtoInteger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FormId** | **int32** | The ID of the form to manage. | 
**Action** | Pointer to [**FormFillingManageAction**](FormFillingManageAction.md) |  | [optional] 

## Methods

### NewManageFormFillingDtoInteger

`func NewManageFormFillingDtoInteger(formId int32, ) *ManageFormFillingDtoInteger`

NewManageFormFillingDtoInteger instantiates a new ManageFormFillingDtoInteger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManageFormFillingDtoIntegerWithDefaults

`func NewManageFormFillingDtoIntegerWithDefaults() *ManageFormFillingDtoInteger`

NewManageFormFillingDtoIntegerWithDefaults instantiates a new ManageFormFillingDtoInteger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormId

`func (o *ManageFormFillingDtoInteger) GetFormId() int32`

GetFormId returns the FormId field if non-nil, zero value otherwise.

### GetFormIdOk

`func (o *ManageFormFillingDtoInteger) GetFormIdOk() (*int32, bool)`

GetFormIdOk returns a tuple with the FormId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormId

`func (o *ManageFormFillingDtoInteger) SetFormId(v int32)`

SetFormId sets FormId field to given value.


### GetAction

`func (o *ManageFormFillingDtoInteger) GetAction() FormFillingManageAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ManageFormFillingDtoInteger) GetActionOk() (*FormFillingManageAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ManageFormFillingDtoInteger) SetAction(v FormFillingManageAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *ManageFormFillingDtoInteger) HasAction() bool`

HasAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


