# SetCustomFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | [**[]CustomFieldRequest**](CustomFieldRequest.md) | The custom fields to set on the entry. A listed field gets the value, a null or empty value removes the field  from the entry, the fields not listed are left alone. A name the portal has not seen yet creates the field. | 

## Methods

### NewSetCustomFields

`func NewSetCustomFields(fields []CustomFieldRequest, ) *SetCustomFields`

NewSetCustomFields instantiates a new SetCustomFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetCustomFieldsWithDefaults

`func NewSetCustomFieldsWithDefaults() *SetCustomFields`

NewSetCustomFieldsWithDefaults instantiates a new SetCustomFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *SetCustomFields) GetFields() []CustomFieldRequest`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *SetCustomFields) GetFieldsOk() (*[]CustomFieldRequest, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *SetCustomFields) SetFields(v []CustomFieldRequest)`

SetFields sets Fields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


