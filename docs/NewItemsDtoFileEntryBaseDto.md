# NewItemsDtoFileEntryBaseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | [**ApiDateTime**](ApiDateTime.md) |  | 
**Items** | [**[]FileEntryBaseDto**](FileEntryBaseDto.md) | The list of items. | 

## Methods

### NewNewItemsDtoFileEntryBaseDto

`func NewNewItemsDtoFileEntryBaseDto(date ApiDateTime, items []FileEntryBaseDto, ) *NewItemsDtoFileEntryBaseDto`

NewNewItemsDtoFileEntryBaseDto instantiates a new NewItemsDtoFileEntryBaseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewItemsDtoFileEntryBaseDtoWithDefaults

`func NewNewItemsDtoFileEntryBaseDtoWithDefaults() *NewItemsDtoFileEntryBaseDto`

NewNewItemsDtoFileEntryBaseDtoWithDefaults instantiates a new NewItemsDtoFileEntryBaseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *NewItemsDtoFileEntryBaseDto) GetDate() ApiDateTime`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *NewItemsDtoFileEntryBaseDto) GetDateOk() (*ApiDateTime, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *NewItemsDtoFileEntryBaseDto) SetDate(v ApiDateTime)`

SetDate sets Date field to given value.


### GetItems

`func (o *NewItemsDtoFileEntryBaseDto) GetItems() []FileEntryBaseDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *NewItemsDtoFileEntryBaseDto) GetItemsOk() (*[]FileEntryBaseDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *NewItemsDtoFileEntryBaseDto) SetItems(v []FileEntryBaseDto)`

SetItems sets Items field to given value.


### SetItemsNil

`func (o *NewItemsDtoFileEntryBaseDto) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *NewItemsDtoFileEntryBaseDto) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


