# DiscountCategoryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The discount category unique identifier. | [optional] 
**ValueDiscount** | Pointer to **float64** | The discount value. | [optional] 
**Description** | Pointer to **NullableString** | The discount category description. | [optional] 
**Created** | Pointer to **time.Time** | The date and time when the discount category was created. | [optional] 

## Methods

### NewDiscountCategoryDto

`func NewDiscountCategoryDto() *DiscountCategoryDto`

NewDiscountCategoryDto instantiates a new DiscountCategoryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscountCategoryDtoWithDefaults

`func NewDiscountCategoryDtoWithDefaults() *DiscountCategoryDto`

NewDiscountCategoryDtoWithDefaults instantiates a new DiscountCategoryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DiscountCategoryDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DiscountCategoryDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DiscountCategoryDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DiscountCategoryDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetValueDiscount

`func (o *DiscountCategoryDto) GetValueDiscount() float64`

GetValueDiscount returns the ValueDiscount field if non-nil, zero value otherwise.

### GetValueDiscountOk

`func (o *DiscountCategoryDto) GetValueDiscountOk() (*float64, bool)`

GetValueDiscountOk returns a tuple with the ValueDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueDiscount

`func (o *DiscountCategoryDto) SetValueDiscount(v float64)`

SetValueDiscount sets ValueDiscount field to given value.

### HasValueDiscount

`func (o *DiscountCategoryDto) HasValueDiscount() bool`

HasValueDiscount returns a boolean if a field has been set.

### GetDescription

`func (o *DiscountCategoryDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DiscountCategoryDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DiscountCategoryDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DiscountCategoryDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DiscountCategoryDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DiscountCategoryDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCreated

`func (o *DiscountCategoryDto) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DiscountCategoryDto) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DiscountCategoryDto) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *DiscountCategoryDto) HasCreated() bool`

HasCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


