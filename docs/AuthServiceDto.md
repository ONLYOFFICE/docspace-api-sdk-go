# AuthServiceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The internal key of the provider, such as `google` or `box`. It is the `name` that  `POST api/2.0/settings/authservice` takes to select the provider. | [optional] 
**Title** | Pointer to **NullableString** | The provider name as it is shown in the interface. | [optional] 
**Description** | Pointer to **NullableString** | A sentence about what connecting the provider gives the portal, shown next to it in the interface. | [optional] 
**Instruction** | Pointer to **NullableString** | The steps an administrator has to take on the provider side to obtain the keys, shown in the interface. | [optional] 
**CanSet** | Pointer to **bool** | Whether this provider accepts keys through the API at all. A provider whose keys are fixed by the  installation reports `false`, and saving keys for it is refused. | [optional] 
**Paid** | Pointer to **bool** | Whether the provider is a paid option. A paid one can only be connected while the portal plan includes  third-party storage or the installation is licensed as self-hosted. | [optional] 
**Props** | Pointer to [**[]AuthKeyDto**](AuthKeyDto.md) | The keys the provider defines, with the values last saved and how the settings form shows each of them.  It is `null` for a provider that forbids changes (`canSet` is `false`): its keys are not read at all. | [optional] 

## Methods

### NewAuthServiceDto

`func NewAuthServiceDto() *AuthServiceDto`

NewAuthServiceDto instantiates a new AuthServiceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthServiceDtoWithDefaults

`func NewAuthServiceDtoWithDefaults() *AuthServiceDto`

NewAuthServiceDtoWithDefaults instantiates a new AuthServiceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AuthServiceDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthServiceDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthServiceDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthServiceDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AuthServiceDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AuthServiceDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTitle

`func (o *AuthServiceDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AuthServiceDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AuthServiceDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AuthServiceDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *AuthServiceDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *AuthServiceDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *AuthServiceDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthServiceDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthServiceDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthServiceDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AuthServiceDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AuthServiceDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInstruction

`func (o *AuthServiceDto) GetInstruction() string`

GetInstruction returns the Instruction field if non-nil, zero value otherwise.

### GetInstructionOk

`func (o *AuthServiceDto) GetInstructionOk() (*string, bool)`

GetInstructionOk returns a tuple with the Instruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstruction

`func (o *AuthServiceDto) SetInstruction(v string)`

SetInstruction sets Instruction field to given value.

### HasInstruction

`func (o *AuthServiceDto) HasInstruction() bool`

HasInstruction returns a boolean if a field has been set.

### SetInstructionNil

`func (o *AuthServiceDto) SetInstructionNil(b bool)`

 SetInstructionNil sets the value for Instruction to be an explicit nil

### UnsetInstruction
`func (o *AuthServiceDto) UnsetInstruction()`

UnsetInstruction ensures that no value is present for Instruction, not even an explicit nil
### GetCanSet

`func (o *AuthServiceDto) GetCanSet() bool`

GetCanSet returns the CanSet field if non-nil, zero value otherwise.

### GetCanSetOk

`func (o *AuthServiceDto) GetCanSetOk() (*bool, bool)`

GetCanSetOk returns a tuple with the CanSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSet

`func (o *AuthServiceDto) SetCanSet(v bool)`

SetCanSet sets CanSet field to given value.

### HasCanSet

`func (o *AuthServiceDto) HasCanSet() bool`

HasCanSet returns a boolean if a field has been set.

### GetPaid

`func (o *AuthServiceDto) GetPaid() bool`

GetPaid returns the Paid field if non-nil, zero value otherwise.

### GetPaidOk

`func (o *AuthServiceDto) GetPaidOk() (*bool, bool)`

GetPaidOk returns a tuple with the Paid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaid

`func (o *AuthServiceDto) SetPaid(v bool)`

SetPaid sets Paid field to given value.

### HasPaid

`func (o *AuthServiceDto) HasPaid() bool`

HasPaid returns a boolean if a field has been set.

### GetProps

`func (o *AuthServiceDto) GetProps() []AuthKeyDto`

GetProps returns the Props field if non-nil, zero value otherwise.

### GetPropsOk

`func (o *AuthServiceDto) GetPropsOk() (*[]AuthKeyDto, bool)`

GetPropsOk returns a tuple with the Props field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProps

`func (o *AuthServiceDto) SetProps(v []AuthKeyDto)`

SetProps sets Props field to given value.

### HasProps

`func (o *AuthServiceDto) HasProps() bool`

HasProps returns a boolean if a field has been set.

### SetPropsNil

`func (o *AuthServiceDto) SetPropsNil(b bool)`

 SetPropsNil sets the value for Props to be an explicit nil

### UnsetProps
`func (o *AuthServiceDto) UnsetProps()`

UnsetProps ensures that no value is present for Props, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


