# ErrorApiResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** | The response status flag. Always 1 on an error, as opposed to 0 on success. | [optional] 
**StatusCode** | Pointer to **int32** | The HTTP status code of the response, repeated in the body. | [optional] 
**Error** | Pointer to [**ErrorApiResponseError**](ErrorApiResponseError.md) |  | [optional] 

## Methods

### NewErrorApiResponse

`func NewErrorApiResponse() *ErrorApiResponse`

NewErrorApiResponse instantiates a new ErrorApiResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorApiResponseWithDefaults

`func NewErrorApiResponseWithDefaults() *ErrorApiResponse`

NewErrorApiResponseWithDefaults instantiates a new ErrorApiResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ErrorApiResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ErrorApiResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ErrorApiResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ErrorApiResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCode

`func (o *ErrorApiResponse) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ErrorApiResponse) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ErrorApiResponse) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *ErrorApiResponse) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetError

`func (o *ErrorApiResponse) GetError() ErrorApiResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ErrorApiResponse) GetErrorOk() (*ErrorApiResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ErrorApiResponse) SetError(v ErrorApiResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *ErrorApiResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


