# TenantQuotaFeatureDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The stable key of the feature - `total_size`, `manager`, `room`, `backup` and so on. It is the value to  branch on, since `title` is prose in the portal language. | [optional] 
**Title** | Pointer to **NullableString** | The feature described in the portal language, with its limit already substituted into the sentence, so it  can be printed as it is. It is empty when this build ships no wording for the feature. | [optional] 
**Image** | Pointer to **NullableString** | The feature's icon as SVG markup to render inline - not a URL to fetch. It is filled in only when the  quota comes from the catalogue, and left empty on the quota the portal is actually on, on a feature that  this quota switches off, and on a feature that ships no icon. | [optional] 
**Value** | Pointer to **interface{}** |  | [optional] 
**Type** | Pointer to **NullableString** | How to read `value` and `used`: `size` for bytes, `count` for a number of things, `flag` for a feature  that is merely on or off. | [optional] 
**Used** | Pointer to [**FeatureUsedDto**](FeatureUsedDto.md) | How much of the limit is already used. It is present only on the quota the portal is actually on, and  only for a feature whose consumption is counted; a guest is shown none of these figures and a plain member  only the one for total size, so an absent value can mean the caller may not see it rather than that  nothing is used. | [optional] 
**PriceTitle** | Pointer to **NullableString** | What the feature is charged as, in the portal language - for instance the per-unit price of an add-on. It  is filled in only for a feature that costs money on top of the plan. | [optional] 

## Methods

### NewTenantQuotaFeatureDto

`func NewTenantQuotaFeatureDto() *TenantQuotaFeatureDto`

NewTenantQuotaFeatureDto instantiates a new TenantQuotaFeatureDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantQuotaFeatureDtoWithDefaults

`func NewTenantQuotaFeatureDtoWithDefaults() *TenantQuotaFeatureDto`

NewTenantQuotaFeatureDtoWithDefaults instantiates a new TenantQuotaFeatureDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TenantQuotaFeatureDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TenantQuotaFeatureDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TenantQuotaFeatureDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TenantQuotaFeatureDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TenantQuotaFeatureDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TenantQuotaFeatureDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetTitle

`func (o *TenantQuotaFeatureDto) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TenantQuotaFeatureDto) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TenantQuotaFeatureDto) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TenantQuotaFeatureDto) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *TenantQuotaFeatureDto) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *TenantQuotaFeatureDto) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetImage

`func (o *TenantQuotaFeatureDto) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *TenantQuotaFeatureDto) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *TenantQuotaFeatureDto) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *TenantQuotaFeatureDto) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *TenantQuotaFeatureDto) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *TenantQuotaFeatureDto) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetValue

`func (o *TenantQuotaFeatureDto) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TenantQuotaFeatureDto) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TenantQuotaFeatureDto) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *TenantQuotaFeatureDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *TenantQuotaFeatureDto) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *TenantQuotaFeatureDto) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetType

`func (o *TenantQuotaFeatureDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TenantQuotaFeatureDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TenantQuotaFeatureDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TenantQuotaFeatureDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *TenantQuotaFeatureDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *TenantQuotaFeatureDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUsed

`func (o *TenantQuotaFeatureDto) GetUsed() FeatureUsedDto`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *TenantQuotaFeatureDto) GetUsedOk() (*FeatureUsedDto, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *TenantQuotaFeatureDto) SetUsed(v FeatureUsedDto)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *TenantQuotaFeatureDto) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### GetPriceTitle

`func (o *TenantQuotaFeatureDto) GetPriceTitle() string`

GetPriceTitle returns the PriceTitle field if non-nil, zero value otherwise.

### GetPriceTitleOk

`func (o *TenantQuotaFeatureDto) GetPriceTitleOk() (*string, bool)`

GetPriceTitleOk returns a tuple with the PriceTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceTitle

`func (o *TenantQuotaFeatureDto) SetPriceTitle(v string)`

SetPriceTitle sets PriceTitle field to given value.

### HasPriceTitle

`func (o *TenantQuotaFeatureDto) HasPriceTitle() bool`

HasPriceTitle returns a boolean if a field has been set.

### SetPriceTitleNil

`func (o *TenantQuotaFeatureDto) SetPriceTitleNil(b bool)`

 SetPriceTitleNil sets the value for PriceTitle to be an explicit nil

### UnsetPriceTitle
`func (o *TenantQuotaFeatureDto) UnsetPriceTitle()`

UnsetPriceTitle ensures that no value is present for PriceTitle, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


