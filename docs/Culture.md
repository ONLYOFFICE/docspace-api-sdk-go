# Culture

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CultureName** | **string** | The user culture name (en-US, de, fr, es, ...). | 

## Methods

### NewCulture

`func NewCulture(cultureName string, ) *Culture`

NewCulture instantiates a new Culture object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCultureWithDefaults

`func NewCultureWithDefaults() *Culture`

NewCultureWithDefaults instantiates a new Culture object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCultureName

`func (o *Culture) GetCultureName() string`

GetCultureName returns the CultureName field if non-nil, zero value otherwise.

### GetCultureNameOk

`func (o *Culture) GetCultureNameOk() (*string, bool)`

GetCultureNameOk returns a tuple with the CultureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCultureName

`func (o *Culture) SetCultureName(v string)`

SetCultureName sets CultureName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


