# AuthServiceRequestsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The provider being configured, by its internal key such as `google` or `box`. Take it from the `name` of  `GET api/2.0/settings/authservice`; it is the only field that selects the provider, and a key this  installation does not know is refused the same way a provider that forbids changes is. | [optional] 
**Title** | Pointer to **NullableString** | The provider name as it is shown in the interface. It is filled in by the portal when the providers are  listed and is ignored when keys are saved. | [optional] 
**Description** | Pointer to **NullableString** | A sentence about what connecting the provider gives the portal, shown next to it in the interface. It is  filled in by the portal and ignored when keys are saved. | [optional] 
**Instruction** | Pointer to **NullableString** | The steps an administrator has to take on the provider side to obtain the keys, shown in the interface. It is  filled in by the portal and ignored when keys are saved. | [optional] 
**CanSet** | Pointer to **bool** | Whether this provider accepts keys through the API at all. A provider whose keys are fixed by the  installation reports `false`, and saving keys for it is refused; the field is reported by the portal and  ignored on the way in. | [optional] 
**Paid** | Pointer to **bool** | Whether the provider is a paid option. A paid one can only be connected while the portal plan includes  third-party storage or the installation is licensed as self-hosted; the field is reported by the portal and  ignored on the way in. | [optional] 
**Props** | Pointer to [**[]AuthKey**](AuthKey.md) | The credentials the portal authenticates to the provider with, as the name and value pairs the provider  defines. Send the whole set the provider expects: leaving every value empty disconnects it, and a set that  fails the provider validation is cleared rather than stored half-applied. The listing operation reports the  values last saved, and a provider that forbids changes reports none at all. | [optional] 

## Methods

### NewAuthServiceRequestsDto

`func NewAuthServiceRequestsDto() *AuthServiceRequestsDto`

NewAuthServiceRequestsDto instantiates a new AuthServiceRequestsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthServiceRequestsDtoWithDefaults

`func NewAuthServiceRequestsDtoWithDefaults() *AuthServiceRequestsDto`

NewAuthServiceRequestsDtoWithDefaults instantiates a new AuthServiceRequestsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AuthServiceRequestsDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthServiceRequestsDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthServiceRequestsDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthServiceRequestsDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AuthServiceRequestsDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AuthServiceRequestsDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTitle

`func (o *AuthServiceRequestsDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AuthServiceRequestsDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AuthServiceRequestsDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AuthServiceRequestsDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *AuthServiceRequestsDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *AuthServiceRequestsDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *AuthServiceRequestsDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthServiceRequestsDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthServiceRequestsDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthServiceRequestsDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AuthServiceRequestsDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AuthServiceRequestsDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInstruction

`func (o *AuthServiceRequestsDto) GetInstruction() string`

GetInstruction returns the Instruction field if non-nil, zero value otherwise.

### GetInstructionOk

`func (o *AuthServiceRequestsDto) GetInstructionOk() (*string, bool)`

GetInstructionOk returns a tuple with the Instruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstruction

`func (o *AuthServiceRequestsDto) SetInstruction(v string)`

SetInstruction sets Instruction field to given value.

### HasInstruction

`func (o *AuthServiceRequestsDto) HasInstruction() bool`

HasInstruction returns a boolean if a field has been set.

### SetInstructionNil

`func (o *AuthServiceRequestsDto) SetInstructionNil(b bool)`

 SetInstructionNil sets the value for Instruction to be an explicit nil

### UnsetInstruction
`func (o *AuthServiceRequestsDto) UnsetInstruction()`

UnsetInstruction ensures that no value is present for Instruction, not even an explicit nil
### GetCanSet

`func (o *AuthServiceRequestsDto) GetCanSet() bool`

GetCanSet returns the CanSet field if non-nil, zero value otherwise.

### GetCanSetOk

`func (o *AuthServiceRequestsDto) GetCanSetOk() (*bool, bool)`

GetCanSetOk returns a tuple with the CanSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSet

`func (o *AuthServiceRequestsDto) SetCanSet(v bool)`

SetCanSet sets CanSet field to given value.

### HasCanSet

`func (o *AuthServiceRequestsDto) HasCanSet() bool`

HasCanSet returns a boolean if a field has been set.

### GetPaid

`func (o *AuthServiceRequestsDto) GetPaid() bool`

GetPaid returns the Paid field if non-nil, zero value otherwise.

### GetPaidOk

`func (o *AuthServiceRequestsDto) GetPaidOk() (*bool, bool)`

GetPaidOk returns a tuple with the Paid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaid

`func (o *AuthServiceRequestsDto) SetPaid(v bool)`

SetPaid sets Paid field to given value.

### HasPaid

`func (o *AuthServiceRequestsDto) HasPaid() bool`

HasPaid returns a boolean if a field has been set.

### GetProps

`func (o *AuthServiceRequestsDto) GetProps() []AuthKey`

GetProps returns the Props field if non-nil, zero value otherwise.

### GetPropsOk

`func (o *AuthServiceRequestsDto) GetPropsOk() (*[]AuthKey, bool)`

GetPropsOk returns a tuple with the Props field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProps

`func (o *AuthServiceRequestsDto) SetProps(v []AuthKey)`

SetProps sets Props field to given value.

### HasProps

`func (o *AuthServiceRequestsDto) HasProps() bool`

HasProps returns a boolean if a field has been set.

### SetPropsNil

`func (o *AuthServiceRequestsDto) SetPropsNil(b bool)`

 SetPropsNil sets the value for Props to be an explicit nil

### UnsetProps
`func (o *AuthServiceRequestsDto) UnsetProps()`

UnsetProps ensures that no value is present for Props, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


