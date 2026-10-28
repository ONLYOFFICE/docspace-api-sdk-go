# FieldError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | Pointer to **string** | The name of the field that failed validation | [optional] 
**Code** | Pointer to **string** | Error code for localization purposes | [optional] 
**Message** | Pointer to **string** | Human readable error message | [optional] 

## Methods

### NewFieldError

`func NewFieldError() *FieldError`

NewFieldError instantiates a new FieldError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldErrorWithDefaults

`func NewFieldErrorWithDefaults() *FieldError`

NewFieldErrorWithDefaults instantiates a new FieldError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *FieldError) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *FieldError) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *FieldError) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *FieldError) HasField() bool`

HasField returns a boolean if a field has been set.

### GetCode

`func (o *FieldError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *FieldError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *FieldError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *FieldError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *FieldError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FieldError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FieldError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *FieldError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


