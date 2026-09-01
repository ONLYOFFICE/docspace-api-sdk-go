# AuditEventDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The audit event ID. | [optional] 
**Date** | Pointer to **NullableTime** | The audit event date. | [optional] 
**User** | Pointer to **NullableString** | The name of the user who triggered the audit event. | [optional] 
**UserId** | Pointer to **string** | The ID of the user who triggered the audit event. | [optional] 
**Action** | Pointer to **NullableString** | The audit event action. | [optional] 
**ActionId** | Pointer to [**MessageAction**](MessageAction.md) | The specific action that occurred within the audit event. | [optional] 
**Ip** | Pointer to **NullableString** | The audit event IP. | [optional] 
**Country** | Pointer to **NullableString** | The audit event country. | [optional] 
**City** | Pointer to **NullableString** | The audit event city. | [optional] 
**Browser** | Pointer to **NullableString** | The audit event browser. | [optional] 
**Platform** | Pointer to **NullableString** | The audit event platform. | [optional] 
**Page** | Pointer to **NullableString** | The audit event page. | [optional] 
**ActionType** | Pointer to [**ActionType**](ActionType.md) | The type of action performed in the audit event (e.g., Create, Update, Delete). | [optional] 
**Product** | Pointer to [**ProductType**](ProductType.md) | The type of product related to the audit event. | [optional] 
**Location** | Pointer to [**LocationType**](LocationType.md) | The location where the audit event occurred. | [optional] 
**Target** | Pointer to **[]string** | The list of target objects affected by the audit event (e.g., document ID, user account). | [optional] 
**Entries** | Pointer to [**[]EntryType**](EntryType.md) | The list of audit entry types (e.g., Folder, User, File). | [optional] 
**Context** | Pointer to **NullableString** | The audit event context. | [optional] 

## Methods

### NewAuditEventDto

`func NewAuditEventDto() *AuditEventDto`

NewAuditEventDto instantiates a new AuditEventDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditEventDtoWithDefaults

`func NewAuditEventDtoWithDefaults() *AuditEventDto`

NewAuditEventDtoWithDefaults instantiates a new AuditEventDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuditEventDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditEventDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditEventDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AuditEventDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDate

`func (o *AuditEventDto) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *AuditEventDto) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *AuditEventDto) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *AuditEventDto) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *AuditEventDto) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *AuditEventDto) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetUser

`func (o *AuditEventDto) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AuditEventDto) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AuditEventDto) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *AuditEventDto) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *AuditEventDto) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *AuditEventDto) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetUserId

`func (o *AuditEventDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AuditEventDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AuditEventDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AuditEventDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetAction

`func (o *AuditEventDto) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AuditEventDto) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AuditEventDto) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *AuditEventDto) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *AuditEventDto) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *AuditEventDto) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetActionId

`func (o *AuditEventDto) GetActionId() MessageAction`

GetActionId returns the ActionId field if non-nil, zero value otherwise.

### GetActionIdOk

`func (o *AuditEventDto) GetActionIdOk() (*MessageAction, bool)`

GetActionIdOk returns a tuple with the ActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionId

`func (o *AuditEventDto) SetActionId(v MessageAction)`

SetActionId sets ActionId field to given value.

### HasActionId

`func (o *AuditEventDto) HasActionId() bool`

HasActionId returns a boolean if a field has been set.

### GetIp

`func (o *AuditEventDto) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *AuditEventDto) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *AuditEventDto) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *AuditEventDto) HasIp() bool`

HasIp returns a boolean if a field has been set.

### SetIpNil

`func (o *AuditEventDto) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *AuditEventDto) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetCountry

`func (o *AuditEventDto) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *AuditEventDto) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *AuditEventDto) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *AuditEventDto) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *AuditEventDto) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *AuditEventDto) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetCity

`func (o *AuditEventDto) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *AuditEventDto) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *AuditEventDto) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *AuditEventDto) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *AuditEventDto) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *AuditEventDto) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetBrowser

`func (o *AuditEventDto) GetBrowser() string`

GetBrowser returns the Browser field if non-nil, zero value otherwise.

### GetBrowserOk

`func (o *AuditEventDto) GetBrowserOk() (*string, bool)`

GetBrowserOk returns a tuple with the Browser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowser

`func (o *AuditEventDto) SetBrowser(v string)`

SetBrowser sets Browser field to given value.

### HasBrowser

`func (o *AuditEventDto) HasBrowser() bool`

HasBrowser returns a boolean if a field has been set.

### SetBrowserNil

`func (o *AuditEventDto) SetBrowserNil(b bool)`

 SetBrowserNil sets the value for Browser to be an explicit nil

### UnsetBrowser
`func (o *AuditEventDto) UnsetBrowser()`

UnsetBrowser ensures that no value is present for Browser, not even an explicit nil
### GetPlatform

`func (o *AuditEventDto) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *AuditEventDto) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *AuditEventDto) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *AuditEventDto) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *AuditEventDto) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *AuditEventDto) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetPage

`func (o *AuditEventDto) GetPage() string`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *AuditEventDto) GetPageOk() (*string, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *AuditEventDto) SetPage(v string)`

SetPage sets Page field to given value.

### HasPage

`func (o *AuditEventDto) HasPage() bool`

HasPage returns a boolean if a field has been set.

### SetPageNil

`func (o *AuditEventDto) SetPageNil(b bool)`

 SetPageNil sets the value for Page to be an explicit nil

### UnsetPage
`func (o *AuditEventDto) UnsetPage()`

UnsetPage ensures that no value is present for Page, not even an explicit nil
### GetActionType

`func (o *AuditEventDto) GetActionType() ActionType`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *AuditEventDto) GetActionTypeOk() (*ActionType, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *AuditEventDto) SetActionType(v ActionType)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *AuditEventDto) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### GetProduct

`func (o *AuditEventDto) GetProduct() ProductType`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *AuditEventDto) GetProductOk() (*ProductType, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *AuditEventDto) SetProduct(v ProductType)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *AuditEventDto) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetLocation

`func (o *AuditEventDto) GetLocation() LocationType`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AuditEventDto) GetLocationOk() (*LocationType, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AuditEventDto) SetLocation(v LocationType)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *AuditEventDto) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetTarget

`func (o *AuditEventDto) GetTarget() []string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *AuditEventDto) GetTargetOk() (*[]string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *AuditEventDto) SetTarget(v []string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *AuditEventDto) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTargetNil

`func (o *AuditEventDto) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *AuditEventDto) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil
### GetEntries

`func (o *AuditEventDto) GetEntries() []EntryType`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *AuditEventDto) GetEntriesOk() (*[]EntryType, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *AuditEventDto) SetEntries(v []EntryType)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *AuditEventDto) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### SetEntriesNil

`func (o *AuditEventDto) SetEntriesNil(b bool)`

 SetEntriesNil sets the value for Entries to be an explicit nil

### UnsetEntries
`func (o *AuditEventDto) UnsetEntries()`

UnsetEntries ensures that no value is present for Entries, not even an explicit nil
### GetContext

`func (o *AuditEventDto) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *AuditEventDto) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *AuditEventDto) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *AuditEventDto) HasContext() bool`

HasContext returns a boolean if a field has been set.

### SetContextNil

`func (o *AuditEventDto) SetContextNil(b bool)`

 SetContextNil sets the value for Context to be an explicit nil

### UnsetContext
`func (o *AuditEventDto) UnsetContext()`

UnsetContext ensures that no value is present for Context, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


