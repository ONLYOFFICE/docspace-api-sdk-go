# MultiSizeLogoCoverDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** | The logo cover ID. | 
**Data** | **map[string]string** | The logo cover data. | 

## Methods

### NewMultiSizeLogoCoverDto

`func NewMultiSizeLogoCoverDto(id NullableString, data map[string]*string, ) *MultiSizeLogoCoverDto`

NewMultiSizeLogoCoverDto instantiates a new MultiSizeLogoCoverDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiSizeLogoCoverDtoWithDefaults

`func NewMultiSizeLogoCoverDtoWithDefaults() *MultiSizeLogoCoverDto`

NewMultiSizeLogoCoverDtoWithDefaults instantiates a new MultiSizeLogoCoverDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MultiSizeLogoCoverDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MultiSizeLogoCoverDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MultiSizeLogoCoverDto) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *MultiSizeLogoCoverDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MultiSizeLogoCoverDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetData

`func (o *MultiSizeLogoCoverDto) GetData() map[string]*string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MultiSizeLogoCoverDto) GetDataOk() (*map[string]*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MultiSizeLogoCoverDto) SetData(v map[string]*string)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


