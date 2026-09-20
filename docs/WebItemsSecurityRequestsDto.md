# WebItemsSecurityRequestsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]ItemKeyValuePairStringBoolean**](ItemKeyValuePairStringBoolean.md) | The modules to switch, each entry pairing a module GUID as its `key` with the new enabled flag as its  `value`. A key that is not a GUID fails the whole request as invalid, and a module listed twice is applied  once, from its first entry. No allow-list travels here: switching a product module on restores the users and  groups it was last restricted to, and everything else is stored as a plain allow or deny for everyone. | [optional] 

## Methods

### NewWebItemsSecurityRequestsDto

`func NewWebItemsSecurityRequestsDto() *WebItemsSecurityRequestsDto`

NewWebItemsSecurityRequestsDto instantiates a new WebItemsSecurityRequestsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebItemsSecurityRequestsDtoWithDefaults

`func NewWebItemsSecurityRequestsDtoWithDefaults() *WebItemsSecurityRequestsDto`

NewWebItemsSecurityRequestsDtoWithDefaults instantiates a new WebItemsSecurityRequestsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *WebItemsSecurityRequestsDto) GetItems() []ItemKeyValuePairStringBoolean`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *WebItemsSecurityRequestsDto) GetItemsOk() (*[]ItemKeyValuePairStringBoolean, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *WebItemsSecurityRequestsDto) SetItems(v []ItemKeyValuePairStringBoolean)`

SetItems sets Items field to given value.

### HasItems

`func (o *WebItemsSecurityRequestsDto) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *WebItemsSecurityRequestsDto) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *WebItemsSecurityRequestsDto) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


