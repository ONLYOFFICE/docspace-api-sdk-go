# LoginEventDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The login event ID. | [optional] 
**Date** | Pointer to **NullableTime** | The login event date. | [optional] 
**User** | Pointer to **NullableString** | The user name of the login event. | [optional] 
**UserId** | Pointer to **string** | The user ID of the login event. | [optional] 
**Login** | Pointer to **NullableString** | The user login of the login event. | [optional] 
**Action** | Pointer to **NullableString** | The login event action. | [optional] 
**ActionId** | Pointer to [**MessageAction**](MessageAction.md) | The login-related action to filter events by. | [optional] 
**Ip** | Pointer to **NullableString** | The login event IP. | [optional] 
**Country** | Pointer to **NullableString** | The login event country. | [optional] 
**City** | Pointer to **NullableString** | The login event city. | [optional] 
**Browser** | Pointer to **NullableString** | The login event browser. | [optional] 
**Platform** | Pointer to **NullableString** | The login event platform. | [optional] 
**Page** | Pointer to **NullableString** | The login event page. | [optional] 

## Methods

### NewLoginEventDto

`func NewLoginEventDto() *LoginEventDto`

NewLoginEventDto instantiates a new LoginEventDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginEventDtoWithDefaults

`func NewLoginEventDtoWithDefaults() *LoginEventDto`

NewLoginEventDtoWithDefaults instantiates a new LoginEventDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LoginEventDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoginEventDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoginEventDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *LoginEventDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDate

`func (o *LoginEventDto) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *LoginEventDto) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *LoginEventDto) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *LoginEventDto) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *LoginEventDto) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *LoginEventDto) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetUser

`func (o *LoginEventDto) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *LoginEventDto) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *LoginEventDto) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *LoginEventDto) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *LoginEventDto) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *LoginEventDto) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetUserId

`func (o *LoginEventDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *LoginEventDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *LoginEventDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *LoginEventDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetLogin

`func (o *LoginEventDto) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *LoginEventDto) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *LoginEventDto) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *LoginEventDto) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### SetLoginNil

`func (o *LoginEventDto) SetLoginNil(b bool)`

 SetLoginNil sets the value for Login to be an explicit nil

### UnsetLogin
`func (o *LoginEventDto) UnsetLogin()`

UnsetLogin ensures that no value is present for Login, not even an explicit nil
### GetAction

`func (o *LoginEventDto) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *LoginEventDto) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *LoginEventDto) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *LoginEventDto) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *LoginEventDto) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *LoginEventDto) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetActionId

`func (o *LoginEventDto) GetActionId() MessageAction`

GetActionId returns the ActionId field if non-nil, zero value otherwise.

### GetActionIdOk

`func (o *LoginEventDto) GetActionIdOk() (*MessageAction, bool)`

GetActionIdOk returns a tuple with the ActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionId

`func (o *LoginEventDto) SetActionId(v MessageAction)`

SetActionId sets ActionId field to given value.

### HasActionId

`func (o *LoginEventDto) HasActionId() bool`

HasActionId returns a boolean if a field has been set.

### GetIp

`func (o *LoginEventDto) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *LoginEventDto) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *LoginEventDto) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *LoginEventDto) HasIp() bool`

HasIp returns a boolean if a field has been set.

### SetIpNil

`func (o *LoginEventDto) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *LoginEventDto) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetCountry

`func (o *LoginEventDto) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *LoginEventDto) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *LoginEventDto) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *LoginEventDto) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *LoginEventDto) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *LoginEventDto) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetCity

`func (o *LoginEventDto) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *LoginEventDto) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *LoginEventDto) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *LoginEventDto) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *LoginEventDto) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *LoginEventDto) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetBrowser

`func (o *LoginEventDto) GetBrowser() string`

GetBrowser returns the Browser field if non-nil, zero value otherwise.

### GetBrowserOk

`func (o *LoginEventDto) GetBrowserOk() (*string, bool)`

GetBrowserOk returns a tuple with the Browser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowser

`func (o *LoginEventDto) SetBrowser(v string)`

SetBrowser sets Browser field to given value.

### HasBrowser

`func (o *LoginEventDto) HasBrowser() bool`

HasBrowser returns a boolean if a field has been set.

### SetBrowserNil

`func (o *LoginEventDto) SetBrowserNil(b bool)`

 SetBrowserNil sets the value for Browser to be an explicit nil

### UnsetBrowser
`func (o *LoginEventDto) UnsetBrowser()`

UnsetBrowser ensures that no value is present for Browser, not even an explicit nil
### GetPlatform

`func (o *LoginEventDto) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *LoginEventDto) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *LoginEventDto) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *LoginEventDto) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *LoginEventDto) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *LoginEventDto) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetPage

`func (o *LoginEventDto) GetPage() string`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *LoginEventDto) GetPageOk() (*string, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *LoginEventDto) SetPage(v string)`

SetPage sets Page field to given value.

### HasPage

`func (o *LoginEventDto) HasPage() bool`

HasPage returns a boolean if a field has been set.

### SetPageNil

`func (o *LoginEventDto) SetPageNil(b bool)`

 SetPageNil sets the value for Page to be an explicit nil

### UnsetPage
`func (o *LoginEventDto) UnsetPage()`

UnsetPage ensures that no value is present for Page, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


