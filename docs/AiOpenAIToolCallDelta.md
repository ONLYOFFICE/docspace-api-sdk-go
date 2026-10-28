# AiOpenAIToolCallDelta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **float32** | The zero-based position of the tool call within the message. | 
**Id** | Pointer to **string** | The tool call identifier, quoted back when its result is submitted. | [optional] 
**Type** | Pointer to **string** | Always `function` - the only tool kind the API defines. | [optional] 
**Function** | Pointer to [**AiOpenAIToolCallDeltaFunction**](AiOpenAIToolCallDeltaFunction.md) |  | [optional] 

## Methods

### NewAiOpenAIToolCallDelta

`func NewAiOpenAIToolCallDelta(index float32, ) *AiOpenAIToolCallDelta`

NewAiOpenAIToolCallDelta instantiates a new AiOpenAIToolCallDelta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiOpenAIToolCallDeltaWithDefaults

`func NewAiOpenAIToolCallDeltaWithDefaults() *AiOpenAIToolCallDelta`

NewAiOpenAIToolCallDeltaWithDefaults instantiates a new AiOpenAIToolCallDelta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *AiOpenAIToolCallDelta) GetIndex() float32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *AiOpenAIToolCallDelta) GetIndexOk() (*float32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *AiOpenAIToolCallDelta) SetIndex(v float32)`

SetIndex sets Index field to given value.


### GetId

`func (o *AiOpenAIToolCallDelta) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiOpenAIToolCallDelta) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiOpenAIToolCallDelta) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AiOpenAIToolCallDelta) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *AiOpenAIToolCallDelta) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AiOpenAIToolCallDelta) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AiOpenAIToolCallDelta) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AiOpenAIToolCallDelta) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFunction

`func (o *AiOpenAIToolCallDelta) GetFunction() AiOpenAIToolCallDeltaFunction`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *AiOpenAIToolCallDelta) GetFunctionOk() (*AiOpenAIToolCallDeltaFunction, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *AiOpenAIToolCallDelta) SetFunction(v AiOpenAIToolCallDeltaFunction)`

SetFunction sets Function field to given value.

### HasFunction

`func (o *AiOpenAIToolCallDelta) HasFunction() bool`

HasFunction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


