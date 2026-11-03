# AiChatErrorPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Stable machine-readable code. One of a canonical set — `auth`, `model_not_found`, `rate_limit`, `timeout`, `network`, `bad_request`, `server`, `not_found`, `insufficient_funds`, `aborted`, `unknown` — or a provider-specific string when the SDK supplies one verbatim. | 
**Message** | **string** | Human-readable, end-user-safe message. | 
**Cause** | Pointer to [**AiChatErrorPayload**](AiChatErrorPayload.md) | Optional nested cause, normalized through the same helper. | [optional] 

## Methods

### NewAiChatErrorPayload

`func NewAiChatErrorPayload(code string, message string, ) *AiChatErrorPayload`

NewAiChatErrorPayload instantiates a new AiChatErrorPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiChatErrorPayloadWithDefaults

`func NewAiChatErrorPayloadWithDefaults() *AiChatErrorPayload`

NewAiChatErrorPayloadWithDefaults instantiates a new AiChatErrorPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *AiChatErrorPayload) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AiChatErrorPayload) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AiChatErrorPayload) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *AiChatErrorPayload) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AiChatErrorPayload) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AiChatErrorPayload) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetCause

`func (o *AiChatErrorPayload) GetCause() AiChatErrorPayload`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *AiChatErrorPayload) GetCauseOk() (*AiChatErrorPayload, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *AiChatErrorPayload) SetCause(v AiChatErrorPayload)`

SetCause sets Cause field to given value.

### HasCause

`func (o *AiChatErrorPayload) HasCause() bool`

HasCause returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


