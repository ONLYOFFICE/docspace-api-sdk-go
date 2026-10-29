# AiEntryMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Templates** | [**[]AiEntryMetadataTemplatesInner**](AiEntryMetadataTemplatesInner.md) |  | 
**CustomFields** | [**[]AiEntryMetadataCustomFieldsInner**](AiEntryMetadataCustomFieldsInner.md) |  | 

## Methods

### NewAiEntryMetadata

`func NewAiEntryMetadata(templates []AiEntryMetadataTemplatesInner, customFields []AiEntryMetadataCustomFieldsInner, ) *AiEntryMetadata`

NewAiEntryMetadata instantiates a new AiEntryMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiEntryMetadataWithDefaults

`func NewAiEntryMetadataWithDefaults() *AiEntryMetadata`

NewAiEntryMetadataWithDefaults instantiates a new AiEntryMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplates

`func (o *AiEntryMetadata) GetTemplates() []AiEntryMetadataTemplatesInner`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *AiEntryMetadata) GetTemplatesOk() (*[]AiEntryMetadataTemplatesInner, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *AiEntryMetadata) SetTemplates(v []AiEntryMetadataTemplatesInner)`

SetTemplates sets Templates field to given value.


### GetCustomFields

`func (o *AiEntryMetadata) GetCustomFields() []AiEntryMetadataCustomFieldsInner`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *AiEntryMetadata) GetCustomFieldsOk() (*[]AiEntryMetadataCustomFieldsInner, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *AiEntryMetadata) SetCustomFields(v []AiEntryMetadataCustomFieldsInner)`

SetCustomFields sets CustomFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


