# WhiteLabelItemDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**WhiteLabelLogoType**](WhiteLabelLogoType.md) | The white label logo type. | [optional] 
**Name** | Pointer to **NullableString** | The white label file name. | [optional] 
**Size** | Pointer to [**WhiteLabelItemSizeDto**](WhiteLabelItemSizeDto.md) | The white label file size. | [optional] 
**Path** | Pointer to [**WhiteLabelItemPathDto**](WhiteLabelItemPathDto.md) | The white label file path. | [optional] 

## Methods

### NewWhiteLabelItemDto

`func NewWhiteLabelItemDto() *WhiteLabelItemDto`

NewWhiteLabelItemDto instantiates a new WhiteLabelItemDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhiteLabelItemDtoWithDefaults

`func NewWhiteLabelItemDtoWithDefaults() *WhiteLabelItemDto`

NewWhiteLabelItemDtoWithDefaults instantiates a new WhiteLabelItemDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WhiteLabelItemDto) GetType() WhiteLabelLogoType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WhiteLabelItemDto) GetTypeOk() (*WhiteLabelLogoType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WhiteLabelItemDto) SetType(v WhiteLabelLogoType)`

SetType sets Type field to given value.

### HasType

`func (o *WhiteLabelItemDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *WhiteLabelItemDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WhiteLabelItemDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WhiteLabelItemDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WhiteLabelItemDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *WhiteLabelItemDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *WhiteLabelItemDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSize

`func (o *WhiteLabelItemDto) GetSize() WhiteLabelItemSizeDto`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *WhiteLabelItemDto) GetSizeOk() (*WhiteLabelItemSizeDto, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *WhiteLabelItemDto) SetSize(v WhiteLabelItemSizeDto)`

SetSize sets Size field to given value.

### HasSize

`func (o *WhiteLabelItemDto) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetPath

`func (o *WhiteLabelItemDto) GetPath() WhiteLabelItemPathDto`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *WhiteLabelItemDto) GetPathOk() (*WhiteLabelItemPathDto, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *WhiteLabelItemDto) SetPath(v WhiteLabelItemPathDto)`

SetPath sets Path field to given value.

### HasPath

`func (o *WhiteLabelItemDto) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


