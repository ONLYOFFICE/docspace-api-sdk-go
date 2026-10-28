# AiApiDateTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UtcTime** | Pointer to **time.Time** | The time in UTC format. | [optional] [readonly] 
**TimeZoneOffset** | Pointer to **string** | The time zone offset. | [optional] [readonly] 

## Methods

### NewAiApiDateTime

`func NewAiApiDateTime() *AiApiDateTime`

NewAiApiDateTime instantiates a new AiApiDateTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiApiDateTimeWithDefaults

`func NewAiApiDateTimeWithDefaults() *AiApiDateTime`

NewAiApiDateTimeWithDefaults instantiates a new AiApiDateTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUtcTime

`func (o *AiApiDateTime) GetUtcTime() time.Time`

GetUtcTime returns the UtcTime field if non-nil, zero value otherwise.

### GetUtcTimeOk

`func (o *AiApiDateTime) GetUtcTimeOk() (*time.Time, bool)`

GetUtcTimeOk returns a tuple with the UtcTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtcTime

`func (o *AiApiDateTime) SetUtcTime(v time.Time)`

SetUtcTime sets UtcTime field to given value.

### HasUtcTime

`func (o *AiApiDateTime) HasUtcTime() bool`

HasUtcTime returns a boolean if a field has been set.

### GetTimeZoneOffset

`func (o *AiApiDateTime) GetTimeZoneOffset() string`

GetTimeZoneOffset returns the TimeZoneOffset field if non-nil, zero value otherwise.

### GetTimeZoneOffsetOk

`func (o *AiApiDateTime) GetTimeZoneOffsetOk() (*string, bool)`

GetTimeZoneOffsetOk returns a tuple with the TimeZoneOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZoneOffset

`func (o *AiApiDateTime) SetTimeZoneOffset(v string)`

SetTimeZoneOffset sets TimeZoneOffset field to given value.

### HasTimeZoneOffset

`func (o *AiApiDateTime) HasTimeZoneOffset() bool`

HasTimeZoneOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


