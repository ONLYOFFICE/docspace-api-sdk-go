# AiThread

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThreadId** | **string** | Unique thread identifier (UUID). | 
**Title** | Pointer to **string** | Optional thread title. Auto-generated from the first message if not set. | [optional] 
**LastEditDate** | Pointer to **float32** | Timestamp (ms since epoch) of the last message in this thread. Used for sorting. | [optional] 
**Provider** | Pointer to [**AiTProvider**](AiTProvider.md) | Provider configuration at the time of last message. Used for thread-level provider display. | [optional] 
**Model** | Pointer to [**AiModel**](AiModel.md) | Model info at the time of last message. | [optional] 
**ProfileId** | Pointer to **string** | ID of the profile used for this thread. Links to `Profile.id`. | [optional] 

## Methods

### NewAiThread

`func NewAiThread(threadId string, ) *AiThread`

NewAiThread instantiates a new AiThread object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiThreadWithDefaults

`func NewAiThreadWithDefaults() *AiThread`

NewAiThreadWithDefaults instantiates a new AiThread object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreadId

`func (o *AiThread) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *AiThread) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *AiThread) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.


### GetTitle

`func (o *AiThread) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AiThread) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AiThread) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AiThread) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetLastEditDate

`func (o *AiThread) GetLastEditDate() float32`

GetLastEditDate returns the LastEditDate field if non-nil, zero value otherwise.

### GetLastEditDateOk

`func (o *AiThread) GetLastEditDateOk() (*float32, bool)`

GetLastEditDateOk returns a tuple with the LastEditDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEditDate

`func (o *AiThread) SetLastEditDate(v float32)`

SetLastEditDate sets LastEditDate field to given value.

### HasLastEditDate

`func (o *AiThread) HasLastEditDate() bool`

HasLastEditDate returns a boolean if a field has been set.

### GetProvider

`func (o *AiThread) GetProvider() AiTProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AiThread) GetProviderOk() (*AiTProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AiThread) SetProvider(v AiTProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *AiThread) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetModel

`func (o *AiThread) GetModel() AiModel`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AiThread) GetModelOk() (*AiModel, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AiThread) SetModel(v AiModel)`

SetModel sets Model field to given value.

### HasModel

`func (o *AiThread) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetProfileId

`func (o *AiThread) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *AiThread) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *AiThread) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *AiThread) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


