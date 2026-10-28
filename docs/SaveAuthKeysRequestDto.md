# SaveAuthKeysRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The internal key of the provider, such as `google` or `box`. Take it from the `name` of  `GET api/2.0/settings/authservice`. | [optional] 
**Title** | Pointer to **NullableString** | Accepted for compatibility with earlier clients and not read: the server keeps its own value. | [optional] 
**Description** | Pointer to **NullableString** | Accepted for compatibility with earlier clients and not read: the server keeps its own value. | [optional] 
**Instruction** | Pointer to **NullableString** | Accepted for compatibility with earlier clients and not read: the server keeps its own value. | [optional] 
**CanSet** | Pointer to **bool** | Accepted for compatibility with earlier clients and not read: the server keeps its own value. | [optional] 
**Paid** | Pointer to **bool** | Accepted for compatibility with earlier clients and not read: the server keeps its own value. | [optional] 
**Props** | Pointer to [**[]AuthKeyRequest**](AuthKeyRequest.md) | The keys of the provider with their new values, by the key names `GET api/2.0/settings/authservice` lists in  `props`. | [optional] 

## Methods

### NewSaveAuthKeysRequestDto

`func NewSaveAuthKeysRequestDto() *SaveAuthKeysRequestDto`

NewSaveAuthKeysRequestDto instantiates a new SaveAuthKeysRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveAuthKeysRequestDtoWithDefaults

`func NewSaveAuthKeysRequestDtoWithDefaults() *SaveAuthKeysRequestDto`

NewSaveAuthKeysRequestDtoWithDefaults instantiates a new SaveAuthKeysRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SaveAuthKeysRequestDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SaveAuthKeysRequestDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SaveAuthKeysRequestDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SaveAuthKeysRequestDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SaveAuthKeysRequestDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SaveAuthKeysRequestDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTitle

`func (o *SaveAuthKeysRequestDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SaveAuthKeysRequestDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SaveAuthKeysRequestDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SaveAuthKeysRequestDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *SaveAuthKeysRequestDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SaveAuthKeysRequestDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *SaveAuthKeysRequestDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SaveAuthKeysRequestDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SaveAuthKeysRequestDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SaveAuthKeysRequestDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SaveAuthKeysRequestDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SaveAuthKeysRequestDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInstruction

`func (o *SaveAuthKeysRequestDto) GetInstruction() string`

GetInstruction returns the Instruction field if non-nil, zero value otherwise.

### GetInstructionOk

`func (o *SaveAuthKeysRequestDto) GetInstructionOk() (*string, bool)`

GetInstructionOk returns a tuple with the Instruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstruction

`func (o *SaveAuthKeysRequestDto) SetInstruction(v string)`

SetInstruction sets Instruction field to given value.

### HasInstruction

`func (o *SaveAuthKeysRequestDto) HasInstruction() bool`

HasInstruction returns a boolean if a field has been set.

### SetInstructionNil

`func (o *SaveAuthKeysRequestDto) SetInstructionNil(b bool)`

 SetInstructionNil sets the value for Instruction to be an explicit nil

### UnsetInstruction
`func (o *SaveAuthKeysRequestDto) UnsetInstruction()`

UnsetInstruction ensures that no value is present for Instruction, not even an explicit nil
### GetCanSet

`func (o *SaveAuthKeysRequestDto) GetCanSet() bool`

GetCanSet returns the CanSet field if non-nil, zero value otherwise.

### GetCanSetOk

`func (o *SaveAuthKeysRequestDto) GetCanSetOk() (*bool, bool)`

GetCanSetOk returns a tuple with the CanSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSet

`func (o *SaveAuthKeysRequestDto) SetCanSet(v bool)`

SetCanSet sets CanSet field to given value.

### HasCanSet

`func (o *SaveAuthKeysRequestDto) HasCanSet() bool`

HasCanSet returns a boolean if a field has been set.

### GetPaid

`func (o *SaveAuthKeysRequestDto) GetPaid() bool`

GetPaid returns the Paid field if non-nil, zero value otherwise.

### GetPaidOk

`func (o *SaveAuthKeysRequestDto) GetPaidOk() (*bool, bool)`

GetPaidOk returns a tuple with the Paid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaid

`func (o *SaveAuthKeysRequestDto) SetPaid(v bool)`

SetPaid sets Paid field to given value.

### HasPaid

`func (o *SaveAuthKeysRequestDto) HasPaid() bool`

HasPaid returns a boolean if a field has been set.

### GetProps

`func (o *SaveAuthKeysRequestDto) GetProps() []AuthKeyRequest`

GetProps returns the Props field if non-nil, zero value otherwise.

### GetPropsOk

`func (o *SaveAuthKeysRequestDto) GetPropsOk() (*[]AuthKeyRequest, bool)`

GetPropsOk returns a tuple with the Props field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProps

`func (o *SaveAuthKeysRequestDto) SetProps(v []AuthKeyRequest)`

SetProps sets Props field to given value.

### HasProps

`func (o *SaveAuthKeysRequestDto) HasProps() bool`

HasProps returns a boolean if a field has been set.

### SetPropsNil

`func (o *SaveAuthKeysRequestDto) SetPropsNil(b bool)`

 SetPropsNil sets the value for Props to be an explicit nil

### UnsetProps
`func (o *SaveAuthKeysRequestDto) UnsetProps()`

UnsetProps ensures that no value is present for Props, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


