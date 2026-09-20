# ManageFormFillingDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FormId** | **int32** | The PDF form the action applies to. This is the value the operation reads, rather than the identifier in its  route, and the two are to be sent the same. | 
**Action** | Pointer to [**FormFillingManageAction**](FormFillingManageAction.md) | The action to apply. | [optional] 

## Methods

### NewManageFormFillingDto

`func NewManageFormFillingDto(formId int32, ) *ManageFormFillingDto`

NewManageFormFillingDto instantiates a new ManageFormFillingDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManageFormFillingDtoWithDefaults

`func NewManageFormFillingDtoWithDefaults() *ManageFormFillingDto`

NewManageFormFillingDtoWithDefaults instantiates a new ManageFormFillingDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormId

`func (o *ManageFormFillingDto) GetFormId() int32`

GetFormId returns the FormId field if non-nil, zero value otherwise.

### GetFormIdOk

`func (o *ManageFormFillingDto) GetFormIdOk() (*int32, bool)`

GetFormIdOk returns a tuple with the FormId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormId

`func (o *ManageFormFillingDto) SetFormId(v int32)`

SetFormId sets FormId field to given value.


### GetAction

`func (o *ManageFormFillingDto) GetAction() FormFillingManageAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ManageFormFillingDto) GetActionOk() (*FormFillingManageAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ManageFormFillingDto) SetAction(v FormFillingManageAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *ManageFormFillingDto) HasAction() bool`

HasAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


