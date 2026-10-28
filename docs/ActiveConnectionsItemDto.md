# ActiveConnectionsItemDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The ID of the sign-in this connection was opened by. Pass it as `loginEventId` to  `PUT api/2.0/security/activeconnections/logout/{loginEventId}` to end this one connection; the item whose  value equals `loginEvent` is the connection the current request uses. | 
**TenantId** | **int32** | The portal the sign-in was made on. The operation never crosses portals, so it is the current one on every  item. | 
**UserId** | **string** | The user the connection belongs to, which is the calling user on every item - the operation cannot report  anyone else's connections. | 
**Mobile** | Pointer to **bool** | Whether the sign-in came from a mobile client. No mobile marker is stored with a connection, so the value  is `false` on every item and tells a caller nothing about the device. | [optional] 
**Ip** | Pointer to **NullableString** | The IP address the sign-in came from, with the port stripped off. On the item that matches `loginEvent` it  is taken from the address the current request arrives from instead of the one stored at sign-in. | [optional] 
**Country** | Pointer to **NullableString** | The English name of the country the IP address is located in. It is empty when the address cannot be  located, which is the normal outcome for private and loopback addresses. | [optional] 
**City** | Pointer to **NullableString** | The city the IP address is located in, empty under the same conditions as `country`. | [optional] 
**Browser** | Pointer to **NullableString** | The browser and its version as parsed from the user agent of the sign-in, empty when the client sent no  recognisable one. It is refreshed from the current request on the item that matches `loginEvent`. | [optional] 
**Platform** | Pointer to **NullableString** | The operating system as parsed from the user agent of the sign-in, refreshed and left empty under the same  conditions as `browser`. | [optional] 
**Date** | Pointer to [**ApiDateTime**](ApiDateTime.md) | When the sign-in happened, in the portal time zone rather than in UTC. | [optional] 
**Page** | Pointer to **NullableString** | Where in the portal the sign-in was made from: the referrer of the request that created it, or that  request's own path when it carried no referrer. Long values are cut off at 512 characters. | [optional] 

## Methods

### NewActiveConnectionsItemDto

`func NewActiveConnectionsItemDto(id int32, tenantId int32, userId string, ) *ActiveConnectionsItemDto`

NewActiveConnectionsItemDto instantiates a new ActiveConnectionsItemDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveConnectionsItemDtoWithDefaults

`func NewActiveConnectionsItemDtoWithDefaults() *ActiveConnectionsItemDto`

NewActiveConnectionsItemDtoWithDefaults instantiates a new ActiveConnectionsItemDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActiveConnectionsItemDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActiveConnectionsItemDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActiveConnectionsItemDto) SetId(v int32)`

SetId sets Id field to given value.


### GetTenantId

`func (o *ActiveConnectionsItemDto) GetTenantId() int32`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ActiveConnectionsItemDto) GetTenantIdOk() (*int32, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ActiveConnectionsItemDto) SetTenantId(v int32)`

SetTenantId sets TenantId field to given value.


### GetUserId

`func (o *ActiveConnectionsItemDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ActiveConnectionsItemDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ActiveConnectionsItemDto) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetMobile

`func (o *ActiveConnectionsItemDto) GetMobile() bool`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *ActiveConnectionsItemDto) GetMobileOk() (*bool, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *ActiveConnectionsItemDto) SetMobile(v bool)`

SetMobile sets Mobile field to given value.

### HasMobile

`func (o *ActiveConnectionsItemDto) HasMobile() bool`

HasMobile returns a boolean if a field has been set.

### GetIp

`func (o *ActiveConnectionsItemDto) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ActiveConnectionsItemDto) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ActiveConnectionsItemDto) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ActiveConnectionsItemDto) HasIp() bool`

HasIp returns a boolean if a field has been set.

### SetIpNil

`func (o *ActiveConnectionsItemDto) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *ActiveConnectionsItemDto) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetCountry

`func (o *ActiveConnectionsItemDto) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ActiveConnectionsItemDto) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ActiveConnectionsItemDto) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ActiveConnectionsItemDto) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *ActiveConnectionsItemDto) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *ActiveConnectionsItemDto) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetCity

`func (o *ActiveConnectionsItemDto) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *ActiveConnectionsItemDto) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *ActiveConnectionsItemDto) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *ActiveConnectionsItemDto) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *ActiveConnectionsItemDto) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *ActiveConnectionsItemDto) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetBrowser

`func (o *ActiveConnectionsItemDto) GetBrowser() string`

GetBrowser returns the Browser field if non-nil, zero value otherwise.

### GetBrowserOk

`func (o *ActiveConnectionsItemDto) GetBrowserOk() (*string, bool)`

GetBrowserOk returns a tuple with the Browser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowser

`func (o *ActiveConnectionsItemDto) SetBrowser(v string)`

SetBrowser sets Browser field to given value.

### HasBrowser

`func (o *ActiveConnectionsItemDto) HasBrowser() bool`

HasBrowser returns a boolean if a field has been set.

### SetBrowserNil

`func (o *ActiveConnectionsItemDto) SetBrowserNil(b bool)`

 SetBrowserNil sets the value for Browser to be an explicit nil

### UnsetBrowser
`func (o *ActiveConnectionsItemDto) UnsetBrowser()`

UnsetBrowser ensures that no value is present for Browser, not even an explicit nil
### GetPlatform

`func (o *ActiveConnectionsItemDto) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ActiveConnectionsItemDto) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ActiveConnectionsItemDto) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ActiveConnectionsItemDto) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *ActiveConnectionsItemDto) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *ActiveConnectionsItemDto) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetDate

`func (o *ActiveConnectionsItemDto) GetDate() ApiDateTime`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ActiveConnectionsItemDto) GetDateOk() (*ApiDateTime, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ActiveConnectionsItemDto) SetDate(v ApiDateTime)`

SetDate sets Date field to given value.

### HasDate

`func (o *ActiveConnectionsItemDto) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetPage

`func (o *ActiveConnectionsItemDto) GetPage() string`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ActiveConnectionsItemDto) GetPageOk() (*string, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ActiveConnectionsItemDto) SetPage(v string)`

SetPage sets Page field to given value.

### HasPage

`func (o *ActiveConnectionsItemDto) HasPage() bool`

HasPage returns a boolean if a field has been set.

### SetPageNil

`func (o *ActiveConnectionsItemDto) SetPageNil(b bool)`

 SetPageNil sets the value for Page to be an explicit nil

### UnsetPage
`func (o *ActiveConnectionsItemDto) UnsetPage()`

UnsetPage ensures that no value is present for Page, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


