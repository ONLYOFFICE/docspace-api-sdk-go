# DefaultTemplateSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]DefaultTemplateItemDto**](DefaultTemplateItemDto.md) | One entry per extension the portal's built-in template set covers, whether or not a custom blank has been  chosen for it, so the list is never empty and its length follows the template set rather than the number of  custom blanks. Entries come in the order an interface shows them: text document, spreadsheet, presentation and  PDF first, everything else by extension. | 

## Methods

### NewDefaultTemplateSettingsDto

`func NewDefaultTemplateSettingsDto(items []DefaultTemplateItemDto, ) *DefaultTemplateSettingsDto`

NewDefaultTemplateSettingsDto instantiates a new DefaultTemplateSettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefaultTemplateSettingsDtoWithDefaults

`func NewDefaultTemplateSettingsDtoWithDefaults() *DefaultTemplateSettingsDto`

NewDefaultTemplateSettingsDtoWithDefaults instantiates a new DefaultTemplateSettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *DefaultTemplateSettingsDto) GetItems() []DefaultTemplateItemDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *DefaultTemplateSettingsDto) GetItemsOk() (*[]DefaultTemplateItemDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *DefaultTemplateSettingsDto) SetItems(v []DefaultTemplateItemDto)`

SetItems sets Items field to given value.


### SetItemsNil

`func (o *DefaultTemplateSettingsDto) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *DefaultTemplateSettingsDto) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


