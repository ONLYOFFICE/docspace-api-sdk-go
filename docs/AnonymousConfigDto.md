# AnonymousConfigDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Request** | **bool** | Whether the editors ask an anonymous participant for a display name before letting them in. It follows the  chat permission of the document, since a nameless participant cannot take part in one. | 

## Methods

### NewAnonymousConfigDto

`func NewAnonymousConfigDto(request bool, ) *AnonymousConfigDto`

NewAnonymousConfigDto instantiates a new AnonymousConfigDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnonymousConfigDtoWithDefaults

`func NewAnonymousConfigDtoWithDefaults() *AnonymousConfigDto`

NewAnonymousConfigDtoWithDefaults instantiates a new AnonymousConfigDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequest

`func (o *AnonymousConfigDto) GetRequest() bool`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *AnonymousConfigDto) GetRequestOk() (*bool, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *AnonymousConfigDto) SetRequest(v bool)`

SetRequest sets Request field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


